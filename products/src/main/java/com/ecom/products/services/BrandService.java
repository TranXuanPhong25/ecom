package com.ecom.products.services;

import com.ecom.products.dtos.BrandDTO;
import com.ecom.products.entities.Brand;
import com.ecom.products.repositories.BrandRepository;
import jakarta.persistence.EntityNotFoundException;
import lombok.RequiredArgsConstructor;
import org.hibernate.exception.ConstraintViolationException;
import org.springframework.dao.DataIntegrityViolationException;
import org.springframework.dao.DuplicateKeyException;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
@RequiredArgsConstructor
public class BrandService {
    private final BrandRepository brandRepository;

    public BrandDTO createBrand(BrandDTO brand) {
        try {
            return toDTO(brandRepository.save(toEntity(brand)));
        } catch (DataIntegrityViolationException ex) {
            if (ex.getCause() instanceof ConstraintViolationException cve) {
                if (cve.getMessage().contains("brands_name_key")) {
                    throw new DuplicateKeyException("Brand with the same name already exists.");
                }
            }
            throw new RuntimeException("Error occur when creating new brand: " + ex.getMessage(), ex);
        }
    }

    public List<BrandDTO> getBrands() {
        return brandRepository.findAll().stream().map(this::toDTO).toList();
    }

    public BrandDTO updateBrand(Long id, BrandDTO brand) {
        if (brandRepository.existsById(id)) {
            brand.setId(id);
            return toDTO(brandRepository.save(toEntity(brand)));
        }
        return null;
    }

    public void deleteBrand(Long id) {
        if (brandRepository.existsById(id)) {
            brandRepository.deleteById(id);
        } else {
            throw new EntityNotFoundException("Brand not found");
        }
    }

    public BrandDTO toDTO(Brand brand) {
        return new BrandDTO(
                brand.getId(),
                brand.getName(),
                brand.getDescription()
        );
    }

    public Brand toEntity(BrandDTO brandDTO) {
        Brand brand = new Brand();
        brand.setId(brandDTO.getId());
        brand.setName(brandDTO.getName());
        brand.setDescription(brandDTO.getDescription());
        return brand;
    }
}
