CREATE TABLE `member_assets_detail` (
  `id` bigint(20) unsigned NOT NULL COMMENT '主键',
  `member_id` bigint(20) unsigned NOT NULL COMMENT '会员ID',
  `bill_no` varchar(50) NOT NULL COMMENT '流水号',
  `type` int(11) NOT NULL COMMENT '类型（1充值，2商品购买）',
  `amount` decimal(38,10) NOT NULL COMMENT '变动金额',
  `balance` decimal(38,10) NOT NULL COMMENT '变动后金额',
  `time` varchar(12) NOT NULL COMMENT '日期',
  `create_by` bigint(20) DEFAULT NULL,
  `update_by` bigint(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='会员资产变更明细';