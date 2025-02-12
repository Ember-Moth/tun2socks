package mobile

import "github.com/xjasonlyu/tun2socks/log"

func init() {
    log.SetLogger(log.Must(log.NewLeveled("info")))
}
