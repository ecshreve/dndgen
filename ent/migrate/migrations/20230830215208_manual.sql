-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_proficiencies" table
CREATE TABLE `new_proficiencies` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `proficiency_category` text NOT NULL, `proficiency_skill` integer NULL, `proficiency_equipment` integer NULL, `proficiency_saving_throw` integer NULL, CONSTRAINT `proficiencies_skills_skill` FOREIGN KEY (`proficiency_skill`) REFERENCES `skills` (`id`) ON DELETE SET NULL, CONSTRAINT `proficiencies_equipment_equipment` FOREIGN KEY (`proficiency_equipment`) REFERENCES `equipment` (`id`) ON DELETE SET NULL, CONSTRAINT `proficiencies_ability_scores_saving_throw` FOREIGN KEY (`proficiency_saving_throw`) REFERENCES `ability_scores` (`id`) ON DELETE SET NULL);
-- Set sequence for "new_proficiencies" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_proficiencies", 47244640256);
-- Copy rows from old table "proficiencies" to new temporary table "new_proficiencies"
INSERT INTO `new_proficiencies` (`id`, `indx`, `name`, `proficiency_category`, `proficiency_skill`, `proficiency_equipment`, `proficiency_saving_throw`) SELECT `id`, `indx`, `name`, `proficiency_category`, `proficiency_skill`, `proficiency_equipment`, `proficiency_saving_throw` FROM `proficiencies`;
-- Drop "proficiencies" table after copying rows
DROP TABLE `proficiencies`;
-- Rename temporary table "new_proficiencies" to "proficiencies"
ALTER TABLE `new_proficiencies` RENAME TO `proficiencies`;
-- Create index "proficiencies_indx_key" to table: "proficiencies"
CREATE UNIQUE INDEX `proficiencies_indx_key` ON `proficiencies` (`indx`);
-- Create "new_races" table
CREATE TABLE `new_races` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `alignment` text NOT NULL, `age` text NOT NULL, `size` text NOT NULL, `size_description` text NOT NULL, `language_desc` text NOT NULL, `speed` integer NOT NULL, `race_starting_proficiency_options` integer NULL, `race_language_options` integer NULL, CONSTRAINT `races_choices_starting_proficiency_options` FOREIGN KEY (`race_starting_proficiency_options`) REFERENCES `choices` (`id`) ON DELETE SET NULL, CONSTRAINT `races_choices_language_options` FOREIGN KEY (`race_language_options`) REFERENCES `choices` (`id`) ON DELETE SET NULL);
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
CREATE TABLE `new_choices` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `choose` integer NOT NULL, `desc` text NULL);
-- Set sequence for "new_choices" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_choices", 98784247808);
-- Copy rows from old table "choices" to new temporary table "new_choices"
INSERT INTO `new_choices` (`id`, `choose`, `desc`) SELECT `id`, `choose`, `desc` FROM `choices`;
-- Drop "choices" table after copying rows
DROP TABLE `choices`;
-- Rename temporary table "new_choices" to "choices"
ALTER TABLE `new_choices` RENAME TO `choices`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
