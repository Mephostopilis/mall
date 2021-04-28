CREATE TABLE `sys_location` (
  `code` varchar(20) NOT NULL COMMENT '编号',
  `name` varchar(50) NOT NULL COMMENT '名称',
  `parent_code` varchar(20) DEFAULT NULL COMMENT '上级编号',
  `level` tinyint(3) NOT NULL COMMENT '层级',
  PRIMARY KEY (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='地区信息';