## What is this repo?
This is a simple implementation of JWT Authorization. There are gin server with mongoDB and test api. 
You can sig-up with your name and password and then test your access token.

## How to use it?
Run from root directory:
    ```shell
    make run-db
    ```
Then:
    ```shell
    make run
    ```
   
You can use PostMan to check if it works: send get request on test/hello with correct token and you will get message.
To generate token you need to sign-up using /auth/sign-up with body:

    ```

    {
    "name": "Name",
    "username": "Username",
    "password": "qwerty"
    }
    ```
Then use /auth/sign-in endpoint with same username and password and you will get your access token.

Stop db:
    ```
    make stop-db
    ```

## Is it done?
No, not yet. Refresh token generated but never used. There is no /refresh endpoint yet.