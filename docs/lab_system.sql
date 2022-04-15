DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `seats`;
DROP TABLE IF EXISTS `reservation`;
DROP TABLE IF EXISTS `laboratory`;
DROP TABLE IF EXISTS `announcement`;
DROP TABLE IF EXISTS `apply`;

CREATE TABLE announcement (
    id int(10) NOT NULL AUTO_INCREMENT COMMENT '主键，唯一标识一个公告',
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    title varchar(50) UNIQUE NOT NULL COMMENT '标题',
    content text DEFAULT NULL COMMENT '公告内容',
    PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
insert into announcement (id, title, content) value (1,'测试公告','hello world!!!!');

CREATE TABLE laboratory (
    lab_id int(10) NOT NULL AUTO_INCREMENT COMMENT '主键，唯一标识一个实验室',
    lab_name varchar(100) unique DEFAULT NULL COMMENT '实验室名字',
    person_number int(10) DEFAULT NULL COMMENT '可容纳人数',
    description varchar(500) DEFAULT NULL COMMENT '实验室信息描述',
    PRIMARY KEY (lab_id)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
insert into laboratory (lab_id, lab_name, person_number) values (1,'b3-234',40);
insert into laboratory (lab_id, lab_name, person_number) values (2,'b3-351',50);

CREATE TABLE seats (
   `id` int unsigned NOT NULL AUTO_INCREMENT,
   `seat_id` int unsigned NOT NULL,
   `lab_name` varchar(100) NOT NULL DEFAULT 0,
   `seat_name` varchar(100) NOT NULL,
   `status` int(10) NOT NULL DEFAULT 0 COMMENT '状态 0-正常 1-故障',
   PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
insert into seats (id, seat_id, lab_name, seat_name) values (0,1,'b3-234','A1');
insert into seats (id, seat_id, lab_name, seat_name) values (0,2,'b3-234','A2');
insert into seats (id, seat_id, lab_name, seat_name) values (0,3,'b3-234','A3');
insert into seats (id, seat_id, lab_name, seat_name) values (0,4,'b3-234','A4');
insert into seats (id, seat_id, lab_name, seat_name) values (0,5,'b3-234','A5');
insert into seats (id, seat_id, lab_name, seat_name) values (0,6,'b3-234','A6');
insert into seats (id, seat_id, lab_name, seat_name) values (0,7,'b3-234','A7');
insert into seats (id, seat_id, lab_name, seat_name) values (0,8,'b3-234','A8');
insert into seats (id, seat_id, lab_name, seat_name) values (0,9,'b3-234','B1');
insert into seats (id, seat_id, lab_name, seat_name) values (0,10,'b3-234','B2');
insert into seats (id, seat_id, lab_name, seat_name) values (0,11,'b3-234','B3');
insert into seats (id, seat_id, lab_name, seat_name) values (0,12,'b3-234','B4');
insert into seats (id, seat_id, lab_name, seat_name) values (0,13,'b3-234','B5');
insert into seats (id, seat_id, lab_name, seat_name) values (0,14,'b3-234','B6');
insert into seats (id, seat_id, lab_name, seat_name) values (0,15,'b3-234','B7');
insert into seats (id, seat_id, lab_name, seat_name) values (0,16,'b3-234','B8');
insert into seats (id, seat_id, lab_name, seat_name) values (0,17,'b3-234','C1');
insert into seats (id, seat_id, lab_name, seat_name) values (0,18,'b3-234','C2');
insert into seats (id, seat_id, lab_name, seat_name) values (0,19,'b3-234','C3');
insert into seats (id, seat_id, lab_name, seat_name) values (0,20,'b3-234','C4');
insert into seats (id, seat_id, lab_name, seat_name) values (0,21,'b3-234','C5');
insert into seats (id, seat_id, lab_name, seat_name) values (0,22,'b3-234','C6');
insert into seats (id, seat_id, lab_name, seat_name) values (0,23,'b3-234','C7');
insert into seats (id, seat_id, lab_name, seat_name) values (0,24,'b3-234','C8');
insert into seats (id, seat_id, lab_name, seat_name) values (0,25,'b3-234','D1');
insert into seats (id, seat_id, lab_name, seat_name) values (0,26,'b3-234','D2');
insert into seats (id, seat_id, lab_name, seat_name) values (0,27,'b3-234','D3');
insert into seats (id, seat_id, lab_name, seat_name) values (0,28,'b3-234','D4');
insert into seats (id, seat_id, lab_name, seat_name) values (0,29,'b3-234','D5');
insert into seats (id, seat_id, lab_name, seat_name) values (0,30,'b3-234','D6');
insert into seats (id, seat_id, lab_name, seat_name) values (0,31,'b3-234','D7');
insert into seats (id, seat_id, lab_name, seat_name) values (0,32,'b3-234','D8');
insert into seats (id, seat_id, lab_name, seat_name) values (0,33,'b3-234','E1');
insert into seats (id, seat_id, lab_name, seat_name) values (0,34,'b3-234','E2');
insert into seats (id, seat_id, lab_name, seat_name) values (0,35,'b3-234','E3');
insert into seats (id, seat_id, lab_name, seat_name) values (0,36,'b3-234','E4');
insert into seats (id, seat_id, lab_name, seat_name) values (0,37,'b3-234','E5');
insert into seats (id, seat_id, lab_name, seat_name) values (0,38,'b3-234','E6');
insert into seats (id, seat_id, lab_name, seat_name) values (0,39,'b3-234','E7');
insert into seats (id, seat_id, lab_name, seat_name) values (0,40,'b3-234','E8');
insert into seats (id, seat_id, lab_name, seat_name) values (0,1,'b3-351','A1');
insert into seats (id, seat_id, lab_name, seat_name) values (0,2,'b3-351','A2');
insert into seats (id, seat_id, lab_name, seat_name) values (0,3,'b3-351','A3');
insert into seats (id, seat_id, lab_name, seat_name) values (0,4,'b3-351','A4');

DROP TABLE IF EXISTS `reservation`;
CREATE TABLE reservation (
     id int(10) NOT NULL AUTO_INCREMENT COMMENT '实验室预约表的主键',
     user_name varchar(50) DEFAULT NULL,
     account char(12) DEFAULT NULL COMMENT '预约人账号',
     lab_name varchar(100) DEFAULT NULL COMMENT '实验室名字',
     seat_name varchar(100) NOT NULL,
     reserve_date varchar(100) DEFAULT NULL COMMENT '预约日期',
     time_interval varchar(100) DEFAULT NULL ,
     weekday int(10) DEFAULT NULL COMMENT '表示星期几，取值为1,2,3,4,5,6,7\r\n1--周一、2—周二\r\n3--周三、4—周四\r\n5--周五、6—周六\r\n7--周日',
     description varchar(500) DEFAULT NULL COMMENT '预约描述',
     PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
insert into reservation (user_name,account, lab_name, seat_name, reserve_date, time_interval) values ('张三','20182022001','B3-234','A1','2022-01-23','8:50:00-10:25:00');
insert into reservation (user_name,account, lab_name, seat_name, reserve_date, time_interval) values ('张三','20182022001','B3-234','A2','2022-01-23','8:50:00-10:25:00');
insert into reservation (user_name,account, lab_name, seat_name, reserve_date, time_interval) values ('张三','20182022001','B3-234','A3','2022-01-23','8:50:00-10:25:00');
insert into reservation (user_name,account, lab_name, seat_name, reserve_date, time_interval) values ('张三','20182022001','B3-234','A4','2022-01-23','8:50:00-10:25:00');

CREATE TABLE `users` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `account` char(12) NOT NULL,
    `name` varchar(50) DEFAULT NULL,
    `email` varchar(50) DEFAULT NULL,
    `telephone` varchar(11) DEFAULT NULL,
    `college` varchar(50) DEFAULT NULL,
    `password` varchar(255) DEFAULT NULL,
    `degree` varchar(10) DEFAULT NULL,
    `grade` varchar(10) DEFAULT NULL,
    `identity` varchar(10) DEFAULT 'student',
    `role_id` varchar(10) DEFAULT '2',
    `enable` smallint DEFAULT '1',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_account` (`account`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
insert into users (account,name,email, password, role_id) value ('201820220001','张三','zhansan@qq.com','123456','1');

CREATE TABLE apply (
      `id` int(10) NOT NULL AUTO_INCREMENT COMMENT '实验室申请表的主键',
      `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
      `user_name` varchar(50) DEFAULT NULL,
      `account` bigint unsigned DEFAULT NULL COMMENT '预约人账号',
      `lab_name` varchar(100) DEFAULT NULL COMMENT '实验室名字',
      `reserve_date` varchar(100) DEFAULT NULL COMMENT '预约日期',
      `status` int(10) NOT NULL DEFAULT 0 COMMENT '状态 0-申请中 1-申请成功 2-申请失败',
      `description` varchar(500) DEFAULT NULL COMMENT '申请描述',
      `dates` varchar(500) DEFAULT NULL COMMENT '申请区间',
      PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

CREATE TABLE `role` (
                        `id` INT PRIMARY KEY auto_increment,
                        `name` VARCHAR(100),
                        `remark` VARCHAR(100)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_role` (
                             `user_id` INT,
                             `role_id` INT
) ENGINE=InnoDB DEFAULT CHARSET=utf8;