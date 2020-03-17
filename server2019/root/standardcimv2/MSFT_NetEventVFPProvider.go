// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetEventVFPProvider struct
type MSFT_NetEventVFPProvider struct {
	MSFT_NetEventProviderBase

	//
	DestinationIPAddresses []string

	//
	DestinationMACAddresses []string

	//
	GREKeys []uint32

	//
	IPProtocols []uint8

	//
	PortIds []uint32

	//
	SourceIPAddresses []string

	//
	SourceMACAddresses []string

	//
	SwitchName string

	//
	TCPPorts []uint16

	//
	TenantIds []uint32

	//
	UDPPorts []uint16

	//
	VFPFlowDirection uint32

	//
	VLANIds []uint16
}

// SetDestinationIPAddresses sets the value of DestinationIPAddresses for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertyDestinationIPAddresses(value []string) (err error) {
	return instance.SetProperty("DestinationIPAddresses", value)
}

// GetDestinationIPAddresses gets the value of DestinationIPAddresses for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertyDestinationIPAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("DestinationIPAddresses")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDestinationMACAddresses sets the value of DestinationMACAddresses for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertyDestinationMACAddresses(value []string) (err error) {
	return instance.SetProperty("DestinationMACAddresses", value)
}

// GetDestinationMACAddresses gets the value of DestinationMACAddresses for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertyDestinationMACAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("DestinationMACAddresses")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGREKeys sets the value of GREKeys for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertyGREKeys(value []uint32) (err error) {
	return instance.SetProperty("GREKeys", value)
}

// GetGREKeys gets the value of GREKeys for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertyGREKeys() (value []uint32, err error) {
	retValue, err := instance.GetProperty("GREKeys")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIPProtocols sets the value of IPProtocols for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertyIPProtocols(value []uint8) (err error) {
	return instance.SetProperty("IPProtocols", value)
}

// GetIPProtocols gets the value of IPProtocols for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertyIPProtocols() (value []uint8, err error) {
	retValue, err := instance.GetProperty("IPProtocols")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPortIds sets the value of PortIds for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertyPortIds(value []uint32) (err error) {
	return instance.SetProperty("PortIds", value)
}

// GetPortIds gets the value of PortIds for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertyPortIds() (value []uint32, err error) {
	retValue, err := instance.GetProperty("PortIds")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceIPAddresses sets the value of SourceIPAddresses for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertySourceIPAddresses(value []string) (err error) {
	return instance.SetProperty("SourceIPAddresses", value)
}

// GetSourceIPAddresses gets the value of SourceIPAddresses for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertySourceIPAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("SourceIPAddresses")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSourceMACAddresses sets the value of SourceMACAddresses for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertySourceMACAddresses(value []string) (err error) {
	return instance.SetProperty("SourceMACAddresses", value)
}

// GetSourceMACAddresses gets the value of SourceMACAddresses for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertySourceMACAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("SourceMACAddresses")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSwitchName sets the value of SwitchName for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertySwitchName(value string) (err error) {
	return instance.SetProperty("SwitchName", value)
}

// GetSwitchName gets the value of SwitchName for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertySwitchName() (value string, err error) {
	retValue, err := instance.GetProperty("SwitchName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTCPPorts sets the value of TCPPorts for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertyTCPPorts(value []uint16) (err error) {
	return instance.SetProperty("TCPPorts", value)
}

// GetTCPPorts gets the value of TCPPorts for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertyTCPPorts() (value []uint16, err error) {
	retValue, err := instance.GetProperty("TCPPorts")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTenantIds sets the value of TenantIds for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertyTenantIds(value []uint32) (err error) {
	return instance.SetProperty("TenantIds", value)
}

// GetTenantIds gets the value of TenantIds for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertyTenantIds() (value []uint32, err error) {
	retValue, err := instance.GetProperty("TenantIds")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUDPPorts sets the value of UDPPorts for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertyUDPPorts(value []uint16) (err error) {
	return instance.SetProperty("UDPPorts", value)
}

// GetUDPPorts gets the value of UDPPorts for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertyUDPPorts() (value []uint16, err error) {
	retValue, err := instance.GetProperty("UDPPorts")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVFPFlowDirection sets the value of VFPFlowDirection for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertyVFPFlowDirection(value uint32) (err error) {
	return instance.SetProperty("VFPFlowDirection", value)
}

// GetVFPFlowDirection gets the value of VFPFlowDirection for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertyVFPFlowDirection() (value uint32, err error) {
	retValue, err := instance.GetProperty("VFPFlowDirection")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetVLANIds sets the value of VLANIds for the instance
func (instance *MSFT_NetEventVFPProvider) SetPropertyVLANIds(value []uint16) (err error) {
	return instance.SetProperty("VLANIds", value)
}

// GetVLANIds gets the value of VLANIds for the instance
func (instance *MSFT_NetEventVFPProvider) GetPropertyVLANIds() (value []uint16, err error) {
	retValue, err := instance.GetProperty("VLANIds")
	if err != nil {
		return
	}
	value, ok := retValue.([]uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}
