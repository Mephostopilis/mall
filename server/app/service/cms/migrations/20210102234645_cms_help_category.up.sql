CREATE TABLE `cms_help_category` (
  `id` bigint(20) unsigned NOT NULL COMMENT 'id',
  `app_id` bigint(20) unsigned NULL DEFAULT NULL ,
  `name` varchar(255) NOT NULL,
  `icon` varchar(255) DEFAULT NULL,
  `help_count` int(11) unsigned DEFAULT NULL,
  `show_status` int(11) DEFAULT NULL,
  `sort` int(11) NULL DEFAULT NULL,
  `create_by` bigint(20) DEFAULT NULL,
  `update_by` bigint(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;