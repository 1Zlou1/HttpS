package repository

import "HttpS/service/entity"

type Repository interface {
	CreateFriend(in entity.Friend) (string, error)

	MakeFriend(in1 string, in2 string) (string, string, error)

	DeleteFriend(in string) (string, error)

	GetFriends(in string) ([]entity.Friend, error)

	UpdateAge(in string, age int) error
}
