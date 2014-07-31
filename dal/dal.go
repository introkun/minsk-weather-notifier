package dal

import (
    "github.com/cznic/ql"
    "time"
)

type Database ql.DB
type DatabaseContext ql.TCtx

type Forecast struct {
    ID              int64 `ql:"index xID"`
    WeatherProvider string
    Timestamp       time.Time `ql:"index xTimestamp"`
    MinTemp         int32
    MaxTemp         int32
}

func openDb(name string) *ql.DB {
    qlOpt := ql.Options{CanCreate: true}
    db, err := ql.OpenFile(name, &qlOpt)

    if err != nil {
        panic(err)
    }

    return db
}

func InitDb(name string) (*Database, *DatabaseContext) {
    schema := ql.MustSchema((*Forecast)(nil), "", nil)

    db := openDb(name)

    ctx := ql.NewRWCtx()
    if _, _, err := db.Execute(ctx, schema); err != nil {
        panic(err)
    }

    return (*Database)(db), (*DatabaseContext)(ctx)
}

func InsertRecord(db *Database, ctx *DatabaseContext, dbItem *Forecast) {
    ins := ql.MustCompile(`
        BEGIN TRANSACTION;
            INSERT INTO Forecast VALUES($1, $2, $3, $4);
        COMMIT;`,
    )

    if _, _, err := (*ql.DB)(db).Execute((*ql.TCtx)(ctx), ins, ql.MustMarshal(dbItem)...); err != nil {
        panic(err)
    }
}

func Flush(db *Database) {
    (*ql.DB)(db).Close()
}
