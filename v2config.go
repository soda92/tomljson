package main

type v2config struct {
	Log       v2Log        `json:"log"`
	Inbounds  []v2Inbound  `json:"inbounds"`
	OutBounds []v2Outbound `json:"outbounds"`
	Routing   v2Routing    `json:"routing"`
}

type v2Log struct {
	LogLevel string `json:"loglevel"`
}

type v2Inbound struct {
	Listen         string           `json:"listen"`
	Port           int              `json:"port"`
	Protocol       string           `json:"protocol"`
	Settings       v2Settings       `json:"settings"`
	StreamSettings v2StreamSettings `json:"streamSettings"`
}

type v2Settings struct {
	Clients    []v2Clients `json:"clients"`
	Decryption string      `json:"decryption"`
}

type v2Clients struct {
	Id string `json:"id"`
}

type v2StreamSettings struct {
	Network   string       `json:"network"`
	WSettings v2WSSettings `json:"wsSettings"`
}

type v2WSSettings struct {
	Path string `json:"path"`
}

type v2Outbound struct {
	Tag      string             `json:"tag"`
	Protocol string             `json:"protocol"`
	Settings v2OutboundSettings `json:"settings"`
}

type v2OutboundSettings struct {
}

type v2Routing struct {
	DomainStrategy string    `json:"domainStrategy"`
	Rules          []v2Rules `json:"rules"`
}

type v2Rules struct {
	Type        string   `json:"type"`
	Ip          []string `json:"ip"`
	OutboundTag string   `json:"outboundTag"`
}
