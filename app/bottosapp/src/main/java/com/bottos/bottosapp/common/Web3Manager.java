package com.bottos.bottosapp.common;

import com.bottos.bottosapp.contract.UserManager;
import com.bottos.bottosapp.mapper.UserRespository;
import org.apache.http.impl.client.CloseableHttpClient;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.web3j.crypto.Credentials;
import org.web3j.crypto.WalletUtils;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.Web3jService;
import org.web3j.protocol.core.methods.response.Web3ClientVersion;
import org.web3j.protocol.http.HttpService;
import org.web3j.tx.RawTransactionManager;
import org.web3j.tx.TransactionManager;

import java.io.File;
import java.io.IOException;
import java.math.BigInteger;
import java.util.HashMap;
import java.util.Map;
import java.util.Set;

/**
 * create by wcj
 */
//@Component
//@Repository
@Service
public class Web3Manager {

    @Autowired
    public ConfigSettings configSettings;
    @Autowired
    private UserRespository userRespository;

//    @Component
//    @ConfigurationProperties(locations = "classpath:application.properties",prefix="test")


    protected Logger logger = LoggerFactory.getLogger(this.getClass().getName());

    private Set<Class<?>> classSet;
    private String gasPrice;
    private String gasLimit;
    private String initialEtherValue;
    private String httpUrl;
    private String registerAddr;
    private String useChain;
    private Credentials credentials;
    private CloseableHttpClient closeableHttpClient;
    private Integer sleepDuration;
    private Integer attempts;
    private String maxTotal;
    private String maxPerRoute;
    private RawTransactionManager transactionManager;
//    private String walletPath;

    private UserManager userManager;

    public Web3Manager() {
    }

    public void doNothing() {
    }


    /**
     * @return connect to BlockChain with web3j service
     */

    public Web3j initWeb3jService(ConfigSettings configSettings) {
        try {
            this.logger.info("start init Web3jServiced");

            HttpClient.maxTotal = configSettings.getMaxTotal();
            HttpClient.maxPerRoute = configSettings.getMaxPerRoute();
            Web3jService httpService = new HttpService(configSettings.getHttpUrl());
            Web3j web3j = Web3j.build(httpService);
            Web3ClientVersion web3ClientVersion = web3j.web3ClientVersion().sendAsync().get();
            String clientVersion = web3ClientVersion.getWeb3ClientVersion();

            return web3j;
        } catch (Exception e) {
            e.printStackTrace();
            logger.error(e.getMessage());
        }
        return null;
    }

    public Map<String, Object> createContractInstance(String contractAddr, String accountName) {
        return null;
    }

    /**
     * @param configSettings
     * @param accountName
     * @return call contract
     */
    public Map<String, Object> createContractInstance(ConfigSettings configSettings, String accountName, String accountPasswd) throws IOException {
        try {
            this.logger.info("start createContractInstance");

            Web3j web3j = initWeb3jService(configSettings);

            String accountFilePath = configSettings.getWalletPath() + File.separator + accountName + ".json";
//            String json = "{\"address\":\"fe67c5731484b044de64a620db511dbdd44201e8\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"21e4df4597a69a95914ab8879c5b4bc997c2fef6e2316fe4bea7a46a78441247\",\"cipherparams\":{\"iv\":\"615bac7f03f404adfed16774a30c9300\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"60a277700f84175f6357d83a6d50d3e504c9e2ecffd872a2a77199bb2f9d2dfb\"},\"mac\":\"659bdbaac9d0dc21a73ffa90c8308292a72132a048e95627a79923ec61276793\"},\"id\":\"063c3411-b4f9-4d9c-bd1a-0fa963775e64\",\"version\":3}";
            String credentialStr = "";
//            String credentialStr = userRespository.findByAccount(accountName).getCredentials();
            this.credentials = WalletUtils.loadCredentials(accountPasswd, credentialStr);
//            this.credentials = WalletUtils.loadCredentials(accountPasswd, accountFilePath);
//            logger.info("credentials: " + this.credentials.toString());

            BigInteger gasPriceBig = new BigInteger(configSettings.getGasPrice());
            BigInteger gasLimitBig = new BigInteger(configSettings.getGasLimit());

            Map<String, Object> contractParaMap = new HashMap<String, Object>();
            contractParaMap.put("contractAddr", configSettings.getUserManagerContractAddr());
            contractParaMap.put("web3j", web3j);
            contractParaMap.put("credentials", credentials);
            contractParaMap.put("gasPrice", gasPriceBig);
            contractParaMap.put("gasLimit", gasLimitBig);

            return contractParaMap;
        } catch (Exception e) {
            logger.error(e.getMessage());
        }
        return null;
    }


