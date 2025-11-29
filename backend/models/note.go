package models

import "time"

type Note struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	UserID    uint      `json:"userId" gorm:"index"`
	User      User      `json:"-" gorm:"foreignKey:UserID"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
}

type CreateNoteInput struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

