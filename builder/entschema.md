AbilityScore:

|-------|-------------------|--------|----------|----------|---------|---------------|-----------|------------------------|------------|---------|
| Field |       Type        | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag        | Validators | Comment |
|-------|-------------------|--------|----------|----------|---------|---------------|-----------|------------------------|------------|---------|
| id    | int               | false  | false    | false    | false   | false         | false     | json:"id,omitempty"    |          0 |         |
| name  | string            | false  | false    | false    | false   | false         | false     | json:"name,omitempty"  |          1 |         |
| abbr  | abilityscore.Abbr | false  | false    | false    | false   | false         | true      | json:"abbr,omitempty"  |          0 |         |
| score | int               | false  | false    | false    | false   | false         | false     | json:"score,omitempty" |          1 |         |
|-------|-------------------|--------|----------|----------|---------|---------------|-----------|------------------------|------------|---------|
	
Character:

| Field |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
|-------|--------|--------|----------|----------|---------|---------------|-----------|-----------------------|------------|---------|
| id    | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
| name  | string | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |

---

|      Edge      |     Type     | Inverse | BackRef | Relation | Unique | Optional | Comment |
|----------------|--------------|---------|---------|----------|--------|----------|---------|
| race           | Race         | false   |         | O2M      | false  | true     |         |
| class          | Class        | false   |         | O2M      | false  | true     |         |
| ability_scores | AbilityScore | false   |         | O2M      | false  | true     |         |
|----------------|--------------|---------|---------|----------|--------|----------|---------|
	
Class:
	|-------|--------|--------|----------|----------|---------|---------------|-----------|-----------------------|------------|---------|
	| Field |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	|-------|--------|--------|----------|----------|---------|---------------|-----------|-----------------------|------------|---------|
	| id    | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| name  | string | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	|-------|--------|--------|----------|----------|---------|---------------|-----------|-----------------------|------------|---------|
	
Race:
	|-------|--------|--------|----------|----------|---------|---------------|-----------|-----------------------|------------|---------|
	| Field |  Type  | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators | Comment |
	|-------|--------|--------|----------|----------|---------|---------------|-----------|-----------------------|------------|---------|
	| id    | int    | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |         |
	| name  | string | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          1 |         |
	|-------|--------|--------|----------|----------|---------|---------------|-----------|-----------------------|------------|---------|
	
