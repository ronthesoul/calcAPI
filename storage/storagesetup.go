package storage

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

var APICSV = "API.csv"

func EnsureCSV() error {
	if _, err := os.Stat(APICSV); os.IsNotExist(err) {
		f, err := os.Create(APICSV)
		if err != nil {
			return err
		}
		defer f.Close()
		w := csv.NewWriter(f)
		defer w.Flush()
		return w.Write([]string{"api_key", "expires_at_unix"})
	}
	return nil
}

func AppendKey(apikey string, ttl time.Duration) error {
	if err := EnsureCSV(); err != nil {
		return err
	}
	f, err := os.OpenFile(APICSV, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	exp := time.Now().Add(ttl).Unix()
	w := csv.NewWriter(f)
	defer w.Flush()
	return w.Write([]string{apikey, strconv.FormatInt(exp, 10)})
}

func IsKeyValid(apikey string) (bool, error) {
	f, err := os.Open(APICSV)
	if err != nil {
		return false, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		return false, err
	}
	now := time.Now().Unix()

	for i := 1; i < len(rows); i++ {
		row := rows[i]
		if len(row) < 2 {
			continue
		}
		if row[0] != apikey {
			continue
		}
		exp, _ := strconv.ParseInt(row[1], 10, 64)
		return now <= exp, nil
	}
	return false, nil
}

func PurgeExpired() error {
	f, err := os.Open(APICSV)
	if err != nil {
		return err
	}
	defer f.Close()

	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		return err
	}

	now := time.Now().Unix()
	kept := [][]string{{"api_key", "expires_at_unix"}}
	for i := 1; i < len(rows); i++ {
		row := rows[i]
		if len(row) < 2 {
			continue
		}
		exp, _ := strconv.ParseInt(row[1], 10, 64)
		if now <= exp {
			kept = append(kept, row)
		}
	}

	tmp := APICSV + ".tmp"
	tf, err := os.Create(tmp)
	if err != nil {
		return err
	}
	defer tf.Close()

	w := csv.NewWriter(tf)
	if err := w.WriteAll(kept); err != nil {
		return err
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return err
	}

	return os.Rename(tmp, APICSV)
}

func SelfTest() error {
	// 1) Exists & stat
	fi, err := os.Stat(APICSV)
	if err != nil {
		return fmt.Errorf("stat: %w", err)
	}
	if fi.IsDir() {
		return errors.New("path is a directory, not a file")
	}

	// 2) Open RW (ensures permissions)
	f, err := os.OpenFile(APICSV, os.O_RDWR, 0)
	if err != nil {
		return fmt.Errorf("open RW: %w", err)
	}
	defer f.Close()

	// 3) Lightweight write-then-truncate (prove write works without polluting)
	off, _ := f.Seek(0, os.SEEK_END)
	testBytes := []byte(fmt.Sprintf("#selftest %d\n", time.Now().UnixNano()))
	if _, err := f.Write(testBytes); err != nil {
		return fmt.Errorf("write: %w", err)
	}
	// rollback
	if err := f.Truncate(off); err != nil {
		return fmt.Errorf("truncate: %w", err)
	}

	// 4) Optional: simple read (prove read works)
	buf := make([]byte, 1)
	if _, err := f.ReadAt(buf, 0); err != nil && !errors.Is(err, os.ErrClosed) {
		// ignore EOF on empty files
	}

	return nil
}
