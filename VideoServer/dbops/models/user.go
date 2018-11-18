package models

type User struct {
	Id       int `json:"id"`
	UserName string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

func CreateUser(name string, pwd string) error {
	stmtIns, err := db.Prepare("INSERT INTO users (user_name, pwd) values (?, ?)")
	if err != nil {
		return err
	}
	stmtIns.Exec(name, pwd)
	stmtIns.Close()
	return nil
}

func GetUser(name string) (User, error) {
	stmtOut, err := db.Prepare("SELECT id,user_name, pwd FROM users WHERE user_name = ?")
	if err != nil {
		return User{}, nil
	}

	user := User{}
	stmtOut.QueryRow(name).Scan(&user.Id, &user.UserName, &user.Pwd)
	stmtOut.Close()
	return user, nil
}

func DeleteUser(name, pwd string)  error {
	stmtIns, err := db.Prepare("Delete from users (user_name, pwd) valuse (?, ?)")
	if err != nil {
		return err
	}
	stmtIns.Exec(name, pwd)
	stmtIns.Close()
	return nil
}