package com.bottos.bottosapp.bean;

import java.io.Serializable;
import java.math.BigInteger;
import java.util.List;
import java.util.Map;

/**
 * Bottos.
 */


public class AccountInfoBean implements Serializable {

    private String accountAlias;
    private String account;
    private String tokenType;
    private BigInteger accountToken;

    private String receiveAccount;




    public String getReceiveAccount() {
        return receiveAccount;
    }

    public void setReceiveAccount(String receiveAccount) {
        this.receiveAccount = receiveAccount;
    }
    public String getAccountAlias() {
        return accountAlias;
    }

    public void setAccountAlias(String accountAlias) {
        this.accountAlias = accountAlias;
    }

    public String getAccount() {
        return account;
    }

    public void setAccount(String account) {
        this.account = account;
    }

    public String getTokenType() {
        return tokenType;
    }

    public void setTokenType(String tokenType) {
        this.tokenType = tokenType;
    }

    public BigInteger getAccountToken() {
        return accountToken;
    }

    public void setAccountToken(BigInteger accountToken) {
        this.accountToken = accountToken;
    }
}

