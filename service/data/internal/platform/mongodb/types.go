/*Copyright 2017~2022 The Bottos Authors
  This file is part of the Bottos Service Layer
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
package mongodb

import (
	"math/big"

	"gopkg.in/mgo.v2/bson"
)

const (
	// BloomByteLength represents the number of bytes used in a header log bloom.
	BloomByteLength = 256

	// BloomBitLength represents the number of bits used in a header log bloom.
	BloomBitLength = 8 * BloomByteLength
)

// Bloom represents a 2048 bit bloom filter.
type Bloom [BloomByteLength]byte

// A BlockNonce is a 64-bit hash which proves (combined with the
// mix-hash) that a sufficient amount of computation has been carried
// out on a block.
type BlockNonce [8]byte

// MBlock represents a block header in Mongodbq.
type MBlock struct {
	Difficulty      bson.Decimal128 `bson:"difficulty"`
	Extra           string          `bson:"extraData"`
	GasLimit        bson.Decimal128 `bson:"gasLimit"`
	GasUsed         bson.Decimal128 `bson:"gasUsed"`
	Hash            string          `bson:"hash"`
	Bloom           string          `bson:"logsBloom"`
	Coinbase        string          `bson:"miner"`
	MixDigest       string          `bson:"mixHash"`
	Nonce           string          `bson:"nonce"`
	Number          bson.Decimal128 `bson:"number"`
	ParentHash      string          `bson:"parentHash"`
	ReceiptHash     string          `bson:"receiptsRoot"`
	UncleHash       string          `bson:"sha3Uncles"`
	Size            bson.Decimal128 `bson:"size"`
	Root            string          `bson:"stateRoot"`
	Time            bson.Decimal128 `bson:"timestamp"`
	TotalDifficulty bson.Decimal128 `bson:"totalDifficulty"`
	TXs             []MTransaction  `bson:"transactions"`
	TxHash          string          `bson:"transactionsRoot"`
	Uncles          []string        `bson:"uncles"`
}

// MTransaction represents a transaction that will serialize to the RPC representation of a transaction
type MTransaction struct {
	BlockHash        string          `bson:"blockHash"`
	BlockNumber      bson.Decimal128 `bson:"blockNumber"`
	From             string          `bson:"from"`
	Gas              bson.Decimal128 `bson:"gas"`
	GasPrice         bson.Decimal128 `bson:"gasPrice"`
	Hash             string          `bson:"hash"`
	Input            string          `bson:"input"`
	Nonce            string          `bson:"nonce"`
	To               string          `bson:"to"`
	TransactionIndex bson.Decimal128 `bson:"transactionIndex"`
	Value            bson.Decimal128 `bson:"value"`
	V                string          `bson:"v"`
	R                string          `bson:"r"`
	S                string          `bson:"s"`
}

// Block represents a block header in the Ethereum blockchain.
type Block struct {
	Difficulty      string        `json:"difficulty"`
	Extra           string        `json:"extraData"`
	GasLimit        string        `json:"gasLimit"`
	GasUsed         string        `json:"gasUsed"`
	Hash            string        `json:"hash"`
	Bloom           string        `json:"logsBloom"`
	Coinbase        string        `json:"miner"`
	MixDigest       string        `json:"mixHash"`
	Nonce           string        `json:"nonce"`
	Number          string        `json:"number"`
	ParentHash      string        `json:"parentHash"`
	ReceiptHash     string        `json:"receiptsRoot"`
	UncleHash       string        `json:"sha3Uncles"`
	Size            string        `json:"size"`
	Root            string        `json:"stateRoot"`
	Time            string        `json:"timestamp"`
	TotalDifficulty string        `json:"totalDifficulty"`
	TXs             []Transaction `json:"transactions"`
	TxHash          string        `json:"transactionsRoot"`
	Uncles          []string      `json:"uncles"`
}

// Transaction represents a transaction that will serialize to the RPC representation of a transaction
type Transaction struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}

func hexToDecimal(s string) bson.Decimal128 {
	bigInt := new(big.Int)
	bigInt.SetString(s, 0)
	bigIntByte, _ := bigInt.MarshalText()
	decimal, _ := bson.ParseDecimal128(string(bigIntByte))
	return decimal
	// return bigInt.Int64()
}

func converToMTransaction(txs []Transaction) []MTransaction {
	var mts []MTransaction

	for _, t := range txs {
		var mt MTransaction

		mt.BlockHash = t.BlockHash
		mt.BlockNumber = hexToDecimal(t.BlockNumber)
		mt.From = t.From
		mt.Gas = hexToDecimal(t.Gas)
		mt.GasPrice = hexToDecimal(t.GasPrice)
		mt.Hash = t.Hash
		mt.Input = t.Input
		mt.Nonce = t.Nonce
		mt.R = t.R
		mt.S = t.S
		mt.To = t.To
		mt.TransactionIndex = hexToDecimal(t.TransactionIndex)
		mt.V = t.V
		mt.Value = hexToDecimal(t.Value)

		mts = append(mts, mt)
	}

	return mts
}

// ToMBlock 转为为mongodb需要的bson格式
func (r *Block) ToMBlock() *MBlock {
	var m = MBlock{}
	// var err error
	m.Bloom = r.Bloom
	m.Coinbase = r.Coinbase
	m.Difficulty = hexToDecimal(r.Difficulty)
	m.Extra = r.Extra
	m.GasLimit = hexToDecimal(r.GasLimit)
	m.GasUsed = hexToDecimal(r.GasUsed)
	m.Hash = r.Hash
	m.MixDigest = r.MixDigest
	m.Nonce = r.Nonce
	m.Number = hexToDecimal(r.Number)
	m.ParentHash = r.ParentHash
	m.ReceiptHash = r.ReceiptHash
	m.Root = r.Root
	m.Size = hexToDecimal(r.Size)
	m.Time = hexToDecimal(r.Time)
	m.TotalDifficulty = hexToDecimal(r.TotalDifficulty)
	m.TxHash = r.TxHash
	m.TXs = converToMTransaction(r.TXs)
	m.UncleHash = r.UncleHash
	m.Uncles = r.Uncles

	return &m
}

func intern(n uint64) []struct{} {
	return make([]struct{}, n)
}
