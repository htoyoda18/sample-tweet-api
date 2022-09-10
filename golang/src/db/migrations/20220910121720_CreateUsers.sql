-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `users` (
      `id` INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
      `email` VARCHAR(255) NOT NULL UNIQUE,
      `name` VARCHAR(63),
      `icon` TEXT,
      `password` VARCHAR(255),
)

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `users`;