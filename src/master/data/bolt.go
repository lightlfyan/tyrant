package data

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"

	"sync"

	"master/data/user"

	"github.com/boltdb/bolt"

	"master/util"
)

var Bucket []byte

func (t TaskReport) ToB() []byte {
	bin, _ := json.Marshal(t)
	return bin
}

func (t *TaskReport) Load(bins []byte) {
	err := json.Unmarshal(bins, t)
	if err != nil {
		log.Println(err)
	}
}

func LoadReport(bins []byte) *TaskReport {
	r := &TaskReport{}
	r.Load(bins)
	return r
}

var TaskReportCache *util.TCache

var db *bolt.DB

func boltInit() {
	Bucket = []byte("task")
	TaskReportCache = &util.TCache{
		Dict: make(map[int64]interface{}, 0),
		Lock: &sync.RWMutex{},
	}

	var err error
	db, err = bolt.Open("tyrant.db", 0600, nil)
	if err != nil {
		panic(err)
	}
	createBucket()
	loadAllReport()

	if k := TaskReportCache.MaxKey(); k > taskId {
		taskId = k
	}
	CurrTaskId = taskId + 1
	user.DB = db

	user.LoadUsers()
}

func BoltClose() {
	if db != nil {
		db.Close()
	}
}

func AddReport(r *TaskReport) {
	TaskReportCache.Set(r.TaskId, r)
	bolt_put(r)
}

func DelReport(id int64) {
	TaskReportCache.Del(id)
	bolt_del(id)
}

func CleanReportCacheByUid(uid string) {

	TaskReportCache.Lock.Lock()
	defer TaskReportCache.Lock.Unlock()

	ks := make([]int64, 0)

	for k, v := range TaskReportCache.Dict {
		v1 := v.(*TaskReport)
		if v1.Uid == uid {
			TaskReportCache.UnsafeDel(k)
			ks = append(ks, k)
		}
	}

	bolt_dels(ks)
}

// private======================================================================

func createBucket() {
	db.Update(func(tx *bolt.Tx) error {
		_, err1 := tx.CreateBucketIfNotExists(Bucket)
		_, err2 := tx.CreateBucketIfNotExists(user.Bucket)

		if err1 != nil {
			return fmt.Errorf("create bucket: %s", err1)
		}
		if err2 != nil {
			return fmt.Errorf("create bucket: %s", err2)
		}
		return nil
	})
}

func loadAllReport() {
	// init function, does not need lock
	f := func(k, v []byte) error {
		TaskReportCache.UnsafeSet(btoi(k), LoadReport(v))
		return nil
	}
	Foreach(f)
}

// bolt db======================================================================

func bolt_put(r *TaskReport) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(Bucket)
		err := b.Put(itob(r.TaskId), r.ToB())
		return err
	})
}

func Get(id int64) (*TaskReport, error) {
	key := itob(id)

	var v []byte

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(Bucket)
		v = b.Get(key)
		return nil
	})

	t := &TaskReport{}

	err := json.Unmarshal(v, t)

	return t, err
}

func bolt_del(id int64) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(Bucket)
		key := itob(id)
		b.Delete(key)
		return nil
	})
}

func bolt_dels(ids []int64) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(Bucket)
		for _, id := range ids {
			key := itob(id)
			b.Delete(key)
		}
		return nil
	})
}

func Foreach(F func(k, v []byte) error) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(Bucket)
		b.ForEach(F)
		return nil
	})
}

func itob(v int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int64 {
	i := binary.BigEndian.Uint64(b)
	return int64(i)
}

// bolt db======================================================================
