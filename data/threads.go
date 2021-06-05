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

func (t *Thread) GetCreateAt() string {
	return t.CreateAt.Format("Jan 2, 2006 at 3:04pm")
}

func (t *Thread) GetUser() (u User) {
	u = User{}
	Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", t.UserId).
		Scan(&u.Id, &u.Uuid, &u.Name, &u.Email, &u.CreatedAt)
	return
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

func GetThreadByUUID(uuid string) (t Thread, err error) {
	t = Thread{}
	err = Db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = $1", uuid).
		Scan(&t.Id, &t.Uuid, &t.Topic, &t.UserId, &t.CreateAt)
	return
}
