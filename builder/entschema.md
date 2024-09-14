Character:
	+-----------+--------+--------+----------+----------+---------+---------------+-----------+----------------------------+------------+---------+
	|   Field   |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |         StructTag          | Validators | Comment |
	+-----------+--------+--------+----------+----------+---------+---------------+-----------+----------------------------+------------+---------+
	| id        | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"        |          0 |         |
	| name      | string | true   | false    | false    | false   | false         | false     | json:"name,omitempty"      |          1 |         |
	| level     | int    | false  | false    | false    | true    | false         | false     | json:"level,omitempty"     |          1 |         |
	| alignment | string | false  | true     | false    | false   | false         | false     | json:"alignment,omitempty" |          0 |         |
	+-----------+--------+--------+----------+----------+---------+---------------+-----------+----------------------------+------------+---------+
	+-------+-------+---------+------------+----------+--------+----------+---------+
	| Edge  | Type  | Inverse |  BackRef   | Relation | Unique | Optional | Comment |
	+-------+-------+---------+------------+----------+--------+----------+---------+
	| race  | Race  | true    | characters | M2O      | true   | false    |         |
	| class | Class | true    | characters | M2O      | true   | false    |         |
	+-------+-------+---------+------------+----------+--------+----------+---------+
	
Class:
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| name  | string | true   | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	+------------+-----------+---------+---------+----------+--------+----------+---------+
	|    Edge    |   Type    | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+------------+-----------+---------+---------+----------+--------+----------+---------+
	| characters | Character | false   |         | O2M      | false  | true     |         |
	+------------+-----------+---------+---------+----------+--------+----------+---------+
	
Race:
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| Field |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	| id    | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| name  | string | true   | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	+-------+--------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+---------+
	+------------+-----------+---------+---------+----------+--------+----------+---------+
	|    Edge    |   Type    | Inverse | BackRef | Relation | Unique | Optional | Comment |
	+------------+-----------+---------+---------+----------+--------+----------+---------+
	| characters | Character | false   |         | O2M      | false  | true     |         |
	+------------+-----------+---------+---------+----------+--------+----------+---------+
	
