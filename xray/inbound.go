package xray

// todo: change the following Equals
type InboundConfig struct {
	Tag  string `json:"tag"`
	Port int    `json:"port"`
	/* Listen         json_util.RawMessage `json:"listen"` // listen cannot be an empty string

	Protocol       string               `json:"protocol"`
	Settings       json_util.RawMessage `json:"settings"`
	StreamSettings json_util.RawMessage `json:"streamSettings"`

	Sniffing       json_util.RawMessage `json:"sniffing"`
	Allocate       json_util.RawMessage `json:"allocate"` */
}

// todo： complete Equals
func (c *InboundConfig) Equals(other *InboundConfig) bool {
	if c.Tag != other.Tag {
		return false
	}
	if c.Port != other.Port {
		return false
	}
	/* 	if !bytes.Equal(c.Listen, other.Listen) {
	   		return false
	   	}

	   	if c.Protocol != other.Protocol {
	   		return false
	   	}
	   	if !bytes.Equal(c.Settings, other.Settings) {
	   		return false
	   	}
	   	if !bytes.Equal(c.StreamSettings, other.StreamSettings) {
	   		return false
	   	}

	   	if !bytes.Equal(c.Sniffing, other.Sniffing) {
	   		return false
	   	}
	   	if !bytes.Equal(c.Allocate, other.Allocate) {
	   		return false
	   	} */
	return true
}
