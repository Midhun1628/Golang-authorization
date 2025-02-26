package models

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
	RoleID   uint
}

type Role struct {
	RollNo           uint   `gorm:"primaryKey"`
	EmployeePosition string `gorm:"unique"`
}
