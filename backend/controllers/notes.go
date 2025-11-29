package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"backend/database"
	"backend/models"
)

// GetNotes
// @Summary Get all notes
// @Description Ambil semua notes milik user yang sedang login
// @Tags Notes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Note
// @Failure 401 {object} map[string]string
// @Router /notes [get]
func GetNotes(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var notes []models.Note
	database.DB.Where("user_id = ?", userID).Order("created_at DESC").Find(&notes)

	c.JSON(http.StatusOK, notes)
}

// CreateNote
// @Summary Create a new note
// @Description Buat note baru
// @Tags Notes
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param note body models.CreateNoteInput true "Note data"
// @Success 200 {object} models.Note
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /notes [post]
func CreateNote(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	note := models.Note{
		Title:  body.Title,
		Content: body.Content,
		UserID: userID.(uint),
	}

	database.DB.Create(&note)

	c.JSON(http.StatusOK, note)
}

