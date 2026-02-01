package com.ecom.products.services;

import com.ecom.products.dtos.VariantDTO;
import com.ecom.products.dtos.VariantWithNameDTO;
import com.ecom.products.entities.ProductVariant;
import com.ecom.products.models.GetProductVariantsResponse;
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

    public GetProductVariantsResponse getVariantByIds(List<Long> ids) {
        List<ProductVariant> foundVariants = variantRepository.findAllByIdWithProduct(ids);
        System.out.println(foundVariants);
        List<VariantWithNameDTO> variantDTOs = foundVariants.stream().map(this::toVariantWithNameDTO)
                .collect(Collectors.toList());
        List<Long> foundIds = foundVariants.stream().map(ProductVariant::getId).collect(Collectors.toList());
        List<Long> notFoundIds = ids.stream().filter(id -> !foundIds.contains(id)).collect(Collectors.toList());
        return new GetProductVariantsResponse(variantDTOs, notFoundIds);
    }

    @Transactional
    public VariantDTO updateVariant(Long id, VariantDTO variantDTO) {
        return variantRepository.findById(id).map(variant -> {
            variant.setOriginalPrice(variantDTO.getOriginalPrice());
            variant.setSalePrice(variantDTO.getSalePrice());
            variant.setStockQuantity(variantDTO.getStockQuantity());
            variant.setAttributes(variantDTO.getAttributes());
            variant.setImages(variantDTO.getImages());
            variant.setActive(variantDTO.isActive());
            variant = variantRepository.save(variant);
            return toDTO(variant);
        }).orElse(null);
    }

    public void deleteVariant(Long id) {
        variantRepository.deleteById(id);
    }

    VariantDTO toDTO(ProductVariant variant) {
        VariantDTO dto = new VariantDTO();
        dto.setId(variant.getId());
        dto.setOriginalPrice(variant.getOriginalPrice());
        dto.setSalePrice(variant.getSalePrice());
        dto.setStockQuantity(variant.getStockQuantity());
        dto.setAttributes(variant.getAttributes());
        dto.setActive(variant.isActive());
        dto.setSku(variant.getSku());
        dto.setImages(variant.getImages());
        // Set product name from the fetched product relationship

        return dto;
    }

    VariantWithNameDTO toVariantWithNameDTO(ProductVariant variant) {
        VariantWithNameDTO dto = new VariantWithNameDTO(
                variant.getId(),
                variant.getProduct() != null ? variant.getProduct().getName() : null,
                variant.getOriginalPrice(),
                variant.getSalePrice(),
                variant.getAttributes(),
                variant.isActive(),
                variant.getStockQuantity(),
                variant.getSku(),
                variant.getImages().toArray(new String[0]));
        return dto;
    }

    ProductVariant toEntity(VariantDTO dto) {
        ProductVariant variant = new ProductVariant();
        variant.setId(dto.getId());
        variant.setImages(dto.getImages());
        variant.setOriginalPrice(dto.getOriginalPrice());
        variant.setSalePrice(dto.getSalePrice());
        variant.setStockQuantity(dto.getStockQuantity());
        variant.setAttributes(dto.getAttributes());
        variant.setActive(dto.isActive());
        variant.setSku(dto.getSku());
        return variant;
    }
}