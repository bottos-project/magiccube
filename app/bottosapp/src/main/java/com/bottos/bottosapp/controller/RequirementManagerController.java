package com.bottos.bottosapp.controller;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.bottos.bottosapp.bean.DataAssetBean;
import com.bottos.bottosapp.bean.DataAssetEntity;
import com.bottos.bottosapp.bean.RequirementManagerBean;
import com.bottos.bottosapp.bean.ResponseMsg;
import com.bottos.bottosapp.common.CommonConst;
import com.bottos.bottosapp.mapper.DataRespository;
import com.bottos.bottosapp.service.DataAssetServiceImpl;
import com.bottos.bottosapp.service.ExchangeServiceImpl;
import com.bottos.bottosapp.service.RequirementManagerServiceImpl;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;
import javax.websocket.server.PathParam;

/**
 * Bottos.
 */
@Controller
@RequestMapping(value = "/requirement")
public class RequirementManagerController {
    protected Logger logger = LoggerFactory.getLogger(this.getClass().getName());
    String errorInfo = "Transaction receipt was not generated";
    @Resource
    RequirementManagerServiceImpl requirementManagerService;

    @Resource
    DataAssetServiceImpl dataAssetManagerService;

    @Resource
    ExchangeServiceImpl exchangeManagerService;
    @Autowired
    private DataRespository dataRespository;

    ResponseMsg responseMsg = new ResponseMsg();

    /**
     * @param json
     * @return
     */
    @ResponseBody
    @RequestMapping(value = "/publish", method = RequestMethod.POST)
    public String register(@RequestBody String json) {

        try {
            logger.info("Start publish requirement. param = " + json);

            RequirementManagerBean requirementManagerBean = JSON.parseObject(json, RequirementManagerBean.class);

            if (requirementManagerBean.getActionAccount().isEmpty()) {
                return responseMsg.retToJson(CommonConst.FAILED, "");
            }

            String result = requirementManagerService.publishRequirement(requirementManagerBean);

            if (result.isEmpty()) {
/*//                write dataAssetID to  dataAssetEntity
                JSONObject jsonArray = JSON.parseObject(result);
                String dataRequirementID = jsonArray.getString("dataRequirementID");
                DataAssetEntity dataAssetEntity = dataRespository.findByStoreID(requirementManagerBean.getDataSampleRef());
                dataAssetEntity.setRequireAssetID(dataRequirementID);
                dataRespository.save(dataAssetEntity);*/
                logger.info("publish Requirement Success." + result);
                return responseMsg.retToJson(CommonConst.SUCCESS, "Publish Requirement OK.");
            } else if (result.contains(errorInfo)) {
                return responseMsg.retToJson(CommonConst.OVERTIME, result);
            } else {
                return responseMsg.retToJson(CommonConst.FAILED, "");
            }

        } catch (Exception e) {

            e.printStackTrace();
            return responseMsg.retToJson(CommonConst.FAILED, e.getMessage());
        }
    }


    


}
