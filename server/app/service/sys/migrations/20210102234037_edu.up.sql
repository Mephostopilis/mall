CREATE TABLE `sys_dept` (
`dept_id`  int(11) NOT NULL AUTO_INCREMENT ,
`app_id` bigint(20) unsigned NULL DEFAULT NULL ,
`parent_id`  int(11) NULL DEFAULT NULL ,
`dept_path`  varchar(255)  NULL DEFAULT NULL ,
`dept_name`  varchar(128)  NULL DEFAULT NULL ,
`sort`  int(11) NULL DEFAULT NULL ,
`leader`  varchar(128)  NULL DEFAULT NULL ,
`phone`  varchar(11)  NULL DEFAULT NULL ,
`email`  varchar(64)  NULL DEFAULT NULL ,
`status`  varchar(4)  NULL DEFAULT NULL ,
`create_by`  varchar(64)  NULL DEFAULT NULL ,
`update_by`  varchar(64)  NULL DEFAULT NULL ,
`created_at`  datetime NULL DEFAULT NULL ,
`updated_at`  datetime NULL DEFAULT NULL ,
`deleted_at`  datetime NULL DEFAULT NULL ,
PRIMARY KEY (`dept_id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=1;