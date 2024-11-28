# PalUp_BE_Test

A RESTful API built in Go that converts an array from the request body into a linked list. This repository demonstrates the implementation of HTTP request handlers, middleware, and unit testing for API endpoints. [👉 coding question](https://github.com/binghuan/PalUp_BE_Test/blob/main/PalUp_BE_Test_readMe.MD)

- [PalUp\_BE\_Test](#palup_be_test)
  - [Features](#features)
  - [Endpoints](#endpoints)
    - [`POST /test-1`](#post-test-1)
      - [Description](#description)
      - [Request](#request)
      - [Response](#response)
  - [Project Structure](#project-structure)
  - [Prerequisites](#prerequisites)
  - [Setup and Usage](#setup-and-usage)
    - [1. Clone the Repository](#1-clone-the-repository)
    - [2. Initialize Go Modules](#2-initialize-go-modules)
    - [3. Run the Server](#3-run-the-server)
    - [4. Run Tests](#4-run-tests)
  - [Middleware](#middleware)
  - [Development and Testing](#development-and-testing)
    - [Add New Features](#add-new-features)
    - [Testing](#testing)


## Features

- Middleware for validating API keys.
- Converts an array from JSON input into a linked list.
- Returns the linked list structure as JSON in the response.
- Comprehensive test coverage for the main functionality.

## Endpoints

### `POST /test-1`

#### Description
This endpoint accepts a JSON payload with an array and converts it into a linked list.

#### Request

**Headers**:
- `Content-Type: application/json`
- `api-key`: The API key for authentication. The key must match the constant `APIKey` defined in the code.

**Body**:
```json
{
  "Array": ["a", "b", "c", "d", "e"]
}
```

#### Response

**Status Code**: 200 OK

**Body**:
```json
{
  "value": "a",
  "next": {
    "value": "b",
    "next": {
      "value": "c",
      "next": {
        "value": "d",
        "next": {
          "value": "e",
          "next": null
        }
      }
    }
  }
}
```

---

**Error Responses**:
- `400 Bad Request`: Invalid JSON format or missing "Array" field.
- `401 Unauthorized`: Missing or invalid `api-key`.

---

## Project Structure

```
.
├── InfraDesign.md     # Documentation for infrastructure design
├── NoSQL.md           # Notes or details related to NoSQL usage
├── SQL.md             # Notes or details related to SQL usage
├── go.sum             # Go dependencies
├── mail.png           # Placeholder image or documentation graphic
├── main.go            # Main entry point for the application
├── readMe.MD          # Documentation (you are reading it)
├── test1.go           # Implementation of the Test1 handler logic
├── main_test.go       # Unit tests for the API handlers
```

## Prerequisites

- Go 1.20+ installed on your system.

## Setup and Usage

### 1. Clone the Repository

```bash
git clone https://github.com/binghuan/PalUp_BE_Test.git
cd PalUp_BE_Test
```

### 2. Initialize Go Modules

```bash
go mod init palup_be_test
go mod tidy
```

### 3. Run the Server

```bash
go run main.go
```

The server will start at `http://localhost:8082`.

### 4. Run Tests

```bash
go test -v
```

## Middleware

The application uses a middleware function `ValidateAPIKeyMiddleware` to validate the `api-key` header before processing requests.

## Development and Testing

### Add New Features
1. Define the new endpoint in `main.go`.
2. Implement the logic in a separate file (e.g., `feature.go`).
3. Add tests in the corresponding `_test.go` file.

### Testing

All tests are located in `main_test.go`. Run the tests with:

```bash
go test ./...
```

---

