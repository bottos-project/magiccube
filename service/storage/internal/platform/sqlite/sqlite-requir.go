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

package sqlite

import (
	"github.com/bottos-project/magiccube/service/storage/util"
)

// CallGetUserRequirementList from db
func (r *SqliteRepository) CallGetUserRequirementList(username string) ([]*util.RequirementDBInfo, error) {
	var reqs = []*util.RequirementDBInfo{}
	dbtag := new(util.RequirementDBInfo)
	dbtag.RequirementId = "idtest"
	dbtag.RequirementName = "requirename"
	dbtag.FeatureTag = 111
	dbtag.SamplePath = "pathtest"
	dbtag.SampleHash = "hashtest"
	dbtag.ExpireTime = 222
	dbtag.Price = 333
	dbtag.Description = "destest"
	dbtag.PublishDate = 444
	reqs = append(reqs, dbtag)
	return reqs, nil
}
