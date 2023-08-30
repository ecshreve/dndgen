-- Create "traits" table
CREATE TABLE `traits` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` json NOT NULL);
-- Set sequence for "traits" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("traits", 94489280512);
-- Create index "traits_indx_key" to table: "traits"
CREATE UNIQUE INDEX `traits_indx_key` ON `traits` (`indx`);
-- Create "race_traits" table
CREATE TABLE `race_traits` (`race_id` integer NOT NULL, `trait_id` integer NOT NULL, PRIMARY KEY (`race_id`, `trait_id`), CONSTRAINT `race_traits_race_id` FOREIGN KEY (`race_id`) REFERENCES `races` (`id`) ON DELETE CASCADE, CONSTRAINT `race_traits_trait_id` FOREIGN KEY (`trait_id`) REFERENCES `traits` (`id`) ON DELETE CASCADE);
-- Create "subrace_traits" table
CREATE TABLE `subrace_traits` (`subrace_id` integer NOT NULL, `trait_id` integer NOT NULL, PRIMARY KEY (`subrace_id`, `trait_id`), CONSTRAINT `subrace_traits_subrace_id` FOREIGN KEY (`subrace_id`) REFERENCES `subraces` (`id`) ON DELETE CASCADE, CONSTRAINT `subrace_traits_trait_id` FOREIGN KEY (`trait_id`) REFERENCES `traits` (`id`) ON DELETE CASCADE);
-- Add pk ranges for ('traits') tables
INSERT INTO `ent_types` (`type`) VALUES ('traits');
