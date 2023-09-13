-- ----------------------------
-- 2023-09-10 修改配置表
-- ----------------------------
ALTER TABLE `sagoo_iot`.`sys_config`
    ADD COLUMN `dict_class_code` varchar(255) NOT NULL DEFAULT 0 COMMENT '所属字典类型数据code' AFTER `config_type`;