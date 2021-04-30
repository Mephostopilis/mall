CREATE TABLE `cms_subject_comment` (
  `id` bigint(20) unsigned NOT NULL,
  `app_id` bigint(20) unsigned NULL DEFAULT NULL ,
  `subject_id` bigint(20) unsigned DEFAULT NULL,
  `member_nick_name` varchar(255) DEFAULT NULL,
  `member_icon` varchar(255) DEFAULT NULL,
  `content` varchar(1000) DEFAULT NULL,
  `show_status` int(1) DEFAULT NULL,
  `create_by` bigint(20) DEFAULT NULL,
  `update_by` bigint(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='专题评论表';