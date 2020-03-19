// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// CIM_NetworkPort struct
type CIM_NetworkPort struct {
	*CIM_LogicalPort

	//
	ActiveMaximumTransmissionUnit uint64

	//
	AutoSense bool

	//
	FullDuplex bool

	//
	LinkTechnology uint16

	//
	NetworkAddresses []string

	//
	OtherLinkTechnology string

	//
	OtherNetworkPortType string

	//
	PermanentAddress string

	//
	PortNumber uint16

	//
	SupportedMaximumTransmissionUnit uint64
}

func NewCIM_NetworkPortEx1(instance *cim.WmiInstance) (newInstance *CIM_NetworkPort, err error) {
	tmp, err := NewCIM_LogicalPortEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_NetworkPort{
		CIM_LogicalPort: tmp,
	}
	return
}

func NewCIM_NetworkPortEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_NetworkPort, err error) {
	tmp, err := NewCIM_LogicalPortEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_NetworkPort{
		CIM_LogicalPort: tmp,
	}
	return
}

// SetActiveMaximumTransmissionUnit sets the value of ActiveMaximumTransmissionUnit for the instance
func (instance *CIM_NetworkPort) SetPropertyActiveMaximumTransmissionUnit(value uint64) (err error) {
	return instance.SetProperty("ActiveMaximumTransmissionUnit", value)
}

// GetActiveMaximumTransmissionUnit gets the value of ActiveMaximumTransmissionUnit for the instance
func (instance *CIM_NetworkPort) GetPropertyActiveMaximumTransmissionUnit() (value uint64, err error) {
	retValue, err := instance.GetProperty("ActiveMaximumTransmissionUnit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetAutoSense sets the value of AutoSense for the instance
func (instance *CIM_NetworkPort) SetPropertyAutoSense(value bool) (err error) {
	return instance.SetProperty("AutoSense", value)
}

// GetAutoSense gets the value of AutoSense for the instance
func (instance *CIM_NetworkPort) GetPropertyAutoSense() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoSense")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetFullDuplex sets the value of FullDuplex for the instance
func (instance *CIM_NetworkPort) SetPropertyFullDuplex(value bool) (err error) {
	return instance.SetProperty("FullDuplex", value)
}

// GetFullDuplex gets the value of FullDuplex for the instance
func (instance *CIM_NetworkPort) GetPropertyFullDuplex() (value bool, err error) {
	retValue, err := instance.GetProperty("FullDuplex")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetLinkTechnology sets the value of LinkTechnology for the instance
func (instance *CIM_NetworkPort) SetPropertyLinkTechnology(value uint16) (err error) {
	return instance.SetProperty("LinkTechnology", value)
}

// GetLinkTechnology gets the value of LinkTechnology for the instance
func (instance *CIM_NetworkPort) GetPropertyLinkTechnology() (value uint16, err error) {
	retValue, err := instance.GetProperty("LinkTechnology")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetNetworkAddresses sets the value of NetworkAddresses for the instance
func (instance *CIM_NetworkPort) SetPropertyNetworkAddresses(value []string) (err error) {
	return instance.SetProperty("NetworkAddresses", value)
}

// GetNetworkAddresses gets the value of NetworkAddresses for the instance
func (instance *CIM_NetworkPort) GetPropertyNetworkAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("NetworkAddresses")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherLinkTechnology sets the value of OtherLinkTechnology for the instance
func (instance *CIM_NetworkPort) SetPropertyOtherLinkTechnology(value string) (err error) {
	return instance.SetProperty("OtherLinkTechnology", value)
}

// GetOtherLinkTechnology gets the value of OtherLinkTechnology for the instance
func (instance *CIM_NetworkPort) GetPropertyOtherLinkTechnology() (value string, err error) {
	retValue, err := instance.GetProperty("OtherLinkTechnology")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetOtherNetworkPortType sets the value of OtherNetworkPortType for the instance
func (instance *CIM_NetworkPort) SetPropertyOtherNetworkPortType(value string) (err error) {
	return instance.SetProperty("OtherNetworkPortType", value)
}

// GetOtherNetworkPortType gets the value of OtherNetworkPortType for the instance
func (instance *CIM_NetworkPort) GetPropertyOtherNetworkPortType() (value string, err error) {
	retValue, err := instance.GetProperty("OtherNetworkPortType")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPermanentAddress sets the value of PermanentAddress for the instance
func (instance *CIM_NetworkPort) SetPropertyPermanentAddress(value string) (err error) {
	return instance.SetProperty("PermanentAddress", value)
}

// GetPermanentAddress gets the value of PermanentAddress for the instance
func (instance *CIM_NetworkPort) GetPropertyPermanentAddress() (value string, err error) {
	retValue, err := instance.GetProperty("PermanentAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPortNumber sets the value of PortNumber for the instance
func (instance *CIM_NetworkPort) SetPropertyPortNumber(value uint16) (err error) {
	return instance.SetProperty("PortNumber", value)
}

// GetPortNumber gets the value of PortNumber for the instance
func (instance *CIM_NetworkPort) GetPropertyPortNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("PortNumber")
	if err != nil {
		return
	}
	value, ok := retValue.(uint16)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSupportedMaximumTransmissionUnit sets the value of SupportedMaximumTransmissionUnit for the instance
func (instance *CIM_NetworkPort) SetPropertySupportedMaximumTransmissionUnit(value uint64) (err error) {
	return instance.SetProperty("SupportedMaximumTransmissionUnit", value)
}

// GetSupportedMaximumTransmissionUnit gets the value of SupportedMaximumTransmissionUnit for the instance
func (instance *CIM_NetworkPort) GetPropertySupportedMaximumTransmissionUnit() (value uint64, err error) {
	retValue, err := instance.GetProperty("SupportedMaximumTransmissionUnit")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}
