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

//AccountInfo struct
type AccountInfo struct {
	AccountName       string `json:"account_name"`
	BtoBalance        string `json:"bto_balance"`
	StakedBalance     string `json:"staked_balance"`
	UnstakingBalance  string `json:"unstaking_balance"`
	LastUnstakingTime string `json:"last_unstaking_time"`
	Permissions       []struct {
		PermName     string `json:"perm_name"`
		Parent       string `json:"parent"`
		RequiredAuth struct {
			Threshold int `json:"threshold"`
			Keys      []struct {
				Key    string `json:"key"`
				Weight int    `json:"weight"`
			} `json:"keys"`
			Accounts []interface{} `json:"accounts"`
		} `json:"required_auth"`
	} `json:"permissions"`
}
