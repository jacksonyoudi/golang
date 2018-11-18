create table `users` (
  `id` int(11) primary KEY auto_increment,
  `user_name` varchar(64)  not null default '',
  `pwd` varchar(1024) not null default '',
  unique key (`user_name`)
)



create table `video_info` (
  `id` int(11) primary KEY auto_increment,
  `user_id` int(11)  not null default 0,
  `name` varchar(500) not null default '',
  `display_ctime` varchar(100) not null default '',
  `image_url` varchar(200) not null default '',
  `file_url` varchar(200) not null default '',
  `create_time` datetime not null default current_timestamp
)


create table `comments` (
  `id` int(11) primary KEY auto_increment,
  `video_id` int(11)  not null default 0,
  `user_id` int(11)  not null default 0,
  `content` varchar(1024) not null default '',
  `create_time` datetime not null default current_timestamp,
  `status` tinyint(2) not null default 1
)

create table `sessions` (
  `ssssion_id` varchar(1024) primary KEY,
  `TTL` varchar(1024),
  `user_name` varchar(500) not null default '',
)


INSERT INTO users (user_name, pwd) valuse (?, ?)