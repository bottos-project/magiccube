package com.bottos.bottosapp.service;


import com.alibaba.fastjson.JSON;
import com.bottos.bottosapp.bean.AccountInfoBean;
import com.bottos.bottosapp.bean.RequirementManagerBean;
import com.bottos.bottosapp.bean.UserInfo;
import com.bottos.bottosapp.bean.WalletInfoBean;
import com.bottos.bottosapp.common.ConfigSettings;
import com.bottos.bottosapp.common.TypeMapping;
import com.bottos.bottosapp.common.Web3Manager;
import com.bottos.bottosapp.contract.RequirementManager;
import com.bottos.bottosapp.contract.TokenManager;
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
import org.web3j.protocol.core.RemoteCall;
import org.web3j.protocol.core.methods.response.EthGetTransactionCount;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.protocol.core.methods.response.Web3ClientVersion;
import org.web3j.protocol.http.HttpService;

import java.io.File;
import java.io.IOException;
import java.math.BigInteger;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

import static com.bottos.bottosapp.common.TypeMapping.convertDisplayString;


@Repository

public class WalletManagerServiceImpl {
    @Autowired
    private UserRespository userRespository;

    @Autowired
    ConfigSettings configSettings;

    protected Logger logger = LoggerFactory.getLogger(this.getClass().getName());


    /**
     * @param name
     * @return
     */
    public String queryWalletInfo(String name) {
        try {
            WalletInfoBean walletInfoBean = new WalletInfoBean();
            String requirementJson = "";
            List<Map<String, Object>> resultList = new ArrayList<>();

            UserInfo userInfo = userRespository.findByName(name);
            String account = userInfo.getAccount();
            TokenManager tokenManagerContract = getContract(configSettings, account);

//            if (requirementManagerBean.getOwnerAccount().equals(requirementManagerBean.getActionAccount())) {
            //query all my account Info
            BigInteger tokenValue = tokenManagerContract.getBalance(account).send();
//            }

            Map<String, Object> map = new HashMap<>();
            map.put("tokenType", userInfo.getTokenType());
            map.put("totalToken", tokenValue);
            resultList.add(map);
            walletInfoBean.setTotalTokenInfo(resultList);

            AccountInfoBean accountInfoBean = new AccountInfoBean();
            accountInfoBean.setAccountAlias(userInfo.getAccountAlias());
            accountInfoBean.setAccount(account);
            accountInfoBean.setTokenType(userInfo.getTokenType());
            accountInfoBean.setAccountToken(tokenValue);

            walletInfoBean.setAccountInfoList(accountInfoBean);

            logger.info("Query WalletInfo End");
            return JSON.toJSONString(walletInfoBean);
        } catch (Exception e) {

            e.printStackTrace();
            return "";
        }
    }

    /**
     * @param accountInfoBean
     * @return
     */
    public String transferToken(AccountInfoBean accountInfoBean) {
        try {
            String requirementJson = "";
            List<Map<String, Object>> resultList = new ArrayList<>();

//            UserInfo userInfo = userRespository.findByName(name);
//            String account = userInfo.getAccount();
            TokenManager tokenManagerContract = getContract(configSettings, accountInfoBean.getAccount());

//            if (requirementManagerBean.getOwnerAccount().equals(requirementManagerBean.getActionAccount())) {

            //query all my account Info
            TransactionReceipt receipt = tokenManagerContract.transfer(accountInfoBean.getAccount(), accountInfoBean.getReceiveAccount(), accountInfoBean.getAccountToken()).send();
//            }

            logger.info("Transfer Amount End");
            return "";

        } catch (Exception e) {

            e.printStackTrace();
            return "";
        }
    }

    /**
     * @param configSettings read the param info
     * @param accountName
     * @return
     * @throws IOException
     */
    private TokenManager getContract(ConfigSettings configSettings, String accountName) throws IOException {
        Credentials credentials = null;
        BigInteger gasPriceBig = null;
        BigInteger gasLimitBig = null;
        Web3j web3j = null;

        String accountPasswd = userRespository.findByAccount(accountName).getPasswd();

        Web3Manager web3Manager = new Web3Manager();

        try {
            this.logger.info("start createContractInstance");

            Web3jService httpService = new HttpService(configSettings.getHttpUrl());
            web3j = Web3j.build(httpService);
            Web3ClientVersion web3ClientVersion = web3j.web3ClientVersion().send();
            String clientVersion = web3ClientVersion.getWeb3ClientVersion();

            String accountFilePath = configSettings.getWalletPath() + File.separator + accountName + ".json";
            credentials = WalletUtils.loadCredentials(accountPasswd, accountFilePath);

            gasPriceBig = new BigInteger(configSettings.getGasPrice());
            gasLimitBig = new BigInteger(configSettings.getGasLimit());

        } catch (Exception e) {
            logger.error(e.getMessage());
        }

        TokenManager manager = new TokenManager(configSettings.getTokenManagerContractAddr(), web3j, credentials, gasPriceBig, gasLimitBig);

        return manager;
    }

    public BigInteger getNonce(String address) throws Exception {

        Web3jService httpService = new HttpService(configSettings.getHttpUrl());
        Web3j web3j = Web3j.build(httpService);

//        Web3ClientVersion web3ClientVersion = web3j.web3ClientVersion().send();
//        String clientVersion = web3ClientVersion.getWeb3ClientVersion();

        EthGetTransactionCount ethGetTransactionCount = web3j.ethGetTransactionCount(address, DefaultBlockParameterName.LATEST).send();

        return ethGetTransactionCount.getTransactionCount();
    }


}
