package service

import (
	. "blog/model"
	. "blog/config"
)

type CourseService struct {} 

func (cs CourseService) GetCourse() ([]Course) {
	db:=GetDB()
	cc:=[]Course{}
	db.First(&cc)
	
	return cc
}
