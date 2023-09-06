-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_ability_scores" table
CREATE TABLE `new_ability_scores` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `full_name` text NOT NULL, `desc` json NOT NULL, `proficiency_saving_throw` integer NULL, CONSTRAINT `ability_scores_proficiencies_saving_throw` FOREIGN KEY (`proficiency_saving_throw`) REFERENCES `proficiencies` (`id`) ON DELETE SET NULL);
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
-- Create "new_equipment" table
CREATE TABLE `new_equipment` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `weight` integer NULL);
-- Set sequence for "new_equipment" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_equipment", 34359738368);
-- Copy rows from old table "equipment" to new temporary table "new_equipment"
INSERT INTO `new_equipment` (`id`, `indx`, `name`) SELECT `id`, `indx`, `name` FROM `equipment`;
-- Drop "equipment" table after copying rows
DROP TABLE `equipment`;
-- Rename temporary table "new_equipment" to "equipment"
ALTER TABLE `new_equipment` RENAME TO `equipment`;
-- Create index "equipment_indx_key" to table: "equipment"
CREATE UNIQUE INDEX `equipment_indx_key` ON `equipment` (`indx`);
-- Create "new_gears" table
CREATE TABLE `new_gears` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `gear_category` text NOT NULL DEFAULT 'other', `quantity` integer NULL, `equipment_id` integer NOT NULL, CONSTRAINT `gears_equipment_gear` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE NO ACTION);
-- Set sequence for "new_gears" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_gears", 38654705664);
-- Copy rows from old table "gears" to new temporary table "new_gears"
INSERT INTO `new_gears` (`id`, `indx`, `name`, `gear_category`, `quantity`, `equipment_id`) SELECT `id`, `indx`, `name`, `gear_category`, `quantity`, `equipment_id` FROM `gears`;
-- Drop "gears" table after copying rows
DROP TABLE `gears`;
-- Rename temporary table "new_gears" to "gears"
ALTER TABLE `new_gears` RENAME TO `gears`;
-- Create index "gears_indx_key" to table: "gears"
CREATE UNIQUE INDEX `gears_indx_key` ON `gears` (`indx`);
-- Create index "gears_equipment_id_key" to table: "gears"
CREATE UNIQUE INDEX `gears_equipment_id_key` ON `gears` (`equipment_id`);
-- Create "new_proficiencies" table
CREATE TABLE `new_proficiencies` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `proficiency_category` text NOT NULL);
-- Set sequence for "new_proficiencies" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_proficiencies", 51539607552);
-- Copy rows from old table "proficiencies" to new temporary table "new_proficiencies"
INSERT INTO `new_proficiencies` (`id`, `indx`, `name`, `proficiency_category`) SELECT `id`, `indx`, `name`, `proficiency_category` FROM `proficiencies`;
-- Drop "proficiencies" table after copying rows
DROP TABLE `proficiencies`;
-- Rename temporary table "new_proficiencies" to "proficiencies"
ALTER TABLE `new_proficiencies` RENAME TO `proficiencies`;
-- Create index "proficiencies_indx_key" to table: "proficiencies"
CREATE UNIQUE INDEX `proficiencies_indx_key` ON `proficiencies` (`indx`);
-- Create "new_races" table
CREATE TABLE `new_races` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `alignment` text NOT NULL, `age` text NOT NULL, `size` text NOT NULL, `size_description` text NOT NULL, `language_desc` text NOT NULL, `speed` integer NOT NULL);
-- Set sequence for "new_races" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_races", 55834574848);
-- Copy rows from old table "races" to new temporary table "new_races"
INSERT INTO `new_races` (`id`, `indx`, `name`, `alignment`, `age`, `size`, `size_description`, `language_desc`, `speed`) SELECT `id`, `indx`, `name`, `alignment`, `age`, `size`, `size_description`, `language_desc`, `speed` FROM `races`;
-- Drop "races" table after copying rows
DROP TABLE `races`;
-- Rename temporary table "new_races" to "races"
ALTER TABLE `new_races` RENAME TO `races`;
-- Create index "races_indx_key" to table: "races"
CREATE UNIQUE INDEX `races_indx_key` ON `races` (`indx`);
-- Create "new_skills" table
CREATE TABLE `new_skills` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` json NOT NULL, `proficiency_skill` integer NULL, `skill_ability_score` integer NULL, CONSTRAINT `skills_proficiencies_skill` FOREIGN KEY (`proficiency_skill`) REFERENCES `proficiencies` (`id`) ON DELETE SET NULL, CONSTRAINT `skills_ability_scores_ability_score` FOREIGN KEY (`skill_ability_score`) REFERENCES `ability_scores` (`id`) ON DELETE SET NULL);
-- Set sequence for "new_skills" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_skills", 68719476736);
-- Copy rows from old table "skills" to new temporary table "new_skills"
INSERT INTO `new_skills` (`id`, `indx`, `name`, `desc`, `skill_ability_score`) SELECT `id`, `indx`, `name`, `desc`, `skill_ability_score` FROM `skills`;
-- Drop "skills" table after copying rows
DROP TABLE `skills`;
-- Rename temporary table "new_skills" to "skills"
ALTER TABLE `new_skills` RENAME TO `skills`;
-- Create index "skills_indx_key" to table: "skills"
CREATE UNIQUE INDEX `skills_indx_key` ON `skills` (`indx`);
-- Create "new_subraces" table
CREATE TABLE `new_subraces` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL, `subrace_race` integer NULL, CONSTRAINT `subraces_races_race` FOREIGN KEY (`subrace_race`) REFERENCES `races` (`id`) ON DELETE SET NULL);
-- Set sequence for "new_subraces" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_subraces", 73014444032);
-- Copy rows from old table "subraces" to new temporary table "new_subraces"
INSERT INTO `new_subraces` (`id`, `indx`, `name`, `desc`) SELECT `id`, `indx`, `name`, `desc` FROM `subraces`;
-- Drop "subraces" table after copying rows
DROP TABLE `subraces`;
-- Rename temporary table "new_subraces" to "subraces"
ALTER TABLE `new_subraces` RENAME TO `subraces`;
-- Create index "subraces_indx_key" to table: "subraces"
CREATE UNIQUE INDEX `subraces_indx_key` ON `subraces` (`indx`);
-- Create "new_class_proficiency_choices" table
CREATE TABLE `new_class_proficiency_choices` (`class_id` integer NOT NULL, `proficiency_choice_id` integer NOT NULL, PRIMARY KEY (`class_id`, `proficiency_choice_id`), CONSTRAINT `class_proficiency_choices_class_id` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`) ON DELETE CASCADE, CONSTRAINT `class_proficiency_choices_proficiency_choice_id` FOREIGN KEY (`proficiency_choice_id`) REFERENCES `proficiency_choices` (`id`) ON DELETE CASCADE);
-- Copy rows from old table "class_proficiency_choices" to new temporary table "new_class_proficiency_choices"
INSERT INTO `new_class_proficiency_choices` (`class_id`) SELECT `class_id` FROM `class_proficiency_choices`;
-- Drop "class_proficiency_choices" table after copying rows
DROP TABLE `class_proficiency_choices`;
-- Rename temporary table "new_class_proficiency_choices" to "class_proficiency_choices"
ALTER TABLE `new_class_proficiency_choices` RENAME TO `class_proficiency_choices`;
-- Create "class_equipments" table
CREATE TABLE `class_equipments` (`quantity` integer NOT NULL, `class_id` integer NOT NULL, `equipment_id` integer NOT NULL, PRIMARY KEY (`class_id`, `equipment_id`), CONSTRAINT `class_equipments_classes_class` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`) ON DELETE NO ACTION, CONSTRAINT `class_equipments_equipment_equipment` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE NO ACTION);
-- Create "coins" table
CREATE TABLE `coins` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `desc` text NOT NULL, `gold_conversion_rate` real NOT NULL);
-- Set sequence for "coins" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("coins", 103079215104);
-- Create index "coins_indx_key" to table: "coins"
CREATE UNIQUE INDEX `coins_indx_key` ON `coins` (`indx`);
-- Create "equipment_categories" table
CREATE TABLE `equipment_categories` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `name` text NOT NULL DEFAULT 'other');
-- Set sequence for "equipment_categories" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("equipment_categories", 107374182400);
-- Create "equipment_choices" table
CREATE TABLE `equipment_choices` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `choose` integer NOT NULL, `desc` text NULL);
-- Set sequence for "equipment_choices" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("equipment_choices", 111669149696);
-- Create "equipment_costs" table
CREATE TABLE `equipment_costs` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `quantity` integer NOT NULL, `gp_value` real NOT NULL, `equipment_id` integer NOT NULL, `coin_id` integer NOT NULL, CONSTRAINT `equipment_costs_equipment_cost` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE NO ACTION, CONSTRAINT `equipment_costs_coins_coin` FOREIGN KEY (`coin_id`) REFERENCES `coins` (`id`) ON DELETE NO ACTION);
-- Set sequence for "equipment_costs" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("equipment_costs", 115964116992);
-- Create index "equipment_costs_equipment_id_key" to table: "equipment_costs"
CREATE UNIQUE INDEX `equipment_costs_equipment_id_key` ON `equipment_costs` (`equipment_id`);
-- Create "proficiency_choices" table
CREATE TABLE `proficiency_choices` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `choose` integer NOT NULL, `desc` text NULL, `proficiency_choice_sub_choice` integer NULL, CONSTRAINT `proficiency_choices_proficiency_choices_sub_choice` FOREIGN KEY (`proficiency_choice_sub_choice`) REFERENCES `proficiency_choices` (`id`) ON DELETE SET NULL);
-- Set sequence for "proficiency_choices" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("proficiency_choices", 120259084288);
-- Create "class_equipment_choices" table
CREATE TABLE `class_equipment_choices` (`class_id` integer NOT NULL, `equipment_choice_id` integer NOT NULL, PRIMARY KEY (`class_id`, `equipment_choice_id`), CONSTRAINT `class_equipment_choices_class_id` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`) ON DELETE CASCADE, CONSTRAINT `class_equipment_choices_equipment_choice_id` FOREIGN KEY (`equipment_choice_id`) REFERENCES `equipment_choices` (`id`) ON DELETE CASCADE);
-- Create "equipment_choice" table
CREATE TABLE `equipment_choice` (`equipment_id` integer NOT NULL, `equipment_choice_id` integer NOT NULL, PRIMARY KEY (`equipment_id`, `equipment_choice_id`), CONSTRAINT `equipment_choice_equipment_id` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE CASCADE, CONSTRAINT `equipment_choice_equipment_choice_id` FOREIGN KEY (`equipment_choice_id`) REFERENCES `equipment_choices` (`id`) ON DELETE CASCADE);
-- Create "equipment_category_equipment" table
CREATE TABLE `equipment_category_equipment` (`equipment_category_id` integer NOT NULL, `equipment_id` integer NOT NULL, PRIMARY KEY (`equipment_category_id`, `equipment_id`), CONSTRAINT `equipment_category_equipment_equipment_category_id` FOREIGN KEY (`equipment_category_id`) REFERENCES `equipment_categories` (`id`) ON DELETE CASCADE, CONSTRAINT `equipment_category_equipment_equipment_id` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE CASCADE);
-- Create "proficiency_choice" table
CREATE TABLE `proficiency_choice` (`proficiency_id` integer NOT NULL, `proficiency_choice_id` integer NOT NULL, PRIMARY KEY (`proficiency_id`, `proficiency_choice_id`), CONSTRAINT `proficiency_choice_proficiency_id` FOREIGN KEY (`proficiency_id`) REFERENCES `proficiencies` (`id`) ON DELETE CASCADE, CONSTRAINT `proficiency_choice_proficiency_choice_id` FOREIGN KEY (`proficiency_choice_id`) REFERENCES `proficiency_choices` (`id`) ON DELETE CASCADE);
-- Create "proficiency_equipment" table
CREATE TABLE `proficiency_equipment` (`proficiency_id` integer NOT NULL, `equipment_id` integer NOT NULL, PRIMARY KEY (`proficiency_id`, `equipment_id`), CONSTRAINT `proficiency_equipment_proficiency_id` FOREIGN KEY (`proficiency_id`) REFERENCES `proficiencies` (`id`) ON DELETE CASCADE, CONSTRAINT `proficiency_equipment_equipment_id` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE CASCADE);
-- Create "race_proficiency_choice" table
CREATE TABLE `race_proficiency_choice` (`race_id` integer NOT NULL, `proficiency_choice_id` integer NOT NULL, PRIMARY KEY (`race_id`, `proficiency_choice_id`), CONSTRAINT `race_proficiency_choice_race_id` FOREIGN KEY (`race_id`) REFERENCES `races` (`id`) ON DELETE CASCADE, CONSTRAINT `race_proficiency_choice_proficiency_choice_id` FOREIGN KEY (`proficiency_choice_id`) REFERENCES `proficiency_choices` (`id`) ON DELETE CASCADE);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
-- Add pk ranges for ('coins'),('equipment_categories'),('equipment_choices'),('equipment_costs'),('proficiency_choices') tables
INSERT INTO `ent_types` (`type`) VALUES ('coins'), ('equipment_categories'), ('equipment_choices'), ('equipment_costs'), ('proficiency_choices');
