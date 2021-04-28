CREATE TABLE `tiku_user_exercise` (
  `id` bigint(20) unsigned NOT NULL COMMENT 'id',
  `app_id` bigint(20) unsigned NULL DEFAULT NULL ,
  `tid` bigint(20) unsigned NOT NULL COMMENT '题目id',
  `uid` bigint(20) unsigned NOT NULL COMMENT 'uid',
  `answer` varchar(255) DEFAULT NULL COMMENT '答案',
  `remark` text NOT NULL COMMENT '自己标记',
  `approve` text NOT NULL COMMENT '批准',

  `create_by` bigint(20) DEFAULT NULL,
  `update_by` bigint(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;