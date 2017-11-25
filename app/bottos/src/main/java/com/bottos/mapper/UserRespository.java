package com.bottos.mapper;

import com.bottos.bean.UserInfo;
import org.springframework.data.jpa.repository.JpaRepository;

import java.util.List;

public interface UserRespository extends JpaRepository<UserInfo, Integer> {

    public UserInfo findByName(String name);
    public UserInfo findByAccount(String account);


}
