package bugui

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

type Screen struct {
	terminal    *Terminal
	framebuffer chan Framebuffer
}

func NewScreen() *Screen {
	return &Screen{
		terminal:    NewTerminal(),
		framebuffer: make(chan Framebuffer),
	}
}

func (screen Screen) Render() *sync.WaitGroup {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			cancel()
			fmt.Println(sig)
		}
	}()

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		defer close(screen.framebuffer)

		for {
			select {
			case <-ctx.Done():
				break
			case buffer := <-screen.framebuffer:
				buffer.Play()
			default:
			}
		}
	}(&wg)

	return &wg
}
