package com.bottos.bottosapp;

import com.bottos.bottosapp.common.ConfigSettings;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.context.properties.EnableConfigurationProperties;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

@Controller
@EnableConfigurationProperties(ConfigSettings.class)
@SpringBootApplication
public class BottosappApplication {

	public static void main(String[] args) {
		SpringApplication.run(BottosappApplication.class, args);
	}

	@RequestMapping(value = "/")
	@ResponseBody
	String home(){
		return "home";
	}
}
