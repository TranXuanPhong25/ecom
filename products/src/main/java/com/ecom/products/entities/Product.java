package com.ecom.products.entities;

import com.ecom.products.enums.ProductStatus;
import jakarta.persistence.*;
import lombok.Data;
import org.hibernate.annotations.CreationTimestamp;
import org.hibernate.annotations.JdbcTypeCode;
import org.hibernate.annotations.UpdateTimestamp;
import org.hibernate.type.SqlTypes;

import java.time.ZonedDateTime;
import java.util.Map;
import java.util.UUID;

@Data
@Entity
@Table(name = "products")
public class Product {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    private UUID shopId;

    private String name;

    @JdbcTypeCode(SqlTypes.JSON)
    @Column(columnDefinition = "jsonb")
    private Map<String, String> specs;

    private String description;

    private Long categoryId;

    @OneToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "brand_id")
    private Brand brand;

    @JdbcTypeCode(SqlTypes.NAMED_ENUM)
    @Basic(optional = false)
    @Column(name = "status", columnDefinition = "product_status")
    private ProductStatus status;


    @CreationTimestamp
    private ZonedDateTime createdAt;

    @UpdateTimestamp
    private ZonedDateTime updatedAt;
}