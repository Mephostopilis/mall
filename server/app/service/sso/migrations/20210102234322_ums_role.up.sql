CREATE TABLE `ums_role` (
`role_id`  int(11) NOT NULL AUTO_INCREMENT ,
`app_id` bigint(20) unsigned NULL DEFAULT NULL ,
`role_name`  varchar(128)  NULL DEFAULT NULL ,
`status`  varchar(4)  NULL DEFAULT NULL ,
`role_key`  varchar(128)  NULL DEFAULT NULL ,
`role_sort`  int(11) NULL DEFAULT NULL ,
`flag`  varchar(128)  NULL DEFAULT NULL ,
`remark`  varchar(255)  NULL DEFAULT NULL ,
`admin`  tinyint(1) NULL DEFAULT NULL ,
`create_by`  varchar(128)  NULL DEFAULT NULL ,
`update_by`  varchar(128)  NULL DEFAULT NULL ,
`created_at`  datetime NULL DEFAULT NULL ,
`updated_at`  datetime NULL DEFAULT NULL ,
`deleted_at`  datetime NULL DEFAULT NULL ,
PRIMARY KEY (`role_id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=1;