

This codebase was created to demonstrate a fully fledged fullstack application built with **Golang/Gin** including CRUD operations, authentication, routing, pagination, and more.


* The project uses:
    * [gin-gonic/gin](https://github.com/gin-gonic/gin) as the web framework
    * [jackc/pgx](https://github.com/jackc/pgx) as the database driver
    * [kyleconroy/sqlc](https://github.com/kyleconroy/sqlc) to generate Go code from SQL queries
    * [golang-migrate/migrate](https://github.com/golang-migrate/migrate) to manage database migrations
    * [golang-jwt/jwt](https://github.com/golang-jwt/jwt) for authentication
    * [spf13/viper](https://github.com/spf13/viper) for configuration
    * [rs/xid](https://github.com/rs/xid) for generating UUIDs
    * [stretchr/testify](https://github.com/stretchr/testify) for testing
    * [golang/mock](https://github.com/golang/mock) for mocking dependencies
* Also uses:
    * *PostgreSQL* for the database
# Getting started 
Running the project locally:

```
    docker-compose up -d // starts postgres container
    go run main.go
```

#Run
```
    make test
```

# Unit testing
```
    make test
```

# TODO
* [x] Add unit tests for handlers (... in progress)
* [ ] Improve error handling and logging
* [x] Improve deployment and testing configuration (Makefile, docker-compose, etc.) ( ... in progress)

