package spec

type WgDump struct {
	Interface     string `json:"interface"`     // e.g., wg0
	PublicKey     string `json:"publicKey"`     // e.g., AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=
	PrivateKey    string `json:"privateKey"`    // e.g., BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB
	Port          string `json:"port"`          // e.g., 51820
	Endpoint      string `json:"endpoint"`      // e.g., 10.10.0.x/32
	TransferRx    string `json:"transferRx"`    // e.g., 0
	TransferTx    string `json:"transferTx"`    // e.g., 0
	LastHandshake string `json:"lastHandshake"` // e.g., 0
	AllowedIPs    string `json:"allowedIps"`    // Allowed IPs in CIDR notation (e.g., 10.10.0.x/32)
	Flags         string `json:"flags"`         // e.g., "off"
}
