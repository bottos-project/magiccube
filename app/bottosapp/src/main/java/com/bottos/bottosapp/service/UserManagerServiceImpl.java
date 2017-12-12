package com.bottos.bottosapp.service;


import com.alibaba.fastjson.JSON;
import com.bottos.bottosapp.bean.UserInfo;
import com.bottos.bottosapp.common.ConfigSettings;
import com.bottos.bottosapp.common.Web3Manager;
import com.bottos.bottosapp.common.Web3jUtils;
import com.bottos.bottosapp.common.exception.EngineExceptionHelper;
import com.bottos.bottosapp.common.utils.UserExcepFactor;
import com.bottos.bottosapp.contract.TokenManager;
import com.bottos.bottosapp.contract.UserManager;
import com.bottos.bottosapp.dao.BaseDaoImpl;
import com.bottos.bottosapp.mapper.UserRespository;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Configuration;
import org.springframework.scheduling.annotation.Async;
import org.springframework.scheduling.annotation.EnableAsync;
import org.springframework.stereotype.Repository;
import org.springframework.stereotype.Service;
import org.springframework.util.StringUtils;
import org.web3j.abi.datatypes.Utf8String;
import org.web3j.crypto.*;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.Web3jService;
import org.web3j.protocol.core.DefaultBlockParameterName;
import org.web3j.protocol.core.methods.response.EthGetTransactionCount;
import org.web3j.protocol.core.methods.response.EthSendTransaction;
import org.web3j.protocol.core.methods.response.TransactionReceipt;
import org.web3j.protocol.core.methods.response.Web3ClientVersion;
import org.web3j.protocol.exceptions.TransactionException;
import org.web3j.protocol.http.HttpService;
import org.web3j.tx.Transfer;
import org.web3j.utils.Convert;
import org.web3j.utils.Numeric;

import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.File;
import java.io.IOException;
import java.math.BigDecimal;
import java.math.BigInteger;
import java.util.Map;
import java.util.UUID;
import java.util.concurrent.ExecutionException;

import static org.web3j.tx.Contract.GAS_LIMIT;
import static org.web3j.tx.ManagedTransaction.GAS_PRICE;

@Repository
public class UserManagerServiceImpl extends BaseDaoImpl<UserManager, Object> {
    @Autowired
    private UserRespository userRespository;
    @Autowired
    ConfigSettings configSettings;

//    @Autowired
//    ConfigSettings configSettings;

    static Integer storeId = 0;
    protected Logger logger = LoggerFactory.getLogger(this.getClass().getName());
//    protected Jedis jedis = new Jedis("172.20.51.141", 6379);

    /**
     * @param configSettings read the parament info
     * @param accountName
     * @return
     * @throws IOException
     */
    private UserManager getContract(ConfigSettings configSettings, String accountName) throws IOException {
        Credentials credentials = null;
        BigInteger gasPriceBig = null;
        BigInteger gasLimitBig = null;
//        Web3j web3j = null;

        String accountPasswd = userRespository.findByAccount(accountName).getPasswd();

        try {
            this.logger.info("start create manager.");
            long startTime1 = System.currentTimeMillis(); //获取开始时间
            Web3j web3j = getWeb3jClient();

            Web3jService httpService = new HttpService(configSettings.getHttpUrl());
            web3j = Web3j.build(httpService);
            Web3ClientVersion web3ClientVersion = web3j.web3ClientVersion().send();
            String clientVersion = web3ClientVersion.getWeb3ClientVersion();

            String accountFilePath = configSettings.getWalletPath() + File.separator + accountName + ".json";
//            String credentialStr = "{\"address\":\"9809eeb4cecc810901b7664a28146f7e4edc5667\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"6fc80581ebf2d64d01b9a7362abc3bbc69fbc75ee1d769724cf49149f08c7808\",\"cipherparams\":{\"iv\":\"7a1d476f34c9e11c699c09ca3a6565d8\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"25a565f3f91a16ff0391bded95eb8493de56f329272a51cec462602fc96a0232\"},\"mac\":\"965a085ab0f3b51c66f5849dd76d5067145501569c36aa06155ccee4c1dd34a5\"},\"id\":\"e3696124-395b-482a-9e6d-34207373c961\",\"version\":3}";
//            String credentialStr="d:/UTC--2017-11-20T09-14-00.597140029Z--9809eeb4cecc810901b7664a28146f7e4edc5667";
            credentials = WalletUtils.loadCredentials(accountPasswd, accountFilePath);

            gasPriceBig = new BigInteger(configSettings.getGasPrice());
            gasLimitBig = new BigInteger(configSettings.getGasLimit());

        } catch (Exception e) {
            logger.error(e.getMessage());
        }


//        Web3Manager web3Manager = new Web3Manager();
//        Map<String, Object> map = web3Manager.createContractInstance(configSettings, accountName, accountPasswd);
//        String contractAddr = (String) map.get("contractAddr");
//        Web3j web3j = (Web3j) map.get("web3j");
//        Credentials credentials = (Credentials) map.get("credentials");
//        BigInteger gasPriceBig = (BigInteger) map.get("gasPrice");
//        BigInteger gasLimitBig = (BigInteger) map.get("gasLimit");

        UserManager manager = new UserManager(configSettings.getUserManagerContractAddr(), web3j, credentials, gasPriceBig, gasLimitBig);
//        logger.info("get userManager :" + manager);
        return manager;
    }

