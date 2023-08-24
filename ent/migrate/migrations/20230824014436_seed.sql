--
-- File generated with SQLiteStudio v3.4.4 on Wed Aug 23 18:57:40 2023
--
-- Text encoding used: UTF-8
--
PRAGMA foreign_keys = off;

-- Table: ability_scores
CREATE TABLE IF NOT EXISTS ability_scores (
    id        INTEGER NOT NULL
                      PRIMARY KEY AUTOINCREMENT,
    indx      TEXT    NOT NULL,
    name      TEXT    NOT NULL,
    full_name TEXT    NOT NULL,
    desc      JSON    NOT NULL
);

INSERT INTO ability_scores (id, indx, name, full_name, desc) VALUES (1, 'str', 'STR', 'Strength', X'5B22537472656E677468206D6561737572657320626F64696C7920706F7765722C206174686C6574696320747261696E696E672C20616E642074686520657874656E7420746F20776869636820796F752063616E2065786572742072617720706879736963616C20666F7263652E222C224120537472656E67746820636865636B2063616E206D6F64656C20616E7920617474656D707420746F206C6966742C20707573682C2070756C6C2C206F7220627265616B20736F6D657468696E672C20746F20666F72636520796F757220626F6479207468726F75676820612073706163652C206F7220746F206F7468657277697365206170706C7920627275746520666F72636520746F206120736974756174696F6E2E20546865204174686C657469637320736B696C6C207265666C6563747320617074697475646520696E206365727461696E206B696E6473206F6620537472656E67746820636865636B732E225D');
INSERT INTO ability_scores (id, indx, name, full_name, desc) VALUES (2, 'dex', 'DEX', 'Dexterity', X'5B22446578746572697479206D65617375726573206167696C6974792C207265666C657865732C20616E642062616C616E63652E222C22412044657874657269747920636865636B2063616E206D6F64656C20616E7920617474656D707420746F206D6F7665206E696D626C792C20717569636B6C792C206F722071756965746C792C206F7220746F206B6565702066726F6D2066616C6C696E67206F6E20747269636B7920666F6F74696E672E20546865204163726F6261746963732C20536C6569676874206F662048616E642C20616E6420537465616C746820736B696C6C73207265666C65637420617074697475646520696E206365727461696E206B696E6473206F662044657874657269747920636865636B732E225D');
INSERT INTO ability_scores (id, indx, name, full_name, desc) VALUES (3, 'con', 'CON', 'Constitution', X'5B22436F6E737469747574696F6E206D65617375726573206865616C74682C207374616D696E612C20616E6420766974616C20666F7263652E222C22436F6E737469747574696F6E20636865636B732061726520756E636F6D6D6F6E2C20616E64206E6F20736B696C6C73206170706C7920746F20436F6E737469747574696F6E20636865636B732C20626563617573652074686520656E647572616E63652074686973206162696C69747920726570726573656E7473206973206C617267656C79207061737369766520726174686572207468616E20696E766F6C76696E672061207370656369666963206566666F7274206F6E207468652070617274206F66206120636861726163746572206F72206D6F6E737465722E225D');
INSERT INTO ability_scores (id, indx, name, full_name, desc) VALUES (4, 'int', 'INT', 'Intelligence', X'5B22496E74656C6C6967656E6365206D65617375726573206D656E74616C206163756974792C206163637572616379206F6620726563616C6C2C20616E6420746865206162696C69747920746F20726561736F6E2E222C22416E20496E74656C6C6967656E636520636865636B20636F6D657320696E746F20706C6179207768656E20796F75206E65656420746F2064726177206F6E206C6F6769632C20656475636174696F6E2C206D656D6F72792C206F722064656475637469766520726561736F6E696E672E2054686520417263616E612C20486973746F72792C20496E7665737469676174696F6E2C204E61747572652C20616E642052656C6967696F6E20736B696C6C73207265666C65637420617074697475646520696E206365727461696E206B696E6473206F6620496E74656C6C6967656E636520636865636B732E225D');
INSERT INTO ability_scores (id, indx, name, full_name, desc) VALUES (5, 'wis', 'WIS', 'Wisdom', X'5B22576973646F6D207265666C6563747320686F7720617474756E656420796F752061726520746F2074686520776F726C642061726F756E6420796F7520616E6420726570726573656E747320706572636570746976656E65737320616E6420696E74756974696F6E2E222C224120576973646F6D20636865636B206D69676874207265666C65637420616E206566666F727420746F207265616420626F6479206C616E67756167652C20756E6465727374616E6420736F6D656F6E652773206665656C696E67732C206E6F74696365207468696E67732061626F75742074686520656E7669726F6E6D656E742C206F72206361726520666F7220616E20696E6A7572656420706572736F6E2E2054686520416E696D616C2048616E646C696E672C20496E73696768742C204D65646963696E652C2050657263657074696F6E2C20616E6420537572766976616C20736B696C6C73207265666C65637420617074697475646520696E206365727461696E206B696E6473206F6620576973646F6D20636865636B732E225D');
INSERT INTO ability_scores (id, indx, name, full_name, desc) VALUES (6, 'cha', 'CHA', 'Charisma', X'5B224368617269736D61206D6561737572657320796F7572206162696C69747920746F20696E746572616374206566666563746976656C792077697468206F74686572732E20497420696E636C75646573207375636820666163746F727320617320636F6E666964656E636520616E6420656C6F7175656E63652C20616E642069742063616E20726570726573656E74206120636861726D696E67206F7220636F6D6D616E64696E6720706572736F6E616C6974792E222C2241204368617269736D6120636865636B206D69676874206172697365207768656E20796F752074727920746F20696E666C75656E6365206F7220656E7465727461696E206F74686572732C207768656E20796F752074727920746F206D616B6520616E20696D7072657373696F6E206F722074656C6C206120636F6E76696E63696E67206C69652C206F72207768656E20796F7520617265206E617669676174696E67206120747269636B7920736F6369616C20736974756174696F6E2E2054686520446563657074696F6E2C20496E74696D69646174696F6E2C20506572666F726D616E63652C20616E642050657273756173696F6E20736B696C6C73207265666C65637420617074697475646520696E206365727461696E206B696E6473206F66204368617269736D6120636865636B732E225D');

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

INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (8589934593, 11, 1, 0, 4294967297);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (8589934594, 11, 1, 0, 4294967298);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (8589934595, 12, 1, 0, 4294967299);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (8589934596, 12, 1, 2, 4294967300);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (8589934597, 13, 1, 2, 4294967301);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (8589934598, 14, 1, 2, 4294967302);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (8589934599, 14, 1, 2, 4294967303);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (8589934600, 15, 1, 2, 4294967304);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (8589934601, 14, 0, 0, 4294967305);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (8589934602, 16, 0, 0, 4294967306);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (8589934603, 17, 0, 0, 4294967307);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (8589934604, 18, 0, 0, 4294967308);
INSERT INTO armor_classes (id, base, dex_bonus, max_bonus, armor_armor_class) VALUES (8589934605, 2, 0, 0, 4294967309);

-- Table: armors
CREATE TABLE IF NOT EXISTS armors (
    id                   INTEGER NOT NULL
                                 PRIMARY KEY AUTOINCREMENT,
    indx                 TEXT    NOT NULL,
    name                 TEXT    NOT NULL,
    stealth_disadvantage BOOL    NOT NULL,
    min_strength         INTEGER NOT NULL
);

INSERT INTO armors (id, indx, name, stealth_disadvantage, min_strength) VALUES (4294967297, 'padded-armor', 'Padded Armor', 1, 0);
INSERT INTO armors (id, indx, name, stealth_disadvantage, min_strength) VALUES (4294967298, 'leather-armor', 'Leather Armor', 0, 0);
INSERT INTO armors (id, indx, name, stealth_disadvantage, min_strength) VALUES (4294967299, 'studded-leather-armor', 'Studded Leather Armor', 0, 0);
INSERT INTO armors (id, indx, name, stealth_disadvantage, min_strength) VALUES (4294967300, 'hide-armor', 'Hide Armor', 0, 0);
INSERT INTO armors (id, indx, name, stealth_disadvantage, min_strength) VALUES (4294967301, 'chain-shirt', 'Chain Shirt', 0, 0);
INSERT INTO armors (id, indx, name, stealth_disadvantage, min_strength) VALUES (4294967302, 'scale-mail', 'Scale Mail', 1, 0);
INSERT INTO armors (id, indx, name, stealth_disadvantage, min_strength) VALUES (4294967303, 'breastplate', 'Breastplate', 0, 0);
INSERT INTO armors (id, indx, name, stealth_disadvantage, min_strength) VALUES (4294967304, 'half-plate-armor', 'Half Plate Armor', 1, 0);
INSERT INTO armors (id, indx, name, stealth_disadvantage, min_strength) VALUES (4294967305, 'ring-mail', 'Ring Mail', 1, 0);
INSERT INTO armors (id, indx, name, stealth_disadvantage, min_strength) VALUES (4294967306, 'chain-mail', 'Chain Mail', 1, 13);
INSERT INTO armors (id, indx, name, stealth_disadvantage, min_strength) VALUES (4294967307, 'splint-armor', 'Splint Armor', 1, 15);
INSERT INTO armors (id, indx, name, stealth_disadvantage, min_strength) VALUES (4294967308, 'plate-armor', 'Plate Armor', 1, 15);
INSERT INTO armors (id, indx, name, stealth_disadvantage, min_strength) VALUES (4294967309, 'shield', 'Shield', 0, 0);

-- Table: class_saving_throws
CREATE TABLE IF NOT EXISTS class_saving_throws (
    class_id         INTEGER NOT NULL,
    ability_score_id INTEGER NOT NULL,
    PRIMARY KEY (
        class_id,
        ability_score_id
    ),
    CONSTRAINT class_saving_throws_class_id FOREIGN KEY (
        class_id
    )
    REFERENCES classes (id) ON DELETE CASCADE,
    CONSTRAINT class_saving_throws_ability_score_id FOREIGN KEY (
        ability_score_id
    )
    REFERENCES ability_scores (id) ON DELETE CASCADE
);


-- Table: classes
CREATE TABLE IF NOT EXISTS classes (
    id      INTEGER NOT NULL
                    PRIMARY KEY AUTOINCREMENT,
    indx    TEXT    NOT NULL,
    name    TEXT    NOT NULL,
    hit_die INTEGER NOT NULL
);

