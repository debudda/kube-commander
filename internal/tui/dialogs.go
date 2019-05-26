package tui

import (
	"strings"

	"github.com/gizak/termui/v3"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

// Buttons text
const (
	ButtonOk     = "OK"
	ButtonCancel = "Cancel"
	ButtonYes    = "Yes"
	ButtonNo     = "No"

	buttonWidth        = 10
	buttonHeight       = 4
	buttonMargingRight = 3
)

// Button represents text button
type Button struct {
	*widgets.Paragraph
}

// NewButton returns button with specified text
func NewButton(text string) Button {
	b := Button{}
	p := widgets.NewParagraph()
	p.Text = text
	b.Paragraph = p
	return b
}

func (b *Button) setRect(x1, y1, x2, y2 int) {
	b.Paragraph.SetRect(x1, y1, x2, y2)
}

// Dialog represents modal dialog window
type Dialog struct {
	*widgets.Paragraph
	Buttons             []Button
	buttonStyle         ui.Style
	selectedButtonStyle ui.Style
	selectedButton      int
}

// Draw draws dialog and buttons
func (dlg *Dialog) Draw(buf *ui.Buffer) {
	dlg.Paragraph.Draw(buf)
	for idx, btn := range dlg.Buttons {
		if idx == dlg.selectedButton {
			btn.Paragraph.BorderStyle = dlg.selectedButtonStyle
		} else {
			btn.Paragraph.BorderStyle = dlg.buttonStyle
		}
		btn.Paragraph.Draw(buf)
	}
}

func (dlg *Dialog) addButton(b Button) {
	dlg.Buttons = append(dlg.Buttons, b)
}

func (dlg *Dialog) setButtonsRect(x1, y1, x2, y2 int) {
	var bx1, by1, bx2, by2, left, buttonsWidth int
	buttonsWidth = (buttonWidth + buttonMargingRight) * len(dlg.Buttons)
	left = x1 + (x2-x1)/2 - buttonsWidth/2 - 1

	for i := 0; i < len(dlg.Buttons); i++ {
		bx1 = left
		by1 = y2 - buttonHeight
		bx2 = bx1 + buttonWidth
		by2 = y2 - 1
		dlg.Buttons[i].setRect(bx1, by1, bx2, by2)
		left += buttonWidth + buttonMargingRight
	}

}

func maxLinesWidth(arr []string) int {
	var maxValue int
	for i := 0; i < len(arr); i++ {
		if len(arr[i]) > maxValue {
			maxValue = len(arr[i])
		}
	}
	return maxValue
}

func (dlg *Dialog) setRect() {
	termWidth, termHeight := ui.TerminalDimensions()
	lines := strings.Split(dlg.Paragraph.Text, "\n")
	lineWidth := maxLinesWidth(lines) + 2
	if len(dlg.Paragraph.Title) > lineWidth {
		lineWidth = len(dlg.Paragraph.Title) + 5
	}

	x1 := termWidth/2 - lineWidth/2 - 1
	y1 := termHeight/2 - 5
	x2 := x1 + lineWidth + dlg.Paragraph.PaddingLeft + dlg.Paragraph.PaddingRight
	y2 := y1 + len(lines) + buttonHeight + 2

	dlg.Paragraph.SetRect(x1, y1, x2, y2)
	dlg.setButtonsRect(x1, y1, x2, y2)
}

func newDialog(title, text string, buttons ...string) *Dialog {
	p := widgets.NewParagraph()
	p.Title = title
	p.Text = text
	p.BorderStyle = theme["dialog"].active
	p.TitleStyle = theme["dialog"].active
	p.TextStyle = theme["dialog"].active
	p.PaddingLeft = 1
	p.PaddingRight = 1

	newDlg := &Dialog{Paragraph: p}
	newDlg.buttonStyle = theme["button"].inactive
	newDlg.selectedButtonStyle = theme["button"].active

	if len(buttons) == 0 {
		newDlg.addButton(NewButton(ButtonOk))
	} else {
		for _, txtButton := range buttons {
			newDlg.addButton(NewButton(txtButton))
		}
	}

	newDlg.setRect()

	return newDlg
}

func (dlg Dialog) dialogResult() string {
	return dlg.Buttons[dlg.selectedButton].Paragraph.Text
}

func (dlg *Dialog) nextButton() {
	if dlg.selectedButton < len(dlg.Buttons)-1 {
		dlg.selectedButton++
	}
}

func (dlg *Dialog) prevButton() {
	if dlg.selectedButton > 0 {
		dlg.selectedButton--
	}
}

func (Dialog) OnEvent(event *termui.Event) bool {
	return false
}

func (Dialog) OnFocusIn() {
}

func (Dialog) OnFocusOut() {
}

// ShowDialog shows modal dialog and waits for button pressed
func ShowDialog(title, text string, buttons ...string) string {
	dlg := newDialog(title, text, buttons...)

	screen.Focus(dlg)
	ui.Render(dlg)
	defer screen.popFocus()

	uiEvents := ui.PollEvents()

	for {
		e := <-uiEvents
		switch e.ID {
		case "<Right>":
			dlg.nextButton()
			ui.Render(dlg)
		case "<Left>":
			dlg.prevButton()
			ui.Render(dlg)
		case "<Enter>":
			return dlg.dialogResult()
		case "<Escape>":
			return ""
		}
	}

}
