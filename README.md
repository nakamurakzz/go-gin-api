# go-gin-api
go-ginで作るWebAPIサーバー

## Webフレームワーク
- go-gin

## ORM
- gorm

# build
```bash
go build -o bin/go-api-server
```

# run
```bash
./bin/go-api-server
```

# install
```bash
go install github.com/k0kubun/sqldef/cmd/mysqldef@latest
```

# run
```bash
make up
```

# migrate
```bash
make migrate
```

# create key
```bash
cd auth/secret
openssl genrsa 4096 > secre.pem
openssl rsa -pubout < secre.pem > public.pem
```