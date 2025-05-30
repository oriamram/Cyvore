# Cyvore Security Scanner

A full-stack web application for security reconnaissance using Amass, built with Go and Vue 3.

## Overview

This application integrates the Amass tool for subdomain enumeration and OSINT gathering, providing a modern web interface for security professionals to visualize and analyze the collected data.

### Features

-   **Authentication**: Secure JWT-based authentication system
-   **Real-time Updates**: WebSocket integration for live data updates
-   **Interactive UI**: Modern, responsive interface built with Vue 3 and Tailwind CSS
-   **Data Visualization**: Organized display of subdomains and their relationships
-   **Search & Filter**: Advanced filtering capabilities for large datasets

## Tech Stack

### Backend

-   Go with Gin framework
-   JWT for authentication
-   WebSocket for real-time updates
-   SQLite for data storage
-   Amass integration for security scanning

### Frontend

-   Vue 3 with TypeScript
-   Vite for build tooling
-   Tailwind CSS for styling
-   ShadCN for UI components
-   WebSocket client for real-time updates

## Setup Instructions

### Prerequisites

-   Go 1.21 or higher
-   Node.js 18 or higher
-   Amass installed and configured
-   SQLite

### Backend Setup

1. Navigate to the backend directory:

    ```bash
    cd backend
    ```

2. Install dependencies:

    ```bash
    go mod download
    ```

3. Configure environment variables:

    ```bash
    cp .env.example .env
    # Edit .env with your configuration
    ```

4. Run the server:
    ```bash
    go run cmd/server/main.go
    ```

### Frontend Setup

1. Navigate to the frontend directory:

    ```bash
    cd frontend
    ```

2. Install dependencies:

    ```bash
    npm install
    ```

3. Run the development server:
    ```bash
    npm run dev
    ```

## Design Choices

### Architecture

-   **Backend**: Clean architecture with separation of concerns

    -   Handlers for HTTP endpoints
    -   Services for business logic
    -   Models for data structures
    -   Middleware for cross-cutting concerns

-   **Frontend**: Component-based architecture
    -   Reusable UI components
    -   Shared layouts
    -   Type-safe API integration
    -   Real-time data updates

### Security

-   JWT-based authentication
-   HTTP-only cookies for refresh tokens
-   CORS configuration
-   Input validation
-   Error handling

### Performance

-   WebSocket for real-time updates
-   Efficient database queries
-   Pagination for large datasets
-   Optimized frontend builds

## Amass Integration

The application integrates Amass for subdomain enumeration and OSINT gathering. The integration:

-   Executes Amass scans
-   Processes JSON output
-   Stores results in SQLite
-   Provides real-time updates via WebSocket
-   Exposes data through REST API

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
