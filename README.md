# Banking Application

## Overview

This is a simple bank application with basic feature like:
- Check Balance
- Calculate compound interest on a saving accounts
- Deposit and Withdray Money
- Interest management _*for admin/superuser only_

This application implements End-to-End Encryption (E2EE) for client-server communication. As a result, some requests and responses are in encrypted format. 

You can use the ***sandbox API*** to simulate E2EE for encryption and decryption purposes.

## Security Features

- **Input Validation:** Validates user inputs to prevent malicious input.
- **Data Encryption:** Uses RSA encryption as **E2EE**. Encrypts sensitive data to protect against unauthorized access.
- **Protection Against Injection Attacks:** Uses parameterized queries to prevent SQL injection.
- **Authentication and Authorization:** Implements JWT-based authentication and role-based authorization.
- **Transaction Integrity:** Ensures consistency and integrity of financial transactions.

## Project Structure
```
bank-app/
├── cmd/
│   ├── main.go
├── common/
│   ├── auth/
│   ├── const/
│   ├── encryption/
│   ├── model/
│   ├── util/
│   ├── validation/
├── domain/
├── middleware/
├── pkg/
│   ├── config/
├── server/
│   ├── http/
│   ├── init.go
├── usecase/
├── config.json
```
#### Structure Desctiption
- **`cmd/`**: Entry point for the application.
  - **`main.go`**: The main application file where execution begins.

- **`common/`**: Common utilities and shared components.
  - **`auth/`**: Authentication-realted utilities and services.
  - **`const/`**: Constant values used across the application.
  - **`encryption/`**: Encryption-related utilities and services.
  - **`model/`**: Data models used throughout the application.
  - **`util/`**: General utility functions.
  - **`validation/`**: Input validation utilities.

- **`domain/`**: Related to Database and external API.

- **`middleware/`**: Middleware components for handling HTTP requests and responses, such as logging, authentication, and error handling.

- **`pkg/`**: External libraries and packages that provide additional functionalities.
  - **`config/`**: Configuration management utilities and services.

- **`server/`**: Server-related configurations and route handlers.
  - **`http/`**: HTTP server handlers and routes.
  - **`init.go`**: Define initial setup tasks that need to be performed when application starts.

- **`usecase/`**: Application-specific use cases and business logic.

# Getting Started

### Prerequisites
- [Go](https://golang.org/dl/) 1.16 or higher
- Postgresql
- Postman, API collection can be found by copy this [LINK](https://elements.getpostman.com/redirect?entityId=20168273-14ed22d6-fd42-4f2c-bbf0-bc9e607c3fac&entityType=collection) then import in your postman.

## Preparation

1. Create your database and table for users, accounts, transactions, and interest.
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL UNIQUE,
    password TEXT NOT NULL,
    role INT NOT NULL,
    public_key text,
    private_key text,
    created_at TIMESTAMP DEFAULT (timezone('utc', now()))
);

CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    account_number INT not null,
    balance NUMERIC(12, 3) NOT NULL,
    created_at TIMESTAMP DEFAULT (timezone('utc', now()))
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    amount NUMERIC(12, 3) NOT NULL,
    type INT NOT NULL, -- 'deposit' or 'withdrawal'
    created_at TIMESTAMP DEFAULT (timezone('utc', now()))
);

CREATE TABLE interest_rates (
    id SERIAL PRIMARY KEY,
    rate NUMERIC(5, 2) NOT NULL,
    compound INT not null
);
```

2. Seed table user for admin and regular user, and interest information
```sql
-- admin
INSERT INTO users (username, password, role)
VALUES ('your_desired_username', 'your_desired_password', 5);

-- regular
INSERT INTO users (username, password, role)
VALUES ('your_desired_username', 'your_desired_password', 1);

-- interest
INSERT INTO interest_rates (rate, compound)
VALUES (0, 0); -- define your own value

-- rate in percentage, e.g 5.1 means 5.1%

-- compound value: 
-- 1 = Annually
-- 2 = Semi-Annually
-- 4 = Quarterly
-- 12 = Monthly
-- 365 = Daily
```

2. Set environment variables in config.json
```json
{
    "db": {
        "user": "",
        "password": "",
        "name": "",
        "host": "",
        "port": ""
    },
    "jwt": {
        "secret": ""
    },
    "rsa-key": { // the service key for encryption and decryption
        "public": "",
        "secret": ""
    }
}
```

## Installation
1. Clone the repository:
    ```sh
    git clone git@github.com:azizyogo/bank-app.git
    ```
2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Running the application
    ```sh
    go run cmd/main.go
    ```

## Keys of Improvement

- One-time encryption code
    - option 1: Add identifier in request body e.g UUID, this can be used as a token.
- Better Input Validation