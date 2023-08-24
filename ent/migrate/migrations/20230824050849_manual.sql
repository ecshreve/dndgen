-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_ability_scores" table
CREATE TABLE `new_ability_scores` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `full_name` text NOT NULL, `desc` json NOT NULL, `proficiency_saving_throw` integer NULL, CONSTRAINT `ability_scores_proficiencies_saving_throw` FOREIGN KEY (`proficiency_saving_throw`) REFERENCES `proficiencies` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "ability_scores" to new temporary table "new_ability_scores"
INSERT INTO `new_ability_scores` (`id`, `indx`, `name`, `full_name`, `desc`) SELECT `id`, `indx`, `name`, `full_name`, `desc` FROM `ability_scores`;
-- Drop "ability_scores" table after copying rows
DROP TABLE `ability_scores`;
-- Rename temporary table "new_ability_scores" to "ability_scores"
ALTER TABLE `new_ability_scores` RENAME TO `ability_scores`;
-- Create index "ability_scores_indx_key" to table: "ability_scores"
CREATE UNIQUE INDEX `ability_scores_indx_key` ON `ability_scores` (`indx`);
-- Create index "ability_scores_proficiency_saving_throw_key" to table: "ability_scores"
CREATE UNIQUE INDEX `ability_scores_proficiency_saving_throw_key` ON `ability_scores` (`proficiency_saving_throw`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
