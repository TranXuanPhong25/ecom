## [3.5.0](https://github.com/TranXuanPhong25/ecom/compare/v3.4.0...v3.5.0) (2025-08-24)

### 🚀 Features

* upload container pull image from hub ([f556fc1](https://github.com/TranXuanPhong25/ecom/commit/f556fc1125efdc10fbb1dc15b61db0fd3dc48d35))
* upload container pull image from hub ([b1ed60b](https://github.com/TranXuanPhong25/ecom/commit/b1ed60b3b990955373cbaf1ce81002774ea4ca44))

## [3.4.0](https://github.com/TranXuanPhong25/ecom/compare/v3.3.0...v3.4.0) (2025-08-08)

### 🚀 Features

* add some default brands ([307cd25](https://github.com/TranXuanPhong25/ecom/commit/307cd2556055d055110d346000b85d599b4f7973))

## [3.3.0](https://github.com/TranXuanPhong25/ecom/compare/v3.2.0...v3.3.0) (2025-08-08)

### 🚀 Features

* add some default brands ([9e871c3](https://github.com/TranXuanPhong25/ecom/commit/9e871c357730d334590427961c1df2699e7a6068))

## [3.2.0](https://github.com/TranXuanPhong25/ecom/compare/v3.1.2...v3.2.0) (2025-08-08)

### 🚀 Features

* cover image for product ([c0e4854](https://github.com/TranXuanPhong25/ecom/commit/c0e4854a31e219cacf448df2bc6f8692aa138674))

## [3.1.2](https://github.com/TranXuanPhong25/ecom/compare/v3.1.1...v3.1.2) (2025-08-08)

### 🐛 Bug Fixes

* product creation and modification tuned ([acf56fd](https://github.com/TranXuanPhong25/ecom/commit/acf56fda1bd6d55e0c69065ca339be11f8674d00))

## [3.1.1](https://github.com/TranXuanPhong25/ecom/compare/v3.1.0...v3.1.1) (2025-08-07)

### 🐛 Bug Fixes

* auto fetch brand when get product ([c79a4e5](https://github.com/TranXuanPhong25/ecom/commit/c79a4e5918a55e5eb61ec44341f2189b3a74f6ee))

## [3.1.0](https://github.com/TranXuanPhong25/ecom/compare/v3.0.0...v3.1.0) (2025-08-07)

### 🚀 Features

* **product:** create product-categories client and make a REST call when get categoryPath ([cb853a3](https://github.com/TranXuanPhong25/ecom/commit/cb853a3270c19bcd19440a83135329af794a9038))
* **product:** new endpoint for delete multiple products ([7ce0e63](https://github.com/TranXuanPhong25/ecom/commit/7ce0e63990a1cc1a09e6cc436933248b5824ac6d))

## [3.0.0](https://github.com/TranXuanPhong25/ecom/compare/v2.3.1...v3.0.0) (2025-08-06)

### ⚠ BREAKING CHANGES

* **products:** - remove variant_images
- add images field to products and product_variants schema
- insert default branch on flyway migration

### 🚀 Features

* **products:** Change Product and Variant schema ([cc1544f](https://github.com/TranXuanPhong25/ecom/commit/cc1544f0d7ead22890749a48b299e8a6b8fc9903))

## [2.3.1](https://github.com/TranXuanPhong25/ecom/compare/v2.3.0...v2.3.1) (2025-08-06)

### 🐛 Bug Fixes

* kong relevant port use ([39bf979](https://github.com/TranXuanPhong25/ecom/commit/39bf979185bdba2b95022a7c83252f384c156a37))

### 🔨 Technical Changes

* don't expose port of each service, set all service to open at default port ([4e4608e](https://github.com/TranXuanPhong25/ecom/commit/4e4608e69be3056e6830438a5e500287ce9d0269))

## [2.3.0](https://github.com/TranXuanPhong25/ecom/compare/v2.2.0...v2.3.0) (2025-08-06)

### 🚀 Features

* **minio:** use nginx as minio-proxy ([9bc2a1e](https://github.com/TranXuanPhong25/ecom/commit/9bc2a1e8edcd55454c2a36652eb44430246c3529))

## [2.2.0](https://github.com/TranXuanPhong25/ecom/compare/v2.1.0...v2.2.0) (2025-08-05)

### 🚀 Features

* **upload:** basic presign url logic for image ([4c08482](https://github.com/TranXuanPhong25/ecom/commit/4c08482172e25cf6a57880e81d5c4add9438a791))

### 🔨 Technical Changes

* **upload-service:** change from image-upload service to uploadservice ([4cf4b32](https://github.com/TranXuanPhong25/ecom/commit/4cf4b325d4ed6ecbb15147c75b0f7acd2e6fa008))

## [2.1.0](https://github.com/TranXuanPhong25/ecom/compare/v2.0.0...v2.1.0) (2025-08-05)

### 🚀 Features

* **product:** add more details to get product response ([2a333f1](https://github.com/TranXuanPhong25/ecom/commit/2a333f1dae7d0a277a09d44db544f2efc1f02ae1))

## [2.0.0](https://github.com/TranXuanPhong25/ecom/compare/v1.7.0...v2.0.0) (2025-08-02)

### ⚠ BREAKING CHANGES

* **product:** remove product_variant_skus then add sku to column definitions of variants table

### 🚀 Features

* **product:** change database schema ([7bfd004](https://github.com/TranXuanPhong25/ecom/commit/7bfd004112d568022c32f683503a058b4212843c))

## [1.7.0](https://github.com/TranXuanPhong25/ecom/compare/v1.6.0...v1.7.0) (2025-08-02)

### 🚀 Features

* **metrics:** integrate prometheus to product-service ([6aa37dd](https://github.com/TranXuanPhong25/ecom/commit/6aa37dd50d9cdd151339565cdbc23b3b33b0e9de))

## [1.6.0](https://github.com/TranXuanPhong25/ecom/compare/v1.5.1...v1.6.0) (2025-08-01)

### 🚀 Features

* **product:** jsonb for specs and enum ProductStatus ([f5c79d3](https://github.com/TranXuanPhong25/ecom/commit/f5c79d3b45fabe5fc642ad36880a18f2ff0798c8))

## [1.5.1](https://github.com/TranXuanPhong25/ecom/compare/v1.5.0...v1.5.1) (2025-07-23)

### 🐛 Bug Fixes

* **eas:** refine opa query ([bf3d9cd](https://github.com/TranXuanPhong25/ecom/commit/bf3d9cd85b63ea545f8ba8e9a93a96f0305a1085))
* **products:** change timezone ([f908a5a](https://github.com/TranXuanPhong25/ecom/commit/f908a5af41c8c148cc3d4ab7187f69d3054870c9))

## [1.5.0](https://github.com/TranXuanPhong25/ecom/compare/v1.4.0...v1.5.0) (2025-07-23)

### 🚀 Features

* **product-categories:** new endpoint for find category path ([64ec266](https://github.com/TranXuanPhong25/ecom/commit/64ec2668c50a912f0e235e64889241ab37c58320))

## [1.4.0](https://github.com/TranXuanPhong25/ecom/compare/v1.3.0...v1.4.0) (2025-07-23)

### 🚀 Features

* **shops:** create and get single shop ([698ca95](https://github.com/TranXuanPhong25/ecom/commit/698ca9539b1e28880f6314ba34c37f5dbb364c3c))

## [1.3.0](https://github.com/TranXuanPhong25/ecom/compare/v1.2.0...v1.3.0) (2025-07-23)

### 🚀 Features

* **shops:** init project ([20d8a8d](https://github.com/TranXuanPhong25/ecom/commit/20d8a8d82baac3cf5128d660e3a088f4d18c64a4))

## [1.2.0](https://github.com/TranXuanPhong25/ecom/compare/v1.1.0...v1.2.0) (2025-07-23)

### 🚀 Features

* **product-service:** add shopId field to table and model ([5b648a7](https://github.com/TranXuanPhong25/ecom/commit/5b648a787840f34afc214c344c6bda2de9d5bbd7))

## [1.1.0](https://github.com/TranXuanPhong25/ecom/compare/v1.0.0...v1.1.0) (2025-07-22)

### 🚀 Features

* **eas:** use Open Policy Agent for validate endpoint access ([785ea26](https://github.com/TranXuanPhong25/ecom/commit/785ea26efcfe4e67239d06e7c8de1030cefcdbf6))

### 🐛 Bug Fixes

* correction use of @RequiredArgConstructor ([33b328d](https://github.com/TranXuanPhong25/ecom/commit/33b328d28975ddf2d5d85b0784c6427543e29723))

## [1.0.0](https://github.com/TranXuanPhong25/ecom/compare/v0.7.0...v1.0.0) (2025-07-22)

### ⚠ BREAKING CHANGES

* use materialized path for categories
* create product-categories logic changed

### 🚀 Features

* create product-categories logic changed ([28879ad](https://github.com/TranXuanPhong25/ecom/commit/28879ad3957bcde5a87984a06686342eebbe3ca9))
* product-categories v2 ([256a8ee](https://github.com/TranXuanPhong25/ecom/commit/256a8ee6ae310a9394de0a75091ecafe5ba36459))

## [0.7.0](https://github.com/TranXuanPhong25/ecom/compare/v0.6.0...v0.7.0) (2025-07-22)

### 🚀 Features

* **authZ:** use opa as authZ server ([80d078d](https://github.com/TranXuanPhong25/ecom/commit/80d078dca6d736cc35b44a6f6ce4293616aac3d4))
* integrated create product ([3100649](https://github.com/TranXuanPhong25/ecom/commit/3100649a3065c88814902b678063e22d57983b31))

## [0.6.0](https://github.com/TranXuanPhong25/ecom/compare/v0.5.0...v0.6.0) (2025-07-20)

### 🚀 Features

* podman-compose.yml for podman ([e0fce96](https://github.com/TranXuanPhong25/ecom/commit/e0fce9696588ae153be2de3d83d604cd84f90c3f))

### 🔨 Technical Changes

* remove validators ([ea7d04e](https://github.com/TranXuanPhong25/ecom/commit/ea7d04e25ac7d9d5642c1bde704e251a0731f95e))

## [0.5.0](https://github.com/TranXuanPhong25/ecom/compare/v0.4.0...v0.5.0) (2025-07-20)

### 🚀 Features

* change back main as the release branch ([0bef7e7](https://github.com/TranXuanPhong25/ecom/commit/0bef7e7784449dabd67911a9c89c864cabeafe49))
* product service ([91f8eb8](https://github.com/TranXuanPhong25/ecom/commit/91f8eb85e554372a03becca77954ee0b78b5c054))

### 🐛 Bug Fixes

* change trigger branch in semantic-release.yml ([3651f07](https://github.com/TranXuanPhong25/ecom/commit/3651f074fd5986dd701663c51376091423a30534))

## [0.4.0](https://github.com/TranXuanPhong25/ecom/compare/v0.3.4...v0.4.0) (2025-07-20)

### ⚠ BREAKING CHANGES

* This is an template
* This is an template
* ...
* remove git-tag

### 🚀 Features

* ... ([27bb853](https://github.com/TranXuanPhong25/ecom/commit/27bb853d30711aa214a7bac7ed105e5ccf708f7a))
* This is an template ([0c3768d](https://github.com/TranXuanPhong25/ecom/commit/0c3768d587a3a168813d3c55f694b8bd17ffb16e))
* This is an template ([4eaf1ee](https://github.com/TranXuanPhong25/ecom/commit/4eaf1ee6de333170e41770c863f59e5504ac474a))
* This is an template ([7ac4455](https://github.com/TranXuanPhong25/ecom/commit/7ac44550cd3a8b60b26d70ffe1d78509c85db981))
* trigger release ([b5b3262](https://github.com/TranXuanPhong25/ecom/commit/b5b326255ac528743a48f5ecbeac57d50ed8198d))

### 🐛 Bug Fixes

* remove git-tag ([e4acf98](https://github.com/TranXuanPhong25/ecom/commit/e4acf98bc5ee21d1cec9c88cdbed9caa3cf0c1fa))
