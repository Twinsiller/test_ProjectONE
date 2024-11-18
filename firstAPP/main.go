package main

import (
	"database/sql"
	"fmt"
	"gin-notes-api/database"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var connStr string

var DbGlobal **sql.DB

func main() {
	connStr := database.StringConnectToBase("database/config_db.json")
	var db, errDb = sql.Open("postgres", connStr)
	if errDb != nil {
		panic(errDb)
	}
	defer db.Close()
	DbGlobal = &db
	fmt.Println("=======db :\n", db)
	fmt.Println("=======DbGlobal :\n", DbGlobal, "\n\n\n\n")
	fmt.Println("=======*DbGlobal :\n", *DbGlobal, "\n\n\n\n")
	fmt.Println("Data base is working!")
	// db, err := sql.Open("postgres", connStr)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Done!")
	// defer db.Close()
	// //пароль который будем проверять
	// password_users := "testing"
	// //создадим хеш пароля
	// hash, err := password.Hash(password_users)
	// if err != nil {
	// 	log.Print(err)
	// }
	// fmt.Println(hash)
	// result, err := db.Exec("insert into profiles (nickname, hash_password, status, access_level, firstname, lastname, created_at) values ( $1, $2, $3, $4, $5, $6, $7)",
	// 	"Twinsiller", hash, true, 5, "Alex", "Boldinov", time.Now())
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(result.LastInsertId()) // не поддерживается
	// fmt.Println(result.RowsAffected()) // количество добавленных строк
	// rows, err := db.Query("select * from profiles")
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()
	// profiles := []Profile{}
	// for rows.Next() {
	// 	p := Profile{}
	// 	err := rows.Scan(&p.id, &p.nickname, &p.hashPassword, &p.status, &p.accessLevel, &p.firstname, &p.lastname, &p.createdAT)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// 	profiles = append(profiles, p)
	// }
	// for _, p := range profiles {
	// 	//fmt.Println(i, "dsadasdasnlda")
	// 	fmt.Println(p.id, p.nickname, p.status, p.accessLevel, p.firstname, p.lastname)
	// }

	fmt.Println("Запуск сервера: ")

	router := gin.Default()

	// Получение всех профилей
	router.GET("/profiles", GetProfiles)

	// Получение профиля по ID
	router.GET("/profiles/:id", GetProfileById)

	// Создание профиля
	router.POST("/profiles", CreateProfile)

	// Обновление существующего профиля
	router.PUT("/profiles/:id", UpdateProfile)

	router.Run(":8080")
}
