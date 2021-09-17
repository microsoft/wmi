// +build windows
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"time"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/cimv2"
)

const (
	SERVICE_STOPPED          string = "Stopped"
	SERVICE_START_PENDING    string = "Start Pending"
	SERVICE_STOP_PENDING     string = "Stop Pending"
	SERVICE_RUNNING          string = "Running"
	SERVICE_CONTINUE_PENDING string = "Continue Pending"
	SERVICE_PAUSE_PENDING    string = "Pause Pending"
	SERVICE_PAUSED           string = "Paused"
)

type Win32Service struct {
	*cimv2.Win32_Service
}

func NewService(winstance *wmi.WmiInstance) (*Win32Service, error) {
	wmisvc, err := cimv2.NewWin32_ServiceEx1(winstance)
	if err != nil {
		return nil, err
	}

	return &Win32Service{wmisvc}, nil
}

func WaitForWin32Service(whost *host.WmiHost, serviceName string, timeout int64) error {
	start := time.Now()
	for {
		isRunning := false
		err := func() error {
			winsvc, err := GetWin32Service(whost, serviceName)
			if err != nil {
				return err
			}

			defer winsvc.Close()

			isRunning, err = winsvc.IsRunning()
			if err != nil {
				return err
			}

			return nil
		}()

		if err != nil {
			return err
		}

		if isRunning {
			return nil
		}

		time.Sleep(5 * time.Second)

		if time.Since(start) > time.Duration(timeout*int64(time.Second)) {
			break
		}
	}

	return errors.Wrapf(errors.Timedout, "The service [%s] is not up after 300 seconds", serviceName)
}

func GetWin32Service(whost *host.WmiHost, serviceName string) (*Win32Service, error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("Win32_Service", "Name", serviceName)
	service, err := cimv2.NewWin32_ServiceEx6(whost.HostName, string(constant.CimV2), creds.UserName, creds.Password, creds.Domain, query)

	if err != nil {
		return nil, err
	}

	return &Win32Service{service}, nil
}

func (ws *Win32Service) IsRunning() (bool, error) {
	state, err := ws.GetPropertyState()
	if err != nil {
		return false, err
	}

	if state == SERVICE_RUNNING {
		return true, nil
	}

	return false, nil
}
