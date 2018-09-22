package treestore

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thetatoken/ukulele/common"
	"github.com/thetatoken/ukulele/store/database/backend"
)

func TestTreeStore(t *testing.T) {
	assert := assert.New(t)

	db, err := backend.NewMgoDatabase()
	assert.Nil(err)
	treestore := NewTreeStore(common.Hash{}, db, false)

	key1 := common.Bytes("test/111")
	value1 := common.Bytes("aaa")

	treestore.Set(key1, value1)
	assert.Equal(value1, treestore.Get(key1))

	key2 := common.Bytes("test/123")
	value2 := common.Bytes("bbb")
	treestore.Set(key2, value2)

	key3 := common.Bytes("test/222")
	value3 := common.Bytes("ccc")
	treestore.Set(key3, value3)

	key4 := common.Bytes("test/333")
	value4 := common.Bytes("ddd")
	treestore.Set(key4, value4)

	key5 := common.Bytes("test/3331")
	value5 := common.Bytes("eee")
	treestore.Set(key5, value5)

	key6 := common.Bytes("test/334")
	value6 := common.Bytes("fff")
	treestore.Set(key6, value6)

	root, _ := treestore.Commit(nil)
	treestore.GetDB().Commit(root, true)

	assert.True(db.Has(root[:]))

	assert.Equal(value2, treestore.Get(key2))
	assert.Equal(value3, treestore.Get(key3))
	assert.Equal(value4, treestore.Get(key4))
	assert.Equal(value5, treestore.Get(key5))
	assert.Equal(value6, treestore.Get(key6))

	var cnt int

	cb := func(prefix common.Bytes) func(k, v common.Bytes) bool {
		cnt = 0
		return func(k, v common.Bytes) bool {
			cnt++
			success := bytes.HasPrefix(k, prefix)
			success = success && (bytes.Compare(v, treestore.Get(k)) == 0)
			return success
		}
	}

	prefix1 := common.Bytes("test/1")
	treestore.Traverse(prefix1, cb(prefix1))
	assert.Equal(2, cnt)

	prefix2 := common.Bytes("test/2")
	treestore.Traverse(prefix2, cb(prefix2))
	assert.Equal(1, cnt)

	prefix3 := common.Bytes("test/333")
	treestore.Traverse(prefix3, cb(prefix3))
	assert.Equal(2, cnt)

	prefix4 := common.Bytes("test/33")
	treestore.Traverse(prefix4, cb(prefix4))
	assert.Equal(3, cnt)

	prefix5 := common.Bytes("test")
	treestore.Traverse(prefix5, cb(prefix5))
	assert.Equal(6, cnt)

	treestore.Set(key1, nil)
	assert.Nil(treestore.Get(key1))
	treestore.Traverse(prefix5, cb(prefix5))
	assert.Equal(5, cnt)

	root, _ = treestore.Commit(nil)
	treestore.Trie.GetDB().Commit(root, true)

	//////////////////////////////

	treestore1 := NewTreeStore(treestore.Hash(), db, false)
	assert.Nil(treestore1.Get(key1))
	assert.Equal(value2, treestore1.Get(key2))

	treestore1.Set(key2, value3)
	assert.Equal(value3, treestore1.Get(key2))
	assert.Equal(value2, treestore.Get(key2))

	//////////////////////////////

	nonpersistentstore := NewTreeStore(common.Hash{}, db, true)

	key7 := common.Bytes("test/000")
	value7 := common.Bytes("zzz")

	nonpersistentstore.Set(key7, value7)
	assert.Equal(value7, nonpersistentstore.Get(key7))

	root, _ = nonpersistentstore.Commit(nil)
	nonpersistentstore.GetDB().Commit(root, true)

	assert.False(db.Has(root[:]))
}
