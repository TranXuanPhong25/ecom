package com.ecom.products.services;

import com.ecom.products.dtos.VariantDTO;
import com.ecom.products.dtos.VariantImageDTO;
import com.ecom.products.entities.ProductVariant;
import com.ecom.products.entities.VariantImage;
import com.ecom.products.repositories.ProductVariantRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
@RequiredArgsConstructor
public class ProductVariantService {

    private final ProductVariantRepository variantRepository;
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

        return toDTO(variant);
    }

    public List<VariantDTO> getVariantsByProductId(Long productId) {
        List<VariantDTO> variantDTOList = variantRepository.findByProductId(productId);
        variantDTOList.forEach(variantDTO -> {
            List<VariantImageDTO> images = variantImageService.getImagesByVariantId(variantDTO.getId());
            variantDTO.setImages(images);
        });
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
        dto.setProductId(variant.getProductId());
        dto.setPrice(variant.getPrice());
        dto.setStockQuantity(variant.getStockQuantity());
        dto.setAttributes(variant.getAttributes());
        dto.setActive(variant.isActive());
        dto.setSku(variant.getSku());
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
        variant.setSku(dto.getSku());
        return variant;
    }
}