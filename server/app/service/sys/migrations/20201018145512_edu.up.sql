CREATE TABLE `casbin_rule` (
`id` bigint(20) NOT NULL AUTO_INCREMENT,
`app_id` bigint(20) unsigned NULL DEFAULT NULL ,
`p_type`  varchar(100) NULL DEFAULT NULL,
`v0`  varchar(100)     NULL DEFAULT NULL,
`v1`  varchar(100)     NULL DEFAULT NULL,
`v2`  varchar(100)     NULL DEFAULT NULL,
`v3`  varchar(100)     NULL DEFAULT NULL,
`v4`  varchar(100)     NULL DEFAULT NULL,
`v5`  varchar(100)     NULL DEFAULT NULL,
PRIMARY KEY (`id`)
)
ENGINE=InnoDB
DEFAULT CHARACTER SET=utf8 COLLATE=utf8_general_ci;