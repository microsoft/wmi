// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.HyperVCluster.v2
//////////////////////////////////////////////
package v2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_System struct
type CIM_System struct {
	*CIM_EnabledLogicalElement

	// CreationClassName indicates the name of the class or the subclass used in the creation of an instance. When used with the other key properties of this class, this property allows all instances of this class and its subclasses to be uniquely identified.
	CreationClassName string

	// An array of free-form strings providing explanations and details behind the entries in the OtherIdentifying Info array. Note, each entry of this array is related to the entry in OtherIdentifyingInfo that is located at the same index.
	IdentifyingDescriptions []string

	// The System object and its derivatives are top-level objects of CIM. They provide the scope for numerous components. Having unique System keys is required. A heuristic can be defined in individual System subclasses to attempt to always generate the same System Name Key. The NameFormat property identifies how the System name was generated, using the heuristic of the subclass.
	NameFormat string

	// OtherIdentifyingInfo captures additional data, beyond System Name information, that could be used to identify a ComputerSystem. One example would be to hold the Fibre Channel World-Wide Name (WWN) of a node. Note that if only the Fibre Channel name is available and is unique (able to be used as the System key), then this property would be NULL and the WWN would become the System key, its data placed in the Name property.
	OtherIdentifyingInfo []string

	// A string that provides information on how the primary system owner can be reached (for example, phone number, e-mail address, and so on).
	PrimaryOwnerContact string

	// The name of the primary system owner. The system owner is the primary user of the system.
	PrimaryOwnerName string

	// An array (bag) of strings that specifies the administrator -defined roles this System plays in the managed environment. Examples might be 'Building 8 print server' or 'Boise user directories'. A single system may perform multiple roles.
	///Note that the instrumentation view of the 'roles' of a System is defined by instantiating a specific subclass of System, or by properties in a subclass, or both. For example, the purpose of a ComputerSystem is defined using the Dedicated and OtherDedicatedDescription properties.
	Roles []string
}

func NewCIM_SystemEx1(instance *cim.WmiInstance) (newInstance *CIM_System, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_System{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

func NewCIM_SystemEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_System, err error) {
	tmp, err := NewCIM_EnabledLogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_System{
		CIM_EnabledLogicalElement: tmp,
	}
	return
}

// SetCreationClassName sets the value of CreationClassName for the instance
func (instance *CIM_System) SetPropertyCreationClassName(value string) (err error) {
	return instance.SetProperty("CreationClassName", (value))
}

// GetCreationClassName gets the value of CreationClassName for the instance
func (instance *CIM_System) GetPropertyCreationClassName() (value string, err error) {
	retValue, err := instance.GetProperty("CreationClassName")
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

// SetIdentifyingDescriptions sets the value of IdentifyingDescriptions for the instance
func (instance *CIM_System) SetPropertyIdentifyingDescriptions(value []string) (err error) {
	return instance.SetProperty("IdentifyingDescriptions", (value))
}

// GetIdentifyingDescriptions gets the value of IdentifyingDescriptions for the instance
func (instance *CIM_System) GetPropertyIdentifyingDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("IdentifyingDescriptions")
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

// SetNameFormat sets the value of NameFormat for the instance
func (instance *CIM_System) SetPropertyNameFormat(value string) (err error) {
	return instance.SetProperty("NameFormat", (value))
}

// GetNameFormat gets the value of NameFormat for the instance
func (instance *CIM_System) GetPropertyNameFormat() (value string, err error) {
	retValue, err := instance.GetProperty("NameFormat")
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

// SetOtherIdentifyingInfo sets the value of OtherIdentifyingInfo for the instance
func (instance *CIM_System) SetPropertyOtherIdentifyingInfo(value []string) (err error) {
	return instance.SetProperty("OtherIdentifyingInfo", (value))
}

// GetOtherIdentifyingInfo gets the value of OtherIdentifyingInfo for the instance
func (instance *CIM_System) GetPropertyOtherIdentifyingInfo() (value []string, err error) {
	retValue, err := instance.GetProperty("OtherIdentifyingInfo")
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

// SetPrimaryOwnerContact sets the value of PrimaryOwnerContact for the instance
func (instance *CIM_System) SetPropertyPrimaryOwnerContact(value string) (err error) {
	return instance.SetProperty("PrimaryOwnerContact", (value))
}

// GetPrimaryOwnerContact gets the value of PrimaryOwnerContact for the instance
func (instance *CIM_System) GetPropertyPrimaryOwnerContact() (value string, err error) {
	retValue, err := instance.GetProperty("PrimaryOwnerContact")
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

// SetPrimaryOwnerName sets the value of PrimaryOwnerName for the instance
func (instance *CIM_System) SetPropertyPrimaryOwnerName(value string) (err error) {
	return instance.SetProperty("PrimaryOwnerName", (value))
}

// GetPrimaryOwnerName gets the value of PrimaryOwnerName for the instance
func (instance *CIM_System) GetPropertyPrimaryOwnerName() (value string, err error) {
	retValue, err := instance.GetProperty("PrimaryOwnerName")
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

// SetRoles sets the value of Roles for the instance
func (instance *CIM_System) SetPropertyRoles(value []string) (err error) {
	return instance.SetProperty("Roles", (value))
}

// GetRoles gets the value of Roles for the instance
func (instance *CIM_System) GetPropertyRoles() (value []string, err error) {
	retValue, err := instance.GetProperty("Roles")
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
