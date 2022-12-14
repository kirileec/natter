package main

import (
	"fmt"
	"github.com/ccding/go-stun/stun"
	"regexp"
)
import (
	"github.com/linxlib/logs"
)

var (
	match      = regexp.MustCompile(`\d`)
	serverList = map[int]string{
		1: "stun.nextcloud.com:3478",
		2: "stun.syncthing.net:3478",
		3: "stun.qq.com:3478",
		4: "stun.miwifi.com:3478",
		5: "stun.bige0.com:3478",
		6: "stun.stunprotocol.org:3478",
	}
)

func main() {
	logs.Println("通过stun服务器探测当前环境NAT类型")
	c := stun.NewClient()
	for i := 0; i < len(serverList)-1; i++ {
		fmt.Printf("%d. %s\n", i+1, serverList[i+1])
		c.SetServerAddr(serverList[i+1])
		nat, host, err := c.Discover()
		if err != nil {
			logs.Error(err)
			return
		}
		logs.Debugln(nat)
		logs.Debugln(host)
	}
	logs.Printf("\n结果供参考\n")
	fmt.Scanln()

}
