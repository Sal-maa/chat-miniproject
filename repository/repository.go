package repository

import (
	"chat-miniproject/config"
	"chat-miniproject/contract"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func FindAll() ([]contract.Users, error) {
	var users []contract.Users

	db, err := config.GormDatabaseConn()

	if err != nil {
		return nil, err
	}

	db.Find(&users)

	return users, nil
}

func Send(user contract.Chat) (contract.Chat, error) {
	db, err := config.GormDatabaseConn()

	if err != nil {
		return contract.Chat{}, err
	}

	db.Create(&user)

	return user, nil
}

func ReadText() ([]contract.Chat, error) {
	var text []contract.Chat

	db, err := config.GormDatabaseConn()

	if err != nil {
		return nil, err
	}

	db.Find(&text)

	return text, nil
}

func ChatRoom() ([]contract.Chat, error) {

	db, _ := config.GormDatabaseConn()
	var name string
	var list []contract.Chat
	var newChat contract.Chat

	err := db.Raw("SELECT sent_to,`text` From chats  where sent_to = ?", name).Find(&list).Error

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	list = append(list, newChat)

	// for _, v := range list {
	// 	// list := append(list,newChat)
	// 	if v.Text!=nil{
	// 		list := append(list,newChat)
	// 	}else{
	// 		fmt.Printf("There is no chat in this room")
	// 	}
	// }
	return list, nil
}

func FindById(id int) (contract.Users, error) {
	var user contract.Users

	db, err := config.GormDatabaseConn()

	if err != nil {
		return contract.Users{}, err
	}

	errorNotFound := db.First(&user, id).Error
	errors.Is(errorNotFound, gorm.ErrRecordNotFound)

	return user, errorNotFound
}

func Save(user contract.Users) (contract.Users, error) {
	db, err := config.GormDatabaseConn()

	if err != nil {
		return contract.Users{}, err
	}

	db.Save(&user)

	return user, nil
}
