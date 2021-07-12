# Env migrate

Use `github.com/rubenv/sql-migrate` and env, auto run migrate.

## 启用 sql-migrate

```go

func main(){
  env_migrate.Auto(db) // db is database/sql *sql.DB
}
```

## 启动服务/迁移数据库

migrate 会读取 `migrations/*.sql` 进行迁移

启动并执行所有迁移

```bash
up_migrate=all go run main.go
```

启动并执行回滚，1 次, 回滚会自动退出

```bash
down_migrate=1 go run main.go
```

跳过 2 次

```bash
skip_migrate=2 go run main.go
```

## 使用其他 migrations 路径

```bash
up_migrate=all dir_migrate=sql/migrations go run main.go
```
