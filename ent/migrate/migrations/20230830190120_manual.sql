-- Create "choices" table
CREATE TABLE `choices` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `choose` integer NOT NULL, `desc` text NOT NULL);
-- Set sequence for "choices" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("choices", 98784247808);
-- Create "choice_proficiencies" table
CREATE TABLE `choice_proficiencies` (`choice_id` integer NOT NULL, `proficiency_id` integer NOT NULL, PRIMARY KEY (`choice_id`, `proficiency_id`), CONSTRAINT `choice_proficiencies_choice_id` FOREIGN KEY (`choice_id`) REFERENCES `choices` (`id`) ON DELETE CASCADE, CONSTRAINT `choice_proficiencies_proficiency_id` FOREIGN KEY (`proficiency_id`) REFERENCES `proficiencies` (`id`) ON DELETE CASCADE);
-- Create "choice_languages" table
CREATE TABLE `choice_languages` (`choice_id` integer NOT NULL, `language_id` integer NOT NULL, PRIMARY KEY (`choice_id`, `language_id`), CONSTRAINT `choice_languages_choice_id` FOREIGN KEY (`choice_id`) REFERENCES `choices` (`id`) ON DELETE CASCADE, CONSTRAINT `choice_languages_language_id` FOREIGN KEY (`language_id`) REFERENCES `languages` (`id`) ON DELETE CASCADE);
-- Create "race_starting_proficiency_options" table
CREATE TABLE `race_starting_proficiency_options` (`race_id` integer NOT NULL, `choice_id` integer NOT NULL, PRIMARY KEY (`race_id`, `choice_id`), CONSTRAINT `race_starting_proficiency_options_race_id` FOREIGN KEY (`race_id`) REFERENCES `races` (`id`) ON DELETE CASCADE, CONSTRAINT `race_starting_proficiency_options_choice_id` FOREIGN KEY (`choice_id`) REFERENCES `choices` (`id`) ON DELETE CASCADE);
-- Add pk ranges for ('choices') tables
INSERT INTO `ent_types` (`type`) VALUES ('choices');
