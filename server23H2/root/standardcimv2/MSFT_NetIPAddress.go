// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetIPAddress struct
type MSFT_NetIPAddress struct {
	*CIM_IPProtocolEndpoint

	//
	AddressFamily uint16

	//
	AddressState uint16

	//
	InterfaceAlias string

	//
	InterfaceIndex uint32

	//
	IPAddress string

	//
	PreferredLifetime string

	//
	PrefixOrigin uint16

	//
	SkipAsSource bool

	//
	Store uint8

	//
	SuffixOrigin uint16

	//
	Type uint8

	//
	ValidLifetime string
}

func NewMSFT_NetIPAddressEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetIPAddress, err error) {
	tmp, err := NewCIM_IPProtocolEndpointEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPAddress{
		CIM_IPProtocolEndpoint: tmp,
	}
	return
}

func NewMSFT_NetIPAddressEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetIPAddress, err error) {
	tmp, err := NewCIM_IPProtocolEndpointEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetIPAddress{
		CIM_IPProtocolEndpoint: tmp,
	}
	return
}

// SetAddressFamily sets the value of AddressFamily for the instance
func (instance *MSFT_NetIPAddress) SetPropertyAddressFamily(value uint16) (err error) {
	return instance.SetProperty("AddressFamily", (value))
}

