# Bythen - Takehome

POSTMAN Collection :
https://postman.co/workspace/My-Workspace~4c24b0d9-bd0b-48c0-866c-9be55b0b931d/collection/13503324-b1f082e1-ae51-4f45-a916-a89d31b9e1ad?action=share&creator=13503324

GitHub Link : 
https://github.com/richardindra/bythen-takehome

## Task: Building a Simple Blog Platform

Create a RESTful API using Golang that allows users to perform CRUD operations on blog posts and comments, with user registration and login functionality.

### Entities

**User**
- id (integer, primary key)
- username (string, unique)
- name (string)
- email (string, unique)
- password_hash (string)
- status (string)
- last_login_at (timestamp)
- created_at (timestamp)
- updated_at (timestamp)

**Blog Post**
- id (integer, primary key)
- title (string)
- content (text)
- author_id (integer, foreign key referencing User)
- status (string)
- view_count (integer)
- created_at (timestamp)
- updated_at (timestamp)

**Comment**
- id (integer, primary key)
- blog_id (integer, foreign key referencing Blog Post)
- author_id (integer, foreign key referencing User)
- content (text)
- created_at (timestamp)
- updated_at (timestamp)

## API Endpoints

**User Registration & Authentication**
- `POST /auth/v1/register` - Register a new user.
- `POST /auth/v1/login` - Login and receive a token for authentication.

**Blog Posts**
- `POST /blog/v1/posts` - Create a new blog post.
- `GET /blog/v1/posts/{id}` - Get blog post details by ID.
- `GET /blog/v1/posts?sort={sort}&page={page}&limit={limit}` - List all blog posts.
- `GET /blog/v1/posts?search=author&author={authorID}&sort={sort}&page={page}&limit={limit}` - List all blog posts by author.
- `PUT /blog/v1/posts/{id}` - Update a blog post.
- `DELETE /blog/v1/posts/{id}` - Delete a blog post.

**Comments**
- `POST /blog/v1/posts/{id}/comments` - Add a comment to a blog post.
- `GET /blog/v1/posts/{id}/comments?sort={sort}&page={page}&limit={limit}` - List all comments for a blog post.

## Design Decisions

- **Layered Architecture**  
  The application is structured into distinct layers (data, service, delivery, and entity) to separate concerns.  

- **Authentication & Authorization**  
  JWT is used for authentication due to its simplicity and stateless nature.  
  Authorization checks are implemented at the service layer to ensure that only resource owners can modify or delete their data.

- **Database Design**  
  Primary keys use `BIGINT` to allow future scalability.  
  Foreign key constraints and indexes are added to ensure data integrity and improve query performance.

- **Error Handling**  
  Sentinel errors are used for cases (e.g. data not found, token expired), allowing error mapping to HTTP status codes at the delivery layer.

## Notes

- **Database Initialization**  
  An `init.sql` file is provided to initialize and create all required MySQL tables.

- **Bearer Token**  
  Blog posts and comments APIs require a Bearer token.  
  The token is automatically set in the Authorization header using the Postman collection script when the Login API is called.

- **Access Control**  
  Update and Delete operations for blog posts are restricted to the post author.  
  If a user attempts to update or delete a post created by another user, the API will return **HTTP 403 (Forbidden)**.

- **Pagination & Sorting**  
  APIs that return lists support basic pagination using `page` and `limit` query parameters.  
  Sorting is implemented using a `sort` parameter with `asc` or `desc` values, currently applied to the `id` field.

- **Token Expiration**  
  Access tokens expire **12 hours** after issuance.  
  Users must re-login once the token has expired.

- **Timezone**  
  All timestamps returned by the API are in **UTC**, following the Docker container configuration.

- **Docker Setup**  
  While I do not work with Docker on a daily basis, I made a best effort to ensure the application can be built and run correctly using Docker and Docker Compose.

- **Project Structure**  
  The repository structure follows the same pattern I use in my day-to-day work, focusing on clear separation of concerns and maintainability.

## Future Improvements

- Implement **refresh tokens** to improve authentication flow and reduce frequent re-login.
- Add **like / reaction counters** for blog posts and comments.
- Improve **error handling consistency**, including more granular HTTP status codes and standardized error responses.
- Sorting is currently limited to the `id` field for simplicity, but can be extended to support additional fields such as `created_at` or `title` in the future.

## Setup Instructions

If you have Docker installed, you can start the app with the following commands:

```
docker-compose build
docker-compose up
```

The server will be up and running at http://localhost:8080.
