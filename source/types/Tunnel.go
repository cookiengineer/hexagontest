package types

type Tunnel struct {
	Address  string   `json:"host"`
	Port     uint16   `json:"port"`
	Protocol Protocol `json:"protocol"`
}

func (tunnel *Tunnel) SetAddress(value string) {

	if IsIPv4(value) {

		ipv4 := ParseIPv4(value)

		if ipv4 != nil {
			tunnel.Address = ipv4.String()
		}

	} else if IsIPv6(value) {

		ipv6 := ParseIPv6(value)

		if ipv6 != nil {
			tunnel.Address = ipv6.String()
		}

	}

}

func (tunnel *Tunnel) SetPort(value uint16) {

	if value > 0 && value < 65535 {
		tunnel.Port = value
	}

}

func (tunnel *Tunnel) SetProtocol(value Protocol) {
	tunnel.Protocol = value
}
