package errors

import (
	"errors"
)

const callStackDepth = 10

type DetailError interface {
	error
	ErrCoder
	CallStacker
	GetRoot()  error
}


func  NewErr(errmsg string) error {
	return errors.New(errmsg)
}

func NewDetailErr(err error,errcode ErrCode,errmsg string) DetailError{
	if err == nil {return nil}

	boterr, ok := err.(botError)
	if !ok {
		boterr.root = err
		boterr.errmsg = err.Error()
		boterr.callstack = getCallStack(0, callStackDepth)
		boterr.code = errcode

	}
	if errmsg != "" {
		boterr.errmsg = errmsg + ": " + boterr.errmsg
	}


	return boterr
}

func RootErr(err error) error {
	if err, ok := err.(DetailError); ok {
		return err.GetRoot()
	}
	return err
}



