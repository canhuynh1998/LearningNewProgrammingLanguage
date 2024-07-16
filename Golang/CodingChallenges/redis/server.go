package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Value struct {
	value     string
	createdAt time.Time
	expiredAt time.Time
}

var REDIS = map[string]Value{}
var lock = sync.RWMutex{}

func main() {
	fmt.Println("Hello")
	l, err := net.Listen("tcp", "0.0.0.0:6380")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go resp(conn)
	}

}

func resp(conn net.Conn) {
	defer conn.Close()
	for {
		buffer := make([]byte, 2048)
		_, err := conn.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				return
			}
			log.Fatal(err)
		}
		commandArgs := deserialize(buffer)
		resp := respond(commandArgs)
		conn.Write([]byte(serialize(resp)))
	}
}

func deserialize(buffer []byte) []string {
	return strings.Split(string(buffer), "\r\n")
}

func serialize(respond string) string {
	return fmt.Sprintf("%s\r\n", respond)
}

func respond(commandArgs []string) string {

	switch strings.ToLower(commandArgs[2]) {
	case "ping":
		return Ping()
	case "echo":
		return Echo(commandArgs[4])
	case "set":
		return Set(commandArgs[4], commandArgs[6])
	case "get":
		return Get(commandArgs[4])
	}
	return fmt.Sprintf("-ERR unknown command '%s', with args beginning with:", commandArgs[2])
}

func Ping() string {
	return "+PONG"
}

func Echo(s string) string {
	return fmt.Sprintf("+%s", s)
}

func Set(key string, valueMeta ...string) string {
	lock.Lock()
	defer lock.Unlock()
	// timeData, err := getExpiredTime(valueMeta[1], valueMeta[2])
	// if err != nil {
	// 	return fmt.Sprintf("-%s", err.Error())
	// }
	// REDIS[key] = Value{value: valueMeta[0], createdAt: timeData["createdAt"], expiredAt: timeData["expiredAt"]}
	REDIS[key] = Value{value: valueMeta[0], createdAt: time.Now()}
	return "+OK"
}

func getExpiredTime(expriedCommand string, expriedAfter string) (map[string]time.Time, error) {
	result := make(map[string]time.Time)
	result["createdAt"] = time.Now()
	expiredTime, e := strconv.Atoi(expriedAfter)
	if e != nil {
		return nil, e
	}
	if expriedCommand == "EX" {
		result["expiredAt"] = result["createdAt"].Add(time.Second * time.Duration(expiredTime))
	} else if expriedCommand == "PX" {
		result["expiredAt"] = result["createdAt"].Add(time.Millisecond * time.Duration(expiredTime))
	} else if expriedCommand == "EXAT" {
		result["expiredAt"] = time.Unix(int64(expiredTime), 0)
	} else if expriedCommand == "PXAT" {
		result["expiredAt"] = time.UnixMilli(int64(expiredTime))
	} else {
		return nil, errors.New("Invalid command")
	}
	return result, nil
}

func Get(key string) string {
	lock.RLock()
	defer lock.RUnlock()
	value, exist := REDIS[key]
	if !exist {
		return "_"
	}
	return fmt.Sprintf("+%s", value.value)
}
