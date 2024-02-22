package friend_servise

import "net/http"

type FriendService interface {
	CreateFriend(w http.ResponseWriter, r *http.Request)

	MakeFriend(iw http.ResponseWriter, r *http.Request)

	DeleteFriend(w http.ResponseWriter, r *http.Request)

	GetFriends(w http.ResponseWriter, r *http.Request)

	UpdateAge(w http.ResponseWriter, r *http.Request)
}
