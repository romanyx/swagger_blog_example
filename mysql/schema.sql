CREATE DATABASE IF NOT EXISTS blog;
GRANT ALL PRIVILEGES ON blog.* to 'blog'@'%';
USE blog;

CREATE TABLE IF NOT EXISTS `blogs` (
	`id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`title` varchar(255) NOT NULL,
	`markdown` text NOT NULL,
	`published_at` int(11) UNSIGNED
)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_general_ci;