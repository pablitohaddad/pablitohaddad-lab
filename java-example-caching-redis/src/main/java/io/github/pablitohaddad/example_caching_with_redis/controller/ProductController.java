package io.github.pablitohaddad.example_caching_with_redis.controller;

import io.github.pablitohaddad.example_caching_with_redis.model.Product;
import io.github.pablitohaddad.example_caching_with_redis.service.ProductService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.List;

@RestController
@RequestMapping("/products")
public class ProductController {

    @Autowired
    private ProductService productService;

    @GetMapping
    public List<Product> getProducts() throws InterruptedException {
        return productService.listAllProducts();
    }
}
