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
    input.method == "POST"
    input.path == "/api/product-categories"
}

allow if {
    presigned_url_pattern := [
    "/api/upload/presigned-url",
    "/products-images"
    ]
    input.method in [ "POST", "GET" , "PUT", "DELETE", "OPTIONS" ]
    regex.match(presigned_url_pattern[_], input.path)
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
    pattern := `\/api\/products`
    regex.match(pattern, input.path)

    input.method in { "POST", "PUT", "DELETE", "OPTIONS" }
    input.authenticated == true
}

#restrict access to seller's products
allow if {
    pattern := `\/api\/products`
    regex.match(pattern, input.path)
#    input.authenticated == false
    input.method == "GET"
}

allow if {
    input.method in { "POST", "PUT", "DELETE", "OPTIONS", "GET" }
    regex.match(`\/api\/shops`, input.path)
#    input.authenticated == true
}

allow if {
    input.method in { "POST", "PUT", "DELETE", "OPTIONS", "GET" }
    regex.match(`\/api\/orders`, input.path)
    input.authenticated == true
}
allow if {
    input.method in { "POST", "PUT", "DELETE", "OPTIONS", "GET" }
    regex.match(`\/api\/carts`, input.path)
    input.authenticated == true
}

