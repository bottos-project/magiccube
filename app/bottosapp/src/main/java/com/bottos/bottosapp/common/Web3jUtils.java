package com.bottos.bottosapp.common;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.bottos.bottosapp.mapper.UserRespository;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.web3j.crypto.*;
import org.web3j.protocol.ObjectMapperFactory;
import org.web3j.protocol.Web3j;
import org.web3j.protocol.http.HttpService;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.IOException;
import java.security.InvalidAlgorithmParameterException;
import java.security.NoSuchAlgorithmException;
import java.security.NoSuchProviderException;
import java.util.HashMap;
import java.util.Map;


public class Web3jUtils {
    @Autowired
    private UserRespository userRespository;
    @Autowired
    ConfigSettings configSettings;

    public static String proUserWallet(String password, String filepath) {
        BufferedReader reader = null;
        String tempString = null;
        try {
            Web3Manager web3Manager = new Web3Manager();
//            String filepath = CommonUtils.getWalletPath();
            File file = new File(filepath);

            String userId = Web3jUtils.generateNewWalletFile(password, file);
            String file1 = filepath + File.separator + userId;

            System.out.println("file1name =" + file1);
            reader = new BufferedReader(new FileReader(file1));
            // Reading into his party at one time until the end of the document read into null
            tempString = reader.readLine();
            reader.close();
            System.out.println(tempString);
            JSONObject jsonArray = JSON.parseObject(tempString);

            return ("0x" + jsonArray.getString("address"));
        } catch (Exception e) {
            e.printStackTrace();
            return "create wallet file error.";
        }
    }

    public static String generateNewWalletFile(String password, File destinationDirectory)
            throws CipherException, IOException, InvalidAlgorithmParameterException, NoSuchAlgorithmException, NoSuchProviderException {
        System.out.println("generateNewWalletFile :");
        ECKeyPair ecKeyPair = Keys.createEcKeyPair();
        System.out.println("generateNewWalletFile3");
        return generateWalletFile(password, ecKeyPair, destinationDirectory);
    }

    public static String generateWalletFile(String password, ECKeyPair ecKeyPair, File destinationDirectory)
            throws CipherException, IOException {
        System.out.println("generateWalletFile:");
        WalletFile walletFile = Wallet.createStandard(password, ecKeyPair);
        String fileName = getWalletFileName(walletFile);
        File destination = new File(destinationDirectory, fileName);

        ObjectMapper objectMapper = ObjectMapperFactory.getObjectMapper();
        objectMapper.writeValue(destination, walletFile);


        Map<String, Object> walletMap = new HashMap<String, Object>();
//        map.put("fileName", fileName);
        walletMap.put("walletFile", walletFile);
        walletMap.put("ecKeyPair", ecKeyPair);
//        JSONObject sas = (JSONObject) JSON.toJSON(walletFile);
//        String sa = JSON.toJSONString(walletFile);
//        JSONObject jsonArray = JSON.parseObject(sa);
//        JSONObject jsonObject = JSONObject.fromObject(walletFile);

        return fileName;
//        return walletMap;
    }


    public static Map<String, Object> proUserWallet1(String password, String filepath) {
        File file = new File(filepath);
        try {
            Map<String, Object> result = generateNewWalletFile1(password, file);
            return result;
        } catch (Exception e) {
            e.printStackTrace();
//            return "create wallet file error.";
        }
        return null;
    }


    public static Map<String, Object> generateNewWalletFile1(String password, File destinationDirectory)
            throws CipherException, IOException, InvalidAlgorithmParameterException, NoSuchAlgorithmException, NoSuchProviderException {
//        System.out.println("generateNewWalletFile :");
        ECKeyPair ecKeyPair = Keys.createEcKeyPair();
//        System.out.println("generateNewWalletFile3");
        return generateWalletFile1(password, ecKeyPair, destinationDirectory);
    }

    public static Map<String, Object> generateWalletFile1(String password, ECKeyPair ecKeyPair, File destinationDirectory)
            throws CipherException, IOException {
//        System.out.println("generateWalletFile:");
        WalletFile walletFile = Wallet.createStandard(password, ecKeyPair);
        String fileName = getWalletFileName(walletFile);
        File destination = new File(destinationDirectory, fileName);

        ObjectMapper objectMapper = ObjectMapperFactory.getObjectMapper();
        objectMapper.writeValue(destination, walletFile);


        Map<String, Object> walletMap = new HashMap<String, Object>();
//        JSONObject sas = (JSONObject) JSON.toJSON(walletFile);
//        String credentials = JSON.toJSONString(walletFile);
//        JSONObject jsonArray = JSON.parseObject(sa);
//        JSONObject jsonObject = JSONObject.fromObject(walletFile);

        walletMap.put("credentials", walletFile);
        walletMap.put("ecKeyPair", ecKeyPair);
//        return fileName;
        return walletMap;
    }


    private static String getWalletFileName(WalletFile walletFile) {
        return "0x" + walletFile.getAddress() + ".json";
    }

/*    private volatile static Web3j web3j;
    public static Web3j getClient(){
        if(web3j==null){
            synchronized (Web3jUtils.class){
                if(web3j==null){
                    web3j = Web3j.build(new HttpService(configSettings.getHttpUrl()));
                }
            }
        }
        return web3j;
    }*/
}
