# GoAI Starter

![Go Version](https://img.shields.io/badge/Go-1.18+-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

A lightweight, open-source starter kit for building secure and intelligent AI applications using Go, Gin, and local LLMs via Ollama.  
Perfect for learning, experimenting, or kickstarting your own AI-powered backend.

---

## ✨ Features

- **JWT Authentication:** Secure login and logout with access/refresh tokens.
- **Local AI Integration:** Connects with local SLMs (like `llama3`, `phi3`) via [Ollama](https://ollama.com/).
- **LangChain for Go:** Uses [Langchaingo](https://github.com/tmc/langchaingo) to prompt and chain LLM logic.
- **RESTful API:** Versioned and cleanly structured endpoints.
- **Scalable Design:** Modular, testable, and SOLID-principled codebase using interfaces and DI.

---

## 🛠️ Tech Stack

- **Language:** Go
- **Web Framework:** [Gin](https://github.com/gin-gonic/gin)
- **ORM:** [GORM](https://gorm.io/)
- **Database:** MySQL
- **AI Runtime:** [Ollama](https://ollama.com/)
- **LLM Wrapper:** [Langchaingo](https://github.com/tmc/langchaingo)
- **Auth:** JWT

---

## 🚀 Getting Started

### 1. Requirements

- Go `1.18+`
- MySQL
- [Ollama](https://ollama.com/download)

### 2. Pull a Local Model

Start Ollama and pull a model like `llama3`:

```bash
ollama pull llama3