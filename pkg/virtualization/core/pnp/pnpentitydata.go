// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package pnp

import (
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	"github.com/microsoft/wmi/server2019/root/cimv2"
)

type PnpEntityData struct {
	*cimv2.Win32_PnPEntity
}

func NewPnpEntityData(instance *wmi.WmiInstance) (*PnpEntityData, error) {
	wmivm, err := cimv2.NewWin32_PnPEntityEx1(instance)
	if err != nil {
		return nil, err
	}
	return &PnpEntityData{wmivm}, nil
}
