package com.bottos.bottosapp.bean;

import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.Id;
import java.io.Serializable;

/**
 *
 */

@Entity
//@Table(name = "t_user")
public class DataAssetEntity implements Serializable {

    @Id
    @GeneratedValue
    private Integer id;
    private String account;
    private String requireAssetID;
    private String assetID;
    private String storeID;
    private String storeAddress;
    private String SignVlaue;


    public String getSignVlaue() {
        return SignVlaue;
    }

    public void setSignVlaue(String signVlaue) {
        SignVlaue = signVlaue;
    }

    public String getRequireAssetID() {
        return requireAssetID;
    }

    public void setRequireAssetID(String requireAssetID) {
        this.requireAssetID = requireAssetID;
    }

    public String getAssetID() {
        return assetID;
    }

    public void setAssetID(String assetID) {
        this.assetID = assetID;
    }

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public String getAccount() {
        return account;
    }

    public void setAccount(String account) {
        this.account = account;
    }

    public String getStoreID() {
        return storeID;
    }

    public void setStoreID(String storeID) {
        this.storeID = storeID;
    }

    public String getStoreAddress() {
        return storeAddress;
    }

    public void setStoreAddress(String storeAddress) {
        this.storeAddress = storeAddress;
    }
}
