-- Create "ability_scores" table
CREATE TABLE `ability_scores` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `abbr` text NOT NULL, `desc` json NOT NULL);
-- Create index "ability_scores_indx_key" to table: "ability_scores"
CREATE UNIQUE INDEX `ability_scores_indx_key` ON `ability_scores` (`indx`);
-- Create index "ability_scores_name_key" to table: "ability_scores"
CREATE UNIQUE INDEX `ability_scores_name_key` ON `ability_scores` (`name`);
-- Create "alignments" table
CREATE TABLE `alignments` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL, `abbr` text NOT NULL);
-- Set sequence for "alignments" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("alignments", 4294967296);
-- Create index "alignments_indx_key" to table: "alignments"
CREATE UNIQUE INDEX `alignments_indx_key` ON `alignments` (`indx`);
-- Create index "alignments_name_key" to table: "alignments"
CREATE UNIQUE INDEX `alignments_name_key` ON `alignments` (`name`);
-- Create "characters" table
CREATE TABLE `characters` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `name` text NOT NULL, `age` integer NOT NULL);
-- Set sequence for "characters" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("characters", 8589934592);
-- Create index "characters_name_key" to table: "characters"
CREATE UNIQUE INDEX `characters_name_key` ON `characters` (`name`);
-- Create "classes" table
CREATE TABLE `classes` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `hit_die` integer NOT NULL);
-- Set sequence for "classes" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("classes", 12884901888);
-- Create index "classes_indx_key" to table: "classes"
CREATE UNIQUE INDEX `classes_indx_key` ON `classes` (`indx`);
-- Create index "classes_name_key" to table: "classes"
CREATE UNIQUE INDEX `classes_name_key` ON `classes` (`name`);
-- Create "languages" table
CREATE TABLE `languages` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `type` text NOT NULL, `script` text NOT NULL);
-- Set sequence for "languages" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("languages", 17179869184);
-- Create index "languages_indx_key" to table: "languages"
CREATE UNIQUE INDEX `languages_indx_key` ON `languages` (`indx`);
-- Create index "languages_name_key" to table: "languages"
CREATE UNIQUE INDEX `languages_name_key` ON `languages` (`name`);
-- Create "magic_schools" table
CREATE TABLE `magic_schools` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `desc` text NOT NULL);
-- Set sequence for "magic_schools" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("magic_schools", 21474836480);
-- Create index "magic_schools_indx_key" to table: "magic_schools"
CREATE UNIQUE INDEX `magic_schools_indx_key` ON `magic_schools` (`indx`);
-- Create index "magic_schools_name_key" to table: "magic_schools"
CREATE UNIQUE INDEX `magic_schools_name_key` ON `magic_schools` (`name`);
-- Create "races" table
CREATE TABLE `races` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL, `speed` integer NOT NULL, `size` text NOT NULL, `size_description` text NOT NULL, `age` text NOT NULL);
-- Set sequence for "races" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("races", 25769803776);
-- Create index "races_indx_key" to table: "races"
CREATE UNIQUE INDEX `races_indx_key` ON `races` (`indx`);
-- Create index "races_name_key" to table: "races"
CREATE UNIQUE INDEX `races_name_key` ON `races` (`name`);
-- Create "skills" table
CREATE TABLE `skills` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `indx` text NOT NULL, `name` text NOT NULL);
-- Set sequence for "skills" table
INSERT INTO sqlite_sequence (name, seq) VALUES ("skills", 30064771072);
-- Create index "skills_indx_key" to table: "skills"
CREATE UNIQUE INDEX `skills_indx_key` ON `skills` (`indx`);
-- Create index "skills_name_key" to table: "skills"
CREATE UNIQUE INDEX `skills_name_key` ON `skills` (`name`);
-- Create "ent_types" table
CREATE TABLE `ent_types` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `type` text NOT NULL);
-- Create index "ent_types_type_key" to table: "ent_types"
CREATE UNIQUE INDEX `ent_types_type_key` ON `ent_types` (`type`);
-- Add pk ranges for ('ability_scores'),('alignments'),('characters'),('classes'),('languages'),('magic_schools'),('races'),('skills') tables
INSERT INTO `ent_types` (`type`) VALUES ('ability_scores'), ('alignments'), ('characters'), ('classes'), ('languages'), ('magic_schools'), ('races'), ('skills');
