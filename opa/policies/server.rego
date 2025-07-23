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

allow if {
    public_auth_pattern := [
    "/api/auth/login",
    "/api/auth/register",
    "/api/auth/refresh",
    "/api/auth/logout",
    "/api/auth/forgot-password",
    "/api/auth/reset-password",
    "/api/auth/verify-email",
    "/api/auth/verify-email-resend"
    ]
    input.method in ["POST", "GET"]
    input.path in public_auth_pattern
}

allow if {
    protected_auth_pattern := [
        "/api/auth/me"
    ]
    input.method == "GET"
    input.path in protected_auth_pattern
}

# restrict access to products for authenticated users
allow if {
    input.method in { "POST", "PUT", "DELETE", "OPTIONS" }
    input.path == "/api/products"
    input.authenticated == true
}

#restrict access to seller's products
allow if {
    private_products_service_get_product_of_shop_pattern := `\/api\/products`
    input.method == "GET"
    regex.match(private_products_service_get_product_of_shop_pattern, input.path)
#    input.authenticated == false
}
