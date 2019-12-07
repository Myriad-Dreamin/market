FORMAT: 1A
HOST: 127.0.0.1

# Minimum-Market
this is the market backend powered by minimum

## Ping [/ping]


 + GET: result


## Post [/v1/goods]


 + POST: Post Goods


### Post [POST]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6442235+08:00",
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
                    "CreatedAt": "2019-12-07T21:54:05.6442235+08:00",
                    "UpdatedAt": "2019-12-07T21:54:05.6442235+08:00",
                    "EndAt": "2019-12-08T21:54:05.6442235+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6442235+08:00",
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
                    "CreatedAt": "2019-12-07T21:54:05.6442235+08:00",
                    "UpdatedAt": "2019-12-07T21:54:05.6442235+08:00",
                    "EndAt": "2019-12-08T21:54:05.6442235+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6442235+08:00",
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
                    "CreatedAt": "2019-12-07T21:54:05.6452299+08:00",
                    "UpdatedAt": "2019-12-07T21:54:05.6452299+08:00",
                    "EndAt": "2019-12-08T21:54:05.6442235+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6442235+08:00",
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
                    "CreatedAt": "2019-12-07T21:54:05.6452299+08:00",
                    "UpdatedAt": "2019-12-07T21:54:05.6452299+08:00",
                    "EndAt": "2019-12-08T21:54:05.6442235+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6442235+08:00",
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
                    "CreatedAt": "2019-12-07T21:54:05.6452299+08:00",
                    "UpdatedAt": "2019-12-07T21:54:05.6452299+08:00",
                    "EndAt": "2019-12-08T21:54:05.6442235+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6442235+08:00",
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
                    "CreatedAt": "2019-12-07T21:54:05.6452299+08:00",
                    "UpdatedAt": "2019-12-07T21:54:05.6452299+08:00",
                    "EndAt": "2019-12-08T21:54:05.6442235+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6442235+08:00",
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
                    "CreatedAt": "2019-12-07T21:54:05.6462297+08:00",
                    "UpdatedAt": "2019-12-07T21:54:05.6462297+08:00",
                    "EndAt": "2019-12-08T21:54:05.6442235+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6442235+08:00",
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
                    "CreatedAt": "2019-12-07T21:54:05.6462297+08:00",
                    "UpdatedAt": "2019-12-07T21:54:05.6462297+08:00",
                    "EndAt": "2019-12-08T21:54:05.6442235+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6442235+08:00",
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
                    "CreatedAt": "2019-12-07T21:54:05.6462297+08:00",
                    "UpdatedAt": "2019-12-07T21:54:05.6462297+08:00",
                    "EndAt": "2019-12-08T21:54:05.6442235+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6442235+08:00",
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
                    "CreatedAt": "2019-12-07T21:54:05.6462297+08:00",
                    "UpdatedAt": "2019-12-07T21:54:05.6462297+08:00",
                    "EndAt": "2019-12-08T21:54:05.6442235+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6442235+08:00",
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
                    "CreatedAt": "2019-12-07T21:54:05.6462297+08:00",
                    "UpdatedAt": "2019-12-07T21:54:05.6462297+08:00",
                    "EndAt": "2019-12-08T21:54:05.6442235+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6472314+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6472314+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6472314+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6472314+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-07T21:54:05.6472314+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-07T22:54:05.9182243+08:00",
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
                    "CreatedAt": "2019-12-07T21:54:05.9182243+08:00",
                    "UpdatedAt": "2019-12-07T21:54:05.9182243+08:00",
                    "EndAt": "2019-12-07T22:54:05.9182243+08:00",
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


## List [/v1/goods-list]


 + GET: List Goods


### List [GET]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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
                        "created_at": "2019-12-07T21:54:05.6442235+08:00",
                        "updated_at": "2019-12-07T21:54:05.6442235+08:00",
                        "end_at": "2019-12-08T21:54:05.6442235+08:00",
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
                        "created_at": "2019-12-07T21:54:05.6442235+08:00",
                        "updated_at": "2019-12-07T21:54:05.6442235+08:00",
                        "end_at": "2019-12-08T21:54:05.6442235+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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
                        "created_at": "2019-12-07T21:54:05.6452299+08:00",
                        "updated_at": "2019-12-07T21:54:05.6452299+08:00",
                        "end_at": "2019-12-08T21:54:05.6442235+08:00",
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
                        "created_at": "2019-12-07T21:54:05.6452299+08:00",
                        "updated_at": "2019-12-07T21:54:05.6452299+08:00",
                        "end_at": "2019-12-08T21:54:05.6442235+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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
                        "created_at": "2019-12-07T21:54:05.6452299+08:00",
                        "updated_at": "2019-12-07T21:54:05.6452299+08:00",
                        "end_at": "2019-12-08T21:54:05.6442235+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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


