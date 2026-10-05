# Setup
Create a new UTF-8 `snippetbox` database. `utf8mb4` is MySQL's UTF-8 charset. 
`COLLATE` defines how text is compared and sorted. The `_ci` means case-insensitive.  
```sql
CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

Switch to using the `snippetbox` database
```sql
USE snippetbox;
```

Create a table named `SNIPPETS`. Each row represents one snippet.  
```sql
CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);
```

Snippet ID, the auto-incrementing primary key.  
```sql
id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT
```

Creates an **index** on the `created` column (makes queries related to this column faster).  
```sql
CREATE INDEX idx_snippets_created ON snippets(created);
```
