package data

import (
	"crypto/rand"
	"fmt"
	"log"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

// 発生したセッションをデータベース内に永久保存(?今のところは)
func (user *User) CreateSession() (s Session, err error) {
	statement := "insert to sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
	stmt, err := Db.Prepare(statement) // 複数SQL文を実行できるように待機する
	if err != nil {
		return
	}
	defer stmt.Close()

	// 今のUserがセッションを持っていない場合に実行される
	err = stmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).Scan(&s.Id, &s.Uuid, &s.Email, &s.UserId, &s.CreatedAt)
	return
}

// セッションが有効かどうか
// 有効 valid = True
func (s *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1", s.Uuid).
		Scan(&s.Id, &s.Uuid, &s.Email, &s.UserId, &s.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if s.Id != 0 {
		valid = true
	}
	return
}

func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// RFC 4122 (0x40)
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}
