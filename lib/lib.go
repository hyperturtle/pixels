package lib

import (
	"flag"
	"github.com/tarm/serial"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"time"
)

const (
	HEIGHT = 10
	WIDTH  = 28
)

type Screen struct {
	Data []byte
}

func NewScreen() Screen {
	s := Screen{
		Data: make([]byte, HEIGHT*WIDTH*3+1),
	}
	s.Data[0] = byte(1)
	return s
}

func (s Screen) Get(x, y int) (r, g, b byte) {
	if y%2 == 1 {
		x = WIDTH - x - 1
	}
	index := (y*WIDTH+x)*3 + 1

	return s.Data[index], s.Data[index+1], s.Data[index+2]
}

func (s Screen) Set(x, y int, r, g, b byte) {
	if x >= WIDTH || y >= HEIGHT {
		panic("out of bounds")
	}
	if y%2 == 1 {
		x = WIDTH - x - 1
	}
	index := (y*WIDTH+x)*3 + 1

	s.Data[index] = r
	s.Data[index+1] = g
	s.Data[index+2] = b
}

func (s Screen) Dump() {
	NextFrame <- &s
}

var NextFrame chan *Screen

var baud int

func init() {
	var serial string
	flag.StringVar(&serial, "serial", "", "serial port")
	flag.IntVar(&baud, "baud", 500000, "baud rate")
	flag.Parse()

	NextFrame = make(chan *Screen)

	if serial == "" {
		go webSocketServer()
		return
	}
	go serialServer(serial)
}

func serialServer(serialPort string) {
	log.Println("starting")
	c := &serial.Config{Name: serialPort, Baud: baud}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	for err == nil {
		nf := <-NextFrame
		_, err = s.Write(nf.Data)
		if err != nil {
			log.Fatal("Write", err.Error())
		}
		time.Sleep(time.Second / 30)
	}

	log.Fatal(err)
}

func webSocketServer() {
	http.Handle("/echo", websocket.Handler(echoServer))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	log.Println("starting")
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

var currentWS *websocket.Conn

func echoServer(ws *websocket.Conn) {
	var err error

	if currentWS != nil {
		currentWS.Close()
	}
	currentWS = ws

	for err == nil {
		nf := <-NextFrame
		err = websocket.Message.Send(ws, nf.Data)
		time.Sleep(time.Second / 30)
	}
	log.Println(err)
}
