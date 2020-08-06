package models

import "github.com/dk-sirius/bulletin/database"

// 发布者发送消息，中心不负责处理数据
type DeliverData struct {
	// 投递数据
	PostData interface{} `json:"PostData"`
	// 主题ID
	TopicID database.UUID `json:"topicID"`
}
