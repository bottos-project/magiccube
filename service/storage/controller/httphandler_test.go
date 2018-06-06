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
	"fmt"
	"github.com/bottos-project/magiccube/service/storage/util"
	"testing"
)

func TestGetInfo(t *testing.T) {
	var rsp *util.Info
	rsp, _ = GetInfo()
	fmt.Println(rsp.HeadBlockNum)
	fmt.Println(rsp.ServerVersion)
}
func TestGetBlock(t *testing.T) {
	var rsp *util.Block
	rsp, _ = GetBlock("0000000445a9f27898383fd7de32835d5d6a978cc14ce40d9f327b5329de796b")
	fmt.Println(rsp.Producer)
	fmt.Println(rsp.ID)
}
func TestGetBlockNum1(t *testing.T) {
	var rsp *util.Block
	rsp, _ = GetBlock("1")
	fmt.Println(rsp.Producer)
	fmt.Println(rsp.ID)
}
func TestGetAccountInfo(t *testing.T) {
	var rsp *util.AccountInfo
	rsp, _ = GetAccountInfo()
	fmt.Println(rsp.AccountName)
}
func TestGetTxInfo(t *testing.T) {
	var rsp *util.TxInfo
	rsp, _ = GetTxInfo()
	fmt.Println(rsp.TransactionID)
	fmt.Println(rsp.Transaction.Expiration)
}
func TestGetCodeInfo(t *testing.T) {
	fmt.Println("ddd")
}
