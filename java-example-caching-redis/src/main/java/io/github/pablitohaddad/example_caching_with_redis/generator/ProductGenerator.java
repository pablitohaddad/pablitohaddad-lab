package io.github.pablitohaddad.example_caching_with_redis.generator;

import io.github.pablitohaddad.example_caching_with_redis.model.Product;
import io.github.pablitohaddad.example_caching_with_redis.repository.ProductRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;

import java.util.ArrayList;
import java.util.List;

@Component
public class ProductGenerator implements CommandLineRunner {

    @Autowired
    private ProductRepository productRepository;

    private static final int TOTAL_PRODUCTS = 500_000;
    private static final int BATCH_SIZE = 100_000;

    @Override
    public void run(String... args) throws Exception {
        generateProducts();
    }

    public void generateProducts() {
        long startTime = System.currentTimeMillis();

        for (int batchStart = 1; batchStart <= TOTAL_PRODUCTS; batchStart += BATCH_SIZE) {
            List<Product> products = new ArrayList<>(BATCH_SIZE);

            for (int i = 0; i < BATCH_SIZE && (batchStart + i) <= TOTAL_PRODUCTS; i++) {
                int productId = batchStart + i;
                Product product = new Product();
                product.setName("Product " + productId);
                product.setCategory("Category " + (productId % 10));
                product.setPrice(String.valueOf(Math.round(Math.random() * 1000)));
                products.add(product);
            }

            productRepository.saveAll(products);
            System.out.printf("Saved products up to: %d%n", batchStart + BATCH_SIZE - 1);
        }

        long endTime = System.currentTimeMillis();
        System.out.printf("Finished inserting %d products in %d ms%n", TOTAL_PRODUCTS, (endTime - startTime));
    }
}
