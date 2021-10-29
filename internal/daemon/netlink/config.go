package netlink

type Config struct {
	InternalBridgeName string
	ExternalBridgeName string

	InternalInterfaceName string
	ExternalInterfaceName string

	InternalIPCIDR string
	ExternalIPCIDR string
}