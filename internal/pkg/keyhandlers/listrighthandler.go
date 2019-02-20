package keyhandlers

import (
	"github.com/jroimartin/gocui"
	"github.com/lawrencegripper/azbrowse/internal/pkg/views"
)

const listRightId = "ListRight"

type ListRightHandler struct {
	List            *views.ListWidget
	EditModeEnabled *bool
}

func (h ListRightHandler) Id() string {
	return listRightId
}

func (h ListRightHandler) Fn() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tmp := false // memory leak?
		h.EditModeEnabled = &tmp
		g.Cursor = true
		g.SetCurrentView("itemWidget")
		return nil
	}
}

func (h ListRightHandler) Widget() string {
	return "itemWidget"
}

func (h ListRightHandler) DefaultKey() gocui.Key {
	return gocui.KeyArrowLeft
}