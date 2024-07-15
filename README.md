# Go Encrypt (AES-256)

This project is a Go application for encryption and decryption. It includes utility functions for encryption and decryption, and a main application entry point. It also includes a Dockerfile for containerization.

## Table of Contents
- [Go Encrypt (AES-256)](#go-encrypt-aes-256)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Docker](#docker)
  - [Executable](#executable)

## Installation

To install and run this application locally, ensure you have Go installed on your machine. Follow the steps below:

1. **Clone the repository:**

    ```sh
    git clone <repository-url>
    cd go-encrypt
    ```

2. **Install dependencies:**

    ```sh
    go mod tidy
    ```

## Usage

The main entry point of the application is in `main.go`. To run the application:

```sh
go run ./cmd/main.go
```

## Docker

You can directly pull the image by this command : 
```sh 
docker pull ocharless/go-encrypt:1
```

Then to run it, simply run : 

```sh
docker run -it ocharless/go-encrypt:1
```


## Executable

To embed the encryption and decryption functions in an executable, run the following command:

```sh
go build -o main ./cmd
```

This will create an executable named `main` in the root directory. To run the executable:

```sh
./main
```

