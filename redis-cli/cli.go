package cli

import (
	"net"
	"sync"
)

type RedisCli struct {
	mu   sync.Mutex
	cn   net.Conn
	vars sync.Map
}
