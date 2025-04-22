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

 Date: 14/04/2025 14:16:32
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for monitor_target
-- ----------------------------
DROP TABLE IF EXISTS `monitor_target`;
CREATE TABLE `monitor_target`  (
                                   `id` int NOT NULL AUTO_INCREMENT,
                                   `ip` char(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '',
                                   `domain` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '域名',
                                   `condition` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '查询条件',
                                   `is_deleted` int NULL DEFAULT 0 COMMENT '是否被删除。0：否；1：是',
                                   `create_time` datetime NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
                                   PRIMARY KEY (`id`) USING BTREE,
                                   UNIQUE INDEX `idx_unique`(`ip` ASC, `domain` ASC, `condition` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of monitor_target
-- ----------------------------
INSERT INTO `monitor_target` VALUES (1, '1.0.0.0', NULL, NULL, 1, '2025-04-14 11:19:22');
INSERT INTO `monitor_target` VALUES (2, '1.0.0.1', '', '', 0, '2025-04-14 12:00:56');
INSERT INTO `monitor_target` VALUES (3, '1.0.0.2', '', '', 1, '2025-04-14 11:39:55');
INSERT INTO `monitor_target` VALUES (4, '220.181.33.115', 'smartprogram.baidu.com', '', 0, '2025-04-14 11:39:28');
INSERT INTO `monitor_target` VALUES (11, '1.0.0.6', '', '', 0, '2025-04-14 12:05:38');

SET FOREIGN_KEY_CHECKS = 1;
