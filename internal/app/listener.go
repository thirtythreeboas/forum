package app

import (
	"fmt"
	"forum/internal/config"
	"net"
)

type listeners struct {
	http net.Listener
}

func newListener(cfg *config.HTTPServer) (*listeners, error) {
	l := &listeners{}

	http, err := listenPublicHTTP(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to creatre listener %w", err)
	}

	l.http = http

	return l, nil
}

func listenPublicHTTP(cfg *config.HTTPServer) (net.Listener, error) {
	listenerAddr := fmt.Sprintf("%s:%v", cfg.AppBindAddress, cfg.AppPort)

	http, err := net.Listen("tcp", listenerAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP listener %w", err)
	}

	return http, nil
}
