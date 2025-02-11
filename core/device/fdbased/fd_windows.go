package fdbased

import (
	"errors"

	"github.com/xjasonlyu/tun2socks/core/device"
)

func Open(name string, mtu uint32, offset int) (device.Device, error) {
	return nil, errors.ErrUnsupported
}