INSERT INTO classes (id, indx, name, hit_die) VALUES (12884901889, 'barbarian', 'Barbarian', 12);
INSERT INTO classes (id, indx, name, hit_die) VALUES (12884901890, 'bard', 'Bard', 8);
INSERT INTO classes (id, indx, name, hit_die) VALUES (12884901891, 'cleric', 'Cleric', 8);
INSERT INTO classes (id, indx, name, hit_die) VALUES (12884901892, 'druid', 'Druid', 8);
INSERT INTO classes (id, indx, name, hit_die) VALUES (12884901893, 'fighter', 'Fighter', 10);
INSERT INTO classes (id, indx, name, hit_die) VALUES (12884901894, 'monk', 'Monk', 8);
INSERT INTO classes (id, indx, name, hit_die) VALUES (12884901895, 'paladin', 'Paladin', 10);
INSERT INTO classes (id, indx, name, hit_die) VALUES (12884901896, 'ranger', 'Ranger', 10);
INSERT INTO classes (id, indx, name, hit_die) VALUES (12884901897, 'rogue', 'Rogue', 8);
INSERT INTO classes (id, indx, name, hit_die) VALUES (12884901898, 'sorcerer', 'Sorcerer', 6);
INSERT INTO classes (id, indx, name, hit_die) VALUES (12884901899, 'warlock', 'Warlock', 8);
INSERT INTO classes (id, indx, name, hit_die) VALUES (12884901900, 'wizard', 'Wizard', 6);

-- Table: costs
CREATE TABLE IF NOT EXISTS costs (
    id       INTEGER NOT NULL
                     PRIMARY KEY AUTOINCREMENT,
    quantity INTEGER NOT NULL,
    unit     TEXT    NOT NULL
);


-- Table: damage_types
CREATE TABLE IF NOT EXISTS damage_types (
    id                        INTEGER NOT NULL
                                      PRIMARY KEY AUTOINCREMENT,
    indx                      TEXT    NOT NULL,
    name                      TEXT    NOT NULL,
    desc                      JSON    NOT NULL,
    weapon_damage_damage_type INTEGER,
    CONSTRAINT damage_types_weapon_damages_damage_type FOREIGN KEY (
        weapon_damage_damage_type
    )
    REFERENCES weapon_damages (id) ON DELETE SET NULL
);

-- Table: equipment
CREATE TABLE IF NOT EXISTS equipment (
    id                 INTEGER NOT NULL
                               PRIMARY KEY AUTOINCREMENT,
    indx               TEXT    NOT NULL,
    name               TEXT    NOT NULL,
    equipment_category TEXT    NOT NULL
                               DEFAULT 'other',
    equipment_weapon   INTEGER,
    equipment_armor    INTEGER,
    equipment_gear     INTEGER,
    equipment_tool     INTEGER,
    equipment_vehicle  INTEGER,
    equipment_cost     INTEGER,
    CONSTRAINT equipment_weapons_weapon FOREIGN KEY (
        equipment_weapon
    )
    REFERENCES weapons (id) ON DELETE SET NULL,
    CONSTRAINT equipment_armors_armor FOREIGN KEY (
        equipment_armor
    )
    REFERENCES armors (id) ON DELETE SET NULL,
    CONSTRAINT equipment_gears_gear FOREIGN KEY (
        equipment_gear
    )
    REFERENCES gears (id) ON DELETE SET NULL,
    CONSTRAINT equipment_tools_tool FOREIGN KEY (
        equipment_tool
    )
    REFERENCES tools (id) ON DELETE SET NULL,
    CONSTRAINT equipment_vehicles_vehicle FOREIGN KEY (
        equipment_vehicle
    )
    REFERENCES vehicles (id) ON DELETE SET NULL,
    CONSTRAINT equipment_costs_cost FOREIGN KEY (
        equipment_cost
    )
    REFERENCES costs (id) ON DELETE SET NULL
);

INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803777, 'club', 'Club', 'weapon', 60129542145, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803778, 'dagger', 'Dagger', 'weapon', 60129542146, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803779, 'greatclub', 'Greatclub', 'weapon', 60129542147, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803780, 'handaxe', 'Handaxe', 'weapon', 60129542148, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803781, 'javelin', 'Javelin', 'weapon', 60129542149, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803782, 'light-hammer', 'Light hammer', 'weapon', 60129542150, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803783, 'mace', 'Mace', 'weapon', 60129542151, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803784, 'quarterstaff', 'Quarterstaff', 'weapon', 60129542152, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803785, 'sickle', 'Sickle', 'weapon', 60129542153, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803786, 'spear', 'Spear', 'weapon', 60129542154, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803787, 'crossbow-light', 'Crossbow, light', 'weapon', 60129542155, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803788, 'dart', 'Dart', 'weapon', 60129542156, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803789, 'shortbow', 'Shortbow', 'weapon', 60129542157, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803790, 'sling', 'Sling', 'weapon', 60129542158, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803791, 'battleaxe', 'Battleaxe', 'weapon', 60129542159, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803792, 'flail', 'Flail', 'weapon', 60129542160, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803793, 'glaive', 'Glaive', 'weapon', 60129542161, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803794, 'greataxe', 'Greataxe', 'weapon', 60129542162, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803795, 'greatsword', 'Greatsword', 'weapon', 60129542163, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803796, 'halberd', 'Halberd', 'weapon', 60129542164, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803797, 'lance', 'Lance', 'weapon', 60129542165, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803798, 'longsword', 'Longsword', 'weapon', 60129542166, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803799, 'maul', 'Maul', 'weapon', 60129542167, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803800, 'morningstar', 'Morningstar', 'weapon', 60129542168, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803801, 'pike', 'Pike', 'weapon', 60129542169, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803802, 'rapier', 'Rapier', 'weapon', 60129542170, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803803, 'scimitar', 'Scimitar', 'weapon', 60129542171, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803804, 'shortsword', 'Shortsword', 'weapon', 60129542172, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803805, 'trident', 'Trident', 'weapon', 60129542173, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803806, 'war-pick', 'War pick', 'weapon', 60129542174, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803807, 'warhammer', 'Warhammer', 'weapon', 60129542175, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803808, 'whip', 'Whip', 'weapon', 60129542176, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803809, 'blowgun', 'Blowgun', 'weapon', 60129542177, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803810, 'crossbow-hand', 'Crossbow, hand', 'weapon', 60129542178, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803811, 'crossbow-heavy', 'Crossbow, heavy', 'weapon', 60129542179, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803812, 'longbow', 'Longbow', 'weapon', 60129542180, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803813, 'net', 'Net', 'weapon', 60129542181, NULL, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803814, 'padded-armor', 'Padded Armor', 'armor', NULL, 4294967297, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803815, 'leather-armor', 'Leather Armor', 'armor', NULL, 4294967298, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803816, 'studded-leather-armor', 'Studded Leather Armor', 'armor', NULL, 4294967299, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803817, 'hide-armor', 'Hide Armor', 'armor', NULL, 4294967300, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803818, 'chain-shirt', 'Chain Shirt', 'armor', NULL, 4294967301, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803819, 'scale-mail', 'Scale Mail', 'armor', NULL, 4294967302, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803820, 'breastplate', 'Breastplate', 'armor', NULL, 4294967303, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803821, 'half-plate-armor', 'Half Plate Armor', 'armor', NULL, 4294967304, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803822, 'ring-mail', 'Ring Mail', 'armor', NULL, 4294967305, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803823, 'chain-mail', 'Chain Mail', 'armor', NULL, 4294967306, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803824, 'splint-armor', 'Splint Armor', 'armor', NULL, 4294967307, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803825, 'plate-armor', 'Plate Armor', 'armor', NULL, 4294967308, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803826, 'shield', 'Shield', 'armor', NULL, 4294967309, NULL, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803827, 'abacus', 'Abacus', 'adventuring_gear', NULL, NULL, 30064771073, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803828, 'acid-vial', 'Acid (vial)', 'adventuring_gear', NULL, NULL, 30064771074, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803829, 'alchemists-fire-flask', 'Alchemist''s fire (flask)', 'adventuring_gear', NULL, NULL, 30064771075, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803830, 'alms-box', 'Alms box', 'adventuring_gear', NULL, NULL, 30064771076, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803831, 'arrow', 'Arrow', 'adventuring_gear', NULL, NULL, 30064771077, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803832, 'block-of-incense', 'Block of incense', 'adventuring_gear', NULL, NULL, 30064771078, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803833, 'blowgun-needle', 'Blowgun needle', 'adventuring_gear', NULL, NULL, 30064771079, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803834, 'censer', 'Censer', 'adventuring_gear', NULL, NULL, 30064771080, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803835, 'crossbow-bolt', 'Crossbow bolt', 'adventuring_gear', NULL, NULL, 30064771081, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803836, 'sling-bullet', 'Sling bullet', 'adventuring_gear', NULL, NULL, 30064771082, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803837, 'amulet', 'Amulet', 'adventuring_gear', NULL, NULL, 30064771083, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803838, 'antitoxin-vial', 'Antitoxin (vial)', 'adventuring_gear', NULL, NULL, 30064771084, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803839, 'crystal', 'Crystal', 'adventuring_gear', NULL, NULL, 30064771085, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803840, 'orb', 'Orb', 'adventuring_gear', NULL, NULL, 30064771086, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803841, 'rod', 'Rod', 'adventuring_gear', NULL, NULL, 30064771087, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803842, 'staff', 'Staff', 'adventuring_gear', NULL, NULL, 30064771088, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803843, 'wand', 'Wand', 'adventuring_gear', NULL, NULL, 30064771089, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803844, 'backpack', 'Backpack', 'adventuring_gear', NULL, NULL, 30064771090, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803845, 'ball-bearings-bag-of-1000', 'Ball bearings (bag of 1,000)', 'adventuring_gear', NULL, NULL, 30064771091, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803846, 'barrel', 'Barrel', 'adventuring_gear', NULL, NULL, 30064771092, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803847, 'basket', 'Basket', 'adventuring_gear', NULL, NULL, 30064771093, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803848, 'bedroll', 'Bedroll', 'adventuring_gear', NULL, NULL, 30064771094, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803849, 'bell', 'Bell', 'adventuring_gear', NULL, NULL, 30064771095, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803850, 'blanket', 'Blanket', 'adventuring_gear', NULL, NULL, 30064771096, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803851, 'block-and-tackle', 'Block and tackle', 'adventuring_gear', NULL, NULL, 30064771097, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803852, 'book', 'Book', 'adventuring_gear', NULL, NULL, 30064771098, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803853, 'bottle-glass', 'Bottle, glass', 'adventuring_gear', NULL, NULL, 30064771099, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803854, 'bucket', 'Bucket', 'adventuring_gear', NULL, NULL, 30064771100, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803855, 'caltrops', 'Caltrops', 'adventuring_gear', NULL, NULL, 30064771101, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803856, 'candle', 'Candle', 'adventuring_gear', NULL, NULL, 30064771102, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803857, 'case-crossbow-bolt', 'Case, crossbow bolt', 'adventuring_gear', NULL, NULL, 30064771103, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803858, 'case-map-or-scroll', 'Case, map or scroll', 'adventuring_gear', NULL, NULL, 30064771104, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803859, 'chain-10-feet', 'Chain (10 feet)', 'adventuring_gear', NULL, NULL, 30064771105, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803860, 'chalk-1-piece', 'Chalk (1 piece)', 'adventuring_gear', NULL, NULL, 30064771106, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803861, 'chest', 'Chest', 'adventuring_gear', NULL, NULL, 30064771107, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803862, 'clothes-common', 'Clothes, common', 'adventuring_gear', NULL, NULL, 30064771108, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803863, 'clothes-costume', 'Clothes, costume', 'adventuring_gear', NULL, NULL, 30064771109, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803864, 'clothes-fine', 'Clothes, fine', 'adventuring_gear', NULL, NULL, 30064771110, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803865, 'clothes-travelers', 'Clothes, traveler''s', 'adventuring_gear', NULL, NULL, 30064771111, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803866, 'component-pouch', 'Component pouch', 'adventuring_gear', NULL, NULL, 30064771112, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803867, 'crowbar', 'Crowbar', 'adventuring_gear', NULL, NULL, 30064771113, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803868, 'sprig-of-mistletoe', 'Sprig of mistletoe', 'adventuring_gear', NULL, NULL, 30064771114, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803869, 'totem', 'Totem', 'adventuring_gear', NULL, NULL, 30064771115, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803870, 'wooden-staff', 'Wooden staff', 'adventuring_gear', NULL, NULL, 30064771116, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803871, 'yew-wand', 'Yew wand', 'adventuring_gear', NULL, NULL, 30064771117, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803872, 'emblem', 'Emblem', 'adventuring_gear', NULL, NULL, 30064771118, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803873, 'fishing-tackle', 'Fishing tackle', 'adventuring_gear', NULL, NULL, 30064771119, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803874, 'flask-or-tankard', 'Flask or tankard', 'adventuring_gear', NULL, NULL, 30064771120, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803875, 'grappling-hook', 'Grappling hook', 'adventuring_gear', NULL, NULL, 30064771121, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803876, 'hammer', 'Hammer', 'adventuring_gear', NULL, NULL, 30064771122, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803877, 'hammer-sledge', 'Hammer, sledge', 'adventuring_gear', NULL, NULL, 30064771123, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803878, 'holy-water-flask', 'Holy water (flask)', 'adventuring_gear', NULL, NULL, 30064771124, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803879, 'hourglass', 'Hourglass', 'adventuring_gear', NULL, NULL, 30064771125, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803880, 'hunting-trap', 'Hunting trap', 'adventuring_gear', NULL, NULL, 30064771126, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803881, 'ink-1-ounce-bottle', 'Ink (1 ounce bottle)', 'adventuring_gear', NULL, NULL, 30064771127, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803882, 'ink-pen', 'Ink pen', 'adventuring_gear', NULL, NULL, 30064771128, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803883, 'jug-or-pitcher', 'Jug or pitcher', 'adventuring_gear', NULL, NULL, 30064771129, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803884, 'climbers-kit', 'Climber''s Kit', 'adventuring_gear', NULL, NULL, 30064771130, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803885, 'disguise-kit', 'Disguise Kit', 'adventuring_gear', NULL, NULL, 30064771131, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803886, 'forgery-kit', 'Forgery Kit', 'adventuring_gear', NULL, NULL, 30064771132, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803887, 'herbalism-kit', 'Herbalism Kit', 'adventuring_gear', NULL, NULL, 30064771133, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803888, 'healers-kit', 'Healer''s Kit', 'adventuring_gear', NULL, NULL, 30064771134, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803889, 'mess-kit', 'Mess Kit', 'adventuring_gear', NULL, NULL, 30064771135, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803890, 'poisoners-kit', 'Poisoner''s Kit', 'adventuring_gear', NULL, NULL, 30064771136, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803891, 'ladder-10-foot', 'Ladder (10-foot)', 'adventuring_gear', NULL, NULL, 30064771137, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803892, 'lamp', 'Lamp', 'adventuring_gear', NULL, NULL, 30064771138, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803893, 'lantern-bullseye', 'Lantern, bullseye', 'adventuring_gear', NULL, NULL, 30064771139, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803894, 'lantern-hooded', 'Lantern, hooded', 'adventuring_gear', NULL, NULL, 30064771140, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803895, 'little-bag-of-sand', 'Little bag of sand', 'adventuring_gear', NULL, NULL, 30064771141, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803896, 'lock', 'Lock', 'adventuring_gear', NULL, NULL, 30064771142, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803897, 'magnifying-glass', 'Magnifying glass', 'adventuring_gear', NULL, NULL, 30064771143, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803898, 'manacles', 'Manacles', 'adventuring_gear', NULL, NULL, 30064771144, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803899, 'mirror-steel', 'Mirror, steel', 'adventuring_gear', NULL, NULL, 30064771145, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803900, 'oil-flask', 'Oil (flask)', 'adventuring_gear', NULL, NULL, 30064771146, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803901, 'paper-one-sheet', 'Paper (one sheet)', 'adventuring_gear', NULL, NULL, 30064771147, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803902, 'parchment-one-sheet', 'Parchment (one sheet)', 'adventuring_gear', NULL, NULL, 30064771148, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803903, 'perfume-vial', 'Perfume (vial)', 'adventuring_gear', NULL, NULL, 30064771149, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803904, 'pick-miners', 'Pick, miner''s', 'adventuring_gear', NULL, NULL, 30064771150, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803905, 'piton', 'Piton', 'adventuring_gear', NULL, NULL, 30064771151, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803906, 'poison-basic-vial', 'Poison, basic (vial)', 'adventuring_gear', NULL, NULL, 30064771152, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803907, 'pole-10-foot', 'Pole (10-foot)', 'adventuring_gear', NULL, NULL, 30064771153, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803908, 'pot-iron', 'Pot, iron', 'adventuring_gear', NULL, NULL, 30064771154, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803909, 'pouch', 'Pouch', 'adventuring_gear', NULL, NULL, 30064771155, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803910, 'quiver', 'Quiver', 'adventuring_gear', NULL, NULL, 30064771156, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803911, 'ram-portable', 'Ram, portable', 'adventuring_gear', NULL, NULL, 30064771157, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803912, 'rations-1-day', 'Rations (1 day)', 'adventuring_gear', NULL, NULL, 30064771158, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803913, 'reliquary', 'Reliquary', 'adventuring_gear', NULL, NULL, 30064771159, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803914, 'robes', 'Robes', 'adventuring_gear', NULL, NULL, 30064771160, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803915, 'rope-hempen-50-feet', 'Rope, hempen (50 feet)', 'adventuring_gear', NULL, NULL, 30064771161, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803916, 'rope-silk-50-feet', 'Rope, silk (50 feet)', 'adventuring_gear', NULL, NULL, 30064771162, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803917, 'sack', 'Sack', 'adventuring_gear', NULL, NULL, 30064771163, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803918, 'scale-merchants', 'Scale, merchant''s', 'adventuring_gear', NULL, NULL, 30064771164, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803919, 'sealing-wax', 'Sealing wax', 'adventuring_gear', NULL, NULL, 30064771165, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803920, 'shovel', 'Shovel', 'adventuring_gear', NULL, NULL, 30064771166, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803921, 'signal-whistle', 'Signal whistle', 'adventuring_gear', NULL, NULL, 30064771167, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803922, 'signet-ring', 'Signet ring', 'adventuring_gear', NULL, NULL, 30064771168, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803923, 'small-knife', 'Small knife', 'adventuring_gear', NULL, NULL, 30064771169, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803924, 'soap', 'Soap', 'adventuring_gear', NULL, NULL, 30064771170, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803925, 'spellbook', 'Spellbook', 'adventuring_gear', NULL, NULL, 30064771171, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803926, 'spike-iron', 'Spike, iron', 'adventuring_gear', NULL, NULL, 30064771172, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803927, 'spyglass', 'Spyglass', 'adventuring_gear', NULL, NULL, 30064771173, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803928, 'string-10-feet', 'String (10 feet)', 'adventuring_gear', NULL, NULL, 30064771174, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803929, 'tent-two-person', 'Tent, two-person', 'adventuring_gear', NULL, NULL, 30064771175, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803930, 'tinderbox', 'Tinderbox', 'adventuring_gear', NULL, NULL, 30064771176, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803931, 'torch', 'Torch', 'adventuring_gear', NULL, NULL, 30064771177, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803932, 'vestments', 'Vestments', 'adventuring_gear', NULL, NULL, 30064771178, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803933, 'vial', 'Vial', 'adventuring_gear', NULL, NULL, 30064771179, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803934, 'waterskin', 'Waterskin', 'adventuring_gear', NULL, NULL, 30064771180, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803935, 'whetstone', 'Whetstone', 'adventuring_gear', NULL, NULL, 30064771181, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803936, 'burglars-pack', 'Burglar''s Pack', 'adventuring_gear', NULL, NULL, 30064771182, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803937, 'diplomats-pack', 'Diplomat''s Pack', 'adventuring_gear', NULL, NULL, 30064771183, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803938, 'dungeoneers-pack', 'Dungeoneer''s Pack', 'adventuring_gear', NULL, NULL, 30064771184, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803939, 'entertainers-pack', 'Entertainer''s Pack', 'adventuring_gear', NULL, NULL, 30064771185, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803940, 'explorers-pack', 'Explorer''s Pack', 'adventuring_gear', NULL, NULL, 30064771186, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803941, 'priests-pack', 'Priest''s Pack', 'adventuring_gear', NULL, NULL, 30064771187, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803942, 'scholars-pack', 'Scholar''s Pack', 'adventuring_gear', NULL, NULL, 30064771188, NULL, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803943, 'alchemists-supplies', 'Alchemist''s Supplies', 'tools', NULL, NULL, NULL, 51539607553, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803944, 'brewers-supplies', 'Brewer''s Supplies', 'tools', NULL, NULL, NULL, 51539607554, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803945, 'calligraphers-supplies', 'Calligrapher''s Supplies', 'tools', NULL, NULL, NULL, 51539607555, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803946, 'carpenters-tools', 'Carpenter''s Tools', 'tools', NULL, NULL, NULL, 51539607556, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803947, 'cartographers-tools', 'Cartographer''s Tools', 'tools', NULL, NULL, NULL, 51539607557, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803948, 'cobblers-tools', 'Cobbler''s Tools', 'tools', NULL, NULL, NULL, 51539607558, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803949, 'cooks-utensils', 'Cook''s utensils', 'tools', NULL, NULL, NULL, 51539607559, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803950, 'glassblowers-tools', 'Glassblower''s Tools', 'tools', NULL, NULL, NULL, 51539607560, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803951, 'jewelers-tools', 'Jeweler''s Tools', 'tools', NULL, NULL, NULL, 51539607561, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803952, 'leatherworkers-tools', 'Leatherworker''s Tools', 'tools', NULL, NULL, NULL, 51539607562, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803953, 'masons-tools', 'Mason''s Tools', 'tools', NULL, NULL, NULL, 51539607563, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803954, 'painters-supplies', 'Painter''s Supplies', 'tools', NULL, NULL, NULL, 51539607564, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803955, 'potters-tools', 'Potter''s Tools', 'tools', NULL, NULL, NULL, 51539607565, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803956, 'smiths-tools', 'Smith''s Tools', 'tools', NULL, NULL, NULL, 51539607566, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803957, 'tinkers-tools', 'Tinker''s Tools', 'tools', NULL, NULL, NULL, 51539607567, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803958, 'weavers-tools', 'Weaver''s Tools', 'tools', NULL, NULL, NULL, 51539607568, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803959, 'woodcarvers-tools', 'Woodcarver''s Tools', 'tools', NULL, NULL, NULL, 51539607569, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803960, 'dice-set', 'Dice Set', 'tools', NULL, NULL, NULL, 51539607570, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803961, 'playing-card-set', 'Playing Card Set', 'tools', NULL, NULL, NULL, 51539607571, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803962, 'bagpipes', 'Bagpipes', 'tools', NULL, NULL, NULL, 51539607572, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803963, 'drum', 'Drum', 'tools', NULL, NULL, NULL, 51539607573, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803964, 'dulcimer', 'Dulcimer', 'tools', NULL, NULL, NULL, 51539607574, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803965, 'flute', 'Flute', 'tools', NULL, NULL, NULL, 51539607575, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803966, 'lute', 'Lute', 'tools', NULL, NULL, NULL, 51539607576, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803967, 'lyre', 'Lyre', 'tools', NULL, NULL, NULL, 51539607577, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803968, 'horn', 'Horn', 'tools', NULL, NULL, NULL, 51539607578, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803969, 'pan-flute', 'Pan flute', 'tools', NULL, NULL, NULL, 51539607579, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803970, 'shawm', 'Shawm', 'tools', NULL, NULL, NULL, 51539607580, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803971, 'viol', 'Viol', 'tools', NULL, NULL, NULL, 51539607581, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803972, 'navigators-tools', 'Navigator''s Tools', 'tools', NULL, NULL, NULL, 51539607582, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803973, 'thieves-tools', 'Thieves'' Tools', 'tools', NULL, NULL, NULL, 51539607583, NULL, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803974, 'camel', 'Camel', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574849, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803975, 'donkey', 'Donkey', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574850, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803976, 'mule', 'Mule', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574851, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803977, 'elephant', 'Elephant', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574852, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803978, 'horse-draft', 'Horse, draft', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574853, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803979, 'horse-riding', 'Horse, riding', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574854, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803980, 'mastiff', 'Mastiff', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574855, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803981, 'pony', 'Pony', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574856, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803982, 'warhorse', 'Warhorse', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574857, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803983, 'barding-padded', 'Barding: Padded', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574858, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803984, 'barding-leather', 'Barding: Leather', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574859, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803985, 'barding-studded-leather', 'Barding: Studded Leather', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574860, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803986, 'barding-hide', 'Barding: Hide', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574861, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803987, 'barding-chain-shirt', 'Barding: Chain shirt', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574862, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803988, 'barding-scale-mail', 'Barding: Scale mail', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574863, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803989, 'barding-breastplate', 'Barding: Breastplate', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574864, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803990, 'barding-half-plate', 'Barding: Half plate', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574865, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803991, 'barding-ring-mail', 'Barding: Ring mail', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574866, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803992, 'barding-chain-mail', 'Barding: Chain mail', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574867, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803993, 'barding-splint', 'Barding: Splint', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574868, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803994, 'barding-plate', 'Barding: Plate', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574869, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803995, 'bit-and-bridle', 'Bit and bridle', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574870, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803996, 'carriage', 'Carriage', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574871, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803997, 'cart', 'Cart', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574872, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803998, 'chariot', 'Chariot', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574873, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769803999, 'animal-feed-1-day', 'Animal Feed (1 day)', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574874, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769804000, 'saddle-exotic', 'Saddle, Exotic', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574875, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769804001, 'saddle-military', 'Saddle, Military', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574876, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769804002, 'saddle-pack', 'Saddle, Pack', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574877, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769804003, 'saddle-riding', 'Saddle, Riding', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574878, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769804004, 'saddlebags', 'Saddlebags', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574879, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769804005, 'sled', 'Sled', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574880, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769804006, 'stabling-1-day', 'Stabling (1 day)', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574881, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769804007, 'wagon', 'Wagon', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574882, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769804008, 'galley', 'Galley', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574883, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769804009, 'keelboat', 'Keelboat', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574884, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769804010, 'longship', 'Longship', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574885, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769804011, 'rowboat', 'Rowboat', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574886, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769804012, 'sailing-ship', 'Sailing ship', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574887, NULL);
INSERT INTO equipment (id, indx, name, equipment_category, equipment_weapon, equipment_armor, equipment_gear, equipment_tool, equipment_vehicle, equipment_cost) VALUES (25769804013, 'warship', 'Warship', 'mounts_and_vehicles', NULL, NULL, NULL, NULL, 55834574888, NULL);

