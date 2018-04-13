package controller
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

func MapToObject(source interface{}, dst interface{}) error {
	b, err := json.Marshal(source)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}

func ParseQuantity(q string) (int64, error) {
	return strconv.ParseInt(q, 0, 64)
}

func EncodeJSON(data interface{}) ([]byte, error) {
	encoded, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return encoded, nil
}

func EncodeJSONString(data interface{}) (string, error) {
	encoded, err := EncodeJSON(data)
	if err != nil {
		return "", err
	}
	return string(encoded), err
}

func EncodeJSONToBuffer(data interface{}, b *bytes.Buffer) error {
	encoded, err := EncodeJSON(data)
	if err != nil {
		return err
	}
	_, err = b.Write(encoded)
	return err
}

type JSON2Request struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Params  interface{} `json:"params,omitempty"`
	Method  string      `json:"method,omitempty"`
}

func (e *JSON2Request) JSONByte() ([]byte, error) {
	return EncodeJSON(e)
}

func (e *JSON2Request) JSONString() (string, error) {
	return EncodeJSONString(e)
}

func (e *JSON2Request) JSONBuffer(b *bytes.Buffer) error {
	return EncodeJSONToBuffer(e, b)
}

func (e *JSON2Request) String() string {
	str, _ := e.JSONString()
	return str
}

func NewJSON2RequestBlank() *JSON2Request {
	j := new(JSON2Request)
	j.JSONRPC = "2.0"
	return j
}

func NewJSON2Request(id, params interface{}, method string) *JSON2Request {
	j := new(JSON2Request)
	j.JSONRPC = "2.0"
	j.ID = id
	j.Params = params
	j.Method = method
	return j
}

func ParseJSON2Request(request string) (*JSON2Request, error) {
	j := new(JSON2Request)
	err := json.Unmarshal([]byte(request), j)
	if err != nil {
		return nil, err
	}
	if j.JSONRPC != "2.0" {
		return nil, fmt.Errorf("Invalid JSON RPC version - `%v`, should be `2.0`", j.JSONRPC)
	}
	return j, nil
}

type JSON2Response struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Error   *JSONError  `json:"error,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

func (e *JSON2Response) JSONByte() ([]byte, error) {
	return EncodeJSON(e)
}

func (e *JSON2Response) JSONString() (string, error) {
	return EncodeJSONString(e)
}

func (e *JSON2Response) JSONBuffer(b *bytes.Buffer) error {
	return EncodeJSONToBuffer(e, b)
}

func (e *JSON2Response) String() string {
	str, _ := e.JSONString()
	return str
}

func NewJSON2Response() *JSON2Response {
	j := new(JSON2Response)
	j.JSONRPC = "2.0"
	return j
}

func (j *JSON2Response) AddError(code int, message string, data interface{}) {
	e := NewJSONError(code, message, data)
	j.Error = e
}

type JSONError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewJSONError(code int, message string, data interface{}) *JSONError {
	j := new(JSONError)
	j.Code = code
	j.Message = message
	j.Data = data
	return j
}