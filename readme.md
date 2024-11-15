
```markdown
# Auth API Service

## Overview

This is an Auth REST API service implemented in Go without using an ORM. It uses MySQL as the database and is fully dockerized for easy setup.

## Requirements

- Docker and Docker Compose installed on your system.

## Setup

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/GoApis.git
   cd GoApis
   ```

2. **Create a `.env` file:**

   Use the provided `.env` template and set your own values for `DB_PASSWORD` and `JWT_SECRET`.

3. **Start the application:**

   ```bash
   docker-compose up --build
   ```

   This command builds and starts both the API service and the MySQL database.

## API Endpoints and cURL Commands

### Sign Up

- **Endpoint:** `POST /signup`

- **cURL Command:**

  ```bash
  curl -X POST http://localhost:8000/signup \
     -H 'Content-Type: application/json' \
     -d '{"email":"test@example.com", "password":"password123"}'
  ```

### Sign In

- **Endpoint:** `POST /signin`

- **cURL Command:**

  ```bash
  curl -X POST http://localhost:8000/signin \
     -H 'Content-Type: application/json' \
     -d '{"email":"test@example.com", "password":"password123"}'
  ```

### Access Protected Route

- **Endpoint:** `GET /protected`

- **cURL Command:**

  ```bash
  curl -X GET http://localhost:8000/protected \
     -H 'Authorization: Bearer YOUR_JWT_TOKEN'
  ```

### Refresh Token

- **Endpoint:** `POST /refresh`

- **cURL Command:**

  ```bash
  curl -X POST http://localhost:8000/refresh \
     -H 'Authorization: Bearer YOUR_JWT_TOKEN'
  ```

### Logout

- **Endpoint:** `POST /logout`

- **cURL Command:**

  ```bash
  curl -X POST http://localhost:8000/logout \
     -H 'Authorization: Bearer YOUR_JWT_TOKEN'
  ```

## Notes

- Replace `YOUR_JWT_TOKEN` with the token received from the Sign In endpoint.
- Ensure that the database credentials in the `.env` file match those in the `docker-compose.yml`.
- The service runs on port `8000` by default.

## Dependencies

- Go 1.20
- MySQL 8.0
- Docker and Docker Compose

```

---

## Running the Project

1. **Ensure Docker and Docker Compose are installed:**

   - [Install Docker](https://docs.docker.com/get-docker/)
   - [Install Docker Compose](https://docs.docker.com/compose/install/)

2. **Navigate to the Project Directory:**

   ```bash
   cd GoApis
   ```

3. **Start the Application:**

   ```bash
   docker-compose up --build
   ```

   - This command will build the Docker image and start the containers for the API service and the MySQL database.
   - The API service will be available at `http://localhost:8000`.

---

.Env 

```env
DB_USER=root
DB_PASSWORD=root
DB_HOST=db
DB_PORT=3306
DB_NAME=authdb
PORT=8000

JWT_SECRET=your_secret_key
```