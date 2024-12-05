package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("my_secret_key")

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func generateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(jwtKey)
	fmt.Println("\n\n", ss)
	return token.SignedString(jwtKey)
}

func login(c *gin.Context) {
	var creds Credentials
	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	// Здесь добавим простую проверку пароля
	if creds.Username != "user" || creds.Password != "password" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	token, err := generateToken(creds.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		fmt.Println("\n\n", tokenString)
		claims := &Claims{}
		fmt.Println("\n\nclaims.Username ", claims.Username)
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

type Profile struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

var countIdProfiles int = 2
var profiles = []Profile{
	{ID: 0, Name: "Alex", Password: "322"},
	{ID: 1, Name: "Valerchik", Password: "52"},
	{ID: 2, Name: "Igar", Password: "894"},
}

func GetProfiles(c *gin.Context) {
	c.JSON(http.StatusOK, profiles)
}

func CreateProfile(c *gin.Context) {
	var newProfile Profile

	if err := c.BindJSON(&newProfile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	countIdProfiles++
	newProfile.ID = countIdProfiles
	profiles = append(profiles, newProfile)
	c.JSON(http.StatusCreated, newProfile)
}

func main() {
	router := gin.Default()

	router.POST("/login", login)

	protected := router.Group("/")
	protected.Use(authMiddleware())
	{
		protected.GET("profiles", GetProfiles)
		protected.POST("profiles", CreateProfile)
		// другие защищенные маршруты
	}

	router.Run(":8080")
}
