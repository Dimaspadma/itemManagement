package user

import (
	"awesomeProject/responsepb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type User struct {
	ID        int
	Username  string
	Password  string
	CreatedAt time.Time
}

func ModelToProto(user *User) *responsepb.User {
	return &responsepb.User{
		Id:        int32(user.ID),
		Username:  user.Username,
		CreatedAt: timestamppb.New(user.CreatedAt),
	}
}

func ModelToProtoList(users []*User) []*responsepb.User {
	var userList []*responsepb.User
	for _, user := range users {
		userList = append(userList, ModelToProto(user))
	}
	return userList
}

func ProtoToModel(user *responsepb.User) *User {
	return &User{
		ID:        int(user.Id),
		Username:  user.Username,
		Password:  "",
		CreatedAt: user.CreatedAt.AsTime(),
	}
}
