-- Create "ability_bonus" table
CREATE TABLE `ability_bonus` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `value` integer NOT NULL);
-- Set sequence for "ability_bonus" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("ability_bonus", 85899345920);
-- Create "subraces" table
CREATE TABLE `subraces` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `race_subrace` integer NULL, CONSTRAINT `subraces_races_subrace` FOREIGN KEY (`race_subrace`) REFERENCES `races` (`id`) ON DELETE SET NULL);
-- Set sequence for "subraces" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("subraces", 90194313216);
-- Create index "subraces_indx_key" to table: "subraces"
CREATE UNIQUE INDEX `subraces_indx_key` ON `subraces` (`indx`);
-- Create index "subraces_race_subrace_key" to table: "subraces"
CREATE UNIQUE INDEX `subraces_race_subrace_key` ON `subraces` (`race_subrace`);
-- Create "ability_bonus_ability_score" table
CREATE TABLE `ability_bonus_ability_score` (`ability_bonus_id` integer NOT NULL, `ability_score_id` integer NOT NULL, PRIMARY KEY (`ability_bonus_id`, `ability_score_id`), CONSTRAINT `ability_bonus_ability_score_ability_bonus_id` FOREIGN KEY (`ability_bonus_id`) REFERENCES `ability_bonus` (`id`) ON DELETE CASCADE, CONSTRAINT `ability_bonus_ability_score_ability_score_id` FOREIGN KEY (`ability_score_id`) REFERENCES `ability_scores` (`id`) ON DELETE CASCADE);
-- Add pk ranges for ('ability_bonus'),('subraces') tables
INSERT INTO `ent_types` (`type`) VALUES ('ability_bonus'), ('subraces');
