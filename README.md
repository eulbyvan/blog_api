# Blog API

Welcome to the Blog API documentation.

## Table of Contents

- [Project Description](#project-description)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Running the Project Locally](#running-the-project-locally)
- [API Documentation](#api-documentation)
- [Additional Resources](#additional-resources)

## Project Description

This API allows you to manage posts and tags for a blog. Below you'll find details on how to set up, use, and interact with the API.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- [Go](https://golang.org/doc/install) (version 1.22.3 or later)
- [PostgreSQL](https://www.postgresql.org/download/) (version 15 or later)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Installation

Follow these steps to install the project on your local machine:

1. **Clone the repository:**
    ```bash
    git clone https://github.com/eulbyvan/blog_api.git
    cd blog_api
    ```

2. **Install Go dependencies:**
    ```bash
    go mod tidy
    ```

3. **Set up PostgreSQL database:**

    - Create a new PostgreSQL database:
        ```bash
        createdb 
        ```

    - Run the database migrations (if any):
        ```bash
        # Use a migration tool like migrate or goose
        migrate -path ./migrations -database "postgres://username:password@localhost/your_database_name?sslmode=disable" up
        ```

4. **Set environment variables:**

    Create a `.env` file in the root of your project and add the necessary environment variables:

    ```dotenv
    DB_HOST=localhost
    DB_PORT=5432
    DB_USER=your_db_user
    DB_PASSWORD=your_db_password
    DB_NAME=your_database_name
    ```

## Running the Project Locally

Follow these steps to run the project on your local development environment:

1. **Start the PostgreSQL server:**
    ```bash
    pg_ctl -D /usr/local/var/postgres start
    ```

2. **Run the application:**
    ```bash
    go run cmd/api/main.go
    ```

3. **Access the application:**

    Open your web browser and navigate to `http://localhost:8080`.

## API Documentation

### Endpoints

#### GET /api/posts?tag={TAG_FILTER_STRING}&page={PAGE_NUMBER_INT}&size={PAGE_SIZE_INT}

- **Description:** Retrieve a paginated list of posts filtered by a specific tag.
- **Response:**
    - `200 OK`: Returns a JSON array of posts

#### POST /api/posts

- **Description:** Create a new example
- **Request:**
    - Body: `{"title":"REST API with Go","content":"Lorem ipsum","tags":["Go","Lang"]}`
- **Response:**
    - `201 Created`: Returns the created post object

#### PUT /api/posts/{id}

- **Description:** Update an existing post
- **Request:**
    - Body: `{"title":"REST API with Go","content":"Lorem ipsum","tags":["Go","Ling"]}`
- **Response:**
    - `200 OK`: Returns the updated post object

#### DELETE /api/posts/{id}

- **Description:** Delete an existing post
- **Response:**
    - `204 No Content`: Indicates that the post was successfully deleted

## Additional Resources

- [Go Documentation](https://golang.org/doc/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [Go Migrate](https://github.com/golang-migrate/migrate)
- [dotenv](https://github.com/joho/godotenv)