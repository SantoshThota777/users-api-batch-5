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

- Install Docker to run the postgres DB as a container

- Spin up the DB container by using

  ```Shell
  docker-compose up
  ```

- Install any SQL client to work with DB, ex: PgAdmin4

- Create the table using the scripts in `./db`

- Get all the packages being used

  ```Shell
  go mod download
  ```

- To run the program

  ```Shell
  go run .\src\main.go
  ```

- Download the `REST Client` extention for VSCode

- Validate the end points using `./http_client/samples.http`

## Assignment-1

- Add an endpoint to delete the user by id from userDB (in-memory map)

## Assignment-2

- Add email validatin while creating the user
- User environment variables for passing the DB password
