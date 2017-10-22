# Error
```
sql: Scan error on column index 0: unsupported Scan, storing driver.Value type []uint8 into type *main.User
```

# resolve dependency
```bash
dep ensure
```
# migrate sqlite3
```bash
sqlite3 db.sqlite < migrate_sqlite3.sql
```

# migrate mysql
```bash
mysql -u root < migrate_mysql.sql
```