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

type FileDBInfo struct {
	FileHash          string `json:"file_hash"`
	Username          string `json:"username"`
	FileName          string `json:"file_name"`
	FileSize          uint64 `json:"file_size"`
	FileNumber        uint64 `json:"file_number"`
	FilePolicy        string `json:"file_policy"`
	AuthorizedStorage string `json:"authorized_storage"`
}

const InsertUserFileSql string = "insert into fileinfo(FileHash,Username,FileName,FileSize,FileNumber,FilePolicy,AuthorizedStorage) values(?,?,?,?,?,?,?)"
