package com.bottos.bottosapp.bean;

/**
 *
 */
public class UserClientMsg {


    public int getTotal() {
        return total;
    }

    public void setTotal(int total) {
        this.total = total;
    }

//    public UserInfo[] getItems() {
//        return items;
//    }
    public UserResponseBean[] getItems() {
        return items;
    }

//    public void setItems(UserInfo[] items) {
//        this.items = items;
//    }

    public void setItems(UserResponseBean[] items) {
        this.items = items;
    }


    public int getReturnCode() {
        return returnCode;
    }

    public void setReturnCode(int returnCode) {
        this.returnCode = returnCode;
    }

    public String getReturnDesc() {
        return returnDesc;
    }

    public void setReturnDesc(String returnDesc) {
        this.returnDesc = returnDesc;
    }

    private int returnCode;
    private String returnDesc;
    private int total;
//    private UserInfo[] items;
    private UserResponseBean[] items;
}