-- Table: gears
CREATE TABLE IF NOT EXISTS gears (
    id            INTEGER NOT NULL
                          PRIMARY KEY AUTOINCREMENT,
    indx          TEXT    NOT NULL,
    name          TEXT    NOT NULL,
    gear_category TEXT    NOT NULL
                          DEFAULT 'other',
    desc          JSON    NOT NULL,
    quantity      INTEGER
);

INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771073, 'abacus', 'Abacus', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771074, 'acid-vial', 'Acid (vial)', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771075, 'alchemists-fire-flask', 'Alchemist''s fire (flask)', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771076, 'alms-box', 'Alms box', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771077, 'arrow', 'Arrow', 'ammunition', X'6E756C6C', 20);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771078, 'block-of-incense', 'Block of incense', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771079, 'blowgun-needle', 'Blowgun needle', 'ammunition', X'6E756C6C', 50);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771080, 'censer', 'Censer', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771081, 'crossbow-bolt', 'Crossbow bolt', 'ammunition', X'6E756C6C', 20);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771082, 'sling-bullet', 'Sling bullet', 'ammunition', X'6E756C6C', 20);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771083, 'amulet', 'Amulet', 'holy_symbols', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771084, 'antitoxin-vial', 'Antitoxin (vial)', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771085, 'crystal', 'Crystal', 'arcane_foci', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771086, 'orb', 'Orb', 'arcane_foci', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771087, 'rod', 'Rod', 'arcane_foci', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771088, 'staff', 'Staff', 'arcane_foci', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771089, 'wand', 'Wand', 'arcane_foci', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771090, 'backpack', 'Backpack', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771091, 'ball-bearings-bag-of-1000', 'Ball bearings (bag of 1,000)', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771092, 'barrel', 'Barrel', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771093, 'basket', 'Basket', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771094, 'bedroll', 'Bedroll', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771095, 'bell', 'Bell', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771096, 'blanket', 'Blanket', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771097, 'block-and-tackle', 'Block and tackle', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771098, 'book', 'Book', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771099, 'bottle-glass', 'Bottle, glass', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771100, 'bucket', 'Bucket', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771101, 'caltrops', 'Caltrops', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771102, 'candle', 'Candle', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771103, 'case-crossbow-bolt', 'Case, crossbow bolt', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771104, 'case-map-or-scroll', 'Case, map or scroll', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771105, 'chain-10-feet', 'Chain (10 feet)', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771106, 'chalk-1-piece', 'Chalk (1 piece)', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771107, 'chest', 'Chest', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771108, 'clothes-common', 'Clothes, common', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771109, 'clothes-costume', 'Clothes, costume', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771110, 'clothes-fine', 'Clothes, fine', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771111, 'clothes-travelers', 'Clothes, traveler''s', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771112, 'component-pouch', 'Component pouch', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771113, 'crowbar', 'Crowbar', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771114, 'sprig-of-mistletoe', 'Sprig of mistletoe', 'druidic_foci', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771115, 'totem', 'Totem', 'druidic_foci', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771116, 'wooden-staff', 'Wooden staff', 'druidic_foci', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771117, 'yew-wand', 'Yew wand', 'druidic_foci', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771118, 'emblem', 'Emblem', 'holy_symbols', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771119, 'fishing-tackle', 'Fishing tackle', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771120, 'flask-or-tankard', 'Flask or tankard', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771121, 'grappling-hook', 'Grappling hook', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771122, 'hammer', 'Hammer', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771123, 'hammer-sledge', 'Hammer, sledge', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771124, 'holy-water-flask', 'Holy water (flask)', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771125, 'hourglass', 'Hourglass', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771126, 'hunting-trap', 'Hunting trap', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771127, 'ink-1-ounce-bottle', 'Ink (1 ounce bottle)', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771128, 'ink-pen', 'Ink pen', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771129, 'jug-or-pitcher', 'Jug or pitcher', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771130, 'climbers-kit', 'Climber''s Kit', 'kits', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771131, 'disguise-kit', 'Disguise Kit', 'kits', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771132, 'forgery-kit', 'Forgery Kit', 'kits', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771133, 'herbalism-kit', 'Herbalism Kit', 'kits', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771134, 'healers-kit', 'Healer''s Kit', 'kits', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771135, 'mess-kit', 'Mess Kit', 'kits', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771136, 'poisoners-kit', 'Poisoner''s Kit', 'kits', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771137, 'ladder-10-foot', 'Ladder (10-foot)', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771138, 'lamp', 'Lamp', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771139, 'lantern-bullseye', 'Lantern, bullseye', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771140, 'lantern-hooded', 'Lantern, hooded', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771141, 'little-bag-of-sand', 'Little bag of sand', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771142, 'lock', 'Lock', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771143, 'magnifying-glass', 'Magnifying glass', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771144, 'manacles', 'Manacles', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771145, 'mirror-steel', 'Mirror, steel', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771146, 'oil-flask', 'Oil (flask)', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771147, 'paper-one-sheet', 'Paper (one sheet)', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771148, 'parchment-one-sheet', 'Parchment (one sheet)', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771149, 'perfume-vial', 'Perfume (vial)', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771150, 'pick-miners', 'Pick, miner''s', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771151, 'piton', 'Piton', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771152, 'poison-basic-vial', 'Poison, basic (vial)', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771153, 'pole-10-foot', 'Pole (10-foot)', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771154, 'pot-iron', 'Pot, iron', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771155, 'pouch', 'Pouch', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771156, 'quiver', 'Quiver', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771157, 'ram-portable', 'Ram, portable', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771158, 'rations-1-day', 'Rations (1 day)', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771159, 'reliquary', 'Reliquary', 'holy_symbols', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771160, 'robes', 'Robes', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771161, 'rope-hempen-50-feet', 'Rope, hempen (50 feet)', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771162, 'rope-silk-50-feet', 'Rope, silk (50 feet)', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771163, 'sack', 'Sack', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771164, 'scale-merchants', 'Scale, merchant''s', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771165, 'sealing-wax', 'Sealing wax', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771166, 'shovel', 'Shovel', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771167, 'signal-whistle', 'Signal whistle', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771168, 'signet-ring', 'Signet ring', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771169, 'small-knife', 'Small knife', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771170, 'soap', 'Soap', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771171, 'spellbook', 'Spellbook', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771172, 'spike-iron', 'Spike, iron', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771173, 'spyglass', 'Spyglass', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771174, 'string-10-feet', 'String (10 feet)', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771175, 'tent-two-person', 'Tent, two-person', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771176, 'tinderbox', 'Tinderbox', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771177, 'torch', 'Torch', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771178, 'vestments', 'Vestments', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771179, 'vial', 'Vial', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771180, 'waterskin', 'Waterskin', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771181, 'whetstone', 'Whetstone', 'standard_gear', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771182, 'burglars-pack', 'Burglar''s Pack', 'equipment_packs', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771183, 'diplomats-pack', 'Diplomat''s Pack', 'equipment_packs', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771184, 'dungeoneers-pack', 'Dungeoneer''s Pack', 'equipment_packs', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771185, 'entertainers-pack', 'Entertainer''s Pack', 'equipment_packs', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771186, 'explorers-pack', 'Explorer''s Pack', 'equipment_packs', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771187, 'priests-pack', 'Priest''s Pack', 'equipment_packs', X'6E756C6C', 0);
INSERT INTO gears (id, indx, name, gear_category, desc, quantity) VALUES (30064771188, 'scholars-pack', 'Scholar''s Pack', 'equipment_packs', X'6E756C6C', 0);

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

INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (34359738369, 'common', 'Common', '', 'STANDARD', 'Common');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (34359738370, 'dwarvish', 'Dwarvish', 'Dwarvish is full of hard consonants and guttural sounds.', 'STANDARD', 'Dwarvish');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (34359738371, 'elvish', 'Elvish', 'Elvish is fluid, with subtle intonations and intricate grammar. Elven literature is rich and varied, and their songs and poems are famous among other races. Many bards learn their language so they can add Elvish ballads to their repertoires.', 'STANDARD', 'Elvish');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (34359738372, 'giant', 'Giant', '', 'STANDARD', 'Dwarvish');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (34359738373, 'gnomish', 'Gnomish', 'The Gnomish language, which uses the Dwarvish script, is renowned for its technical treatises and its catalogs of knowledge about the natural world.', 'STANDARD', 'Dwarvish');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (34359738374, 'goblin', 'Goblin', '', 'STANDARD', 'Dwarvish');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (34359738375, 'halfling', 'Halfling', 'The Halfling language isn''t secret, but halflings are loath to share it with others. They write very little, so they don''t have a rich body of literature. Their oral tradition, however, is very strong.', 'STANDARD', 'Common');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (34359738376, 'orc', 'Orc', 'Orc is a harsh, grating language with hard consonants. It has no script of its own but is written in the Dwarvish script.', 'STANDARD', 'Dwarvish');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (34359738377, 'abyssal', 'Abyssal', '', 'EXOTIC', 'Infernal');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (34359738378, 'celestial', 'Celestial', '', 'EXOTIC', 'Celestial');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (34359738379, 'draconic', 'Draconic', 'Draconic is thought to be one of the oldest languages and is often used in the study of magic. The language sounds harsh to most other creatures and includes numerous hard consonants and sibilants.', 'EXOTIC', 'Draconic');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (34359738380, 'deep-speech', 'Deep Speech', 'An eerie tongue, Deep Speech is a clicking language with unsettling vocalizations that can almost sound like noise. Vibrating, creaking, and trembling sounds complement this mode of communication, making it hard for humanoid vocal cords to mimic.', 'EXOTIC', 'Other');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (34359738381, 'infernal', 'Infernal', '', 'EXOTIC', 'Infernal');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (34359738382, 'primordial', 'Primordial', '', 'EXOTIC', 'Dwarvish');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (34359738383, 'sylvan', 'Sylvan', '', 'EXOTIC', 'Elvish');
INSERT INTO languages (id, indx, name, desc, language_type, script) VALUES (34359738384, 'undercommon', 'Undercommon', '', 'EXOTIC', 'Elvish');

-- Table: proficiencies
CREATE TABLE IF NOT EXISTS proficiencies (
    id                   INTEGER NOT NULL
                                 PRIMARY KEY AUTOINCREMENT,
    indx                 TEXT    NOT NULL,
    name                 TEXT    NOT NULL,
    proficiency_category TEXT    NOT NULL
                                 DEFAULT 'other'
);

INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705665, 'light-armor', 'Light Armor', 'armor');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705666, 'medium-armor', 'Medium Armor', 'armor');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705667, 'heavy-armor', 'Heavy Armor', 'armor');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705668, 'all-armor', 'All armor', 'armor');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705669, 'padded-armor', 'Padded Armor', 'armor');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705670, 'leather-armor', 'Leather Armor', 'armor');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705671, 'studded-leather-armor', 'Studded Leather Armor', 'armor');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705672, 'hide-armor', 'Hide Armor', 'armor');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705673, 'chain-shirt', 'Chain Shirt', 'armor');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705674, 'scale-mail', 'Scale Mail', 'armor');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705675, 'breastplate', 'Breastplate', 'armor');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705676, 'half-plate-armor', 'Half Plate Armor', 'armor');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705677, 'ring-mail', 'Ring Mail', 'armor');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705678, 'chain-mail', 'Chain Mail', 'armor');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705679, 'splint-armor', 'Splint Armor', 'armor');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705680, 'plate-armor', 'Plate Armor', 'armor');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705681, 'shields', 'Shields', 'armor');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705682, 'simple-weapons', 'Simple Weapons', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705683, 'martial-weapons', 'Martial Weapons', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705684, 'clubs', 'Clubs', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705685, 'daggers', 'Daggers', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705686, 'greatclubs', 'Greatclubs', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705687, 'handaxes', 'Handaxes', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705688, 'javelins', 'Javelins', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705689, 'light-hammers', 'Light hammers', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705690, 'maces', 'Maces', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705691, 'quarterstaffs', 'Quarterstaffs', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705692, 'sickles', 'Sickles', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705693, 'spears', 'Spears', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705694, 'crossbows-light', 'Crossbows, light', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705695, 'darts', 'Darts', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705696, 'shortbows', 'Shortbows', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705697, 'slings', 'Slings', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705698, 'battleaxes', 'Battleaxes', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705699, 'flails', 'Flails', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705700, 'glaives', 'Glaives', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705701, 'greataxes', 'Greataxes', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705702, 'greatswords', 'Greatswords', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705703, 'halberds', 'Halberds', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705704, 'lances', 'Lances', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705705, 'longswords', 'Longswords', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705706, 'mauls', 'Mauls', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705707, 'morningstars', 'Morningstars', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705708, 'pikes', 'Pikes', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705709, 'rapiers', 'Rapiers', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705710, 'scimitars', 'Scimitars', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705711, 'shortswords', 'Shortswords', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705712, 'tridents', 'Tridents', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705713, 'war-picks', 'War picks', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705714, 'warhammers', 'Warhammers', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705715, 'whips', 'Whips', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705716, 'blowguns', 'Blowguns', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705717, 'hand-crossbows', 'Hand crossbows', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705718, 'crossbows-heavy', 'Crossbows, heavy', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705719, 'longbows', 'Longbows', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705720, 'nets', 'Nets', 'weapons');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705721, 'alchemists-supplies', 'Alchemist''s Supplies', 'artisans_tools');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705722, 'brewers-supplies', 'Brewer''s Supplies', 'artisans_tools');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705723, 'calligraphers-supplies', 'Calligrapher''s Supplies', 'artisans_tools');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705724, 'carpenters-tools', 'Carpenter''s Tools', 'artisans_tools');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705725, 'cartographers-tools', 'Cartographer''s Tools', 'artisans_tools');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705726, 'cobblers-tools', 'Cobbler''s Tools', 'artisans_tools');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705727, 'cooks-utensils', 'Cook''s utensils', 'artisans_tools');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705728, 'glassblowers-tools', 'Glassblower''s Tools', 'artisans_tools');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705729, 'jewelers-tools', 'Jeweler''s Tools', 'artisans_tools');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705730, 'leatherworkers-tools', 'Leatherworker''s Tools', 'artisans_tools');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705731, 'masons-tools', 'Mason''s Tools', 'artisans_tools');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705732, 'painters-supplies', 'Painter''s Supplies', 'artisans_tools');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705733, 'potters-tools', 'Potter''s Tools', 'artisans_tools');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705734, 'smiths-tools', 'Smith''s Tools', 'artisans_tools');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705735, 'tinkers-tools', 'Tinker''s Tools', 'artisans_tools');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705736, 'weavers-tools', 'Weaver''s Tools', 'artisans_tools');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705737, 'woodcarvers-tools', 'Woodcarver''s Tools', 'artisans_tools');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705738, 'disguise-kit', 'Disguise Kit', 'artisans_tools');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705739, 'forgery-kit', 'Forgery Kit', 'artisans_tools');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705740, 'dice-set', 'Dice Set', 'gaming_sets');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705741, 'playing-card-set', 'Playing Card Set', 'gaming_sets');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705742, 'bagpipes', 'Bagpipes', 'musical_instruments');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705743, 'drum', 'Drum', 'musical_instruments');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705744, 'dulcimer', 'Dulcimer', 'musical_instruments');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705745, 'flute', 'Flute', 'musical_instruments');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705746, 'lute', 'Lute', 'musical_instruments');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705747, 'lyre', 'Lyre', 'musical_instruments');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705748, 'horn', 'Horn', 'musical_instruments');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705749, 'pan-flute', 'Pan flute', 'musical_instruments');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705750, 'shawm', 'Shawm', 'musical_instruments');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705751, 'viol', 'Viol', 'musical_instruments');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705752, 'herbalism-kit', 'Herbalism Kit', 'other');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705753, 'navigators-tools', 'Navigator''s Tools', 'other');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705754, 'poisoners-kit', 'Poisoner''s Kit', 'other');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705755, 'thieves-tools', 'Thieves'' Tools', 'other');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705756, 'land-vehicles', 'Land Vehicles', 'vehicles');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705757, 'water-vehicles', 'Water Vehicles', 'vehicles');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705758, 'saving-throw-str', 'Saving Throw: STR', 'saving_throws');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705759, 'saving-throw-dex', 'Saving Throw: DEX', 'saving_throws');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705760, 'saving-throw-con', 'Saving Throw: CON', 'saving_throws');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705761, 'saving-throw-int', 'Saving Throw: INT', 'saving_throws');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705762, 'saving-throw-wis', 'Saving Throw: WIS', 'saving_throws');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705763, 'saving-throw-cha', 'Saving Throw: CHA', 'saving_throws');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705764, 'skill-acrobatics', 'Skill: Acrobatics', 'skills');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705765, 'skill-animal-handling', 'Skill: Animal Handling', 'skills');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705766, 'skill-arcana', 'Skill: Arcana', 'skills');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705767, 'skill-athletics', 'Skill: Athletics', 'skills');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705768, 'skill-deception', 'Skill: Deception', 'skills');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705769, 'skill-history', 'Skill: History', 'skills');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705770, 'skill-insight', 'Skill: Insight', 'skills');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705771, 'skill-intimidation', 'Skill: Intimidation', 'skills');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705772, 'skill-investigation', 'Skill: Investigation', 'skills');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705773, 'skill-medicine', 'Skill: Medicine', 'skills');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705774, 'skill-nature', 'Skill: Nature', 'skills');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705775, 'skill-perception', 'Skill: Perception', 'skills');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705776, 'skill-performance', 'Skill: Performance', 'skills');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705777, 'skill-persuasion', 'Skill: Persuasion', 'skills');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705778, 'skill-religion', 'Skill: Religion', 'skills');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705779, 'skill-sleight-of-hand', 'Skill: Sleight of Hand', 'skills');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705780, 'skill-stealth', 'Skill: Stealth', 'skills');
INSERT INTO proficiencies (id, indx, name, proficiency_category) VALUES (38654705781, 'skill-survival', 'Skill: Survival', 'skills');

