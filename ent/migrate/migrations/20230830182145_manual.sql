-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_ability_scores" table
CREATE TABLE `new_ability_scores` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `full_name` text NOT NULL, `desc` json NOT NULL);
-- Set sequence for "new_ability_scores" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_ability_scores", 4294967296);
-- Copy rows from old table "ability_scores" to new temporary table "new_ability_scores"
INSERT INTO `new_ability_scores` (`id`, `indx`, `name`, `full_name`, `desc`) SELECT `id`, `indx`, `name`, `full_name`, `desc` FROM `ability_scores`;
-- Drop "ability_scores" table after copying rows
DROP TABLE `ability_scores`;
-- Rename temporary table "new_ability_scores" to "ability_scores"
ALTER TABLE `new_ability_scores` RENAME TO `ability_scores`;
-- Create index "ability_scores_indx_key" to table: "ability_scores"
CREATE UNIQUE INDEX `ability_scores_indx_key` ON `ability_scores` (`indx`);
-- Create "new_ability_bonus" table
CREATE TABLE `new_ability_bonus` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `bonus` integer NOT NULL, `ability_score_id` integer NOT NULL, `race_race_ability_bonuses` integer NULL, `subrace_subrace_ability_bonuses` integer NULL, CONSTRAINT `ability_bonus_ability_scores_ability_bonuses` FOREIGN KEY (`ability_score_id`) REFERENCES `ability_scores` (`id`) ON DELETE NO ACTION, CONSTRAINT `ability_bonus_races_race_ability_bonuses` FOREIGN KEY (`race_race_ability_bonuses`) REFERENCES `races` (`id`) ON DELETE SET NULL, CONSTRAINT `ability_bonus_subraces_subrace_ability_bonuses` FOREIGN KEY (`subrace_subrace_ability_bonuses`) REFERENCES `subraces` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "ability_bonus" to new temporary table "new_ability_bonus"
INSERT INTO `new_ability_bonus` (`id`, `bonus`) SELECT `id`, `bonus` FROM `ability_bonus`;
-- Drop "ability_bonus" table after copying rows
DROP TABLE `ability_bonus`;
-- Rename temporary table "new_ability_bonus" to "ability_bonus"
ALTER TABLE `new_ability_bonus` RENAME TO `ability_bonus`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
