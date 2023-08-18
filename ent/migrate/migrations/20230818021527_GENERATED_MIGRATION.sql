-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_languages" table
CREATE TABLE `new_languages` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL, `category` text NOT NULL DEFAULT 'standard', `script` text NOT NULL);
-- Copy rows from old table "languages" to new temporary table "new_languages"
INSERT INTO `new_languages` (`id`, `indx`, `name`, `desc`, `script`) SELECT `id`, `indx`, `name`, `desc`, `script` FROM `languages`;
-- Drop "languages" table after copying rows
DROP TABLE `languages`;
-- Rename temporary table "new_languages" to "languages"
ALTER TABLE `new_languages` RENAME TO `languages`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
