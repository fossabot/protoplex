package protocols

func NewOpenVPNProtocol(targetAddress string) *Protocol {
	return &Protocol{
		Name:                    "OpenVPN",
		Target:                  targetAddress,
		MatchBytes:              [][]byte{{0x00, 0x0e, 0x38}},
		NoComparisonBeforeBytes: 3,
	}
}