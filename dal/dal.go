package dal

import (
    "github.com/cznic/ql"
    "time"
)

type Forecast struct {
    ID              int64 `ql:"index xID"`
    WeatherProvider string
    Timestamp       time.Time `ql:"index xTimestamp"`
    MinTemp         int32
    MaxTemp         int32
}

func InitDb(name string, ctx *ql.TCtx) *ql.DB {
    schema := ql.MustSchema((*Forecast)(nil), "", nil)

    qlOpt := ql.Options{CanCreate: true}
    db, err := ql.OpenFile(name, &qlOpt)

    if err != nil {
        panic(err)
    }

    if _, _, err := db.Execute(ctx, schema); err != nil {
        panic(err)
    }

    return db
}

func InsertRecord(db *ql.DB, ctx *ql.TCtx, dbItem *Forecast) {
    ins := ql.MustCompile(`
        BEGIN TRANSACTION;
            INSERT INTO forecast VALUES($1, $2, $3, $4);
        COMMIT;`,
    )

    if _, _, err := db.Execute(ctx, ins, ql.MustMarshal(dbItem)...); err != nil {
        panic(err)
    }
}
