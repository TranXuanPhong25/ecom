# eCommerce Microservices Platform

A modern, microservices-based eCommerce platform with frontend applications for customers, sellers, and administrators.

## Project Overview

This project implements a complete eCommerce ecosystem with multiple services and frontend applications:

### Backend Microservices

- **Product Service**: Manages product information and inventory
  - Endpoint: `http://localhost:8080`
  - Built with Spring Boot

- **Product Variance Service**: Handles product variants (sizes, colors, etc.)
  - Endpoint: `http://localhost:8081`
  - Built with Spring Boot

- **Product Category Service**: 
  - Uses Closure table pattern for efficient hierarchical category management
  - Endpoint: `GET /api/product-category/hierarchy`

- **Product Reviews Service**:
  - Handles customer reviews and ratings
  - Built with Go (Echo framework)
  - Endpoint: `http://localhost:8200`
  - Routes:
    - `GET /reviews`: Retrieve product reviews
    - `POST /reviews`: Create new reviews

### Frontend Applications

1. **Customer Storefront** ([Shopiew](https://github.com/TranXuanPhong25/shopiew)):
   - Built with Next.js 
   - Features:
     - Product discovery grid with filtering
     - Product detail pages with specifications
     - Customer reviews and ratings
     - Shopping cart functionality
     - Category-based browsing

2. **Seller Portal** ([Shopiew Seller](https://github.com/TranXuanPhong25/shopiew-seller)):
   - Built with Next.js
   - Allows sellers to manage products and inventory

3. **Admin Dashboard** ([Shopiew Admin](https://github.com/TranXuanPhong25/shopiew-admin)):
   - Built with Vue 3 + TypeScript + Vite
   - Provides administrative controls for the platform

## Getting Started

### Running the Microservices

Using Docker Compose:
```bash
docker-compose build
docker-compose up