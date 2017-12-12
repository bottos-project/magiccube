package com.bottos.bottosapp.bean;

import java.io.Serializable;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;

/**
 * Bottos.
 */


public class WalletInfoBean implements Serializable {

    private List<Map<String,Object>> totalTokenInfo;

    private List<AccountInfoBean> accountInfoList;



    public List<Map<String, Object>> getTotalTokenInfo() {
        return totalTokenInfo;
    }

    public void setTotalTokenInfo(List<Map<String, Object>> totalTokenInfo) {
        this.totalTokenInfo = totalTokenInfo;
    }

    public List<AccountInfoBean> getAccountInfoList() {
        return accountInfoList;
    }

    public void setAccountInfoList(AccountInfoBean accountInfoList) {
        List<AccountInfoBean> list = new ArrayList<>();
        list.add(accountInfoList);
        this.accountInfoList = list;
    }
}

