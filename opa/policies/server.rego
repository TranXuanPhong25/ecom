package route

default allow = false



allow if {
    public_products_service_get_product_pattern := `/api/products/(\w|\d)+`
    input.request.method == "GET"
    regex.match(public_products_service_get_product_pattern, input.request.path)
}


# restrict access to products for authenticated users
allow if {
    input.request.method in { "POST", "PUT", "DELETE", "OPTIONS" }
    input.request.path == "/api/products"
    input.request.authenticated == true
}

#restrict access to seller's products
allow if {
    private_products_service_get_product_of_seller_pattern := `/api/products/(\w|\d)+\?seller_id=(\w|\d)`
    input.request.method == "GET"
    regex.match(private_products_service_get_product_of_seller_pattern, input.request.path)
    input.request.authenticated == false
}
