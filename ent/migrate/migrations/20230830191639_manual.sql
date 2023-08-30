-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_proficiencies" table
CREATE TABLE `new_proficiencies` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `proficiency_category` text NOT NULL, `choice_proficiencies` integer NULL, `proficiency_skill` integer NULL, `proficiency_equipment` integer NULL, `proficiency_saving_throw` integer NULL, CONSTRAINT `proficiencies_choices_proficiencies` FOREIGN KEY (`choice_proficiencies`) REFERENCES `choices` (`id`) ON DELETE SET NULL, CONSTRAINT `proficiencies_skills_skill` FOREIGN KEY (`proficiency_skill`) REFERENCES `skills` (`id`) ON DELETE SET NULL, CONSTRAINT `proficiencies_equipment_equipment` FOREIGN KEY (`proficiency_equipment`) REFERENCES `equipment` (`id`) ON DELETE SET NULL, CONSTRAINT `proficiencies_ability_scores_saving_throw` FOREIGN KEY (`proficiency_saving_throw`) REFERENCES `ability_scores` (`id`) ON DELETE SET NULL);
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
-- Create "new_languages" table
CREATE TABLE `new_languages` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL, `language_type` text NOT NULL DEFAULT 'STANDARD', `script` text NULL DEFAULT 'Common');
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
