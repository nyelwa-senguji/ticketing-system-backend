CREATE TABLE `roles` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `role_name` varchar(255) NOT NULL,
  `status` varchar(255) NOT NULL
);

CREATE TABLE `permission` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `permission_name` varchar(255) NOT NULL,
  `status` varchar(255) NOT NULL
);

CREATE TABLE `users` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role_id` int 
);

CREATE TABLE `permission_roles` (
  `permission_id` int ,
  `role_id` int 
);

CREATE TABLE `category` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `category_name` varchar(255) NOT NULL,
  `status` varchar(255) NOT NULL
);

CREATE TABLE `tickets` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `subject` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `status` varchar(255) NOT NULL,
  `user_id` int ,
  `category_id` int 
);

CREATE INDEX `users_index_0` ON `users` (`role_id`);

CREATE INDEX `tickets_index_1` ON `tickets` (`user_id`);

CREATE INDEX `tickets_index_2` ON `tickets` (`category_id`);

ALTER TABLE `users` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

ALTER TABLE `permission_roles` ADD FOREIGN KEY (`permission_id`) REFERENCES `permission` (`id`);

ALTER TABLE `permission_roles` ADD FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

ALTER TABLE `tickets` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `tickets` ADD FOREIGN KEY (`category_id`) REFERENCES `category` (`id`);