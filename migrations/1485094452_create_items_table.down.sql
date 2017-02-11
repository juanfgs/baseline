TRUNCATE `items`;
ALTER TABLE `items` DROP FOREIGN KEY fk_section;
DROP TABLE `items`;
