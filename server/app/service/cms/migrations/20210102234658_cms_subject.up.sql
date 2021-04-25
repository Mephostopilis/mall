CREATE TABLE `cms_subject` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `app_id` bigint(20) unsigned NULL DEFAULT NULL ,
  `category_id` int(11) DEFAULT NULL COMMENT '客户端id',
  `title` varchar(255) DEFAULT NULL COMMENT '标题',
  `pic` varchar(255) DEFAULT NULL COMMENT 'pic',
  `product_count` varchar(255) DEFAULT NULL COMMENT '产品数量',

  `create_by` bigint(20) DEFAULT NULL,
  `update_by` bigint(20) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;