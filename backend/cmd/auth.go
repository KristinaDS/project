package main

import (
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
	"golang.org/x/crypto/bcrypt"
)

var db *pgx.Conn

var jwtSecret = []byte("your_secret_key")

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
	CreatedAt    string `json:"created_at"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

func initDB() {
	var err error
	connString := "postgres://postgres:postgres@localhost:5432/quiz"
	db, err = pgx.Connect(context.Background(), connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	fmt.Println("Database connection established")
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func generateJWT(userID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   fmt.Sprintf("%d", userID),
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func isValidEmail(email string) bool {
	var re = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func registerUser(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	if !isValidEmail(input.Email) {
		c.JSON(400, gin.H{"error": "Invalid email format"})
		return
	}

	var existingUser User
	err := db.QueryRow(context.Background(), "SELECT id, email FROM users WHERE email = $1", input.Email).Scan(&existingUser.ID, &existingUser.Email)
	if err == nil {
		c.JSON(400, gin.H{"error": "Email already exists"})
		return
	}

	hashedPassword, err := hashPassword(input.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	_, err = db.Exec(context.Background(), "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)",
		input.Username, input.Email, hashedPassword)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(201, gin.H{"message": "User created successfully"})
}

func authenticateUser(c *gin.Context) {
	var input AuthInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	var user User
	err := db.QueryRow(context.Background(), "SELECT id, username, email, password_hash FROM users WHERE email = $1", input.Email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid email or password"})
		return
	}

	if !checkPasswordHash(input.Password, user.PasswordHash) {
		c.JSON(400, gin.H{"error": "Invalid email or password"})
		return
	}

	token, err := generateJWT(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, AuthResponse{Token: token})
}

func main() {
	initDB()
	defer db.Close(context.Background())

	r := gin.Default()

	r.Use(cors.Default())

	r.POST("/register", registerUser)
	r.POST("/authenticate", authenticateUser)

	r.Run(":1234")
}
