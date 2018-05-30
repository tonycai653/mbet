package websock

import (
	"fmt"
	"github.com/gorilla/websocket"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

const (
	e = "abcdefghijklmnopqrstuvwxyz012345"
)

func NumberString(num int) string {
	ll := len(strconv.Itoa(num - 1))
	s := strings.Repeat("0", ll+1) + strconv.Itoa(Number(num))
	return s[len(s)-ll:]
}

func Number(num int) int {
	return int(math.Floor(rand.Float64() * float64(num)))
}

func RandomString(num int) string {
	c := RandomBytes(num)
	f := make([]string, num)
	for g := 0; g < num; g++ {
		ind := int(c[g]) % len(e)
		f = append(f, e[ind:ind+1])
	}
	return strings.Join(f, "")
}

func RandomBytes(num int) []byte {
	b := make([]byte, num)
	for ind, _ := range b {
		b[ind] = byte(math.Floor(256 * rand.Float64()))
	}
	return b[:]
}

func WebsocketUrl() string {
	sv := NumberString(1e3)
	rs := RandomString(8)
	return fmt.Sprintf("wss://www.mbet.com/zh/websocket/endpoint/%s/%s/websocket", sv, rs)
}

func CreateWebsocketConn(url string, reqHead http.Header) (*websocket.Conn, *http.Response, error) {
	fmt.Printf("connecing to %s\n", url)
	return websocket.DefaultDialer.Dial(url, reqHead)
}
