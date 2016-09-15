package check

import "errors"

type Check struct{}
type Str struct{}

func C() *Check                { return &Check{} }
func (check *Check) Str() *Str { return &Str{} }

func (str *Str) NotEmpty(val, msg string) {
	if val == "" {
		panicErr(errors.New(msg))
	}
}

func panicErr(err error) {
	panic(err)
}
