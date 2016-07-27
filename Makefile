.PHONY: all
all: deps
	go build github.com/hyperturtle/pixels/scenes/loop 
	go build github.com/hyperturtle/pixels/scenes/main 
	go build github.com/hyperturtle/pixels/scenes/reaction-diffusion 
	go build github.com/hyperturtle/pixels/scenes/sparkle 
	go build github.com/hyperturtle/pixels/scenes/walk

.PHONY: deps
deps:
	go get ./...