## Get [/v1/goods/:goid]


 + GET: Get Goods
 + DELETE: Delete Goods
 + PUT: Put Goods


### Get [GET]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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
                    "CreatedAt": "2019-12-07T21:54:05.6442235+08:00",
                    "UpdatedAt": "2019-12-07T21:54:05.6442235+08:00",
                    "EndAt": "2019-12-08T21:54:05.6442235+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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
                    "CreatedAt": "2019-12-07T21:54:05.9182243+08:00",
                    "UpdatedAt": "2019-12-07T21:54:05.9182243+08:00",
                    "EndAt": "2019-12-07T22:54:05.9182243+08:00",
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


### Delete [DELETE]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0
            }


## Force Delete [/v1/goods/:goid/force]


 + DELETE: Force Delete Goods


### Force Delete [DELETE]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzYzMzE2NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjF9fQ.86iszmQzPOSEcHs8t1CP_PFm1ofGSSpqYcq277AHBEY"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.jtl6XRKgE9BbuIaXB37ZZjlU0vkRhyULVQNJqIuP7I8",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzYzMzE2NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.dLCEylejwrSF7rxh-LZmenjGcYnFMLMBILMej4aBym8"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.jtl6XRKgE9BbuIaXB37ZZjlU0vkRhyULVQNJqIuP7I8",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzYzMzE2NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.dLCEylejwrSF7rxh-LZmenjGcYnFMLMBILMej4aBym8"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.jtl6XRKgE9BbuIaXB37ZZjlU0vkRhyULVQNJqIuP7I8",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzYzMzE2NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.dLCEylejwrSF7rxh-LZmenjGcYnFMLMBILMej4aBym8"
            }


## Post [/v1/needs]


 + POST: Post Needs


### Post [POST]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6522297+08:00",
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
                    "created_at": "2019-12-07T21:54:05.6522297+08:00",
                    "updated_at": "2019-12-07T21:54:05.6522297+08:00",
                    "EndAt": "2019-12-08T21:54:05.6522297+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6522297+08:00",
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
                    "created_at": "2019-12-07T21:54:05.6522297+08:00",
                    "updated_at": "2019-12-07T21:54:05.6522297+08:00",
                    "EndAt": "2019-12-08T21:54:05.6522297+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6522297+08:00",
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
                    "created_at": "2019-12-07T21:54:05.653224+08:00",
                    "updated_at": "2019-12-07T21:54:05.653224+08:00",
                    "EndAt": "2019-12-08T21:54:05.6522297+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6522297+08:00",
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
                    "created_at": "2019-12-07T21:54:05.653224+08:00",
                    "updated_at": "2019-12-07T21:54:05.653224+08:00",
                    "EndAt": "2019-12-08T21:54:05.6522297+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6522297+08:00",
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
                    "created_at": "2019-12-07T21:54:05.6542304+08:00",
                    "updated_at": "2019-12-07T21:54:05.6542304+08:00",
                    "EndAt": "2019-12-08T21:54:05.6522297+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6522297+08:00",
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
                    "created_at": "2019-12-07T21:54:05.6542304+08:00",
                    "updated_at": "2019-12-07T21:54:05.6542304+08:00",
                    "EndAt": "2019-12-08T21:54:05.6522297+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6522297+08:00",
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
                    "created_at": "2019-12-07T21:54:05.6552294+08:00",
                    "updated_at": "2019-12-07T21:54:05.6552294+08:00",
                    "EndAt": "2019-12-08T21:54:05.6522297+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6522297+08:00",
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
                    "created_at": "2019-12-07T21:54:05.6552294+08:00",
                    "updated_at": "2019-12-07T21:54:05.6552294+08:00",
                    "EndAt": "2019-12-08T21:54:05.6522297+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6522297+08:00",
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
                    "created_at": "2019-12-07T21:54:05.6562307+08:00",
                    "updated_at": "2019-12-07T21:54:05.6562307+08:00",
                    "EndAt": "2019-12-08T21:54:05.6522297+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6522297+08:00",
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
                    "created_at": "2019-12-07T21:54:05.6562307+08:00",
                    "updated_at": "2019-12-07T21:54:05.6562307+08:00",
                    "EndAt": "2019-12-08T21:54:05.6522297+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-08T21:54:05.6522297+08:00",
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
                    "created_at": "2019-12-07T21:54:05.6572301+08:00",
                    "updated_at": "2019-12-07T21:54:05.6572301+08:00",
                    "EndAt": "2019-12-08T21:54:05.6522297+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-12-07T22:54:05.9192276+08:00",
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
                    "created_at": "2019-12-07T21:54:05.9192276+08:00",
                    "updated_at": "2019-12-07T21:54:05.9192276+08:00",
                    "EndAt": "2019-12-07T22:54:05.9192276+08:00",
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


