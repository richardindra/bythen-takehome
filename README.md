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

### API Endpoints

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

## Notes
- Bearer Token: Automatically set using the Postman collection script when Login API is called.
- Access Control: Update/Delete APIs are restricted to the author of the post. Any other user will receive a 403 error.
- Token Expiration: 12 hours from issuance; you need to re-login after expiration.
- Timezone: All timestamps in responses are UTC, as per Docker configuration.

## Setup Instructions

If you have Docker installed, you can start the app with the following commands:

```
docker-compose build
docker-compose up
```

The server will be up and running at http://localhost:8080.

## Submission Instructions

Push your code to a Git repository and send us the link.
