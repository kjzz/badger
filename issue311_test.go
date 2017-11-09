package badger

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/dgraph-io/badger/options"
)

func TestPersistentCache_DirectBadger(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	config := DefaultOptions
	config.TableLoadingMode = options.MemoryMap
	config.ValueLogFileSize = 16 << 20
	config.LevelOneSize = 8 << 20
	config.MaxTableSize = 2 << 20
	config.Dir = dir
	config.ValueDir = dir
	config.SyncWrites = false

	db, err := Open(config)
	if err != nil {
		t.Fatalf("cannot open db at location %s: %v", dir, err)
	}

	err = db.View(func(txn *Txn) error { return nil })

	if err != nil {
		t.Fatal(err)
	}

	if err = db.Close(); err != nil {
		t.Fatal(err)
	}
}