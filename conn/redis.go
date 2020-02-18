package conn

/*
	集成了redis 连接，操作等方法
	Redis的连接方式
	1. 普通连接
	2. SSH普通连接
	3. 连接池连接
	4. SSH连接池连接
*/

import (
	_ "encoding/json"
	"fmt"
	"net"
	_ "reflect"
	_ "strconv"
	_ "strings"
	"time"
	_ "unsafe"

	"github.com/garyburd/redigo/redis"
	"golang.org/x/crypto/ssh"
)

const (
	RedisMaxIdle        = 10
	RedisMaxActive      = 10
	RedisIdleTimeoutSec = 20
)

// getSSHClient 连接ssh
// addr : 主机地址, 如: 127.0.0.1:22
// user : 用户
// pass : 密码
// 返回 ssh连接
func getSSHClient(user string, pass string, addr string) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}
	sshConn, err := net.Dial("tcp", addr)
	if nil != err {
		fmt.Println("net dial err: ", err)
		return nil, err
	}

	clientConn, chans, reqs, err := ssh.NewClientConn(sshConn, addr, config)
	if nil != err {
		sshConn.Close()
		fmt.Println("ssh client conn err: ", err)
		return nil, err
	}

	client := ssh.NewClient(clientConn, chans, reqs)
	return client, nil
}

// RConn 普通连接
// ip : redis服务地址
// port:  Redis 服务端口
// password  Redis 服务密码
// 返回redis连接
func RConn(ip string, port int, password string) (redis.Conn, error) {
	host := fmt.Sprintf("%s:%d", ip, port)
	conn, err := redis.Dial("tcp", host)
	if nil != err {
		fmt.Println("dial to redis addr err: ", err)
		return nil, err
	}

	fmt.Println(conn)
	if _, authErr := conn.Do("AUTH", password); authErr != nil {
		fmt.Println("redis auth password error: ", authErr)
		return nil, fmt.Errorf("redis auth password error: %s", authErr)
	}

	return conn, nil
}

// RSSHConn SSH普通连接
// addr : SSH主机地址, 如: 127.0.0.1:22
// user : SSH用户
// pass : SSH密码
// ip : redis服务地址
// port:  Redis 服务端口
// password  Redis 服务密码
// 返回redis连接
func RSSHConn(user string, pass string, addr string, ip string, port int, password string) (redis.Conn, error) {
	sshClient, err := getSSHClient(user, pass, addr)
	if nil != err {
		fmt.Println(err)
		return nil, err
	}

	host := fmt.Sprintf("%s:%d", ip, port)
	conn, err := sshClient.Dial("tcp", host)
	if nil != err {
		fmt.Println("dial to redis addr err: ", err)
		return nil, err
	}

	redisConn := redis.NewConn(conn, -1, -1)

	if _, authErr := redisConn.Do("AUTH", password); authErr != nil {
		fmt.Println("redis auth password error: ", authErr)
		return nil, fmt.Errorf("redis auth password error: %s", authErr)
	}

	return redisConn, nil
}

// RPool 连接池连接
// ip : redis服务地址
// port:  Redis 服务端口
// password  Redis 服务密码
// 配置参数  RedisMaxIdle 最大连接
// 配置参数  RedisMaxActive 最大连接数
// 配置参数  RedisIdleTimeoutSec 设置超时
// 返回redis连接池  调用: c := RPool().Get() 返回redis连接
func RPool(ip string, port int, password string) *redis.Pool {
	redisURL := fmt.Sprintf("redis://%s:%d", ip, port)
	fmt.Println("redisURL -> ", redisURL)
	return &redis.Pool{
		MaxIdle:     RedisMaxIdle,
		MaxActive:   RedisMaxActive,
		IdleTimeout: time.Duration(RedisIdleTimeoutSec) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(redisURL)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}

			//验证redis密码
			if _, authErr := c.Do("AUTH", password); authErr != nil {
				return nil, fmt.Errorf("redis auth password error: %s", authErr)
			}

			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}

			return nil
		},
	}
}

// RSSHPool SSH连接池连接
// addr : SSH主机地址, 如: 127.0.0.1:22
// user : SSH用户
// pass : SSH密码
// ip : redis服务地址
// port:  Redis 服务端口
// password  Redis 服务密码
// 配置参数  RedisMaxIdle 最大连接
// 配置参数  RedisMaxActive 最大连接数
// 配置参数  RedisIdleTimeoutSec 设置超时
// 返回redis连接池  调用: c := RSSHPool().Get() 返回redis连接
func RSSHPool(user string, pass string, addr string, ip string, port int, password string) *redis.Pool {
	sshClient, err := getSSHClient(user, pass, addr)
	if nil != err {
		fmt.Println(err)
		return nil
	}

	redisURL := fmt.Sprintf("%s:%d", ip, port)
	return &redis.Pool{
		MaxIdle:     RedisMaxIdle,
		MaxActive:   RedisMaxActive,
		IdleTimeout: time.Duration(RedisIdleTimeoutSec) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := sshClient.Dial("tcp", redisURL)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}

			fmt.Println(c)
			redisc := redis.NewConn(c, -1, -1)
			if _, authErr := redisc.Do("AUTH", password); authErr != nil {
				return nil, fmt.Errorf("redis auth password error: %s", authErr)
			}

			return redisc, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}

			return nil
		},
	}
}

func RedisConn() redis.Conn {
	dbnumber := 2
	conn := RPool("127.0.0.1",6379,"123").Get()
	_, err := conn.Do("select", fmt.Sprintf("%d", dbnumber))
	fmt.Println("Redis Error : ", err)
	return conn
}