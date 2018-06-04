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
package util

type DataDBInfo struct {
	Filehash           string     `json:"filehash"`
	Username string     `json:"username"`
	Filename       string     `json:"filename"`
	Filesize       uint64     `json:"filesize"`
	Filepolicy       string     `json:"filepolicy"`
	Filenumber       uint64     `json:"filenumber"`
	Simorass     uint64     `json:"simorass"`
	Optype     uint64     `json:"optype"`
	Storeaddr         string `json:"storeaddr"`
}

const InsertDataSql string = "insert into datainfo(Guid,MerkleRootHash,Username,FileName,FileSize,FileNumber,FilePolicy,Fslice) values(?,?,?,?,?,?,?,?)"
