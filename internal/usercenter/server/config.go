package server

import (
	genericoptions "github.com/mindmatterlab/acex/pkg/options"
)

// Config represents the configuration options for a service, including HTTP, GRPC, and TLS settings.
type Config struct {
	HTTP genericoptions.HTTPOptions
	GRPC genericoptions.GRPCOptions
	TLS  genericoptions.TLSOptions
}
