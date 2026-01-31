## [3.23.0](https://github.com/TranXuanPhong25/ecom/compare/v3.22.0...v3.23.0) (2026-01-31)

### 🚀 Features

* **auth:** build scripts ([2e7090f](https://github.com/TranXuanPhong25/ecom/commit/2e7090fe51538807b6a6153cb66307da25833ec0))
* go build scripts ([ab6eb35](https://github.com/TranXuanPhong25/ecom/commit/ab6eb351091f2208b6b8c50f73d9497b83ac5903))
* **k8s:** expose kafka external, config debezium, orders routes ([331eb76](https://github.com/TranXuanPhong25/ecom/commit/331eb7638bf3308d5d51531aa684fe5e169a6ecd))
* **k8s:** order, order-placement manifest, k3s config ([31229d8](https://github.com/TranXuanPhong25/ecom/commit/31229d899798e187ef3430f80d2072e1411d759c))
* **order-placement:** init service ([46c86b4](https://github.com/TranXuanPhong25/ecom/commit/46c86b47674d07f26241780c062345a008873a34))
* **order-placement:** outbox ([ba170c2](https://github.com/TranXuanPhong25/ecom/commit/ba170c24abcd3e98ca1903afe6b6cb58f85bbb85))
* **orders:** hexagonal frame, event template ([6c17af6](https://github.com/TranXuanPhong25/ecom/commit/6c17af66b199a5a95d411de50fe66c1a71bc69be))
* **orders:** listen order-placement create event and init base apis ([d66e31d](https://github.com/TranXuanPhong25/ecom/commit/d66e31d9857244345e11bb524f3ded4e1d280a99))
* **recommendation:** init deps ([49b01e2](https://github.com/TranXuanPhong25/ecom/commit/49b01e24bc50397d06bcc996ad02e12aec864497))

### 🐛 Bug Fixes

* **k8s:** lb ([e7f70fe](https://github.com/TranXuanPhong25/ecom/commit/e7f70fe9964172d29a5d362033ad3267c2598b1f))
* **k8s:** pvc ([3896aa5](https://github.com/TranXuanPhong25/ecom/commit/3896aa5f6122da57cce97a144d083ab620e9ba0f))
* **k8s:** storages ([bd5a897](https://github.com/TranXuanPhong25/ecom/commit/bd5a89701950913a068a0b30fc9c1692d9ba55f1))
* upload service imports ([fade70d](https://github.com/TranXuanPhong25/ecom/commit/fade70df1d868d40ffb94300663a09f07a990969))

### 📝 Documentation

* **products:** meta-inf for config.yaml ([89a8b95](https://github.com/TranXuanPhong25/ecom/commit/89a8b95270eef00e82b1f9f4449fa61091733e07))

### 🔨 Technical Changes

* **jwt:** change expire time to 7 days ([f93ca9e](https://github.com/TranXuanPhong25/ecom/commit/f93ca9eed6628802e9697edda5dea002da5a767e))
* **order-placement:** more standart layout ([81c98b0](https://github.com/TranXuanPhong25/ecom/commit/81c98b0b7e37562ca3cec6e6a6ca9ec424cbf438))
* **order-placement:** remove outbox, sync with orders service ([f6fe9cb](https://github.com/TranXuanPhong25/ecom/commit/f6fe9cba68fcc2f28f893deb1644f8b540acdf28))

## [3.22.0](https://github.com/TranXuanPhong25/ecom/compare/v3.21.0...v3.22.0) (2026-01-07)

### 🚀 Features

* **k8s:** streamzi kafka and kafka-connect cdc installation ([43471ac](https://github.com/TranXuanPhong25/ecom/commit/43471acac591bc3ec1d600be9d6952de9ab3b46c))
* **orders:** init order service ([5b59958](https://github.com/TranXuanPhong25/ecom/commit/5b59958151d87a1a179ed156003d294c15aecb6c))
* **promo:** new ([bf24c51](https://github.com/TranXuanPhong25/ecom/commit/bf24c5104cf538ffe1571d41784d8e6fe59b7c85))
* **recommendation:** init service ([8674cad](https://github.com/TranXuanPhong25/ecom/commit/8674cad569fe59522089535d5fd5a1b832b05183))

### 🔨 Technical Changes

* **chatbot:** structure simplified ([cb99ebf](https://github.com/TranXuanPhong25/ecom/commit/cb99ebf23a704f6b01e006fd59d1bfc3b9b0abdc))
* **jwt:** long expire time for development convenience ([35dfd9a](https://github.com/TranXuanPhong25/ecom/commit/35dfd9a640b21de8a4bb648613671dff5435184c))
* **k8s:** correctly use StateFulSet for db ([8909700](https://github.com/TranXuanPhong25/ecom/commit/890970037c75e625888a09c1649ed52d5efe42c4))
* **promotion:** response format update ([e90cb57](https://github.com/TranXuanPhong25/ecom/commit/e90cb5791ee62ad9d810102d9fa93db36c5361f7))

## [3.21.0](https://github.com/TranXuanPhong25/ecom/compare/v3.20.1...v3.21.0) (2025-12-26)

### 🚀 Features

* **carts:** totalcount ([f9da8e1](https://github.com/TranXuanPhong25/ecom/commit/f9da8e1cf7024fa7f5de39deb7c8303f4a0d4323))

### 🐛 Bug Fixes

* cart params ([7c6ded3](https://github.com/TranXuanPhong25/ecom/commit/7c6ded3ac73657fb4f522e53beb733c5faaafabd))

### 🔨 Technical Changes

* **carts:** migrate to postgres ([e483de9](https://github.com/TranXuanPhong25/ecom/commit/e483de99619aa60178eea3c72d57cdbc261ca150))

## [3.20.1](https://github.com/TranXuanPhong25/ecom/compare/v3.20.0...v3.20.1) (2025-12-26)

### 🐛 Bug Fixes

* imports ([f5f92a0](https://github.com/TranXuanPhong25/ecom/commit/f5f92a0d4bcdefd227ef526112b47681d755deca))
* update cart imports ([c1e26d6](https://github.com/TranXuanPhong25/ecom/commit/c1e26d6c867154a156729a956ccba20ebcabc28b))

### 🔨 Technical Changes

* update k8s ([132644c](https://github.com/TranXuanPhong25/ecom/commit/132644c02acfd583bc419498a2a582adf7191833))

## [3.20.0](https://github.com/TranXuanPhong25/ecom/compare/v3.19.0...v3.20.0) (2025-12-26)

### 🚀 Features

* update go.mod to match with new location ([52af5fe](https://github.com/TranXuanPhong25/ecom/commit/52af5feca355c2550ff10677f83fd29181edbfe5))

## [3.19.0](https://github.com/TranXuanPhong25/ecom/compare/v3.18.0...v3.19.0) (2025-12-17)

### 🚀 Features

* **k8s:** update README.md ([fe50270](https://github.com/TranXuanPhong25/ecom/commit/fe50270f58b2a2dcf138d886c135daf73a52e1c9))

### 🔨 Technical Changes

* **carts:** remove replace module ([9e8a522](https://github.com/TranXuanPhong25/ecom/commit/9e8a522df57c8dc7f047f6bfc4eaa6c469782a08))
* move all services into services folder ([915fe35](https://github.com/TranXuanPhong25/ecom/commit/915fe357090a24adb69ad72ea6180de9164b863d))
* move all services into services folder ([48d7f75](https://github.com/TranXuanPhong25/ecom/commit/48d7f75d4d4fd70dd929bf5768cf9a6e82842cbb))

## [3.18.0](https://github.com/TranXuanPhong25/ecom/compare/v3.17.0...v3.18.0) (2025-10-04)

### 🚀 Features

* **opa:** rules for carts ([2d0716d](https://github.com/TranXuanPhong25/ecom/commit/2d0716d80b2eb60be9585c071e5929e6b21716ca))
* **shop:** shop proto ([8a28904](https://github.com/TranXuanPhong25/ecom/commit/8a289046863068cc8b1922d63227268f92db3b09))
* **shop:** shop proto ([3c1e56c](https://github.com/TranXuanPhong25/ecom/commit/3c1e56c0a5a00d64d94eaab89857f4a33252bf07))
* **shops:** rpc server ([6ba075a](https://github.com/TranXuanPhong25/ecom/commit/6ba075ab6a7ace3642550bb5828e7d6c94e92d06))

### 🐛 Bug Fixes

* **carts:** correct mapping response from product service ([3c55356](https://github.com/TranXuanPhong25/ecom/commit/3c55356021affcdd88964a3103bb8ae858a876da))

### 🔨 Technical Changes

* carts - shops logic, run server in goroutine to graceful shutdown ([68fd04a](https://github.com/TranXuanPhong25/ecom/commit/68fd04a344e0fd6e455f14327bb42987558c484e))
* **carts:** remove replace module ([734392e](https://github.com/TranXuanPhong25/ecom/commit/734392e0e4e3d5b6cfbb030bcc0d8f4421634b1b))

## [3.17.0](https://github.com/TranXuanPhong25/ecom/compare/v3.16.0...v3.17.0) (2025-10-03)

### 🚀 Features

* **carts:** connect with product-variants service ([c14ec1a](https://github.com/TranXuanPhong25/ecom/commit/c14ec1abbb1e1f5f3bb9783bcd056c48f807b1b3))
* **chatbots:** basic agent graph ([1b027f9](https://github.com/TranXuanPhong25/ecom/commit/1b027f98970c981000f6ddf71f104d404c4fa7a4))
* **products:** product-variants retrieval logic ([36c5724](https://github.com/TranXuanPhong25/ecom/commit/36c5724c507613ba8dd63b97a7ddc13526adc4b2))

### 🔨 Technical Changes

* change image pull rule to default(never) ([6b1fd90](https://github.com/TranXuanPhong25/ecom/commit/6b1fd905ed0dd8975bd05fb4cfe7a1c58595243f))
* suppress compiler warning ([6f8ea84](https://github.com/TranXuanPhong25/ecom/commit/6f8ea846f5277ddf4e711c46e745fe1f80e3f159))

## [3.16.0](https://github.com/TranXuanPhong25/ecom/compare/v3.15.1...v3.16.0) (2025-09-23)

### 🚀 Features

* cart service with scyllaBD and simple get cart endpoint ([5390fc2](https://github.com/TranXuanPhong25/ecom/commit/5390fc225f20c61377c7f0b5c0fe43ffccc9871b))
* **carts:** complete CRUD ([03ec710](https://github.com/TranXuanPhong25/ecom/commit/03ec710c52aca95358d16557a9033856eb54b24d))

### 🐛 Bug Fixes

* correct mapping response from CategoryService ([6e6c8ce](https://github.com/TranXuanPhong25/ecom/commit/6e6c8ce8c8170b048787b69b16d5359e1a34f904))
* remove naive drop cluster and wrong use of SelectRelease, BindMap ([e245a06](https://github.com/TranXuanPhong25/ecom/commit/e245a0639dd9bb3c1df1c53d8f5160de52032370))
* rewrite host injection ([8dcf813](https://github.com/TranXuanPhong25/ecom/commit/8dcf813c5a0c10676bb64397d58263d85b360dfd))

## [3.15.1](https://github.com/TranXuanPhong25/ecom/compare/v3.15.0...v3.15.1) (2025-09-22)

### 🐛 Bug Fixes

* correct path for minio ([9c26635](https://github.com/TranXuanPhong25/ecom/commit/9c26635daf12e64fa9a27e8391de443f31909f97))

## [3.15.0](https://github.com/TranXuanPhong25/ecom/compare/v3.14.0...v3.15.0) (2025-09-22)

### 🚀 Features

* minio routing ([9772870](https://github.com/TranXuanPhong25/ecom/commit/9772870b8c30188b7ba7f16e311c62a4ee8e7b65))

## [3.14.0](https://github.com/TranXuanPhong25/ecom/compare/v3.13.0...v3.14.0) (2025-09-21)

### 🚀 Features

* chatbot with Tavily search tool ([4e39753](https://github.com/TranXuanPhong25/ecom/commit/4e39753633e573f21fc9e9b36c3a0a30523bbf1a))
* parse category path to object ([845ef51](https://github.com/TranXuanPhong25/ecom/commit/845ef51b18705076b801377815eb0e3b23fb1acb))

### 🔨 Technical Changes

* remove docker-compose ([a14079c](https://github.com/TranXuanPhong25/ecom/commit/a14079cc39894509b02ba2de8eceb4fa0f81c78e))
* temporary allow post product-categories ([9dab403](https://github.com/TranXuanPhong25/ecom/commit/9dab403b645f9c1a28542af25ff3d2eecf3c32df))

## [3.13.0](https://github.com/TranXuanPhong25/ecom/compare/v3.12.0...v3.13.0) (2025-09-19)

### 🚀 Features

* **auth:** logout endpoint ([0daf2c5](https://github.com/TranXuanPhong25/ecom/commit/0daf2c587c0da3fe5892b374ef7f31abb598c768))

### 🔨 Technical Changes

* **shops:** shops model refined ([c67fd8f](https://github.com/TranXuanPhong25/ecom/commit/c67fd8f225a088942f9b45a61169264876bb4ed2))

## [3.12.0](https://github.com/TranXuanPhong25/ecom/compare/v3.11.0...v3.12.0) (2025-09-05)

### 🚀 Features

* architecture overview ([8bc8b49](https://github.com/TranXuanPhong25/ecom/commit/8bc8b49d58195b5cdb28bbe251d272f857eb1865))

### 🔨 Technical Changes

* **shops:** return correct status code for entity notfound ([1704bb2](https://github.com/TranXuanPhong25/ecom/commit/1704bb2402f151c0814561b5de34a95a39f6f733))

## [3.11.0](https://github.com/TranXuanPhong25/ecom/compare/v3.10.0...v3.11.0) (2025-09-05)

### 🚀 Features

* add ClusterIP services for auth, product-categories, products, and users ([cefc8bd](https://github.com/TranXuanPhong25/ecom/commit/cefc8bd93e62bf2901754e1e7fd7f839b0600d13))
* add cors policy ([27dd026](https://github.com/TranXuanPhong25/ecom/commit/27dd0260f680f4d2843c09af8359fe4dab0cb55f))
* add Dockerfile for eas-inbound-filter plugin and update service addresses ([22a8131](https://github.com/TranXuanPhong25/ecom/commit/22a8131235531e5e0716b83c4fa515e384040ee1))
* add gateway-wide auth policy and inbound filter service configurations ([94d521d](https://github.com/TranXuanPhong25/ecom/commit/94d521de810315aa78a1c2f51c218b8ec650f832))
* add HTTPRoute configurations for auth, brands, product-categories, product-reviews, products, shops, and upload services ([5c1b507](https://github.com/TranXuanPhong25/ecom/commit/5c1b5077c5e79c1955c5c27ebdb20d2a9d344fb1))
* auth config integrated ([0dffa5a](https://github.com/TranXuanPhong25/ecom/commit/0dffa5a2a086b1f89258cda158eb65c7d2d6fe3f))
* configs go modules ([7b2887e](https://github.com/TranXuanPhong25/ecom/commit/7b2887e92eb0b6e143524dadf42434349332d9a1))
* implement inbound filter service ([d93ecdf](https://github.com/TranXuanPhong25/ecom/commit/d93ecdfad4e2a26225a9db10cf633025ec8b57f6))
* integrated shops config ([d637c84](https://github.com/TranXuanPhong25/ecom/commit/d637c8448fc897871fff2736e692c4a9ffe2fb7d))
* jwt config integrated ([1054fce](https://github.com/TranXuanPhong25/ecom/commit/1054fce934b960e802a22a62e7726b2b0616f53d))
* jwt-service ([d756145](https://github.com/TranXuanPhong25/ecom/commit/d756145dd16458ce34397ec451c80491ea588bde))
* merge cors and extAuth into one sec-policy ([158d283](https://github.com/TranXuanPhong25/ecom/commit/158d2835413b5f17c7f793d5d0a003852216fdaa))
* minio server ([19b4889](https://github.com/TranXuanPhong25/ecom/commit/19b4889ee43cc9f4e0469f102b68ca4f45d457ad))
* opa-server ([83f957a](https://github.com/TranXuanPhong25/ecom/commit/83f957adecc0c33805688423c33deba052878c06))
* ran auth and product-categories service ([912a294](https://github.com/TranXuanPhong25/ecom/commit/912a294108beba4b2c47fcb7e9394360c7ea3994))
* upload service ([dc5b38c](https://github.com/TranXuanPhong25/ecom/commit/dc5b38ca43749c31a71c1c8f25f285f9d4a4338c))
* upload-svc config integrated ([2837985](https://github.com/TranXuanPhong25/ecom/commit/2837985f3ba00cd8f0a9106223a05f7716ad2d04))
* users config integrated ([5bbaedd](https://github.com/TranXuanPhong25/ecom/commit/5bbaedd169c623ac650b0d9c082471de327645a1))

### 🐛 Bug Fixes

* correct listen on port users service ([395b2f5](https://github.com/TranXuanPhong25/ecom/commit/395b2f526a01d536d6e26215f435cb129a99a13a))
* namespace headless service correction ([a01e4a4](https://github.com/TranXuanPhong25/ecom/commit/a01e4a495b841efb896214128ab4878c5ce0f629))

### 🔨 Technical Changes

* change dns name of product-categories pass to products service ([9a1e276](https://github.com/TranXuanPhong25/ecom/commit/9a1e276c900142f3640cdfc6f774a5717c058487))
* consistently use 8080 load balancer port ([2cce82b](https://github.com/TranXuanPhong25/ecom/commit/2cce82bdf3a72047a7189f4db6aa5dfe9c7a100d))
* **k8s:** always pull image ([eed40df](https://github.com/TranXuanPhong25/ecom/commit/eed40df1d939c24a0b38d55513885c5a43b57b25))
* **k8s:** remove deployments manifest ([6764027](https://github.com/TranXuanPhong25/ecom/commit/6764027a1a1f7f4913412a5453f083d236df7ac4))
* **k8s:** tidy load balancer service ([386b380](https://github.com/TranXuanPhong25/ecom/commit/386b380f9b981ebe48659811fcb4f736482f32bd))
* rename configs folder ([c85227f](https://github.com/TranXuanPhong25/ecom/commit/c85227f852dad9244426581679aace69d28188ce))
* use other cookies extraction logic for eas inbound service ([7be9ad6](https://github.com/TranXuanPhong25/ecom/commit/7be9ad6c65e75ffb015577a6fe2e132f1273b791))

## [3.10.0](https://github.com/TranXuanPhong25/ecom/compare/v3.9.0...v3.10.0) (2025-08-26)

### 🚀 Features

* **k8s:** worked manifest for users and products service ([e2f6d0c](https://github.com/TranXuanPhong25/ecom/commit/e2f6d0c572a88122b2239b18884a1adfefca68f0))

## [3.9.0](https://github.com/TranXuanPhong25/ecom/compare/v3.8.0...v3.9.0) (2025-08-25)

### 🚀 Features

* init k8s with Kong Ingress controller ([0567d6d](https://github.com/TranXuanPhong25/ecom/commit/0567d6db0016d55d6fda01f31d8d86b9dadf892e))

## [3.8.0](https://github.com/TranXuanPhong25/ecom/compare/v3.7.0...v3.8.0) (2025-08-25)

### 🚀 Features

* use custom kong image ([f472831](https://github.com/TranXuanPhong25/ecom/commit/f4728316df432da0981c2c728f7e00dad49d5b4e))

## [3.7.0](https://github.com/TranXuanPhong25/ecom/compare/v3.6.0...v3.7.0) (2025-08-24)

### 🚀 Features

* use kong db-less mode ([4bf0076](https://github.com/TranXuanPhong25/ecom/commit/4bf007654ff633d7cf22d1cdd06051cab78b229b))

### 🔨 Technical Changes

* remove product test ([fb732ee](https://github.com/TranXuanPhong25/ecom/commit/fb732ee28d8bb9924934ce5187398e308dd7420f))

## [3.6.0](https://github.com/TranXuanPhong25/ecom/compare/v3.5.1...v3.6.0) (2025-08-24)

### 🚀 Features

* upload decksync.sh for convenient ([83a3b73](https://github.com/TranXuanPhong25/ecom/commit/83a3b7340649fdca4c2412056a0bd1cd1160eb90))

## [3.5.1](https://github.com/TranXuanPhong25/ecom/compare/v3.5.0...v3.5.1) (2025-08-24)

### 🐛 Bug Fixes

* use declarative config for Kong ([6d11bc5](https://github.com/TranXuanPhong25/ecom/commit/6d11bc58d111d10ea506f1848c92192f4c8f2a3f))

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
