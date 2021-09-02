package bugui

import (
	"github.com/containerd/console"
	"golang.org/x/crypto/ssh/terminal"
)

type Terminal struct {
	prev console.Console
	tty  *terminal.Terminal
}

func NewTerminal() *Terminal {
	current := console.Current()
	newterm := Terminal{prev: current}

	if err := newterm.prev.SetRaw(); err != nil {
		panic(err)
	}

	newterm.tty = terminal.NewTerminal(newterm.prev, "")
	newterm.tty.AutoCompleteCallback = newterm.callback

	return &newterm
}

func (terminal Terminal) Callback(fn func(string, int, rune) (string, int, bool)) {
	terminal.tty.AutoCompleteCallback = fn
}

func (terminal *Terminal) Readline() (string, error) {
	return terminal.tty.ReadLine()
}

func (terminal *Terminal) Close() {
	terminal.prev.Reset()
}

func (terminal *Terminal) callback(line string, pos int, key rune) (newline string, newpos int, ok bool) {
	//fmt.Println("callback: ", line, pos, key)
	return "", 0, false
}
