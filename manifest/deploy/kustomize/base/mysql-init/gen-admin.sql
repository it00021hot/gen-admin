-- MySQL dump 10.13  Distrib 8.0.12, for Win64 (x86_64)
--
-- Host: 10.168.1.11    Database: gfast
-- ------------------------------------------------------
-- Server version	8.0.27

/*!40101 SET @OLD_CHARACTER_SET_CLIENT = @@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS = @@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION = @@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40103 SET @OLD_TIME_ZONE = @@TIME_ZONE */;
/*!40103 SET TIME_ZONE = '+08:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS = @@UNIQUE_CHECKS, UNIQUE_CHECKS = 0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS = 0 */;
/*!40101 SET @OLD_SQL_MODE = @@SQL_MODE, SQL_MODE = 'NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES = @@SQL_NOTES, SQL_NOTES = 0 */;

CREATE DATABASE IF NOT EXISTS gfast DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

use gfast;
--
-- Table structure for table `casbin_rule`
--

DROP TABLE IF EXISTS `casbin_rule`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
SET character_set_client = utf8mb4;
CREATE TABLE `casbin_rule`
(
    `ptype` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  DEFAULT NULL,
    `v0`    varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
    `v1`    varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
    `v2`    varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
    `v3`    varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
    `v4`    varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
    `v5`    varchar(256) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = COMPACT;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `casbin_rule`
--

LOCK TABLES `casbin_rule` WRITE;
/*!40000 ALTER TABLE `casbin_rule`
    DISABLE KEYS */;
INSERT INTO `casbin_rule`
VALUES ('p', '1', '10', 'All', '', '', ''),
       ('p', '1', '12', 'All', '', '', ''),
       ('p', '1', '13', 'All', '', '', ''),
       ('p', '1', '14', 'All', '', '', ''),
       ('g', '42', '1', '', '', '', ''),
       ('g', '42', '2', '', '', '', ''),
       ('g', '1', '1', '', '', '', ''),
       ('g', '1', '2', '', '', '', ''),
       ('g', '2', '3', '', '', '', ''),
       ('g', '2', '2', '', '', '', ''),
       ('g', '4', '2', '', '', '', ''),
       ('g', '5', '2', '', '', '', ''),
       ('g', '7', '2', '', '', '', ''),
       ('g', '8', '2', '', '', '', ''),
       ('g', '10', '2', '', '', '', ''),
       ('g', '14', '2', '', '', '', ''),
       ('g', '15', '2', '', '', '', ''),
       ('g', '16', '2', '', '', '', ''),
       ('p', '1', '1', 'All', '', '', ''),
       ('g', '6', '2', '', '', '', ''),
       ('g', '3', '2', '', '', '', '');
/*!40000 ALTER TABLE `casbin_rule`
    ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_auth_rule`
--

DROP TABLE IF EXISTS `sys_auth_rule`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
SET character_set_client = utf8mb4;
CREATE TABLE `sys_auth_rule`
(
    `id`          int unsigned                                                  NOT NULL AUTO_INCREMENT,
    `pid`         int unsigned                                                  NOT NULL DEFAULT '0' COMMENT '???ID',
    `name`        varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '????????????',
    `title`       varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '????????????',
    `icon`        varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '??????',
    `condition`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '??????',
    `remark`      varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '??????',
    `menu_type`   tinyint unsigned                                              NOT NULL DEFAULT '0' COMMENT '?????? 0?????? 1?????? 2??????',
    `weigh`       int                                                           NOT NULL DEFAULT '0' COMMENT '??????',
    `is_hide`     tinyint unsigned                                              NOT NULL DEFAULT '0' COMMENT '????????????',
    `path`        varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '????????????',
    `component`   varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '????????????',
    `is_link`     tinyint unsigned                                              NOT NULL DEFAULT '0' COMMENT '???????????? 1??? 0???',
    `module_type` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '????????????',
    `model_id`    int unsigned                                                  NOT NULL DEFAULT '0' COMMENT '??????ID',
    `is_iframe`   tinyint unsigned                                              NOT NULL DEFAULT '0' COMMENT '????????????iframe',
    `is_cached`   tinyint unsigned                                              NOT NULL DEFAULT '0' COMMENT '????????????',
    `redirect`    varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '?????????????????????',
    `is_affix`    tinyint unsigned                                              NOT NULL DEFAULT '0' COMMENT '????????????',
    `link_url`    varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '????????????',
    `created_at`  datetime                                                               DEFAULT NULL COMMENT '????????????',
    `updated_at`  datetime                                                               DEFAULT NULL COMMENT '????????????',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `name` (`name`) USING BTREE,
    KEY `pid` (`pid`) USING BTREE,
    KEY `weigh` (`weigh`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 34
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = DYNAMIC COMMENT ='???????????????';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_auth_rule`
--

LOCK TABLES `sys_auth_rule` WRITE;
/*!40000 ALTER TABLE `sys_auth_rule`
    DISABLE KEYS */;
INSERT INTO `sys_auth_rule`
VALUES (1, 0, 'api/v1/system/auth', '????????????', 'ele-Stamp', '', '', 0, 30, 0, '/system/auth',
        'layout/routerView/parent', 0, '', 0, 0, 1, '0', 0, '', '2022-03-24 15:03:37', '2022-04-14 16:29:19'),
       (2, 1, 'api/v1/system/auth/menuList', '????????????', 'ele-Calendar', '', '', 1, 0, 0, '/system/auth/menuList',
        'system/menu/index', 0, '', 0, 0, 1, '', 0, '', '2022-03-24 17:24:13', '2022-03-29 10:54:49'),
       (3, 2, 'api/v1/system/menu/add', '????????????', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '',
        '2022-03-29 16:48:43', '2022-03-29 17:05:19'),
       (4, 2, 'api/v1/system/menu/update', '????????????', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '',
        '2022-03-29 17:04:25', '2022-03-29 18:11:36'),
       (10, 1, 'api/v1/system/role/list', '????????????', 'iconfont icon-juxingkaobei', '', '', 1, 0, 0,
        '/system/auth/roleList', 'system/role/index', 0, '', 0, 0, 1, '', 0, '', '2022-03-29 18:15:03',
        '2022-03-30 10:25:34'),
       (11, 2, 'api/v1/system/menu/delete', '????????????', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '',
        '2022-04-06 14:49:10', '2022-04-06 14:49:17'),
       (12, 10, 'api/v1/system/role/add', '????????????', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '',
        '2022-04-06 14:49:46', '2022-04-06 14:49:46'),
       (13, 10, '/api/v1/system/role/edit', '????????????', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '',
        '2022-04-06 14:50:08', '2022-04-06 14:50:08'),
       (14, 10, '/api/v1/system/role/delete', '????????????', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '',
        '2022-04-06 14:50:22', '2022-04-06 14:50:22'),
       (15, 1, 'api/v1/system/dept/list', '????????????', 'iconfont icon-siweidaotu', '', '', 1, 0, 0,
        '/system/auth/deptList', 'system/dept/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-06 14:52:23',
        '2022-04-07 22:59:20'),
       (16, 17, 'aliyun', '?????????-iframe', 'iconfont icon-diannao1', '', '', 1, 0, 0, '/demo/outLink/aliyun',
        'layout/routerView/iframes', 1, '', 0, 1, 1, '', 0,
        'https://www.aliyun.com/daily-act/ecs/activity_selection?spm=5176.8789780.J_3965641470.5.568845b58KHj51',
        '2022-04-06 17:26:29', '2022-04-07 15:27:17'),
       (17, 0, 'outLink', '????????????', 'iconfont icon-zhongduancanshu', '', '', 0, 20, 0, '/demo/outLink',
        'layout/routerView/parent', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 15:20:51', '2022-04-14 16:29:07'),
       (18, 17, 'tenyun', '?????????-??????', 'iconfont icon-shouye_dongtaihui', '', '', 1, 0, 0, '/demo/outLink/tenyun',
        'layout/routerView/link', 1, '', 0, 0, 1, '', 0,
        'https://cloud.tencent.com/act/new?cps_key=20b1c3842f74986b2894e2c5fcde7ea2&fromSource=gwzcw.3775555.3775555.3775555&utm_id=gwzcw.3775555.3775555.3775555&utm_medium=cpc',
        '2022-04-07 15:23:52', '2022-04-07 15:27:25'),
       (19, 15, 'api/v1/system/dept/add', '????????????', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '',
        '2022-04-07 22:56:39', '2022-04-07 22:56:39'),
       (20, 15, 'api/v1/system/dept/edit', '????????????', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '',
        '2022-04-07 22:57:00', '2022-04-07 22:57:00'),
       (21, 15, 'api/v1/system/dept/delete', '????????????', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '',
        '2022-04-07 22:57:30', '2022-04-07 22:57:30'),
       (22, 1, 'api/v1/system/post/list', '????????????', 'iconfont icon-neiqianshujuchucun', '', '', 1, 0, 0,
        '/system/auth/postList', 'system/post/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-07 22:58:46',
        '2022-04-09 14:26:15'),
       (23, 22, 'api/v1/system/post/add', '????????????', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '',
        '2022-04-09 14:14:49', '2022-04-09 14:14:49'),
       (24, 22, 'api/v1/system/post/edit', '????????????', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '',
        '2022-04-09 14:15:25', '2022-04-09 14:15:25'),
       (25, 22, 'api/v1/system/post/delete', '????????????', '', '', '', 2, 0, 0, '', '', 0, '', 0, 0, 1, '', 0, '',
        '2022-04-09 14:15:47', '2022-04-09 14:15:47'),
       (26, 1, 'api/v1/system/user/list', '????????????', 'ele-User', '', '', 1, 0, 0, '/system/auth/user/list',
        'system/user/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-09 14:19:10', '2022-04-09 14:19:58'),
       (27, 0, 'api/v1/system/dict', '????????????', 'iconfont icon-shuxingtu', '', '', 0, 40, 0, '/system/dict',
        'layout/routerView/parent', 0, '', 0, 0, 1, '654', 0, '', '2022-04-14 16:28:51', '2022-04-18 14:40:56'),
       (28, 27, 'api/v1/system/dict/type/list', '????????????', 'iconfont icon-crew_feature', '', '', 1, 0, 0,
        '/system/dict/type/list', 'system/dict/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-14 16:32:10',
        '2022-04-16 17:02:50'),
       (29, 27, 'api/v1/system/dict/dataList', '??????????????????', 'iconfont icon-putong', '', '', 1, 0, 1,
        '/system/dict/data/list/:dictType', 'system/dict/dataList', 0, '', 0, 0, 1, '', 0, '', '2022-04-18 12:04:17',
        '2022-04-18 14:58:43'),
       (30, 27, 'api/v1/system/config/list', '????????????', 'ele-Cherry', '', '', 1, 0, 0, '/system/config/list',
        'system/config/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-18 21:05:20', '2022-04-18 21:13:19'),
       (31, 0, 'api/v1/system/monitor', '????????????', 'iconfont icon-xuanzeqi', '', '', 0, 30, 0, '/system/monitor',
        'layout/routerView/parent', 0, '', 0, 0, 1, '', 0, '', '2022-04-19 10:40:19', '2022-04-19 10:44:38'),
       (32, 31, 'api/v1/system/monitor/server', '????????????', 'iconfont icon-shuju', '', '', 1, 0, 0,
        '/system/monitor/server', 'system/monitor/server/index', 0, '', 0, 0, 1, '', 0, '', '2022-04-19 10:43:32',
        '2022-04-19 10:44:47'),
       (33, 31, 'api/swagger', 'api??????', 'iconfont icon--chaifenlie', '', '', 1, 0, 0, '/system/swagger',
        'layout/routerView/iframes', 1, '', 0, 1, 1, '', 0, 'http://localhost:8201/swagger', '2022-04-21 09:23:43',
        '2022-04-21 11:19:49');
/*!40000 ALTER TABLE `sys_auth_rule`
    ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_config`
--

DROP TABLE IF EXISTS `sys_config`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
SET character_set_client = utf8mb4;
CREATE TABLE `sys_config`
(
    `config_id`    int unsigned NOT NULL AUTO_INCREMENT COMMENT '????????????',
    `config_name`  varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '????????????',
    `config_key`   varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '????????????',
    `config_value` varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '????????????',
    `config_type`  tinyint(1)                                                    DEFAULT '0' COMMENT '???????????????Y??? N??????',
    `create_by`    int unsigned                                                  DEFAULT '0' COMMENT '?????????',
    `update_by`    int unsigned                                                  DEFAULT '0' COMMENT '?????????',
    `remark`       varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '??????',
    `created_at`   datetime                                                      DEFAULT NULL COMMENT '????????????',
    `updated_at`   datetime                                                      DEFAULT NULL COMMENT '????????????',
    PRIMARY KEY (`config_id`) USING BTREE,
    UNIQUE KEY `uni_config_key` (`config_key`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 15
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_config`
--

LOCK TABLES `sys_config` WRITE;
/*!40000 ALTER TABLE `sys_config`
    DISABLE KEYS */;
INSERT INTO `sys_config`
VALUES (1, '????????????-????????????', 'sys.uploadFile.fileSize', '50M', 1, 31, 31, '????????????????????????', NULL,
        '2021-07-06 14:57:35'),
       (2, '????????????-????????????', 'sys.uploadFile.fileType', 'doc,docx,zip,xls,xlsx,rar,jpg,jpeg,gif,npm,png', 1, 31, 31,
        '??????????????????????????????', NULL, NULL),
       (3, '????????????-????????????', 'sys.uploadFile.imageType', 'jpg,jpeg,gif,npm,png', 1, 31, 0, '??????????????????????????????',
        NULL, NULL),
       (4, '????????????-????????????', 'sys.uploadFile.imageSize', '50M', 1, 31, 31, '????????????????????????', NULL, NULL),
       (11, '????????????', 'static.resource', '/', 1, 2, 0, '', NULL, NULL);
/*!40000 ALTER TABLE `sys_config`
    ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dept`
--

DROP TABLE IF EXISTS `sys_dept`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
SET character_set_client = utf8mb4;
CREATE TABLE `sys_dept`
(
    `dept_id`    bigint NOT NULL AUTO_INCREMENT COMMENT '??????id',
    `parent_id`  bigint                                                       DEFAULT '0' COMMENT '?????????id',
    `ancestors`  varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '????????????',
    `dept_name`  varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '????????????',
    `order_num`  int                                                          DEFAULT '0' COMMENT '????????????',
    `leader`     varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '?????????',
    `phone`      varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '????????????',
    `email`      varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '??????',
    `status`     tinyint unsigned                                             DEFAULT '0' COMMENT '???????????????0?????? 1?????????',
    `created_by` bigint unsigned                                              DEFAULT '0' COMMENT '?????????',
    `updated_by` bigint                                                       DEFAULT NULL COMMENT '?????????',
    `created_at` datetime                                                     DEFAULT NULL COMMENT '????????????',
    `updated_at` datetime                                                     DEFAULT NULL COMMENT '????????????',
    `deleted_at` datetime                                                     DEFAULT NULL COMMENT '????????????',
    PRIMARY KEY (`dept_id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 204
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = COMPACT COMMENT ='?????????';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dept`
--

LOCK TABLES `sys_dept` WRITE;
/*!40000 ALTER TABLE `sys_dept`
    DISABLE KEYS */;
INSERT INTO `sys_dept`
VALUES (100, 0, '0', '????????????', 0, '??????', '15888888888', 'ry@qq.com', 1, 0, 31, '2021-07-13 15:56:52',
        '2021-07-13 15:57:05', NULL),
       (101, 100, '0,100', '???????????????', 1, '??????', '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52',
        '2021-07-13 15:56:52', NULL),
       (102, 100, '0,100', '???????????????', 2, '??????', '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52',
        '2021-07-13 15:56:52', NULL),
       (103, 101, '0,100,101', '????????????', 1, '??????', '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52',
        '2021-07-13 15:56:52', NULL),
       (104, 101, '0,100,101', '????????????', 2, '??????', '15888888888', 'ry@qq.com', 1, 0, 31, '2021-07-13 15:56:52',
        '2021-11-04 09:16:38', NULL),
       (105, 101, '0,100,101', '????????????', 3, '??????', '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52',
        '2021-07-13 15:56:52', NULL),
       (106, 101, '0,100,101', '????????????', 4, '??????', '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52',
        '2021-07-13 15:56:52', NULL),
       (107, 101, '0,100,101', '????????????', 5, '??????', '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52',
        '2021-07-13 15:56:52', NULL),
       (108, 102, '0,100,102', '????????????', 1, '??????', '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52',
        '2021-07-13 15:56:52', NULL),
       (109, 102, '0,100,102', '????????????', 2, '??????', '15888888888', 'ry@qq.com', 1, 0, NULL, '2021-07-13 15:56:52',
        '2021-07-13 15:56:52', NULL),
       (200, 100, '', '?????????', 1, '??????', '18888888888', 'liou@qq.com', 1, 0, NULL, '2021-07-13 15:56:52',
        '2021-07-13 15:56:52', NULL),
       (201, 100, '', '??????', 1, '??????', '18888888888', 'li@qq.com', 0, 31, NULL, '2021-07-13 15:56:52',
        '2022-04-07 22:35:21', NULL),
       (202, 108, '', '??????', 1, '???a', '18888888888', 'aa@qq.com', 1, 0, NULL, '2021-07-13 15:56:52',
        '2021-07-13 15:56:52', NULL),
       (203, 108, '', '??????', 0, 'aa', '18888888888', 'aa@qq.com', 0, 0, NULL, '2021-07-13 15:56:52',
        '2021-07-13 15:56:52', NULL);
/*!40000 ALTER TABLE `sys_dept`
    ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dict_data`
--

DROP TABLE IF EXISTS `sys_dict_data`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
SET character_set_client = utf8mb4;
CREATE TABLE `sys_dict_data`
(
    `dict_code`  bigint NOT NULL AUTO_INCREMENT COMMENT '????????????',
    `dict_sort`  int                                                           DEFAULT '0' COMMENT '????????????',
    `dict_label` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '????????????',
    `dict_value` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '????????????',
    `dict_type`  varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '????????????',
    `css_class`  varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '????????????????????????????????????',
    `list_class` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '??????????????????',
    `is_default` tinyint(1)                                                    DEFAULT '0' COMMENT '???????????????1??? 0??????',
    `status`     tinyint(1)                                                    DEFAULT '0' COMMENT '?????????0?????? 1?????????',
    `create_by`  bigint unsigned                                               DEFAULT '0' COMMENT '?????????',
    `update_by`  bigint unsigned                                               DEFAULT '0' COMMENT '?????????',
    `remark`     varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '??????',
    `created_at` datetime                                                      DEFAULT NULL COMMENT '????????????',
    `updated_at` datetime                                                      DEFAULT NULL COMMENT '????????????',
    PRIMARY KEY (`dict_code`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 102
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = COMPACT COMMENT ='???????????????';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dict_data`
--

LOCK TABLES `sys_dict_data` WRITE;
/*!40000 ALTER TABLE `sys_dict_data`
    DISABLE KEYS */;
INSERT INTO `sys_dict_data`
VALUES (1, 0, '???', '1', 'sys_user_sex', '', '', 0, 1, 31, 2, '????????????', '2022-04-18 16:46:22', NULL),
       (2, 0, '???', '2', 'sys_user_sex', '', '', 0, 1, 31, 31, '????????????', NULL, NULL),
       (3, 0, '??????', '0', 'sys_user_sex', '', '', 1, 1, 31, 31, '????????????', NULL, NULL),
       (24, 0, '?????????', '1', 'cms_category_type', '', '', 0, 1, 31, 31,
        '????????????????????????????????????????????????????????????????????????', NULL, '2021-07-21 10:54:22'),
       (25, 0, '????????????', '2', 'cms_category_type', '', '', 0, 1, 31, 31, '????????????????????????????????????', NULL,
        '2021-07-21 10:54:22'),
       (26, 0, '????????????', '3', 'cms_category_type', '', '', 0, 1, 31, 31, '??????????????????????????????????????????', NULL,
        '2021-07-21 10:54:22'),
       (27, 0, '????????????', '4', 'cms_category_type', '', '', 0, 1, 31, 31, '?????????????????????????????????????????????', NULL,
        '2021-07-21 10:54:22'),
       (28, 0, '??????', '0', 'sys_job_status', '', 'default', 1, 1, 31, 0, '', NULL, NULL),
       (29, 0, '??????', '1', 'sys_job_status', '', 'default', 0, 1, 31, 31, '', NULL, NULL),
       (30, 0, '??????', 'DEFAULT', 'sys_job_group', '', 'default', 1, 1, 31, 0, '', NULL, NULL),
       (31, 0, '??????', 'SYSTEM', 'sys_job_group', '', 'default', 0, 1, 31, 0, '', NULL, NULL),
       (32, 0, '??????', '1', 'admin_login_status', '', 'default', 0, 1, 31, 31, '', NULL, NULL),
       (33, 0, '??????', '0', 'admin_login_status', '', 'default', 0, 1, 31, 0, '', NULL, NULL),
       (34, 0, '??????', '1', 'sys_oper_log_status', '', 'default', 0, 1, 31, 0, '', NULL, NULL),
       (35, 0, '??????', '0', 'sys_oper_log_status', '', 'default', 0, 1, 31, 0, '', NULL, NULL),
       (36, 0, '????????????', '1', 'sys_job_policy', '', 'default', 1, 1, 31, 0, '', NULL, NULL),
       (37, 0, '????????????', '2', 'sys_job_policy', '', 'default', 1, 1, 31, 0, '', NULL, NULL),
       (38, 0, '??????', '0', 'sys_show_hide', NULL, 'default', 1, 1, 31, 0, NULL, NULL, NULL),
       (39, 0, '??????', '1', 'sys_show_hide', NULL, 'default', 0, 1, 31, 0, NULL, NULL, NULL),
       (40, 0, '??????', '1', 'sys_normal_disable', '', 'default', 1, 1, 31, 0, '', NULL, NULL),
       (41, 0, '??????', '0', 'sys_normal_disable', '', 'default', 0, 1, 31, 0, '', NULL, NULL),
       (49, 0, '???', '1', 'sys_yes_no', '', '', 1, 1, 31, 0, '', NULL, NULL),
       (50, 0, '???', '0', 'sys_yes_no', '', '', 0, 1, 31, 0, '', NULL, NULL),
       (51, 0, '?????????', '1', 'cms_article_pub_type', '', '', 1, 1, 31, 31, '', NULL, NULL),
       (54, 0, '?????????', '0', 'cms_article_pub_type', '', '', 0, 1, 31, 0, '', NULL, NULL),
       (55, 0, '??????', '1', 'cms_article_attr', '', '', 0, 1, 31, 0, '', NULL, NULL),
       (56, 0, '??????', '2', 'cms_article_attr', '', '', 0, 1, 31, 0, '', NULL, NULL),
       (57, 0, '????????????', '0', 'cms_article_type', '', '', 0, 1, 31, 31, '', NULL, NULL),
       (58, 0, '????????????', '1', 'cms_article_type', '', '', 0, 1, 31, 31, '', NULL, NULL),
       (59, 0, 'cms??????', '6', 'cms_cate_models', '', '', 0, 1, 1, 1, '', NULL, NULL),
       (61, 0, '??????????????????', '1', 'gov_cate_models', '', '', 0, 1, 2, 0, '', NULL, NULL),
       (62, 0, '????????????', 'sys_admin', 'menu_module_type', '', '', 1, 1, 2, 0, '', NULL, NULL),
       (63, 0, '????????????', 'gov_work', 'menu_module_type', '', '', 0, 1, 2, 0, '', NULL, NULL),
       (64, 0, '??????', '3', 'cms_article_attr', '', '', 0, 1, 31, 0, '', NULL, NULL),
       (65, 0, '[work]???????????????', 'wf_news', 'flow_type', '', '', 0, 1, 2, 2, '', NULL, NULL),
       (66, 0, '????????????', '-1', 'flow_status', '', '', 0, 1, 31, 0, '', NULL, NULL),
       (67, 0, '?????????', '0', 'flow_status', '', '', 0, 1, 31, 0, '', NULL, NULL),
       (68, 0, '?????????', '1', 'flow_status', '', '', 0, 1, 31, 0, '', NULL, NULL),
       (69, 0, '????????????', '2', 'flow_status', '', '', 0, 1, 31, 2, '', NULL, NULL),
       (70, 2, '????????????', '2', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL),
       (71, 3, '????????????', '3', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL),
       (72, 4, '????????????', '4', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL),
       (73, 2, '??????', '1', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL),
       (74, 3, '??????', '2', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL),
       (75, 4, '??????', '3', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL),
       (76, 1, '??????', '0', 'sys_log_sign', '', '', 0, 1, 31, 31, '', NULL, NULL),
       (77, 1, '?????????', '1', 'sys_blog_sign', '', '', 0, 1, 31, 31, '', NULL, NULL),
       (78, 0, '??????', '0', 'flow_level', '', '', 0, 1, 31, 0, '', NULL, '2021-07-20 08:55:20'),
       (79, 0, '??????', '1', 'flow_level', '', '', 0, 1, 31, 0, '', NULL, '2021-07-20 08:55:20'),
       (80, 0, '??????', '2', 'flow_level', '', '', 0, 1, 31, 0, '', NULL, '2021-07-20 08:55:20'),
       (81, 0, '??????', '3', 'flow_level', '', '', 0, 1, 31, 31, '', NULL, '2021-07-20 08:55:25'),
       (82, 0, '?????????', '1', 'sys_blog_type', '', '', 0, 1, 31, 0, '', NULL, NULL),
       (83, 0, '????????????', '2', 'sys_blog_type', '', '', 0, 1, 31, 0, '', NULL, NULL),
       (84, 0, '????????????', '3', 'sys_blog_type', '', '', 0, 1, 31, 31, '', NULL, NULL),
       (85, 0, '????????????', '4', 'sys_blog_type', '', '', 0, 1, 31, 0, '', NULL, NULL),
       (87, 0, '[cms]?????????', 'cms_news', 'flow_type', '', '', 0, 1, 31, 0, '', NULL, NULL),
       (91, 0, '????????????', '666', 'cms_article_type', '', '', 0, 1, 31, 0, '', '2021-08-03 17:04:12',
        '2021-08-03 17:04:12'),
       (92, 0, '????????????222', '33333', 'cms_article_type', '', '', 0, 1, 31, 31, '', '2021-08-03 17:16:45',
        '2021-08-03 17:19:41'),
       (93, 0, '????????????222', '11111', 'cms_article_type', '', '', 0, 1, 31, 31, '', '2021-08-03 17:26:14',
        '2021-08-03 17:26:26'),
       (94, 0, '1???', '10', 'plugin_store_discount', '', '', 0, 1, 31, 0, '', '2021-08-14 11:59:38',
        '2021-08-14 11:59:38'),
       (95, 0, '5???', '50', 'plugin_store_discount', '', '', 0, 1, 31, 0, '', '2021-08-14 11:59:49',
        '2021-08-14 11:59:49'),
       (96, 0, '8???', '80', 'plugin_store_discount', '', '', 0, 1, 31, 0, '', '2021-08-14 12:00:00',
        '2021-08-14 12:00:00'),
       (97, 0, '9???', '90', 'plugin_store_discount', '', '', 0, 1, 31, 0, '', '2021-08-14 12:00:07',
        '2021-08-14 12:00:07'),
       (98, 0, '?????????', '100', 'plugin_store_discount', '', '', 0, 1, 31, 0, '', '2021-08-14 12:00:16',
        '2021-08-14 12:00:16'),
       (99, 0, '?????????', 'none', 'cms_nav_position', '', '', 1, 1, 22, 0, '', '2021-08-31 15:37:35',
        '2021-08-31 15:37:35'),
       (100, 0, '????????????', 'top', 'cms_nav_position', '', '', 0, 1, 22, 0, '', '2021-08-31 15:37:57',
        '2021-08-31 15:37:57'),
       (101, 0, '????????????', 'bottom', 'cms_nav_position', '', '', 0, 1, 22, 0, '', '2021-08-31 15:38:08',
        '2021-08-31 15:38:08');
/*!40000 ALTER TABLE `sys_dict_data`
    ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_dict_type`
--

DROP TABLE IF EXISTS `sys_dict_type`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
SET character_set_client = utf8mb4;
CREATE TABLE `sys_dict_type`
(
    `dict_id`    bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '????????????',
    `dict_name`  varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '????????????',
    `dict_type`  varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '????????????',
    `status`     tinyint unsigned                                              DEFAULT '0' COMMENT '?????????0?????? 1?????????',
    `create_by`  int unsigned                                                  DEFAULT '0' COMMENT '?????????',
    `update_by`  int unsigned                                                  DEFAULT '0' COMMENT '?????????',
    `remark`     varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '??????',
    `created_at` datetime                                                      DEFAULT NULL COMMENT '????????????',
    `updated_at` datetime                                                      DEFAULT NULL COMMENT '????????????',
    PRIMARY KEY (`dict_id`) USING BTREE,
    UNIQUE KEY `dict_type` (`dict_type`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 50
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = COMPACT COMMENT ='???????????????';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_dict_type`
--

LOCK TABLES `sys_dict_type` WRITE;
/*!40000 ALTER TABLE `sys_dict_type`
    DISABLE KEYS */;
INSERT INTO `sys_dict_type`
VALUES (1, '????????????', 'sys_user_sex', 1, 31, 1, '????????????????????????', NULL, NULL),
       (2, '????????????', 'cms_category_type', 1, 31, 3, '??????????????????', NULL, '2021-07-21 10:54:22'),
       (3, '????????????', 'sys_job_status', 1, 31, 31, '??????????????????', NULL, NULL),
       (13, '????????????', 'sys_job_group', 1, 31, 0, '', NULL, NULL),
       (14, '?????????????????????', 'admin_login_status', 1, 31, 0, '', NULL, NULL),
       (15, '??????????????????', 'sys_oper_log_status', 1, 31, 0, '', NULL, NULL),
       (16, '????????????', 'sys_job_policy', 1, 31, 0, '', NULL, NULL),
       (17, '????????????', 'sys_show_hide', 1, 31, 0, '????????????', NULL, NULL),
       (18, '????????????', 'sys_normal_disable', 1, 31, 31, '????????????', NULL, NULL),
       (24, '????????????', 'sys_yes_no', 1, 31, 0, '', NULL, NULL),
       (25, '??????????????????', 'cms_article_pub_type', 1, 31, 0, '', NULL, NULL),
       (26, '??????????????????', 'cms_article_attr', 1, 31, 0, '', NULL, NULL),
       (27, '????????????', 'cms_article_type', 1, 31, 0, '', NULL, NULL),
       (28, '????????????????????????', 'cms_cate_models', 1, 1, 0, '', NULL, NULL),
       (29, '????????????????????????', 'gov_cate_models', 1, 2, 0, '', NULL, NULL),
       (30, '??????????????????', 'menu_module_type', 1, 2, 0, '', NULL, NULL),
       (31, '??????????????????', 'flow_type', 1, 2, 0, '', NULL, NULL),
       (32, '????????????????????????', 'flow_status', 1, 31, 0, '????????????????????????', NULL, NULL),
       (33, '??????????????????', 'sys_blog_type', 1, 31, 31, '????????????????????????', NULL, NULL),
       (34, '??????????????????', 'sys_log_sign', 1, 31, 0, '??????????????????????????????????????????', NULL, NULL),
       (35, '?????????????????????', 'flow_level', 1, 31, 31, '', NULL, '2021-07-20 08:55:20'),
       (48, '??????????????????', 'plugin_store_discount', 1, 31, 0, '', '2021-08-14 11:59:26', '2021-08-14 11:59:26'),
       (49, 'CMS??????????????????', 'cms_nav_position', 1, 22, 0, '', '2021-08-31 15:37:04', '2021-08-31 15:37:04');
/*!40000 ALTER TABLE `sys_dict_type`
    ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_login_log`
--

DROP TABLE IF EXISTS `sys_login_log`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
SET character_set_client = utf8mb4;
CREATE TABLE `sys_login_log`
(
    `info_id`        bigint NOT NULL AUTO_INCREMENT COMMENT '??????ID',
    `login_name`     varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  DEFAULT '' COMMENT '????????????',
    `ipaddr`         varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  DEFAULT '' COMMENT '??????IP??????',
    `login_location` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '????????????',
    `browser`        varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  DEFAULT '' COMMENT '???????????????',
    `os`             varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  DEFAULT '' COMMENT '????????????',
    `status`         tinyint                                                       DEFAULT '0' COMMENT '???????????????0?????? 1?????????',
    `msg`            varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '????????????',
    `login_time`     datetime                                                      DEFAULT NULL COMMENT '????????????',
    `module`         varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  DEFAULT '' COMMENT '????????????',
    PRIMARY KEY (`info_id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 886
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = DYNAMIC COMMENT ='??????????????????';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_login_log`
--

LOCK TABLES `sys_login_log` WRITE;
/*!40000 ALTER TABLE `sys_login_log`
    DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_login_log`
    ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_oper_log`
--

DROP TABLE IF EXISTS `sys_oper_log`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
SET character_set_client = utf8mb4;
CREATE TABLE `sys_oper_log`
(
    `oper_id`        bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '????????????',
    `title`          varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci   DEFAULT '' COMMENT '????????????',
    `business_type`  int                                                            DEFAULT '0' COMMENT '???????????????0?????? 1?????? 2?????? 3?????????',
    `method`         varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  DEFAULT '' COMMENT '????????????',
    `request_method` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci   DEFAULT '' COMMENT '????????????',
    `operator_type`  int                                                            DEFAULT '0' COMMENT '???????????????0?????? 1???????????? 2??????????????????',
    `oper_name`      varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci   DEFAULT '' COMMENT '????????????',
    `dept_name`      varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci   DEFAULT '' COMMENT '????????????',
    `oper_url`       varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  DEFAULT '' COMMENT '??????URL',
    `oper_ip`        varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci   DEFAULT '' COMMENT '????????????',
    `oper_location`  varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  DEFAULT '' COMMENT '????????????',
    `oper_param`     text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '????????????',
    `json_result`    text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci COMMENT '????????????',
    `status`         int                                                            DEFAULT '0' COMMENT '???????????????0?????? 1?????????',
    `error_msg`      varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT '' COMMENT '????????????',
    `oper_time`      datetime                                                       DEFAULT NULL COMMENT '????????????',
    PRIMARY KEY (`oper_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = DYNAMIC COMMENT ='??????????????????';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_oper_log`
--

LOCK TABLES `sys_oper_log` WRITE;
/*!40000 ALTER TABLE `sys_oper_log`
    DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_oper_log`
    ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_post`
--

DROP TABLE IF EXISTS `sys_post`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
SET character_set_client = utf8mb4;
CREATE TABLE `sys_post`
(
    `post_id`    bigint unsigned                                              NOT NULL AUTO_INCREMENT COMMENT '??????ID',
    `post_code`  varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '????????????',
    `post_name`  varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '????????????',
    `post_sort`  int                                                          NOT NULL COMMENT '????????????',
    `status`     tinyint unsigned                                             NOT NULL DEFAULT '0' COMMENT '?????????0?????? 1?????????',
    `remark`     varchar(500) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci         DEFAULT NULL COMMENT '??????',
    `created_by` bigint unsigned                                              NOT NULL DEFAULT '0' COMMENT '?????????',
    `updated_by` bigint unsigned                                              NOT NULL DEFAULT '0' COMMENT '?????????',
    `created_at` datetime                                                              DEFAULT NULL COMMENT '????????????',
    `updated_at` datetime                                                              DEFAULT NULL COMMENT '????????????',
    `deleted_at` datetime                                                              DEFAULT NULL COMMENT '????????????',
    PRIMARY KEY (`post_id`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 10
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = COMPACT COMMENT ='???????????????';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_post`
--

LOCK TABLES `sys_post` WRITE;
/*!40000 ALTER TABLE `sys_post`
    DISABLE KEYS */;
INSERT INTO `sys_post`
VALUES (1, 'ceo', '?????????', 1, 1, '?????????', 0, 0, '2021-07-11 11:32:58', NULL, NULL),
       (2, 'se', '????????????', 2, 1, '????????????', 0, 0, '2021-07-12 11:01:26', NULL, NULL),
       (3, 'hr', '????????????', 3, 1, '????????????', 0, 0, '2021-07-12 11:01:30', NULL, NULL),
       (4, 'user', '????????????', 4, 1, '????????????', 0, 1, '2021-07-12 11:01:33', '2022-10-24 21:03:54', NULL),
       (5, 'it', 'IT???', 5, 1, '?????????', 31, 31, '2021-07-12 11:09:42', '2022-04-09 12:59:12', NULL);
/*!40000 ALTER TABLE `sys_post`
    ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role`
--

DROP TABLE IF EXISTS `sys_role`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
SET character_set_client = utf8mb4;
CREATE TABLE `sys_role`
(
    `id`         int unsigned                                                  NOT NULL AUTO_INCREMENT,
    `status`     tinyint unsigned                                              NOT NULL DEFAULT '0' COMMENT '??????;0:??????;1:??????',
    `list_order` int unsigned                                                  NOT NULL DEFAULT '0' COMMENT '??????',
    `name`       varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '????????????',
    `remark`     varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '??????',
    `data_scope` tinyint unsigned                                              NOT NULL DEFAULT '3' COMMENT '???????????????1????????????????????? 2????????????????????? 3???????????????????????? 4????????????????????????????????????',
    `created_at` datetime                                                               DEFAULT NULL COMMENT '????????????',
    `updated_at` datetime                                                               DEFAULT NULL COMMENT '????????????',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `status` (`status`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 9
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = COMPACT COMMENT ='?????????';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role`
--

LOCK TABLES `sys_role` WRITE;
/*!40000 ALTER TABLE `sys_role`
    DISABLE KEYS */;
INSERT INTO `sys_role`
VALUES (1, 1, 0, '???????????????', '???????????????', 3, '2022-04-01 11:38:39', '2022-04-09 12:59:28');
/*!40000 ALTER TABLE `sys_role`
    ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_role_dept`
--

DROP TABLE IF EXISTS `sys_role_dept`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
SET character_set_client = utf8mb4;
CREATE TABLE `sys_role_dept`
(
    `role_id` bigint NOT NULL COMMENT '??????ID',
    `dept_id` bigint NOT NULL COMMENT '??????ID',
    PRIMARY KEY (`role_id`, `dept_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = COMPACT COMMENT ='????????????????????????';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_role_dept`
--

LOCK TABLES `sys_role_dept` WRITE;
/*!40000 ALTER TABLE `sys_role_dept`
    DISABLE KEYS */;
INSERT INTO `sys_role_dept`
VALUES (5, 103),
       (5, 104),
       (5, 105),
       (8, 105),
       (8, 106);
/*!40000 ALTER TABLE `sys_role_dept`
    ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user`
--

DROP TABLE IF EXISTS `sys_user`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
SET character_set_client = utf8mb4;
CREATE TABLE `sys_user`
(
    `id`              bigint unsigned                                               NOT NULL AUTO_INCREMENT,
    `user_name`       varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '?????????',
    `mobile`          varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '????????????????????????????????????????????????????????????????????????-?????????',
    `user_nickname`   varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci  NOT NULL DEFAULT '' COMMENT '????????????',
    `birthday`        int                                                           NOT NULL DEFAULT '0' COMMENT '??????',
    `user_password`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '????????????;cmf_password??????',
    `user_salt`       char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci     NOT NULL COMMENT '?????????',
    `user_status`     tinyint unsigned                                              NOT NULL DEFAULT '1' COMMENT '????????????;0:??????,1:??????,2:?????????',
    `user_email`      varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '??????????????????',
    `sex`             tinyint                                                       NOT NULL DEFAULT '0' COMMENT '??????;0:??????,1:???,2:???',
    `avatar`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '????????????',
    `dept_id`         bigint unsigned                                               NOT NULL DEFAULT '0' COMMENT '??????id',
    `remark`          varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '??????',
    `is_admin`        tinyint                                                       NOT NULL DEFAULT '1' COMMENT '????????????????????? 1 ???  0   ???',
    `address`         varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '????????????',
    `describe`        varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT ' ????????????',
    `last_login_ip`   varchar(15) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '????????????ip',
    `last_login_time` datetime                                                               DEFAULT NULL COMMENT '??????????????????',
    `created_at`      datetime                                                               DEFAULT NULL COMMENT '????????????',
    `updated_at`      datetime                                                               DEFAULT NULL COMMENT '????????????',
    `deleted_at`      datetime                                                               DEFAULT NULL COMMENT '????????????',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `user_login` (`user_name`, `deleted_at`) USING BTREE,
    UNIQUE KEY `mobile` (`mobile`, `deleted_at`) USING BTREE,
    KEY `user_nickname` (`user_nickname`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 43
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = COMPACT COMMENT ='?????????';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user`
--

LOCK TABLES `sys_user` WRITE;
/*!40000 ALTER TABLE `sys_user`
    DISABLE KEYS */;
INSERT INTO `sys_user`
VALUES (1, 'admin', '13578342363', '???????????????', 0, '39978de67915a11e94bfe9c879b2d9a1', 'gqwLs4n95E', 1,
        'yxh669@qq.com', 1,
        'https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-07-19/ccwpeuqz1i2s769hua.jpeg', 101, '', 1,
        'asdasfdsaf?????????????????????????????????', '????????????', '::1', '2022-04-19 16:38:37', '2021-06-22 17:58:00',
        '2022-04-19 16:38:37', NULL),
       (2, 'yixiaohu', '13699885599', '??????', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, 'yxh@qq.com', 1,
        'pub_upload/2020-11-02/c6sntzg7r96c7p9gqf.jpeg', 102, '??????', 1, '', '', '[::1]', '2022-02-14 18:10:40',
        '2021-06-22 17:58:00', '2022-04-13 11:20:03', '2022-10-24 20:56:30'),
       (3, 'zs', '16399669855', '??????', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, 'zs@qq.com', 0,
        'https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-08-02/cd8nif79egjg9kbkgk.jpeg', 101, '', 1, '',
        '', '127.0.0.1', '2022-03-18 15:22:13', '2021-06-22 17:58:00', '2022-04-21 11:20:06', '2022-10-24 20:56:30'),
       (4, 'qlgl', '13758596696', '??????c', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, 'qlgl@qq.com', 0, '',
        102, '', 1, '', '', '127.0.0.1', NULL, '2021-06-22 17:58:00', '2022-04-12 17:55:28', '2022-10-24 20:56:30'),
       (5, 'test', '13845696696', '??????2', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '123@qq.com', 0, '',
        101, '', 0, '', '', '::1', '2022-03-30 10:50:39', '2021-06-22 17:58:00', '2022-04-12 17:55:31',
        '2022-10-24 20:56:30'),
       (6, '18999998889', '13755866654', '?????????', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1,
        '1223@qq.com', 0, '', 103, '', 1, '', '', '[::1]', '2022-02-25 14:29:22', '2021-06-22 17:58:00',
        '2022-04-14 21:11:06', '2022-10-24 20:56:30'),
       (7, 'zmm', '13788566696', '?????????', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '11123@qq.com', 0,
        '', 104, '', 1, '', '', '127.0.0.1', NULL, '2021-06-22 17:58:00', '2022-04-12 17:55:42', '2022-10-24 20:56:30'),
       (8, 'lxx', '13756566696', '?????????', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '123333@qq.com', 0,
        '', 101, '', 1, '', '', '127.0.0.1', NULL, '2021-06-22 17:58:00', '2022-04-12 17:55:45', '2022-10-24 20:56:30'),
       (10, 'xmm', '13588999969', '?????????', 0, '2de2a8df703bfc634cfda2cb2f6a59be', 'Frz7LJY7SE', 1, '696@qq.com', 0, '',
        101, '', 1, '', '', '[::1]', '2021-07-22 17:08:53', '2021-06-22 17:58:00', '2022-04-12 17:55:50',
        '2022-10-24 20:56:30'),
       (14, 'cd_19', '13699888899', '???????????????', 0, '1169d5fe4119fd4277a95f02d7036171', '7paigEoedh', 1, '', 0, '',
        102, '', 1, '', '', '', NULL, '2021-06-22 17:58:00', '2022-04-12 18:13:22', '2022-10-24 20:56:30'),
       (15, 'lmm', '13587754545', '?????????', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, 'a@coc.com', 0, '',
        201, '', 1, '', '', '127.0.0.1', NULL, '2021-06-22 17:58:00', '2022-04-12 17:56:23', '2022-10-24 20:56:30'),
       (16, 'ldn', '13899658874', '?????????', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, 'a@ll.con', 0, '',
        102, '', 1, '', '', '127.0.0.1', NULL, '2021-06-22 17:58:00', '2022-04-12 17:56:27', '2022-10-24 20:56:30'),
       (20, 'dbc', '13877555566', '?????????', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '', 0, '', 0, '', 1,
        '', '', '', NULL, '2021-06-22 17:58:00', '2021-06-22 17:58:00', '2022-10-24 20:56:30'),
       (22, 'yxfmlbb', '15969423326', '?????????????????????', 0, '66f89b40ee4a10aabaf70c15756429ea', 'mvd2OtUe8f', 1,
        'yxh6691@qq.com', 0,
        'https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-09-29/cem20k3fdciosy7nwo.jpeg', 200, '', 1,
        '2222233', '1222', '[::1]', '2021-10-28 11:36:07', '2021-06-22 17:58:00', '2021-06-22 17:58:00',
        '2022-10-24 20:56:30'),
       (23, 'wangming', '13699888855', '??????', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '', 0, '', 0, '',
        1, '', '', '', NULL, '2021-06-22 17:58:00', '2021-06-22 17:58:00', '2022-10-24 20:56:30'),
       (24, 'zhk', '13699885591', '?????????', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '', 0, '', 0, '', 1,
        '', '', '192.168.0.146', NULL, '2021-06-22 17:58:00', '2021-06-22 17:58:00', '2022-10-24 20:56:30'),
       (28, 'demo3', '18699888855', '????????????1', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1,
        '123132@qq.com', 0, '', 109, '', 1, '', '', '192.168.0.229', NULL, '2021-06-22 17:58:00', '2021-06-22 17:58:00',
        '2022-10-24 20:56:30'),
       (31, 'demo', '15334455789', '??????', 0, '39978de67915a11e94bfe9c879b2d9a1', 'gqwLs4n95E', 1, '223@qq.com', 2,
        'https://yxh-1301841944.cos.ap-chongqing.myqcloud.com/gfast/2021-11-30/cg30rab8myj85vjzcf.jpeg', 109, '', 1,
        '??????????????????22223', '12345', '127.0.0.1', '2022-04-21 17:28:09', '2021-06-22 17:58:00', '2022-04-21 17:28:09',
        '2022-10-24 20:56:30'),
       (32, 'demo100', '18699888859', '????????????1', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '', 0, '', 0,
        '', 1, '', '', '[::1]', '2021-11-24 18:01:21', '2021-06-22 17:58:00', '2021-06-22 17:58:00',
        '2022-10-24 20:56:30'),
       (33, 'demo110', '18699888853', '????????????1', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '', 0, '', 0,
        '', 1, '', '', '', NULL, '2021-06-22 17:58:00', '2021-06-22 17:58:00', '2022-10-24 20:56:30'),
       (34, 'yxfmlbb2', '15969423327', '??????????????????', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1,
        '1111@qqq.com', 1, '', 103, '', 0, '', '', '127.0.0.1', NULL, '2021-06-22 17:58:00', '2021-06-22 17:58:00',
        '2022-10-24 20:56:30'),
       (35, 'wk666', '18888888888', 'wk', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '396861557@qq.com', 1,
        '', 100, '', 1, '', '', '[::1]', '2021-12-09 14:52:37', '2021-06-22 17:58:00', '2021-06-22 17:58:00',
        '2022-10-24 20:56:30'),
       (36, 'zxd', '13699885565', '?????????', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, 'zxk@qq.com', 1, '',
        201, '666', 1, '', '', '', NULL, '2021-06-22 17:58:00', '2021-06-22 17:58:00', '2022-10-24 20:56:30'),
       (37, 'yxfmlbb3', '13513513511', '??????', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '111@qq.com', 0,
        '', 204, '', 1, '', '', '[::1]', '2021-07-26 14:49:25', '2021-06-22 17:58:00', '2021-07-26 14:49:18',
        '2022-10-24 20:56:30'),
       (38, 'test_user', '18888888880', 'test', 0, '542a6e44dbac171f260fc4a032cd5522', 'dlqVVBTADg', 1, '11@qq.com', 1,
        '', 200, '111', 0, '', '', '', NULL, '2021-06-22 17:58:00', '2021-07-12 22:05:29', '2022-10-24 20:56:30'),
       (39, 'asan', '18687460555', '??????', 0, '2354837137115700e2adf870ac113dcf', 'drdDvbtYZW', 1, '456654@qq.com', 1,
        '', 201, '666666', 1, '', '', '', NULL, '2021-07-12 17:21:43', '2021-07-12 21:13:31', '2021-07-12 22:00:44'),
       (40, 'asi', '13655888888', '??????', 0, 'fbb755b35d48759dad47bb1540249fd1', '9dfUstcxrz', 1, '5464@qq.com', 1, '',
        201, 'adsaasd', 1, '', '', '', '2021-07-12 17:54:31', '2021-07-12 17:46:27', '2021-07-12 21:29:41',
        '2021-07-12 22:00:44'),
       (41, 'awu', '13578556546', '??????', 0, '3b36a96afa0dfd66aa915e0816e0e9f6', '9gHRa9ho4U', 0, '132321@qq.com', 1,
        '', 201, 'asdasdasd', 1, '', '', '', NULL, '2021-07-12 17:54:31', '2021-07-12 21:46:34', '2021-07-12 21:59:56'),
       (42, 'demo01', '13699888556', '??????01222', 0, '048dc94116558fb40920f3553ecd5fe8', 'KiVrfzKJQx', 1, '456@qq.com',
        2, '', 109, '????????????', 1, '', '', '', NULL, '2022-04-12 16:15:23', '2022-04-12 17:54:49',
        '2022-10-24 20:56:30');
/*!40000 ALTER TABLE `sys_user`
    ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_online`
--

DROP TABLE IF EXISTS `sys_user_online`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
SET character_set_client = utf8mb4;
CREATE TABLE `sys_user_online`
(
    `id`          bigint unsigned                                               NOT NULL AUTO_INCREMENT,
    `uuid`        char(32) CHARACTER SET latin1 COLLATE latin1_general_ci       NOT NULL DEFAULT '' COMMENT '????????????',
    `token`       varchar(255) CHARACTER SET latin1 COLLATE latin1_general_ci   NOT NULL DEFAULT '' COMMENT '??????token',
    `create_time` datetime                                                               DEFAULT NULL COMMENT '????????????',
    `user_name`   varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '?????????',
    `ip`          varchar(120) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '??????ip',
    `explorer`    varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '?????????',
    `os`          varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '' COMMENT '????????????',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uni_token` (`token`) USING BTREE
) ENGINE = InnoDB
  AUTO_INCREMENT = 17387
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = DYNAMIC COMMENT ='?????????????????????';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_online`
--

LOCK TABLES `sys_user_online` WRITE;
/*!40000 ALTER TABLE `sys_user_online`
    DISABLE KEYS */;
/*!40000 ALTER TABLE `sys_user_online`
    ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sys_user_post`
--

DROP TABLE IF EXISTS `sys_user_post`;
/*!40101 SET @saved_cs_client = @@character_set_client */;
SET character_set_client = utf8mb4;
CREATE TABLE `sys_user_post`
(
    `user_id` bigint NOT NULL COMMENT '??????ID',
    `post_id` bigint NOT NULL COMMENT '??????ID',
    PRIMARY KEY (`user_id`, `post_id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci
  ROW_FORMAT = COMPACT COMMENT ='????????????????????????';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sys_user_post`
--

LOCK TABLES `sys_user_post` WRITE;
/*!40000 ALTER TABLE `sys_user_post`
    DISABLE KEYS */;
INSERT INTO `sys_user_post`
VALUES (1, 2),
       (1, 3);
/*!40000 ALTER TABLE `sys_user_post`
    ENABLE KEYS */;
UNLOCK TABLES;


create table if not exists gen_database
(
    id          bigint auto_increment comment '??????'
        primary key,
    `group`     varchar(30)                  not null comment '????????????',
    host        varchar(200) default ''      null comment '??????',
    port        varchar(10)  default ''      null comment '??????',
    user        varchar(64)                  null comment '??????',
    pass        varchar(64)                  null comment '??????',
    name        varchar(20)  default ''      null comment '???????????????',
    type        varchar(20)  default 'mysql' null comment '???????????????',
    create_by   varchar(64)  default ''      null comment '?????????',
    create_time datetime                     null comment '????????????',
    update_by   varchar(64)  default ''      null comment '?????????',
    update_time datetime                     null comment '????????????',
    remark      varchar(500)                 null comment '??????',
    constraint gen_database_group_uindex
        unique (`group`)
)
    comment '????????????????????????';

create table if not exists gen_table
(
    table_id          bigint auto_increment comment '??????'
        primary key,
    table_name        varchar(200) default ''     null comment '?????????',
    table_comment     varchar(500) default ''     null comment '?????????',
    tpl_category      varchar(200) default 'crud' null comment '??????????????????crud???????????? tree???????????? sub??????????????????',
    sub_table_name    varchar(64)                 null comment '?????????????????????',
    sub_table_fk_name varchar(64)                 null comment '????????????????????????',
    tree_code         varchar(50)                 null comment '???????????????',
    tree_name         varchar(50)                 null comment '???????????????',
    tree_parent_code  varchar(50)                 null comment '???????????????',
    class_name        varchar(100) default ''     null comment '???????????????',
    system_name       varchar(50)                 null comment '????????????',
    module_name       varchar(100)                null comment '???????????????',
    package_name      varchar(100)                null comment '???????????????',
    business_name     varchar(100)                null comment '???????????????',
    function_name     varchar(50)                 null comment '???????????????',
    function_author   varchar(50)                 null comment '??????????????????',
    gen_type          char         default '0'    null comment '?????????????????????0zip????????? 1??????????????????',
    gen_path          varchar(200) default '/'    null comment '??????????????????????????????????????????',
    params            text                        null comment '????????????',
    create_by         varchar(64)  default ''     null comment '?????????',
    create_time       datetime                    null comment '????????????',
    update_by         varchar(64)  default ''     null comment '?????????',
    update_time       datetime                    null comment '????????????',
    remark            varchar(500)                null comment '??????'
)
    comment '?????????????????????';

create table if not exists gen_table_column
(
    column_id      bigint auto_increment comment '??????'
        primary key,
    table_id       varchar(64)               null comment '???????????????',
    column_name    varchar(200)              null comment '?????????',
    column_comment varchar(500)              null comment '?????????',
    column_type    varchar(100)              null comment '?????????',
    column_length  int                       null comment '?????????',
    java_type      varchar(500)              null comment 'JAVA??????',
    java_field     varchar(200)              null comment 'JAVA?????????',
    go_type        varchar(500)              null comment 'GO??????',
    go_field       varchar(200)              null comment 'GO?????????',
    is_pk          char                      null comment '???????????????1??????',
    is_increment   char                      null comment '???????????????1??????',
    is_required    char                      null comment '???????????????1??????',
    is_insert      char                      null comment '????????????????????????1??????',
    is_edit        char                      null comment '?????????????????????1??????',
    is_list        char                      null comment '?????????????????????1??????',
    is_query       char                      null comment '?????????????????????1??????',
    query_type     varchar(200) default 'EQ' null comment '???????????????????????????????????????????????????????????????',
    html_type      varchar(200)              null comment '??????????????????????????????????????????????????????????????????????????????????????????',
    dict_type      varchar(200) default ''   null comment '????????????',
    sort           int                       null comment '??????',
    create_by      varchar(64)  default ''   null comment '?????????',
    create_time    datetime                  null comment '????????????',
    update_by      varchar(64)  default ''   null comment '?????????',
    update_time    datetime                  null comment '????????????'
)
    comment '???????????????????????????';



/*!40103 SET TIME_ZONE = @OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE = @OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS = @OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS = @OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT = @OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS = @OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION = @OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES = @OLD_SQL_NOTES */;

-- Dump completed on 2022-10-24 21:08:15
