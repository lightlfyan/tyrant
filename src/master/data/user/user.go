package user

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"master/util"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"strconv"

	"sync/atomic"

	"github.com/boltdb/bolt"
)

type User struct {
	Uid      string
	Name     string
	Password string
}

var Bucket []byte

var UserId int64

var configByUid map[string]*User
var configByName map[string]*User

var lock *sync.RWMutex

var DB *bolt.DB

var UserCache *util.TCache

var flag int32 = 0

func init() {

	if flag > 0 {
		log.Println("init")
		return
	}

	atomic.AddInt32(&flag, 1)

	Bucket = []byte("user")

	lock = &sync.RWMutex{}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	users, err := ioutil.ReadFile(dir + "/user.txt")
	if err != nil {
		panic(err)
	}

	configByUid = make(map[string]*User, 0)
	configByName = make(map[string]*User, 0)

	for _, u := range strings.Split(string(users), "\n") {

		ul := strings.Split(strings.TrimSpace(u), ",")
		if len(ul) < 3 {
			continue
		}

		u1 := &User{
			Uid:      ul[0],
			Name:     ul[1],
			Password: ul[2],
		}

		configByUid[u1.Uid] = u1
		configByName[u1.Name] = u1
	}

	UserCache = &util.TCache{
		Dict: make(map[int64]interface{}, 0),
		Lock: &sync.RWMutex{},
	}

}

func LoadUsers() {
	loadAllUser()
	UserId = UserCache.MaxKey()

	for uid := range configByUid {
		u := configByUid[uid]
		AddUser(&User{Uid: uid, Name: u.Name, Password: u.Password})
	}

	loadAllUser()
	UserId = UserCache.MaxKey()

	//log.Println(UserCache)
}

func GetUserByName1(name string) (*User, bool) {
	lock.RLock()
	defer lock.RUnlock()
	u, b := configByName[name]
	return u, b

}

func GetUserByName(name string) (*User, bool) {
	var result *User

	UserCache.ForeachRead(func(k int64, i interface{}) {
		u := i.(*User)
		if u.Name == name {
			result = u
			return
		}
	})

	return result, result != nil
}

func GetUserByUid1(uid string) (*User, bool) {
	lock.RLock()
	defer lock.RUnlock()

	u, b := configByUid[uid]
	return u, b

}

func GetUserByUid(uid string) (*User, bool) {
	i64, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return nil, false
	}

	i, b := UserCache.Get(i64)
	if i == nil {
		return nil, false
	}

	result := i.(*User)

	return result, b
}

func AddUser(u *User) {
	lock.Lock()
	defer lock.Unlock()

	i64, _ := strconv.ParseInt(u.Uid, 10, 64)

	if _, ok := UserCache.Get(i64); ok {
		log.Println("Adduser dumped", u.Uid)
		return
	}

	UserCache.Set(i64, u)
	bolt_put(u)
}

func AddUserWithId(u *User) *User {
	lock.Lock()
	defer lock.Unlock()

	UserId += 1
	u.Uid = strconv.FormatInt(UserId, 10)

	UserCache.Set(UserId, u)
	bolt_put(u)

	log.Println(UserCache.Get(UserId))

	return u
}

func loadAllUser() {
	// init function, does not need lock
	f := func(k, v []byte) error {
		u := &User{}
		err := json.Unmarshal(v, u)
		if err != nil {
			log.Println(err)
		}

		UserCache.UnsafeSet(util.B2I(k), u)
		return nil
	}
	Foreach(f)
}

// bolt db======================================================================

func bolt_put(r *User) {
	DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(Bucket)
		bin, _ := json.Marshal(r)
		i64, _ := strconv.ParseInt(r.Uid, 10, 64)
		err := b.Put(util.I2B(int64(i64)), bin)
		return err
	})
}

func Get(id int64) (*User, error) {
	key := util.I2B(id)

	var v []byte

	DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(Bucket)
		v = b.Get(key)
		return nil
	})

	t := &User{}

	err := json.Unmarshal(v, t)

	return t, err
}

func bolt_del(id int64) {
	DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(Bucket)
		key := util.I2B(id)
		b.Delete(key)
		return nil
	})
}

func bolt_dels(ids []int64) {
	DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(Bucket)
		for _, id := range ids {
			key := util.I2B(id)
			b.Delete(key)
		}
		return nil
	})
}

func Foreach(F func(k, v []byte) error) {
	DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(Bucket)
		b.ForEach(F)
		return nil
	})
}
