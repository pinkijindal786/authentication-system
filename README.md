# Authentication_System

This is a RESTful Authentication System built using [Go (Golang)](https://golang.org/) and the [Gin Web Framework](https://gin-gonic.com/). It supports user registration, login, token management, and secure endpoints protected by authentication middleware.

---

## Table of Contents

1. [Features](#features)
2. [Prerequisites](#prerequisites)
3. [Installation](#installation)
4. [Usage](#usage)
5. [API Endpoints](#api-endpoints)
6. [Folder Structure](#folder-structure)
7. [Project Components](#project-components)
8. [Contributing](#contributing)
9. [License](#license)

---

## Features

- **User Registration**: Create new user accounts with email and password.
- **User Login**: Authenticate users and generate JWT tokens.
- **Token Management**:
  - Refresh expired tokens.
  - Revoke tokens.
- **Secure Endpoints**: Protect routes using middleware that validates JWTs.

---

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.20 or later recommended)
- [Docker](https://www.docker.com/) (optional, for running dependencies)
- A relational database (e.g., PostgreSQL)

---

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/your-repo/authentication-system.git
   cd authentication-system
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure the database connection in `internal/database`.

4. Run the application:
   ```bash
   go run main.go
   ```

---

## Usage

- The server runs on port `8080` by default.
- Use an API client like [Postman](https://www.postman.com/) or [curl](https://curl.se/) to interact with the endpoints.

---

## API Endpoints

### **Authentication**

| Method | Endpoint       | Description                          |
|--------|----------------|--------------------------------------|
| POST   | `/auth/signup` | Register a new user.                |
| POST   | `/auth/signin` | Log in and obtain a JWT token.       |
| POST   | `/auth/refresh`| Refresh an expired access token.     |
| POST   | `/auth/revoke` | Revoke a JWT token.                 |

### **Protected Routes**

| Method | Endpoint      | Description                       |
|--------|---------------|-----------------------------------|
| GET    | `/secure`     | Access a protected route example.|

---

## Folder Structure

```plaintext
.
├── internal
│   ├── database         # Database connection and initialization
│   ├── handlers         # HTTP request handlers
│   ├── middlewares      # Authentication and authorization middleware
│   ├── repositories     # Database interaction and queries
│   ├── services         # Business logic and token management
├── main.go              # Application entry point
├── go.mod               # Go module file
├── go.sum               # Dependency checksum file
```

---

## Project Components

### 1. **Handler Layer**
   - **File**: `internal/handlers/auth_handler.go`
   - **Purpose**: Processes HTTP requests and delegates logic to the service layer.
   - **Endpoints**:
     - `SignUp`: Registers new users.
     - `SignIn`: Authenticates users and generates JWT tokens.
     - `RefreshToken`: Renews expired tokens.
     - `RevokeToken`: Invalidates tokens.
     - `SecureEndpoint`: Demonstrates a protected endpoint.

### 2. **Service Layer**
   - Handles business logic like authentication, token management, and user validation.

### 3. **Repository Layer**
   - Interacts with the database to perform CRUD operations for users and tokens.

### 4. **Middleware**
   - Validates JWT tokens for protected routes.

### 5. **Database**
   - Manages database connections and configurations.

---

## Contributing

1. Fork the repository.
2. Create a feature branch:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add feature description"
   ```
4. Push the branch:
   ```bash
   git push origin feature-name
   ```
5. Create a pull request.

---

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.