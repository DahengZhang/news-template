DROP TABLE IF EXISTS `news`;

CREATE TABLE IF NOT EXISTS `news` (
    `nid` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `title` VARCHAR(100) NOT NULL,
    `preview` TEXT NOT NULL,
    `content` TEXT NOT NULL,
    `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,
    `update_time` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `news` (title, preview, content) VALUES ("title", "hello;你好啊", "<h1>Hello</h1><p>你好啊</p>")