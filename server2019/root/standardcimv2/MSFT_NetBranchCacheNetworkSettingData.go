// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetBranchCacheNetworkSettingData struct
type MSFT_NetBranchCacheNetworkSettingData struct {
	MSFT_NetBranchCacheSettingData

	//
	ContentDownloadConnectPort uint16

	//
	ContentDownloadListenPort uint16

	//
	ContentRetrievalFirewallRulesEnabled bool

	//
	ContentRetrievalUrlReservationEnabled bool

	//
	HostedCacheClientFirewallRulesEnabled bool

	//
	HostedCacheHttpConnectPort uint16

	//
	HostedCacheHttpListenPort uint16

	//
	HostedCacheHttpsConnectPort uint16

	//
	HostedCacheHttpsListenPort uint16

	//
	HostedCacheHttpsUrlReservationEnabled bool

	//
	HostedCacheHttpUrlReservationEnabled bool

	//
	HostedCacheServerFirewallRulesEnabled bool

	//
	PeerDiscoveryFirewallRulesEnabled bool
}

// SetContentDownloadConnectPort sets the value of ContentDownloadConnectPort for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) SetPropertyContentDownloadConnectPort(value uint16) (err error) {
	return instance.SetProperty("ContentDownloadConnectPort", value)
}

// GetContentDownloadConnectPort gets the value of ContentDownloadConnectPort for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) GetPropertyContentDownloadConnectPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("ContentDownloadConnectPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetContentDownloadListenPort sets the value of ContentDownloadListenPort for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) SetPropertyContentDownloadListenPort(value uint16) (err error) {
	return instance.SetProperty("ContentDownloadListenPort", value)
}

// GetContentDownloadListenPort gets the value of ContentDownloadListenPort for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) GetPropertyContentDownloadListenPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("ContentDownloadListenPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetContentRetrievalFirewallRulesEnabled sets the value of ContentRetrievalFirewallRulesEnabled for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) SetPropertyContentRetrievalFirewallRulesEnabled(value bool) (err error) {
	return instance.SetProperty("ContentRetrievalFirewallRulesEnabled", value)
}

// GetContentRetrievalFirewallRulesEnabled gets the value of ContentRetrievalFirewallRulesEnabled for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) GetPropertyContentRetrievalFirewallRulesEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("ContentRetrievalFirewallRulesEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetContentRetrievalUrlReservationEnabled sets the value of ContentRetrievalUrlReservationEnabled for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) SetPropertyContentRetrievalUrlReservationEnabled(value bool) (err error) {
	return instance.SetProperty("ContentRetrievalUrlReservationEnabled", value)
}

// GetContentRetrievalUrlReservationEnabled gets the value of ContentRetrievalUrlReservationEnabled for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) GetPropertyContentRetrievalUrlReservationEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("ContentRetrievalUrlReservationEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostedCacheClientFirewallRulesEnabled sets the value of HostedCacheClientFirewallRulesEnabled for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) SetPropertyHostedCacheClientFirewallRulesEnabled(value bool) (err error) {
	return instance.SetProperty("HostedCacheClientFirewallRulesEnabled", value)
}

// GetHostedCacheClientFirewallRulesEnabled gets the value of HostedCacheClientFirewallRulesEnabled for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) GetPropertyHostedCacheClientFirewallRulesEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("HostedCacheClientFirewallRulesEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostedCacheHttpConnectPort sets the value of HostedCacheHttpConnectPort for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) SetPropertyHostedCacheHttpConnectPort(value uint16) (err error) {
	return instance.SetProperty("HostedCacheHttpConnectPort", value)
}

// GetHostedCacheHttpConnectPort gets the value of HostedCacheHttpConnectPort for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) GetPropertyHostedCacheHttpConnectPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("HostedCacheHttpConnectPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostedCacheHttpListenPort sets the value of HostedCacheHttpListenPort for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) SetPropertyHostedCacheHttpListenPort(value uint16) (err error) {
	return instance.SetProperty("HostedCacheHttpListenPort", value)
}

// GetHostedCacheHttpListenPort gets the value of HostedCacheHttpListenPort for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) GetPropertyHostedCacheHttpListenPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("HostedCacheHttpListenPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostedCacheHttpsConnectPort sets the value of HostedCacheHttpsConnectPort for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) SetPropertyHostedCacheHttpsConnectPort(value uint16) (err error) {
	return instance.SetProperty("HostedCacheHttpsConnectPort", value)
}

// GetHostedCacheHttpsConnectPort gets the value of HostedCacheHttpsConnectPort for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) GetPropertyHostedCacheHttpsConnectPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("HostedCacheHttpsConnectPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostedCacheHttpsListenPort sets the value of HostedCacheHttpsListenPort for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) SetPropertyHostedCacheHttpsListenPort(value uint16) (err error) {
	return instance.SetProperty("HostedCacheHttpsListenPort", value)
}

// GetHostedCacheHttpsListenPort gets the value of HostedCacheHttpsListenPort for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) GetPropertyHostedCacheHttpsListenPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("HostedCacheHttpsListenPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostedCacheHttpsUrlReservationEnabled sets the value of HostedCacheHttpsUrlReservationEnabled for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) SetPropertyHostedCacheHttpsUrlReservationEnabled(value bool) (err error) {
	return instance.SetProperty("HostedCacheHttpsUrlReservationEnabled", value)
}

// GetHostedCacheHttpsUrlReservationEnabled gets the value of HostedCacheHttpsUrlReservationEnabled for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) GetPropertyHostedCacheHttpsUrlReservationEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("HostedCacheHttpsUrlReservationEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostedCacheHttpUrlReservationEnabled sets the value of HostedCacheHttpUrlReservationEnabled for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) SetPropertyHostedCacheHttpUrlReservationEnabled(value bool) (err error) {
	return instance.SetProperty("HostedCacheHttpUrlReservationEnabled", value)
}

// GetHostedCacheHttpUrlReservationEnabled gets the value of HostedCacheHttpUrlReservationEnabled for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) GetPropertyHostedCacheHttpUrlReservationEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("HostedCacheHttpUrlReservationEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetHostedCacheServerFirewallRulesEnabled sets the value of HostedCacheServerFirewallRulesEnabled for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) SetPropertyHostedCacheServerFirewallRulesEnabled(value bool) (err error) {
	return instance.SetProperty("HostedCacheServerFirewallRulesEnabled", value)
}

// GetHostedCacheServerFirewallRulesEnabled gets the value of HostedCacheServerFirewallRulesEnabled for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) GetPropertyHostedCacheServerFirewallRulesEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("HostedCacheServerFirewallRulesEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPeerDiscoveryFirewallRulesEnabled sets the value of PeerDiscoveryFirewallRulesEnabled for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) SetPropertyPeerDiscoveryFirewallRulesEnabled(value bool) (err error) {
	return instance.SetProperty("PeerDiscoveryFirewallRulesEnabled", value)
}

// GetPeerDiscoveryFirewallRulesEnabled gets the value of PeerDiscoveryFirewallRulesEnabled for the instance
func (instance *MSFT_NetBranchCacheNetworkSettingData) GetPropertyPeerDiscoveryFirewallRulesEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("PeerDiscoveryFirewallRulesEnabled")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
