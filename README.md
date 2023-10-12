# Golang Simple Clean REST API

A simple REST API built with Golang, Echo, Gorm, and following Clean Architecture principles. It includes JWT authentication and connects to a MySQL database.

## Technologies Used:
- Golang
- Echo (Web Framework)
- Gorm (ORM for Golang)
- MySQL
- Clean Architecture
- JWT

## Environment Variables:
Ensure the following environment variables are set before running the application:

- `DB_HOST`: MySQL database host
- `DB_PORT`: MySQL database port
- `DB_USERNAME`: MySQL database username
- `DB_PASSWORD`: MySQL database password
- `DB_NAME`: MySQL database name
- `JWT_SECRET`: Secret key for JWT authentication

## How to Run Example:

### Local
1. Create .env file, with environment variables
2. Start Database
3. Run the Project using ``` go run . ```

### With Docker 
Run the Docker container with the provided example command:

```bash
docker run --name go-simple-clean-rest-api -p 8080:8080 -d \
  -e DB_HOST=localhost \
  -e DB_PORT=3306 \
  -e DB_NAME=go-simple-clean-rest-api \
  -e DB_USERNAME=root \
  -e DB_PASSWORD=12345678 \
  -e JWT_SECRET=kayokoonikata \
  aszaychik/go-simple-clean-rest-api:latest
```

## Endpoints

### Get All Users
- **Method:** GET
- **URL:** `http://localhost:8080/users`
- **Auth:** `Bearer Token required`

### Get User by Id
- **Method:** GET
- **URL:** `http://localhost:8080/users/:id`
- **Auth:** `Bearer Token required`

### Create User
- **Method:** POST
- **URL:** `http://localhost:8080/users/`
- **Body:**
  ```json
  {
      "name": "AsZ",
      "email": "asz@gmail.com",
      "password": "12345678"
  }

### Update User
- **Method:** PUT
- **URL:** `http://localhost:8080/users/:id`
- **Auth:** `Bearer Token required`
- **Body:**
  ```json
  {
      "name": "AsZ",
      "email": "asz@gmail.com",
      "password": "12345678"
  }

### Delete User
- **Method:** DELETE
- **URL:** `http://localhost:8080/users/:id`
- **Auth:** `Bearer Token required`

### Login User
- **Method:** POST
- **URL:** `http://localhost:8080/users/login`
- **Body:**
  ```json
  {
      "email": "asz@gmail.com",
      "password": "12345678"
  }
