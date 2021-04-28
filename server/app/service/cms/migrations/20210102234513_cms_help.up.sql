CREATE TABLE `cms_help` (
  `id` bigint(20) unsigned NOT NULL COMMENT 'id',
  `app_id` bigint(20) unsigned NULL DEFAULT NULL ,
  `category_id` bigint(20) DEFAULT NULL COMMENT '分类',
  `icon` varchar(255) DEFAULT NULL COMMENT 'icon',
  `title` varchar(255) DEFAULT NULL COMMENT 'title',
  `show_status` int(11) DEFAULT NULL COMMENT '显示状态',
  `read_count` int(11) DEFAULT NULL COMMENT '读取数',
  `content` text DEFAULT NULL COMMENT '内容',
  `create_by` bigint(20) DEFAULT NULL,
  `update_by` bigint(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;