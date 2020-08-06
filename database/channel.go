package database

// 通道
type Chanel struct {
	Common
	RefChannelID
	RefSpeakerID
}

type RefChannelID struct {
	// 频道ID
	ChannelID UUID `db:"f_channel_id"`
}
