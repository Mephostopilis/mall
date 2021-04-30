
CREATE TABLE `member_assets` (
  `member_id` bigint(20) NOT NULL COMMENT '会员ID',
  `balance` decimal(38,10) NOT NULL DEFAULT '0.0000000000' COMMENT '余额',
  `frozen` decimal(38,10) NOT NULL DEFAULT '0.0000000000' COMMENT '冻结',
  `create_by` bigint(20) DEFAULT NULL,
  `update_by` bigint(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`member_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='会员资产';