package com.ecom.products.exceptions;

import com.ecom.products.controllers.ProductController;
import org.springframework.dao.DuplicateKeyException;
import org.springframework.http.*;
import org.springframework.util.ObjectUtils;
import org.springframework.validation.FieldError;
import org.springframework.validation.method.MethodValidationException;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestControllerAdvice;
import org.springframework.web.context.request.WebRequest;
import org.springframework.web.servlet.mvc.method.annotation.ResponseEntityExceptionHandler;

import java.net.URI;
import java.time.Instant;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.Optional;

@RestControllerAdvice
public class    ProductExceptionHandler {
    @ExceptionHandler(DuplicateKeyException.class)
    ResponseEntity<ProblemDetail> handleDuplicateKeyException(DuplicateKeyException ex, WebRequest request) {
        ProblemDetail problemDetail = ProblemDetail.forStatusAndDetail(HttpStatus.CONFLICT, ex.getMessage());
        problemDetail.setTitle("Duplicate Key Error");
        problemDetail.setInstance(URI.create(request.getContextPath()));
        problemDetail.setProperty("timestamp", Instant.now()); // adding more data using the Map properties of the ProblemDetail
        return ResponseEntity.status(HttpStatus.CONFLICT)
                .contentType(MediaType.APPLICATION_PROBLEM_JSON)
                .body(problemDetail);
    }

    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<Object> handleMethodArgumentNotValidException(MethodArgumentNotValidException ex,  WebRequest request) {
        ProblemDetail problemDetail = ProblemDetail.forStatusAndDetail(HttpStatus.BAD_REQUEST, handleMethodArgumentNotValidMessage(ex));
        problemDetail.setTitle("Bad request");
        problemDetail.setInstance(URI.create(request.getContextPath()));
        problemDetail.setProperty("timestamp", Instant.now()); // adding more data using the Map properties of the ProblemDetail
        return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(problemDetail);
    }

    protected String handleMethodArgumentNotValidMessage(MethodArgumentNotValidException ex) {
        return Optional.ofNullable(ex.getBindingResult().getFieldError())
                .map(FieldError::getDefaultMessage)
                .orElse("Invalid request");
    }
}
