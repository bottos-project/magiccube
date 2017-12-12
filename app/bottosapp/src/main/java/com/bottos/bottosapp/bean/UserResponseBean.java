package com.bottos.bottosapp.bean;

import java.io.Serializable;

/**
 *
 */

public class UserResponseBean implements Serializable {
    public enum UserError {
        NO_ERROR,
        BAD_PARAMETER,
        NAME_EMPTY,
        DEPT_NOT_EXISTS,
        ROLE_ID_INVALID,
        ROLE_ID_EXCEED_DEPT,
        USER_NOT_EXISTS,
        ROLE_ID_ALREADY_EXISTS,
        ADDRESS_ALREADY_EXISTS,
        ACCOUNT_ALREDY_EXISTS,
        ACCOUNT_CANNOT_UPDATE,
        USER_LOGIN_FAILED,
        DEPT_CANNOT_UPDATE,
        NO_PERMISSION
    }


    public String getName() {
        return name;
    }


    public String getEmail() {
        return email;
    }

    public String getMobile() {
        return mobile;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public void setMobile(String mobile) {
        this.mobile = mobile;
    }

    public String getVerificationCode() {
        return verificationCode;
    }

    public void setVerificationCode(String verificationCode) {
        this.verificationCode = verificationCode;
    }

    public String getUserType() {
        return userType;
    }

    public void setUserType(String userType) {
        this.userType = userType;
    }

    public String getRoleType() {
        return roleType;
    }

    public void setRoleType(String roleType) {
        this.roleType = roleType;
    }

    public String getCompanyName() {
        return companyName;
    }

    public void setCompanyName(String companyName) {
        this.companyName = companyName;
    }

    public String getCompanyAddr() {
        return companyAddr;
    }

    public void setCompanyAddr(String companyAddr) {
        this.companyAddr = companyAddr;
    }

    public String getOrgCode() {
        return orgCode;
    }

    public void setOrgCode(String orgCode) {
        this.orgCode = orgCode;
    }
    public String getAccount() {
        return account;
    }

    public void setAccount(String account) {
        this.account = account;
    }

    private String name;
    //    private String passwd;
    private String account;
    private String mobile;
    private String email;
    private String userType;
    private String roleType;
    private String companyName;
    private String companyAddr;
    private String orgCode;
    private String verificationCode;


//    private String userAddr;
//    private String departmentId;
//    private int accountStatus;
//    private int passwordStatus;          // -
//    private int deleteStatus;            // -
//    private String uuid;
//    private String publicKey;
//    private String cipherGroupKey;
//    private long createTime;
//    private long updateTime;
//    private long loginTime;                //-
//    private int loginStatus;
//    private String[] roleIdList;

}
