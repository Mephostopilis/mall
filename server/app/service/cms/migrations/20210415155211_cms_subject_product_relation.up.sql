CREATE TABLE `cms_subject_product_relation` (
  `id` bigint(20) unsigned NOT NULL,
  `app_id` bigint(20) unsigned NULL DEFAULT NULL ,
  `subject_id` bigint(20) unsigned DEFAULT NULL,
  `product_id` bigint(20) unsigned DEFAULT NULL,
  `create_by` bigint(20) unsigned DEFAULT NULL,
  `update_by` bigint(20) unsigned DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8 COMMENT='专题商品关系表';