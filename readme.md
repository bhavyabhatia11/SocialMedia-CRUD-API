# Social Media CRUD API

This project is a CRUD (Create, Read, Update, Delete) API for a social media application. It is built using the Go programming language and the Echo web framework, and uses a MySQL database to store data.

## Setup

To use this project, you will need to have Go and MySQL installed on your system. Once you have these installed, follow the steps below:

1. Clone this repository to your local machine.
2. Navigate to the project directory and run `go mod download` to download the required dependencies.
3. Create a new MySQL database and Update the following details in db.go file inside the database folder: 
    ```
    const DB_USERNAME = <your-database-username>
    const DB_PASSWORD = <your-database-password>
    const DB_NAME = <your-database-name>
    const DB_HOST = <your-database-host>
    const DB_PORT = <your-database-port>
    ```

4. Start the server by running `go run server.go`. The Echo server should be up on port `1336`.


## Architechture

### File Structure
```
📦goCRUD-MySQL
 ┣ 📂 database
 ┃ ┗ 📜 db.go
 ┣ 📂 handlers
 ┃ ┣ 📜 comments.go
 ┃ ┣ 📜 handler.go
 ┃ ┣ 📜 posts.go
 ┃ ┗ 📜 users.go
 ┣ 📂 models
 ┃ ┗ 📜 models.go
 ┣ 📂 modules
 ┃ ┣ 📜 comment.go
 ┃ ┣ 📜 post.go
 ┃ ┗ 📜 user.go
 ┣ 📜 go.mod
 ┣ 📜 go.sum
 ┣ 📜 readme.md
 ┗ 📜 server.go
```

### Code Flow 


The `server.go` file is the entrypoint for the project. This is the main package of a web server built with the Go programming language using the Echo framework.
Here are the main components of this code:

- The database package: `database.InitDb()` initializes a database connection and returns a pointer to the `gorm.DB` object. gorm.`AutoMigrate()` method is called on the DB object to create database tables for the three models.

- The handler package: `h := &handler.Handler{DB: db}` creates a new handler object, which takes the DB object created in the database package as input. The handler methods implement CRUD functionality for the three models.

- The Echo router: `e := echo.New()` creates a new instance of the Echo router. `e.GET()`, `e.POST()`, `e.PUT()`, and `e.DELETE()` methods define HTTP endpoints that handle GET, POST, PUT, and DELETE requests respectively. Each endpoint is mapped to a handler method.

- HTTP endpoints: There are several HTTP endpoints defined in this code. For example, `e.GET("/users", h.GetAllUsers)` maps the GET request for the `"/users"` URL to the `GetAllUsers()` method of the handler object. Similarly, `e.POST("/posts/create-post", h.CreatePost)` maps the POST request for the `"/posts/create-post"` URL to the `CreatePost()` method of the handler object.


This project is divided into 4 main folders.
1. `Database` : This folder contains the db.go file which is responsible for initializing and connecting to a MySQL database using the GORM library in Go. It defines the necessary connection details such as the database username, password, host, port, and database name. The `InitDb()` function calls the connectDB() function to establish a connection to the database and return the `*gorm.DB` object, which is used to interact with the database. The `connectDB()` function creates a database source name (DSN) string and uses it to open a connection to the database, and returns the `*gorm.DB` object.
2. `Modules` : This folder contains 3 files (For each model). The code defines various CRUD (Create, Read, Update, Delete) functions for interacting with the These models in the database. These functions take a database connection *gorm.DB object as their first argument and a pointer to a model object or a slice of model objects as their second argument, depending on the function.
3. `Models` : This Folder contains a models.go file. This code defines the model structs for the application's database schema. Here's an ER Diagram for the same:

    ```
        +-------------+
        |    User     |
        +-------------+
        | UID (PK)    |
        | Name        |
        | Email       |
        +-------------+

            | 1
            |
            | N
        +-------------+
        |   Comment   |
        +-------------+
        | UID (PK)    |
        | UserUID (FK)|--------+
        | PostUID (FK)|        |
        | Content     |        |
        +-------------+        |
                                |
            | 1                 | N
            |                   |
        +-------------+        |
        |    Post     |        |
        +-------------+        |
        | UID (PK)    |        |
        | UserUID (FK)|<-------+
        | Title       |
        | Content     |
        | Likes       |
        | Dislikes    |
        +-------------+
    ```
4. `Handlers` : This folder contains files separated model wise to handle API requests coming from the server.go file. The handler package contains the HTTP handlers that are responsible for handling incoming HTTP requests and generating appropriate responses.



## Endpoints

The API exposes the following endpoints:

1. `POST /users`: Create a new user. The request body should contain a JSON object with the following fields: `name`, `email`, and `password`.
2. `GET /users/:id`: Get a user by ID.
3. `PUT /users/:id`: Update a user by ID. The request body should contain a JSON object with the fields you want to update.
4. `DELETE /users/:id`: Delete a user by ID.
5. `POST /posts`: Create a new post. The request body should contain a JSON object with the following fields: `title`, `body`, and `user_id`.
6. `GET /posts/:id`: Get a post by ID.
7. `PUT /posts/:id`: Update a post by ID. The request body should contain a JSON object with the fields you want to update.
8. `DELETE /posts/:id`: Delete a post by ID.

## Contributing

Contributions to this project are welcome! If you have any ideas or find any issues, please open an issue or pull request on the GitHub repository.

## License

This project is licensed under the MIT License. See the LICENSE file for more information.
