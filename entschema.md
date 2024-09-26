AbilityBonus:
	+-------+------+--------+----------+----------+---------+---------------+-----------+------------------------+------------+---------+
	| Field | Type | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag        | Validators | Comment |
	+-------+------+--------+----------+----------+---------+---------------+-----------+------------------------+------------+---------+
	| id    | int  | false  | false    | false    | false   | false         | false     | json:"id,omitempty"    |          0 |         |
	| bonus | int  | false  | false    | false    | false   | false         | false     | json:"bonus,omitempty" |          1 |         |
	+-------+------+--------+----------+----------+---------+---------------+-----------+------------------------+------------+---------+
	+---------------+--------------------+---------+-----------------+----------+--------+----------+---------+
	|     Edge      |        Type        | Inverse |     BackRef     | Relation | Unique | Optional | Comment |
	+---------------+--------------------+---------+-----------------+----------+--------+----------+---------+
	| ability_score | AbilityScore       | false   |                 | M2O      | true   | false    |         |
	| race          | Race               | true    | ability_bonuses | M2M      | false  | true     |         |
	| options       | AbilityBonusChoice | true    | ability_bonuses | M2M      | false  | true     |         |
	| subrace       | Subrace            | true    | ability_bonuses | M2M      | false  | true     |         |
	+---------------+--------------------+---------+-----------------+----------+--------+----------+---------+
	
AbilityBonusChoice:
	+--------+------+--------+----------+----------+---------+---------------+-----------+-------------------------+------------+---------+
	| Field  | Type | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |        StructTag        | Validators | Comment |
	+--------+------+--------+----------+----------+---------+---------------+-----------+-------------------------+------------+---------+
	| id     | int  | false  | false    | false    | false   | false         | false     | json:"id,omitempty"     |          0 |         |
	| choose | int  | false  | false    | false    | false   | false         | false     | json:"choose,omitempty" |          1 |         |
	+--------+------+--------+----------+----------+---------+---------------+-----------+-------------------------+------------+---------+
	+-----------------+--------------+---------+-----------------------+----------+--------+----------+---------+
	|      Edge       |     Type     | Inverse |        BackRef        | Relation | Unique | Optional | Comment |
	+-----------------+--------------+---------+-----------------------+----------+--------+----------+---------+
	| ability_bonuses | AbilityBonus | false   |                       | M2M      | false  | true     |         |
	| race            | Race         | true    | ability_bonus_options | O2M      | false  | true     |         |
	+-----------------+--------------+---------+-----------------------+----------+--------+----------+---------+
	
AbilityScore:
	+-----------+----------+--------+----------+----------+---------+---------------+-----------+----------------------------+------------+---------+
	|   Field   |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |         StructTag          | Validators | Comment |
	+-----------+----------+--------+----------+----------+---------+---------------+-----------+----------------------------+------------+---------+
	| id        | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"        |          0 |         |
	| indx      | string   | true   | false    | false    | false   | false         | false     | json:"index"               |          1 |         |
	| name      | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty"      |          1 |         |
	| desc      | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty"      |          0 |         |
	| full_name | string   | false  | false    | false    | false   | false         | false     | json:"full_name,omitempty" |          0 |         |
	+-----------+----------+--------+----------+----------+---------+---------------+-----------+----------------------------+------------+---------+
	+-----------------+--------------+---------+---------------+----------+--------+----------+---------+
	|      Edge       |     Type     | Inverse |    BackRef    | Relation | Unique | Optional | Comment |
	+-----------------+--------------+---------+---------------+----------+--------+----------+---------+
	| skills          | Skill        | false   |               | O2M      | false  | true     |         |
	| ability_bonuses | AbilityBonus | true    | ability_score | O2M      | false  | true     |         |
	| classes         | Class        | true    | saving_throws | M2M      | false  | true     |         |
	+-----------------+--------------+---------+---------------+----------+--------+----------+---------+
	
Alignment:
	+-------+----------------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |      Type      | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+----------------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int            | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string         | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string         | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| desc  | []string       | false  | true     | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	| abbr  | alignment.Abbr | false  | false    | false    | false   | false         | false     | json:"abbreviation"   |          0 |         |
	+-------+----------------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	
