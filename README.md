# Golang Authentication API

A simple backend authentication service built using Golang, Gin, GORM, SQLite, JWT, and bcrypt.

## Features

- User Signup API
- User Login API
- Password Hashing using bcrypt
- JWT Authentication
- Role-Based Access Control
- Protected Routes
- Admin and User Roles
- SQLite Database
- Middleware for JWT Validation

---

# Tech Stack

- Golang
- Gin Framework
- GORM
- SQLite
- JWT
- bcrypt

---

# Project Structure

golang-auth/
│
├── main.go
├── go.mod
├── database/
│   └── db.go
├── handlers/
│   ├── auth.go
│   └── user.go
├── middleware/
│   └── auth.go
├── models/
│   └── user.go
├── utils/
│   └── jwt.go
└── README.md

---

# Installation & Setup

## Clone Repository

```bash
git clone https://github.com/Sashriya/Golong-authentication-api.git
