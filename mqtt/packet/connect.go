// Direct conversion from my Elixir implementation
package packet

import "bytes"

const (
	packetType    = 1
	protocolName  = "MQTT"
	protocolLevel = 4 // This is MQTT v3.1.1
	clientID      = "GoMQTT"
)

// Connect MQTT 3.1.1 control packet
type Connect struct {
	keepalive int
	clientID  string
}

// PacketType return packet type numerical value
func (c *Connect) PacketType() int {
	return packetType
}

// Marshall return buffer containing serialized CONNECT MQTT control packet
func (c *Connect) Marshall() bytes.Buffer {
	var variablePart bytes.Buffer
	var packet bytes.Buffer

	fixedHeaderFlags := 0
	connectFlags := 0         // TODO: support connect flag definition
	var keepalive uint16 = 30 // TODO: make it configurable

	variablePart.Write(encodeString(protocolName))
	variablePart.WriteByte(byte(protocolLevel))
	variablePart.WriteByte(byte(connectFlags))
	variablePart.Write(encodeUint16(keepalive))
	variablePart.Write(encodeString(clientID))

	fixedHeader := (packetType<<4 | fixedHeaderFlags)
	packet.WriteByte(byte(fixedHeader))
	packet.WriteByte(byte(variablePart.Len()))
	packet.Write(variablePart.Bytes())

	return packet
}
