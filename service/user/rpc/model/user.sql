create table `fs_user`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `username`    varchar(255) not null default '' unique comment '用户名',
    `mobile`      varchar(11)  not null default '' unique comment '用户电话',
    `password`    varchar(255) not null default '' comment '密码',
    `create_time` timestamp null default current_timestamp,
    `update_time` timestamp null default current_timestamp on update current_timestamp,
    primary key (`id`)
) engine=InnoDB default charset=utf8mb4;