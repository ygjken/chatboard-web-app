package data

import (
	"time"
)

type Thread struct {
	Id       int
	Uuid     string
	Topic    string
	UserId   int
	CreateAt time.Time
}

func GetThreads() (t []Thread, err error) {
	rows, err := Db.Query("select id, uuid, topic, user_id, created_at from threads order by created_at desc")
	if err != nil {
		return
	}
	for rows.Next() {
		c := Thread{}
		err = rows.Scan(&c.Id, &c.Uuid, &c.Topic, &c.UserId, &c.CreateAt)
		if err != nil {
			return
		}
		t = append(t, c)
	}
	rows.Close()
	return
}
