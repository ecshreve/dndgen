--
-- File generated with SQLiteStudio v3.4.4 on Thu Aug 31 15:08:32 2023
--
-- Text encoding used: UTF-8
--
PRAGMA foreign_keys = off;
BEGIN TRANSACTION;

-- Table: ability_bonus
CREATE TABLE IF NOT EXISTS ability_bonus (
    id                      INTEGER NOT NULL
                                    PRIMARY KEY AUTOINCREMENT,
    bonus                   INTEGER NOT NULL,
    ability_score_id        INTEGER NOT NULL,
    race_ability_bonuses    INTEGER,
    subrace_ability_bonuses INTEGER,
    CONSTRAINT ability_bonus_ability_scores_ability_bonuses FOREIGN KEY (
        ability_score_id
    )
    REFERENCES ability_scores (id) ON DELETE NO ACTION,
    CONSTRAINT ability_bonus_races_ability_bonuses FOREIGN KEY (
        race_ability_bonuses
    )
    REFERENCES races (id) ON DELETE SET NULL,
    CONSTRAINT ability_bonus_subraces_ability_bonuses FOREIGN KEY (
        subrace_ability_bonuses
    )
    REFERENCES subraces (id) ON DELETE SET NULL
);

INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (1, 2, 4294967299, 55834574849, NULL);
INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (2, 2, 4294967298, 55834574850, NULL);
INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (3, 2, 4294967298, 55834574851, NULL);
INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (4, 1, 4294967297, 55834574852, NULL);
INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (5, 1, 4294967298, 55834574852, NULL);
INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (6, 1, 4294967299, 55834574852, NULL);
INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (7, 1, 4294967300, 55834574852, NULL);
INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (8, 1, 4294967301, 55834574852, NULL);
INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (9, 1, 4294967302, 55834574852, NULL);
INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (10, 2, 4294967297, 55834574853, NULL);
INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (11, 1, 4294967302, 55834574853, NULL);
INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (12, 2, 4294967300, 55834574854, NULL);
INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (13, 2, 4294967302, 55834574855, NULL);
INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (14, 2, 4294967297, 55834574856, NULL);
INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (15, 1, 4294967299, 55834574856, NULL);
INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (16, 1, 4294967300, 55834574857, NULL);
INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (17, 2, 4294967302, 55834574857, NULL);
INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (18, 1, 4294967301, NULL, 73014444033);
INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (19, 1, 4294967300, NULL, 73014444034);
INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (20, 1, 4294967302, NULL, 73014444035);
INSERT INTO ability_bonus (id, bonus, ability_score_id, race_ability_bonuses, subrace_ability_bonuses) VALUES (21, 1, 4294967299, NULL, 73014444036);

-- Table: ability_scores
CREATE TABLE IF NOT EXISTS ability_scores (
    id        INTEGER NOT NULL
                      PRIMARY KEY AUTOINCREMENT,
    indx      TEXT    NOT NULL,
    name      TEXT    NOT NULL,
    full_name TEXT    NOT NULL,
    desc      JSON    NOT NULL
);

INSERT INTO ability_scores (id, indx, name, full_name, desc) VALUES (4294967297, 'str', 'STR', 'Strength', X'5B22537472656E677468206D6561737572657320626F64696C7920706F7765722C206174686C6574696320747261696E696E672C20616E642074686520657874656E7420746F20776869636820796F752063616E2065786572742072617720706879736963616C20666F7263652E222C224120537472656E67746820636865636B2063616E206D6F64656C20616E7920617474656D707420746F206C6966742C20707573682C2070756C6C2C206F7220627265616B20736F6D657468696E672C20746F20666F72636520796F757220626F6479207468726F75676820612073706163652C206F7220746F206F7468657277697365206170706C7920627275746520666F72636520746F206120736974756174696F6E2E20546865204174686C657469637320736B696C6C207265666C6563747320617074697475646520696E206365727461696E206B696E6473206F6620537472656E67746820636865636B732E225D');
INSERT INTO ability_scores (id, indx, name, full_name, desc) VALUES (4294967298, 'dex', 'DEX', 'Dexterity', X'5B22446578746572697479206D65617375726573206167696C6974792C207265666C657865732C20616E642062616C616E63652E222C22412044657874657269747920636865636B2063616E206D6F64656C20616E7920617474656D707420746F206D6F7665206E696D626C792C20717569636B6C792C206F722071756965746C792C206F7220746F206B6565702066726F6D2066616C6C696E67206F6E20747269636B7920666F6F74696E672E20546865204163726F6261746963732C20536C6569676874206F662048616E642C20616E6420537465616C746820736B696C6C73207265666C65637420617074697475646520696E206365727461696E206B696E6473206F662044657874657269747920636865636B732E225D');
INSERT INTO ability_scores (id, indx, name, full_name, desc) VALUES (4294967299, 'con', 'CON', 'Constitution', X'5B22436F6E737469747574696F6E206D65617375726573206865616C74682C207374616D696E612C20616E6420766974616C20666F7263652E222C22436F6E737469747574696F6E20636865636B732061726520756E636F6D6D6F6E2C20616E64206E6F20736B696C6C73206170706C7920746F20436F6E737469747574696F6E20636865636B732C20626563617573652074686520656E647572616E63652074686973206162696C69747920726570726573656E7473206973206C617267656C79207061737369766520726174686572207468616E20696E766F6C76696E672061207370656369666963206566666F7274206F6E207468652070617274206F66206120636861726163746572206F72206D6F6E737465722E225D');
INSERT INTO ability_scores (id, indx, name, full_name, desc) VALUES (4294967300, 'int', 'INT', 'Intelligence', X'5B22496E74656C6C6967656E6365206D65617375726573206D656E74616C206163756974792C206163637572616379206F6620726563616C6C2C20616E6420746865206162696C69747920746F20726561736F6E2E222C22416E20496E74656C6C6967656E636520636865636B20636F6D657320696E746F20706C6179207768656E20796F75206E65656420746F2064726177206F6E206C6F6769632C20656475636174696F6E2C206D656D6F72792C206F722064656475637469766520726561736F6E696E672E2054686520417263616E612C20486973746F72792C20496E7665737469676174696F6E2C204E61747572652C20616E642052656C6967696F6E20736B696C6C73207265666C65637420617074697475646520696E206365727461696E206B696E6473206F6620496E74656C6C6967656E636520636865636B732E225D');
INSERT INTO ability_scores (id, indx, name, full_name, desc) VALUES (4294967301, 'wis', 'WIS', 'Wisdom', X'5B22576973646F6D207265666C6563747320686F7720617474756E656420796F752061726520746F2074686520776F726C642061726F756E6420796F7520616E6420726570726573656E747320706572636570746976656E65737320616E6420696E74756974696F6E2E222C224120576973646F6D20636865636B206D69676874207265666C65637420616E206566666F727420746F207265616420626F6479206C616E67756167652C20756E6465727374616E6420736F6D656F6E652773206665656C696E67732C206E6F74696365207468696E67732061626F75742074686520656E7669726F6E6D656E742C206F72206361726520666F7220616E20696E6A7572656420706572736F6E2E2054686520416E696D616C2048616E646C696E672C20496E73696768742C204D65646963696E652C2050657263657074696F6E2C20616E6420537572766976616C20736B696C6C73207265666C65637420617074697475646520696E206365727461696E206B696E6473206F6620576973646F6D20636865636B732E225D');
INSERT INTO ability_scores (id, indx, name, full_name, desc) VALUES (4294967302, 'cha', 'CHA', 'Charisma', X'5B224368617269736D61206D6561737572657320796F7572206162696C69747920746F20696E746572616374206566666563746976656C792077697468206F74686572732E20497420696E636C75646573207375636820666163746F727320617320636F6E666964656E636520616E6420656C6F7175656E63652C20616E642069742063616E20726570726573656E74206120636861726D696E67206F7220636F6D6D616E64696E6720706572736F6E616C6974792E222C2241204368617269736D6120636865636B206D69676874206172697365207768656E20796F752074727920746F20696E666C75656E6365206F7220656E7465727461696E206F74686572732C207768656E20796F752074727920746F206D616B6520616E20696D7072657373696F6E206F722074656C6C206120636F6E76696E63696E67206C69652C206F72207768656E20796F7520617265206E617669676174696E67206120747269636B7920736F6369616C20736974756174696F6E2E2054686520446563657074696F6E2C20496E74696D69646174696F6E2C20506572666F726D616E63652C20616E642050657273756173696F6E20736B696C6C73207265666C65637420617074697475646520696E206365727461696E206B696E6473206F66204368617269736D6120636865636B732E225D');

-- Table: armor_classes
CREATE TABLE IF NOT EXISTS armor_classes (
    id                INTEGER NOT NULL
                              PRIMARY KEY AUTOINCREMENT,
    base              INTEGER NOT NULL,
    dex_bonus         BOOL    NOT NULL,
    max_bonus         INTEGER,
    armor_armor_class INTEGER,
    CONSTRAINT armor_classes_armors_armor_class FOREIGN KEY (
        armor_armor_class
    )
    REFERENCES armors (id) ON DELETE SET NULL
);

INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (12884901889, 11, 1, 0, 8589934593);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (12884901890, 11, 1, 0, 8589934594);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (12884901891, 12, 1, 0, 8589934595);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (12884901892, 12, 1, 2, 8589934596);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (12884901893, 13, 1, 2, 8589934597);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (12884901894, 14, 1, 2, 8589934598);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (12884901895, 14, 1, 2, 8589934599);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (12884901896, 15, 1, 2, 8589934600);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (12884901897, 14, 0, 0, 8589934601);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (12884901898, 16, 0, 0, 8589934602);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (12884901899, 17, 0, 0, 8589934603);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (12884901900, 18, 0, 0, 8589934604);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (12884901901, 2, 0, 0, 8589934605);

-- Table: armors
CREATE TABLE IF NOT EXISTS armors (
    id                   INTEGER NOT NULL
                                 PRIMARY KEY AUTOINCREMENT,
    indx                 TEXT    NOT NULL,
    name                 TEXT    NOT NULL,
    armor_category       TEXT    NOT NULL,
    stealth_disadvantage BOOL    NOT NULL,
    min_strength         INTEGER NOT NULL,
    equipment_id         INTEGER NOT NULL,
    CONSTRAINT armors_equipment_armor FOREIGN KEY (
        equipment_id
    )
    REFERENCES equipment (id) ON DELETE NO ACTION
);

INSERT INTO armors (id, indx, name, armor_category, stealth_disadvantage, min_strength, equipment_id) VALUES (8589934593, 'padded-armor', 'Padded Armor', 'light', 1, 0, 34359738406);
INSERT INTO armors (id, indx, name, armor_category, stealth_disadvantage, min_strength, equipment_id) VALUES (8589934594, 'leather-armor', 'Leather Armor', 'light', 0, 0, 34359738407);
INSERT INTO armors (id, indx, name, armor_category, stealth_disadvantage, min_strength, equipment_id) VALUES (8589934595, 'studded-leather-armor', 'Studded Leather Armor', 'light', 0, 0, 34359738408);
INSERT INTO armors (id, indx, name, armor_category, stealth_disadvantage, min_strength, equipment_id) VALUES (8589934596, 'hide-armor', 'Hide Armor', 'medium', 0, 0, 34359738409);
INSERT INTO armors (id, indx, name, armor_category, stealth_disadvantage, min_strength, equipment_id) VALUES (8589934597, 'chain-shirt', 'Chain Shirt', 'medium', 0, 0, 34359738410);
INSERT INTO armors (id, indx, name, armor_category, stealth_disadvantage, min_strength, equipment_id) VALUES (8589934598, 'scale-mail', 'Scale Mail', 'medium', 1, 0, 34359738411);
INSERT INTO armors (id, indx, name, armor_category, stealth_disadvantage, min_strength, equipment_id) VALUES (8589934599, 'breastplate', 'Breastplate', 'medium', 0, 0, 34359738412);
INSERT INTO armors (id, indx, name, armor_category, stealth_disadvantage, min_strength, equipment_id) VALUES (8589934600, 'half-plate-armor', 'Half Plate Armor', 'medium', 1, 0, 34359738413);
INSERT INTO armors (id, indx, name, armor_category, stealth_disadvantage, min_strength, equipment_id) VALUES (8589934601, 'ring-mail', 'Ring Mail', 'heavy', 1, 0, 34359738414);
INSERT INTO armors (id, indx, name, armor_category, stealth_disadvantage, min_strength, equipment_id) VALUES (8589934602, 'chain-mail', 'Chain Mail', 'heavy', 1, 13, 34359738415);
INSERT INTO armors (id, indx, name, armor_category, stealth_disadvantage, min_strength, equipment_id) VALUES (8589934603, 'splint-armor', 'Splint Armor', 'heavy', 1, 15, 34359738416);
INSERT INTO armors (id, indx, name, armor_category, stealth_disadvantage, min_strength, equipment_id) VALUES (8589934604, 'plate-armor', 'Plate Armor', 'heavy', 1, 15, 34359738417);
INSERT INTO armors (id, indx, name, armor_category, stealth_disadvantage, min_strength, equipment_id) VALUES (8589934605, 'shield', 'Shield', 'shield', 0, 0, 34359738418);

-- Table: choice_proficiencies
CREATE TABLE IF NOT EXISTS choice_proficiencies (
    choice_id      INTEGER NOT NULL,
    proficiency_id INTEGER NOT NULL,
    PRIMARY KEY (
        choice_id,
        proficiency_id
    ),
    CONSTRAINT choice_proficiencies_choice_id FOREIGN KEY (
        choice_id
    )
    REFERENCES choices (id) ON DELETE CASCADE,
    CONSTRAINT choice_proficiencies_proficiency_id FOREIGN KEY (
        proficiency_id
    )
    REFERENCES proficiencies (id) ON DELETE CASCADE
);


-- Table: choices
CREATE TABLE IF NOT EXISTS choices (
    id      INTEGER NOT NULL
                    PRIMARY KEY AUTOINCREMENT,
    choose  INTEGER NOT NULL,
    race_id INTEGER,
    CONSTRAINT choices_races_starting_proficiency_option FOREIGN KEY (
        race_id
    )
    REFERENCES races (id) ON DELETE SET NULL
);


-- Table: class_proficiencies
CREATE TABLE IF NOT EXISTS class_proficiencies (
    class_id       INTEGER NOT NULL,
    proficiency_id INTEGER NOT NULL,
    PRIMARY KEY (
        class_id,
        proficiency_id
    ),
    CONSTRAINT class_proficiencies_class_id FOREIGN KEY (
        class_id
    )
    REFERENCES classes (id) ON DELETE CASCADE,
    CONSTRAINT class_proficiencies_proficiency_id FOREIGN KEY (
        proficiency_id
    )
    REFERENCES proficiencies (id) ON DELETE CASCADE
);

INSERT INTO class_proficiencies (class_id, proficiency_id) VALUES (21474836489, 51539607553);
INSERT INTO class_proficiencies (class_id, proficiency_id) VALUES (21474836491, 51539607553);
INSERT INTO class_proficiencies (class_id, proficiency_id) VALUES (21474836481, 51539607553);
INSERT INTO class_proficiencies (class_id, proficiency_id) VALUES (21474836482, 51539607553);
INSERT INTO class_proficiencies (class_id, proficiency_id) VALUES (21474836483, 51539607553);
INSERT INTO class_proficiencies (class_id, proficiency_id) VALUES (21474836484, 51539607553);
INSERT INTO class_proficiencies (class_id, proficiency_id) VALUES (21474836488, 51539607553);
INSERT INTO class_proficiencies (class_id, proficiency_id) VALUES (21474836481, 51539607554);
INSERT INTO class_proficiencies (class_id, proficiency_id) VALUES (21474836483, 51539607554);
INSERT INTO class_proficiencies (class_id, proficiency_id) VALUES (21474836484, 51539607554);
INSERT INTO class_proficiencies (class_id, proficiency_id) VALUES (21474836488, 51539607554);
INSERT INTO class_proficiencies (class_id, proficiency_id) VALUES (21474836487, 51539607556);
INSERT INTO class_proficiencies (class_id, proficiency_id) VALUES (21474836485, 51539607556);