Armor:
	+----------------------+---------------------+--------+----------+----------+---------+---------------+-----------+---------------------------------------+------------+---------+
	|        Field         |        Type         | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |               StructTag               | Validators | Comment |
	+----------------------+---------------------+--------+----------+----------+---------+---------------+-----------+---------------------------------------+------------+---------+
	| id                   | int                 | false  | false    | false    | false   | false         | false     | json:"id,omitempty"                   |          0 |         |
	| armor_category       | armor.ArmorCategory | false  | false    | false    | false   | false         | false     | json:"armor_category,omitempty"       |          0 |         |
	| str_minimum          | int                 | false  | false    | false    | false   | false         | false     | json:"str_minimum,omitempty"          |          0 |         |
	| stealth_disadvantage | bool                | false  | false    | false    | false   | false         | false     | json:"stealth_disadvantage,omitempty" |          0 |         |
	| ac_base              | int                 | false  | false    | false    | false   | false         | false     | json:"ac_base,omitempty"              |          1 |         |
	| ac_dex_bonus         | bool                | false  | false    | false    | true    | false         | false     | json:"ac_dex_bonus,omitempty"         |          0 |         |
	| ac_max_bonus         | int                 | false  | false    | false    | true    | false         | false     | json:"ac_max_bonus,omitempty"         |          0 |         |
	+----------------------+---------------------+--------+----------+----------+---------+---------------+-----------+---------------------------------------+------------+---------+
	+-----------+-----------+---------+---------+----------+--------+----------+---------+
	|   Edge    |   Type    | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+-----------+-----------+---------+---------+----------+--------+----------+---------+
	| equipment | Equipment | true    | armor   | O2O      | true   | false    |         |
	+-----------+-----------+---------+---------+----------+--------+----------+---------+
	
Character:
	+-------+--------+--------+----------+----------+---------+---------------+-----------+------------------------+------------+---------+
	| Field |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag        | Validators | Comment |
	+-------+--------+--------+----------+----------+---------+---------------+-----------+------------------------+------------+---------+
	| id    | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"    |          0 |         |
	| name  | string | false  | false    | false    | false   | false         | false     | json:"name,omitempty"  |          1 |         |
	| age   | int    | false  | false    | false    | true    | false         | false     | json:"age,omitempty"   |          1 |         |
	| level | int    | false  | false    | false    | true    | false         | false     | json:"level,omitempty" |          1 |         |
	+-------+--------+--------+----------+----------+---------+---------------+-----------+------------------------+------------+---------+
	+----------------+-----------------------+---------+-----------+----------+--------+----------+---------+
	|      Edge      |         Type          | Inverse |  BackRef  | Relation | Unique | Optional | Comment |
	+----------------+-----------------------+---------+-----------+----------+--------+----------+---------+
	| race           | Race                  | false   |           | M2O      | true   | true     |         |
	| class          | Class                 | false   |           | M2O      | true   | true     |         |
	| alignment      | Alignment             | false   |           | M2O      | true   | true     |         |
	| traits         | Trait                 | false   |           | O2M      | false  | true     |         |
	| languages      | Language              | false   |           | O2M      | false  | true     |         |
	| proficiencies  | Proficiency           | false   |           | O2M      | false  | true     |         |
	| ability_scores | CharacterAbilityScore | true    | character | O2M      | false  | true     |         |
	+----------------+-----------------------+---------+-----------+----------+--------+----------+---------+
	
CharacterAbilityScore:
	+-------+------+--------+----------+----------+---------+---------------+-----------+------------------------+------------+---------+
	| Field | Type | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag        | Validators | Comment |
	+-------+------+--------+----------+----------+---------+---------------+-----------+------------------------+------------+---------+
	| id    | int  | false  | false    | false    | false   | false         | false     | json:"id,omitempty"    |          0 |         |
	| score | int  | false  | false    | false    | false   | false         | false     | json:"score,omitempty" |          1 |         |
	+-------+------+--------+----------+----------+---------+---------------+-----------+------------------------+------------+---------+
	+---------------+--------------+---------+---------+----------+--------+----------+---------+
	|     Edge      |     Type     | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+---------------+--------------+---------+---------+----------+--------+----------+---------+
	| character     | Character    | false   |         | M2O      | true   | true     |         |
	| ability_score | AbilityScore | false   |         | M2O      | true   | false    |         |
	+---------------+--------------+---------+---------+----------+--------+----------+---------+
	
