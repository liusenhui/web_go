package models

import "time"

type User struct {
	Id        uint      `gorm:"primary_key;auto_increment"`
	Account   string    `gorm:"type:varchar(50);unique_index"`
	Name      string    `gorm:"type:varchar(50);not null"`
	Password  string    `gorm:"type:varchar(100);not null"`
	CreatedBy string    `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time `gorm:"type:timestamp;not null"`
}

// TableName 设置表名（可选）
func (User) TableName() string {
	return "users"
}
