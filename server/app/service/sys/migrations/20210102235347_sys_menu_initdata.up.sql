INSERT INTO sys_menu VALUES 
(2,   1, 'Upms',     '系统管理', 'example', '/upms', '/0/2', 'M', '无', '', 0, 1, '', 'Layout', 1, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(3,   1, 'Sysuser',  '用户管理', 'user', 'sysuser',                '/0/2/3', 'C', '无', 'system:sysuser:list', 2, NULL, NULL, '/sys/user/index', 1, '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:10:42', NULL),
(50,  1, 'Res',      '资源管理', 'tree-table', 'resource',         '/0/2/50', 'C', '无', 'system:sysresource:list', 2, 1, '', '/sys/resource/index', 9, '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:16:30', NULL),
(51,  1, 'Menu',     '菜单管理', 'tree-table', 'menu',             '/0/2/51', 'C', '无', 'system:sysmenu:list', 2, 1, '', '/sys/menu/index', 3, '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:16:30', NULL),
(52,  1, 'Role',     '角色管理', 'peoples', 'role',                '/0/2/52', 'C', '无', 'system:sysrole:list', 2, 1, '', '/sys/role/index', 2, '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:16:19', NULL),
(56,  1, 'Dept',     '部门管理', 'tree', 'dept',                   '/0/2/56', 'C', '无', 'system:sysdept:list', 2, 0, '', '/sys/dept/index', 4, '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:16:47', NULL),
(57,  1, 'post',     '岗位管理', 'pass', 'post',                   '/0/2/57', 'C', '无', 'system:syspost:list', 2, 0, '', '/sys/post/index', 5, '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:16:53', NULL),
(58,  1, 'Dict',     '字典管理', 'education', 'dict',              '/0/2/58', 'C', '无', 'system:sysdicttype:list', 2, 0, '', '/sys/dict/index', 6, '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:17:04', NULL),
(59,  1, 'DictData', '字典数据', 'education', 'dict/data/:dictId', '/0/2/59', 'C', '无', 'system:sysdictdata:list', 2, 0, '', '/sys/dict/data', 100, '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:17:25', NULL),
(60,  1, 'Tools',    '系统工具', 'component', '/tools',            '/0/60', 'M', '无', '', 0, 0, '', 'Layout', 3, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(61,  1, 'Swagger',  '系统接口', 'guide', 'swagger',               '/0/60/61', 'C', '无', '', 60, 0, '', '/tools/swagger/index', 1, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(62,  1, 'Config',   '参数设置', 'list', '/config',                '/0/2/62', 'C', '无', 'system:sysconfig:list', 2, 0, '', '/sys/config/index', 7, '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:17:16', NULL),
(211, 1, 'Log',      '日志管理', 'log', '/log', '/0/2/211', 'M', '', '', 2, 0, '', '/sys/log/index', 8, '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:15:32', NULL),
(212, 1, 'LoginLog', '登录日志', 'logininfor', '/loginlog',        '/0/2/211/212', 'C', '', 'system:sysloginlog:list', 211, 0, '', '/sys/log/loginlog', 1, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(216, 1, 'OperLog',  '操作日志', 'skill', '/operlog',              '/0/2/211/216', 'C', '', 'system:sysoperlog:list', 211, 0, '', '/sys/log/operlog', 1, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(220, 1, '', '新增菜单', '', '', '/0/2/51/220', 'F', '', 'system:sysmenu:add', 51, 0, '', '', 1, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(221, 1, '', '修改菜单', 'edit', '', '/0/2/51/221', 'F', '', 'system:sysmenu:edit', 51, 0, '', '', 1, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(222, 1, '', '查询菜单', 'search', '', '/0/2/51/222', 'F', '', 'system:sysmenu:query', 51, 0, '', '', 1, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(223, 1, '', '删除菜单', '', '', '/0/2/51/223', 'F', '', 'system:sysmenu:remove', 51, 0, '', '', 1, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(224, 1, '', '新增角色', '', '', '/0/2/52/224', 'F', '', 'system:sysrole:add', 52, 0, '', '', 1, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(225, 1, '', '查询角色', '', '', '/0/2/52/225', 'F', '', 'system:sysrole:query', 52, 0, '', '', 1, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(226, 1, '', '修改角色', '', '', '/0/2/52/226', 'F', '', 'system:sysrole:edit', 52, 0, '', '', 1, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(227, 1, '', '删除角色', '', '', '/0/2/52/227', 'F', '', 'system:sysrole:remove', 52, 0, '', '', 1, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(228, 1, '', '查询部门', '', '', '/0/2/56/228', 'F', '', 'system:sysdept:query', 56, 0, '', '', 1, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(229, 1, '', '新增部门', '', '', '/0/2/56/229', 'F', '', 'system:sysdept:add', 56, 0, '', '', 1, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(230, 1, '', '修改部门', '', '', '/0/2/56/230', 'F', '', 'system:sysdept:edit', 56, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(231, 1, '', '删除部门', '', '', '/0/2/56/231', 'F', '', 'system:sysdept:remove', 56, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(232, 1, '', '查询岗位', '', '', '/0/2/57/232', 'F', '', 'system:syspost:query', 57, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(233, 1, '', '新增岗位', '', '', '/0/2/57/233', 'F', '', 'system:syspost:add', 57, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(234, 1, '', '修改岗位', '', '', '/0/2/57/234', 'F', '', 'system:syspost:edit', 57, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(235, 1, '', '删除岗位', '', '', '/0/2/57/235', 'F', '', 'system:syspost:remove', 57, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(236, 1, '', '字典查询', '', '', '/0/2/58/236', 'F', '', 'system:sysdicttype:query', 58, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(237, 1, '', '新增类型', '', '', '/0/2/58/237', 'F', '', 'system:sysdicttype:add', 58, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(238, 1, '', '修改类型', '', '', '/0/2/58/238', 'F', '', 'system:sysdicttype:edit', 58, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(239, 1, '', '删除类型', '', '', '/0/2/58/239', 'F', '', 'system:sysdicttype:remove', 58, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(240, 1, '', '查询数据', '', '', '/0/2/59/240', 'F', '', 'system:sysdictdata:query', 59, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(241, 1, '', '新增数据', '', '', '/0/2/59/241', 'F', '', 'system:sysdictdata:add', 59, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(242, 1, '', '修改数据', '', '', '/0/2/59/242', 'F', '', 'system:sysdictdata:edit', 59, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(243, 1, '', '删除数据', '', '', '/0/2/59/243', 'F', '', 'system:sysdictdata:remove', 59, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(244, 1, '', '查询参数', '', '', '/0/2/62/244', 'F', '', 'system:sysconfig:query', 62, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(245, 1, '', '新增参数', '', '', '/0/2/62/245', 'F', '', 'system:sysconfig:add', 62, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(246, 1, '', '修改参数', '', '', '/0/2/62/246', 'F', '', 'system:sysconfig:edit', 62, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(247, 1, '', '删除参数', '', '', '/0/2/62/247', 'F', '', 'system:sysconfig:remove', 62, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(248, 1, '', '查询登录日志', '', '', '/0/2/211/212/248', 'F', '', 'system:sysloginlog:query', 212, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(249, 1, '', '删除登录日志', '', '', '/0/2/211/212/249', 'F', '', 'system:sysloginlog:remove', 212, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(250, 1, '', '查询操作日志', '', '', '/0/2/211/216/250', 'F', '', 'system:sysoperlog:query', 216, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(251, 1, '', '删除操作日志', '', '', '/0/2/211/216/251', 'F', '', 'system:sysoperlog:remove', 216, 0, '', '', 0, '0', '1', '1', '0', '2020-04-11 15:52:48', NULL, NULL),
(261, 1, 'Gen', '代码生成', 'code', 'gen', '/0/60/261', 'C', '', '', 60, 0, '', '/tools/gen/index', 2, '0', '1', '1', '0', '2020-04-11 15:52:48', '2020-04-12 11:18:12', NULL),
(262, 1, 'EditTable', '代码生成修改', 'build', 'editTable', '/0/60/262', 'C', '', '', 60, 0, '', '/tools/gen/editTable', 100, '1', '1', '1', '0', '2020-04-11 15:52:48', '2020-05-03 20:38:36', NULL),
(264, 1, 'Build', '表单构建', 'build', 'build', '/0/60/264', 'C', '', '', 60, 0, '', '/tools/build/index', 1, '0', '1', '1', '1', '2020-04-11 15:52:48', '2020-07-18 13:51:47', NULL),
(269, 1, 'Server', '服务监控', 'druid', 'system/monitor', '/0/60/269', 'C', '', 'monitor:server:list', 60, 0, '', '/tools/system/monitor', 0, '0', '1', '1', '1', '2020-04-14 00:28:19', '2020-08-09 02:07:53', NULL),
(460, 1, 'sys_job管理', '定时任务', 'tool', 'sys_job', '/0/2/460', 'C', '无', 'sysjob:sysjob:list', 2, 0, '', '/sys/job/index', 10, '0', '1', '1', '0', '2020-08-03 09:17:37', '2020-08-04 15:18:32', NULL),
(461, 1, 'sys_job', '分页获取定时任务', 'pass', 'sys_job', '/0/2/460/461', 'F', '无', 'sysjob:sysjob:query', 460, 0, '', '', 0, '0', '1', '1', '0', '2020-08-03 09:17:37', '2020-08-03 09:17:37', NULL),
(462, 1, 'sys_job', '创建定时任务', 'pass', 'sys_job', '/0/2/460/462', 'F', '无', 'sysjob:sysjob:add', 460, 0, '', '', 0, '0', '1', '1', '0', '2020-08-03 09:17:37', '2020-08-03 09:17:37', NULL),
(463, 1, 'sys_job', '修改定时任务', 'pass', 'sys_job', '/0/2/460/463', 'F', '无', 'sysjob:sysjob:edit', 460, 0, '', '', 0, '0', '1', '1', '0', '2020-08-03 09:17:37', '2020-08-03 09:17:37', NULL),
(464, 1, 'sys_job', '删除定时任务', 'pass', 'sys_job', '/0/2/460/464', 'F', '无', 'sysjob:sysjob:remove', 460, 0, '', '', 0, '0', '1', '1', '0', '2020-08-03 09:17:37', '2020-08-03 09:17:37', NULL),
(471, 1, 'job_log', '日志', 'bug', 'job_log', '/0/459/471', 'C', '', '', 459, 0, '', '/sysjob/log', 0, '1', '1', '1', '1', '2020-08-05 21:24:46', '2020-08-06 00:02:20', NULL),
(473, 1, 'sysSetting', '系统配置', 'form', 'syssettings', '/0/60/473', 'C', '无', 'syssetting:syssetting:list', 60, 0, '', '/tools/system/settings', 0, '0', '1', '1', '0', '2020-08-09 01:05:22', '2020-08-09 02:17:10', NULL),
(474, 1, 'sys_setting', '分页获取SysSetting', 'pass', 'sys_setting', '/0/60/473/474', 'F', '无', 'syssetting:syssetting:query', 473, 0, '', '', 0, '0', '1', '1', '0', '2020-08-09 01:05:22', '2020-08-09 01:05:22', NULL),
(475, 1, 'sys_setting', '创建SysSetting', 'pass', 'sys_setting', '/0/60/473/475', 'F', '无', 'syssetting:syssetting:add', 473, 0, '', '', 0, '0', '1', '1', '0', '2020-08-09 01:05:22', '2020-08-09 01:05:22', NULL),
(476, 1, 'sys_setting', '修改SysSetting', 'pass', 'sys_setting', '/0/60/473/476', 'F', '无', 'syssetting:syssetting:edit', 473, 0, '', '', 0, '0', '1', '1', '0', '2020-08-09 01:05:22', '2020-08-09 01:05:22', NULL),
(477, 1, 'sys_setting', '删除SysSetting', 'pass', 'sys_setting', '/0/60/473/477', 'F', '无', 'syssetting:syssetting:remove', 473, 0, '', '', 0, '0', '1', '1', '0', '2020-08-09 01:05:22', '2020-08-09 01:05:22', NULL),
(485, 1, 'App',    '应用管理',    'code',    'app',            '/0/2/485', 'C', '', '', 2, 0, '', '/sys/sso/apps',                  11, '0', '1', '0', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(486, 1, 'Tiku',     '题库管理',   'example', '/tiku',    '/0/486',    'M', '', '', 0,   0, '', 'Layout', 5, '0', '1', '1', '1', '2020-10-22 20:52:49', '2020-10-22 20:53:10', NULL),
(487, 1, 'Choice',       '选择题',     'code',    'choice',         '/0/486/487', 'C', '', '', 486, 0, '', '/tiku/choices/index',  1, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(488, 1, 'ChoiceDetail', '选择题详情', 'code',    'choiceDetail',   '/0/486/488', 'C', '', '', 486, 0, '', '/tiku/choices/detail', 2, '1', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(489, 1, 'Exercise',     '练习题',     'code',    'exercise',       '/0/486/489', 'C', '', '', 486, 0, '', '/tiku/exercise/index', 3, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(490, 1, 'Cms',      '内容管理',   'example', '/cms',    '/0/490',     'M', '', '', 0,   0, '', 'Layout', 6, '0', '1', '1', '1', '2020-10-22 20:52:49', '2020-10-22 20:53:10', NULL),
(491, 1, 'Pms',      '商品管理',   'example', '/pms',    '/0/491',     'M', '', '', 0,   0, '', 'Layout', 7, '0', '1', '1', '1', '2020-10-22 20:52:49', '2020-10-22 20:53:10', NULL),
(492, 1, 'Oms',      '订单管理',   'example', '/oms',    '/0/492',     'M', '', '', 0,   0, '', 'Layout', 8, '0', '1', '1', '1', '2020-10-22 20:52:49', '2020-10-22 20:53:10', NULL),
(493, 1, 'Sms',      '营销管理',   'example', '/sms',    '/0/493',     'M', '', '', 0,   0, '', 'Layout', 9, '0', '1', '1', '1', '2020-10-22 20:52:49', '2020-10-22 20:53:10', NULL),
(494, 1, 'Member',   '会员管理',   'example', '/member', '/0/494',     'M', '', '', 0,   0, '', 'Layout', 10, '0', '1', '1', '1', '2020-10-22 20:52:49', '2020-10-22 20:53:10', NULL),
(495, 1, 'CmsHelp',           '帮助',     'code',    'help',           '/0/490/495', 'C', '', '', 490, 0, '', '/cms/cmshelp/index',           1, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(496, 1, 'CmsHelpCategory',   '帮助分类', 'code',    'helpcategory',   '/0/490/496', 'C', '', '', 490, 0, '', '/cms/cmshelpcategory/index',   2, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(497, 1, 'CmsSubject',        '主题',     'code',    'subject',        '/0/490/497', 'C', '', '', 490, 0, '', '/cms/cmssubject/index',        3, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(498, 1, 'CmsSubjectCategory','主题分类', 'code',    'subjectcategory','/0/490/498', 'C', '', '', 490, 0, '', '/cms/cmssubjectcategory/index',4, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(499, 1, 'CmsSubjectComment', '主题评论', 'code',    'subjectcomment', '/0/490/499', 'C', '', '', 490, 0, '', '/cms/cmssubjectcomment/index', 5, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(500, 1, 'PmsAlbum',          '专辑',     'code',    'album',          '/0/491/500', 'C', '', '', 491, 0, '', '/pms/pmsalbum/index',          1, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(501, 1, 'OmsCompanyAddress', '公司地址', 'code',    'companyaddress', '/0/492/501', 'C', '', '', 492, 0, '', '/oms/omscompanyaddress/index', 1, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(502, 1, 'SmsCoupon',         '优惠价',   'code',    'coupon',         '/0/493/502', 'C', '', '', 493, 0, '', '/sms/smscoupon/index',         1, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(503, 1, 'SmsCouponHistory',  '优惠卷历史','code',   'couponhistory',  '/0/493/503', 'C', '', '', 493, 0, '', '/sms/smscouponhistory/index',  2, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(504, 1, 'UmsMember',         '会员',     'code',    'list',           '/0/494/504', 'C', '', '', 494, 0, '', '/member/member/index',      1, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(505, 1, 'UmsMemberAssets',   '会员资产',     'code', 'memberassets',  '/0/494/505', 'C', '', '', 494, 0, '', '/member/memberassets/index', 2, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(506, 1, 'Chat',   '客服管理',   'example', '/chat', '/0/506',     'M', '', '', 0,   0, '', 'Layout', 11, '0', '1', '1', '1', '2020-10-22 20:52:49', '2020-10-22 20:53:10', NULL),
(507, 1, 'ChatSession',       '会话',     'code', 'session',  '/0/506/507', 'C', '', '', 506, 0, '', '/im/chat', 1, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(509, 1, 'PmsBrand',          '品牌',     'code',    'brand',           '/0/491/509', 'C', '', '', 491, 0, '', '/pms/pmsbrand/index',          3, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(510, 1, 'PmsComment',        '商品评论', 'code',    'comment',         '/0/491/510', 'C', '', '', 491, 0, '', '/pms/pmscomment/index',         4, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(511, 1, 'PmsCommentReplay',  '评论回复', 'code',    'commentreplay',   '/0/491/511', 'C', '', '', 491, 0, '', '/pms/pmscommentreplay/index',   5, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(512, 1, 'PmsProductAttributeCategory',  '特性分类', 'code',    'productattributecategory',   '/0/491/512', 'C', '', '', 491, 0, '', '/pms/pmsproductattributecategory/index',   6, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(513, 1, 'PmsProductCategory', '产品分类', 'code',    'productcategory', '/0/491/513', 'C', '', '', 491, 0, '', '/pms/pmsproductcategory/index',7, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(514, 1, 'PmsProduct',         '产品',    'code',    'product',         '/0/491/514',  'C', '', '', 491, 0, '', '/pms/pmsproduct/index',        8, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(515, 1, 'OmsOrder',           '订单',    'code',    'order',           '/0/492/515',  'C', '', '', 492, 0, '', '/oms/omsorder/index',          2, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL),
(516, 1, 'PmsProductAttribute','产品特性','code',    'productattribute','/0/491/516',  'C', '', '', 491, 0, '', '/pms/pmsproductattribute/index',9, '0', '1', '1', '0', '2020-10-22 20:54:51', '2020-10-22 20:55:40', NULL);
