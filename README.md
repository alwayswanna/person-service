# Person Service (Go)

A lightweight REST API for managing person records built with Go and PostgreSQL.

## Features

- **CRUD Operations**: Create, Read, Update, Delete person records
- **PostgreSQL Integration**: Using `pgx` driver
- **Docker Support**: Containerized app + database
- **Clean Architecture**: Separated layers (handlers, services, repositories)
- **Environment Configuration**: Using `cleanenv`
- **Migrations**: Database schema management

## Structure

1. app - backend module
2. app-vue - simple client