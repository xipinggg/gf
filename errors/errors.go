package errors

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"time"
)

var skip = 1

func Join(errs ...error) error {
	return errors.Join(errs...)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}

func Unwrap(err error) error {
	return errors.Unwrap(err)
}

func Cause(err error) error {
	return gerror.Cause(err)
}

func New(text string) error {
	return gerror.NewSkip(skip, text)
}

func NewT(text string) error {
	return gerror.NewSkip(skip, t()+text)
}

func Newf(format string, a ...any) error {
	return gerror.NewSkipf(skip, format, a...)
}

func NewfT(format string, a ...any) error {
	return gerror.NewSkipf(skip, fmt.Sprintf(t()+format, a...))
}

func Wrap(err error, text string) error {
	return gerror.WrapSkip(skip, err, text)
}

func WrapT(err error, text string) error {
	return gerror.WrapSkip(skip, err, fmt.Sprintf(t()+text))
}

func Wrapf(err error, format string, a ...any) error {
	return gerror.WrapSkipf(skip, err, format, a...)
}

func WrapfT(err error, format string, a ...any) error {
	return gerror.WrapSkipf(skip, err, fmt.Sprintf(t()+format, a...))
}

func WithMassage(err error, text string) error {
	return gerror.NewOption(gerror.Option{
		Error: err,
		Stack: false,
		Text:  text,
		Code:  nil,
	})
}

func WithMassageT(err error, text string) error {
	return gerror.NewOption(gerror.Option{
		Error: err,
		Stack: false,
		Text:  t() + " " + text,
		Code:  nil,
	})
}

func WithMassagef(err error, format string, a ...any) error {
	return WithMassage(err, fmt.Sprintf(format, a...))
}

func WithMassagefT(err error, format string, a ...any) error {
	return WithMassage(err, fmt.Sprintf(t()+format, a...))
}

func t() string {
	return time.Now().Format(time.DateTime) + " "
}