Class:
	+---------+--------+--------+----------+----------+---------+---------------+-----------+--------------------------+------------+---------+
	|  Field  |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |        StructTag         | Validators | Comment |
	+---------+--------+--------+----------+----------+---------+---------------+-----------+--------------------------+------------+---------+
	| id      | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"      |          0 |         |
	| indx    | string | true   | false    | false    | false   | false         | false     | json:"index"             |          1 |         |
	| name    | string | false  | false    | false    | false   | false         | false     | json:"name,omitempty"    |          1 |         |
	| hit_die | int    | false  | false    | false    | false   | false         | false     | json:"hit_die,omitempty" |          1 |         |
	+---------+--------+--------+----------+----------+---------+---------------+-----------+--------------------------+------------+---------+
	+---------------------+-------------------+---------+---------+----------+--------+----------+---------+
	|        Edge         |       Type        | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+---------------------+-------------------+---------+---------+----------+--------+----------+---------+
	| proficiencies       | Proficiency       | false   |         | M2M      | false  | true     |         |
	| proficiency_options | ProficiencyChoice | false   |         | O2M      | false  | true     |         |
	| starting_equipment  | EquipmentEntry    | false   |         | M2M      | false  | true     |         |
	| saving_throws       | AbilityScore      | false   |         | M2M      | false  | true     |         |
	| characters          | Character         | true    | class   | O2M      | false  | true     |         |
	+---------------------+-------------------+---------+---------+----------+--------+----------+---------+
	
Coin:
	+----------------------+----------+--------+----------+----------+---------+---------------+-----------+---------------------------------------+------------+---------+
	|        Field         |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |               StructTag               | Validators | Comment |
	+----------------------+----------+--------+----------+----------+---------+---------------+-----------+---------------------------------------+------------+---------+
	| id                   | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"                   |          0 |         |
	| indx                 | string   | true   | false    | false    | false   | false         | false     | json:"index"                          |          1 |         |
	| name                 | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty"                 |          1 |         |
	| desc                 | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty"                 |          0 |         |
	| gold_conversion_rate | float64  | false  | false    | false    | false   | false         | false     | json:"gold_conversion_rate,omitempty" |          0 |         |
	+----------------------+----------+--------+----------+----------+---------+---------------+-----------+---------------------------------------+------------+---------+
	+-------+------+---------+---------+----------+--------+----------+---------+
	| Edge  | Type | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+-------+------+---------+---------+----------+--------+----------+---------+
	| costs | Cost | true    | coin    | O2M      | false  | true     |         |
	+-------+------+---------+---------+----------+--------+----------+---------+
	
Condition:
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string   | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| desc  | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	
Cost:
	+----------+------+--------+----------+----------+---------+---------------+-----------+---------------------------+------------+---------+
	|  Field   | Type | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |         StructTag         | Validators | Comment |
	+----------+------+--------+----------+----------+---------+---------------+-----------+---------------------------+------------+---------+
	| id       | int  | false  | false    | false    | false   | false         | false     | json:"id,omitempty"       |          0 |         |
	| quantity | int  | false  | false    | false    | true    | false         | false     | json:"quantity,omitempty" |          0 |         |
	+----------+------+--------+----------+----------+---------+---------------+-----------+---------------------------+------------+---------+
	+-----------+-----------+---------+---------+----------+--------+----------+---------+
	|   Edge    |   Type    | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+-----------+-----------+---------+---------+----------+--------+----------+---------+
	| coin      | Coin      | false   |         | M2O      | true   | false    |         |
	| equipment | Equipment | true    | cost    | O2O      | true   | true     |         |
	+-----------+-----------+---------+---------+----------+--------+----------+---------+
	
DamageType:
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string   | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| desc  | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	+---------+--------+---------+-------------+----------+--------+----------+---------+
	|  Edge   |  Type  | Inverse |   BackRef   | Relation | Unique | Optional | Comment |
	+---------+--------+---------+-------------+----------+--------+----------+---------+
	| weapons | Weapon | true    | damage_type | O2M      | false  | true     |         |
	+---------+--------+---------+-------------+----------+--------+----------+---------+
	
