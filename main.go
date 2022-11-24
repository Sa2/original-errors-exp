package main

import (
	"errors"
	"fmt"
	"github.com/Sa2/original-errors/originalErrors"
	"golang.org/x/xerrors"
)

func main() {
	basicErr := newError()
	xErr := newXError()
	orignalErr := newOriginalError()

	// standard
	fmt.Println(basicErr.Error())
	fmt.Println(xErr.Error())
	fmt.Println(orignalErr.Error())

	fmt.Println(orignalErr.Error())
	fmt.Println(orignalErr)

	fmt.Println(xerrors.Errorf("xerrors.Errorf", xErr))
	fmt.Println(xerrors.Errorf("xerrors.Errorf", orignalErr))

	return
}

func newError() error {
	return errors.New("errors!")
}

func newXError() error {
	return xerrors.New("xerrors!")
}

func newOriginalError() error {
	return originalErrors.New("originalErrors!", originalErrors.ErrorType0)
}