    public static Credentials loadCredentials(String password, String credentialStr)
            throws IOException, CipherException {
        WalletFile walletFile = JSON.parseObject(credentialStr, WalletFile.class);
//        WalletFile walletFile = objectMapper.readValue(source, WalletFile.class);
        return Credentials.create(Wallet.decrypt(password, walletFile));
    }

    private UserManager getContract(String contractAddr, String accountName) throws IOException {
//        Web3Manager web3Manager = (Web3Manager) SpringContextUtil.getBean("web3Manager");
        Web3Manager web3Manager = new Web3Manager();
        Map<String, Object> map = web3Manager.createContractInstance(contractAddr, accountName);
//        String contractAddr = (String) map.get("contractAddr");
//        Web3j web3j = (Web3j) map.get("web3j");
//        Credentials credentials = (Credentials) map.get("credentials");
//        BigInteger gasPriceBig = (BigInteger) map.get("gasPriceBig");
//        BigInteger gasLimitBig = (BigInteger) map.get("gasLimitBig");

//        UserManager userManager = web3Manager.createContractInstance(contractName, accountName);

        //            add by wcj
        Credentials credentials = null;
        try {
            credentials = WalletUtils.loadCredentials(
                    "1", "d:/0xfe67c5731484b044de64a620db511dbdd44201e8.json");
        } catch (IOException e) {
            e.printStackTrace();
        } catch (CipherException e) {
            e.printStackTrace();
        }
//            build web3j
//        Web3Manager web3Manager = (Web3Manager) SpringContextUtil.getBean("web3Manager");
//        Web3j web3j = web3Manager.initWeb3jService();

        String desUrl = "http://10.104.11.249:9001";
        Web3j web3j = Web3j.build(new HttpService(desUrl));  // defaults to http://localhost:8545/
        Web3ClientVersion web3ClientVersion = web3j.web3ClientVersion().send();
        String clientVersion = web3ClientVersion.getWeb3ClientVersion();

        UserManager userManager = new UserManager("0xede092a3c2a3bb69ff3c7a1e2e3153d73f25a0f3", web3j, credentials, GAS_PRICE, GAS_LIMIT);

//            end test

//        UserManager userManager = web3Manager.createContractInstance(contractName, accountName);
        ;
//        userManager = new UserManager(contractAddr, web3j, credentials, gasPriceBig, gasLimitBig);
        return userManager;
//        return getContract4Unlogin("1", accountName,UserManager.class );
    }


    //TODO::支持单用户多账户模式
    public String addAccount(UserInfo userInfo) {
        return "";
    }

