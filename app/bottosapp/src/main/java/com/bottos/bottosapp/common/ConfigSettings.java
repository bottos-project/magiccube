package com.bottos.bottosapp.common;

import org.springframework.boot.context.properties.ConfigurationProperties;

import java.math.BigInteger;

@ConfigurationProperties(prefix = "bottos")
public class ConfigSettings {

    private String httpUrl;
    private String walletPath;
    private String gasPrice;
    private String gasLimit;

    private String userManagerContractAddr;
    private String assetManagerContractAddr;
    private String requirementManagerContractAddr;
    private String exchangeManagerContractAddr;

    private String tokenManagerContractAddr;

    private String maxWait;
    private String filters;
    private String maxTotal;
    private String maxPerRoute;

    private String adminAccount;
    private String adminAccountPasswd;
    private String adminAccountEth;



    private BigInteger defaultAmountBo;
    private double defaultAmountEth;

    private String adminAccountPasswdEth;
    private String uploadPathRequire;
    private String uploadPathDataAsset;


    public BigInteger getDefaultAmountBo() {
        return defaultAmountBo;
    }

    public void setDefaultAmountBo(BigInteger defaultAmountBo) {
        this.defaultAmountBo = defaultAmountBo;
    }

    public double getDefaultAmountEth() {
        return defaultAmountEth;
    }

    public void setDefaultAmountEth(double defaultAmountEth) {
        this.defaultAmountEth = defaultAmountEth;
    }

    public String getAdminAccountEth() {
        return adminAccountEth;
    }

    public void setAdminAccountEth(String adminAccountEth) {
        this.adminAccountEth = adminAccountEth;
    }

    public String getAdminAccountPasswdEth() {
        return adminAccountPasswdEth;
    }

    public void setAdminAccountPasswdEth(String adminAccountPasswdEth) {
        this.adminAccountPasswdEth = adminAccountPasswdEth;
    }




    public String getUploadPathRequire() {
        return uploadPathRequire;
    }

    public void setUploadPathRequire(String uploadPathRequire) {
        this.uploadPathRequire = uploadPathRequire;
    }

    public String getUploadPathDataAsset() {
        return uploadPathDataAsset;
    }

    public void setUploadPathDataAsset(String uploadPathDataAsset) {
        this.uploadPathDataAsset = uploadPathDataAsset;
    }

    public String getTokenManagerContractAddr() {
        return tokenManagerContractAddr;
    }

    public void setTokenManagerContractAddr(String tokenManagerContractAddr) {
        this.tokenManagerContractAddr = tokenManagerContractAddr;
    }

    public String getAdminAccount() {
        return adminAccount;
    }

    public void setAdminAccount(String adminAccount) {
        this.adminAccount = adminAccount;
    }

    public String getAdminAccountPasswd() {
        return adminAccountPasswd;
    }

    public void setAdminAccountPasswd(String adminAccountPasswd) {
        this.adminAccountPasswd = adminAccountPasswd;
    }

    public String getMaxWait() {
        return maxWait;
    }

    public void setMaxWait(String maxWait) {
        this.maxWait = maxWait;
    }

    public String getFilters() {
        return filters;
    }

    public void setFilters(String filters) {
        this.filters = filters;
    }

    public String getMaxTotal() {
        return maxTotal;
    }

    public void setMaxTotal(String maxTotal) {
        this.maxTotal = maxTotal;
    }

    public String getMaxPerRoute() {
        return maxPerRoute;
    }

    public void setMaxPerRoute(String maxPerRoute) {
        this.maxPerRoute = maxPerRoute;
    }

    public String getUserManagerContractAddr() {
        return userManagerContractAddr;
    }

    public void setUserManagerContractAddr(String userManagerContractAddr) {
        this.userManagerContractAddr = userManagerContractAddr;
    }

    public String getAssetManagerContractAddr() {
        return assetManagerContractAddr;
    }

    public void setAssetManagerContractAddr(String assetManagerContractAddr) {
        this.assetManagerContractAddr = assetManagerContractAddr;
    }

    public String getRequirementManagerContractAddr() {
        return requirementManagerContractAddr;
    }

    public void setRequirementManagerContractAddr(String requirementManagerContractAddr) {
        this.requirementManagerContractAddr = requirementManagerContractAddr;
    }

    public String getExchangeManagerContractAddr() {
        return exchangeManagerContractAddr;
    }

    public void setExchangeManagerContractAddr(String exchangeManagerContractAddr) {
        this.exchangeManagerContractAddr = exchangeManagerContractAddr;
    }

    public String getHttpUrl() {
        return httpUrl;
    }

    public void setHttpUrl(String httpUrl) {
        this.httpUrl = httpUrl;
    }

    public String getGasPrice() {
        return gasPrice;
    }

    public void setGasPrice(String gasPrice) {
        this.gasPrice = gasPrice;
    }

    public String getGasLimit() {
        return gasLimit;
    }

    public void setGasLimit(String gasLimit) {
        this.gasLimit = gasLimit;
    }



    public String getWalletPath() {
        return walletPath;
    }

    public void setWalletPath(String walletPath) {
        this.walletPath = walletPath;
    }
}
