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

package util
//TxInfo struct
type TxInfo struct {
	TransactionID string `json:"transaction_id"`
	Transaction   struct {
		RefBlockNum    uint64        `json:"ref_block_num"`
		RefBlockPrefix uint64        `json:"ref_block_prefix"`
		Expiration     string        `json:"expiration"`
		Scope          []string      `json:"scope"`
		Signatures     []interface{} `json:"signatures"`
		Messages       []struct {
			Code          string        `json:"code"`
			Type          string        `json:"type"`
			Authorization []interface{} `json:"authorization"`
			Data          struct {
				UserName  string `json:"user_name"`
				BasicInfo struct {
					Info string `json:"info"`
				} `json:"basic_info"`
			} `json:"data"`
			HexData string `json:"hex_data"`
		} `json:"messages"`
		Output []struct {
			Notify       []interface{} `json:"notify"`
			DeferredTrxs []interface{} `json:"deferred_trxs"`
		} `json:"output"`
	} `json:"transaction"`
}
//TxDBInfo struct
type TxDBInfo struct {
	TransactionID string `json:"transaction_id"`
	From          string `json:"from"`
	To            string `json:"to"`
	Price         uint64 `json:"price"`
	Type          uint64 `json:"type"`
	Date          string `json:"date"`
	BlockId       uint64 `json:"block_id"`
}
//TransferDBInfo struct
type TransferDBInfo struct {
	TransactionID string `json:"tx_id"`
	From          string `json:"from"`
	To            string `json:"to"`
	Price         uint64 `json:"price"`
	TxTime        string `json:"tx_time"`
	BlockNum      uint64 `json:"block_num"`
}
//InserTxInfoSql const
const InserTxInfoSql string = "INSERT INTO txinfo(TransactionID,Price,Type,From,To,Date,BlockId) values(?,?,?,?,?,?,?)"