    /**
     * @param accountPasswd
     * @param accountName
     * @param c
     * @return call contract
     */
    public Object initContract(String accountPasswd, String accountName, Class c) {
        try {
            this.logger.info("initContract");
            BigInteger gasPriceBig = new BigInteger(this.gasPrice);
            BigInteger gasLimitBig = new BigInteger(this.gasLimit);
            HttpClient.maxTotal = this.maxPerRoute;
            HttpClient.maxPerRoute = this.maxPerRoute;
            Web3jService httpService = new HttpService(this.httpUrl);
            Web3j web3j = Web3j.build(httpService);
            Web3ClientVersion web3ClientVersion = web3j.web3ClientVersion().send();
            String clientVersion = web3ClientVersion.getWeb3ClientVersion();

            String accountFileName = accountName + ".json";
            this.credentials = null;
//            this.credentials = WalletUtils.loadCredentials(accountPasswd, this.walletPath + File.separator + accountFileName);

            userManager = new UserManager(this.registerAddr, web3j, this.credentials, gasPriceBig, gasLimitBig);

            return userManager;

        } catch (Exception e) {
            logger.error(e.getMessage());
        }
        return null;
    }


    public Set<Class<?>> getClassSet() {
        return this.classSet;
    }

    public void setClassSet(Set<Class<?>> classSet) {
        this.classSet = classSet;
    }

    public void setGasPrice(String gasPrice) {
        this.gasPrice = gasPrice;
    }

    public void setGasLimit(String gasLimit) {
        this.gasLimit = gasLimit;
    }

    public void setInitialEtherValue(String initialEtherValue) {
        this.initialEtherValue = initialEtherValue;
    }

    public void setHttpUrl(String httpUrl) {
        this.httpUrl = httpUrl;
    }

    public String getHttpUrl() {
        return httpUrl;
    }


    public Credentials getCredentials() {
        return this.credentials;
    }

    public void setCredentials(Credentials credentials) {
        this.credentials = credentials;
    }

    public void setRegisterAddr(String registerAddr) {
        this.registerAddr = registerAddr;
    }

    public void setUseChain(String useChain) {
        this.useChain = useChain;
    }

    public Integer getAttempts() {
        return this.attempts;
    }

    public void setAttempts(Integer attempts) {
        this.attempts = attempts;
    }

    public Integer getSleepDuration() {
        return this.sleepDuration;
    }

    public void setSleepDuration(Integer sleepDuration) {
        this.sleepDuration = sleepDuration;
    }

    public CloseableHttpClient getCloseableHttpClient() {
        return this.closeableHttpClient;
    }

    public void setCloseableHttpClient(CloseableHttpClient closeableHttpClient) {
        this.closeableHttpClient = closeableHttpClient;
    }


    public String getMaxTotal() {
        return this.maxTotal;
    }

    public void setMaxTotal(String maxTotal) {
        this.maxTotal = maxTotal;
    }

    public String getMaxPerRoute() {
        return this.maxPerRoute;
    }

    public void setMaxPerRoute(String maxPerRoute) {
        this.maxPerRoute = maxPerRoute;
    }

    public String getGasPrice() {
        return this.gasPrice;
    }

    public String getGasLimit() {
        return this.gasLimit;
    }

    public TransactionManager getTransactionManager() {
        return transactionManager;
    }


}
