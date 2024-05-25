# Blog API Documentation

Welcome to the Blog API documentation. This API allows you to manage posts and tags for a blog. Below you'll find details on how to set up, use, and interact with the API.

## Table of Contents
- [Getting Started](#getting-started)
- [API Endpoints](#api-endpoints)
  - [Posts](#posts)
    - [Create Post](#create-post)
    - [Update Post](#update-post)
    - [Delete Post](#delete-post)
    - [Get Post by ID](#get-post-by-id)
    - [Get Paged Posts](#get-paged-posts)
  - [Tags](#tags)
    - [Create Tag](#create-tag)
    - [Update Tag](#update-tag)
    - [Delete Tag](#delete-tag)
    - [Get Tag by ID](#get-tag-by-id)
    - [Get All Tags](#get-all-tags)
- [Database Schema](#database-schema)
- [Running the Application](#running-the-application)

## Getting Started

### Prerequisites
- Go 1.16 or higher
- PostgreSQL

### Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/your_username/blog_api.git
    cd blog_api
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Set up your PostgreSQL database and update the connection string in `database.go` file:
    ```go
    const (
        host     = "localhost"
        port     = 5432
        user     = "your_db_user"
        password = "your_db_password"
        dbname   = "your_db_name"
    )
    ```

4. Run the database migrations to set up the schema.

5. Start the server:
    ```bash
    go run main.go
    ```

## API Endpoints

### Posts

#### Create Post
- **URL:** `/api/v1/posts`
- **Method:** `POST`
- **Request Body:**
    ```json
    {
      "title": "REST API with Go",
      "content": "Lorem ipsum",
      "status": "draft",
      "publish_date": "2024-05-25T10:00:00Z",
      "tags": [
        {
          "label": "Go"
        },
        {
          "label": "API"
        }
      ]
    }
    ```
- **Response:**
    - **Success (201 Created):**
        ```json
        {
          "status": "success",
          "message": "Post created successfully",
          "data": {
            "id": 1,
            "title": "REST API with Go",
            "content": "Lorem ipsum",
            "status": "draft",
            "publish_date": "2024-05-25T10:00:00Z",
            "tags": [
              {
                "id": 1,
                "label": "Go"
              },
              {
                "id": 2,
                "label": "API"
              }
            ]
          }
        }
        ```
    - **Error (500 Internal Server Error):**
        ```json
        {
          "status": "error",
          "message": "Error message",
          "data": null
        }
        ```

#### Update Post
- **URL:** `/api/v1/posts/{id}`
- **Method:** `PUT`
- **Request Body:**
    ```json
    {
      "title": "Updated Title",
      "content": "Updated content",
      "status": "publish",
      "publish_date": "2024-06-01T10:00:00Z",
      "tags": [
        {
          "label": "Go"
        },
        {
          "label": "API"
        }
      ]
    }
    ```
- **Response:**
    - **Success (200 OK):**
        ```json
        {
          "status": "success",
          "message": "Post updated successfully",
          "data": {
            "id": 1,
            "title": "Updated Title",
            "content": "Updated content",
            "status": "publish",
            "publish_date": "2024-06-01T10:00:00Z",
            "tags": [
              {
                "id": 1,
                "label": "Go"
              },
              {
                "id": 2,
                "label": "API"
              }
            ]
          }
        }
        ```
    - **Error (500 Internal Server Error):**
        ```json
        {
          "status": "error",
          "message": "Error message",
          "data": null
        }
        ```

#### Delete Post
- **URL:** `/api/v1/posts/{id}`
- **Method:** `DELETE`
- **Response:**
    - **Success (200 OK):**
        ```json
        {
          "status": "success",
          "message": "Post deleted successfully",
          "data": null
        }
        ```
    - **Error (500 Internal Server Error):**
        ```json
        {
          "status": "error",
          "message": "Error message",
          "data": null
        }
        ```

#### Get Post by ID
- **URL:** `/api/v1/posts/{id}`
- **Method:** `GET`
- **Response:**
    - **Success (200 OK):**
        ```json
        {
          "status": "success",
          "message": "Post retrieved successfully",
          "data": {
            "id": 1,
            "title": "REST API with Go",
            "content": "Lorem ipsum",
            "status": "draft",
            "publish_date": "2024-05-25T10:00:00Z",
            "tags": [
              {
                "id": 1,
                "label": "Go"
              },
              {
                "id": 2,
                "label": "API"
              }
            ]
          }
        }
        ```
    - **Error (500 Internal Server Error):**
        ```json
        {
          "status": "error",
          "message": "Error message",
          "data": null
        }
        ```

#### Get Paged Posts
- **URL:** `/api/v1/posts`
- **Method:** `GET`
- **Query Parameters:**
    - `page` (optional, default: 1): The page number.
    - `size` (optional, default: 10): The number of posts per page.
- **Response:**
    - **Success (200 OK):**
        ```json
        {
          "status": "success",
          "message": "Posts retrieved successfully",
          "data": [
            {
              "id": 1,
              "title": "REST API with Go",
              "content": "Lorem ipsum",
              "status": "draft",
              "publish_date": "2024-05-25T10:00:00Z",
              "tags": [
                {
                  "id": 1,
                  "label": "Go"
                },
                {
                  "id": 2,
                  "label": "API"
                }
              ]
            },
            ...
          ]
        }
        ```
    - **Error (500 Internal Server Error):**
        ```json
        {
          "status": "error",
          "message": "Error message",
          "data": null
        }
        ```

### Tags

#### Create Tag
- **URL:** `/api/v1/tags`
- **Method:** `POST`
- **Request Body:**
    ```json
    {
      "label": "Tech"
    }
    ```
- **Response:**
    - **Success (201 Created):**
        ```json
        {
          "status": "success",
          "message": "Tag created successfully",
          "data": {
            "id": 1,
            "label": "Tech"
          }
        }
        ```
    - **Error (500 Internal Server Error):**
        ```json
        {
          "status": "error",
          "message": "Error message",
          "data": null
        }
        ```

#### Update Tag
- **URL:** `/api/v1/tags/{id}`
- **Method:** `PUT`
- **Request Body:**
    ```json
    {
      "label": "Updated Label"
    }
    ```
- **Response:**
    - **Success (200 OK):**
        ```json
        {
          "status": "success",
          "message": "Tag updated successfully",
          "data": {
            "id": 1,
            "label": "Updated Label"
          }
        }
        ```
    - **Error (500 Internal Server Error):**
        ```json
        {
          "status": "error",
          "message": "Error message",
          "data": null
        }
        ```

#### Delete Tag
- **URL:** `/api/v1/tags/{id}`
- **Method:** `DELETE`
- **Response:**
    - **Success (200 OK):**
        ```json
        {
          "status": "success",
          "message": "Tag deleted successfully",
          "data": null
        }
        ```
    - **Error (500 Internal Server Error):**
        ```json
        {
          "status": "error",
          "message": "Error message",
          "data": null
        }
        ```

#### Get Tag by ID
- **URL:** `/api/v1/tags/{id}`
- **Method:** `GET`
- **Response:**
    - **Success (200 OK):**
        ```json
        {
          "status": "success",
          "message": "Tag retrieved successfully",
          "data": {
            "id": 1,
            "label": "Tech"
          }
        }
        ```
    - **Error (500 Internal Server Error):**
        ```json
        {
          "status": "error",
          "message": "Error message",
          "data": null
        }
        ```

#### Get All Tags
- **URL:** `/api/v1/tags`
- **Method:** `GET`
- **Response:**
    - **Success (200 OK):**
        ```json
        {
          "status": "success",
          "message": "Tags retrieved successfully",
          "data": [
            {
              "id": 1,
              "label": "Tech"
            },
            ...
          ]
        }
        ```
    - **Error (500 Internal Server Error):**
        ```json
        {
          "status": "error",
          "message": "Error message",
          "data": null
        }
        ```

## Database Schema

```sql
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    status VARCHAR(50) NOT NULL,
    publish_date TIMESTAMP
);

CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    label VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE post_tags (
    post_id INT NOT NULL,
    tag_id INT NOT NULL,
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE ON UPDATE CASCADE,
    PRIMARY KEY (post_id, tag_id)
);