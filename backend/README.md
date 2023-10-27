# Backend Setup

This project is a backend service that manages user search history using a PostgreSQL database. Follow the steps below to set up and run the backend.

## Prerequisites

- PostgreSQL installed on your local machine.
- Go (Golang) installed on your local machine.

## Database Setup

1. Install PostgreSQL on your system.

2. Run the schema.sql file located inside the `database` folder using the SQL shell (psql) to create the necessary tables and schema for the project.

3. Set DB_HOST, DB_PORT, DB_USER, DB_PASSWORD in .env file according to your postgre sql local environment

## Backend Setup

1. Clone the repository to your local machine.

2. Navigate to the backend folder.

3. Set your own OPEN_WEATHER_API_KEY and JWT_SECRET_KEY if you want to modify

3. Run the following command to install all necessary libraries:

```
go mod tidy

```

4. Run the main.go file to start the backend server:

```
go run main.go

```

The backend will start running on the default port (usually 8080).

## Testing Golang APIs

- Refer to postmancollection folder and import the json file using postman to use Golang APIs