# RESTful API User And Book

## Prerequisite

- [Go 1.18](https://golang.org/dl/)
- [GORM](https://gorm.io/index.html)
- [Echo Framework](https://echo.labstack.com/)

## Run the API Server

- Run the server

  ```shell
  go run main.go
  ```

- Run the Go Test

  - Uncomment DB DSN in config/config.go
  - Comment line code `dsnMaster := os.Getenv("DB_DSN")`
  - Go to folder handlers:

    ```shell
    cd handlers
    ```

  - Run command below for running the test:

    ```shell
    go test
    ```
