package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	password "github.com/vzglad-smerti/password_hash"
)

//	func main() {
//		router := gin.Default()
//		// Получение всех профилей
//		router.GET("/profiles", GetProfiles)
//		// Получение всех постов
//		router.GET("/posts", GetPosts)
//		// Получение поста по ID
//		router.GET("/profiles/:id", GetProfileByID)
//		// Получение профиля по ID
//		router.GET("/posts/:id", GetPostByID)
//		// Создание нового профиля
//		router.POST("/profiles", CreateProfile)
//		// Создание новой поста
//		router.POST("/posts", CreatePost)
//		// Обновление существующего профиля
//		router.PUT("/profiles/:id", UpdateProfile)
//		// Обновление существующего поста
//		router.PUT("/posts/:id", UpdatePost)
//		// Удаление профиля
//		router.DELETE("/profiles/:id", DeleteProfile)
//		// Удаление поста
//		router.DELETE("/posts/:id", DeletePost)
//		router.Run(":8080")
//	}
type Profile struct {
	id           int       `json:"id"`
	nickname     string    `json:"nickname"`
	hashPassword string    `json:"hashPassword"`
	status       bool      `json:"status"`
	accessLevel  int       `json:"accessLevel"`
	name         string    `json:"name"`
	lastname     string    `json:"lastname"`
	createdAT    time.Time `json:"createdAT"`
}

func main() {
	connStr := "user=postgres password=1706 dbname=ProjectONE sslmode=disable host=localhost port=5432"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Done!")
	defer db.Close()

	//пароль который будем проверять
	password_users := "testing"
	//создадим хеш пароля
	hash, err := password.Hash(password_users)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(hash)

	result, err := db.Exec("insert into profiles (nickname, hash_password, status, access_level, firstname, lastname, created_at) values ( $1, $2, $3, $4, $5, $6, $7)",
		"Twinsiller", hash, true, 5, "Alex", "Boldinov", time.Now())
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // не поддерживается
	fmt.Println(result.RowsAffected()) // количество добавленных строк

	rows, err := db.Query("select * from profiles")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	profiles := []Profile{}

	for rows.Next() {
		p := Profile{}
		err := rows.Scan(&p.id, &p.nickname, &p.hashPassword, &p.status, &p.accessLevel, &p.name, &p.lastname, &p.createdAT)
		if err != nil {
			fmt.Println(err)
			continue
		}
		profiles = append(profiles, p)
	}
	for i, p := range profiles {
		fmt.Println(i, "dsadasdasnlda")
		fmt.Println(p.id, p.nickname, p.status, p.accessLevel, p.name, p.lastname)
	}
	fmt.Println("Exit")
}
