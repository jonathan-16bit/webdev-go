# Setup
Create a new UTF-8 `snippetbox` database. `utf8mb4` is MySQL's UTF-8 charset. 
"Collate" is for comparison and sorting test. The `_ci` means case-insensitive.  
```sql
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

Switch to using the `snippetbox` database
```sql
USE snippetbox;
```
