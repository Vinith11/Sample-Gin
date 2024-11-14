# 🛍️ Product API with Go and Gin

This is a sample RESTful API built using [Go](https://golang.org/) and the [Gin framework](https://github.com/gin-gonic/gin). The API performs CRUD operations on a list of products, stored in a JSON file.

## 🚀 Deployment

The project is deployed on Render.

## 📂 Project Structure

- `main.go`: Entry point of the application.
- `products.json`: JSON file that stores the product data.

## 🛠️ Features

- **Health Check Endpoint**: Check API status with `GET /`
- **List Products**: Retrieve all products with `GET /products`
- **Get Product by ID**: Retrieve a single product by ID with `GET /products/:id`
- **Add New Product**: Add a product with `POST /products`
- **Update Product**: Update an existing product with `PUT /products/:id`
- **Delete Product**: Delete a product by ID with `DELETE /products/:id`

## ⚙️ Installation and Running Locally

1. Clone the repository
   ```bash
   git clone https://github.com/yourusername/your-repo.git
   cd your-repo
   ```

2. Install dependencies
   ```bash
   go mod init
   go mod tidy
   ```

3. Run the API
   ```bash
   go run main.go
   ```

4. The API will be available at `http://localhost:8080`.

## 📝 API Endpoints

### Health Check
- **`GET /`**
  - **Description**: Returns a health-check message.
  - **Response**:
    ```json
    "Health-check"
    ```

### List Products
- **`GET /products`**
  - **Description**: Returns a list of all products.
  - **Response**:
    ```json
    [
      {
        "id": 1,
        "name": "Sample Product",
        "price": 10.99
      }
    ]
    ```

### Get Product by ID
- **`GET /products/:id`**
  - **Description**: Returns a product by its ID.
  - **Response**:
    ```json
    {
      "id": 1,
      "name": "Sample Product",
      "price": 10.99
    }
    ```

### Add New Product
- **`POST /products`**
  - **Description**: Adds a new product to the list.
  - **Request**:
    ```json
    {
      "name": "New Product",
      "price": 12.99
    }
    ```
  - **Response**:
    ```json
    {
      "id": 2,
      "name": "New Product",
      "price": 12.99
    }
    ```

### Update Product
- **`PUT /products/:id`**
  - **Description**: Updates an existing product.
  - **Request**:
    ```json
    {
      "name": "Updated Product",
      "price": 15.99
    }
    ```
  - **Response**:
    ```json
    {
      "id": 1,
      "name": "Updated Product",
      "price": 15.99
    }
    ```

### Delete Product
- **`DELETE /products/:id`**
  - **Description**: Deletes a product by ID.
  - **Response**:
    ```json
    {
      "message": "Product deleted"
    }
    ```

## 📦 Deployment

The application can be deployed on Render or any platform that supports Go applications. Configure environment variables and dependencies as needed.

## 📂 JSON File Data Persistence

Product data is stored in `products.json`. The file is read on startup and updated whenever products are added, modified, or deleted.



Happy coding! 😊