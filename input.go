package termloop

import "github.com/nsf/termbox-go"

type Input struct {
	endKey termbox.Key
	eventQ chan termbox.Event
	ctrl   chan bool
}

func (input Input) GetEventQ() chan termbox.Event {
	return input.eventQ
}

func newInput() *Input {
	i := Input{eventQ: make(chan termbox.Event),
		ctrl:   make(chan bool, 2),
		endKey: termbox.KeyCtrlC}
	return &i
}

func (i *Input) start() {
	go poll(i)
}

func (i *Input) stop() {
	i.ctrl <- true
}

func poll(i *Input) {
loop:
	for {
		select {
		case <-i.ctrl:
			break loop
		default:
			i.eventQ <- termbox.PollEvent()
		}
	}
}
