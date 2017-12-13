package com.bottos.bottosapp.bean;

import java.io.Serializable;
import java.math.BigInteger;

/**
 * Bottos.
 */


public class DataAssetBean implements Serializable {


    public String getActionAccount() {
        return actionAccount;
    }

    public void setActionAccount(String actionAccount) {
        this.actionAccount = actionAccount;
    }

    public String getOwner() {
        return owner;
    }

    public void setOwner(String owner) {
        this.owner = owner;
    }

    public String getDataRequirementID() {
        return dataRequirementID;
    }

    public void setDataRequirementID(String dataRequirementID) {
        this.dataRequirementID = dataRequirementID;
    }

    public String getAssetType() {
        return assetType;
    }

    public void setAssetType(String assetType) {
        this.assetType = assetType;
    }

    public String getSubType() {
        return subType;
    }

    public void setSubType(String subType) {
        this.subType = subType;
    }

    public String getApplicationDomain() {
        return applicationDomain;
    }

    public void setApplicationDomain(String applicationDomain) {
        this.applicationDomain = applicationDomain;
    }

    public BigInteger getPriceHigh() {
        return priceHigh;
    }

    public void setPriceHigh(BigInteger priceHigh) {
        this.priceHigh = priceHigh;
    }

    public BigInteger getPriceLow() {
        return priceLow;
    }

    public void setPriceLow(BigInteger priceLow) {
        this.priceLow = priceLow;
    }

    public String getFeatureLabel1() {
        return featureLabel1;
    }

    public void setFeatureLabel1(String featureLabel1) {
        this.featureLabel1 = featureLabel1;
    }

    public String getFeatureLabel2() {
        return featureLabel2;
    }

    public void setFeatureLabel2(String featureLabel2) {
        this.featureLabel2 = featureLabel2;
    }

    public String getFeatureLabel3() {
        return featureLabel3;
    }

    public void setFeatureLabel3(String featureLabel3) {
        this.featureLabel3 = featureLabel3;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public String getDataStoreID() {
        return dataStoreID;
    }

    public void setDataStoreID(String dataStoreID) {
        this.dataStoreID = dataStoreID;
    }

    public String getAssetStatus() {
        return assetStatus;
    }

    public void setAssetStatus(String assetStatus) {
        this.assetStatus = assetStatus;
    }

    public String getAssetID() {
        return assetID;
    }

    public void setAssetID(String assetID) {
        this.assetID = assetID;
    }

    public String getSize() {
        return size;
    }

    public void setSize(String size) {
        this.size = size;
    }

    public String getNonce() {
        return nonce;
    }

    public void setNonce(String nonce) {
        this.nonce = nonce;
    }


    public BigInteger getRegisterTime() {
        return registerTime;
    }

    public void setRegisterTime(BigInteger registerTime) {
        this.registerTime = registerTime;
    }

    public BigInteger getExpirationTime() {
        return expirationTime;
    }

    public void setExpirationTime(BigInteger expirationTime) {
        this.expirationTime = expirationTime;
    }

    private String actionAccount;

    private String owner;

    private String dataRequirementID;
    private String assetType;
    private String subType;
    private String applicationDomain;
    private BigInteger priceHigh;
    private BigInteger priceLow;
    private String featureLabel1;
    private String featureLabel2;
    private String featureLabel3;
    private String description;
    private String dataStoreID;

    private String assetStatus;
    private String assetID;
    private String size;
    private String nonce;

    private BigInteger registerTime;
    private BigInteger expirationTime;

}

