-- Create "ability_bonus" table
CREATE TABLE `ability_bonus` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `bonus` integer NOT NULL);
-- Create "ability_scores" table
CREATE TABLE `ability_scores` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL, `full_name` text NOT NULL, `ability_bonus_ability_score` integer NULL, `class_saving_throws` integer NULL, `prerequisite_ability_score` integer NULL, CONSTRAINT `ability_scores_ability_bonus_ability_score` FOREIGN KEY (`ability_bonus_ability_score`) REFERENCES `ability_bonus` (`id`) ON DELETE SET NULL, CONSTRAINT `ability_scores_classes_saving_throws` FOREIGN KEY (`class_saving_throws`) REFERENCES `classes` (`id`) ON DELETE SET NULL, CONSTRAINT `ability_scores_prerequisites_ability_score` FOREIGN KEY (`prerequisite_ability_score`) REFERENCES `prerequisites` (`id`) ON DELETE SET NULL);
-- Create "alignments" table
CREATE TABLE `alignments` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL, `abbr` text NOT NULL);
-- Create "ammunitions" table
CREATE TABLE `ammunitions` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `quantity` integer NOT NULL);
-- Create "armors" table
CREATE TABLE `armors` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `stealth_disadvantage` bool NOT NULL, `armor_class` text NOT NULL, `min_strength` integer NOT NULL);
-- Create "classes" table
CREATE TABLE `classes` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL, `hit_die` integer NOT NULL);
-- Create "conditions" table
CREATE TABLE `conditions` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL);
-- Create "damage_types" table
CREATE TABLE `damage_types` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL, `weapon_damage_damage_type` integer NULL, CONSTRAINT `damage_types_weapon_damages_damage_type` FOREIGN KEY (`weapon_damage_damage_type`) REFERENCES `weapon_damages` (`id`) ON DELETE SET NULL);
-- Create "equipment" table
CREATE TABLE `equipment` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL, `cost` text NOT NULL, `weight` text NOT NULL, `class_starting_equipment` integer NULL, CONSTRAINT `equipment_classes_starting_equipment` FOREIGN KEY (`class_starting_equipment`) REFERENCES `classes` (`id`) ON DELETE SET NULL);
-- Create "equipment_categories" table
CREATE TABLE `equipment_categories` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL, `equipment_subcategory` integer NULL, CONSTRAINT `equipment_categories_equipment_subcategory` FOREIGN KEY (`equipment_subcategory`) REFERENCES `equipment` (`id`) ON DELETE SET NULL);
-- Create "gears" table
CREATE TABLE `gears` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT);
-- Create "languages" table
CREATE TABLE `languages` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL, `tier` text NOT NULL DEFAULT 'standard', `script` text NOT NULL);
-- Create "magic_items" table
CREATE TABLE `magic_items` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `rarity` text NOT NULL);
-- Create "magic_schools" table
CREATE TABLE `magic_schools` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL);
-- Create "packs" table
CREATE TABLE `packs` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `contents` text NOT NULL);
-- Create "prerequisites" table
CREATE TABLE `prerequisites` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `minimum` integer NOT NULL);
-- Create "proficiencies" table
CREATE TABLE `proficiencies` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL, `tier` text NOT NULL);
-- Create "races" table
CREATE TABLE `races` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL, `speed` integer NOT NULL);
-- Create "skills" table
CREATE TABLE `skills` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL, `ability_score_skills` integer NULL, CONSTRAINT `skills_ability_scores_skills` FOREIGN KEY (`ability_score_skills`) REFERENCES `ability_scores` (`id`) ON DELETE SET NULL);
-- Create "vehicles" table
CREATE TABLE `vehicles` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `speed` text NOT NULL, `capacity` text NOT NULL);
-- Create "weapons" table
CREATE TABLE `weapons` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `properties` text NOT NULL);
-- Create "weapon_damages" table
CREATE TABLE `weapon_damages` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `dice` text NOT NULL, `weapon_two_handed_damage` integer NULL, CONSTRAINT `weapon_damages_weapons_two_handed_damage` FOREIGN KEY (`weapon_two_handed_damage`) REFERENCES `weapons` (`id`) ON DELETE SET NULL);
-- Create "weapon_ranges" table
CREATE TABLE `weapon_ranges` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `desc` text NOT NULL, `normal` integer NOT NULL, `long` integer NOT NULL);
-- Create "class_starting_proficiencies" table
CREATE TABLE `class_starting_proficiencies` (`class_id` integer NOT NULL, `proficiency_id` integer NOT NULL, PRIMARY KEY (`class_id`, `proficiency_id`), CONSTRAINT `class_starting_proficiencies_class_id` FOREIGN KEY (`class_id`) REFERENCES `classes` (`id`) ON DELETE CASCADE, CONSTRAINT `class_starting_proficiencies_proficiency_id` FOREIGN KEY (`proficiency_id`) REFERENCES `proficiencies` (`id`) ON DELETE CASCADE);
-- Create "equipment_weapon" table
CREATE TABLE `equipment_weapon` (`equipment_id` integer NOT NULL, `weapon_id` integer NOT NULL, PRIMARY KEY (`equipment_id`, `weapon_id`), CONSTRAINT `equipment_weapon_equipment_id` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE CASCADE, CONSTRAINT `equipment_weapon_weapon_id` FOREIGN KEY (`weapon_id`) REFERENCES `weapons` (`id`) ON DELETE CASCADE);
-- Create "equipment_armor" table
CREATE TABLE `equipment_armor` (`equipment_id` integer NOT NULL, `armor_id` integer NOT NULL, PRIMARY KEY (`equipment_id`, `armor_id`), CONSTRAINT `equipment_armor_equipment_id` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE CASCADE, CONSTRAINT `equipment_armor_armor_id` FOREIGN KEY (`armor_id`) REFERENCES `armors` (`id`) ON DELETE CASCADE);
-- Create "equipment_gear" table
CREATE TABLE `equipment_gear` (`equipment_id` integer NOT NULL, `gear_id` integer NOT NULL, PRIMARY KEY (`equipment_id`, `gear_id`), CONSTRAINT `equipment_gear_equipment_id` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE CASCADE, CONSTRAINT `equipment_gear_gear_id` FOREIGN KEY (`gear_id`) REFERENCES `gears` (`id`) ON DELETE CASCADE);
-- Create "equipment_pack" table
CREATE TABLE `equipment_pack` (`equipment_id` integer NOT NULL, `pack_id` integer NOT NULL, PRIMARY KEY (`equipment_id`, `pack_id`), CONSTRAINT `equipment_pack_equipment_id` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE CASCADE, CONSTRAINT `equipment_pack_pack_id` FOREIGN KEY (`pack_id`) REFERENCES `packs` (`id`) ON DELETE CASCADE);
-- Create "equipment_ammunition" table
CREATE TABLE `equipment_ammunition` (`equipment_id` integer NOT NULL, `ammunition_id` integer NOT NULL, PRIMARY KEY (`equipment_id`, `ammunition_id`), CONSTRAINT `equipment_ammunition_equipment_id` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE CASCADE, CONSTRAINT `equipment_ammunition_ammunition_id` FOREIGN KEY (`ammunition_id`) REFERENCES `ammunitions` (`id`) ON DELETE CASCADE);
-- Create "equipment_vehicle" table
CREATE TABLE `equipment_vehicle` (`equipment_id` integer NOT NULL, `vehicle_id` integer NOT NULL, PRIMARY KEY (`equipment_id`, `vehicle_id`), CONSTRAINT `equipment_vehicle_equipment_id` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE CASCADE, CONSTRAINT `equipment_vehicle_vehicle_id` FOREIGN KEY (`vehicle_id`) REFERENCES `vehicles` (`id`) ON DELETE CASCADE);
-- Create "equipment_magic_item" table
CREATE TABLE `equipment_magic_item` (`equipment_id` integer NOT NULL, `magic_item_id` integer NOT NULL, PRIMARY KEY (`equipment_id`, `magic_item_id`), CONSTRAINT `equipment_magic_item_equipment_id` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE CASCADE, CONSTRAINT `equipment_magic_item_magic_item_id` FOREIGN KEY (`magic_item_id`) REFERENCES `magic_items` (`id`) ON DELETE CASCADE);
-- Create "equipment_category" table
CREATE TABLE `equipment_category` (`equipment_id` integer NOT NULL, `equipment_category_id` integer NOT NULL, PRIMARY KEY (`equipment_id`, `equipment_category_id`), CONSTRAINT `equipment_category_equipment_id` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE CASCADE, CONSTRAINT `equipment_category_equipment_category_id` FOREIGN KEY (`equipment_category_id`) REFERENCES `equipment_categories` (`id`) ON DELETE CASCADE);
-- Create "proficiency_skill" table
CREATE TABLE `proficiency_skill` (`proficiency_id` integer NOT NULL, `skill_id` integer NOT NULL, PRIMARY KEY (`proficiency_id`, `skill_id`), CONSTRAINT `proficiency_skill_proficiency_id` FOREIGN KEY (`proficiency_id`) REFERENCES `proficiencies` (`id`) ON DELETE CASCADE, CONSTRAINT `proficiency_skill_skill_id` FOREIGN KEY (`skill_id`) REFERENCES `skills` (`id`) ON DELETE CASCADE);
-- Create "proficiency_ability_score" table
CREATE TABLE `proficiency_ability_score` (`proficiency_id` integer NOT NULL, `ability_score_id` integer NOT NULL, PRIMARY KEY (`proficiency_id`, `ability_score_id`), CONSTRAINT `proficiency_ability_score_proficiency_id` FOREIGN KEY (`proficiency_id`) REFERENCES `proficiencies` (`id`) ON DELETE CASCADE, CONSTRAINT `proficiency_ability_score_ability_score_id` FOREIGN KEY (`ability_score_id`) REFERENCES `ability_scores` (`id`) ON DELETE CASCADE);
-- Create "proficiency_equipment" table
CREATE TABLE `proficiency_equipment` (`proficiency_id` integer NOT NULL, `equipment_id` integer NOT NULL, PRIMARY KEY (`proficiency_id`, `equipment_id`), CONSTRAINT `proficiency_equipment_proficiency_id` FOREIGN KEY (`proficiency_id`) REFERENCES `proficiencies` (`id`) ON DELETE CASCADE, CONSTRAINT `proficiency_equipment_equipment_id` FOREIGN KEY (`equipment_id`) REFERENCES `equipment` (`id`) ON DELETE CASCADE);
-- Create "race_languages" table
CREATE TABLE `race_languages` (`race_id` integer NOT NULL, `language_id` integer NOT NULL, PRIMARY KEY (`race_id`, `language_id`), CONSTRAINT `race_languages_race_id` FOREIGN KEY (`race_id`) REFERENCES `races` (`id`) ON DELETE CASCADE, CONSTRAINT `race_languages_language_id` FOREIGN KEY (`language_id`) REFERENCES `languages` (`id`) ON DELETE CASCADE);
-- Create "race_ability_bonuses" table
CREATE TABLE `race_ability_bonuses` (`race_id` integer NOT NULL, `ability_bonus_id` integer NOT NULL, PRIMARY KEY (`race_id`, `ability_bonus_id`), CONSTRAINT `race_ability_bonuses_race_id` FOREIGN KEY (`race_id`) REFERENCES `races` (`id`) ON DELETE CASCADE, CONSTRAINT `race_ability_bonuses_ability_bonus_id` FOREIGN KEY (`ability_bonus_id`) REFERENCES `ability_bonus` (`id`) ON DELETE CASCADE);
-- Create "race_starting_proficiencies" table
CREATE TABLE `race_starting_proficiencies` (`race_id` integer NOT NULL, `proficiency_id` integer NOT NULL, PRIMARY KEY (`race_id`, `proficiency_id`), CONSTRAINT `race_starting_proficiencies_race_id` FOREIGN KEY (`race_id`) REFERENCES `races` (`id`) ON DELETE CASCADE, CONSTRAINT `race_starting_proficiencies_proficiency_id` FOREIGN KEY (`proficiency_id`) REFERENCES `proficiencies` (`id`) ON DELETE CASCADE);
-- Create "weapon_range" table
CREATE TABLE `weapon_range` (`weapon_id` integer NOT NULL, `weapon_range_id` integer NOT NULL, PRIMARY KEY (`weapon_id`, `weapon_range_id`), CONSTRAINT `weapon_range_weapon_id` FOREIGN KEY (`weapon_id`) REFERENCES `weapons` (`id`) ON DELETE CASCADE, CONSTRAINT `weapon_range_weapon_range_id` FOREIGN KEY (`weapon_range_id`) REFERENCES `weapon_ranges` (`id`) ON DELETE CASCADE);
-- Create "weapon_damage" table
CREATE TABLE `weapon_damage` (`weapon_id` integer NOT NULL, `weapon_damage_id` integer NOT NULL, PRIMARY KEY (`weapon_id`, `weapon_damage_id`), CONSTRAINT `weapon_damage_weapon_id` FOREIGN KEY (`weapon_id`) REFERENCES `weapons` (`id`) ON DELETE CASCADE, CONSTRAINT `weapon_damage_weapon_damage_id` FOREIGN KEY (`weapon_damage_id`) REFERENCES `weapon_damages` (`id`) ON DELETE CASCADE);
