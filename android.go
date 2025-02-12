package tun2socks

import (
    "go.uber.org/zap/zapcore"
    "github.com/Ember-Moth/tun2socks/log"
)

func init() {
    log.SetLogger(log.Must(log.NewLeveled(zapcore.InfoLevel)))
}
