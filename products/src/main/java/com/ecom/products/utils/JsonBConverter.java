package com.ecom.products.utils;

import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.ObjectMapper;
import jakarta.persistence.AttributeConverter;
import jakarta.persistence.Converter;
import org.postgresql.util.PGobject;

import java.io.IOException;
import java.sql.SQLException;
import java.util.Map;

@Converter(autoApply = false)
public class JsonBConverter implements AttributeConverter<Map<String, String>, PGobject> {
    private final ObjectMapper objectMapper = new ObjectMapper();

    @Override
    public PGobject convertToDatabaseColumn(Map<String, String> attribute) {
        if (attribute == null) {
            return null;
        }
        try {
            PGobject pgObject = new PGobject();
            pgObject.setType("jsonb");
            pgObject.setValue(objectMapper.writeValueAsString(attribute));
            return pgObject;
        } catch (JsonProcessingException | SQLException e) {
            throw new IllegalArgumentException("Failed to convert to JSONB", e);
        }
    }

    @Override
    public Map<String, String> convertToEntityAttribute(PGobject dbData) {
        if (dbData == null || dbData.getValue() == null) {
            return null;
        }
        try {
            return objectMapper.readValue(dbData.getValue(), new TypeReference<>() {
            });
        } catch (IOException e) {
            throw new IllegalArgumentException("Failed to deserialize JSONB", e);
        }
    }
}
