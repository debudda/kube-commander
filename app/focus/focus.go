package focus

import (
	"github.com/AnatolyRugalev/kube-commander/commander"
	"github.com/gdamore/tcell"
)

type manager struct {
	stack []commander.Widget
}

func NewFocusManager(root commander.Widget) *manager {
	root.OnFocus()
	manager := &manager{
		stack: []commander.Widget{root},
	}
	return manager
}

func (f *manager) HandleEvent(e tcell.Event, useStack bool) bool {
	for i := len(f.stack) - 1; i >= 0; i-- {
		if f.stack[i].HandleEvent(e) {
			return true
		}
		if !useStack {
			break
		}
	}
	switch ev := e.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyESC, tcell.KeyBackspace2:
			f.Blur()
			return true
		}
	}
	return false
}

func (f *manager) StackSize() int {
	return len(f.stack)
}

func (f *manager) Current() commander.Widget {
	return f.stack[len(f.stack)-1]
}

func (f *manager) Root() commander.Widget {
	return f.stack[0]
}

func (f *manager) Focus(widget commander.Widget) {
	current := f.Current()
	if current == widget {
		return
	}
	current.OnBlur()
	widget.OnFocus()

	// Find if this widget persist in focus stack
	for i, w := range f.stack {
		if w == widget {
			// Reduce focus stack to focus on found widget
			f.stack = f.stack[0 : i+1]
			return
		}
	}
	f.stack = append(f.stack, widget)
}

func (f *manager) Blur() {
	if len(f.stack) <= 1 {
		return
	}
	f.Current().OnBlur()
	f.stack = f.stack[0 : len(f.stack)-1]
	f.Current().OnFocus()
}

type Focusable struct {
	focus   bool
	visible bool
}

func (f *Focusable) OnFocus() {
	f.focus = true
}

func (f *Focusable) OnBlur() {
	f.focus = false
}

func (f *Focusable) IsFocused() bool {
	return f.focus
}

func (f *Focusable) OnShow() {
	f.visible = true
}

func (f *Focusable) OnHide() {
	f.visible = false
}

func (f *Focusable) IsVisible() bool {
	return f.visible
}

func NewFocusable() *Focusable {
	return &Focusable{}
}
