package register

import "github.com/dk-sirius/bulletin/database"

// 注册演讲者
type RegSpeakerParams struct {
	// 名称
	Name string `json:"name"`
	// 备注
	Remark string `json:"remark"`
}

// 注册演讲者返回
type RegSpeakerResp struct {
	// 身份ID
	SpeakerID database.UUID `json:"speakerID"`
	// 身份签名
	Identity string `json:"identity"`
}

// 注册通道参数
type RegChannelParams struct {
	Callback string `json:"callback"`
	Token    string `json:"token"`
}

// 注册通道接受者返回
type RegChannelSubscriberResp struct {
	// 当前通道演讲者
	SpeakerID database.UUID `json:"speakerID"`
	// 当前通道演讲者名称
	SpeakerName string `json:"speakerName"`
}
