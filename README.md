# ChrisMuladiRianto CRUD Test Assessment

This API service is for my Tech Assessment number 2

## Table of Contents

- [Features](#features)
- [Installation](#installation)

## Features

- Fully integrated Account permission (User, Role, Permission, etc)
- CRUD for example table (clients)
- Using GIN, swagger, and ORM
- Example of Domain-Driven Design standard that I created for standardization and maintainable purposes

## Installation

The installation provided will be designated to Mac machines
1. Install Swagger
    - Please follow installation on this link https://github.com/swaggo/swag
2. Install Make command
    - `xcode-select --install`
    - In the windows that pops up, click `Install`, and agree to the Terms of Service.
3. Install MySQL
    - Download MySQL from this link https://dev.mysql.com/downloads/mysql/
    - create user and database based on file `config.toml` for easier setup
4. Run the following commands
    ```
    make init
    make tidy
    ```
5. Run with `make run`
