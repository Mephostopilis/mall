CREATE TABLE `member_member_tag_relation` (
  `id` bigint(20) unsigned NOT NULL,
  `member_id` bigint(20) unsigned DEFAULT NULL,
  `tag_id` bigint(20) unsigned DEFAULT NULL,
  `create_by` bigint(20) DEFAULT NULL,
  `update_by` bigint(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户和标签关系表';

