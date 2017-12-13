package com.bottos.bottosapp.controller;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.bottos.bottosapp.bean.AccountInfoBean;
import com.bottos.bottosapp.bean.RequirementManagerBean;
import com.bottos.bottosapp.bean.ResponseMsg;
import com.bottos.bottosapp.common.CommonConst;
import com.bottos.bottosapp.service.DataAssetServiceImpl;
import com.bottos.bottosapp.service.ExchangeServiceImpl;
import com.bottos.bottosapp.service.RequirementManagerServiceImpl;
import com.bottos.bottosapp.service.WalletManagerServiceImpl;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;

/**
 * Bottos.
 */
@Controller
@RequestMapping(value = "/wallet")
public class WalletManagerController {
    protected Logger logger = LoggerFactory.getLogger(this.getClass().getName());

    @Resource
    WalletManagerServiceImpl walletManagerService;

    @Resource
    DataAssetServiceImpl dataAssetManagerService;

    @Resource
    ExchangeServiceImpl exchangeManagerService;

    ResponseMsg responseMsg = new ResponseMsg();

  

    /**
     * @param json
     * @return Query My Requirement
     */
    @ResponseBody
    @RequestMapping(value = "/queryAccountInfo", method = RequestMethod.POST)
    public String queryAccountInfo(@RequestBody String json) {

        try {
            logger.info("Start query my Wallet Info. param= " + json);

            JSONObject jsonArray = JSON.parseObject(json);
            String result = walletManagerService.queryWalletInfo(jsonArray.getString("name"));

            if (!result.isEmpty()) {
                return responseMsg.retToJson(CommonConst.SUCCESS, result);
            } else {
                return responseMsg.retToJson(CommonConst.FAILED, "");
            }

        } catch (Exception e) {

            e.printStackTrace();
            return responseMsg.retToJson(CommonConst.FAILED, "");
        }
    }

    /**
     * @param json
     * @return Query My Requirement
     */
    @ResponseBody
    @RequestMapping(value = "/transferToken", method = RequestMethod.POST)
    public String transferToken(@RequestBody String json) {

        try {
            logger.info("Start transfer Token . ", "param = ", json);

            AccountInfoBean accountInfoBean = JSON.parseObject(json, AccountInfoBean.class);

            if (accountInfoBean.getAccount().isEmpty()) {
                return responseMsg.retToJson(CommonConst.FAILED, "");
            }

            String result = walletManagerService.transferToken(accountInfoBean);


            if (result.isEmpty()) {
                return responseMsg.retToJson(CommonConst.SUCCESS, result);
            } else {
                return responseMsg.retToJson(CommonConst.FAILED, "");
            }

        } catch (Exception e) {

            e.printStackTrace();
            return responseMsg.retToJson(CommonConst.FAILED, "");
        }
    }


}
