# Shop Creation Step Validation

This document describes the step-by-step validation feature for shop creation.

## Overview

The shop creation process is divided into multiple steps, and each step requires certain fields to be filled before moving to the next step. This feature provides an API endpoint to validate the required fields for each step.

## API Endpoint

**POST** `/api/shops/validate-step`

### Request Body

```json
{
  "step": 1,
  "name": "My Shop",
  "ownerId": "550e8400-e29b-41d4-a716-446655440000",
  "location": "123 Main St",
  "logo": "https://example.com/logo.png",
  "banner": "https://example.com/banner.png",
  "email": "test@example.com",
  "phone": "1234567890",
  "businessType": "individual"
}
```

### Step Validation Rules

#### Step 1: Basic Information
Required fields:
- `name` - Shop name (must be provided)
- `businessType` - Type of business (e.g., "individual" or "business")

#### Step 2: Contact Information
Required fields (includes Step 1 fields):
- All fields from Step 1
- `location` - Physical address of the shop
- `email` - Valid email address

#### Step 3: Branding (Optional)
Required fields (includes Step 1 and 2 fields):
- All fields from Steps 1 and 2
- `logo` - URL to the shop logo (optional but validated if provided)
- `banner` - URL to the shop banner (optional but validated if provided)

### Response

#### Success Response (200 OK)
```json
{
  "status": "200",
  "message": "Validation successful",
  "step": 1
}
```

#### Error Response (400 Bad Request)
```json
{
  "status": "400",
  "detail": "required fields missing: name, businessType"
}
```

## Usage Example

### Frontend Implementation

```javascript
// Example: Validate step 1 before moving to step 2
async function validateStep(stepNumber, formData) {
  const response = await fetch('/api/shops/validate-step', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      step: stepNumber,
      ...formData
    })
  });

  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.detail);
  }

  return await response.json();
}

// Usage in a form handler
async function handleNextStep(currentStep, formData) {
  try {
    await validateStep(currentStep, formData);
    // Move to next step
    goToStep(currentStep + 1);
  } catch (error) {
    // Show error message to user
    showError(error.message);
  }
}
```

## Validation Logic

The validation is cumulative, meaning:
- Step 2 validation checks all fields from Step 1 + Step 2
- Step 3 validation checks all fields from Steps 1, 2 + Step 3

This ensures that users cannot skip steps and that all previous data is valid before proceeding.

## Field Validation Rules

- **name**: Required, 3-100 characters
- **ownerId**: Valid UUID format
- **location**: Required, 3-255 characters
- **email**: Required, valid email format
- **phone**: Optional, numeric only
- **businessType**: Required, string value
- **logo**: Optional, valid URL format
- **banner**: Optional, valid URL format
