-- Create "ability_bonus" table
CREATE TABLE `ability_bonus` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `bonus` integer NOT NULL, `ability_score_id` integer NOT NULL, `race_ability_bonuses` integer NULL, `subrace_ability_bonuses` integer NULL, CONSTRAINT `ability_bonus_ability_scores_ability_bonuses` FOREIGN KEY (`ability_score_id`) REFERENCES `ability_scores` (`id`) ON DELETE NO ACTION, CONSTRAINT `ability_bonus_races_ability_bonuses` FOREIGN KEY (`race_ability_bonuses`) REFERENCES `races` (`id`) ON DELETE SET NULL, CONSTRAINT `ability_bonus_subraces_ability_bonuses` FOREIGN KEY (`subrace_ability_bonuses`) REFERENCES `subraces` (`id`) ON DELETE SET NULL);
-- Create "ability_scores" table
CREATE TABLE `ability_scores` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `full_name` text NOT NULL, `desc` json NOT NULL);
-- Set sequence for "ability_scores" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("ability_scores", 4294967296);
-- Create index "ability_scores_indx_key" to table: "ability_scores"
CREATE UNIQUE INDEX `ability_scores_indx_key` ON `ability_scores` (`indx`);
-- Create "armors" table
CREATE TABLE `armors` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `armor_category` text NOT NULL, `stealth_disadvantage` bool NOT NULL, `min_strength` integer NOT NULL, `equipment_id` integer NOT NULL, CONSTRAINT `armors_equipment_armor` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE NO ACTION);
-- Set sequence for "armors" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("armors", 8589934592);
-- Create index "armors_indx_key" to table: "armors"
CREATE UNIQUE INDEX `armors_indx_key` ON `armors` (`indx`);
-- Create index "armors_equipment_id_key" to table: "armors"
CREATE UNIQUE INDEX `armors_equipment_id_key` ON `armors` (`equipment_id`);
-- Create "armor_classes" table
CREATE TABLE `armor_classes` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `base` integer NOT NULL, `dex_bonus` bool NOT NULL, `max_bonus` integer NULL, `armor_armor_class` integer NULL, CONSTRAINT `armor_classes_armors_armor_class` FOREIGN KEY (`armor_armor_class`) REFERENCES `armors` (`id`) ON DELETE SET NULL);
-- Set sequence for "armor_classes" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("armor_classes", 12884901888);
-- Create "classes" table
CREATE TABLE `classes` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `hit_die` integer NOT NULL);
-- Set sequence for "classes" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("classes", 17179869184);
-- Create index "classes_indx_key" to table: "classes"
CREATE UNIQUE INDEX `classes_indx_key` ON `classes` (`indx`);
-- Create "costs" table
CREATE TABLE `costs` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `quantity` integer NOT NULL, `unit` text NOT NULL);
-- Set sequence for "costs" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("costs", 21474836480);
-- Create "damage_types" table
CREATE TABLE `damage_types` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` json NOT NULL);
-- Set sequence for "damage_types" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("damage_types", 25769803776);
-- Create index "damage_types_indx_key" to table: "damage_types"
CREATE UNIQUE INDEX `damage_types_indx_key` ON `damage_types` (`indx`);
-- Create "equipment" table
CREATE TABLE `equipment` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `equipment_category` text NOT NULL DEFAULT 'other', `equipment_cost` integer NULL, CONSTRAINT `equipment_costs_cost` FOREIGN KEY (`equipment_cost`) REFERENCES `costs` (`id`) ON DELETE SET NULL);
-- Set sequence for "equipment" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("equipment", 30064771072);
-- Create index "equipment_indx_key" to table: "equipment"
CREATE UNIQUE INDEX `equipment_indx_key` ON `equipment` (`indx`);
-- Create "gears" table
CREATE TABLE `gears` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `gear_category` text NOT NULL DEFAULT 'other', `desc` json NOT NULL, `quantity` integer NULL, `equipment_id` integer NOT NULL, CONSTRAINT `gears_equipment_gear` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE NO ACTION);
-- Set sequence for "gears" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("gears", 34359738368);
-- Create index "gears_indx_key" to table: "gears"
CREATE UNIQUE INDEX `gears_indx_key` ON `gears` (`indx`);
-- Create index "gears_equipment_id_key" to table: "gears"
CREATE UNIQUE INDEX `gears_equipment_id_key` ON `gears` (`equipment_id`);
-- Create "languages" table
CREATE TABLE `languages` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL, `language_type` text NOT NULL DEFAULT 'STANDARD', `script` text NULL DEFAULT 'Common');
-- Set sequence for "languages" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("languages", 38654705664);
-- Create index "languages_indx_key" to table: "languages"
CREATE UNIQUE INDEX `languages_indx_key` ON `languages` (`indx`);
-- Create "magic_schools" table
CREATE TABLE `magic_schools` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL);
-- Set sequence for "magic_schools" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("magic_schools", 42949672960);
-- Create index "magic_schools_indx_key" to table: "magic_schools"
CREATE UNIQUE INDEX `magic_schools_indx_key` ON `magic_schools` (`indx`);
-- Create "proficiencies" table
CREATE TABLE `proficiencies` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `proficiency_category` text NOT NULL, `proficiency_skill` integer NULL, `proficiency_equipment` integer NULL, `proficiency_saving_throw` integer NULL, CONSTRAINT `proficiencies_skills_skill` FOREIGN KEY (`proficiency_skill`) REFERENCES `skills` (`id`) ON DELETE SET NULL, CONSTRAINT `proficiencies_equipment_equipment` FOREIGN KEY (`proficiency_equipment`) REFERENCES `equipment` (`id`) ON DELETE SET NULL, CONSTRAINT `proficiencies_ability_scores_saving_throw` FOREIGN KEY (`proficiency_saving_throw`) REFERENCES `ability_scores` (`id`) ON DELETE SET NULL);
-- Set sequence for "proficiencies" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("proficiencies", 47244640256);
-- Create index "proficiencies_indx_key" to table: "proficiencies"
CREATE UNIQUE INDEX `proficiencies_indx_key` ON `proficiencies` (`indx`);
-- Create "proficiency_choices" table
CREATE TABLE `proficiency_choices` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `desc` text NULL, `choose` integer NOT NULL);
-- Set sequence for "proficiency_choices" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("proficiency_choices", 51539607552);
-- Create "races" table
CREATE TABLE `races` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `alignment` text NOT NULL, `age` text NOT NULL, `size` text NOT NULL, `size_description` text NOT NULL, `language_desc` text NOT NULL, `speed` integer NOT NULL, `race_starting_proficiency_option` integer NULL, CONSTRAINT `races_proficiency_choices_starting_proficiency_option` FOREIGN KEY (`race_starting_proficiency_option`) REFERENCES `proficiency_choices` (`id`) ON DELETE SET NULL);
-- Set sequence for "races" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("races", 55834574848);
-- Create index "races_indx_key" to table: "races"
CREATE UNIQUE INDEX `races_indx_key` ON `races` (`indx`);
-- Create "rules" table
CREATE TABLE `rules` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL);
-- Set sequence for "rules" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("rules", 60129542144);
-- Create index "rules_indx_key" to table: "rules"
CREATE UNIQUE INDEX `rules_indx_key` ON `rules` (`indx`);
-- Create "rule_sections" table
CREATE TABLE `rule_sections` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL);
-- Set sequence for "rule_sections" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("rule_sections", 64424509440);
-- Create index "rule_sections_indx_key" to table: "rule_sections"
CREATE UNIQUE INDEX `rule_sections_indx_key` ON `rule_sections` (`indx`);
-- Create "skills" table
CREATE TABLE `skills` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` json NOT NULL, `skill_ability_score` integer NULL, CONSTRAINT `skills_ability_scores_ability_score` FOREIGN KEY (`skill_ability_score`) REFERENCES `ability_scores` (`id`) ON DELETE SET NULL);
-- Set sequence for "skills" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("skills", 68719476736);
-- Create index "skills_indx_key" to table: "skills"
CREATE UNIQUE INDEX `skills_indx_key` ON `skills` (`indx`);
-- Create "subraces" table
CREATE TABLE `subraces` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL, `race_subraces` integer NULL, CONSTRAINT `subraces_races_subraces` FOREIGN KEY (`race_subraces`) REFERENCES `races` (`id`) ON DELETE SET NULL);
-- Set sequence for "subraces" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("subraces", 73014444032);
-- Create index "subraces_indx_key" to table: "subraces"
CREATE UNIQUE INDEX `subraces_indx_key` ON `subraces` (`indx`);
-- Create "tools" table
CREATE TABLE `tools` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `tool_category` text NOT NULL, `equipment_id` integer NOT NULL, CONSTRAINT `tools_equipment_tool` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE NO ACTION);
-- Set sequence for "tools" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("tools", 77309411328);
-- Create index "tools_indx_key" to table: "tools"
CREATE UNIQUE INDEX `tools_indx_key` ON `tools` (`indx`);
-- Create index "tools_equipment_id_key" to table: "tools"
CREATE UNIQUE INDEX `tools_equipment_id_key` ON `tools` (`equipment_id`);
-- Create "traits" table
CREATE TABLE `traits` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` json NOT NULL);
-- Set sequence for "traits" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("traits", 81604378624);
-- Create index "traits_indx_key" to table: "traits"
CREATE UNIQUE INDEX `traits_indx_key` ON `traits` (`indx`);
-- Create "vehicles" table
CREATE TABLE `vehicles` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `vehicle_category` text NOT NULL, `capacity` text NOT NULL, `equipment_id` integer NOT NULL, CONSTRAINT `vehicles_equipment_vehicle` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE NO ACTION);
-- Set sequence for "vehicles" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("vehicles", 85899345920);
-- Create index "vehicles_indx_key" to table: "vehicles"
CREATE UNIQUE INDEX `vehicles_indx_key` ON `vehicles` (`indx`);
-- Create index "vehicles_equipment_id_key" to table: "vehicles"
CREATE UNIQUE INDEX `vehicles_equipment_id_key` ON `vehicles` (`equipment_id`);
-- Create "weapons" table
CREATE TABLE `weapons` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `weapon_category` text NOT NULL, `weapon_range` text NOT NULL, `equipment_id` integer NOT NULL, CONSTRAINT `weapons_equipment_weapon` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE NO ACTION);
-- Set sequence for "weapons" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("weapons", 90194313216);
-- Create index "weapons_indx_key" to table: "weapons"
CREATE UNIQUE INDEX `weapons_indx_key` ON `weapons` (`indx`);
-- Create index "weapons_equipment_id_key" to table: "weapons"
CREATE UNIQUE INDEX `weapons_equipment_id_key` ON `weapons` (`equipment_id`);
-- Create "weapon_damages" table
CREATE TABLE `weapon_damages` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `dice` text NOT NULL, `weapon_id` integer NOT NULL, `damage_type_id` integer NOT NULL, CONSTRAINT `weapon_damages_weapons_weapon_damage` FOREIGN KEY (`weapon_id`) REFERENCES `weapons` (`id`) ON DELETE NO ACTION, CONSTRAINT `weapon_damages_damage_types_damage_type` FOREIGN KEY (`damage_type_id`) REFERENCES `damage_types` (`id`) ON DELETE NO ACTION);
-- Set sequence for "weapon_damages" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("weapon_damages", 94489280512);
-- Create "weapon_properties" table
CREATE TABLE `weapon_properties` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` json NOT NULL);
-- Set sequence for "weapon_properties" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("weapon_properties", 98784247808);
-- Create index "weapon_properties_indx_key" to table: "weapon_properties"
CREATE UNIQUE INDEX `weapon_properties_indx_key` ON `weapon_properties` (`indx`);
-- Create "class_proficiencies" table
CREATE TABLE `class_proficiencies` (`class_id` integer NOT NULL, `proficiency_id` integer NOT NULL, PRIMARY KEY (`class_id`, `proficiency_id`), CONSTRAINT `class_proficiencies_class_id` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`) ON DELETE CASCADE, CONSTRAINT `class_proficiencies_proficiency_id` FOREIGN KEY (`proficiency_id`) REFERENCES `proficiencies` (`id`) ON DELETE CASCADE);
-- Create "class_proficiency_choices" table
CREATE TABLE `class_proficiency_choices` (`class_id` integer NOT NULL, `proficiency_choice_id` integer NOT NULL, PRIMARY KEY (`class_id`, `proficiency_choice_id`), CONSTRAINT `class_proficiency_choices_class_id` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`) ON DELETE CASCADE, CONSTRAINT `class_proficiency_choices_proficiency_choice_id` FOREIGN KEY (`proficiency_choice_id`) REFERENCES `proficiency_choices` (`id`) ON DELETE CASCADE);
-- Create "proficiency_choice_options" table
CREATE TABLE `proficiency_choice_options` (`proficiency_choice_id` integer NOT NULL, `proficiency_id` integer NOT NULL, PRIMARY KEY (`proficiency_choice_id`, `proficiency_id`), CONSTRAINT `proficiency_choice_options_proficiency_choice_id` FOREIGN KEY (`proficiency_choice_id`) REFERENCES `proficiency_choices` (`id`) ON DELETE CASCADE, CONSTRAINT `proficiency_choice_options_proficiency_id` FOREIGN KEY (`proficiency_id`) REFERENCES `proficiencies` (`id`) ON DELETE CASCADE);
-- Create "race_languages" table
CREATE TABLE `race_languages` (`race_id` integer NOT NULL, `language_id` integer NOT NULL, PRIMARY KEY (`race_id`, `language_id`), CONSTRAINT `race_languages_race_id` FOREIGN KEY (`race_id`) REFERENCES `races` (`id`) ON DELETE CASCADE, CONSTRAINT `race_languages_language_id` FOREIGN KEY (`language_id`) REFERENCES `languages` (`id`) ON DELETE CASCADE);
-- Create "race_proficiencies" table
CREATE TABLE `race_proficiencies` (`race_id` integer NOT NULL, `proficiency_id` integer NOT NULL, PRIMARY KEY (`race_id`, `proficiency_id`), CONSTRAINT `race_proficiencies_race_id` FOREIGN KEY (`race_id`) REFERENCES `races` (`id`) ON DELETE CASCADE, CONSTRAINT `race_proficiencies_proficiency_id` FOREIGN KEY (`proficiency_id`) REFERENCES `proficiencies` (`id`) ON DELETE CASCADE);
-- Create "race_traits" table
CREATE TABLE `race_traits` (`race_id` integer NOT NULL, `trait_id` integer NOT NULL, PRIMARY KEY (`race_id`, `trait_id`), CONSTRAINT `race_traits_race_id` FOREIGN KEY (`race_id`) REFERENCES `races` (`id`) ON DELETE CASCADE, CONSTRAINT `race_traits_trait_id` FOREIGN KEY (`trait_id`) REFERENCES `traits` (`id`) ON DELETE CASCADE);
-- Create "rule_rule_sections" table
CREATE TABLE `rule_rule_sections` (`rule_id` integer NOT NULL, `rule_section_id` integer NOT NULL, PRIMARY KEY (`rule_id`, `rule_section_id`), CONSTRAINT `rule_rule_sections_rule_id` FOREIGN KEY (`rule_id`) REFERENCES `rules` (`id`) ON DELETE CASCADE, CONSTRAINT `rule_rule_sections_rule_section_id` FOREIGN KEY (`rule_section_id`) REFERENCES `rule_sections` (`id`) ON DELETE CASCADE);
-- Create "subrace_proficiencies" table
CREATE TABLE `subrace_proficiencies` (`subrace_id` integer NOT NULL, `proficiency_id` integer NOT NULL, PRIMARY KEY (`subrace_id`, `proficiency_id`), CONSTRAINT `subrace_proficiencies_subrace_id` FOREIGN KEY (`subrace_id`) REFERENCES `subraces` (`id`) ON DELETE CASCADE, CONSTRAINT `subrace_proficiencies_proficiency_id` FOREIGN KEY (`proficiency_id`) REFERENCES `proficiencies` (`id`) ON DELETE CASCADE);
-- Create "subrace_traits" table
CREATE TABLE `subrace_traits` (`subrace_id` integer NOT NULL, `trait_id` integer NOT NULL, PRIMARY KEY (`subrace_id`, `trait_id`), CONSTRAINT `subrace_traits_subrace_id` FOREIGN KEY (`subrace_id`) REFERENCES `subraces` (`id`) ON DELETE CASCADE, CONSTRAINT `subrace_traits_trait_id` FOREIGN KEY (`trait_id`) REFERENCES `traits` (`id`) ON DELETE CASCADE);
-- Create "weapon_weapon_properties" table
CREATE TABLE `weapon_weapon_properties` (`weapon_id` integer NOT NULL, `weapon_property_id` integer NOT NULL, PRIMARY KEY (`weapon_id`, `weapon_property_id`), CONSTRAINT `weapon_weapon_properties_weapon_id` FOREIGN KEY (`weapon_id`) REFERENCES `weapons` (`id`) ON DELETE CASCADE, CONSTRAINT `weapon_weapon_properties_weapon_property_id` FOREIGN KEY (`weapon_property_id`) REFERENCES `weapon_properties` (`id`) ON DELETE CASCADE);
-- Create "ent_types" table
CREATE TABLE `ent_types` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `type` text NOT NULL);
-- Create index "ent_types_type_key" to table: "ent_types"
CREATE UNIQUE INDEX `ent_types_type_key` ON `ent_types` (`type`);
-- Add pk ranges for ('ability_bonus'),('ability_scores'),('armors'),('armor_classes'),('classes'),('costs'),('damage_types'),('equipment'),('gears'),('languages'),('magic_schools'),('proficiencies'),('proficiency_choices'),('races'),('rules'),('rule_sections'),('skills'),('subraces'),('tools'),('traits'),('vehicles'),('weapons'),('weapon_damages'),('weapon_properties') tables
INSERT INTO `ent_types` (`type`) VALUES ('ability_bonus'), ('ability_scores'), ('armors'), ('armor_classes'), ('classes'), ('costs'), ('damage_types'), ('equipment'), ('gears'), ('languages'), ('magic_schools'), ('proficiencies'), ('proficiency_choices'), ('races'), ('rules'), ('rule_sections'), ('skills'), ('subraces'), ('tools'), ('traits'), ('vehicles'), ('weapons'), ('weapon_damages'), ('weapon_properties');
