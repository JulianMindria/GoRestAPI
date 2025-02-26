# GoRestAPI - Voucher & Transaction Management API

GoRestAPI is a RESTful API built using Golang with **DDD (Domain-Driven Design)** and **Onion Architecture**. It provides endpoints for managing users, brands, products, vouchers, and transactions. The API uses **PostgreSQL** as the database and **golang-migrate** for database migrations.

## Features
✅ User management with point system
✅ Product & Brand management
✅ Voucher system (discount & cashback) with expiration
✅ Transaction tracking (includes used vouchers & total price calculation)
✅ Clean architecture (DDD, Onion Architecture)
✅ Swagger integration for API documentation

## Tech Stack
- **Golang** (Gin Framework)
- **PostgreSQL** (Database)
- **GORM** (ORM)
- **golang-migrate** (Database migrations)
- **Swaggo** (Swagger documentation)
- **Makefile** (Task automation)

---

## Installation
### 1️⃣ Clone the Repository
```sh
git clone https://github.com/julianmindria/GoRestAPI.git
cd GoRestAPI
```

### 2️⃣ Set Up Environment Variables
Create a `.env` file in the root directory:
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=gorestdb
```

### 3️⃣ Install Dependencies
```sh
go mod tidy
```

### 4️⃣ Run Database Migrations
Ensure PostgreSQL is running, then execute:
```sh
make migrate-up
```
To rollback migrations:
```sh
make migrate-down
```

### 5️⃣ Run the Application
```sh
go run main.go
```

---

## API Endpoints
### 🌟 Base URL: `/api`


## Endpoints

### 🌟 **Voucher Endpoints**
- **Create Voucher**  
  `POST /api/vouchers`  
  Create a new voucher.

- **Get All Vouchers**  
  `GET /api/vouchers`  
  Retrieve all vouchers.

### 👤 **User Endpoints**
- **Create User**  
  `POST /api/users`  
  Create a new user.

- **Update User Points**  
  `PUT /api/users/:id/update`  
  Update the points of a user by their ID.

- **Get All Users**  
  `GET /api/users`  
  Retrieve all users.

- **Get Available Vouchers for User**  
  `GET /api/users/:id/vouchers`  
  Retrieve available vouchers for a specific user by their ID.

### 💳 **Transaction Endpoints**
- **Create Transaction**  
  `POST /api/transactions`  
  Create a new transaction.

- **Get Transaction by ID**  
  `GET /api/transactions/:id`  
  Retrieve transaction details by transaction ID.

- **Get All Transactions**  
  `GET /api/transactions`  
  Retrieve all transactions.

- **Delete Transaction**  
  `DELETE /api/transactions/:id`  
  Delete a transaction by its ID.

- **Get Transaction Details by User**  
  `GET /api/transactions/user/:user_id`  
  Retrieve transaction details for a specific user.

### 📦 **Product Endpoints**
- **Create Product**  
  `POST /api/products`  
  Create a new product.

- **Get All Products**  
  `GET /api/products`  
  Retrieve all products.

- **Get Product by ID**  
  `GET /api/products/:id`  
  Retrieve a single product by its ID.

- **Delete Product**  
  `DELETE /api/products/:id`  
  Delete a product by its ID.

### 🏷 **Brand Endpoints**
- **Create Brand**  
  `POST /api/brands`  
  Create a new brand.

- **Get All Brands**  
  `GET /api/brands`  
  Retrieve all brands.

- **Get Brand by ID**  
  `GET /api/brands/:id`  
  Retrieve a single brand by its ID.

- **Delete Brand**  
  `DELETE /api/brands/:id`  
  Delete a brand by its ID.

---

## API Response Structure
```json
{
  "status": 200,
  "message": "Success",
  "data": {...}
}
```
For list responses:
```json
{
  "status": 200,
  "message": "Success",
  "data": [{...}, {...}]
}
```

---

## Swagger Documentation
Swagger is available at:
```
http://localhost:8080/swagger/index.html
```
Run `swag init -g cmd/main.go` if you need to regenerate Swagger docs.

---

## Makefile Commands
To simplify commands, use:
```sh
make migrate-up   # Run migrations
make migrate-down # Rollback migrations
make run          # Run the API
```

---

## Contributing
Feel free to fork this repository and submit pull requests! 🚀

---

## License
This project is licensed under the **MIT License**.

