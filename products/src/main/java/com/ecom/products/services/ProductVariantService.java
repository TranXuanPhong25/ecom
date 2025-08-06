package com.ecom.products.services;

import com.ecom.products.dtos.VariantDTO;
import com.ecom.products.entities.ProductVariant;
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

    public ProductVariant createVariant(ProductVariant variant) {
        return variantRepository.save(variant);
    }

    public List<VariantDTO> getVariantsByProductId(Long productId) {
        List<VariantDTO> variantDTOList = variantRepository.findByProductId(productId)
                .stream()
                .map(this::toDTO)
                .collect(Collectors.toList());
        return variantDTOList;
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
        dto.setPrice(variant.getPrice());
        dto.setStockQuantity(variant.getStockQuantity());
        dto.setAttributes(variant.getAttributes());
        dto.setActive(variant.isActive());
        dto.setSku(variant.getSku());
        dto.setImages(variant.getImages());
        return dto;
    }

    ProductVariant toEntity(VariantDTO dto) {
        ProductVariant variant = new ProductVariant();
        variant.setId(dto.getId());
        variant.setImages(dto.getImages());
        variant.setPrice(dto.getPrice());
        variant.setStockQuantity(dto.getStockQuantity());
        variant.setAttributes(dto.getAttributes());
        variant.setActive(dto.isActive());
        variant.setSku(dto.getSku());
        return variant;
    }
}