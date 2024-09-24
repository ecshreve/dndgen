AbilityBonus:
	+-------+------+--------+----------+----------+---------+---------------+-----------+------------------------+------------+---------+
	| Field | Type | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag        | Validators | Comment |
	+-------+------+--------+----------+----------+---------+---------------+-----------+------------------------+------------+---------+
	| id    | int  | false  | false    | false    | false   | false         | false     | json:"id,omitempty"    |          0 |         |
	| bonus | int  | false  | false    | false    | false   | false         | false     | json:"bonus,omitempty" |          1 |         |
	+-------+------+--------+----------+----------+---------+---------------+-----------+------------------------+------------+---------+
	+---------------+--------------+---------+---------+----------+--------+----------+---------+
	|     Edge      |     Type     | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+---------------+--------------+---------+---------+----------+--------+----------+---------+
	| ability_score | AbilityScore | false   |         | M2O      | true   | false    |         |
	| race          | Race         | false   |         | M2O      | true   | false    |         |
	+---------------+--------------+---------+---------+----------+--------+----------+---------+
	
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
	+-----------------+--------------+---------+---------------+----------+--------+----------+---------+
	
Alignment:
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string   | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| desc  | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	| abbr  | string   | false  | false    | false    | false   | false         | false     | json:"abbreviation"   |          0 |         |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	
Armor:
	+----------------------+---------------------+--------+----------+----------+---------+---------------+-----------+---------------------------------------+------------+---------+
	|        Field         |        Type         | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |               StructTag               | Validators | Comment |
	+----------------------+---------------------+--------+----------+----------+---------+---------------+-----------+---------------------------------------+------------+---------+
	| id                   | int                 | false  | false    | false    | false   | false         | false     | json:"id,omitempty"                   |          0 |         |
	| armor_category       | armor.ArmorCategory | false  | false    | false    | false   | false         | false     | json:"armor_category,omitempty"       |          0 |         |
	| str_minimum          | int                 | false  | false    | false    | false   | false         | false     | json:"str_minimum,omitempty"          |          0 |         |
	| stealth_disadvantage | bool                | false  | false    | false    | false   | false         | false     | json:"stealth_disadvantage,omitempty" |          0 |         |
	+----------------------+---------------------+--------+----------+----------+---------+---------------+-----------+---------------------------------------+------------+---------+
	+-------------+------------+---------+---------+----------+--------+----------+---------+
	|    Edge     |    Type    | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+-------------+------------+---------+---------+----------+--------+----------+---------+
	| equipment   | Equipment  | true    | armor   | O2O      | true   | false    |         |
	| armor_class | ArmorClass | false   |         | O2O      | true   | true     |         |
	+-------------+------------+---------+---------+----------+--------+----------+---------+
	
ArmorClass:
	+-----------+------+--------+----------+----------+---------+---------------+-----------+----------------------------+------------+---------+
	|   Field   | Type | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |         StructTag          | Validators | Comment |
	+-----------+------+--------+----------+----------+---------+---------------+-----------+----------------------------+------------+---------+
	| id        | int  | false  | false    | false    | false   | false         | false     | json:"id,omitempty"        |          0 |         |
	| base      | int  | false  | false    | false    | false   | false         | false     | json:"base,omitempty"      |          1 |         |
	| dex_bonus | bool | false  | false    | false    | true    | false         | false     | json:"dex_bonus,omitempty" |          0 |         |
	+-----------+------+--------+----------+----------+---------+---------------+-----------+----------------------------+------------+---------+
	+-------+-------+---------+-------------+----------+--------+----------+---------+
	| Edge  | Type  | Inverse |   BackRef   | Relation | Unique | Optional | Comment |
	+-------+-------+---------+-------------+----------+--------+----------+---------+
	| armor | Armor | true    | armor_class | O2O      | true   | true     |         |
	+-------+-------+---------+-------------+----------+--------+----------+---------+
	
