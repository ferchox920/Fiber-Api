# Fiber-Api: User, Product, and Order Management in Go

This project is a simple Go application that allows you to manage users, products, and orders. It utilizes the Fiber framework for handling routes and GORM for interacting with the SQLite database.

## Prerequisites

- Go (must be installed on your system)

## Installation

1. Clone this repository:

   ```sh
   git clone https://github.com/ferchox920/fiber-api.git
   cd fiber-api
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Run the application:

   ```sh
   go run main.go
   ```

## Usage

The application provides routes to create users, products, and orders:

- **Create User**: Send a POST request to `/api/users` with user details in the JSON body.

- **Create Product**: Send a POST request to `/api/products` with product details in the JSON body.

- **Create Order**: Send a POST request to `/api/orders` with order details in the JSON body, including user and product references.

## SQLite Database

This project uses SQLite as the database engine. The database file `api.db` will be created in the root directory of the project when you run the application.

## Database Configuration

The database connection is configured in the `database.ConnectDb()` function. You can modify this function in the `database/database.go` file to suit your requirements.

## Contribution

If you wish to contribute to this project, you are welcome to do so. You can fork the repository, make your changes, and submit a pull request.


Please ensure that you adjust any references, paths, or additional details in the README to match your project's structure and setup accurately.