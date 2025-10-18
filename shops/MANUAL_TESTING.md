# Manual Testing Guide for Step Validation

This guide demonstrates how to manually test the step validation endpoint.

## Prerequisites

- The shops service must be running
- You'll need a tool like `curl` or Postman to make HTTP requests

## Test Cases

### Test 1: Validate Step 1 with Valid Data

```bash
curl -X POST http://localhost:8080/api/shops/validate-step \
  -H "Content-Type: application/json" \
  -d '{
    "step": 1,
    "name": "My Shop",
    "businessType": "individual"
  }'
```

**Expected Response:**
```json
{
  "status": "200",
  "message": "Validation successful",
  "step": 1
}
```

### Test 2: Validate Step 1 with Missing Name

```bash
curl -X POST http://localhost:8080/api/shops/validate-step \
  -H "Content-Type: application/json" \
  -d '{
    "step": 1,
    "businessType": "individual"
  }'
```

**Expected Response:**
```json
{
  "status": "400",
  "detail": "required fields missing: name"
}
```

### Test 3: Validate Step 2 with Valid Data

```bash
curl -X POST http://localhost:8080/api/shops/validate-step \
  -H "Content-Type: application/json" \
  -d '{
    "step": 2,
    "name": "My Shop",
    "businessType": "individual",
    "location": "123 Main St, City, Country",
    "email": "shop@example.com"
  }'
```

**Expected Response:**
```json
{
  "status": "200",
  "message": "Validation successful",
  "step": 2
}
```

### Test 4: Validate Step 2 with Missing Step 1 Fields

```bash
curl -X POST http://localhost:8080/api/shops/validate-step \
  -H "Content-Type: application/json" \
  -d '{
    "step": 2,
    "location": "123 Main St",
    "email": "shop@example.com"
  }'
```

**Expected Response:**
```json
{
  "status": "400",
  "detail": "required fields missing: name, businessType"
}
```

### Test 5: Validate Step 2 with Invalid Email

```bash
curl -X POST http://localhost:8080/api/shops/validate-step \
  -H "Content-Type: application/json" \
  -d '{
    "step": 2,
    "name": "My Shop",
    "businessType": "individual",
    "location": "123 Main St",
    "email": "invalid-email"
  }'
```

**Expected Response:**
```json
{
  "status": "400",
  "detail": "invalid email format"
}
```

### Test 6: Validate Step 3 with Valid Data

```bash
curl -X POST http://localhost:8080/api/shops/validate-step \
  -H "Content-Type: application/json" \
  -d '{
    "step": 3,
    "name": "My Shop",
    "businessType": "individual",
    "location": "123 Main St",
    "email": "shop@example.com",
    "logo": "https://example.com/logo.png",
    "banner": "https://example.com/banner.png"
  }'
```

**Expected Response:**
```json
{
  "status": "200",
  "message": "Validation successful",
  "step": 3
}
```

### Test 7: Validate Step 3 with Invalid Logo URL

```bash
curl -X POST http://localhost:8080/api/shops/validate-step \
  -H "Content-Type: application/json" \
  -d '{
    "step": 3,
    "name": "My Shop",
    "businessType": "individual",
    "location": "123 Main St",
    "email": "shop@example.com",
    "logo": "not-a-valid-url",
    "banner": "https://example.com/banner.png"
  }'
```

**Expected Response:**
```json
{
  "status": "400",
  "detail": "invalid logo URL format"
}
```

## Notes

- The validation is cumulative: validating step 2 also validates all step 1 fields
- The validation is cumulative: validating step 3 also validates all step 1 and 2 fields
- You can skip optional fields in later steps (like logo and banner in step 3)
- Invalid step numbers (0, 4, etc.) will return an error
