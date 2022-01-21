DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `seats`;
DROP TABLE IF EXISTS `reservation`;
DROP TABLE IF EXISTS `laboratory`;

DROP TABLE IF EXISTS laboratory;
CREATE TABLE laboratory (
    lab_id int(10) NOT NULL AUTO_INCREMENT COMMENT '主键，唯一标识一个实验室',
    lab_name varchar(50) DEFAULT NULL COMMENT '实验室名字',
    person_number int(10) DEFAULT NULL COMMENT '可容纳人数',
    description varchar(500) DEFAULT NULL COMMENT '实验室信息描述',
    PRIMARY KEY (lab_id)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
insert into laboratory (lab_id, lab_name, lab_id) values (0,'b3-111',50);
insert into laboratory (lab_id, lab_name, lab_id) values (1,'b3-222',50);

CREATE TABLE seats (
   `id` int unsigned NOT NULL AUTO_INCREMENT,
   `seat_id` int unsigned NOT NULL,
   `lab_id` int unsigned NOT NULL DEFAULT 0,
   PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
insert into seats (id, seat_id, lab_id) values (0,1,0);
insert into seats (id, seat_id, lab_id) values (1,2,0);
insert into seats (id, seat_id, lab_id) values (2,3,0);
insert into seats (id, seat_id, lab_id) values (3,1,1);
insert into seats (id, seat_id, lab_id) values (4,2,1);
insert into seats (id, seat_id, lab_id) values (5,3,1);
insert into seats (id, seat_id, lab_id) values (6,4,1);
insert into seats (id, seat_id, lab_id) values (7,5,1);


CREATE TABLE reservation (
     id int(10) NOT NULL AUTO_INCREMENT COMMENT '实验室预约表的主键',
     user_id int(10) DEFAULT NULL COMMENT '预约人id',
     lab_id int(10) DEFAULT NULL COMMENT '实验室id',
     seat_id int(10) DEFAULT NULL COMMENT '座位id',
     reserve_date varchar(100) DEFAULT NULL COMMENT '预约日期',
     time_interval varchar(100) DEFAULT NULL ,
     weekday int(10) DEFAULT NULL COMMENT '表示星期几，取值为1,2,3,4,5,6,7\r\n1--周一、2—周二\r\n3--周三、4—周四\r\n5--周五、6—周六\r\n7--周日',
     description varchar(500) DEFAULT NULL COMMENT '预约描述',
     PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

CREATE TABLE `user` (
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
    `role_name` varchar(10) DEFAULT 'user',
    `enable` smallint DEFAULT '1',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_account` (`account`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8;
insert into user (email, password, role_name) value ('zhansan@qq.com','123456','student');