// GetAddressFamily gets the value of AddressFamily for the instance
func (instance *MSFT_NetIPAddress) GetPropertyAddressFamily() (value uint16, err error) {
	retValue, err := instance.GetProperty("AddressFamily")
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

// SetAddressState sets the value of AddressState for the instance
func (instance *MSFT_NetIPAddress) SetPropertyAddressState(value uint16) (err error) {
	return instance.SetProperty("AddressState", (value))
}

// GetAddressState gets the value of AddressState for the instance
func (instance *MSFT_NetIPAddress) GetPropertyAddressState() (value uint16, err error) {
	retValue, err := instance.GetProperty("AddressState")
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

// SetInterfaceAlias sets the value of InterfaceAlias for the instance
func (instance *MSFT_NetIPAddress) SetPropertyInterfaceAlias(value string) (err error) {
	return instance.SetProperty("InterfaceAlias", (value))
}

// GetInterfaceAlias gets the value of InterfaceAlias for the instance
func (instance *MSFT_NetIPAddress) GetPropertyInterfaceAlias() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceAlias")
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

// SetInterfaceIndex sets the value of InterfaceIndex for the instance
func (instance *MSFT_NetIPAddress) SetPropertyInterfaceIndex(value uint32) (err error) {
	return instance.SetProperty("InterfaceIndex", (value))
}

// GetInterfaceIndex gets the value of InterfaceIndex for the instance
func (instance *MSFT_NetIPAddress) GetPropertyInterfaceIndex() (value uint32, err error) {
	retValue, err := instance.GetProperty("InterfaceIndex")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint32)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint32 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint32(valuetmp)

	return
}

// SetIPAddress sets the value of IPAddress for the instance
func (instance *MSFT_NetIPAddress) SetPropertyIPAddress(value string) (err error) {
	return instance.SetProperty("IPAddress", (value))
}

// GetIPAddress gets the value of IPAddress for the instance
func (instance *MSFT_NetIPAddress) GetPropertyIPAddress() (value string, err error) {
	retValue, err := instance.GetProperty("IPAddress")
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

// SetPreferredLifetime sets the value of PreferredLifetime for the instance
func (instance *MSFT_NetIPAddress) SetPropertyPreferredLifetime(value string) (err error) {
	return instance.SetProperty("PreferredLifetime", (value))
}

// GetPreferredLifetime gets the value of PreferredLifetime for the instance
func (instance *MSFT_NetIPAddress) GetPropertyPreferredLifetime() (value string, err error) {
	retValue, err := instance.GetProperty("PreferredLifetime")
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

// SetPrefixOrigin sets the value of PrefixOrigin for the instance
func (instance *MSFT_NetIPAddress) SetPropertyPrefixOrigin(value uint16) (err error) {
	return instance.SetProperty("PrefixOrigin", (value))
}

// GetPrefixOrigin gets the value of PrefixOrigin for the instance
func (instance *MSFT_NetIPAddress) GetPropertyPrefixOrigin() (value uint16, err error) {
	retValue, err := instance.GetProperty("PrefixOrigin")
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

// SetSkipAsSource sets the value of SkipAsSource for the instance
func (instance *MSFT_NetIPAddress) SetPropertySkipAsSource(value bool) (err error) {
	return instance.SetProperty("SkipAsSource", (value))
}

// GetSkipAsSource gets the value of SkipAsSource for the instance
func (instance *MSFT_NetIPAddress) GetPropertySkipAsSource() (value bool, err error) {
	retValue, err := instance.GetProperty("SkipAsSource")
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

// SetStore sets the value of Store for the instance
func (instance *MSFT_NetIPAddress) SetPropertyStore(value uint8) (err error) {
	return instance.SetProperty("Store", (value))
}

// GetStore gets the value of Store for the instance
func (instance *MSFT_NetIPAddress) GetPropertyStore() (value uint8, err error) {
	retValue, err := instance.GetProperty("Store")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetSuffixOrigin sets the value of SuffixOrigin for the instance
func (instance *MSFT_NetIPAddress) SetPropertySuffixOrigin(value uint16) (err error) {
	return instance.SetProperty("SuffixOrigin", (value))
}

// GetSuffixOrigin gets the value of SuffixOrigin for the instance
func (instance *MSFT_NetIPAddress) GetPropertySuffixOrigin() (value uint16, err error) {
	retValue, err := instance.GetProperty("SuffixOrigin")
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

// SetType sets the value of Type for the instance
func (instance *MSFT_NetIPAddress) SetPropertyType(value uint8) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSFT_NetIPAddress) GetPropertyType() (value uint8, err error) {
	retValue, err := instance.GetProperty("Type")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(uint8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = uint8(valuetmp)

	return
}

// SetValidLifetime sets the value of ValidLifetime for the instance
func (instance *MSFT_NetIPAddress) SetPropertyValidLifetime(value string) (err error) {
	return instance.SetProperty("ValidLifetime", (value))
}

// GetValidLifetime gets the value of ValidLifetime for the instance
func (instance *MSFT_NetIPAddress) GetPropertyValidLifetime() (value string, err error) {
	retValue, err := instance.GetProperty("ValidLifetime")
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

//

// <param name="AddressFamily" type="uint16 "></param>
// <param name="AddressState" type="uint16 "></param>
// <param name="DefaultGateway" type="string "></param>
// <param name="InterfaceAlias" type="string "></param>
// <param name="InterfaceIndex" type="uint32 "></param>
// <param name="IPAddress" type="string "></param>
// <param name="PassThru" type="bool "></param>
// <param name="PolicyStore" type="string "></param>
// <param name="PreferredLifetime" type="string "></param>
// <param name="PrefixLength" type="uint8 "></param>
// <param name="PrefixOrigin" type="uint16 "></param>
// <param name="SkipAsSource" type="bool "></param>
// <param name="SuffixOrigin" type="uint16 "></param>
// <param name="Type" type="uint8 "></param>
// <param name="ValidLifetime" type="string "></param>

// <param name="CmdletOutput" type="MSFT_NetIPAddress []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_NetIPAddress) Create( /* IN */ InterfaceIndex uint32,
	/* IN */ InterfaceAlias string,
	/* IN */ IPAddress string,
	/* IN */ AddressFamily uint16,
	/* IN */ PrefixLength uint8,
	/* IN */ Type uint8,
	/* IN */ PrefixOrigin uint16,
	/* IN */ SuffixOrigin uint16,
	/* IN */ AddressState uint16,
	/* IN */ ValidLifetime string,
	/* IN */ PreferredLifetime string,
	/* IN */ SkipAsSource bool,
	/* IN */ DefaultGateway string,
	/* IN */ PolicyStore string,
	/* IN */ PassThru bool,
	/* OUT */ CmdletOutput []MSFT_NetIPAddress) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("Create", InterfaceIndex, InterfaceAlias, IPAddress, AddressFamily, PrefixLength, Type, PrefixOrigin, SuffixOrigin, AddressState, ValidLifetime, PreferredLifetime, SkipAsSource, DefaultGateway, PolicyStore, PassThru)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
