-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_languages" table
CREATE TABLE `new_languages` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL, `language_type` text NOT NULL DEFAULT 'STANDARD', `script` text NULL DEFAULT 'Common', `choice_languages` integer NULL, CONSTRAINT `languages_choices_languages` FOREIGN KEY (`choice_languages`) REFERENCES `choices` (`id`) ON DELETE SET NULL);
-- Set sequence for "new_languages" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_languages", 38654705664);
-- Copy rows from old table "languages" to new temporary table "new_languages"
INSERT INTO `new_languages` (`id`, `indx`, `name`, `desc`, `language_type`, `script`) SELECT `id`, `indx`, `name`, `desc`, `language_type`, `script` FROM `languages`;
-- Drop "languages" table after copying rows
DROP TABLE `languages`;
-- Rename temporary table "new_languages" to "languages"
ALTER TABLE `new_languages` RENAME TO `languages`;
-- Create index "languages_indx_key" to table: "languages"
CREATE UNIQUE INDEX `languages_indx_key` ON `languages` (`indx`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
