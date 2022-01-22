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
insert into laboratory (lab_id, lab_name, person_number) values (1,'b3-234',40);
insert into laboratory (lab_id, lab_name, person_number) values (2,'b3-351',50);

CREATE TABLE seats (
   `id` int unsigned NOT NULL AUTO_INCREMENT,
   `seat_id` int unsigned NOT NULL,
   `lab_id` int unsigned NOT NULL DEFAULT 0,
   `seat_name` varchar(100) NOT NULL,
   PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
insert into seats (id, seat_id, lab_id, seat_name) values (0,1,1,'A1');
insert into seats (id, seat_id, lab_id, seat_name) values (0,2,1,'A2');
insert into seats (id, seat_id, lab_id, seat_name) values (0,3,1,'A3');
insert into seats (id, seat_id, lab_id, seat_name) values (0,4,1,'A4');
insert into seats (id, seat_id, lab_id, seat_name) values (0,5,1,'A5');
insert into seats (id, seat_id, lab_id, seat_name) values (0,6,1,'A6');
insert into seats (id, seat_id, lab_id, seat_name) values (0,7,1,'A7');
insert into seats (id, seat_id, lab_id, seat_name) values (0,8,1,'A8');
insert into seats (id, seat_id, lab_id, seat_name) values (0,9,1,'B1');
insert into seats (id, seat_id, lab_id, seat_name) values (0,10,1,'B2');
insert into seats (id, seat_id, lab_id, seat_name) values (0,11,1,'B3');
insert into seats (id, seat_id, lab_id, seat_name) values (0,12,1,'B4');
insert into seats (id, seat_id, lab_id, seat_name) values (0,13,1,'B5');
insert into seats (id, seat_id, lab_id, seat_name) values (0,14,1,'B6');
insert into seats (id, seat_id, lab_id, seat_name) values (0,15,1,'B7');
insert into seats (id, seat_id, lab_id, seat_name) values (0,16,1,'B8');
insert into seats (id, seat_id, lab_id, seat_name) values (0,17,1,'C1');
insert into seats (id, seat_id, lab_id, seat_name) values (0,18,1,'C2');
insert into seats (id, seat_id, lab_id, seat_name) values (0,19,1,'C3');
insert into seats (id, seat_id, lab_id, seat_name) values (0,20,1,'C4');
insert into seats (id, seat_id, lab_id, seat_name) values (0,21,1,'C5');
insert into seats (id, seat_id, lab_id, seat_name) values (0,22,1,'C6');
insert into seats (id, seat_id, lab_id, seat_name) values (0,23,1,'C7');
insert into seats (id, seat_id, lab_id, seat_name) values (0,24,1,'C8');
insert into seats (id, seat_id, lab_id, seat_name) values (0,25,1,'D1');
insert into seats (id, seat_id, lab_id, seat_name) values (0,26,1,'D2');
insert into seats (id, seat_id, lab_id, seat_name) values (0,27,1,'D3');
insert into seats (id, seat_id, lab_id, seat_name) values (0,28,1,'D4');
insert into seats (id, seat_id, lab_id, seat_name) values (0,29,1,'D5');
insert into seats (id, seat_id, lab_id, seat_name) values (0,30,1,'D6');
insert into seats (id, seat_id, lab_id, seat_name) values (0,31,1,'D7');
insert into seats (id, seat_id, lab_id, seat_name) values (0,32,1,'D8');
insert into seats (id, seat_id, lab_id, seat_name) values (0,33,1,'E1');
insert into seats (id, seat_id, lab_id, seat_name) values (0,34,1,'E2');
insert into seats (id, seat_id, lab_id, seat_name) values (0,35,1,'E3');
insert into seats (id, seat_id, lab_id, seat_name) values (0,36,1,'E4');
insert into seats (id, seat_id, lab_id, seat_name) values (0,37,1,'E5');
insert into seats (id, seat_id, lab_id, seat_name) values (0,38,1,'E6');
insert into seats (id, seat_id, lab_id, seat_name) values (0,39,1,'E7');
insert into seats (id, seat_id, lab_id, seat_name) values (0,40,1,'E8');



CREATE TABLE reservation (
     id int(10) NOT NULL AUTO_INCREMENT COMMENT '实验室预约表的主键',
     user_id int(10) DEFAULT NULL COMMENT '预约人id',
     lab_id int(10) DEFAULT NULL COMMENT '实验室id',
     seat_id int(10) DEFAULT NULL COMMENT '座位id',
     seat_name varchar(100) NOT NULL,
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