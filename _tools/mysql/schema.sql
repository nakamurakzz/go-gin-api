CREATE TABLE
    `site` (
        `id` BIGINT NOT NULL auto_increment,
        `name` VARCHAR(255) NOT NULL,
        `isEnable` TINYINT NOT NULL,
        `createdAt` DATETIME NOT NULL,
        `updatedAt` DATETIME NOT NULL,
        PRIMARY KEY (`id`) UNIQUE KEY `name` (`name`)
    ) Engine = InnoDB DEFAULT CHARSET = utf8 COMMENT = 'Site';