package topology

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Topology 拓扑图数据结构
type Topology struct {
	ID bson.ObjectId `json:"id" bson:"_id"`

	// 历史记录用，表示源文件id
	Src bson.ObjectId `json:"src,omitempty" bson:"src,omitempty"`

	Name  string `json:"name"`
	Desc  string `json:"desc"`
	Data  bson.M `json:"data"`
	Image string `json:"image"`
	Class string `json:"class"`

	UserID   string `json:"userId" bson:"userId"`
	Username string `json:"username" `

	EditorID   string `json:"editorId"  bson:"editorId"`
	EditorName string `json:"editorName" bson:"editorName"`

	Shared    bool  `json:"shared"`
	Star      uint8 `json:"star" bson:"star,omitempty"`
	View      uint8 `json:"view" bson:"view,omitempty" `
	Recommend uint8 `json:"recommend" bson:"recommend,omitempty" `

	Tags []string `json:"tags"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt,omitempty"`
	DeletedAt time.Time `json:"deletedAt,omitempty" bson:"deletedAt,omitempty"`
}

// Favorite 用户收藏
type Favorite struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	UserID    string        `json:"userId" bson:"userId"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt,omitempty"`
}

// Star 点赞
type Star struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	UserID    string        `json:"userId" bson:"userId"`
	CreatedAt time.Time     `json:"createdAt" bson:"createdAt,omitempty"`
}
