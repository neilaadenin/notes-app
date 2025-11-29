package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"backend/database"
	"backend/models"
	"backend/utils"
)

// Register user
// @Summary Register user
// @Description Daftarkan user baru
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.RegisterInput true "User data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	c.Bind(&body)

	// Cek email sudah terdaftar
	var existing models.User
	database.DB.Where("email = ?", body.Email).First(&existing)
	if existing.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email sudah terdaftar"})
		return
	}

	// Hash password
	hash, _ := utils.HashPassword(body.Password)

	user := models.User{
		Username: body.Username,
		Email:    body.Email,
		Password: hash,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		log.Printf("Error creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan user ke database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Register berhasil",
		"user_id": user.ID,
	})
}

// Login user
// @Summary Login user
// @Description Login dan dapatkan token JWT
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body models.LoginInput true "Login data"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	c.Bind(&body)

	var user models.User
	database.DB.Where("email = ?", body.Email).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email tidak ditemukan"})
		return
	}

	// cek password
	if !utils.CheckPassword(user.Password, body.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password salah"})
		return
	}

	// buat token
	token, _ := utils.GenerateToken(user.Email)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"token":   token,
	})
}

