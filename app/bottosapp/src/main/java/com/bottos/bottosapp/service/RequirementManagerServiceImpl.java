package com.bottos.bottosapp.service;


import com.alibaba.fastjson.JSON;
import com.bottos.bottosapp.bean.RequirementManagerBean;
import com.bottos.bottosapp.common.ConfigSettings;
import com.bottos.bottosapp.common.TypeMapping;
import com.bottos.bottosapp.common.Web3jUtils;
import com.bottos.bottosapp.contract.RequirementManager;
import com.bottos.bottosapp.common.Web3Manager;
import com.bottos.bottosapp.mapper.UserRespository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;
import org.web3j.crypto.Credentials;
import org.web3j.crypto.WalletUtils;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.Web3jService;
import org.web3j.protocol.core.DefaultBlockParameterName;
import org.web3j.protocol.core.methods.response.EthGetTransactionCount;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.protocol.core.methods.response.Web3ClientVersion;
import org.web3j.protocol.http.HttpService;

import java.io.File;
import java.io.IOException;
import java.math.BigInteger;
import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.CompletableFuture;

import static com.bottos.bottosapp.common.TypeMapping.convertDisplayString;


@Repository

public class RequirementManagerServiceImpl {
    @Autowired
    private UserRespository userRespository;

    @Autowired
    ConfigSettings configSettings;

    protected Logger logger = LoggerFactory.getLogger(this.getClass().getName());


    // publish requirement
    public String publishRequirement(RequirementManagerBean requirementManagerBean) {
        try {

            Map<String, Object> requirementCreateParaMap = new HashMap<String, Object>();
            requirementCreateParaMap.put("dataRequirementName", requirementManagerBean.getRequirementName());
            requirementCreateParaMap.put("expirationTime", requirementManagerBean.getExpirationTime());
            requirementCreateParaMap.put("publishedTime", System.currentTimeMillis());
            requirementCreateParaMap.put("bidMoney", requirementManagerBean.getBidMoney());
            requirementCreateParaMap.put("requirementType", TypeMapping.getTypeValue(requirementManagerBean.getRequirementType()));
            requirementCreateParaMap.put("applicationDomain", TypeMapping.getTypeValue(requirementManagerBean.getApplicationDomain()));
            requirementCreateParaMap.put("description", requirementManagerBean.getDescription());
            requirementCreateParaMap.put("dataType", TypeMapping.getTypeValue(requirementManagerBean.getDataType()));
            requirementCreateParaMap.put("specifications", requirementManagerBean.getSpecifications());
            requirementCreateParaMap.put("DataSampleRef", requirementManagerBean.getDataSampleRef());
            requirementCreateParaMap.put("dataSample1", requirementManagerBean.getDataSample1());
            requirementCreateParaMap.put("dataSample2", "cj2");
            requirementCreateParaMap.put("dataSample3", "cj3");
//            requirementCreateParaMap.put("dataSample2", requirementManagerBean.getDataSample2());
//            requirementCreateParaMap.put("dataSample3", requirementManagerBean.getDataSample3());
            requirementCreateParaMap.put("nonce", getNonce(requirementManagerBean.getActionAccount()).toString());
            requirementCreateParaMap.put("requirementSignature", String.valueOf(System.currentTimeMillis()));// TODO !!!!!!!!

            String createRequirementParaJson = JSON.toJSONString(requirementCreateParaMap);

            RequirementManager requirementManagerContract = getContract(configSettings, requirementManagerBean.getActionAccount());

            TransactionReceipt receipt = requirementManagerContract.createDataRequirement(createRequirementParaJson).send();

            // TODO: need check whether the requirement is published succeeded
            logger.info("Publish Requirement OK.");
            return "";

        } catch (Exception e) {

            e.printStackTrace();
            return e.getMessage();
        }
    }

    // query my requirement
    public String queryMyRequirement(RequirementManagerBean requirementManagerBean) {
        try {

            String requirementJson = "";
            RequirementManager requirementManagerContract = getContract(configSettings, requirementManagerBean.getActionAccount());

            if (requirementManagerBean.getOwnerAccount().equals(requirementManagerBean.getActionAccount())) {

                //query all my requirement
                requirementJson = requirementManagerContract.queryDataRequirementbyOwner(requirementManagerBean.getOwnerAccount()).send();
            }
            logger.info("query My Requirement End.");
            return convertDisplayString(requirementJson);

        } catch (Exception e) {

            e.printStackTrace();
            return "";
        }
    }


}
