package main

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

func HelloVecty() {
	//vecty.SetTitle("Hello Vecty")
	vecty.RenderBody(&view{})
}

type view struct {
	vecty.Core
}

func (v *view) Render() vecty.ComponentOrHTML {
	markdown := HelloGoldMark()
	return elem.Body(
		vecty.Text("Hello Vecty!"),
		vecty.Markup(
			vecty.UnsafeHTML(markdown),
		),
	)
}