-- Table: proficiency_classes
CREATE TABLE IF NOT EXISTS proficiency_classes (
    proficiency_id INTEGER NOT NULL,
    class_id       INTEGER NOT NULL,
    PRIMARY KEY (
        proficiency_id,
        class_id
    ),
    CONSTRAINT proficiency_classes_proficiency_id FOREIGN KEY (
        proficiency_id
    )
    REFERENCES proficiencies (id) ON DELETE CASCADE,
    CONSTRAINT proficiency_classes_class_id FOREIGN KEY (
        class_id
    )
    REFERENCES classes (id) ON DELETE CASCADE
);

INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705665, 12884901899);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705665, 12884901889);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705665, 12884901890);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705665, 12884901891);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705665, 12884901892);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705665, 12884901896);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705665, 12884901897);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705666, 12884901892);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705666, 12884901896);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705666, 12884901889);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705666, 12884901891);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705668, 12884901893);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705668, 12884901895);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705681, 12884901892);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705681, 12884901893);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705681, 12884901895);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705681, 12884901896);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705681, 12884901889);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705681, 12884901891);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705682, 12884901891);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705682, 12884901894);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705682, 12884901896);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705682, 12884901897);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705682, 12884901889);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705682, 12884901890);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705682, 12884901893);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705682, 12884901895);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705682, 12884901899);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705683, 12884901889);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705683, 12884901893);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705683, 12884901895);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705683, 12884901896);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705684, 12884901892);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705685, 12884901892);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705685, 12884901898);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705685, 12884901900);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705688, 12884901892);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705690, 12884901892);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705691, 12884901892);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705691, 12884901898);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705691, 12884901900);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705692, 12884901892);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705693, 12884901892);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705694, 12884901898);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705694, 12884901900);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705695, 12884901900);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705695, 12884901892);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705695, 12884901898);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705697, 12884901892);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705697, 12884901898);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705697, 12884901900);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705705, 12884901890);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705705, 12884901897);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705709, 12884901890);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705709, 12884901897);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705710, 12884901892);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705711, 12884901894);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705711, 12884901897);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705711, 12884901890);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705717, 12884901897);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705717, 12884901890);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705752, 12884901892);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705755, 12884901897);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705758, 12884901896);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705758, 12884901889);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705758, 12884901893);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705758, 12884901894);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705759, 12884901890);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705759, 12884901894);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705759, 12884901896);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705759, 12884901897);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705760, 12884901893);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705760, 12884901898);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705760, 12884901889);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705761, 12884901892);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705761, 12884901897);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705761, 12884901900);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705762, 12884901891);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705762, 12884901892);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705762, 12884901895);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705762, 12884901899);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705762, 12884901900);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705763, 12884901890);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705763, 12884901891);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705763, 12884901895);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705763, 12884901898);
INSERT INTO proficiency_classes (proficiency_id, class_id) VALUES (38654705763, 12884901899);

-- Table: proficiency_races
CREATE TABLE IF NOT EXISTS proficiency_races (
    proficiency_id INTEGER NOT NULL,
    race_id        INTEGER NOT NULL,
    PRIMARY KEY (
        proficiency_id,
        race_id
    ),
    CONSTRAINT proficiency_races_proficiency_id FOREIGN KEY (
        proficiency_id
    )
    REFERENCES proficiencies (id) ON DELETE CASCADE,
    CONSTRAINT proficiency_races_race_id FOREIGN KEY (
        race_id
    )
    REFERENCES races (id) ON DELETE CASCADE
);

INSERT INTO proficiency_races (proficiency_id, race_id) VALUES (38654705687, 42949672961);
INSERT INTO proficiency_races (proficiency_id, race_id) VALUES (38654705689, 42949672961);
INSERT INTO proficiency_races (proficiency_id, race_id) VALUES (38654705698, 42949672961);
INSERT INTO proficiency_races (proficiency_id, race_id) VALUES (38654705714, 42949672961);
INSERT INTO proficiency_races (proficiency_id, race_id) VALUES (38654705775, 42949672962);

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

INSERT INTO race_languages (race_id, language_id) VALUES (42949672961, 34359738369);
INSERT INTO race_languages (race_id, language_id) VALUES (42949672961, 34359738370);
INSERT INTO race_languages (race_id, language_id) VALUES (42949672962, 34359738369);
INSERT INTO race_languages (race_id, language_id) VALUES (42949672962, 34359738371);
INSERT INTO race_languages (race_id, language_id) VALUES (42949672963, 34359738369);
INSERT INTO race_languages (race_id, language_id) VALUES (42949672963, 34359738375);
INSERT INTO race_languages (race_id, language_id) VALUES (42949672964, 34359738369);
INSERT INTO race_languages (race_id, language_id) VALUES (42949672965, 34359738369);
INSERT INTO race_languages (race_id, language_id) VALUES (42949672965, 34359738379);
INSERT INTO race_languages (race_id, language_id) VALUES (42949672966, 34359738369);
INSERT INTO race_languages (race_id, language_id) VALUES (42949672966, 34359738373);
INSERT INTO race_languages (race_id, language_id) VALUES (42949672967, 34359738369);
INSERT INTO race_languages (race_id, language_id) VALUES (42949672967, 34359738371);
INSERT INTO race_languages (race_id, language_id) VALUES (42949672968, 34359738369);
INSERT INTO race_languages (race_id, language_id) VALUES (42949672968, 34359738376);
INSERT INTO race_languages (race_id, language_id) VALUES (42949672969, 34359738369);
INSERT INTO race_languages (race_id, language_id) VALUES (42949672969, 34359738381);

-- Table: races
CREATE TABLE IF NOT EXISTS races (
    id    INTEGER NOT NULL
                  PRIMARY KEY AUTOINCREMENT,
    indx  TEXT    NOT NULL,
    name  TEXT    NOT NULL,
    speed INTEGER NOT NULL
);

INSERT INTO races (id, indx, name, speed) VALUES (42949672961, 'dwarf', 'Dwarf', 25);
INSERT INTO races (id, indx, name, speed) VALUES (42949672962, 'elf', 'Elf', 30);
INSERT INTO races (id, indx, name, speed) VALUES (42949672963, 'halfling', 'Halfling', 25);
INSERT INTO races (id, indx, name, speed) VALUES (42949672964, 'human', 'Human', 30);
INSERT INTO races (id, indx, name, speed) VALUES (42949672965, 'dragonborn', 'Dragonborn', 30);
INSERT INTO races (id, indx, name, speed) VALUES (42949672966, 'gnome', 'Gnome', 25);
INSERT INTO races (id, indx, name, speed) VALUES (42949672967, 'half-elf', 'Half-Elf', 30);
INSERT INTO races (id, indx, name, speed) VALUES (42949672968, 'half-orc', 'Half-Orc', 30);
INSERT INTO races (id, indx, name, speed) VALUES (42949672969, 'tiefling', 'Tiefling', 30);

-- Table: skills
CREATE TABLE IF NOT EXISTS skills (
    id               INTEGER NOT NULL
                             PRIMARY KEY AUTOINCREMENT,
    indx             TEXT    NOT NULL,
    name             TEXT    NOT NULL,
    desc             JSON    NOT NULL,
    ability_score_id INTEGER NOT NULL,
    CONSTRAINT skills_ability_scores_ability_score FOREIGN KEY (
        ability_score_id
    )
    REFERENCES ability_scores (id) ON DELETE NO ACTION
);

