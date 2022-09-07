package model

import (
	"demo_cms_1/dao"
	"demo_cms_1/hash"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DemoAdmin struct {
	gorm.Model
	Name     string
	PossWord string
	Status   bool
	Email    string
	Phone    string
}

func AdminPostMode(username string, password string, phone string, email string) {
	adminPost := DemoAdmin{
		Name:     username,
		PossWord: hash.ScryptPw(password),
		Status:   true,
		Email:    email,
		Phone:    phone,
	}
	dao.DB.Create(&adminPost)
	return
}

func AdminPutMode(username string, password string, phone string, email string, id string) {
	adminPost := DemoAdmin{
		Name:     username,
		PossWord: hash.ScryptPw(password),
		Phone:    phone,
		Email:    email,
	}
	dao.DB.Model(&DemoAdmin{}).Where("ID = ?", id).Updates(adminPost)
	return
}

func LoginModel(username, password string) int {
	var adminList []DemoAdmin
	dataName := dao.DB.Where("BINARY name = ?", username).First(&adminList)

	if dataName.RowsAffected == 0 {
		return 1
	}
	dataPw := dao.DB.Where("BINARY poss_word = ?", hash.ScryptPw(password)).First(&adminList)
	if dataPw.RowsAffected == 0 {
		return 2
	}
	return 0
}
