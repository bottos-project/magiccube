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

            CompletableFuture<TransactionReceipt> receipt = requirementManagerContract.createDataRequirement(createRequirementParaJson).sendAsync();

            // TODO: need check whether the requirement is published succeeded
            logger.info("Publish Requirement OK.");
            return "";

        } catch (Exception e) {

            e.printStackTrace();
            return e.getMessage();
        }
    }

    

    /**
     * @param configSettings read the param info
     * @param accountName
     * @return
     * @throws IOException
     */
    private RequirementManager getContract(ConfigSettings configSettings, String accountName) throws IOException {
        Credentials credentials = null;
        BigInteger gasPriceBig = null;
        BigInteger gasLimitBig = null;
        Web3j web3j = null;

        String accountPasswd = userRespository.findByAccount(accountName).getPasswd();

        Web3Manager web3Manager = new Web3Manager();

        try {
            this.logger.info("start createContractInstance");

//            Web3jService httpService = new HttpService(configSettings.getHttpUrl());
//            web3j = Web3j.build(httpService);
//            Web3ClientVersion web3ClientVersion = web3j.web3ClientVersion().send();
//            String clientVersion = web3ClientVersion.getWeb3ClientVersion();
            web3j = getWeb3jClient();

            String accountFilePath = configSettings.getWalletPath() + File.separator + accountName + ".json";
            credentials = WalletUtils.loadCredentials(accountPasswd, accountFilePath);

            gasPriceBig = new BigInteger(configSettings.getGasPrice());
            gasLimitBig = new BigInteger(configSettings.getGasLimit());

        } catch (Exception e) {
            logger.error(e.getMessage());
        }

        RequirementManager manager = new RequirementManager(configSettings.getRequirementManagerContractAddr(), web3j, credentials, gasPriceBig, gasLimitBig);

        return manager;
    }

    public BigInteger getNonce(String address) throws Exception {

//        Web3jService httpService = new HttpService(configSettings.getHttpUrl());
//        Web3j web3j = Web3j.build(httpService);

        Web3j web3j = getWeb3jClient();

//        Web3ClientVersion web3ClientVersion = web3j.web3ClientVersion().send();
//        String clientVersion = web3ClientVersion.getWeb3ClientVersion();

        EthGetTransactionCount ethGetTransactionCount = web3j.ethGetTransactionCount(address, DefaultBlockParameterName.LATEST).send();

        return ethGetTransactionCount.getTransactionCount();
    }

    private volatile static Web3j web3j;

    public Web3j getWeb3jClient() {
        if (web3j == null) {
            synchronized (Web3jUtils.class) {
                if (web3j == null) {
                    web3j = Web3j.build(new HttpService(configSettings.getHttpUrl()));
                }
            }
        }
        return web3j;
    }

}
