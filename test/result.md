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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzM5NjE5MjYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTczOTU4MzE2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.M_xBsrw4IVemKBYHl_WlN9GdtEU4NyRUJImCYbaAmD4
            Content-Type: application/json

    + Body

            {
                "end_at": "2019-11-18T10:38:46.049523821+08:00",
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
                    "CreatedAt": "2019-11-17T10:38:46.049799845+08:00",
                    "UpdatedAt": "2019-11-17T10:38:46.049799845+08:00",
                    "EndAt": "2019-11-18T10:38:46.049523821+08:00",
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


## base-service.(*ListService).List-fm [/v1/goods-list]


 + GET: 


## base-service.(*CRUDService).Get-fm [/v1/goods/:goid]


 + GET: 
 + DELETE: 
 + PUT: 


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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzM5NjE5MjYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTczOTU4MzE2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.M_xBsrw4IVemKBYHl_WlN9GdtEU4NyRUJImCYbaAmD4",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzQ1NjMxMjYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTczOTU4MzE2LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzM5NjE5MjYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTczOTU4MzE2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjF9fQ.emebmC8lscUy1iA6asPoZSpoN7J7msMcKvbLuWhRNz4"
            }


## base-service.(*CRUDService).Post-fm [/v1/needs]


 + POST: 


## base-service.(*ListService).List-fm [/v1/needs-list]


 + GET: 


## base-service.(*CRUDService).Delete-fm [/v1/needs/:nid]


 + DELETE: 
 + GET: 
 + PUT: 


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


## List [/v1/user-list]


 + GET: List User


## Get [/v1/user/:id]


 + GET: Get User
 + DELETE: Delete User
 + PUT: Put User


## ChangePassword [/v1/user/:id/password]


 + POST: change password of user

