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

type OperatingSystem struct {
	*cimv2.Win32_OperatingSystem
}

// GetOperatingSystem
func GetOperatingSystem(whost *host.WmiHost) (*OperatingSystem, error) {
	query := query.NewWmiQuery("Win32_OperatingSystem")
	osinstance, err := instance.GetWmiInstanceEx(whost, string(constant.CimV2), query)
	if err != nil {
		return nil, err
	}

	tmpinstance, err := cimv2.NewWin32_OperatingSystemEx1(osinstance)
	if err != nil {
		return nil, err
	}
	return &OperatingSystem{tmpinstance}, nil
}
