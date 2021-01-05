package redis

/*
	集成了redis的操作方法
*/

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
	conn "github.com/mangenotwork/ManGe-Notes/conn"
)

type RDB struct{}

//string类型设置值
func (this *RDB) StringSet(key string, value interface{}) {
	rcon := conn.RedisConn()
	defer rcon.Close()
	_, err := rcon.Do("SET", key, value)
	if err != nil {
		fmt.Println("set error", err.Error())
	} else {
		fmt.Println("set ok.")
	}

}

//string 判断值
func (this *RDB) StringJudge(key string, value interface{}) (bool, error) {
	rcon := conn.RedisConn()
	defer rcon.Close()
	res, err := redis.String(rcon.Do("GET", key))
	if err != nil {
		fmt.Println("GET error", err.Error())
		return false, err
	}
	fmt.Println("res： ", res)
	fmt.Println("value ", value)
	if res == value {
		return true, nil
	}
	return false, nil

}

//创建 Hash数据
func (this *RDB) HashSet(k string, f string, v interface{}) {
	rcon := conn.RedisConn()
	defer rcon.Close()
	_, err := rcon.Do("hset", k, f, v)
	if err != nil {
		fmt.Println("hset error", err.Error())
	} else {
		fmt.Println("hset ok")
	}
}

//获取 Hash数据
func (this *RDB) HashGet(k, f string) string {
	rcon := conn.RedisConn()
	defer rcon.Close()
	res, err := redis.String(rcon.Do("hget", k, f))
	if err != nil {
		fmt.Println("hget failed", err.Error())
		return ""
	} else {
		fmt.Printf("hget value :%s\n", res)
		return res
	}
}

//删除key
func (this *RDB) DELKey(k string) {
	rcon := conn.RedisConn()
	defer rcon.Close()
	_, err := rcon.Do("del", k)
	if err != nil {
		fmt.Println("del key error", err.Error())
	} else {
		fmt.Println("del key ok")
	}

}
