package httperror

import (
	"context"
)

type Error struct {
	message string
	code    string
	status  int
	ctx     context.Context
}

func (e Error) Error() string {
	return e.message
}

func (e Error) Code() string {
	return e.code
}

func (e Error) Status() int {
	return e.status
}

func (e Error) Context() context.Context {
	return e.ctx
}
