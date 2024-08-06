package features

import (
	tele "gopkg.in/telebot.v3"
)

type Feature struct {
	Handler    tele.HandlerFunc
	Middleware []tele.MiddlewareFunc
}

var features = make(map[interface{}]*Feature)

func RegisterFeature(endpoint interface{}, handler tele.HandlerFunc, middleware ...tele.MiddlewareFunc) {
	features[endpoint] = &Feature{
		Handler:    handler,
		Middleware: middleware,
	}
}

func Handle(b *tele.Bot) {
	for endpoint, feature := range features {
		b.Handle(endpoint, feature.Handler, feature.Middleware...)
	}
}