Equipment:
	+--------------------+-----------------------------+--------+----------+----------+---------+---------------+-----------+-------------------------------------+------------+---------+
	|       Field        |            Type             | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |              StructTag              | Validators | Comment |
	+--------------------+-----------------------------+--------+----------+----------+---------+---------------+-----------+-------------------------------------+------------+---------+
	| id                 | int                         | false  | false    | false    | false   | false         | false     | json:"id,omitempty"                 |          0 |         |
	| indx               | string                      | true   | false    | false    | false   | false         | false     | json:"indx,omitempty"               |          1 |         |
	| name               | string                      | false  | false    | false    | false   | false         | false     | json:"name,omitempty"               |          1 |         |
	| equipment_category | equipment.EquipmentCategory | false  | false    | false    | false   | false         | false     | json:"equipment_category,omitempty" |          0 |         |
	| weight             | float64                     | false  | true     | false    | false   | false         | false     | json:"weight,omitempty"             |          0 |         |
	+--------------------+-----------------------------+--------+----------+----------+---------+---------------+-----------+-------------------------------------+------------+---------+
	+-------------------+----------------+---------+-----------+----------+--------+----------+---------+
	|       Edge        |      Type      | Inverse |  BackRef  | Relation | Unique | Optional | Comment |
	+-------------------+----------------+---------+-----------+----------+--------+----------+---------+
	| cost              | Cost           | false   |           | O2O      | true   | true     |         |
	| gear              | Gear           | false   |           | O2O      | true   | true     |         |
	| tool              | Tool           | false   |           | O2O      | true   | true     |         |
	| weapon            | Weapon         | false   |           | O2O      | true   | true     |         |
	| vehicle           | Vehicle        | false   |           | O2O      | true   | true     |         |
	| armor             | Armor          | false   |           | O2O      | true   | true     |         |
	| equipment_entries | EquipmentEntry | true    | equipment | O2M      | false  | true     |         |
	+-------------------+----------------+---------+-----------+----------+--------+----------+---------+
	
EquipmentEntry:
	+----------+------+--------+----------+----------+---------+---------------+-----------+---------------------------+------------+---------+
	|  Field   | Type | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |         StructTag         | Validators | Comment |
	+----------+------+--------+----------+----------+---------+---------------+-----------+---------------------------+------------+---------+
	| id       | int  | false  | false    | false    | false   | false         | false     | json:"id,omitempty"       |          0 |         |
	| quantity | int  | false  | false    | false    | false   | false         | false     | json:"quantity,omitempty" |          1 |         |
	+----------+------+--------+----------+----------+---------+---------------+-----------+---------------------------+------------+---------+
	+-----------+-----------+---------+--------------------+----------+--------+----------+---------+
	|   Edge    |   Type    | Inverse |      BackRef       | Relation | Unique | Optional | Comment |
	+-----------+-----------+---------+--------------------+----------+--------+----------+---------+
	| class     | Class     | true    | starting_equipment | M2M      | false  | true     |         |
	| equipment | Equipment | false   |                    | M2O      | true   | false    |         |
	+-----------+-----------+---------+--------------------+----------+--------+----------+---------+
	
Feat:
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string   | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| desc  | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	
Feature:
	+-------+----------+--------+----------+----------+---------+---------------+-----------+------------------------+------------+---------+
	| Field |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag        | Validators | Comment |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+------------------------+------------+---------+
	| id    | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"    |          0 |         |
	| indx  | string   | true   | false    | false    | false   | false         | false     | json:"index"           |          1 |         |
	| name  | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty"  |          1 |         |
	| desc  | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty"  |          0 |         |
	| level | int      | false  | false    | false    | false   | false         | false     | json:"level,omitempty" |          1 |         |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+------------------------+------------+---------+
	+---------------+--------------+---------+---------+----------+--------+----------+---------+
	|     Edge      |     Type     | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+---------------+--------------+---------+---------+----------+--------+----------+---------+
	| prerequisites | Prerequisite | false   |         | O2M      | false  | true     |         |
	+---------------+--------------+---------+---------+----------+--------+----------+---------+
	
