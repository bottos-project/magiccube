package error

import "testing"

func Test_getErrorInfo(t *testing.T) {
	t.Log(GetErrorInfo(1002, ""))
}

func Test_GetAllErrorInfos(t *testing.T) {
	t.Log(GetAllErrorInfos(""))
}