Class:
	+---------+--------+--------+----------+----------+---------+---------------+-----------+--------------------------+------------+---------+
	|  Field  |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |        StructTag         | Validators | Comment |
	+---------+--------+--------+----------+----------+---------+---------------+-----------+--------------------------+------------+---------+
	| id      | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"      |          0 |         |
	| indx    | string | true   | false    | false    | false   | false         | false     | json:"index"             |          1 |         |
	| name    | string | false  | false    | false    | false   | false         | false     | json:"name,omitempty"    |          1 |         |
	| hit_die | int    | false  | false    | false    | false   | false         | false     | json:"hit_die,omitempty" |          1 |         |
	+---------+--------+--------+----------+----------+---------+---------------+-----------+--------------------------+------------+---------+
	
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
	
Damage:
	+-------------+--------+--------+----------+----------+---------+---------------+-----------+------------------------------+------------+---------+
	|    Field    |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |          StructTag           | Validators | Comment |
	+-------------+--------+--------+----------+----------+---------+---------------+-----------+------------------------------+------------+---------+
	| id          | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"          |          0 |         |
	| damage_dice | string | false  | false    | false    | false   | false         | false     | json:"damage_dice,omitempty" |          0 |         |
	+-------------+--------+--------+----------+----------+---------+---------------+-----------+------------------------------+------------+---------+
	+-------------+------------+---------+---------+----------+--------+----------+---------+
	|    Edge     |    Type    | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+-------------+------------+---------+---------+----------+--------+----------+---------+
	| damage_type | DamageType | false   |         | M2O      | true   | true     |         |
	+-------------+------------+---------+---------+----------+--------+----------+---------+
	
DamageType:
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string   | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| desc  | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	
Equipment:
	+--------------------+-----------------------------+--------+----------+----------+---------+---------------+-----------+-------------------------------------+------------+---------+
	|       Field        |            Type             | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |              StructTag              | Validators | Comment |
	+--------------------+-----------------------------+--------+----------+----------+---------+---------------+-----------+-------------------------------------+------------+---------+
	| id                 | int                         | false  | false    | false    | false   | false         | false     | json:"id,omitempty"                 |          0 |         |
	| indx               | string                      | true   | false    | false    | false   | false         | false     | json:"index"                        |          1 |         |
	| name               | string                      | false  | false    | false    | false   | false         | false     | json:"name,omitempty"               |          1 |         |
	| desc               | []string                    | false  | true     | false    | false   | false         | false     | json:"desc,omitempty"               |          0 |         |
	| equipment_category | equipment.EquipmentCategory | false  | false    | false    | false   | false         | false     | json:"equipment_category,omitempty" |          0 |         |
	| weight             | float64                     | false  | false    | false    | false   | false         | false     | json:"weight,omitempty"             |          0 |         |
	+--------------------+-----------------------------+--------+----------+----------+---------+---------------+-----------+-------------------------------------+------------+---------+
	+---------+---------+---------+---------+----------+--------+----------+---------+
	|  Edge   |  Type   | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+---------+---------+---------+---------+----------+--------+----------+---------+
	| cost    | Cost    | false   |         | O2O      | true   | true     |         |
	| tool    | Tool    | false   |         | O2O      | true   | true     |         |
	| gear    | Gear    | false   |         | O2O      | true   | true     |         |
	| armor   | Armor   | false   |         | O2O      | true   | true     |         |
	| weapon  | Weapon  | false   |         | O2O      | true   | true     |         |
	| vehicle | Vehicle | false   |         | O2O      | true   | true     |         |
	+---------+---------+---------+---------+----------+--------+----------+---------+
	
Feat:
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string   | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| desc  | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	
Gear:
	+---------------+--------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
	|     Field     |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |           StructTag            | Validators | Comment |
	+---------------+--------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
	| id            | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"            |          0 |         |
	| gear_category | string | false  | false    | false    | false   | false         | false     | json:"gear_category,omitempty" |          0 |         |
	+---------------+--------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
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
	+-------+------+---------+-----------+----------+--------+----------+---------+
	| Edge  | Type | Inverse |  BackRef  | Relation | Unique | Optional | Comment |
	+-------+------+---------+-----------+----------+--------+----------+---------+
	| races | Race | true    | languages | M2M      | false  | true     |         |
	+-------+------+---------+-----------+----------+--------+----------+---------+
	
