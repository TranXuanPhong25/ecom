package com.ecom.productcategory.services;

import com.ecom.productcategory.dto.ProductCategoryNodeDTO;
import com.ecom.productcategory.entities.ProductCategoryEntity;
import com.ecom.productcategory.repositories.ProductCategoryRepository;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;

import java.util.*;

import static org.junit.jupiter.api.Assertions.*;
import static org.mockito.Mockito.*;

@ExtendWith(MockitoExtension.class)
class ProductCategoryServiceImplTest {

    @Mock
    private ProductCategoryRepository productCategoryRepository;

    @Mock
    private ProductCategoryClosureService productCategoryClosureService;

    @InjectMocks
    private ProductCategoryServiceImpl productCategoryService;

    private List<ProductCategoryEntity> mockCategories;
    private List<Map<String, Object>> mockRelationships;

    @BeforeEach
    void setUp() {
        // Create mock category entities
        mockCategories = Arrays.asList(
            new ProductCategoryEntity(1, "Electronics", "electronics.jpg"),
            new ProductCategoryEntity(2, "Computers", "computers.jpg"),
            new ProductCategoryEntity(3, "Laptops", "laptops.jpg"),
            new ProductCategoryEntity(4, "Desktops", "desktops.jpg"),
            new ProductCategoryEntity(5, "Phones", "phones.jpg")
        );

        // Create mock relationships
        // Electronics (1) -> Computers (2)
        // Electronics (1) -> Phones (5)
        // Computers (2) -> Laptops (3)
        // Computers (2) -> Desktops (4)
        mockRelationships = new ArrayList<>();
        mockRelationships.add(createRelationship(1, 2));
        mockRelationships.add(createRelationship(1, 5));
        mockRelationships.add(createRelationship(2, 3));
        mockRelationships.add(createRelationship(2, 4));
    }

    private Map<String, Object> createRelationship(Integer parentId, Integer childId) {
        Map<String, Object> rel = new HashMap<>();
        rel.put("parentId", parentId);
        rel.put("childId", childId);
        return rel;
    }

    @Test
    void testGetProductCategoriesTree_BuildsCorrectHierarchy() {
        // Arrange
        when(productCategoryRepository.findAll()).thenReturn(mockCategories);
        when(productCategoryRepository.findAllDirectRelationships()).thenReturn(mockRelationships);

        // Act
        List<ProductCategoryNodeDTO> result = productCategoryService.getProductCategoriesTree();

        // Assert
        assertNotNull(result);
        assertEquals(1, result.size(), "Should have only one root category");
        
        ProductCategoryNodeDTO root = result.get(0);
        assertEquals(1, root.getId(), "Root should be Electronics");
        assertEquals("Electronics", root.getName());
        assertEquals(2, root.getChildren().size(), "Electronics should have 2 children");

        // Verify the tree structure
        ProductCategoryNodeDTO computers = root.getChildren().stream()
            .filter(c -> c.getId().equals(2))
            .findFirst()
            .orElse(null);
        assertNotNull(computers, "Computers should be a child of Electronics");
        assertEquals(2, computers.getChildren().size(), "Computers should have 2 children");

        ProductCategoryNodeDTO phones = root.getChildren().stream()
            .filter(c -> c.getId().equals(5))
            .findFirst()
            .orElse(null);
        assertNotNull(phones, "Phones should be a child of Electronics");
        assertEquals(0, phones.getChildren().size(), "Phones should have no children");

        // Verify only 2 queries were made (optimized)
        verify(productCategoryRepository, times(1)).findAll();
        verify(productCategoryRepository, times(1)).findAllDirectRelationships();
    }

    @Test
    void testGetProductCategoriesTree_EmptyCategories() {
        // Arrange
        when(productCategoryRepository.findAll()).thenReturn(Collections.emptyList());
        when(productCategoryRepository.findAllDirectRelationships()).thenReturn(Collections.emptyList());

        // Act
        List<ProductCategoryNodeDTO> result = productCategoryService.getProductCategoriesTree();

        // Assert
        assertNotNull(result);
        assertTrue(result.isEmpty(), "Should return empty list for no categories");
        verify(productCategoryRepository, times(1)).findAll();
        verify(productCategoryRepository, times(1)).findAllDirectRelationships();
    }

    @Test
    void testGetALlRootProductCategories() {
        // Arrange
        List<ProductCategoryEntity> rootCategories = Arrays.asList(
            new ProductCategoryEntity(1, "Electronics", "electronics.jpg"),
            new ProductCategoryEntity(6, "Clothing", "clothing.jpg")
        );
        when(productCategoryRepository.findAllRootCategories()).thenReturn(rootCategories);

        // Act
        List<ProductCategoryEntity> result = productCategoryService.getALlRootProductCategories();

        // Assert
        assertNotNull(result);
        assertEquals(2, result.size());
        verify(productCategoryRepository, times(1)).findAllRootCategories();
    }
}
