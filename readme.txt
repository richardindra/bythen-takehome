How to run & test
- Create new database name "blog" and create table users, jobs, applications
- Run go run cmd/http/main.go on Terminal
- Wait until HTTP server running on port :8888

POSTMAN Collenction
https://api.postman.com/collections/13503324-a422f4a8-219c-4548-8347-4724bbef533a?access_key=PMAT-01HTWEK81B03KT5RA209Z5W0W9

# Take-Home Test for Backend Engineer

**Notice:** You are not required to complete 100% of the task. Please do your best within the given time frame, and focus on demonstrating your skills and approach to problem-solving. We are interested in seeing your thought process and how you tackle the core aspects of the task.

## Task: Building a Simple Blog Platform

Create a RESTful API using Golang that allows users to perform CRUD operations on blog posts and comments, with user registration and login functionality. The data should be stored in a MySQL database.

### Entities

**User**
- id (integer, primary key)
- name (string)
- email (string, unique)
- password_hash (string)
- created_at (timestamp)
- updated_at (timestamp)

**Blog Post**
- id (integer, primary key)
- title (string)
- content (text)
- author_id (integer, foreign key referencing User)
- created_at (timestamp)
- updated_at (timestamp)

**Comment**
- id (integer, primary key)
- post_id (integer, foreign key referencing Blog Post)
- author_name (string)
- content (text)
- created_at (timestamp)

### API Endpoints

**User Registration & Authentication**
- `POST /register` - Register a new user.
- `POST /login` - Login and receive a token for authentication.

**Blog Posts**
- `POST /posts` - Create a new blog post.
- `GET /posts/{id}` - Get blog post details by ID.
- `GET /posts` - List all blog posts.
- `PUT /posts/{id}` - Update a blog post.
- `DELETE /posts/{id}` - Delete a blog post.

**Comments**
- `POST /posts/{id}/comments` - Add a comment to a blog post.
- `GET /posts/{id}/comments` - List all comments for a blog post.

### Database Designs

Provide a MySQL schema design that reflects the above entities and their relationships.
Ensure proper indexing for performance optimization.

## Evaluation Criteria

- Code quality and organization.
- Completeness of the required features.
- Security measures (e.g., authentication implementation).
- Creativity and problem-solving approach, especially if modifications to the entities were made.

## Setup Instructions

### Option 1: Using Docker

If you have Docker installed, you can start the app with the following commands:

```
docker-compose build
docker-compose up
```

The server will be up and running at http://localhost:8080.

### Option 2: Manual Setup

If you prefer to set up the web server manually, ensure you have the following prerequisites:

- Go version 1.21.0
- MySQL version 8.0

Once the prerequisites are ready:

1. Install [Air](https://github.com/air-verse/air), a live reload tool for Go.
2. Navigate to the `./app` directory.
3. Start the server by running `air`.

## Submission Instructions

Push your code to a Git repository and send us the link.



CREATE TABLE `m_users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL,
  `name` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password_hash` TEXT NOT NULL,
  `status` char(1) DEFAULT 'Y',
  `last_login_at` timestamp NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `m_blog_posts` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `title` varchar(200) NOT NULL,
  `content` text NOT NULL,
  `author_id` bigint NOT NULL,
  `status` char(1) DEFAULT 'Y',
  `view_count` int NOT NULL DEFAULT '0',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_blog_user` (`author_id`),
  CONSTRAINT `fk_blog_user` FOREIGN KEY (`author_id`) REFERENCES `m_users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

CREATE TABLE `m_comments` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `blog_id` bigint NOT NULL,
  `author_id` bigint NOT NULL,
  `content` text NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `fk_comment_user` (`author_id`),
  KEY `fk_comment_post` (`blog_id`),
  CONSTRAINT `fk_comment_user` FOREIGN KEY (`author_id`) REFERENCES `m_users` (`id`) ON DELETE CASCADE,
  CONSTRAINT `fk_comment_post` FOREIGN KEY (`blog_id`) REFERENCES `m_blog_posts` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci

////
CREATE TABLE m_users (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(50) NOT NULL,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL,
  password_hash TEXT NOT NULL,
  status TINYINT NOT NULL DEFAULT 1,
  last_login_at TIMESTAMP NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  UNIQUE KEY uk_users_username (username),
  UNIQUE KEY uk_users_email (email)
) ENGINE=InnoDB;

CREATE TABLE m_blog_posts (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(200) NOT NULL,
  content TEXT NOT NULL,
  author_id BIGINT NOT NULL,
  status TINYINT NOT NULL DEFAULT 1,
  view_count INT NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  
  KEY idx_blog_author (author_id),
  CONSTRAINT fk_blog_user
    FOREIGN KEY (author_id) REFERENCES m_users(id)
    ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE m_comments (
  id BIGINT AUTO_INCREMENT PRIMARY KEY,
  blog_id BIGINT NOT NULL,
  author_id BIGINT NOT NULL,
  content TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  
  KEY idx_comment_blog (blog_id),
  KEY idx_comment_author (author_id),
  CONSTRAINT fk_comment_user
    FOREIGN KEY (author_id) REFERENCES m_users(id)
    ON DELETE CASCADE,
  CONSTRAINT fk_comment_post
    FOREIGN KEY (blog_id) REFERENCES m_blog_posts(id)
    ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;