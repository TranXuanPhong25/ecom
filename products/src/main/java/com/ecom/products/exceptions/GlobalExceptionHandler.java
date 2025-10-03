package com.ecom.products.exceptions;

import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.dao.DuplicateKeyException;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ProblemDetail;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.FieldError;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;
import org.springframework.web.context.request.WebRequest;
import org.springframework.web.method.annotation.MethodArgumentTypeMismatchException;

import java.net.URI;
import java.time.Instant;
import java.util.Optional;
import java.util.UUID;

@RestControllerAdvice
public class GlobalExceptionHandler {
    @ExceptionHandler(MethodArgumentTypeMismatchException.class)
    ResponseEntity<ProblemDetail> handleMethodArgumentTypeMismatchException(
            MethodArgumentTypeMismatchException ex, WebRequest request) {
        String detailMessage = "";
        if (ex.getRequiredType() != null) {
            if (ex.getRequiredType().getName().equals(UUID.class.getName())) {
                detailMessage = "Invalid UUID format for parameter '" + ex.getName() + "': " + ex.getValue();
            }
        } else {
            detailMessage = "Invalid value for parameter '" + ex.getName() + "': " + ex.getValue();
        }
        ProblemDetail problemDetail = ProblemDetail.forStatusAndDetail(HttpStatus.BAD_REQUEST, detailMessage);
        problemDetail.setTitle("Bad Request");
        problemDetail.setInstance(URI.create(request.getContextPath()));
        problemDetail.setProperty("timestamp", Instant.now()); // adding more data using the Map properties of the
                                                               // ProblemDetail
        return ResponseEntity.status(HttpStatus.BAD_REQUEST)
                .contentType(MediaType.APPLICATION_PROBLEM_JSON)
                .body(problemDetail);
    }

    @ExceptionHandler(DuplicateKeyException.class)
    ResponseEntity<ProblemDetail> handleDuplicateKeyException(DuplicateKeyException ex, WebRequest request) {
        ProblemDetail problemDetail = ProblemDetail.forStatusAndDetail(HttpStatus.CONFLICT, ex.getMessage());
        problemDetail.setTitle("Duplicate Key Error");
        problemDetail.setInstance(URI.create(request.getContextPath()));
        problemDetail.setProperty("timestamp", Instant.now()); // adding more data using the Map properties of the
                                                               // ProblemDetail
        return ResponseEntity.status(HttpStatus.CONFLICT)
                .contentType(MediaType.APPLICATION_PROBLEM_JSON)
                .body(problemDetail);
    }

    @ExceptionHandler(DataIntegrityViolationException.class)
    ResponseEntity<ProblemDetail> handleDataIntegrityViolationException(DataIntegrityViolationException ex,
            WebRequest request) {
        ProblemDetail problemDetail = ProblemDetail.forStatusAndDetail(HttpStatus.UNPROCESSABLE_ENTITY,
                ex.getMessage());
        problemDetail.setTitle("Data Integrity Violation");
        problemDetail.setInstance(URI.create(request.getContextPath()));
        problemDetail.setProperty("timestamp", Instant.now()); // adding more data using the Map properties of the
                                                               // ProblemDetail
        return ResponseEntity.status(HttpStatus.UNPROCESSABLE_ENTITY)
                .contentType(MediaType.APPLICATION_PROBLEM_JSON)
                .body(problemDetail);
    }

    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<Object> handleMethodArgumentNotValidException(MethodArgumentNotValidException ex,
            WebRequest request) {
        ProblemDetail problemDetail = ProblemDetail.forStatusAndDetail(HttpStatus.BAD_REQUEST,
                handleMethodArgumentNotValidMessage(ex));
        problemDetail.setTitle("Bad request");
        problemDetail.setInstance(URI.create(request.getContextPath()));
        problemDetail.setProperty("timestamp", Instant.now()); // adding more data using the Map properties of the
                                                               // ProblemDetail
        return ResponseEntity.status(HttpStatus.BAD_REQUEST).body(problemDetail);
    }

    protected String handleMethodArgumentNotValidMessage(MethodArgumentNotValidException ex) {
        return Optional.ofNullable(ex.getBindingResult().getFieldError())
                .map(FieldError::getDefaultMessage)
                .orElse("Invalid request");
    }
}