## List [/v1/needs-list]


 + GET: List Needs


### List [GET]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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
                        "created_at": "2019-12-07T21:54:05.6522297+08:00",
                        "updated_at": "2019-12-07T21:54:05.6522297+08:00",
                        "end_at": "2019-12-08T21:54:05.6522297+08:00",
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
                        "created_at": "2019-12-07T21:54:05.6522297+08:00",
                        "updated_at": "2019-12-07T21:54:05.6522297+08:00",
                        "end_at": "2019-12-08T21:54:05.6522297+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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
                        "created_at": "2019-12-07T21:54:05.653224+08:00",
                        "updated_at": "2019-12-07T21:54:05.653224+08:00",
                        "end_at": "2019-12-08T21:54:05.6522297+08:00",
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
                        "created_at": "2019-12-07T21:54:05.653224+08:00",
                        "updated_at": "2019-12-07T21:54:05.653224+08:00",
                        "end_at": "2019-12-08T21:54:05.6522297+08:00",
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


## Put [/v1/needs/:nid]


 + PUT: Put Needs
 + DELETE: Delete Needs
 + GET: Get Needs


### Delete [DELETE]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0
            }


### Get [GET]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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
                    "created_at": "2019-12-07T21:54:05.6522297+08:00",
                    "updated_at": "2019-12-07T21:54:05.6522297+08:00",
                    "EndAt": "2019-12-08T21:54:05.6522297+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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
                    "created_at": "2019-12-07T21:54:05.9192276+08:00",
                    "updated_at": "2019-12-07T21:54:05.9192276+08:00",
                    "EndAt": "2019-12-07T22:54:05.9192276+08:00",
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


## Force Delete [/v1/needs/:nid/force]


 + DELETE: Force Delete Needs


### Force Delete [DELETE]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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
                        "created_at": "2019-12-07T21:54:05.5782226+08:00",
                        "updated_at": "2019-12-07T21:54:05.5782226+08:00",
                        "last_login": "2019-12-07T21:54:05.642252+08:00",
                        "NickName": "admin_context",
                        "Name": "admin_context",
                        "Password": "$2a$10$SoUTh19pxRRZNXKORPqqFu8qkKTTaMOvUHEXubr1MmoRkgQpFtLxe",
                        "Phone": "1234567891011",
                        "RegisterCity": "Qing Dao S.D."
                    }
                ]
            }


## Delete [/v1/user/:id]


 + DELETE: Delete User
 + PUT: Put User
 + GET: Get User


### Delete [DELETE]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0
            }


### Get [GET]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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
                    "created_at": "2019-12-07T21:54:05.5782226+08:00",
                    "updated_at": "2019-12-07T21:54:05.5782226+08:00",
                    "last_login": "2019-12-07T21:54:05.642252+08:00",
                    "NickName": "admin_context",
                    "Name": "admin_context",
                    "Password": "$2a$10$SoUTh19pxRRZNXKORPqqFu8qkKTTaMOvUHEXubr1MmoRkgQpFtLxe",
                    "Phone": "1234567891011",
                    "RegisterCity": "Qing Dao S.D."
                }
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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


## Buy [/v1/user/:id/goods/:goid/buy]


 + POST: Buy Goods


### Buy [POST]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: text/plain

    + Body

            

+ Response 200

    + Headers

            Content-Type: application/json; charset=utf-8

    + Body

            {
                "code": 0
            }


## ConfirmBuy [/v1/user/:id/goods/:goid/confirm-buy]


 + POST: ConfirmBuy Goods


## ConfirmSell [/v1/user/:id/needs/:nid/confirm-sell]


 + POST: ConfirmSell Needs


## Sell [/v1/user/:id/needs/:nid/sell]


 + POST: Sell Needs


### Sell [POST]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
            Content-Type: text/plain

    + Body

            

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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzU3MzA0NDUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc1NzI2ODM1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.1oPYi0LdOk-ptCLK0h_rmBuPsUZ0YnJUgsu6Hc35rns
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

