CREATE TABLE `ums_loginlog` (
`info_id`  int(11) NOT NULL AUTO_INCREMENT ,
`app_id` bigint(20) unsigned NULL DEFAULT NULL ,
`username`  varchar(128)  NULL DEFAULT NULL ,
`status`  varchar(4)  NULL DEFAULT NULL ,
`ipaddr`  varchar(255)  NULL DEFAULT NULL ,
`login_location`  varchar(255)  NULL DEFAULT NULL ,
`browser`  varchar(255)  NULL DEFAULT NULL ,
`os`  varchar(255)  NULL DEFAULT NULL ,
`platform`  varchar(255)  NULL DEFAULT NULL ,
`login_time`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP ,
`create_by`  varchar(128)  NULL DEFAULT NULL ,
`update_by`  varchar(128)  NULL DEFAULT NULL ,
`remark`  varchar(255)  NULL DEFAULT NULL ,
`msg`  varchar(255)  NULL DEFAULT NULL ,
`created_at`  datetime NULL DEFAULT NULL ,
`updated_at`  datetime NULL DEFAULT NULL ,
`deleted_at`  datetime NULL DEFAULT NULL ,
PRIMARY KEY (`info_id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=1;