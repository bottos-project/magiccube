package test

import(
	"testing"
)

func TestRegitser(t *testing.T) {

}

type data struct {
	code int64
	msg string

}

func TestS(t *testing.T){
	var d = data{msg:"23432"}
	t.Log(d.code)
}
