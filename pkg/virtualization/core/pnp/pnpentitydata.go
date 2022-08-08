// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package pnp

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/cimv2"
)

type PnpEntityData struct {
	*cimv2.Win32_PnPEntity
}

// NewPnpEntityData
func NewPnpEntityData(instance *wmi.WmiInstance) (*PnpEntityData, error) {
	wmivm, err := cimv2.NewWin32_PnPEntityEx1(instance)
	if err != nil {
		return nil, err
	}
	return &PnpEntityData{wmivm}, nil
}

func GetDefaultPnpEntityData(whost *host.WmiHost) (*PnpEntityData, error) {
	inst, err := instance.CreateWmiInstance(whost, string(constant.CimV2), "Win32_PnPEntity")
	if err != nil {
		return nil, err
	}
	return NewPnpEntityData(inst)
}
