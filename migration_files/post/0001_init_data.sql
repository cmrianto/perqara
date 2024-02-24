INSERT INTO `clients` (`name`, `created_at`, `updated_at`, `deleted_at`)
VALUES ('Client Name', '2024-02-23 12:34:56', '2024-02-23 12:34:56', NULL);
INSERT INTO `users` (`name`, `user_name`, `type`, `password`, `created_at`, `updated_at`, `deleted_at`, `client_id`)
VALUES ('Admin User', 'admin', 1, '$2a$10$SA5o/5hb.yqBLoz/ttqiV.shiQCK/lr4qNhq6kvDz/mau4et6SO/y', '2024-02-23 12:34:56', '2024-02-23 12:34:56', NULL, NULL);
INSERT INTO `roles` (`name`, `created_at`, `updated_at`, `deleted_at`)
VALUES ('admin', '2024-02-23 12:34:56', '2024-02-23 12:34:56', NULL);
INSERT INTO `user_roles` (`user_id`, `role_id`, `created_at`, `updated_at`, `deleted_at`)
VALUES (1, 1, '2024-02-23 12:34:56', '2024-02-23 12:34:56', NULL);
INSERT INTO `role_permissions` (`role_id`, `permission_id`, `created_at`, `updated_at`, `deleted_at`)
VALUES (1, 1, '2024-02-23 12:34:56', '2024-02-23 12:34:56', NULL),
(1, 2, '2024-02-23 12:34:56', '2024-02-23 12:34:56', NULL),
(1, 3, '2024-02-23 12:34:56', '2024-02-23 12:34:56', NULL),
(1, 4, '2024-02-23 12:34:56', '2024-02-23 12:34:56', NULL);