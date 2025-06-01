package io.github.pablitohaddad.example_caching_with_redis;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.cache.annotation.EnableCaching;

@SpringBootApplication
// Enabled the app for caching :)
@EnableCaching
public class ExampleCachingWithRedisApplication {

	public static void main(String[] args) {
		SpringApplication.run(ExampleCachingWithRedisApplication.class, args);
	}

}
