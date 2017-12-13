package com.bottos.bottosapp.mapper;

import com.bottos.bottosapp.bean.UserInfo;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.JpaSpecificationExecutor;
import org.springframework.data.jpa.repository.Modifying;
import org.springframework.data.jpa.repository.Query;

import java.io.Serializable;

public interface UserRespository extends JpaRepository<UserInfo, Integer>,
        JpaSpecificationExecutor<UserInfo>,
        Serializable {

    public UserInfo findByName(String name);

    public UserInfo findByAccount(String account);

//    public List<UserInfo> findByName(String name);

    //根据用户名、密码删除一条数据
    @Modifying
    @Query(value = "delete from t_user where t_name = ?1 and t_pwd = ?2",nativeQuery = true)
    public void deleteQuery(String name, String pwd);


}
