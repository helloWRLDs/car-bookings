# Car Booking

## Installation
- Clone repository
    ```bash
    git clone https://github.com/helloWRLDs/car-bookings.git
    cd car-bookings 
    ```
- Install dependencies
    ```bash
    go mod tidy
    ```
- Install goose for migrations
    ```bash
    go install github.com/pressly/goose/v3/cmd/goose@v3.19.2
    ```

## Run
- Start your db or use [docker-compose](./docker-compose.yml)
    ```bash
    docker-compose up -d
    ```
- Build application
    ```bash
    go build ./cmd/api
    ```
and run api.exe

## Migrate
migrations are listed in [makefile](./Makefile)
```bash
make migrate-up
```
```bash
make migrate-down
```

## Endpoints
- GET http://localhost:9090/api/v1/cars?offset=5&limit=5&sort=id&order=asc&color=Blue
- GET http://localhost:9090/api/v1/cars/1
- POST http://localhost:9090/api/v1/cars
- DELETE http://localhost:9090/api/v1/cars/2
- PUT http://localhost:9090/api/v1/cars/3

