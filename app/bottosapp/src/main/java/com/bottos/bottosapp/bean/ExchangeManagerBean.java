package com.bottos.bottosapp.bean;

import java.io.Serializable;
import java.math.BigInteger;

/**
 * Bottos.
 */


public class ExchangeManagerBean implements Serializable {

    public String getActionAccount() {
        return actionAccount;
    }

    public void setActionAccount(String actionAccount) {
        this.actionAccount = actionAccount;
    }

    public String getExchangeID() {
        return exchangeID;
    }

    public void setExchangeID(String exchangeID) {
        this.exchangeID = exchangeID;
    }

    public String getExchangeTime() {
        return exchangeTime;
    }

    public void setExchangeTime(String exchangeTime) {
        this.exchangeTime = exchangeTime;
    }

    public String getRequirementID() {
        return requirementID;
    }

    public void setRequirementID(String requirementID) {
        this.requirementID = requirementID;
    }

    public String getAssetID() {
        return assetID;
    }

    public void setAssetID(String assetID) {
        this.assetID = assetID;
    }

    public BigInteger getPrice() {
        return price;
    }

    public void setPrice(BigInteger price) {
        this.price = price;
    }

    public String getRequirementOwnerAccount() {
        return requirementOwnerAccount;
    }

    public void setRequirementOwnerAccount(String requirementOwnerAccount) {
        this.requirementOwnerAccount = requirementOwnerAccount;
    }

    public String getAssetOwnerAccount() {
        return assetOwnerAccount;
    }

    public void setAssetOwnerAccount(String assetOwnerAccount) {
        this.assetOwnerAccount = assetOwnerAccount;
    }

    public BigInteger getStatus() {
        return status;
    }

    public void setStatus(BigInteger status) {
        this.status = status;
    }

    private String actionAccount;

    private String exchangeID;
    private String exchangeTime;
    private String requirementID;
    private String assetID;
    private BigInteger status;
    private BigInteger price;
    private String requirementOwnerAccount;
    private String assetOwnerAccount;


}

