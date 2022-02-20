package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Accepted 某一时刻该用户ac的题数
type Accepted struct {
	Meta   `bson:"inline"`
	UserID primitive.ObjectID `bson:"user_id"`
	Hard   int                `bson:"hard"`
	Medium int                `bson:"medium"`
	Easy   int                `bson:"easy"`
}

func NewAccepted() *Accepted {
	return &Accepted{
		Meta: *NewMeta(),
	}
}
