package com.bottos.bottosapp.bean;

import java.io.Serializable;
import java.math.BigInteger;

/**
 * Bottos.
 *
**/




public class RequirementManagerBean implements Serializable {


    public String getActionAccount() {
        return actionAccount;
    }

    public void setActionAccount(String actionAccount) {
        this.actionAccount = actionAccount;
    }

    public String getOwnerAccount() {
        return ownerAccount;
    }

    public void setOwnerAccount(String ownerAccount) {
        this.ownerAccount = ownerAccount;
    }

    public String getRequirementName() {
        return requirementName;
    }

    public void setRequirementName(String requirementName) {
        this.requirementName = requirementName;
    }

    public String getRequirementType() {
        return requirementType;
    }

    public void setRequirementType(String requirementType) {
        this.requirementType = requirementType;
    }

    public String getDataType() {
        return dataType;
    }

    public void setDataType(String dataType) {
        this.dataType = dataType;
    }

    public String getApplicationDomain() {
        return applicationDomain;
    }

    public void setApplicationDomain(String applicationDomain) {
        this.applicationDomain = applicationDomain;
    }

    public BigInteger getExpirationTime() {
        return expirationTime;
    }

    public void setExpirationTime(BigInteger expirationTime) {
        this.expirationTime = expirationTime;
    }

    public BigInteger getBidMoney() {
        return bidMoney;
    }

    public void setBidMoney(BigInteger bidMoney) {
        this.bidMoney = bidMoney;
    }

    public String getDataSampleRef() {
        return dataSampleRef;
    }

    public void setDataSampleRef(String dataSampleRef) {
        this.dataSampleRef = dataSampleRef;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public String getSpecifications() {
        return specifications;
    }

    public void setSpecifications(String specifications) {
        this.specifications = specifications;
    }

    public String getRequirementID() {
        return requirementID;
    }

    public void setRequirementID(String requirementID) {
        this.requirementID = requirementID;
    }

    public String getCollectionNum() {
        return collectionNum;
    }

    public void setCollectionNum(String collectionNum) {
        this.collectionNum = collectionNum;
    }

    public BigInteger getPublishedTime() {
        return publishedTime;
    }

    public void setPublishedTime(BigInteger publishedTime) {
        this.publishedTime = publishedTime;
    }


    public String getDataSample1() {
        return dataSample1;
    }

    public void setDataSample1(String dataSample1) {
        this.dataSample1 = dataSample1;
    }

    public String getDataSample2() {
        return dataSample2;
    }

    public void setDataSample2(String dataSample2) {
        this.dataSample2 = dataSample2;
    }

    public String getDataSample3() {
        return dataSample3;
    }

    public void setDataSample3(String dataSample3) {
        this.dataSample3 = dataSample3;
    }



    private String actionAccount;

    private String ownerAccount;
    private String requirementName;
    private String requirementType;
    private String dataType;
    private String applicationDomain;
    private BigInteger expirationTime;
    private BigInteger bidMoney;
    private String dataSampleRef;
    private String dataSample1;
    private String dataSample2;
    private String dataSample3;
    private String description;
    private String specifications;


    private String requirementID;
    private String collectionNum;
    private BigInteger publishedTime;

}

