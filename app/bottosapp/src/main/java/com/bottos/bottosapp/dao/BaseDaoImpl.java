package com.bottos.bottosapp.dao;

import com.bottos.bottosapp.common.SpringContextUtil;
import com.bottos.bottosapp.common.Web3Manager;
import com.bottos.bottosapp.common.exception.DaoException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.lang.reflect.ParameterizedType;
import java.lang.reflect.Type;
import java.util.concurrent.ExecutionException;

/**
 *
 * @param <T>
 * @param <Model>
 */
public abstract class BaseDaoImpl<T, Model> implements BaseDao<T, Model> {
    protected Logger logger = LoggerFactory.getLogger(this.getClass().getName());
    private Class<T> entityClass;

    public BaseDaoImpl() {
        Type genType = this.getClass().getGenericSuperclass();
        Type[] params = ((ParameterizedType) genType).getActualTypeArguments();
        this.entityClass = (Class) params[0];
    }

//    public RegisterManager getRegisterContract() {
//        Web3Manager web3Manager = (Web3Manager) SpringContextUtil.getBean("web3Manager");
//        return web3Manager.getRegisterManager();
//    }

    public T getContract() {
//        Session session = SecurityUtils.getSubject().getSession();
//        Web3Manager web3Manager = (Web3Manager) SpringContextUtil.getBean("web3Manager");
//        Map<String, Object> constractMap = (Map) web3Manager.getCurrentContractMap().get(session.getId().toString());
//        if (constractMap != null) {
//            return constractMap.get(this.entityClass.getName());
//        } else {
//            this.logger.error("getContract()--->please reLogin，【sessionId=" + session.getId() + "】");
//            throw new IllegalArgumentException("请重新登录");
//        }
        return null;
    }

    public T getContract4Unlogin(String accountPasswd, String accountFile, Class c) {
        Web3Manager web3Manager = (Web3Manager) SpringContextUtil.getBean("web3Manager");
        return (T) web3Manager.initContract(accountPasswd, accountFile, c);

    }


    protected boolean exceptionDetailDeal(Exception e) {
        if (e instanceof ExecutionException) {
            if (e.getMessage().contains("org.web3j.protocol.exceptions.TransactionTimeoutException")) {
                throw new DaoException("交易超时");
            } else {
                throw new DaoException("系统出错，请联系管理员");
            }
        } else if (e instanceof DaoException) {
            throw (DaoException) e;
        } else {
            throw new DaoException("系统出错，请联系管理员");
        }
    }
}