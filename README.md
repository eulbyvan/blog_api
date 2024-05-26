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
        createdb your_database_name
        ```

    - Run the following SQL script to set up the database schema and insert example data:
        ```sql
        -- Create the posts table
        CREATE TABLE posts (
            id SERIAL PRIMARY KEY,
            title VARCHAR(255) NOT NULL,
            content TEXT NOT NULL,
            status VARCHAR(50) CHECK (status IN ('draft', 'publish')) NOT NULL,
            publish_date TIMESTAMP
        );

        -- Create the tags table
        CREATE TABLE tags (
            id SERIAL PRIMARY KEY,
            label VARCHAR(100) UNIQUE NOT NULL
        );

        -- Create the junction table for the many-to-many relationship
        CREATE TABLE post_tags (
            post_id INT REFERENCES posts(id) ON DELETE CASCADE,
            tag_id INT REFERENCES tags(id) ON DELETE CASCADE,
            PRIMARY KEY (post_id, tag_id)
        );

        -- Insert some example data
        INSERT INTO posts (title, content, status, publish_date)
        VALUES
        ('First Post', 'Content of the first post', 'draft', NULL),
        ('Second Post', 'Content of the second post', 'publish', '2023-05-25 10:00:00');

        INSERT INTO tags (label)
        VALUES
        ('Go'),
        ('API');

        INSERT INTO post_tags (post_id, tag_id)
        VALUES
        (1, 1),
        (1, 2),
        (2, 1);
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
1. **Get a paginated list of posts filtered by a specific tag**

    ```http
    GET /api/posts?tag={TAG_FILTER_STRING}&page={PAGE_NUMBER_INT}&size={PAGE_SIZE_INT}
    ```

    **Response:**
    `200 OK`: Returns a JSON array of posts

2. **Create a new post**

    ```http
    POST /api/posts
    ```
    **Request Body:**
    ```json
    {
        "title": "REST API with Go",
        "content": "Lorem ipsum",
        "tags": ["Go", "Lang"]
    }
    ```

    **Response:**
    `201` Created: Returns the created post object
3. **Update a post**
    ```http
    PUT /api/posts/{id}
    ```
    **Request Body:**
    ```json
    {
        "title": "REST API with Go",
        "content": "Lorem ipsum dolor sit amet",
        "tags": ["Go", "Lang"]
    }
    ```
    **Response:**
    `200 OK`: Returns the updated post object
4. **Delete an existing post**
    ```http
    DELETE /api/posts/{id}
    ```
    **Response:**
    `204 No Content`: Indicates that the post was successfully deleted

## Additional Resources

- [Go Documentation](https://golang.org/doc/)
- [PostgreSQL Documentation](https://www.postgresql.org/docs/)
- [dotenv](https://github.com/joho/godotenv)