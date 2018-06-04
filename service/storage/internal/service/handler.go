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
package service

//import (
//	"encoding/json"
//	"net/http"
//	"time"

//	"github.com/micro/go-micro/client"

//	"golang.org/x/net/context"
//)

//type StorageHandler interface {
//	GetTx(context.Context, *DetailRequest, *DetailResponse) error
//}

//func StorageCall(w http.ResponseWriter, r *http.Request) {
//	// decode the incoming request as json
//	var request map[string]interface{}
//	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}

//	// call the backend service
//	storageClient := service.NewStorageClient("bottos.storage", client.DefaultClient)
//	rsp, err := storageClient.Call(context.TODO(), &storage.Request{
//		Name: request["name"].(string),
//	})
//	if err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}

//	// we want to augment the response
//	response := map[string]interface{}{
//		"msg": rsp.Msg,
//		"ref": time.Now().UnixNano(),
//	}

//	// encode and write the response as json
//	if err := json.NewEncoder(w).Encode(response); err != nil {
//		http.Error(w, err.Error(), 500)
//		return
//	}
//}
