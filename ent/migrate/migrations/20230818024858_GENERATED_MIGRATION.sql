-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_ability_scores" table
CREATE TABLE `new_ability_scores` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NULL, `full_name` text NOT NULL, `class_saving_throws` integer NULL, `prerequisite_ability_score` integer NULL, CONSTRAINT `ability_scores_classes_saving_throws` FOREIGN KEY (`class_saving_throws`) REFERENCES `classes` (`id`) ON DELETE SET NULL, CONSTRAINT `ability_scores_prerequisites_ability_score` FOREIGN KEY (`prerequisite_ability_score`) REFERENCES `prerequisites` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "ability_scores" to new temporary table "new_ability_scores"
INSERT INTO `new_ability_scores` (`id`, `indx`, `name`, `desc`, `full_name`, `class_saving_throws`, `prerequisite_ability_score`) SELECT `id`, `indx`, `name`, `desc`, `full_name`, `class_saving_throws`, `prerequisite_ability_score` FROM `ability_scores`;
-- Drop "ability_scores" table after copying rows
DROP TABLE `ability_scores`;
-- Rename temporary table "new_ability_scores" to "ability_scores"
ALTER TABLE `new_ability_scores` RENAME TO `ability_scores`;
-- Create "ability_bonus_ability_score" table
CREATE TABLE `ability_bonus_ability_score` (`ability_bonus_id` integer NOT NULL, `ability_score_id` integer NOT NULL, PRIMARY KEY (`ability_bonus_id`, `ability_score_id`), CONSTRAINT `ability_bonus_ability_score_ability_bonus_id` FOREIGN KEY (`ability_bonus_id`) REFERENCES `ability_bonus` (`id`) ON DELETE CASCADE, CONSTRAINT `ability_bonus_ability_score_ability_score_id` FOREIGN KEY (`ability_score_id`) REFERENCES `ability_scores` (`id`) ON DELETE CASCADE);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
