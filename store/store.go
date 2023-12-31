/*
Copyright © 2023 Alfred Dobradi <alfreddobradi@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package store

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"time"

	"github.com/alfreddobradi/pwdm/crypto"
	"github.com/dgraph-io/badger/v3"
)

const keySessionKey string = "session-key"

var store *Store

type Store struct {
	*badger.DB
}

func Init(path string) error {
	if store == nil {
		opts := badger.DefaultOptions(path).WithLoggingLevel(badger.WARNING)
		db, err := badger.Open(opts)
		if err != nil {
			return err
		}

		store = &Store{
			db,
		}
	}

	return nil
}

func getSessionKey() ([]byte, error) {
	if store == nil {
		return nil, fmt.Errorf("store not created")
	}

	var sessionKey []byte
	err := store.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(keySessionKey))
		if err != nil {
			return err
		}

		sessionKey, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}

		return err
	})

	return sessionKey, err
}

func SetSessionKey(key []byte, ttl time.Duration) error {
	if store == nil {
		return fmt.Errorf("store not created")
	}

	sum := md5.Sum(key) // AES encryption takes fixed-length secrets
	err := store.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(keySessionKey), []byte(fmt.Sprintf("%x", sum))).
			WithTTL(ttl)
		err := txn.SetEntry(e)
		return err
	})

	return err
}

func Set(key, val []byte) error {
	if bytes.Equal(key, []byte(keySessionKey)) {
		return fmt.Errorf("key %s is reserved for storing the current session key", keySessionKey)
	}

	sessionKey, err := getSessionKey()
	if err != nil {
		return fmt.Errorf("failed to retrieve session-key: %w", err)
	}

	encryptedValue, err := crypto.Encrypt(sessionKey, val)
	if err != nil {
		return fmt.Errorf("failed to encrypt value: %w", err)
	}

	err = store.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry(key, encryptedValue)
		err := txn.SetEntry(e)
		return err
	})

	if err != nil {
		return fmt.Errorf("failed to store encrypted value: %w", err)
	}

	return nil
}

func Get(key []byte) ([]byte, error) {
	if store == nil {
		return nil, fmt.Errorf("store not created")
	}

	var val []byte
	err := store.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		val, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}

		return err
	})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve value for key %s: %w", string(key), err)
	}

	sessionKey, err := getSessionKey()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve session-key: %w", err)
	}
	val, err = crypto.Decrypt(sessionKey, val)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt value: %w", err)
	}

	return val, nil
}
