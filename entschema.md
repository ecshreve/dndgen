AbilityBonus:
	+-------+------+--------+----------+----------+---------+---------------+-----------+------------------------+------------+---------+
	| Field | Type | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag        | Validators | Comment |
	+-------+------+--------+----------+----------+---------+---------------+-----------+------------------------+------------+---------+
	| id    | int  | false  | false    | false    | false   | false         | false     | json:"id,omitempty"    |          0 |         |
	| bonus | int  | false  | false    | false    | false   | false         | false     | json:"bonus,omitempty" |          1 |         |
	+-------+------+--------+----------+----------+---------+---------------+-----------+------------------------+------------+---------+
	+---------------+--------------+---------+-----------------+----------+--------+----------+---------+
	|     Edge      |     Type     | Inverse |     BackRef     | Relation | Unique | Optional | Comment |
	+---------------+--------------+---------+-----------------+----------+--------+----------+---------+
	| ability_score | AbilityScore | true    | ability_bonuses | M2O      | true   | true     |         |
	| race          | Race         | true    | ability_bonuses | M2O      | true   | true     |         |
	+---------------+--------------+---------+-----------------+----------+--------+----------+---------+
	
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
	+-----------------+--------------+---------+---------+----------+--------+----------+---------+
	|      Edge       |     Type     | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+-----------------+--------------+---------+---------+----------+--------+----------+---------+
	| skills          | Skill        | false   |         | O2M      | false  | true     |         |
	| ability_bonuses | AbilityBonus | false   |         | O2M      | false  | true     |         |
	+-----------------+--------------+---------+---------+----------+--------+----------+---------+
	
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
	
Condition:
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string   | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| desc  | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	
DamageType:
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string   | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| desc  | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	
Feat:
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string   | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| desc  | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	
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
	| ability_bonuses | AbilityBonus | false   |         | O2M      | false  | true     |         |
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
	
WeaponProperty:
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |   Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int      | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string   | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string   | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| desc  | []string | false  | true     | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	+-------+----------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	
