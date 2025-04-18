# golang-echo-mongodb


## ⚙️ Requirements

- Go 1.18+
- MongoDB (local or cloud)
- Git

---

## 🚀 Getting Started


### 1. Clone the repo

```bash
git clone https://github.com/rzmsv/golang-echo-mongodb.git
```
### 2. Run application
1. Rename `.env.example` to `.env` and fill in the necessary configurations.
2. For install dependencies run this command :
    ```bash
    go mod tidy
    ```
3. For run app run this command :
    ```bash
    make start
    ```

### 3. APIs
## 🔥 Example API Endpoints

| Method | Endpoint                | Description          |
|--------|-------------------------|----------------------|
| GET    | `/api/v1/products`      | List all products    |
| POST   | `/api/v1/products/new-products`      | Create a product     |
| GET    | `/api/v1/products/:id`  | Get product by ID    |
| PATCH    | `/api/v1/products/:id`  | Update product by ID |
| DELETE | `/api/v1/products/:id`  | Delete product       |