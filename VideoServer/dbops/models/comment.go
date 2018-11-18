package models


/*
create table `comments` (
  `id` int(11) primary KEY auto_increment,
  `video_id` int(11)  not null default 0,
  `user_id` int(11)  not null default 0,
  `content` varchar(1024) not null default '',
  `create_time` datetime not null default current_timestamp,
  `status` tinyint(2) not null default 1
)

 */


type Commnet struct {
	Id int
	VideoId int
	UserId int
	Content string
	Status int
}




func CreateCommet( string, pwd string) error {
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