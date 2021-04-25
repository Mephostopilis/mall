CREATE TABLE `ums_user` (
`user_id`  int(11) NOT NULL AUTO_INCREMENT ,
`app_id` bigint(20) unsigned NULL DEFAULT NULL ,
`nick_name`  varchar(128)  NULL DEFAULT NULL ,
`phone`  varchar(11)  NULL DEFAULT NULL ,
`role_id`  int(11) NULL DEFAULT NULL ,
`salt`  varchar(255)  NULL DEFAULT NULL ,
`avatar`  varchar(255)  NULL DEFAULT NULL ,
`sex`  varchar(255)  NULL DEFAULT NULL ,
`email`  varchar(128)  NULL DEFAULT NULL ,
`dept_id`  int(11) NULL DEFAULT NULL ,
`post_id`  int(11) NULL DEFAULT NULL ,
`remark`  varchar(255)  NULL DEFAULT NULL ,
`status`  varchar(4)  NULL DEFAULT NULL ,
`create_by`  varchar(128)  NULL DEFAULT NULL ,
`update_by`  varchar(128)  NULL DEFAULT NULL ,
`created_at`  datetime NULL DEFAULT NULL ,
`updated_at`  datetime NULL DEFAULT NULL ,
`deleted_at`  datetime NULL DEFAULT NULL ,
`username`  varchar(64)  NULL DEFAULT NULL ,
`password`  varchar(128)  NULL DEFAULT NULL ,
PRIMARY KEY (`user_id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=1;