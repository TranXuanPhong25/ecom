package com.ecom.products.services;

import com.ecom.products.dtos.ProductDTO;
import com.ecom.products.entities.Product;
import com.ecom.products.models.CreateProductRequest;
import com.ecom.products.repositories.ProductRepository;
import lombok.RequiredArgsConstructor;
import org.hibernate.exception.ConstraintViolationException;
import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.dao.DuplicateKeyException;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;

import java.util.UUID;

@Service
@RequiredArgsConstructor
public class ProductService {
    private final ProductVariantService productVariantService;
    private final ProductRepository productRepository;
    private final BrandService brandService;

    public ProductDTO createProduct(CreateProductRequest createProductRequest) {
        Product product = toEntity(createProductRequest.product());
        try {
            product = productRepository.save(product);
        } catch (DataIntegrityViolationException ex) {
            if (ex.getCause() instanceof ConstraintViolationException cve) {
                if (cve.getMessage().contains("products_name_key")) {
                    throw new DuplicateKeyException("Product with the same name already exists.");
                }
            }
            throw new RuntimeException("Error occur when creating new product: " + ex.getMessage(), ex);
        }

        if (createProductRequest.variants() != null) {
            Product finalProduct = product;
            createProductRequest.variants().forEach(variant -> {
                variant.setProductId(finalProduct.getId());
                productVariantService.createVariant(variant);
            });
        }
        ;
        return toDTO(product);
    }

    public Page<ProductDTO> getProducts(Pageable pageable) {
        return productRepository.findAll(pageable).map(this::toDTO);
    }

    public ProductDTO getProductById(Long id) {
        return productRepository.findById(id).map(this::toDTO).orElse(null);
    }

    public ProductDTO updateProduct(Long id, ProductDTO productDTO) {
        return productRepository.findById(id).map(product -> {
            product.setName(productDTO.getName());
            product.setDescription(productDTO.getDescription());
            product.setCategoryId(productDTO.getCategoryId());
            if (product.getBrand() != null) {
                product.setBrand(brandService.toEntity(productDTO.getBrand()));
            }
            product.setStatus(productDTO.getStatus());
            product = productRepository.save(product);
            return toDTO(product);
        }).orElse(null);
    }

    public void deleteProduct(Long id) {
        productRepository.deleteById(id);
    }

    private ProductDTO toDTO(Product product) {
        ProductDTO dto = new ProductDTO();
        dto.setShopId(product.getShopId().toString());
        dto.setId(product.getId());
        dto.setName(product.getName());
        dto.setDescription(product.getDescription());
        dto.setCategoryId(product.getCategoryId());
        dto.setSpecs(product.getSpecs());
        if (product.getBrand() != null) {
            dto.setBrand(brandService.toDTO(product.getBrand()));
        }
        dto.setStatus(product.getStatus());
        dto.setCreatedAt(product.getCreatedAt());
        dto.setUpdatedAt(product.getUpdatedAt());
        return dto;
    }

    private Product toEntity(ProductDTO dto) {
        Product product = new Product();
        product.setShopId(UUID.fromString(dto.getShopId()));
        product.setId(dto.getId());
        product.setSpecs(dto.getSpecs());
        product.setName(dto.getName());
        product.setDescription(dto.getDescription());
        product.setCategoryId(dto.getCategoryId());
        if (dto.getBrand() != null) {
            product.setBrand(brandService.toEntity(dto.getBrand()));
        }
        product.setStatus(dto.getStatus());
        return product;
    }

    public Page<ProductDTO> getProductsByShopId(UUID shopId, Pageable pageable) {
        return productRepository.findByShopId(shopId, pageable).map(this::toDTO);
    }
}