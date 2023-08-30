-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_ability_bonus" table
CREATE TABLE `new_ability_bonus` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `bonus` integer NOT NULL, `ability_bonus_ability_score` integer NULL, `race_ability_bonuses` integer NULL, `subrace_ability_bonuses` integer NULL, CONSTRAINT `ability_bonus_ability_scores_ability_score` FOREIGN KEY (`ability_bonus_ability_score`) REFERENCES `ability_scores` (`id`) ON DELETE SET NULL, CONSTRAINT `ability_bonus_races_ability_bonuses` FOREIGN KEY (`race_ability_bonuses`) REFERENCES `races` (`id`) ON DELETE SET NULL, CONSTRAINT `ability_bonus_subraces_ability_bonuses` FOREIGN KEY (`subrace_ability_bonuses`) REFERENCES `subraces` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "ability_bonus" to new temporary table "new_ability_bonus"
INSERT INTO `new_ability_bonus` (`id`) SELECT `id` FROM `ability_bonus`;
-- Drop "ability_bonus" table after copying rows
DROP TABLE `ability_bonus`;
-- Rename temporary table "new_ability_bonus" to "ability_bonus"
ALTER TABLE `new_ability_bonus` RENAME TO `ability_bonus`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
