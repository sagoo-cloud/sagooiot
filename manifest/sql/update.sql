-- ----------------------------
-- 2023-09-08 修改字典类型表
-- ----------------------------
ALTER TABLE `sagoo_iot`.`sys_dict_data`
    ADD COLUMN `dict_class` varchar(10) NOT NULL DEFAULT 0 COMMENT '字典分类：取值从字典类型' AFTER `remark`;
ALTER TABLE `sagoo_iot`.`sys_dict_type`
    MODIFY COLUMN `is_deleted` int NOT NULL DEFAULT 0 COMMENT '是否删除 0未删除 1已删除' AFTER `status`;
ALTER TABLE `sagoo_iot`.`sys_dict_type`
    MODIFY COLUMN `module_classify` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '模块分类' AFTER `dict_type`;
