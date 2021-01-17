package recordDb

import (
    "sort"
    "strings"
    "tianshi25.github.com/record"
)

type RecordDb struct {
    records []record.Record
}

func (db *RecordDb) Append (r record.Record) {
    db.records = append(db.records, r)
}

func (db *RecordDb) Concate (db2 RecordDb) {
    db.records = append(db.records, db2.records...)
}

func (db RecordDb) GetStr() string {
    sArray := make([]string, len(db.records))
    for i:=0; i<len(db.records); i++ {
        sArray[i] = db.records[i].GetStr() + "\n"
    }
    sort.Strings(sArray)
    return strings.Join(sArray, "")
}

func (db RecordDb) Len() int {
    return len(db.records)
}
