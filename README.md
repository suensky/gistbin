# gistbin


## MySQL
### Start MySQL
To start mysql now and restart at login:
```SQL
brew services start mysql
```

Or, if you don't want/need a background service you can just run:
```SQL
/opt/homebrew/opt/mysql/bin/mysqld_safe --datadir\=/opt/homebrew/var/mysql
```


### Connect to MySQL
```SQL
mysql -u root

-- OR

mysql -D snippetbox -u web -p
-- PASSWORD: pass
```

### SQL

