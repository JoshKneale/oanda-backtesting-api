# Oanda Backtesting API

This repository contains a mock implementation of the Oanda trading API, developed using Go. The project serves as a means for backtesting trading algorithms and strategies without interacting with the real Oanda API. The mock API provides endpoints for managing trades, orders, and positions, simulating core functionalities of the Oanda trading platform.

## Features

- Simulates core functionalities of the Oanda trading API - TODO
- Supports account and trading operations - TODO
- Provides endpoints for managing trades, orders, and positions - TODO
- Utilizes Go's standard library and common packages for HTTP handling and JSON processing - TODO
- Allows for passing of a timestamp in HTTP headers that can simulate price changes over time - TODO

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/joshkneale/oanda-backtesting-api.git
    ```
2. Navigate to the project directory:
    ```sh
    cd oanda-backtesting-api
    ```
3. Install the dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Build the server:
    ```sh
    make build
    ```
2. Run the server:
    ```sh
    make run
    ```
2. The server will start on `http://localhost:8080` (configurable by environment variables). You can use tools like `curl` or Postman to interact with the API.

## Endpoints - WIP

### Accounts

- `GET /accounts` - Retrieve a list of accounts
- `GET /accounts/{account_id}` - Retrieve details of a specific account

### Trades

- `GET /accounts/{account_id}/trades` - List all trades for an account
- `POST /accounts/{account_id}/trades` - Create a new trade

### Orders

- `GET /accounts/{account_id}/orders` - List all orders for an account
- `POST /accounts/{account_id}/orders` - Create a new order

### Positions

- `GET /accounts/{account_id}/positions` - List all positions for an account


## Directory and File Descriptions

### `cmd/`: The entry point for the application.
- **Subdirectory `api/`**
  - **`api/main.go`**: The main Go file that initializes and starts the server.

### `internal/`: Holds the internal application code.
- **Subdirectory `handlers/`**: Contains handler functions for endpoints.
- **Subdirectory `models/`**: Defines the data models used in the application.
- **Subdirectory `services/`**: Contains business logic and service layer functions.
- **Subdirectory `repository/`**: Handles database interactions.
- **Subdirectory `config/`**: Manages application configuration and environment variables.
- **Subdirectory `middleware/`**: Contains middleware functions.

### `/tests`: Test files for the application.

### `Makefile`: Commands to streamline common tasks such as running, building, and testing the application.
  - **`run`**: Command to run the application.
  - **`build`**: Command to build the application into an executable.
  - **`test-api`**: Command to run tests for the API.


## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

If you have any questions or feedback, feel free to reach out via GitHub issues.