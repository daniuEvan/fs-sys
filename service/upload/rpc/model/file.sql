-- 创建文件表
CREATE TABLE `fs_files`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `file_hash`   char(40)        NOT NULL DEFAULT '' COMMENT '文件hash',
#     `file_name`   varchar(256)  NOT NULL DEFAULT '' COMMENT '文件名',
    `file_size`   bigint          NOT NULL COMMENT '文件大小',
    `file_addr`   varchar(1024)   NOT NULL DEFAULT '' COMMENT '文件存储位置',
    `create_time` timestamp       null     default current_timestamp,
    `update_time` timestamp       null     default current_timestamp on update current_timestamp,
    `status`      int             NOT NULL DEFAULT '0' COMMENT '状态(可用/禁用/已删除等状态)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_file_hash` (`file_hash`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

-- 创建用户文件表
CREATE TABLE `fs_user_file`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`     bigint          NOT NULL,
    `file_hash`   varchar(64)     NOT NULL DEFAULT '' COMMENT '文件hash',
    `file_size`   bigint          NOT NULL COMMENT '文件大小',
    `file_name`   varchar(256)    NOT NULL DEFAULT '' COMMENT '文件名',
    `create_time` timestamp       null     default current_timestamp,
    `update_time` timestamp       null     default current_timestamp on update current_timestamp,
    `status`      int             NOT NULL DEFAULT '0' COMMENT '文件状态(0正常1已删除2禁用)',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_file` (`user_id`, `file_hash`,`file_name`),
    KEY `idx_status` (`status`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_file_name` (`file_name`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
