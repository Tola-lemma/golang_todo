# 📝 Go RESTful TODO API

A clean, modular **TODO management REST API** built with Golang, **Gorilla Mux**, **PostgreSQL**, and **Viper** for configuration.  
Designed to be simple for beginners yet scalable for real-world use.


## 🚀 Features

✅ RESTful CRUD endpoints (`Create`, `Read`, `Update`, `Delete`)  
✅ PostgreSQL as database (with simple migrations in Go)  
✅ Environment-based config using **Viper**  
✅ Input validation and error handling  
✅ Optional field updates (PATCH-like behavior)  
✅ Clean, layered folder structure  

---

## 📁 Project Structure


        .
        ├── Config/
        │   └── config.go   # Configuration
        ├── controllers/
        │   └── todoController.go   # All handler logic (CRUD)
        ├── database/
        │   ├── connection.go      # Connects to PostgreSQL
        │───migrations/
        │   └── migrate.go
        │───Schema/
        │   └── createTodoTable.go
        ├── models/
        │   └── todo.go            # Structs for Todo
        ├── validations/
        │   └── todoValidation.go  # Validation rules
        ├── routes/
        │   └── routes.go          # Gorilla Mux route definitions
        ├── main.go
        └── go.mod


---
  ## ⚙️ Installation & Setup
### 1️⃣ Clone the repository  
        git clone https://github.com/Tola-lemma/golang_todo
        cd golang_todo 


### 2️⃣ Install dependencies

            go mod tidy

### 3️⃣ Set up environment variables

Create a `.env` file in your project root:


            SERVER_PORT=5000
            SSL_MODE=false
            DB_DRIVER=postgres
            DB_HOST=localhost
            DB_PORT=5432
            DB_USER=postgres
            DB_PASS=yourpassword
            DB_NAME=todo_db


---

## 🧱 Database Migration

Migrations are handled automatically by Go code (no `.sql` needed).
Each migration file (like `createTodoTable.go`) runs when the app starts.

Example migration:

            
            const createTodoTable = `
            CREATE TABLE IF NOT EXISTS todos (
                id SERIAL PRIMARY KEY,
                title VARCHAR(255) NOT NULL,
                completed BOOLEAN DEFAULT FALSE,
                created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
            );
            `

            _, err := DB.Exec(createTodoTable)



## ▶️ Run the Server

            go run main.go


Server starts on:

            http://localhost:8080


---

## 🧩 API Endpoints

| Method   | Endpoint          | Description                      | Example Body                                  |
| -------- | ----------------- | -------------------------------- | --------------------------------------------- |
| `POST`   | `/api/todos`      | Create a new todo                | `{ "title": "Buy milk" }`                     |
| `GET`    | `/api/todos`      | Get all todos                    | —                                             |
| `GET`    | `/api/todos/{id}` | Get todo by ID                   | —                                             |
| `PUT`    | `/api/todos/{id}` | Update a todo (replace fields)   | `{ "title": "Buy bread", "completed": true }` |
| `DELETE` | `/api/todos/{id}` | Delete a todo                    | —                                             |

✅ Supports partial updates with optional fields.

---

## 🧠 Example Request (cURL)

### Create a todo


    curl -X POST http://localhost:8080/api/todos \
        -H "Content-Type: application/json" \
        -d '{"title":"Learn Go"}'


### Update completion status


    curl -X PATCH http://localhost:8080/api/todos/1 \
        -H "Content-Type: application/json" \
        -d '{"completed":true}'




## 🧩 Validation Rules

* Title must be **at least 3 characters**
* Completed is optional (`true/false`)
* Empty update requests (`{}`) return 400



## 🧰 Tech Stack

| Layer       | Tool                 |
| ----------- | -------------------- |
| Language    | Go 1.25.2             |
| Router      | Gorilla Mux          |
| Config      | Viper                |
| Database    | PostgreSQL           |
| Validation  | Custom Go validation |
| Environment | `.env` file          |

---

## 🧪 Sample Response


    {
    "message": "Todo updated successfully"
    }



## 🧱 Example Folder Workflow


        POST /api/todos → controllers.CreateTodo → database.DB.Exec(INSERT)
        GET /api/todos  → controllers.GetTodos  → database.DB.Query(SELECT)
        PATCH /api/todos/:id → controllers.UpdateTodo → dynamic SQL builder




## ❤️ Contributing

Pull requests are welcome!



## 📜 License

This project is licensed under the MIT License.
Feel free to use and modify it for your own learning or production use.

---

## 👨‍💻 Author

**Tola Lemma**
        🌐 [github.com/Tola-lemma](https://github.com/Tola-lemma)

> 🦋 *"Code should be simple enough for a beginner to read,
> yet powerful enough for a senior to extend."* – Go Philosophy

