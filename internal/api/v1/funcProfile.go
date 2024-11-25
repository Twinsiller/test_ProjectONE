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

var profiles = []Profile{}

func GetProfiles(c *gin.Context) {
	rows, err := (*DbGlobal).Query("select * from profiles")
	if err != nil {
		log.Panic(err)
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

	c.JSON(http.StatusOK, profiles)
	profiles = []Profile{}
}

func GetProfileById(c *gin.Context) {
	var Db = *DbGlobal
	id := c.Param("id")

	row := Db.QueryRow("select * from profiles WHERE id = $1;", id)

	p := Profile{}
	err := row.Scan(&p.Id, &p.Nickname, &p.HashPassword, &p.Status, &p.AccessLevel, &p.Firstname, &p.Lastname, &p.CreatedAt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "profile not found"})
		return
		//panic(err)
	}
	c.JSON(http.StatusOK, p)
}

func CreateProfile(c *gin.Context) {
	var Db = *DbGlobal
	p := Profile{}

	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	if hash, err := password.Hash(p.HashPassword); err != nil {
		log.Print(err)
		return
	} else {
		p.HashPassword = hash
	}

	result, err := Db.Exec("insert into profiles (nickname, hash_password, status, access_level, firstname, lastname) values ( $1, $2, $3, $4, $5, $6)",
		p.Nickname, p.HashPassword, p.Status, p.AccessLevel, p.Firstname, p.Lastname,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // не поддерживается
	fmt.Println(result.RowsAffected()) // количество добавленных строк
	c.JSON(http.StatusCreated, p)
}

func UpdateProfile(c *gin.Context) {
	var Db = *DbGlobal
	id := c.Param("id")
	var p Profile

	if err := c.BindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	if hash, err := password.Hash(p.HashPassword); err != nil {
		log.Print(err)
		return
	} else {
		p.HashPassword = hash
	}

	result, err := Db.Exec("UPDATE profiles SET nickname = $1, hash_password = $2, status = $3, access_level = $4, firstname = $5, lastname = $6  WHERE id = $7",
		p.Nickname, p.HashPassword, p.Status, p.AccessLevel, p.Firstname, p.Lastname, id,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.LastInsertId()) // не поддерживается
	fmt.Println(result.RowsAffected()) // количество добавленных строк
	c.JSON(http.StatusAccepted, p)
}

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
