package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	password "github.com/vzglad-smerti/password_hash"
)

type Profile struct {
	Id           int       `json:"id"`
	Nickname     string    `json:"nickname"`
	HashPassword string    `json:"hashpassword"`
	Status       bool      `json:"status"`
	AccessLevel  int       `json:"accesslevel"`
	Firstname    string    `json:"firstname"`
	Lastname     string    `json:"lastname"`
	CreatedAt    time.Time `json:"createdat"`
}

var profiles = []Profile{
	//{id: 87, nickname: "Check", hashPassword: "dsawqcxs", status: true, accessLevel: 23, firstname: "Tom", lastname: "qwerty", t: time.Now()},
}
var countidProfiles int = 2

func GetProfiles(c *gin.Context) {

	rows, err := Db.Query("select * from profiles")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		p := Profile{}
		err := rows.Scan(&p.Id, &p.Nickname, &p.HashPassword, &p.Status, &p.AccessLevel, &p.Firstname, &p.Lastname, &p.CreatedAt)
		if err != nil {
			fmt.Println(err)
			continue
		}
		profiles = append(profiles, p)
	}
	// fmt.Println("Выгруженные данные из базы по профилям")
	// for _, p := range profiles {
	// 	fmt.Println(p.id, p.nickname, p.status, p.accessLevel, p.firstname, p.lastname)
	// }
	c.JSON(http.StatusOK, profiles)
}

func GetProfileById(c *gin.Context) {
	id := c.Param("id")

	row := Db.QueryRow("select * from profiles WHERE id = $1;", id)

	p := Profile{}
	err := row.Scan(&p.Id, &p.Nickname, &p.HashPassword, &p.Status, &p.AccessLevel, &p.Firstname, &p.Lastname, &p.CreatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "profile not found"})
		panic(err)
	}
	c.JSON(http.StatusOK, p)
}

func CreateProfile(c *gin.Context) {
	p := Profile{}

	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	if hash, err := password.Hash(p.HashPassword); err != nil {
		log.Print(err)
	} else {
		p.HashPassword = hash
	}
	result, err := Db.Exec("insert into profiles (nickname, hash_password, status, access_level, firstname, lastname, created_at) values ( $1, $2, $3, $4, $5, $6, $7)",
		p.Nickname, p.HashPassword, p.Status, p.AccessLevel, p.Firstname, p.Lastname, time.Now(),
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // не поддерживается
	fmt.Println(result.RowsAffected()) // количество добавленных строк
	profiles = append(profiles, p)
	c.JSON(http.StatusCreated, p)
}

// func UpdateProfile(c *gin.Context) {
// 	idStr := c.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
// 		return
// 	}
// 	var updatedProfile Profile

// 	if err := c.BindJSON(&updatedProfile); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
// 		return
// 	}

// 	for i, profile := range profiles {
// 		if profile.id == id {
// 			profiles[i] = updatedProfile
// 			c.JSON(http.StatusOK, updatedProfile)
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
// }

// func DeleteProfile(c *gin.Context) {
// 	idStr := c.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
// 		return
// 	}
// 	for i, profile := range profiles {
// 		if profile.id == id {
// 			profiles = append(profiles[:i], profiles[i+1:]...)
// 			c.JSON(http.StatusOK, gin.H{"message": "profile was deleted"})
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusNotFound, gin.H{"message": "profile not found for delete"})
// }