-- Table: classes
CREATE TABLE IF NOT EXISTS classes (
    id      INTEGER NOT NULL
                    PRIMARY KEY AUTOINCREMENT,
    indx    TEXT    NOT NULL,
    name    TEXT    NOT NULL,
    hit_die INTEGER NOT NULL
);

INSERT INTO classes (id, indx, name, hit_die) VALUES (21474836481, 'barbarian', 'Barbarian', 12);
INSERT INTO classes (id, indx, name, hit_die) VALUES (21474836482, 'bard', 'Bard', 8);
INSERT INTO classes (id, indx, name, hit_die) VALUES (21474836483, 'cleric', 'Cleric', 8);
INSERT INTO classes (id, indx, name, hit_die) VALUES (21474836484, 'druid', 'Druid', 8);
INSERT INTO classes (id, indx, name, hit_die) VALUES (21474836485, 'fighter', 'Fighter', 10);
INSERT INTO classes (id, indx, name, hit_die) VALUES (21474836486, 'monk', 'Monk', 8);
INSERT INTO classes (id, indx, name, hit_die) VALUES (21474836487, 'paladin', 'Paladin', 10);
INSERT INTO classes (id, indx, name, hit_die) VALUES (21474836488, 'ranger', 'Ranger', 10);
INSERT INTO classes (id, indx, name, hit_die) VALUES (21474836489, 'rogue', 'Rogue', 8);
INSERT INTO classes (id, indx, name, hit_die) VALUES (21474836490, 'sorcerer', 'Sorcerer', 6);
INSERT INTO classes (id, indx, name, hit_die) VALUES (21474836491, 'warlock', 'Warlock', 8);
INSERT INTO classes (id, indx, name, hit_die) VALUES (21474836492, 'wizard', 'Wizard', 6);

-- Table: costs
CREATE TABLE IF NOT EXISTS costs (
    id       INTEGER NOT NULL
                     PRIMARY KEY AUTOINCREMENT,
    quantity INTEGER NOT NULL,
    unit     TEXT    NOT NULL
);

INSERT INTO costs (id, quantity, unit) VALUES (25769803777, 1, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803778, 2, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803779, 2, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803780, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803781, 5, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803782, 2, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803783, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803784, 2, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803785, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803786, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803787, 25, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803788, 5, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803789, 25, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803790, 1, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803791, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803792, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803793, 20, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803794, 30, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803795, 50, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803796, 20, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803797, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803798, 15, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803799, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803800, 15, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803801, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803802, 25, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803803, 25, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803804, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803805, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803806, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803807, 15, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803808, 2, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803809, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803810, 75, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803811, 50, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803812, 50, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803813, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803814, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803815, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803816, 45, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803817, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803818, 50, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803819, 50, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803820, 400, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803821, 750, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803822, 30, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803823, 75, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803824, 200, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803825, 1500, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803826, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803827, 2, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803828, 25, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803829, 50, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803830, 0, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803831, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803832, 0, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803833, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803834, 0, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803835, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803836, 4, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803837, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803838, 50, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803839, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803840, 20, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803841, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803842, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803843, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803844, 2, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803845, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803846, 2, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803847, 4, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803848, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803849, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803850, 5, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803851, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803852, 25, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803853, 2, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803854, 5, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803855, 5, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803856, 1, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803857, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803858, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803859, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803860, 1, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803861, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803862, 5, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803863, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803864, 15, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803865, 2, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803866, 25, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803867, 2, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803868, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803869, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803870, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803871, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803872, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803873, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803874, 2, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803875, 2, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803876, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803877, 2, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803878, 25, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803879, 25, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803880, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803881, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803882, 2, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803883, 2, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803884, 25, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803885, 25, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803886, 15, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803887, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803888, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803889, 2, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803890, 50, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803891, 1, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803892, 5, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803893, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803894, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803895, 0, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803896, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803897, 100, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803898, 2, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803899, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803900, 1, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803901, 2, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803902, 1, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803903, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803904, 2, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803905, 5, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803906, 100, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803907, 5, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803908, 2, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803909, 5, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803910, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803911, 4, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803912, 5, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803913, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803914, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803915, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803916, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803917, 1, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803918, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803919, 5, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803920, 2, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803921, 5, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803922, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803923, 0, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803924, 2, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803925, 50, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803926, 1, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803927, 1000, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803928, 0, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803929, 2, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803930, 5, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803931, 1, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803932, 0, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803933, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803934, 2, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803935, 1, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803936, 16, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803937, 39, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803938, 12, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803939, 40, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803940, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803941, 19, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803942, 40, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803943, 50, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803944, 20, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803945, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803946, 8, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803947, 15, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803948, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803949, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803950, 30, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803951, 25, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803952, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803953, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803954, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803955, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803956, 20, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803957, 50, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803958, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803959, 1, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803960, 1, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803961, 5, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803962, 30, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803963, 6, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803964, 25, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803965, 2, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803966, 35, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803967, 30, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803968, 3, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803969, 12, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803970, 2, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803971, 30, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803972, 25, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803973, 25, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803974, 50, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803975, 8, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803976, 8, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803977, 200, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803978, 50, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803979, 75, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803980, 25, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803981, 30, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803982, 400, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803983, 20, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803984, 40, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803985, 180, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803986, 40, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803987, 200, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803988, 200, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803989, 1600, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803990, 3000, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803991, 12, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803992, 300, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803993, 800, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803994, 6000, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803995, 2, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803996, 100, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803997, 15, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803998, 250, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769803999, 5, 'cp');
INSERT INTO costs (id, quantity, unit) VALUES (25769804000, 60, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769804001, 20, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769804002, 5, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769804003, 10, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769804004, 4, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769804005, 20, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769804006, 5, 'sp');
INSERT INTO costs (id, quantity, unit) VALUES (25769804007, 35, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769804008, 30000, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769804009, 3000, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769804010, 10000, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769804011, 50, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769804012, 10000, 'gp');
INSERT INTO costs (id, quantity, unit) VALUES (25769804013, 25000, 'gp');

-- Table: damage_types
CREATE TABLE IF NOT EXISTS damage_types (
    id   INTEGER NOT NULL
                 PRIMARY KEY AUTOINCREMENT,
    indx TEXT    NOT NULL,
    name TEXT    NOT NULL,
    desc JSON    NOT NULL
);

INSERT INTO damage_types (id, indx, name, desc) VALUES (30064771073, 'acid', 'Acid', X'5B2254686520636F72726F73697665207370726179206F66206120626C61636B20647261676F6E27732062726561746820616E642074686520646973736F6C76696E6720656E7A796D6573207365637265746564206279206120626C61636B2070756464696E67206465616C20616369642064616D6167652E225D');
INSERT INTO damage_types (id, indx, name, desc) VALUES (30064771074, 'bludgeoning', 'Bludgeoning', X'5B22426C756E7420666F7263652061747461636B732C2066616C6C696E672C20636F6E737472696374696F6E2C20616E6420746865206C696B65206465616C20626C756467656F6E696E672064616D6167652E225D');
INSERT INTO damage_types (id, indx, name, desc) VALUES (30064771075, 'cold', 'Cold', X'5B2254686520696E6665726E616C206368696C6C20726164696174696E672066726F6D20616E2069636520646576696C277320737065617220616E64207468652066726967696420626C617374206F66206120776869746520647261676F6E277320627265617468206465616C20636F6C642064616D6167652E225D');
INSERT INTO damage_types (id, indx, name, desc) VALUES (30064771076, 'fire', 'Fire', X'5B2252656420647261676F6E73206272656174686520666972652C20616E64206D616E79207370656C6C7320636F6E6A75726520666C616D657320746F206465616C20666972652064616D6167652E225D');
INSERT INTO damage_types (id, indx, name, desc) VALUES (30064771077, 'force', 'Force', X'5B22466F7263652069732070757265206D61676963616C20656E6572677920666F637573656420696E746F20612064616D6167696E6720666F726D2E204D6F737420656666656374732074686174206465616C20666F7263652064616D61676520617265207370656C6C732C20696E636C7564696E67206D61676963206D697373696C6520616E642073706972697475616C20776561706F6E2E225D');
INSERT INTO damage_types (id, indx, name, desc) VALUES (30064771078, 'lightning', 'Lightning', X'5B2241206C696768746E696E6720626F6C74207370656C6C20616E64206120626C756520647261676F6E277320627265617468206465616C206C696768746E696E672064616D6167652E225D');
INSERT INTO damage_types (id, indx, name, desc) VALUES (30064771079, 'necrotic', 'Necrotic', X'5B224E6563726F7469632064616D6167652C206465616C74206279206365727461696E20756E6465616420616E642061207370656C6C2073756368206173206368696C6C20746F7563682C2077697468657273206D617474657220616E64206576656E2074686520736F756C2E225D');
INSERT INTO damage_types (id, indx, name, desc) VALUES (30064771080, 'piercing', 'Piercing', X'5B2250756E63747572696E6720616E6420696D70616C696E672061747461636B732C20696E636C7564696E672073706561727320616E64206D6F6E7374657273272062697465732C206465616C207069657263696E672064616D6167652E225D');
INSERT INTO damage_types (id, indx, name, desc) VALUES (30064771081, 'poison', 'Poison', X'5B2256656E6F6D6F7573207374696E677320616E642074686520746F78696320676173206F66206120677265656E20647261676F6E277320627265617468206465616C20706F69736F6E2064616D6167652E225D');
INSERT INTO damage_types (id, indx, name, desc) VALUES (30064771082, 'psychic', 'Psychic', X'5B224D656E74616C206162696C697469657320737563682061732061207073696F6E696320626C617374206465616C20707379636869632064616D6167652E225D');
INSERT INTO damage_types (id, indx, name, desc) VALUES (30064771083, 'radiant', 'Radiant', X'5B2252616469616E742064616D6167652C206465616C74206279206120636C65726963277320666C616D6520737472696B65207370656C6C206F7220616E20616E67656C277320736D6974696E6720776561706F6E2C2073656172732074686520666C657368206C696B65206669726520616E64206F7665726C6F6164732074686520737069726974207769746820706F7765722E225D');
INSERT INTO damage_types (id, indx, name, desc) VALUES (30064771084, 'slashing', 'Slashing', X'5B2253776F7264732C20617865732C20616E64206D6F6E73746572732720636C617773206465616C20736C617368696E672064616D6167652E225D');
INSERT INTO damage_types (id, indx, name, desc) VALUES (30064771085, 'thunder', 'Thunder', X'5B224120636F6E63757373697665206275727374206F6620736F756E642C20737563682061732074686520656666656374206F6620746865207468756E64657277617665207370656C6C2C206465616C73207468756E6465722064616D6167652E225D');

-- Table: ent_types
CREATE TABLE IF NOT EXISTS ent_types (
    id   INTEGER NOT NULL
                 PRIMARY KEY AUTOINCREMENT,
    type TEXT    NOT NULL
);

INSERT INTO ent_types (id, type) VALUES (1, 'ability_bonus');
INSERT INTO ent_types (id, type) VALUES (2, 'ability_scores');
INSERT INTO ent_types (id, type) VALUES (3, 'armors');
INSERT INTO ent_types (id, type) VALUES (4, 'armor_classes');
INSERT INTO ent_types (id, type) VALUES (5, 'choices');
INSERT INTO ent_types (id, type) VALUES (6, 'classes');
INSERT INTO ent_types (id, type) VALUES (7, 'costs');
INSERT INTO ent_types (id, type) VALUES (8, 'damage_types');
INSERT INTO ent_types (id, type) VALUES (9, 'equipment');
INSERT INTO ent_types (id, type) VALUES (10, 'gears');
INSERT INTO ent_types (id, type) VALUES (11, 'languages');
INSERT INTO ent_types (id, type) VALUES (12, 'magic_schools');
INSERT INTO ent_types (id, type) VALUES (13, 'proficiencies');
INSERT INTO ent_types (id, type) VALUES (14, 'races');
INSERT INTO ent_types (id, type) VALUES (15, 'rules');
INSERT INTO ent_types (id, type) VALUES (16, 'rule_sections');
INSERT INTO ent_types (id, type) VALUES (17, 'skills');
INSERT INTO ent_types (id, type) VALUES (18, 'subraces');
INSERT INTO ent_types (id, type) VALUES (19, 'tools');
INSERT INTO ent_types (id, type) VALUES (20, 'traits');
INSERT INTO ent_types (id, type) VALUES (21, 'vehicles');
INSERT INTO ent_types (id, type) VALUES (22, 'weapons');
INSERT INTO ent_types (id, type) VALUES (23, 'weapon_damages');
INSERT INTO ent_types (id, type) VALUES (24, 'weapon_properties');

-- Table: equipment
CREATE TABLE IF NOT EXISTS equipment (
    id                 INTEGER NOT NULL
                               PRIMARY KEY AUTOINCREMENT,
    indx               TEXT    NOT NULL,
    name               TEXT    NOT NULL,
    equipment_category TEXT    NOT NULL
                               DEFAULT 'other',
    equipment_cost     INTEGER,
    CONSTRAINT equipment_costs_cost FOREIGN KEY (
        equipment_cost
    )
    REFERENCES costs (id) ON DELETE SET NULL
);

INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738369, 'club', 'Club', 'weapon', 25769803777);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738370, 'dagger', 'Dagger', 'weapon', 25769803778);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738371, 'greatclub', 'Greatclub', 'weapon', 25769803779);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738372, 'handaxe', 'Handaxe', 'weapon', 25769803780);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738373, 'javelin', 'Javelin', 'weapon', 25769803781);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738374, 'light-hammer', 'Light hammer', 'weapon', 25769803782);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738375, 'mace', 'Mace', 'weapon', 25769803783);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738376, 'quarterstaff', 'Quarterstaff', 'weapon', 25769803784);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738377, 'sickle', 'Sickle', 'weapon', 25769803785);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738378, 'spear', 'Spear', 'weapon', 25769803786);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738379, 'crossbow-light', 'Crossbow, light', 'weapon', 25769803787);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738380, 'dart', 'Dart', 'weapon', 25769803788);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738381, 'shortbow', 'Shortbow', 'weapon', 25769803789);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738382, 'sling', 'Sling', 'weapon', 25769803790);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738383, 'battleaxe', 'Battleaxe', 'weapon', 25769803791);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738384, 'flail', 'Flail', 'weapon', 25769803792);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738385, 'glaive', 'Glaive', 'weapon', 25769803793);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738386, 'greataxe', 'Greataxe', 'weapon', 25769803794);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738387, 'greatsword', 'Greatsword', 'weapon', 25769803795);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738388, 'halberd', 'Halberd', 'weapon', 25769803796);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738389, 'lance', 'Lance', 'weapon', 25769803797);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738390, 'longsword', 'Longsword', 'weapon', 25769803798);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738391, 'maul', 'Maul', 'weapon', 25769803799);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738392, 'morningstar', 'Morningstar', 'weapon', 25769803800);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738393, 'pike', 'Pike', 'weapon', 25769803801);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738394, 'rapier', 'Rapier', 'weapon', 25769803802);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738395, 'scimitar', 'Scimitar', 'weapon', 25769803803);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738396, 'shortsword', 'Shortsword', 'weapon', 25769803804);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738397, 'trident', 'Trident', 'weapon', 25769803805);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738398, 'war-pick', 'War pick', 'weapon', 25769803806);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738399, 'warhammer', 'Warhammer', 'weapon', 25769803807);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738400, 'whip', 'Whip', 'weapon', 25769803808);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738401, 'blowgun', 'Blowgun', 'weapon', 25769803809);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738402, 'crossbow-hand', 'Crossbow, hand', 'weapon', 25769803810);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738403, 'crossbow-heavy', 'Crossbow, heavy', 'weapon', 25769803811);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738404, 'longbow', 'Longbow', 'weapon', 25769803812);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738405, 'net', 'Net', 'weapon', 25769803813);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738406, 'padded-armor', 'Padded Armor', 'armor', 25769803814);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738407, 'leather-armor', 'Leather Armor', 'armor', 25769803815);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738408, 'studded-leather-armor', 'Studded Leather Armor', 'armor', 25769803816);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738409, 'hide-armor', 'Hide Armor', 'armor', 25769803817);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738410, 'chain-shirt', 'Chain Shirt', 'armor', 25769803818);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738411, 'scale-mail', 'Scale Mail', 'armor', 25769803819);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738412, 'breastplate', 'Breastplate', 'armor', 25769803820);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738413, 'half-plate-armor', 'Half Plate Armor', 'armor', 25769803821);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738414, 'ring-mail', 'Ring Mail', 'armor', 25769803822);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738415, 'chain-mail', 'Chain Mail', 'armor', 25769803823);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738416, 'splint-armor', 'Splint Armor', 'armor', 25769803824);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738417, 'plate-armor', 'Plate Armor', 'armor', 25769803825);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738418, 'shield', 'Shield', 'armor', 25769803826);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738419, 'abacus', 'Abacus', 'adventuring_gear', 25769803827);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738420, 'acid-vial', 'Acid (vial)', 'adventuring_gear', 25769803828);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738421, 'alchemists-fire-flask', 'Alchemist''s fire (flask)', 'adventuring_gear', 25769803829);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738422, 'alms-box', 'Alms box', 'adventuring_gear', 25769803830);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738423, 'arrow', 'Arrow', 'adventuring_gear', 25769803831);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738424, 'block-of-incense', 'Block of incense', 'adventuring_gear', 25769803832);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738425, 'blowgun-needle', 'Blowgun needle', 'adventuring_gear', 25769803833);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738426, 'censer', 'Censer', 'adventuring_gear', 25769803834);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738427, 'crossbow-bolt', 'Crossbow bolt', 'adventuring_gear', 25769803835);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738428, 'sling-bullet', 'Sling bullet', 'adventuring_gear', 25769803836);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738429, 'amulet', 'Amulet', 'adventuring_gear', 25769803837);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738430, 'antitoxin-vial', 'Antitoxin (vial)', 'adventuring_gear', 25769803838);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738431, 'crystal', 'Crystal', 'adventuring_gear', 25769803839);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738432, 'orb', 'Orb', 'adventuring_gear', 25769803840);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738433, 'rod', 'Rod', 'adventuring_gear', 25769803841);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738434, 'staff', 'Staff', 'adventuring_gear', 25769803842);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738435, 'wand', 'Wand', 'adventuring_gear', 25769803843);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738436, 'backpack', 'Backpack', 'adventuring_gear', 25769803844);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738437, 'ball-bearings-bag-of-1000', 'Ball bearings (bag of 1,000)', 'adventuring_gear', 25769803845);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738438, 'barrel', 'Barrel', 'adventuring_gear', 25769803846);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738439, 'basket', 'Basket', 'adventuring_gear', 25769803847);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738440, 'bedroll', 'Bedroll', 'adventuring_gear', 25769803848);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738441, 'bell', 'Bell', 'adventuring_gear', 25769803849);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738442, 'blanket', 'Blanket', 'adventuring_gear', 25769803850);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738443, 'block-and-tackle', 'Block and tackle', 'adventuring_gear', 25769803851);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738444, 'book', 'Book', 'adventuring_gear', 25769803852);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738445, 'bottle-glass', 'Bottle, glass', 'adventuring_gear', 25769803853);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738446, 'bucket', 'Bucket', 'adventuring_gear', 25769803854);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738447, 'caltrops', 'Caltrops', 'adventuring_gear', 25769803855);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738448, 'candle', 'Candle', 'adventuring_gear', 25769803856);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738449, 'case-crossbow-bolt', 'Case, crossbow bolt', 'adventuring_gear', 25769803857);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738450, 'case-map-or-scroll', 'Case, map or scroll', 'adventuring_gear', 25769803858);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738451, 'chain-10-feet', 'Chain (10 feet)', 'adventuring_gear', 25769803859);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738452, 'chalk-1-piece', 'Chalk (1 piece)', 'adventuring_gear', 25769803860);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738453, 'chest', 'Chest', 'adventuring_gear', 25769803861);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738454, 'clothes-common', 'Clothes, common', 'adventuring_gear', 25769803862);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738455, 'clothes-costume', 'Clothes, costume', 'adventuring_gear', 25769803863);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738456, 'clothes-fine', 'Clothes, fine', 'adventuring_gear', 25769803864);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738457, 'clothes-travelers', 'Clothes, traveler''s', 'adventuring_gear', 25769803865);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738458, 'component-pouch', 'Component pouch', 'adventuring_gear', 25769803866);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738459, 'crowbar', 'Crowbar', 'adventuring_gear', 25769803867);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738460, 'sprig-of-mistletoe', 'Sprig of mistletoe', 'adventuring_gear', 25769803868);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738461, 'totem', 'Totem', 'adventuring_gear', 25769803869);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738462, 'wooden-staff', 'Wooden staff', 'adventuring_gear', 25769803870);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738463, 'yew-wand', 'Yew wand', 'adventuring_gear', 25769803871);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738464, 'emblem', 'Emblem', 'adventuring_gear', 25769803872);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738465, 'fishing-tackle', 'Fishing tackle', 'adventuring_gear', 25769803873);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738466, 'flask-or-tankard', 'Flask or tankard', 'adventuring_gear', 25769803874);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738467, 'grappling-hook', 'Grappling hook', 'adventuring_gear', 25769803875);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738468, 'hammer', 'Hammer', 'adventuring_gear', 25769803876);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738469, 'hammer-sledge', 'Hammer, sledge', 'adventuring_gear', 25769803877);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738470, 'holy-water-flask', 'Holy water (flask)', 'adventuring_gear', 25769803878);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738471, 'hourglass', 'Hourglass', 'adventuring_gear', 25769803879);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738472, 'hunting-trap', 'Hunting trap', 'adventuring_gear', 25769803880);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738473, 'ink-1-ounce-bottle', 'Ink (1 ounce bottle)', 'adventuring_gear', 25769803881);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738474, 'ink-pen', 'Ink pen', 'adventuring_gear', 25769803882);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738475, 'jug-or-pitcher', 'Jug or pitcher', 'adventuring_gear', 25769803883);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738476, 'climbers-kit', 'Climber''s Kit', 'adventuring_gear', 25769803884);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738477, 'disguise-kit', 'Disguise Kit', 'adventuring_gear', 25769803885);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738478, 'forgery-kit', 'Forgery Kit', 'adventuring_gear', 25769803886);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738479, 'herbalism-kit', 'Herbalism Kit', 'adventuring_gear', 25769803887);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738480, 'healers-kit', 'Healer''s Kit', 'adventuring_gear', 25769803888);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738481, 'mess-kit', 'Mess Kit', 'adventuring_gear', 25769803889);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738482, 'poisoners-kit', 'Poisoner''s Kit', 'adventuring_gear', 25769803890);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738483, 'ladder-10-foot', 'Ladder (10-foot)', 'adventuring_gear', 25769803891);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738484, 'lamp', 'Lamp', 'adventuring_gear', 25769803892);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738485, 'lantern-bullseye', 'Lantern, bullseye', 'adventuring_gear', 25769803893);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738486, 'lantern-hooded', 'Lantern, hooded', 'adventuring_gear', 25769803894);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738487, 'little-bag-of-sand', 'Little bag of sand', 'adventuring_gear', 25769803895);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738488, 'lock', 'Lock', 'adventuring_gear', 25769803896);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738489, 'magnifying-glass', 'Magnifying glass', 'adventuring_gear', 25769803897);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738490, 'manacles', 'Manacles', 'adventuring_gear', 25769803898);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738491, 'mirror-steel', 'Mirror, steel', 'adventuring_gear', 25769803899);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738492, 'oil-flask', 'Oil (flask)', 'adventuring_gear', 25769803900);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738493, 'paper-one-sheet', 'Paper (one sheet)', 'adventuring_gear', 25769803901);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738494, 'parchment-one-sheet', 'Parchment (one sheet)', 'adventuring_gear', 25769803902);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738495, 'perfume-vial', 'Perfume (vial)', 'adventuring_gear', 25769803903);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738496, 'pick-miners', 'Pick, miner''s', 'adventuring_gear', 25769803904);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738497, 'piton', 'Piton', 'adventuring_gear', 25769803905);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738498, 'poison-basic-vial', 'Poison, basic (vial)', 'adventuring_gear', 25769803906);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738499, 'pole-10-foot', 'Pole (10-foot)', 'adventuring_gear', 25769803907);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738500, 'pot-iron', 'Pot, iron', 'adventuring_gear', 25769803908);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738501, 'pouch', 'Pouch', 'adventuring_gear', 25769803909);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738502, 'quiver', 'Quiver', 'adventuring_gear', 25769803910);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738503, 'ram-portable', 'Ram, portable', 'adventuring_gear', 25769803911);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738504, 'rations-1-day', 'Rations (1 day)', 'adventuring_gear', 25769803912);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738505, 'reliquary', 'Reliquary', 'adventuring_gear', 25769803913);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738506, 'robes', 'Robes', 'adventuring_gear', 25769803914);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738507, 'rope-hempen-50-feet', 'Rope, hempen (50 feet)', 'adventuring_gear', 25769803915);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738508, 'rope-silk-50-feet', 'Rope, silk (50 feet)', 'adventuring_gear', 25769803916);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738509, 'sack', 'Sack', 'adventuring_gear', 25769803917);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738510, 'scale-merchants', 'Scale, merchant''s', 'adventuring_gear', 25769803918);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738511, 'sealing-wax', 'Sealing wax', 'adventuring_gear', 25769803919);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738512, 'shovel', 'Shovel', 'adventuring_gear', 25769803920);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738513, 'signal-whistle', 'Signal whistle', 'adventuring_gear', 25769803921);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738514, 'signet-ring', 'Signet ring', 'adventuring_gear', 25769803922);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738515, 'small-knife', 'Small knife', 'adventuring_gear', 25769803923);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738516, 'soap', 'Soap', 'adventuring_gear', 25769803924);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738517, 'spellbook', 'Spellbook', 'adventuring_gear', 25769803925);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738518, 'spike-iron', 'Spike, iron', 'adventuring_gear', 25769803926);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738519, 'spyglass', 'Spyglass', 'adventuring_gear', 25769803927);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738520, 'string-10-feet', 'String (10 feet)', 'adventuring_gear', 25769803928);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738521, 'tent-two-person', 'Tent, two-person', 'adventuring_gear', 25769803929);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738522, 'tinderbox', 'Tinderbox', 'adventuring_gear', 25769803930);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738523, 'torch', 'Torch', 'adventuring_gear', 25769803931);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738524, 'vestments', 'Vestments', 'adventuring_gear', 25769803932);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738525, 'vial', 'Vial', 'adventuring_gear', 25769803933);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738526, 'waterskin', 'Waterskin', 'adventuring_gear', 25769803934);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738527, 'whetstone', 'Whetstone', 'adventuring_gear', 25769803935);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738528, 'burglars-pack', 'Burglar''s Pack', 'adventuring_gear', 25769803936);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738529, 'diplomats-pack', 'Diplomat''s Pack', 'adventuring_gear', 25769803937);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738530, 'dungeoneers-pack', 'Dungeoneer''s Pack', 'adventuring_gear', 25769803938);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738531, 'entertainers-pack', 'Entertainer''s Pack', 'adventuring_gear', 25769803939);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738532, 'explorers-pack', 'Explorer''s Pack', 'adventuring_gear', 25769803940);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738533, 'priests-pack', 'Priest''s Pack', 'adventuring_gear', 25769803941);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738534, 'scholars-pack', 'Scholar''s Pack', 'adventuring_gear', 25769803942);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738535, 'alchemists-supplies', 'Alchemist''s Supplies', 'tools', 25769803943);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738536, 'brewers-supplies', 'Brewer''s Supplies', 'tools', 25769803944);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738537, 'calligraphers-supplies', 'Calligrapher''s Supplies', 'tools', 25769803945);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738538, 'carpenters-tools', 'Carpenter''s Tools', 'tools', 25769803946);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738539, 'cartographers-tools', 'Cartographer''s Tools', 'tools', 25769803947);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738540, 'cobblers-tools', 'Cobbler''s Tools', 'tools', 25769803948);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738541, 'cooks-utensils', 'Cook''s utensils', 'tools', 25769803949);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738542, 'glassblowers-tools', 'Glassblower''s Tools', 'tools', 25769803950);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738543, 'jewelers-tools', 'Jeweler''s Tools', 'tools', 25769803951);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738544, 'leatherworkers-tools', 'Leatherworker''s Tools', 'tools', 25769803952);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738545, 'masons-tools', 'Mason''s Tools', 'tools', 25769803953);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738546, 'painters-supplies', 'Painter''s Supplies', 'tools', 25769803954);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738547, 'potters-tools', 'Potter''s Tools', 'tools', 25769803955);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738548, 'smiths-tools', 'Smith''s Tools', 'tools', 25769803956);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738549, 'tinkers-tools', 'Tinker''s Tools', 'tools', 25769803957);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738550, 'weavers-tools', 'Weaver''s Tools', 'tools', 25769803958);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738551, 'woodcarvers-tools', 'Woodcarver''s Tools', 'tools', 25769803959);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738552, 'dice-set', 'Dice Set', 'tools', 25769803960);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738553, 'playing-card-set', 'Playing Card Set', 'tools', 25769803961);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738554, 'bagpipes', 'Bagpipes', 'tools', 25769803962);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738555, 'drum', 'Drum', 'tools', 25769803963);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738556, 'dulcimer', 'Dulcimer', 'tools', 25769803964);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738557, 'flute', 'Flute', 'tools', 25769803965);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738558, 'lute', 'Lute', 'tools', 25769803966);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738559, 'lyre', 'Lyre', 'tools', 25769803967);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738560, 'horn', 'Horn', 'tools', 25769803968);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738561, 'pan-flute', 'Pan flute', 'tools', 25769803969);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738562, 'shawm', 'Shawm', 'tools', 25769803970);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738563, 'viol', 'Viol', 'tools', 25769803971);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738564, 'navigators-tools', 'Navigator''s Tools', 'tools', 25769803972);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738565, 'thieves-tools', 'Thieves'' Tools', 'tools', 25769803973);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738566, 'camel', 'Camel', 'mounts_and_vehicles', 25769803974);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738567, 'donkey', 'Donkey', 'mounts_and_vehicles', 25769803975);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738568, 'mule', 'Mule', 'mounts_and_vehicles', 25769803976);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738569, 'elephant', 'Elephant', 'mounts_and_vehicles', 25769803977);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738570, 'horse-draft', 'Horse, draft', 'mounts_and_vehicles', 25769803978);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738571, 'horse-riding', 'Horse, riding', 'mounts_and_vehicles', 25769803979);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738572, 'mastiff', 'Mastiff', 'mounts_and_vehicles', 25769803980);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738573, 'pony', 'Pony', 'mounts_and_vehicles', 25769803981);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738574, 'warhorse', 'Warhorse', 'mounts_and_vehicles', 25769803982);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738575, 'barding-padded', 'Barding: Padded', 'mounts_and_vehicles', 25769803983);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738576, 'barding-leather', 'Barding: Leather', 'mounts_and_vehicles', 25769803984);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738577, 'barding-studded-leather', 'Barding: Studded Leather', 'mounts_and_vehicles', 25769803985);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738578, 'barding-hide', 'Barding: Hide', 'mounts_and_vehicles', 25769803986);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738579, 'barding-chain-shirt', 'Barding: Chain shirt', 'mounts_and_vehicles', 25769803987);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738580, 'barding-scale-mail', 'Barding: Scale mail', 'mounts_and_vehicles', 25769803988);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738581, 'barding-breastplate', 'Barding: Breastplate', 'mounts_and_vehicles', 25769803989);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738582, 'barding-half-plate', 'Barding: Half plate', 'mounts_and_vehicles', 25769803990);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738583, 'barding-ring-mail', 'Barding: Ring mail', 'mounts_and_vehicles', 25769803991);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738584, 'barding-chain-mail', 'Barding: Chain mail', 'mounts_and_vehicles', 25769803992);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738585, 'barding-splint', 'Barding: Splint', 'mounts_and_vehicles', 25769803993);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738586, 'barding-plate', 'Barding: Plate', 'mounts_and_vehicles', 25769803994);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738587, 'bit-and-bridle', 'Bit and bridle', 'mounts_and_vehicles', 25769803995);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738588, 'carriage', 'Carriage', 'mounts_and_vehicles', 25769803996);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738589, 'cart', 'Cart', 'mounts_and_vehicles', 25769803997);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738590, 'chariot', 'Chariot', 'mounts_and_vehicles', 25769803998);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738591, 'animal-feed-1-day', 'Animal Feed (1 day)', 'mounts_and_vehicles', 25769803999);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738592, 'saddle-exotic', 'Saddle, Exotic', 'mounts_and_vehicles', 25769804000);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738593, 'saddle-military', 'Saddle, Military', 'mounts_and_vehicles', 25769804001);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738594, 'saddle-pack', 'Saddle, Pack', 'mounts_and_vehicles', 25769804002);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738595, 'saddle-riding', 'Saddle, Riding', 'mounts_and_vehicles', 25769804003);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738596, 'saddlebags', 'Saddlebags', 'mounts_and_vehicles', 25769804004);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738597, 'sled', 'Sled', 'mounts_and_vehicles', 25769804005);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738598, 'stabling-1-day', 'Stabling (1 day)', 'mounts_and_vehicles', 25769804006);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738599, 'wagon', 'Wagon', 'mounts_and_vehicles', 25769804007);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738600, 'galley', 'Galley', 'mounts_and_vehicles', 25769804008);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738601, 'keelboat', 'Keelboat', 'mounts_and_vehicles', 25769804009);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738602, 'longship', 'Longship', 'mounts_and_vehicles', 25769804010);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738603, 'rowboat', 'Rowboat', 'mounts_and_vehicles', 25769804011);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738604, 'sailing-ship', 'Sailing ship', 'mounts_and_vehicles', 25769804012);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_cost) VALUES (34359738605, 'warship', 'Warship', 'mounts_and_vehicles', 25769804013);

