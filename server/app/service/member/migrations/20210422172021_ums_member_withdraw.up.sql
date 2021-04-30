CREATE TABLE `member_withdraw` (
  `id` bigint(20) NOT NULL COMMENT 'Id',
  `user_id` bigint(20) NOT NULL COMMENT '用户Id',
  `order_no` char(28) NOT NULL COMMENT '提现单号',
  `type` char(6) NOT NULL COMMENT '类型，alipay，wxpay, bank',
  `amount` decimal(38,4) NOT NULL COMMENT '金额',
  `poundage` decimal(38,4) NOT NULL COMMENT '手续费',
  `gold` decimal(38,4) NOT NULL COMMENT '扣除金币数量',
  `account` varchar(50) NOT NULL COMMENT '提现账号',
  `name` varchar(50) NOT NULL COMMENT '姓名',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态，0申请中，10已通过，100提现成功，-10已拒绝，-100提现失败',
  `reason` varchar(200) DEFAULT NULL COMMENT '拒绝、失败原因',
  `remark` varchar(200) DEFAULT NULL COMMENT '说明',
  `paid_time` datetime DEFAULT NULL COMMENT '提现完成时间',
  `audit_time` datetime DEFAULT NULL COMMENT '审核时间',
  `audit_by` varchar(50) DEFAULT NULL COMMENT '审核人',

  `create_by` bigint(20) DEFAULT NULL,
  `update_by` bigint(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_order` (`order_no`) USING BTREE,
  KEY `idx_user` (`user_id`,`created_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='提现记录';