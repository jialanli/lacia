package lacia

import (
	"fmt"
)

/*
	concurrency goroutines for you. many tasks...
*/

type takeHelper []error

func (ts takeHelper) Error() (res string) {
	if len(ts) == 0 {
		return ""
	}

	if len(ts) == 1 {
		return ts[0].Error()
	}

	res = fmt.Sprintf("[%s", ts[0].Error())
	for i := 1; i < len(ts); i++ {
		res += fmt.Sprintf("; %s", ts[i].Error())
	}
	res += "]"
	return
}

func (ts takeHelper) Errors() []error {
	return []error(ts)
}

func NewTakeHelper(errList []error) ErrInterface {
	if len(errList) == 0 {
		return nil
	}
	var errs []error
	for _, e := range errList {
		if e != nil {
			errs = append(errs, e)
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return takeHelper(errs) //转换类型
}

//all exec success, return nil
func ConcurrencyGos(execList ...func() error) ErrInterface {
	errChan := make(chan error, len(execList))
	for _, f := range execList {
		go func(exec func() error) {
			errChan <- exec()
		}(f)
	}

	errs := make([]error, 0)
	for i := 0; i < cap(errChan); i++ {
		if err := <-errChan; err != nil {
			errs = append(errs, err)
		}
	}

	return NewTakeHelper(errs)
}
