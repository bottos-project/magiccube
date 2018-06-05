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
 
package mongodb

import (
	"fmt"
	"testing"
)

func TestMongodbAsset_CallGetDayAssetNum(t *testing.T) {
	ins := MongoRepository{"47.98.47.148:27017"}

	code, err := ins.CallGetDayAssetNum(6)
	fmt.Println(code)
	fmt.Println(err)
	fmt.Println("code")
}
