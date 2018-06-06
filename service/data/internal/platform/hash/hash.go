// Copyright 2017~2022 The Bottos Authors
// This file is part of the Bottos Chain library.
// Created by Rocket Core Team of Bottos.

//This program is free software: you can distribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.

//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.

//You should have received a copy of the GNU General Public License
// along with bottos.  If not, see <http://www.gnu.org/licenses/>.

/*
 * file description:  general Hash type
 * @Author: Gong Zibin
 * @Date:   2017-12-05
 * @Last Modified by:
 * @Last Modified time:
 */

package common

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

const (
	//HashLength the length of hash bytes
	HashLength = 32
)

type (
	//Hash the hash data
	Hash [HashLength]byte
)

//Sha256 do sha256
func Sha256(data []byte) Hash {
	hash := sha256.Sum256(data)
	return hash
}

//DoubleSha256 do sha256 twice
func DoubleSha256(data []byte) Hash {
	temp := sha256.Sum256(data)
	hash := sha256.Sum256(temp[:])
	return hash
}

//BytesToHash set hash data
func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}

//StringToHash hash string to hash data
func StringToHash(s string) Hash {
	return BytesToHash([]byte(s))
}

//EmptyHash chech hash data
func EmptyHash(h Hash) bool {
	return h == Hash{}
}

//HexToHash hex hash string to hash data
func HexToHash(s string) Hash {
	return BytesToHash(HexStringToBytes(s))
}

//ToString hash data to string
func (h Hash) ToString() string {
	return string(h[:])
}

//Bytes get hash data
func (h Hash) Bytes() []byte {
	return h[:]
}

//ToHexString hash data to hex string
func (h Hash) ToHexString() string {
	return BytesToHex(h[:])
}

//SetString set hash data from string
func (h *Hash) SetString(s string) {
	h.SetBytes([]byte(s))
}

//SetBytes  set hash data from bytes
func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-HashLength:]
	}

	copy(h[HashLength-len(b):], b)
}

//Label calc lable from hash data
func (h *Hash) Label() uint32 {
	var chainCursorLabel uint32
	chainCursorLabel = (uint32)(h[HashLength-1]) + (uint32)(h[HashLength-2])<<8 + (uint32)(h[HashLength-3])<<16 + (uint32)(h[HashLength-4])<<24

	return chainCursorLabel
}

//BytesToHex hex data to string
func BytesToHex(d []byte) string {
	return hex.EncodeToString(d)
}

//HexToBytes hex string to bytes
func HexToBytes(str string) ([]byte, error) {
	h, err := hex.DecodeString(str)

	return h, err
}

//NumberToBytes number code covert to bytes
func NumberToBytes(num interface{}, bits int) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, num)
	if err != nil {
		fmt.Println("NumberToBytes failed:", err)
	}

	return buf.Bytes()[buf.Len()-(bits/8):]
}

//HexStringToBytes hex string to bytes
func HexStringToBytes(s string) []byte {
	if len(s) > 1 {
		if s[0:2] == "0x" {
			s = s[2:]
		}
		if len(s)%2 == 1 {
			s = "0" + s
		}
		b, err := HexToBytes(s)
		if err == nil {
			return b
		}
	}
	return nil
}
