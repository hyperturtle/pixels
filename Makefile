all: bin/loop bin/main bin/reaction-diffusion bin/sparkle bin/walk

bin/%: scenes/%.go deps
	go build -o $@ $<

deps:
	go get ./...