### Окружение.
```env
DB_PASSWORD
```

### Конфиг.
*обязателен!!!*
- Путь до файла с конфигом вынесен в константу в `cmd/main.go` название переменной - `configPath`

``` json
{
  "server": {
    "port": ""
  },
  "db": {
    "host": "",
    "port": "",
    "username": "",
    "db_name": "",
    "ssl_mode": ""
  }
}
```