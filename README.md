# eCommerceService

This project is an eCommerce service built with Go. It includes functionalities such as user management, recommendations, and email verification.

## Prerequisites

- Go 1.23 or later
- Docker
- Docker Compose
- Redis
- MySQL

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/hhhhp52/eCommerceService.git
    cd eCommerceService
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Configuration

Update the configuration files with your database and email server details.

## Running the Application

### Using Docker Compose

1. Build and start the services:
    ```sh
    docker-compose up --build
    ```

2. The application will be available at `http://localhost:8080`.

### Without Docker

1. Start Redis and MySQL services.

2. Update the database connection details in the code.

3. Run the application:
    ```sh
    go run main.go
    ```

## Usage

### API Endpoints

- **User Registration**: `POST /api/register`
- **Verify Email**: `GET /api/verify-email`
- **User Login**: `POST /api/login`
- **Get Recommendations**: `GET /api/recommendations`

### Example Requests

#### User Registration

```sh
curl -X POST http://localhost:8080/api/register -d '{
    "email": "user@example.com",
    "password": "Password123!"
    "password_confirm": "Password123!"
}'
```

#### Verify Email

```sh 
curl -X POST http://localhost:8080/api/verify-email -d '{
    "email": "user@example.com",
    "password": "Password123!"
}'
```


#### User Login

```sh
curl -X POST http://localhost:8080/api/login -d '{
    "email": "user1@example.com",
    "password": "password1"
}'
```

#### Get Recommendations

```sh
curl -X GET http://localhost:8080/api/recommendations -H "Authorization: <access_token>"
```