-- Table: gears
CREATE TABLE IF NOT EXISTS gears (
    id            INTEGER NOT NULL
                          PRIMARY KEY AUTOINCREMENT,
    indx          TEXT    NOT NULL,
    name          TEXT    NOT NULL,
    gear_category TEXT    NOT NULL
                          DEFAULT 'other',
    desc          JSON    NOT NULL,
    quantity      INTEGER,
    equipment_id  INTEGER NOT NULL,
    CONSTRAINT gears_equipment_gear FOREIGN KEY (
        equipment_id
    )
    REFERENCES equipment (id) ON DELETE NO ACTION
);

INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705665, 'abacus', 'Abacus', 'standard_gear', X'6E756C6C', 0, 34359738419);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705666, 'acid-vial', 'Acid (vial)', 'standard_gear', X'6E756C6C', 0, 34359738420);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705667, 'alchemists-fire-flask', 'Alchemist''s fire (flask)', 'standard_gear', X'6E756C6C', 0, 34359738421);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705668, 'alms-box', 'Alms box', 'standard_gear', X'6E756C6C', 0, 34359738422);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705669, 'arrow', 'Arrow', 'ammunition', X'6E756C6C', 20, 34359738423);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705670, 'block-of-incense', 'Block of incense', 'standard_gear', X'6E756C6C', 0, 34359738424);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705671, 'blowgun-needle', 'Blowgun needle', 'ammunition', X'6E756C6C', 50, 34359738425);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705672, 'censer', 'Censer', 'standard_gear', X'6E756C6C', 0, 34359738426);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705673, 'crossbow-bolt', 'Crossbow bolt', 'ammunition', X'6E756C6C', 20, 34359738427);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705674, 'sling-bullet', 'Sling bullet', 'ammunition', X'6E756C6C', 20, 34359738428);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705675, 'amulet', 'Amulet', 'holy_symbols', X'6E756C6C', 0, 34359738429);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705676, 'antitoxin-vial', 'Antitoxin (vial)', 'standard_gear', X'6E756C6C', 0, 34359738430);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705677, 'crystal', 'Crystal', 'arcane_foci', X'6E756C6C', 0, 34359738431);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705678, 'orb', 'Orb', 'arcane_foci', X'6E756C6C', 0, 34359738432);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705679, 'rod', 'Rod', 'arcane_foci', X'6E756C6C', 0, 34359738433);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705680, 'staff', 'Staff', 'arcane_foci', X'6E756C6C', 0, 34359738434);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705681, 'wand', 'Wand', 'arcane_foci', X'6E756C6C', 0, 34359738435);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705682, 'backpack', 'Backpack', 'standard_gear', X'6E756C6C', 0, 34359738436);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705683, 'ball-bearings-bag-of-1000', 'Ball bearings (bag of 1,000)', 'standard_gear', X'6E756C6C', 0, 34359738437);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705684, 'barrel', 'Barrel', 'standard_gear', X'6E756C6C', 0, 34359738438);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705685, 'basket', 'Basket', 'standard_gear', X'6E756C6C', 0, 34359738439);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705686, 'bedroll', 'Bedroll', 'standard_gear', X'6E756C6C', 0, 34359738440);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705687, 'bell', 'Bell', 'standard_gear', X'6E756C6C', 0, 34359738441);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705688, 'blanket', 'Blanket', 'standard_gear', X'6E756C6C', 0, 34359738442);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705689, 'block-and-tackle', 'Block and tackle', 'standard_gear', X'6E756C6C', 0, 34359738443);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705690, 'book', 'Book', 'standard_gear', X'6E756C6C', 0, 34359738444);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705691, 'bottle-glass', 'Bottle, glass', 'standard_gear', X'6E756C6C', 0, 34359738445);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705692, 'bucket', 'Bucket', 'standard_gear', X'6E756C6C', 0, 34359738446);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705693, 'caltrops', 'Caltrops', 'standard_gear', X'6E756C6C', 0, 34359738447);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705694, 'candle', 'Candle', 'standard_gear', X'6E756C6C', 0, 34359738448);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705695, 'case-crossbow-bolt', 'Case, crossbow bolt', 'standard_gear', X'6E756C6C', 0, 34359738449);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705696, 'case-map-or-scroll', 'Case, map or scroll', 'standard_gear', X'6E756C6C', 0, 34359738450);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705697, 'chain-10-feet', 'Chain (10 feet)', 'standard_gear', X'6E756C6C', 0, 34359738451);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705698, 'chalk-1-piece', 'Chalk (1 piece)', 'standard_gear', X'6E756C6C', 0, 34359738452);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705699, 'chest', 'Chest', 'standard_gear', X'6E756C6C', 0, 34359738453);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705700, 'clothes-common', 'Clothes, common', 'standard_gear', X'6E756C6C', 0, 34359738454);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705701, 'clothes-costume', 'Clothes, costume', 'standard_gear', X'6E756C6C', 0, 34359738455);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705702, 'clothes-fine', 'Clothes, fine', 'standard_gear', X'6E756C6C', 0, 34359738456);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705703, 'clothes-travelers', 'Clothes, traveler''s', 'standard_gear', X'6E756C6C', 0, 34359738457);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705704, 'component-pouch', 'Component pouch', 'standard_gear', X'6E756C6C', 0, 34359738458);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705705, 'crowbar', 'Crowbar', 'standard_gear', X'6E756C6C', 0, 34359738459);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705706, 'sprig-of-mistletoe', 'Sprig of mistletoe', 'druidic_foci', X'6E756C6C', 0, 34359738460);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705707, 'totem', 'Totem', 'druidic_foci', X'6E756C6C', 0, 34359738461);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705708, 'wooden-staff', 'Wooden staff', 'druidic_foci', X'6E756C6C', 0, 34359738462);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705709, 'yew-wand', 'Yew wand', 'druidic_foci', X'6E756C6C', 0, 34359738463);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705710, 'emblem', 'Emblem', 'holy_symbols', X'6E756C6C', 0, 34359738464);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705711, 'fishing-tackle', 'Fishing tackle', 'standard_gear', X'6E756C6C', 0, 34359738465);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705712, 'flask-or-tankard', 'Flask or tankard', 'standard_gear', X'6E756C6C', 0, 34359738466);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705713, 'grappling-hook', 'Grappling hook', 'standard_gear', X'6E756C6C', 0, 34359738467);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705714, 'hammer', 'Hammer', 'standard_gear', X'6E756C6C', 0, 34359738468);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705715, 'hammer-sledge', 'Hammer, sledge', 'standard_gear', X'6E756C6C', 0, 34359738469);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705716, 'holy-water-flask', 'Holy water (flask)', 'standard_gear', X'6E756C6C', 0, 34359738470);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705717, 'hourglass', 'Hourglass', 'standard_gear', X'6E756C6C', 0, 34359738471);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705718, 'hunting-trap', 'Hunting trap', 'standard_gear', X'6E756C6C', 0, 34359738472);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705719, 'ink-1-ounce-bottle', 'Ink (1 ounce bottle)', 'standard_gear', X'6E756C6C', 0, 34359738473);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705720, 'ink-pen', 'Ink pen', 'standard_gear', X'6E756C6C', 0, 34359738474);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705721, 'jug-or-pitcher', 'Jug or pitcher', 'standard_gear', X'6E756C6C', 0, 34359738475);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705722, 'climbers-kit', 'Climber''s Kit', 'kits', X'6E756C6C', 0, 34359738476);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705723, 'disguise-kit', 'Disguise Kit', 'kits', X'6E756C6C', 0, 34359738477);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705724, 'forgery-kit', 'Forgery Kit', 'kits', X'6E756C6C', 0, 34359738478);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705725, 'herbalism-kit', 'Herbalism Kit', 'kits', X'6E756C6C', 0, 34359738479);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705726, 'healers-kit', 'Healer''s Kit', 'kits', X'6E756C6C', 0, 34359738480);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705727, 'mess-kit', 'Mess Kit', 'kits', X'6E756C6C', 0, 34359738481);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705728, 'poisoners-kit', 'Poisoner''s Kit', 'kits', X'6E756C6C', 0, 34359738482);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705729, 'ladder-10-foot', 'Ladder (10-foot)', 'standard_gear', X'6E756C6C', 0, 34359738483);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705730, 'lamp', 'Lamp', 'standard_gear', X'6E756C6C', 0, 34359738484);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705731, 'lantern-bullseye', 'Lantern, bullseye', 'standard_gear', X'6E756C6C', 0, 34359738485);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705732, 'lantern-hooded', 'Lantern, hooded', 'standard_gear', X'6E756C6C', 0, 34359738486);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705733, 'little-bag-of-sand', 'Little bag of sand', 'standard_gear', X'6E756C6C', 0, 34359738487);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705734, 'lock', 'Lock', 'standard_gear', X'6E756C6C', 0, 34359738488);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705735, 'magnifying-glass', 'Magnifying glass', 'standard_gear', X'6E756C6C', 0, 34359738489);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705736, 'manacles', 'Manacles', 'standard_gear', X'6E756C6C', 0, 34359738490);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705737, 'mirror-steel', 'Mirror, steel', 'standard_gear', X'6E756C6C', 0, 34359738491);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705738, 'oil-flask', 'Oil (flask)', 'standard_gear', X'6E756C6C', 0, 34359738492);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705739, 'paper-one-sheet', 'Paper (one sheet)', 'standard_gear', X'6E756C6C', 0, 34359738493);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705740, 'parchment-one-sheet', 'Parchment (one sheet)', 'standard_gear', X'6E756C6C', 0, 34359738494);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705741, 'perfume-vial', 'Perfume (vial)', 'standard_gear', X'6E756C6C', 0, 34359738495);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705742, 'pick-miners', 'Pick, miner''s', 'standard_gear', X'6E756C6C', 0, 34359738496);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705743, 'piton', 'Piton', 'standard_gear', X'6E756C6C', 0, 34359738497);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705744, 'poison-basic-vial', 'Poison, basic (vial)', 'standard_gear', X'6E756C6C', 0, 34359738498);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705745, 'pole-10-foot', 'Pole (10-foot)', 'standard_gear', X'6E756C6C', 0, 34359738499);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705746, 'pot-iron', 'Pot, iron', 'standard_gear', X'6E756C6C', 0, 34359738500);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705747, 'pouch', 'Pouch', 'standard_gear', X'6E756C6C', 0, 34359738501);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705748, 'quiver', 'Quiver', 'standard_gear', X'6E756C6C', 0, 34359738502);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705749, 'ram-portable', 'Ram, portable', 'standard_gear', X'6E756C6C', 0, 34359738503);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705750, 'rations-1-day', 'Rations (1 day)', 'standard_gear', X'6E756C6C', 0, 34359738504);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705751, 'reliquary', 'Reliquary', 'holy_symbols', X'6E756C6C', 0, 34359738505);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705752, 'robes', 'Robes', 'standard_gear', X'6E756C6C', 0, 34359738506);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705753, 'rope-hempen-50-feet', 'Rope, hempen (50 feet)', 'standard_gear', X'6E756C6C', 0, 34359738507);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705754, 'rope-silk-50-feet', 'Rope, silk (50 feet)', 'standard_gear', X'6E756C6C', 0, 34359738508);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705755, 'sack', 'Sack', 'standard_gear', X'6E756C6C', 0, 34359738509);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705756, 'scale-merchants', 'Scale, merchant''s', 'standard_gear', X'6E756C6C', 0, 34359738510);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705757, 'sealing-wax', 'Sealing wax', 'standard_gear', X'6E756C6C', 0, 34359738511);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705758, 'shovel', 'Shovel', 'standard_gear', X'6E756C6C', 0, 34359738512);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705759, 'signal-whistle', 'Signal whistle', 'standard_gear', X'6E756C6C', 0, 34359738513);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705760, 'signet-ring', 'Signet ring', 'standard_gear', X'6E756C6C', 0, 34359738514);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705761, 'small-knife', 'Small knife', 'standard_gear', X'6E756C6C', 0, 34359738515);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705762, 'soap', 'Soap', 'standard_gear', X'6E756C6C', 0, 34359738516);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705763, 'spellbook', 'Spellbook', 'standard_gear', X'6E756C6C', 0, 34359738517);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705764, 'spike-iron', 'Spike, iron', 'standard_gear', X'6E756C6C', 0, 34359738518);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705765, 'spyglass', 'Spyglass', 'standard_gear', X'6E756C6C', 0, 34359738519);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705766, 'string-10-feet', 'String (10 feet)', 'standard_gear', X'6E756C6C', 0, 34359738520);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705767, 'tent-two-person', 'Tent, two-person', 'standard_gear', X'6E756C6C', 0, 34359738521);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705768, 'tinderbox', 'Tinderbox', 'standard_gear', X'6E756C6C', 0, 34359738522);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705769, 'torch', 'Torch', 'standard_gear', X'6E756C6C', 0, 34359738523);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705770, 'vestments', 'Vestments', 'standard_gear', X'6E756C6C', 0, 34359738524);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705771, 'vial', 'Vial', 'standard_gear', X'6E756C6C', 0, 34359738525);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705772, 'waterskin', 'Waterskin', 'standard_gear', X'6E756C6C', 0, 34359738526);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705773, 'whetstone', 'Whetstone', 'standard_gear', X'6E756C6C', 0, 34359738527);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705774, 'burglars-pack', 'Burglar''s Pack', 'equipment_packs', X'6E756C6C', 0, 34359738528);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705775, 'diplomats-pack', 'Diplomat''s Pack', 'equipment_packs', X'6E756C6C', 0, 34359738529);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705776, 'dungeoneers-pack', 'Dungeoneer''s Pack', 'equipment_packs', X'6E756C6C', 0, 34359738530);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705777, 'entertainers-pack', 'Entertainer''s Pack', 'equipment_packs', X'6E756C6C', 0, 34359738531);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705778, 'explorers-pack', 'Explorer''s Pack', 'equipment_packs', X'6E756C6C', 0, 34359738532);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705779, 'priests-pack', 'Priest''s Pack', 'equipment_packs', X'6E756C6C', 0, 34359738533);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity, equipment_id) VALUES (38654705780, 'scholars-pack', 'Scholar''s Pack', 'equipment_packs', X'6E756C6C', 0, 34359738534);

