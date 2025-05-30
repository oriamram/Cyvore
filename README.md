# Cyvore Setup Guide

This guide provides steps to set up and run the Cyvore project.

## Getting Started

1.  Clone the repository:

    ```bash
    git clone https://github.com/oriamram/Cyvore.git
    ```

2.  Pull the Amass Docker image:
    ```bash
    docker pull caffix/amass
    ```

## Setting up Dependencies

Due to not using a monorepo manager, you will need to set up dependencies for the backend and frontend separately.

1.  Navigate to the `backend` directory and download dependencies:

    ```bash
    cd backend
    go mod download
    go mod tidy
    ```

2.  Navigate to the `frontend` directory and install dependencies:
    ```bash
    cd ../frontend
    npm ci
    ```

## Running the Services

Each service needs to be run in its respective folder.

1.  **Frontend:** Run the development server from the `frontend` directory:

    ```bash
    cd frontend # If not already there
    npm run dev
    ```

2.  **Server:** Run the backend server from the `backend` directory:

    ```bash
    cd backend # If not already there
    go run cmd/server/main.go
    ```

3.  **Websocket:** Run the websocket server from the `backend` directory:
    ```bash
    cd backend # If not already there
    go run cmd/server/main.go
    ```

Once all services are running, you can access the login page at `localhost:5173`.
