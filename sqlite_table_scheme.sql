CREATE DATABASE IF NOT EXISTS `your-project-database`;
USE `your-project-database`;

CREATE TABLE  IF NOT EXISTS `music` (
    `id` INT NOT NULL,
    `iswc` VARCHAR(11) NOT NULL,
    `title` VARCHAR(250) NOT NULL,
    `time` VARCHAR(100) NOT NULL,
    `genre` VARCHAR(100) NOT NULL,
    `artist_id` INT,
    `company_id` INT,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE  IF NOT EXISTS `artist` (
    `id` INT NOT NULL,
    `name` VARCHAR(100) NOT NULL,
    `email` VARCHAR(250) NOT NULL,
    `birth` VARCHAR(20) NOT NULL,
    `gender` VARCHAR(20) NOT NULL,
    `address` VARCHAR(200) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP,
    PRIMARY KEY (`id`)
);

CREATE TABLE  IF NOT EXISTS `company` (
    `id` INT NOT NULL,
    `name` VARCHAR(100) NOT NULL,
    `email` VARCHAR(250) NOT NULL,
    `address` VARCHAR(200) NOT NULL,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP,
    PRIMARY KEY (`id`)
);