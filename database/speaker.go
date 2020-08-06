package database

// 发布服务
type Speaker struct {
	Common
	RefSpeakerID
	// 演讲者名称
	Name string `db:"f_name"`
	// 备注
	Remark string `db:"f_remark"`
	// 频道授权共钥
	Identify string `db:"f_identify"`
}

type RefSpeakerID struct {
	// 演讲者ID
	SpeakerId UUID `db:"f_speaker_id"`
}
