// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

// MSFT_NetNatSession struct
type MSFT_NetNatSession struct {
	MSFT_NetSettingData

	//
	CreationTime string

	//
	ExternalDestinationAddress string

	//
	ExternalDestinationPort uint16

	//
	ExternalSourceAddress string

	//
	ExternalSourcePort uint16

	//
	InternalDestinationAddress string

	//
	InternalDestinationPort uint16

	//
	InternalRoutingDomainId string

	//
	InternalSourceAddress string

	//
	InternalSourcePort uint16

	//
	NatName string

	//
	Protocol uint32
}

// SetCreationTime sets the value of CreationTime for the instance
func (instance *MSFT_NetNatSession) SetPropertyCreationTime(value string) (err error) {
	return instance.SetProperty("CreationTime", value)
}

// GetCreationTime gets the value of CreationTime for the instance
func (instance *MSFT_NetNatSession) GetPropertyCreationTime() (value string, err error) {
	retValue, err := instance.GetProperty("CreationTime")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExternalDestinationAddress sets the value of ExternalDestinationAddress for the instance
func (instance *MSFT_NetNatSession) SetPropertyExternalDestinationAddress(value string) (err error) {
	return instance.SetProperty("ExternalDestinationAddress", value)
}

// GetExternalDestinationAddress gets the value of ExternalDestinationAddress for the instance
func (instance *MSFT_NetNatSession) GetPropertyExternalDestinationAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ExternalDestinationAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExternalDestinationPort sets the value of ExternalDestinationPort for the instance
func (instance *MSFT_NetNatSession) SetPropertyExternalDestinationPort(value uint16) (err error) {
	return instance.SetProperty("ExternalDestinationPort", value)
}

// GetExternalDestinationPort gets the value of ExternalDestinationPort for the instance
func (instance *MSFT_NetNatSession) GetPropertyExternalDestinationPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("ExternalDestinationPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExternalSourceAddress sets the value of ExternalSourceAddress for the instance
func (instance *MSFT_NetNatSession) SetPropertyExternalSourceAddress(value string) (err error) {
	return instance.SetProperty("ExternalSourceAddress", value)
}

// GetExternalSourceAddress gets the value of ExternalSourceAddress for the instance
func (instance *MSFT_NetNatSession) GetPropertyExternalSourceAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ExternalSourceAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetExternalSourcePort sets the value of ExternalSourcePort for the instance
func (instance *MSFT_NetNatSession) SetPropertyExternalSourcePort(value uint16) (err error) {
	return instance.SetProperty("ExternalSourcePort", value)
}

// GetExternalSourcePort gets the value of ExternalSourcePort for the instance
func (instance *MSFT_NetNatSession) GetPropertyExternalSourcePort() (value uint16, err error) {
	retValue, err := instance.GetProperty("ExternalSourcePort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternalDestinationAddress sets the value of InternalDestinationAddress for the instance
func (instance *MSFT_NetNatSession) SetPropertyInternalDestinationAddress(value string) (err error) {
	return instance.SetProperty("InternalDestinationAddress", value)
}

// GetInternalDestinationAddress gets the value of InternalDestinationAddress for the instance
func (instance *MSFT_NetNatSession) GetPropertyInternalDestinationAddress() (value string, err error) {
	retValue, err := instance.GetProperty("InternalDestinationAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternalDestinationPort sets the value of InternalDestinationPort for the instance
func (instance *MSFT_NetNatSession) SetPropertyInternalDestinationPort(value uint16) (err error) {
	return instance.SetProperty("InternalDestinationPort", value)
}

// GetInternalDestinationPort gets the value of InternalDestinationPort for the instance
func (instance *MSFT_NetNatSession) GetPropertyInternalDestinationPort() (value uint16, err error) {
	retValue, err := instance.GetProperty("InternalDestinationPort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternalRoutingDomainId sets the value of InternalRoutingDomainId for the instance
func (instance *MSFT_NetNatSession) SetPropertyInternalRoutingDomainId(value string) (err error) {
	return instance.SetProperty("InternalRoutingDomainId", value)
}

// GetInternalRoutingDomainId gets the value of InternalRoutingDomainId for the instance
func (instance *MSFT_NetNatSession) GetPropertyInternalRoutingDomainId() (value string, err error) {
	retValue, err := instance.GetProperty("InternalRoutingDomainId")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternalSourceAddress sets the value of InternalSourceAddress for the instance
func (instance *MSFT_NetNatSession) SetPropertyInternalSourceAddress(value string) (err error) {
	return instance.SetProperty("InternalSourceAddress", value)
}

// GetInternalSourceAddress gets the value of InternalSourceAddress for the instance
func (instance *MSFT_NetNatSession) GetPropertyInternalSourceAddress() (value string, err error) {
	retValue, err := instance.GetProperty("InternalSourceAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInternalSourcePort sets the value of InternalSourcePort for the instance
func (instance *MSFT_NetNatSession) SetPropertyInternalSourcePort(value uint16) (err error) {
	return instance.SetProperty("InternalSourcePort", value)
}

// GetInternalSourcePort gets the value of InternalSourcePort for the instance
func (instance *MSFT_NetNatSession) GetPropertyInternalSourcePort() (value uint16, err error) {
	retValue, err := instance.GetProperty("InternalSourcePort")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNatName sets the value of NatName for the instance
func (instance *MSFT_NetNatSession) SetPropertyNatName(value string) (err error) {
	return instance.SetProperty("NatName", value)
}

// GetNatName gets the value of NatName for the instance
func (instance *MSFT_NetNatSession) GetPropertyNatName() (value string, err error) {
	retValue, err := instance.GetProperty("NatName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetProtocol sets the value of Protocol for the instance
func (instance *MSFT_NetNatSession) SetPropertyProtocol(value uint32) (err error) {
	return instance.SetProperty("Protocol", value)
}

// GetProtocol gets the value of Protocol for the instance
func (instance *MSFT_NetNatSession) GetPropertyProtocol() (value uint32, err error) {
	retValue, err := instance.GetProperty("Protocol")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
