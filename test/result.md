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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.DPJQE0fV4oqeHVgNQO_LAzTPvPZjV9wp3HS2pEHTuvM
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-11-23T19:36:43.132993711+08:00",
                "g_type": 1,
                "name": "es",
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
                    "CreatedAt": "2019-11-22T19:36:43.133268781+08:00",
                    "UpdatedAt": "2019-11-22T19:36:43.133268781+08:00",
                    "EndAt": "2019-11-23T19:36:43.132993711+08:00",
                    "Seller": 1,
                    "Buyer": 0,
                    "Type": 1,
                    "Name": "es",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.DPJQE0fV4oqeHVgNQO_LAzTPvPZjV9wp3HS2pEHTuvM
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-11-23T19:36:43.132993711+08:00",
                "g_type": 1,
                "name": "es00000",
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
                    "CreatedAt": "2019-11-22T19:36:43.133792914+08:00",
                    "UpdatedAt": "2019-11-22T19:36:43.133792914+08:00",
                    "EndAt": "2019-11-23T19:36:43.132993711+08:00",
                    "Seller": 1,
                    "Buyer": 0,
                    "Type": 1,
                    "Name": "es00000",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.DPJQE0fV4oqeHVgNQO_LAzTPvPZjV9wp3HS2pEHTuvM
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-11-23T19:36:43.132993711+08:00",
                "g_type": 1,
                "name": "es00001",
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
                    "CreatedAt": "2019-11-22T19:36:43.134157962+08:00",
                    "UpdatedAt": "2019-11-22T19:36:43.134157962+08:00",
                    "EndAt": "2019-11-23T19:36:43.132993711+08:00",
                    "Seller": 1,
                    "Buyer": 0,
                    "Type": 1,
                    "Name": "es00001",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.DPJQE0fV4oqeHVgNQO_LAzTPvPZjV9wp3HS2pEHTuvM
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-11-23T19:36:43.134414682+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.DPJQE0fV4oqeHVgNQO_LAzTPvPZjV9wp3HS2pEHTuvM
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-11-23T19:36:43.134414682+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.DPJQE0fV4oqeHVgNQO_LAzTPvPZjV9wp3HS2pEHTuvM
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-11-23T19:36:43.134414682+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.DPJQE0fV4oqeHVgNQO_LAzTPvPZjV9wp3HS2pEHTuvM
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-11-23T19:36:43.134414682+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.DPJQE0fV4oqeHVgNQO_LAzTPvPZjV9wp3HS2pEHTuvM
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-11-22T19:36:43.135083853+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.DPJQE0fV4oqeHVgNQO_LAzTPvPZjV9wp3HS2pEHTuvM
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
                        "ID": 1,
                        "CreatedAt": "2019-11-22T19:36:43.133268781+08:00",
                        "UpdatedAt": "2019-11-22T19:36:43.133268781+08:00",
                        "EndAt": "2019-11-23T19:36:43.132993711+08:00",
                        "Seller": 1,
                        "Buyer": 0,
                        "Type": 1,
                        "Name": "es",
                        "MinPrice": 100,
                        "IsFixed": false,
                        "Description": "",
                        "Status": 1,
                        "BuyerFee": 0,
                        "SellerFee": 0
                    },
                    {
                        "ID": 2,
                        "CreatedAt": "2019-11-22T19:36:43.133792914+08:00",
                        "UpdatedAt": "2019-11-22T19:36:43.133792914+08:00",
                        "EndAt": "2019-11-23T19:36:43.132993711+08:00",
                        "Seller": 1,
                        "Buyer": 0,
                        "Type": 1,
                        "Name": "es00000",
                        "MinPrice": 100,
                        "IsFixed": false,
                        "Description": "",
                        "Status": 1,
                        "BuyerFee": 0,
                        "SellerFee": 0
                    }
                ]
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.DPJQE0fV4oqeHVgNQO_LAzTPvPZjV9wp3HS2pEHTuvM
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
                        "ID": 3,
                        "CreatedAt": "2019-11-22T19:36:43.134157962+08:00",
                        "UpdatedAt": "2019-11-22T19:36:43.134157962+08:00",
                        "EndAt": "2019-11-23T19:36:43.132993711+08:00",
                        "Seller": 1,
                        "Buyer": 0,
                        "Type": 1,
                        "Name": "es00001",
                        "MinPrice": 100,
                        "IsFixed": false,
                        "Description": "",
                        "Status": 1,
                        "BuyerFee": 0,
                        "SellerFee": 0
                    }
                ]
            }


