package config

import (
	"strconv"
	"strings"
)

type HostPort struct {
	Host string
	Port int
}

type QueueHostPort struct {
	HostPort

	QueueName string

	Maxconn int

	Timeout int
}

func NewHostPort(hp string) HostPort {
	hparr := strings.Split(hp, ":")
	port, _ := strconv.Atoi(hparr[1])
	return HostPort{Host: hparr[0], Port: port}
}

//启动设置的选项
type Option struct {
	LogPath    string
	Businesses []string

	Zkhost string //zookeeper的Host

	QueueHostPorts []QueueHostPort //redis队列Pop

}

func NewOption(logPath string, businesses []string, zkhosts string, hostPorts []QueueHostPort) *Option {

	return &Option{LogPath: logPath, Businesses: businesses, Zkhost: zkhosts, QueueHostPorts: hostPorts}
}

//command
type Command struct {
	Action string `json:"action"`

	Params map[string]interface{} `json:"params"`
}
