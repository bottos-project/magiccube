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
public class UserInfo implements Serializable {
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

    @Id
    @GeneratedValue
    private Integer id;
    private String name;
    private String passwd;
    private String mobile;
    private String email;
    private String userType;
    private String roleType;
    private String companyName;
    private String companyAddr;
    private String orgCode;
    private String verificationCode;

    private String publicKey;
    private String praviteKey;
    private String credentials;

    private String account;
    private String userAddr;
    private String accountAlias;
    private String tokenType;
    private int accountStatus;
    private int passwordStatus;          // -
    private int deleteStatus;            // -
    private String uuid;


    private String cipherGroupKey;
    private long createTime;
    private long updateTime;
    private long loginTime;                //-
    private int loginStatus;
    private String departmentId;

//    private String[] roleIdList;

    public UserInfo() {
    }

    public String getUserAddr() {
        return userAddr;
    }

    public void setUserAddr(String userAddr) {
        this.userAddr = userAddr;
    }

    public String getDepartmentId() {
        return departmentId;
    }

    public void setDepartmentId(String departmentId) {
        this.departmentId = departmentId;
    }

    public String getTokenType() {
        return tokenType;
    }

    public void setTokenType(String tokenType) {
        this.tokenType = tokenType;
    }

    public String getAccountAlias() {
        return accountAlias;
    }

    public void setAccountAlias(String accountAlias) {
        this.accountAlias = accountAlias;
    }

    public String getPraviteKey() {
        return praviteKey;
    }

    public void setPraviteKey(String praviteKey) {
        this.praviteKey = praviteKey;
    }

    public String getCredentials() {
        return credentials;
    }

    public void setCredentials(String credentials) {
        this.credentials = credentials;
    }

    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public String getAccount() {
        return account;
    }

    public String getEmail() {
        return email;
    }

    public String getMobile() {
        return mobile;
    }

    public int getAccountStatus() {
        return accountStatus;
    }

    public int getPasswordStatus() {
        return passwordStatus;
    }

    public int getDeleteStatus() {
        return deleteStatus;
    }

    public String getUuid() {
        return uuid;
    }

    public String getPublicKey() {
        return publicKey;
    }

    public String getCipherGroupKey() {
        return cipherGroupKey;
    }

    public long getCreateTime() {
        return createTime;
    }

    public long getUpdateTime() {
        return updateTime;
    }

    public long getLoginTime() {
        return loginTime;
    }

    public void setName(String name) {
        this.name = name;
    }

    public void setAccount(String account) {
        this.account = account;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public void setMobile(String mobile) {
        this.mobile = mobile;
    }

    public void setAccountStatus(int accountStatus) {
        this.accountStatus = accountStatus;
    }

    public void setPasswordStatus(int passwordStatus) {
        this.passwordStatus = passwordStatus;
    }

    public void setDeleteStatus(int deleteStatus) {
        this.deleteStatus = deleteStatus;
    }

    public void setUuid(String uuid) {
        this.uuid = uuid;
    }

    public void setPublicKey(String publicKey) {
        this.publicKey = publicKey;
    }

    public void setCipherGroupKey(String cipherGroupKey) {
        this.cipherGroupKey = cipherGroupKey;
    }

    public void setCreateTime(long createTime) {
        this.createTime = createTime;
    }

    public void setUpdateTime(long updateTime) {
        this.updateTime = updateTime;
    }

    public void setLoginTime(long loginTime) {
        this.loginTime = loginTime;
    }

    public void setPasswd(String passwd) {
        this.passwd = passwd;
    }

    public String getPasswd() {
        return passwd;
    }

    public int getLoginStatus() {
        return loginStatus;
    }

    public void setLoginStatus(int loginStatus) {
        this.loginStatus = loginStatus;
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
}
