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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQwNzM0MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.BlHL99pzkr2r4jfXdU4j92TbHtmItarb1xuONh5JBts
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-11-19T17:37:13.5904686+08:00",
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
                    "CreatedAt": "2019-11-18T17:37:13.5904686+08:00",
                    "UpdatedAt": "2019-11-18T17:37:13.5904686+08:00",
                    "EndAt": "2019-11-19T17:37:13.5904686+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQwNzM0MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.BlHL99pzkr2r4jfXdU4j92TbHtmItarb1xuONh5JBts
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-11-19T17:37:13.5914756+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQwNzM0MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.BlHL99pzkr2r4jfXdU4j92TbHtmItarb1xuONh5JBts
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-11-19T17:37:13.5914756+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQwNzM0MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.BlHL99pzkr2r4jfXdU4j92TbHtmItarb1xuONh5JBts
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-11-19T17:37:13.5914756+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQwNzM0MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.BlHL99pzkr2r4jfXdU4j92TbHtmItarb1xuONh5JBts
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-11-19T17:37:13.5914756+08:00",
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQwNzM0MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.BlHL99pzkr2r4jfXdU4j92TbHtmItarb1xuONh5JBts
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-11-18T17:37:13.5914756+08:00",
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


## base-service.(*CRUDService).Get-fm [/v1/goods/:goid]


 + GET: 
 + PUT: 
 + DELETE: 


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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQwNzM0MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.BlHL99pzkr2r4jfXdU4j92TbHtmItarb1xuONh5JBts",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ2NzQ2MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzQwNzM0MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjF9fQ.MVp-QzpelViPcpxWhdra2M6Bm4KZMPFsYJsX2Xd1qZU"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQwNzM0MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.BlHL99pzkr2r4jfXdU4j92TbHtmItarb1xuONh5JBts
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQwNzM0MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.Ptrs3BdJU1NkgGRJanKFDVjrOmuJn4cCjv18Ly-zmSY",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ2NzQ2MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzQwNzM0MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.CBi-kXxM72mkF6aEJfXeiPDGGjsGM8TJTk_aHBWox_s"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQwNzM0MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.BlHL99pzkr2r4jfXdU4j92TbHtmItarb1xuONh5JBts
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQwNzM0MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.Ptrs3BdJU1NkgGRJanKFDVjrOmuJn4cCjv18Ly-zmSY",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ2NzQ2MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzQwNzM0MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.CBi-kXxM72mkF6aEJfXeiPDGGjsGM8TJTk_aHBWox_s"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQwNzM0MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.BlHL99pzkr2r4jfXdU4j92TbHtmItarb1xuONh5JBts
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQwNzM0MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.Ptrs3BdJU1NkgGRJanKFDVjrOmuJn4cCjv18Ly-zmSY",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ2NzQ2MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzQwNzM0MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.CBi-kXxM72mkF6aEJfXeiPDGGjsGM8TJTk_aHBWox_s"
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQwNzM0MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.BlHL99pzkr2r4jfXdU4j92TbHtmItarb1xuONh5JBts
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQwNzM0MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.BlHL99pzkr2r4jfXdU4j92TbHtmItarb1xuONh5JBts
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
                        "created_at": "2019-11-18T17:37:13.5244725+08:00",
                        "updated_at": "2019-11-18T17:37:13.5244725+08:00",
                        "last_login": "2019-11-18T09:37:13Z",
                        "NickName": "admin_context",
                        "Name": "admin_context",
                        "Password": "$2a$10$j7uuXA38tfaa6slOvgfg5OaY5FxM1tIB/9HkVGZ8n76eGgpgx0fv6",
                        "Phone": "1234567891011",
                        "RegisterCity": "Qing Dao S.D."
                    }
                ]
            }


## Delete [/v1/user/:id]


 + DELETE: Delete User
 + GET: Get User
 + PUT: Put User


### Get [GET]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQwNzM0MzMsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTc0MDY5ODIzLCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.BlHL99pzkr2r4jfXdU4j92TbHtmItarb1xuONh5JBts
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
                    "created_at": "2019-11-18T17:37:13.5244725+08:00",
                    "updated_at": "2019-11-18T17:37:13.5244725+08:00",
                    "last_login": "2019-11-18T09:37:13Z",
                    "NickName": "admin_context",
                    "Name": "admin_context",
                    "Password": "$2a$10$j7uuXA38tfaa6slOvgfg5OaY5FxM1tIB/9HkVGZ8n76eGgpgx0fv6",
                    "Phone": "1234567891011",
                    "RegisterCity": "Qing Dao S.D."
                }
            }


## ChangePassword [/v1/user/:id/password]


 + POST: change password of user

