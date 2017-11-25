package com.bottos;

import com.bottos.common.ConfigSettings;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.context.properties.EnableConfigurationProperties;

@SpringBootApplication
@EnableConfigurationProperties(ConfigSettings.class)

public class Application {

	public static void main(String[] args) {
		SpringApplication.run(Application.class, args);
	}
}
