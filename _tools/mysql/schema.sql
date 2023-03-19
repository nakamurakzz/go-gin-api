CREATE TABLE
    `site` (
        `id` BIGINT NOT NULL AUTO_INCREMENT,
        `name` VARCHAR(255) NOT NULL,
        `isEnabled` TINYINT(1) NOT NULL,
        `createdAt` DATETIME NOT NULL,
        `updatedAt` DATETIME NOT NULL,
        PRIMARY KEY (`id`),
        UNIQUE KEY `name` (`name`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8 COMMENT = 'Site';