-- Table: languages
CREATE TABLE IF NOT EXISTS languages (
    id            INTEGER NOT NULL
                          PRIMARY KEY AUTOINCREMENT,
    indx          TEXT    NOT NULL,
    name          TEXT    NOT NULL,
    desc          TEXT    NOT NULL,
    language_type TEXT    NOT NULL
                          DEFAULT 'STANDARD',
    script        TEXT
                          DEFAULT 'Common'
);

INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (42949672961, 'common', 'Common', '', 'STANDARD', 'Common');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (42949672962, 'dwarvish', 'Dwarvish', 'Dwarvish is full of hard consonants and guttural sounds.', 'STANDARD', 'Dwarvish');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (42949672963, 'elvish', 'Elvish', 'Elvish is fluid, with subtle intonations and intricate grammar. Elven literature is rich and varied, and their songs and poems are famous among other races. Many bards learn their language so they can add Elvish ballads to their repertoires.', 'STANDARD', 'Elvish');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (42949672964, 'giant', 'Giant', '', 'STANDARD', 'Dwarvish');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (42949672965, 'gnomish', 'Gnomish', 'The Gnomish language, which uses the Dwarvish script, is renowned for its technical treatises and its catalogs of knowledge about the natural world.', 'STANDARD', 'Dwarvish');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (42949672966, 'goblin', 'Goblin', '', 'STANDARD', 'Dwarvish');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (42949672967, 'halfling', 'Halfling', 'The Halfling language isn''t secret, but halflings are loath to share it with others. They write very little, so they don''t have a rich body of literature. Their oral tradition, however, is very strong.', 'STANDARD', 'Common');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (42949672968, 'orc', 'Orc', 'Orc is a harsh, grating language with hard consonants. It has no script of its own but is written in the Dwarvish script.', 'STANDARD', 'Dwarvish');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (42949672969, 'abyssal', 'Abyssal', '', 'EXOTIC', 'Infernal');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (42949672970, 'celestial', 'Celestial', '', 'EXOTIC', 'Celestial');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (42949672971, 'draconic', 'Draconic', 'Draconic is thought to be one of the oldest languages and is often used in the study of magic. The language sounds harsh to most other creatures and includes numerous hard consonants and sibilants.', 'EXOTIC', 'Draconic');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (42949672972, 'deep-speech', 'Deep Speech', 'An eerie tongue, Deep Speech is a clicking language with unsettling vocalizations that can almost sound like noise. Vibrating, creaking, and trembling sounds complement this mode of communication, making it hard for humanoid vocal cords to mimic.', 'EXOTIC', 'Other');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (42949672973, 'infernal', 'Infernal', '', 'EXOTIC', 'Infernal');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (42949672974, 'primordial', 'Primordial', '', 'EXOTIC', 'Dwarvish');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (42949672975, 'sylvan', 'Sylvan', '', 'EXOTIC', 'Elvish');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (42949672976, 'undercommon', 'Undercommon', '', 'EXOTIC', 'Elvish');

-- Table: magic_schools
CREATE TABLE IF NOT EXISTS magic_schools (
    id   INTEGER NOT NULL
                 PRIMARY KEY AUTOINCREMENT,
    indx TEXT    NOT NULL,
    name TEXT    NOT NULL,
    desc TEXT    NOT NULL
);


-- Table: proficiencies
CREATE TABLE IF NOT EXISTS proficiencies (
    id                       INTEGER NOT NULL
                                     PRIMARY KEY AUTOINCREMENT,
    indx                     TEXT    NOT NULL,
    name                     TEXT    NOT NULL,
    proficiency_category     TEXT    NOT NULL,
    proficiency_skill        INTEGER,
    proficiency_equipment    INTEGER,
    proficiency_saving_throw INTEGER,
    CONSTRAINT proficiencies_skills_skill FOREIGN KEY (
        proficiency_skill
    )
    REFERENCES skills (id) ON DELETE SET NULL,
    CONSTRAINT proficiencies_equipment_equipment FOREIGN KEY (
        proficiency_equipment
    )
    REFERENCES equipment (id) ON DELETE SET NULL,
    CONSTRAINT proficiencies_ability_scores_saving_throw FOREIGN KEY (
        proficiency_saving_throw
    )
    REFERENCES ability_scores (id) ON DELETE SET NULL
);

INSERT INTO proficiencies (id, indx, name, proficiency_category, proficiency_skill, proficiency_equipment, proficiency_saving_throw) VALUES (51539607553, 'light-armor', 'Light Armor', 'equipment_categories', NULL, NULL, NULL);
INSERT INTO proficiencies (id, indx, name, proficiency_category, proficiency_skill, proficiency_equipment, proficiency_saving_throw) VALUES (51539607554, 'medium-armor', 'Medium Armor', 'equipment_categories', NULL, NULL, NULL);
INSERT INTO proficiencies (id, indx, name, proficiency_category, proficiency_skill, proficiency_equipment, proficiency_saving_throw) VALUES (51539607555, 'heavy-armor', 'Heavy Armor', 'equipment_categories', NULL, NULL, NULL);
INSERT INTO proficiencies (id, indx, name, proficiency_category, proficiency_skill, proficiency_equipment, proficiency_saving_throw) VALUES (51539607556, 'all-armor', 'All armor', 'equipment_categories', NULL, NULL, NULL);
INSERT INTO proficiencies (id, indx, name, proficiency_category, proficiency_skill, proficiency_equipment, proficiency_saving_throw) VALUES (51539607557, 'padded-armor', 'Padded Armor', 'equipment', NULL, NULL, NULL);

-- Table: race_languages
CREATE TABLE IF NOT EXISTS race_languages (
    race_id     INTEGER NOT NULL,
    language_id INTEGER NOT NULL,
    PRIMARY KEY (
        race_id,
        language_id
    ),
    CONSTRAINT race_languages_race_id FOREIGN KEY (
        race_id
    )
    REFERENCES races (id) ON DELETE CASCADE,
    CONSTRAINT race_languages_language_id FOREIGN KEY (
        language_id
    )
    REFERENCES languages (id) ON DELETE CASCADE
);

INSERT INTO race_languages (race_id, language_id) VALUES (55834574849, 42949672961);
INSERT INTO race_languages (race_id, language_id) VALUES (55834574849, 42949672962);
INSERT INTO race_languages (race_id, language_id) VALUES (55834574850, 42949672961);
INSERT INTO race_languages (race_id, language_id) VALUES (55834574850, 42949672963);
INSERT INTO race_languages (race_id, language_id) VALUES (55834574851, 42949672961);
INSERT INTO race_languages (race_id, language_id) VALUES (55834574851, 42949672967);
INSERT INTO race_languages (race_id, language_id) VALUES (55834574852, 42949672961);
INSERT INTO race_languages (race_id, language_id) VALUES (55834574853, 42949672961);
INSERT INTO race_languages (race_id, language_id) VALUES (55834574853, 42949672971);
INSERT INTO race_languages (race_id, language_id) VALUES (55834574854, 42949672961);
INSERT INTO race_languages (race_id, language_id) VALUES (55834574854, 42949672965);
INSERT INTO race_languages (race_id, language_id) VALUES (55834574855, 42949672961);
INSERT INTO race_languages (race_id, language_id) VALUES (55834574855, 42949672963);
INSERT INTO race_languages (race_id, language_id) VALUES (55834574856, 42949672961);
INSERT INTO race_languages (race_id, language_id) VALUES (55834574856, 42949672968);
INSERT INTO race_languages (race_id, language_id) VALUES (55834574857, 42949672961);
INSERT INTO race_languages (race_id, language_id) VALUES (55834574857, 42949672973);

-- Table: race_proficiencies
CREATE TABLE IF NOT EXISTS race_proficiencies (
    race_id        INTEGER NOT NULL,
    proficiency_id INTEGER NOT NULL,
    PRIMARY KEY (
        race_id,
        proficiency_id
    ),
    CONSTRAINT race_proficiencies_race_id FOREIGN KEY (
        race_id
    )
    REFERENCES races (id) ON DELETE CASCADE,
    CONSTRAINT race_proficiencies_proficiency_id FOREIGN KEY (
        proficiency_id
    )
    REFERENCES proficiencies (id) ON DELETE CASCADE
);


-- Table: race_traits
CREATE TABLE IF NOT EXISTS race_traits (
    race_id  INTEGER NOT NULL,
    trait_id INTEGER NOT NULL,
    PRIMARY KEY (
        race_id,
        trait_id
    ),
    CONSTRAINT race_traits_race_id FOREIGN KEY (
        race_id
    )
    REFERENCES races (id) ON DELETE CASCADE,
    CONSTRAINT race_traits_trait_id FOREIGN KEY (
        trait_id
    )
    REFERENCES traits (id) ON DELETE CASCADE
);


-- Table: races
CREATE TABLE IF NOT EXISTS races (
    id               INTEGER NOT NULL
                             PRIMARY KEY AUTOINCREMENT,
    indx             TEXT    NOT NULL,
    name             TEXT    NOT NULL,
    alignment        TEXT    NOT NULL,
    age              TEXT    NOT NULL,
    size             TEXT    NOT NULL,
    size_description TEXT    NOT NULL,
    language_desc    TEXT    NOT NULL,
    speed            INTEGER NOT NULL
);

INSERT INTO races (id, indx, name, alignment, age, size, size_description, language_desc, speed) VALUES (55834574849, 'dwarf', 'Dwarf', 'Most dwarves are lawful, believing firmly in the benefits of a well-ordered society. They tend toward good as well, with a strong sense of fair play and a belief that everyone deserves to share in the benefits of a just order.', 'Dwarves mature at the same rate as humans, but they''re considered young until they reach the age of 50. On average, they live about 350 years.', 'Medium', 'Dwarves stand between 4 and 5 feet tall and average about 150 pounds. Your size is Medium.', 'You can speak, read, and write Common and Dwarvish. Dwarvish is full of hard consonants and guttural sounds, and those characteristics spill over into whatever other language a dwarf might speak.', 25);
INSERT INTO races (id, indx, name, alignment, age, size, size_description, language_desc, speed) VALUES (55834574850, 'elf', 'Elf', 'Elves love freedom, variety, and self-expression, so they lean strongly toward the gentler aspects of chaos. They value and protect others'' freedom as well as their own, and they are more often good than not.', 'Although elves reach physical maturity at about the same age as humans, the elven understanding of adulthood goes beyond physical growth to encompass worldly experience. An elf typically claims adulthood and an adult name around the age of 100 and can live to be 750 years old.', 'Medium', 'Elves range from under 5 to over 6 feet tall and have slender builds. Your size is Medium.', 'You can speak, read, and write Common and Elvish. Elvish is fluid, with subtle intonations and intricate grammar. Elven literature is rich and varied, and their songs and poems are famous among other races. Many bards learn their language so they can add Elvish ballads to their repertoires.', 30);
INSERT INTO races (id, indx, name, alignment, age, size, size_description, language_desc, speed) VALUES (55834574851, 'halfling', 'Halfling', 'Most halflings are lawful good. As a rule, they are good-hearted and kind, hate to see others in pain, and have no tolerance for oppression. They are also very orderly and traditional, leaning heavily on the support of their community and the comfort of their old ways.', 'A halfling reaches adulthood at the age of 20 and generally lives into the middle of his or her second century.', 'Small', 'Halflings average about 3 feet tall and weigh about 40 pounds. Your size is Small.', 'You can speak, read, and write Common and Halfling. The Halfling language isn''t secret, but halflings are loath to share it with others. They write very little, so they don''t have a rich body of literature. Their oral tradition, however, is very strong. Almost all halflings speak Common to converse with the people in whose lands they dwell or through which they are traveling.', 25);
INSERT INTO races (id, indx, name, alignment, age, size, size_description, language_desc, speed) VALUES (55834574852, 'human', 'Human', 'Humans tend toward no particular alignment. The best and the worst are found among them.', 'Humans reach adulthood in their late teens and live less than a century.', 'Medium', 'Humans vary widely in height and build, from barely 5 feet to well over 6 feet tall. Regardless of your position in that range, your size is Medium.', 'You can speak, read, and write Common and one extra language of your choice. Humans typically learn the languages of other peoples they deal with, including obscure dialects. They are fond of sprinkling their speech with words borrowed from other tongues: Orc curses, Elvish musical expressions, Dwarvish military phrases, and so on.', 30);
INSERT INTO races (id, indx, name, alignment, age, size, size_description, language_desc, speed) VALUES (55834574853, 'dragonborn', 'Dragonborn', 'Dragonborn tend to extremes, making a conscious choice for one side or the other in the cosmic war between good and evil. Most dragonborn are good, but those who side with evil can be terrible villains.', 'Young dragonborn grow quickly. They walk hours after hatching, attain the size and development of a 10-year-old human child by the age of 3, and reach adulthood by 15. They live to be around 80.', 'Medium', 'Dragonborn are taller and heavier than humans, standing well over 6 feet tall and averaging almost 250 pounds. Your size is Medium.', 'You can speak, read, and write Common and Draconic. Draconic is thought to be one of the oldest languages and is often used in the study of magic. The language sounds harsh to most other creatures and includes numerous hard consonants and sibilants.', 30);
INSERT INTO races (id, indx, name, alignment, age, size, size_description, language_desc, speed) VALUES (55834574854, 'gnome', 'Gnome', 'Gnomes are most often good. Those who tend toward law are sages, engineers, researchers, scholars, investigators, or inventors. Those who tend toward chaos are minstrels, tricksters, wanderers, or fanciful jewelers. Gnomes are good-hearted, and even the tricksters among them are more playful than vicious.', 'Gnomes mature at the same rate humans do, and most are expected to settle down into an adult life by around age 40. They can live 350 to almost 500 years.', 'Small', 'Gnomes are between 3 and 4 feet tall and average about 40 pounds. Your size is Small.', 'You can speak, read, and write Common and Gnomish. The Gnomish language, which uses the Dwarvish script, is renowned for its technical treatises and its catalogs of knowledge about the natural world.', 25);
INSERT INTO races (id, indx, name, alignment, age, size, size_description, language_desc, speed) VALUES (55834574855, 'half-elf', 'Half-Elf', 'Half-elves share the chaotic bent of their elven heritage. They value both personal freedom and creative expression, demonstrating neither love of leaders nor desire for followers. They chafe at rules, resent others'' demands, and sometimes prove unreliable, or at least unpredictable.', 'Half-elves mature at the same rate humans do and reach adulthood around the age of 20. They live much longer than humans, however, often exceeding 180 years.', 'Medium', 'Half-elves are about the same size as humans, ranging from 5 to 6 feet tall. Your size is Medium.', 'You can speak, read, and write Common, Elvish, and one extra language of your choice.', 30);
INSERT INTO races (id, indx, name, alignment, age, size, size_description, language_desc, speed) VALUES (55834574856, 'half-orc', 'Half-Orc', 'Half-orcs inherit a tendency toward chaos from their orc parents and are not strongly inclined toward good. Half-orcs raised among orcs and willing to live out their lives among them are usually evil.', 'Half-orcs mature a little faster than humans, reaching adulthood around age 14. They age noticeably faster and rarely live longer than 75 years.', 'Medium', 'Half-orcs are somewhat larger and bulkier than humans, and they range from 5 to well over 6 feet tall. Your size is Medium.', 'You can speak, read, and write Common and Orc. Orc is a harsh, grating language with hard consonants. It has no script of its own but is written in the Dwarvish script.', 30);
INSERT INTO races (id, indx, name, alignment, age, size, size_description, language_desc, speed) VALUES (55834574857, 'tiefling', 'Tiefling', 'Tieflings might not have an innate tendency toward evil, but many of them end up there. Evil or not, an independent nature inclines many tieflings toward a chaotic alignment.', 'Tieflings mature at the same rate as humans but live a few years longer.', 'Medium', 'Tieflings are about the same size and build as humans. Your size is Medium.', 'You can speak, read, and write Common and Infernal.', 30);

