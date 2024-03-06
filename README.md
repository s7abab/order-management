# Order Management Application

This is a Go application for managing orders via a REST API and Next.js for the client-side.

## Getting Started

### Prerequisites

- Go installed on your machine. [Download Go](https://golang.org/dl/)
- Node.js installed on your machine. [Download Node.js](https://nodejs.org/en/)
- MySQL database server installed and running

### Installation

1. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/s7abab/order-management.git
    ```

2. Navigate to the project directory:

    ```bash
    cd order-management
    ```

3. Navigate to the server directory:

    ```bash
    cd server
    ```

4. Install server dependencies:

    ```bash
    go mod tidy
    ```

5. Navigate to the client directory:

    ```bash
    cd ../client
    ```

6. Install client dependencies:

    ```bash
    npm install
    ```

### Configuration

1. Configure MySQL Connection:
    - Navigate to the `server/db` directory.
    - Update the `connection.go` file with your MySQL connection details.

### Running the Application

#### Running the Server

To run the server, execute the following command from the `server` directory:

```bash
go run main.go
```
#### Running the Client

To run the client, execute the following command from the `client` directory:

```bash
npm run dev
