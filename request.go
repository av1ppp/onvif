package onvif

import "github.com/av1ppp/logx"

type Req[B any] struct {
	Body   B
	Header any
	Logger *logx.Logger
}

func Request[T any](body T) *Req[T] {
	return &Req[T]{
		Body: body,
	}
}
