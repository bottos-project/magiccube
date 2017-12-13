package com.bottos.bottosapp.bean;

import com.alibaba.fastjson.JSON;
import com.bottos.bottosapp.common.CommonConst;

/**
 * Bottos
 */
public class ResponseMsg {

    public int getRetCode() {
        return retCode;
    }

    public void setRetCode(int retCode) {
        this.retCode = retCode;
    }

    public String getRetResult() {
        return retResult;
    }

    public void setRetResult(String retResult) {
        this.retResult = retResult;
    }

    public ResponseMsg() {
        this.retCode = CommonConst.FAILED;
        this.retResult = "";
    }

    public ResponseMsg(int code, String result) {
        this.retCode = code;
        this.retResult = result;
    }

    private int retCode = CommonConst.FAILED;
    private String retResult;

    public String retToJson(int code, String result) {

        ResponseMsg responseMsg = new ResponseMsg(code, result);
        return JSON.toJSONString(responseMsg);
    }


}
