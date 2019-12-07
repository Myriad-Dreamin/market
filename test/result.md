FORMAT: 1A
HOST: 127.0.0.1

# Minimum-Market
this is the market backend powered by minimum

## Ping [/ping]


 + GET: result


## base-service.(*CRUDService).Post-fm [/v1/goods]


 + POST: 


### base-service.(*CRUDService).Post-fm [POST]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7856602+08:00",
                "g_type": 1,
                "name": "es000",
                "min_price": 100,
                "is_fixed": false,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "goods": {
                    "ID": 1,
                    "CreatedAt": "2019-12-07T15:36:24.7856602+08:00",
                    "UpdatedAt": "2019-12-07T15:36:24.7856602+08:00",
                    "EndAt": "2019-12-08T15:36:24.7856602+08:00",
                    "Seller": 1,
                    "Buyer": 0,
                    "Type": 1,
                    "Name": "es000",
                    "MinPrice": 100,
                    "IsFixed": false,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7856602+08:00",
                "g_type": 1,
                "name": "es0000",
                "min_price": 100,
                "is_fixed": false,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "goods": {
                    "ID": 2,
                    "CreatedAt": "2019-12-07T15:36:24.7856602+08:00",
                    "UpdatedAt": "2019-12-07T15:36:24.7856602+08:00",
                    "EndAt": "2019-12-08T15:36:24.7856602+08:00",
                    "Seller": 1,
                    "Buyer": 0,
                    "Type": 1,
                    "Name": "es0000",
                    "MinPrice": 100,
                    "IsFixed": false,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7856602+08:00",
                "g_type": 1,
                "name": "es0001",
                "min_price": 100,
                "is_fixed": false,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "goods": {
                    "ID": 3,
                    "CreatedAt": "2019-12-07T15:36:24.7866598+08:00",
                    "UpdatedAt": "2019-12-07T15:36:24.7866598+08:00",
                    "EndAt": "2019-12-08T15:36:24.7856602+08:00",
                    "Seller": 1,
                    "Buyer": 0,
                    "Type": 1,
                    "Name": "es0001",
                    "MinPrice": 100,
                    "IsFixed": false,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7856602+08:00",
                "g_type": 1,
                "name": "es0002",
                "min_price": 100,
                "is_fixed": false,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "goods": {
                    "ID": 4,
                    "CreatedAt": "2019-12-07T15:36:24.7866598+08:00",
                    "UpdatedAt": "2019-12-07T15:36:24.7866598+08:00",
                    "EndAt": "2019-12-08T15:36:24.7856602+08:00",
                    "Seller": 1,
                    "Buyer": 0,
                    "Type": 1,
                    "Name": "es0002",
                    "MinPrice": 100,
                    "IsFixed": false,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7856602+08:00",
                "g_type": 1,
                "name": "es0003",
                "min_price": 100,
                "is_fixed": false,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "goods": {
                    "ID": 5,
                    "CreatedAt": "2019-12-07T15:36:24.7866598+08:00",
                    "UpdatedAt": "2019-12-07T15:36:24.7866598+08:00",
                    "EndAt": "2019-12-08T15:36:24.7856602+08:00",
                    "Seller": 1,
                    "Buyer": 0,
                    "Type": 1,
                    "Name": "es0003",
                    "MinPrice": 100,
                    "IsFixed": false,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7856602+08:00",
                "g_type": 1,
                "name": "es0004",
                "min_price": 100,
                "is_fixed": false,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "goods": {
                    "ID": 6,
                    "CreatedAt": "2019-12-07T15:36:24.7866598+08:00",
                    "UpdatedAt": "2019-12-07T15:36:24.7866598+08:00",
                    "EndAt": "2019-12-08T15:36:24.7856602+08:00",
                    "Seller": 1,
                    "Buyer": 0,
                    "Type": 1,
                    "Name": "es0004",
                    "MinPrice": 100,
                    "IsFixed": false,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7856602+08:00",
                "g_type": 1,
                "name": "es0005",
                "min_price": 100,
                "is_fixed": false,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "goods": {
                    "ID": 7,
                    "CreatedAt": "2019-12-07T15:36:24.7876593+08:00",
                    "UpdatedAt": "2019-12-07T15:36:24.7876593+08:00",
                    "EndAt": "2019-12-08T15:36:24.7856602+08:00",
                    "Seller": 1,
                    "Buyer": 0,
                    "Type": 1,
                    "Name": "es0005",
                    "MinPrice": 100,
                    "IsFixed": false,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7856602+08:00",
                "g_type": 1,
                "name": "es0006",
                "min_price": 100,
                "is_fixed": false,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "goods": {
                    "ID": 8,
                    "CreatedAt": "2019-12-07T15:36:24.7876593+08:00",
                    "UpdatedAt": "2019-12-07T15:36:24.7876593+08:00",
                    "EndAt": "2019-12-08T15:36:24.7856602+08:00",
                    "Seller": 1,
                    "Buyer": 0,
                    "Type": 1,
                    "Name": "es0006",
                    "MinPrice": 100,
                    "IsFixed": false,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7856602+08:00",
                "g_type": 1,
                "name": "es0007",
                "min_price": 100,
                "is_fixed": false,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "goods": {
                    "ID": 9,
                    "CreatedAt": "2019-12-07T15:36:24.7876593+08:00",
                    "UpdatedAt": "2019-12-07T15:36:24.7876593+08:00",
                    "EndAt": "2019-12-08T15:36:24.7856602+08:00",
                    "Seller": 1,
                    "Buyer": 0,
                    "Type": 1,
                    "Name": "es0007",
                    "MinPrice": 100,
                    "IsFixed": false,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7856602+08:00",
                "g_type": 1,
                "name": "es0008",
                "min_price": 100,
                "is_fixed": false,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "goods": {
                    "ID": 10,
                    "CreatedAt": "2019-12-07T15:36:24.7876593+08:00",
                    "UpdatedAt": "2019-12-07T15:36:24.7876593+08:00",
                    "EndAt": "2019-12-08T15:36:24.7856602+08:00",
                    "Seller": 1,
                    "Buyer": 0,
                    "Type": 1,
                    "Name": "es0008",
                    "MinPrice": 100,
                    "IsFixed": false,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7856602+08:00",
                "g_type": 1,
                "name": "es0009",
                "min_price": 100,
                "is_fixed": false,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "goods": {
                    "ID": 11,
                    "CreatedAt": "2019-12-07T15:36:24.7876593+08:00",
                    "UpdatedAt": "2019-12-07T15:36:24.7876593+08:00",
                    "EndAt": "2019-12-08T15:36:24.7856602+08:00",
                    "Seller": 1,
                    "Buyer": 0,
                    "Type": 1,
                    "Name": "es0009",
                    "MinPrice": 100,
                    "IsFixed": false,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7886595+08:00",
                "g_type": 0,
                "name": "es0",
                "min_price": 100,
                "is_fixed": false,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 3,
                "error": "Key: 'PostRequest.Type' Error:Field validation for 'Type' failed on the 'required' tag"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7886595+08:00",
                "g_type": 1,
                "name": "",
                "min_price": 100,
                "is_fixed": false,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 3,
                "error": "Key: 'PostRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7886595+08:00",
                "g_type": 1,
                "name": "",
                "min_price": 100,
                "is_fixed": false,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 3,
                "error": "Key: 'PostRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7886595+08:00",
                "g_type": 1,
                "is_fixed": false,
                "min_price": -1,
                "name": "es1"
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 3,
                "error": "json: cannot unmarshal number -1 into Go struct field PostRequest.min_price of type uint64"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-07T15:36:24.7886595+08:00",
                "g_type": 1,
                "name": "es0",
                "min_price": 100,
                "is_fixed": false,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 3,
                "error": "could not set end time before a duration shorter than minimum end duration"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-07T16:36:25.0556702+08:00",
                "g_type": 1,
                "name": "es000000000000000",
                "min_price": 1,
                "is_fixed": false,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "goods": {
                    "ID": 12,
                    "CreatedAt": "2019-12-07T15:36:25.0556702+08:00",
                    "UpdatedAt": "2019-12-07T15:36:25.0556702+08:00",
                    "EndAt": "2019-12-07T16:36:25.0556702+08:00",
                    "Seller": 1,
                    "Buyer": 0,
                    "Type": 1,
                    "Name": "es000000000000000",
                    "MinPrice": 1,
                    "IsFixed": false,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


## base-service.(*ListService).List-fm [/v1/goods-list]


 + GET: 


### base-service.(*ListService).List-fm [GET]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "goodss": [
                    {
                        "id": 1,
                        "created_at": "2019-12-07T15:36:24.7856602+08:00",
                        "updated_at": "2019-12-07T15:36:24.7856602+08:00",
                        "end_at": "2019-12-08T15:36:24.7856602+08:00",
                        "seller": {
                            "id": 1,
                            "nick_name": "admin_context",
                            "register_city": "Qing Dao S.D."
                        },
                        "buyer": null,
                        "g_type": 1,
                        "name": "es000",
                        "min_price": 100,
                        "is_fixed": false,
                        "description": "",
                        "status": 1
                    },
                    {
                        "id": 2,
                        "created_at": "2019-12-07T15:36:24.7856602+08:00",
                        "updated_at": "2019-12-07T15:36:24.7856602+08:00",
                        "end_at": "2019-12-08T15:36:24.7856602+08:00",
                        "seller": {
                            "id": 1,
                            "nick_name": "admin_context",
                            "register_city": "Qing Dao S.D."
                        },
                        "buyer": null,
                        "g_type": 1,
                        "name": "es0000",
                        "min_price": 100,
                        "is_fixed": false,
                        "description": "",
                        "status": 1
                    }
                ]
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "goodss": [
                    {
                        "id": 3,
                        "created_at": "2019-12-07T15:36:24.7866598+08:00",
                        "updated_at": "2019-12-07T15:36:24.7866598+08:00",
                        "end_at": "2019-12-08T15:36:24.7856602+08:00",
                        "seller": {
                            "id": 1,
                            "nick_name": "admin_context",
                            "register_city": "Qing Dao S.D."
                        },
                        "buyer": null,
                        "g_type": 1,
                        "name": "es0001",
                        "min_price": 100,
                        "is_fixed": false,
                        "description": "",
                        "status": 1
                    },
                    {
                        "id": 4,
                        "created_at": "2019-12-07T15:36:24.7866598+08:00",
                        "updated_at": "2019-12-07T15:36:24.7866598+08:00",
                        "end_at": "2019-12-08T15:36:24.7856602+08:00",
                        "seller": {
                            "id": 1,
                            "nick_name": "admin_context",
                            "register_city": "Qing Dao S.D."
                        },
                        "buyer": null,
                        "g_type": 1,
                        "name": "es0002",
                        "min_price": 100,
                        "is_fixed": false,
                        "description": "",
                        "status": 1
                    }
                ]
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "goodss": [
                    {
                        "id": 3,
                        "created_at": "2019-12-07T15:36:24.7866598+08:00",
                        "updated_at": "2019-12-07T15:36:24.7866598+08:00",
                        "end_at": "2019-12-08T15:36:24.7856602+08:00",
                        "seller": {
                            "id": 1,
                            "nick_name": "admin_context",
                            "register_city": "Qing Dao S.D."
                        },
                        "buyer": null,
                        "g_type": 1,
                        "name": "es0001",
                        "min_price": 100,
                        "is_fixed": false,
                        "description": "",
                        "status": 1
                    }
                ]
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "goodss": null
            }


## base-service.(*CRUDService).Get-fm [/v1/goods/:goid]


 + GET: 
 + DELETE: 
 + PUT: 


### base-service.(*CRUDService).Get-fm [GET]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "goods": {
                    "ID": 1,
                    "CreatedAt": "2019-12-07T15:36:24.7856602+08:00",
                    "UpdatedAt": "2019-12-07T15:36:24.7856602+08:00",
                    "EndAt": "2019-12-08T15:36:24.7856602+08:00",
                    "Seller": 1,
                    "Buyer": 0,
                    "Type": 1,
                    "Name": "es000",
                    "MinPrice": 100,
                    "IsFixed": false,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 102,
                "error": "not found"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "goods": {
                    "ID": 12,
                    "CreatedAt": "2019-12-07T15:36:25.0556702+08:00",
                    "UpdatedAt": "2019-12-07T15:36:25.0566704+08:00",
                    "EndAt": "2019-12-07T16:36:25.0556702+08:00",
                    "Seller": 1,
                    "Buyer": 1,
                    "Type": 1,
                    "Name": "es000000000000000",
                    "MinPrice": 1,
                    "IsFixed": false,
                    "Description": "",
                    "Status": 2,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


### base-service.(*CRUDService).Delete-fm [DELETE]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0
            }


## Login [/v1/login]


 + POST: Login


### Login [POST]

+ Request admin login for test

    + Headers

            Content-Type: application/json

    + Body

            {
                "id": 1,
                "user_name": "",
                "phone": "",
                "password": "admin"
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "identity": null,
                "phone": "1234567891011",
                "id": 1,
                "nick_name": "admin_context",
                "name": "admin_context",
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzYzMDg5ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjF9fQ.QnHtEsIBVBeW_HFEp2oxbVKIVUT4pBIbjDCZUGX0c2U"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "id": 2,
                "user_name": "",
                "phone": "",
                "password": "11112222"
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "identity": null,
                "phone": "10086111",
                "id": 2,
                "nick_name": "tan chan",
                "name": "chan tan",
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.D7iYGFKwu6yF7-sXAffJoFEGng77CeZ_vzguLuDhjSc",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzYzMDg5ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.HrKpU0_FpzXk7e6peUShQZ4ySkWX4lBSqcBYsWBZrzU"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "id": 0,
                "user_name": "chan tan",
                "phone": "",
                "password": "11112222"
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "identity": null,
                "phone": "10086111",
                "id": 2,
                "nick_name": "tan chan",
                "name": "chan tan",
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.D7iYGFKwu6yF7-sXAffJoFEGng77CeZ_vzguLuDhjSc",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzYzMDg5ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.HrKpU0_FpzXk7e6peUShQZ4ySkWX4lBSqcBYsWBZrzU"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "id": 0,
                "user_name": "",
                "phone": "10086111",
                "password": "11112222"
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "identity": null,
                "phone": "10086111",
                "id": 2,
                "nick_name": "tan chan",
                "name": "chan tan",
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.t6_NehYKB7jXM0xmD9wCsd9os2oMktvYun4485u7FjA",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzYzMDg5ODUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc1LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzU3MDc3ODUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.hVei3O0B5AyK-eAjs0fPMJi5Exa75b5E2YXDK4lyQvw"
            }


## base-service.(*CRUDService).Post-fm [/v1/needs]


 + POST: 


### base-service.(*CRUDService).Post-fm [POST]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7916604+08:00",
                "g_type": 1,
                "name": "es000",
                "min_price": 100,
                "max_price": 100,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "needs": {
                    "ID": 1,
                    "created_at": "2019-12-07T15:36:24.7916604+08:00",
                    "updated_at": "2019-12-07T15:36:24.7916604+08:00",
                    "EndAt": "2019-12-08T15:36:24.7916604+08:00",
                    "Buyer": 1,
                    "Seller": 0,
                    "Type": 1,
                    "Name": "es000",
                    "MinPrice": 100,
                    "MaxPrice": 100,
                    "EndDuration": 0,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7916604+08:00",
                "g_type": 1,
                "name": "es0000",
                "min_price": 100,
                "max_price": 100,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "needs": {
                    "ID": 2,
                    "created_at": "2019-12-07T15:36:24.7926615+08:00",
                    "updated_at": "2019-12-07T15:36:24.7926615+08:00",
                    "EndAt": "2019-12-08T15:36:24.7916604+08:00",
                    "Buyer": 1,
                    "Seller": 0,
                    "Type": 1,
                    "Name": "es0000",
                    "MinPrice": 100,
                    "MaxPrice": 100,
                    "EndDuration": 0,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7916604+08:00",
                "g_type": 1,
                "name": "es0001",
                "min_price": 100,
                "max_price": 100,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "needs": {
                    "ID": 3,
                    "created_at": "2019-12-07T15:36:24.7926615+08:00",
                    "updated_at": "2019-12-07T15:36:24.7926615+08:00",
                    "EndAt": "2019-12-08T15:36:24.7916604+08:00",
                    "Buyer": 1,
                    "Seller": 0,
                    "Type": 1,
                    "Name": "es0001",
                    "MinPrice": 100,
                    "MaxPrice": 100,
                    "EndDuration": 0,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7916604+08:00",
                "g_type": 1,
                "name": "es0002",
                "min_price": 100,
                "max_price": 100,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "needs": {
                    "ID": 4,
                    "created_at": "2019-12-07T15:36:24.7926615+08:00",
                    "updated_at": "2019-12-07T15:36:24.7926615+08:00",
                    "EndAt": "2019-12-08T15:36:24.7916604+08:00",
                    "Buyer": 1,
                    "Seller": 0,
                    "Type": 1,
                    "Name": "es0002",
                    "MinPrice": 100,
                    "MaxPrice": 100,
                    "EndDuration": 0,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7916604+08:00",
                "g_type": 1,
                "name": "es0003",
                "min_price": 100,
                "max_price": 100,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "needs": {
                    "ID": 5,
                    "created_at": "2019-12-07T15:36:24.7946618+08:00",
                    "updated_at": "2019-12-07T15:36:24.7946618+08:00",
                    "EndAt": "2019-12-08T15:36:24.7916604+08:00",
                    "Buyer": 1,
                    "Seller": 0,
                    "Type": 1,
                    "Name": "es0003",
                    "MinPrice": 100,
                    "MaxPrice": 100,
                    "EndDuration": 0,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7916604+08:00",
                "g_type": 1,
                "name": "es0004",
                "min_price": 100,
                "max_price": 100,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "needs": {
                    "ID": 6,
                    "created_at": "2019-12-07T15:36:24.7946618+08:00",
                    "updated_at": "2019-12-07T15:36:24.7946618+08:00",
                    "EndAt": "2019-12-08T15:36:24.7916604+08:00",
                    "Buyer": 1,
                    "Seller": 0,
                    "Type": 1,
                    "Name": "es0004",
                    "MinPrice": 100,
                    "MaxPrice": 100,
                    "EndDuration": 0,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7916604+08:00",
                "g_type": 1,
                "name": "es0005",
                "min_price": 100,
                "max_price": 100,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "needs": {
                    "ID": 7,
                    "created_at": "2019-12-07T15:36:24.7956606+08:00",
                    "updated_at": "2019-12-07T15:36:24.7956606+08:00",
                    "EndAt": "2019-12-08T15:36:24.7916604+08:00",
                    "Buyer": 1,
                    "Seller": 0,
                    "Type": 1,
                    "Name": "es0005",
                    "MinPrice": 100,
                    "MaxPrice": 100,
                    "EndDuration": 0,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7916604+08:00",
                "g_type": 1,
                "name": "es0006",
                "min_price": 100,
                "max_price": 100,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "needs": {
                    "ID": 8,
                    "created_at": "2019-12-07T15:36:24.7956606+08:00",
                    "updated_at": "2019-12-07T15:36:24.7956606+08:00",
                    "EndAt": "2019-12-08T15:36:24.7916604+08:00",
                    "Buyer": 1,
                    "Seller": 0,
                    "Type": 1,
                    "Name": "es0006",
                    "MinPrice": 100,
                    "MaxPrice": 100,
                    "EndDuration": 0,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7916604+08:00",
                "g_type": 1,
                "name": "es0007",
                "min_price": 100,
                "max_price": 100,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "needs": {
                    "ID": 9,
                    "created_at": "2019-12-07T15:36:24.7966628+08:00",
                    "updated_at": "2019-12-07T15:36:24.7966628+08:00",
                    "EndAt": "2019-12-08T15:36:24.7916604+08:00",
                    "Buyer": 1,
                    "Seller": 0,
                    "Type": 1,
                    "Name": "es0007",
                    "MinPrice": 100,
                    "MaxPrice": 100,
                    "EndDuration": 0,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7916604+08:00",
                "g_type": 1,
                "name": "es0008",
                "min_price": 100,
                "max_price": 100,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "needs": {
                    "ID": 10,
                    "created_at": "2019-12-07T15:36:24.7966628+08:00",
                    "updated_at": "2019-12-07T15:36:24.7966628+08:00",
                    "EndAt": "2019-12-08T15:36:24.7916604+08:00",
                    "Buyer": 1,
                    "Seller": 0,
                    "Type": 1,
                    "Name": "es0008",
                    "MinPrice": 100,
                    "MaxPrice": 100,
                    "EndDuration": 0,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T15:36:24.7916604+08:00",
                "g_type": 1,
                "name": "es0009",
                "min_price": 100,
                "max_price": 100,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "needs": {
                    "ID": 11,
                    "created_at": "2019-12-07T15:36:24.7966628+08:00",
                    "updated_at": "2019-12-07T15:36:24.7966628+08:00",
                    "EndAt": "2019-12-08T15:36:24.7916604+08:00",
                    "Buyer": 1,
                    "Seller": 0,
                    "Type": 1,
                    "Name": "es0009",
                    "MinPrice": 100,
                    "MaxPrice": 100,
                    "EndDuration": 0,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-07T16:36:25.0566704+08:00",
                "g_type": 1,
                "name": "es000000000000000",
                "min_price": 1,
                "max_price": 2,
                "description": ""
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "needs": {
                    "ID": 12,
                    "created_at": "2019-12-07T15:36:25.0566704+08:00",
                    "updated_at": "2019-12-07T15:36:25.0566704+08:00",
                    "EndAt": "2019-12-07T16:36:25.0566704+08:00",
                    "Buyer": 1,
                    "Seller": 0,
                    "Type": 1,
                    "Name": "es000000000000000",
                    "MinPrice": 1,
                    "MaxPrice": 2,
                    "EndDuration": 0,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


## base-service.(*ListService).List-fm [/v1/needs-list]


 + GET: 


### base-service.(*ListService).List-fm [GET]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "needss": [
                    {
                        "id": 1,
                        "created_at": "2019-12-07T15:36:24.7916604+08:00",
                        "updated_at": "2019-12-07T15:36:24.7916604+08:00",
                        "end_at": "2019-12-08T15:36:24.7916604+08:00",
                        "seller": null,
                        "buyer": {
                            "id": 1,
                            "nick_name": "admin_context",
                            "register_city": "Qing Dao S.D."
                        },
                        "ddd": 0,
                        "g_type": 1,
                        "name": "es000",
                        "is_fixed": false,
                        "description": "",
                        "status": 1
                    },
                    {
                        "id": 2,
                        "created_at": "2019-12-07T15:36:24.7926615+08:00",
                        "updated_at": "2019-12-07T15:36:24.7926615+08:00",
                        "end_at": "2019-12-08T15:36:24.7916604+08:00",
                        "seller": null,
                        "buyer": {
                            "id": 1,
                            "nick_name": "admin_context",
                            "register_city": "Qing Dao S.D."
                        },
                        "ddd": 0,
                        "g_type": 1,
                        "name": "es0000",
                        "is_fixed": false,
                        "description": "",
                        "status": 1
                    }
                ]
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "needss": [
                    {
                        "id": 3,
                        "created_at": "2019-12-07T15:36:24.7926615+08:00",
                        "updated_at": "2019-12-07T15:36:24.7926615+08:00",
                        "end_at": "2019-12-08T15:36:24.7916604+08:00",
                        "seller": null,
                        "buyer": {
                            "id": 1,
                            "nick_name": "admin_context",
                            "register_city": "Qing Dao S.D."
                        },
                        "ddd": 0,
                        "g_type": 1,
                        "name": "es0001",
                        "is_fixed": false,
                        "description": "",
                        "status": 1
                    },
                    {
                        "id": 4,
                        "created_at": "2019-12-07T15:36:24.7926615+08:00",
                        "updated_at": "2019-12-07T15:36:24.7926615+08:00",
                        "end_at": "2019-12-08T15:36:24.7916604+08:00",
                        "seller": null,
                        "buyer": {
                            "id": 1,
                            "nick_name": "admin_context",
                            "register_city": "Qing Dao S.D."
                        },
                        "ddd": 0,
                        "g_type": 1,
                        "name": "es0002",
                        "is_fixed": false,
                        "description": "",
                        "status": 1
                    }
                ]
            }


## base-service.(*CRUDService).Put-fm [/v1/needs/:nid]


 + PUT: 
 + GET: 
 + DELETE: 


### base-service.(*CRUDService).Get-fm [GET]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "needs": {
                    "ID": 1,
                    "created_at": "2019-12-07T15:36:24.7916604+08:00",
                    "updated_at": "2019-12-07T15:36:24.7916604+08:00",
                    "EndAt": "2019-12-08T15:36:24.7916604+08:00",
                    "Buyer": 1,
                    "Seller": 0,
                    "Type": 1,
                    "Name": "es000",
                    "MinPrice": 100,
                    "MaxPrice": 100,
                    "EndDuration": 0,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 102,
                "error": "not found"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "needs": {
                    "ID": 12,
                    "created_at": "2019-12-07T15:36:25.0566704+08:00",
                    "updated_at": "2019-12-07T15:36:25.0576713+08:00",
                    "EndAt": "2019-12-07T16:36:25.0566704+08:00",
                    "Buyer": 1,
                    "Seller": 1,
                    "Type": 1,
                    "Name": "es000000000000000",
                    "MinPrice": 1,
                    "MaxPrice": 2,
                    "EndDuration": 0,
                    "Description": "",
                    "Status": 2,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
            }


### base-service.(*CRUDService).Delete-fm [DELETE]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0
            }


## Register [/v1/user]


 + POST: Register


### Register [POST]

+ Request 

    + Headers

            Content-Type: application/json

    + Body

            {
                "name": "admin_context",
                "password": "admin",
                "nick_name": "admin_context",
                "phone": "1234567891011",
                "register_city": "Qing Dao S.D."
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "id": 1
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "name": "chan tan",
                "password": "11112222",
                "nick_name": "tan chan",
                "phone": "10086111",
                "register_city": "tan arch"
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "id": 2
            }


## List [/v1/user-list]


 + GET: List User


### List [GET]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "users": [
                    {
                        "ID": 1,
                        "created_at": "2019-12-07T15:36:24.7196618+08:00",
                        "updated_at": "2019-12-07T15:36:24.7196618+08:00",
                        "last_login": "2019-12-07T15:36:24.7836939+08:00",
                        "NickName": "admin_context",
                        "Name": "admin_context",
                        "Password": "$2a$10$C9Oh/.nKGeLdF8OwD/jZtOwoNtEta3zNVjmN5aQjT8MDKGglm5oMm",
                        "Phone": "1234567891011",
                        "RegisterCity": "Qing Dao S.D."
                    }
                ]
            }


## Put [/v1/user/:id]


 + PUT: Put User
 + GET: Get User
 + DELETE: Delete User


### Get [GET]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0,
                "user": {
                    "ID": 1,
                    "created_at": "2019-12-07T15:36:24.7196618+08:00",
                    "updated_at": "2019-12-07T15:36:24.7196618+08:00",
                    "last_login": "2019-12-07T15:36:24.7836939+08:00",
                    "NickName": "admin_context",
                    "Name": "admin_context",
                    "Password": "$2a$10$C9Oh/.nKGeLdF8OwD/jZtOwoNtEta3zNVjmN5aQjT8MDKGglm5oMm",
                    "Phone": "1234567891011",
                    "RegisterCity": "Qing Dao S.D."
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 102,
                "error": "not found"
            }


### Delete [DELETE]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0
            }


## Buy [/v1/user/:id/buy]


 + POST: Buy Goods


### Buy [POST]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "goid": 12
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0
            }


## ChangePassword [/v1/user/:id/password]


 + PUT: change password of user


### ChangePassword [PUT]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "old-password": "11112222",
                "new-password": "111122222"
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0
            }


## Sell [/v1/user/:id/sell]


 + POST: Sell Needs


### Sell [POST]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MDc3ODQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzA0MTc0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.V5WA8vGCqtepdw_YuVMCwlM7o5pDoLKQHgftddR9jNg
            Content-Type: application/json

    + Body

            {
                "nid": 12
            }


+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0
            }

