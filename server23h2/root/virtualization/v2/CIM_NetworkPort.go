// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.virtualization.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_NetworkPort struct
type CIM_NetworkPort struct {
	*CIM_LogicalPort

	// The active or negotiated maximum transmission unit (MTU) that can be supported.
	ActiveMaximumTransmissionUnit uint64

	// A Boolean that indicates whether the NetworkPort is capable of automatically determining the speed or other communications characteristics of the attached network media.
	AutoSense bool

	// Boolean that indicates that the port is operating in full duplex mode.
	FullDuplex bool

	// An enumeration of the types of links. When set to 1 ("Other"), the related property OtherLinkTechnology contains a string description of the type of link.
	LinkTechnology NetworkPort_LinkTechnology

	// An array of strings that indicates the network addresses for the port.
	NetworkAddresses []string

	// A string value that describes LinkTechnology when it is set to 1, "Other".
	OtherLinkTechnology string

	// Note: The use of this property is deprecated in lieu of CIM_LogicalPort.PortType.
	///Deprecated description: The type of module, when PortType is set to 1 ("Other".)
	OtherNetworkPortType string

	// PermanentAddress defines the network address that is hardcoded into a port. This 'hardcoded' address can be changed using a firmware upgrade or a software configuration. When this change is made, the field should be updated at the same time. PermanentAddress should be left blank if no 'hardcoded' address exists for the NetworkAdapter.
	PermanentAddress string

	// NetworkPorts are often numbered relative to either a logical module or a network element.
	PortNumber uint16

	// The maximum transmission unit (MTU) that can be supported.
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
	return instance.SetProperty("ActiveMaximumTransmissionUnit", (value))
}

// GetActiveMaximumTransmissionUnit gets the value of ActiveMaximumTransmissionUnit for the instance
func (instance *CIM_NetworkPort) GetPropertyActiveMaximumTransmissionUnit() (value uint64, err error) {
	retValue, err := instance.GetProperty("ActiveMaximumTransmissionUnit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}

// SetAutoSense sets the value of AutoSense for the instance
func (instance *CIM_NetworkPort) SetPropertyAutoSense(value bool) (err error) {
	return instance.SetProperty("AutoSense", (value))
}

// GetAutoSense gets the value of AutoSense for the instance
func (instance *CIM_NetworkPort) GetPropertyAutoSense() (value bool, err error) {
	retValue, err := instance.GetProperty("AutoSense")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetFullDuplex sets the value of FullDuplex for the instance
func (instance *CIM_NetworkPort) SetPropertyFullDuplex(value bool) (err error) {
	return instance.SetProperty("FullDuplex", (value))
}

// GetFullDuplex gets the value of FullDuplex for the instance
func (instance *CIM_NetworkPort) GetPropertyFullDuplex() (value bool, err error) {
	retValue, err := instance.GetProperty("FullDuplex")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(bool)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = bool(valuetmp)

	return
}

// SetLinkTechnology sets the value of LinkTechnology for the instance
func (instance *CIM_NetworkPort) SetPropertyLinkTechnology(value NetworkPort_LinkTechnology) (err error) {
	return instance.SetProperty("LinkTechnology", (value))
}

// GetLinkTechnology gets the value of LinkTechnology for the instance
func (instance *CIM_NetworkPort) GetPropertyLinkTechnology() (value NetworkPort_LinkTechnology, err error) {
	retValue, err := instance.GetProperty("LinkTechnology")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = NetworkPort_LinkTechnology(valuetmp)

	return
}

// SetNetworkAddresses sets the value of NetworkAddresses for the instance
func (instance *CIM_NetworkPort) SetPropertyNetworkAddresses(value []string) (err error) {
	return instance.SetProperty("NetworkAddresses", (value))
}

// GetNetworkAddresses gets the value of NetworkAddresses for the instance
func (instance *CIM_NetworkPort) GetPropertyNetworkAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("NetworkAddresses")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(string)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, string(valuetmp))
	}

	return
}

// SetOtherLinkTechnology sets the value of OtherLinkTechnology for the instance
func (instance *CIM_NetworkPort) SetPropertyOtherLinkTechnology(value string) (err error) {
	return instance.SetProperty("OtherLinkTechnology", (value))
}

// GetOtherLinkTechnology gets the value of OtherLinkTechnology for the instance
func (instance *CIM_NetworkPort) GetPropertyOtherLinkTechnology() (value string, err error) {
	retValue, err := instance.GetProperty("OtherLinkTechnology")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetOtherNetworkPortType sets the value of OtherNetworkPortType for the instance
func (instance *CIM_NetworkPort) SetPropertyOtherNetworkPortType(value string) (err error) {
	return instance.SetProperty("OtherNetworkPortType", (value))
}

// GetOtherNetworkPortType gets the value of OtherNetworkPortType for the instance
func (instance *CIM_NetworkPort) GetPropertyOtherNetworkPortType() (value string, err error) {
	retValue, err := instance.GetProperty("OtherNetworkPortType")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetPermanentAddress sets the value of PermanentAddress for the instance
func (instance *CIM_NetworkPort) SetPropertyPermanentAddress(value string) (err error) {
	return instance.SetProperty("PermanentAddress", (value))
}

// GetPermanentAddress gets the value of PermanentAddress for the instance
func (instance *CIM_NetworkPort) GetPropertyPermanentAddress() (value string, err error) {
	retValue, err := instance.GetProperty("PermanentAddress")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(string)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " string is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = string(valuetmp)

	return
}

// SetPortNumber sets the value of PortNumber for the instance
func (instance *CIM_NetworkPort) SetPropertyPortNumber(value uint16) (err error) {
	return instance.SetProperty("PortNumber", (value))
}

// GetPortNumber gets the value of PortNumber for the instance
func (instance *CIM_NetworkPort) GetPropertyPortNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("PortNumber")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint16)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint16(valuetmp)

	return
}

// SetSupportedMaximumTransmissionUnit sets the value of SupportedMaximumTransmissionUnit for the instance
func (instance *CIM_NetworkPort) SetPropertySupportedMaximumTransmissionUnit(value uint64) (err error) {
	return instance.SetProperty("SupportedMaximumTransmissionUnit", (value))
}

// GetSupportedMaximumTransmissionUnit gets the value of SupportedMaximumTransmissionUnit for the instance
func (instance *CIM_NetworkPort) GetPropertySupportedMaximumTransmissionUnit() (value uint64, err error) {
	retValue, err := instance.GetProperty("SupportedMaximumTransmissionUnit")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint64(valuetmp)

	return
}
