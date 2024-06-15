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

// MSFT_SmbServerCertificateMappingAccessControlEntry struct
type MSFT_SmbServerCertificateMappingAccessControlEntry struct {
	*cim.WmiInstance

	//
	AccessControlType SmbServerCertificateMappingAccessControlEntry_AccessControlType

	//
	Description string

	//
	Identifier string

	//
	IdentifierType SmbServerCertificateMappingAccessControlEntry_IdentifierType

	//
	Name string
}

func NewMSFT_SmbServerCertificateMappingAccessControlEntryEx1(instance *cim.WmiInstance) (newInstance *MSFT_SmbServerCertificateMappingAccessControlEntry, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbServerCertificateMappingAccessControlEntry{
		WmiInstance: tmp,
	}
	return
}

func NewMSFT_SmbServerCertificateMappingAccessControlEntryEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_SmbServerCertificateMappingAccessControlEntry, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_SmbServerCertificateMappingAccessControlEntry{
		WmiInstance: tmp,
	}
	return
}

// SetAccessControlType sets the value of AccessControlType for the instance
func (instance *MSFT_SmbServerCertificateMappingAccessControlEntry) SetPropertyAccessControlType(value SmbServerCertificateMappingAccessControlEntry_AccessControlType) (err error) {
	return instance.SetProperty("AccessControlType", (value))
}

// GetAccessControlType gets the value of AccessControlType for the instance
func (instance *MSFT_SmbServerCertificateMappingAccessControlEntry) GetPropertyAccessControlType() (value SmbServerCertificateMappingAccessControlEntry_AccessControlType, err error) {
	retValue, err := instance.GetProperty("AccessControlType")
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

	value = SmbServerCertificateMappingAccessControlEntry_AccessControlType(valuetmp)

	return
}

// SetDescription sets the value of Description for the instance
func (instance *MSFT_SmbServerCertificateMappingAccessControlEntry) SetPropertyDescription(value string) (err error) {
	return instance.SetProperty("Description", (value))
}

// GetDescription gets the value of Description for the instance
func (instance *MSFT_SmbServerCertificateMappingAccessControlEntry) GetPropertyDescription() (value string, err error) {
	retValue, err := instance.GetProperty("Description")
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

// SetIdentifier sets the value of Identifier for the instance
func (instance *MSFT_SmbServerCertificateMappingAccessControlEntry) SetPropertyIdentifier(value string) (err error) {
	return instance.SetProperty("Identifier", (value))
}

// GetIdentifier gets the value of Identifier for the instance
func (instance *MSFT_SmbServerCertificateMappingAccessControlEntry) GetPropertyIdentifier() (value string, err error) {
	retValue, err := instance.GetProperty("Identifier")
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

// SetIdentifierType sets the value of IdentifierType for the instance
func (instance *MSFT_SmbServerCertificateMappingAccessControlEntry) SetPropertyIdentifierType(value SmbServerCertificateMappingAccessControlEntry_IdentifierType) (err error) {
	return instance.SetProperty("IdentifierType", (value))
}

// GetIdentifierType gets the value of IdentifierType for the instance
func (instance *MSFT_SmbServerCertificateMappingAccessControlEntry) GetPropertyIdentifierType() (value SmbServerCertificateMappingAccessControlEntry_IdentifierType, err error) {
	retValue, err := instance.GetProperty("IdentifierType")
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

	value = SmbServerCertificateMappingAccessControlEntry_IdentifierType(valuetmp)

	return
}

// SetName sets the value of Name for the instance
func (instance *MSFT_SmbServerCertificateMappingAccessControlEntry) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_SmbServerCertificateMappingAccessControlEntry) GetPropertyName() (value string, err error) {
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
