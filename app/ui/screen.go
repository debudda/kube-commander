package ui

import (
	"github.com/AnatolyRugalev/kube-commander/app/focus"
	"github.com/AnatolyRugalev/kube-commander/app/ui/theme"
	"github.com/AnatolyRugalev/kube-commander/commander"
	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/views"
)

type Screen struct {
	*views.Panel
	*focus.Focusable

	app       commander.App
	workspace commander.Workspace
	status    commander.StatusReporter
	view      commander.View
}

func (s *Screen) View() commander.View {
	return s.view
}

func (s *Screen) SetView(view views.View) {
	s.view = view
	s.Panel.SetView(view)
}

func (s *Screen) SetStatus(stat commander.StatusReporter) {
	s.status = stat
	s.Panel.SetStatus(s.status)
}

func (s *Screen) UpdateScreen() {
	if s.app != nil {
		s.app.Update()
	}
}

func (s *Screen) SetWorkspace(workspace commander.Workspace) {
	s.workspace = workspace
	s.SetContent(s.workspace)
}

func (s Screen) Workspace() commander.Workspace {
	return s.workspace
}

func NewScreen(app commander.App) *Screen {
	s := Screen{
		Panel:     views.NewPanel(),
		Focusable: focus.NewFocusable(),
		app:       app,
	}

	title := views.NewTextBar()
	title.SetStyle(tcell.StyleDefault.
		Background(tcell.ColorTeal).
		Foreground(tcell.ColorWhite))
	title.SetCenter("kube-commander", theme.Default)
	title.SetRight(app.Config().Context(), theme.Default)

	s.SetTitle(title)

	return &s
}

func (s Screen) HandleEvent(e tcell.Event) bool {
	switch ev := e.(type) {
	case *tcell.EventKey:
		if ev.Rune() == 'q' && ev.Modifiers() == tcell.ModNone {
			s.app.Quit()
			return true
		}
	}
	return s.BoxLayout.HandleEvent(e)
}
