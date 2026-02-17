// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package pcie

import (
	"github.com/microsoft/wmi/pkg/errors"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
)

type PcieDevice struct {
	*PciExpressSettingData
}

func NewPcieDevice(instance *wmi.WmiInstance) (*PcieDevice, error) {
	wmivm, err := NewPciExpressSettingData(instance)
	if err != nil {
		return nil, err
	}
	return &PcieDevice{wmivm}, nil
}

func (pcie *PcieDevice) CloneEx1() (*PcieDevice, error) {
	tmp, err := pcie.Clone()
	if err != nil {
		return nil, err
	}
	return NewPcieDevice(tmp)
}

func (device *PcieDevice) GetPath() (string, error) {
	value, err := device.GetProperty("HostResource")
	if err != nil {
		return "", err
	}
	hostResources, ok := value.([]interface{})
	if !ok || hostResources == nil {
		return "", errors.Wrapf(errors.NotFound, "Unable to get host resource for given PCIe device [%s]", device)
	}
	for _, hr := range hostResources {
		if hrStr, ok := hr.(string); ok {
			return hrStr, nil
		}
	}
	return "", errors.Wrapf(errors.NotFound, "Unable to get host resource for given PCIe device [%s]", device)
}