-- Table: rule_rule_sections
CREATE TABLE IF NOT EXISTS rule_rule_sections (
    rule_id         INTEGER NOT NULL,
    rule_section_id INTEGER NOT NULL,
    PRIMARY KEY (
        rule_id,
        rule_section_id
    ),
    CONSTRAINT rule_rule_sections_rule_id FOREIGN KEY (
        rule_id
    )
    REFERENCES rules (id) ON DELETE CASCADE,
    CONSTRAINT rule_rule_sections_rule_section_id FOREIGN KEY (
        rule_section_id
    )
    REFERENCES rule_sections (id) ON DELETE CASCADE
);


-- Table: rule_sections
CREATE TABLE IF NOT EXISTS rule_sections (
    id   INTEGER NOT NULL
                 PRIMARY KEY AUTOINCREMENT,
    indx TEXT    NOT NULL,
    name TEXT    NOT NULL,
    desc TEXT    NOT NULL
);


-- Table: rules
CREATE TABLE IF NOT EXISTS rules (
    id   INTEGER NOT NULL
                 PRIMARY KEY AUTOINCREMENT,
    indx TEXT    NOT NULL,
    name TEXT    NOT NULL,
    desc TEXT    NOT NULL
);


-- Table: skills
CREATE TABLE IF NOT EXISTS skills (
    id                  INTEGER NOT NULL
                                PRIMARY KEY AUTOINCREMENT,
    indx                TEXT    NOT NULL,
    name                TEXT    NOT NULL,
    desc                JSON    NOT NULL,
    skill_ability_score INTEGER,
    CONSTRAINT skills_ability_scores_ability_score FOREIGN KEY (
        skill_ability_score
    )
    REFERENCES ability_scores (id) ON DELETE SET NULL
);

INSERT INTO skills (id, indx, name, desc, skill_ability_score) VALUES (68719476737, 'acrobatics', 'Acrobatics', X'5B22596F75722044657874657269747920284163726F6261746963732920636865636B20636F7665727320796F757220617474656D707420746F2073746179206F6E20796F7572206665657420696E206120747269636B7920736974756174696F6E2C2073756368206173207768656E20796F7527726520747279696E6720746F2072756E206163726F73732061207368656574206F66206963652C2062616C616E6365206F6E2061207469676874726F70652C206F7220737461792075707269676874206F6E206120726F636B696E6720736869702773206465636B2E2054686520474D206D6967687420616C736F2063616C6C20666F7220612044657874657269747920284163726F6261746963732920636865636B20746F2073656520696620796F752063616E20706572666F726D206163726F6261746963207374756E74732C20696E636C7564696E672064697665732C20726F6C6C732C20736F6D65727361756C74732C20616E6420666C6970732E225D', 4294967298);
INSERT INTO skills (id, indx, name, desc, skill_ability_score) VALUES (68719476738, 'animal-handling', 'Animal Handling', X'5B225768656E20746865726520697320616E79207175657374696F6E207768657468657220796F752063616E2063616C6D20646F776E206120646F6D65737469636174656420616E696D616C2C206B6565702061206D6F756E742066726F6D2067657474696E672073706F6F6B65642C206F7220696E7475697420616E20616E696D616C277320696E74656E74696F6E732C2074686520474D206D696768742063616C6C20666F72206120576973646F6D2028416E696D616C2048616E646C696E672920636865636B2E20596F7520616C736F206D616B65206120576973646F6D2028416E696D616C2048616E646C696E672920636865636B20746F20636F6E74726F6C20796F7572206D6F756E74207768656E20796F7520617474656D70742061207269736B79206D616E65757665722E225D', 4294967301);
INSERT INTO skills (id, indx, name, desc, skill_ability_score) VALUES (68719476739, 'arcana', 'Arcana', X'5B22596F757220496E74656C6C6967656E63652028417263616E612920636865636B206D6561737572657320796F7572206162696C69747920746F20726563616C6C206C6F72652061626F7574207370656C6C732C206D61676963206974656D732C20656C6472697463682073796D626F6C732C206D61676963616C20747261646974696F6E732C2074686520706C616E6573206F66206578697374656E63652C20616E642074686520696E6861626974616E7473206F662074686F736520706C616E65732E225D', 4294967300);
INSERT INTO skills (id, indx, name, desc, skill_ability_score) VALUES (68719476740, 'athletics', 'Athletics', X'5B22596F757220537472656E67746820284174686C65746963732920636865636B20636F7665727320646966666963756C7420736974756174696F6E7320796F7520656E636F756E746572207768696C6520636C696D62696E672C206A756D70696E672C206F72207377696D6D696E672E225D', 4294967297);
INSERT INTO skills (id, indx, name, desc, skill_ability_score) VALUES (68719476741, 'deception', 'Deception', X'5B22596F7572204368617269736D612028446563657074696F6E2920636865636B2064657465726D696E6573207768657468657220796F752063616E20636F6E76696E63696E676C792068696465207468652074727574682C206569746865722076657262616C6C79206F72207468726F75676820796F757220616374696F6E732E205468697320646563657074696F6E2063616E20656E636F6D706173732065766572797468696E672066726F6D206D69736C656164696E67206F7468657273207468726F75676820616D6269677569747920746F2074656C6C696E67206F75747269676874206C6965732E205479706963616C20736974756174696F6E7320696E636C75646520747279696E6720746F20666173742D2074616C6B20612067756172642C20636F6E2061206D65726368616E742C206561726E206D6F6E6579207468726F7567682067616D626C696E672C207061737320796F757273656C66206F666620696E20612064697367756973652C2064756C6C20736F6D656F6E65277320737573706963696F6E7320776974682066616C7365206173737572616E6365732C206F72206D61696E7461696E20612073747261696768742066616365207768696C652074656C6C696E67206120626C6174616E74206C69652E225D', 4294967302);
INSERT INTO skills (id, indx, name, desc, skill_ability_score) VALUES (68719476742, 'history', 'History', X'5B22596F757220496E74656C6C6967656E63652028486973746F72792920636865636B206D6561737572657320796F7572206162696C69747920746F20726563616C6C206C6F72652061626F757420686973746F726963616C206576656E74732C206C6567656E646172792070656F706C652C20616E6369656E74206B696E67646F6D732C20706173742064697370757465732C20726563656E7420776172732C20616E64206C6F737420636976696C697A6174696F6E732E225D', 4294967300);
INSERT INTO skills (id, indx, name, desc, skill_ability_score) VALUES (68719476743, 'insight', 'Insight', X'5B22596F757220576973646F6D2028496E73696768742920636865636B2064656369646573207768657468657220796F752063616E2064657465726D696E6520746865207472756520696E74656E74696F6E73206F6620612063726561747572652C2073756368206173207768656E20736561726368696E67206F75742061206C6965206F722070726564696374696E6720736F6D656F6E652773206E657874206D6F76652E20446F696E6720736F20696E766F6C76657320676C65616E696E6720636C7565732066726F6D20626F6479206C616E67756167652C20737065656368206861626974732C20616E64206368616E67657320696E206D616E6E657269736D732E225D', 4294967301);
INSERT INTO skills (id, indx, name, desc, skill_ability_score) VALUES (68719476744, 'intimidation', 'Intimidation', X'5B225768656E20796F7520617474656D707420746F20696E666C75656E636520736F6D656F6E65207468726F756768206F7665727420746872656174732C20686F7374696C6520616374696F6E732C20616E6420706879736963616C2076696F6C656E63652C2074686520474D206D696768742061736B20796F7520746F206D616B652061204368617269736D612028496E74696D69646174696F6E2920636865636B2E204578616D706C657320696E636C75646520747279696E6720746F2070727920696E666F726D6174696F6E206F7574206F66206120707269736F6E65722C20636F6E76696E63696E672073747265657420746875677320746F206261636B20646F776E2066726F6D206120636F6E66726F6E746174696F6E2C206F72207573696E67207468652065646765206F6620612062726F6B656E20626F74746C6520746F20636F6E76696E6365206120736E656572696E672076697A69657220746F207265636F6E73696465722061206465636973696F6E2E225D', 4294967302);
INSERT INTO skills (id, indx, name, desc, skill_ability_score) VALUES (68719476745, 'investigation', 'Investigation', X'5B225768656E20796F75206C6F6F6B2061726F756E6420666F7220636C75657320616E64206D616B6520646564756374696F6E73206261736564206F6E2074686F736520636C7565732C20796F75206D616B6520616E20496E74656C6C6967656E63652028496E7665737469676174696F6E2920636865636B2E20596F75206D696768742064656475636520746865206C6F636174696F6E206F6620612068696464656E206F626A6563742C206469736365726E2066726F6D2074686520617070656172616E6365206F66206120776F756E642077686174206B696E64206F6620776561706F6E206465616C742069742C206F722064657465726D696E6520746865207765616B65737420706F696E7420696E20612074756E6E656C207468617420636F756C6420636175736520697420746F20636F6C6C617073652E20506F72696E67207468726F75676820616E6369656E74207363726F6C6C7320696E20736561726368206F6620612068696464656E20667261676D656E74206F66206B6E6F776C65646765206D6967687420616C736F2063616C6C20666F7220616E20496E74656C6C6967656E63652028496E7665737469676174696F6E2920636865636B2E225D', 4294967300);
INSERT INTO skills (id, indx, name, desc, skill_ability_score) VALUES (68719476746, 'medicine', 'Medicine', X'5B224120576973646F6D20284D65646963696E652920636865636B206C65747320796F752074727920746F2073746162696C697A652061206479696E6720636F6D70616E696F6E206F7220646961676E6F736520616E20696C6C6E6573732E225D', 4294967301);
INSERT INTO skills (id, indx, name, desc, skill_ability_score) VALUES (68719476747, 'nature', 'Nature', X'5B22596F757220496E74656C6C6967656E636520284E61747572652920636865636B206D6561737572657320796F7572206162696C69747920746F20726563616C6C206C6F72652061626F7574207465727261696E2C20706C616E747320616E6420616E696D616C732C2074686520776561746865722C20616E64206E61747572616C206379636C65732E225D', 4294967300);
INSERT INTO skills (id, indx, name, desc, skill_ability_score) VALUES (68719476748, 'perception', 'Perception', X'5B22596F757220576973646F6D202850657263657074696F6E2920636865636B206C65747320796F752073706F742C20686561722C206F72206F746865727769736520646574656374207468652070726573656E6365206F6620736F6D657468696E672E204974206D6561737572657320796F75722067656E6572616C2061776172656E657373206F6620796F757220737572726F756E64696E677320616E6420746865206B65656E6E657373206F6620796F75722073656E7365732E20466F72206578616D706C652C20796F75206D696768742074727920746F2068656172206120636F6E766572736174696F6E207468726F756768206120636C6F73656420646F6F722C20656176657364726F7020756E64657220616E206F70656E2077696E646F772C206F722068656172206D6F6E7374657273206D6F76696E6720737465616C7468696C7920696E2074686520666F726573742E204F7220796F75206D696768742074727920746F2073706F74207468696E6773207468617420617265206F62736375726564206F72206561737920746F206D6973732C2077686574686572207468657920617265206F726373206C79696E6720696E20616D62757368206F6E206120726F61642C20746875677320686964696E6720696E2074686520736861646F7773206F6620616E20616C6C65792C206F722063616E646C656C6967687420756E646572206120636C6F7365642073656372657420646F6F722E225D', 4294967301);
INSERT INTO skills (id, indx, name, desc, skill_ability_score) VALUES (68719476749, 'performance', 'Performance', X'5B22596F7572204368617269736D612028506572666F726D616E63652920636865636B2064657465726D696E657320686F772077656C6C20796F752063616E2064656C6967687420616E2061756469656E63652077697468206D757369632C2064616E63652C20616374696E672C2073746F727974656C6C696E672C206F7220736F6D65206F7468657220666F726D206F6620656E7465727461696E6D656E742E225D', 4294967302);
INSERT INTO skills (id, indx, name, desc, skill_ability_score) VALUES (68719476750, 'persuasion', 'Persuasion', X'5B225768656E20796F7520617474656D707420746F20696E666C75656E636520736F6D656F6E65206F7220612067726F7570206F662070656F706C65207769746820746163742C20736F6369616C206772616365732C206F7220676F6F64206E61747572652C2074686520474D206D696768742061736B20796F7520746F206D616B652061204368617269736D61202850657273756173696F6E2920636865636B2E205479706963616C6C792C20796F75207573652070657273756173696F6E207768656E20616374696E6720696E20676F6F642066616974682C20746F20666F7374657220667269656E6473686970732C206D616B6520636F726469616C2072657175657374732C206F7220657868696269742070726F706572206574697175657474652E204578616D706C6573206F662070657273756164696E67206F746865727320696E636C75646520636F6E76696E63696E672061206368616D6265726C61696E20746F206C657420796F75722070617274792073656520746865206B696E672C206E65676F74696174696E67207065616365206265747765656E2077617272696E67207472696265732C206F7220696E73706972696E6720612063726F7764206F6620746F776E73666F6C6B2E225D', 4294967302);
INSERT INTO skills (id, indx, name, desc, skill_ability_score) VALUES (68719476751, 'religion', 'Religion', X'5B22596F757220496E74656C6C6967656E6365202852656C6967696F6E2920636865636B206D6561737572657320796F7572206162696C69747920746F20726563616C6C206C6F72652061626F757420646569746965732C20726974657320616E6420707261796572732C2072656C6967696F75732068696572617263686965732C20686F6C792073796D626F6C732C20616E642074686520707261637469636573206F66207365637265742063756C74732E225D', 4294967300);
INSERT INTO skills (id, indx, name, desc, skill_ability_score) VALUES (68719476752, 'sleight-of-hand', 'Sleight of Hand', X'5B225768656E6576657220796F7520617474656D707420616E20616374206F66206C6567657264656D61696E206F72206D616E75616C20747269636B6572792C207375636820617320706C616E74696E6720736F6D657468696E67206F6E20736F6D656F6E6520656C7365206F7220636F6E6365616C696E6720616E206F626A656374206F6E20796F757220706572736F6E2C206D616B652061204465787465726974792028536C6569676874206F662048616E642920636865636B2E2054686520474D206D6967687420616C736F2063616C6C20666F722061204465787465726974792028536C6569676874206F662048616E642920636865636B20746F2064657465726D696E65207768657468657220796F752063616E206C696674206120636F696E207075727365206F666620616E6F7468657220706572736F6E206F7220736C697020736F6D657468696E67206F7574206F6620616E6F7468657220706572736F6E277320706F636B65742E225D', 4294967298);
INSERT INTO skills (id, indx, name, desc, skill_ability_score) VALUES (68719476753, 'stealth', 'Stealth', X'5B224D616B652061204465787465726974792028537465616C74682920636865636B207768656E20796F7520617474656D707420746F20636F6E6365616C20796F757273656C662066726F6D20656E656D6965732C20736C696E6B2070617374206775617264732C20736C6970206177617920776974686F7574206265696E67206E6F74696365642C206F7220736E65616B207570206F6E20736F6D656F6E6520776974686F7574206265696E67207365656E206F722068656172642E225D', 4294967298);
INSERT INTO skills (id, indx, name, desc, skill_ability_score) VALUES (68719476754, 'survival', 'Survival', X'5B2254686520474D206D696768742061736B20796F7520746F206D616B65206120576973646F6D2028537572766976616C2920636865636B20746F20666F6C6C6F7720747261636B732C2068756E742077696C642067616D652C20677569646520796F75722067726F7570207468726F7567682066726F7A656E2077617374656C616E64732C206964656E74696679207369676E732074686174206F776C6265617273206C697665206E65617262792C20707265646963742074686520776561746865722C206F722061766F696420717569636B73616E6420616E64206F74686572206E61747572616C2068617A617264732E225D', 4294967301);

