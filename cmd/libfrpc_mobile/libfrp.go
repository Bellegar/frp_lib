package libfrp

import (
	"runtime/debug"

	frpc "frp_lib/cmd/frpc/sub"
	"frp_lib/pkg/util/version"
)
func RunFrpc(cfgFilePath string) (err error) {
	return frpc.RunFrpc(cfgFilePath)
}

func StopFrpc() (err error) {
	defer debug.FreeOSMemory()
	return frpc.StopFrp()
}

func IsFrpcRunning() bool {
	return frpc.IsFrpRunning()
}

// Version frp
func Version() string {
	return version.Full()
}
