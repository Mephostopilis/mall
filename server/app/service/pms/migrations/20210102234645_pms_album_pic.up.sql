CREATE TABLE `pms_album_pic` (
  `id` bigint(20) unsigned NOT NULL,
  `app_id` bigint(20) unsigned NULL DEFAULT NULL ,
  `album_id` bigint(20) unsigned DEFAULT NULL,
  `pic` varchar(1000) DEFAULT NULL,
  `create_by` bigint(20) DEFAULT NULL,
  `update_by` bigint(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='画册图片表';