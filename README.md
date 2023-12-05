# gistbin


## MySQL
### Installation
```bash
brew install mysql
```

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

### Browse
#### Curl
```
curl -i "http://localhost:4000/gist/view?id=1”
```

## Security

### TLS
Generate a TLS certificate
```bash
go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```
