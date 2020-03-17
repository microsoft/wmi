// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_NetworkProtocol struct
type Win32_NetworkProtocol struct {
	CIM_LogicalElement

	//
	ConnectionlessService bool

	//
	GuaranteesDelivery bool

	//
	GuaranteesSequencing bool

	//
	MaximumAddressSize uint32

	//
	MaximumMessageSize uint32

	//
	MessageOriented bool

	//
	MinimumAddressSize uint32

	//
	PseudoStreamOriented bool

	//
	SupportsBroadcasting bool

	//
	SupportsConnectData bool

	//
	SupportsDisconnectData bool

	//
	SupportsEncryption bool

	//
	SupportsExpeditedData bool

	//
	SupportsFragmentation bool

	//
	SupportsGracefulClosing bool

	//
	SupportsGuaranteedBandwidth bool

	//
	SupportsMulticasting bool

	//
	SupportsQualityofService bool
}

// SetConnectionlessService sets the value of ConnectionlessService for the instance
func (instance *Win32_NetworkProtocol) SetPropertyConnectionlessService(value bool) (err error) {
	return instance.SetProperty("ConnectionlessService", value)
}

// GetConnectionlessService gets the value of ConnectionlessService for the instance
func (instance *Win32_NetworkProtocol) GetPropertyConnectionlessService() (value bool, err error) {
	retValue, err := instance.GetProperty("ConnectionlessService")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGuaranteesDelivery sets the value of GuaranteesDelivery for the instance
func (instance *Win32_NetworkProtocol) SetPropertyGuaranteesDelivery(value bool) (err error) {
	return instance.SetProperty("GuaranteesDelivery", value)
}

// GetGuaranteesDelivery gets the value of GuaranteesDelivery for the instance
func (instance *Win32_NetworkProtocol) GetPropertyGuaranteesDelivery() (value bool, err error) {
	retValue, err := instance.GetProperty("GuaranteesDelivery")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetGuaranteesSequencing sets the value of GuaranteesSequencing for the instance
func (instance *Win32_NetworkProtocol) SetPropertyGuaranteesSequencing(value bool) (err error) {
	return instance.SetProperty("GuaranteesSequencing", value)
}

// GetGuaranteesSequencing gets the value of GuaranteesSequencing for the instance
func (instance *Win32_NetworkProtocol) GetPropertyGuaranteesSequencing() (value bool, err error) {
	retValue, err := instance.GetProperty("GuaranteesSequencing")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumAddressSize sets the value of MaximumAddressSize for the instance
func (instance *Win32_NetworkProtocol) SetPropertyMaximumAddressSize(value uint32) (err error) {
	return instance.SetProperty("MaximumAddressSize", value)
}

// GetMaximumAddressSize gets the value of MaximumAddressSize for the instance
func (instance *Win32_NetworkProtocol) GetPropertyMaximumAddressSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumAddressSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMaximumMessageSize sets the value of MaximumMessageSize for the instance
func (instance *Win32_NetworkProtocol) SetPropertyMaximumMessageSize(value uint32) (err error) {
	return instance.SetProperty("MaximumMessageSize", value)
}

// GetMaximumMessageSize gets the value of MaximumMessageSize for the instance
func (instance *Win32_NetworkProtocol) GetPropertyMaximumMessageSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MaximumMessageSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMessageOriented sets the value of MessageOriented for the instance
func (instance *Win32_NetworkProtocol) SetPropertyMessageOriented(value bool) (err error) {
	return instance.SetProperty("MessageOriented", value)
}

// GetMessageOriented gets the value of MessageOriented for the instance
func (instance *Win32_NetworkProtocol) GetPropertyMessageOriented() (value bool, err error) {
	retValue, err := instance.GetProperty("MessageOriented")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetMinimumAddressSize sets the value of MinimumAddressSize for the instance
func (instance *Win32_NetworkProtocol) SetPropertyMinimumAddressSize(value uint32) (err error) {
	return instance.SetProperty("MinimumAddressSize", value)
}

// GetMinimumAddressSize gets the value of MinimumAddressSize for the instance
func (instance *Win32_NetworkProtocol) GetPropertyMinimumAddressSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("MinimumAddressSize")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPseudoStreamOriented sets the value of PseudoStreamOriented for the instance
func (instance *Win32_NetworkProtocol) SetPropertyPseudoStreamOriented(value bool) (err error) {
	return instance.SetProperty("PseudoStreamOriented", value)
}

// GetPseudoStreamOriented gets the value of PseudoStreamOriented for the instance
func (instance *Win32_NetworkProtocol) GetPropertyPseudoStreamOriented() (value bool, err error) {
	retValue, err := instance.GetProperty("PseudoStreamOriented")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsBroadcasting sets the value of SupportsBroadcasting for the instance
func (instance *Win32_NetworkProtocol) SetPropertySupportsBroadcasting(value bool) (err error) {
	return instance.SetProperty("SupportsBroadcasting", value)
}

// GetSupportsBroadcasting gets the value of SupportsBroadcasting for the instance
func (instance *Win32_NetworkProtocol) GetPropertySupportsBroadcasting() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsBroadcasting")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsConnectData sets the value of SupportsConnectData for the instance
func (instance *Win32_NetworkProtocol) SetPropertySupportsConnectData(value bool) (err error) {
	return instance.SetProperty("SupportsConnectData", value)
}

// GetSupportsConnectData gets the value of SupportsConnectData for the instance
func (instance *Win32_NetworkProtocol) GetPropertySupportsConnectData() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsConnectData")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsDisconnectData sets the value of SupportsDisconnectData for the instance
func (instance *Win32_NetworkProtocol) SetPropertySupportsDisconnectData(value bool) (err error) {
	return instance.SetProperty("SupportsDisconnectData", value)
}

// GetSupportsDisconnectData gets the value of SupportsDisconnectData for the instance
func (instance *Win32_NetworkProtocol) GetPropertySupportsDisconnectData() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsDisconnectData")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsEncryption sets the value of SupportsEncryption for the instance
func (instance *Win32_NetworkProtocol) SetPropertySupportsEncryption(value bool) (err error) {
	return instance.SetProperty("SupportsEncryption", value)
}

// GetSupportsEncryption gets the value of SupportsEncryption for the instance
func (instance *Win32_NetworkProtocol) GetPropertySupportsEncryption() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsEncryption")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsExpeditedData sets the value of SupportsExpeditedData for the instance
func (instance *Win32_NetworkProtocol) SetPropertySupportsExpeditedData(value bool) (err error) {
	return instance.SetProperty("SupportsExpeditedData", value)
}

// GetSupportsExpeditedData gets the value of SupportsExpeditedData for the instance
func (instance *Win32_NetworkProtocol) GetPropertySupportsExpeditedData() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsExpeditedData")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsFragmentation sets the value of SupportsFragmentation for the instance
func (instance *Win32_NetworkProtocol) SetPropertySupportsFragmentation(value bool) (err error) {
	return instance.SetProperty("SupportsFragmentation", value)
}

// GetSupportsFragmentation gets the value of SupportsFragmentation for the instance
func (instance *Win32_NetworkProtocol) GetPropertySupportsFragmentation() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsFragmentation")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsGracefulClosing sets the value of SupportsGracefulClosing for the instance
func (instance *Win32_NetworkProtocol) SetPropertySupportsGracefulClosing(value bool) (err error) {
	return instance.SetProperty("SupportsGracefulClosing", value)
}

// GetSupportsGracefulClosing gets the value of SupportsGracefulClosing for the instance
func (instance *Win32_NetworkProtocol) GetPropertySupportsGracefulClosing() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsGracefulClosing")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsGuaranteedBandwidth sets the value of SupportsGuaranteedBandwidth for the instance
func (instance *Win32_NetworkProtocol) SetPropertySupportsGuaranteedBandwidth(value bool) (err error) {
	return instance.SetProperty("SupportsGuaranteedBandwidth", value)
}

// GetSupportsGuaranteedBandwidth gets the value of SupportsGuaranteedBandwidth for the instance
func (instance *Win32_NetworkProtocol) GetPropertySupportsGuaranteedBandwidth() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsGuaranteedBandwidth")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsMulticasting sets the value of SupportsMulticasting for the instance
func (instance *Win32_NetworkProtocol) SetPropertySupportsMulticasting(value bool) (err error) {
	return instance.SetProperty("SupportsMulticasting", value)
}

// GetSupportsMulticasting gets the value of SupportsMulticasting for the instance
func (instance *Win32_NetworkProtocol) GetPropertySupportsMulticasting() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsMulticasting")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportsQualityofService sets the value of SupportsQualityofService for the instance
func (instance *Win32_NetworkProtocol) SetPropertySupportsQualityofService(value bool) (err error) {
	return instance.SetProperty("SupportsQualityofService", value)
}

// GetSupportsQualityofService gets the value of SupportsQualityofService for the instance
func (instance *Win32_NetworkProtocol) GetPropertySupportsQualityofService() (value bool, err error) {
	retValue, err := instance.GetProperty("SupportsQualityofService")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}
