package errors

type botError struct {
	errmsg string
	callstack *CallStack
	root error
	code ErrCode
}

func (e botError) Error() string {
	return e.errmsg
}

func (e botError) GetErrCode()  ErrCode {
	return e.code
}

func (e botError) GetRoot()  error {
	return e.root
}

func (e botError) GetCallStack()  *CallStack {
	return e.callstack
}
