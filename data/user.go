package data

import "time"

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
	// err = stmt.QueryRow(CreateUUID)
	return
}

func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", email).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	return
}
