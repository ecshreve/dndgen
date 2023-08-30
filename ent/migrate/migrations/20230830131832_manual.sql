-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_weapon_damages" table
CREATE TABLE `new_weapon_damages` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `dice` text NOT NULL, `weapon_id` integer NOT NULL, `damage_type_id` integer NOT NULL, CONSTRAINT `weapon_damages_weapons_weapon_damage` FOREIGN KEY (`weapon_id`) REFERENCES `weapons` (`id`) ON DELETE NO ACTION, CONSTRAINT `weapon_damages_damage_types_damage_type` FOREIGN KEY (`damage_type_id`) REFERENCES `damage_types` (`id`) ON DELETE NO ACTION);
-- Set sequence for "new_weapon_damages" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("new_weapon_damages", 81604378624);
-- Copy rows from old table "weapon_damages" to new temporary table "new_weapon_damages"
INSERT INTO `new_weapon_damages` (`dice`, `weapon_id`, `damage_type_id`) SELECT `dice`, `weapon_id`, `damage_type_id` FROM `weapon_damages`;
-- Drop "weapon_damages" table after copying rows
DROP TABLE `weapon_damages`;
-- Rename temporary table "new_weapon_damages" to "weapon_damages"
ALTER TABLE `new_weapon_damages` RENAME TO `weapon_damages`;
-- Create "magic_schools" table
CREATE TABLE `magic_schools` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL);
-- Set sequence for "magic_schools" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("magic_schools", 68719476736);
-- Create index "magic_schools_indx_key" to table: "magic_schools"
CREATE UNIQUE INDEX `magic_schools_indx_key` ON `magic_schools` (`indx`);
-- Create "rules" table
CREATE TABLE `rules` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL);
-- Set sequence for "rules" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("rules", 73014444032);
-- Create index "rules_indx_key" to table: "rules"
CREATE UNIQUE INDEX `rules_indx_key` ON `rules` (`indx`);
-- Create "rule_sections" table
CREATE TABLE `rule_sections` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL);
-- Set sequence for "rule_sections" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("rule_sections", 77309411328);
-- Create index "rule_sections_indx_key" to table: "rule_sections"
CREATE UNIQUE INDEX `rule_sections_indx_key` ON `rule_sections` (`indx`);
-- Create "rule_rule_sections" table
CREATE TABLE `rule_rule_sections` (`rule_id` integer NOT NULL, `rule_section_id` integer NOT NULL, PRIMARY KEY (`rule_id`, `rule_section_id`), CONSTRAINT `rule_rule_sections_rule_id` FOREIGN KEY (`rule_id`) REFERENCES `rules` (`id`) ON DELETE CASCADE, CONSTRAINT `rule_rule_sections_rule_section_id` FOREIGN KEY (`rule_section_id`) REFERENCES `rule_sections` (`id`) ON DELETE CASCADE);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
-- Add pk ranges for ('magic_schools'),('rules'),('rule_sections'),('weapon_damages') tables
INSERT INTO `ent_types` (`type`) VALUES ('magic_schools'), ('rules'), ('rule_sections'), ('weapon_damages');