-- Table: subrace_proficiencies
CREATE TABLE IF NOT EXISTS subrace_proficiencies (
    subrace_id     INTEGER NOT NULL,
    proficiency_id INTEGER NOT NULL,
    PRIMARY KEY (
        subrace_id,
        proficiency_id
    ),
    CONSTRAINT subrace_proficiencies_subrace_id FOREIGN KEY (
        subrace_id
    )
    REFERENCES subraces (id) ON DELETE CASCADE,
    CONSTRAINT subrace_proficiencies_proficiency_id FOREIGN KEY (
        proficiency_id
    )
    REFERENCES proficiencies (id) ON DELETE CASCADE
);


-- Table: subrace_traits
CREATE TABLE IF NOT EXISTS subrace_traits (
    subrace_id INTEGER NOT NULL,
    trait_id   INTEGER NOT NULL,
    PRIMARY KEY (
        subrace_id,
        trait_id
    ),
    CONSTRAINT subrace_traits_subrace_id FOREIGN KEY (
        subrace_id
    )
    REFERENCES subraces (id) ON DELETE CASCADE,
    CONSTRAINT subrace_traits_trait_id FOREIGN KEY (
        trait_id
    )
    REFERENCES traits (id) ON DELETE CASCADE
);


-- Table: subraces
CREATE TABLE IF NOT EXISTS subraces (
    id            INTEGER NOT NULL
                          PRIMARY KEY AUTOINCREMENT,
    indx          TEXT    NOT NULL,
    name          TEXT    NOT NULL,
    desc          TEXT    NOT NULL,
    race_subraces INTEGER,
    CONSTRAINT subraces_races_subraces FOREIGN KEY (
        race_subraces
    )
    REFERENCES races (id) ON DELETE SET NULL
);

INSERT INTO subraces (id, indx, name, desc, race_subraces) VALUES (73014444033, 'hill-dwarf', 'Hill Dwarf', 'As a hill dwarf, you have keen senses, deep intuition, and remarkable resilience.', 55834574849);
INSERT INTO subraces (id, indx, name, desc, race_subraces) VALUES (73014444034, 'high-elf', 'High Elf', 'As a high elf, you have a keen mind and a mastery of at least the basics of magic. In many fantasy gaming worlds, there are two kinds of high elves. One type is haughty and reclusive, believing themselves to be superior to non-elves and even other elves. The other type is more common and more friendly, and often encountered among humans and other races.', 55834574850);
INSERT INTO subraces (id, indx, name, desc, race_subraces) VALUES (73014444035, 'lightfoot-halfling', 'Lightfoot Halfling', 'As a lightfoot halfling, you can easily hide from notice, even using other people as cover. You''re inclined to be affable and get along well with others. Lightfoots are more prone to wanderlust than other halflings, and often dwell alongside other races or take up a nomadic life.', 55834574851);
INSERT INTO subraces (id, indx, name, desc, race_subraces) VALUES (73014444036, 'rock-gnome', 'Rock Gnome', 'As a rock gnome, you have a natural inventiveness and hardiness beyond that of other gnomes.', 55834574854);

-- Table: tools
CREATE TABLE IF NOT EXISTS tools (
    id            INTEGER NOT NULL
                          PRIMARY KEY AUTOINCREMENT,
    indx          TEXT    NOT NULL,
    name          TEXT    NOT NULL,
    tool_category TEXT    NOT NULL,
    equipment_id  INTEGER NOT NULL,
    CONSTRAINT tools_equipment_tool FOREIGN KEY (
        equipment_id
    )
    REFERENCES equipment (id) ON DELETE NO ACTION
);

INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411329, 'alchemists-supplies', 'Alchemist''s Supplies', '', 34359738535);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411330, 'brewers-supplies', 'Brewer''s Supplies', '', 34359738536);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411331, 'calligraphers-supplies', 'Calligrapher''s Supplies', '', 34359738537);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411332, 'carpenters-tools', 'Carpenter''s Tools', '', 34359738538);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411333, 'cartographers-tools', 'Cartographer''s Tools', '', 34359738539);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411334, 'cobblers-tools', 'Cobbler''s Tools', '', 34359738540);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411335, 'cooks-utensils', 'Cook''s utensils', '', 34359738541);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411336, 'glassblowers-tools', 'Glassblower''s Tools', '', 34359738542);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411337, 'jewelers-tools', 'Jeweler''s Tools', '', 34359738543);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411338, 'leatherworkers-tools', 'Leatherworker''s Tools', '', 34359738544);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411339, 'masons-tools', 'Mason''s Tools', '', 34359738545);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411340, 'painters-supplies', 'Painter''s Supplies', '', 34359738546);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411341, 'potters-tools', 'Potter''s Tools', '', 34359738547);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411342, 'smiths-tools', 'Smith''s Tools', '', 34359738548);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411343, 'tinkers-tools', 'Tinker''s Tools', '', 34359738549);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411344, 'weavers-tools', 'Weaver''s Tools', '', 34359738550);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411345, 'woodcarvers-tools', 'Woodcarver''s Tools', '', 34359738551);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411346, 'dice-set', 'Dice Set', '', 34359738552);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411347, 'playing-card-set', 'Playing Card Set', '', 34359738553);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411348, 'bagpipes', 'Bagpipes', '', 34359738554);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411349, 'drum', 'Drum', '', 34359738555);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411350, 'dulcimer', 'Dulcimer', '', 34359738556);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411351, 'flute', 'Flute', '', 34359738557);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411352, 'lute', 'Lute', '', 34359738558);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411353, 'lyre', 'Lyre', '', 34359738559);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411354, 'horn', 'Horn', '', 34359738560);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411355, 'pan-flute', 'Pan flute', '', 34359738561);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411356, 'shawm', 'Shawm', '', 34359738562);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411357, 'viol', 'Viol', '', 34359738563);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411358, 'navigators-tools', 'Navigator''s Tools', '', 34359738564);
INSERT INTO tools (id, indx, name, tool_category, equipment_id) VALUES (77309411359, 'thieves-tools', 'Thieves'' Tools', '', 34359738565);

-- Table: traits
CREATE TABLE IF NOT EXISTS traits (
    id   INTEGER NOT NULL
                 PRIMARY KEY AUTOINCREMENT,
    indx TEXT    NOT NULL,
    name TEXT    NOT NULL,
    desc JSON    NOT NULL
);


-- Table: vehicles
CREATE TABLE IF NOT EXISTS vehicles (
    id               INTEGER NOT NULL
                             PRIMARY KEY AUTOINCREMENT,
    indx             TEXT    NOT NULL,
    name             TEXT    NOT NULL,
    vehicle_category TEXT    NOT NULL,
    capacity         TEXT    NOT NULL,
    equipment_id     INTEGER NOT NULL,
    CONSTRAINT vehicles_equipment_vehicle FOREIGN KEY (
        equipment_id
    )
    REFERENCES equipment (id) ON DELETE NO ACTION
);

INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345921, 'camel', 'Camel', 'Mounts and Other Animals', '', 34359738566);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345922, 'donkey', 'Donkey', 'Mounts and Other Animals', '', 34359738567);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345923, 'mule', 'Mule', 'Mounts and Other Animals', '', 34359738568);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345924, 'elephant', 'Elephant', 'Mounts and Other Animals', '', 34359738569);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345925, 'horse-draft', 'Horse, draft', 'Mounts and Other Animals', '', 34359738570);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345926, 'horse-riding', 'Horse, riding', 'Mounts and Other Animals', '', 34359738571);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345927, 'mastiff', 'Mastiff', 'Mounts and Other Animals', '', 34359738572);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345928, 'pony', 'Pony', 'Mounts and Other Animals', '', 34359738573);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345929, 'warhorse', 'Warhorse', 'Mounts and Other Animals', '', 34359738574);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345930, 'barding-padded', 'Barding: Padded', 'Tack, Harness, and Drawn Vehicles', '', 34359738575);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345931, 'barding-leather', 'Barding: Leather', 'Tack, Harness, and Drawn Vehicles', '', 34359738576);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345932, 'barding-studded-leather', 'Barding: Studded Leather', 'Tack, Harness, and Drawn Vehicles', '', 34359738577);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345933, 'barding-hide', 'Barding: Hide', 'Tack, Harness, and Drawn Vehicles', '', 34359738578);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345934, 'barding-chain-shirt', 'Barding: Chain shirt', 'Tack, Harness, and Drawn Vehicles', '', 34359738579);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345935, 'barding-scale-mail', 'Barding: Scale mail', 'Tack, Harness, and Drawn Vehicles', '', 34359738580);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345936, 'barding-breastplate', 'Barding: Breastplate', 'Tack, Harness, and Drawn Vehicles', '', 34359738581);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345937, 'barding-half-plate', 'Barding: Half plate', 'Tack, Harness, and Drawn Vehicles', '', 34359738582);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345938, 'barding-ring-mail', 'Barding: Ring mail', 'Tack, Harness, and Drawn Vehicles', '', 34359738583);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345939, 'barding-chain-mail', 'Barding: Chain mail', 'Tack, Harness, and Drawn Vehicles', '', 34359738584);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345940, 'barding-splint', 'Barding: Splint', 'Tack, Harness, and Drawn Vehicles', '', 34359738585);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345941, 'barding-plate', 'Barding: Plate', 'Tack, Harness, and Drawn Vehicles', '', 34359738586);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345942, 'bit-and-bridle', 'Bit and bridle', 'Tack, Harness, and Drawn Vehicles', '', 34359738587);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345943, 'carriage', 'Carriage', 'Tack, Harness, and Drawn Vehicles', '', 34359738588);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345944, 'cart', 'Cart', 'Tack, Harness, and Drawn Vehicles', '', 34359738589);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345945, 'chariot', 'Chariot', 'Tack, Harness, and Drawn Vehicles', '', 34359738590);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345946, 'animal-feed-1-day', 'Animal Feed (1 day)', 'Tack, Harness, and Drawn Vehicles', '', 34359738591);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345947, 'saddle-exotic', 'Saddle, Exotic', 'Tack, Harness, and Drawn Vehicles', '', 34359738592);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345948, 'saddle-military', 'Saddle, Military', 'Tack, Harness, and Drawn Vehicles', '', 34359738593);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345949, 'saddle-pack', 'Saddle, Pack', 'Tack, Harness, and Drawn Vehicles', '', 34359738594);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345950, 'saddle-riding', 'Saddle, Riding', 'Tack, Harness, and Drawn Vehicles', '', 34359738595);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345951, 'saddlebags', 'Saddlebags', 'Tack, Harness, and Drawn Vehicles', '', 34359738596);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345952, 'sled', 'Sled', 'Tack, Harness, and Drawn Vehicles', '', 34359738597);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345953, 'stabling-1-day', 'Stabling (1 day)', 'Tack, Harness, and Drawn Vehicles', '', 34359738598);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345954, 'wagon', 'Wagon', 'Tack, Harness, and Drawn Vehicles', '', 34359738599);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345955, 'galley', 'Galley', 'Waterborne Vehicles', '', 34359738600);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345956, 'keelboat', 'Keelboat', 'Waterborne Vehicles', '', 34359738601);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345957, 'longship', 'Longship', 'Waterborne Vehicles', '', 34359738602);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345958, 'rowboat', 'Rowboat', 'Waterborne Vehicles', '', 34359738603);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345959, 'sailing-ship', 'Sailing ship', 'Waterborne Vehicles', '', 34359738604);
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity, equipment_id) VALUES (85899345960, 'warship', 'Warship', 'Waterborne Vehicles', '', 34359738605);

-- Table: weapon_damages
CREATE TABLE IF NOT EXISTS weapon_damages (
    id             INTEGER NOT NULL
                           PRIMARY KEY AUTOINCREMENT,
    dice           TEXT    NOT NULL,
    weapon_id      INTEGER NOT NULL,
    damage_type_id INTEGER NOT NULL,
    CONSTRAINT weapon_damages_weapons_weapon_damage FOREIGN KEY (
        weapon_id
    )
    REFERENCES weapons (id) ON DELETE NO ACTION,
    CONSTRAINT weapon_damages_damage_types_damage_type FOREIGN KEY (
        damage_type_id
    )
    REFERENCES damage_types (id) ON DELETE NO ACTION
);

INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280513, '1d4', 90194313217, 30064771074);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280514, '1d4', 90194313218, 30064771080);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280515, '1d8', 90194313219, 30064771074);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280516, '1d6', 90194313220, 30064771084);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280517, '1d6', 90194313221, 30064771080);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280518, '1d4', 90194313222, 30064771074);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280519, '1d6', 90194313223, 30064771074);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280520, '1d6', 90194313224, 30064771074);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280521, '1d4', 90194313225, 30064771084);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280522, '1d6', 90194313226, 30064771080);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280523, '1d8', 90194313227, 30064771080);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280524, '1d4', 90194313228, 30064771080);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280525, '1d6', 90194313229, 30064771080);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280526, '1d4', 90194313230, 30064771074);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280527, '1d8', 90194313231, 30064771084);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280528, '1d8', 90194313232, 30064771074);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280529, '1d10', 90194313233, 30064771084);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280530, '1d12', 90194313234, 30064771084);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280531, '2d6', 90194313235, 30064771084);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280532, '1d10', 90194313236, 30064771084);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280533, '1d12', 90194313237, 30064771080);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280534, '1d8', 90194313238, 30064771084);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280535, '2d6', 90194313239, 30064771074);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280536, '1d8', 90194313240, 30064771080);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280537, '1d10', 90194313241, 30064771080);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280538, '1d8', 90194313242, 30064771080);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280539, '1d6', 90194313243, 30064771084);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280540, '1d6', 90194313244, 30064771080);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280541, '1d6', 90194313245, 30064771084);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280542, '1d8', 90194313246, 30064771080);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280543, '1d8', 90194313247, 30064771074);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280544, '1d4', 90194313248, 30064771084);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280545, '1', 90194313249, 30064771080);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280546, '1d6', 90194313250, 30064771080);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280547, '1d10', 90194313251, 30064771080);
INSERT INTO weapon_damages (id, dice, weapon_id, damage_type_id) VALUES (94489280548, '1d8', 90194313252, 30064771080);

-- Table: weapon_properties
CREATE TABLE IF NOT EXISTS weapon_properties (
    id   INTEGER NOT NULL
                 PRIMARY KEY AUTOINCREMENT,
    indx TEXT    NOT NULL,
    name TEXT    NOT NULL,
    desc JSON    NOT NULL
);

