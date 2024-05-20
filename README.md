# todoproject01

 Todo Project
This is a simple TO-DO list application built using Go and the Gin framework. It allows users to create, update, delete, and view their TO-DO lists, as well as add items to these lists. The application includes JWT-based authentication and role-based authorization.

Features
User authentication with JWT.
Create, read, update, and delete TO-DO lists.
Add items to TO-DO lists.
Soft delete functionality for TO-DO lists and items.
Role-based access control for users and admins.

Installation
1: Clone the repository:
git clone https://github.com/yourusername/todoproject01.git
cd todoproject01
2: Install dependencies:
go mod tidy
3: Run the application:
go run main.go

API Endpoints
Authentication
POST /login
Request Body:
{
  "username": "user1",
  "password": "password"
}
Response:
{
  "token": "jwt_token"
}
TO-DO List
POST /todos

Request Body:
{
  "id": "unique_todo_list_id"
}
Response:
{
  "id": "unique_todo_list_id",
  "creationDate": "2024-05-20T00:00:00Z",
  "modificationDate": "2024-05-20T00:00:00Z",
  "completionPercent": 0,
  "messages": {}
}
GET /todos/:id
Response:
{
  "id": "unique_todo_list_id",
  "creationDate": "2024-05-20T00:00:00Z",
  "modificationDate": "2024-05-20T00:00:00Z",
  "completionPercent": 0,
  "messages": {}
}
PUT /todos/:id

Request Body:
{
  "completionPercent": 50,
  "messages": {}
}
Response:
{
  "id": "unique_todo_list_id",
  "creationDate": "2024-05-20T00:00:00Z",
  "modificationDate": "2024-05-20T00:00:00Z",
  "completionPercent": 50,
  "messages": {}
}
DELETE /todos/:id

Response:
{
  "message": "Todo list deleted"
}
TO-DO Items
POST /todos/:id/items
Request Body:
{
  "id": "unique_item_id",
  "content": "Item content",
  "completionStatus": false
}
Response:
{
  "id": "unique_todo_list_id",
  "creationDate": "2024-05-20T00:00:00Z",
  "modificationDate": "2024-05-20T00:00:00Z",
  "completionPercent": 0,
  "messages": {
    "unique_item_id": {
      "id": "unique_item_id",
      "todoListId": "unique_todo_list_id",
      "creationDate": "2024-05-20T00:00:00Z",
      "modificationDate": "2024-05-20T00:00:00Z",
      "content": "Item content",
      "completionStatus": false
    }
  }
}

User Roles
User: Can only manage their own TO-DO lists.
Admin: Can view and manage all users' TO-DO lists.
Notes
Soft delete is implemented by setting a deletion date instead of removing the record from the database.
Make sure to replace the placeholder data (like jwt_key) with actual secure values for a production environment.
