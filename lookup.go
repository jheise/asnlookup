package main

import (
	"io/ioutil"
	"net"
	"strings"
	"time"
)

func QueryServer(server string, port string, query string, timeout int) ([]string, error) {
	var data []byte
	var fields []string

	duration := time.Duration(timeout) * time.Second
	connection, err := net.DialTimeout("tcp", net.JoinHostPort(server, port), duration)
	if err != nil {
		return fields, err
	}
	defer connection.Close()

	connection.Write([]byte(query + "\r\n"))
	data, err = ioutil.ReadAll(connection)
	if err != nil {
		return fields, err
	}

	info := strings.TrimSpace(string(data[:]))
	for _, line := range strings.Split(info, "\n") {
		for _, field := range strings.Split(line, "|") {
			fields = append(fields, strings.TrimSpace(field))
		}
	}

	return fields, nil
}
