package controllers

import (
	"assignment-3/database"
	"assignment-3/models"
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateStatus(c *gin.Context) {
	db, err := database.InitDB()
	if err != nil {
		fmt.Println("Error initializing database", err)
		return
	}

	var status models.Status
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
			fmt.Println("Error:", err.Error())
			return
	}

	if err := db.Model(&models.Status{}).Where("id = ?", 1).Updates(&status).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update data in database"})
			fmt.Println("Error Updating data in database:", err)
			return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data has been updated"})
}

func UpdateStatusPeriodically() {
	for {
		water := rand.Intn(100) + 1
		wind := rand.Intn(100) + 1

		status := models.Status{
			Water: water,
			Wind:  wind,
		}

		apiURL := "http://localhost:8080/status/update" 
		payload, _ := json.Marshal(status)
		req, err := http.NewRequest("PUT", apiURL, bytes.NewBuffer(payload))
		if err != nil {
			fmt.Println("Error creating request:", err)
			time.Sleep(15 * time.Second)
			continue
		}
	
		req.Header.Set("Content-Type", "application/json")
	
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error making request:", err)
			time.Sleep(15 * time.Second)
			continue
		}
		defer resp.Body.Close()
	
		if resp.StatusCode != http.StatusOK {
			fmt.Println("API returned non-OK status:", resp.Status)
			time.Sleep(15 * time.Second)
			continue
		}

		waterStatus := getStatus(water, 6, 8)
		windStatus := getStatus(wind, 7, 15)

		logStatus := models.Status{
			Water: water,
			Wind:  wind,
		}

		logJSON, _ := json.MarshalIndent(logStatus, "", "  ")
		fmt.Printf("%s\nstatus water: %s\nstatus wind: %s\n", string(logJSON), waterStatus, windStatus)

		time.Sleep(15 * time.Second)
	}
}

func getStatus(value, safe, danger int) string {
	if value < safe {
		return "aman"
	} else if value >= safe && value <= danger {
		return "siaga"
	} else {
		return "bahaya"
	}
}