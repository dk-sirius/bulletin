package database

import (
	"strconv"
	"time"
)

// 唯一ID
type UUID uint64

func (uuid *UUID) Value() uint64 {
	return uint64(*uuid)
}

func (uuid *UUID) String() string {
	return strconv.FormatUint(uuid.Value(), 10)
}

// 时间
type SQLTime time.Time

func (t *SQLTime) Unix() int64 {
	return time.Time(*t).Unix()
}

func (t *SQLTime) String() string {
	if t.Unix() == 0 {
		return ""
	}
	return time.Time(*t).In(time.UTC).Format(time.RFC3339)
}

func (t *SQLTime) Format(layout string) string {
	return time.Time(*t).In(time.UTC).Format(layout)
}

// 数据库通用
type Common struct {
	ID UUID `db:"f_id"`
	TimeUnit
}

type TimeUnit struct {
	CreatedTime SQLTime `db:"created_at"`
	UpdatedTime SQLTime `db:"updated_at"`
	DeletedTime SQLTime `db:"deleted_at"`
}
