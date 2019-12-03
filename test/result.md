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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-04T13:42:25.854995+08:00",
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
                    "CreatedAt": "2019-12-03T13:42:25.854995+08:00",
                    "UpdatedAt": "2019-12-03T13:42:25.854995+08:00",
                    "EndAt": "2019-12-04T13:42:25.854995+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-04T13:42:25.854995+08:00",
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
                    "CreatedAt": "2019-12-03T13:42:25.854995+08:00",
                    "UpdatedAt": "2019-12-03T13:42:25.854995+08:00",
                    "EndAt": "2019-12-04T13:42:25.854995+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-04T13:42:25.854995+08:00",
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
                    "CreatedAt": "2019-12-03T13:42:25.854995+08:00",
                    "UpdatedAt": "2019-12-03T13:42:25.854995+08:00",
                    "EndAt": "2019-12-04T13:42:25.854995+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-04T13:42:25.854995+08:00",
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
                    "CreatedAt": "2019-12-03T13:42:25.8559931+08:00",
                    "UpdatedAt": "2019-12-03T13:42:25.8559931+08:00",
                    "EndAt": "2019-12-04T13:42:25.854995+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-04T13:42:25.854995+08:00",
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
                    "CreatedAt": "2019-12-03T13:42:25.8559931+08:00",
                    "UpdatedAt": "2019-12-03T13:42:25.8559931+08:00",
                    "EndAt": "2019-12-04T13:42:25.854995+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-04T13:42:25.854995+08:00",
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
                    "CreatedAt": "2019-12-03T13:42:25.8559931+08:00",
                    "UpdatedAt": "2019-12-03T13:42:25.8559931+08:00",
                    "EndAt": "2019-12-04T13:42:25.854995+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-04T13:42:25.854995+08:00",
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
                    "CreatedAt": "2019-12-03T13:42:25.8559931+08:00",
                    "UpdatedAt": "2019-12-03T13:42:25.8559931+08:00",
                    "EndAt": "2019-12-04T13:42:25.854995+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-04T13:42:25.854995+08:00",
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
                    "CreatedAt": "2019-12-03T13:42:25.8569935+08:00",
                    "UpdatedAt": "2019-12-03T13:42:25.8569935+08:00",
                    "EndAt": "2019-12-04T13:42:25.854995+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-04T13:42:25.854995+08:00",
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
                    "CreatedAt": "2019-12-03T13:42:25.8569935+08:00",
                    "UpdatedAt": "2019-12-03T13:42:25.8569935+08:00",
                    "EndAt": "2019-12-04T13:42:25.854995+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-04T13:42:25.854995+08:00",
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
                    "CreatedAt": "2019-12-03T13:42:25.8569935+08:00",
                    "UpdatedAt": "2019-12-03T13:42:25.8569935+08:00",
                    "EndAt": "2019-12-04T13:42:25.854995+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-04T13:42:25.854995+08:00",
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
                    "CreatedAt": "2019-12-03T13:42:25.8569935+08:00",
                    "UpdatedAt": "2019-12-03T13:42:25.8569935+08:00",
                    "EndAt": "2019-12-04T13:42:25.854995+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-04T13:42:25.8569935+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-04T13:42:25.8569935+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-04T13:42:25.8569935+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-04T13:42:25.8569935+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-03T13:42:25.8579934+08:00",
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


## base-service.(*ListService).List-fm [/v1/goods-list]


 + GET: 


### base-service.(*ListService).List-fm [GET]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
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
                        "created_at": "2019-12-03T13:42:25.854995+08:00",
                        "updated_at": "2019-12-03T13:42:25.854995+08:00",
                        "end_at": "2019-12-04T13:42:25.854995+08:00",
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
                        "created_at": "2019-12-03T13:42:25.854995+08:00",
                        "updated_at": "2019-12-03T13:42:25.854995+08:00",
                        "end_at": "2019-12-04T13:42:25.854995+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
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
                        "created_at": "2019-12-03T13:42:25.854995+08:00",
                        "updated_at": "2019-12-03T13:42:25.854995+08:00",
                        "end_at": "2019-12-04T13:42:25.854995+08:00",
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
                        "created_at": "2019-12-03T13:42:25.8559931+08:00",
                        "updated_at": "2019-12-03T13:42:25.8559931+08:00",
                        "end_at": "2019-12-04T13:42:25.854995+08:00",
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


