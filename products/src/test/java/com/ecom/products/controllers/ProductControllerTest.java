package com.ecom.products.controllers;

import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.status;

import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.TestInfo;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.test.web.servlet.MockMvc;

import com.ecom.products.services.ProductService;
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders;
import org.springframework.test.web.servlet.result.MockMvcResultMatchers;
@SpringBootTest
@AutoConfigureMockMvc
class ProductControllerTest {

    @Autowired
    private MockMvc mockMvc;

    @Autowired
    private ProductService service;

    @Test
    void getAllProducts_returnsListOfProducts(TestInfo testInfo) throws Exception {
        mockMvc.perform(MockMvcRequestBuilders.get("/api/products"))
            .andExpect(status().isOk())
            .andExpect(MockMvcResultMatchers.jsonPath("content").isArray());
        System.out.println("================Test: " + testInfo.getDisplayName() + " executed successfully.===============");
    }
}
