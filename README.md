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
- **User Login**: `POST /api/login`
- **Get Recommendations**: `GET /api/recommendations`

### Example Requests

#### User Registration

```sh
curl -X POST http://localhost:8080/api/register -d '{
    "email": "user@example.com",
    "password": "Password123!"
}'
```

#### User Login

```sh
curl -X POST http://localhost:8080/api/login -d '{
    "email": "user@example.com",
    "password": "Password123!"
}'
```

#### Get Recommendations

```sh
curl -X GET http://localhost:8080/api/recommendations -H "Authorization: Bearer <access_token>"
```

## Testing

Run the tests using:
```sh
go test ./...
```

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature-branch`).
5. Create a new Pull Request.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.