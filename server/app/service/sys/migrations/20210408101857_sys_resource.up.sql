CREATE TABLE `sys_resource` (
`id`  int(11) NOT NULL AUTO_INCREMENT ,
`app_id` bigint(20) unsigned NULL DEFAULT NULL ,
`module` varchar(128)  NULL DEFAULT NULL ,
`title`  varchar(128)  NULL DEFAULT NULL ,
`path`   varchar(128)  NULL DEFAULT NULL ,
`action`  varchar(16)  NULL DEFAULT NULL ,
`create_by`  varchar(128)  NULL DEFAULT NULL ,
`update_by`  varchar(128)  NULL DEFAULT NULL ,
`created_at`  datetime NULL DEFAULT NULL ,
`updated_at`  datetime NULL DEFAULT NULL ,
`deleted_at`  datetime NULL DEFAULT NULL ,
PRIMARY KEY (`id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=1;