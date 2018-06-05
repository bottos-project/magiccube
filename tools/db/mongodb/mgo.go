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
package mgo

import (
	"github.com/bottos-project/magiccube/config"
	"gopkg.in/mgo.v2"
)

var mgoSession *mgo.Session

func Session() *mgo.Session {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(config.BASE_MONGODB_ADDR)
		if err != nil {
			panic(err)
		}
	}

	mgoSession.SetMode(mgo.Monotonic, true)
	mgoSession.SetPoolLimit(200)
	return mgoSession.Clone()
}
