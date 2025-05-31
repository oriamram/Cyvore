# Cyvore Setup Guide

This guide provides steps to set up and run the Cyvore project.
Here is a link to a walkthrough video:
https://drive.google.com/file/d/14lXxw1jiYoyuh2C1mUhO5WfhbqHdpyIA/view?usp=sharing

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

![image](https://github.com/user-attachments/assets/27eebc2c-f160-494c-91cb-214f3737c4e6)
![image](https://github.com/user-attachments/assets/d3b2b9fd-42a7-410a-ace9-9f45cf11e62d)
![image](https://github.com/user-attachments/assets/329e9083-ab4a-4c64-a776-da37cc55e953)

