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
  - Retrieve a list of all blog posts.

- **Get a Post by ID**

  - `GET /posts/{id}`
  - Retrieves a specific blog post by its ID.
  - Parameter: id (integer) - The unique identifier of the blog post.
  - Response: Array of a single blog post objects.

- **Create a New Post**

  - `POST /posts`
  - Creates a blog new post.
  - Request Body: Blog post object (title, content, author). Example request body:
    ```json
    {
      "title": "Post Title",
      "content": "Post content",
      "author": "Author Name"
    }
    ```
  - Response: The created blog post object with its assigned ID.

- **Update a Post by ID**

  - `PUT /posts/{id}`
  - Update an existing blog post by its ID.
  - Parameter: id (integer) - The unique identifier of the blog post.
  - Request Body: Updated blog post object (title, content, author).Example request body:
    ```json
    {
      "title": "Updated Title",
      "content": "Updated content",
      "author": "Author Name"
    }
    ```
  - Response: The updated blog post object.

- **Delete a Post by ID**

  - `DELETE /posts/{id}`
  - Deletes a blog post by its ID.
  - Parameter: id (integer) - The unique identifier of the blog post.
  - Response: Success message or status.

- **Get Comments for a Post**

  - `GET /posts/{postId}/comments`
  - Retrieves all comments for a specified blog post.
  - Parameter: postId (integer) - The unique identifier of the blog post.
  - Response: Array of comment objects.

- **Create a Comment for a Post**
  - `POST /posts/{postId}/comments`
  - Adds a new comment to a specified blog post.
  - Parameter: postId (integer) - The unique identifier of the blog post.
  - Request Body: Comment object (content, author).Example request body:
    ```json
    {
      "content": "Comment content",
      "author": "Author Name"
    }
    ```
  - Response: The created comment object with its assigned ID.
