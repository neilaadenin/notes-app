package models

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Notes    []Note `json:"-" gorm:"foreignKey:UserID"`
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

