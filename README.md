# Learning Companion

![Go Version](https://img.shields.io/badge/Go-1.18+-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

A modern, AI-powered backend service designed to be your personal learning companion. This application leverages local Large Language Models (LLMs) through Ollama to provide a powerful and private learning experience.

This project is built with a clean and scalable architecture, following SOLID principles to ensure maintainability and testability.

## ✨ Features

- **Secure Authentication:** JWT-based authentication (Login/Logout) for secure user access.
- **User & Admin Roles:** Clear separation of roles for managing the platform.
- **AI-Powered Learning:** Integrates with local LLMs via Ollama and `langchaingo` for intelligent learning interactions.
- **RESTful API:** A well-structured, versioned API for both frontend and backend services.
- **Scalable Architecture:** Built with a modular design using interfaces and dependency injection for easy extension and testing.

## 🛠️ Technologies Used

- **Backend:** Go
- **Framework:** Gin Gonic
- **Database:** MySQL
- **AI Integration:** [Ollama](https://ollama.com/) for running local LLMs/SLMs.
- **Go LangChain:** [langchaingo](https://github.com/tmc/langchaingo) for interacting with LLMs.
- **Authentication:** JWT (JSON Web Tokens)
- **Database ORM:** GORM

## 🚀 Getting Started

Follow these instructions to get the Learning Companion server up and running on your local machine.

### Prerequisites

- **Go:** Version 1.18 or higher.
- **MySQL:** A running instance of MySQL.
- **Ollama:** You must have [Ollama installed](https://ollama.com/download) and running.

### 1. Set up Ollama and Pull a Model

First, ensure your Ollama server is running. Then, pull a model that you want to use. We recommend starting with `llama3` or a smaller model like `phi3`.

```bash
ollama pull llama3
```

Verify that the model is available:

```bash
ollama list
```

### 2. Clone the Repository

```bash
git clone <your-repository-url>
cd learning-companion
```

### 3. Configure the Application

Copy the example configuration file:

```bash
cp configs/config.example.yaml configs/config.yaml
```

Now, edit `configs/config.yaml` with your specific settings:

- **Database:** Update the `host`, `port`, `user`, `password`, and `dbname` for your MySQL instance.
- **JWT:** Change the `secret` to a long, random string for security.
- **Ollama:** Ensure the `base_url` points to your Ollama instance and the `model` is set to the one you downloaded.

### 4. Install Dependencies

```bash
go mod tidy
```

### 5. Run Database Migrations

(Assuming you have a migration tool set up. If not, you will need to manually create the database and tables based on the models in `internal/model`.)

```bash
# Example command - replace with your actual migration command
go run migrations/migrate.go
```

### 6. Run the Server

```bash
go run cmd/server/main.go
```

The server should now be running on the port specified in your `config.yaml` (default is `8080`).

## 🧪 API Usage & Testing

You can test the API endpoints using any API client like [Postman](https://www.postman.com/), [Insomnia](https://insomnia.rest/), or `curl`.

**Base URL:** `http://localhost:8080/api/v1`

### Authentication

#### 1. User Login

- **Endpoint:** `POST /auth/login`
- **Headers:** `Content-Type: application/json`
- **Body:**

```json
{
  "username": "your_username",
  "password": "your_password"
}
```

**Success Response (200 OK):**

```json
{
    "status": "success",
    "message": "Login successful",
    "data": {
        "access_token": "ey...",
        "refresh_token": "ey...",
        "expires_at": "2025-07-23T18:00:00Z"
    }
}
```

Copy the `access_token` for use in authenticated requests.

#### 2. Authenticated Requests

For endpoints that require authentication, you must include the `Authorization` header.

- **Header:** `Authorization: Bearer <your_access_token>`

#### 3. User Logout

- **Endpoint:** `POST /auth/logout`
- **Headers:** `Authorization: Bearer <your_access_token>`

**Success Response (200 OK):**

```json
{
    "status": "success",
    "message": "Logout successful",
    "data": null
}
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.

## 📄 License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
