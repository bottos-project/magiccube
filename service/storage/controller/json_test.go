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
	"testing"

	//. "github.com/xiaoping378/blockchain-on-sql/parser"
	"fmt"
)
func TestMapToObject(t *testing.T) {
	fmt.Print("dddd")
}
/*
func TestParseQuantity(t *testing.T) {
	i, err := ParseQuantity("0x41")
	if err != nil {
		t.Errorf("%v", err)
	}
	if i != 65 {
		t.Errorf("Returned wrong int value")
	}

	i, err = ParseQuantity("0x400")
	if err != nil {
		t.Errorf("%v", err)
	}
	if i != 1024 {
		t.Errorf("Returned wrong int value")
	}

	i, err = ParseQuantity("0x0")
	if err != nil {
		t.Errorf("%v", err)
	}
	if i != 0 {
		t.Errorf("Returned wrong int value")
	}
}

func TestMapToObject(t *testing.T) {
	list := []string{"Hello", "bye"}
	l2 := []string{}

	err := MapToObject(list, &l2)
	if err != nil {
		t.Errorf("%v", err)
	}

	if list[0] != l2[0] || list[1] != l2[1] {
		t.Errorf("Invalid list returned")
	}
}
*/