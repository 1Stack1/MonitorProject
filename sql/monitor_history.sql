/*
 Navicat Premium Dump SQL

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80034 (8.0.34)
 Source Host           : localhost:3306
 Source Schema         : monitor

 Target Server Type    : MySQL
 Target Server Version : 80034 (8.0.34)
 File Encoding         : 65001

 Date: 14/04/2025 14:16:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for monitor_history
-- ----------------------------
DROP TABLE IF EXISTS `monitor_history`;
CREATE TABLE `monitor_history`  (
                                    `id` int(10) UNSIGNED ZEROFILL NOT NULL AUTO_INCREMENT,
                                    `target_id` int NULL DEFAULT NULL COMMENT '监控目标id',
                                    `monitor_start_time` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '监控开始时间',
                                    `asset_count` int NULL DEFAULT NULL COMMENT '每日资产数量',
                                    `changed_count` int NULL DEFAULT NULL COMMENT '变化数量',
                                    `is_deleted` int NULL DEFAULT 0 COMMENT '是否被删除。0：否；1：是',
                                    `create_time` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
                                    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 41 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of monitor_history
-- ----------------------------
INSERT INTO `monitor_history` VALUES (0000000001, 2, '2025-04-11 10:19:14', 3524, 0, 0, '2025-04-11 10:19:14');
INSERT INTO `monitor_history` VALUES (0000000002, 2, '2025-04-11 10:21:23', 3524, 0, 0, '2025-04-11 10:21:23');
INSERT INTO `monitor_history` VALUES (0000000003, 3, '2025-04-11 10:21:23', 3524, 0, 0, '2025-04-11 10:21:23');
INSERT INTO `monitor_history` VALUES (0000000004, 2, '2025-04-11 10:35:33', 3524, 0, 0, '2025-04-11 10:35:33');
INSERT INTO `monitor_history` VALUES (0000000005, 3, '2025-04-11 10:35:33', 0, 0, 0, '2025-04-11 10:35:33');
INSERT INTO `monitor_history` VALUES (0000000006, 2, '2025-04-11 10:35:57', 3524, 0, 0, '2025-04-11 10:35:57');
INSERT INTO `monitor_history` VALUES (0000000007, 3, '2025-04-11 10:36:02', 3524, 0, 0, '2025-04-11 10:36:02');
INSERT INTO `monitor_history` VALUES (0000000008, 2, '2025-04-11 10:36:38', 3524, 0, 0, '2025-04-11 10:36:38');
INSERT INTO `monitor_history` VALUES (0000000009, 3, '2025-04-11 10:36:44', 3524, 0, 0, '2025-04-11 10:36:44');
INSERT INTO `monitor_history` VALUES (0000000011, 3, '2025-04-11 11:09:12', 3524, 0, 0, '2025-04-11 11:09:12');
INSERT INTO `monitor_history` VALUES (0000000014, 4, '2025-04-14 09:42:54', 106, 20, 0, '2025-04-14 09:42:54');
INSERT INTO `monitor_history` VALUES (0000000015, 4, '2025-04-14 09:51:02', 106, 20, 0, '2025-04-14 09:51:02');
INSERT INTO `monitor_history` VALUES (0000000016, 4, '2025-04-14 09:52:22', 106, 20, 0, '2025-04-14 09:52:22');
INSERT INTO `monitor_history` VALUES (0000000017, 4, '2025-04-14 11:40:41', 106, 0, 0, '2025-04-14 11:40:41');
INSERT INTO `monitor_history` VALUES (0000000018, 4, '2025-04-14 11:59:01', 106, 0, 0, '2025-04-14 11:59:01');
INSERT INTO `monitor_history` VALUES (0000000019, 4, '2025-04-14 11:59:11', 106, 0, 0, '2025-04-14 11:59:11');
INSERT INTO `monitor_history` VALUES (0000000020, 2, '2025-04-14 12:01:11', 3540, 16, 0, '2025-04-14 12:01:11');
INSERT INTO `monitor_history` VALUES (0000000021, 4, '2025-04-14 12:01:11', 0, -106, 0, '2025-04-14 12:01:11');
INSERT INTO `monitor_history` VALUES (0000000022, 2, '2025-04-14 12:03:51', 3541, 1, 0, '2025-04-14 12:03:51');
INSERT INTO `monitor_history` VALUES (0000000023, 4, '2025-04-14 12:03:51', 0, 0, 0, '2025-04-14 12:03:51');
INSERT INTO `monitor_history` VALUES (0000000024, 2, '2025-04-14 12:04:01', 3541, 0, 0, '2025-04-14 12:04:01');
INSERT INTO `monitor_history` VALUES (0000000025, 4, '2025-04-14 12:04:02', 106, 0, 0, '2025-04-14 12:04:02');
INSERT INTO `monitor_history` VALUES (0000000026, 2, '2025-04-14 12:05:50', 3541, 0, 0, '2025-04-14 12:05:50');
INSERT INTO `monitor_history` VALUES (0000000027, 4, '2025-04-14 12:05:54', 106, 0, 0, '2025-04-14 12:05:54');
INSERT INTO `monitor_history` VALUES (0000000028, 11, '2025-04-14 12:05:55', 170, 0, 0, '2025-04-14 12:05:55');
INSERT INTO `monitor_history` VALUES (0000000029, 2, '2025-04-14 12:06:01', 3541, 0, 0, '2025-04-14 12:06:01');
INSERT INTO `monitor_history` VALUES (0000000030, 2, '2025-04-14 12:06:01', 0, -3541, 0, '2025-04-14 12:06:01');
INSERT INTO `monitor_history` VALUES (0000000031, 4, '2025-04-14 12:06:02', 106, 0, 0, '2025-04-14 12:06:02');
INSERT INTO `monitor_history` VALUES (0000000032, 4, '2025-04-14 12:06:03', 0, -106, 0, '2025-04-14 12:06:03');
INSERT INTO `monitor_history` VALUES (0000000033, 11, '2025-04-14 12:06:03', 0, -170, 0, '2025-04-14 12:06:03');
INSERT INTO `monitor_history` VALUES (0000000034, 11, '2025-04-14 12:06:03', 170, 0, 0, '2025-04-14 12:06:03');
INSERT INTO `monitor_history` VALUES (0000000035, 2, '2025-04-14 12:06:11', 3541, 0, 0, '2025-04-14 12:06:11');
INSERT INTO `monitor_history` VALUES (0000000036, 4, '2025-04-14 12:06:12', 106, 0, 0, '2025-04-14 12:06:12');
INSERT INTO `monitor_history` VALUES (0000000037, 11, '2025-04-14 12:06:13', 0, -170, 0, '2025-04-14 12:06:13');
INSERT INTO `monitor_history` VALUES (0000000038, 2, '2025-04-14 12:09:01', 3541, 0, 0, '2025-04-14 12:09:01');
INSERT INTO `monitor_history` VALUES (0000000039, 4, '2025-04-14 12:09:02', 0, -106, 0, '2025-04-14 12:09:02');
INSERT INTO `monitor_history` VALUES (0000000040, 11, '2025-04-14 12:09:03', 170, 0, 0, '2025-04-14 12:09:03');

SET FOREIGN_KEY_CHECKS = 1;
