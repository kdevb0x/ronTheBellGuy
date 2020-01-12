// Copyright (C) 2019-2020 Kdevb0x Ltd.
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package server

import (
	"net"
	"net/http"
)

type Server struct {
	http.Server
	T TemplateServer
}

func NewServer(addr string) *Server {
	s := new(Server)
	s.Addr = addr
	s.T = *NewTemplateServer()
	s.Handler = s.T.Mux

	// TODO: set timeouts

	return s
}

var _ net.Listener

var _ http.FileSystem
