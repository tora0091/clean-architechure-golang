CREATE DATABASE IF NOT EXISTS `your-project-database`;
USE `your-project-database`;

create table if not exists `music` (
    `id` integer primary key,
    `iswc` varchar(11) not null,
    `title` varchar(250) not null,
    `time` varchar(100) not null,
    `genre` varchar(100) not null,
    `artist_id` integer,
    `company_id` integer,
    `created_at` datetime default current_timestamp,
    `updated_at` datetime default current_timestamp,
    `deleted_at` datetime
);

create table  if not exists `artist` (
    `id` integer primary key,
    `name` varchar(100) not null,
    `email` varchar(250) not null,
    `birth` varchar(20) not null,
    `gender` varchar(20) not null,
    `address` varchar(200) not null,
    `created_at` datetime default current_timestamp,
    `updated_at` datetime default current_timestamp,
    `deleted_at` datetime
);

create table if not exists `company` (
    `id` integer not null primary key,
    `name` varchar(100) not null,
    `email` varchar(250) not null,
    `address` varchar(200) not null,
    `created_at` datetime default current_timestamp,
    `updated_at` datetime default current_timestamp,
    `deleted_at` datetime
);
