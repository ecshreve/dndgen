-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_ability_scores" table
CREATE TABLE `new_ability_scores` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NULL, `full_name` text NOT NULL, `ability_bonus_ability_score` integer NULL, `class_saving_throws` integer NULL, `prerequisite_ability_score` integer NULL, CONSTRAINT `ability_scores_ability_bonus_ability_score` FOREIGN KEY (`ability_bonus_ability_score`) REFERENCES `ability_bonus` (`id`) ON DELETE SET NULL, CONSTRAINT `ability_scores_classes_saving_throws` FOREIGN KEY (`class_saving_throws`) REFERENCES `classes` (`id`) ON DELETE SET NULL, CONSTRAINT `ability_scores_prerequisites_ability_score` FOREIGN KEY (`prerequisite_ability_score`) REFERENCES `prerequisites` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "ability_scores" to new temporary table "new_ability_scores"
INSERT INTO `new_ability_scores` (`id`, `indx`, `name`, `desc`, `full_name`, `ability_bonus_ability_score`, `class_saving_throws`, `prerequisite_ability_score`) SELECT `id`, `indx`, `name`, `desc`, `full_name`, `ability_bonus_ability_score`, `class_saving_throws`, `prerequisite_ability_score` FROM `ability_scores`;
-- Drop "ability_scores" table after copying rows
DROP TABLE `ability_scores`;
-- Rename temporary table "new_ability_scores" to "ability_scores"
ALTER TABLE `new_ability_scores` RENAME TO `ability_scores`;
-- Create "new_alignments" table
CREATE TABLE `new_alignments` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NULL, `abbr` text NOT NULL);
-- Copy rows from old table "alignments" to new temporary table "new_alignments"
INSERT INTO `new_alignments` (`id`, `indx`, `name`, `desc`, `abbr`) SELECT `id`, `indx`, `name`, `desc`, `abbr` FROM `alignments`;
-- Drop "alignments" table after copying rows
DROP TABLE `alignments`;
-- Rename temporary table "new_alignments" to "alignments"
ALTER TABLE `new_alignments` RENAME TO `alignments`;
-- Create "new_classes" table
CREATE TABLE `new_classes` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NULL, `hit_die` integer NOT NULL);
-- Copy rows from old table "classes" to new temporary table "new_classes"
INSERT INTO `new_classes` (`id`, `indx`, `name`, `desc`, `hit_die`) SELECT `id`, `indx`, `name`, `desc`, `hit_die` FROM `classes`;
-- Drop "classes" table after copying rows
DROP TABLE `classes`;
-- Rename temporary table "new_classes" to "classes"
ALTER TABLE `new_classes` RENAME TO `classes`;
-- Create "new_conditions" table
CREATE TABLE `new_conditions` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NULL);
-- Copy rows from old table "conditions" to new temporary table "new_conditions"
INSERT INTO `new_conditions` (`id`, `indx`, `name`, `desc`) SELECT `id`, `indx`, `name`, `desc` FROM `conditions`;
-- Drop "conditions" table after copying rows
DROP TABLE `conditions`;
-- Rename temporary table "new_conditions" to "conditions"
ALTER TABLE `new_conditions` RENAME TO `conditions`;
-- Create "new_damage_types" table
CREATE TABLE `new_damage_types` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NULL, `weapon_damage_damage_type` integer NULL, CONSTRAINT `damage_types_weapon_damages_damage_type` FOREIGN KEY (`weapon_damage_damage_type`) REFERENCES `weapon_damages` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "damage_types" to new temporary table "new_damage_types"
INSERT INTO `new_damage_types` (`id`, `indx`, `name`, `desc`, `weapon_damage_damage_type`) SELECT `id`, `indx`, `name`, `desc`, `weapon_damage_damage_type` FROM `damage_types`;
-- Drop "damage_types" table after copying rows
DROP TABLE `damage_types`;
-- Rename temporary table "new_damage_types" to "damage_types"
ALTER TABLE `new_damage_types` RENAME TO `damage_types`;
-- Create "new_equipment" table
CREATE TABLE `new_equipment` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NULL, `cost` text NOT NULL, `weight` text NOT NULL, `class_starting_equipment` integer NULL, CONSTRAINT `equipment_classes_starting_equipment` FOREIGN KEY (`class_starting_equipment`) REFERENCES `classes` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "equipment" to new temporary table "new_equipment"
INSERT INTO `new_equipment` (`id`, `indx`, `name`, `desc`, `cost`, `weight`, `class_starting_equipment`) SELECT `id`, `indx`, `name`, `desc`, `cost`, `weight`, `class_starting_equipment` FROM `equipment`;
-- Drop "equipment" table after copying rows
DROP TABLE `equipment`;
-- Rename temporary table "new_equipment" to "equipment"
ALTER TABLE `new_equipment` RENAME TO `equipment`;
-- Create "new_equipment_categories" table
CREATE TABLE `new_equipment_categories` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NULL, `equipment_subcategory` integer NULL, CONSTRAINT `equipment_categories_equipment_subcategory` FOREIGN KEY (`equipment_subcategory`) REFERENCES `equipment` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "equipment_categories" to new temporary table "new_equipment_categories"
INSERT INTO `new_equipment_categories` (`id`, `indx`, `name`, `desc`, `equipment_subcategory`) SELECT `id`, `indx`, `name`, `desc`, `equipment_subcategory` FROM `equipment_categories`;
-- Drop "equipment_categories" table after copying rows
DROP TABLE `equipment_categories`;
-- Rename temporary table "new_equipment_categories" to "equipment_categories"
ALTER TABLE `new_equipment_categories` RENAME TO `equipment_categories`;
-- Create "new_magic_schools" table
CREATE TABLE `new_magic_schools` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NULL);
-- Copy rows from old table "magic_schools" to new temporary table "new_magic_schools"
INSERT INTO `new_magic_schools` (`id`, `indx`, `name`, `desc`) SELECT `id`, `indx`, `name`, `desc` FROM `magic_schools`;
-- Drop "magic_schools" table after copying rows
DROP TABLE `magic_schools`;
-- Rename temporary table "new_magic_schools" to "magic_schools"
ALTER TABLE `new_magic_schools` RENAME TO `magic_schools`;
-- Create "new_proficiencies" table
CREATE TABLE `new_proficiencies` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NULL, `tier` text NOT NULL);
-- Copy rows from old table "proficiencies" to new temporary table "new_proficiencies"
INSERT INTO `new_proficiencies` (`id`, `indx`, `name`, `desc`, `tier`) SELECT `id`, `indx`, `name`, `desc`, `tier` FROM `proficiencies`;
-- Drop "proficiencies" table after copying rows
DROP TABLE `proficiencies`;
-- Rename temporary table "new_proficiencies" to "proficiencies"
ALTER TABLE `new_proficiencies` RENAME TO `proficiencies`;
-- Create "new_races" table
CREATE TABLE `new_races` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NULL, `speed` integer NOT NULL);
-- Copy rows from old table "races" to new temporary table "new_races"
INSERT INTO `new_races` (`id`, `indx`, `name`, `desc`, `speed`) SELECT `id`, `indx`, `name`, `desc`, `speed` FROM `races`;
-- Drop "races" table after copying rows
DROP TABLE `races`;
-- Rename temporary table "new_races" to "races"
ALTER TABLE `new_races` RENAME TO `races`;
-- Create "new_skills" table
CREATE TABLE `new_skills` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NULL, `skill_ability_score` integer NULL, CONSTRAINT `skills_ability_scores_ability_score` FOREIGN KEY (`skill_ability_score`) REFERENCES `ability_scores` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "skills" to new temporary table "new_skills"
INSERT INTO `new_skills` (`id`, `indx`, `name`, `desc`, `skill_ability_score`) SELECT `id`, `indx`, `name`, `desc`, `skill_ability_score` FROM `skills`;
-- Drop "skills" table after copying rows
DROP TABLE `skills`;
-- Rename temporary table "new_skills" to "skills"
ALTER TABLE `new_skills` RENAME TO `skills`;
-- Create "new_languages" table
CREATE TABLE `new_languages` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NULL, `category` text NOT NULL DEFAULT 'standard', `script` text NULL);
-- Copy rows from old table "languages" to new temporary table "new_languages"
INSERT INTO `new_languages` (`id`, `indx`, `name`, `desc`, `category`, `script`) SELECT `id`, `indx`, `name`, `desc`, `category`, `script` FROM `languages`;
-- Drop "languages" table after copying rows
DROP TABLE `languages`;
-- Rename temporary table "new_languages" to "languages"
ALTER TABLE `new_languages` RENAME TO `languages`;
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
