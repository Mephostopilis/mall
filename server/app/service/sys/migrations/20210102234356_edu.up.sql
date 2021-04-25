CREATE TABLE `sys_role_menu` (
`id` int(11) NOT NULL AUTO_INCREMENT ,
`app_id` bigint(20) unsigned NULL DEFAULT NULL ,
`role_id`  int(11) NULL DEFAULT NULL ,
`menu_id`  int(11) NULL DEFAULT NULL ,
`role_name`  longtext  NULL ,
`create_by`  longtext  NULL ,
`update_by`  longtext  NULL ,
PRIMARY KEY (`id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci;