AbilityScore:
	+-------+-------------------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |       Type        | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+-------------------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int               | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string            | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string            | true   | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| abbr  | abilityscore.Abbr | false  | false    | false    | false   | false         | false     | json:"abbr,omitempty" |          0 |         |
	| desc  | []string          | false  | false    | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	+-------+-------------------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	
Alignment:
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string | true   | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| desc  | string | false  | false    | false    | false   | false         | false     | json:"desc,omitempty" |          0 |         |
	| abbr  | string | false  | false    | false    | false   | false         | false     | json:"abbreviation"   |          0 |         |
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	
Character:
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| name  | string | true   | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	| age   | int    | false  | false    | false    | false   | false         | false     | json:"age,omitempty"  |          1 |         |
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	
Class:
	+---------+--------+--------+----------+----------+---------+---------------+-----------+--------------------------+------------+---------+
	|  Field  |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |        StructTag         | Validators | Comment |
	+---------+--------+--------+----------+----------+---------+---------------+-----------+--------------------------+------------+---------+
	| id      | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"      |          0 |         |
	| indx    | string | true   | false    | false    | false   | false         | false     | json:"index"             |          1 |         |
	| name    | string | true   | false    | false    | false   | false         | false     | json:"name,omitempty"    |          1 |         |
	| hit_die | int    | false  | false    | false    | false   | false         | false     | json:"hit_die,omitempty" |          1 |         |
	+---------+--------+--------+----------+----------+---------+---------------+-----------+--------------------------+------------+---------+
	
Race:
	+------------------+-----------+--------+----------+----------+---------+---------------+-----------+-----------------------------------+------------+---------+
	|      Field       |   Type    | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |             StructTag             | Validators | Comment |
	+------------------+-----------+--------+----------+----------+---------+---------------+-----------+-----------------------------------+------------+---------+
	| id               | int       | false  | false    | false    | false   | false         | false     | json:"id,omitempty"               |          0 |         |
	| indx             | string    | true   | false    | false    | false   | false         | false     | json:"index"                      |          1 |         |
	| name             | string    | true   | false    | false    | false   | false         | false     | json:"name,omitempty"             |          1 |         |
	| speed            | int       | false  | false    | false    | false   | false         | false     | json:"speed,omitempty"            |          1 |         |
	| size             | race.Size | false  | false    | false    | false   | false         | false     | json:"size,omitempty"             |          0 |         |
	| size_description | string    | false  | false    | false    | false   | false         | false     | json:"size_description,omitempty" |          0 |         |
	| age              | string    | false  | false    | false    | false   | false         | false     | json:"age,omitempty"              |          0 |         |
	+------------------+-----------+--------+----------+----------+---------+---------------+-----------+-----------------------------------+------------+---------+
	
Skill:
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| indx  | string | true   | false    | false    | false   | false         | false     | json:"index"          |          1 |         |
	| name  | string | true   | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	
