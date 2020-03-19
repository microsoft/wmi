// +build windows
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package session

import (
	"fmt"
	"os"
	"strings"

	"github.com/microsoft/wmi/pkg/base/credential"
	"github.com/microsoft/wmi/pkg/base/host"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

var (
	sessionManager *wmi.WmiSessionManager
	sessionsMap    map[string]*wmi.WmiSession
)

// StartWMI
func StartWMI() {
	sessionsMap = make(map[string]*wmi.WmiSession)
	sessionManager = wmi.NewWmiSessionManager()
}

// StopWMI
func StopWMI() {
	for key := range sessionsMap {
		sessionsMap[key].Dispose()
		sessionsMap[key] = nil
	}

	if sessionManager != nil {
		sessionManager.Dispose()
		sessionManager = nil
	}
}

// GetHostSession
func GetHostSession(sessionName string, whost *host.WmiHost) (*wmi.WmiSession, error) {
	cred := whost.GetCredential()
	return GetSession(sessionName, whost.HostName, cred.Domain, cred.UserName, cred.Password)
}

func GetHostSessionWithCredentials(sessionName string, whost *host.WmiHost, cred *credential.WmiCredential) (*wmi.WmiSession, error) {
	return GetSession(sessionName, whost.HostName, cred.Domain, cred.UserName, cred.Password)
}

// GetSession
func GetSession(sessionName string, serverName string, domain string, userName string, password string) (*wmi.WmiSession, error) {
	sessionsMapId := strings.Join([]string{sessionName, serverName, domain}, "_")
	if sessionsMap[sessionsMapId] == nil {
		var err error
		sessionsMap[sessionsMapId], err = createSession(sessionName, serverName, domain, userName, password)
		if err != nil {
			return nil, err
		}
	}

	return sessionsMap[sessionsMapId], nil
}

var (
	localHostName string
)

////////////// Private functions ////////////////////////////
func createSession(sessionName string, serverName string, domain string, username string, password string) (*wmi.WmiSession, error) {
	if len(localHostName) == 0 {
		localHostName, _ = os.Hostname()
	}

	// TODO: ideally, we should also compare the domain here.
	// that said, this is low priority as cross-domain WMI calls are rare
	if strings.EqualFold(localHostName, serverName) {
		// Optimization for local clusters: connecting to the local cluster through remote WMI results in a much longer
		// response than connecting directly. When providing the cluster name, the cluster has to go through a
		// long sequence of connection/authentication. Not providing the name allows the cluster to skip that
		// expensive sequence.
		serverName = ""
		domain = ""
	}

	session, err := sessionManager.GetSession(sessionName, serverName, domain, username, password)
	if err != nil {
		return nil, fmt.Errorf("Failed getting the WMI session for " + sessionName)
	}

	connected, err := session.Connect()

	if !connected || err != nil {
		return nil, fmt.Errorf("Failed connecting to the WMI session for " + sessionName)
	}

	return session, nil
}
