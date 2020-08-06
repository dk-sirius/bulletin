package database

// 通道订阅者
type ChannelSubscriber struct {
	Common
	RefChannelID
	// 通道通知回调
	Callback string `db:"f_callback"`
}
