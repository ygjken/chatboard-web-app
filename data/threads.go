package data

import "time"

type Thread struct {
	Id       int
	Uuid     string
	Topic    string
	UserId   int
	CreateAt time.Time
}

func Threads() (t Thread, err error) {
	// テストコード
	// TODO: スライスに変更する
	t = Thread{Id: 0, Uuid: "460", Topic: "Test", UserId: 100, CreateAt: time.Now()}
	err = nil
	return
}
