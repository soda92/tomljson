package main

type v2struct struct {
	log       v2Log
	inbounds  []v2Inbound
	outBounds []v2Outbound
}

type v2Log struct {
	logLevel string
}

type v2Inbound struct {
	listen         string
	port           int
	protocol       string
	settings       v2Settings
	streamSettings v2StreamSettings
}

type v2Settings struct {
	clients    []v2Clients
	decryption string
}

type v2Clients struct {
	id string
}

type v2StreamSettings struct {
	network    string
	wsSettings v2WSSettings
}

type v2WSSettings struct {
	path string
}

type v2Outbound struct {
	tag      string
	protocol string
}

type v2Routing struct {
	domainStrategy string
	roles          []v2Rules
}

type v2Rules struct {
	_type       string
	ip          []string
	outboundTag string
}
