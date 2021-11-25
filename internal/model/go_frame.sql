/*
 Navicat Premium Data Transfer

 Source Server         : go
 Source Server Type    : MySQL
 Source Server Version : 80025
 Source Host           : localhost:13306
 Source Schema         : go_frame

 Target Server Type    : MySQL
 Target Server Version : 80025
 File Encoding         : 65001

 Date: 16/09/2021 21:58:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for blog_article
-- ----------------------------
DROP TABLE IF EXISTS `blog_article`;
CREATE TABLE `blog_article`  (
  `id` int(10) UNSIGNED ZEROFILL NOT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `cover_image_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `status` tinyint(0) NULL DEFAULT 1 COMMENT '0 禁用 1 启用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for blog_article_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_article_tag`;
CREATE TABLE `blog_article_tag`  (
  `id` int(10) UNSIGNED ZEROFILL NOT NULL,
  `article_id` int(0) NOT NULL,
  `tag_id` int(0) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for blog_service
-- ----------------------------
DROP TABLE IF EXISTS `blog_service`;
CREATE TABLE `blog_service`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `create_on` int(0) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `create_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '创建人',
  `modifyed_on` int(0) NOT NULL DEFAULT 0 COMMENT '修改时间',
  `modifyed_by` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '修改人',
  `deleted_on` int(0) NOT NULL DEFAULT 0 COMMENT '删除时间',
  `is_del` tinyint(1) NULL DEFAULT 0 COMMENT '是否删除 0 未删除 1 以删除',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for blog_tag
-- ----------------------------
DROP TABLE IF EXISTS `blog_tag`;
CREATE TABLE `blog_tag`  (
  `id` int(0) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '标签名',
  `state` tinyint(0) UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态 1 启用 0 禁用',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of blog_tag
-- ----------------------------
INSERT INTO `blog_tag` VALUES (1, 'test', 1);
INSERT INTO `blog_tag` VALUES (2, 'testAno', 1);

-- ----------------------------
-- Table structure for company
-- ----------------------------
CREATE TABLE `company` (
   `id` bigint unsigned NOT NULL AUTO_INCREMENT,
   `industry` int DEFAULT '0',
   `name` varchar(255) DEFAULT '',
   `job` varchar(255) DEFAULT '',
   `user_id` bigint DEFAULT NULL,
   PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci