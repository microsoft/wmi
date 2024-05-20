// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package virtualnetworkadapter

import (
	"fmt"

	"github.com/google/uuid"

	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type SyntheticNetworkAdapter struct {
	*v2.Msvm_SyntheticEthernetPortSettingData
}

// NewSyntheticNetworkAdapter
func NewSyntheticNetworkAdapter(instance *wmi.WmiInstance) (*SyntheticNetworkAdapter, error) {
	wmisna, err := v2.NewMsvm_SyntheticEthernetPortSettingDataEx1(instance)
	if err != nil {
		return nil, err
	}
	return &SyntheticNetworkAdapter{wmisna}, nil
}

func (sna *SyntheticNetworkAdapter) InitializeIdentifier() (err error) {
	g, err := uuid.NewUUID()
	if err != nil {
		return
	}

	vsid := fmt.Sprintf("{%s}", g.String())

	return sna.SetPropertyVirtualSystemIdentifiers([]string{vsid})
}

func (sna *SyntheticNetworkAdapter) SetMACAddress(mac string) (err error) {
	err = sna.SetPropertyStaticMacAddress(true)
	if err != nil {
		return
	}
	return sna.SetPropertyAddress(mac)
}

func (sna *SyntheticNetworkAdapter) Rename(newName string) (err error) {
	return sna.SetPropertyElementName(newName)
}

func (sna *SyntheticNetworkAdapter) GetGuestNetworkAdapterConfiguration() (guestConfiguration *GuestNetworkAdapterConfiguration, err error) {
	wmiGuestConfig, err := sna.GetRelated("Msvm_GuestNetworkAdapterConfiguration")
	if err != nil {
		return
	}
	guestConfiguration, err = NewGuestNetworkAdapterConfiguration(wmiGuestConfig)
	return guestConfiguration, err
}
