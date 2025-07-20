package com.ecom.products.exceptions;


import org.springframework.http.HttpStatus;
import org.springframework.http.ProblemDetail;
import org.springframework.web.ErrorResponseException;

import java.net.URI;
import java.time.Instant;

public class ProductNotFoundException extends ErrorResponseException {

    public ProductNotFoundException(Long productId, String path) {
        super(HttpStatus.NOT_FOUND, problemDetailFrom("Product with id " + productId + " not found", path), null);
    }

    private static ProblemDetail problemDetailFrom(String message, String path) {
        ProblemDetail problemDetail = ProblemDetail.forStatusAndDetail(HttpStatus.NOT_FOUND, message);
        problemDetail.setType(URI.create("/api/products" + path));
        problemDetail.setTitle("Product not found");
        problemDetail.setInstance(URI.create(path));
        problemDetail.setProperty("timestamp", Instant.now()); // additional data
        return problemDetail;
    }
}