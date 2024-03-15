// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package snapshot

import (
	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/constant"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type VirtualSystemSnapshotSettingData struct {
	*v2.Msvm_VirtualSystemSnapshotSettingData
}

func NewVirtualSystemSnapshotSettingData(instance *wmi.WmiInstance) (*VirtualSystemSnapshotSettingData, error) {
	wmientity, err := v2.NewMsvm_VirtualSystemSnapshotSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &VirtualSystemSnapshotSettingData{wmientity}, nil
}

func GetDefaultVirtualSystemSnapshotSettingData(whost *host.WmiHost) (*VirtualSystemSnapshotSettingData, error) {
	inst, err := instance.CreateWmiInstance(whost, string(constant.Virtualization), "Msvm_VirtualSystemSnapshotSettingData")
	if err != nil {
		return nil, err
	}
	return NewVirtualSystemSnapshotSettingData(inst)

}

func GetVirtualSystemSnapshotSettingData(whost *host.WmiHost, name string) (*VirtualSystemSnapshotSettingData, error) {
	snapshotSettings, err := GetDefaultVirtualSystemSnapshotSettingData(whost)
	if err != nil {
		return nil, err
	}
	snapshotSettings.SetPropertyElementName(name)

	return snapshotSettings, err
}
