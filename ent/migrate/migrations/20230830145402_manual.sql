-- Add column "desc" to table: "subraces"
ALTER TABLE `subraces` ADD COLUMN `desc` text NOT NULL;
-- Create "subrace_languages" table
CREATE TABLE `subrace_languages` (`subrace_id` integer NOT NULL, `language_id` integer NOT NULL, PRIMARY KEY (`subrace_id`, `language_id`), CONSTRAINT `subrace_languages_subrace_id` FOREIGN KEY (`subrace_id`) REFERENCES `subraces` (`id`) ON DELETE CASCADE, CONSTRAINT `subrace_languages_language_id` FOREIGN KEY (`language_id`) REFERENCES `languages` (`id`) ON DELETE CASCADE);
-- Create "subrace_proficiencies" table
CREATE TABLE `subrace_proficiencies` (`subrace_id` integer NOT NULL, `proficiency_id` integer NOT NULL, PRIMARY KEY (`subrace_id`, `proficiency_id`), CONSTRAINT `subrace_proficiencies_subrace_id` FOREIGN KEY (`subrace_id`) REFERENCES `subraces` (`id`) ON DELETE CASCADE, CONSTRAINT `subrace_proficiencies_proficiency_id` FOREIGN KEY (`proficiency_id`) REFERENCES `proficiencies` (`id`) ON DELETE CASCADE);
