/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : localhost:3306
 Source Schema         : demo_cms_1

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 06/09/2022 15:18:11
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for demo_admins
-- ----------------------------
DROP TABLE IF EXISTS `demo_admins`;
CREATE TABLE `demo_admins`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `deleted_at` datetime NULL DEFAULT NULL,
  `name` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `poss_word` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT NULL,
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `phone` varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_demo_admins_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = MyISAM AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of demo_admins
-- ----------------------------
INSERT INTO `demo_admins` VALUES (1, '2022-08-29 19:25:21', '2022-09-01 16:40:45', NULL, 'admin', 'CTm0p4GN6/BtMg==', 0, '105@qq.com', '15999999999');
INSERT INTO `demo_admins` VALUES (6, '2022-08-29 19:25:21', '2022-09-01 16:40:45', NULL, 'admin', 'CTm0p4GN6/BtMg==', 0, '105@qq.com', '15999999999');
INSERT INTO `demo_admins` VALUES (7, '2022-08-31 16:46:04', '2022-09-06 12:40:57', NULL, 'admin', 'CTm0p4GN6/BtMg==', 1, '105@qq.com', '15999999999');
INSERT INTO `demo_admins` VALUES (8, '2022-09-01 17:15:00', '2022-09-01 17:15:00', NULL, '朝阳', 'wMHzsyqb/q0FlA==', 1, '105@qq.com', '15999999999');
INSERT INTO `demo_admins` VALUES (9, '2022-09-01 17:38:23', '2022-09-01 17:38:23', NULL, '张三', 'wMHzsyqb/q0FlA==', 1, '1@qq.com', '12444444444');
INSERT INTO `demo_admins` VALUES (10, '2022-09-01 17:39:20', '2022-09-01 17:39:20', NULL, 'tom', 'CTm0p4GN6/BtMg==', 1, '105@qq.com', '15999999999');

SET FOREIGN_KEY_CHECKS = 1;
