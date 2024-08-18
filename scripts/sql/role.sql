CREATE TABLE roles_table (
                             id BIGINT PRIMARY KEY AUTO_INCREMENT,
                             created_at TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
                             status TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态 1 正常 2 禁用',
                             name VARCHAR(30) NOT NULL COMMENT '角色名',
                             code VARCHAR(12) NOT NULL COMMENT '角色码，用于前端权限控制',
                             remark TEXT COMMENT '备注',
                             sort INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序编号'
);
