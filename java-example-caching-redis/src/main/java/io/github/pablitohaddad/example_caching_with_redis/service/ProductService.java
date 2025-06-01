package io.github.pablitohaddad.example_caching_with_redis.service;

import io.github.pablitohaddad.example_caching_with_redis.model.Product;
import io.github.pablitohaddad.example_caching_with_redis.repository.ProductRepository;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cache.annotation.Cacheable;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class ProductService {
    @Autowired
    private ProductRepository productRespository;

    @Cacheable("products")
    public List<Product> listAllProducts() throws InterruptedException {
        return productRespository.findAll();
    }
}
