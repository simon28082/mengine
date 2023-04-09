package errors

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	defaultId         = `default`
	defaultStatus     = -1
	defaultStatusText = `status error`
)

type Error interface {
	error

	fmt.Stringer

	Status() int

	StatusText() string

	ID() string

	Unwrap() []error

	Is(err error) bool

	As(err error) bool
}

type basicError struct {
	wraps      []error
	id         string
	message    string
	status     int
	statusText string
}

func (b *basicError) String() string {
	var builder strings.Builder

	if len(b.id) > 0 {
		builder.WriteString(`(`)
		builder.WriteString(b.id)
		if b.status > 0 {
			builder.WriteString(`:`)
			builder.WriteString(strconv.Itoa(b.status))
		}
		builder.WriteString(`)`)
	}

	if len(b.message) > 0 {
		builder.WriteString(` `)
		builder.WriteString(b.message)
	}

	wrapLength := len(b.wraps)
	if wrapLength > 0 {
		builder.WriteString(` <- `)

		builder.WriteString(`[`)
		for i := 0; i < wrapLength; i++ {
			builder.WriteString(b.wraps[i].Error())
			if i != wrapLength-1 {
				builder.WriteString(`,`)
			}
		}
		builder.WriteString(`]`)
	}

	return builder.String()
}

func (b *basicError) Unwrap() []error {
	return b.wraps
}

func (b *basicError) Error() string {
	return b.String()
}

func (b *basicError) Status() int {
	return b.status
}

func (b *basicError) StatusText() string {
	return b.statusText
}

func (b *basicError) ID() string {
	return b.id
}

func (b *basicError) Is(err error) bool {
	if len(b.wraps) > 0 {
		for i := 0; i < len(b.wraps); i++ {
			if errors.Is(err, b.wraps[i]) {
				return true
			}
		}
	}
	return false
}

func (b *basicError) As(target any) bool {
	if len(b.wraps) > 0 {
		for i := 0; i < len(b.wraps); i++ {
			if errors.As(b.wraps[i], target) {
				return true
			}
		}
	}
	return false
}

func New(id string, status int, statusText, message string) error {
	return &basicError{
		id:         id,
		status:     status,
		statusText: statusText,
		message:    message,
	}
}

func NewDefault(message string) error {
	return New(defaultId, defaultStatus, defaultStatusText, message)
}

func Errorf(format string, args ...any) error {
	return NewDefault(fmt.Sprintf(format, args...))
}

//func WithMessage(message string, errs ...error) error {
//	return &basicError{
//		wraps:   errs,
//		message: message,
//	}
//}

func WithErrorf(format string, args ...any) error {
	var (
		errs   []error
		others []any
	)
	if len(args) > 0 {
		for i := range args {
			if target, ok := args[i].(error); ok {
				errs = append(errs, target)
			} else {
				others = append(others, target)
			}
		}
	}
	return &basicError{
		wraps:   errs,
		message: fmt.Sprintf(format, others...),
	}
}
