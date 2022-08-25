// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package host

import (
	"strconv"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/server2019/root/cimv2"
)

type OperatingSystemInfo struct {
	OperatingSystemSKU uint64
}

// GetFreePhysicalMemory
func GetOperatingSystemInfo(whost *host.WmiHost) (osInfo *OperatingSystemInfo, err error) {
	query := query.NewWmiQuery("Win32_OperatingSystem")
	osinstance, err := instance.GetWmiInstanceEx(whost, string(constant.CimV2), query)
	if err != nil {
		return
	}
	defer osinstance.Close()

	tmpInstance, err1 := cimv2.NewWin32_OperatingSystemEx1(osinstance)
	if err1 != nil {
		err = err1
		return
	}

	osSku, err1 := tmpInstance.GetProperty("OperatingSystemSKU")
	if err1 != nil {
		err = err1
		return
	}

	osSkuSize, err1 := strconv.ParseUint(osSku.(string), 10, 64)
	if err1 != nil {
		err = err1
		return
	}

	return &OperatingSystemInfo{
		OperatingSystemSKU: osSkuSize,
	}, nil
}
