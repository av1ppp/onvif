package onvif

type Req[B any] struct {
	Body     B
	Header   any
	Endpoint string
}

func Request[T any](body T) *Req[T] {
	return &Req[T]{
		Body: body,
	}
}
