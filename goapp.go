package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"log"
	"net/http"
)

func HelloGoApp() {
	app.Route("/", &hello{})
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

type hello struct {
	app.Compo
	name string
}

func (h *hello) Render() app.UI {
	markdown := HelloGoldMark()

	return app.Div().Body(
		app.H1().Body(
			app.Text("Hello, "),
			app.If(
				len(h.name) != 0,
				app.Text(h.name),
			).Else(
				app.Text("World!"),
			),
		),
		app.P().Body(
			app.Input().
				Type("text").
				Value(h.name).
				Placeholder("What is your name?").
				AutoFocus(true).
				OnChange(h.ValueTo(&h.name)),
		),
		app.Mark().Text(markdown),
	)
}