## base-service.(*CRUDService).Put-fm [/v1/goods/:goid]


 + PUT: 
 + GET: 
 + DELETE: 


### base-service.(*CRUDService).Get-fm [GET]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
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
                    "CreatedAt": "2019-12-03T13:42:25.854995+08:00",
                    "UpdatedAt": "2019-12-03T13:42:25.854995+08:00",
                    "EndAt": "2019-12-04T13:42:25.854995+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
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


### base-service.(*CRUDService).Delete-fm [DELETE]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU5NTY1NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjF9fQ.gMJa-4299GhCeFauqhFOGQEuIiIK3lHFjsILNrfqTzw"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
            Content-Type: application/json

    + Body

            {
                "id": 2,
                "user_name": "",
                "phone": "",
                "password": "11122222"
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.roHNRO3x3h8zHNAblKvFLco8m7ZOS6kekTohxrag7A0",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU5NTY1NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.0BBhvXjRpbPSrxFVF0rvXjcqmurAmUo-m7eTSSqpNU4"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
            Content-Type: application/json

    + Body

            {
                "id": 0,
                "user_name": "chan tan",
                "phone": "",
                "password": "11122222"
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.Tnqjo-sH0DGMumERIlUJea74UC6dM5Nges9LHo5unTY",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU5NTY1NDYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM2LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzUzNTUzNDYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.etB1rdYDqojeO2-WyS1JAAzK9Yuq1i6rv-RuhUtPVZI"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
            Content-Type: application/json

    + Body

            {
                "id": 0,
                "user_name": "",
                "phone": "10086111",
                "password": "11122222"
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.Tnqjo-sH0DGMumERIlUJea74UC6dM5Nges9LHo5unTY",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU5NTY1NDYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM2LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzUzNTUzNDYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.etB1rdYDqojeO2-WyS1JAAzK9Yuq1i6rv-RuhUtPVZI"
            }


## base-service.(*CRUDService).Post-fm [/v1/needs]


 + POST: 


## base-service.(*ListService).List-fm [/v1/needs-list]


 + GET: 


## base-service.(*CRUDService).Get-fm [/v1/needs/:nid]


 + GET: 
 + PUT: 
 + DELETE: 


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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
            Content-Type: application/json

    + Body

            {
                "name": "chan tan",
                "password": "11122222",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
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
                        "created_at": "2019-12-03T13:42:25.7889647+08:00",
                        "updated_at": "2019-12-03T13:42:25.7889647+08:00",
                        "last_login": "2019-12-03T05:42:25Z",
                        "NickName": "admin_context",
                        "Name": "admin_context",
                        "Password": "$2a$10$3zI8VTaBvA1Z4TYDtVMzPOSMNqb5TyPHjH4LX2IdjkJT8EyVemhnO",
                        "Phone": "1234567891011",
                        "RegisterCity": "Qing Dao S.D."
                    }
                ]
            }


## Get [/v1/user/:id]


 + GET: Get User
 + PUT: Put User
 + DELETE: Delete User


### Get [GET]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUzNTUzNDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1MzUxNzM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.vrOjOad_bMf9bUZmEd477THgbESA4tnwwO9hZPKCbQg
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
                    "created_at": "2019-12-03T13:42:25.7889647+08:00",
                    "updated_at": "2019-12-03T13:42:25.7889647+08:00",
                    "last_login": "2019-12-03T05:42:25Z",
                    "NickName": "admin_context",
                    "Name": "admin_context",
                    "Password": "$2a$10$3zI8VTaBvA1Z4TYDtVMzPOSMNqb5TyPHjH4LX2IdjkJT8EyVemhnO",
                    "Phone": "1234567891011",
                    "RegisterCity": "Qing Dao S.D."
                }
            }


## ChangePassword [/v1/user/:id/password]


 + POST: change password of user

