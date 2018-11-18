package models

/*
create table `video_info` (
`id` int(11) primary KEY auto_increment,
`user_id` int(11)  not null default 0,
`name` varchar(500) not null default '',
`display_ctime` varchar(100) not null default '',
`image_url` varchar(200) not null default '',
`file_url` varchar(200) not null default '',
`create_time` datetime not null default current_timestamp
)
*/




type VideoInfo struct {
	Id int `json:"id"`
	UserId  int `json:"user_id"`
	Name string `json:"name"`
	DisplayCtime string `json:"display_ctime"`
	FileUrl string `json:"file_url"`
}



func CreateVideo(name string, pwd string) error {
	stmtIns, err := db.Prepare("INSERT INTO video_info (user_id, name, display_ctime, file_url) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	stmtIns.Exec(name, pwd)
	stmtIns.Close()
	return nil
}

func GetVideo(id int) (VideoInfo, error) {
	stmtOut, err := db.Prepare("SELECT id,user_id, name, display_ctime, file_url FROM video_info WHERE id = ?")
	if err != nil {
		return VideoInfo{}, nil
	}

	video := VideoInfo{}
	stmtOut.QueryRow(id).Scan(&video.Id, &video.Name, &video.DisplayCtime, &video.FileUrl)
	stmtOut.Close()
	return video, nil
}

func DeleteVideo(id int)  error {
	stmtIns, err := db.Prepare("Delete from video_info where id = ?")
	if err != nil {
		return err
	}
	stmtIns.Exec(id)
	stmtIns.Close()
	return nil
}