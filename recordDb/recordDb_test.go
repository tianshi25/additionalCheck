package recordDb

import (
    "testing"
    "tianshi25.github.com/record"
)

func TestAppend(t *testing.T) {
    tests := []struct {
        db RecordDb
        r record.Record
        res int
    } {
        {
            RecordDb {
                [] record.Record {
                    record.NewRecord("path1", 1, 1, []string{}),
                    record.NewRecord("path2", 1, 1, []string{}),
                },
            },
            record.NewRecord("path3", 1, 1, []string{}),
            3,
        },
    }

    for _, test := range(tests) {
        test.db.Append(test.r)
        if (len(test.db.records) != test.res) {
            t.Errorf("test fail, input:%#v %#v expected:%#v output:%#v",
                test.db, test.r, test.res, test.db.Len())
        }
    }
}

func TestConcate(t *testing.T) {
    tests := []struct {
        db1 RecordDb
        db2 RecordDb
        res int
    } {
        {
            RecordDb {
                [] record.Record {
                    record.NewRecord("path1", 1, 1, []string{}),
                    record.NewRecord("path2", 1, 1, []string{}),
                },
            },
            RecordDb {
                [] record.Record {
                    record.NewRecord("path3", 1, 1, []string{}),
                    record.NewRecord("path4", 1, 1, []string{}),
                },
            },
            4,
        },
    }

    for _, test := range(tests) {
        test.db1.Concate(test.db2)
        if (len(test.db1.records) != test.res) {
            t.Errorf("test fail, input:%#v %#v expected:%#v output:%#v",
                test.db1, test.db2, test.res, test.db1.Len())
        }
    }
}