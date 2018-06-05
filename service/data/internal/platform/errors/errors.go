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

package errors

// Error is a handy type alias that lets us create constant error messages
type Error string

// Error is the "stringer" function, which completes our error interface implementation.
func (e Error) Error() string {
	return string(e)
}

const (
	// BadSearchTerm indicates a search term validation failure
	BadSearchTerm = Error("Bad search term")
	// NoSuchCategory indicates a request for a category that doesn't exist.
	NoSuchCategory = Error("No such category")
	// NoSuchProduct indicates a request for a non-existent product
	NoSuchProduct = Error("No such product")
)
