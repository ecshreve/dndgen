-- Create "ability_scores" table
CREATE TABLE `ability_scores` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `full_name` text NOT NULL, `desc` json NOT NULL);
-- Create index "ability_scores_indx_key" to table: "ability_scores"
CREATE UNIQUE INDEX `ability_scores_indx_key` ON `ability_scores` (`indx`);
-- Create "armors" table
CREATE TABLE `armors` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `stealth_disadvantage` bool NOT NULL, `min_strength` integer NOT NULL);
-- Create index "armors_indx_key" to table: "armors"
CREATE UNIQUE INDEX `armors_indx_key` ON `armors` (`indx`);
-- Create "armor_classes" table
CREATE TABLE `armor_classes` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `base` integer NOT NULL, `dex_bonus` bool NOT NULL, `max_bonus` integer NULL, `armor_armor_class` integer NULL, CONSTRAINT `armor_classes_armors_armor_class` FOREIGN KEY (`armor_armor_class`) REFERENCES `armors` (`id`) ON DELETE SET NULL);
-- Create "classes" table
CREATE TABLE `classes` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `hit_die` integer NOT NULL);
-- Create index "classes_indx_key" to table: "classes"
CREATE UNIQUE INDEX `classes_indx_key` ON `classes` (`indx`);
-- Create "damage_types" table
CREATE TABLE `damage_types` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` json NOT NULL, `weapon_damage_damage_type` integer NULL, CONSTRAINT `damage_types_weapon_damages_damage_type` FOREIGN KEY (`weapon_damage_damage_type`) REFERENCES `weapon_damages` (`id`) ON DELETE SET NULL);
-- Create index "damage_types_indx_key" to table: "damage_types"
CREATE UNIQUE INDEX `damage_types_indx_key` ON `damage_types` (`indx`);
-- Create "equipment" table
CREATE TABLE `equipment` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `equipment_category` text NOT NULL DEFAULT 'other', `equipment_weapon` integer NULL, `equipment_armor` integer NULL, `equipment_gear` integer NULL, `equipment_tool` integer NULL, `equipment_vehicle` integer NULL, CONSTRAINT `equipment_weapons_weapon` FOREIGN KEY (`equipment_weapon`) REFERENCES `weapons` (`id`) ON DELETE SET NULL, CONSTRAINT `equipment_armors_armor` FOREIGN KEY (`equipment_armor`) REFERENCES `armors` (`id`) ON DELETE SET NULL, CONSTRAINT `equipment_gears_gear` FOREIGN KEY (`equipment_gear`) REFERENCES `gears` (`id`) ON DELETE SET NULL, CONSTRAINT `equipment_tools_tool` FOREIGN KEY (`equipment_tool`) REFERENCES `tools` (`id`) ON DELETE SET NULL, CONSTRAINT `equipment_vehicles_vehicle` FOREIGN KEY (`equipment_vehicle`) REFERENCES `vehicles` (`id`) ON DELETE SET NULL);
-- Create index "equipment_indx_key" to table: "equipment"
CREATE UNIQUE INDEX `equipment_indx_key` ON `equipment` (`indx`);
-- Create "gears" table
CREATE TABLE `gears` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `gear_category` text NOT NULL DEFAULT 'other', `desc` json NOT NULL, `quantity` integer NULL);
-- Create index "gears_indx_key" to table: "gears"
CREATE UNIQUE INDEX `gears_indx_key` ON `gears` (`indx`);
-- Create "races" table
CREATE TABLE `races` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `speed` integer NOT NULL);
-- Create index "races_indx_key" to table: "races"
CREATE UNIQUE INDEX `races_indx_key` ON `races` (`indx`);
-- Create "skills" table
CREATE TABLE `skills` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` json NOT NULL, `skill_ability_score` integer NULL, CONSTRAINT `skills_ability_scores_ability_score` FOREIGN KEY (`skill_ability_score`) REFERENCES `ability_scores` (`id`) ON DELETE SET NULL);
-- Create index "skills_indx_key" to table: "skills"
CREATE UNIQUE INDEX `skills_indx_key` ON `skills` (`indx`);
-- Create "tools" table
CREATE TABLE `tools` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `tool_category` text NOT NULL);
-- Create index "tools_indx_key" to table: "tools"
CREATE UNIQUE INDEX `tools_indx_key` ON `tools` (`indx`);
-- Create "vehicles" table
CREATE TABLE `vehicles` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `vehicle_category` text NOT NULL, `capacity` text NOT NULL);
-- Create index "vehicles_indx_key" to table: "vehicles"
CREATE UNIQUE INDEX `vehicles_indx_key` ON `vehicles` (`indx`);
-- Create "weapons" table
CREATE TABLE `weapons` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `weapon_range` text NOT NULL);
-- Create index "weapons_indx_key" to table: "weapons"
CREATE UNIQUE INDEX `weapons_indx_key` ON `weapons` (`indx`);
-- Create "weapon_damages" table
CREATE TABLE `weapon_damages` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `dice` text NOT NULL);
-- Create "class_saving_throws" table
CREATE TABLE `class_saving_throws` (`class_id` integer NOT NULL, `ability_score_id` integer NOT NULL, PRIMARY KEY (`class_id`, `ability_score_id`), CONSTRAINT `class_saving_throws_class_id` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`) ON DELETE CASCADE, CONSTRAINT `class_saving_throws_ability_score_id` FOREIGN KEY (`ability_score_id`) REFERENCES `ability_scores` (`id`) ON DELETE CASCADE);
