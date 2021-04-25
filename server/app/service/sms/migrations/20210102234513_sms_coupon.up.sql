CREATE TABLE `sms_coupon` (
  `id` bigint(20) unsigned NOT NULL COMMENT 'id',
  `app_id` bigint(20) unsigned NULL DEFAULT NULL ,
  `type` int DEFAULT NULL COMMENT '类型',
  `name` varchar(255) DEFAULT NULL COMMENT 'name',
  `platform` int DEFAULT NULL COMMENT '用户id',
  `count` int(11) DEFAULT NULL COMMENT '数量',
  `amount` decimal DEFAULT NULL COMMENT '价格',
  `create_by` bigint(20) DEFAULT NULL,
  `update_by` bigint(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;