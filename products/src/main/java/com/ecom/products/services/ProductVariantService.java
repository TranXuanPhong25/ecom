package com.ecom.products.services;

import com.ecom.products.dtos.VariantDTO;
import com.ecom.products.entities.ProductVariant;
import com.ecom.products.entities.ProductVariantSku;
import com.ecom.products.entities.VariantImage;
import com.ecom.products.repositories.ProductVariantRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;
import java.util.stream.Collectors;

@Service
@RequiredArgsConstructor
public class ProductVariantService {

    private final ProductVariantRepository variantRepository;
    private final ProductVariantSkuService productVariantSkuService;
    private final VariantImageService variantImageService;

    public VariantDTO createVariant(VariantDTO variantDTO) {
        ProductVariant variant = toEntity(variantDTO);
        variant = variantRepository.save(variant);
        Long variantId = variant.getId();

        if (variantDTO.getImages() != null) {
            variantDTO.getImages().forEach(imageDTO -> {
                VariantImage variantImage = variantImageService.toEntity(imageDTO);
                variantImage.setVariantId(variantId);
                variantImageService.createVariantImage(variantImage);
            });
        }

        if(variantDTO.getSku() != null) {
            ProductVariantSku productVariantSku = new ProductVariantSku();
            productVariantSku.setVariantId(variantId);
            productVariantSku.setSku(variantDTO.getSku());
            productVariantSkuService.createProductVariantSku(productVariantSku);
        }

        return toDTO(variant);
    }

    public List<VariantDTO> getVariantsByProductId(Long productId) {
        return variantRepository.findByProductId(productId).stream()
                .map(this::toDTO)
                .collect(Collectors.toList());
    }

    public VariantDTO getVariantById(Long id) {
        return variantRepository.findById(id).map(this::toDTO).orElse(null);
    }

    @Transactional
    public VariantDTO updateVariant(Long id, VariantDTO variantDTO) {
        return variantRepository.findById(id).map(variant -> {
            variant.setPrice(variantDTO.getPrice());
            variant.setStockQuantity(variantDTO.getStockQuantity());
            variant.setAttributes(variantDTO.getAttributes());
            variant.setActive(variantDTO.isActive());
            variant = variantRepository.save(variant);
            return toDTO(variant);
        }).orElse(null);
    }

    public void deleteVariant(Long id) {
        variantRepository.deleteById(id);
    }

    private VariantDTO toDTO(ProductVariant variant) {
        VariantDTO dto = new VariantDTO();
        dto.setId(variant.getId());
        dto.setProductId(variant.getProductId());
        dto.setPrice(variant.getPrice());
        dto.setStockQuantity(variant.getStockQuantity());
        dto.setAttributes(variant.getAttributes());
        dto.setActive(variant.isActive());
        return dto;
    }

    private ProductVariant toEntity(VariantDTO dto) {
        ProductVariant variant = new ProductVariant();
        variant.setId(dto.getId());
        variant.setProductId(dto.getProductId());
        variant.setPrice(dto.getPrice());
        variant.setStockQuantity(dto.getStockQuantity());
        variant.setAttributes(dto.getAttributes());
        variant.setActive(dto.isActive());
        return variant;
    }
}