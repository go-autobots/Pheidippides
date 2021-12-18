package server

import (
	"github.com/google/wire"
)

// ProviderSet is server providers.
var SrvSet = wire.NewSet(NewGRPCServer)
