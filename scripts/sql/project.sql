CREATE TABLE  projects_table (
                             id BIGINT PRIMARY KEY AUTO_INCREMENT,
                             created_at TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
                             updated_at TIMESTAMP NOT NULL default CURRENT_TIMESTAMP,
                             deleted_at TIMESTAMP default NULL,
                             project_status TINYINT UNSIGNED NOT NULL DEFAULT 1 COMMENT '状态 1 正常 2 禁用',
                             project_name VARCHAR(30) NOT NULL COMMENT '项目名',
                             remark TEXT COMMENT '备注',
                             sort INT UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序编号'
);
