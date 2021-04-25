CREATE TABLE `sys_config` (
`config_id`  int(11) NOT NULL AUTO_INCREMENT ,
`app_id` bigint(20) unsigned NULL DEFAULT NULL ,
`config_name`  varchar(128)  NULL DEFAULT NULL ,
`config_key`  varchar(128)  NULL DEFAULT NULL ,
`config_value`  varchar(255)  NULL DEFAULT NULL ,
`config_type`  varchar(64)  NULL DEFAULT NULL ,
`remark`  varchar(128)  NULL DEFAULT NULL ,
`create_by`  varchar(128)  NULL DEFAULT NULL ,
`update_by`  varchar(128)  NULL DEFAULT NULL ,
`created_at`  datetime NULL DEFAULT NULL ,
`updated_at`  datetime NULL DEFAULT NULL ,
`deleted_at`  datetime NULL DEFAULT NULL ,
PRIMARY KEY (`config_id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=2;