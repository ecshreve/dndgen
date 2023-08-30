-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_subraces" table
CREATE TABLE `new_subraces` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL, `race_subraces` integer NULL, CONSTRAINT `subraces_races_subraces` FOREIGN KEY (`race_subraces`) REFERENCES `races` (`id`) ON DELETE SET NULL);
-- Set sequence for "new_subraces" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_subraces", 68719476736);
-- Copy rows from old table "subraces" to new temporary table "new_subraces"
INSERT INTO `new_subraces` (`id`, `indx`, `name`, `desc`) SELECT `id`, `indx`, `name`, `desc` FROM `subraces`;
-- Drop "subraces" table after copying rows
DROP TABLE `subraces`;
-- Rename temporary table "new_subraces" to "subraces"
ALTER TABLE `new_subraces` RENAME TO `subraces`;
-- Create index "subraces_indx_key" to table: "subraces"
CREATE UNIQUE INDEX `subraces_indx_key` ON `subraces` (`indx`);
-- Create "new_ability_bonus" table
CREATE TABLE `new_ability_bonus` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `bonus` integer NOT NULL, `ability_bonus_ability_score` integer NULL, `race_ability_bonuses` integer NULL, CONSTRAINT `ability_bonus_ability_scores_ability_score` FOREIGN KEY (`ability_bonus_ability_score`) REFERENCES `ability_scores` (`id`) ON DELETE SET NULL, CONSTRAINT `ability_bonus_races_ability_bonuses` FOREIGN KEY (`race_ability_bonuses`) REFERENCES `races` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "ability_bonus" to new temporary table "new_ability_bonus"
INSERT INTO `new_ability_bonus` (`id`, `bonus`, `ability_bonus_ability_score`, `race_ability_bonuses`) SELECT `id`, `bonus`, `ability_bonus_ability_score`, `race_ability_bonuses` FROM `ability_bonus`;
-- Drop "ability_bonus" table after copying rows
DROP TABLE `ability_bonus`;
-- Rename temporary table "new_ability_bonus" to "ability_bonus"
ALTER TABLE `new_ability_bonus` RENAME TO `ability_bonus`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
