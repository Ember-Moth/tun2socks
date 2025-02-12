package tun2socks

import (
    "fmt"
    "time"
    "github.com/xjasonlyu/tun2socks/engine"
)

// Tun2socks Android绑定类
type Tun2socks struct {
    config  *AndroidConfig
    started bool
}

// NewTun2socks 创建新实例
func NewTun2socks() *Tun2socks {
    return &Tun2socks{}
}

// Configure 配置服务
func (t *Tun2socks) Configure(config *AndroidConfig) error {
    if t.started {
        return fmt.Errorf("cannot configure while running")
    }
    
    if config.FileDescriptor <= 0 {
        return fmt.Errorf("invalid file descriptor: %d", config.FileDescriptor)
    }
    if config.ProxyURL == "" {
        return fmt.Errorf("proxy URL is required")
    }
    
    if config.MTU <= 0 {
        config.MTU = 1500
    }
    if config.LogLevel == "" {
        config.LogLevel = "info"
    }
    if config.UDPTimeout <= 0 {
        config.UDPTimeout = 30
    }
    
    t.config = config
    return nil
}

// Start 启动服务
func (t *Tun2socks) Start() error {
    if t.started {
        return fmt.Errorf("already started")
    }
    if t.config == nil {
        return fmt.Errorf("configuration required")
    }

    key := &engine.Key{
        Device: fmt.Sprintf("fd://%d", t.config.FileDescriptor),
        MTU:    t.config.MTU,
        Proxy:  t.config.ProxyURL,
        
        TCPModerateReceiveBuffer: t.config.EnableTCPAutoTuning,
        TCPSendBufferSize:       t.config.TCPSendBufferSize,
        TCPReceiveBufferSize:    t.config.TCPReceiveBufferSize,
        
        UDPTimeout: time.Duration(t.config.UDPTimeout) * time.Second,
        LogLevel:   t.config.LogLevel,
        
        // 禁用不需要的功能
        RestAPI: "",
        Interface: "",
        TUNPreUp: "",
        TUNPostUp: "",
        Mark: 0,
    }

    engine.Insert(key)
    engine.Start()

    t.started = true
    return nil
}

// Stop 停止服务
func (t *Tun2socks) Stop() {
    if !t.started {
        return
    }
    engine.Stop()
    t.started = false
    t.config = nil
}

// IsRunning 检查服务状态
func (t *Tun2socks) IsRunning() bool {
    return t.started
}
