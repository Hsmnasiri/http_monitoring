# HTTP Monitoring

HTTP Monitoring is a Go application for tracking the availability and response codes of HTTP endpoints. It exposes a REST API for managing users, registering URLs to monitor, and storing the results of each request. The service can run locally or in Docker alongside a PostgreSQL database and an optional Nginx proxy.

## Features
- User registration and login with JWT-based authentication.
- CRUD endpoints for managed URLs and users.
- Records responses for monitored URLs and keeps a history of calls.
- Thresholds that can trigger alarms when consecutive failures exceed a limit.

## Project Layout
- `server/` – Go API service built with Gorilla Mux and GORM.
- `nginx/` – Nginx reverse proxy forwarding requests to the API.
- `database/` – Stand‑alone PostgreSQL compose file for development.
- `docker-compose.yml` – Runs PostgreSQL, the API service and Nginx together.
- `API-POSTMAN.json` – Postman collection with example requests.

## Environment Variables
The API reads configuration from the environment (or a `.env` file):

```
DB_DRIVER
DB_USER
DB_PASSWORD
DB_PORT
DB_HOST
DB_NAME
API_SECRET
```

If no environment file is found, default PostgreSQL credentials are used and the API listens on port `8080`.

## Running with Docker
```bash
docker-compose up --build
```
This starts PostgreSQL, the API service and Nginx. The API becomes available via the proxy on [http://localhost:8080](http://localhost:8080).

## Running Locally
```bash
cd server
go run main.go
```
Ensure a PostgreSQL instance is running and the environment variables above are set.

## API Overview
Key endpoints:

- `GET /` – Home route
- `POST /login` – Obtain a JWT token
- CRUD on `/users`
- CRUD on `/urls`
- `/calls` for recording and querying endpoint call details
- `POST /callsByTime` to retrieve call history within a time range

Refer to the Postman collection for sample requests and responses.

## License
This project is provided as-is without an explicit license.

