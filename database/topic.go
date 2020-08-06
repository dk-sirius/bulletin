package database

// 话题
type Topic struct {
	Common
	RefChannelID
	RefTopicID
	Name string `db:"f_name"`
}

type RefTopicID struct {
	TopicID UUID `db:"f_topic_id"`
}
