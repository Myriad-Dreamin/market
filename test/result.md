FORMAT: 1A
HOST: 127.0.0.1

# Minimum-Market
this is the market backend powered by minimum

## router.NewRootRouter.func1 [/ping]


 + GET: 


## base-service.(*CRUDService).Post-fm [/v1/goods]


 + POST: 


## base-service.(*ListService).List-fm [/v1/goods-list]


 + GET: 


## base-service.(*CRUDService).Put-fm [/v1/goods/:goid]


 + PUT: 
 + DELETE: 
 + GET: 


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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2NzM3MTQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjcwMTA0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.9YzAoUS1ut-1-hfwAcQ7OhkP5UhYtZ32yrIkQ-uwpNs",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzMyNzQ5MTQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjcwMTA0LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzI2NzM3MTQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjcwMTA0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjF9fQ.Ar3BDt9GhwM-bcTLXIRzKDtVmjopu1KhQdl6lhCTee8"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2NzM3MTQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjcwMTA0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.9YzAoUS1ut-1-hfwAcQ7OhkP5UhYtZ32yrIkQ-uwpNs
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2NzM3MTQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjcwMTA0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.Z154ziyd7WZhFli3G820e1RXw-l3u68y4MxdzN3sQt8",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzMyNzQ5MTQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjcwMTA0LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzI2NzM3MTQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjcwMTA0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.IHxd_35Pdi27LpOLhOmc6lmRYOmP0EyvOpQON14msos"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2NzM3MTQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjcwMTA0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.9YzAoUS1ut-1-hfwAcQ7OhkP5UhYtZ32yrIkQ-uwpNs
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2NzM3MTQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjcwMTA0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.Z154ziyd7WZhFli3G820e1RXw-l3u68y4MxdzN3sQt8",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzMyNzQ5MTQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjcwMTA0LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzI2NzM3MTQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjcwMTA0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.IHxd_35Pdi27LpOLhOmc6lmRYOmP0EyvOpQON14msos"
            }


+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2NzM3MTQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjcwMTA0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.9YzAoUS1ut-1-hfwAcQ7OhkP5UhYtZ32yrIkQ-uwpNs
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2NzM3MTUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjcwMTA1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.VN4JNNUUcagKtaKXvGTp__ZPq7pv7P4Ow7_T1AKks-I",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzMyNzQ5MTUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjcwMTA1LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzI2NzM3MTUsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjcwMTA1LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.iyEdH2ko8NwMWEHv0U66lWgYbIoXn3pqO0D-NjgtLsM"
            }


## base-service.(*CRUDService).Post-fm [/v1/needs]


 + POST: 


## base-service.(*ListService).List-fm [/v1/needs-list]


 + GET: 


## base-service.(*CRUDService).Delete-fm [/v1/needs/:nid]


 + DELETE: 
 + PUT: 
 + GET: 


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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2NzM3MTQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjcwMTA0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.9YzAoUS1ut-1-hfwAcQ7OhkP5UhYtZ32yrIkQ-uwpNs
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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2NzM3MTQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjcwMTA0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.9YzAoUS1ut-1-hfwAcQ7OhkP5UhYtZ32yrIkQ-uwpNs
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
                        "created_at": "2019-11-02T12:48:34.666321642+08:00",
                        "updated_at": "2019-11-02T12:48:34.666321642+08:00",
                        "last_login": "2019-11-02T04:48:34Z",
                        "NickName": "admin_context",
                        "Name": "admin_context",
                        "Password": "$2a$10$sqauHEoB41DxIWMiSMSJ4ezQCJph9x9fLPUgwWN1U7yBMscJW5yta",
                        "Phone": "1234567891011",
                        "RegisterCity": "Qing Dao S.D."
                    }
                ]
            }


## Put [/v1/user/:id]


 + PUT: Put User
 + DELETE: Delete User
 + GET: Get User


### Get [GET]

+ Request 

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2NzM3MTQsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjcwMTA0LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.9YzAoUS1ut-1-hfwAcQ7OhkP5UhYtZ32yrIkQ-uwpNs
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
                    "created_at": "2019-11-02T12:48:34.666321642+08:00",
                    "updated_at": "2019-11-02T12:48:34.666321642+08:00",
                    "last_login": "2019-11-02T04:48:34Z",
                    "NickName": "admin_context",
                    "Name": "admin_context",
                    "Password": "$2a$10$sqauHEoB41DxIWMiSMSJ4ezQCJph9x9fLPUgwWN1U7yBMscJW5yta",
                    "Phone": "1234567891011",
                    "RegisterCity": "Qing Dao S.D."
                }
            }


## ChangePassword [/v1/user/:id/password]


 + POST: change password of user

