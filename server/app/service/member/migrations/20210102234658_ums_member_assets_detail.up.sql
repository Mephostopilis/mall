CREATE TABLE `ums_member_assets_detail` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `client_id` varchar(255) DEFAULT NULL COMMENT '客户端id',
  `user_id` varchar(255) DEFAULT NULL COMMENT '用户id',
  `redirect_uri` varchar(255) DEFAULT NULL COMMENT 'redirect',
  `scope` varchar(255) DEFAULT NULL COMMENT '范围',
  `code` varchar(255) DEFAULT NULL COMMENT 'code',
  `code_created_at` timestamp NULL DEFAULT NULL COMMENT 'code创建时',
  `code_expires_in` varchar(255) DEFAULT NULL COMMENT 'code创建时',
  `access` varchar(255) DEFAULT NULL COMMENT 'code',
  `access_created_at` timestamp NULL DEFAULT NULL COMMENT 'code创建时',
  `access_expires_in` varchar(255) DEFAULT NULL COMMENT 'code创建时',
  `refresh` varchar(255) DEFAULT NULL COMMENT 'code',
  `refresh_created_at` timestamp NULL DEFAULT NULL COMMENT 'code创建时',
  `refresh_expires_in` varchar(255) DEFAULT NULL COMMENT 'code创建时',
  
  `create_by` bigint(20) DEFAULT NULL,
  `update_by` bigint(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;