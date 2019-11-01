FORMAT: 1A
HOST: 127.0.0.1

# Minimum-Market

method description

## control.UserService.Get-fm [/v1/user/:id]

### Get
// [GET]

Get User

+ Request

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2MjQzMTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjIwNzA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.burV0LOHCU4_-4EEi0B7RrOw7qht-ubjj6G3Xl1IHe0
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
                    "created_at": "2019-11-01T23:05:16.0974441+08:00",
                    "updated_at": "2019-11-01T23:05:16.0974441+08:00",
                    "last_login": "2019-11-01T15:05:16Z",
                    "NickName": "admin_context",
                    "Name": "admin_context",
                    "Password": "$2a$10$AB/rAFCLzjfVEYfmC7VleuJcdmerIbqy9tSEGSceeTLe0h0eIvDfe",
                    "Phone": "1234567891011",
                    "RegisterCity": "Qing Dao S.D."
                }
            }



## control.UserService.Register-fm [/v1/user]

### Register
// [POST]

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

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2MjQzMTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjIwNzA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.burV0LOHCU4_-4EEi0B7RrOw7qht-ubjj6G3Xl1IHe0
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



## control.UserService.List-fm [/v1/user-list]

### List
// [GET]

List User

+ Request

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2MjQzMTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjIwNzA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.burV0LOHCU4_-4EEi0B7RrOw7qht-ubjj6G3Xl1IHe0
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
                        "created_at": "2019-11-01T23:05:16.0974441+08:00",
                        "updated_at": "2019-11-01T23:05:16.0974441+08:00",
                        "last_login": "2019-11-01T15:05:16Z",
                        "NickName": "admin_context",
                        "Name": "admin_context",
                        "Password": "$2a$10$AB/rAFCLzjfVEYfmC7VleuJcdmerIbqy9tSEGSceeTLe0h0eIvDfe",
                        "Phone": "1234567891011",
                        "RegisterCity": "Qing Dao S.D."
                    }
                ]
            }



## control.UserService.Login-fm [/v1/login]

### Login
// [POST]

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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2MjQzMTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjIwNzA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.burV0LOHCU4_-4EEi0B7RrOw7qht-ubjj6G3Xl1IHe0",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzMyMjU1MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjIwNzA2LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzI2MjQzMTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjIwNzA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjF9fQ.xIAZ6L-zmweiMC3eeFivE05IJr_cutXicnRXlHAFrsk"
            }



+ Request

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2MjQzMTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjIwNzA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.burV0LOHCU4_-4EEi0B7RrOw7qht-ubjj6G3Xl1IHe0
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2MjQzMTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjIwNzA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.v7nDS0lgPvk_O8VMSO42j5C05RWJuFRkser07mPs3Qw",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzMyMjU1MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjIwNzA2LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzI2MjQzMTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjIwNzA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.V2Sef7UtPr74Pa45ZH6058vrPsT2kFNiXbaE9deiuRA"
            }



+ Request

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2MjQzMTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjIwNzA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.burV0LOHCU4_-4EEi0B7RrOw7qht-ubjj6G3Xl1IHe0
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2MjQzMTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjIwNzA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.v7nDS0lgPvk_O8VMSO42j5C05RWJuFRkser07mPs3Qw",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzMyMjU1MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjIwNzA2LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzI2MjQzMTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjIwNzA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.V2Sef7UtPr74Pa45ZH6058vrPsT2kFNiXbaE9deiuRA"
            }



+ Request

    + Headers

            Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2MjQzMTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjIwNzA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6MX19.burV0LOHCU4_-4EEi0B7RrOw7qht-ubjj6G3Xl1IHe0
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
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzI2MjQzMTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjIwNzA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19.v7nDS0lgPvk_O8VMSO42j5C05RWJuFRkser07mPs3Qw",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzMyMjU1MTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjIwNzA2LCJJc1JlZnJlc2hUb2tlbiI6dHJ1ZSwiUmVmcmVzaFRhcmdldCI6eyJleHAiOjE1NzI2MjQzMTYsImlzcyI6Ik15cmlhZC1EcmVhbWluIiwibmJmIjoxNTcyNjIwNzA2LCJJc1JlZnJlc2hUb2tlbiI6ZmFsc2UsIlJlZnJlc2hUYXJnZXQiOm51bGwsIkN1c3RvbUZpZWxkIjp7IlVJRCI6Mn19LCJDdXN0b21GaWVsZCI6eyJVSUQiOjJ9fQ.V2Sef7UtPr74Pa45ZH6058vrPsT2kFNiXbaE9deiuRA"
            }


