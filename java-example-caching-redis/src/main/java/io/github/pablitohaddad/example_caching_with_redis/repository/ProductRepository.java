package io.github.pablitohaddad.example_caching_with_redis.repository;

import io.github.pablitohaddad.example_caching_with_redis.model.Product;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;

@Repository
public interface ProductRepository extends JpaRepository<Product, Long> {

    Optional<Product> findByNameAndCategory(String name, String category);
}

