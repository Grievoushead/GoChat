package redis

import (
    "fmt"
    "time"

    "gopkg.in/go-redis/cache.v1"
    "gopkg.in/redis.v3"
    "gopkg.in/vmihailenco/msgpack.v2"
)

type Object struct {
    Str string
    Num int
}

func GetResult() int {
    ring := redis.NewRing(&redis.RingOptions{
        Addrs: map[string]string{
            "localhost": ":6379",
            //"server2": ":6379",
        },

        DialTimeout:  3 * time.Second,
        ReadTimeout:  time.Second,
        WriteTimeout: time.Second,
    })

    codec := &cache.Codec{
        Ring: ring,

        Marshal: func(v interface{}) ([]byte, error) {
            return msgpack.Marshal(v)
        },
        Unmarshal: func(b []byte, v interface{}) error {
            return msgpack.Unmarshal(b, v)
        },
    }

    key := "mykey"
    obj := &Object{
        Str: "mystring",
        Num: 42,
    }

    codec.Set(&cache.Item{
        Key:        key,
        Object:     obj,
        Expiration: time.Hour,
    })

    var wanted Object
    if err := codec.Get(key, &wanted); err == nil {
        fmt.Println(wanted)
    }

    return wanted.Num
    // Output: {mystring 42}
}
