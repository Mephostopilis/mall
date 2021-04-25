CREATE TABLE `sys_menu` (
`menu_id`  int(11) NOT NULL AUTO_INCREMENT ,
`app_id` bigint(20) unsigned NULL DEFAULT NULL ,
`menu_name`  varchar(128)  NULL DEFAULT NULL ,
`title`  varchar(128)  NULL DEFAULT NULL ,
`icon`  varchar(128)  NULL DEFAULT NULL ,
`path`  varchar(128)  NULL DEFAULT NULL ,
`paths`  varchar(128)  NULL DEFAULT NULL ,
`menu_type`  varchar(1)  NULL DEFAULT NULL ,
`action`  varchar(16)  NULL DEFAULT NULL ,
`permission`  varchar(255)  NULL DEFAULT NULL ,
`parent_id`  int(11) NULL DEFAULT NULL ,
`no_cache`  tinyint(1) NULL DEFAULT NULL ,
`breadcrumb`  varchar(255)  NULL DEFAULT NULL ,
`component`  varchar(255)  NULL DEFAULT NULL ,
`sort`  int(11) NULL DEFAULT NULL ,
`visible`  boolean  NULL DEFAULT NULL ,
`is_frame`  boolean  NULL DEFAULT '0' ,
`create_by`  varchar(128)  NULL DEFAULT NULL ,
`update_by`  varchar(128)  NULL DEFAULT NULL ,
`created_at`  datetime NULL DEFAULT NULL ,
`updated_at`  datetime NULL DEFAULT NULL ,
`deleted_at`  datetime NULL DEFAULT NULL ,
PRIMARY KEY (`menu_id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=1;