CREATE TABLE `sys_job` (
`job_id`  int(11) NOT NULL AUTO_INCREMENT ,
`app_id` bigint(20) unsigned NULL DEFAULT NULL ,
`job_name`  varchar(255)  NULL DEFAULT NULL ,
`job_group`  varchar(255)  NULL DEFAULT NULL ,
`job_type`  int(11) NULL DEFAULT NULL ,
`cron_expression`  varchar(255)  NULL DEFAULT NULL ,
`invoke_target`  varchar(255)  NULL DEFAULT NULL ,
`misfire_policy`  int(11) NULL DEFAULT NULL ,
`concurrent`  int(11) NULL DEFAULT NULL ,
`status`  int(11) NULL DEFAULT NULL ,
`entry_id`  int(11) NULL DEFAULT NULL ,
`create_by`  varchar(128)  NULL DEFAULT NULL ,
`update_by`  varchar(128)  NULL DEFAULT NULL ,
`created_at`  datetime NULL DEFAULT NULL ,
`updated_at`  datetime NULL DEFAULT NULL ,
`deleted_at`  datetime NULL DEFAULT NULL ,
PRIMARY KEY (`job_id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=1;