    /**
     * @param userInfo user info
     * @return login result
     */
    public String login(UserInfo userInfo) {
        try {
            String name = userInfo.getName();
            String passwd = userInfo.getPasswd();

            if (name == null || "".equals(name) || passwd == null || "".equals(passwd)) {
                return "The userName or password is invalid.";
            } else {
                UserInfo userTmp = userRespository.findByName(name);
                if (userTmp != null && passwd.equals(userTmp.getPasswd())) {
                    return "";
                } else {
                    logger.info("The userName or password is invalid.");
                    return "The userName or password is invalid.";
                }
            }
/*
            String walletfile = userInfo.getUserAddr() + ".json";
            System.out.println("wallet file:" + walletfile);
            String accountName = userInfo.getUserAddr();

//            ConfigSettings aa = new ConfigSettings();

            //调用链上UserManager合约
            UserManager userManager = getContract(configSettings.getUserManagerContractAddr(), accountName);
            if (userManager == null) {
                return "user wallet file error.";
            }
            Utf8String userAccount = new Utf8String(userInfo.getAccount());
            Utf8String user = null;
//            Utf8String user = userManager.findByAccount(userAccount).get();

            //Just fot 0726 test version
            switch (userInfo.getAccount()) {
                case "admin":
                    userInfo.setDepartmentId("org0001");
                    String[] roleList = {"userRole0001"};
//                    userInfo.setRoleIdList(roleList);
                    userInfo.setLoginStatus(1);
                    break;
                case "finance001":
                    userInfo.setDepartmentId("org0003");
                    String[] roleList1 = {"userRole0001"};
//                    userInfo.setRoleIdList(roleList1);
                    userInfo.setLoginStatus(1);
                    break;
                case "center001":
                    userInfo.setDepartmentId("org0002");
                    String[] roleList2 = {"userRole0001"};
//                    userInfo.setRoleIdList(roleList2);
                    userInfo.setLoginStatus(1);
                    break;
                case "supply001":
                    userInfo.setDepartmentId("org0004");
                    String[] roleList3 = {"userRole0001"};
//                    userInfo.setRoleIdList(roleList3);
                    userInfo.setLoginStatus(1);
                    break;
                default:
                    return "invalid user.";
            }
            //TODO::解析并记录链上用户的信息
            // UserChainMsg userChainMsg = JSON.parseObject(user.getValue().toString(), UserChainMsg.class);
            //UserInfo[] retUserInfo = userChainMsg.getData().getItems();
            //userInfo.setDepartmentId(retUserInfo[0].getDepartmentId());
            //userInfo.setRoleIdList(retUserInfo[0].getRoleIdList());
            //userInfo.setLoginStatus(1);

            //用户的信息记录到redis
            userInfo.setUuid(UUID.randomUUID().toString());
            System.out.println(userInfo.getUuid());

            String userJson = JSON.toJSONString(userInfo);
//            jedis.set(userInfo.getUuid(), userJson);
            return "";*/

        } catch (Exception e) {
            e.printStackTrace();
            return "login failed.";
        }
    }

