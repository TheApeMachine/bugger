package bugui

type Layout interface {
	Buffer(chan Framebuffer)
}

func NewLayout(layoutType Layout) Layout {
	return layoutType
}

type Base struct {
	framebuffer chan Framebuffer
}

func (layout Base) Buffer() chan Framebuffer {
	return layout.framebuffer
}