INSERT INTO skills (id, indx, name, desc, ability_score_id) VALUES (47244640257, 'acrobatics', 'Acrobatics', X'5B22596F75722044657874657269747920284163726F6261746963732920636865636B20636F7665727320796F757220617474656D707420746F2073746179206F6E20796F7572206665657420696E206120747269636B7920736974756174696F6E2C2073756368206173207768656E20796F7527726520747279696E6720746F2072756E206163726F73732061207368656574206F66206963652C2062616C616E6365206F6E2061207469676874726F70652C206F7220737461792075707269676874206F6E206120726F636B696E6720736869702773206465636B2E2054686520474D206D6967687420616C736F2063616C6C20666F7220612044657874657269747920284163726F6261746963732920636865636B20746F2073656520696620796F752063616E20706572666F726D206163726F6261746963207374756E74732C20696E636C7564696E672064697665732C20726F6C6C732C20736F6D65727361756C74732C20616E6420666C6970732E225D', 2);
INSERT INTO skills (id, indx, name, desc, ability_score_id) VALUES (47244640258, 'animal-handling', 'Animal Handling', X'5B225768656E20746865726520697320616E79207175657374696F6E207768657468657220796F752063616E2063616C6D20646F776E206120646F6D65737469636174656420616E696D616C2C206B6565702061206D6F756E742066726F6D2067657474696E672073706F6F6B65642C206F7220696E7475697420616E20616E696D616C277320696E74656E74696F6E732C2074686520474D206D696768742063616C6C20666F72206120576973646F6D2028416E696D616C2048616E646C696E672920636865636B2E20596F7520616C736F206D616B65206120576973646F6D2028416E696D616C2048616E646C696E672920636865636B20746F20636F6E74726F6C20796F7572206D6F756E74207768656E20796F7520617474656D70742061207269736B79206D616E65757665722E225D', 5);
INSERT INTO skills (id, indx, name, desc, ability_score_id) VALUES (47244640259, 'arcana', 'Arcana', X'5B22596F757220496E74656C6C6967656E63652028417263616E612920636865636B206D6561737572657320796F7572206162696C69747920746F20726563616C6C206C6F72652061626F7574207370656C6C732C206D61676963206974656D732C20656C6472697463682073796D626F6C732C206D61676963616C20747261646974696F6E732C2074686520706C616E6573206F66206578697374656E63652C20616E642074686520696E6861626974616E7473206F662074686F736520706C616E65732E225D', 4);
INSERT INTO skills (id, indx, name, desc, ability_score_id) VALUES (47244640260, 'athletics', 'Athletics', X'5B22596F757220537472656E67746820284174686C65746963732920636865636B20636F7665727320646966666963756C7420736974756174696F6E7320796F7520656E636F756E746572207768696C6520636C696D62696E672C206A756D70696E672C206F72207377696D6D696E672E225D', 1);
INSERT INTO skills (id, indx, name, desc, ability_score_id) VALUES (47244640261, 'deception', 'Deception', X'5B22596F7572204368617269736D612028446563657074696F6E2920636865636B2064657465726D696E6573207768657468657220796F752063616E20636F6E76696E63696E676C792068696465207468652074727574682C206569746865722076657262616C6C79206F72207468726F75676820796F757220616374696F6E732E205468697320646563657074696F6E2063616E20656E636F6D706173732065766572797468696E672066726F6D206D69736C656164696E67206F7468657273207468726F75676820616D6269677569747920746F2074656C6C696E67206F75747269676874206C6965732E205479706963616C20736974756174696F6E7320696E636C75646520747279696E6720746F20666173742D2074616C6B20612067756172642C20636F6E2061206D65726368616E742C206561726E206D6F6E6579207468726F7567682067616D626C696E672C207061737320796F757273656C66206F666620696E20612064697367756973652C2064756C6C20736F6D656F6E65277320737573706963696F6E7320776974682066616C7365206173737572616E6365732C206F72206D61696E7461696E20612073747261696768742066616365207768696C652074656C6C696E67206120626C6174616E74206C69652E225D', 6);
INSERT INTO skills (id, indx, name, desc, ability_score_id) VALUES (47244640262, 'history', 'History', X'5B22596F757220496E74656C6C6967656E63652028486973746F72792920636865636B206D6561737572657320796F7572206162696C69747920746F20726563616C6C206C6F72652061626F757420686973746F726963616C206576656E74732C206C6567656E646172792070656F706C652C20616E6369656E74206B696E67646F6D732C20706173742064697370757465732C20726563656E7420776172732C20616E64206C6F737420636976696C697A6174696F6E732E225D', 4);
INSERT INTO skills (id, indx, name, desc, ability_score_id) VALUES (47244640263, 'insight', 'Insight', X'5B22596F757220576973646F6D2028496E73696768742920636865636B2064656369646573207768657468657220796F752063616E2064657465726D696E6520746865207472756520696E74656E74696F6E73206F6620612063726561747572652C2073756368206173207768656E20736561726368696E67206F75742061206C6965206F722070726564696374696E6720736F6D656F6E652773206E657874206D6F76652E20446F696E6720736F20696E766F6C76657320676C65616E696E6720636C7565732066726F6D20626F6479206C616E67756167652C20737065656368206861626974732C20616E64206368616E67657320696E206D616E6E657269736D732E225D', 5);
INSERT INTO skills (id, indx, name, desc, ability_score_id) VALUES (47244640264, 'intimidation', 'Intimidation', X'5B225768656E20796F7520617474656D707420746F20696E666C75656E636520736F6D656F6E65207468726F756768206F7665727420746872656174732C20686F7374696C6520616374696F6E732C20616E6420706879736963616C2076696F6C656E63652C2074686520474D206D696768742061736B20796F7520746F206D616B652061204368617269736D612028496E74696D69646174696F6E2920636865636B2E204578616D706C657320696E636C75646520747279696E6720746F2070727920696E666F726D6174696F6E206F7574206F66206120707269736F6E65722C20636F6E76696E63696E672073747265657420746875677320746F206261636B20646F776E2066726F6D206120636F6E66726F6E746174696F6E2C206F72207573696E67207468652065646765206F6620612062726F6B656E20626F74746C6520746F20636F6E76696E6365206120736E656572696E672076697A69657220746F207265636F6E73696465722061206465636973696F6E2E225D', 6);
INSERT INTO skills (id, indx, name, desc, ability_score_id) VALUES (47244640265, 'investigation', 'Investigation', X'5B225768656E20796F75206C6F6F6B2061726F756E6420666F7220636C75657320616E64206D616B6520646564756374696F6E73206261736564206F6E2074686F736520636C7565732C20796F75206D616B6520616E20496E74656C6C6967656E63652028496E7665737469676174696F6E2920636865636B2E20596F75206D696768742064656475636520746865206C6F636174696F6E206F6620612068696464656E206F626A6563742C206469736365726E2066726F6D2074686520617070656172616E6365206F66206120776F756E642077686174206B696E64206F6620776561706F6E206465616C742069742C206F722064657465726D696E6520746865207765616B65737420706F696E7420696E20612074756E6E656C207468617420636F756C6420636175736520697420746F20636F6C6C617073652E20506F72696E67207468726F75676820616E6369656E74207363726F6C6C7320696E20736561726368206F6620612068696464656E20667261676D656E74206F66206B6E6F776C65646765206D6967687420616C736F2063616C6C20666F7220616E20496E74656C6C6967656E63652028496E7665737469676174696F6E2920636865636B2E225D', 4);
INSERT INTO skills (id, indx, name, desc, ability_score_id) VALUES (47244640266, 'medicine', 'Medicine', X'5B224120576973646F6D20284D65646963696E652920636865636B206C65747320796F752074727920746F2073746162696C697A652061206479696E6720636F6D70616E696F6E206F7220646961676E6F736520616E20696C6C6E6573732E225D', 5);
INSERT INTO skills (id, indx, name, desc, ability_score_id) VALUES (47244640267, 'nature', 'Nature', X'5B22596F757220496E74656C6C6967656E636520284E61747572652920636865636B206D6561737572657320796F7572206162696C69747920746F20726563616C6C206C6F72652061626F7574207465727261696E2C20706C616E747320616E6420616E696D616C732C2074686520776561746865722C20616E64206E61747572616C206379636C65732E225D', 4);
INSERT INTO skills (id, indx, name, desc, ability_score_id) VALUES (47244640268, 'perception', 'Perception', X'5B22596F757220576973646F6D202850657263657074696F6E2920636865636B206C65747320796F752073706F742C20686561722C206F72206F746865727769736520646574656374207468652070726573656E6365206F6620736F6D657468696E672E204974206D6561737572657320796F75722067656E6572616C2061776172656E657373206F6620796F757220737572726F756E64696E677320616E6420746865206B65656E6E657373206F6620796F75722073656E7365732E20466F72206578616D706C652C20796F75206D696768742074727920746F2068656172206120636F6E766572736174696F6E207468726F756768206120636C6F73656420646F6F722C20656176657364726F7020756E64657220616E206F70656E2077696E646F772C206F722068656172206D6F6E7374657273206D6F76696E6720737465616C7468696C7920696E2074686520666F726573742E204F7220796F75206D696768742074727920746F2073706F74207468696E6773207468617420617265206F62736375726564206F72206561737920746F206D6973732C2077686574686572207468657920617265206F726373206C79696E6720696E20616D62757368206F6E206120726F61642C20746875677320686964696E6720696E2074686520736861646F7773206F6620616E20616C6C65792C206F722063616E646C656C6967687420756E646572206120636C6F7365642073656372657420646F6F722E225D', 5);
INSERT INTO skills (id, indx, name, desc, ability_score_id) VALUES (47244640269, 'performance', 'Performance', X'5B22596F7572204368617269736D612028506572666F726D616E63652920636865636B2064657465726D696E657320686F772077656C6C20796F752063616E2064656C6967687420616E2061756469656E63652077697468206D757369632C2064616E63652C20616374696E672C2073746F727974656C6C696E672C206F7220736F6D65206F7468657220666F726D206F6620656E7465727461696E6D656E742E225D', 6);
INSERT INTO skills (id, indx, name, desc, ability_score_id) VALUES (47244640270, 'persuasion', 'Persuasion', X'5B225768656E20796F7520617474656D707420746F20696E666C75656E636520736F6D656F6E65206F7220612067726F7570206F662070656F706C65207769746820746163742C20736F6369616C206772616365732C206F7220676F6F64206E61747572652C2074686520474D206D696768742061736B20796F7520746F206D616B652061204368617269736D61202850657273756173696F6E2920636865636B2E205479706963616C6C792C20796F75207573652070657273756173696F6E207768656E20616374696E6720696E20676F6F642066616974682C20746F20666F7374657220667269656E6473686970732C206D616B6520636F726469616C2072657175657374732C206F7220657868696269742070726F706572206574697175657474652E204578616D706C6573206F662070657273756164696E67206F746865727320696E636C75646520636F6E76696E63696E672061206368616D6265726C61696E20746F206C657420796F75722070617274792073656520746865206B696E672C206E65676F74696174696E67207065616365206265747765656E2077617272696E67207472696265732C206F7220696E73706972696E6720612063726F7764206F6620746F776E73666F6C6B2E225D', 6);
INSERT INTO skills (id, indx, name, desc, ability_score_id) VALUES (47244640271, 'religion', 'Religion', X'5B22596F757220496E74656C6C6967656E6365202852656C6967696F6E2920636865636B206D6561737572657320796F7572206162696C69747920746F20726563616C6C206C6F72652061626F757420646569746965732C20726974657320616E6420707261796572732C2072656C6967696F75732068696572617263686965732C20686F6C792073796D626F6C732C20616E642074686520707261637469636573206F66207365637265742063756C74732E225D', 4);
INSERT INTO skills (id, indx, name, desc, ability_score_id) VALUES (47244640272, 'sleight-of-hand', 'Sleight of Hand', X'5B225768656E6576657220796F7520617474656D707420616E20616374206F66206C6567657264656D61696E206F72206D616E75616C20747269636B6572792C207375636820617320706C616E74696E6720736F6D657468696E67206F6E20736F6D656F6E6520656C7365206F7220636F6E6365616C696E6720616E206F626A656374206F6E20796F757220706572736F6E2C206D616B652061204465787465726974792028536C6569676874206F662048616E642920636865636B2E2054686520474D206D6967687420616C736F2063616C6C20666F722061204465787465726974792028536C6569676874206F662048616E642920636865636B20746F2064657465726D696E65207768657468657220796F752063616E206C696674206120636F696E207075727365206F666620616E6F7468657220706572736F6E206F7220736C697020736F6D657468696E67206F7574206F6620616E6F7468657220706572736F6E277320706F636B65742E225D', 2);
INSERT INTO skills (id, indx, name, desc, ability_score_id) VALUES (47244640273, 'stealth', 'Stealth', X'5B224D616B652061204465787465726974792028537465616C74682920636865636B207768656E20796F7520617474656D707420746F20636F6E6365616C20796F757273656C662066726F6D20656E656D6965732C20736C696E6B2070617374206775617264732C20736C6970206177617920776974686F7574206265696E67206E6F74696365642C206F7220736E65616B207570206F6E20736F6D656F6E6520776974686F7574206265696E67207365656E206F722068656172642E225D', 2);
INSERT INTO skills (id, indx, name, desc, ability_score_id) VALUES (47244640274, 'survival', 'Survival', X'5B2254686520474D206D696768742061736B20796F7520746F206D616B65206120576973646F6D2028537572766976616C2920636865636B20746F20666F6C6C6F7720747261636B732C2068756E742077696C642067616D652C20677569646520796F75722067726F7570207468726F7567682066726F7A656E2077617374656C616E64732C206964656E74696679207369676E732074686174206F776C6265617273206C697665206E65617262792C20707265646963742074686520776561746865722C206F722061766F696420717569636B73616E6420616E64206F74686572206E61747572616C2068617A617264732E225D', 5);

