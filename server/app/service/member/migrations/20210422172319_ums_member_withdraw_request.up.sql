CREATE TABLE `member_withdraw_request` (
  `id` bigint(20) NOT NULL COMMENT 'Id',
  `order_no` char(28) NOT NULL COMMENT '单号',
  `type` tinyint(1) NOT NULL COMMENT '1：提现 2：退款',
  `required_parameter` text COMMENT '请求参数',
  `return_parameter` text NOT NULL COMMENT '返回参数',

  `create_by` bigint(20) DEFAULT NULL,
  `update_by` bigint(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,

  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_order` (`order_no`) USING BTREE,
  KEY `idx_user` (`created_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='提币记录';
