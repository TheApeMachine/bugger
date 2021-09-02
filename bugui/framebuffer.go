package bugui

import "fmt"

type Framebuffer struct {
	frame string
}

func (framebuffer Framebuffer) Play() {
	fmt.Println(framebuffer.frame)
}
