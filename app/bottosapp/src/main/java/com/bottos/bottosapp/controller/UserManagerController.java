package com.bottos.bottosapp.controller;

import com.alibaba.fastjson.JSON;
import com.bottos.bottosapp.bean.UserClientMsg;
import com.bottos.bottosapp.bean.UserInfo;
import com.bottos.bottosapp.bean.UserResponseBean;
import com.bottos.bottosapp.mapper.UserRespository;
import com.bottos.bottosapp.service.UserManagerServiceImpl;
import org.apache.catalina.servlet4preview.http.HttpServletRequest;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;

@Controller
@RequestMapping(value = "/user")
public class UserManagerController {
    protected Logger logger = LoggerFactory.getLogger(this.getClass().getName());

    @Resource
    UserManagerServiceImpl userManagerService;
    @Autowired
    private UserRespository userRespository;

//    @ResponseBody
//    @GetMapping(produces = "text/plain;charset=UTF-8")
////    @RequestMapping(produces = "text/plain;charset=UTF-8")//produces描述的是响应的头信息的Content-Type字段
//    public String user(HttpServletRequest request) {
//        //url:http://localhost:8080/user can access
//        return "url:" + request.getRequestURL() + " can access";
//    }

    //    test
    @RequestMapping(value = "/")
    @ResponseBody
    String hello() {
        return "hello";
    }

    /**
     * @param json the request body
     * @return login result
     */

    @ResponseBody
    @RequestMapping(value = "/login", method = RequestMethod.POST)
    public String login(HttpServletRequest request, @RequestBody String json) {
        logger.info("start login. json=" + json.toString());
        int ret = -1;
        String retDes = "";

        try {
            UserInfo user = JSON.parseObject(json, UserInfo.class);
//            Verification Code
            boolean flag = userManagerService.checkTcode(request, user.getVerificationCode());
            if (!flag) {
                return retToJson(-2, "login failed. verification code error.", null);
            }
            String result = userManagerService.login(user);
            //判断执行结果
            String returnJson = "";
            if (result.equals("")) {
                String account = userRespository.findByName(user.getName()).getAccount();
                UserResponseBean userRespObj = new UserResponseBean();
                userRespObj.setName(user.getName());
                userRespObj.setAccount(account);

                returnJson = retToJson(0, "", userRespObj);
            } else {
                returnJson = retToJson(-1, result, null);
            }
            logger.info("end login.", "result=", result);
            return returnJson;
        } catch (Exception e) {
            e.printStackTrace();
            return retToJson(-1, "login failed.", null);
        }
    }


    /**
     * @param json the request
     * @return register user result
     */
    @ResponseBody
    @RequestMapping(value = "/register", method = RequestMethod.POST)
    public String registerUser(HttpServletRequest request, @RequestBody String json) {
        logger.info("start register User. json=" + json.toString());

        try {
            UserInfo newUser = JSON.parseObject(json, UserInfo.class);
//            Verification Code
            boolean flag = userManagerService.checkTcode(request, newUser.getVerificationCode());
            if (!flag) {
//                code=-2,verification code error
                return retToJson(-2, "login failed. verification code error.", null);
            }
//      store user info in DB
//            UserInfo storeResult = userManagerService.saveUser(newUser);
//            Integer storeId = storeResult.getId();

            logger.info("start addUser.", "data=", newUser.getUuid());
//            String userSession = jedis.get(newUser.getUuid());
//            if(userSession.equals("")){
//                return "please login first.";
//            }
            //判断执行结果
//            UserInfo serverUsr = JSON.parseObject(userSession, UserInfo.class);
            String result = userManagerService.addUser(newUser, newUser);

            String returnJson = "";
            if (result.equals("")) {
                UserResponseBean resultObj = new UserResponseBean();
                resultObj.setName(newUser.getName());
                resultObj.setAccount(newUser.getAccount());
//                resultObj.setMobile(newUser.getMobile());
//                resultObj.setEmail(newUser.getEmail());
//                resultObj.setUserType(newUser.getUserType());
//                resultObj.setRoleType(newUser.getRoleType());
//                resultObj.setCompanyName(newUser.getCompanyName());
//                resultObj.setCompanyAddr(newUser.getCompanyAddr());
//                resultObj.setOrgCode(newUser.getOrgCode());

                returnJson = retToJson(0, "", resultObj);
//                returnJson = retToJson(0, "", newUser);
            } else {
//      If adding a user on BlockChain fails, the user is deleted from the DB
                if (!result.equals("account exists")) {
                    userManagerService.deleteOneUser();
                } else {
//                    code=11,account exists
                    return retToJson(11, "register User failed, " + result, null);
                }
                returnJson = retToJson(-1, result, null);
            }
            logger.info("end register User.", "result=", result);
            return returnJson;
        } catch (Exception e) {
            e.printStackTrace();
            logger.info("end register User.", "addUser failed.");
            return retToJson(-1, "addUser failed.", null);
        }
    }


    /**
     * @param json the request
     * @return loginOut result.
     */
    @ResponseBody
    @RequestMapping(value = "/logOut", method = RequestMethod.POST)
    public String logOut(@RequestBody String json) {

        return null;
    }


    private String retToJson(int ret, String retDes, UserResponseBean userInfo) {
        UserClientMsg userClientMsg = new UserClientMsg();

        userClientMsg.setReturnCode(ret);
        userClientMsg.setReturnDesc(retDes);
        if (userInfo != null) {
            UserResponseBean[] returnInfo = {userInfo};
            userClientMsg.setItems(returnInfo);
        }
        String returnJson = JSON.toJSONString(userClientMsg);
        return returnJson;
    }
}
