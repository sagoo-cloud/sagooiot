# TDengine 数据库迁移工具

> 简单的 TDengine 数据库迁移工具。  
> 该工具在数据迁移时，会删除目标数据库的超级表、子表的表结构和数据，请注意备份！！

使用方式：直接动行即可。

## 一、TDengine 配置
1. tdengineSrc：源数据库连接配置
2. tdengineDest：目标数据库连接配置

## 二、程序运行时迁移步骤说明
1. 创建超级表
    - 获取源数据库的超级表结构
    - 删除目标数据库超级表结构
    - 创建新超级表
2. 创建子表
    - 获取源数据库的子表结构
    - 删除目标数据库子表结构
    - 创建新子表
3. 数据迁移
