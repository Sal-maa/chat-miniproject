package repository

import (
	"chat-miniproject/config"
	"chat-miniproject/contract"
	"fmt"
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
	return list, nil
}

func Save(user contract.Users) (contract.Users, error) {
	db, err := config.GormDatabaseConn()

	if err != nil {
		return contract.Users{}, err
	}

	db.Save(&user)

	return user, nil
}
