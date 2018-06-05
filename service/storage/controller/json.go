/*Copyright 2017~2022 The Bottos Authors
  This file is part of the Bottos Data Exchange Client
  Created by Developers Team of Bottos.

  This program is free software: you can distribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with Bottos. If not, see <http://www.gnu.org/licenses/>.
*/

package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

// MapToObject is to conver data type
func MapToObject(source interface{}, dst interface{}) error {
	b, err := json.Marshal(source)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}

// ParseQuantity is to parse quantity
func ParseQuantity(q string) (int64, error) {
	return strconv.ParseInt(q, 0, 64)
}

//EncodeJSON is to encode json
func EncodeJSON(data interface{}) ([]byte, error) {
	encoded, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return encoded, nil
}

// EncodeJSONString is to encode json string
func EncodeJSONString(data interface{}) (string, error) {
	encoded, err := EncodeJSON(data)
	if err != nil {
		return "", err
	}
	return string(encoded), err
}

// EncodeJSONToBuffer is to encode json to buffer
func EncodeJSONToBuffer(data interface{}, b *bytes.Buffer) error {
	encoded, err := EncodeJSON(data)
	if err != nil {
		return err
	}
	_, err = b.Write(encoded)
	return err
}

// JSON2Request is definition of json rpc struct
type JSON2Request struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Params  interface{} `json:"params,omitempty"`
	Method  string      `json:"method,omitempty"`
}

// JSONByte is to convert json to byte
func (e *JSON2Request) JSONByte() ([]byte, error) {
	return EncodeJSON(e)
}

// JSONString is to covert json to byte
func (e *JSON2Request) JSONString() (string, error) {
	return EncodeJSONString(e)
}

// JSONBuffer is to convert json to buffer
func (e *JSON2Request) JSONBuffer(b *bytes.Buffer) error {
	return EncodeJSONToBuffer(e, b)
}

// String is to get string format
func (e *JSON2Request) String() string {
	str, _ := e.JSONString()
	return str
}

// NewJSON2RequestBlank is to construct a request struct
func NewJSON2RequestBlank() *JSON2Request {
	j := new(JSON2Request)
	j.JSONRPC = "2.0"
	return j
}

// NewJSON2Request is to construct a request struct
func NewJSON2Request(id, params interface{}, method string) *JSON2Request {
	j := new(JSON2Request)
	j.JSONRPC = "2.0"
	j.ID = id
	j.Params = params
	j.Method = method
	return j
}

// ParseJSON2Request is to parse json request
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

// JSON2Response is definition of jsonrpc rsp
type JSON2Response struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Error   *JSONError  `json:"error,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}

// JSONByte is to encode json
func (e *JSON2Response) JSONByte() ([]byte, error) {
	return EncodeJSON(e)
}

// JSONString is to encode json to string
func (e *JSON2Response) JSONString() (string, error) {
	return EncodeJSONString(e)
}

// JSONBuffer is to encode json to buffer
func (e *JSON2Response) JSONBuffer(b *bytes.Buffer) error {
	return EncodeJSONToBuffer(e, b)
}

// String is to get string format data
func (e *JSON2Response) String() string {
	str, _ := e.JSONString()
	return str
}

// NewJSON2Response is to construct json rpc rsp
func NewJSON2Response() *JSON2Response {
	j := new(JSON2Response)
	j.JSONRPC = "2.0"
	return j
}

// AddError is to add error
func (e *JSON2Response) AddError(code int, message string, data interface{}) {
	jsonError := NewJSONError(code, message, data)
	e.Error = jsonError
}

// JSONError is definition of json error
type JSONError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// NewJSONError is to construct json erro
func NewJSONError(code int, message string, data interface{}) *JSONError {
	j := new(JSONError)
	j.Code = code
	j.Message = message
	j.Data = data
	return j
}
