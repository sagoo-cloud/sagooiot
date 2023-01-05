/*
 Navicat Premium Data Transfer

 Source Server         : Sagoo IOT
 Source Server Type    : MySQL
 Source Server Version : 50650 (5.6.50-log)
 Source Host           : 101.200.198.249:3306
 Source Schema         : sagoo-iot

 Target Server Type    : MySQL
 Target Server Version : 50650 (5.6.50-log)
 File Encoding         : 65001

 Date: 01/01/2023 22:52:58
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for alarm_level
-- ----------------------------
DROP TABLE IF EXISTS `alarm_level`;
CREATE TABLE `alarm_level` (
  `level` int(10) unsigned NOT NULL COMMENT '告警级别',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '名称',
  PRIMARY KEY (`level`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of alarm_level
-- ----------------------------
BEGIN;
INSERT INTO `alarm_level` (`level`, `name`) VALUES (1, '超紧急');
INSERT INTO `alarm_level` (`level`, `name`) VALUES (2, '紧急');
INSERT INTO `alarm_level` (`level`, `name`) VALUES (3, '严重');
INSERT INTO `alarm_level` (`level`, `name`) VALUES (4, '一般');
INSERT INTO `alarm_level` (`level`, `name`) VALUES (5, '提醒');
COMMIT;

-- ----------------------------
-- Table structure for alarm_log
-- ----------------------------
DROP TABLE IF EXISTS `alarm_log`;
CREATE TABLE `alarm_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `type` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '告警类型：1=规则告警，2=设备自主告警',
  `rule_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '规则id',
  `rule_name` varchar(255) NOT NULL DEFAULT '' COMMENT '规则名称',
  `level` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '告警级别',
  `data` text COMMENT '触发告警的数据',
  `product_key` varchar(255) NOT NULL DEFAULT '' COMMENT '产品标识',
  `device_key` varchar(255) NOT NULL DEFAULT '' COMMENT '设备标识',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '告警状态：0=未处理，1=已处理',
  `created_at` datetime DEFAULT NULL COMMENT '告警时间',
  `update_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '告警处理人员',
  `updated_at` datetime DEFAULT NULL COMMENT '处理时间',
  `content` varchar(500) NOT NULL COMMENT '处理意见',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=136208 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of alarm_log
-- ----------------------------
BEGIN;
INSERT INTO `alarm_log` (`id`, `type`, `rule_id`, `rule_name`, `level`, `data`, `product_key`, `device_key`, `status`, `created_at`, `update_by`, `updated_at`, `content`) VALUES (1, 1, 6, '南门电表', 2, '{\"ia\":138.87,\"ib\":112.03,\"ic\":119.54,\"pa\":17.17,\"pb\":52.86,\"pc\":21.04,\"ts\":\"2022-11-14 16:54:28\",\"va\":145.4,\"vab\":216.47,\"vb\":153.33,\"vbc\":201.05,\"vc\":162.82,\"vca\":120.46}', 'monipower20221103', 't20221333', 0, '2022-11-14 16:54:28', 0, '2022-11-14 16:54:28', '');
COMMIT;

-- ----------------------------
-- Table structure for alarm_rule
-- ----------------------------
DROP TABLE IF EXISTS `alarm_rule`;
CREATE TABLE `alarm_rule` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '告警规则名称',
  `level` int(10) unsigned NOT NULL DEFAULT '4' COMMENT '告警级别，默认：4（一般）',
  `product_key` varchar(255) NOT NULL DEFAULT '' COMMENT '产品标识',
  `device_key` varchar(255) NOT NULL DEFAULT '' COMMENT '设备标识',
  `trigger_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '触发类型：1=上线，2=离线，3=属性上报',
  `trigger_condition` text COMMENT '触发条件',
  `action` text COMMENT '执行动作',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态：0=未启用，1=已启用',
  `create_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新者',
  `deleted_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of alarm_rule
-- ----------------------------
BEGIN;
INSERT INTO `alarm_rule` (`id`, `name`, `level`, `product_key`, `device_key`, `trigger_type`, `trigger_condition`, `action`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, '测试', 1, '25', '15', 1, '{\"triggerCondition\":[{\"filters\":[{\"key\":\"sysTime\",\"operator\":\"eq\",\"value\":[\"1\"],\"andOr\":0},{\"key\":\"sysReportTime\",\"operator\":\"ne\",\"value\":[\"2\"],\"andOr\":1}],\"andOr\":0}]}', '{\"action\":[{\"sendGateway\":\"\",\"noticeConfig\":\"\",\"noticeTemplate\":\"\",\"addressee\":[]}]}', 0, 1, 0, 1, '2022-11-13 12:27:24', '2022-11-13 12:27:24', '2022-11-13 05:06:00');
INSERT INTO `alarm_rule` (`id`, `name`, `level`, `product_key`, `device_key`, `trigger_type`, `trigger_condition`, `action`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (2, '测试2', 3, 'monipower20221103', 't20221333', 3, '{\"triggerCondition\":[{\"filters\":[{\"key\":\"va\",\"operator\":\"gt\",\"value\":[\"200\"],\"andOr\":0}],\"andOr\":0}]}', '{\"action\":[{\"sendGateway\":\"mail\",\"noticeConfig\":\"tu0rkg0g4d0couk34jdudr7100b5f41f\",\"noticeTemplate\":\"tu0rkg0g4d0couk3cgkrqft200fm009z\",\"addressee\":[\"xinjy@qq.com\"]},{\"sendGateway\":\"sms\",\"noticeConfig\":\"tu0rkg0hq90cojuowfh931wi00vbhnck\",\"noticeTemplate\":\"tu0rkg018o0colupp3s7vyv900vk3s6n\",\"addressee\":[\"13700005102\"]}]}', 1, 1, 1, 0, '2022-11-13 13:00:05', '2022-12-09 14:46:41', NULL);
INSERT INTO `alarm_rule` (`id`, `name`, `level`, `product_key`, `device_key`, `trigger_type`, `trigger_condition`, `action`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (3, 'axtest', 2, '25', '15', 1, '{\"triggerCondition\":[{\"filters\":[{\"key\":\"sysReportTime\",\"operator\":\"gt\",\"value\":[\"2022-11-14:10\"],\"andOr\":0},{\"key\":\"\",\"operator\":\"\",\"value\":[],\"andOr\":0}],\"andOr\":0}]}', '{\"action\":[{\"sendGateway\":\"\",\"noticeConfig\":\"\",\"noticeTemplate\":\"\",\"addressee\":[]}]}', 0, 6, 0, 0, '2022-11-14 09:26:49', '2022-11-14 09:26:49', NULL);
INSERT INTO `alarm_rule` (`id`, `name`, `level`, `product_key`, `device_key`, `trigger_type`, `trigger_condition`, `action`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (4, 'test2', 4, '25', '14', 1, '{\"triggerCondition\":[{\"filters\":[{\"key\":\"sysReportTime\",\"operator\":\"eq\",\"value\":[\"1\"],\"andOr\":0}],\"andOr\":0}]}', '{\"action\":[{\"sendGateway\":\"\",\"noticeConfig\":\"\",\"noticeTemplate\":\"\",\"addressee\":[]}]}', 0, 1, 0, 0, '2022-11-14 11:20:40', '2022-11-14 20:24:47', NULL);
INSERT INTO `alarm_rule` (`id`, `name`, `level`, `product_key`, `device_key`, `trigger_type`, `trigger_condition`, `action`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 'test3', 4, 'monipower20221103', 't20221333', 3, '{\"triggerCondition\":[{\"filters\":[{\"key\":\"va\",\"operator\":\"lt\",\"value\":[\"120\"],\"andOr\":0}],\"andOr\":0}]}', '{\"action\":[{\"sendGateway\":\"sms\",\"noticeConfig\":\"tu0rkg0hq90cojuowfh931wi00vbhnck\",\"noticeTemplate\":\"tu0rkg018o0colupp3s7vyv900vk3s6n\",\"addressee\":[\"13700005102\"]}]}', 1, 1, 1, 0, '2022-11-14 11:21:44', '2022-12-07 07:54:00', NULL);
INSERT INTO `alarm_rule` (`id`, `name`, `level`, `product_key`, `device_key`, `trigger_type`, `trigger_condition`, `action`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, '南门电表', 2, 'monipower20221103', 't20221333', 3, '{\"triggerCondition\":[{\"filters\":[{\"key\":\"va\",\"operator\":\"lt\",\"value\":[\"120\"],\"andOr\":0}],\"andOr\":0}]}', '{\"action\":[{\"sendGateway\":\"mail\",\"noticeConfig\":\"tu0rkg0g4d0couk34jdudr7100b5f41f\",\"noticeTemplate\":\"tu0rkg0ijj0cozi9ct6u7jr500vs6shu\",\"addressee\":[\"xhpeng11@163.com\"]}]}', 1, 6, 6, 0, '2022-11-14 14:49:51', '2022-12-12 11:18:14', NULL);
INSERT INTO `alarm_rule` (`id`, `name`, `level`, `product_key`, `device_key`, `trigger_type`, `trigger_condition`, `action`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, '测试名称测试名称测试名称测试名称测试名称', 3, 'ww202211', 'w88991111', 1, '{\"triggerCondition\":[{\"filters\":[{\"key\":\"sysReportTime\",\"operator\":\"gt\",\"value\":[\"1\"],\"andOr\":0}],\"andOr\":0}]}', '{\"action\":[{\"sendGateway\":\"\",\"noticeConfig\":\"\",\"noticeTemplate\":\"\",\"addressee\":[]}]}', 0, 1, 0, 1, '2022-11-17 14:10:12', '2022-11-17 14:10:12', '2022-11-19 01:40:01');
INSERT INTO `alarm_rule` (`id`, `name`, `level`, `product_key`, `device_key`, `trigger_type`, `trigger_condition`, `action`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, '测试', 2, 'monipower20221103', 't20221333', 2, '{\"triggerCondition\":[{\"filters\":[{\"key\":\"va\",\"operator\":\"gt\",\"value\":[\"1\"],\"andOr\":0}],\"andOr\":0}]}', '{\"action\":[{\"sendGateway\":\"\",\"noticeConfig\":\"\",\"noticeTemplate\":\"\",\"addressee\":[]}]}', 0, 1, 0, 0, '2022-12-05 17:26:51', '2022-12-05 17:26:51', NULL);
INSERT INTO `alarm_rule` (`id`, `name`, `level`, `product_key`, `device_key`, `trigger_type`, `trigger_condition`, `action`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, '1221', 1, 'monipower20221103', 't20221333', 2, '{\"triggerCondition\":[{\"filters\":[{\"key\":\"va\",\"operator\":\"eq\",\"value\":[\"1\"],\"andOr\":0}],\"andOr\":0}]}', '{\"action\":[{\"sendGateway\":\"work_weixin\",\"noticeConfig\":\"tu0rkg0en80coq6dc5v5ppy1000ecyo1\",\"noticeTemplate\":\"\",\"addressee\":[\"11\"]},{\"sendGateway\":\"work_weixin\",\"noticeConfig\":\"tu0rkg0en80coq6dc5v5ppy1000ecyo1\",\"noticeTemplate\":\"\",\"addressee\":[\"111\"]}]}', 0, 1, 0, 0, '2022-12-05 17:33:47', '2022-12-05 17:33:47', NULL);
INSERT INTO `alarm_rule` (`id`, `name`, `level`, `product_key`, `device_key`, `trigger_type`, `trigger_condition`, `action`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, '测试', 1, 'monipower20221103', 't20221222', 2, '{\"triggerCondition\":[{\"filters\":[{\"key\":\"va\",\"operator\":\"gt\",\"value\":[\"180\"],\"andOr\":0}],\"andOr\":0}]}', '{\"action\":[{\"sendGateway\":\"sms\",\"noticeConfig\":\"tu0rkg0hq90cojuowfh931wi00vbhnck\",\"noticeTemplate\":\"tu0rkg018o0colupp3s7vyv900vk3s6n\",\"addressee\":[\"15163999200\"]}]}', 0, 1, 1, 1, '2022-12-05 17:36:41', '2022-12-07 07:46:07', '2022-12-08 12:40:10');
INSERT INTO `alarm_rule` (`id`, `name`, `level`, `product_key`, `device_key`, `trigger_type`, `trigger_condition`, `action`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (11, '2222', 1, 'ww202211', 'w88991111', 2, '{\"triggerCondition\":[{\"filters\":[{\"key\":\"sysReportTime\",\"operator\":\"eq\",\"value\":[\"1\"],\"andOr\":0}],\"andOr\":0}]}', '{\"action\":[{\"sendGateway\":\"work_weixin\",\"noticeConfig\":\"tu0rkg0en80coq6dc5v5ppy1000ecyo1\",\"noticeTemplate\":\"\",\"addressee\":[\"1111\"]},{\"sendGateway\":\"work_weixin\",\"noticeConfig\":\"tu0rkg0ari0coq6cnpy9w5z400j0k7dm\",\"noticeTemplate\":\"\",\"addressee\":[\"222\"]}]}', 0, 1, 0, 1, '2022-12-05 17:45:54', '2022-12-05 17:45:54', '2022-12-08 12:40:04');
INSERT INTO `alarm_rule` (`id`, `name`, `level`, `product_key`, `device_key`, `trigger_type`, `trigger_condition`, `action`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (12, '123123', 1, 'ww202211', 'w88991111', 2, '{\"triggerCondition\":[{\"filters\":[{\"key\":\"sysReportTime\",\"operator\":\"eq\",\"value\":[\"1\"],\"andOr\":0}],\"andOr\":0}]}', '{\"action\":[{\"sendGateway\":\"work_weixin\",\"noticeConfig\":\"tu0rkg0en80coq6dc5v5ppy1000ecyo1\",\"noticeTemplate\":\"\",\"addressee\":[\"11\"]},{\"sendGateway\":\"work_weixin\",\"noticeConfig\":\"tu0rkg0ari0coq6cnpy9w5z400j0k7dm\",\"noticeTemplate\":\"\",\"addressee\":[\"111\"]}]}', 1, 1, 0, 0, '2022-12-05 17:47:24', '2022-12-30 13:38:36', NULL);
INSERT INTO `alarm_rule` (`id`, `name`, `level`, `product_key`, `device_key`, `trigger_type`, `trigger_condition`, `action`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (13, '测试1', 1, 'ww202211', 'w88991111', 2, '{\"triggerCondition\":[{\"filters\":[{\"key\":\"sysReportTime\",\"operator\":\"eq\",\"value\":[\"11\"],\"andOr\":0}],\"andOr\":0}]}', '{\"action\":[{\"sendGateway\":\"work_weixin\",\"noticeConfig\":\"tu0rkg0en80coq6dc5v5ppy1000ecyo1\",\"noticeTemplate\":\"\",\"addressee\":[\"111\"]}]}', 0, 1, 0, 1, '2022-12-05 17:50:37', '2022-12-05 17:50:37', '2022-12-08 12:39:56');
INSERT INTO `alarm_rule` (`id`, `name`, `level`, `product_key`, `device_key`, `trigger_type`, `trigger_condition`, `action`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (14, '冲冲冲', 1, 'ww202211', 'w88991111', 2, '{\"triggerCondition\":[{\"filters\":[{\"key\":\"sysReportTime\",\"operator\":\"eq\",\"value\":[\"1\"],\"andOr\":0}],\"andOr\":0}]}', '{\"action\":[{\"sendGateway\":\"work_weixin\",\"noticeConfig\":\"tu0rkg0en80coq6dc5v5ppy1000ecyo1\",\"noticeTemplate\":\"1vfwi9pb2p0coqh3cc49od4400nvviwj\",\"addressee\":\"111\"},{\"sendGateway\":\"work_weixin\",\"noticeConfig\":\"tu0rkg0ari0coq6cnpy9w5z400j0k7dm\",\"noticeTemplate\":\"tu0rkg06g30coqhxytbh8ly500ybwlre\",\"addressee\":[\"222\"]}]}', 0, 1, 0, 1, '2022-12-05 17:53:24', '2022-12-05 17:53:24', '2022-12-07 04:45:12');
INSERT INTO `alarm_rule` (`id`, `name`, `level`, `product_key`, `device_key`, `trigger_type`, `trigger_condition`, `action`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (15, 'test11', 1, 'ww202211', 'w88991111', 2, '{\"triggerCondition\":[{\"filters\":[{\"key\":\"sysReportTime\",\"operator\":\"eq\",\"value\":[\"11\"],\"andOr\":0}],\"andOr\":0}]}', '{\"action\":[{\"sendGateway\":\"work_weixin\",\"noticeConfig\":\"tu0rkg0ari0coq6cnpy9w5z400j0k7dm\",\"noticeTemplate\":\"tu0rkg06g30coqhxytbh8ly500ybwlre\",\"addressee\":[\"2222\"]},{\"sendGateway\":\"mail\",\"noticeConfig\":\"tu0rkg0g4d0couk34jdudr7100b5f41f\",\"noticeTemplate\":\"tu0rkg0g4d0couk3cgkrqft200fm009z\",\"addressee\":[\"2222\"]}]}', 0, 1, 1, 1, '2022-12-07 12:50:57', '2022-12-07 13:06:20', '2022-12-08 12:39:50');
INSERT INTO `alarm_rule` (`id`, `name`, `level`, `product_key`, `device_key`, `trigger_type`, `trigger_condition`, `action`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (16, '测试多号码', 1, 'ww202211', 'w88991111', 1, '{\"triggerCondition\":[{\"filters\":[{\"key\":\"sysReportTime\",\"operator\":\"eq\",\"value\":[\"1\"],\"andOr\":0}],\"andOr\":0}]}', '{\"action\":[{\"sendGateway\":\"work_weixin\",\"noticeConfig\":\"tu0rkg0en80coq6dc5v5ppy1000ecyo1\",\"noticeTemplate\":\"1vfwi9pb2p0coqh3cc49od4400nvviwj\",\"addressee\":[\"1\",\"2\"]},{\"sendGateway\":\"mail\",\"noticeConfig\":\"tu0rkg0g4d0couk34jdudr7100b5f41f\",\"noticeTemplate\":\"tu0rkg0g4d0couk3cgkrqft200fm009z\",\"addressee\":[\"3\",\"4\",\"7\"]}]}', 0, 1, 1, 1, '2022-12-07 14:22:28', '2022-12-07 14:28:51', '2022-12-08 12:39:45');
COMMIT;

-- ----------------------------
-- Table structure for base_db_link
-- ----------------------------
DROP TABLE IF EXISTS `base_db_link`;
CREATE TABLE `base_db_link` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '名称',
  `types` varchar(20) NOT NULL COMMENT '驱动类型 mysql或oracle',
  `host` varchar(255) NOT NULL COMMENT '主机地址',
  `port` int(11) NOT NULL COMMENT '端口号',
  `user_name` varchar(30) NOT NULL COMMENT '用户名称',
  `password` varchar(50) NOT NULL COMMENT '密码',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `status` int(11) NOT NULL COMMENT '状态 0 停用 1启用',
  `is_deleted` int(11) NOT NULL COMMENT '是否删除 0未删除 1已删除',
  `created_by` int(10) unsigned DEFAULT '0' COMMENT '创建人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` int(11) DEFAULT NULL COMMENT '修改人',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_by` int(11) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='数据源连接';

-- ----------------------------
-- Records of base_db_link
-- ----------------------------
BEGIN;
INSERT INTO `base_db_link` (`id`, `name`, `types`, `host`, `port`, `user_name`, `password`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (1, 'TDEngine', 'taosRestful', '101.200.198.249', 6041, 'zhgy_iot', 'zh5365a', '时序数据库', 1, 0, 1, '2022-09-03 17:55:38', NULL, NULL, NULL, NULL);
INSERT INTO `base_db_link` (`id`, `name`, `types`, `host`, `port`, `user_name`, `password`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (2, 'hddata', 'sqlServer', '124.70.37.242', 1433, 'sa', 'hrjn2020ldy#', '微软sql-server', 1, 0, 0, '2022-09-21 20:24:49', NULL, NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `ptype` varchar(10) DEFAULT NULL,
  `v0` varchar(256) DEFAULT NULL,
  `v1` varchar(256) DEFAULT NULL,
  `v2` varchar(256) DEFAULT NULL,
  `v3` varchar(256) DEFAULT NULL,
  `v4` varchar(256) DEFAULT NULL,
  `v5` varchar(256) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for city_data
-- ----------------------------
DROP TABLE IF EXISTS `city_data`;
CREATE TABLE `city_data` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL COMMENT '名字',
  `code` varchar(50) NOT NULL COMMENT '编码',
  `parent_id` int(11) NOT NULL COMMENT '父ID',
  `sort` int(11) DEFAULT NULL COMMENT '排序',
  `status` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '状态;0:禁用;1:正常',
  `is_deleted` int(11) NOT NULL COMMENT '是否删除 0未删除 1已删除',
  `created_by` int(11) unsigned DEFAULT '0' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建日期',
  `updated_by` int(11) unsigned DEFAULT '0' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '修改日期',
  `deleted_by` int(11) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8 COMMENT='城市结构表';

-- ----------------------------
-- Records of city_data
-- ----------------------------
BEGIN;
INSERT INTO `city_data` (`id`, `name`, `code`, `parent_id`, `sort`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (1, '北京市', '110000', -1, NULL, 1, 1, 1, '2022-09-25 06:49:40', 1, '2022-10-17 16:44:57', 1, '2022-10-17 08:44:57');
INSERT INTO `city_data` (`id`, `name`, `code`, `parent_id`, `sort`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (2, '东城区', '110101', 1, NULL, 1, 1, 1, '2022-09-26 03:27:51', 0, '2022-10-08 19:22:16', 1, '2022-10-08 11:22:16');
INSERT INTO `city_data` (`id`, `name`, `code`, `parent_id`, `sort`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (3, '陕西省', '610000', -1, NULL, 1, 1, 1, '2022-09-25 20:53:05', 1, '2022-10-08 19:22:14', 1, '2022-10-08 11:22:14');
INSERT INTO `city_data` (`id`, `name`, `code`, `parent_id`, `sort`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (4, '汉中市', '610700', 3, NULL, 1, 1, 1, '2022-09-25 20:53:22', 1, '2022-10-08 19:22:12', 1, '2022-10-08 11:22:12');
INSERT INTO `city_data` (`id`, `name`, `code`, `parent_id`, `sort`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (5, '西安市', '610100', 3, NULL, 1, 1, 1, '2022-09-26 04:54:11', 0, '2022-10-08 19:22:06', 1, '2022-10-08 11:22:06');
INSERT INTO `city_data` (`id`, `name`, `code`, `parent_id`, `sort`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (6, '四川省', '510000', -1, NULL, 1, 1, 1, '2022-09-26 04:55:13', 0, '2022-10-08 19:22:04', 1, '2022-10-08 11:22:04');
INSERT INTO `city_data` (`id`, `name`, `code`, `parent_id`, `sort`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (7, '成都市', '510100', 6, NULL, 1, 1, 1, '2022-09-26 04:55:31', 0, '2022-10-08 19:22:02', 1, '2022-10-08 11:22:02');
INSERT INTO `city_data` (`id`, `name`, `code`, `parent_id`, `sort`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (8, '德阳市', '510600', 6, NULL, 1, 1, 1, '2022-09-26 04:55:54', 0, '2022-09-26 13:02:21', 1, '2022-09-26 05:02:21');
INSERT INTO `city_data` (`id`, `name`, `code`, `parent_id`, `sort`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (9, '绵阳市', '510700', 6, NULL, 1, 1, 1, '2022-09-26 04:56:19', 0, '2022-10-08 19:22:00', 1, '2022-10-08 11:22:00');
INSERT INTO `city_data` (`id`, `name`, `code`, `parent_id`, `sort`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (10, '乐山市', '511100', 6, NULL, 1, 1, 1, '2022-09-26 04:56:39', 0, '2022-10-08 19:21:57', 1, '2022-10-08 11:21:57');
INSERT INTO `city_data` (`id`, `name`, `code`, `parent_id`, `sort`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (11, '汉台区', '610702', 4, NULL, 1, 1, 1, '2022-09-26 04:57:14', 0, '2022-10-08 19:22:10', 1, '2022-10-08 11:22:10');
INSERT INTO `city_data` (`id`, `name`, `code`, `parent_id`, `sort`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (12, '南郑区', '610703', 4, NULL, 1, 1, 1, '2022-09-26 04:57:42', 0, '2022-10-08 19:22:08', 1, '2022-10-08 11:22:08');
INSERT INTO `city_data` (`id`, `name`, `code`, `parent_id`, `sort`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (13, '沈阳市', '210100', -1, 2, 1, 0, 1, '2022-10-08 03:22:42', 1, '2022-10-21 22:23:58', 0, NULL);
INSERT INTO `city_data` (`id`, `name`, `code`, `parent_id`, `sort`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (14, '丹东', '210600', -1, 1, 1, 0, 1, '2022-10-08 03:22:57', 1, '2022-10-21 22:24:05', 0, NULL);
INSERT INTO `city_data` (`id`, `name`, `code`, `parent_id`, `sort`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (15, '长春市', '220100', -1, NULL, 1, 0, 1, '2022-10-08 11:23:27', 1, '2022-10-08 19:52:12', 0, NULL);
INSERT INTO `city_data` (`id`, `name`, `code`, `parent_id`, `sort`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (16, '辽宁省', '211000', -1, NULL, 1, 0, 1, '2022-10-13 22:30:49', 1, '2022-10-14 06:34:26', 0, NULL);
INSERT INTO `city_data` (`id`, `name`, `code`, `parent_id`, `sort`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (17, '开原市', '211282', -1, NULL, 1, 0, 1, '2022-10-14 06:31:13', 0, '2022-10-14 06:31:13', 0, NULL);
INSERT INTO `city_data` (`id`, `name`, `code`, `parent_id`, `sort`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (18, '铁岭市', '211200', -1, NULL, 1, 0, 1, '2022-10-14 06:31:34', 0, '2022-10-14 06:31:34', 0, NULL);
INSERT INTO `city_data` (`id`, `name`, `code`, `parent_id`, `sort`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (19, '西安', '610000', -1, 4, 1, 1, 1, '2022-10-21 14:25:17', 0, '2022-11-03 21:57:35', 1, '2022-11-03 13:57:35');
COMMIT;

-- ----------------------------
-- Table structure for data_node
-- ----------------------------
DROP TABLE IF EXISTS `data_node`;
CREATE TABLE `data_node` (
  `node_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `source_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '数据源ID',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '数据节点名称',
  `key` varchar(255) NOT NULL DEFAULT '' COMMENT '数据节点标识',
  `data_type` varchar(255) NOT NULL DEFAULT '' COMMENT '数据类型',
  `value` varchar(255) NOT NULL DEFAULT '' COMMENT '取值项',
  `is_pk` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否主键：0=否，1=是',
  `rule` text COMMENT '规则配置json',
  `create_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新者',
  `deleted_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`node_id`)
) ENGINE=InnoDB AUTO_INCREMENT=389 DEFAULT CHARSET=utf8 COMMENT='数据节点';

-- ----------------------------
-- Records of data_node
-- ----------------------------
BEGIN;
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (314, 78, '区域编码', 'adcode', 'int', 'lives.0.adcode', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-10-21 00:48:59', '2022-10-21 00:48:59', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (315, 78, '天气现象', 'weather', 'string', 'lives.0.weather', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-10-21 00:49:21', '2022-10-21 00:49:21', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (316, 78, '实时气温，单位：摄氏度', 'temperature', 'string', 'lives.0.temperature', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-10-21 00:49:41', '2022-10-21 00:49:41', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (317, 78, '风向描述', 'winddirection', 'string', 'lives.0.winddirection', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-10-21 00:50:06', '2022-10-21 00:50:06', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (318, 78, '风力级别', 'windpower', 'string', 'lives.0.windpower', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-10-21 00:50:23', '2022-10-21 00:50:23', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (319, 78, '空气湿度', 'humidity', 'string', 'lives.0.humidity', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-10-21 00:50:40', '2022-10-21 00:50:40', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (320, 78, '数据发布时间', 'reporttime', 'date', 'lives.0.reporttime', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-10-21 00:50:57', '2022-10-21 00:50:57', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (321, 78, '省份名', 'province', 'string', 'lives.0.province', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-10-21 00:51:14', '2022-10-21 00:51:14', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (322, 78, '城市名', 'city', 'string', 'lives.0.city', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-10-21 00:51:31', '2022-10-21 00:51:31', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (323, 79, '日出时间', 'sunrise', 'string', 'result.daily.0.sunrise', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-10-21 01:07:03', '2022-10-21 01:07:03', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (324, 79, '日落时间', 'sunset', 'string', 'result.daily.0.sunset', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-10-21 01:07:20', '2022-10-21 01:07:20', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (325, 79, '风向', 'winddirect', 'string', 'result.daily.0.day.winddirect', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-10-21 01:07:42', '2022-10-21 01:07:42', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (326, 79, '风力等级', 'windpower', 'string', 'result.daily.0.day.windpower', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-10-21 01:08:04', '2022-10-21 01:08:04', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (352, 79, '预报下一日的温度', 'next_day_temp', 'string', 'result.daily.1.day.temphigh', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-11-09 19:53:13', '2022-11-09 19:53:13', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (353, 79, '预报下一日的风力等级', 'next_day_windpower', 'string', 'result.daily.1.day.windpower', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-11-09 19:53:58', '2022-11-09 19:53:58', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (354, 79, '预报第三日的温度', 'next_three_day_temp', 'string', 'result.daily.2.day.temphigh', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-11-09 19:55:03', '2022-11-09 19:55:03', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (355, 79, '预报第三日的风力等级', 'next_three_day_windpower', 'string', 'result.daily.2.day.windpower', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-11-09 19:55:46', '2022-11-09 19:55:46', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (356, 79, '预报第四日的温度', 'next_four_day_temp', 'string', 'result.daily.3.day.temphigh', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-11-09 19:56:24', '2022-11-09 19:56:24', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (357, 79, '预报第四日的风力等级', 'next_four_day_windpower', 'string', 'result.daily.3.day.windpower', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-11-09 19:57:09', '2022-11-09 19:57:09', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (358, 79, '预报第五日的温度', 'next_five_day_temp', 'string', 'result.daily.4.day.temphigh', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-11-09 19:58:11', '2022-11-09 19:58:11', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (359, 79, '预报第五日的风力等级', 'next_five_day_windpower', 'string', 'result.daily.4.day.windpower', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-11-09 19:58:47', '2022-11-09 19:58:47', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (360, 79, '预报第六日的温度', 'next_six_day_temp', 'string', 'result.daily.5.day.temphigh', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-11-09 19:59:40', '2022-11-09 19:59:40', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (361, 79, '预报第六日的风力等级', 'next_six_day_windpower', 'string', 'result.daily.5.day.windpower', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-11-09 20:00:13', '2022-11-09 20:00:13', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (362, 79, '预报第七日的温度', 'next_seven_day_temp', 'string', 'result.daily.6.day.temphigh', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-11-09 20:01:14', '2022-11-09 20:01:14', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (363, 79, '预报第七日的风力等级', 'next_seven_day_windpower', 'string', 'result.daily.6.day.windpower', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 0, 0, '2022-11-09 20:01:52', '2022-11-09 20:01:52', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (385, 79, '城市', 'city', 'int', 'result.city', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 6, 6, 6, '2022-11-29 11:40:14', '2022-11-29 11:40:56', '2022-11-29 03:41:47');
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (386, 79, '城市', 'city', 'string', 'result.city', 0, '[{\"expression\":\"\",\"replace\":\"\"}]', 6, 0, 0, '2022-11-29 11:42:04', '2022-11-29 11:42:04', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (387, 78, '关联字段', 'rcity', 'string', 'lives.0.city', 0, '[{\"expression\":\"市|省\",\"replace\":\"\"}]', 6, 0, 0, '2022-12-12 15:42:51', '2022-12-12 15:42:51', NULL);
INSERT INTO `data_node` (`node_id`, `source_id`, `name`, `key`, `data_type`, `value`, `is_pk`, `rule`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (388, 79, '关联字段', 'rcity', 'string', 'result.city', 0, '[{\"expression\":\"市|省\",\"replace\":\"\"}]', 6, 0, 0, '2022-12-12 15:44:16', '2022-12-12 15:44:16', NULL);
COMMIT;

-- ----------------------------
-- Table structure for data_source
-- ----------------------------
DROP TABLE IF EXISTS `data_source`;
CREATE TABLE `data_source` (
  `source_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '数据源名称',
  `key` varchar(255) NOT NULL DEFAULT '' COMMENT '数据源标识',
  `desc` varchar(500) NOT NULL DEFAULT '' COMMENT '描述',
  `from` tinyint(1) NOT NULL DEFAULT '1' COMMENT '数据来源：1=api导入，2=数据库，3=文件，4=设备',
  `config` text COMMENT '数据源配置json：api配置、数据库配置、文件配置',
  `rule` text COMMENT '规则配置json',
  `lock_key` tinyint(1) NOT NULL DEFAULT '1' COMMENT '锁定key标识：0=未锁定，1=锁定，不允许修改',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态：0=未发布，1=已发布',
  `data_table` varchar(255) NOT NULL DEFAULT '' COMMENT '数据表名称',
  `create_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新者',
  `deleted_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`source_id`)
) ENGINE=InnoDB AUTO_INCREMENT=85 DEFAULT CHARSET=utf8 COMMENT='数据源';

-- ----------------------------
-- Records of data_source
-- ----------------------------
BEGIN;
INSERT INTO `data_source` (`source_id`, `name`, `key`, `desc`, `from`, `config`, `rule`, `lock_key`, `status`, `data_table`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (78, '天气预报', 'weather_report', '天气预报', 1, '{\"method\":\"get\",\"url\":\"https://restapi.amap.com/v3/weather/weatherInfo\",\"requestParams\":[[{\"type\":\"param\",\"key\":\"key\",\"name\":\"key\",\"value\":\"adfcbd60fbd6315e350be34ee67fe8d0\"},{\"type\":\"param\",\"key\":\"city\",\"name\":\"城市编码\",\"value\":\"210100\"},{\"type\":\"param\",\"key\":\"extensions\",\"name\":\"气象类型\",\"value\":\"base\"}],[{\"type\":\"param\",\"key\":\"key\",\"name\":\"key\",\"value\":\"adfcbd60fbd6315e350be34ee67fe8d0\"},{\"type\":\"param\",\"key\":\"city\",\"name\":\"城市编码\",\"value\":\"210600\"},{\"type\":\"param\",\"key\":\"extensions\",\"name\":\"气象类型\",\"value\":\"base\"}],[{\"type\":\"param\",\"key\":\"key\",\"name\":\"key\",\"value\":\"adfcbd60fbd6315e350be34ee67fe8d0\"},{\"type\":\"param\",\"key\":\"city\",\"name\":\"城市编码\",\"value\":\"220100\"},{\"type\":\"param\",\"key\":\"extensions\",\"name\":\"气象类型\",\"value\":\"base\"}],[{\"type\":\"param\",\"key\":\"key\",\"name\":\"key\",\"value\":\"adfcbd60fbd6315e350be34ee67fe8d0\"},{\"type\":\"param\",\"key\":\"city\",\"name\":\"城市编码\",\"value\":\"211000\"},{\"type\":\"param\",\"key\":\"extensions\",\"name\":\"气象类型\",\"value\":\"base\"}],[{\"type\":\"param\",\"key\":\"key\",\"name\":\"key\",\"value\":\"adfcbd60fbd6315e350be34ee67fe8d0\"},{\"type\":\"param\",\"key\":\"city\",\"name\":\"城市编码\",\"value\":\"211282\"},{\"type\":\"param\",\"key\":\"extensions\",\"name\":\"气象类型\",\"value\":\"base\"}],[{\"type\":\"param\",\"key\":\"key\",\"name\":\"key\",\"value\":\"adfcbd60fbd6315e350be34ee67fe8d0\"},{\"type\":\"param\",\"key\":\"city\",\"name\":\"城市编码\",\"value\":\"211200\"},{\"type\":\"param\",\"key\":\"extensions\",\"name\":\"气象类型\",\"value\":\"base\"}]],\"cronExpression\":\"0 0 */1 * * ?\"}', '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 1, 'data_source_78', 1, 1, 0, '2022-10-21 00:46:49', '2022-12-12 15:42:58', NULL);
INSERT INTO `data_source` (`source_id`, `name`, `key`, `desc`, `from`, `config`, `rule`, `lock_key`, `status`, `data_table`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (79, '天气预报日出与日落', 'weather_Sunrise_And_Sunset', '天气预报日出与日落', 1, '{\"method\":\"get\",\"url\":\"https://api.jisuapi.com/weather/query\",\"requestParams\":[[{\"type\":\"param\",\"key\":\"appkey\",\"name\":\"appkey\",\"value\":\"d69906c51c4b812a\"},{\"type\":\"param\",\"key\":\"city\",\"name\":\"城市名\",\"value\":\"沈阳市\"}],[{\"type\":\"param\",\"key\":\"appkey\",\"name\":\"appkey\",\"value\":\"d69906c51c4b812a\"},{\"type\":\"param\",\"key\":\"city\",\"name\":\"城市名\",\"value\":\"丹东市\"}],[{\"type\":\"param\",\"key\":\"appkey\",\"name\":\"appkey\",\"value\":\"d69906c51c4b812a\"},{\"type\":\"param\",\"key\":\"city\",\"name\":\"城市名\",\"value\":\"长春市\"}],[{\"type\":\"param\",\"key\":\"appkey\",\"name\":\"appkey\",\"value\":\"d69906c51c4b812a\"},{\"type\":\"param\",\"key\":\"city\",\"name\":\"城市名\",\"value\":\"辽宁省\"}],[{\"type\":\"param\",\"key\":\"appkey\",\"name\":\"appkey\",\"value\":\"d69906c51c4b812a\"},{\"type\":\"param\",\"key\":\"city\",\"name\":\"城市名\",\"value\":\"开原市\"}],[{\"type\":\"param\",\"key\":\"appkey\",\"name\":\"appkey\",\"value\":\"d69906c51c4b812a\"},{\"type\":\"param\",\"key\":\"city\",\"name\":\"城市名\",\"value\":\"铁岭市\"}]],\"cronExpression\":\"0 0 */1 * * ?\"}', '[{\"expression\":\"\",\"replace\":\"\"}]', 1, 1, 'data_source_79', 1, 0, 0, '2022-10-21 01:06:10', '2022-12-12 15:44:18', NULL);
COMMIT;

-- ----------------------------
-- Table structure for data_source_78
-- ----------------------------
DROP TABLE IF EXISTS `data_source_78`;
CREATE TABLE `data_source_78` (
  `adcode` int(11) DEFAULT '0' COMMENT '区域编码',
  `weather` varchar(500) DEFAULT '' COMMENT '天气现象',
  `temperature` varchar(500) DEFAULT '' COMMENT '实时气温，单位：摄氏度',
  `winddirection` varchar(500) DEFAULT '' COMMENT '风向描述',
  `windpower` varchar(500) DEFAULT '' COMMENT '风力级别',
  `humidity` varchar(500) DEFAULT '' COMMENT '空气湿度',
  `reporttime` datetime DEFAULT NULL COMMENT '数据发布时间',
  `province` varchar(500) DEFAULT '' COMMENT '省份名',
  `city` varchar(500) DEFAULT '' COMMENT '城市名',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `rcity` varchar(255) DEFAULT '' COMMENT '关联字段'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for data_source_79
-- ----------------------------
DROP TABLE IF EXISTS `data_source_79`;
CREATE TABLE `data_source_79` (
  `sunrise` varchar(500) DEFAULT '' COMMENT '日出时间',
  `sunset` varchar(500) DEFAULT '' COMMENT '日落时间',
  `winddirect` varchar(500) DEFAULT '' COMMENT '风向',
  `windpower` varchar(500) DEFAULT '' COMMENT '风力等级',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `next_day_temp` varchar(255) DEFAULT '' COMMENT '预报下一日的温度',
  `next_day_windpower` varchar(255) DEFAULT '' COMMENT '预报下一日的风力等级',
  `next_three_day_temp` varchar(255) DEFAULT '' COMMENT '预报第三日的温度',
  `next_three_day_windpower` varchar(255) DEFAULT '' COMMENT '预报第三日的风力等级',
  `next_four_day_temp` varchar(255) DEFAULT '' COMMENT '预报第四日的温度',
  `next_four_day_windpower` varchar(255) DEFAULT '' COMMENT '预报第四日的风力等级',
  `next_five_day_temp` varchar(255) DEFAULT '' COMMENT '预报第五日的温度',
  `next_five_day_windpower` varchar(255) DEFAULT '' COMMENT '预报第五日的风力等级',
  `next_six_day_temp` varchar(255) DEFAULT '' COMMENT '预报第六日的温度',
  `next_six_day_windpower` varchar(255) DEFAULT '' COMMENT '预报第六日的风力等级',
  `next_seven_day_temp` varchar(255) DEFAULT '' COMMENT '预报第七日的温度',
  `next_seven_day_windpower` varchar(255) DEFAULT '' COMMENT '预报第七日的风力等级',
  `city` varchar(255) DEFAULT '' COMMENT '城市',
  `rcity` varchar(255) DEFAULT '' COMMENT '关联字段'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for data_template
-- ----------------------------
DROP TABLE IF EXISTS `data_template`;
CREATE TABLE `data_template` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '名称',
  `key` varchar(50) NOT NULL DEFAULT '' COMMENT '标识',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态：0=未发布，1=已发布',
  `cron_expression` varchar(255) NOT NULL DEFAULT '' COMMENT 'cron执行表达式',
  `sort_node_key` varchar(50) NOT NULL DEFAULT '' COMMENT '排序节点标识',
  `sort_desc` tinyint(1) NOT NULL DEFAULT '0' COMMENT '排序方式：1=倒序，2=正序',
  `data_table` varchar(255) NOT NULL DEFAULT '' COMMENT '数据表名称',
  `lock_key` tinyint(1) NOT NULL DEFAULT '1' COMMENT '锁定key标识：0=未锁定，1=锁定，不允许修改',
  `main_source_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '主数据源',
  `source_node_key` varchar(255) NOT NULL DEFAULT '' COMMENT '数据源关联节点',
  `create_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新者',
  `deleted_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8 COMMENT='数据建模';

-- ----------------------------
-- Records of data_template
-- ----------------------------
BEGIN;
INSERT INTO `data_template` (`id`, `name`, `key`, `desc`, `status`, `cron_expression`, `sort_node_key`, `sort_desc`, `data_table`, `lock_key`, `main_source_id`, `source_node_key`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (36, '天气预报数据建模', 'weather_temp_20221212154619', '', 0, '0 0 */1 * * ?', '', 0, 'data_template_36', 1, 78, 'rcity', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-13 10:31:08', NULL);
COMMIT;

-- ----------------------------
-- Table structure for data_template_36
-- ----------------------------
DROP TABLE IF EXISTS `data_template_36`;
CREATE TABLE `data_template_36` (
  `adcode` varchar(255) DEFAULT '' COMMENT '区域编码',
  `weather` varchar(255) DEFAULT '' COMMENT '天气现象',
  `temperature` varchar(255) DEFAULT '' COMMENT '实时气温，单位：摄氏度',
  `winddirection` varchar(255) DEFAULT '' COMMENT '风向描述',
  `humidity` varchar(255) DEFAULT '' COMMENT '空气湿度',
  `reporttime` datetime DEFAULT NULL COMMENT '数据发布时间',
  `sunrise` varchar(255) DEFAULT '' COMMENT '日出时间',
  `sunset` varchar(255) DEFAULT '' COMMENT '日落时间',
  `windpower` varchar(255) DEFAULT '' COMMENT '风力等级',
  `next_day_temp` varchar(255) DEFAULT '' COMMENT '预报下一日的温度',
  `next_day_windpower` varchar(255) DEFAULT '' COMMENT '预报下一日的风力等级',
  `next_three_day_temp` varchar(255) DEFAULT '' COMMENT '预报第三日的温度',
  `next_three_day_windpower` varchar(255) DEFAULT '' COMMENT '预报第三日的风力等级',
  `next_four_day_temp` varchar(255) DEFAULT '' COMMENT '预报第四日的温度',
  `next_four_day_windpower` varchar(255) DEFAULT '' COMMENT '预报第四日的风力等级',
  `next_five_day_temp` varchar(255) DEFAULT '' COMMENT '预报第五日的温度',
  `next_five_day_windpower` varchar(255) DEFAULT '' COMMENT '预报第五日的风力等级',
  `next_six_day_temp` varchar(255) DEFAULT '' COMMENT '预报第六日的温度',
  `next_six_day_windpower` varchar(255) DEFAULT '' COMMENT '预报第六日的风力等级',
  `next_seven_day_temp` varchar(255) DEFAULT '' COMMENT '预报第七日的温度',
  `next_seven_day_windpower` varchar(255) DEFAULT '' COMMENT '预报第七日的风力等级',
  `rcity` varchar(255) DEFAULT '' COMMENT '关联字段',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for data_template_busi
-- ----------------------------
DROP TABLE IF EXISTS `data_template_busi`;
CREATE TABLE `data_template_busi` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `data_template_id` int(11) NOT NULL COMMENT '数据建模ID',
  `busi_types` int(11) NOT NULL COMMENT '业务单元',
  `is_deleted` int(11) NOT NULL COMMENT '0未删除 1已删除',
  `created_by` int(11) unsigned DEFAULT '0' COMMENT '创建人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `deleted_by` int(11) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of data_template_busi
-- ----------------------------
BEGIN;
INSERT INTO `data_template_busi` (`id`, `data_template_id`, `busi_types`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (1, 36, 1, 0, 1, '2022-10-27 16:03:18', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for data_template_node
-- ----------------------------
DROP TABLE IF EXISTS `data_template_node`;
CREATE TABLE `data_template_node` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `tid` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '模型ID',
  `from` tinyint(2) NOT NULL COMMENT '字段生成方式:1=自动生成,2=数据源',
  `source_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '数据源ID',
  `node_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '数据源ID',
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '节点名称',
  `key` varchar(50) NOT NULL DEFAULT '' COMMENT '节点标识',
  `data_type` varchar(50) NOT NULL DEFAULT '' COMMENT '数据类型',
  `default` varchar(255) NOT NULL DEFAULT '' COMMENT '默认值',
  `method` enum('max','min','avg') DEFAULT NULL COMMENT '数值类型，取值方式',
  `is_pk` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否主键：0=否，1=是',
  `desc` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `create_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新者',
  `deleted_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=377 DEFAULT CHARSET=utf8 COMMENT='数据模型节点';

-- ----------------------------
-- Records of data_template_node
-- ----------------------------
BEGIN;
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (355, 36, 2, 78, 314, '区域编码', 'adcode', 'string', '', NULL, 0, '', 6, 6, 0, '2022-12-12 15:46:19', '2022-12-26 11:27:22', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (356, 36, 2, 78, 315, '天气现象', 'weather', 'string', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-12 15:46:19', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (357, 36, 2, 78, 316, '实时气温，单位：摄氏度', 'temperature', 'string', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-12 15:46:19', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (358, 36, 2, 78, 317, '风向描述', 'winddirection', 'string', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-12 15:46:19', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (359, 36, 2, 78, 319, '空气湿度', 'humidity', 'string', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-12 15:46:19', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (360, 36, 2, 78, 320, '数据发布时间', 'reporttime', 'date', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-12 15:46:19', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (361, 36, 2, 79, 323, '日出时间', 'sunrise', 'string', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-12 15:46:19', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (362, 36, 2, 79, 324, '日落时间', 'sunset', 'string', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-12 15:46:19', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (363, 36, 2, 79, 326, '风力等级', 'windpower', 'string', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-12 15:46:19', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (364, 36, 2, 79, 352, '预报下一日的温度', 'next_day_temp', 'string', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-12 15:46:19', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (365, 36, 2, 79, 353, '预报下一日的风力等级', 'next_day_windpower', 'string', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-12 15:46:19', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (366, 36, 2, 79, 354, '预报第三日的温度', 'next_three_day_temp', 'string', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-12 15:46:19', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (367, 36, 2, 79, 355, '预报第三日的风力等级', 'next_three_day_windpower', 'string', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-12 15:46:19', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (368, 36, 2, 79, 356, '预报第四日的温度', 'next_four_day_temp', 'string', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-12 15:46:19', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (369, 36, 2, 79, 357, '预报第四日的风力等级', 'next_four_day_windpower', 'string', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-12 15:46:19', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (370, 36, 2, 79, 358, '预报第五日的温度', 'next_five_day_temp', 'string', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-12 15:46:19', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (371, 36, 2, 79, 359, '预报第五日的风力等级', 'next_five_day_windpower', 'string', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-12 15:46:19', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (372, 36, 2, 79, 360, '预报第六日的温度', 'next_six_day_temp', 'string', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-12 15:46:19', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (373, 36, 2, 79, 361, '预报第六日的风力等级', 'next_six_day_windpower', 'string', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-12 15:46:19', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (374, 36, 2, 79, 362, '预报第七日的温度', 'next_seven_day_temp', 'string', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-12 15:46:19', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (375, 36, 2, 79, 363, '预报第七日的风力等级', 'next_seven_day_windpower', 'string', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:46:19', '2022-12-12 15:46:19', NULL);
INSERT INTO `data_template_node` (`id`, `tid`, `from`, `source_id`, `node_id`, `name`, `key`, `data_type`, `default`, `method`, `is_pk`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (376, 36, 2, 78, 387, '关联字段', 'rcity', 'string', '', NULL, 0, '', 6, 0, 0, '2022-12-12 15:47:46', '2022-12-12 15:47:46', NULL);
COMMIT;

-- ----------------------------
-- Table structure for dev_device
-- ----------------------------
DROP TABLE IF EXISTS `dev_device`;
CREATE TABLE `dev_device` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(255) NOT NULL DEFAULT '' COMMENT '设备标识',
  `name` varchar(255) DEFAULT NULL COMMENT '设备名称',
  `product_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '所属产品',
  `desc` varchar(255) DEFAULT NULL COMMENT '描述',
  `metadata_table` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否生成物模型子表：0=否，1=是',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态：0=未启用,1=离线,2=在线',
  `registry_time` datetime DEFAULT NULL COMMENT '激活时间',
  `last_online_time` datetime DEFAULT NULL COMMENT '最后上线时间',
  `certificate` varchar(255) NOT NULL DEFAULT '' COMMENT '设备证书',
  `secure_key` varchar(255) NOT NULL DEFAULT '' COMMENT '设备密钥',
  `version` varchar(255) NOT NULL DEFAULT '' COMMENT '固件版本号',
  `tunnel_id` int(11) NOT NULL DEFAULT '0' COMMENT 'tunnelId',
  `create_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新者',
  `deleted_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COMMENT='设备';

-- ----------------------------
-- Records of dev_device
-- ----------------------------
BEGIN;
INSERT INTO `dev_device` (`id`, `key`, `name`, `product_id`, `desc`, `metadata_table`, `status`, `registry_time`, `last_online_time`, `certificate`, `secure_key`, `version`, `tunnel_id`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (7, 'k213213', '设备名称', 18, '备注备注备注', 1, 1, NULL, NULL, '备证书', '设备秘钥1', '固件版本号', 0, 1, 1, 0, '2022-08-15 19:51:01', '2022-11-24 15:50:48', NULL);
INSERT INTO `dev_device` (`id`, `key`, `name`, `product_id`, `desc`, `metadata_table`, `status`, `registry_time`, `last_online_time`, `certificate`, `secure_key`, `version`, `tunnel_id`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, 'device1', '翱翔测试设备', 22, '', 1, 1, NULL, NULL, 'aaaa', 'bbbb', '1.1.1', 0, 1, 0, 0, '2022-09-16 18:39:49', '2022-11-24 15:50:47', NULL);
INSERT INTO `dev_device` (`id`, `key`, `name`, `product_id`, `desc`, `metadata_table`, `status`, `registry_time`, `last_online_time`, `certificate`, `secure_key`, `version`, `tunnel_id`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, 'w42134342', '角度传感设备', 23, '', 1, 1, NULL, NULL, '', '', '', 0, 1, 0, 0, '2022-09-17 15:37:55', '2022-11-24 15:50:44', NULL);
INSERT INTO `dev_device` (`id`, `key`, `name`, `product_id`, `desc`, `metadata_table`, `status`, `registry_time`, `last_online_time`, `certificate`, `secure_key`, `version`, `tunnel_id`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (11, 'testshebei', '设备测试', 22, '', 1, 1, NULL, NULL, '', '', '', 0, 1, 0, 0, '2022-09-20 09:24:02', '2022-11-24 15:50:42', NULL);
INSERT INTO `dev_device` (`id`, `key`, `name`, `product_id`, `desc`, `metadata_table`, `status`, `registry_time`, `last_online_time`, `certificate`, `secure_key`, `version`, `tunnel_id`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (12, 't20221103', '测试电表设备', 24, '', 1, 2, NULL, NULL, '', '', '', 0, 1, 0, 0, '2022-11-03 16:46:48', '2022-11-24 15:50:39', NULL);
INSERT INTO `dev_device` (`id`, `key`, `name`, `product_id`, `desc`, `metadata_table`, `status`, `registry_time`, `last_online_time`, `certificate`, `secure_key`, `version`, `tunnel_id`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (13, 'w88991111', '办公室环境监测设备', 25, '', 1, 2, NULL, NULL, '', '', '', 0, 1, 0, 0, '2022-11-04 09:14:01', '2022-11-24 15:50:37', NULL);
INSERT INTO `dev_device` (`id`, `key`, `name`, `product_id`, `desc`, `metadata_table`, `status`, `registry_time`, `last_online_time`, `certificate`, `secure_key`, `version`, `tunnel_id`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (14, 't20221222', '北门电表02', 24, '', 1, 2, NULL, NULL, '', '', '', 0, 0, 1, 0, '2022-11-04 15:49:17', '2022-11-24 15:50:35', NULL);
INSERT INTO `dev_device` (`id`, `key`, `name`, `product_id`, `desc`, `metadata_table`, `status`, `registry_time`, `last_online_time`, `certificate`, `secure_key`, `version`, `tunnel_id`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (15, 't20221333', '南门电表03', 24, '', 1, 2, NULL, NULL, '', '', '', 0, 0, 1, 0, '2022-11-04 15:49:48', '2022-11-24 15:50:31', NULL);
INSERT INTO `dev_device` (`id`, `key`, `name`, `product_id`, `desc`, `metadata_table`, `status`, `registry_time`, `last_online_time`, `certificate`, `secure_key`, `version`, `tunnel_id`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (16, 'tgn22021201', '室内测温插座001', 25, '', 1, 1, NULL, NULL, '', '', '', 0, 1, 1, 0, '2022-12-12 15:39:55', '2022-12-12 16:01:20', NULL);
COMMIT;

-- ----------------------------
-- Table structure for dev_device_tag
-- ----------------------------
DROP TABLE IF EXISTS `dev_device_tag`;
CREATE TABLE `dev_device_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `device_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '设备ID',
  `device_key` varchar(255) NOT NULL DEFAULT '' COMMENT '设备标识',
  `key` varchar(255) DEFAULT NULL COMMENT '标签标识',
  `name` varchar(255) DEFAULT NULL COMMENT '标签名称',
  `value` varchar(255) DEFAULT NULL COMMENT '标签值',
  `create_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新者',
  `deleted_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='设备标签';

-- ----------------------------
-- Records of dev_device_tag
-- ----------------------------
BEGIN;
INSERT INTO `dev_device_tag` (`id`, `device_id`, `device_key`, `key`, `name`, `value`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (1, 3, 'device-3', 'tag-1', '长度', '100', 0, 0, 0, '2022-08-11 16:52:28', '2022-08-11 17:08:32', NULL);
COMMIT;

-- ----------------------------
-- Table structure for dev_product
-- ----------------------------
DROP TABLE IF EXISTS `dev_product`;
CREATE TABLE `dev_product` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `key` varchar(255) NOT NULL DEFAULT '' COMMENT '产品标识',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '产品名称',
  `category_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '所属品类',
  `message_protocol` varchar(255) DEFAULT '' COMMENT '消息协议',
  `transport_protocol` varchar(255) DEFAULT '' COMMENT '传输协议: MQTT,COAP,UDP',
  `protocol_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '协议id',
  `device_type` varchar(255) DEFAULT '' COMMENT '设备类型: 网关，设备',
  `desc` varchar(255) DEFAULT NULL COMMENT '描述',
  `icon` varchar(1000) DEFAULT NULL COMMENT '图片地址',
  `metadata` text COMMENT '物模型',
  `metadata_table` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否生成物模型表：0=否，1=是',
  `policy` varchar(255) NOT NULL DEFAULT '' COMMENT '采集策略',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '发布状态：0=未发布，1=已发布',
  `create_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新者',
  `deleted_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COMMENT='产品表';

-- ----------------------------
-- Records of dev_product
-- ----------------------------
BEGIN;
INSERT INTO `dev_product` (`id`, `key`, `name`, `category_id`, `message_protocol`, `transport_protocol`, `protocol_id`, `device_type`, `desc`, `icon`, `metadata`, `metadata_table`, `policy`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (18, 'dianji', '电机', 22, 'modbus', 'tcp-server', 0, '网关', '水房电机', 'http://101.200.198.249:8899/upload_file/2022-08-23/cmd37oai1w9scjl2co.png', '{\"key\":\"dianji\",\"name\":\"电机\",\"properties\":[{\"key\":\"property_99\",\"name\":\"属性-9\",\"accessMode\":1,\"valueType\":{\"type\":\"int\",\"maxLength\":22},\"desc\":\"描述\"},{\"key\":\"property_98\",\"name\":\"属性-9\",\"accessMode\":1,\"valueType\":{\"type\":\"int\",\"maxLength\":22},\"desc\":\"描述\"},{\"key\":\"property_97\",\"name\":\"属性-9\",\"accessMode\":1,\"valueType\":{\"type\":\"int\",\"maxLength\":22},\"desc\":\"描述\"},{\"key\":\"property_96\",\"name\":\"属性-9\",\"accessMode\":1,\"valueType\":{\"type\":\"int\",\"maxLength\":22},\"desc\":\"描述\"},{\"key\":\"property_95\",\"name\":\"属性-9\",\"accessMode\":1,\"valueType\":{\"type\":\"int\",\"maxLength\":22},\"desc\":\"描述\"}],\"functions\":null,\"events\":null,\"tags\":[]}', 1, '', 1, 1, 1, 0, '2022-08-16 08:50:40', '2022-11-24 15:49:23', NULL);
INSERT INTO `dev_product` (`id`, `key`, `name`, `category_id`, `message_protocol`, `transport_protocol`, `protocol_id`, `device_type`, `desc`, `icon`, `metadata`, `metadata_table`, `policy`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (20, 'hhdhhsj', '测试', 22, 'modbus', 'tcp-client', 0, '设备', '', '', '{\"key\":\"hhdhhsj\",\"name\":\"测试\",\"properties\":[{\"key\":\"asaaaaa\",\"name\":\"测试e\",\"accessMode\":1,\"valueType\":{\"type\":\"int\",\"unit\":\"11\"},\"desc\":\"\"},{\"key\":\"csss\",\"name\":\"测试\",\"accessMode\":1,\"valueType\":{\"type\":\"int\",\"unit\":\"1\"},\"desc\":\"\"}],\"functions\":null,\"events\":null,\"tags\":null}', 0, '', 0, 1, 0, 0, '2022-09-14 09:12:02', '2022-09-15 16:38:45', NULL);
INSERT INTO `dev_product` (`id`, `key`, `name`, `category_id`, `message_protocol`, `transport_protocol`, `protocol_id`, `device_type`, `desc`, `icon`, `metadata`, `metadata_table`, `policy`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (22, 'aoxiang', '翱翔测试产品', 22, 'modbus', 'tcp-server', 0, '设备', '', '', '{\"key\":\"aoxiang\",\"name\":\"翱翔测试产品\",\"properties\":[{\"key\":\"pr1\",\"name\":\"属性1\",\"accessMode\":1,\"valueType\":{\"type\":\"int\",\"unit\":\"个\"},\"desc\":\"\"},{\"key\":\"pr2\",\"name\":\"属性2\",\"accessMode\":1,\"valueType\":{\"type\":\"string\",\"maxLength\":200},\"desc\":\"\"},{\"key\":\"cesgi\",\"name\":\"测试1\",\"accessMode\":1,\"valueType\":{\"type\":\"int\",\"unit\":\"1\"},\"desc\":\"\"}],\"functions\":null,\"events\":null,\"tags\":[{\"key\":\"author\",\"name\":\"作者\",\"accessMode\":1,\"valueType\":{\"type\":\"string\",\"maxLength\":50},\"desc\":\"\"}]}', 1, '', 1, 1, 0, 0, '2022-09-16 18:35:51', '2022-11-24 15:49:18', NULL);
INSERT INTO `dev_product` (`id`, `key`, `name`, `category_id`, `message_protocol`, `transport_protocol`, `protocol_id`, `device_type`, `desc`, `icon`, `metadata`, `metadata_table`, `policy`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (23, 'wt61np', '物联姿态传感器', 9, 'WT61NP', 'udp', 0, '设备', '', 'http://zhgy.sagoo.cn:8899/upload_file/2022-09-17/cmy7zn612zvc9r60cr.jpg', '{\"key\":\"wt61np\",\"name\":\"WT61NP物联姿态传感器\",\"properties\":[{\"key\":\"ax\",\"name\":\"X轴\",\"accessMode\":1,\"valueType\":{\"type\":\"float\",\"decimals\":5,\"unit\":\"度\"},\"desc\":\"滚转角\"},{\"key\":\"ay\",\"name\":\"Y轴\",\"accessMode\":1,\"valueType\":{\"type\":\"float\",\"decimals\":5,\"unit\":\"度\"},\"desc\":\"俯仰角\"},{\"key\":\"az\",\"name\":\"Z轴\",\"accessMode\":1,\"valueType\":{\"type\":\"float\",\"decimals\":5,\"unit\":\"度\"},\"desc\":\"偏航角\"},{\"key\":\"t1\",\"name\":\"温度\",\"accessMode\":1,\"valueType\":{\"type\":\"float\",\"decimals\":2,\"unit\":\"度\"},\"desc\":\"温度\"}],\"functions\":null,\"events\":null,\"tags\":null}', 1, '', 1, 1, 1, 0, '2022-09-17 07:39:26', '2022-11-24 15:49:13', NULL);
INSERT INTO `dev_product` (`id`, `key`, `name`, `category_id`, `message_protocol`, `transport_protocol`, `protocol_id`, `device_type`, `desc`, `icon`, `metadata`, `metadata_table`, `policy`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (24, 'monipower20221103', '模拟测试电表2022', 9, '', 'mqtt_server', 0, '设备', '', '', '{\"key\":\"monipower20221103\",\"name\":\"模拟测试电表2022\",\"properties\":[{\"key\":\"va\",\"name\":\"A相电压\",\"accessMode\":1,\"valueType\":{\"type\":\"float\",\"decimals\":2,\"unit\":\"V\"},\"desc\":\"\"},{\"key\":\"vb\",\"name\":\"B相电压\",\"accessMode\":1,\"valueType\":{\"type\":\"float\",\"decimals\":2,\"unit\":\"V\"},\"desc\":\"\"},{\"key\":\"vc\",\"name\":\"C相电压\",\"accessMode\":1,\"valueType\":{\"type\":\"float\",\"decimals\":2,\"unit\":\"V\"},\"desc\":\"\"},{\"key\":\"ia\",\"name\":\"A相电流\",\"accessMode\":1,\"valueType\":{\"type\":\"float\",\"decimals\":2,\"unit\":\"A\"},\"desc\":\"\"},{\"key\":\"ib\",\"name\":\"A相电流\",\"accessMode\":1,\"valueType\":{\"type\":\"float\",\"decimals\":2,\"unit\":\"A\"},\"desc\":\"\"},{\"key\":\"ic\",\"name\":\"C相电流\",\"accessMode\":1,\"valueType\":{\"type\":\"float\",\"decimals\":2,\"unit\":\"A\"},\"desc\":\"\"},{\"key\":\"vab\",\"name\":\"AB相电压\",\"accessMode\":1,\"valueType\":{\"type\":\"float\",\"decimals\":2,\"unit\":\"V\"},\"desc\":\"\"},{\"key\":\"vbc\",\"name\":\"BC线电压\",\"accessMode\":1,\"valueType\":{\"type\":\"float\",\"decimals\":2,\"unit\":\"V\"},\"desc\":\"\"},{\"key\":\"vca\",\"name\":\"CA线电压\",\"accessMode\":1,\"valueType\":{\"type\":\"float\",\"decimals\":2,\"unit\":\"V\"},\"desc\":\"\"},{\"key\":\"pa\",\"name\":\"A相有功功率\\t\",\"accessMode\":1,\"valueType\":{\"type\":\"float\",\"decimals\":2,\"unit\":\"kW\"},\"desc\":\"\"},{\"key\":\"pb\",\"name\":\"B相有功功率\",\"accessMode\":1,\"valueType\":{\"type\":\"float\",\"decimals\":2,\"unit\":\"kW\"},\"desc\":\"\"},{\"key\":\"pc\",\"name\":\"C相有功功率\",\"accessMode\":1,\"valueType\":{\"type\":\"float\",\"decimals\":2,\"unit\":\"kW\"},\"desc\":\"\"}],\"functions\":null,\"events\":null,\"tags\":null}', 1, '', 1, 1, 0, 0, '2022-11-03 16:45:50', '2022-11-24 15:49:07', NULL);
INSERT INTO `dev_product` (`id`, `key`, `name`, `category_id`, `message_protocol`, `transport_protocol`, `protocol_id`, `device_type`, `desc`, `icon`, `metadata`, `metadata_table`, `policy`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (25, 'ww202211', '室内环境监测设备01', 9, '', 'mqtt_server', 0, '设备', '', '', '{\"key\":\"ww202211\",\"name\":\"室内环境监测设备01\",\"properties\":[{\"key\":\"t\",\"name\":\"温度\",\"accessMode\":1,\"valueType\":{\"type\":\"int\",\"unit\":\"度\"},\"desc\":\"\"},{\"key\":\"h\",\"name\":\"湿度\",\"accessMode\":1,\"valueType\":{\"type\":\"float\",\"decimals\":2,\"unit\":\"湿度\"},\"desc\":\"\"},{\"key\":\"pa\",\"name\":\"气压\",\"accessMode\":1,\"valueType\":{\"type\":\"int\"},\"desc\":\"\"}],\"functions\":null,\"events\":null,\"tags\":null}', 1, '', 1, 1, 0, 0, '2022-11-04 09:11:44', '2022-11-24 15:49:00', NULL);
INSERT INTO `dev_product` (`id`, `key`, `name`, `category_id`, `message_protocol`, `transport_protocol`, `protocol_id`, `device_type`, `desc`, `icon`, `metadata`, `metadata_table`, `policy`, `status`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (26, 'tgnsncw', '室内测温插座', 22, 'tgn52', 'tcp', 0, '设备', '', '', '{\"key\":\"tgnsncw\",\"name\":\"室内测温插座\",\"properties\":[{\"key\":\"signal\",\"name\":\"信号质量\",\"accessMode\":1,\"valueType\":{\"type\":\"int\"},\"desc\":\"1:低；2:中；3:高\"},{\"key\":\"battery\",\"name\":\"电池质量\",\"accessMode\":1,\"valueType\":{\"type\":\"int\"},\"desc\":\"1:低；2:中；3:高\"},{\"key\":\"temperature\",\"name\":\"温度\",\"accessMode\":1,\"valueType\":{\"type\":\"float\",\"decimals\":1,\"unit\":\"℃\"},\"desc\":\"\"},{\"key\":\"humidity\",\"name\":\"湿度\",\"accessMode\":1,\"valueType\":{\"type\":\"float\",\"decimals\":1,\"unit\":\"%rh\"},\"desc\":\"\"}],\"functions\":null,\"events\":null,\"tags\":null}', 1, '', 1, 1, 1, 0, '2022-12-12 15:38:51', '2022-12-12 16:01:02', NULL);
COMMIT;

-- ----------------------------
-- Table structure for dev_product_category
-- ----------------------------
DROP TABLE IF EXISTS `dev_product_category`;
CREATE TABLE `dev_product_category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '父ID',
  `key` varchar(20) NOT NULL DEFAULT '' COMMENT '分类标识',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '分类名称',
  `desc` varchar(255) DEFAULT NULL COMMENT '描述',
  `create_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建者',
  `update_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '更新者',
  `deleted_by` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8mb4 COMMENT='产品分类';

-- ----------------------------
-- Records of dev_product_category
-- ----------------------------
BEGIN;
INSERT INTO `dev_product_category` (`id`, `parent_id`, `key`, `name`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (8, 0, 'key-1', '分类-1', '', 0, 0, 0, '2022-08-04 15:58:48', '2022-08-04 15:58:48', NULL);
INSERT INTO `dev_product_category` (`id`, `parent_id`, `key`, `name`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (9, 0, 'key-12', '分类-12', '12', 0, 1, 0, '2022-08-04 15:59:00', '2022-11-08 23:46:34', NULL);
INSERT INTO `dev_product_category` (`id`, `parent_id`, `key`, `name`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (10, 8, 'key-2', '分类-2', '', 0, 0, 0, '2022-08-04 15:59:28', '2022-08-04 15:59:28', NULL);
INSERT INTO `dev_product_category` (`id`, `parent_id`, `key`, `name`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (11, 8, 'key-3', '分类-3', '', 0, 0, 0, '2022-08-04 15:59:38', '2022-08-10 18:25:48', '2022-08-10 10:30:43');
INSERT INTO `dev_product_category` (`id`, `parent_id`, `key`, `name`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (12, 10, 'key-3', '分类-3', '', 0, 0, 0, '2022-08-04 15:59:49', '2022-08-04 15:59:49', NULL);
INSERT INTO `dev_product_category` (`id`, `parent_id`, `key`, `name`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (21, 10, 'string', 'string', 'string', 0, 0, 0, '2022-08-11 11:18:12', '2022-08-11 11:23:50', '2022-08-11 03:24:53');
INSERT INTO `dev_product_category` (`id`, `parent_id`, `key`, `name`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (22, 9, '标识', '测试', '11', 1, 0, 0, '2022-08-13 14:16:19', '2022-08-13 14:16:19', NULL);
INSERT INTO `dev_product_category` (`id`, `parent_id`, `key`, `name`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (23, 4294967295, '1111', '111', '', 1, 0, 0, '2022-08-13 14:25:32', '2022-08-13 14:25:32', NULL);
INSERT INTO `dev_product_category` (`id`, `parent_id`, `key`, `name`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (24, 4294967295, '11111', '1111', '', 1, 0, 0, '2022-08-13 14:25:59', '2022-08-13 14:25:59', NULL);
INSERT INTO `dev_product_category` (`id`, `parent_id`, `key`, `name`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (25, 0, '1111', '111122', '', 1, 1, 0, '2022-08-13 14:26:55', '2022-08-13 14:27:57', '2022-08-13 07:19:41');
INSERT INTO `dev_product_category` (`id`, `parent_id`, `key`, `name`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (26, 0, '22', '11122111', '11', 1, 0, 0, '2022-08-13 14:28:06', '2022-08-13 14:28:06', '2022-08-13 07:19:37');
INSERT INTO `dev_product_category` (`id`, `parent_id`, `key`, `name`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (27, 0, '333', '333', '333', 1, 0, 0, '2022-08-13 14:28:42', '2022-08-13 15:19:31', '2022-08-13 07:19:35');
INSERT INTO `dev_product_category` (`id`, `parent_id`, `key`, `name`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (28, 0, '44', '444', '', 1, 0, 1, '2022-08-13 14:29:59', '2022-08-13 14:29:59', '2022-08-13 07:03:28');
INSERT INTO `dev_product_category` (`id`, `parent_id`, `key`, `name`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (29, 0, '55', '55', '', 1, 0, 1, '2022-08-13 14:30:34', '2022-08-13 14:30:34', '2022-08-13 07:03:25');
INSERT INTO `dev_product_category` (`id`, `parent_id`, `key`, `name`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (30, 0, '111', '111', '', 1, 0, 1, '2022-08-15 18:13:56', '2022-08-15 18:13:56', '2022-08-15 10:14:00');
INSERT INTO `dev_product_category` (`id`, `parent_id`, `key`, `name`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (31, 22, 'test', 'test', '', 1, 0, 0, '2022-09-16 18:31:33', '2022-09-16 18:31:33', NULL);
INSERT INTO `dev_product_category` (`id`, `parent_id`, `key`, `name`, `desc`, `create_by`, `update_by`, `deleted_by`, `created_at`, `updated_at`, `deleted_at`) VALUES (32, 8, 'sdss', '测试', '', 1, 0, 0, '2022-09-16 18:34:26', '2022-09-16 18:34:26', NULL);
COMMIT;

-- ----------------------------
-- Table structure for guestbook
-- ----------------------------
DROP TABLE IF EXISTS `guestbook`;
CREATE TABLE `guestbook` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(50) NOT NULL COMMENT '留言标题',
  `content` varchar(200) NOT NULL COMMENT '留言内容',
  `contacts` varchar(50) NOT NULL COMMENT '联系人',
  `telephone` varchar(50) NOT NULL COMMENT '联系方式',
  `created_at` datetime NOT NULL COMMENT '留言时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='访客留言表';

-- ----------------------------
-- Records of guestbook
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for network_server
-- ----------------------------
DROP TABLE IF EXISTS `network_server`;
CREATE TABLE `network_server` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `types` varchar(50) NOT NULL COMMENT 'tcp/udp',
  `addr` varchar(50) NOT NULL,
  `register` varchar(200) NOT NULL COMMENT '注册包',
  `heartbeat` varchar(200) NOT NULL COMMENT '心跳包',
  `protocol` varchar(200) NOT NULL COMMENT '协议',
  `devices` varchar(200) NOT NULL COMMENT '默认设备',
  `status` tinyint(1) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `create_by` int(11) NOT NULL,
  `remark` varchar(200) NOT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8 COMMENT='网络组件服务表';

-- ----------------------------
-- Records of network_server
-- ----------------------------
BEGIN;
INSERT INTO `network_server` (`id`, `name`, `types`, `addr`, `register`, `heartbeat`, `protocol`, `devices`, `status`, `created_at`, `updated_at`, `create_by`, `remark`) VALUES (22, '新建服务器', 'tcp', '9000', '{\"regex\":\"^w+$\"}', '{\"enable\":false,\"hex\":\"\",\"regex\":\"^\\\\w+$\",\"text\":\"\",\"timeout\":30}', '{\"name\":\"ModbusTCP\",\"options\":{}}', '[]', 1, '2022-10-23 15:46:22', '2022-12-24 14:13:01', 0, '');
INSERT INTO `network_server` (`id`, `name`, `types`, `addr`, `register`, `heartbeat`, `protocol`, `devices`, `status`, `created_at`, `updated_at`, `create_by`, `remark`) VALUES (23, '室内测温服务', 'tcp', '5001', '{\"regex\":\"^w+$\"}', '{\"enable\":false,\"hex\":\"\",\"regex\":\"^\\\\w+$\",\"text\":\"\",\"timeout\":30}', '{\"name\":\"tgn52\",\"options\":{}}', '[]', 1, '2022-12-12 16:37:37', '2022-12-12 16:37:48', 0, '');
COMMIT;

-- ----------------------------
-- Table structure for network_tunnel
-- ----------------------------
DROP TABLE IF EXISTS `network_tunnel`;
CREATE TABLE `network_tunnel` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `server_id` int(11) NOT NULL DEFAULT '0' COMMENT '服务ID',
  `name` varchar(50) NOT NULL,
  `types` varchar(50) NOT NULL,
  `addr` varchar(50) NOT NULL,
  `remote` varchar(50) NOT NULL,
  `retry` varchar(200) NOT NULL COMMENT '断线重连',
  `heartbeat` varchar(200) NOT NULL COMMENT '心跳包',
  `serial` varchar(200) NOT NULL COMMENT '串口参数',
  `protoccol` varchar(200) NOT NULL COMMENT '适配协议',
  `device_key` varchar(255) NOT NULL DEFAULT '' COMMENT '设备标识',
  `status` tinyint(1) NOT NULL,
  `last` datetime DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `remark` varchar(200) NOT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8 COMMENT='网络通道表';

-- ----------------------------
-- Records of network_tunnel
-- ----------------------------
BEGIN;
INSERT INTO `network_tunnel` (`id`, `server_id`, `name`, `types`, `addr`, `remote`, `retry`, `heartbeat`, `serial`, `protoccol`, `device_key`, `status`, `last`, `created_at`, `updated_at`, `remark`) VALUES (19, 0, '新建通道', 'tcp-client', '192.168.10.5:49621', '47.56.247.127:80', '{\"enable\":true,\"maximum\":0,\"timeout\":30}', '{\"enable\":true,\"hex\":\"\",\"regex\":\"^\\\\w+$\",\"text\":\"\",\"timeout\":30}', '{\"baud_rate\":\"9600\",\"data_bits\":\"6\",\"parity\":\"0\",\"rs485\":false,\"stop_bits\":\"1\"}', '{\"name\":\"ModbusTCP\",\"options\":{}}', '', 0, '2022-10-23 07:39:45', '2022-10-23 14:47:17', '2022-10-23 15:39:45', '');
INSERT INTO `network_tunnel` (`id`, `server_id`, `name`, `types`, `addr`, `remote`, `retry`, `heartbeat`, `serial`, `protoccol`, `device_key`, `status`, `last`, `created_at`, `updated_at`, `remark`) VALUES (20, 0, '新建通道', 'udp-client', '127.0.0.1:53351', '', '{\"enable\":true,\"maximum\":0,\"timeout\":30}', '{\"enable\":true,\"hex\":\"\",\"regex\":\"^\\\\w+$\",\"text\":\"\",\"timeout\":30}', '{\"baud_rate\":\"9600\",\"data_bits\":\"6\",\"parity\":\"0\",\"rs485\":false,\"stop_bits\":\"1\"}', '{\"name\":\"ModbusTCP\",\"options\":{}}', '', 0, '2022-10-23 07:41:03', '2022-10-23 15:36:09', '2022-10-24 11:26:07', '');
INSERT INTO `network_tunnel` (`id`, `server_id`, `name`, `types`, `addr`, `remote`, `retry`, `heartbeat`, `serial`, `protoccol`, `device_key`, `status`, `last`, `created_at`, `updated_at`, `remark`) VALUES (21, 0, '新建通道', 'tcp-client', '127.0.0.1:9090', '', '{\"enable\":true,\"maximum\":0,\"timeout\":30}', '{\"enable\":false,\"hex\":\"\",\"regex\":\"^\\\\w+$\",\"text\":\"\",\"timeout\":30}', '{\"baud_rate\":\"9600\",\"data_bits\":\"6\",\"parity\":\"0\",\"rs485\":false,\"stop_bits\":\"1\"}', '{\"name\":\"ModbusTCP\",\"options\":{}}', '', 0, NULL, '2022-10-26 00:13:14', '2022-10-26 11:03:02', '');
COMMIT;

-- ----------------------------
-- Table structure for notice_config
-- ----------------------------
DROP TABLE IF EXISTS `notice_config`;
CREATE TABLE `notice_config` (
  `id` varchar(32) NOT NULL,
  `title` varchar(100) NOT NULL,
  `send_gateway` varchar(200) NOT NULL,
  `types` tinyint(1) NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='通知服务配置表';

-- ----------------------------
-- Records of notice_config
-- ----------------------------
BEGIN;
INSERT INTO `notice_config` (`id`, `title`, `send_gateway`, `types`, `created_at`) VALUES ('tu0rkg03iz0coqyeyhni297200eh7xkf', 'ereqrqe', 'dingding', 1, '2022-12-02 09:51:37');
INSERT INTO `notice_config` (`id`, `title`, `send_gateway`, `types`, `created_at`) VALUES ('tu0rkg0ari0coq6cnpy9w5z400j0k7dm', '测试', 'work_weixin', 2, '2022-12-01 11:52:06');
INSERT INTO `notice_config` (`id`, `title`, `send_gateway`, `types`, `created_at`) VALUES ('tu0rkg0en80coq6dc5v5ppy1000ecyo1', '测试类型', 'work_weixin', 0, '2022-12-01 11:52:59');
INSERT INTO `notice_config` (`id`, `title`, `send_gateway`, `types`, `created_at`) VALUES ('tu0rkg0g4d0couk34jdudr7100b5f41f', '告警测试', 'mail', 1, '2022-12-06 15:28:34');
INSERT INTO `notice_config` (`id`, `title`, `send_gateway`, `types`, `created_at`) VALUES ('tu0rkg0hq90cojukvo1u0oyd00vawtmc', 'text3', 'work_weixin', 0, '2022-11-24 01:22:37');
INSERT INTO `notice_config` (`id`, `title`, `send_gateway`, `types`, `created_at`) VALUES ('tu0rkg0hq90cojulmymp2pne00czr9m4', '111', 'work_weixin', 0, '2022-11-24 01:23:36');
INSERT INTO `notice_config` (`id`, `title`, `send_gateway`, `types`, `created_at`) VALUES ('tu0rkg0hq90cojummccsd4ug00u4m7mp', '11', 'work_weixin', 1, '2022-11-24 01:24:53');
INSERT INTO `notice_config` (`id`, `title`, `send_gateway`, `types`, `created_at`) VALUES ('tu0rkg0hq90cojumuehlvx0h00mlssjy', '33', 'work_weixin', 2, '2022-11-24 01:25:11');
INSERT INTO `notice_config` (`id`, `title`, `send_gateway`, `types`, `created_at`) VALUES ('tu0rkg0hq90cojuowfh931wi00vbhnck', '天天', 'sms', 1, '2022-11-24 01:27:52');
COMMIT;

-- ----------------------------
-- Table structure for notice_info
-- ----------------------------
DROP TABLE IF EXISTS `notice_info`;
CREATE TABLE `notice_info` (
  `id` bigint(10) NOT NULL AUTO_INCREMENT,
  `config_id` varchar(32) NOT NULL,
  `come_from` varchar(100) NOT NULL,
  `method` varchar(30) NOT NULL,
  `msg_title` varchar(100) NOT NULL,
  `msg_body` varchar(300) NOT NULL,
  `msg_url` varchar(200) NOT NULL,
  `user_ids` varchar(500) NOT NULL,
  `org_ids` varchar(500) NOT NULL,
  `totag` varchar(100) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `method_cron` varchar(50) NOT NULL,
  `method_num` int(10) NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='通知服务发送信息表';

-- ----------------------------
-- Records of notice_info
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for notice_log
-- ----------------------------
DROP TABLE IF EXISTS `notice_log`;
CREATE TABLE `notice_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `send_gateway` varchar(32) NOT NULL DEFAULT '' COMMENT '通知渠道',
  `template_id` varchar(32) NOT NULL DEFAULT '' COMMENT '通知模板ID',
  `addressee` varchar(255) NOT NULL DEFAULT '' COMMENT '收信人列表',
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '通知标题',
  `content` varchar(500) NOT NULL DEFAULT '' COMMENT '通知内容',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '发送状态：0=失败，1=成功',
  `fail_msg` varchar(500) NOT NULL DEFAULT '' COMMENT '失败信息',
  `send_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '发送时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=81985 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of notice_log
-- ----------------------------
BEGIN;
INSERT INTO `notice_log` (`id`, `send_gateway`, `template_id`, `addressee`, `title`, `content`, `status`, `fail_msg`, `send_time`) VALUES (25591, 'sms', 'tu0rkg018o0colupp3s7vyv900vk3s6n', '13700005102', '5', '6', 1, '', '2022-12-20 00:00:46');
INSERT INTO `notice_log` (`id`, `send_gateway`, `template_id`, `addressee`, `title`, `content`, `status`, `fail_msg`, `send_time`) VALUES (25592, 'sms', 'tu0rkg018o0colupp3s7vyv900vk3s6n', '13700005102', '5', '6', 1, '', '2022-12-20 00:00:46');
INSERT INTO `notice_log` (`id`, `send_gateway`, `template_id`, `addressee`, `title`, `content`, `status`, `fail_msg`, `send_time`) VALUES (25593, 'mail', 'tu0rkg0ijj0cozi9ct6u7jr500vs6shu', 'xhpeng11@163.com', '告警邮件模板', '<div>你好，你的系统有如下告警：</div>\n<div> </div>\n<div>产品：模拟测试电表2022 </div>\n<div>设备：南门电表03 </div>\n<div>级别：紧急 </div>\n<div>触发规则：南门电表 (va < 120) </div>', 1, '', '2022-12-20 00:00:47');
INSERT INTO `notice_log` (`id`, `send_gateway`, `template_id`, `addressee`, `title`, `content`, `status`, `fail_msg`, `send_time`) VALUES (25594, 'mail', 'tu0rkg0ijj0cozi9ct6u7jr500vs6shu', 'xhpeng11@163.com', '告警邮件模板', '<div>你好，你的系统有如下告警：</div>\n<div> </div>\n<div>产品：模拟测试电表2022 </div>\n<div>设备：南门电表03 </div>\n<div>级别：紧急 </div>\n<div>触发规则：南门电表 (va < 120) </div>', 1, '', '2022-12-20 00:00:47');
INSERT INTO `notice_log` (`id`, `send_gateway`, `template_id`, `addressee`, `title`, `content`, `status`, `fail_msg`, `send_time`) VALUES (25595, 'sms', 'tu0rkg018o0colupp3s7vyv900vk3s6n', '13700005102', '5', '6', 1, '', '2022-12-20 00:00:52');
INSERT INTO `notice_log` (`id`, `send_gateway`, `template_id`, `addressee`, `title`, `content`, `status`, `fail_msg`, `send_time`) VALUES (25596, 'sms', 'tu0rkg018o0colupp3s7vyv900vk3s6n', '13700005102', '5', '6', 1, '', '2022-12-20 00:00:53');
INSERT INTO `notice_log` (`id`, `send_gateway`, `template_id`, `addressee`, `title`, `content`, `status`, `fail_msg`, `send_time`) VALUES (25597, 'mail', 'tu0rkg0ijj0cozi9ct6u7jr500vs6shu', 'xhpeng11@163.com', '告警邮件模板', '<div>你好，你的系统有如下告警：</div>\n<div> </div>\n<div>产品：模拟测试电表2022 </div>\n<div>设备：南门电表03 </div>\n<div>级别：紧急 </div>\n<div>触发规则：南门电表 (va < 120) </div>', 1, '', '2022-12-20 00:00:54');
INSERT INTO `notice_log` (`id`, `send_gateway`, `template_id`, `addressee`, `title`, `content`, `status`, `fail_msg`, `send_time`) VALUES (25598, 'mail', 'tu0rkg0ijj0cozi9ct6u7jr500vs6shu', 'xhpeng11@163.com', '告警邮件模板', '<div>你好，你的系统有如下告警：</div>\n<div> </div>\n<div>产品：模拟测试电表2022 </div>\n<div>设备：南门电表03 </div>\n<div>级别：紧急 </div>\n<div>触发规则：南门电表 (va < 120) </div>', 1, '', '2022-12-20 00:00:54');
COMMIT;


-- ----------------------------
-- Table structure for notice_template
-- ----------------------------
DROP TABLE IF EXISTS `notice_template`;
CREATE TABLE `notice_template` (
  `id` varchar(32) NOT NULL,
  `config_id` varchar(32) NOT NULL,
  `send_gateway` varchar(32) NOT NULL,
  `code` varchar(32) NOT NULL,
  `title` varchar(100) NOT NULL,
  `content` varchar(500) NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `config_id` (`config_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='通知模版表';

-- ----------------------------
-- Records of notice_template
-- ----------------------------
BEGIN;
INSERT INTO `notice_template` (`id`, `config_id`, `send_gateway`, `code`, `title`, `content`, `created_at`) VALUES ('1vfwi9pb2p0coqh3cc49od4400nvviwj', 'tu0rkg0en80coq6dc5v5ppy1000ecyo1', 'work_weixin', '', '某项地ssss', 'erreeerrere222', '2022-12-01 20:08:04');
INSERT INTO `notice_template` (`id`, `config_id`, `send_gateway`, `code`, `title`, `content`, `created_at`) VALUES ('tu0rkg018o0colum0is6t5w700rkqkf0', 'tu0rkg0hq90cojumuehlvx0h00mlssjy', 'work_weixin', '', '注册通知', '注册通知注册通知注册通知', '2022-11-26 09:49:26');
INSERT INTO `notice_template` (`id`, `config_id`, `send_gateway`, `code`, `title`, `content`, `created_at`) VALUES ('tu0rkg018o0colupp3s7vyv900vk3s6n', 'tu0rkg0hq90cojuowfh931wi00vbhnck', 'sms', '', '5', '6', '2022-11-26 09:54:14');
INSERT INTO `notice_template` (`id`, `config_id`, `send_gateway`, `code`, `title`, `content`, `created_at`) VALUES ('tu0rkg06g30coqhxytbh8ly500ybwlre', 'tu0rkg0ari0coq6cnpy9w5z400j0k7dm', 'work_weixin', '', '234', '234324', '2022-12-01 20:57:08');
INSERT INTO `notice_template` (`id`, `config_id`, `send_gateway`, `code`, `title`, `content`, `created_at`) VALUES ('tu0rkg0ijj0cozi9ct6u7jr500vs6shu', 'tu0rkg0g4d0couk34jdudr7100b5f41f', 'mail', '', '告警邮件模板', '<div>你好，你的系统有如下告警：</div>\n<div> </div>\n<div>产品：{{.Product}} </div>\n<div>设备：{{.Device}} </div>\n<div>级别：{{.Level}} </div>\n<div>触发规则：{{.Rule}} </div>', '2022-12-06 15:28:51');
COMMIT;

-- ----------------------------
-- Table structure for sys_api
-- ----------------------------
DROP TABLE IF EXISTS `sys_api`;
CREATE TABLE `sys_api` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL COMMENT '名称',
  `types` int(11) NOT NULL COMMENT '1 分类 2接口',
  `method` varchar(255) DEFAULT NULL COMMENT '请求方式(数据字典维护)',
  `address` varchar(255) DEFAULT NULL COMMENT '接口地址',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `status` int(11) NOT NULL COMMENT '状态 0 停用 1启用',
  `sort` int(11) DEFAULT NULL COMMENT '排序',
  `is_deleted` int(11) NOT NULL COMMENT '是否删除 0未删除 1已删除',
  `create_by` int(10) unsigned DEFAULT '0' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` int(10) unsigned DEFAULT '0' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_by` int(11) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=311 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='API 接口管理表';

-- ----------------------------
-- Records of sys_api
-- ----------------------------
BEGIN;
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (1, 90, '产品详情接口', 2, NULL, '/api/v1/system/device/getAll', 'asdfasdf', 1, NULL, 1, 1, '2022-08-09 23:56:20', 1, '2022-11-03 11:19:25', 1, '2022-11-03 03:19:25');
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (2, 90, '123', 2, NULL, '123123222', '', 1, NULL, 1, 1, '2022-08-14 16:55:41', 1, '2022-08-15 10:18:21', 1, '2022-08-15 02:18:21');
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (3, 90, '234', 2, NULL, '12123123123', '', 1, NULL, 1, 1, '2022-08-15 09:26:17', 1, '2022-08-15 10:27:51', 1, '2022-08-15 02:27:51');
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (4, 90, '产品分类借楼', 2, NULL, '123123', '', 1, NULL, 1, 1, '2022-08-17 14:37:28', 1, '2022-11-03 11:19:28', 1, '2022-11-03 03:19:28');
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (55, 90, '获取城市的风力及日照时长', 2, NULL, '/api/v1/envirotronics/weather/cityWeatherList', '', 1, NULL, 0, 1, '2022-11-09 23:09:21', 0, '2022-11-09 23:09:21', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (56, 90, '根据ID获取指定城市的风力图表', 2, NULL, '/api/v1/envirotronics/weather/getWindpowerEchartById', '', 1, NULL, 0, 1, '2022-11-09 23:11:05', 0, '2022-11-09 23:11:05', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (57, 90, '根据ID获取指定城市的温度图表', 2, NULL, '/api/v1/envirotronics/weather/getTemperatureEchartById', '', 1, NULL, 0, 1, '2022-11-09 23:11:34', 0, '2022-11-09 23:11:34', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (58, 90, '根据ID获取指定城市的天气', 2, NULL, '/api/v1/envirotronics/weather/getInfoById', '', 1, NULL, 0, 1, '2022-11-09 23:11:58', 0, '2022-11-09 23:11:58', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (59, 90, '产品搜索列表（分页）', 2, NULL, '/api/v1/product/page_list', '', 1, NULL, 0, 1, '2022-11-09 15:18:02', 1, '2022-11-09 23:28:13', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (60, 90, '产品分类列表', 2, NULL, '/api/v1/product/category/list', '', 1, NULL, 0, 1, '2022-11-09 23:29:19', 0, '2022-11-09 23:29:19', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (61, 90, '获取部门列表', 2, NULL, '/api/v1/system/dept/tree', '', 1, NULL, 0, 1, '2022-11-09 15:30:37', 1, '2022-11-10 22:47:24', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (83, 90, '用户列表', 2, NULL, '/api/v1/system/user/list', '', 1, NULL, 0, 1, '2022-11-10 00:13:04', 0, '2022-11-10 00:13:04', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (84, 90, '添加用户', 2, NULL, '/api/v1/system/user/add', '', 1, NULL, 0, 1, '2022-11-10 00:13:38', 0, '2022-11-10 00:13:38', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (85, 90, '根据ID删除用户', 2, NULL, '/api/v1/system/user/delInfoById', '', 1, NULL, 0, 1, '2022-11-10 00:14:14', 0, '2022-11-10 00:14:14', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (86, 90, '编辑用户', 2, NULL, '/api/v1/system/user/edit', '', 1, NULL, 0, 1, '2022-11-10 00:14:34', 0, '2022-11-10 00:14:34', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (87, 90, '根据ID获取用户', 2, NULL, '/api/v1/system/user/getInfoById', '', 1, NULL, 0, 1, '2022-11-10 00:15:04', 0, '2022-11-10 00:15:04', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (88, 90, '所有用户列表', 2, NULL, '/api/v1/system/user/getAll', '', 1, NULL, 0, 1, '2022-11-10 00:15:39', 0, '2022-11-10 00:15:39', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (89, 90, '重置用户密码', 2, NULL, '/api/v1/system/user/resetPassword', '', 1, NULL, 0, 1, '2022-11-10 00:16:04', 0, '2022-11-10 00:16:04', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (90, -1, '接口分类', 1, NULL, '', NULL, 1, NULL, 0, 10, '2022-11-10 01:09:20', 0, NULL, NULL, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (91, 90, '数据源：搜索列表', 2, NULL, 'source/search', '', 1, NULL, 0, 1, '2022-11-10 11:13:53', 0, '2022-11-10 11:13:53', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (92, 90, '产品分类：添加', 2, NULL, '/api/v1/product/category/add', '', 1, NULL, 0, 1, '2022-11-10 11:39:54', 0, '2022-11-10 11:39:54', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (93, 90, '产品分类：编辑', 2, NULL, '/api/v1/product/category/edit', '', 1, NULL, 0, 1, '2022-11-10 11:40:25', 0, '2022-11-10 11:40:25', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (94, 90, '产品分类：删除', 2, NULL, '/api/v1/product/category/del', '', 1, NULL, 0, 1, '2022-11-10 11:41:08', 0, '2022-11-10 11:41:08', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (95, 90, 'asdf ', 2, NULL, 'afd', '', 1, NULL, 0, 1, '2022-11-10 12:03:49', 0, '2022-11-10 12:03:49', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (96, 90, '产品：添加', 2, NULL, '/api/v1/product/add', '', 1, NULL, 0, 1, '2022-11-10 13:01:09', 0, '2022-11-10 13:01:09', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (97, 90, '产品：编辑', 2, NULL, '/api/v1/product/edit', '', 1, NULL, 0, 1, '2022-11-10 13:01:37', 0, '2022-11-10 13:01:37', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (98, 90, '产品：删除', 2, NULL, '/api/v1/product/del', '', 1, NULL, 0, 1, '2022-11-10 13:01:54', 0, '2022-11-10 13:01:54', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (99, 90, '产品：发布', 2, NULL, '/api/v1/product/deploy', '', 1, NULL, 0, 1, '2022-11-10 13:02:10', 0, '2022-11-10 13:02:10', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (100, 90, '产品：停用', 2, NULL, '/api/v1/product/undeploy', '', 1, NULL, 0, 1, '2022-11-10 13:02:24', 0, '2022-11-10 13:02:24', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (101, 90, '产品：图标上传', 2, NULL, '/api/v1/product/icon/upload', '', 1, NULL, 0, 1, '2022-11-10 13:02:41', 0, '2022-11-10 13:02:41', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (102, 90, '产品：列表', 2, NULL, '/api/v1/product/list', '', 1, NULL, 0, 1, '2022-11-10 13:02:57', 0, '2022-11-10 13:02:57', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (103, 90, '产品：详情', 2, NULL, '/api/v1/product/detail', '', 1, NULL, 0, 1, '2022-11-10 13:03:17', 0, '2022-11-10 13:03:17', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (104, 90, '产品：详情，按key查询', 2, NULL, '/api/v1/product/get', '', 1, NULL, 0, 1, '2022-11-10 13:03:43', 0, '2022-11-10 13:03:43', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (105, 90, '设备：添加', 2, NULL, '/api/v1/product/device/add', '', 1, NULL, 0, 1, '2022-11-10 13:04:19', 0, '2022-11-10 13:04:19', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (106, 90, '设备：编辑', 2, NULL, '/api/v1/product/device/edit', '', 1, NULL, 0, 1, '2022-11-10 13:04:35', 0, '2022-11-10 13:04:35', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (107, 90, '设备：删除', 2, NULL, '/api/v1/product/device/del', '', 1, NULL, 0, 1, '2022-11-10 13:04:58', 0, '2022-11-10 13:04:58', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (108, 90, '设备：启用', 2, NULL, '/api/v1/product/device/deploy', '', 1, NULL, 0, 1, '2022-11-10 13:05:19', 0, '2022-11-10 13:05:19', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (109, 90, '设备：停用', 2, NULL, '/api/v1/product/device/undeploy', '', 1, NULL, 0, 1, '2022-11-10 13:05:34', 0, '2022-11-10 13:05:34', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (110, 90, '设备：上线', 2, NULL, '/api/v1/product/device/online', '', 1, NULL, 0, 1, '2022-11-10 13:06:13', 0, '2022-11-10 13:06:13', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (111, 90, '设备：下线', 2, NULL, '/api/v1/product/device/offline', '', 1, NULL, 0, 1, '2022-11-10 13:06:30', 0, '2022-11-10 13:06:30', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (112, 90, '设备：相关统计', 2, NULL, '/api/v1/product/device/statistics', '', 1, NULL, 0, 1, '2022-11-10 13:06:54', 0, '2022-11-10 13:06:54', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (113, 90, '设备：运行状态', 2, NULL, '/api/v1/product/device/run_status', '', 1, NULL, 0, 1, '2022-11-10 13:07:11', 0, '2022-11-10 13:07:11', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (114, 90, '设备：获取指定属性值', 2, NULL, '/api/v1/product/device/property/get', '', 1, NULL, 0, 1, '2022-11-10 13:07:26', 0, '2022-11-10 13:07:26', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (115, 90, '设备：属性详情列表', 2, NULL, '/api/v1/product/device/property/list', '', 1, NULL, 0, 1, '2022-11-10 13:07:43', 0, '2022-11-10 13:07:43', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (116, 90, '设备：已发布产品设备列表', 2, NULL, '/api/v1/product/device/list', '', 1, NULL, 0, 1, '2022-11-10 13:08:00', 0, '2022-11-10 13:08:00', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (117, 90, '设备：搜索列表(分页)', 2, NULL, '/api/v1/product/device/page_list', '', 1, NULL, 0, 1, '2022-11-10 13:08:18', 0, '2022-11-10 13:08:18', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (118, 90, '设备：详情', 2, NULL, '/api/v1/product/device/detail', '', 1, NULL, 0, 1, '2022-11-10 13:08:37', 0, '2022-11-10 13:08:37', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (119, 90, '设备：详情，按key查询', 2, NULL, '/api/v1/product/device/get', '', 1, NULL, 0, 1, '2022-11-10 13:09:01', 0, '2022-11-10 13:09:01', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (120, 90, '日志：类型列表', 2, NULL, '/api/v1/product/log/type', '', 1, NULL, 0, 1, '2022-11-10 13:09:33', 0, '2022-11-10 13:09:33', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (121, 90, '日志：搜索', 2, NULL, '/api/v1/product/log/search', '', 1, NULL, 0, 1, '2022-11-10 13:09:49', 0, '2022-11-10 13:09:49', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (122, 90, '数据类型：列表', 2, NULL, '/api/v1/product/tsl/data_type', '', 1, NULL, 0, 1, '2022-11-10 13:10:39', 0, '2022-11-10 13:10:39', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (123, 90, '属性：添加', 2, NULL, '/api/v1/product/tsl/property/add', '', 1, NULL, 0, 1, '2022-11-10 13:10:55', 0, '2022-11-10 13:10:55', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (124, 90, '属性：编辑', 2, NULL, '/api/v1/product/tsl/property/edit', '', 1, NULL, 0, 1, '2022-11-10 13:11:09', 0, '2022-11-10 13:11:09', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (125, 90, '属性：删除', 2, NULL, '/api/v1/product/tsl/property/del', '', 1, NULL, 0, 1, '2022-11-10 13:11:30', 0, '2022-11-10 13:11:30', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (126, 90, '属性：列表', 2, NULL, '/api/v1/product/tsl/property/list', '', 1, NULL, 0, 1, '2022-11-10 13:11:47', 0, '2022-11-10 13:11:47', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (127, 90, '属性：所有属性列表', 2, NULL, '/api/v1/product/tsl/property/all', '', 1, NULL, 0, 1, '2022-11-10 13:12:01', 0, '2022-11-10 13:12:01', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (128, 90, '功能：添加', 2, NULL, '/api/v1/product/tsl/function/add', '', 1, NULL, 0, 1, '2022-11-10 13:12:13', 0, '2022-11-10 13:12:13', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (129, 90, '功能：编辑', 2, NULL, '/api/v1/product/tsl/function/edit', '', 1, NULL, 0, 1, '2022-11-10 13:12:26', 0, '2022-11-10 13:12:26', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (130, 90, '功能：删除', 2, NULL, '/api/v1/product/tsl/function/del', '', 1, NULL, 0, 1, '2022-11-10 13:12:44', 0, '2022-11-10 13:12:44', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (131, 90, '功能：列表', 2, NULL, '/api/v1/product/tsl/function/list', '', 1, NULL, 0, 1, '2022-11-10 13:13:01', 0, '2022-11-10 13:13:01', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (132, 90, '事件：添加', 2, NULL, '/api/v1/product/tsl/event/add', '', 1, NULL, 0, 1, '2022-11-10 13:13:13', 0, '2022-11-10 13:13:13', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (133, 90, '事件：编辑', 2, NULL, '/api/v1/product/tsl/event/edit', '', 1, NULL, 0, 1, '2022-11-10 13:13:25', 0, '2022-11-10 13:13:25', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (134, 90, '事件：删除', 2, NULL, '/api/v1/product/tsl/event/del', '', 1, NULL, 0, 1, '2022-11-10 13:13:42', 0, '2022-11-10 13:13:42', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (135, 90, '事件：列表', 2, NULL, '/api/v1/product/tsl/event/list', '', 1, NULL, 0, 1, '2022-11-10 13:13:59', 0, '2022-11-10 13:13:59', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (136, 90, '标签：添加', 2, NULL, '/api/v1/product/tsl/tag/add', '', 1, NULL, 0, 1, '2022-11-10 13:14:13', 0, '2022-11-10 13:14:13', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (137, 90, '标签：编辑', 2, NULL, '/api/v1/product/tsl/tag/edit', '', 1, NULL, 0, 1, '2022-11-10 13:14:28', 0, '2022-11-10 13:14:28', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (138, 90, '标签：删除', 2, NULL, '/api/v1/product/tsl/tag/del', '', 1, NULL, 0, 1, '2022-11-10 13:14:43', 0, '2022-11-10 13:14:43', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (139, 90, '标签：列表', 2, NULL, '/api/v1/product/tsl/tag/list', '', 1, NULL, 0, 1, '2022-11-10 13:14:58', 0, '2022-11-10 13:14:58', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (140, 90, '数据库数据源：添加', 2, NULL, '/api/v1/source/db/add', '', 1, NULL, 0, 1, '2022-11-10 13:16:03', 0, '2022-11-10 13:16:03', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (141, 90, '数据库数据源：编辑', 2, NULL, '/api/v1/source/db/edit', '', 1, NULL, 0, 1, '2022-11-10 13:16:16', 0, '2022-11-10 13:16:16', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (142, 90, '数据库数据源：获取字段', 2, NULL, '/api/v1/source/db/fields', '', 1, NULL, 0, 1, '2022-11-10 13:16:42', 0, '2022-11-10 13:16:42', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (143, 90, '数据库数据源：获取数据', 2, NULL, '/api/v1/source/db/get', '', 1, NULL, 0, 1, '2022-11-10 13:16:59', 0, '2022-11-10 13:16:59', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (144, 90, '设备数据源：添加', 2, NULL, '/api/v1/source/device/add', '', 1, NULL, 0, 1, '2022-11-10 13:17:21', 0, '2022-11-10 13:17:21', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (145, 90, '设备数据源：编辑', 2, NULL, '/api/v1/source/device/edit', '', 1, NULL, 0, 1, '2022-11-10 13:17:37', 0, '2022-11-10 13:17:37', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (146, 90, '设备数据源：获取设备数据', 2, NULL, '/api/v1/source/device/get', '', 1, NULL, 0, 1, '2022-11-10 13:18:02', 0, '2022-11-10 13:18:02', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (147, 90, 'API 数据源：添加', 2, NULL, '/api/v1/source/api/add', '', 1, NULL, 0, 1, '2022-11-10 13:18:17', 0, '2022-11-10 13:18:17', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (148, 90, 'API 数据源：编辑', 2, NULL, '/api/v1/source/api/edit', '', 1, NULL, 0, 1, '2022-11-10 13:18:30', 0, '2022-11-10 13:18:30', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (149, 90, 'API 数据源：获取api数据', 2, NULL, '/api/v1/source/api/get', '', 1, NULL, 0, 1, '2022-11-10 13:18:47', 0, '2022-11-10 13:18:47', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (150, 90, '数据源：删除', 2, NULL, '/api/v1/source/del', '', 1, NULL, 0, 1, '2022-11-10 13:19:00', 0, '2022-11-10 13:19:00', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (151, 90, '数据源：列表', 2, NULL, '/api/v1/source/list', '', 1, NULL, 0, 1, '2022-11-10 13:19:28', 0, '2022-11-10 13:19:28', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (152, 90, '数据源：获取源数据记录', 2, NULL, '/api/v1/source/getdata', '', 1, NULL, 0, 1, '2022-11-10 13:19:43', 0, '2022-11-10 13:19:43', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (153, 90, '数据源：详情', 2, NULL, '/api/v1/source/detail', '', 1, NULL, 0, 1, '2022-11-10 13:20:00', 0, '2022-11-10 13:20:00', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (154, 90, '数据源：发布', 2, NULL, '/api/v1/source/deploy', '', 1, NULL, 0, 1, '2022-11-10 13:20:15', 0, '2022-11-10 13:20:15', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (155, 90, '数据源：停用', 2, NULL, '/api/v1/source/undeploy', '', 1, NULL, 0, 1, '2022-11-10 13:20:39', 0, '2022-11-10 13:20:39', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (156, 90, '数据源：复制数据源', 2, NULL, '/api/v1/source/copy', '', 1, NULL, 0, 1, '2022-11-10 13:20:55', 0, '2022-11-10 13:20:55', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (157, 90, '数据节点：添加', 2, NULL, '/api/v1/source/node/add', '', 1, NULL, 0, 1, '2022-11-10 13:21:16', 0, '2022-11-10 13:21:16', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (158, 90, '数据节点：编辑', 2, NULL, '/api/v1/source/node/edit', '', 1, NULL, 0, 1, '2022-11-10 13:21:31', 0, '2022-11-10 13:21:31', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (159, 90, '数据节点：删除', 2, NULL, '/api/v1/source/node/del', '', 1, NULL, 0, 1, '2022-11-10 13:21:48', 0, '2022-11-10 13:21:48', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (160, 90, '数据节点：列表', 2, NULL, '/api/v1/source/node/list', '', 1, NULL, 0, 1, '2022-11-10 13:22:07', 0, '2022-11-10 13:22:07', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (161, 90, '模型：添加', 2, NULL, '/api/v1/source/template/add', '', 1, NULL, 0, 1, '2022-11-10 13:22:52', 0, '2022-11-10 13:22:52', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (162, 90, '模型：编辑', 2, NULL, '/api/v1/source/template/edit', '', 1, NULL, 0, 1, '2022-11-10 13:23:06', 0, '2022-11-10 13:23:06', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (163, 90, '模型：删除', 2, NULL, '/api/v1/source/template/del', '', 1, NULL, 0, 1, '2022-11-10 13:23:20', 0, '2022-11-10 13:23:20', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (164, 90, '模型：搜索列表', 2, NULL, '/api/v1/source/template/search', '', 1, NULL, 0, 1, '2022-11-10 13:23:35', 0, '2022-11-10 13:23:35', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (165, 90, '模型：已发布列表', 2, NULL, '/api/v1/source/template/list', '', 1, NULL, 0, 1, '2022-11-10 13:23:48', 0, '2022-11-10 13:23:48', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (166, 90, '模型：详情', 2, NULL, '/api/v1/source/template/detail', '', 1, NULL, 0, 1, '2022-11-10 13:24:08', 0, '2022-11-10 13:24:08', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (167, 90, '模型：获取模型数据记录', 2, NULL, '/api/v1/source/template/getdata', '', 1, NULL, 0, 1, '2022-11-10 13:24:32', 0, '2022-11-10 13:24:32', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (168, 90, '模型：发布', 2, NULL, '/api/v1/source/template/deploy', '', 1, NULL, 0, 1, '2022-11-10 13:24:47', 0, '2022-11-10 13:24:47', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (169, 90, '模型：停用', 2, NULL, '/api/v1/source/template/undeploy', '', 1, NULL, 0, 1, '2022-11-10 13:25:10', 0, '2022-11-10 13:25:10', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (170, 90, '模型：复制数据模型', 2, NULL, '/api/v1/source/template/copy', '', 1, NULL, 0, 1, '2022-11-10 13:25:24', 0, '2022-11-10 13:25:24', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (171, 90, '节点：添加', 2, NULL, '/api/v1/source/template/node/add', '', 1, NULL, 0, 1, '2022-11-10 13:25:38', 0, '2022-11-10 13:25:38', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (172, 90, '节点：编辑', 2, NULL, '/api/v1/source/template/node/edit', '', 1, NULL, 0, 1, '2022-11-10 13:25:55', 0, '2022-11-10 13:25:55', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (173, 90, '节点：删除', 2, NULL, '/api/v1/source/template/node/del', '', 1, NULL, 0, 1, '2022-11-10 13:26:11', 0, '2022-11-10 13:26:11', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (174, 90, '节点：列表', 2, NULL, '/api/v1/source/template/node/list', '', 1, NULL, 0, 1, '2022-11-10 13:26:31', 0, '2022-11-10 13:26:31', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (175, 90, '获取组织列表', 2, NULL, '/api/v1/system/organization/tree', '', 1, NULL, 0, 1, '2022-11-10 22:38:27', 0, '2022-11-10 22:38:27', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (176, 90, '添加组织', 2, NULL, '/api/v1/system/organization/add', '', 1, NULL, 0, 1, '2022-11-10 22:39:19', 0, '2022-11-10 22:39:19', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (177, 90, '根据ID删除组织', 2, NULL, '/api/v1/system/organization/del', '', 1, NULL, 0, 1, '2022-11-10 22:39:38', 0, '2022-11-10 22:39:38', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (178, 90, '根据ID获取组织详情', 2, NULL, '/api/v1/system/organization/detail', '', 1, NULL, 0, 1, '2022-11-10 22:39:57', 0, '2022-11-10 22:39:57', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (179, 90, '编辑组织', 2, NULL, '/api/v1/system/organization/edit', '', 1, NULL, 0, 1, '2022-11-10 22:40:17', 0, '2022-11-10 22:40:17', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (180, 90, '添加部门', 2, NULL, '/api/v1/system/dept/add', '', 1, NULL, 0, 1, '2022-11-10 22:40:50', 0, '2022-11-10 22:40:50', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (181, 90, '根据ID删除部门', 2, NULL, '/api/v1/system/dept/del', '', 1, NULL, 0, 1, '2022-11-10 22:41:18', 0, '2022-11-10 22:41:18', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (182, 90, '根据ID获取部门详情', 2, NULL, '/api/v1/system/dept/detail', '', 1, NULL, 0, 1, '2022-11-10 22:41:38', 0, '2022-11-10 22:41:38', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (183, 90, '编辑部门', 2, NULL, '/api/v1/system/dept/edit', '', 1, NULL, 0, 1, '2022-11-10 22:42:03', 0, '2022-11-10 22:42:03', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (184, 90, '添加岗位', 2, NULL, '/api/v1/system/post/add', '', 1, NULL, 0, 1, '2022-11-10 22:47:57', 0, '2022-11-10 22:47:57', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (185, 90, '根据ID删除岗位', 2, NULL, '/api/v1/system/post/del', '', 1, NULL, 0, 1, '2022-11-10 22:48:17', 0, '2022-11-10 22:48:17', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (186, 90, '根据ID获取岗位详情', 2, NULL, '/api/v1/system/post/detail', '', 1, NULL, 0, 1, '2022-11-10 22:48:36', 0, '2022-11-10 22:48:36', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (187, 90, '编辑岗位', 2, NULL, '/api/v1/system/post/edit', '', 1, NULL, 0, 1, '2022-11-10 22:49:02', 0, '2022-11-10 22:49:02', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (188, 90, '获取岗位列表', 2, NULL, '/api/v1/system/post/tree', '', 1, NULL, 0, 1, '2022-11-10 22:49:30', 0, '2022-11-10 22:49:30', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (189, 90, '添加角色', 2, NULL, '/api/v1/system/role/add', '', 1, NULL, 0, 1, '2022-11-10 22:51:06', 0, '2022-11-10 22:51:06', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (190, 90, '角色数据权限授权', 2, NULL, '/api/v1/system/role/dataScope', '', 1, NULL, 0, 1, '2022-11-10 22:52:12', 0, '2022-11-10 22:52:12', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (191, 90, '根据ID删除角色', 2, NULL, '/api/v1/system/role/delInfoById', '\n', 1, NULL, 0, 1, '2022-11-10 22:52:34', 0, '2022-11-10 22:52:34', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (192, 90, '编辑角色', 2, NULL, '/api/v1/system/role/edit', '\n', 1, NULL, 0, 1, '2022-11-10 22:57:24', 0, '2022-11-10 22:57:24', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (193, 90, '根据ID获取权限信息', 2, NULL, '/api/v1/system/role/getAuthorizeById', '', 1, NULL, 0, 1, '2022-11-10 22:57:45', 0, '2022-11-10 22:57:45', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (194, 90, '根据ID获取角色', 2, NULL, '/api/v1/system/role/getInfoById', '', 1, NULL, 0, 1, '2022-11-10 22:58:06', 0, '2022-11-10 22:58:06', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (195, 90, '角色树状列表', 2, NULL, '/api/v1/system/role/tree', '', 1, NULL, 0, 1, '2022-11-10 22:58:24', 0, '2022-11-10 22:58:24', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (196, 90, '服务监控', 2, NULL, '/api/v1/system/monitor/server', '', 1, NULL, 0, 1, '2022-11-10 22:59:01', 0, '2022-11-10 22:59:01', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (197, 90, '访问日志列表', 2, NULL, '/api/v1/system/login/log/list', '', 1, NULL, 0, 1, '2022-11-10 23:00:17', 0, '2022-11-10 23:00:17', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (198, 90, '根据ID删除访问日志', 2, NULL, '/api/v1/system/login/log/del', '', 1, NULL, 0, 1, '2022-11-10 23:00:44', 0, '2022-11-10 23:00:44', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (199, 90, '根据ID获取访问日志详情', 2, NULL, '/api/v1/system/login/log/detail', '', 1, NULL, 0, 1, '2022-11-10 23:01:09', 0, '2022-11-10 23:01:09', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (200, 90, '根据ID删除操作日志', 2, NULL, '/api/v1/system/oper/log/del', '', 1, NULL, 0, 1, '2022-11-10 23:01:40', 0, '2022-11-10 23:01:40', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (201, 90, '根据ID获取操作日志详情', 2, NULL, '/api/v1/system/oper/log/detail', '', 1, NULL, 0, 1, '2022-11-10 23:03:07', 0, '2022-11-10 23:03:07', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (202, 90, '操作日志列表', 2, NULL, '/api/v1/system/oper/log/list', '', 1, NULL, 0, 1, '2022-11-10 23:03:30', 0, '2022-11-10 23:03:30', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (203, 90, '在线用户列表', 2, NULL, '/api/v1/system/userOnline/list', '', 1, NULL, 0, 1, '2022-11-10 23:05:00', 0, '2022-11-10 23:05:00', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (204, 90, '在线用户强退', 2, NULL, '/api/v1/system/userOnline/strongBack', '', 1, NULL, 0, 1, '2022-11-10 23:05:28', 0, '2022-11-10 23:05:28', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (205, 90, '获取插件', 2, NULL, '/api/v1/system/plugins/get', '', 1, NULL, 0, 1, '2022-11-10 23:07:42', 0, '2022-11-10 23:07:42', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (206, 90, '获取插件列表', 2, NULL, '/api/v1/system/plugins/list', '', 1, NULL, 0, 1, '2022-11-10 23:08:06', 0, '2022-11-10 23:08:06', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (207, 90, '设置插件状态', 2, NULL, '/api/v1/system/plugins/set', '', 1, NULL, 0, 1, '2022-11-10 23:09:41', 0, '2022-11-10 23:09:41', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (208, 90, '添加菜单', 2, NULL, '/api/v1/system/menu/add', '', 1, NULL, 0, 1, '2022-11-10 23:12:10', 0, '2022-11-10 23:12:10', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (209, 90, '添加菜单与按钮相关关联', 2, NULL, '/api/v1/system/menu/button/add', '', 1, NULL, 0, 1, '2022-11-10 23:12:33', 0, '2022-11-10 23:12:33', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (210, 90, '根据ID删除菜单按钮', 2, NULL, '/api/v1/system/menu/button/del', '', 1, NULL, 0, 1, '2022-11-10 23:12:49', 0, '2022-11-10 23:12:49', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (211, 90, '根据ID获取菜单按钮详情', 2, NULL, '/api/v1/system/menu/button/detail', '', 1, NULL, 0, 1, '2022-11-10 23:13:07', 0, '2022-11-10 23:13:07', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (212, 90, '编辑菜单按钮', 2, NULL, '/api/v1/system/menu/button/edit', '', 1, NULL, 0, 1, '2022-11-10 23:13:24', 0, '2022-11-10 23:13:24', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (213, 90, '编辑菜单按钮状态', 2, NULL, '/api/v1/system/menu/button/editStatus', '', 1, NULL, 0, 1, '2022-11-10 23:13:42', 0, '2022-11-10 23:13:42', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (214, 90, '菜单与按钮树列表', 2, NULL, '/api/v1/system/menu/button/tree', '', 1, NULL, 0, 1, '2022-11-10 23:14:04', 0, '2022-11-10 23:14:04', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (215, 90, '添加菜单与列表相关关联', 2, NULL, '/api/v1/system/menu/column/add', '', 1, NULL, 0, 1, '2022-11-10 23:14:30', 0, '2022-11-10 23:14:30', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (216, 90, '根据ID删除菜单列表', 2, NULL, '/api/v1/system/menu/column/del', '', 1, NULL, 0, 1, '2022-11-10 23:15:04', 0, '2022-11-10 23:15:04', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (217, 90, '根据ID获取菜单列表详情', 2, NULL, '/api/v1/system/menu/column/detail', '', 1, NULL, 0, 1, '2022-11-10 23:54:26', 0, '2022-11-10 23:54:26', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (218, 90, '编辑菜单列表', 2, NULL, '/api/v1/system/menu/column/edit', '', 1, NULL, 0, 1, '2022-11-10 23:57:04', 0, '2022-11-10 23:57:04', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (219, 90, '编辑菜单列表状态', 2, NULL, '/api/v1/system/menu/column/editStatus', '', 1, NULL, 0, 1, '2022-11-10 23:57:35', 0, '2022-11-10 23:57:35', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (220, 90, '菜单与列表树列表', 2, NULL, '/api/v1/system/menu/column/tree', '', 1, NULL, 0, 1, '2022-11-10 23:57:59', 0, '2022-11-10 23:57:59', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (221, 90, '根据ID删除菜单', 2, NULL, '/api/v1/system/menu/del', '', 1, NULL, 0, 1, '2022-11-10 23:58:34', 0, '2022-11-10 23:58:34', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (222, 90, '根据ID获取菜单详情', 2, NULL, '/api/v1/system/menu/detail', '', 1, NULL, 0, 1, '2022-11-10 23:58:55', 0, '2022-11-10 23:58:55', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (223, 90, '编辑菜单', 2, NULL, '/api/v1/system/menu/edit', '', 1, NULL, 0, 1, '2022-11-10 23:59:17', 0, '2022-11-10 23:59:17', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (224, 90, '菜单列表', 2, NULL, '/api/v1/system/menu/tree', '', 1, NULL, 0, 1, '2022-11-10 23:59:36', 0, '2022-11-10 23:59:36', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (225, -1, 'test', 2, NULL, 'sdasfasf', '', 1, NULL, 1, 1, '2022-11-11 00:14:59', 0, '2022-11-11 00:15:05', 1, '2022-11-10 16:15:05');
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (226, 90, '菜单相关', 2, NULL, '234234234', '', 1, NULL, 1, 1, '2022-11-11 00:15:14', 0, '2022-11-11 00:15:26', 1, '2022-11-10 16:15:26');
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (227, 90, '子分类', 1, NULL, '', '', 1, NULL, 1, 1, '2022-11-11 17:54:09', 0, '2022-11-11 18:01:02', 1, '2022-11-11 10:01:02');
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (228, -1, '12123', 2, NULL, '123123123', '123123', 1, NULL, 1, 1, '2022-11-11 01:54:34', 1, '2022-11-11 17:56:23', 1, '2022-11-11 09:56:23');
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (229, 228, '123', 2, NULL, '123123123', '', 1, NULL, 1, 1, '2022-11-11 17:55:19', 0, '2022-11-11 17:55:25', 1, '2022-11-11 09:55:25');
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (230, 227, '接口分类1123', 1, NULL, '', '123123', 1, NULL, 1, 1, '2022-11-11 17:56:44', 0, '2022-11-11 17:57:02', 1, '2022-11-11 09:57:02');
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (231, 227, '123', 1, NULL, '', '123123123', 1, NULL, 0, 1, '2022-11-11 18:00:05', 0, '2022-11-11 18:00:05', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (232, 227, '123', 1, NULL, '', '123123123', 1, NULL, 0, 1, '2022-11-11 18:00:05', 0, '2022-11-11 18:00:05', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (233, 231, '123123', 2, NULL, '123123', '', 1, NULL, 0, 1, '2022-11-11 18:00:54', 0, '2022-11-11 18:00:54', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (237, -1, '天气监测', 1, NULL, '', '', 1, NULL, 0, 1, '2022-11-11 21:23:33', 0, '2022-11-11 21:23:33', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (239, -1, '网络组件', 1, NULL, '', '', 1, NULL, 0, 1, '2022-11-11 22:32:37', 0, '2022-11-11 22:32:37', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (240, 239, '通道列表', 2, NULL, '/api/v1/network/tunnel/list', '', 1, NULL, 0, 1, '2022-11-11 22:35:50', 0, '2022-11-11 22:35:50', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (241, 239, '新增通道', 2, NULL, '/api/v1/network/tunnel/add', '', 1, NULL, 0, 1, '2022-11-11 14:38:03', 1, '2022-11-11 22:38:37', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (242, 239, '通道详情', 2, NULL, '/api/v1/network/tunnel/get', '', 1, NULL, 0, 1, '2022-11-11 22:39:17', 0, '2022-11-11 22:39:17', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (243, 239, '编辑通道', 2, NULL, '/api/v1/network/tunnel/edit', '', 1, NULL, 0, 1, '2022-11-11 14:39:52', 1, '2022-11-11 22:40:07', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (244, 239, '服务器管理列表', 2, NULL, '/api/v1/network/server/list', '', 1, NULL, 0, 1, '2022-11-11 22:41:50', 0, '2022-11-11 22:41:50', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (245, 239, '服务器详情', 2, NULL, '/api/v1/network/get', '', 1, NULL, 0, 1, '2022-11-11 14:42:19', 1, '2022-11-11 22:42:34', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (246, 239, '新增服务', 2, NULL, '/api/v1/network/server/add', '', 1, NULL, 0, 1, '2022-11-11 22:43:04', 0, '2022-11-11 22:43:04', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (247, 239, '编辑服务', 2, NULL, '/api/v1/network/server/edit', '', 1, NULL, 0, 1, '2022-11-11 22:43:53', 0, '2022-11-11 22:43:53', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (248, -1, '城市管理', 1, NULL, '', '', 1, NULL, 0, 1, '2022-11-11 22:45:36', 0, '2022-11-11 22:45:36', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (249, 248, '城市管理列表页', 2, NULL, '/api/v1/common/city/tree', '', 1, NULL, 0, 1, '2022-11-11 22:46:15', 0, '2022-11-11 22:46:15', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (250, 248, '添加城市', 2, NULL, '/api/v1/common/city/add', '', 1, NULL, 0, 1, '2022-11-11 22:47:06', 0, '2022-11-11 22:47:06', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (251, 248, '编辑城市', 2, NULL, '/api/v1/common/city/edit', '', 1, NULL, 0, 1, '2022-11-11 22:47:46', 0, '2022-11-11 22:47:46', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (252, 248, '删除城市', 2, NULL, '/api/v1/common/city/del', '', 1, NULL, 0, 1, '2022-11-11 22:48:30', 0, '2022-11-11 22:48:30', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (253, 248, '获取城市详情', 2, NULL, '/api/v1/common/city/getInfoById', '', 1, NULL, 0, 1, '2022-11-11 22:49:09', 0, '2022-11-11 22:49:09', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (254, -1, '系统配置', 1, NULL, '', '', 1, NULL, 0, 1, '2022-11-11 23:08:52', 0, '2022-11-11 23:08:52', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (255, 254, '添加系统参数', 2, NULL, '/api/v1/common/config/add', '', 1, NULL, 0, 1, '2022-11-11 23:09:43', 0, '2022-11-11 23:09:43', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (256, 254, '删除系统参数', 2, NULL, '/api/v1/common/config/delete', '', 1, NULL, 0, 1, '2022-11-11 23:10:13', 0, '2022-11-11 23:10:13', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (257, -1, '指数管理', 1, NULL, '', '', 1, NULL, 0, 1, '2022-11-11 23:10:21', 0, '2022-11-11 23:10:21', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (258, 254, '修改系统参数', 2, NULL, '/api/v1/common/config/edit', '', 1, NULL, 0, 1, '2022-11-11 23:10:35', 0, '2022-11-11 23:10:35', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (259, 257, '指数管理列表', 2, NULL, '/assess/v1/setup', '', 1, 0, 0, 1, '2022-11-11 23:10:49', 1, '2022-12-15 23:44:18', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (260, 254, '获取系统参数', 2, NULL, '/api/v1/common/config/get', '', 1, NULL, 0, 1, '2022-11-11 23:10:56', 0, '2022-11-11 23:10:56', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (261, 254, '系统参数列表', 2, NULL, '/api/v1/common/config/list', '', 1, NULL, 0, 1, '2022-11-11 23:11:16', 0, '2022-11-11 23:11:16', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (262, 257, '获取数据源信息', 2, NULL, '/assess/v1/datasetup/target', '', 1, 0, 0, 1, '2022-11-11 23:11:29', 1, '2022-12-20 21:03:54', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (263, 257, '测试数据源', 2, NULL, '/assess/v1/datasetup/test', '', 1, 0, 0, 1, '2022-11-11 23:12:03', 1, '2022-12-20 21:04:06', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (264, 254, '添加字典类型', 2, NULL, '/api/v1/common/dict/type/add', '', 1, NULL, 0, 1, '2022-11-11 23:12:19', 0, '2022-11-11 23:12:19', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (265, 257, '添加/编辑数据源信息', 2, 'GET', '/assess/v1/datasetup', '', 1, 0, 0, 1, '2022-11-11 23:12:44', 1, '2022-12-20 21:14:33', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (266, 254, '删除字典类型', 2, NULL, '/api/v1/common/dict/type/delete', '', 1, NULL, 0, 1, '2022-11-11 23:13:13', 0, '2022-11-11 23:13:13', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (267, 254, '修改字典类型', 2, NULL, '/api/v1/common/dict/type/edit', '', 1, NULL, 0, 1, '2022-11-11 23:13:53', 0, '2022-11-11 23:13:53', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (268, 254, '获取字典类型', 2, NULL, '/api/v1/common/dict/type/get', '', 1, NULL, 0, 1, '2022-11-11 23:14:16', 0, '2022-11-11 23:14:16', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (269, 254, '字典类型列表', 2, NULL, '/api/v1/common/dict/type/list', '', 1, NULL, 0, 1, '2022-11-11 23:14:39', 0, '2022-11-11 23:14:39', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (270, -1, '数据源管理', 1, NULL, '', '', 1, NULL, 0, 1, '2022-11-11 23:14:51', 0, '2022-11-11 23:14:51', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (271, 254, '添加字典数据', 2, NULL, '/api/v1/common/dict/data/add', '', 1, NULL, 0, 1, '2022-11-11 23:15:01', 0, '2022-11-11 23:15:01', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (272, 270, '数据源列表', 2, NULL, '/api/v1/source/search', '', 1, NULL, 0, 1, '2022-11-11 23:15:23', 0, '2022-11-11 23:15:23', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (273, 254, '删除字典数据', 2, NULL, '/api/v1/common/dict/data/delete', '', 1, NULL, 0, 1, '2022-11-11 23:15:26', 0, '2022-11-11 23:15:26', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (274, 254, '修改字典数据', 2, NULL, '/api/v1/common/dict/data/edit', '', 1, NULL, 0, 1, '2022-11-11 23:15:50', 0, '2022-11-11 23:15:50', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (275, 254, '获取字典数据', 2, NULL, '/api/v1/common/dict/data/get', '', 1, NULL, 0, 1, '2022-11-11 23:16:24', 0, '2022-11-11 23:16:24', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (276, 254, '获取字典数据公共方法', 2, NULL, '/api/v1/common/dict/data/getDictData', '', 1, NULL, 0, 1, '2022-11-11 23:16:45', 0, '2022-11-11 23:16:45', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (277, 254, '字典数据列表', 2, 'GET', '/api/v1/common/dict/data/list', '', 1, 0, 0, 1, '2022-11-11 23:17:07', 1, '2022-12-22 13:02:13', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (278, 254, '添加定时任务', 2, NULL, '/api/v1/system/job/add', '', 1, NULL, 0, 1, '2022-11-11 23:20:30', 0, '2022-11-11 23:20:30', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (279, 254, '根据ID删除任务', 2, NULL, '/api/v1/system/job/delJobById', '', 1, NULL, 0, 1, '2022-11-11 23:20:52', 0, '2022-11-11 23:20:52', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (280, 254, '编辑定时任务', 2, NULL, '/api/v1/system/job/edit', '', 1, NULL, 0, 1, '2022-11-11 23:21:12', 0, '2022-11-11 23:21:12', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (281, 254, '根据ID获取任务', 2, NULL, '/api/v1/system/job/getJobById', '', 1, NULL, 0, 1, '2022-11-11 23:21:35', 0, '2022-11-11 23:21:35', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (282, 254, '获取任务列表', 2, NULL, '/api/v1/system/job/list', '', 1, NULL, 0, 1, '2022-11-11 23:21:56', 0, '2022-11-11 23:21:56', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (283, 254, '执行一个任务', 2, NULL, '/api/v1/system/job/run', '', 1, NULL, 0, 1, '2022-11-11 23:22:19', 0, '2022-11-11 23:22:19', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (284, 254, '开始一个任务', 2, NULL, '/api/v1/system/job/start', '', 1, NULL, 0, 1, '2022-11-11 23:22:40', 0, '2022-11-11 23:22:40', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (285, 254, '结束一个任务', 2, NULL, '/api/v1/system/job/stop', '', 1, NULL, 0, 1, '2022-11-11 23:23:00', 0, '2022-11-11 23:23:00', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (286, 254, '获取所有接口', 2, NULL, '/api/v1/system/api/GetAll', '', 1, NULL, 0, 1, '2022-11-11 23:25:29', 0, '2022-11-11 23:25:29', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (287, 254, '添加Api', 2, NULL, '/api/v1/system/api/add', '', 1, NULL, 0, 1, '2022-11-11 23:25:57', 0, '2022-11-11 23:25:57', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (288, 254, '根据ID删除Api', 2, NULL, '/api/v1/system/api/del', '', 1, NULL, 0, 1, '2022-11-11 23:26:19', 0, '2022-11-11 23:26:19', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (289, 254, '根据ID获取Api详情', 2, NULL, '/api/v1/system/api/detail', '', 1, NULL, 0, 1, '2022-11-11 23:26:42', 0, '2022-11-11 23:26:42', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (290, 254, '编辑Api', 2, NULL, '/api/v1/system/api/edit', '', 1, NULL, 0, 1, '2022-11-11 23:27:03', 0, '2022-11-11 23:27:03', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (291, 254, '编辑Api状态', 2, NULL, '/api/v1/system/api/editStatus', '', 1, NULL, 0, 1, '2022-11-11 23:27:23', 0, '2022-11-11 23:27:23', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (292, 254, '获取接口列表', 2, NULL, '/api/v1/system/api/tree', '', 1, NULL, 0, 1, '2022-11-11 23:27:44', 0, '2022-11-11 23:27:44', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (293, -1, '告警管理', 1, NULL, '', '', 1, 0, 0, 1, '2022-11-23 10:37:37', 0, '2022-11-23 10:37:37', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (294, 293, '告警级别：列表', 2, NULL, '/api/v1/alarm/level/all', '', 1, 0, 0, 1, '2022-11-23 10:38:23', 0, '2022-11-23 10:38:23', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (295, 293, '告警级别：配置', 2, NULL, '/api/v1/alarm/level/edit', '', 1, 0, 0, 1, '2022-11-23 10:39:02', 0, '2022-11-23 10:39:02', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (296, 293, '告警规则：操作符列表', 2, NULL, '/api/v1/alarm/rule/operator', '', 1, 0, 0, 1, '2022-11-23 10:39:30', 0, '2022-11-23 10:39:30', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (297, 293, '告警规则：触发类型列表', 2, NULL, '/api/v1/alarm/rule/trigger_type', '', 1, 0, 0, 1, '2022-11-23 10:39:53', 0, '2022-11-23 10:39:53', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (298, 293, '告警规则：触发条件参数列表', 2, NULL, '/api/v1/alarm/rule/trigger_param', '', 1, 0, 0, 1, '2022-11-23 10:40:14', 0, '2022-11-23 10:40:14', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (299, 293, '告警规则：详情', 2, NULL, '/api/v1/alarm/rule/detail', '', 1, 0, 0, 1, '2022-11-23 10:40:37', 0, '2022-11-23 10:40:37', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (300, 293, '告警规则：列表', 2, NULL, '/api/v1/alarm/rule/list', '', 1, 0, 0, 1, '2022-11-23 10:41:01', 0, '2022-11-23 10:41:01', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (301, 293, '告警规则：添加', 2, NULL, '/api/v1/alarm/rule/add', '', 1, 0, 0, 1, '2022-11-23 10:41:20', 0, '2022-11-23 10:41:20', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (302, 293, '告警规则：编辑', 2, NULL, '/api/v1/alarm/rule/edit', '', 1, 0, 0, 1, '2022-11-23 10:41:45', 0, '2022-11-23 10:41:45', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (303, 293, '告警规则：启用', 2, NULL, '/api/v1/alarm/rule/deploy', '', 1, 0, 0, 1, '2022-11-23 10:42:06', 0, '2022-11-23 10:42:06', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (304, 293, '告警规则：禁用', 2, NULL, '/api/v1/alarm/rule/undeploy', '', 1, 0, 0, 1, '2022-11-23 10:42:25', 0, '2022-11-23 10:42:25', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (305, 293, '告警规则：删除', 2, NULL, '/api/v1/alarm/rule/del', '', 1, 0, 0, 1, '2022-11-23 10:42:41', 0, '2022-11-23 10:42:41', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (306, 293, '告警日志：详情', 2, NULL, '/api/v1/alarm/log/detail', '', 1, 0, 0, 1, '2022-11-23 10:43:05', 0, '2022-11-23 10:43:05', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (307, 293, '告警日志：列表', 2, NULL, '/api/v1/alarm/log/list', '', 1, 0, 0, 1, '2022-11-23 10:43:31', 0, '2022-11-23 10:43:31', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (308, 293, '告警日志：告警处理', 2, NULL, '/api/v1/alarm/log/handle', '', 1, 0, 0, 1, '2022-11-23 10:45:49', 0, '2022-11-23 10:45:49', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (309, -1, '物联概览', 1, NULL, '', '', 1, 0, 0, 1, '2022-11-28 10:29:39', 0, '2022-11-28 10:29:39', 0, NULL);
INSERT INTO `sys_api` (`id`, `parent_id`, `name`, `types`, `method`, `address`, `remark`, `status`, `sort`, `is_deleted`, `create_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (310, 309, '物联概览:查询', 2, NULL, '/api/v1/statistics/thing/overview', '', 1, 0, 0, 1, '2022-11-28 07:45:40', 1, '2022-11-28 15:45:58', 0, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_authorize
-- ----------------------------
DROP TABLE IF EXISTS `sys_authorize`;
CREATE TABLE `sys_authorize` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL COMMENT '角色ID',
  `items_type` varchar(50) NOT NULL COMMENT '项目类型 menu菜单 button按钮 column列表字段 api接口',
  `items_id` int(11) NOT NULL COMMENT '项目ID',
  `is_check_all` int(11) NOT NULL COMMENT '是否全选 1是 0否',
  `is_deleted` int(11) NOT NULL COMMENT '是否删除 0未删除 1已删除',
  `created_by` int(10) unsigned DEFAULT '0' COMMENT '创建人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `deleted_by` int(11) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8133 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='角色与菜单、按钮、列表权限关系';

-- ----------------------------
-- Records of sys_authorize
-- ----------------------------
BEGIN;
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7739, 7, 'menu', 113, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7740, 7, 'menu', 56, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7741, 7, 'menu', 57, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7742, 7, 'menu', 84, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7743, 7, 'menu', 95, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7744, 7, 'menu', 101, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7745, 7, 'menu', 104, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7746, 7, 'menu', 105, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7747, 7, 'menu', 106, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7748, 7, 'menu', 107, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7749, 7, 'menu', 75, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7750, 7, 'menu', 76, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7751, 7, 'menu', 77, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7752, 7, 'menu', 78, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7753, 7, 'menu', 79, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7754, 7, 'menu', 118, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7755, 7, 'menu', 53, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7756, 7, 'menu', 55, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7757, 7, 'menu', 54, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7758, 7, 'menu', 100, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7759, 7, 'menu', 58, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7760, 7, 'menu', 59, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7761, 7, 'menu', 60, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7762, 7, 'menu', 62, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7763, 7, 'menu', 63, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7764, 7, 'menu', 92, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7765, 7, 'menu', 11, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7766, 7, 'menu', 34, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7767, 7, 'menu', 52, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7768, 7, 'menu', 36, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7769, 7, 'menu', 51, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7770, 7, 'menu', 35, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7771, 7, 'menu', 65, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7772, 7, 'menu', 66, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7773, 7, 'menu', 68, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7774, 7, 'menu', 69, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7775, 7, 'menu', 74, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7776, 7, 'menu', 70, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7777, 7, 'menu', 71, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7778, 7, 'menu', 72, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7779, 7, 'menu', 73, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7780, 7, 'menu', 23, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7781, 7, 'menu', 24, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7782, 7, 'menu', 25, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7783, 7, 'menu', 44, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7784, 7, 'menu', 40, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7785, 7, 'menu', 41, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7786, 7, 'menu', 116, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7787, 7, 'menu', 117, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7788, 7, 'menu', 42, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7789, 7, 'menu', 80, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7790, 7, 'menu', 114, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7791, 7, 'menu', 115, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7792, 7, 'menu', 1, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7793, 7, 'menu', 15, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7794, 7, 'menu', 13, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7795, 7, 'menu', 12, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7796, 7, 'menu', 14, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7797, 7, 'menu', 10, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7798, 7, 'menu', 17, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7799, 7, 'menu', 18, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7800, 7, 'menu', 49, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7801, 7, 'menu', 50, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7802, 7, 'menu', 110, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7803, 7, 'menu', 16, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7804, 7, 'menu', 21, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7805, 7, 'menu', 22, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7806, 7, 'menu', 46, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7807, 7, 'menu', 85, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7808, 7, 'menu', 91, 0, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7809, 7, 'menu', 19, 0, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7810, 7, 'menu', 20, 0, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7811, 7, 'button', 156, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7812, 7, 'button', 12, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7813, 7, 'button', 62, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7814, 7, 'button', 96, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7815, 7, 'button', 42, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7816, 7, 'button', 43, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7817, 7, 'button', 136, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7818, 7, 'button', 46, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7819, 7, 'button', 63, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7820, 7, 'button', 65, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7821, 7, 'button', 101, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7822, 7, 'button', 103, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7823, 7, 'button', 137, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7824, 7, 'button', 143, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7825, 7, 'button', 56, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7826, 7, 'button', 57, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7827, 7, 'button', 58, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7828, 7, 'button', 59, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7829, 7, 'button', 145, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7830, 7, 'button', 146, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7831, 7, 'button', 67, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7832, 7, 'button', 68, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7833, 7, 'button', 70, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7834, 7, 'button', 71, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7835, 7, 'button', 73, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7836, 7, 'button', 74, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7837, 7, 'button', 76, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7838, 7, 'button', 77, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7839, 7, 'button', 52, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7840, 7, 'button', 53, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7841, 7, 'button', 149, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7842, 7, 'column', 4, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7843, 7, 'column', 5, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7844, 7, 'column', 7, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7845, 7, 'column', 13, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7846, 7, 'column', 14, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7847, 7, 'column', 15, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7848, 7, 'column', 16, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7849, 7, 'column', 21, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7850, 7, 'column', 224, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7851, 7, 'column', 225, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7852, 7, 'column', 226, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7853, 7, 'column', 227, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7854, 7, 'column', 228, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7855, 7, 'column', 229, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7856, 7, 'column', 230, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7857, 7, 'column', 231, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7858, 7, 'column', 8, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7859, 7, 'column', 9, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7860, 7, 'column', 10, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7861, 7, 'column', 11, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7862, 7, 'column', 12, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7863, 7, 'column', 24, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7864, 7, 'column', 25, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7865, 7, 'column', 26, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7866, 7, 'column', 233, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7867, 7, 'column', 28, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7868, 7, 'column', 29, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7869, 7, 'column', 30, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7870, 7, 'column', 31, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7871, 7, 'column', 32, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7872, 7, 'column', 33, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7873, 7, 'column', 1, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7874, 7, 'column', 3, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7875, 7, 'column', 128, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7876, 7, 'column', 129, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7877, 7, 'column', 130, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7878, 7, 'column', 131, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7879, 7, 'column', 132, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7880, 7, 'column', 133, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7881, 7, 'column', 134, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7882, 7, 'column', 67, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7883, 7, 'column', 68, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7884, 7, 'column', 69, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7885, 7, 'column', 71, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7886, 7, 'column', 72, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7887, 7, 'column', 73, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7888, 7, 'column', 137, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7889, 7, 'column', 138, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7890, 7, 'column', 139, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7891, 7, 'column', 140, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7892, 7, 'column', 142, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7893, 7, 'column', 143, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7894, 7, 'column', 144, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7895, 7, 'column', 34, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7896, 7, 'column', 35, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7897, 7, 'column', 37, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7898, 7, 'column', 38, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7899, 7, 'column', 39, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7900, 7, 'column', 159, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7901, 7, 'column', 160, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7902, 7, 'column', 161, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7903, 7, 'column', 162, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7904, 7, 'column', 163, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7905, 7, 'column', 164, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7906, 7, 'column', 40, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7907, 7, 'column', 41, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7908, 7, 'column', 42, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7909, 7, 'column', 43, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7910, 7, 'column', 44, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7911, 7, 'column', 166, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7912, 7, 'column', 167, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7913, 7, 'column', 168, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7914, 7, 'column', 169, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7915, 7, 'column', 170, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7916, 7, 'column', 172, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7917, 7, 'column', 173, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7918, 7, 'column', 174, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7919, 7, 'column', 175, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7920, 7, 'column', 176, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7921, 7, 'column', 177, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7922, 7, 'column', 135, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7923, 7, 'column', 136, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7924, 7, 'column', 234, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7925, 7, 'column', 235, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7926, 7, 'column', 236, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7927, 7, 'column', 238, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7928, 7, 'column', 239, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7929, 7, 'column', 240, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7930, 7, 'column', 241, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7931, 7, 'column', 182, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7932, 7, 'column', 183, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7933, 7, 'column', 184, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7934, 7, 'column', 185, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7935, 7, 'column', 186, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7936, 7, 'column', 187, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7937, 7, 'column', 188, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7938, 7, 'column', 46, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7939, 7, 'column', 47, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7940, 7, 'column', 48, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7941, 7, 'column', 49, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7942, 7, 'column', 50, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7943, 7, 'column', 52, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7944, 7, 'column', 53, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7945, 7, 'column', 54, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7946, 7, 'column', 75, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7947, 7, 'column', 76, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7948, 7, 'column', 77, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7949, 7, 'column', 78, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7950, 7, 'column', 79, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7951, 7, 'column', 80, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7952, 7, 'column', 81, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7953, 7, 'column', 82, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7954, 7, 'column', 83, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7955, 7, 'column', 84, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7956, 7, 'column', 146, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7957, 7, 'column', 147, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7958, 7, 'column', 148, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7959, 7, 'column', 149, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7960, 7, 'column', 203, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7961, 7, 'column', 204, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7962, 7, 'column', 205, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7963, 7, 'column', 206, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7964, 7, 'column', 207, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7965, 7, 'column', 208, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7966, 7, 'column', 63, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7967, 7, 'column', 64, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7968, 7, 'column', 65, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7969, 7, 'column', 66, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7970, 7, 'column', 194, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7971, 7, 'column', 195, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7972, 7, 'column', 196, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7973, 7, 'column', 197, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7974, 7, 'column', 86, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7975, 7, 'column', 87, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7976, 7, 'column', 88, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7977, 7, 'column', 89, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7978, 7, 'column', 91, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7979, 7, 'column', 92, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7980, 7, 'column', 93, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7981, 7, 'column', 94, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7982, 7, 'column', 96, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7983, 7, 'column', 97, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7984, 7, 'column', 98, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7985, 7, 'column', 99, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7986, 7, 'column', 101, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7987, 7, 'column', 102, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7988, 7, 'column', 103, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7989, 7, 'column', 104, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7990, 7, 'column', 105, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7991, 7, 'column', 106, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7992, 7, 'column', 107, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7993, 7, 'column', 108, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7994, 7, 'column', 60, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7995, 7, 'column', 61, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7996, 7, 'column', 62, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7997, 7, 'column', 56, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7998, 7, 'column', 57, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7999, 7, 'column', 58, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8000, 7, 'column', 59, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8001, 7, 'column', 220, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8002, 7, 'column', 221, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8003, 7, 'column', 222, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8004, 7, 'column', 211, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8005, 7, 'column', 212, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8006, 7, 'column', 213, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8007, 7, 'column', 214, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8008, 7, 'column', 215, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8009, 7, 'column', 216, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8010, 7, 'column', 217, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8011, 7, 'column', 218, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8012, 7, 'api', 161, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8013, 7, 'api', 165, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8014, 7, 'api', 166, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8015, 7, 'api', 384, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8016, 7, 'api', 111, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8017, 7, 'api', 113, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8018, 7, 'api', 114, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8019, 7, 'api', 115, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8020, 7, 'api', 252, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8021, 7, 'api', 255, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8022, 7, 'api', 259, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8023, 7, 'api', 262, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8024, 7, 'api', 265, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8025, 7, 'api', 267, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8026, 7, 'api', 131, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8027, 7, 'api', 132, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8028, 7, 'api', 179, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8029, 7, 'api', 180, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8030, 7, 'api', 181, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8031, 7, 'api', 261, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8032, 7, 'api', 70, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8033, 7, 'api', 104, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8034, 7, 'api', 194, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8035, 7, 'api', 195, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8036, 7, 'api', 196, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8037, 7, 'api', 199, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8038, 7, 'api', 272, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8039, 7, 'api', 273, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8040, 7, 'api', 274, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8041, 7, 'api', 335, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8042, 7, 'api', 336, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8043, 7, 'api', 275, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8044, 7, 'api', 276, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8045, 7, 'api', 278, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8046, 7, 'api', 343, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8047, 7, 'api', 344, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8048, 7, 'api', 349, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8049, 7, 'api', 350, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8050, 7, 'api', 351, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8051, 7, 'api', 139, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8052, 7, 'api', 141, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8053, 7, 'api', 142, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8054, 7, 'api', 148, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8055, 7, 'api', 149, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8056, 7, 'api', 133, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8057, 7, 'api', 150, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8058, 7, 'api', 154, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8059, 7, 'api', 155, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8060, 7, 'api', 355, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8061, 7, 'api', 356, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8062, 7, 'api', 280, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8063, 7, 'api', 281, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8064, 7, 'api', 282, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8065, 7, 'api', 189, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8066, 7, 'api', 190, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8067, 7, 'api', 191, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8068, 7, 'api', 192, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8069, 7, 'api', 193, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8070, 7, 'api', 197, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8071, 7, 'api', 198, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8072, 7, 'api', 203, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8073, 7, 'api', 204, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8074, 7, 'api', 208, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8075, 7, 'api', 212, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8076, 7, 'api', 216, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8077, 7, 'api', 61, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8078, 7, 'api', 65, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8079, 7, 'api', 58, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8080, 7, 'api', 60, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8081, 7, 'api', 311, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8082, 7, 'api', 316, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8083, 7, 'api', 312, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8084, 7, 'api', 319, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8085, 7, 'api', 323, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8086, 7, 'api', 320, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8087, 7, 'api', 321, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8088, 7, 'api', 313, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8089, 7, 'api', 112, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8090, 7, 'api', 105, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8091, 7, 'api', 106, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8092, 7, 'api', 109, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8093, 7, 'api', 110, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8094, 7, 'api', 116, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8095, 7, 'api', 117, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8096, 7, 'api', 118, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8097, 7, 'api', 119, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8098, 7, 'api', 310, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8099, 7, 'api', 326, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8100, 7, 'api', 330, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8101, 7, 'api', 71, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8102, 7, 'api', 72, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8103, 7, 'api', 73, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8104, 7, 'api', 74, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8105, 7, 'api', 75, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8106, 7, 'api', 82, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8107, 7, 'api', 83, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8108, 7, 'api', 87, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8109, 7, 'api', 89, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8110, 7, 'api', 90, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8111, 7, 'api', 94, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8112, 7, 'api', 96, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8113, 7, 'api', 100, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8114, 7, 'api', 101, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8115, 7, 'api', 126, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8116, 7, 'api', 127, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8117, 7, 'api', 128, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8118, 7, 'api', 129, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8119, 7, 'api', 130, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8120, 7, 'api', 284, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8121, 7, 'api', 285, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8122, 7, 'api', 286, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8123, 7, 'api', 367, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8124, 7, 'api', 369, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8125, 7, 'api', 370, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8126, 7, 'api', 371, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8127, 7, 'api', 372, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8128, 7, 'api', 373, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8129, 7, 'api', 374, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8130, 7, 'api', 368, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8131, 7, 'api', 380, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
INSERT INTO `sys_authorize` (`id`, `role_id`, `items_type`, `items_id`, `is_check_all`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8132, 7, 'api', 381, 1, 0, 0, '2022-12-21 15:39:33', 0, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_config`;
CREATE TABLE `sys_config` (
  `config_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '参数主键',
  `config_name` varchar(100) NOT NULL DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) NOT NULL DEFAULT '' COMMENT '参数键名',
  `config_value` varchar(500) NOT NULL DEFAULT '' COMMENT '参数键值',
  `config_type` int(11) NOT NULL DEFAULT '0' COMMENT '系统内置（1是 2否）',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `status` int(11) NOT NULL COMMENT '状态 0 停用 1启用',
  `is_deleted` int(11) NOT NULL COMMENT '是否删除 0未删除 1已删除',
  `create_by` int(10) unsigned DEFAULT '0' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` int(10) unsigned DEFAULT '0' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_by` int(11) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`config_id`) USING BTREE,
  UNIQUE KEY `uni_config_key` (`config_key`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统参数';

-- ----------------------------
-- Records of sys_config
-- ----------------------------
BEGIN;
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (1, '文件上传-文件大小', 'sys.uploadFile.fileSize', '50M', 1, '文件上传大小限制', 0, 0, 0, '2022-08-06 12:42:19', 1, '2022-08-11 15:07:39', NULL, NULL);
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (2, '文件上传-文件类型', 'sys.uploadFile.fileType', 'doc,docx,zip,xls,xlsx,rar,jpg,jpeg,gif,npm,png', 1, '文件上传后缀类型限制', 0, 0, 0, '2022-08-06 12:44:00', 0, '2022-08-06 12:44:00', NULL, NULL);
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (3, '图片上传-图片类型', 'sys.uploadFile.imageType', 'jpg,jpeg,gif,npm,png,svg', 1, '图片上传后缀类型限制', 0, 0, 0, '2022-08-06 12:44:46', 0, '2022-08-06 12:44:46', NULL, NULL);
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (4, '图片上传-图片大小', 'sys.uploadFile.imageSize', '50M', 1, '图片上传大小限制', 0, 0, 0, '2022-08-06 12:45:35', 0, '2022-08-06 12:45:35', NULL, NULL);
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (5, '设置项1', '阿什顿发12123123123', '阿什顿发 大师傅', 0, '阿什顿发阿什顿发大师傅 ', 0, 0, 1, '2022-08-11 12:02:01', 1, '2022-08-11 12:02:32', NULL, '2022-08-11 12:02:37');
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (6, '123', '123', '123', 0, '123123123', 0, 0, 1, '2022-08-11 12:02:46', 0, '2022-08-11 12:02:46', NULL, '2022-08-11 12:02:49');
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (7, '系统名称', 'sys.system.name', '沙果IOT', 1, '', 0, 0, 1, '2022-08-12 21:54:54', 1, '2022-11-10 17:20:37', NULL, NULL);
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (8, '系统版本', 'sys.system.copyright', 'Sagoo.cn', 1, '', 0, 0, 1, '2022-08-27 13:21:55', 0, '2022-08-27 13:21:55', NULL, NULL);
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (9, '换热站失水数据建模ID', 'energy.loss.water', '10', 1, '换热站失水数据建模ID', 0, 0, 1, '2022-09-15 11:10:53', 1, '2022-09-17 17:48:23', NULL, NULL);
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (10, '换热站报警数据建模ID', 'energy.station.early.warning', '11', 0, '换热站温度压力报警数据建模', 0, 0, 1, '2022-09-17 17:46:38', 1, '2022-09-18 22:54:37', NULL, '2022-09-19 22:44:50');
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (11, '换热站列表', 'energy.station.lists', '11,12', 0, '通过API获取数据源的换热站列表', 0, 0, 1, '2022-09-18 22:56:05', 1, '2022-09-18 22:56:28', NULL, NULL);
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (12, '换热站基础数据', 'energy.station.infos', '14_B00114', 0, '格式为:数建模ID_换热站编号;多个使用英文逗号隔开', 0, 0, 1, '2022-09-19 22:46:11', 1, '2022-09-20 11:30:24', NULL, NULL);
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (13, '地图访问密钥', 'sys.map.access.key', 'sGgaYnGYF87fy9dVDvLF5GemYMH02aax', 1, '', 0, 0, 1, '2022-10-15 05:03:20', 1, '2022-10-15 05:05:29', NULL, NULL);
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (14, '地图中心点经纬度', 'sys.map.lngAndLat', '124.383044, 40.124296', 0, '地图中心点的经纬度', 0, 0, 1, '2022-11-15 21:32:17', 0, '2022-11-15 21:32:17', NULL, NULL);
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (15, '是否开启访问权限', 'sys.access.control', '0', 1, '0 否 1是', 0, 0, 1, '2022-11-22 10:06:41', 1, '2022-12-16 02:54:58', NULL, NULL);
INSERT INTO `sys_config` (`config_id`, `config_name`, `config_key`, `config_value`, `config_type`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (16, '是否自启定时任务', 'sys.auto.run.job', '1', 1, '0 否 1是', 0, 0, 1, '2022-11-22 22:40:32', 1, '2022-12-16 02:54:54', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_dept`;
CREATE TABLE `sys_dept` (
  `dept_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '部门id',
  `organization_id` int(11) NOT NULL COMMENT '组织ID',
  `parent_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '父部门id',
  `ancestors` varchar(1000) NOT NULL DEFAULT '' COMMENT '祖级列表',
  `dept_name` varchar(30) NOT NULL DEFAULT '' COMMENT '部门名称',
  `order_num` int(11) DEFAULT '0' COMMENT '显示顺序',
  `leader` varchar(20) DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '部门状态（0停用 1正常）',
  `is_deleted` int(11) NOT NULL COMMENT '是否删除 0未删除 1已删除',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `created_by` int(10) unsigned DEFAULT '0' COMMENT '创建人',
  `updated_by` int(11) DEFAULT NULL COMMENT '修改人',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_by` int(11) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='部门表';

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_dept` (`dept_id`, `organization_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `is_deleted`, `created_at`, `created_by`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (1, 1, -1, '1', '开发部', 1, '15888888888', 'ry@qq.com', '1', 1, 1, '2021-07-13 15:56:52', 1, NULL, '2022-08-10 23:43:00', 1, '2022-08-10 23:43:00');
INSERT INTO `sys_dept` (`dept_id`, `organization_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `is_deleted`, `created_at`, `created_by`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (2, 7, -1, '2', '技术管理部', 0, '张发财', '', '', 1, 0, '2022-08-10 07:38:36', 1, 1, '2022-08-12 20:30:11', 0, NULL);
INSERT INTO `sys_dept` (`dept_id`, `organization_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `is_deleted`, `created_at`, `created_by`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (3, 5, 2, '2,3', '维护组', 0, '阿什顿发', '', '', 1, 0, '2022-08-09 23:43:21', 1, 1, '2022-12-15 23:42:20', 0, NULL);
INSERT INTO `sys_dept` (`dept_id`, `organization_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `is_deleted`, `created_at`, `created_by`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (4, 5, -1, '4', '运营管理部', 0, '123', '', '', 1, 0, '2022-08-10 15:59:00', 1, 1, '2022-08-12 20:30:47', 0, NULL);
INSERT INTO `sys_dept` (`dept_id`, `organization_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `is_deleted`, `created_at`, `created_by`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (5, 11, -1, '5', '行政办公室', 0, '沙果管理员', '', '', 1, 0, '2022-08-11 01:35:55', 1, 1, '2022-08-12 20:31:00', 0, NULL);
INSERT INTO `sys_dept` (`dept_id`, `organization_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `is_deleted`, `created_at`, `created_by`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (6, 11, -1, '6', '管理员', 0, '管理员', '', '', 1, 0, '2022-08-12 23:15:01', 1, 0, '2022-08-12 23:15:01', 0, NULL);
INSERT INTO `sys_dept` (`dept_id`, `organization_id`, `parent_id`, `ancestors`, `dept_name`, `order_num`, `leader`, `phone`, `email`, `status`, `is_deleted`, `created_at`, `created_by`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (7, 9, 6, '6,7', 'test', 0, 'aaa', '', '', 1, 1, '2022-09-16 18:30:52', 1, 0, '2022-09-16 18:30:58', 1, '2022-09-16 18:30:58');
COMMIT;

-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_data`;
CREATE TABLE `sys_dict_data` (
  `dict_code` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '字典编码',
  `dict_sort` int(11) DEFAULT '0' COMMENT '字典排序',
  `dict_label` varchar(100) DEFAULT '' COMMENT '字典标签',
  `dict_value` varchar(100) DEFAULT '' COMMENT '字典键值',
  `dict_type` varchar(100) DEFAULT '' COMMENT '字典类型',
  `css_class` varchar(100) DEFAULT NULL COMMENT '样式属性（其他样式扩展）',
  `list_class` varchar(100) DEFAULT NULL COMMENT '表格回显样式',
  `is_default` tinyint(1) DEFAULT '0' COMMENT '是否默认（1是 0否）',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `is_deleted` int(11) NOT NULL COMMENT '是否删除 0未删除 1已删除',
  `create_by` int(10) unsigned DEFAULT '0' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` int(10) unsigned DEFAULT '0' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_by` int(11) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=59 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='字典数据表';

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (1, 0, '333333工aa工工工', '334455555333', 'testtype', '2222', 'stri222ng', 0, 'stri3333ng', 0, 0, 0, '2022-08-10 13:48:13', 0, '2022-08-10 13:51:42', NULL, '2022-08-10 13:53:50');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (2, 0, '6666工aa工工工', '66666', 'testtype', '2222', 'stri222ng', 0, 'stri3333ng', 0, 0, 0, '2022-08-10 13:53:32', 0, '2022-08-10 13:53:32', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (3, 0, 'sex1', '1', 'sex', '', '', 0, '', 1, 0, 1, '2022-08-11 11:41:07', 0, '2022-08-11 11:41:07', NULL, '2022-08-11 14:37:51');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (4, 0, '123', '123', 'sex', '', '', 0, '', 1, 0, 1, '2022-08-11 11:42:15', 0, '2022-08-11 11:42:15', NULL, '2022-08-11 14:37:47');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (5, 0, '男人', '1', 'sex', '', '', 0, '', 1, 0, 1, '2022-08-11 11:43:31', 0, '2022-08-11 11:43:31', NULL, '2022-08-11 14:29:40');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (6, 0, '男人', 'man', 'sex', '', '', 0, 'nanren ', 1, 0, 1, '2022-08-11 11:45:34', 0, '2022-08-11 11:45:34', NULL, '2022-08-11 14:29:35');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (7, 120, 'aaaaaa', '1111111', 'dashifu', '', '', 0, '大地方阿什顿发大师傅阿什顿发阿什顿发阿什顿发', 1, 0, 1, '2022-08-11 14:03:27', 1, '2022-08-11 14:49:23', NULL, '2022-08-11 22:46:21');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (8, 1, 'bbbbb', '222222', 'dashifu', '', '', 0, '', 1, 0, 1, '2022-08-11 14:03:37', 1, '2022-08-11 14:38:48', NULL, '2022-08-11 22:46:21');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (9, 3, '123', '123', 'dashifu', '', '', 1, '', 0, 0, 1, '2022-08-11 14:38:11', 1, '2022-08-11 14:39:29', NULL, '2022-08-11 22:46:21');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (10, 0, '是', '1', 'sys_yes_no', '', '', 1, '', 1, 0, 1, '2022-08-11 14:48:49', 0, '2022-08-11 14:48:49', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (11, 0, '否', '0', 'sys_yes_no', '', '', 0, '', 1, 0, 1, '2022-08-11 14:49:03', 0, '2022-08-11 14:49:03', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (12, 0, '成功', '1', 'admin_login_status', '', '', 0, '', 1, 0, 1, '2022-08-11 22:46:46', 0, '2022-08-11 22:46:46', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (13, 0, '失败', '0', 'admin_login_status', '', '', 0, '', 1, 0, 1, '2022-08-11 22:46:58', 0, '2022-08-11 22:46:58', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (14, 0, '默认', 'DEFAULT', 'sys_job_group', '', '', 1, '', 1, 0, 1, '2022-08-15 09:42:37', 0, '2022-08-15 09:42:37', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (15, 0, '系统', 'SYSTEM', 'sys_job_group', '', '', 1, '', 1, 0, 1, '2022-08-15 09:42:49', 0, '2022-08-15 09:42:49', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (16, 0, '数据中心', 'dataSourceJob', 'sys_job_group', '', '', 0, '', 1, 0, 1, '2022-08-27 20:36:34', 0, '2022-08-27 20:36:34', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (17, 0, '串口', 'serial', 'network_tunnel_type', '', '', 1, '', 1, 0, 1, '2022-09-05 16:32:54', 1, '2022-09-07 17:57:50', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (18, 0, 'TCP客户端', 'tcp-client', 'network_tunnel_type', '', '', 0, '', 1, 0, 1, '2022-09-05 16:33:07', 0, '2022-09-05 16:33:07', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (19, 0, 'TCP服务端', 'tcp-server', 'network_tunnel_type', '', '', 0, '', 1, 0, 1, '2022-09-05 16:33:19', 0, '2022-09-05 16:33:19', NULL, '2022-10-15 10:13:22');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (20, 0, 'UDP客户端', 'udp-client', 'network_tunnel_type', '', '', 0, '', 1, 0, 1, '2022-09-05 16:33:37', 0, '2022-09-05 16:33:37', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (21, 0, 'UDP服务端', 'udp-server', 'network_tunnel_type', '', '', 0, '', 1, 0, 1, '2022-09-05 16:33:51', 1, '2022-09-07 19:58:10', NULL, '2022-10-15 10:13:26');
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (22, 0, 'TCP服务器', 'tcp', 'network_server_type', '', '', 0, '', 1, 0, 1, '2022-09-05 16:38:34', 0, '2022-09-05 16:39:07', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (23, 0, 'UDP服务器', 'udp', 'network_server_type', '', '', 0, '', 1, 0, 1, '2022-09-05 16:38:48', 0, '2022-09-05 16:39:07', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (24, 0, '9600', '9600', 'tunnel_serial_baudrate', '', '', 0, '', 1, 0, 1, '2022-09-08 10:15:22', 0, '2022-09-08 10:15:22', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (25, 0, '14400', '14400', 'tunnel_serial_baudrate', '', '', 0, '', 1, 0, 1, '2022-09-08 10:15:32', 0, '2022-09-08 10:15:32', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (26, 0, '19200', '19200', 'tunnel_serial_baudrate', '', '', 0, '', 1, 0, 1, '2022-09-08 10:15:44', 0, '2022-09-08 10:15:44', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (27, 0, '5', '5', 'tunnel_serial_databits', '', '', 0, '', 1, 0, 1, '2022-09-08 10:18:32', 0, '2022-09-08 10:18:32', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (28, 0, '6', '6', 'tunnel_serial_databits', '', '', 0, '', 1, 0, 1, '2022-09-08 10:18:36', 0, '2022-09-08 10:18:36', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (29, 0, '7', '7', 'tunnel_serial_databits', '', '', 0, '', 1, 0, 1, '2022-09-08 10:18:41', 0, '2022-09-08 10:18:41', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (30, 0, '1', '1', 'tunnel_serial_stopbits', '', '', 0, '', 1, 0, 1, '2022-09-08 10:19:54', 0, '2022-09-08 10:19:54', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (31, 0, '2', '2', 'tunnel_serial_stopbits', '', '', 0, '', 1, 0, 1, '2022-09-08 10:19:58', 0, '2022-09-08 10:19:58', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (32, 0, '无', '0', 'tunnel_serial_parity', '', '', 0, '', 1, 0, 1, '2022-09-08 10:21:23', 0, '2022-09-08 10:21:23', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (33, 0, '奇', '1', 'tunnel_serial_parity', '', '', 0, '', 1, 0, 1, '2022-09-08 10:21:36', 0, '2022-09-08 10:21:36', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (34, 0, '偶', '2', 'tunnel_serial_parity', '', '', 0, '', 1, 0, 1, '2022-09-08 10:21:43', 0, '2022-09-08 10:21:43', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (35, 2, 'Modbus RTU', 'ModbusRTU', 'network_protocols', '', '', 0, '', 1, 0, 1, '2022-09-08 10:34:16', 1, '2022-09-08 10:34:53', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (36, 1, 'Modbus TCP', 'ModbusTCP', 'network_protocols', '', '', 0, '', 1, 0, 1, '2022-09-08 10:34:37', 1, '2022-09-08 10:34:47', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (37, 0, 'MQTT 服务', 'mqtt_server', 'network_server_type', '', '', 0, '', 1, 0, 1, '2022-09-09 10:22:05', 0, '2022-09-09 10:22:05', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (38, 0, 'WT61NP', 'WT61NP', 'network_protocols', '', '', 0, '', 1, 0, 1, '2022-09-17 07:40:24', 0, '2022-09-17 07:40:24', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (39, 0, '系统消息', '1', 'message_scope', '', '', 0, '', 1, 0, 1, '2022-10-09 04:25:10', 0, '2022-10-09 04:25:10', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (40, 0, '组织消息', '2', 'message_scope', '', '', 0, '', 1, 0, 1, '2022-10-09 04:25:18', 0, '2022-10-09 04:25:18', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (41, 0, '部门消息', '3', 'message_scope', '', '', 0, '', 1, 0, 1, '2022-10-09 04:25:26', 0, '2022-10-09 04:25:26', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (42, 0, '用户消息', '4', 'message_scope', '', '', 0, '', 1, 0, 1, '2022-10-09 04:25:33', 0, '2022-10-09 04:25:33', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (43, 0, '天气', '1', 'busi_types', '', '', 0, '', 1, 0, 1, '2022-10-11 10:59:21', 0, '2022-10-11 10:59:21', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (44, 0, 'Sagoo Mqtt', 'SagooMqtt', 'network_protocols', '', '', 1, '', 1, 0, 1, '2022-11-03 16:44:17', 1, '2022-11-03 16:44:42', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (45, 0, '环路监管', '2', 'busi_types', '', '', 0, '', 1, 0, 1, '2022-11-15 04:43:06', 0, '2022-11-15 04:43:06', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (46, 0, '微信', 'wework', 'notice_send_gateway', '', '', 0, '支持腾讯企业消息、服务号消息类型', 1, 0, 1, '2022-11-16 16:33:07', 1, '2022-12-07 21:07:10', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (47, 0, '钉钉', 'dingding', 'notice_send_gateway', '', '', 0, '支持钉钉消息、群机器人消息类型', 1, 0, 1, '2022-11-16 16:33:26', 1, '2022-11-16 16:36:41', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (48, 0, '邮件', 'mail', 'notice_send_gateway', '', '', 0, '支持国内通用和自定义邮件类型', 1, 0, 1, '2022-11-16 16:33:41', 1, '2022-12-06 09:14:44', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (49, 0, '语音', 'voice', 'notice_send_gateway', '', '', 0, '支持语音消息类型', 1, 0, 1, '2022-11-16 16:34:16', 1, '2022-11-16 16:37:12', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (50, 0, '短信', 'sms', 'notice_send_gateway', '', '', 0, '支持短信消息类型', 1, 0, 1, '2022-11-16 16:34:46', 1, '2022-11-16 16:37:29', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (51, 0, 'Webhook', 'webhook', 'notice_send_gateway', '', '', 0, '支持http消息通知', 1, 0, 1, '2022-11-16 16:35:26', 1, '2022-11-16 16:37:41', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (52, 0, '分布图', '3', 'busi_types', '', '', 0, '', 1, 0, 1, '2022-12-08 23:10:56', 0, '2022-12-08 23:10:56', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (53, 0, '能耗分析', '4', 'busi_types', '', '', 0, '', 1, 0, 1, '2022-12-08 23:11:02', 0, '2022-12-08 23:11:02', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (54, 0, 'TGN52', 'tgn52', 'network_protocols', '', '', 0, '', 1, 0, 1, '2022-12-12 15:37:47', 0, '2022-12-12 15:37:47', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (55, 0, 'GET', 'GET', 'api_methods', '', '', 0, '', 1, 0, 1, '2022-12-19 21:15:41', 1, '2022-12-19 21:15:58', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (56, 0, 'POST', 'POST', 'api_methods', '', '', 0, '', 1, 0, 1, '2022-12-19 21:15:52', 0, '2022-12-19 21:15:52', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (57, 0, 'PUT', 'PUT', 'api_methods', '', '', 0, '', 1, 0, 1, '2022-12-19 21:16:05', 0, '2022-12-19 21:16:05', NULL, NULL);
INSERT INTO `sys_dict_data` (`dict_code`, `dict_sort`, `dict_label`, `dict_value`, `dict_type`, `css_class`, `list_class`, `is_default`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (58, 0, 'DELETE', 'DELETE', 'api_methods', '', '', 0, '', 1, 0, 1, '2022-12-19 21:16:18', 0, '2022-12-19 21:16:18', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS `sys_dict_type`;
CREATE TABLE `sys_dict_type` (
  `dict_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '字典主键',
  `parent_id` int(11) NOT NULL COMMENT '父主键ID',
  `dict_name` varchar(100) NOT NULL DEFAULT '' COMMENT '字典名称',
  `dict_type` varchar(100) NOT NULL DEFAULT '' COMMENT '字典类型',
  `module_classify` varchar(255) NOT NULL COMMENT '模块分类',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `is_deleted` int(11) NOT NULL COMMENT '是否删除 0未删除 1已删除',
  `create_by` int(10) unsigned DEFAULT '0' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建日期',
  `update_by` int(10) unsigned DEFAULT '0' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '修改日期',
  `deleted_by` int(11) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`dict_id`) USING BTREE,
  UNIQUE KEY `dict_type` (`dict_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='字典类型表';

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
BEGIN;
INSERT INTO `sys_dict_type` (`dict_id`, `parent_id`, `dict_name`, `dict_type`, `module_classify`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (1, 0, '测试类型2222', 'testtype222', '', '这是一个介绍的内容222', 0, 0, 0, '2022-08-10 13:32:07', 1, '2022-08-11 10:49:01', NULL, '2022-08-11 10:51:14');
INSERT INTO `sys_dict_type` (`dict_id`, `parent_id`, `dict_name`, `dict_type`, `module_classify`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (2, 0, '测试类型55555', 'testtype', '', '这是一个介绍的内容', 0, 0, 0, '2022-08-10 13:34:09', 0, '2022-08-10 13:34:09', NULL, '2022-08-10 13:35:42');
INSERT INTO `sys_dict_type` (`dict_id`, `parent_id`, `dict_name`, `dict_type`, `module_classify`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (3, 0, 'asdf', 'asdf', '', 'asdfasdf', 1, 0, 1, '2022-08-11 10:51:22', 0, '2022-08-11 10:51:22', NULL, '2022-08-11 10:51:28');
INSERT INTO `sys_dict_type` (`dict_id`, `parent_id`, `dict_name`, `dict_type`, `module_classify`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (4, 0, '地点1', 'dashifu', '', '大师傅', 1, 0, 1, '2022-08-11 10:51:36', 1, '2022-08-11 14:03:14', NULL, '2022-08-11 22:46:21');
INSERT INTO `sys_dict_type` (`dict_id`, `parent_id`, `dict_name`, `dict_type`, `module_classify`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (5, 0, '性别', 'sex', '', '', 1, 0, 1, '2022-08-11 11:15:50', 1, '2022-08-11 11:16:13', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `parent_id`, `dict_name`, `dict_type`, `module_classify`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (6, 0, '系统内置', 'sys_yes_no', '', '', 1, 0, 1, '2022-08-11 14:48:02', 0, '2022-08-11 14:48:02', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `parent_id`, `dict_name`, `dict_type`, `module_classify`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (7, 0, '管理员登录状态', 'admin_login_status', '', '', 1, 0, 1, '2022-08-11 22:46:30', 0, '2022-08-11 22:46:30', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `parent_id`, `dict_name`, `dict_type`, `module_classify`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (8, 0, '任务分组', 'sys_job_group', '', '', 1, 0, 1, '2022-08-15 09:42:12', 0, '2022-08-15 09:42:12', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `parent_id`, `dict_name`, `dict_type`, `module_classify`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (9, 0, '网络组件通道类型', 'network_tunnel_type', '', '', 1, 0, 1, '2022-09-05 16:31:53', 0, '2022-09-05 16:31:53', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `parent_id`, `dict_name`, `dict_type`, `module_classify`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (10, 0, '网络组件服务器类型', 'network_server_type', '', '', 1, 0, 1, '2022-09-05 16:37:28', 1, '2022-09-05 16:39:07', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `parent_id`, `dict_name`, `dict_type`, `module_classify`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (11, 0, '通道串口参数波特率', 'tunnel_serial_baudrate', '', '', 1, 0, 1, '2022-09-08 10:14:55', 0, '2022-09-08 10:14:55', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `parent_id`, `dict_name`, `dict_type`, `module_classify`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (12, 0, '通道串口数据位', 'tunnel_serial_databits', '', '', 1, 0, 1, '2022-09-08 10:18:15', 0, '2022-09-08 10:18:15', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `parent_id`, `dict_name`, `dict_type`, `module_classify`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (13, 0, '通道串口停止位', 'tunnel_serial_stopbits', '', '', 1, 0, 1, '2022-09-08 10:19:37', 0, '2022-09-08 10:19:37', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `parent_id`, `dict_name`, `dict_type`, `module_classify`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (14, 0, '通道串口检验位', 'tunnel_serial_parity', '', '', 1, 0, 1, '2022-09-08 10:20:48', 0, '2022-09-08 10:20:48', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `parent_id`, `dict_name`, `dict_type`, `module_classify`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (15, 0, '网络组件通讯协议', 'network_protocols', '', '', 1, 0, 1, '2022-09-08 10:33:28', 0, '2022-09-08 10:33:28', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `parent_id`, `dict_name`, `dict_type`, `module_classify`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (16, 0, '消息范围', 'message_scope', '', '', 1, 0, 1, '2022-10-09 04:24:40', 0, '2022-10-09 04:24:40', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `parent_id`, `dict_name`, `dict_type`, `module_classify`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (17, 0, '业务单元类型', 'busi_types', '', '', 1, 0, 1, '2022-10-11 10:58:59', 0, '2022-10-11 10:58:59', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `parent_id`, `dict_name`, `dict_type`, `module_classify`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (18, 0, '通知发送通道', 'notice_send_gateway', '', '', 1, 0, 1, '2022-11-16 16:30:24', 0, '2022-11-16 16:30:24', NULL, NULL);
INSERT INTO `sys_dict_type` (`dict_id`, `parent_id`, `dict_name`, `dict_type`, `module_classify`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (19, 0, '接口请求方式', 'api_methods', '', '', 1, 0, 1, '2022-12-19 21:15:08', 0, '2022-12-19 21:15:08', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS `sys_job`;
CREATE TABLE `sys_job` (
  `job_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '任务ID',
  `job_name` varchar(64) NOT NULL DEFAULT '' COMMENT '任务名称',
  `job_params` varchar(200) DEFAULT '' COMMENT '参数',
  `job_group` varchar(64) NOT NULL DEFAULT 'DEFAULT' COMMENT '任务组名',
  `invoke_target` varchar(500) NOT NULL COMMENT '调用目标字符串',
  `cron_expression` varchar(255) DEFAULT '' COMMENT 'cron执行表达式',
  `misfire_policy` tinyint(4) DEFAULT '1' COMMENT '计划执行策略（1多次执行 2执行一次）',
  `concurrent` tinyint(4) DEFAULT '1' COMMENT '是否并发执行（0允许 1禁止）',
  `status` tinyint(4) DEFAULT '1' COMMENT '状态（0正常 1暂停）',
  `create_by` bigint(64) unsigned DEFAULT '0' COMMENT '创建者',
  `update_by` bigint(64) unsigned DEFAULT '0' COMMENT '更新者',
  `remark` varchar(500) DEFAULT '' COMMENT '备注信息',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`job_id`,`job_name`,`job_group`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=96 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='定时任务调度表';

-- ----------------------------
-- Records of sys_job
-- ----------------------------
BEGIN;
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (19, 'test2', 'test2', 'DEFAULT', 'test2', 'test2', 1, 0, 1, 1, 0, '', '2022-08-15 10:36:25', '2022-08-15 10:36:25', '2022-08-15 10:36:30');
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (20, 'test2', '1111111', 'SYSTEM', 'test2', '*/5 * * * * ?', 1, 0, 1, 1, 1, '', '2022-08-15 10:36:41', '2022-08-15 20:50:53', '2022-12-19 19:50:38');
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (21, '定时清理操作日志', '3', 'SYSTEM', 'clearOperationLogByDays', '0 0 1 * * ?', 1, 0, 0, 1, 1, '清理3天前的操作日志', '2022-08-15 12:05:56', '2022-12-19 19:45:50', NULL);
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (22, '访问URL', 'http://www.baidu.com', 'DEFAULT', 'getAccessURL', '*/5 * * * * ?', 1, 0, 1, 1, 0, '', '2022-08-15 21:07:27', '2022-08-15 21:07:27', NULL);
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (78, 'dataSource-78', '78', 'dataSourceJob', 'dataSource', '0 0 */1 * * ?', 1, 0, 0, 1, 1, '天气预报', '2022-10-21 00:51:41', '2022-11-03 21:56:32', NULL);
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (79, 'dataSource-79', '79', 'dataSourceJob', 'dataSource', '0 0 */1 * * ?', 1, 0, 0, 1, 0, '天气预报日出与日落', '2022-10-21 01:08:06', '2022-10-21 01:08:06', NULL);
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (90, '设备日志清理', '', 'SYSTEM', 'deviceLogClear', '0 0 2 * * ?', 1, 0, 0, 6, 6, '清理7天前日志', '2022-11-25 11:41:29', '2022-11-25 11:43:05', NULL);
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (93, 'dataTemplate-36', '36', 'dataSourceJob', 'dataTemplate', '0 0 */1 * * ?', 1, 0, 1, 6, 0, '天气预报数据建模', '2022-12-12 15:52:20', '2022-12-12 15:52:20', NULL);
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (94, '告警日志清理', '30', 'SYSTEM', 'clearAlarmLogByDays', '0 0 1 1 * ?', 1, 0, 0, 1, 0, '清理30天前的日志', '2022-12-19 19:48:08', '2022-12-19 19:48:08', NULL);
INSERT INTO `sys_job` (`job_id`, `job_name`, `job_params`, `job_group`, `invoke_target`, `cron_expression`, `misfire_policy`, `concurrent`, `status`, `create_by`, `update_by`, `remark`, `created_at`, `updated_at`, `deleted_at`) VALUES (95, '通知服务日志清理', '10', 'SYSTEM', 'clearNoticeLogByDays', '0 0 3 * * ?', 1, 0, 0, 1, 0, '清理10天前的日志', '2022-12-19 19:50:21', '2022-12-19 19:50:21', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log` (
  `info_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `login_name` varchar(50) DEFAULT '' COMMENT '登录账号',
  `ipaddr` varchar(50) DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) DEFAULT '' COMMENT '操作系统',
  `status` tinyint(4) DEFAULT '0' COMMENT '登录状态（0失败 1成功）',
  `msg` varchar(255) DEFAULT '' COMMENT '提示消息',
  `login_time` datetime DEFAULT NULL COMMENT '登录时间',
  `module` varchar(30) DEFAULT '' COMMENT '登录模块',
  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2472 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='系统访问记录';

-- ----------------------------
-- Records of sys_login_log
-- ----------------------------
BEGIN;
INSERT INTO `sys_login_log` (`info_id`, `login_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`, `module`) VALUES (344, 'admin', '127.0.0.1', '内网IP', 'Chrome', 'Intel Mac OS X 10_15_7', 1, '登录成功', '2022-08-25 01:36:25', '系统后台');
INSERT INTO `sys_login_log` (`info_id`, `login_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`, `module`) VALUES (345, 'admin', '117.22.82.22', '陕西省 西安市', 'Chrome', 'Windows 10', 1, '登录成功', '2022-08-25 02:18:51', '系统后台');
INSERT INTO `sys_login_log` (`info_id`, `login_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`, `module`) VALUES (346, 'admin', '175.0.186.240', '湖南省 长沙市', 'Chrome', 'Windows 10', 1, '登录成功', '2022-08-25 03:27:12', '系统后台');
INSERT INTO `sys_login_log` (`info_id`, `login_name`, `ipaddr`, `login_location`, `browser`, `os`, `status`, `msg`, `login_time`, `module`) VALUES (347, 'admin', '113.232.143.120', '辽宁省 沈阳市', 'Chrome', 'Intel Mac OS X 10_15_7', 1, '登录成功', '2022-08-25 03:30:20', '系统后台');
COMMIT;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父ID',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '规则名称',
  `title` varchar(50) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `icon` varchar(300) NOT NULL DEFAULT '' COMMENT '图标',
  `condition` varchar(255) NOT NULL DEFAULT '' COMMENT '条件',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `menu_type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '类型 0目录 1菜单 2按钮',
  `weigh` int(11) NOT NULL DEFAULT '0' COMMENT '权重',
  `is_hide` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '显示状态',
  `path` varchar(100) NOT NULL DEFAULT '' COMMENT '路由地址',
  `component` varchar(100) NOT NULL DEFAULT '' COMMENT '组件路径',
  `is_link` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否外链 1是 0否',
  `module_type` varchar(30) NOT NULL DEFAULT '' COMMENT '所属模块 system 运维 company企业',
  `model_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '模型ID',
  `is_iframe` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否内嵌iframe',
  `is_cached` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否缓存',
  `redirect` varchar(255) NOT NULL DEFAULT '' COMMENT '路由重定向地址',
  `is_affix` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否固定',
  `link_url` varchar(500) NOT NULL DEFAULT '' COMMENT '链接地址',
  `status` int(11) NOT NULL COMMENT '状态 0 停用 1启用',
  `is_deleted` int(11) NOT NULL COMMENT '是否删除 0未删除 1已删除',
  `created_by` int(10) unsigned DEFAULT '0' COMMENT '创建人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` int(11) DEFAULT NULL COMMENT '修改人',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_by` int(11) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE,
  KEY `pid` (`parent_id`) USING BTREE,
  KEY `weigh` (`weigh`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=119 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='菜单节点表';

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (1, -1, 'api/v1/system/auth', '系统管理', 'iconfont icon-xitongshezhi', '', '', 0, 30, 0, '/system', 'layout/routerView/parent', 0, 'system', 0, 0, 1, '0', 0, '', 1, 0, 1, '2022-08-04 22:13:54', 1, '2022-12-09 23:57:48', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (2, 20, 'api/v1/system/auth/menuList', '菜单管理', 'iconfont icon-juxingkaobei', '', '', 1, 9, 0, '/config/menuList', 'system/menu/index', 0, '', 0, 0, 1, '', 0, '', 1, 0, 1, '2022-08-05 06:15:26', 1, '2022-12-10 00:01:38', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (9, 19, 'api/swagger', 'api文档', 'iconfont icon-jiliandongxuanzeqi', '', '', 1, 0, 0, '/monitor/iframes', 'layout/routerView/iframes', 1, '', 0, 1, 0, '', 0, '/base-api/swagger', 1, 0, 1, '2022-08-04 10:26:43', 1, '2022-12-10 00:00:22', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (10, 1, '角色管理', '角色管理', 'ele-Guide', '', '', 1, 12, 0, '/system/roleList', 'system/manage/role/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-08 06:59:54', 1, '2022-12-09 23:59:06', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (11, 91, '设备管理', '设备管理', 'iconfont icon-siweidaotu', '', '', 0, 300, 0, '/iotmanager/device', '/iot/device', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-04 15:44:01', 1, '2022-12-08 23:04:17', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (12, 1, '部门管理', '部门管理', 'iconfont icon-juxingkaobei', '', '', 1, 920, 0, '/system/deptList', 'system/manage/dept/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-08 20:49:57', 1, '2022-12-09 23:58:50', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (13, 1, '区域组织管理', '区域管理', 'iconfont icon-shuxingtu', '', '', 1, 950, 0, '/system/orgList', 'system/manage/org/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-08 04:51:33', 1, '2022-12-09 23:58:43', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (14, 1, '岗位管理', '岗位管理', 'iconfont icon-gerenzhongxin', '', '', 1, 900, 0, '/system/postList', 'system/manage/post/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-09 04:52:35', 1, '2022-12-09 23:58:58', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (15, 1, '用户管理', '用户管理', 'ele-UserFilled', '', '', 1, 1000, 0, '/system/user', 'system/manage/user/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-09 04:53:23', 1, '2022-12-09 23:58:36', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (16, 20, '系统参数管理', '参数管理', 'ele-Tickets', '', '', 1, 2, 0, '/config/list', 'system/config/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-09 04:54:55', 1, '2022-12-10 00:01:45', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (17, 19, '服务监测', '服务监测', 'iconfont icon-putong', '', '', 1, 0, 0, '/monitor/server', 'system/monitor/server/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-10 06:50:33', 1, '2022-12-10 00:00:31', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (18, 19, '登录日志', '登录日志', 'iconfont icon-neiqianshujuchucun', '', '', 1, 0, 0, '/monitor/loginLog', 'system/monitor/loginLog/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-10 14:53:36', 1, '2022-12-10 00:00:39', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (19, -1, '系统监控', '系统监控', 'iconfont icon-neiqianshujuchucun', '', '', 0, 0, 0, '/monitor', 'layout/routerView/parent', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-10 22:55:17', 1, '2022-12-09 23:59:48', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (20, -1, '系统配置', '系统配置', 'ele-SetUp', '', '', 0, 0, 0, '/config', 'layout/routerView/parent', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-10 07:01:59', 1, '2022-12-09 23:59:57', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (21, 20, '字典管理', '字典管理', 'iconfont icon--chaifenlie', '', '', 1, 0, 0, '/config/dict', 'system/dict/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-09 15:04:21', 1, '2022-12-10 00:02:00', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (22, 20, '字典数据管理', '字典数据管理', 'ele-Collection', '', '', 1, 0, 1, '/config/dict/:dictType', 'system/dict/dataList', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-08 02:21:34', 1, '2022-12-10 00:02:25', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (27, 20, '数据中心', '数据中心', 'ele-Coin', '', '', 0, 0, 0, '/config/datahub', '/system/datahub', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-09 08:02:39', 1, '2022-12-10 00:03:36', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (34, 11, '产品', '产品', 'ele-CreditCard', '', '', 1, 100, 0, '/iotmanager/device/product', '/iot/device/product/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-07 09:09:30', 1, '2022-12-08 23:04:35', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (35, 11, '产品分类', '产品分类', 'ele-DocumentCopy', '', '', 1, 0, 0, '/iotmanager/device/category', '/iot/device/category/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-11 01:10:38', 1, '2022-12-08 23:04:50', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (36, 11, '设备', '设备', 'iconfont icon-diannao', '', '', 1, 80, 0, '/iotmanager/device/instance', '/iot/device/instance/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-11 01:11:24', 1, '2022-12-08 23:04:43', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (37, 27, '数据源管理', '数据源管理', 'ele-Connection', '', '', 1, 0, 0, '/config/datahub/source', '/system/datahub/source/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-11 01:11:36', 1, '2022-12-10 00:04:16', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (38, 27, '数据管理', '数据管理', 'ele-DocumentCopy', '', '', 1, 0, 0, '/datahub/datas', '/datahub/datas', 0, '', 0, 0, 0, '', 0, '', 1, 1, 1, '2022-08-12 17:14:56', 0, '2022-08-27 20:24:44', 1, '2022-08-27 20:24:44');
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (39, 27, '数据建模', '数据建模', 'iconfont icon-juxingkaobei', '', '', 1, 0, 0, '/config/datahub/modeling', '/system/datahub/modeling/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-11 20:54:33', 1, '2022-12-10 00:04:23', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (40, 91, '通知服务', '通知服务', 'iconfont icon-dongtai', '', '', 0, 80, 0, '/iotmanager/noticeservices', '/noticeservices', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-11 20:57:26', 1, '2022-12-09 23:51:54', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (41, 40, '通知配置', '通知配置', 'iconfont icon-jiliandongxuanzeqi', '', '', 1, 0, 0, '/iotmanager/noticeservices/config', '/iot/noticeservices/config/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-11 21:01:00', 1, '2022-12-09 23:52:18', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (42, 40, '通知日志', '通知日志', 'ele-Postcard', '', '', 1, 0, 0, '/iotmanager/noticeservices/log', '/iot/noticeservices/log/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-11 05:02:15', 1, '2022-12-09 23:52:33', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (45, 43, '模版市场', '模版市场', 'iconfont icon-crew_feature', '', '', 1, 0, 1, '/bigscreen/template-market', '	 layout/routerView/link', 1, '', 0, 0, 0, '', 0, 'http://home.yanglizhi.cn:10003/#/project/template-market', 1, 1, 1, '2022-08-11 13:22:39', 1, '2022-09-17 21:28:31', 1, '2022-09-17 21:28:31');
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (46, 20, '定时任务', '定时任务', 'ele-Money', '', '', 1, 0, 0, '/config/task', '/system/task/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-12 05:59:18', 1, '2022-12-10 00:11:30', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (47, -1, '1', '1', '', '', '', 0, 0, 0, '1', '1', 0, '', 0, 0, 0, '', 0, '', 1, 1, 1, '2022-08-13 22:08:08', 0, '2022-08-13 22:08:11', 1, '2022-08-13 22:08:11');
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (48, 20, '接口管理', '接口管理', 'iconfont icon-zujian', '', '', 1, 0, 0, '/config/api', '/system/api/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-13 06:43:34', 1, '2022-12-10 00:11:38', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (49, 19, 'operLog', '操作日志', 'iconfont icon-LoggedinPC', '', '', 1, 0, 0, '/monitor/operLog', '/system/monitor/operLog/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-14 20:25:37', 1, '2022-12-10 00:00:50', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (50, 19, 'online', '在线用户', 'iconfont icon-diannao-shuju', '', '', 1, 0, 0, '/monitor/online', '/system/monitor/online/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-15 06:40:43', 1, '2022-12-10 00:00:58', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (51, 36, '设备详情', '设备详情', 'iconfont icon-biaodan', '', '', 1, 0, 1, '/iotmanager/device/instance/:id', '/iot/device/instance/detail', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-12 12:12:59', 1, '2022-12-08 23:09:21', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (52, 34, '产品详情', '产品详情', 'iconfont icon-biaodan', '', '', 1, 0, 1, '/iotmanager/device/product/detail/:id', '/iot/device/product/detail', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-13 21:33:48', 1, '2022-12-08 23:10:27', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (61, 27, '数据源详情', '数据源详情', '', '', '', 1, 0, 1, '/config/datahub/source/:sourceId', '/system/datahub/source/detail', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-26 07:30:33', 1, '2022-12-10 00:05:11', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (64, 27, '数据模型详情', '数据模型详情', '', '', '', 1, 0, 1, '/config/datahub/modeling/:id', '/system/datahub/modeling/detail', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-26 19:37:24', 1, '2022-12-10 00:05:52', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (65, 91, '网络组件管理', '网络组件', 'ele-Guide', '', '', 0, 200, 0, '/iotmanager/network', '/iot/network', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-30 02:53:26', 1, '2022-12-09 23:40:36', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (66, 65, '通道管理', '通道管理', 'ele-ScaleToOriginal', '', '', 1, 0, 0, '/iotmanager/network/tunnel', '/iot/network/tunnel/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-08-31 02:55:36', 1, '2022-12-09 23:40:55', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (67, -1, 'tunnulManage', '通道管理', '', '', '', 1, 0, 0, '/tunnulManage/list', '/tunnulManage/index', 0, '', 0, 0, 0, '', 0, '', 1, 1, 1, '2022-09-01 09:21:05', 0, '2022-09-01 09:25:53', 1, '2022-09-01 09:25:53');
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (68, 66, '通道详情', '通道详情', '', '', '', 1, 0, 1, '/iotmanager/network/tunnel/detail/:id', '/iot/network/tunnel/detail', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-09-01 07:05:44', 1, '2022-12-09 23:43:40', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (69, 66, '新增通道', '新增通道', '', '', '', 1, 0, 1, '/iotmanager/network/tunnel/create', '/iot/network/tunnel/create', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-09-04 18:10:41', 1, '2022-12-09 23:43:48', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (70, 65, '服务器管理', '服务器管理', 'ele-SetUp', '', '', 1, 0, 0, '/iotmanager/network/server', '/iot/network/server/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-09-04 02:31:20', 1, '2022-12-09 23:41:05', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (71, 70, '服务详情', '服务详情', '', '', '', 1, 0, 1, '/iotmanager/network/server/detail/:id', '/iot/network/server/detail', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-09-05 10:36:11', 1, '2022-12-09 23:46:49', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (72, 70, '新增服务', '新增服务', '', '', '', 1, 0, 1, '/iotmanager/network/server/create', '/iot/network/server/create', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-09-06 16:21:10', 1, '2022-12-09 23:46:56', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (73, 70, '编辑服务', '编辑服务', '', '', '', 1, 0, 1, '/iotmanager/network/server/edit/:id', '/iot/network/server/edit', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-09-06 20:46:52', 1, '2022-12-09 23:47:02', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (74, 66, '编辑通道', '编辑通道', '', '', '', 1, 0, 1, '/iotmanager/network/tunnel/edit/:id', '/iot/network/tunnel/edit', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-09-06 20:48:19', 1, '2022-12-09 23:43:55', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (80, 91, '告警中心', '告警中心', 'iconfont icon-tongzhi2', '', '', 0, 0, 0, '/iotmanager/alarm', '/iot/alarm', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-09-11 07:50:06', 1, '2022-12-09 23:54:21', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (81, 27, '数据源数据列表', '数据源数据列表', '', '', '', 1, 0, 1, '/config/datahub/source/:id', '/system/datahub/source/list', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-09-15 10:13:26', 1, '2022-12-10 00:07:54', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (82, -1, 'dataOverview', '数据概览', 'iconfont icon-shouye', '', '', 0, 458, 0, '/dataOverview', 'dataOverview/index', 0, '', 0, 0, 0, '', 0, '', 1, 1, 1, '2022-09-19 13:40:42', 1, '2022-10-09 14:18:34', 1, '2022-10-09 14:18:34');
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (83, 82, '平台数据概览', '平台数据概览', 'iconfont icon-ico_shuju', '', '', 1, 0, 0, '/dataOverview', 'dataOverview/index', 0, '', 0, 0, 0, '', 0, '', 1, 1, 1, '2022-09-20 05:43:57', 1, '2022-10-09 14:18:29', 1, '2022-10-09 14:18:29');
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (84, 56, '天气监测', '天气监测', 'ele-Drizzling', '', '', 1, 0, 0, '/heating-monitor/weather', '/heating/monitor/weather', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-09-21 12:50:32', 1, '2022-12-08 22:57:15', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (85, 20, '城市管理', '城市管理', 'ele-DeleteLocation', '', '', 1, 0, 0, '/config/city', 'system/city/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-09-25 18:41:55', 1, '2022-12-10 00:11:45', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (87, 86, '功能列表', '功能列表', 'ele-SetUp', '', '', 1, 0, 0, '/code/list', '/code/list', 0, '', 0, 0, 0, '', 0, '', 1, 1, 1, '2022-09-28 14:32:48', 1, '2022-09-29 10:35:03', 1, '2022-09-29 10:35:03');
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (91, -1, 'iotmanager', '物联管理', 'ele-UploadFilled', '', '', 0, 100, 0, '/iotmanager', '/iotmanager', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-10-08 23:11:42', 1, '2022-10-09 07:15:53', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (92, 91, 'iotmanager-dashboard', '物联概览', 'iconfont icon-diannao-shuju', '', '', 1, 1000, 0, '/iotmanager/dashboard', '/iot/iotmanager/dashboard', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-10-07 23:31:31', 1, '2022-11-28 10:19:28', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (110, 19, 'plugin', '插件监控', 'iconfont icon-diannao', '', '', 1, 0, 0, '/monitor/plugin', '/system/monitor/plugin/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-10-28 17:01:38', 1, '2022-12-10 00:01:05', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (114, 80, '告警配置', '告警配置', 'iconfont icon-shuju', '', '', 1, 0, 0, '/iotmanager/alarm/setting', '/iot/alarm/setting/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-11-10 23:14:24', 1, '2022-12-09 23:54:34', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (115, 80, '告警日志', '告警日志', 'iconfont icon-biaodan', '', '', 1, 0, 0, '/iotmanager/alarm/log', '/iot/alarm/log/index', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-11-11 23:33:09', 1, '2022-12-09 23:54:42', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (116, 41, '通知配置管理', '通知配置管理', '', '', '', 1, 0, 1, '/iotmanager/noticeservices/config/setting/:id', '/iot/noticeservices/config/setting', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-11-15 05:59:08', 1, '2022-12-09 23:53:00', 0, NULL);
INSERT INTO `sys_menu` (`id`, `parent_id`, `name`, `title`, `icon`, `condition`, `remark`, `menu_type`, `weigh`, `is_hide`, `path`, `component`, `is_link`, `module_type`, `model_id`, `is_iframe`, `is_cached`, `redirect`, `is_affix`, `link_url`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (117, 41, '通知模板管理', '通知模板管理', '', '', '', 1, 0, 1, '/iotmanager/noticeservices/config/template/:id', '/iot/noticeservices/config/template', 0, '', 0, 0, 0, '', 0, '', 1, 0, 1, '2022-11-14 22:01:07', 1, '2022-12-09 23:53:07', 0, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_menu_api
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu_api`;
CREATE TABLE `sys_menu_api` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `menu_id` int(11) NOT NULL COMMENT '菜单ID',
  `api_id` int(11) NOT NULL COMMENT 'apiId',
  `is_deleted` int(11) NOT NULL COMMENT '是否删除 0未删除 1已删除',
  `created_by` int(10) unsigned DEFAULT '0' COMMENT '创建人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `deleted_by` int(11) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=390 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='菜单与接口关系表';

-- ----------------------------
-- Records of sys_menu_api
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (1, 34, 1, 1, 1, '2022-08-14 23:56:20', 1, '2022-08-15 00:13:02');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (2, 35, 1, 1, 1, '2022-08-14 23:56:20', 1, '2022-08-15 00:13:02');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (3, 36, 1, 1, 1, '2022-08-14 23:56:20', 1, '2022-08-15 00:13:02');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (4, 34, 1, 1, 1, '2022-08-15 00:13:02', 1, '2022-08-15 08:29:18');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (5, 35, 1, 1, 1, '2022-08-15 00:13:02', 1, '2022-08-15 08:29:18');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (6, 36, 1, 1, 1, '2022-08-15 00:13:02', 1, '2022-08-15 08:29:18');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (7, 9, 1, 1, 1, '2022-08-15 08:29:18', 1, '2022-08-15 08:29:42');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (8, 17, 1, 1, 1, '2022-08-15 08:29:18', 1, '2022-08-15 08:29:42');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (9, 18, 1, 1, 1, '2022-08-15 08:29:18', 1, '2022-08-15 08:29:42');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (10, 24, 1, 1, 1, '2022-08-15 08:29:42', 1, '2022-08-15 09:25:57');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (11, 25, 1, 1, 1, '2022-08-15 08:29:42', 1, '2022-08-15 09:25:57');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (12, 24, 2, 1, 1, '2022-08-15 08:55:41', 1, '2022-08-15 10:13:26');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (13, 25, 2, 1, 1, '2022-08-15 08:55:41', 1, '2022-08-15 10:13:26');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (14, 30, 1, 1, 1, '2022-08-15 09:25:57', 1, '2022-08-15 10:27:56');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (15, 31, 1, 1, 1, '2022-08-15 09:25:57', 1, '2022-08-15 10:27:56');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (16, 29, 1, 1, 1, '2022-08-15 09:25:57', 1, '2022-08-15 10:27:56');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (17, 37, 3, 1, 1, '2022-08-15 09:26:17', 1, '2022-08-15 02:27:51');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (18, 38, 3, 1, 1, '2022-08-15 09:26:17', 1, '2022-08-15 02:27:51');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (19, 39, 3, 1, 1, '2022-08-15 09:26:17', 1, '2022-08-15 02:27:51');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (20, 24, 2, 1, 1, '2022-08-15 10:13:26', 1, '2022-08-15 10:13:27');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (21, 25, 2, 1, 1, '2022-08-15 10:13:26', 1, '2022-08-15 10:13:27');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (22, 24, 2, 1, 1, '2022-08-15 10:13:27', 1, '2022-08-15 02:18:21');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (23, 25, 2, 1, 1, '2022-08-15 10:13:27', 1, '2022-08-15 02:18:21');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (24, 30, 1, 1, 1, '2022-08-15 10:27:56', 1, '2022-08-15 10:34:20');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (25, 31, 1, 1, 1, '2022-08-15 10:27:56', 1, '2022-08-15 10:34:20');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (26, 29, 1, 1, 1, '2022-08-15 10:27:56', 1, '2022-08-15 10:34:20');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (27, 34, 1, 1, 1, '2022-08-15 10:34:21', 1, '2022-08-15 10:36:12');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (28, 35, 1, 1, 1, '2022-08-15 10:34:21', 1, '2022-08-15 10:36:12');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (29, 36, 1, 1, 1, '2022-08-15 10:34:21', 1, '2022-08-15 10:36:12');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (30, 34, 1, 1, 0, '2022-08-15 10:36:12', 1, '2022-08-15 10:36:21');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (31, 35, 1, 1, 0, '2022-08-15 10:36:12', 1, '2022-08-15 10:36:21');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (32, 36, 1, 1, 0, '2022-08-15 10:36:12', 1, '2022-08-15 10:36:21');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (33, 34, 1, 1, 0, '2022-08-15 10:36:21', 1, '2022-08-15 13:44:16');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (34, 35, 1, 1, 0, '2022-08-15 10:36:21', 1, '2022-08-15 13:44:16');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (35, 36, 1, 1, 0, '2022-08-15 10:36:21', 1, '2022-08-15 13:44:16');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (36, 34, 1, 1, 1, '2022-08-15 13:44:17', 1, '2022-08-15 13:44:26');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (37, 35, 1, 1, 1, '2022-08-15 13:44:17', 1, '2022-08-15 13:44:26');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (38, 36, 1, 1, 1, '2022-08-15 13:44:17', 1, '2022-08-15 13:44:26');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (39, 34, 1, 1, 1, '2022-08-15 13:44:26', 1, '2022-08-15 13:44:39');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (40, 35, 1, 1, 1, '2022-08-15 13:44:26', 1, '2022-08-15 13:44:39');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (41, 36, 1, 1, 1, '2022-08-15 13:44:26', 1, '2022-08-15 13:44:39');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (42, 34, 1, 1, 1, '2022-08-15 13:44:40', 1, '2022-08-15 13:45:07');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (43, 35, 1, 1, 1, '2022-08-15 13:44:40', 1, '2022-08-15 13:45:07');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (44, 36, 1, 1, 1, '2022-08-15 13:44:40', 1, '2022-08-15 13:45:07');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (45, 34, 1, 1, 1, '2022-08-15 13:45:08', 1, '2022-08-15 13:46:08');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (46, 35, 1, 1, 1, '2022-08-15 13:45:08', 1, '2022-08-15 13:46:08');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (47, 36, 1, 1, 1, '2022-08-15 13:45:08', 1, '2022-08-15 13:46:08');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (48, 34, 1, 1, 1, '2022-08-15 13:46:08', 1, '2022-08-17 22:37:06');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (49, 35, 1, 1, 1, '2022-08-15 13:46:08', 1, '2022-08-17 22:37:06');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (50, 36, 1, 1, 1, '2022-08-15 13:46:08', 1, '2022-08-17 22:37:06');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (51, 52, 1, 1, 1, '2022-08-17 22:37:06', 1, '2022-08-17 22:37:18');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (52, 52, 1, 0, 1, '2022-08-17 22:37:18', 1, '2022-11-03 03:19:25');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (53, 35, 4, 1, 1, '2022-08-17 22:37:28', 1, '2022-08-17 22:52:39');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (54, 35, 4, 0, 1, '2022-08-17 22:52:39', 1, '2022-11-03 03:19:28');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (55, 36, 4, 0, 1, '2022-08-17 22:52:39', 1, '2022-11-03 03:19:28');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (56, 57, 5, 0, 1, '2022-11-03 00:35:21', 1, '2022-11-03 03:19:30');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (57, 2, 6, 1, 1, '2022-11-03 11:25:16', 1, '2022-11-04 23:56:48');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (58, 57, 6, 0, 1, '2022-11-04 23:56:48', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (59, 57, 7, 1, 1, '2022-11-04 23:57:53', 1, '2022-11-07 20:57:00');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (60, 57, 7, 0, 1, '2022-11-07 20:57:00', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (61, 54, 7, 0, 1, '2022-11-07 20:57:00', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (62, 54, 8, 0, 1, '2022-11-07 20:57:37', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (63, 54, 9, 0, 1, '2022-11-07 20:58:15', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (64, 54, 10, 0, 1, '2022-11-07 20:58:36', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (65, 54, 11, 0, 1, '2022-11-07 20:59:02', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (66, 55, 12, 1, 1, '2022-11-07 21:00:29', 1, '2022-11-07 21:20:14');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (67, 55, 13, 0, 1, '2022-11-07 21:01:01', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (68, 55, 14, 0, 1, '2022-11-07 21:01:20', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (69, 55, 15, 0, 1, '2022-11-07 21:01:37', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (70, 55, 16, 0, 1, '2022-11-07 21:01:55', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (71, 100, 17, 0, 1, '2022-11-07 21:02:43', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (72, 100, 18, 0, 1, '2022-11-07 21:03:11', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (73, 100, 19, 0, 1, '2022-11-07 21:03:32', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (74, 100, 20, 0, 1, '2022-11-07 21:03:50', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (75, 100, 21, 0, 1, '2022-11-07 21:04:20', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (76, 100, 22, 1, 1, '2022-11-07 21:04:47', 1, '2022-11-09 23:26:33');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (77, 100, 23, 0, 1, '2022-11-07 21:05:12', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (78, 100, 24, 0, 1, '2022-11-07 21:05:28', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (79, 100, 25, 0, 1, '2022-11-07 21:05:48', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (80, 100, 26, 0, 1, '2022-11-07 21:06:02', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (81, 100, 27, 0, 1, '2022-11-07 21:06:22', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (82, 100, 28, 0, 1, '2022-11-07 21:06:49', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (83, 100, 29, 0, 1, '2022-11-07 21:07:03', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (84, 100, 30, 0, 1, '2022-11-07 21:07:24', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (85, 100, 31, 0, 1, '2022-11-07 21:07:39', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (86, 100, 32, 0, 1, '2022-11-07 21:07:56', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (87, 100, 33, 0, 1, '2022-11-07 21:08:11', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (88, 100, 34, 0, 1, '2022-11-07 21:08:25', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (89, 100, 35, 0, 1, '2022-11-07 21:08:45', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (90, 100, 36, 0, 1, '2022-11-07 21:09:14', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (91, 100, 37, 0, 1, '2022-11-07 21:10:37', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (92, 100, 38, 0, 1, '2022-11-07 21:10:52', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (93, 100, 39, 0, 1, '2022-11-07 21:11:08', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (94, 100, 40, 0, 1, '2022-11-07 21:11:21', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (95, 100, 41, 0, 1, '2022-11-07 21:11:39', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (96, 100, 42, 0, 1, '2022-11-07 21:12:44', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (97, 100, 43, 0, 1, '2022-11-07 21:13:01', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (98, 100, 44, 0, 1, '2022-11-07 21:13:15', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (99, 100, 45, 0, 1, '2022-11-07 21:13:31', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (100, 100, 46, 0, 1, '2022-11-07 21:13:49', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (101, 100, 47, 0, 1, '2022-11-07 21:15:49', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (102, 55, 12, 1, 1, '2022-11-07 21:20:14', 1, '2022-11-07 21:20:46');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (103, 77, 12, 1, 1, '2022-11-07 21:20:14', 1, '2022-11-07 21:20:46');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (104, 55, 12, 0, 1, '2022-11-07 21:20:46', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (105, 77, 12, 0, 1, '2022-11-07 21:20:46', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (106, 78, 12, 0, 1, '2022-11-07 21:20:46', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (107, 76, 48, 1, 1, '2022-11-07 21:23:39', 1, '2022-11-09 23:05:44');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (108, 77, 49, 0, 1, '2022-11-07 21:23:51', 1, '2022-11-08 12:53:34');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (109, 78, 50, 0, 1, '2022-11-07 21:24:04', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (110, 78, 51, 0, 1, '2022-11-07 21:24:17', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (111, 113, 52, 0, 1, '2022-11-09 23:00:06', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (112, 76, 48, 0, 1, '2022-11-09 23:05:44', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (113, 113, 48, 0, 1, '2022-11-09 23:05:44', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (114, 113, 53, 0, 1, '2022-11-09 23:06:39', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (115, 113, 54, 0, 1, '2022-11-09 23:07:16', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (116, 84, 55, 0, 1, '2022-11-09 23:09:21', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (117, 84, 56, 0, 1, '2022-11-09 23:11:05', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (118, 84, 57, 0, 1, '2022-11-09 23:11:34', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (119, 84, 58, 0, 1, '2022-11-09 23:11:58', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (120, 104, 59, 1, 1, '2022-11-09 23:18:02', 1, '2022-11-09 23:28:13');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (121, 105, 59, 1, 1, '2022-11-09 23:18:02', 1, '2022-11-09 23:28:13');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (122, 106, 59, 1, 1, '2022-11-09 23:18:02', 1, '2022-11-09 23:28:13');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (123, 107, 59, 1, 1, '2022-11-09 23:18:02', 1, '2022-11-09 23:28:13');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (124, 100, 22, 1, 1, '2022-11-09 23:26:33', 1, '2022-11-09 23:27:24');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (125, 92, 22, 1, 1, '2022-11-09 23:26:33', 1, '2022-11-09 23:27:24');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (126, 100, 22, 0, 1, '2022-11-09 23:27:24', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (127, 104, 59, 0, 1, '2022-11-09 23:28:13', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (128, 105, 59, 0, 1, '2022-11-09 23:28:13', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (129, 106, 59, 0, 1, '2022-11-09 23:28:13', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (130, 107, 59, 0, 1, '2022-11-09 23:28:13', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (131, 34, 59, 0, 1, '2022-11-09 23:28:13', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (132, 34, 60, 0, 1, '2022-11-09 23:29:19', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (133, 35, 60, 0, 1, '2022-11-09 23:29:19', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (134, 34, 61, 1, 1, '2022-11-09 23:30:37', 1, '2022-11-10 22:47:24');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (135, 24, 62, 1, 1, '2022-11-09 23:37:55', 1, '2022-11-09 23:40:23');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (136, 24, 63, 0, 1, '2022-11-09 23:38:37', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (137, 24, 64, 1, 1, '2022-11-09 23:39:00', 1, '2022-11-09 23:40:16');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (138, 25, 64, 0, 1, '2022-11-09 23:40:16', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (139, 25, 62, 0, 1, '2022-11-09 23:40:23', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (140, 25, 65, 0, 1, '2022-11-09 23:43:04', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (141, 25, 66, 0, 1, '2022-11-09 23:43:27', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (142, 25, 67, 0, 1, '2022-11-09 23:43:53', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (143, 25, 68, 0, 1, '2022-11-09 23:45:04', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (144, 25, 69, 0, 1, '2022-11-09 23:45:22', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (145, 25, 70, 0, 1, '2022-11-09 23:45:42', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (146, 25, 71, 1, 1, '2022-11-09 23:46:48', 1, '2022-11-09 23:47:27');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (147, 25, 71, 0, 1, '2022-11-09 23:47:27', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (148, 25, 72, 0, 1, '2022-11-09 23:47:53', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (149, 25, 73, 0, 1, '2022-11-09 23:48:14', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (150, 44, 74, 0, 1, '2022-11-09 23:49:20', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (151, 44, 75, 0, 1, '2022-11-09 23:49:46', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (152, 44, 76, 0, 1, '2022-11-09 23:50:04', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (153, 44, 77, 0, 1, '2022-11-09 23:50:28', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (154, 44, 78, 0, 1, '2022-11-09 23:51:03', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (155, 44, 79, 0, 1, '2022-11-09 23:51:27', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (156, 44, 80, 0, 1, '2022-11-09 23:51:47', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (157, 89, 81, 0, 1, '2022-11-10 00:00:30', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (158, 90, 81, 0, 1, '2022-11-10 00:00:30', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (159, 89, 82, 0, 1, '2022-11-10 00:11:46', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (160, 90, 82, 0, 1, '2022-11-10 00:11:46', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (161, 15, 83, 0, 1, '2022-11-10 00:13:04', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (162, 15, 84, 0, 1, '2022-11-10 00:13:39', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (163, 15, 85, 0, 1, '2022-11-10 00:14:14', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (164, 15, 86, 0, 1, '2022-11-10 00:14:34', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (165, 15, 87, 0, 1, '2022-11-10 00:15:04', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (166, 15, 88, 0, 1, '2022-11-10 00:15:39', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (167, 15, 89, 0, 1, '2022-11-10 00:16:04', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (168, 37, 91, 0, 1, '2022-11-10 11:13:53', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (169, 35, 92, 0, 1, '2022-11-10 11:39:54', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (170, 35, 93, 0, 1, '2022-11-10 11:40:25', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (171, 35, 94, 0, 1, '2022-11-10 11:41:08', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (172, 113, 95, 0, 1, '2022-11-10 12:03:49', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (173, 34, 96, 0, 1, '2022-11-10 13:01:10', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (174, 34, 97, 0, 1, '2022-11-10 13:01:37', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (175, 34, 98, 0, 1, '2022-11-10 13:01:54', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (176, 34, 99, 0, 1, '2022-11-10 13:02:10', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (177, 34, 100, 0, 1, '2022-11-10 13:02:24', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (178, 34, 101, 0, 1, '2022-11-10 13:02:41', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (179, 34, 102, 0, 1, '2022-11-10 13:02:57', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (180, 34, 103, 0, 1, '2022-11-10 13:03:17', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (181, 34, 104, 0, 1, '2022-11-10 13:03:43', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (182, 34, 105, 0, 1, '2022-11-10 13:04:19', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (183, 36, 106, 0, 1, '2022-11-10 13:04:35', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (184, 36, 107, 0, 1, '2022-11-10 13:04:58', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (185, 36, 108, 0, 1, '2022-11-10 13:05:19', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (186, 36, 109, 0, 1, '2022-11-10 13:05:34', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (187, 36, 110, 0, 1, '2022-11-10 13:06:13', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (188, 36, 111, 0, 1, '2022-11-10 13:06:30', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (189, 51, 112, 0, 1, '2022-11-10 13:06:54', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (190, 51, 113, 0, 1, '2022-11-10 13:07:11', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (191, 51, 114, 0, 1, '2022-11-10 13:07:26', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (192, 51, 115, 0, 1, '2022-11-10 13:07:43', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (193, 51, 116, 0, 1, '2022-11-10 13:08:00', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (194, 36, 117, 0, 1, '2022-11-10 13:08:18', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (195, 36, 118, 0, 1, '2022-11-10 13:08:37', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (196, 36, 119, 0, 1, '2022-11-10 13:09:01', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (197, 51, 120, 0, 1, '2022-11-10 13:09:33', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (198, 51, 121, 0, 1, '2022-11-10 13:09:49', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (199, 36, 122, 0, 1, '2022-11-10 13:10:39', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (200, 51, 123, 0, 1, '2022-11-10 13:10:55', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (201, 51, 124, 0, 1, '2022-11-10 13:11:09', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (202, 51, 125, 0, 1, '2022-11-10 13:11:30', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (203, 51, 126, 0, 1, '2022-11-10 13:11:47', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (204, 51, 127, 0, 1, '2022-11-10 13:12:01', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (205, 51, 128, 0, 1, '2022-11-10 13:12:13', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (206, 51, 129, 0, 1, '2022-11-10 13:12:26', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (207, 51, 130, 0, 1, '2022-11-10 13:12:44', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (208, 51, 131, 0, 1, '2022-11-10 13:13:01', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (209, 51, 132, 0, 1, '2022-11-10 13:13:13', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (210, 51, 133, 0, 1, '2022-11-10 13:13:25', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (211, 51, 134, 0, 1, '2022-11-10 13:13:42', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (212, 51, 135, 0, 1, '2022-11-10 13:13:59', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (213, 51, 136, 0, 1, '2022-11-10 13:14:13', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (214, 51, 137, 0, 1, '2022-11-10 13:14:28', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (215, 51, 138, 0, 1, '2022-11-10 13:14:43', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (216, 51, 139, 0, 1, '2022-11-10 13:14:58', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (217, 37, 140, 0, 1, '2022-11-10 13:16:03', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (218, 37, 141, 0, 1, '2022-11-10 13:16:16', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (219, 61, 142, 0, 1, '2022-11-10 13:16:42', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (220, 61, 143, 0, 1, '2022-11-10 13:16:59', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (221, 37, 144, 0, 1, '2022-11-10 13:17:21', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (222, 37, 145, 0, 1, '2022-11-10 13:17:37', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (223, 37, 146, 0, 1, '2022-11-10 13:18:02', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (224, 37, 147, 0, 1, '2022-11-10 13:18:17', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (225, 37, 148, 0, 1, '2022-11-10 13:18:30', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (226, 37, 149, 0, 1, '2022-11-10 13:18:47', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (227, 37, 150, 0, 1, '2022-11-10 13:19:00', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (228, 37, 151, 0, 1, '2022-11-10 13:19:28', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (229, 37, 152, 0, 1, '2022-11-10 13:19:43', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (230, 61, 153, 0, 1, '2022-11-10 13:20:00', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (231, 61, 154, 0, 1, '2022-11-10 13:20:15', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (232, 61, 155, 0, 1, '2022-11-10 13:20:39', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (233, 37, 156, 0, 1, '2022-11-10 13:20:55', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (234, 61, 157, 0, 1, '2022-11-10 13:21:16', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (235, 61, 158, 0, 1, '2022-11-10 13:21:31', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (236, 61, 159, 0, 1, '2022-11-10 13:21:48', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (237, 61, 160, 0, 1, '2022-11-10 13:22:07', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (238, 39, 161, 0, 1, '2022-11-10 13:22:52', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (239, 39, 162, 0, 1, '2022-11-10 13:23:06', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (240, 39, 163, 0, 1, '2022-11-10 13:23:20', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (241, 39, 164, 0, 1, '2022-11-10 13:23:35', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (242, 39, 165, 0, 1, '2022-11-10 13:23:48', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (243, 64, 166, 0, 1, '2022-11-10 13:24:08', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (244, 39, 167, 0, 1, '2022-11-10 13:24:32', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (245, 64, 168, 0, 1, '2022-11-10 13:24:47', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (246, 64, 169, 0, 1, '2022-11-10 13:25:10', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (247, 39, 170, 0, 1, '2022-11-10 13:25:24', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (248, 64, 171, 0, 1, '2022-11-10 13:25:38', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (249, 64, 172, 0, 1, '2022-11-10 13:25:55', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (250, 64, 173, 0, 1, '2022-11-10 13:26:11', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (251, 64, 174, 0, 1, '2022-11-10 13:26:31', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (252, 13, 175, 0, 1, '2022-11-10 22:38:27', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (253, 13, 176, 0, 1, '2022-11-10 22:39:19', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (254, 13, 177, 0, 1, '2022-11-10 22:39:38', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (255, 13, 178, 0, 1, '2022-11-10 22:39:57', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (256, 13, 179, 0, 1, '2022-11-10 22:40:17', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (257, 12, 180, 0, 1, '2022-11-10 22:40:50', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (258, 12, 181, 0, 1, '2022-11-10 22:41:18', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (259, 12, 182, 0, 1, '2022-11-10 22:41:38', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (260, 12, 183, 0, 1, '2022-11-10 22:42:03', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (261, 34, 61, 0, 1, '2022-11-10 22:47:24', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (262, 12, 61, 0, 1, '2022-11-10 22:47:24', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (263, 14, 184, 0, 1, '2022-11-10 22:47:57', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (264, 14, 185, 0, 1, '2022-11-10 22:48:17', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (265, 14, 186, 0, 1, '2022-11-10 22:48:36', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (266, 14, 187, 0, 1, '2022-11-10 22:49:02', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (267, 14, 188, 0, 1, '2022-11-10 22:49:30', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (268, 10, 189, 0, 1, '2022-11-10 22:51:06', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (269, 10, 190, 0, 1, '2022-11-10 22:52:12', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (270, 10, 191, 0, 1, '2022-11-10 22:52:34', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (271, 10, 192, 0, 1, '2022-11-10 22:57:24', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (272, 10, 193, 0, 1, '2022-11-10 22:57:46', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (273, 10, 194, 0, 1, '2022-11-10 22:58:06', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (274, 10, 195, 0, 1, '2022-11-10 22:58:24', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (275, 17, 196, 0, 1, '2022-11-10 22:59:01', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (276, 18, 197, 0, 1, '2022-11-10 23:00:17', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (277, 18, 198, 0, 1, '2022-11-10 23:00:44', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (278, 18, 199, 0, 1, '2022-11-10 23:01:09', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (279, 49, 200, 0, 1, '2022-11-10 23:01:40', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (280, 49, 201, 0, 1, '2022-11-10 23:03:07', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (281, 49, 202, 0, 1, '2022-11-10 23:03:31', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (282, 50, 203, 0, 1, '2022-11-10 23:05:00', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (283, 50, 204, 0, 1, '2022-11-10 23:05:28', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (284, 110, 205, 0, 1, '2022-11-10 23:07:42', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (285, 110, 206, 0, 1, '2022-11-10 23:08:06', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (286, 110, 207, 0, 1, '2022-11-10 23:09:41', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (287, 2, 208, 0, 1, '2022-11-10 23:12:10', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (288, 2, 209, 0, 1, '2022-11-10 23:12:33', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (289, 2, 210, 0, 1, '2022-11-10 23:12:49', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (290, 2, 211, 0, 1, '2022-11-10 23:13:07', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (291, 2, 212, 0, 1, '2022-11-10 23:13:24', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (292, 2, 213, 0, 1, '2022-11-10 23:13:42', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (293, 2, 214, 0, 1, '2022-11-10 23:14:04', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (294, 2, 215, 0, 1, '2022-11-10 23:14:30', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (295, 2, 216, 0, 1, '2022-11-10 23:15:05', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (296, 2, 217, 0, 1, '2022-11-10 23:54:26', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (297, 2, 218, 0, 1, '2022-11-10 23:57:04', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (298, 2, 219, 0, 1, '2022-11-10 23:57:35', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (299, 2, 220, 0, 1, '2022-11-10 23:57:59', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (300, 2, 221, 0, 1, '2022-11-10 23:58:34', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (301, 2, 222, 0, 1, '2022-11-10 23:58:55', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (302, 2, 223, 0, 1, '2022-11-10 23:59:17', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (303, 2, 224, 0, 1, '2022-11-10 23:59:36', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (304, 113, 225, 0, 1, '2022-11-11 00:14:59', 1, '2022-11-10 16:15:05');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (305, 113, 226, 0, 1, '2022-11-11 00:15:14', 1, '2022-11-10 16:15:26');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (306, 113, 229, 0, 1, '2022-11-11 17:55:19', 1, '2022-11-11 09:55:25');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (307, 113, 228, 0, 1, '2022-11-11 17:56:18', 1, '2022-11-11 09:56:23');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (308, 113, 233, 0, 1, '2022-11-11 18:00:54', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (309, 113, 235, 0, 1, '2022-11-11 18:13:29', 1, '2022-11-11 10:16:05');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (310, 84, 238, 0, 1, '2022-11-11 21:28:22', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (311, 68, 240, 0, 1, '2022-11-11 22:35:50', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (312, 69, 240, 0, 1, '2022-11-11 22:35:50', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (313, 74, 240, 0, 1, '2022-11-11 22:35:50', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (314, 68, 241, 1, 1, '2022-11-11 22:38:03', 1, '2022-11-11 22:38:37');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (315, 69, 241, 0, 1, '2022-11-11 22:38:37', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (316, 68, 242, 0, 1, '2022-11-11 22:39:17', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (317, 74, 243, 1, 1, '2022-11-11 22:39:52', 1, '2022-11-11 22:40:07');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (318, 74, 243, 0, 1, '2022-11-11 22:40:07', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (319, 71, 244, 0, 1, '2022-11-11 22:41:50', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (320, 72, 244, 0, 1, '2022-11-11 22:41:50', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (321, 73, 244, 0, 1, '2022-11-11 22:41:50', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (322, 71, 245, 1, 1, '2022-11-11 22:42:19', 1, '2022-11-11 22:42:34');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (323, 71, 245, 0, 1, '2022-11-11 22:42:34', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (324, 72, 246, 0, 1, '2022-11-11 22:43:04', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (325, 73, 247, 0, 1, '2022-11-11 22:43:53', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (326, 85, 249, 0, 1, '2022-11-11 22:46:15', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (327, 85, 250, 0, 1, '2022-11-11 22:47:06', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (328, 85, 251, 0, 1, '2022-11-11 22:47:46', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (329, 85, 252, 0, 1, '2022-11-11 22:48:30', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (330, 85, 253, 0, 1, '2022-11-11 22:49:09', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (331, 16, 255, 0, 1, '2022-11-11 23:09:43', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (332, 16, 256, 0, 1, '2022-11-11 23:10:13', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (333, 16, 258, 0, 1, '2022-11-11 23:10:35', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (334, 32, 259, 1, 1, '2022-11-11 23:10:49', 1, '2022-12-15 23:44:18');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (335, 16, 260, 0, 1, '2022-11-11 23:10:56', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (336, 16, 261, 0, 1, '2022-11-11 23:11:16', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (337, 32, 262, 1, 1, '2022-11-11 23:11:29', 1, '2022-12-20 21:03:54');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (338, 32, 263, 1, 1, '2022-11-11 23:12:03', 1, '2022-12-20 21:04:06');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (339, 21, 264, 0, 1, '2022-11-11 23:12:19', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (340, 32, 265, 1, 1, '2022-11-11 23:12:44', 1, '2022-12-20 21:14:33');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (341, 21, 266, 0, 1, '2022-11-11 23:13:14', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (342, 21, 267, 0, 1, '2022-11-11 23:13:53', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (343, 21, 268, 0, 1, '2022-11-11 23:14:16', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (344, 21, 269, 0, 1, '2022-11-11 23:14:39', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (345, 22, 271, 0, 1, '2022-11-11 23:15:01', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (346, 37, 272, 0, 1, '2022-11-11 23:15:23', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (347, 22, 273, 0, 1, '2022-11-11 23:15:26', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (348, 22, 274, 0, 1, '2022-11-11 23:15:50', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (349, 22, 275, 0, 1, '2022-11-11 23:16:24', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (350, 22, 276, 0, 1, '2022-11-11 23:16:45', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (351, 22, 277, 1, 1, '2022-11-11 23:17:07', 1, '2022-12-22 13:02:13');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (352, 46, 278, 0, 1, '2022-11-11 23:20:30', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (353, 46, 279, 0, 1, '2022-11-11 23:20:52', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (354, 46, 280, 0, 1, '2022-11-11 23:21:12', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (355, 46, 281, 0, 1, '2022-11-11 23:21:35', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (356, 46, 282, 0, 1, '2022-11-11 23:21:56', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (357, 46, 283, 0, 1, '2022-11-11 23:22:19', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (358, 46, 284, 0, 1, '2022-11-11 23:22:40', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (359, 46, 285, 0, 1, '2022-11-11 23:23:00', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (360, 48, 286, 0, 1, '2022-11-11 23:25:29', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (361, 48, 287, 0, 1, '2022-11-11 23:25:57', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (362, 48, 288, 0, 1, '2022-11-11 23:26:19', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (363, 48, 289, 0, 1, '2022-11-11 23:26:42', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (364, 48, 290, 0, 1, '2022-11-11 23:27:03', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (365, 48, 291, 0, 1, '2022-11-11 23:27:23', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (366, 48, 292, 0, 1, '2022-11-11 23:27:44', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (367, 114, 294, 0, 1, '2022-11-23 10:38:23', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (368, 115, 294, 0, 1, '2022-11-23 10:38:23', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (369, 114, 295, 0, 1, '2022-11-23 10:39:02', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (370, 114, 296, 0, 1, '2022-11-23 10:39:30', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (371, 114, 297, 0, 1, '2022-11-23 10:39:53', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (372, 114, 298, 0, 1, '2022-11-23 10:40:14', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (373, 114, 299, 0, 1, '2022-11-23 10:40:37', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (374, 114, 300, 0, 1, '2022-11-23 10:41:01', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (375, 114, 301, 0, 1, '2022-11-23 10:41:21', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (376, 114, 302, 0, 1, '2022-11-23 10:41:45', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (377, 114, 303, 0, 1, '2022-11-23 10:42:06', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (378, 114, 304, 0, 1, '2022-11-23 10:42:25', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (379, 114, 305, 0, 1, '2022-11-23 10:42:41', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (380, 115, 306, 0, 1, '2022-11-23 10:43:05', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (381, 115, 307, 0, 1, '2022-11-23 10:43:31', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (382, 115, 308, 0, 1, '2022-11-23 10:45:49', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (383, 92, 310, 1, 1, '2022-11-28 15:45:40', 1, '2022-11-28 15:45:58');
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (384, 92, 310, 0, 1, '2022-11-28 15:45:58', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (385, 32, 259, 0, 1, '2022-12-15 23:44:18', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (386, 32, 262, 0, 1, '2022-12-20 21:03:54', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (387, 32, 263, 0, 1, '2022-12-20 21:04:06', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (388, 32, 265, 0, 1, '2022-12-20 21:14:33', 0, NULL);
INSERT INTO `sys_menu_api` (`id`, `menu_id`, `api_id`, `is_deleted`, `created_by`, `created_at`, `deleted_by`, `deleted_at`) VALUES (389, 22, 277, 0, 1, '2022-12-22 13:02:13', 0, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_menu_button
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu_button`;
CREATE TABLE `sys_menu_button` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL COMMENT '父ID',
  `menu_id` int(11) NOT NULL COMMENT '菜单ID',
  `name` varchar(255) NOT NULL COMMENT '名称',
  `types` varchar(20) NOT NULL COMMENT '类型 自定义 add添加 edit编辑 del 删除',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `status` int(11) NOT NULL COMMENT '状态 0 停用 1启用',
  `is_deleted` int(11) NOT NULL COMMENT '是否删除 0未删除 1已删除',
  `created_by` int(10) unsigned DEFAULT '0' COMMENT '创建人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` int(11) DEFAULT NULL COMMENT '修改人',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_by` int(11) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=163 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='菜单与按钮关系表';

-- ----------------------------
-- Records of sys_menu_button
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (1, -1, 34, '添加', 'add', 'adsf asdf 地方敖德萨手动阀', 1, 1, 1, '2022-08-14 15:49:55', 0, '2022-08-14 16:04:01', 1, '2022-08-14 16:04:01');
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (2, -1, 34, '213', '123', '123123', 1, 1, 1, '2022-08-14 15:58:14', 0, '2022-08-14 16:02:25', 1, '2022-08-14 16:02:25');
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (3, 2, 34, 'safd', '爱的方式', '1123  ', 1, 1, 1, '2022-08-14 15:58:21', 0, '2022-08-14 16:02:23', 1, '2022-08-14 16:02:23');
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (4, -1, 34, '123-1111', '123123', '123', 0, 1, 1, '2022-08-13 08:04:10', 1, '2022-08-14 20:56:10', 0, '2022-08-14 20:56:10');
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (5, -1, 34, '12312', '受访人人情味儿去', '', 1, 1, 1, '2022-08-14 16:04:17', 0, '2022-08-14 16:04:34', 1, '2022-08-14 16:04:34');
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (6, -1, 34, '1232', 'asdf', '', 1, 1, 1, '2022-08-14 16:12:21', 0, '2022-08-14 20:56:12', 0, '2022-08-14 20:56:12');
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (7, -1, 34, '12', '1212', '', 1, 1, 0, '2022-08-14 20:47:59', 0, '2022-08-14 20:56:15', 0, '2022-08-14 20:56:15');
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (8, -1, 34, '121', '12122', '', 1, 1, 0, '2022-08-14 20:48:17', 0, '2022-08-14 20:48:37', 0, '2022-08-14 20:48:37');
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (9, -1, 34, '曲儿群翁', '121222', '', 1, 1, 0, '2022-08-14 20:48:29', 0, '2022-08-14 20:48:34', 0, '2022-08-14 20:48:34');
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (10, -1, 34, '新增', 'add', '', 1, 0, 0, '2022-08-13 04:55:53', 1, '2022-08-14 22:17:55', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (11, -1, 34, '编辑', 'edit', '', 1, 0, 0, '2022-08-13 12:55:57', 1, '2022-08-14 22:17:53', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (12, -1, 34, '详情', 'detail', '', 1, 0, 0, '2022-08-13 20:55:59', 1, '2022-08-14 22:17:47', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (13, -1, 34, '删除', 'del', '', 1, 0, 0, '2022-08-14 20:56:02', 0, '2022-08-14 20:56:02', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (14, -1, 34, '导入', 'upload', '', 1, 0, 0, '2022-08-14 20:56:04', 0, '2022-08-14 20:56:04', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (15, -1, 34, '导出', 'dwonload', '', 1, 0, 0, '2022-08-14 20:56:06', 0, '2022-08-14 20:56:06', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (16, -1, 35, '新增', 'add', '', 1, 0, 0, '2022-08-14 21:14:48', 0, '2022-08-14 21:14:48', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (17, -1, 15, '新增', 'add', '', 1, 0, 1, '2022-08-19 06:17:58', 1, '2022-08-19 23:27:07', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (18, -1, 15, '编辑', 'edit', '', 1, 0, 1, '2022-08-19 06:18:02', 1, '2022-08-19 23:27:42', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (19, -1, 15, '删除', 'del', '', 1, 0, 1, '2022-08-19 22:18:04', 0, '2022-08-19 22:18:04', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (20, -1, 15, '重置', 'reset', '', 1, 1, 1, '2022-08-19 22:18:35', 0, '2022-08-19 22:36:03', 1, '2022-08-19 22:36:03');
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (21, -1, 15, '重置', 'reset', '', 1, 0, 1, '2022-08-19 22:36:06', 0, '2022-08-19 22:36:06', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (22, -1, 57, '新增', 'add', '', 1, 1, 1, '2022-10-16 11:04:18', 0, '2022-10-16 11:04:36', 1, '2022-10-16 11:04:36');
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (23, -1, 57, '编辑', 'edit', '', 1, 1, 1, '2022-10-16 11:04:20', 0, '2022-10-16 11:04:33', 1, '2022-10-16 11:04:33');
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (24, -1, 57, '删除', 'del', '', 1, 1, 1, '2022-10-16 11:04:23', 0, '2022-10-16 11:04:30', 1, '2022-10-16 11:04:30');
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (25, -1, 13, '新增', 'add', '', 1, 0, 1, '2022-11-03 00:03:46', 0, '2022-11-03 00:03:46', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (26, -1, 13, '删除', 'del', '', 1, 0, 1, '2022-11-03 00:03:48', 0, '2022-11-03 00:03:48', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (27, -1, 13, '编辑', 'edit', '', 1, 0, 1, '2022-11-03 00:03:51', 0, '2022-11-03 00:03:51', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (28, -1, 37, '新增', 'add', '', 1, 0, 1, '2022-11-03 11:00:38', 0, '2022-11-03 11:00:38', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (29, -1, 37, '编辑', 'edit', '', 1, 0, 1, '2022-11-03 11:00:41', 0, '2022-11-03 11:00:41', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (30, -1, 37, '详情', 'detail', '', 1, 0, 1, '2022-11-03 11:00:42', 0, '2022-11-03 11:00:42', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (31, -1, 37, '删除', 'del', '', 1, 0, 1, '2022-11-03 11:00:44', 0, '2022-11-03 11:00:44', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (32, -1, 37, '复制', 'copy', '', 1, 0, 1, '2022-11-03 03:04:35', 1, '2022-11-03 11:04:50', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (33, -1, 12, '新增', 'add', '', 1, 0, 1, '2022-11-03 20:04:31', 0, '2022-11-03 20:04:31', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (34, -1, 12, '删除', 'del', '', 1, 0, 1, '2022-11-03 20:04:33', 0, '2022-11-03 20:04:33', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (35, -1, 12, '编辑', 'edit', '', 1, 0, 1, '2022-11-03 20:04:36', 0, '2022-11-03 20:04:36', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (36, -1, 14, '新增', 'add', '', 1, 0, 1, '2022-11-03 20:10:54', 0, '2022-11-03 20:10:54', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (37, -1, 14, '编辑', 'edit', '', 1, 0, 1, '2022-11-03 20:10:57', 0, '2022-11-03 20:10:57', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (38, -1, 14, '删除', 'del', '', 1, 0, 1, '2022-11-03 20:10:59', 0, '2022-11-03 20:10:59', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (39, -1, 10, '新增', 'add', '', 1, 0, 1, '2022-11-03 20:30:49', 0, '2022-11-03 20:30:49', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (40, -1, 10, '编辑', 'edit', '', 1, 0, 1, '2022-11-03 20:30:50', 0, '2022-11-03 20:30:50', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (41, -1, 10, '删除', 'del', '', 1, 0, 1, '2022-11-03 20:30:53', 0, '2022-11-03 20:30:53', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (42, -1, 10, '角色权限', 'role-premission', '', 1, 0, 1, '2022-11-03 20:31:17', 0, '2022-11-03 20:31:17', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (43, -1, 10, '数据权限', 'data-premission', '', 1, 0, 1, '2022-11-03 20:31:31', 0, '2022-11-03 20:31:31', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (44, -1, 18, '删除', 'del', '', 1, 0, 1, '2022-11-03 20:35:13', 0, '2022-11-03 20:35:13', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (45, -1, 49, '删除', 'del', '', 1, 0, 1, '2022-11-03 20:35:35', 0, '2022-11-03 20:35:35', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (46, -1, 49, '详情', 'detail', '', 1, 0, 1, '2022-11-03 20:42:52', 0, '2022-11-03 20:42:52', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (47, -1, 50, '强退', 'out', '', 1, 0, 1, '2022-11-03 20:48:01', 0, '2022-11-03 20:48:01', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (48, -1, 110, '启用', 'start', '', 1, 0, 1, '2022-11-03 21:01:05', 0, '2022-11-03 21:01:05', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (49, -1, 110, '停用', 'stop', '', 1, 0, 1, '2022-11-03 21:01:12', 0, '2022-11-03 21:01:12', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (50, -1, 32, '新增', 'add', '', 1, 0, 1, '2022-11-03 21:41:08', 0, '2022-11-03 21:41:08', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (51, -1, 2, '新增', 'add', '', 1, 0, 1, '2022-11-03 21:41:20', 0, '2022-11-03 21:41:20', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (52, -1, 101, '重置', 'reset', '', 1, 0, 1, '2022-11-04 06:46:59', 1, '2022-11-04 22:49:36', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (53, -1, 101, '查询', 'query', '', 1, 0, 1, '2022-11-04 06:47:12', 1, '2022-11-04 22:49:37', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (54, -1, 101, '换热站', 'heatStation', '', 1, 0, 1, '2022-11-04 06:47:51', 1, '2022-11-04 22:49:38', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (55, -1, 101, '环路', 'loop', '', 1, 0, 1, '2022-11-04 06:47:58', 1, '2022-11-04 22:49:40', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (56, -1, 77, '重置', 'reset', '', 1, 0, 1, '2022-11-04 22:53:59', 0, '2022-11-04 22:53:59', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (57, -1, 77, '查询', 'query', '', 1, 0, 1, '2022-11-04 22:54:08', 0, '2022-11-04 22:54:08', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (58, -1, 79, '重置', 'reset', '', 1, 0, 1, '2022-11-04 22:56:11', 0, '2022-11-04 22:56:11', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (59, -1, 79, '查询', 'query', '', 1, 0, 1, '2022-11-04 22:56:20', 0, '2022-11-04 22:56:20', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (60, -1, 55, '新增', 'add', '', 1, 0, 1, '2022-11-04 22:58:26', 0, '2022-11-04 22:58:26', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (61, -1, 55, '重置', 'reset', '', 1, 0, 1, '2022-11-04 22:58:29', 0, '2022-11-04 22:58:29', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (62, -1, 55, '查询', 'query', '', 1, 0, 1, '2022-11-04 22:58:39', 0, '2022-11-04 22:58:39', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (63, -1, 54, '查询', 'query', '', 1, 0, 1, '2022-11-04 23:02:08', 0, '2022-11-04 23:02:08', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (64, -1, 54, '新增', 'add', '', 1, 0, 1, '2022-11-04 23:02:10', 0, '2022-11-04 23:02:10', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (65, -1, 54, '重置', 'reset', '', 1, 0, 1, '2022-11-04 23:02:13', 0, '2022-11-04 23:02:13', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (66, -1, 100, '小区-新增', 'regional-add', '', 1, 0, 1, '2022-11-03 07:06:53', 1, '2022-11-04 23:48:21', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (67, -1, 100, '小区-查询', 'regional-query', '', 1, 0, 1, '2022-11-03 23:07:03', 1, '2022-11-04 23:48:22', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (68, -1, 100, '小区-重置', 'regional-reset', '', 1, 0, 1, '2022-11-03 23:07:55', 1, '2022-11-04 23:48:24', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (69, -1, 100, '楼宇-新增', 'floor-add', '', 1, 0, 1, '2022-11-04 23:19:28', 0, '2022-11-04 23:19:28', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (70, -1, 100, '楼宇-查询', 'floor-query', '', 1, 0, 1, '2022-11-04 23:19:38', 0, '2022-11-04 23:19:38', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (71, -1, 100, '楼宇-重置', 'floor-reset', '', 1, 0, 1, '2022-11-04 23:19:51', 0, '2022-11-04 23:19:51', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (72, -1, 100, '单元-新增', 'unit-add', '', 1, 0, 1, '2022-11-04 23:31:20', 0, '2022-11-04 23:31:20', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (73, -1, 100, '单元-查询', 'unit-query', '', 1, 0, 1, '2022-11-04 23:31:32', 0, '2022-11-04 23:31:32', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (74, -1, 100, '单元-重置', 'unit-reset', '', 1, 0, 1, '2022-11-04 23:31:45', 0, '2022-11-04 23:31:45', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (75, -1, 100, '住户-新增', 'resident-add', '', 1, 0, 1, '2022-11-04 23:36:35', 0, '2022-11-04 23:36:35', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (76, -1, 100, '住户-查询', 'resident-query', '', 1, 0, 1, '2022-11-04 23:37:06', 0, '2022-11-04 23:37:06', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (77, -1, 100, '住户-重置', 'resident-reset', '', 1, 0, 1, '2022-11-04 23:37:20', 0, '2022-11-04 23:37:20', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (78, -1, 39, '新增', 'add', '', 1, 0, 1, '2022-11-06 09:44:42', 0, '2022-11-06 09:44:42', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (79, -1, 39, '编辑', 'edit', '', 1, 0, 1, '2022-11-06 09:44:44', 0, '2022-11-06 09:44:44', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (80, -1, 39, '详情', 'detail', '', 1, 0, 1, '2022-11-06 09:44:45', 0, '2022-11-06 09:44:45', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (81, -1, 39, '删除', 'del', '', 1, 0, 1, '2022-11-06 09:44:47', 0, '2022-11-06 09:44:47', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (82, -1, 39, '复制', 'copy', '', 1, 0, 1, '2022-11-06 09:46:46', 0, '2022-11-06 09:46:46', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (83, -1, 39, '数据记录', 'record', '', 1, 0, 1, '2022-11-06 09:47:41', 0, '2022-11-06 09:47:41', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (84, -1, 61, '新增', 'add', '', 1, 0, 1, '2022-11-06 09:52:32', 0, '2022-11-06 09:52:32', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (85, -1, 61, '编辑', 'edit', '', 1, 0, 1, '2022-11-06 09:52:33', 0, '2022-11-06 09:52:33', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (86, -1, 61, '删除', 'del', '', 1, 0, 1, '2022-11-06 09:52:38', 0, '2022-11-06 09:52:38', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (87, -1, 61, '发布与停止', 'pro-status', '', 1, 0, 1, '2022-11-06 09:54:09', 0, '2022-11-06 09:54:09', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (88, -1, 64, '新增', 'add', '', 1, 0, 1, '2022-11-06 09:58:15', 0, '2022-11-06 09:58:15', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (89, -1, 64, '编辑', 'edit', '', 1, 0, 1, '2022-11-06 09:58:16', 0, '2022-11-06 09:58:16', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (90, -1, 64, '详情', 'detail', '', 1, 0, 1, '2022-11-06 09:58:18', 0, '2022-11-06 09:58:18', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (91, -1, 64, '删除', 'del', '', 1, 0, 1, '2022-11-06 09:58:21', 0, '2022-11-06 09:58:21', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (92, -1, 35, '编辑', 'edit', '', 1, 0, 1, '2022-11-06 10:13:13', 0, '2022-11-06 10:13:13', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (93, -1, 35, '删除', 'del', '', 1, 0, 1, '2022-11-06 10:13:14', 0, '2022-11-06 10:13:14', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (94, -1, 36, '新增', 'add', '', 1, 0, 1, '2022-11-06 10:13:44', 0, '2022-11-06 10:13:44', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (95, -1, 36, '编辑', 'edit', '', 1, 0, 1, '2022-11-06 10:13:46', 0, '2022-11-06 10:13:46', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (96, -1, 36, '详情', 'detail', '', 1, 0, 1, '2022-11-06 10:13:48', 0, '2022-11-06 10:13:48', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (97, -1, 36, '删除', 'del', '', 1, 0, 1, '2022-11-06 10:13:49', 0, '2022-11-06 10:13:49', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (98, -1, 36, '启用与停止', 'status', '', 1, 0, 1, '2022-11-06 10:15:15', 0, '2022-11-06 10:15:15', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (99, -1, 84, '新增', 'add', '', 1, 1, 1, '2022-11-06 23:35:47', 0, '2022-11-06 23:35:53', 1, '2022-11-06 23:35:53');
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (100, -1, 66, '新增', 'add', '', 1, 0, 1, '2022-11-05 15:42:25', 1, '2022-11-11 15:14:45', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (101, -1, 66, '详情', 'detail', '', 1, 0, 1, '2022-11-06 23:43:23', 0, '2022-11-06 23:43:23', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (102, -1, 66, '编辑', 'edit', '', 1, 0, 1, '2022-11-06 23:43:26', 0, '2022-11-06 23:43:26', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (103, -1, 66, '查询', 'query', '', 1, 0, 1, '2022-11-05 15:45:29', 1, '2022-11-11 15:14:47', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (104, 105, 66, '删除', 'del', '', 1, 0, 1, '2022-11-06 15:46:02', 1, '2022-12-12 22:38:08', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (105, -1, 66, '更多', 'more', '', 1, 0, 1, '2022-11-06 07:46:18', 1, '2022-11-11 15:13:45', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (106, 105, 66, '启用', 'on', '', 1, 0, 1, '2022-11-05 23:46:34', 1, '2022-12-12 22:37:51', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (107, 105, 66, '禁用', 'off', '', 1, 0, 1, '2022-11-06 15:46:48', 1, '2022-12-12 22:38:01', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (108, -1, 2, '编辑', 'edit', '', 1, 0, 1, '2022-11-07 21:20:17', 0, '2022-11-07 21:20:17', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (109, -1, 2, '删除', 'del', '', 1, 0, 1, '2022-11-07 21:20:19', 0, '2022-11-07 21:20:19', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (110, -1, 2, '更多', 'more', '', 1, 0, 1, '2022-11-07 21:20:47', 0, '2022-11-07 21:20:47', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (111, 110, 2, '按钮权限', 'btn', '', 1, 0, 1, '2022-11-07 21:22:22', 0, '2022-11-07 21:22:22', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (112, 110, 2, '列表权限', 'list', '', 1, 0, 1, '2022-11-07 21:22:35', 0, '2022-11-07 21:22:35', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (113, -1, 16, '新增', 'add', '', 1, 0, 1, '2022-11-07 21:28:48', 0, '2022-11-07 21:28:48', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (114, -1, 16, '删除', 'del', '', 1, 0, 1, '2022-11-07 21:28:50', 0, '2022-11-07 21:28:50', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (115, -1, 16, '编辑', 'edit', '', 1, 0, 1, '2022-11-07 21:28:58', 0, '2022-11-07 21:28:58', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (116, -1, 21, '新增', 'add', '', 1, 0, 1, '2022-11-07 21:47:20', 0, '2022-11-07 21:47:20', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (117, -1, 21, '编辑', 'edit', '', 1, 0, 1, '2022-11-07 21:47:22', 0, '2022-11-07 21:47:22', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (118, -1, 21, '删除', 'del', '', 1, 0, 1, '2022-11-07 21:47:24', 0, '2022-11-07 21:47:24', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (119, -1, 22, '新增', 'add', '', 1, 0, 1, '2022-11-07 21:56:01', 0, '2022-11-07 21:56:01', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (120, -1, 22, '编辑', 'edit', '', 1, 0, 1, '2022-11-07 21:56:05', 0, '2022-11-07 21:56:05', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (121, -1, 22, '删除', 'del', '', 1, 0, 1, '2022-11-07 21:56:07', 0, '2022-11-07 21:56:07', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (122, -1, 32, '编辑', 'edit', '', 1, 0, 1, '2022-11-07 22:00:53', 0, '2022-11-07 22:00:53', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (123, -1, 32, '详情', 'detail', '', 1, 0, 1, '2022-11-07 22:00:55', 0, '2022-11-07 22:00:55', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (124, -1, 32, '删除', 'del', '', 1, 0, 1, '2022-11-07 22:00:57', 0, '2022-11-07 22:00:57', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (125, -1, 46, '新增', 'add', '', 1, 0, 1, '2022-11-07 22:38:32', 0, '2022-11-07 22:38:32', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (126, -1, 46, '编辑', 'edit', '', 1, 0, 1, '2022-11-07 22:38:33', 0, '2022-11-07 22:38:33', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (127, -1, 46, '删除', 'del', '', 1, 0, 1, '2022-11-07 22:38:35', 0, '2022-11-07 22:38:35', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (128, -1, 46, '执行一次', 'do', '', 1, 0, 1, '2022-11-07 22:38:43', 0, '2022-11-07 22:38:43', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (129, -1, 48, '新增', 'add', '', 1, 0, 1, '2022-11-07 22:44:22', 0, '2022-11-07 22:44:22', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (130, -1, 48, '编辑', 'edit', '', 1, 0, 1, '2022-11-07 22:44:23', 0, '2022-11-07 22:44:23', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (131, -1, 48, '删除', 'del', '', 1, 0, 1, '2022-11-07 22:44:25', 0, '2022-11-07 22:44:25', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (132, -1, 85, '新增', 'add', '', 1, 0, 1, '2022-11-07 22:47:24', 0, '2022-11-07 22:47:24', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (133, -1, 85, '详情', 'detail', '', 1, 1, 1, '2022-11-07 22:47:26', 0, '2022-11-07 22:48:46', 1, '2022-11-07 22:48:46');
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (134, -1, 85, '删除', 'del', '', 1, 0, 1, '2022-11-07 22:47:28', 0, '2022-11-07 22:47:28', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (135, -1, 85, '编辑', 'edit', '', 1, 0, 1, '2022-11-07 22:48:43', 0, '2022-11-07 22:48:43', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (136, -1, 35, '描述', 'desc', '', 1, 0, 1, '2022-11-10 09:44:20', 0, '2022-11-10 09:44:20', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (137, -1, 70, '详情', 'detail', '', 1, 0, 1, '2022-11-11 17:34:48', 0, '2022-11-11 17:34:48', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (138, -1, 70, '编辑', 'edit', '', 1, 0, 1, '2022-11-11 17:34:57', 0, '2022-11-11 17:34:57', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (139, -1, 70, '更多', 'more', '', 1, 0, 1, '2022-11-11 17:35:06', 0, '2022-11-11 17:35:06', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (140, 139, 70, '启用', 'on', '', 1, 0, 1, '2022-11-11 09:35:23', 1, '2022-12-12 22:39:12', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (141, 139, 70, '禁用', 'off', '', 1, 0, 1, '2022-11-11 09:35:35', 1, '2022-12-12 22:39:18', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (142, 139, 70, '删除', 'delete', '', 1, 0, 1, '2022-11-11 09:35:50', 1, '2022-12-12 22:39:06', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (143, -1, 70, '查询', 'query', '', 1, 0, 1, '2022-11-11 17:36:51', 0, '2022-11-11 17:36:51', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (144, -1, 70, '新增', 'add', '', 1, 0, 1, '2022-11-11 17:36:56', 0, '2022-11-11 17:36:56', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (145, -1, 85, '重置', 'reset', '', 1, 0, 1, '2022-11-11 17:45:25', 0, '2022-11-11 17:45:25', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (146, -1, 85, '查询', 'query', '', 1, 0, 1, '2022-11-11 17:45:38', 0, '2022-11-11 17:45:38', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (147, -1, 37, '查询', 'query', '', 1, 0, 1, '2022-11-11 22:53:40', 0, '2022-11-11 22:53:40', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (148, -1, 37, '重置', 'reset', '', 1, 0, 1, '2022-11-11 22:54:02', 0, '2022-11-11 22:54:02', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (149, -1, 115, '详情', 'detail', '', 1, 0, 1, '2022-11-15 14:44:01', 0, '2022-11-15 14:44:01', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (150, -1, 115, '处理', 'edit', '', 1, 0, 1, '2022-11-15 06:44:11', 1, '2022-11-15 14:44:21', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (151, -1, 114, '新增', 'add', '', 1, 0, 1, '2022-11-23 10:27:50', 0, '2022-11-23 10:27:50', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (152, -1, 114, '编辑', 'edit', '', 1, 0, 1, '2022-11-23 10:27:53', 0, '2022-11-23 10:27:53', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (153, -1, 114, '删除', 'del', '', 1, 0, 1, '2022-11-23 10:28:01', 0, '2022-11-23 10:28:01', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (154, -1, 114, '状态', 'status', '', 1, 0, 1, '2022-11-23 10:29:26', 0, '2022-11-23 10:29:26', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (155, -1, 114, '级别设置', 'level', '', 1, 0, 1, '2022-11-23 10:30:07', 0, '2022-11-23 10:30:07', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (156, -1, 92, '详情', 'detail', '', 1, 0, 1, '2022-11-28 10:20:53', 0, '2022-11-28 10:20:53', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (157, -1, 92, '处理', 'edit', '', 1, 0, 1, '2022-11-28 10:21:11', 0, '2022-11-28 10:21:11', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (158, -1, 41, '管理', 'setting', '', 1, 0, 1, '2022-12-07 14:36:08', 0, '2022-12-07 14:36:08', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (159, -1, 41, '配置', 'edit', '', 1, 0, 1, '2022-12-07 14:36:23', 0, '2022-12-07 14:36:23', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (160, -1, 42, '详情', 'detail', '', 1, 0, 1, '2022-12-15 20:19:34', 0, '2022-12-15 20:19:34', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (161, -1, 42, '删除', 'del', '', 1, 0, 1, '2022-12-16 13:58:00', 0, '2022-12-16 13:58:00', 0, NULL);
INSERT INTO `sys_menu_button` (`id`, `parent_id`, `menu_id`, `name`, `types`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (162, -1, 15, '修改用户状态', 'change-status', '', 1, 0, 1, '2022-12-20 21:27:06', 0, '2022-12-20 21:27:06', 0, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_menu_column
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu_column`;
CREATE TABLE `sys_menu_column` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL COMMENT '父ID',
  `menu_id` int(11) NOT NULL COMMENT '菜单ID',
  `name` varchar(255) NOT NULL COMMENT '名称',
  `code` varchar(50) NOT NULL COMMENT '代表字段',
  `description` varchar(255) DEFAULT NULL COMMENT '描述',
  `status` int(11) NOT NULL COMMENT '状态 0 停用 1启用',
  `is_deleted` int(11) NOT NULL COMMENT '是否删除 0未删除 1已删除',
  `created_by` int(10) unsigned DEFAULT '0' COMMENT '创建人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` int(11) DEFAULT NULL COMMENT '修改人',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_by` int(11) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=245 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='菜单列表字段';

-- ----------------------------
-- Records of sys_menu_column
-- ----------------------------
BEGIN;
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (1, -1, 34, 'ID', 'id', '', 1, 0, 1, '2022-08-13 21:57:56', 1, '2022-11-06 10:06:24', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (2, 1, 34, '213', '阿什顿发', '', 1, 1, 1, '2022-08-14 22:09:17', 0, '2022-08-14 22:41:18', 1, '2022-08-14 22:41:18');
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (3, -1, 34, '标识', 'key', '', 1, 0, 1, '2022-08-13 06:13:48', 1, '2022-11-06 10:06:35', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (4, -1, 15, '手机号', 'mobile', '', 0, 0, 1, '2022-08-19 06:42:39', 1, '2022-11-04 23:49:40', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (5, -1, 15, '用户昵称', 'userNickname', '', 0, 0, 1, '2022-08-19 14:54:11', 1, '2022-11-04 23:49:42', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (6, -1, 15, '1', '1', '', 1, 1, 1, '2022-08-19 23:04:02', 0, '2022-08-19 23:04:18', 1, '2022-08-19 23:04:18');
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (7, -1, 15, '部门', 'deptName', '', 0, 0, 1, '2022-11-02 16:02:17', 1, '2022-11-04 23:49:46', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (8, -1, 13, '组织名称', 'name', '', 1, 0, 1, '2022-11-03 00:05:00', 0, '2022-11-03 00:05:00', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (9, -1, 13, '组织编号', 'number', '', 1, 0, 1, '2022-11-03 00:05:24', 0, '2022-11-03 00:05:24', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (10, -1, 13, '组织状态', 'status', '', 1, 0, 1, '2022-11-03 00:05:47', 0, '2022-11-03 00:05:47', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (11, -1, 13, '排序', 'orderNum', '', 1, 0, 1, '2022-11-03 00:06:11', 0, '2022-11-03 00:06:11', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (12, -1, 13, '创建时间', 'createdAt', '', 1, 0, 1, '2022-11-03 00:06:28', 0, '2022-11-03 00:06:28', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (13, -1, 15, '账户名称', 'userName', '', 1, 0, 1, '2022-11-03 00:19:30', 0, '2022-11-03 00:19:30', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (14, -1, 15, '角色', 'rolesNames', '', 1, 0, 1, '2022-11-03 00:23:49', 0, '2022-11-03 00:23:49', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (15, -1, 15, '状态', 'status', '', 1, 0, 1, '2022-11-03 00:24:08', 0, '2022-11-03 00:24:08', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (16, -1, 15, '创建时间', 'createdAt', '', 1, 0, 1, '2022-11-03 00:24:28', 0, '2022-11-03 00:24:28', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (17, -1, 37, '数据源名称', 'name', '', 1, 0, 1, '2022-11-03 11:10:18', 0, '2022-11-03 11:10:18', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (18, -1, 37, '数据源类型', 'from', '', 1, 0, 1, '2022-11-03 11:10:27', 0, '2022-11-03 11:10:27', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (19, -1, 37, '状态', 'status', '', 1, 0, 1, '2022-11-03 11:10:37', 0, '2022-11-03 11:10:37', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (20, -1, 37, 'ID', 'sourceId', '', 1, 0, 1, '2022-11-03 11:11:36', 0, '2022-11-03 11:11:36', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (21, -1, 15, '操作', 'handle', '', 1, 0, 1, '2022-11-03 20:05:26', 0, '2022-11-03 20:05:26', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (22, -1, 13, '操作', 'handle', '', 1, 0, 1, '2022-11-03 20:06:26', 0, '2022-11-03 20:06:26', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (23, -1, 12, '操作', 'handle', '', 1, 0, 1, '2022-11-03 20:06:55', 0, '2022-11-03 20:06:55', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (24, -1, 12, '部门状态', 'status', '', 1, 0, 1, '2022-11-03 20:07:46', 0, '2022-11-03 20:07:46', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (25, -1, 12, '排序', 'orderNum', '', 1, 0, 1, '2022-11-03 20:08:00', 0, '2022-11-03 20:08:00', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (26, -1, 12, '创建时间', 'createdAt', '', 1, 0, 1, '2022-11-03 20:08:25', 0, '2022-11-03 20:08:25', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (27, -1, 14, '操作', 'handle', '', 1, 0, 1, '2022-11-03 20:20:47', 0, '2022-11-03 20:20:47', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (28, -1, 14, '创建时间', 'createdAt', '', 1, 0, 1, '2022-11-03 20:20:51', 0, '2022-11-03 20:20:51', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (29, -1, 14, '岗位描述', 'remark', '', 1, 0, 1, '2022-11-03 20:21:46', 0, '2022-11-03 20:21:46', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (30, -1, 14, '状态', 'status', '', 1, 0, 1, '2022-11-03 20:26:21', 0, '2022-11-03 20:26:21', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (31, -1, 14, '排序', 'postSort', '', 1, 0, 1, '2022-11-03 20:26:35', 0, '2022-11-03 20:26:35', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (32, -1, 14, '岗位名称', 'postName', '', 1, 0, 1, '2022-11-03 20:27:21', 0, '2022-11-03 20:27:21', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (33, -1, 14, '岗位编码', 'postCode', '', 1, 0, 1, '2022-11-03 20:27:32', 0, '2022-11-03 20:27:32', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (34, -1, 10, '状态', 'status', '', 1, 0, 1, '2022-11-03 20:32:19', 0, '2022-11-03 20:32:19', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (35, -1, 10, '创建时间', 'createdAt', '', 1, 0, 1, '2022-11-03 20:32:20', 0, '2022-11-03 20:32:20', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (36, -1, 10, '操作', 'handle', '', 1, 0, 1, '2022-11-03 20:32:21', 0, '2022-11-03 20:32:21', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (37, -1, 10, '排序', 'listOrder', '', 1, 0, 1, '2022-11-03 20:32:34', 0, '2022-11-03 20:32:34', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (38, -1, 10, '角色描述', 'remark', '', 1, 0, 1, '2022-11-03 20:32:44', 0, '2022-11-03 20:32:44', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (39, -1, 10, '角色名称', 'name', '', 1, 0, 1, '2022-11-03 20:32:54', 0, '2022-11-03 20:32:54', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (40, -1, 18, '状态', 'status', '', 1, 0, 1, '2022-11-03 20:39:18', 0, '2022-11-03 20:39:18', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (41, -1, 18, '登录地点', 'loginLocation', '', 1, 0, 1, '2022-11-03 20:39:34', 0, '2022-11-03 20:39:34', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (42, -1, 18, '操作信息', 'msg', '', 1, 0, 1, '2022-11-03 20:39:50', 0, '2022-11-03 20:39:50', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (43, -1, 18, '登录日期', 'loginTime', '', 1, 0, 1, '2022-11-03 12:40:03', 1, '2022-11-03 20:40:23', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (44, -1, 18, '登录模块', 'module', '', 1, 0, 1, '2022-11-03 20:40:15', 0, '2022-11-03 20:40:15', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (45, -1, 49, '操作', 'handle', '', 1, 0, 1, '2022-11-03 20:43:02', 0, '2022-11-03 20:43:02', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (46, -1, 49, '状态', 'status', '', 1, 0, 1, '2022-11-03 20:43:15', 0, '2022-11-03 20:43:15', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (47, -1, 49, '操作类型', 'operatorType', '', 1, 0, 1, '2022-11-03 20:44:09', 0, '2022-11-03 20:44:09', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (48, -1, 49, '操作人员', 'operName', '', 1, 0, 1, '2022-11-03 20:45:06', 0, '2022-11-03 20:45:06', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (49, -1, 49, '操作地点', 'operLocation', '', 1, 0, 1, '2022-11-03 20:45:20', 0, '2022-11-03 20:45:20', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (50, -1, 49, '操作时间', 'operTime', '', 1, 0, 1, '2022-11-03 20:45:31', 0, '2022-11-03 20:45:31', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (51, -1, 50, '操作', 'handle', '', 1, 0, 1, '2022-11-03 20:48:09', 0, '2022-11-03 20:48:09', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (52, -1, 50, '用户名', 'userName', '', 1, 0, 1, '2022-11-03 20:49:36', 0, '2022-11-03 20:49:36', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (53, -1, 50, '登录地址', 'ip', '', 1, 0, 1, '2022-11-03 20:49:45', 0, '2022-11-03 20:49:45', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (54, -1, 50, '操作系统', 'os', '', 1, 0, 1, '2022-11-03 20:49:53', 0, '2022-11-03 20:49:53', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (55, -1, 110, '操作', 'handle', '', 1, 0, 1, '2022-11-03 21:01:29', 0, '2022-11-03 21:01:29', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (56, -1, 110, '状态', 'status', '', 1, 0, 1, '2022-11-03 21:01:57', 0, '2022-11-03 21:01:57', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (57, -1, 110, '作者 ', 'author', '', 1, 0, 1, '2022-11-03 21:02:05', 0, '2022-11-03 21:02:05', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (58, -1, 110, '介绍', 'intro', '', 1, 0, 1, '2022-11-03 21:02:13', 0, '2022-11-03 21:02:13', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (59, -1, 110, '名称', 'name', '', 1, 0, 1, '2022-11-03 21:02:21', 0, '2022-11-03 21:02:21', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (110, -1, 39, 'ID', 'id', '', 1, 0, 1, '2022-11-06 09:49:11', 0, '2022-11-06 09:49:11', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (111, -1, 39, '模型名称', 'name', '', 1, 0, 1, '2022-11-06 09:49:27', 0, '2022-11-06 09:49:27', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (112, -1, 39, '创建时间', 'createdAt', '', 1, 0, 1, '2022-11-06 09:49:37', 0, '2022-11-06 09:49:37', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (113, -1, 39, '操作', 'handle', '', 1, 0, 1, '2022-11-06 09:49:40', 0, '2022-11-06 09:49:40', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (114, -1, 61, 'ID', 'id', '', 1, 0, 1, '2022-11-06 09:55:37', 0, '2022-11-06 09:55:37', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (115, -1, 61, '数据标识', 'key', '', 1, 0, 1, '2022-11-05 17:55:43', 1, '2022-11-06 09:55:59', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (116, -1, 61, '数据名称', 'name', '', 1, 0, 1, '2022-11-06 09:56:07', 0, '2022-11-06 09:56:07', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (117, -1, 61, '数据类型', 'dataType', '', 1, 0, 1, '2022-11-06 09:56:15', 0, '2022-11-06 09:56:15', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (118, -1, 61, '数据取值项', 'value', '', 1, 0, 1, '2022-11-06 09:56:22', 0, '2022-11-06 09:56:22', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (119, -1, 61, '创建时间', 'createdAt', '', 1, 0, 1, '2022-11-06 09:56:31', 0, '2022-11-06 09:56:31', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (120, -1, 64, 'ID', 'id', '', 1, 0, 1, '2022-11-06 10:01:29', 0, '2022-11-06 10:01:29', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (121, -1, 64, '字段名称', 'key', '', 1, 0, 1, '2022-11-06 10:01:38', 0, '2022-11-06 10:01:38', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (122, -1, 64, '字段标题', 'name', '', 1, 0, 1, '2022-11-06 10:01:45', 0, '2022-11-06 10:01:45', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (123, -1, 64, '类型', 'dataType', '', 1, 0, 1, '2022-11-06 10:01:51', 0, '2022-11-06 10:01:51', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (124, -1, 64, '数据源名称', 'from', '', 1, 0, 1, '2022-11-06 10:01:59', 0, '2022-11-06 10:01:59', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (125, -1, 64, '默认值', 'default', '', 1, 0, 1, '2022-11-06 10:02:10', 0, '2022-11-06 10:02:10', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (126, -1, 64, '备注说明', 'value', '', 1, 0, 1, '2022-11-06 10:02:17', 0, '2022-11-06 10:02:17', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (127, -1, 64, '创建时间', 'createdAt', '', 1, 0, 1, '2022-11-06 10:02:24', 0, '2022-11-06 10:02:24', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (128, -1, 34, '名称', 'name', '', 1, 0, 1, '2022-11-06 10:06:42', 0, '2022-11-06 10:06:42', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (129, -1, 34, '分类', 'categoryName', '', 1, 0, 1, '2022-11-06 10:06:49', 0, '2022-11-06 10:06:49', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (130, -1, 34, '部门', 'deptName', '', 1, 0, 1, '2022-11-06 10:06:56', 0, '2022-11-06 10:06:56', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (131, -1, 34, '消息协议', 'messageProtocol', '', 1, 0, 1, '2022-11-06 10:07:11', 0, '2022-11-06 10:07:11', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (132, -1, 34, '传输协议', 'transportProtocol', '', 1, 0, 1, '2022-11-06 10:07:23', 0, '2022-11-06 10:07:23', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (133, -1, 34, '类型', 'deviceType', '', 1, 0, 1, '2022-11-06 10:07:33', 0, '2022-11-06 10:07:33', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (134, -1, 34, '状态', 'status', '', 1, 0, 1, '2022-11-06 10:07:41', 0, '2022-11-06 10:07:41', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (135, -1, 35, '分类名称', 'name', '', 1, 0, 1, '2022-11-06 10:12:59', 0, '2022-11-06 10:12:59', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (136, -1, 35, '描述', 'desc', '', 1, 0, 1, '2022-11-06 02:13:03', 1, '2022-11-08 23:46:21', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (137, -1, 36, 'ID', 'id', '', 1, 0, 1, '2022-11-06 10:15:48', 0, '2022-11-06 10:15:48', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (138, -1, 36, '标识', 'key', '', 1, 0, 1, '2022-11-06 10:15:58', 0, '2022-11-06 10:15:58', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (139, -1, 36, '名称', 'name', '', 1, 0, 1, '2022-11-06 10:16:06', 0, '2022-11-06 10:16:06', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (140, -1, 36, '产品名称', 'productName', '', 1, 0, 1, '2022-11-06 10:16:13', 0, '2022-11-06 10:16:13', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (141, -1, 36, '部门', 'deptName', '', 1, 1, 1, '2022-11-06 10:16:22', 0, '2022-11-08 23:44:21', 1, '2022-11-08 23:44:21');
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (142, -1, 36, '状态', 'status', '', 1, 0, 1, '2022-11-06 02:16:30', 1, '2022-11-06 10:16:54', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (143, -1, 36, '激活时间', 'registryTime', '', 1, 0, 1, '2022-11-06 10:16:38', 0, '2022-11-06 10:16:38', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (144, -1, 36, '最后上线时间', 'lastOnlineTime', '', 1, 0, 1, '2022-11-06 10:16:47', 0, '2022-11-06 10:16:47', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (145, -1, 66, '状态', 'status', '', 1, 0, 1, '2022-11-07 00:01:36', 0, '2022-11-07 00:01:36', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (146, -1, 66, '名称', 'name', '', 1, 0, 1, '2022-11-07 00:01:52', 0, '2022-11-07 00:01:52', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (147, -1, 66, '类型', 'type', '', 1, 0, 1, '2022-11-07 00:02:08', 0, '2022-11-07 00:02:08', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (148, -1, 66, '地址', 'address', '', 1, 0, 1, '2022-11-07 00:02:21', 0, '2022-11-07 00:02:21', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (149, -1, 66, '创建时间', 'createTime', '', 1, 0, 1, '2022-11-06 16:02:46', 1, '2022-11-07 00:02:55', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (150, -1, 2, '操作', 'handle', '', 1, 0, 1, '2022-11-07 21:23:51', 0, '2022-11-07 21:23:51', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (151, -1, 2, '状态', 'status', '', 1, 1, 1, '2022-11-07 21:23:53', 0, '2022-11-07 21:24:10', 1, '2022-11-07 21:24:10');
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (152, -1, 2, '菜单名称', 'title', '', 1, 0, 1, '2022-11-07 21:24:01', 0, '2022-11-07 21:24:01', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (153, -1, 2, '路由路径', 'path', '', 1, 0, 1, '2022-11-07 21:25:02', 0, '2022-11-07 21:25:02', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (154, -1, 2, '组件路径', 'component', '', 1, 0, 1, '2022-11-07 21:25:13', 0, '2022-11-07 21:25:13', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (155, -1, 2, '排序', 'weigh', '', 1, 0, 1, '2022-11-07 21:25:20', 0, '2022-11-07 21:25:20', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (156, -1, 2, '排序配型', 'menuType', '', 1, 0, 1, '2022-11-07 21:25:34', 0, '2022-11-07 21:25:34', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (157, -1, 2, '显示状态', 'isHide', '', 1, 0, 1, '2022-11-07 21:25:48', 0, '2022-11-07 21:25:48', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (158, -1, 16, '操作', 'handle', '', 1, 0, 1, '2022-11-07 21:42:05', 0, '2022-11-07 21:42:05', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (159, -1, 16, 'ID', 'configId', '', 1, 0, 1, '2022-11-07 21:42:35', 0, '2022-11-07 21:42:35', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (160, -1, 16, '参数名称', 'configName', '', 1, 0, 1, '2022-11-07 21:42:45', 0, '2022-11-07 21:42:45', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (161, -1, 16, '参数键名', 'configKey', '', 1, 0, 1, '2022-11-07 21:42:57', 0, '2022-11-07 21:42:57', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (162, -1, 16, '参数键值', 'configValue', '', 1, 0, 1, '2022-11-07 21:43:07', 0, '2022-11-07 21:43:07', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (163, -1, 16, '备注', 'remark', '', 1, 0, 1, '2022-11-07 21:43:15', 0, '2022-11-07 21:43:15', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (164, -1, 16, '系统内置', 'configType', '', 1, 0, 1, '2022-11-07 21:43:26', 0, '2022-11-07 21:43:26', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (165, -1, 21, '操作', 'handle', '', 1, 0, 1, '2022-11-07 21:53:45', 0, '2022-11-07 21:53:45', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (166, -1, 21, '状态', 'status', '', 1, 0, 1, '2022-11-07 21:53:49', 0, '2022-11-07 21:53:49', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (167, -1, 21, '字典ID', 'dictId', '', 1, 0, 1, '2022-11-07 21:54:01', 0, '2022-11-07 21:54:01', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (168, -1, 21, '字典名称', 'dictName', '', 1, 0, 1, '2022-11-07 21:54:09', 0, '2022-11-07 21:54:09', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (169, -1, 21, '字典类型', 'dictType', '', 1, 0, 1, '2022-11-07 21:54:19', 0, '2022-11-07 21:54:19', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (170, -1, 21, '字典描述', 'remark', '', 1, 0, 1, '2022-11-07 21:54:32', 0, '2022-11-07 21:54:32', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (171, -1, 22, '操作', 'handle', '', 1, 0, 1, '2022-11-07 21:56:33', 0, '2022-11-07 21:56:33', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (172, -1, 22, '状态', 'status', '', 1, 0, 1, '2022-11-07 21:57:20', 0, '2022-11-07 21:57:20', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (173, -1, 22, '备注', 'remark', '', 1, 0, 1, '2022-11-07 21:57:35', 0, '2022-11-07 21:57:35', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (174, -1, 22, '字典排序', 'dictSort', '', 1, 0, 1, '2022-11-07 21:57:44', 0, '2022-11-07 21:57:44', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (175, -1, 22, '字典键值', 'dictValue', '', 1, 0, 1, '2022-11-07 21:58:00', 0, '2022-11-07 21:58:00', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (176, -1, 22, '字典标签', 'dictLabel', '', 1, 0, 1, '2022-11-07 21:58:11', 0, '2022-11-07 21:58:11', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (177, -1, 22, '字典编码', 'dictCode', '', 1, 0, 1, '2022-11-07 21:58:20', 0, '2022-11-07 21:58:20', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (178, -1, 32, '操作', 'handle', '', 1, 0, 1, '2022-11-07 22:01:31', 0, '2022-11-07 22:01:31', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (179, -1, 32, '名称', 'title', '', 1, 0, 1, '2022-11-07 22:01:40', 0, '2022-11-07 22:01:40', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (180, -1, 32, '描述', 'explain', '', 1, 0, 1, '2022-11-07 22:01:48', 0, '2022-11-07 22:01:48', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (181, -1, 46, '操作', 'handle', '', 1, 0, 1, '2022-11-07 22:42:33', 0, '2022-11-07 22:42:33', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (182, -1, 46, '状态', 'status', '', 1, 0, 1, '2022-11-07 22:42:34', 0, '2022-11-07 22:42:34', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (183, -1, 46, 'ID', 'jobId', '', 1, 0, 1, '2022-11-07 22:42:46', 0, '2022-11-07 22:42:46', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (184, -1, 46, '任务名称', 'jobName', '', 1, 0, 1, '2022-11-07 22:42:57', 0, '2022-11-07 22:42:57', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (185, -1, 46, '任务描述', 'remark', '', 1, 0, 1, '2022-11-07 22:43:06', 0, '2022-11-07 22:43:06', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (186, -1, 46, '任务分组', 'jobGroup', '', 1, 0, 1, '2022-11-07 22:43:13', 0, '2022-11-07 22:43:13', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (187, -1, 46, '任务方法名', 'invokeTarget', '', 1, 0, 1, '2022-11-07 22:43:24', 0, '2022-11-07 22:43:24', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (188, -1, 46, 'cron执行表达式', 'cronExpression', '', 1, 0, 1, '2022-11-07 22:43:36', 0, '2022-11-07 22:43:36', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (189, -1, 48, '操作', 'handle', '', 1, 0, 1, '2022-11-07 22:44:56', 0, '2022-11-07 22:44:56', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (190, -1, 48, '状态', 'status', '', 1, 0, 1, '2022-11-07 22:44:57', 0, '2022-11-07 22:44:57', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (191, -1, 48, '接口名称', 'name', '', 1, 0, 1, '2022-11-07 22:45:03', 0, '2022-11-07 22:45:03', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (192, -1, 48, '接口地址', 'address', '', 1, 0, 1, '2022-11-07 22:45:13', 0, '2022-11-07 22:45:13', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (193, -1, 85, '操作', 'handle', '', 1, 0, 1, '2022-11-07 22:47:35', 0, '2022-11-07 22:47:35', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (194, -1, 85, '状态', 'status', '', 1, 0, 1, '2022-11-07 22:48:56', 0, '2022-11-07 22:48:56', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (195, -1, 85, '创建时间', 'createdAt', '', 1, 0, 1, '2022-11-07 22:48:57', 0, '2022-11-07 22:48:57', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (196, -1, 85, '城市名称', 'name', '', 1, 0, 1, '2022-11-07 22:49:11', 0, '2022-11-07 22:49:11', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (197, -1, 85, '城市编号', 'code', '', 1, 0, 1, '2022-11-07 22:49:19', 0, '2022-11-07 22:49:19', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (198, -1, 39, '状态', 'status', '', 1, 0, 1, '2022-11-08 23:38:11', 0, '2022-11-08 23:38:11', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (199, -1, 39, '描述', 'desc', '', 1, 0, 1, '2022-11-08 23:38:21', 0, '2022-11-08 23:38:21', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (200, -1, 66, 'ID', 'id', '', 1, 0, 1, '2022-11-11 15:10:17', 0, '2022-11-11 15:10:17', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (201, -1, 66, '操作', 'auth', '', 1, 0, 1, '2022-11-11 15:12:19', 0, '2022-11-11 15:12:19', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (202, -1, 70, '操作', 'auth', '', 1, 0, 1, '2022-11-11 15:35:32', 0, '2022-11-11 15:35:32', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (203, -1, 70, 'ID', 'id', '', 1, 0, 1, '2022-11-11 15:36:02', 0, '2022-11-11 15:36:02', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (204, -1, 70, '名称', 'name', '', 1, 0, 1, '2022-11-11 15:36:11', 0, '2022-11-11 15:36:11', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (205, -1, 70, '类型', 'types', '', 1, 0, 1, '2022-11-11 15:36:22', 0, '2022-11-11 15:36:22', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (206, -1, 70, '地址', 'addr', '', 1, 0, 1, '2022-11-11 15:36:32', 0, '2022-11-11 15:36:32', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (207, -1, 70, '创建时间', 'createdAt', '', 1, 0, 1, '2022-11-11 15:36:51', 0, '2022-11-11 15:36:51', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (208, -1, 70, '状态', 'status', '', 1, 0, 1, '2022-11-11 15:37:43', 0, '2022-11-11 15:37:43', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (209, -1, 37, '标识', 'key', '', 1, 0, 1, '2022-11-11 17:12:39', 1, '2022-11-12 01:12:53', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (210, -1, 39, '标识', 'key', '', 1, 0, 1, '2022-11-12 01:13:13', 0, '2022-11-12 01:13:13', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (211, -1, 115, 'ID', 'id', '', 1, 0, 1, '2022-11-15 14:45:10', 0, '2022-11-15 14:45:10', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (212, -1, 115, '告警类型', 'type', '', 1, 0, 1, '2022-11-15 14:45:23', 0, '2022-11-15 14:45:23', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (213, -1, 115, '规则名称', 'ruleName', '', 1, 0, 1, '2022-11-15 14:45:34', 0, '2022-11-15 14:45:34', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (214, -1, 115, '规则级别', 'alarmLevel', '', 1, 0, 1, '2022-11-15 14:45:41', 0, '2022-11-15 14:45:41', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (215, -1, 115, '产品标识', 'productKey', '', 1, 0, 1, '2022-11-15 14:45:49', 0, '2022-11-15 14:45:49', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (216, -1, 115, '设备标识', 'deviceKey', '', 1, 0, 1, '2022-11-15 14:45:57', 0, '2022-11-15 14:45:57', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (217, -1, 115, '告警状态', 'status', '', 1, 0, 1, '2022-11-15 14:46:06', 0, '2022-11-15 14:46:06', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (218, -1, 115, '告警时间', 'createdAt', '', 1, 0, 1, '2022-11-15 14:46:15', 0, '2022-11-15 14:46:15', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (219, -1, 115, '操作', 'handle', '', 1, 0, 1, '2022-11-15 14:46:23', 0, '2022-11-15 14:46:23', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (220, -1, 114, '图标', 'image', '', 1, 0, 1, '2022-11-23 10:31:35', 0, '2022-11-23 10:31:35', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (221, -1, 114, '名称', 'name', '', 1, 0, 1, '2022-11-23 10:32:14', 0, '2022-11-23 10:32:14', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (222, -1, 114, '属性', 'alarm', '', 1, 0, 1, '2022-11-23 10:33:14', 0, '2022-11-23 10:33:14', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (223, -1, 114, '状态', 'liststatus', '', 1, 0, 1, '2022-11-23 10:33:36', 0, '2022-11-23 10:33:36', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (224, -1, 92, 'ID', 'id', '', 1, 0, 1, '2022-11-28 10:21:34', 0, '2022-11-28 10:21:34', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (225, -1, 92, '告警类型', 'type', '', 1, 0, 1, '2022-11-28 10:21:46', 0, '2022-11-28 10:21:46', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (226, -1, 92, '规则名称', 'ruleName', '', 1, 0, 1, '2022-11-28 10:22:05', 0, '2022-11-28 10:22:05', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (227, -1, 92, '规则级别', 'alarmLevel', '', 1, 0, 1, '2022-11-28 10:22:27', 0, '2022-11-28 10:22:27', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (228, -1, 92, '产品标识', 'productKey', '', 1, 0, 1, '2022-11-28 10:22:52', 0, '2022-11-28 10:22:52', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (229, -1, 92, '设别标识', 'deviceKey', '', 1, 0, 1, '2022-11-28 10:23:24', 0, '2022-11-28 10:23:24', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (230, -1, 92, '告警状态', 'status', '', 1, 0, 1, '2022-11-28 10:23:39', 0, '2022-11-28 10:23:39', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (231, -1, 92, '告警时间', 'createdAt', '', 1, 0, 1, '2022-11-28 02:23:54', 1, '2022-11-28 10:24:01', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (232, -1, 92, '操作', 'handle', '', 1, 0, 1, '2022-11-28 10:24:14', 0, '2022-11-28 10:24:14', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (233, -1, 12, '部门名称', 'deptName', '', 1, 0, 1, '2022-12-04 21:57:22', 0, '2022-12-04 21:57:22', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (234, -1, 41, '标题', 'title', '', 1, 0, 1, '2022-12-07 14:39:28', 0, '2022-12-07 14:39:28', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (235, -1, 41, '描述', 'desc', '', 1, 0, 1, '2022-12-07 14:39:38', 0, '2022-12-07 14:39:38', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (236, -1, 41, '图标', 'image', '', 1, 0, 1, '2022-12-07 14:40:56', 0, '2022-12-07 14:40:56', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (237, -1, 41, '操作', 'handle', '', 1, 0, 1, '2022-12-07 14:40:58', 0, '2022-12-07 14:40:58', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (238, -1, 42, 'ID', 'id', '', 1, 0, 1, '2022-12-15 19:39:00', 1, '2022-12-15 19:39:19', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (239, -1, 42, '标题', 'title', '', 1, 0, 1, '2022-12-15 19:39:34', 0, '2022-12-15 19:39:34', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (240, -1, 42, '状态', 'status', '', 1, 0, 1, '2022-12-15 19:39:54', 0, '2022-12-15 19:39:54', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (241, -1, 42, '创建时间', 'createdAt', '', 1, 0, 1, '2022-12-15 19:39:56', 0, '2022-12-15 19:39:56', 0, NULL);
INSERT INTO `sys_menu_column` (`id`, `parent_id`, `menu_id`, `name`, `code`, `description`, `status`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (242, -1, 42, '操作', 'handle', '', 1, 0, 1, '2022-12-15 19:39:58', 0, '2022-12-15 19:39:58', 0, NULL);

COMMIT;

-- ----------------------------
-- Table structure for sys_notifications
-- ----------------------------
DROP TABLE IF EXISTS `sys_notifications`;
CREATE TABLE `sys_notifications` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) NOT NULL COMMENT '标题',
  `doc` varchar(200) NOT NULL COMMENT '描述',
  `source` varchar(50) NOT NULL COMMENT '消息来源',
  `types` varchar(50) NOT NULL COMMENT '类型',
  `created_at` datetime NOT NULL COMMENT '发送时间',
  `status` tinyint(1) NOT NULL COMMENT '0，未读，1，已读',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='消息中心';


-- ----------------------------
-- Table structure for sys_oper_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_oper_log`;
CREATE TABLE `sys_oper_log` (
  `oper_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '日志主键',
  `title` varchar(50) DEFAULT '' COMMENT '模块标题',
  `business_type` int(11) DEFAULT '0' COMMENT '业务类型（0其它 1新增 2修改 3删除）',
  `method` varchar(100) DEFAULT '' COMMENT '方法名称',
  `request_method` varchar(10) DEFAULT '' COMMENT '请求方式',
  `operator_type` int(11) DEFAULT '0' COMMENT '操作类别（0其它 1后台用户 2手机端用户）',
  `oper_name` varchar(50) DEFAULT '' COMMENT '操作人员',
  `dept_name` varchar(50) DEFAULT '' COMMENT '部门名称',
  `oper_url` varchar(500) DEFAULT '' COMMENT '请求URL',
  `oper_ip` varchar(50) DEFAULT '' COMMENT '主机地址',
  `oper_location` varchar(255) DEFAULT '' COMMENT '操作地点',
  `oper_param` text COMMENT '请求参数',
  `json_result` text COMMENT '返回参数',
  `status` int(11) DEFAULT '0' COMMENT '操作状态（0异常 1正常）',
  `error_msg` varchar(2000) DEFAULT '' COMMENT '错误消息',
  `oper_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '操作时间',
  PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=68867 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='操作日志记录';


-- ----------------------------
-- Table structure for sys_organization
-- ----------------------------
DROP TABLE IF EXISTS `sys_organization`;
CREATE TABLE `sys_organization` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '组织ID',
  `parent_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '父组织id',
  `ancestors` varchar(1000) NOT NULL DEFAULT '' COMMENT '祖级列表',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '组织名称',
  `number` varchar(50) NOT NULL COMMENT '组织编号',
  `order_num` int(11) DEFAULT '0' COMMENT '显示顺序',
  `leader` varchar(20) DEFAULT NULL COMMENT '负责人',
  `phone` varchar(11) DEFAULT NULL COMMENT '联系电话',
  `email` varchar(50) DEFAULT NULL COMMENT '邮箱',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '组织状态（0停用 1正常）',
  `is_deleted` int(11) NOT NULL COMMENT '是否删除 0未删除 1已删除',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `created_by` int(10) unsigned DEFAULT '0' COMMENT '创建人',
  `updated_by` int(11) DEFAULT NULL COMMENT '修改人',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_by` int(11) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='组织表';

-- ----------------------------
-- Records of sys_organization
-- ----------------------------
BEGIN;
INSERT INTO `sys_organization` (`id`, `parent_id`, `ancestors`, `name`, `number`, `order_num`, `leader`, `phone`, `email`, `status`, `is_deleted`, `created_at`, `created_by`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (1, -1, '1', 'Sagoo', '', 1, 'Sagoo', '15888888888', 'Sagoo@qq.com', 1, 1, '2022-08-07 21:22:24', 1, NULL, '2022-08-10 22:50:14', 1, '2022-08-10 22:50:14');
INSERT INTO `sys_organization` (`id`, `parent_id`, `ancestors`, `name`, `number`, `order_num`, `leader`, `phone`, `email`, `status`, `is_deleted`, `created_at`, `created_by`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (2, -1, '2', '', '', 0, 'asdf', '13354231895', '123', 1, 1, '2022-08-10 22:37:23', 1, 0, '2022-08-10 22:50:28', 1, '2022-08-10 22:50:28');
INSERT INTO `sys_organization` (`id`, `parent_id`, `ancestors`, `name`, `number`, `order_num`, `leader`, `phone`, `email`, `status`, `is_deleted`, `created_at`, `created_by`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (3, -1, '3', '测试组织', '', 0, '杨立志', '13354231895', '12123', 1, 1, '2022-08-10 22:39:36', 1, 0, '2022-08-10 23:11:22', 1, '2022-08-10 23:11:22');
INSERT INTO `sys_organization` (`id`, `parent_id`, `ancestors`, `name`, `number`, `order_num`, `leader`, `phone`, `email`, `status`, `is_deleted`, `created_at`, `created_by`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (4, -1, '4', '测试组织2', '', 0, '杨立志', '13354231895', '111@qq.com', 1, 1, '2022-08-10 22:41:46', 1, 0, '2022-08-10 23:11:24', 1, '2022-08-10 23:11:24');
INSERT INTO `sys_organization` (`id`, `parent_id`, `ancestors`, `name`, `number`, `order_num`, `leader`, `phone`, `email`, `status`, `is_deleted`, `created_at`, `created_by`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (5, -1, '5', '辽宁区域', '', 2, '123', '', '', 1, 0, '2022-08-08 22:54:57', 1, 1, '2022-08-10 23:19:16', 0, NULL);
INSERT INTO `sys_organization` (`id`, `parent_id`, `ancestors`, `name`, `number`, `order_num`, `leader`, `phone`, `email`, `status`, `is_deleted`, `created_at`, `created_by`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (6, 5, '5,6', '测试2', '', 0, '一直', '', '', 1, 1, '2022-08-10 23:07:58', 1, 0, '2022-08-10 23:10:35', 1, '2022-08-10 23:10:35');
INSERT INTO `sys_organization` (`id`, `parent_id`, `ancestors`, `name`, `number`, `order_num`, `leader`, `phone`, `email`, `status`, `is_deleted`, `created_at`, `created_by`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (7, 5, '5,7', '沈阳沙果', '', 0, '12123', '', '', 1, 0, '2022-08-09 07:10:42', 1, 1, '2022-12-15 23:42:28', 0, NULL);
INSERT INTO `sys_organization` (`id`, `parent_id`, `ancestors`, `name`, `number`, `order_num`, `leader`, `phone`, `email`, `status`, `is_deleted`, `created_at`, `created_by`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (8, -1, '8', '辽宁区域', '', 3, '阿什顿发', '', '', 1, 1, '2022-08-09 23:11:27', 1, 1, '2022-08-10 23:19:08', 1, '2022-08-10 23:19:08');
INSERT INTO `sys_organization` (`id`, `parent_id`, `ancestors`, `name`, `number`, `order_num`, `leader`, `phone`, `email`, `status`, `is_deleted`, `created_at`, `created_by`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (9, -1, '9', '333', '', 1, '3333', '', '', 1, 0, '2022-08-10 15:11:34', 1, 1, '2022-08-10 23:11:46', 0, NULL);
INSERT INTO `sys_organization` (`id`, `parent_id`, `ancestors`, `name`, `number`, `order_num`, `leader`, `phone`, `email`, `status`, `is_deleted`, `created_at`, `created_by`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (10, -1, '10', 'c收到', '', 3, '212', '', '', 1, 1, '2022-08-10 23:34:54', 1, 0, '2022-08-12 20:31:16', 1, '2022-08-12 20:31:16');
INSERT INTO `sys_organization` (`id`, `parent_id`, `ancestors`, `name`, `number`, `order_num`, `leader`, `phone`, `email`, `status`, `is_deleted`, `created_at`, `created_by`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (11, -1, '11', '集团总部', '', 0, '沙果', '', '', 1, 0, '2022-08-11 09:35:28', 1, 0, '2022-08-11 09:35:28', 0, NULL);
INSERT INTO `sys_organization` (`id`, `parent_id`, `ancestors`, `name`, `number`, `order_num`, `leader`, `phone`, `email`, `status`, `is_deleted`, `created_at`, `created_by`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (12, -1, '12', '丹东区域', '', 0, '王馨墨', '', '', 1, 0, '2022-09-19 18:40:35', 1, 1, '2022-11-15 11:45:26', 0, NULL);
INSERT INTO `sys_organization` (`id`, `parent_id`, `ancestors`, `name`, `number`, `order_num`, `leader`, `phone`, `email`, `status`, `is_deleted`, `created_at`, `created_by`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (13, -1, '13', '1111', '', 0, '111', '', '', 1, 1, '2022-10-19 00:16:15', 1, 0, '2022-10-19 00:16:23', 1, '2022-10-19 00:16:23');
COMMIT;

-- ----------------------------
-- Table structure for sys_plugins
-- ----------------------------
DROP TABLE IF EXISTS `sys_plugins`;
CREATE TABLE `sys_plugins` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `types` varchar(50) NOT NULL COMMENT '插件类型',
  `name` varchar(100) NOT NULL COMMENT '名称',
  `title` varchar(100) NOT NULL COMMENT '标题',
  `intro` varchar(200) NOT NULL COMMENT '介绍',
  `version` varchar(50) NOT NULL COMMENT '版本',
  `author` varchar(100) NOT NULL COMMENT '作者',
  `status` tinyint(1) NOT NULL COMMENT '状态',
  `start_time` datetime NOT NULL COMMENT '启动时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='系统插件表';

-- ----------------------------
-- Records of sys_plugins
-- ----------------------------
BEGIN;
INSERT INTO `sys_plugins` (`id`, `types`, `name`, `title`, `intro`, `version`, `author`, `status`, `start_time`) VALUES (2, 'protocol', 'modbus', 'Modbus TCP协议', '对modbus TCP模式的设备进行数据采集', '0.01', 'Microrain', 1, '2022-12-03 10:37:45');
INSERT INTO `sys_plugins` (`id`, `types`, `name`, `title`, `intro`, `version`, `author`, `status`, `start_time`) VALUES (3, 'protocol', 'tgn52', 'TG-N5 v2设备协议', '对TG-N5插座设备进行数据采集v2', '0.01', 'Microrain', 1, '2022-12-03 10:37:45');
COMMIT;

-- ----------------------------
-- Table structure for sys_plugins_config
-- ----------------------------
DROP TABLE IF EXISTS `sys_plugins_config`;
CREATE TABLE `sys_plugins_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` varchar(50) NOT NULL COMMENT '插件类型',
  `name` varchar(50) NOT NULL COMMENT '插件名称',
  `value` varchar(300) NOT NULL COMMENT '配置内容',
  `doc` varchar(300) NOT NULL COMMENT '配置说明',
  PRIMARY KEY (`id`),
  UNIQUE KEY `typeandname` (`type`,`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8 COMMENT='插件配置表';

-- ----------------------------
-- Records of sys_plugins_config
-- ----------------------------
BEGIN;
INSERT INTO `sys_plugins_config` (`id`, `type`, `name`, `value`, `doc`) VALUES (1, 'aaaa', 'bbbb', '1231sssssseeeeee31231231231', 'fadsfsdfsdf');
INSERT INTO `sys_plugins_config` (`id`, `type`, `name`, `value`, `doc`) VALUES (4, 'notice', 'work_weixin', 'Corpid: \"ww3da9f97a7f1babc0\"\nAgentID: 1000002\nSecret: \"bo6nFN38dcYpZyKD7tBXnXDMCkaDCta2m3pPO6lC1U8\"\nToken: \"mnRLL3LBiiaxskhQD5\"\nEncodingAESKey: \"4975mx8v13S8QTNyOVcjLfw2EyJPyJTurXPVpympjkC\"', '<p>配置内容配置内容</p>');
INSERT INTO `sys_plugins_config` (`id`, `type`, `name`, `value`, `doc`) VALUES (6, 'notice', 'dingding', 'AppKey: \"sdfadfasdfasdfasdfasdf\"\nAppSecret: \"ewrerwerwerrwerwerwerwer\"', '<p>width: 40px;</p>');
INSERT INTO `sys_plugins_config` (`id`, `type`, `name`, `value`, `doc`) VALUES (7, 'notice', 'sms', 'ProviderName: \"alisms\"\nTitle: \"阿里云\"\nRegionId: \"cn-hangzhou\"\nAccessKeyId: \"LTAI7TGCAqUzuI6B\"\naccessSecret: \"AseYR5htpq7BTwfH6w73wqLZIwYjxa\"\nSignName: \"沙果学习\"', '<p>ProviderName 的参数为短信提供商的代码；alisms：阿里云；<span style=\"color: var(--el-text-color-regular); font-size: var(--el-dialog-content-font-size);\">tencentcloud：腾讯云</span></p>');
INSERT INTO `sys_plugins_config` (`id`, `type`, `name`, `value`, `doc`) VALUES (10, 'notice', 'mail', 'MailHost: \"smtp.qq.com\"\nMailPort: \"465\"\nMailUser: \"xinjy@qq.com\"\nMailPass: \"zdqkqrdzplnabiig\"', '');
INSERT INTO `sys_plugins_config` (`id`, `type`, `name`, `value`, `doc`) VALUES (12, 'notice', 'webhook', 'webhook:\n  - PayloadURL: \"http://127.0.0.1:8180/test/webhook\"\n    Secret: \"aaaadfasdfasf\"\n  - PayloadURL: \"http://127.0.0.1:8180/test/webhook22\"\n    Secret: \"aaaadfasdfasf22\"\n  - PayloadURL: \"http://127.0.0.1:8180/test/webhook33\"\n    Secret: \"aaaadfasdfasf333\"', '');
INSERT INTO `sys_plugins_config` (`id`, `type`, `name`, `value`, `doc`) VALUES (13, 'notice', 'voice', 'wwwwww', '');
COMMIT;

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post` (
  `post_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `parent_id` int(11) NOT NULL COMMENT '父ID',
  `post_code` varchar(64) NOT NULL COMMENT '岗位编码',
  `post_name` varchar(50) NOT NULL COMMENT '岗位名称',
  `post_sort` int(11) NOT NULL COMMENT '显示顺序',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态（0正常 1停用）',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  `is_deleted` int(11) NOT NULL COMMENT '是否删除 0未删除 1已删除',
  `created_by` int(10) unsigned DEFAULT '0' COMMENT '创建人',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_by` int(10) unsigned DEFAULT '0' COMMENT '修改人',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_by` int(11) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='岗位信息表';

-- ----------------------------
-- Records of sys_post
-- ----------------------------
BEGIN;
INSERT INTO `sys_post` (`post_id`, `parent_id`, `post_code`, `post_name`, `post_sort`, `status`, `remark`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (1, -1, '01', '技术总监', 1, 0, 'Sagoo IOT技术负责人', 1, 1, '2022-07-09 00:25:06', 1, '2022-08-11 15:25:00', 1, '2022-08-11 15:25:01');
INSERT INTO `sys_post` (`post_id`, `parent_id`, `post_code`, `post_name`, `post_sort`, `status`, `remark`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (2, -1, 'G20220811121447', 'ass 打发斯蒂芬', 0, 1, '123123', 1, 1, '2022-08-11 04:14:47', 1, '2022-08-11 12:18:45', 1, '2022-08-11 12:18:46');
INSERT INTO `sys_post` (`post_id`, `parent_id`, `post_code`, `post_name`, `post_sort`, `status`, `remark`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (3, -1, 'G20220811121542', '阿什顿发', 1, 1, '123123', 1, 1, '2022-08-11 12:15:42', 0, '2022-08-11 12:19:27', 1, '2022-08-11 12:19:27');
INSERT INTO `sys_post` (`post_id`, `parent_id`, `post_code`, `post_name`, `post_sort`, `status`, `remark`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (4, -1, 'G20220811121552', '阿什顿发22', 33, 1, '3', 0, 1, '2022-08-10 12:15:52', 1, '2022-09-15 17:04:07', 0, NULL);
INSERT INTO `sys_post` (`post_id`, `parent_id`, `post_code`, `post_name`, `post_sort`, `status`, `remark`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (5, 4, 'G20220811152510', 'gangwei', 0, 1, '123', 0, 1, '2022-08-10 23:25:10', 1, '2022-12-15 23:42:34', 0, NULL);
INSERT INTO `sys_post` (`post_id`, `parent_id`, `post_code`, `post_name`, `post_sort`, `status`, `remark`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (6, -1, 'G20220812235425', '岗位2', 0, 1, '', 0, 1, '2022-08-12 23:54:25', 0, '2022-08-12 23:54:25', 0, NULL);
INSERT INTO `sys_post` (`post_id`, `parent_id`, `post_code`, `post_name`, `post_sort`, `status`, `remark`, `is_deleted`, `created_by`, `created_at`, `updated_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (7, -1, 'G20220812235433', '岗位223', 0, 1, '', 0, 1, '2022-08-12 23:54:33', 0, '2022-08-12 23:54:33', 0, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `parent_id` int(11) NOT NULL COMMENT '父ID',
  `list_order` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `data_scope` tinyint(3) unsigned DEFAULT '3' COMMENT '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `status` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '状态;0:禁用;1:正常',
  `is_deleted` int(11) NOT NULL COMMENT '是否删除 0未删除 1已删除',
  `create_by` int(11) unsigned DEFAULT '0' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建日期',
  `update_by` int(11) unsigned DEFAULT '0' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '修改日期',
  `deleted_by` int(11) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='角色表';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_role` (`id`, `parent_id`, `list_order`, `name`, `data_scope`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (1, -1, 1, '超级管理员', 1, '超级管理员拥有所有权限', 1, 0, 1, '2022-08-07 21:30:28', 0, NULL, NULL, NULL);
INSERT INTO `sys_role` (`id`, `parent_id`, `list_order`, `name`, `data_scope`, `remark`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (7, 3, 0, '示例演示', 1, '示例演示勿动', 1, 0, 1, '2022-08-09 08:22:03', 1, '2022-11-06 01:00:32', 0, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_dept`;
CREATE TABLE `sys_role_dept` (
  `role_id` bigint(20) NOT NULL COMMENT '角色ID',
  `dept_id` bigint(20) NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`role_id`,`dept_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='角色和部门关联表';

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------
BEGIN;
INSERT INTO `sys_role_dept` (`role_id`, `dept_id`) VALUES (3, 6);
COMMIT;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `user_types` varchar(255) DEFAULT NULL COMMENT '系统 system 企业 company',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '中国手机不带国家代码，国际手机号格式为：国家代码-手机号',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `birthday` int(11) NOT NULL DEFAULT '0' COMMENT '生日',
  `user_password` varchar(255) NOT NULL DEFAULT '' COMMENT '登录密码;cmf_password加密',
  `user_salt` char(10) NOT NULL COMMENT '加密盐',
  `user_email` varchar(100) NOT NULL DEFAULT '' COMMENT '用户登录邮箱',
  `sex` tinyint(4) NOT NULL DEFAULT '0' COMMENT '性别;0:保密,1:男,2:女',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像',
  `dept_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '部门id',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `is_admin` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否后台管理员 1 是  0   否',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '联系地址',
  `describe` varchar(255) NOT NULL DEFAULT '' COMMENT ' 描述信息',
  `last_login_ip` varchar(15) NOT NULL DEFAULT '' COMMENT '最后登录ip',
  `last_login_time` datetime DEFAULT NULL COMMENT '最后登录时间',
  `status` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '用户状态;0:禁用,1:正常,2:未验证',
  `is_deleted` int(11) NOT NULL COMMENT '是否删除 0未删除 1已删除',
  `create_by` int(10) unsigned DEFAULT '0' COMMENT '创建者',
  `created_at` datetime DEFAULT NULL COMMENT '创建日期',
  `update_by` int(10) unsigned DEFAULT '0' COMMENT '更新者',
  `updated_at` datetime DEFAULT NULL COMMENT '修改日期',
  `deleted_by` int(11) DEFAULT NULL COMMENT '删除人',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_login` (`user_name`,`deleted_at`) USING BTREE,
  UNIQUE KEY `mobile` (`mobile`,`deleted_at`) USING BTREE,
  KEY `user_nickname` (`user_nickname`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='用户表';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO `sys_user` (`id`, `user_name`, `user_types`, `mobile`, `user_nickname`, `birthday`, `user_password`, `user_salt`, `user_email`, `sex`, `avatar`, `dept_id`, `remark`, `is_admin`, `address`, `describe`, `last_login_ip`, `last_login_time`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (1, 'admin', '1', '15888888889', '超级管理员', 1990, '3a66af5d8753bf48bf97dd3be53d1956', 'IIsQ2vKfaV', 'yxh669@qq.com', 1, 'https://zhgy.sagoo.cn/base-api/upload_file/2022-11-11/co9dilmc7lt8ahwaa6.png', 6, '', 1, 'asdasfdsaf大发放打发士大夫发按时', '描述信息', '112.41.4.221', '2023-01-01 18:22:18', 1, 0, 1, '2022-08-03 21:33:30', 10, '2023-01-01 18:22:18', 0, NULL);
INSERT INTO `sys_user` (`id`, `user_name`, `user_types`, `mobile`, `user_nickname`, `birthday`, `user_password`, `user_salt`, `user_email`, `sex`, `avatar`, `dept_id`, `remark`, `is_admin`, `address`, `describe`, `last_login_ip`, `last_login_time`, `status`, `is_deleted`, `create_by`, `created_at`, `update_by`, `updated_at`, `deleted_by`, `deleted_at`) VALUES (10, 'demo', 'system', '18711111111', '演示示例账号', 0, 'e1f6ef3bd5cbf35fdb05b686b111d16f', 'k5jjjB3VU8', '', 0, 'https://zhgy.sagoo.cn/base-api/upload_file/2022-12-04/cot10byfqd39ykjnno.png', 3, '', 1, '', '', '119.85.96.1', '2022-12-31 18:00:04', 1, 0, 1, '2022-11-03 00:19:50', 1, '2022-12-31 18:00:04', 0, NULL);

COMMIT;

-- ----------------------------
-- Table structure for sys_user_online
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_online`;
CREATE TABLE `sys_user_online` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `uuid` char(32) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户标识',
  `key` varchar(255) NOT NULL,
  `token` varchar(255) CHARACTER SET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT '' COMMENT '用户token',
  `created_at` datetime NOT NULL COMMENT '登录时间',
  `user_name` varchar(255) NOT NULL COMMENT '用户名',
  `ip` varchar(120) NOT NULL DEFAULT '' COMMENT '登录ip',
  `explorer` varchar(30) NOT NULL DEFAULT '' COMMENT '浏览器',
  `os` varchar(30) NOT NULL DEFAULT '' COMMENT '操作系统',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_token` (`token`) USING BTREE,
  UNIQUE KEY `token` (`token`)
) ENGINE=InnoDB AUTO_INCREMENT=849 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户在线状态表';

-- ----------------------------
-- Records of sys_user_online
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_online` (`id`, `uuid`, `key`, `token`, `created_at`, `user_name`, `ip`, `explorer`, `os`) VALUES (848, 'tu0rkg0fwt0cpgs2b69xkdo1006fzb9h', '1-21232f297a57a5a743894a0e4a801fc3-e1db92661861f7772d7eabe1e4328e90', '5rrLCPtzPM4tnvlHq+0iav2BDmIrd9QCru7zhgXMkRfHyvxQikCz+SxJ1qEgNa5kL7XO5cyg7CCmQCP5lPgqCyzikx8yzR1capYtD08oQHlt9WScmPvI+BS2JCQgp0BOZZcDnNS9YL46Xi06nz1BXA==', '2023-01-01 18:22:18', 'admin', '112.41.4.221', 'Chrome', 'Intel Mac OS X 10_15_7');
COMMIT;

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_post`;
CREATE TABLE `sys_user_post` (
  `user_id` int(11) NOT NULL COMMENT '用户ID',
  `post_id` int(11) NOT NULL COMMENT '岗位ID',
  PRIMARY KEY (`user_id`,`post_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=COMPACT COMMENT='用户与岗位关联表';

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_post` (`user_id`, `post_id`) VALUES (1, 6);
INSERT INTO `sys_user_post` (`user_id`, `post_id`) VALUES (10, 5);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role` (
  `user_id` int(11) NOT NULL,
  `role_id` int(11) NOT NULL,
  PRIMARY KEY (`user_id`,`role_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='用户和角色关联表';

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_role` (`user_id`, `role_id`) VALUES (1, 1);
INSERT INTO `sys_user_role` (`user_id`, `role_id`) VALUES (10, 7);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
