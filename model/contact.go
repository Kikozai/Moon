package model

import (
	"fmt"

	u "github.com/Kikozai/Moon/utils"
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	UserID uint   `json:"user_id"`
}

/*
  Эта функция структуры проверяет необходимые параметры, отправленные через тело http-запроса.
возвращает сообщение и true, если требование выполнено
*/
func (contact *Contact) Valide() (map[string]interface{}, bool) {
	if contact.Name == "" {
		return u.Message(false, "Contact name should be on the payload"), false
	}

	if contact.Phone == "" {
		return u.Message(false, "Phone number should be on the payload"), false
	}

	if contact.UserID <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (contact *Contact) Create() map[string]interface{} {

	if resp, ok := contact.Valide(); !ok {
		return resp
	}

	GetDB().Create(contact)

	resp := u.Message(true, "success")
	resp["contact"] = contact
	return resp
}

func getContact(id uint) *Contact {
	contact := &Contact{}
	err := GetDB().Table("contacts").Where("id =?", id).First(contact).Error
	if err != nil {
		return nil
	}
	return contact
}

func GetContacts(user uint) []*Contact {
	contacts := make([]*Contact, 0)
	err := GetDB().Table("contacts").Where("user_id = ?", user).First(&contacts).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return contacts
}
