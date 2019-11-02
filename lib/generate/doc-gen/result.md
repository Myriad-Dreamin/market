FORMAT: 1A
HOST: 127.0.0.1

# Maximum Document
controllerS

## Accounts [/v1/user]
controllerS

 + PUT: put accounts
 + GET: get accounts


### Put Accounts [PUT]

+ Request 

    + Headers

            Content-Type: application/json

    + Body

            {
                "1": 1
            }

+ Response 200

    + Headers

            Content-Type: application/json

    + Body

            {
                "2": 2
            }

### Get AccountsA [GET]

+ Request 

    + Headers

            Content-Type: application/json

    + Body

            {
                "1": 1
            }

+ Response 200

    + Headers

            Content-Type: application/json

    + Body

            {
                "2": 2
            }
