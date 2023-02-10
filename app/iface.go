package app

import (
	"fmt"
	"net"
	"strings"
)

func printAddressesWithPort(port int) {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(errStyle.Render(fmt.Sprintf("Err: Unable to find interfaces of this machine: %v\n", err)))
	} else {
		for i := range ifaces {
			if ifaces[i].Name == "lo" || strings.HasPrefix(ifaces[i].Name, "docker") || strings.HasPrefix(ifaces[i].Name, "br-") || strings.HasPrefix(ifaces[i].Name, "lxcbr") {
				continue
			}
			var tmpl strings.Builder
			addrs, _ := ifaces[i].Addrs()

      var ipv4Addresses []string

      for j := range addrs {
			  if !strings.Contains(addrs[j].String(), ":") {
				  ipv4Addresses = append(ipv4Addresses, addrs[j].String())
			  }
      }

			if len(ipv4Addresses) == 0 {
				continue
			}

      tmpl.WriteString(fmt.Sprintf("%v:\n", ifaces[i].Name))
			
      for j := range ipv4Addresses {
				ip, _, _ := strings.Cut(addrs[j].String(), "/")

				tmpl.WriteString(fmt.Sprintf("- %v:%d\n", ip, port))

			}
			fmt.Println(interfaceStyle.Render(strings.TrimSpace(tmpl.String())))
		}
	}
}
