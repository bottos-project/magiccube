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
type Info struct {
	ServerVersion            string `json:"server_version"`
	HeadBlockNum             uint64    `json:"head_block_num"`
	LastIrreversibleBlockNum uint64    `json:"last_irreversible_block_num"`
	HeadBlockID              string `json:"head_block_id"`
	HeadBlockTime            string `json:"head_block_time"`
	HeadBlockProducer        string `json:"head_block_producer"`
	RecentSlots              string `json:"recent_slots"`
	ParticipationRate        string `json:"participation_rate"`
}
type Block struct {
	Previous              string        `json:"previous"`
	Timestamp             string        `json:"timestamp"`
	TransactionMerkleRoot string        `json:"transaction_merkle_root"`
	Producer              string        `json:"producer"`
	ProducerChanges       []interface{} `json:"producer_changes"`
	ProducerSignature     string        `json:"producer_signature"`
	Cycles                []interface{} `json:"cycles"`
	ID                    string        `json:"id"`
	BlockNum              uint64           `json:"block_num"`
	RefBlockPrefix        uint64           `json:"ref_block_prefix"`
}