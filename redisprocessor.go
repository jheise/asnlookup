package main

import (
	"strings"
	// external
	"github.com/garyburd/redigo/redis"
)

type RedisASNProcessor struct {
	requests chan *ASNRequest
	rconn    redis.Conn
	connstr  string
	timeout  string
	server   string
	port     string
}

func NewRedisASNProcessor(requests chan *ASNRequest, conn string, timeout string, server string, port string) *RedisASNProcessor {
	processor := new(RedisASNProcessor)
	processor.requests = requests
	processor.connstr = conn
	processor.timeout = timeout
	processor.server = server
	processor.port = port

	rconn, err := redis.Dial("tcp", conn)
	if err != nil {
		panic(err)
	}

	processor.rconn = rconn
	return processor
}

func (self *RedisASNProcessor) Process() {
	for request := range self.requests {
		var response *ASNResponse

		ipaddr := request.IPAddr
		data, err := redis.String(self.rconn.Do("GET", ipaddr))
		if err != nil {
			if strings.Compare(err.Error(), "redigo: nil returned") == 0 {
				answers, err := QueryServer(self.server, self.port, request.IPAddr, 30)
				if err != nil {
					self.rconn.Close()
					panic(err)
				}

				_, err = self.rconn.Do("SET", ipaddr, answers[3])
				if err != nil {
					self.rconn.Close()
					panic(err)
				}

				_, err = self.rconn.Do("EXPIRE", ipaddr, timeout)
				if err != nil {
					self.rconn.Close()
					panic(err)
				}
				response = NewASNResponse(answers[3])

			} else {
				panic(err)
			}
		} else {
			response = NewASNResponse(data)
		}
		request.SendResponse(response)
	}
}
