-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_races" table
CREATE TABLE `new_races` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `alignment` text NOT NULL, `age` text NOT NULL, `size` text NOT NULL, `size_description` text NOT NULL, `language_desc` text NOT NULL, `speed` integer NOT NULL);
-- Set sequence for "new_races" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_races", 51539607552);
-- Copy rows from old table "races" to new temporary table "new_races"
INSERT INTO `new_races` (`id`, `indx`, `name`, `alignment`, `age`, `size`, `size_description`, `language_desc`, `speed`) SELECT `id`, `indx`, `name`, `alignment`, `age`, `size`, `size_description`, `language_desc`, `speed` FROM `races`;
-- Drop "races" table after copying rows
DROP TABLE `races`;
-- Rename temporary table "new_races" to "races"
ALTER TABLE `new_races` RENAME TO `races`;
-- Create index "races_indx_key" to table: "races"
CREATE UNIQUE INDEX `races_indx_key` ON `races` (`indx`);
-- Create "new_choices" table
CREATE TABLE `new_choices` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `choose` integer NOT NULL, `desc` text NOT NULL, `race_starting_proficiency_options` integer NULL, CONSTRAINT `choices_races_starting_proficiency_options` FOREIGN KEY (`race_starting_proficiency_options`) REFERENCES `races` (`id`) ON DELETE SET NULL);
-- Set sequence for "new_choices" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_choices", 98784247808);
-- Copy rows from old table "choices" to new temporary table "new_choices"
INSERT INTO `new_choices` (`id`, `choose`, `desc`) SELECT `id`, `choose`, `desc` FROM `choices`;
-- Drop "choices" table after copying rows
DROP TABLE `choices`;
-- Rename temporary table "new_choices" to "choices"
ALTER TABLE `new_choices` RENAME TO `choices`;
-- Create index "choices_race_starting_proficiency_options_key" to table: "choices"
CREATE UNIQUE INDEX `choices_race_starting_proficiency_options_key` ON `choices` (`race_starting_proficiency_options`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
