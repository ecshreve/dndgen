-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Add column "alignment" to table: "races"
ALTER TABLE `races` ADD COLUMN `alignment` text NOT NULL;
-- Add column "age" to table: "races"
ALTER TABLE `races` ADD COLUMN `age` text NOT NULL;
-- Add column "size" to table: "races"
ALTER TABLE `races` ADD COLUMN `size` text NOT NULL;
-- Add column "size_description" to table: "races"
ALTER TABLE `races` ADD COLUMN `size_description` text NOT NULL;
-- Add column "language_desc" to table: "races"
ALTER TABLE `races` ADD COLUMN `language_desc` text NOT NULL;
-- Create "new_ability_bonus" table
CREATE TABLE `new_ability_bonus` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `bonus` integer NOT NULL, `ability_score_id` integer NOT NULL, `race_ability_bonuses` integer NULL, `subrace_ability_bonuses` integer NULL, CONSTRAINT `ability_bonus_ability_scores_ability_bonuses` FOREIGN KEY (`ability_score_id`) REFERENCES `ability_scores` (`id`) ON DELETE NO ACTION, CONSTRAINT `ability_bonus_races_ability_bonuses` FOREIGN KEY (`race_ability_bonuses`) REFERENCES `races` (`id`) ON DELETE SET NULL, CONSTRAINT `ability_bonus_subraces_ability_bonuses` FOREIGN KEY (`subrace_ability_bonuses`) REFERENCES `subraces` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "ability_bonus" to new temporary table "new_ability_bonus"
INSERT INTO `new_ability_bonus` (`id`, `bonus`, `ability_score_id`) SELECT `id`, `bonus`, `ability_score_id` FROM `ability_bonus`;
-- Drop "ability_bonus" table after copying rows
DROP TABLE `ability_bonus`;
-- Rename temporary table "new_ability_bonus" to "ability_bonus"
ALTER TABLE `new_ability_bonus` RENAME TO `ability_bonus`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
