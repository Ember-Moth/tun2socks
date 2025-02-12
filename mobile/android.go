package mobile

import "github.com/xjasonlyu/tun2socks/log"
import "go.uber.org/zap/zapcore"

func init() {
    log.SetLogger(log.Must(log.NewLeveled("zapcore.InfoLevel")))
}
