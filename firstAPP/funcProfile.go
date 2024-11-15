package main

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// )

// type Profile struct {
// 	id           int    `json:"id"`
// 	nickname     int    `json:"nickname"`
// 	hashPassword string `json:"hashPassword"`
// 	status       string `json:"status"`
// 	accessLevel  string `json:"accessLevel"`
// 	name         string `json:"name"`
// 	surname      int    `json:"surname"`
// }

// var countidProfiles int = 2
// var profiles = []Profile{
// 	{id: 0, name: "Alex", hashPassword: "322"},
// 	{id: 1, name: "Valerchik", hashPassword: "52"},
// 	{id: 2, name: "Igar", hashPassword: "894"},
// }

// func GetProfiles(c *gin.Context) {
// 	c.JSON(http.StatusOK, profiles)
// }

// func GetProfileByid(c *gin.Context) {
// 	idStr := c.Param("id")
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
// 		return
// 	}

// 	for _, profile := range profiles {
// 		if profile.id == id {
// 			c.JSON(http.StatusOK, profile)
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusNotFound, gin.H{"message": "profile not found"})
// }

// func CreateProfile(c *gin.Context) {
// 	var newProfile Profile

// 	if err := c.BindJSON(&newProfile); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
// 		return
// 	}
// 	countidProfiles++
// 	newProfile.id = countidProfiles
// 	profiles = append(profiles, newProfile)
// 	c.JSON(http.StatusCreated, newProfile)
// }

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
