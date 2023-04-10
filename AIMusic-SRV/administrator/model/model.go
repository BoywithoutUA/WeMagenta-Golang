package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        int            `gorm:"primary_key"`
	CreatedAt time.Time      `gorm:"column:add_time;index:idx_created_time"`
	UpdatedAt time.Time      `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt `gorm:"column:delete_time"`
	IsDeleted bool           `gorm:"column:is_deleted"`
}

type User struct {
	BaseModel
	Mobile   string `gorm:"type:varchar(16);unique;not null"`
	Password string `gorm:"type:varchar(128);not null"`
	NickName string `gorm:"type:varchar(32);not null"`
	Motto    string `gorm:"type:varchar(32)"`
	Avatar   string `gorm:"type:varchar(512)"`
	Gender   int32  `gorm:"type:tinyint;not null"`
	Birthday string `gorm:"type:varchar(16)"`
	Likes    int64  `gorm:"type:int;default:0;index:idx_likes;not null"`
	Report   int64  `gorm:"type:int;default:0;index:idx_likes;not null"`
	Banned   int32  `gorm:"type:tinyint;default:0"`
}

type Composition struct {
	BaseModel
	CreatorMobile   string `gorm:"type:varchar(16);index:idx_composition_name,priority:1;index:idx_composition_room,priority:1;index:idx_composition_time,priority:1;not null"`
	CompositionName string `gorm:"type:varchar(16);index:idx_composition_name;not null"`
	CreatorNickname string `gorm:"type:varchar(32);not null"`
	Avatar          string `gorm:"type:varchar(512)"`
	CreatedFor      string `gorm:"type:varchar(16);index:idx_composition_for;not null"`
	CreatedForType  int32  `gorm:"type:tinyint;not null"`
	Detail          string `gorm:"type:varchar(256)"`
	MP3             string `gorm:"type:varchar(256);not null"`
	Status          int32  `gorm:"type:tinyint;not null"`
	Likes           int64  `gorm:"type:int;default:0;index:idx_likes;not null"`
	Report          int64  `gorm:"type:int;default:0;index:idx_likes;not null"`
}

type Template struct {
	BaseModel
	Name   string `gorm:"type:varchar(16);unique;not null"`
	Type   int32  `gorm:"type:tinyint;not null"`
	Detail string `gorm:"type:text;not null"`
	Pic    string `gorm:"type:varchar(256);not null"`
}

type Feedback struct {
	BaseModel
	Solved       int32  `gorm:"type:tinyint"`
	UserNickName string `gorm:"type:varchar(32);not null"`
	Avatar       string `gorm:"type:varchar(512);"`
	UserMobile   string `gorm:"type:varchar(16);not null"`
	Target       string `gorm:"type:varchar(32);not null"`
	Evidence     string `gorm:"type:varchar(256);not null"`
	Advice       string `gorm:"type:varchar(256);not null"`
}

type Usage struct {
	BaseModel
	Scratch     int64 `gorm:"type:int;default:0;not null"`
	Composition int64 `gorm:"type:int;default:0;not null"`
}

func (User) TableName() string {
	return "user"
}

func (Composition) TableName() string {
	return "composition"
}

func (Template) TableName() string {
	return "template"
}

func (Feedback) TableName() string {
	return "feedback"
}

func (Usage) TableName() string {
	return "usage"
}