## base-service.(*CRUDService).Delete-fm [/v1/goods/:goid]


 + DELETE: 
 + PUT: 
 + GET: 


### base-service.(*CRUDService).Get-fm [GET]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.DPJQE0fV4oqeHVgNQO_LAzTPvPZjV9wp3HS2pEHTuvM
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
                    "CreatedAt": "2019-11-22T19:36:43.133268781+08:00",
                    "UpdatedAt": "2019-11-22T19:36:43.133268781+08:00",
                    "EndAt": "2019-11-23T19:36:43.132993711+08:00",
                    "Seller": 1,
                    "Buyer": 0,
                    "Type": 1,
                    "Name": "es",
                    "MinPrice": 100,
                    "IsFixed": false,
                    "Description": "",
                    "Status": 1,
                    "BuyerFee": 0,
                    "SellerFee": 0
                }
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.DPJQE0fV4oqeHVgNQO_LAzTPvPZjV9wp3HS2pEHTuvM",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUwMjc0MDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjF9fQ.mk8wYK98cGtujfJvF8DzTj_waVJPSydSGOfSFrubsnA"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.DPJQE0fV4oqeHVgNQO_LAzTPvPZjV9wp3HS2pEHTuvM
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.Rqn4qIVDX0VItxDpfUGn5zhqYYR-EUx8-lusGgao-Hc",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUwMjc0MDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.1FA-Yo6gAmz4jXcK_pUvo4SI3_6492OTKVthfu5DMg8"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.DPJQE0fV4oqeHVgNQO_LAzTPvPZjV9wp3HS2pEHTuvM
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.Rqn4qIVDX0VItxDpfUGn5zhqYYR-EUx8-lusGgao-Hc",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUwMjc0MDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.1FA-Yo6gAmz4jXcK_pUvo4SI3_6492OTKVthfu5DMg8"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.DPJQE0fV4oqeHVgNQO_LAzTPvPZjV9wp3HS2pEHTuvM
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.Rqn4qIVDX0VItxDpfUGn5zhqYYR-EUx8-lusGgao-Hc",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzUwMjc0MDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.1FA-Yo6gAmz4jXcK_pUvo4SI3_6492OTKVthfu5DMg8"
            }


## base-service.(*CRUDService).Post-fm [/v1/needs]


 + POST: 


## base-service.(*ListService).List-fm [/v1/needs-list]


 + GET: 


## base-service.(*CRUDService).Put-fm [/v1/needs/:nid]


 + PUT: 
 + GET: 
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.DPJQE0fV4oqeHVgNQO_LAzTPvPZjV9wp3HS2pEHTuvM
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.DPJQE0fV4oqeHVgNQO_LAzTPvPZjV9wp3HS2pEHTuvM
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
                        "created_at": "2019-11-22T19:36:43.052672744+08:00",
                        "updated_at": "2019-11-22T19:36:43.052672744+08:00",
                        "last_login": "2019-11-22T11:36:43Z",
                        "NickName": "admin_context",
                        "Name": "admin_context",
                        "Password": "$2a$10$2Sfxzd1fwhw.pzy0ve1x2.emRMDHUTYoRRM7u1M0tq/REnL1UAyMK",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ0MjYyMDMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0NDIyNTkzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.DPJQE0fV4oqeHVgNQO_LAzTPvPZjV9wp3HS2pEHTuvM
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
                    "created_at": "2019-11-22T19:36:43.052672744+08:00",
                    "updated_at": "2019-11-22T19:36:43.052672744+08:00",
                    "last_login": "2019-11-22T11:36:43Z",
                    "NickName": "admin_context",
                    "Name": "admin_context",
                    "Password": "$2a$10$2Sfxzd1fwhw.pzy0ve1x2.emRMDHUTYoRRM7u1M0tq/REnL1UAyMK",
                    "Phone": "1234567891011",
                    "RegisterCity": "Qing Dao S.D."
                }
            }


## ChangePassword [/v1/user/:id/password]


 + POST: change password of user

