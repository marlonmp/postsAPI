package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type Model struct {
	Id *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`

	NickName string `json:"nick_name,omitempty" bson:"nick_name,omitempty"`
	UserName string `json:"user_name,omitempty" bson:"user_name,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}