Gear:
	+---------------+----------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
	|     Field     |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |           StructTag            | Validators | Comment |
	+---------------+----------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
	| id            | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"            |          0 |         |
	| gear_category | string   | false  | false    | false    | false   | false         | false     | json:"gear_category,omitempty" |          0 |         |
	| desc          | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty"          |          0 |         |
	+---------------+----------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
	+-----------+-----------+---------+---------+----------+--------+----------+---------+
	|   Edge    |   Type    | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+-----------+-----------+---------+---------+----------+--------+----------+---------+
	| equipment | Equipment | true    | gear    | O2O      | true   | false    |         |
	+-----------+-----------+---------+---------+----------+--------+----------+---------+
	
Language:
	+---------------+-----------------------+--------+----------+----------+---------+---------------+-----------+-------------------------+------------+---------+
	|     Field     |         Type          | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |        StructTag        | Validators | Comment |
	+---------------+-----------------------+--------+----------+----------+---------+---------------+-----------+-------------------------+------------+---------+
	| id            | int                   | false  | false    | false    | false   | false         | false     | json:"id,omitempty"     |          0 |         |
	| indx          | string                | true   | false    | false    | false   | false         | false     | json:"index"            |          1 |         |
	| name          | string                | false  | false    | false    | false   | false         | false     | json:"name,omitempty"   |          1 |         |
	| desc          | []string              | false  | true     | false    | false   | false         | false     | json:"desc,omitempty"   |          0 |         |
	| language_type | language.LanguageType | false  | false    | false    | true    | false         | false     | json:"type"             |          0 |         |
	| script        | language.Script       | false  | false    | false    | true    | false         | false     | json:"script,omitempty" |          0 |         |
	+---------------+-----------------------+--------+----------+----------+---------+---------------+-----------+-------------------------+------------+---------+
	+---------+----------------+---------+-----------+----------+--------+----------+---------+
	|  Edge   |      Type      | Inverse |  BackRef  | Relation | Unique | Optional | Comment |
	+---------+----------------+---------+-----------+----------+--------+----------+---------+
	| race    | Race           | true    | languages | M2M      | false  | true     |         |
	| options | LanguageChoice | true    | languages | M2M      | false  | true     |         |
	+---------+----------------+---------+-----------+----------+--------+----------+---------+
	
LanguageChoice:
	+--------+------+--------+----------+----------+---------+---------------+-----------+-------------------------+------------+---------+
	| Field  | Type | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |        StructTag        | Validators | Comment |
	+--------+------+--------+----------+----------+---------+---------------+-----------+-------------------------+------------+---------+
	| id     | int  | false  | false    | false    | false   | false         | false     | json:"id,omitempty"     |          0 |         |
	| choose | int  | false  | false    | false    | false   | false         | false     | json:"choose,omitempty" |          1 |         |
	+--------+------+--------+----------+----------+---------+---------------+-----------+-------------------------+------------+---------+
	+-----------+----------+---------+------------------+----------+--------+----------+---------+
	|   Edge    |   Type   | Inverse |     BackRef      | Relation | Unique | Optional | Comment |
	+-----------+----------+---------+------------------+----------+--------+----------+---------+
	| languages | Language | false   |                  | M2M      | false  | true     |         |
	| race      | Race     | true    | language_options | O2O      | true   | true     |         |
	| subrace   | Subrace  | true    | language_options | M2O      | true   | true     |         |
	+-----------+----------+---------+------------------+----------+--------+----------+---------+
	
MagicSchool:
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string   | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| desc  | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	
Prerequisite:
	+-------------------+-------------------------------+--------+----------+----------+---------+---------------+-----------+--------------------------+------------+---------+
	|       Field       |             Type              | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |        StructTag         | Validators | Comment |
	+-------------------+-------------------------------+--------+----------+----------+---------+---------------+-----------+--------------------------+------------+---------+
	| id                | int                           | false  | false    | false    | false   | false         | false     | json:"id,omitempty"      |          0 |         |
	| prerequisite_type | prerequisite.PrerequisiteType | false  | false    | false    | false   | false         | false     | json:"type,omitempty"    |          0 |         |
	| level_value       | int                           | false  | true     | false    | false   | false         | false     | json:"level,omitempty"   |          0 |         |
	| feature_value     | string                        | false  | true     | false    | false   | false         | false     | json:"feature,omitempty" |          0 |         |
	| spell_value       | string                        | false  | true     | false    | false   | false         | false     | json:"spell,omitempty"   |          0 |         |
	+-------------------+-------------------------------+--------+----------+----------+---------+---------------+-----------+--------------------------+------------+---------+
	+---------+---------+---------+---------------+----------+--------+----------+---------+
	|  Edge   |  Type   | Inverse |    BackRef    | Relation | Unique | Optional | Comment |
	+---------+---------+---------+---------------+----------+--------+----------+---------+
	| feature | Feature | true    | prerequisites | M2O      | true   | true     |         |
	+---------+---------+---------+---------------+----------+--------+----------+---------+
	
