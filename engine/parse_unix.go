//go:build unix

package engine

import (
	"net/url"

	"github.com/Ember-Moth/tun2socks/core/device"
	"github.com/Ember-Moth/tun2socks/core/device/tun"
)

func parseTUN(u *url.URL, mtu uint32) (device.Device, error) {
	return tun.Open(u.Host, mtu)
}
