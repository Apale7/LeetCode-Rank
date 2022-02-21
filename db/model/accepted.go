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

func (a *Accepted) Sub(b *Accepted) *Accepted {
	return &Accepted{
		Meta:   a.Meta,
		UserID: a.UserID,
		Hard:   a.Hard - b.Hard,
		Medium: a.Medium - b.Medium,
		Easy:   a.Easy - b.Easy,
	}
}

func NewAccepted() *Accepted {
	return &Accepted{
		Meta: *NewMeta(),
	}
}