Proficiency:
	+-----------+--------+--------+----------+----------+---------+---------------+-----------+----------------------------+------------+---------+
	|   Field   |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |         StructTag          | Validators | Comment |
	+-----------+--------+--------+----------+----------+---------+---------------+-----------+----------------------------+------------+---------+
	| id        | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"        |          0 |         |
	| indx      | string | true   | false    | false    | false   | false         | false     | json:"indx,omitempty"      |          1 |         |
	| name      | string | false  | false    | false    | false   | false         | false     | json:"name,omitempty"      |          1 |         |
	| reference | string | false  | false    | false    | false   | false         | false     | json:"reference,omitempty" |          1 |         |
	+-----------+--------+--------+----------+----------+---------+---------------+-----------+----------------------------+------------+---------+
	+---------+-------------------+---------+------------------------+----------+--------+----------+---------+
	|  Edge   |       Type        | Inverse |        BackRef         | Relation | Unique | Optional | Comment |
	+---------+-------------------+---------+------------------------+----------+--------+----------+---------+
	| race    | Race              | true    | starting_proficiencies | M2M      | false  | true     |         |
	| options | ProficiencyChoice | true    | proficiencies          | M2M      | false  | true     |         |
	| subrace | Subrace           | true    | proficiencies          | M2M      | false  | true     |         |
	| class   | Class             | true    | proficiencies          | M2M      | false  | true     |         |
	+---------+-------------------+---------+------------------------+----------+--------+----------+---------+
	
ProficiencyChoice:
	+--------+----------+--------+----------+----------+---------+---------------+-----------+-------------------------+------------+---------+
	| Field  |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |        StructTag        | Validators | Comment |
	+--------+----------+--------+----------+----------+---------+---------------+-----------+-------------------------+------------+---------+
	| id     | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"     |          0 |         |
	| choose | int      | false  | false    | false    | false   | false         | false     | json:"choose,omitempty" |          1 |         |
	| desc   | []string | false  | false    | false    | false   | false         | false     | json:"desc,omitempty"   |          0 |         |
	+--------+----------+--------+----------+----------+---------+---------------+-----------+-------------------------+------------+---------+
	+---------------+-------------+---------+------------------------------+----------+--------+----------+---------+
	|     Edge      |    Type     | Inverse |           BackRef            | Relation | Unique | Optional | Comment |
	+---------------+-------------+---------+------------------------------+----------+--------+----------+---------+
	| proficiencies | Proficiency | false   |                              | M2M      | false  | true     |         |
	| race          | Race        | true    | starting_proficiency_options | O2O      | true   | true     |         |
	| class         | Class       | true    | proficiency_options          | M2O      | true   | true     |         |
	+---------------+-------------+---------+------------------------------+----------+--------+----------+---------+
	
Property:
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string   | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| desc  | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	+---------+--------+---------+------------+----------+--------+----------+---------+
	|  Edge   |  Type  | Inverse |  BackRef   | Relation | Unique | Optional | Comment |
	+---------+--------+---------+------------+----------+--------+----------+---------+
	| weapons | Weapon | true    | properties | M2M      | false  | true     |         |
	+---------+--------+---------+------------+----------+--------+----------+---------+
	
