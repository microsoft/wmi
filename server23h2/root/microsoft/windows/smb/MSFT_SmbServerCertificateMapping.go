// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.Microsoft.Windows.SMB
//////////////////////////////////////////////
package smb

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_SmbServerCertificateMapping struct
type MSFT_SmbServerCertificateMapping struct {
	*cim.WmiInstance

	//
	DisplayName string

	//
	Flags SmbServerCertificateMapping_Flags

	//
	MappingStatus uint32

	//
	Name string

	//
	RenewalChain string

	//
	RequireClientAuthentication bool

	//
	SkipClientCertificateAccessCheck bool

	//
	StoreName string

	//
	Subject string

	//
	Thumbprint string

	//
	Type SmbServerCertificateMapping_Type
}

func NewMSFT_SmbServerCertificateMappingEx1(instance *cim.WmiInstance) (newInstance *MSFT_SmbServerCertificateMapping, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbServerCertificateMapping{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_SmbServerCertificateMappingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_SmbServerCertificateMapping, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbServerCertificateMapping{
		WmiInstance: tmp,
	}
	return
}

// SetDisplayName sets the value of DisplayName for the instance
func (instance *MSFT_SmbServerCertificateMapping) SetPropertyDisplayName(value string) (err error) {
	return instance.SetProperty("DisplayName", (value))
}

// GetDisplayName gets the value of DisplayName for the instance
func (instance *MSFT_SmbServerCertificateMapping) GetPropertyDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DisplayName")
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

// SetFlags sets the value of Flags for the instance
func (instance *MSFT_SmbServerCertificateMapping) SetPropertyFlags(value SmbServerCertificateMapping_Flags) (err error) {
	return instance.SetProperty("Flags", (value))
}

// GetFlags gets the value of Flags for the instance
func (instance *MSFT_SmbServerCertificateMapping) GetPropertyFlags() (value SmbServerCertificateMapping_Flags, err error) {
	retValue, err := instance.GetProperty("Flags")
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

	value = SmbServerCertificateMapping_Flags(valuetmp)

	return
}

// SetMappingStatus sets the value of MappingStatus for the instance
func (instance *MSFT_SmbServerCertificateMapping) SetPropertyMappingStatus(value uint32) (err error) {
	return instance.SetProperty("MappingStatus", (value))
}

// GetMappingStatus gets the value of MappingStatus for the instance
func (instance *MSFT_SmbServerCertificateMapping) GetPropertyMappingStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("MappingStatus")
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

// SetName sets the value of Name for the instance
func (instance *MSFT_SmbServerCertificateMapping) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_SmbServerCertificateMapping) GetPropertyName() (value string, err error) {
	retValue, err := instance.GetProperty("Name")
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

// SetRenewalChain sets the value of RenewalChain for the instance
func (instance *MSFT_SmbServerCertificateMapping) SetPropertyRenewalChain(value string) (err error) {
	return instance.SetProperty("RenewalChain", (value))
}

// GetRenewalChain gets the value of RenewalChain for the instance
func (instance *MSFT_SmbServerCertificateMapping) GetPropertyRenewalChain() (value string, err error) {
	retValue, err := instance.GetProperty("RenewalChain")
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

// SetRequireClientAuthentication sets the value of RequireClientAuthentication for the instance
func (instance *MSFT_SmbServerCertificateMapping) SetPropertyRequireClientAuthentication(value bool) (err error) {
	return instance.SetProperty("RequireClientAuthentication", (value))
}

// GetRequireClientAuthentication gets the value of RequireClientAuthentication for the instance
func (instance *MSFT_SmbServerCertificateMapping) GetPropertyRequireClientAuthentication() (value bool, err error) {
	retValue, err := instance.GetProperty("RequireClientAuthentication")
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

// SetSkipClientCertificateAccessCheck sets the value of SkipClientCertificateAccessCheck for the instance
func (instance *MSFT_SmbServerCertificateMapping) SetPropertySkipClientCertificateAccessCheck(value bool) (err error) {
	return instance.SetProperty("SkipClientCertificateAccessCheck", (value))
}

// GetSkipClientCertificateAccessCheck gets the value of SkipClientCertificateAccessCheck for the instance
func (instance *MSFT_SmbServerCertificateMapping) GetPropertySkipClientCertificateAccessCheck() (value bool, err error) {
	retValue, err := instance.GetProperty("SkipClientCertificateAccessCheck")
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

// SetStoreName sets the value of StoreName for the instance
func (instance *MSFT_SmbServerCertificateMapping) SetPropertyStoreName(value string) (err error) {
	return instance.SetProperty("StoreName", (value))
}

// GetStoreName gets the value of StoreName for the instance
func (instance *MSFT_SmbServerCertificateMapping) GetPropertyStoreName() (value string, err error) {
	retValue, err := instance.GetProperty("StoreName")
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

// SetSubject sets the value of Subject for the instance
func (instance *MSFT_SmbServerCertificateMapping) SetPropertySubject(value string) (err error) {
	return instance.SetProperty("Subject", (value))
}

// GetSubject gets the value of Subject for the instance
func (instance *MSFT_SmbServerCertificateMapping) GetPropertySubject() (value string, err error) {
	retValue, err := instance.GetProperty("Subject")
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

// SetThumbprint sets the value of Thumbprint for the instance
func (instance *MSFT_SmbServerCertificateMapping) SetPropertyThumbprint(value string) (err error) {
	return instance.SetProperty("Thumbprint", (value))
}

// GetThumbprint gets the value of Thumbprint for the instance
func (instance *MSFT_SmbServerCertificateMapping) GetPropertyThumbprint() (value string, err error) {
	retValue, err := instance.GetProperty("Thumbprint")
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

// SetType sets the value of Type for the instance
func (instance *MSFT_SmbServerCertificateMapping) SetPropertyType(value SmbServerCertificateMapping_Type) (err error) {
	return instance.SetProperty("Type", (value))
}

// GetType gets the value of Type for the instance
func (instance *MSFT_SmbServerCertificateMapping) GetPropertyType() (value SmbServerCertificateMapping_Type, err error) {
	retValue, err := instance.GetProperty("Type")
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

	value = SmbServerCertificateMapping_Type(valuetmp)

	return
}

//

// <param name="DisplayName" type="string "></param>
// <param name="Flags" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="RequireClientAuthentication" type="bool "></param>
// <param name="SkipClientCertificateAccessCheck" type="bool "></param>
// <param name="StoreName" type="string "></param>
// <param name="Subject" type="string "></param>
// <param name="Thumbprint" type="string "></param>
// <param name="Type" type="uint32 "></param>

// <param name="CreatedMapping" type="MSFT_SmbServerCertificateMapping "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbServerCertificateMapping) AddServerCertificateMapping( /* IN */ Name string,
	/* IN */ Subject string,
	/* IN */ Thumbprint string,
	/* IN */ DisplayName string,
	/* IN */ StoreName string,
	/* IN */ Type uint32,
	/* IN */ Flags uint32,
	/* IN */ RequireClientAuthentication bool,
	/* IN */ SkipClientCertificateAccessCheck bool,
	/* OUT */ CreatedMapping MSFT_SmbServerCertificateMapping) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("AddServerCertificateMapping", Name, Subject, Thumbprint, DisplayName, StoreName, Type, Flags, RequireClientAuthentication, SkipClientCertificateAccessCheck)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Flags" type="uint32 "></param>
// <param name="Name" type="string "></param>
// <param name="RequireClientAuthentication" type="bool "></param>
// <param name="SkipClientCertificateAccessCheck" type="bool "></param>
// <param name="StoreName" type="string "></param>
// <param name="Thumbprint" type="string "></param>

// <param name="CreatedMapping" type="MSFT_SmbServerCertificateMapping "></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbServerCertificateMapping) SetServerCertificateMapping( /* IN */ Name string,
	/* IN */ Thumbprint string,
	/* IN */ StoreName string,
	/* IN */ Flags uint32,
	/* IN */ RequireClientAuthentication bool,
	/* IN */ SkipClientCertificateAccessCheck bool,
	/* OUT */ CreatedMapping MSFT_SmbServerCertificateMapping) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("SetServerCertificateMapping", Name, Thumbprint, StoreName, Flags, RequireClientAuthentication, SkipClientCertificateAccessCheck)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Description" type="string "></param>
// <param name="Identifier" type="string "></param>
// <param name="IdentifierType" type="uint32 "></param>

// <param name="Output" type="MSFT_SmbServerCertificateMappingAccessControlEntry []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbServerCertificateMapping) GrantClientAccessToServer( /* IN */ IdentifierType uint32,
	/* IN */ Identifier string,
	/* IN */ Description string,
	/* OUT */ Output []MSFT_SmbServerCertificateMappingAccessControlEntry) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GrantClientAccessToServer", IdentifierType, Identifier, Description)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Identifier" type="string "></param>
// <param name="IdentifierType" type="uint32 "></param>

// <param name="Output" type="MSFT_SmbServerCertificateMappingAccessControlEntry []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbServerCertificateMapping) RevokeClientAccessToServer( /* IN */ IdentifierType uint32,
	/* IN */ Identifier string,
	/* OUT */ Output []MSFT_SmbServerCertificateMappingAccessControlEntry) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("RevokeClientAccessToServer", IdentifierType, Identifier)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Description" type="string "></param>
// <param name="Identifier" type="string "></param>
// <param name="IdentifierType" type="uint32 "></param>

// <param name="Output" type="MSFT_SmbServerCertificateMappingAccessControlEntry []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbServerCertificateMapping) BlockClientAccessToServer( /* IN */ IdentifierType uint32,
	/* IN */ Identifier string,
	/* IN */ Description string,
	/* OUT */ Output []MSFT_SmbServerCertificateMappingAccessControlEntry) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("BlockClientAccessToServer", IdentifierType, Identifier, Description)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Identifier" type="string "></param>
// <param name="IdentifierType" type="uint32 "></param>

// <param name="Output" type="MSFT_SmbServerCertificateMappingAccessControlEntry []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbServerCertificateMapping) UnblockClientAccessToServer( /* IN */ IdentifierType uint32,
	/* IN */ Identifier string,
	/* OUT */ Output []MSFT_SmbServerCertificateMappingAccessControlEntry) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("UnblockClientAccessToServer", IdentifierType, Identifier)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Output" type="MSFT_SmbServerCertificateMappingAccessControlEntry []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_SmbServerCertificateMapping) GetAccessControlEntries( /* OUT */ Output []MSFT_SmbServerCertificateMappingAccessControlEntry) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("GetAccessControlEntries")
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
