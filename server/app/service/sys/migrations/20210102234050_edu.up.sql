CREATE TABLE `sys_dict_data` (
`dict_code`  int(11) NOT NULL AUTO_INCREMENT ,
`app_id` bigint(20) unsigned NULL DEFAULT NULL ,
`dict_sort`  int(11) NULL DEFAULT NULL ,
`dict_label`  varchar(128)  NULL DEFAULT NULL ,
`dict_value`  varchar(255)  NULL DEFAULT NULL ,
`dict_type`  varchar(64)  NULL DEFAULT NULL ,
`css_class`  varchar(128)  NULL DEFAULT NULL ,
`list_class`  varchar(128)  NULL DEFAULT NULL ,
`is_default`  varchar(8)  NULL DEFAULT NULL ,
`status`  varchar(4)  NULL DEFAULT NULL ,
`default`  varchar(8)  NULL DEFAULT NULL ,
`create_by`  varchar(64)  NULL DEFAULT NULL ,
`update_by`  varchar(64)  NULL DEFAULT NULL ,
`remark`  varchar(255)  NULL DEFAULT NULL ,
`created_at`  datetime NULL DEFAULT NULL ,
`updated_at`  datetime NULL DEFAULT NULL ,
`deleted_at`  datetime NULL DEFAULT NULL ,
PRIMARY KEY (`dict_code`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=1;