-- Table: tools
CREATE TABLE IF NOT EXISTS tools (
    id            INTEGER NOT NULL
                          PRIMARY KEY AUTOINCREMENT,
    indx          TEXT    NOT NULL,
    name          TEXT    NOT NULL,
    tool_category TEXT    NOT NULL
);

INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607553, 'alchemists-supplies', 'Alchemist''s Supplies', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607554, 'brewers-supplies', 'Brewer''s Supplies', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607555, 'calligraphers-supplies', 'Calligrapher''s Supplies', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607556, 'carpenters-tools', 'Carpenter''s Tools', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607557, 'cartographers-tools', 'Cartographer''s Tools', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607558, 'cobblers-tools', 'Cobbler''s Tools', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607559, 'cooks-utensils', 'Cook''s utensils', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607560, 'glassblowers-tools', 'Glassblower''s Tools', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607561, 'jewelers-tools', 'Jeweler''s Tools', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607562, 'leatherworkers-tools', 'Leatherworker''s Tools', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607563, 'masons-tools', 'Mason''s Tools', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607564, 'painters-supplies', 'Painter''s Supplies', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607565, 'potters-tools', 'Potter''s Tools', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607566, 'smiths-tools', 'Smith''s Tools', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607567, 'tinkers-tools', 'Tinker''s Tools', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607568, 'weavers-tools', 'Weaver''s Tools', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607569, 'woodcarvers-tools', 'Woodcarver''s Tools', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607570, 'dice-set', 'Dice Set', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607571, 'playing-card-set', 'Playing Card Set', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607572, 'bagpipes', 'Bagpipes', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607573, 'drum', 'Drum', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607574, 'dulcimer', 'Dulcimer', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607575, 'flute', 'Flute', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607576, 'lute', 'Lute', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607577, 'lyre', 'Lyre', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607578, 'horn', 'Horn', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607579, 'pan-flute', 'Pan flute', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607580, 'shawm', 'Shawm', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607581, 'viol', 'Viol', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607582, 'navigators-tools', 'Navigator''s Tools', '');
INSERT INTO tools (id, indx, name, tool_category) VALUES (51539607583, 'thieves-tools', 'Thieves'' Tools', '');

-- Table: vehicles
CREATE TABLE IF NOT EXISTS vehicles (
    id               INTEGER NOT NULL
                             PRIMARY KEY AUTOINCREMENT,
    indx             TEXT    NOT NULL,
    name             TEXT    NOT NULL,
    vehicle_category TEXT    NOT NULL,
    capacity         TEXT    NOT NULL
);

INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574849, 'camel', 'Camel', 'Mounts and Other Animals', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574850, 'donkey', 'Donkey', 'Mounts and Other Animals', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574851, 'mule', 'Mule', 'Mounts and Other Animals', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574852, 'elephant', 'Elephant', 'Mounts and Other Animals', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574853, 'horse-draft', 'Horse, draft', 'Mounts and Other Animals', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574854, 'horse-riding', 'Horse, riding', 'Mounts and Other Animals', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574855, 'mastiff', 'Mastiff', 'Mounts and Other Animals', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574856, 'pony', 'Pony', 'Mounts and Other Animals', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574857, 'warhorse', 'Warhorse', 'Mounts and Other Animals', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574858, 'barding-padded', 'Barding: Padded', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574859, 'barding-leather', 'Barding: Leather', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574860, 'barding-studded-leather', 'Barding: Studded Leather', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574861, 'barding-hide', 'Barding: Hide', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574862, 'barding-chain-shirt', 'Barding: Chain shirt', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574863, 'barding-scale-mail', 'Barding: Scale mail', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574864, 'barding-breastplate', 'Barding: Breastplate', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574865, 'barding-half-plate', 'Barding: Half plate', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574866, 'barding-ring-mail', 'Barding: Ring mail', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574867, 'barding-chain-mail', 'Barding: Chain mail', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574868, 'barding-splint', 'Barding: Splint', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574869, 'barding-plate', 'Barding: Plate', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574870, 'bit-and-bridle', 'Bit and bridle', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574871, 'carriage', 'Carriage', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574872, 'cart', 'Cart', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574873, 'chariot', 'Chariot', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574874, 'animal-feed-1-day', 'Animal Feed (1 day)', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574875, 'saddle-exotic', 'Saddle, Exotic', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574876, 'saddle-military', 'Saddle, Military', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574877, 'saddle-pack', 'Saddle, Pack', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574878, 'saddle-riding', 'Saddle, Riding', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574879, 'saddlebags', 'Saddlebags', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574880, 'sled', 'Sled', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574881, 'stabling-1-day', 'Stabling (1 day)', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574882, 'wagon', 'Wagon', 'Tack, Harness, and Drawn Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574883, 'galley', 'Galley', 'Waterborne Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574884, 'keelboat', 'Keelboat', 'Waterborne Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574885, 'longship', 'Longship', 'Waterborne Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574886, 'rowboat', 'Rowboat', 'Waterborne Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574887, 'sailing-ship', 'Sailing ship', 'Waterborne Vehicles', '');
INSERT INTO vehicles (id, indx, name, vehicle_category, capacity) VALUES (55834574888, 'warship', 'Warship', 'Waterborne Vehicles', '');

-- Table: weapon_damages
CREATE TABLE IF NOT EXISTS weapon_damages (
    id   INTEGER NOT NULL
                 PRIMARY KEY AUTOINCREMENT,
    dice TEXT    NOT NULL
);


-- Table: weapons
CREATE TABLE IF NOT EXISTS weapons (
    id           INTEGER NOT NULL
                         PRIMARY KEY AUTOINCREMENT,
    indx         TEXT    NOT NULL,
    name         TEXT    NOT NULL,
    weapon_range TEXT    NOT NULL
);

INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542145, 'club', 'Club', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542146, 'dagger', 'Dagger', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542147, 'greatclub', 'Greatclub', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542148, 'handaxe', 'Handaxe', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542149, 'javelin', 'Javelin', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542150, 'light-hammer', 'Light hammer', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542151, 'mace', 'Mace', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542152, 'quarterstaff', 'Quarterstaff', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542153, 'sickle', 'Sickle', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542154, 'spear', 'Spear', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542155, 'crossbow-light', 'Crossbow, light', 'Ranged');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542156, 'dart', 'Dart', 'Ranged');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542157, 'shortbow', 'Shortbow', 'Ranged');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542158, 'sling', 'Sling', 'Ranged');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542159, 'battleaxe', 'Battleaxe', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542160, 'flail', 'Flail', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542161, 'glaive', 'Glaive', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542162, 'greataxe', 'Greataxe', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542163, 'greatsword', 'Greatsword', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542164, 'halberd', 'Halberd', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542165, 'lance', 'Lance', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542166, 'longsword', 'Longsword', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542167, 'maul', 'Maul', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542168, 'morningstar', 'Morningstar', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542169, 'pike', 'Pike', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542170, 'rapier', 'Rapier', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542171, 'scimitar', 'Scimitar', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542172, 'shortsword', 'Shortsword', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542173, 'trident', 'Trident', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542174, 'war-pick', 'War pick', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542175, 'warhammer', 'Warhammer', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542176, 'whip', 'Whip', 'Melee');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542177, 'blowgun', 'Blowgun', 'Ranged');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542178, 'crossbow-hand', 'Crossbow, hand', 'Ranged');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542179, 'crossbow-heavy', 'Crossbow, heavy', 'Ranged');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542180, 'longbow', 'Longbow', 'Ranged');
INSERT INTO weapons (id, indx, name, weapon_range) VALUES (60129542181, 'net', 'Net', 'Ranged');

-- Index: ability_scores_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS ability_scores_indx_key ON ability_scores (
    indx
);


-- Index: armors_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS armors_indx_key ON armors (
    indx
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


-- Index: gears_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS gears_indx_key ON gears (
    indx
);


-- Index: languages_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS languages_indx_key ON languages (
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


-- Index: skills_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS skills_indx_key ON skills (
    indx
);


-- Index: tools_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS tools_indx_key ON tools (
    indx
);


-- Index: vehicles_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS vehicles_indx_key ON vehicles (
    indx
);


-- Index: weapons_indx_key
CREATE UNIQUE INDEX IF NOT EXISTS weapons_indx_key ON weapons (
    indx
);


PRAGMA foreign_keys = on;