    public String addUser(UserInfo creator, UserInfo newUser) {

        UserInfo existsUser = userRespository.findByName(newUser.getName());
        if (existsUser != null) {
            return "account exists";
//                throw EngineExceptionHelper.localException(UserExcepFactor.ACCOUNT_EXISTS);
        }
        String accountName = null;
        try {
            //创建用户的wallet文件
            Map walletInfoMap = Web3jUtils.proUserWallet1(newUser.getPasswd(), configSettings.getWalletPath());
//            String accountName = Web3jUtils.proUserWallet(newUser.getPasswd(), configSettings.getWalletPath());
//            System.out.println(accountName);
            WalletFile walletFile = (WalletFile) walletInfoMap.get("credentials");
            newUser.setCredentials(JSON.toJSONString(walletFile));
            accountName = "0x" + walletFile.getAddress();
            newUser.setAccount(accountName);
            newUser.setAccountAlias("Default Account");
            newUser.setTokenType("bobi");

            ECKeyPair ecKeyPair = (ECKeyPair) walletInfoMap.get("ecKeyPair");
            newUser.setPublicKey(ecKeyPair.getPublicKey().toString());
            newUser.setPraviteKey(ecKeyPair.getPrivateKey().toString());

            //      store user info in DB
            UserInfo storeResult = saveUser(newUser);
            storeId = storeResult.getId();

            logger.info("storeId :" + storeId);

            //调用userManager链上合约
//            String accountFilePath = configSettings.getWalletPath() + File.separator + accountName + ".json";
            UserManager userManager = getContract(configSettings, configSettings.getAdminAccount());

            if (userManager == null) {
                return "user wallet file error.";
            }

//            String tmpJson2 = "{\"name\":\"wcj21\", \"mobile\":\"13522222222\", \"email\":\"zw@test.com\", \"userTpye\":0, \"gender\":0, \"taxId\":\"123456\", \"companyAddr\":\"Grand Avenue No.1\"}";
            newUser.setPasswd("null");
            String tmpJson = JSON.toJSONString(newUser);
            Utf8String chainUserJson = new Utf8String(tmpJson);
//            logger.info("tmpJson: " + 0);
//         同步调用
            logger.info("Start register user on contract");
            long startTime = System.currentTimeMillis(); //获取开始时间
            TransactionReceipt receipt = (userManager.registerUser(tmpJson)).send();

            long endTime = System.currentTimeMillis(); //获取结束时间
            logger.info("get receipt " + "Time:" + (endTime - startTime) + "ms : ");
//            String user = userManager.ListAllUser().send();
//            logger.info("register user Ok.: " + user);

//            JSONObject jsonObject = JSON.parseObject(user);
//            String aa=jsonObject.getString("account");

//            异步调用
//            TransactionReceipt receiptAsync = (userManager.registerUser(tmpJson)).sendAsync().get();

//            TransactionReceipt receipt = (UserManager_sol_UserManager.registerUser(chainUserJson)).get();
//            userManager.insert(chainUserJson);
//            TransactionReceipt receipt = (userManager.registerUser(chainUserJson)).get();
//            TransactionReceipt receipt = (userManager.insert(chainUserJson)).get();
//            UserManager.NotifyEventResponse response = userManager.getNotifyEvents(receipt).get(0);
//            System.out.println( "errno = " + response._errno.getValue().intValue() + " info:" + response._info);

//            if(response._errno.getValue().intValue() != 0){
//                return  response._info.getValue().toString();
//            }
//            add bobi amount
//            add test
            final String accountName2 = accountName;

//            new Thread() {
//                public void run() {
//            addBobiAmount(accountName);
//                }
//            }.start();

            new Thread() {
                public void run() {
                    try {
                        addBobiAmount(accountName2);
                        String resTransfer = transfer(configSettings.getAdminAccountEth(), accountName2, configSettings.getDefaultAmountEth());
                    } catch (InterruptedException e) {
                        e.printStackTrace();
                    } catch (ExecutionException e) {
                        e.printStackTrace();
                    } catch (IOException e) {
                        e.printStackTrace();
                    } catch (CipherException e) {
                        e.printStackTrace();
                    } catch (TransactionException e) {
                        e.printStackTrace();
                    }
                }
            }.start();
            logger.info("eth nonce:" + getNonce(accountName).toString());
//            logger.info("eth nonce:"+ getNonce(accountName).toString());
//            new Thread(){
//                public void run(){
//                    addBobiAmount(accountName2);
//                }
//            }.start();
//            logger.info("bobi nonce:"+ getNonce(accountName).toString());

//            addBobiAmount(accountName);
//          add eth amount
//            String resTransfer = transfer1(configSettings.getAdminAccountEth(), accountName, configSettings.getDefaultAmountEth());
            String resTransfer = "";
            if (resTransfer.equals("")) {
                return "";
            } else {
                return resTransfer;
            }

        } catch (Exception e) {
            e.printStackTrace();
            return "register User failed.";
        }
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

    public String transfer(String fromAccount, String toAccount, double amount) throws InterruptedException, ExecutionException, IOException, CipherException, TransactionException {
        logger.info("Start transfer Eth. Amount is:" + amount);
        Web3j web3 = getWeb3jClient();
        String accountPasswd = userRespository.findByAccount(fromAccount).getPasswd();
        String accountFilePath = configSettings.getWalletPath() + System.getProperty("file.separator") + configSettings.getAdminAccountEth() + ".json";
        Credentials credentials = WalletUtils.loadCredentials(accountPasswd, accountFilePath);

        try {
            TransactionReceipt transactionReceipt = Transfer.sendFunds(
                    web3, credentials, toAccount, BigDecimal.valueOf(amount), Convert.Unit.ETHER)
                    .send();
            logger.info("Transfer Eth OK. Amount is: " + amount);
            return "";
           /* // get the next available nonce
            EthGetTransactionCount ethGetTransactionCount = web3j.ethGetTransactionCount(
                    configSettings.getAdminAccountEth(), DefaultBlockParameterName.LATEST).send();
            BigInteger nonce = ethGetTransactionCount.getTransactionCount();

// create our transaction
            RawTransaction rawTransaction  = RawTransaction.createEtherTransaction(
                    nonce, GAS_PRICE, GAS_LIMIT, "0x48553960b9e1AC72cEf473f3BB1d552fC2a0850e", new BigInteger(String.valueOf(1)));

// sign & send our transaction
            byte[] signedMessage = TransactionEncoder.signMessage(rawTransaction, credentials);
            String hexValue = Numeric.toHexString(signedMessage);
            EthSendTransaction ethSendTransaction = web3j.ethSendRawTransaction(hexValue).send();*/
        } catch (Exception e) {
            logger.info("Transfer Eth Fail. " + e.getMessage());
            return e.getMessage();
        }
/*
        EthGetTransactionCount ethGetTransactionCount = web3j.ethGetTransactionCount(
                from, DefaultBlockParameterName.LATEST).sendAsync().get();

        BigInteger nonce = ethGetTransactionCount.getTransactionCount();
        System.out.println(nonce);

        RawTransaction rawTransaction = RawTransaction.createEtherTransaction(
                nonce, Convert.toWei("22", Convert.Unit.MWEI).toBigInteger(), Convert.toWei("44", Convert.Unit.GWEI).toBigInteger(), to, amount);
        byte[] signedMessage = TransactionEncoder.signMessage(rawTransaction, credentials1);
        String hexValue = Numeric.toHexString(signedMessage);

        EthSendTransaction ethSendTransaction = web3j.ethSendRawTransaction(hexValue).sendAsync().get();
        String transactionHash = ethSendTransaction.getTransactionHash();*/
    }


    public UserInfo saveUser(UserInfo user) {
        return userRespository.save(user);
    }

    public void deleteOneUser() {
        userRespository.delete(storeId);
    }

    /**
     * @param account
     * @return call contract
     */
    public String getAccountPasswd(String account) {
        userRespository.findAll();
        return userRespository.findByAccount(account).getPasswd();
    }

/*    public UserInfo userFindOne(UserInfo user, String) {
        UserInfo user = new UserInfo();
        user.setName(name);
        user.setEmail(email);
        return userRespository.save(user);
    }*/

    private String addBobiAmount(String toAccout) {

        TokenManager tokenManagerContract = null;
        try {
            tokenManagerContract = getContractToken(configSettings, configSettings.getAdminAccount());
            TransactionReceipt receipt = tokenManagerContract.addTokentoAccount(toAccout, configSettings.getDefaultAmountBo())
                    .send();
//        add test
            BigInteger tokenValue = tokenManagerContract.getBalance(toAccout).send();
            logger.info("Add BObi OK. Amount is :" + tokenValue);
            return "Add Amount success.";
        } catch (IOException e) {
            e.printStackTrace();
        } catch (Exception e) {
            e.printStackTrace();
            return e.getMessage();
        }
        return "";


    }

    /**
     * @param configSettings read the param info
     * @param accountName
     * @return
     * @throws IOException
     */
    private TokenManager getContractToken(ConfigSettings configSettings, String accountName) throws IOException {
        Credentials credentials = null;
        BigInteger gasPriceBig = null;
        BigInteger gasLimitBig = null;
//        Web3j web3j = null;

        Web3j web3j = getWeb3jClient();
        String accountPasswd = userRespository.findByAccount(accountName).getPasswd();

        try {
            this.logger.info("start create TokenManager ContractInstance");

//            Web3jService httpService = new HttpService(configSettings.getHttpUrl());
//            web3j = Web3j.build(httpService);
//            Web3ClientVersion web3ClientVersion = web3j.web3ClientVersion().send();
//            String clientVersion = web3ClientVersion.getWeb3ClientVersion();


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

    /**
     * @param request
     * @param inputCode
     * @return
     */
    public boolean checkTcode(HttpServletRequest request, String inputCode) {
        logger.info("Start check verification Code.");
        String code = null;
//1:Get the cookie inside the verification code information
        Cookie[] cookies = request.getCookies();
        for (Cookie cookie : cookies) {
            if ("imagecode".equals(cookie.getName())) {
                code = cookie.getValue();
                break;
            }
        }
//1:Get session verification code information
// String code1 = (String) request.getSession().getAttribute("");
//2:Determine the verification code is correct
        if (!StringUtils.isEmpty(inputCode) && inputCode.equalsIgnoreCase(code)) {
            logger.info("verification code OK.");
            return true;
        }
        logger.info("verification code error. writeCode=" + inputCode + " validateCode:" + code);
        return false;
    }

    public BigInteger getNonce(String address) throws Exception {

//        Web3jService httpService = new HttpService(configSettings.getHttpUrl());
//        Web3j web3j = Web3j.build(httpService);

        Web3j web3j = getWeb3jClient();

//        Web3ClientVersion web3ClientVersion = web3j.web3ClientVersion().send();
//        String clientVersion = web3ClientVersion.getWeb3ClientVersion();

        EthGetTransactionCount ethGetTransactionCount = web3j.ethGetTransactionCount(address, DefaultBlockParameterName.LATEST).send();
        BigInteger nonce = ethGetTransactionCount.getTransactionCount();

        return nonce;
    }
}
