# Blog API

Welcome to the Blog API documentation.

## Table of Contents

- [Project Description](#project-description)
- [Prerequisites](#prerequisites)
- [Running the Project Locally](#running-the-project-locally)
- [API Documentation](#api-documentation)
- [Additional Resources](#additional-resources)

## Project Description

This API allows you to manage posts and tags for a blog. Below you'll find details on how to set up, use, and interact with the API.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- [Docker](https://docs.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Running the Project Locally

Follow these steps to run the project on your local development environment:

1. **Pull the Image**
    ```bash
    docker pull eulbyvan/blog_api
    ```

2. **Run the application:**
    ```bash
    docker compose up
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