// Copyright 2017~2022 The Bottos Authors
// This file is part of the Bottos Chain library.
// Created by Rocket Core Team of Bottos.

//This program is free software: you can distribute it and/or modify
//it under the terms of the GNU General Public License as published by
//the Free Software Foundation, either version 3 of the License, or
//(at your option) any later version.

//This program is distributed in the hope that it will be useful,
//but WITHOUT ANY WARRANTY; without even the implied warranty of
//MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//GNU General Public License for more details.

//You should have received a copy of the GNU General Public License
// along with bottos.  If not, see <http://www.gnu.org/licenses/>.

/*
 * file description:  provide a interface error definition for all modules
 * @Author: Charles
 * @Date:   2017-12-29
 * @Last Modified by:
 * @Last Modified time:
*/
package PackageError

import (
	_ "errors"
	_ "reflect"
	_ "strings"
	_ "encoding/json"
	  "map"
)

type errorCodeId unit64 /* error code : fatal,serious, tip, notification*/
type errorLevel unit8 /* error level: fatal,error,  notification*/
type errorMessage struct{
	  lanague unit8
      description string

}


type errorObj struct {

		errorCodeId
		errorLevel
		errorMessageMap[unit8] errorMessage
	    reason string

}


/*collection for one module's errors

{
  {
     "343890040505008"
     "error"
     {
     {
      "en"
      "this is an error,please check your input"
     }
     {
       "chi"
      "这是一个错误"
     }
     }

     "over length"
  }

 {
 
   "343893434545454545"
     "error"
     {
     {
      "en"
      "this is an notification,please check your input"
     }
     {
       "chi"
      "仅仅是一个提示"
     }
     }

     "input tooshort"
 
 
 }


}



*/
type eCollection map[errorCodeId]errorObj




/*init errors collection*/

func (*eCollection) initMapErrorObj(){



}

/* 生成全局errorCODE ID*/
func (errorCodeId) createErrorCodeId(serviceid UUID,innerErrorIndex unit16){

	return  serviceid&&0XFFFFFF00+innerErrorIndex


}



/* 更新错误描述*/
func insertErrorCode(code ErrorCode, imessage string,obj errorObj ){



}


func InsertUndefinedError(errCode, object errObj,uuid serviceID ){

}



/*根据error code查询error描述*/

 func CheckError(errorID unit64) string{



}




