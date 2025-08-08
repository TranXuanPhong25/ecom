package com.ecom.products.services;

import com.ecom.products.dtos.ProductDTO;
import com.ecom.products.dtos.VariantDTO;
import com.ecom.products.entities.Product;
import com.ecom.products.entities.ProductVariant;
import com.ecom.products.models.CreateProductRequest;
import com.ecom.products.repositories.ProductRepository;
import jakarta.transaction.Transactional;
import lombok.RequiredArgsConstructor;
import org.hibernate.exception.ConstraintViolationException;
import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.dao.DuplicateKeyException;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;
import reactor.core.publisher.Mono;

import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

@Service
@RequiredArgsConstructor
public class ProductService {
    private final ProductVariantService productVariantService;
    private final ProductRepository productRepository;
    private final BrandService brandService;
    private final CategoryService categoryService;

    @Transactional
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
        List<VariantDTO> variantDTOs = new ArrayList<>();
        Product finalProduct = product;
        createProductRequest.variants().forEach(variant -> {
            ProductVariant newVariant = productVariantService.toEntity(variant);
            newVariant.setProductId(finalProduct.getId());
            newVariant = productVariantService.createVariant(newVariant);
            VariantDTO variantDTO = productVariantService.toDTO(newVariant);
            variantDTOs.add(variantDTO);
        });
        return toDTO(product, variantDTOs);
    }

    public Page<ProductDTO> getProducts(Pageable pageable) {
        return productRepository.findAll(pageable).map(this::toDTO);
    }

    public ProductDTO getProductById(Long id) {
        Product product = productRepository.findById(id).orElse(null);
        if (product == null) {
            return null;
        }

        List<VariantDTO> variantDTOs = productVariantService.getVariantsByProductId(id);
        ProductDTO productDTO = toDTO(product, variantDTOs);
        if (product.getCategoryId() != null) {
            Mono<String> categoryPathMono = categoryService.getCategoryPath(product.getCategoryId());
            productDTO.setCategoryPath(categoryPathMono.block());
        }
        return productDTO;

    }

    public ProductDTO updateProduct(Long id, CreateProductRequest updateProductRequest) {
        ProductDTO productDTO = updateProductRequest.product();
        ProductDTO updatedProduct =  productRepository.findById(id).map(product -> {
            product.setName(productDTO.getName());
            product.setDescription(productDTO.getDescription());
            product.setCategoryId(productDTO.getCategoryId());
            product.setImages(productDTO.getImages());
            product.setBrand(brandService.toEntity(productDTO.getBrand()));
            product.setStatus(productDTO.getStatus());
            product = productRepository.save(product);
            return toDTO(product);
        }).orElse(null);

        // Update variants
        if (updatedProduct != null) {
            updateProductRequest.variants().forEach(variant -> {
                productVariantService.updateVariant(variant.getId(), variant);
            });
        }
        return updatedProduct;
    }

    public void deleteProduct(Long id) {
        productRepository.deleteById(id);
    }

    private ProductDTO toDTO(Product product) {
        ProductDTO dto = new ProductDTO();
        dto.setShopId(product.getShopId().toString());
        dto.setId(product.getId());
        dto.setName(product.getName());
        dto.setCoverImage(product.getCoverImage());
        dto.setDescription(product.getDescription());
        dto.setCategoryId(product.getCategoryId());
        dto.setSpecs(product.getSpecs());
        if (product.getBrand() != null) {
            dto.setBrand(brandService.toDTO(product.getBrand()));
        }
        dto.setImages(product.getImages());
        dto.setStatus(product.getStatus());
        dto.setCreatedAt(product.getCreatedAt());
        dto.setUpdatedAt(product.getUpdatedAt());
        return dto;
    }

    private ProductDTO toDTO(Product product, List<VariantDTO> variantDTOs) {
        ProductDTO dto = toDTO(product);
        dto.setVariants(variantDTOs);
        return dto;
    }

    private Product toEntity(ProductDTO dto) {

        Product product = new Product();
        product.setShopId(UUID.fromString(dto.getShopId()));
        product.setId(dto.getId());
        product.setCoverImage(dto.getCoverImage());
        product.setSpecs(dto.getSpecs());
        product.setName(dto.getName());
        product.setImages(dto.getImages());
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

    public void deleteProductsByIds(List<Long> ids) {
        if (ids == null || ids.isEmpty()) {
            return;
        }
        try {
            productRepository.deleteAllById(ids);
        } catch (DataIntegrityViolationException ex) {
            throw new RuntimeException("Error occur when deleting products: " + ex.getMessage(), ex);
        }
    }
}