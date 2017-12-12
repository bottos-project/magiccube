package com.bottos.bottosapp.dao;

/**
 *
 * @param <T>
 * @param <Model>
 */
public interface BaseDao<T, Model> {
//    RegisterManager getRegisterContract();

    T getContract4Unlogin(String walletPassword, String walletAdmin, Class c);

//    PageInfo<Model> pageByExample(Integer var1, Integer var2, List<Model> var3);
}