package common

import (
  "net"
)

// LocalIP returns the non loopback IP of the host
// https://stackoverflow.com/a/31551220/6998584
func LocalIP() string {
  addrs, err := net.InterfaceAddrs()
  if err != nil {
    return ""
  }
  for _, add := range addrs {
    if ipnet, ok := add.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
      if ipnet.IP.To4() != nil {
        return ipnet.IP.String()
      }
    }
  }
  return ""
}
