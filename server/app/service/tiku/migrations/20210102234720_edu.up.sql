CREATE TABLE `tiku_choices` (
  `id` bigint(20) unsigned NOT NULL COMMENT 'id',
  `app_id` bigint(20) unsigned NULL DEFAULT NULL ,
  `title` text DEFAULT NULL COMMENT '标题',
  `pics` text DEFAULT NULL COMMENT '图片',
  `ty` int(11) NULL COMMENT '类型',
  `level` int(11) NULL COMMENT '等级',
  `answer` varchar(255) DEFAULT NULL COMMENT '答案',

  `create_by` bigint(20) DEFAULT NULL,
  `update_by` bigint(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;