INSERT INTO weapon_properties (id, indx, name, desc) VALUES (98784247809, 'ammunition', 'Ammunition', X'5B22596F752063616E20757365206120776561706F6E2074686174206861732074686520616D6D756E6974696F6E2070726F706572747920746F206D616B6520612072616E6765642061747461636B206F6E6C7920696620796F75206861766520616D6D756E6974696F6E20746F20666972652066726F6D2074686520776561706F6E2E20456163682074696D6520796F752061747461636B20776974682074686520776561706F6E2C20796F7520657870656E64206F6E65207069656365206F6620616D6D756E6974696F6E2E2044726177696E672074686520616D6D756E6974696F6E2066726F6D2061207175697665722C20636173652C206F72206F7468657220636F6E7461696E65722069732070617274206F66207468652061747461636B2028796F75206E656564206120667265652068616E6420746F206C6F61642061206F6E652D68616E64656420776561706F6E292E222C2241742074686520656E64206F662074686520626174746C652C20796F752063616E207265636F7665722068616C6620796F757220657870656E64656420616D6D756E6974696F6E2062792074616B696E672061206D696E75746520746F207365617263682074686520626174746C656669656C642E20496620796F7520757365206120776561706F6E2074686174206861732074686520616D6D756E6974696F6E2070726F706572747920746F206D616B652061206D656C65652061747461636B2C20796F752074726561742074686520776561706F6E20617320616E20696D70726F766973656420776561706F6E2028736565205C22496D70726F766973656420576561706F6E735C22206C6174657220696E207468652073656374696F6E292E204120736C696E67206D757374206265206C6F6164656420746F206465616C20616E792064616D616765207768656E207573656420696E2074686973207761792E225D');
INSERT INTO weapon_properties (id, indx, name, desc) VALUES (98784247810, 'finesse', 'Finesse', X'5B225768656E206D616B696E6720616E2061747461636B207769746820612066696E6573736520776561706F6E2C20796F752075736520796F75722063686F696365206F6620796F757220537472656E677468206F7220446578746572697479206D6F64696669657220666F72207468652061747461636B20616E642064616D61676520726F6C6C732E20596F75206D75737420757365207468652073616D65206D6F64696669657220666F7220626F746820726F6C6C732E225D');
INSERT INTO weapon_properties (id, indx, name, desc) VALUES (98784247811, 'heavy', 'Heavy', X'5B22536D616C6C20637265617475726573206861766520646973616476616E74616765206F6E2061747461636B20726F6C6C73207769746820686561767920776561706F6E732E204120686561767920776561706F6E27732073697A6520616E642062756C6B206D616B6520697420746F6F206C6172676520666F72206120536D616C6C20637265617475726520746F20757365206566666563746976656C792E225D');
INSERT INTO weapon_properties (id, indx, name, desc) VALUES (98784247812, 'light', 'Light', X'5B2241206C6967687420776561706F6E20697320736D616C6C20616E64206561737920746F2068616E646C652C206D616B696E6720697420696465616C20666F7220757365207768656E206669676874696E6720776974682074776F20776561706F6E732E225D');
INSERT INTO weapon_properties (id, indx, name, desc) VALUES (98784247813, 'loading', 'Loading', X'5B2242656361757365206F66207468652074696D6520726571756972656420746F206C6F6164207468697320776561706F6E2C20796F752063616E2066697265206F6E6C79206F6E65207069656365206F6620616D6D756E6974696F6E2066726F6D206974207768656E20796F752075736520616E20616374696F6E2C20626F6E757320616374696F6E2C206F72207265616374696F6E20746F20666972652069742C207265676172646C657373206F6620746865206E756D626572206F662061747461636B7320796F752063616E206E6F726D616C6C79206D616B652E225D');
INSERT INTO weapon_properties (id, indx, name, desc) VALUES (98784247814, 'reach', 'Reach', X'5B225468697320776561706F6E20616464732035206665657420746F20796F7572207265616368207768656E20796F752061747461636B20776974682069742C2061732077656C6C206173207768656E2064657465726D696E696E6720796F757220726561636820666F72206F70706F7274756E6974792061747461636B7320776974682069742E225D');
INSERT INTO weapon_properties (id, indx, name, desc) VALUES (98784247815, 'special', 'Special', X'5B224120776561706F6E207769746820746865207370656369616C2070726F70657274792068617320756E757375616C2072756C657320676F7665726E696E6720697473207573652C206578706C61696E656420696E2074686520776561706F6E2773206465736372697074696F6E2028736565205C225370656369616C20576561706F6E735C22206C6174657220696E20746869732073656374696F6E292E225D');
INSERT INTO weapon_properties (id, indx, name, desc) VALUES (98784247816, 'thrown', 'Thrown', X'5B224966206120776561706F6E2068617320746865207468726F776E2070726F70657274792C20796F752063616E207468726F772074686520776561706F6E20746F206D616B6520612072616E6765642061747461636B2E2049662074686520776561706F6E2069732061206D656C656520776561706F6E2C20796F7520757365207468652073616D65206162696C697479206D6F64696669657220666F7220746861742061747461636B20726F6C6C20616E642064616D61676520726F6C6C207468617420796F7520776F756C642075736520666F722061206D656C65652061747461636B20776974682074686520776561706F6E2E20466F72206578616D706C652C20696620796F75207468726F7720612068616E646178652C20796F752075736520796F757220537472656E6774682C2062757420696620796F75207468726F772061206461676765722C20796F752063616E207573652065697468657220796F757220537472656E677468206F7220796F7572204465787465726974792C2073696E6365207468652064616767657220686173207468652066696E657373652070726F70657274792E225D');
INSERT INTO weapon_properties (id, indx, name, desc) VALUES (98784247817, 'two-handed', 'Two-Handed', X'5B225468697320776561706F6E2072657175697265732074776F2068616E6473207768656E20796F752061747461636B20776974682069742E225D');
INSERT INTO weapon_properties (id, indx, name, desc) VALUES (98784247818, 'versatile', 'Versatile', X'5B225468697320776561706F6E2063616E20626520757365642077697468206F6E65206F722074776F2068616E64732E20412064616D6167652076616C756520696E20706172656E74686573657320617070656172732077697468207468652070726F70657274792D2D7468652064616D616765207768656E2074686520776561706F6E206973207573656420776974682074776F2068616E647320746F206D616B652061206D656C65652061747461636B2E225D');
INSERT INTO weapon_properties (id, indx, name, desc) VALUES (98784247819, 'monk', 'Monk', X'5B224D6F6E6B73206761696E207365766572616C2062656E6566697473207768696C6520756E61726D6564206F72207769656C64696E67206F6E6C79206D6F6E6B20776561706F6E73207768696C652074686579206172656E27742077656172696E672061726D6F72206F72207769656C64696E6720736869656C64732E225D');

-- Table: weapon_weapon_properties
CREATE TABLE IF NOT EXISTS weapon_weapon_properties (
    weapon_id          INTEGER NOT NULL,
    weapon_property_id INTEGER NOT NULL,
    PRIMARY KEY (
        weapon_id,
        weapon_property_id
    ),
    CONSTRAINT weapon_weapon_properties_weapon_id FOREIGN KEY (
        weapon_id
    )
    REFERENCES weapons (id) ON DELETE CASCADE,
    CONSTRAINT weapon_weapon_properties_weapon_property_id FOREIGN KEY (
        weapon_property_id
    )
    REFERENCES weapon_properties (id) ON DELETE CASCADE
);

INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313217, 98784247812);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313217, 98784247819);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313218, 98784247810);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313218, 98784247812);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313218, 98784247816);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313218, 98784247819);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313219, 98784247817);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313220, 98784247812);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313220, 98784247816);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313220, 98784247819);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313221, 98784247816);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313221, 98784247819);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313222, 98784247812);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313222, 98784247816);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313222, 98784247819);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313223, 98784247819);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313224, 98784247818);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313224, 98784247819);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313225, 98784247812);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313225, 98784247819);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313226, 98784247816);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313226, 98784247818);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313226, 98784247819);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313227, 98784247813);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313227, 98784247817);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313227, 98784247809);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313228, 98784247810);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313228, 98784247816);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313229, 98784247809);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313229, 98784247817);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313230, 98784247809);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313231, 98784247818);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313233, 98784247811);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313233, 98784247814);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313233, 98784247817);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313234, 98784247811);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313234, 98784247817);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313235, 98784247811);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313235, 98784247817);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313236, 98784247814);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313236, 98784247817);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313236, 98784247811);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313237, 98784247814);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313237, 98784247815);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313238, 98784247818);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313239, 98784247811);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313239, 98784247817);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313241, 98784247811);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313241, 98784247814);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313241, 98784247817);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313242, 98784247810);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313243, 98784247810);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313243, 98784247812);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313244, 98784247810);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313244, 98784247812);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313244, 98784247819);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313245, 98784247816);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313245, 98784247818);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313247, 98784247818);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313248, 98784247810);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313248, 98784247814);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313249, 98784247809);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313249, 98784247813);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313250, 98784247809);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313250, 98784247812);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313250, 98784247813);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313251, 98784247809);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313251, 98784247811);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313251, 98784247813);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313251, 98784247817);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313252, 98784247809);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313252, 98784247811);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313252, 98784247817);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313253, 98784247816);
INSERT INTO weapon_weapon_properties (weapon_id, weapon_property_id) VALUES (90194313253, 98784247815);

-- Table: weapons
CREATE TABLE IF NOT EXISTS weapons (
    id              INTEGER NOT NULL
                            PRIMARY KEY AUTOINCREMENT,
    indx            TEXT    NOT NULL,
    name            TEXT    NOT NULL,
    weapon_category TEXT    NOT NULL,
    weapon_range    TEXT    NOT NULL,
    equipment_id    INTEGER NOT NULL,
    CONSTRAINT weapons_equipment_weapon FOREIGN KEY (
        equipment_id
    )
    REFERENCES equipment (id) ON DELETE NO ACTION
);

INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313217, 'club', 'Club', 'Simple', 'Melee', 34359738369);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313218, 'dagger', 'Dagger', 'Simple', 'Melee', 34359738370);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313219, 'greatclub', 'Greatclub', 'Simple', 'Melee', 34359738371);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313220, 'handaxe', 'Handaxe', 'Simple', 'Melee', 34359738372);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313221, 'javelin', 'Javelin', 'Simple', 'Melee', 34359738373);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313222, 'light-hammer', 'Light hammer', 'Simple', 'Melee', 34359738374);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313223, 'mace', 'Mace', 'Simple', 'Melee', 34359738375);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313224, 'quarterstaff', 'Quarterstaff', 'Simple', 'Melee', 34359738376);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313225, 'sickle', 'Sickle', 'Simple', 'Melee', 34359738377);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313226, 'spear', 'Spear', 'Simple', 'Melee', 34359738378);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313227, 'crossbow-light', 'Crossbow, light', 'Simple', 'Ranged', 34359738379);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313228, 'dart', 'Dart', 'Simple', 'Ranged', 34359738380);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313229, 'shortbow', 'Shortbow', 'Simple', 'Ranged', 34359738381);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313230, 'sling', 'Sling', 'Simple', 'Ranged', 34359738382);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313231, 'battleaxe', 'Battleaxe', 'Martial', 'Melee', 34359738383);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313232, 'flail', 'Flail', 'Martial', 'Melee', 34359738384);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313233, 'glaive', 'Glaive', 'Martial', 'Melee', 34359738385);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313234, 'greataxe', 'Greataxe', 'Martial', 'Melee', 34359738386);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313235, 'greatsword', 'Greatsword', 'Martial', 'Melee', 34359738387);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313236, 'halberd', 'Halberd', 'Martial', 'Melee', 34359738388);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313237, 'lance', 'Lance', 'Martial', 'Melee', 34359738389);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313238, 'longsword', 'Longsword', 'Martial', 'Melee', 34359738390);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313239, 'maul', 'Maul', 'Martial', 'Melee', 34359738391);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313240, 'morningstar', 'Morningstar', 'Martial', 'Melee', 34359738392);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313241, 'pike', 'Pike', 'Martial', 'Melee', 34359738393);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313242, 'rapier', 'Rapier', 'Martial', 'Melee', 34359738394);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313243, 'scimitar', 'Scimitar', 'Martial', 'Melee', 34359738395);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313244, 'shortsword', 'Shortsword', 'Martial', 'Melee', 34359738396);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313245, 'trident', 'Trident', 'Martial', 'Melee', 34359738397);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313246, 'war-pick', 'War pick', 'Martial', 'Melee', 34359738398);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313247, 'warhammer', 'Warhammer', 'Martial', 'Melee', 34359738399);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313248, 'whip', 'Whip', 'Martial', 'Melee', 34359738400);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313249, 'blowgun', 'Blowgun', 'Martial', 'Ranged', 34359738401);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313250, 'crossbow-hand', 'Crossbow, hand', 'Martial', 'Ranged', 34359738402);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313251, 'crossbow-heavy', 'Crossbow, heavy', 'Martial', 'Ranged', 34359738403);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313252, 'longbow', 'Longbow', 'Martial', 'Ranged', 34359738404);
INSERT INTO weapons (id, indx, name, weapon_category, weapon_range, equipment_id) VALUES (90194313253, 'net', 'Net', 'Martial', 'Ranged', 34359738405);

-- Index: ability_scores_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS ability_scores_indx_key ON ability_scores (
    indx
);


-- Index: armors_equipment_id_key
CREATE UNIQUE INDEX IF NOT EXISTS armors_equipment_id_key ON armors (
    equipment_id
);


-- Index: armors_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS armors_indx_key ON armors (
    indx
);


-- Index: choices_race_id_key
CREATE UNIQUE INDEX IF NOT EXISTS choices_race_id_key ON choices (
    race_id
);


-- Index: classes_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS classes_indx_key ON classes (
    indx
);


-- Index: damage_types_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS damage_types_indx_key ON damage_types (
    indx
);


-- Index: ent_types_type_key
CREATE UNIQUE INDEX IF NOT EXISTS ent_types_type_key ON ent_types (
    type
);


-- Index: equipment_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS equipment_indx_key ON equipment (
    indx
);


-- Index: gears_equipment_id_key
CREATE UNIQUE INDEX IF NOT EXISTS gears_equipment_id_key ON gears (
    equipment_id
);


-- Index: gears_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS gears_indx_key ON gears (
    indx
);


-- Index: languages_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS languages_indx_key ON languages (
    indx
);


-- Index: magic_schools_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS magic_schools_indx_key ON magic_schools (
    indx
);


-- Index: proficiencies_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS proficiencies_indx_key ON proficiencies (
    indx
);


-- Index: races_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS races_indx_key ON races (
    indx
);


-- Index: rule_sections_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS rule_sections_indx_key ON rule_sections (
    indx
);


-- Index: rules_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS rules_indx_key ON rules (
    indx
);


-- Index: skills_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS skills_indx_key ON skills (
    indx
);


-- Index: subraces_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS subraces_indx_key ON subraces (
    indx
);


-- Index: tools_equipment_id_key
CREATE UNIQUE INDEX IF NOT EXISTS tools_equipment_id_key ON tools (
    equipment_id
);


-- Index: tools_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS tools_indx_key ON tools (
    indx
);


-- Index: traits_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS traits_indx_key ON traits (
    indx
);


-- Index: vehicles_equipment_id_key
CREATE UNIQUE INDEX IF NOT EXISTS vehicles_equipment_id_key ON vehicles (
    equipment_id
);


-- Index: vehicles_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS vehicles_indx_key ON vehicles (
    indx
);


-- Index: weapon_properties_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS weapon_properties_indx_key ON weapon_properties (
    indx
);


-- Index: weapons_equipment_id_key
CREATE UNIQUE INDEX IF NOT EXISTS weapons_equipment_id_key ON weapons (
    equipment_id
);


-- Index: weapons_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS weapons_indx_key ON weapons (
    indx
);


COMMIT TRANSACTION;
PRAGMA foreign_keys = on;