Race:
	+----------------+-----------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
	|     Field      |   Type    | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |           StructTag            | Validators | Comment |
	+----------------+-----------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
	| id             | int       | false  | false    | false    | false   | false         | false     | json:"id,omitempty"            |          0 |         |
	| indx           | string    | true   | false    | false    | false   | false         | false     | json:"index"                   |          1 |         |
	| name           | string    | false  | false    | false    | false   | false         | false     | json:"name,omitempty"          |          1 |         |
	| speed          | int       | false  | false    | false    | false   | false         | false     | json:"speed,omitempty"         |          1 |         |
	| size           | race.Size | false  | false    | false    | true    | false         | false     | json:"size,omitempty"          |          0 |         |
	| size_desc      | string    | false  | false    | false    | false   | false         | false     | json:"size_description"        |          0 |         |
	| alignment_desc | string    | false  | false    | false    | false   | false         | false     | json:"alignment"               |          0 |         |
	| age_desc       | string    | false  | false    | false    | false   | false         | false     | json:"age"                     |          0 |         |
	| language_desc  | string    | false  | false    | false    | false   | false         | false     | json:"language_desc,omitempty" |          0 |         |
	+----------------+-----------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
	+------------------------------+--------------------+---------+---------+----------+--------+----------+---------+
	|             Edge             |        Type        | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+------------------------------+--------------------+---------+---------+----------+--------+----------+---------+
	| traits                       | Trait              | false   |         | M2M      | false  | true     |         |
	| starting_proficiencies       | Proficiency        | false   |         | M2M      | false  | true     |         |
	| starting_proficiency_options | ProficiencyChoice  | false   |         | O2O      | true   | true     |         |
	| ability_bonuses              | AbilityBonus       | false   |         | M2M      | false  | true     |         |
	| ability_bonus_options        | AbilityBonusChoice | false   |         | M2O      | true   | true     |         |
	| languages                    | Language           | false   |         | M2M      | false  | true     |         |
	| language_options             | LanguageChoice     | false   |         | O2O      | true   | true     |         |
	| subraces                     | Subrace            | false   |         | O2M      | false  | true     |         |
	| characters                   | Character          | true    | race    | O2M      | false  | true     |         |
	+------------------------------+--------------------+---------+---------+----------+--------+----------+---------+
	
Rule:
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string   | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| desc  | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	+----------+-------------+---------+---------+----------+--------+----------+---------+
	|   Edge   |    Type     | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+----------+-------------+---------+---------+----------+--------+----------+---------+
	| sections | RuleSection | false   |         | O2M      | false  | true     |         |
	+----------+-------------+---------+---------+----------+--------+----------+---------+
	
RuleSection:
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string   | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| desc  | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	+------+------+---------+----------+----------+--------+----------+---------+
	| Edge | Type | Inverse | BackRef  | Relation | Unique | Optional | Comment |
	+------+------+---------+----------+----------+--------+----------+---------+
	| rule | Rule | true    | sections | M2O      | true   | true     |         |
	+------+------+---------+----------+----------+--------+----------+---------+
	
Skill:
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string   | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| desc  | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	+---------------+--------------+---------+---------+----------+--------+----------+---------+
	|     Edge      |     Type     | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+---------------+--------------+---------+---------+----------+--------+----------+---------+
	| ability_score | AbilityScore | true    | skills  | M2O      | true   | true     |         |
	+---------------+--------------+---------+---------+----------+--------+----------+---------+
	
Subrace:
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string   | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| desc  | []string | false  | false    | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	+------------------+----------------+---------+----------+----------+--------+----------+---------+
	|       Edge       |      Type      | Inverse | BackRef  | Relation | Unique | Optional | Comment |
	+------------------+----------------+---------+----------+----------+--------+----------+---------+
	| race             | Race           | true    | subraces | M2O      | true   | true     |         |
	| ability_bonuses  | AbilityBonus   | false   |          | M2M      | false  | true     |         |
	| proficiencies    | Proficiency    | false   |          | M2M      | false  | true     |         |
	| traits           | Trait          | false   |          | M2M      | false  | true     |         |
	| language_options | LanguageChoice | false   |          | O2M      | false  | true     |         |
	+------------------+----------------+---------+----------+----------+--------+----------+---------+
	
Tool:
	+---------------+----------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
	|     Field     |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |           StructTag            | Validators | Comment |
	+---------------+----------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
	| id            | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"            |          0 |         |
	| tool_category | string   | false  | false    | false    | false   | false         | false     | json:"tool_category,omitempty" |          0 |         |
	| desc          | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty"          |          0 |         |
	+---------------+----------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
	+-----------+-----------+---------+---------+----------+--------+----------+---------+
	|   Edge    |   Type    | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+-----------+-----------+---------+---------+----------+--------+----------+---------+
	| equipment | Equipment | true    | tool    | O2O      | true   | false    |         |
	+-----------+-----------+---------+---------+----------+--------+----------+---------+
	
