# This is me - 01
gRPC Server with standard REST API methods (CRUD):
- Create.
- Get (Read through 'get' method).
- Update.
- Delete.
- List (Read through 'post' method).

### How to test it
Execute, on project's root parent directory:
- Launch a dockerized Postgres DB service ('-d' flag means a background running):
    > `$ docker-compose build && docker-compose up -d`
- Launch go gRPC server main:
    > `$ go run cmd/server/main.go`
- Test gRPC methods with 'Postman', starting a new 'gRPC Request'
- Host direction is 'localhost:5070', and use reflection to fetch available methods.

After you're done:
- Stop go execution with 'ctrl+c' signal.
- Stop dockerized postgres db service with ('--rmi local' flag means delete built
images after stop):
    > `$ docker-compose down --rmi local`
