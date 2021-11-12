package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;column:name;type:varchar(64);not null;default:'';comment:姓名"`
	Age  int    `gorm:"index;column:age;type:int(8);not null;default:0;comment:年龄"`
}

func (u User) TableName() string {
	return "user"
}
