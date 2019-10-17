package model

type Course struct {
	Cid int `gorm:"PRIMARY_KEY"`
	TeacherId int
	Cname string
}

func (c Course) TableName() string {
	return "course"
}