# Users API (Batch 5)

RESTFul webservice developed using Go

## Different use cases

- Create User
- Read / List Users
- Update User
- Delete User

## Packages

- gorilla/mux

## Development instructions

- Clone the repository using ssh url (With Ssh key configured in your machine)

- Get all the packages being used
    ``` Shell
        go mod download
    ```

- To run the program
    ``` Shell
    go run .\src\main.go
    ```

- Download the `REST Client` extention for VSCode

- Validate the end points using `./http_client/samples.http`
