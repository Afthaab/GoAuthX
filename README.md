# 🔐 GoAuthX - Lightweight Stateless Authentication Library for Golang

GoAuthX is a **lightweight, stateless authentication library** for Golang, designed to simplify secure authentication using **JWT (JSON Web Tokens)**. With no database dependencies, GoAuthX provides a flexible, high-performance authentication system for modern web applications and microservices.

---

## 🚀 Features
✅ **Stateless Authentication** - No session storage required, works purely with JWTs.  
✅ **Secure Password Hashing** - Uses bcrypt for strong password encryption.  
✅ **JWT Generation & Validation** - Easy-to-use functions for issuing and verifying tokens.  
✅ **Middleware for Authentication** - Protect your API routes effortlessly.  
✅ **Role-Based Access Control (RBAC)** - Enforce access restrictions based on user roles.  
✅ **Token Expiry Handling** - Built-in expiration support for better security.  
✅ **Highly Configurable** - Custom secret keys, expiration times, and hashing algorithms.  

---

## 📌 Why Choose GoAuthX?
🔹 **Lightweight & Fast** - No database needed, making it ideal for **microservices** and **serverless** applications.  
🔹 **Highly Secure** - Uses industry-standard cryptography for password hashing and JWT signing.  
🔹 **Simple & Developer-Friendly** - Minimal setup with a clean API design.  
🔹 **Production-Ready** - Designed to scale for real-world applications.

---

## 📦 Installation

Install GoAuthX using:
```sh
go get github.com/yourusername/GoAuthX
```

---

## 🛠️ Usage

### 1️⃣ Hashing a Password
```go
hashedPassword := auth.HashPassword("my_secure_password")
```

### 2️⃣ Verifying a Password
```go
isValid := auth.VerifyPassword(hashedPassword, "user_input_password")
```

### 3️⃣ Generating a JWT Token
```go
token, err := auth.GenerateToken("userID123", "admin")
```

### 4️⃣ Validating a JWT Token
```go
claims, err := auth.ValidateToken(token)
```

### 5️⃣ Protecting API Routes with Middleware
```go
router.Use(auth.AuthMiddleware)
```

---

## ⚡ API Reference

### **🔹 HashPassword(password string) string**
Hashes a password using bcrypt.

### **🔹 VerifyPassword(hashed, plain string) bool**
Checks if a password matches the stored hash.

### **🔹 GenerateToken(userID, role string) (string, error)**
Creates a JWT token with user ID and role claims.

### **🔹 ValidateToken(token string) (*Claims, error)**
Validates a JWT and extracts user claims.

### **🔹 AuthMiddleware(next http.Handler) http.Handler**
Middleware to protect API routes.

---

## 🔒 Security Best Practices
- **Use HTTPS** - Always encrypt communication between clients and your server.
- **Short-lived Tokens** - Set a reasonable expiration time for JWTs.
- **Secure Secret Key** - Store your JWT secret key securely.
- **Rotate Secrets Periodically** - Change your signing key periodically to enhance security.

---

## 📜 License
MIT License. See `LICENSE` for details.

---

## 📧 Contact & Support
For any issues or feature requests, please open an [issue](https://github.com/yourusername/GoAuthX/issues) on GitHub. Contributions are welcome! 🎉
