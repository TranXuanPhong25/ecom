package com.ecom.products.services;

import com.ecom.products.dtos.VariantImageDTO;
import com.ecom.products.entities.VariantImage;
import com.ecom.products.repositories.VariantImageRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class VariantImageService {
    private final VariantImageRepository variantImageRepository;

    public VariantImage createVariantImage(VariantImage variantImage) {
        return variantImageRepository.save(variantImage);
    }

    public VariantImageDTO toDTO(VariantImage variantImage) {
         VariantImageDTO dto = new VariantImageDTO();
         dto.setImageUrl(variantImage.getImageUrl());
         dto.setPrimary(variantImage.isPrimary());
         return dto;
    }

    public VariantImage toEntity(VariantImageDTO dto) {
        VariantImage variantImage = new VariantImage();
        variantImage.setImageUrl(dto.getImageUrl());
        variantImage.setPrimary(dto.isPrimary());
        return variantImage;
    }
}