Trait:
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string   | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| desc  | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	+---------+---------+---------+---------+----------+--------+----------+---------+
	|  Edge   |  Type   | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+---------+---------+---------+---------+----------+--------+----------+---------+
	| race    | Race    | true    | traits  | M2M      | false  | true     |         |
	| subrace | Subrace | true    | traits  | M2M      | false  | true     |         |
	+---------+---------+---------+---------+----------+--------+----------+---------+
	
Vehicle:
	+------------------+-------------------------+--------+----------+----------+---------+---------------+-----------+-----------------------------------+------------+---------+
	|      Field       |          Type           | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |             StructTag             | Validators | Comment |
	+------------------+-------------------------+--------+----------+----------+---------+---------------+-----------+-----------------------------------+------------+---------+
	| id               | int                     | false  | false    | false    | false   | false         | false     | json:"id,omitempty"               |          0 |         |
	| vehicle_category | vehicle.VehicleCategory | false  | false    | false    | false   | false         | false     | json:"vehicle_category,omitempty" |          0 |         |
	| capacity         | string                  | false  | true     | false    | false   | false         | false     | json:"capacity,omitempty"         |          0 |         |
	| desc             | []string                | false  | true     | false    | false   | false         | false     | json:"desc,omitempty"             |          0 |         |
	| speed_quantity   | float64                 | false  | true     | false    | false   | false         | false     | json:"speed_quantity,omitempty"   |          0 |         |
	| speed_units      | vehicle.SpeedUnits      | false  | true     | false    | false   | false         | false     | json:"speed_units,omitempty"      |          0 |         |
	+------------------+-------------------------+--------+----------+----------+---------+---------------+-----------+-----------------------------------+------------+---------+
	+-----------+-----------+---------+---------+----------+--------+----------+---------+
	|   Edge    |   Type    | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+-----------+-----------+---------+---------+----------+--------+----------+---------+
	| equipment | Equipment | true    | vehicle | O2O      | true   | false    |         |
	+-----------+-----------+---------+---------+----------+--------+----------+---------+
	
Weapon:
	+--------------------+--------------------------+--------+----------+----------+---------+---------------+-----------+-------------------------------------+------------+---------+
	|       Field        |           Type           | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |              StructTag              | Validators | Comment |
	+--------------------+--------------------------+--------+----------+----------+---------+---------------+-----------+-------------------------------------+------------+---------+
	| id                 | int                      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"                 |          0 |         |
	| weapon_category    | weapon.WeaponCategory    | false  | false    | false    | false   | false         | false     | json:"weapon_category,omitempty"    |          0 |         |
	| weapon_subcategory | weapon.WeaponSubcategory | false  | false    | false    | false   | false         | false     | json:"weapon_subcategory,omitempty" |          0 |         |
	| range_normal       | int                      | false  | true     | false    | false   | false         | false     | json:"range_normal,omitempty"       |          0 |         |
	| range_long         | int                      | false  | true     | false    | false   | false         | false     | json:"range_long,omitempty"         |          0 |         |
	| throw_range_normal | int                      | false  | true     | false    | false   | false         | false     | json:"throw_range_normal,omitempty" |          0 |         |
	| throw_range_long   | int                      | false  | true     | false    | false   | false         | false     | json:"throw_range_long,omitempty"   |          0 |         |
	| damage_dice        | string                   | false  | true     | false    | false   | false         | false     | json:"damage_dice,omitempty"        |          0 |         |
	+--------------------+--------------------------+--------+----------+----------+---------+---------------+-----------+-------------------------------------+------------+---------+
	+-------------+------------+---------+---------+----------+--------+----------+---------+
	|    Edge     |    Type    | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+-------------+------------+---------+---------+----------+--------+----------+---------+
	| properties  | Property   | false   |         | M2M      | false  | true     |         |
	| damage_type | DamageType | false   |         | M2O      | true   | true     |         |
	| equipment   | Equipment  | true    | weapon  | O2O      | true   | false    |         |
	+-------------+------------+---------+---------+----------+--------+----------+---------+
	
