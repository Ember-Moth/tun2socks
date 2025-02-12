package tun2socks

// AndroidConfig Android配置参数
type AndroidConfig struct {
    // 基础配置
    FileDescriptor int    
    MTU           int    
    
    // 代理配置
    ProxyURL      string 
    
    // TCP配置
    EnableTCPAutoTuning bool   
    TCPSendBufferSize   string 
    TCPReceiveBufferSize string 
    
    // UDP配置
    UDPTimeout int    
    
    // 日志配置
    LogLevel string 
}
