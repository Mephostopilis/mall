CREATE TABLE `member_product_category_relation` (
  `id` bigint(20) unsigned NOT NULL,
  `member_id` bigint(20) unsigned DEFAULT NULL,
  `product_category_id` bigint(20) unsigned DEFAULT NULL,
  `create_by` bigint(20) DEFAULT NULL,
  `update_by` bigint(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='会员与产品分类关系表（用户喜欢的分类）';