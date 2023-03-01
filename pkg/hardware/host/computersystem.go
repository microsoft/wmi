// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package host

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"

	"github.com/microsoft/wmi/server2019/root/cimv2"
)

type ComputerSystem struct {
	*cimv2.Win32_ComputerSystem
}

// GetComputerSystem
func GetComputerSystem(whost *host.WmiHost) (cs *ComputerSystem, err error) {
	queryComputerSystem := query.NewWmiQuery("Win32_ComputerSystem") //query for fields from Win32_ComputerSystem class

	winComputerSystemInfo, err := instance.GetWmiInstanceEx(whost, string(constant.CimV2), queryComputerSystem)
	if err != nil {
		return nil, err
	}
	winCompSystemInstance, err := cimv2.NewWin32_ComputerSystemEx1(winComputerSystemInfo)
	if err != nil {
		return nil, err
	}
	return &ComputerSystem{winCompSystemInstance}, nil
}
