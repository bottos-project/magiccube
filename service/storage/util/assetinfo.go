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

//AssetDBInfo struct
type AssetDBInfo struct {
	AssetID     string `bson:"asset_id" json:"asset_id"`
	UserName    string `bson:"user_name" json:"user_name"`
	AssetName   string `bson:"asset_name" json:"asset_name"`
	FeatureTag  uint64 `bson:"feature_tag" json:"feature_tag"`
	SamplePath  string `bson:"sample_path" json:"sample_path"`
	SampleHash  string `bson:"sample_hash" json:"sample_hash"`
	StoragePath string `bson:"storage_path" json:"storage_path"`
	StorageHash string `bson:"storage_hash" json:"storage_hash"`
	ExpireTime  uint32 `bson:"expire_time" json:"expire_time"`
	Price       uint64 `bson:"price" json:"price"`
	Description string `bson:"description" json:"description"`
	UploadDate  uint32 `bson:"upload_date" json:"upload_date"`
}
