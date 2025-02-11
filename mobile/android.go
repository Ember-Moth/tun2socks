package mobile

import "github.com/xjasonlyu/tun2socks/v2/log"

func init() {
    log.SetLogger(log.Must(log.NewLeveled("info")))
}
