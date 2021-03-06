package main

import (
	"encoding/json"
	"flag"
	"github.com/hyperturtle/pixels/lib"
	"github.com/lucasb-eyer/go-colorful"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var (
	ga            string
	refresh_token string
	client_secret string
	client_id     string
)

type jwtToken struct {
	Access_Token  string `json:"access_token"`
	Refresh_Token string `json:"refresh_token"`
}

func getToken(refresh string) string {
	resp, err := http.PostForm("https://www.googleapis.com/oauth2/v3/token", url.Values{
		"client_secret": {client_secret},
		"grant_type":    {"refresh_token"},
		"refresh_token": {refresh},
		"client_id":     {client_id},
	})
	if err != nil {
		panic(err)
	}

	var v jwtToken
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&v)
	return v.Access_Token
}

type analyticsResponse struct {
	TotalsForAllResults totalsForAllResults `json:"totalsForAllResults"`
}

type totalsForAllResults struct {
	ActiveUsers string `json:"rt:activeUsers"`
}

var (
	histogram []int
)

func getCount() int {
	r, err := http.NewRequest("GET", "https://content.googleapis.com/analytics/v3/data/realtime?ids="+ga+"&metrics=rt%3AactiveUsers", nil)
	if err != nil {
		panic(err)
	}
	// r.Header.Add("Authorization", "Bearer "+token)
	r.Header.Add("Authorization", "Bearer "+getToken(refresh_token))
	// r.Header.Add("X-Origin", domain)
	// r.Header.Add("X-Referer", domain)
	resp, err := (&http.Client{}).Do(r)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(resp.Body)
	var v analyticsResponse
	decoder.Decode(&v)
	i, _ := strconv.Atoi(v.TotalsForAllResults.ActiveUsers)
	return i
}

var breakpoints = []int{
	20,
	30,
	40,
	50,
	60,
	60 * 2,
	60 * 3,
	60 * 4,
	60 * 5,
	60 * 6,
	60 * 7,
	60 * 8,
	60 * 9,
	60 * 10,
	60 * 11,
	60 * 12,
	60 * 24,
	60 * 24 * 2,
	60 * 24 * 3,
	60 * 24 * 4,
}

func logScale(x int) int {
	if x <= 2 {
		return x
	}
	if x < 4 {
		return 3
	}
	if x < 8 {
		return 4
	}
	if x < 16 {
		return 5
	}
	if x < 32 {
		return 6
	}
	if x < 64 {
		return 7
	}
	if x < 128 {
		return 8
	}
	return 9
}

func display() [28]int {
	d := [28]int{}
	s := 0
	dIndex := 0
	for index, v := range histogram {
		s += v
		if index < 10 {
			d[dIndex] = logScale(v)
			dIndex++
		}
		for _, bp := range breakpoints {
			if index == bp {
				d[dIndex] = logScale(s / bp)
				dIndex++
			}
		}
	}
	return d
}

var palette [10]colorful.Color
var paletteBg [10]colorful.Color

func init() {
	{
		c1, _ := colorful.Hex("#FF0000")
		c2, _ := colorful.Hex("#00AA00")
		for i := 0; i < 10; i++ {
			d := float64(i) / float64(10)
			palette[i] = c2.BlendHsv(c1, d)
		}
	}

	{
		c1 := colorful.Color{0.0, 0.0, 0.5}
		c2 := colorful.Color{0.1, 0.0, 0.0}
		for i := 0; i < 10; i++ {
			d := float64(i) / float64(10)
			paletteBg[i] = c2.BlendHsv(c1, d)
		}
	}

	flag.StringVar(&ga, "ga", "", "ga")
	flag.StringVar(&refresh_token, "refresh_token", "", "refresh_token")
	flag.StringVar(&client_secret, "client_secret", "", "client_secret")
	flag.StringVar(&client_id, "client_id", "", "client_id")
	//flag.Parse()
}

var blink = true

func setScreen(screen lib.Screen) {
	dis := display()
	max := 0
	for _, amt := range dis {
		if amt > max {
			max = amt
		}
	}
	isLog := max > 9
	for x, amt := range dis {
		var a int
		if isLog {
			a = logScale(amt)
		} else {
			a = amt
		}
		for y := 0; y < 10; y++ {
			r, g, b := byte(0), byte(0), byte(0)
			if y < a {
				r, g, b = palette[y].RGB255()
			} else if x < 10 {
				r, g, b = paletteBg[x].RGB255()
			} else if x < 15 {
				r, g, b = paletteBg[(x-10)*2].RGB255()
			} else if x < 25 {
				r, g, b = paletteBg[x-15].RGB255()
			} else {
				r, g, b = paletteBg[(x-25)*3].RGB255()
			}
			screen.Set(27-x, 9-y, r, g, b)
		}
	}
	blink = !blink
	if blink {
		if isLog {
			screen.Set(0, 0, 255, 0, 0)
		} else {
			screen.Set(0, 0, 0, 255, 0)
		}
	}
	screen.Dump()
}

func main() {
	screen := lib.NewScreen()
	histogram = make([]int, 4320)
	for {
		for i := 0; i < 11; i++ {
			histogram[0] = getCount()
			setScreen(screen)
			time.Sleep(time.Second * 5)
		}
		histogram = append([]int{getCount()}, histogram...)[:4320]
		setScreen(screen)
		log.Println(display())
		time.Sleep(time.Second * 5)
	}
}
