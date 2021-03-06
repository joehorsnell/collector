// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !windows,!plan9

package syslog

import (
	"crypto/tls"
	"errors"
	"net"
)

// unixSyslog opens a connection to the syslog daemon running on the
// local machine using a Unix domain socket.

func localSyslog() (conn serverConn, err error) {
	logTypes := []string{"unixgram", "unix"}
	logPaths := []string{"/dev/log", "/var/run/syslog"}
	for _, network := range logTypes {
		for _, path := range logPaths {
			conn, err := net.Dial(network, path)
			if err != nil {
				continue
			} else {
				return &netConn{conn: conn, local: true}, nil
			}
		}
	}
	return nil, errors.New("Unix syslog delivery error")
}

func dial(network, address string, tlsCfg *tls.Config) (net.Conn, error) {
	if tlsCfg != nil && network == "tcp" {
		return tls.Dial(network, address, tlsCfg)
	} else {
		return net.Dial(network, address)
	}
}
