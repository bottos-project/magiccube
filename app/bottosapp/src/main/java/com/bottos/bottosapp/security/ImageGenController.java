package com.bottos.bottosapp.security;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Controller;
import org.springframework.util.StringUtils;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.Cookie;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

@Controller
public class ImageGenController {
    protected Logger logger = LoggerFactory.getLogger(this.getClass().getName());

//    @RequestMapping(value = "/toImg")
//    public String toImg() {
//        return "image";
//    }

    //Login to get the verification code
    @RequestMapping("/getSysManageLoginCode")
    @ResponseBody
    public String getSysManageLoginCode(HttpServletResponse response,
                                        HttpServletRequest request) {
        response.setContentType("image/jpeg");// Set the appropriate type, tell the browser output is the picture
        response.setHeader("Pragma", "No-cache");// Set the response header to tell the browser not to cache this content
        response.setHeader("Cache-Control", "no-cache");
        response.setHeader("Set-Cookie", "name=value; HttpOnly");//Set HttpOnly property to prevent Xss attacks
        response.setDateHeader("Expire", 0);
        RandomValidateCode randomValidateCode = new RandomValidateCode();
        try {
            randomValidateCode.getRandcode(request, response, "imagecode");// Output picture method
        } catch (Exception e) {
            e.printStackTrace();
        }
        return "";
    }

    //Verification code verification
    @RequestMapping(value = "/checkimagecode")
    @ResponseBody
    public String checkTcode(HttpServletRequest request, HttpServletResponse response) {
        String validateCode = request.getParameter("validateCode");
        String code = null;
//1:Get the cookie inside the verification code information
        Cookie[] cookies = request.getCookies();
        for (Cookie cookie : cookies) {
            if ("imagecode".equals(cookie.getName())) {
                code = cookie.getValue();
                break;
            }
        }
//1:Get session verification code information
// String code1 = (String) request.getSession().getAttribute("");
//2:Determine the verification code is correct
        if (!StringUtils.isEmpty(validateCode) && validateCode.equalsIgnoreCase(code)) {
            logger.info("verification code OK.");
            return "ok";
        }
        logger.error("verification code error. writeCode=" + validateCode + " validateCode:" + code);
        return "error";
    }

}