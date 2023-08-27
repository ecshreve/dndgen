-- Add column "weapon_category" to table: "weapons"
ALTER TABLE `weapons` ADD COLUMN `weapon_category` text NOT NULL;
-- Create "weapon_properties" table
CREATE TABLE `weapon_properties` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` json NOT NULL);
-- Set sequence for "weapon_properties" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("weapon_properties", 64424509440);
-- Create index "weapon_properties_indx_key" to table: "weapon_properties"
CREATE UNIQUE INDEX `weapon_properties_indx_key` ON `weapon_properties` (`indx`);
-- Create "weapon_weapon_properties" table
CREATE TABLE `weapon_weapon_properties` (`weapon_id` integer NOT NULL, `weapon_property_id` integer NOT NULL, PRIMARY KEY (`weapon_id`, `weapon_property_id`), CONSTRAINT `weapon_weapon_properties_weapon_id` FOREIGN KEY (`weapon_id`) REFERENCES `weapons` (`id`) ON DELETE CASCADE, CONSTRAINT `weapon_weapon_properties_weapon_property_id` FOREIGN KEY (`weapon_property_id`) REFERENCES `weapon_properties` (`id`) ON DELETE CASCADE);
-- Add pk ranges for ('weapon_properties') tables
INSERT INTO `ent_types` (`type`) VALUES ('weapon_properties');
