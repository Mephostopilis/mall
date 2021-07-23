CREATE TABLE `ums_user` (
  `user_id`  int(11) NOT NULL,
  `app_id` bigint(20) unsigned NULL DEFAULT NULL ,
  `username`  varchar(64)  NULL DEFAULT NULL,
  `nick_name` varchar(128)  NULL DEFAULT NULL,
  `phone`     varchar(11)  NULL DEFAULT NULL,
  `role_id`  int(11) NULL DEFAULT NULL,
  `salt`    varchar(255)  NULL DEFAULT NULL ,
  `avatar`  varchar(255)  NULL DEFAULT NULL ,
  `sex`     int(11) NULL DEFAULT NULL ,
  `email`   varchar(128)  NULL DEFAULT NULL ,
  `dept_id` int(11) NULL DEFAULT NULL ,
  `post_id` int(11) NULL DEFAULT NULL ,
  `remark`  varchar(255)  NULL DEFAULT NULL ,
  `status`  int(11) NULL DEFAULT NULL COMMENT '状态',
  `create_by`  varchar(128)  NULL DEFAULT NULL ,
  `update_by`  varchar(128)  NULL DEFAULT NULL ,
  `created_at`  datetime NULL DEFAULT NULL ,
  `updated_at`  datetime NULL DEFAULT NULL ,
  `deleted_at`  datetime NULL DEFAULT NULL ,

  PRIMARY KEY (`user_id`),
  UNIQUE KEY `idx_username` (`username`) USING BTREE,
  UNIQUE KEY `idx_phone` (`phone`) USING BTREE
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8mb4 COLLATE=utf8mb4_general_ci;
