## 資料庫包＋Logrus
***
### MongoDB

```go
// ReplicaSet
host := "192.168.10.79:27017,192.168.10.80:27017,192.168.10.81:27017"
db, err := mongoz.New(host).
    SetReplicaSet("sh-rs-3").
    SetPool(1, 5, 10).
    // SetPoolMonitor().
    Connect()

// Direct
db, _ := mongoz.New(host).
    SetDirect(true).
    SetPool(1, 5, 10).
    // SetPoolMonitor().
    Connect()
```

### PostgreSQL

```go
pgdb, err := postgrez.New("192.168.10.101", "6432", "postgres", "1qaz2wsx", "gpsa_tx").
    // SetTimeZone("PRC").
    // SetLogger(logrusz.New().SetLevel("debug").Writer()).
    Connect(postgrez.Pool(1, 10, 10))
```

### MySQL

```go

mydb, err := mysqlz.New("127.0.0.1", "3306", "root", "iLove5566", "line").
    // SetAppendParameter(mysqlz.NewParamsmeter()).
    // SetCharset("utf8").
    // SetLoc("UTC").
    // SetLogger(logrusz.New().SetLevel("debug").Writer()).
    Connect(mysqlz.Pool(1, 2, 180))

```

### Logrus 

```go
l := logrusz.New().
    // SetLevel("debug").
    // SetPath("./logs").
    // SetPrefix("gf-").
    Writer()

// example
l.Println()
l.Info()
l.Warn()
l.Panic()
...
...
```

