package route

default allow = false



allow if {
    public_get_pattern := [
    "\/api\/products\/(\\w|\\d)+",
     "\/api\/product-categories.*",
    ]
#    input.path == "/api/products"
    input.method == "GET"
    regex.match(public_get_pattern[_], input.path)
}


# restrict access to products for authenticated users
allow if {
    input.method in { "POST", "PUT", "DELETE", "OPTIONS" }
    input.path == "/api/products"
    input.authenticated == true
}

#restrict access to seller's products
allow if {
    private_products_service_get_product_of_seller_pattern := `/api/products/(\w|\d)+\?seller_id=(\w|\d)`
    input.method == "GET"
    regex.match(private_products_service_get_product_of_seller_pattern, input.path)
    input.authenticated == false
}