MagicSchool:
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string   | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| desc  | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	
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
	+-----------------+--------------+---------+---------+----------+--------+----------+---------+
	|      Edge       |     Type     | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+-----------------+--------------+---------+---------+----------+--------+----------+---------+
	| ability_bonuses | AbilityBonus | true    | race    | O2M      | false  | true     |         |
	| languages       | Language     | false   |         | M2M      | false  | true     |         |
	+-----------------+--------------+---------+---------+----------+--------+----------+---------+
	
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
	
Tool:
	+---------------+--------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
	|     Field     |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |           StructTag            | Validators | Comment |
	+---------------+--------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
	| id            | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"            |          0 |         |
	| tool_category | string | false  | false    | false    | false   | false         | false     | json:"tool_category,omitempty" |          0 |         |
	+---------------+--------+--------+----------+----------+---------+---------------+-----------+--------------------------------+------------+---------+
	+-----------+-----------+---------+---------+----------+--------+----------+---------+
	|   Edge    |   Type    | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+-----------+-----------+---------+---------+----------+--------+----------+---------+
	| equipment | Equipment | true    | tool    | O2O      | true   | false    |         |
	+-----------+-----------+---------+---------+----------+--------+----------+---------+
	
Vehicle:
	+------------------+-------------------------+--------+----------+----------+---------+---------------+-----------+-----------------------------------+------------+---------+
	|      Field       |          Type           | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |             StructTag             | Validators | Comment |
	+------------------+-------------------------+--------+----------+----------+---------+---------------+-----------+-----------------------------------+------------+---------+
	| id               | int                     | false  | false    | false    | false   | false         | false     | json:"id,omitempty"               |          0 |         |
	| vehicle_category | vehicle.VehicleCategory | false  | false    | false    | false   | false         | false     | json:"vehicle_category,omitempty" |          0 |         |
	| capacity         | string                  | false  | true     | false    | false   | false         | false     | json:"capacity,omitempty"         |          0 |         |
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
	+--------------------+--------------------------+--------+----------+----------+---------+---------------+-----------+-------------------------------------+------------+---------+
	+--------------+-------------+---------+---------+----------+--------+----------+---------+
	|     Edge     |    Type     | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+--------------+-------------+---------+---------+----------+--------+----------+---------+
	| damage       | Damage      | false   |         | M2O      | true   | true     |         |
	| properties   | Property    | false   |         | M2M      | false  | true     |         |
	| equipment    | Equipment   | true    | weapon  | O2O      | true   | false    |         |
	| weapon_range | WeaponRange | false   |         | M2O      | true   | true     |         |
	+--------------+-------------+---------+---------+----------+--------+----------+---------+
	
WeaponRange:
	+--------------------+------+--------+----------+----------+---------+---------------+-----------+-------------------------------------+------------+---------+
	|       Field        | Type | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |              StructTag              | Validators | Comment |
	+--------------------+------+--------+----------+----------+---------+---------------+-----------+-------------------------------------+------------+---------+
	| id                 | int  | false  | false    | false    | false   | false         | false     | json:"id,omitempty"                 |          0 |         |
	| range_normal       | int  | false  | true     | false    | false   | false         | false     | json:"range_normal,omitempty"       |          0 |         |
	| range_long         | int  | false  | true     | false    | false   | false         | false     | json:"range_long,omitempty"         |          0 |         |
	| throw_range_normal | int  | false  | true     | false    | false   | false         | false     | json:"throw_range_normal,omitempty" |          0 |         |
	| throw_range_long   | int  | false  | true     | false    | false   | false         | false     | json:"throw_range_long,omitempty"   |          0 |         |
	+--------------------+------+--------+----------+----------+---------+---------------+-----------+-------------------------------------+------------+---------+
	
