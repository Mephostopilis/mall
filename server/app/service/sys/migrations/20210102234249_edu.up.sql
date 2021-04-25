CREATE TABLE `sys_operlog` (
`oper_id`  int(11) NOT NULL AUTO_INCREMENT ,
`app_id` bigint(20) unsigned NULL DEFAULT NULL ,
`title`  varchar(255)  NULL DEFAULT NULL ,
`business_type`  varchar(128)  NULL DEFAULT NULL ,
`business_types`  varchar(128)  NULL DEFAULT NULL ,
`method`  varchar(128)  NULL DEFAULT NULL ,
`request_method`  varchar(128)  NULL DEFAULT NULL ,
`operator_type`  varchar(128)  NULL DEFAULT NULL ,
`oper_name`  varchar(128)  NULL DEFAULT NULL ,
`dept_name`  varchar(128)  NULL DEFAULT NULL ,
`oper_url`  varchar(255)  NULL DEFAULT NULL ,
`oper_ip`  varchar(128)  NULL DEFAULT NULL ,
`oper_location`  varchar(128)  NULL DEFAULT NULL ,
`oper_param`  varchar(255)  NULL DEFAULT NULL ,
`status`  varchar(4)  NULL DEFAULT NULL ,
`oper_time`  timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP ,
`json_result`  varchar(255)  NULL DEFAULT NULL ,
`create_by`  varchar(128)  NULL DEFAULT NULL ,
`update_by`  varchar(128)  NULL DEFAULT NULL ,
`remark`  varchar(255)  NULL DEFAULT NULL ,
`latency_time`  varchar(128)  NULL DEFAULT NULL ,
`user_agent`  varchar(255)  NULL DEFAULT NULL ,
`created_at`  datetime NULL DEFAULT NULL ,
`updated_at`  datetime NULL DEFAULT NULL ,
`deleted_at`  datetime NULL DEFAULT NULL ,
PRIMARY KEY (`oper_id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci
AUTO_INCREMENT=1;