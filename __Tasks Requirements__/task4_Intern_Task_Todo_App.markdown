# Intern Task: Build a Todo App with REST APIs in Go

## Overview
You are tasked with designing and implementing a simple RESTful API for a Todo App using Go and the `gin-gonic/gin` framework for routing. The app will manage todo items stored in memory (using a slice) and support basic CRUD operations (Create, Read, Update, Delete). This project is designed for a beginner-level intern to learn REST API development, HTTP request handling, and modular design in Go. The focus is on creating a clean, functional API that can be tested via a command-line tool like cURL or a GUI tool like Postman.

## Requirements

### 1. Todo Structure
- Create a `Todo` struct with attributes:
  - `id` (int): A unique identifier for each todo.
  - `title` (string): A description of the task.
  - `completed` (bool): Indicates whether the task is completed.
- Use a slice to store todos in memory, ensuring no data persistence is required.

### 2. REST API Endpoints
- Implement the following endpoints using the `gin-gonic/gin` framework:
  - **GET `/todos`**: Retrieve all todos as a JSON array.
  - **GET `/todos/:id`**: Retrieve a specific todo by its ID (return 404 if not found).
  - **POST `/todos`**: Create a new todo with a title and completion status.
  - **PUT `/todos/:id`**: Update a todo’s title or completion status by its ID (return 404 if not found).
  - **DELETE `/todos/:id`**: Delete a todo by its ID (return 404 if not found).
- Responses must be in JSON format with appropriate HTTP status codes.

### 3. Input Handling
- For POST and PUT requests, accept a JSON body with `title` (string) and `completed` (bool).
- Validate input to ensure the `title` is non-empty; return a 400 status code for invalid input.

### 4. In-Memory Storage
- Store todos in a slice within the Go program.
- Assign incremental IDs to new todos (e.g., 1, 2, 3, ...).
- Data should persist only during the server’s runtime.

### 5. Error Handling
- Return appropriate HTTP status codes:
  - 200: Successful GET, POST, PUT, or DELETE.
  - 400: Invalid JSON or empty title.
  - 404: Todo not found for GET, PUT, or DELETE.
- Include a JSON error message (e.g., `{"error": "Todo not found"}`).

### 6. RESTful Interface
- Provide a REST API served at `http://localhost:8080`.
- Ensure the API can be tested with tools like cURL or Postman, allowing the user to:
  - Create a todo with a title and status.
  - Retrieve all todos or a specific todo.
  - Update a todo’s details.
  - Delete a todo.
  - View error messages for invalid operations.

## Example Usage
To illustrate how the system should function, the following example walks through creating todos, retrieving them, updating, and deleting, using a REST client like cURL or Postman.

### Step 1: Starting the Server
- The intern runs the Go program, starting the API server on `http://localhost:8080`.
- The in-memory store (slice) is initialized as empty.

### Step 2: Creating a Todo
- The user sends a POST request to `http://localhost:8080/todos` with the JSON body:
  ```
  {"title": "Finish Go project", "completed": false}
  ```
- The system assigns an ID (e.g., 1) and stores the todo in the slice.
- The response is:
  ```
  {
    "id": 1,
    "title": "Finish Go project",
    "completed": false
  }
  ```

### Step 3: Retrieving Todos
- The user sends a GET request to `http://localhost:8080/todos`.
- The system returns a JSON array of all todos:
  ```
  [
    {
      "id": 1,
      "title": "Finish Go project",
      "completed": false
    }
  ]
  ```
- The user sends a GET request to `http://localhost:8080/todos/1`.
- The response is the specific todo:
  ```
  {
    "id": 1,
    "title": "Finish Go project",
    "completed": false
  }
  ```

### Step 4: Updating a Todo
- The user sends a PUT request to `http://localhost:8080/todos/1` with:
  ```
  {"title": "Complete Go project", "completed": true}
  ```
- The system updates the todo with ID 1.
- The response is:
  ```
  {
    "id": 1,
    "title": "Complete Go project",
    "completed": true
  }
  ```

### Step 5: Deleting a Todo
- The user sends a DELETE request to `http://localhost:8080/todos/1`.
- The system removes the todo with ID 1 from the slice.
- The response is:
  ```
  {"message": "Todo deleted"}
  ```

### Step 6: Handling Errors
- The user sends a GET request to `http://localhost:8080/todos/999` (non-existent ID).
- The system returns:
  ```
  {"error": "Todo not found"}
  ```
  with a 404 status code.
- The user sends a POST request with an empty title:
  ```
  {"title": "", "completed": false}
  ```
- The system returns:
  ```
  {"error": "Title cannot be empty"}
  ```
  with a 400 status code.

This example demonstrates how the API handles todo creation, retrieval, updates, deletion, and error cases, ensuring the requirements are met in a clear and RESTful way.

## Guidelines
- **Modularity**: Structure the code to separate concerns (e.g., routing, data management, error handling).
- **Encapsulation**: Protect the todo slice within the program, using functions or methods to access or modify it.
- **Gin Framework**: Use `gin-gonic/gin` for routing and JSON handling, leveraging its features for simplicity.
- **REST Principles**: Follow REST conventions (e.g., use appropriate HTTP methods, status codes, and JSON responses).
- **Simplicity**: Keep the implementation minimal, focusing on the required functionality without unnecessary


### Resources:
https://gin-gonic.com/en/docs/