package service

import (
	"chat-miniproject/contract"
	"chat-miniproject/repository"
)

func CreateUser(user contract.Users) (contract.Users, error) {
	inserted, err := repository.Save(user)

	if err != nil {
		return contract.Users{}, err
	}

	return inserted, nil
}

func FindAllUser() ([]contract.Users, error) {
	result, err := repository.FindAll()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func SendText(user contract.Chat) (contract.Chat, error) {
	inserted, err := repository.Send(user)

	if err != nil {
		return contract.Chat{}, err
	}

	return inserted, nil
}

func ReadText() ([]contract.Chat, error) {
	result, err := repository.ReadText()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func ChatRoom(name string) ([]contract.Chat, error) {
	result, err := repository.ChatRoom()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func FindById(id int) (contract.Users, error) {
	result, err := repository.FindById(id)
	if err != nil {
		return contract.Users{}, err
	}

	return result, nil
}
