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
	deserializedStr := strings.Split(string(buffer), "\r\n")
	return deserializedStr[:len(deserializedStr) - 1]
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
		return Set(commandArgs[4], commandArgs[6:])
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

func Set(key string, valueMetaData []string) string {
	lock.Lock()
	defer lock.Unlock()
	value := Value{value: valueMetaData[0], createdAt: time.Now()}
	if len(valueMetaData) > 1 {
		fmt.Println(valueMetaData[2])
		expiryInfo := map[string]string{"expiredAfter": valueMetaData[4], "expriedCommand": valueMetaData[2]}

		expiredAt, e := getExpiredTime(expiryInfo, value.createdAt)
		if e != nil {
			log.Fatal(e.Error())
		}
		value.expiredAt = expiredAt
	}

	REDIS[key] = value
	return "+OK"
}

func getExpiredTime(expryInfo map[string]string, createdAt time.Time) (time.Time, error) {

	expiredTime, e := strconv.Atoi(expryInfo["expiredAfter"])
	if e != nil {
		log.Fatal(e.Error())
	}

	switch strings.ToLower(expryInfo["expriedCommand"]) {
	case "ex":
		return createdAt.Add(time.Second * time.Duration(expiredTime)), nil
	case "px":
		return createdAt.Add(time.Millisecond * time.Duration(expiredTime)), nil
	case "exat":
		return time.Unix(int64(expiredTime), 0), nil
	case "pxat":
		return time.UnixMilli(int64(expiredTime)), nil
	}
	return time.Time{}, errors.New("Invalid command")
}

func Get(key string) string {
	lock.RLock()
	defer lock.RUnlock()
	value, exist := REDIS[key]
	if !exist || time.Now().After(value.expiredAt) {
		return "_"
	}

	return fmt.Sprintf("+%s", value.value)
}
