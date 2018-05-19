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
	"encoding/binary"
	"encoding/hex"
	"crypto/sha256"
	"fmt"
)

const (
	HashLength = 32
)

type (
	Hash [HashLength]byte
)

func Sha256(data []byte) Hash {
	hash := sha256.Sum256(data)
	return hash
}

func DoubleSha256(data []byte) Hash {
	temp := sha256.Sum256(data)
	hash := sha256.Sum256(temp[:])
	return hash
}

func BytesToHash(b []byte) Hash {
	var h Hash
	h.SetBytes(b)
	return h
}

func StringToHash(s string) Hash {
	return BytesToHash([]byte(s))
}

func EmptyHash(h Hash) bool {
	return h == Hash{}
}

func HexToHash(s string) Hash {
	return BytesToHash(HexStringToBytes(s))
}

func (h Hash) ToString() string {
	return string(h[:])
}

func (h Hash) Bytes() []byte {
	return h[:]
}

func (h Hash) ToHexString() string {
	return BytesToHex(h[:])
}

func (h *Hash) SetString(s string) {
	h.SetBytes([]byte(s))
}

func (h *Hash) SetBytes(b []byte) {
	if len(b) > len(h) {
		b = b[len(b)-HashLength:]
	}

	copy(h[HashLength-len(b):], b)
}

func BytesToHex(d []byte) string {
	return hex.EncodeToString(d)
}

func HexToBytes(str string) []byte {
	h, _ := hex.DecodeString(str)

	return h
}

func NumberToBytes(num interface{}, bits int) []byte {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, num)
	if err != nil {
		fmt.Println("NumberToBytes failed:", err)
	}

	return buf.Bytes()[buf.Len()-(bits/8):]
}

func HexStringToBytes(s string) []byte {
	if len(s) > 1 {
		if s[0:2] == "0x" {
			s = s[2:]
		}
		if len(s)%2 == 1 {
			s = "0" + s
		}
		return HexToBytes(s)
	}
	return nil
}
