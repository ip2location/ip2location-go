package ip2location

type meta struct {
	databaseType      uint8
	databaseColumn    uint8
	databaseDay       uint8
	databaseMonth     uint8
	databaseYear      uint8
	ipv4DatabaseCount uint32
	ipv4DatabaseAddr  uint32
	ipv6DatabaseCount uint32
	ipv6DatabaseAddr  uint32
	ipv4IndexBaseAddr uint32
	ipv6IndexBaseAddr uint32
	ipv4ColumnSize    uint32
	ipv6ColumnSize    uint32
}
