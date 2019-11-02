FORMAT: 1A
HOST: 127.0.0.1

# Minimum-Market

this is the market backend powered by minimum


## router.NewRootRouter.func1 [/ping]



## base-service.(*CRUDService).Post-fm [/v1/goods]



## base-service.(*ListService).List-fm [/v1/goods-list]



## base-service.(*CRUDService).Put-fm [/v1/goods/:goid]



## Login [/v1/login]



### Login [POST]

Login

+ Request

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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2NzA5OTgsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjY3Mzg4LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.NVtYGeCI6hk5lSawil6AA8bnhV7S_EN7oZqGj5mpXZY",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzMyNzIxOTgsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjY3Mzg4LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzI2NzA5OTgsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjY3Mzg4LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjF9fQ.KjDNhW0R8Gm80IS20EvddVh_r6QPnHLPQ-lWQ8d8gnk"
            }



+ Request

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2NzA5OTgsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjY3Mzg4LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.NVtYGeCI6hk5lSawil6AA8bnhV7S_EN7oZqGj5mpXZY
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2NzA5OTgsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjY3Mzg4LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.ib_tW6mV9o3-hxZfpCxK-yJcyhVVDqRlTbZjvpfLTfg",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzMyNzIxOTgsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjY3Mzg4LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzI2NzA5OTgsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjY3Mzg4LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.lZNRRbbdi-1SigBBh_-qx5J9WvQhPz0l_IDgoV-cRnw"
            }



+ Request

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2NzA5OTgsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjY3Mzg4LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.NVtYGeCI6hk5lSawil6AA8bnhV7S_EN7oZqGj5mpXZY
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2NzA5OTgsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjY3Mzg4LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.ib_tW6mV9o3-hxZfpCxK-yJcyhVVDqRlTbZjvpfLTfg",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzMyNzIxOTgsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjY3Mzg4LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzI2NzA5OTgsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjY3Mzg4LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.lZNRRbbdi-1SigBBh_-qx5J9WvQhPz0l_IDgoV-cRnw"
            }



+ Request

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2NzA5OTgsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjY3Mzg4LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.NVtYGeCI6hk5lSawil6AA8bnhV7S_EN7oZqGj5mpXZY
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2NzA5OTgsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjY3Mzg4LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.ib_tW6mV9o3-hxZfpCxK-yJcyhVVDqRlTbZjvpfLTfg",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzMyNzIxOTgsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjY3Mzg4LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzI2NzA5OTgsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjY3Mzg4LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.lZNRRbbdi-1SigBBh_-qx5J9WvQhPz0l_IDgoV-cRnw"
            }



## base-service.(*CRUDService).Post-fm [/v1/needs]



## base-service.(*ListService).List-fm [/v1/needs-list]



## base-service.(*CRUDService).Delete-fm [/v1/needs/:nid]



## Register [/v1/user]



### Register [POST]

Register

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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2NzA5OTgsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjY3Mzg4LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.NVtYGeCI6hk5lSawil6AA8bnhV7S_EN7oZqGj5mpXZY
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



### List [GET]

List User

+ Request

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2NzA5OTgsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjY3Mzg4LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.NVtYGeCI6hk5lSawil6AA8bnhV7S_EN7oZqGj5mpXZY
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
                        "created_at": "2019-11-02T12:03:18.066890431+08:00",
                        "updated_at": "2019-11-02T12:03:18.066890431+08:00",
                        "last_login": "2019-11-02T04:03:18Z",
                        "NickName": "admin_context",
                        "Name": "admin_context",
                        "Password": "$2a$10$6.70G2oJ8Yk.uDR4LgQVY.m3EDd38RfWGUnu7RfJQ.oq6y.TiTPnm",
                        "Phone": "1234567891011",
                        "RegisterCity": "Qing Dao S.D."
                    }
                ]
            }



## Delete [/v1/user/:id]



### Get [GET]

Get User

+ Request

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2NzA5OTgsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjY3Mzg4LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.NVtYGeCI6hk5lSawil6AA8bnhV7S_EN7oZqGj5mpXZY
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
                    "created_at": "2019-11-02T12:03:18.066890431+08:00",
                    "updated_at": "2019-11-02T12:03:18.066890431+08:00",
                    "last_login": "2019-11-02T04:03:18Z",
                    "NickName": "admin_context",
                    "Name": "admin_context",
                    "Password": "$2a$10$6.70G2oJ8Yk.uDR4LgQVY.m3EDd38RfWGUnu7RfJQ.oq6y.TiTPnm",
                    "Phone": "1234567891011",
                    "RegisterCity": "Qing Dao S.D."
                }
            }



## ChangePassword [/v1/user/:id/password]


