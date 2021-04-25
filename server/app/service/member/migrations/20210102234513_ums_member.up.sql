CREATE TABLE `ums_member` (
  `id` bigint(20) NOT NULL COMMENT 'id',
  `secret` varchar(255) DEFAULT NULL COMMENT '密钥',
  `domain` varchar(255) DEFAULT NULL COMMENT 'domain',
  `user_id` varchar(255) DEFAULT NULL COMMENT '用户id',
  `scope` varchar(255) DEFAULT NULL COMMENT '权限范围',
  `grant_type` varchar(255) DEFAULT NULL COMMENT '授权类型',
  `create_by` bigint(20) DEFAULT NULL,
  `update_by` bigint(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;