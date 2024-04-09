# Seam-backend-assignment

It's the repository for the take home assignment of the backend intern role at Seam. I create a simple API for a blog application that allows users to manage blog posts and comments.

## Installation

To install Project Name, follow these steps:

1. **Clone the repository**

   ```bash
   git clone https://github.com/SylviaZhang16/Seam-backend.git
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

## Running the Application

To run Project Name locally, follow these steps:

1. **Build the project** (optional, if you prefer to compile before running):

   ```bash
   go build
   ```

2. **Run the project:**
   ```bash
   go run main.go
   ```

Access the application by navigating to `http://localhost:8000` in your web browser or using a tool like Postman to interact with the API.

## API Endpoints

Below are the endpoints available in this API:

- **Get All Posts**

  - `GET /posts`
  - Retrieves a list of all posts.

- **Get a Post by ID**

  - `GET /posts/{id}`
  - Retrieves a specific post by its ID.

- **Create a New Post**

  - `POST /posts`
  - Creates a new post. Example request body:
    ```json
    {
      "title": "Post Title",
      "content": "Post content",
      "author": "Author Name"
    }
    ```

- **Update a Post by ID**

  - `PUT /posts/{id}`
  - Updates the specified post. Example request body:
    ```json
    {
      "title": "Updated Title",
      "content": "Updated content",
      "author": "Author Name"
    }
    ```

- **Delete a Post by ID**

  - `DELETE /posts/{id}`
  - Deletes the specified post.

- **Get Comments for a Post**

  - `GET /posts/{postId}/comments`
  - Retrieves all comments for the specified post.

- **Create a Comment for a Post**
  - `POST /posts/{postId}/comments`
  - Adds a new comment to the specified post. Example request body:
    ```json
    {
      "content": "Comment content",
      "author": "Author Name"
    }
    ```
