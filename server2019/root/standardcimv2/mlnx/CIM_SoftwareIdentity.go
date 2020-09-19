// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_SoftwareIdentity struct
type CIM_SoftwareIdentity struct {
	*CIM_LogicalElement

	//
	BuildNumber uint16

	//
	ClassificationDescriptions []string

	//
	Classifications []SoftwareIdentity_Classifications

	//
	ExtendedResourceType SoftwareIdentity_ExtendedResourceType

	//
	IdentityInfoType []string

	//
	IdentityInfoValue []string

	//
	IsEntity bool

	//
	IsLargeBuildNumber bool

	//
	Languages []string

	//
	LargeBuildNumber uint64

	//
	MajorVersion uint16

	//
	Manufacturer string

	//
	MinExtendedResourceTypeBuildNumber uint16

	//
	MinExtendedResourceTypeMajorVersion uint16

	//
	MinExtendedResourceTypeMinorVersion uint16

	//
	MinExtendedResourceTypeRevisionNumber uint16

	//
	MinorVersion uint16

	//
	OtherExtendedResourceTypeDescription string

	//
	ReleaseDate string

	//
	RevisionNumber uint16

	//
	SerialNumber string

	//
	TargetOperatingSystems []string

	//
	TargetOSTypes []uint16

	//
	TargetTypes []string

	//
	VersionString string
}

func NewCIM_SoftwareIdentityEx1(instance *cim.WmiInstance) (newInstance *CIM_SoftwareIdentity, err error) {
	tmp, err := NewCIM_LogicalElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_SoftwareIdentity{
		CIM_LogicalElement: tmp,
	}
	return
}

func NewCIM_SoftwareIdentityEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_SoftwareIdentity, err error) {
	tmp, err := NewCIM_LogicalElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_SoftwareIdentity{
		CIM_LogicalElement: tmp,
	}
	return
}

// SetBuildNumber sets the value of BuildNumber for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyBuildNumber(value uint16) (err error) {
	return instance.SetProperty("BuildNumber", (value))
}

// GetBuildNumber gets the value of BuildNumber for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyBuildNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("BuildNumber")
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

// SetClassificationDescriptions sets the value of ClassificationDescriptions for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyClassificationDescriptions(value []string) (err error) {
	return instance.SetProperty("ClassificationDescriptions", (value))
}

// GetClassificationDescriptions gets the value of ClassificationDescriptions for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyClassificationDescriptions() (value []string, err error) {
	retValue, err := instance.GetProperty("ClassificationDescriptions")
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

// SetClassifications sets the value of Classifications for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyClassifications(value []SoftwareIdentity_Classifications) (err error) {
	return instance.SetProperty("Classifications", (value))
}

// GetClassifications gets the value of Classifications for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyClassifications() (value []SoftwareIdentity_Classifications, err error) {
	retValue, err := instance.GetProperty("Classifications")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(int32)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " int32 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, SoftwareIdentity_Classifications(valuetmp))
	}

	return
}

// SetExtendedResourceType sets the value of ExtendedResourceType for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyExtendedResourceType(value SoftwareIdentity_ExtendedResourceType) (err error) {
	return instance.SetProperty("ExtendedResourceType", (value))
}

// GetExtendedResourceType gets the value of ExtendedResourceType for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyExtendedResourceType() (value SoftwareIdentity_ExtendedResourceType, err error) {
	retValue, err := instance.GetProperty("ExtendedResourceType")
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

	value = SoftwareIdentity_ExtendedResourceType(valuetmp)

	return
}

// SetIdentityInfoType sets the value of IdentityInfoType for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyIdentityInfoType(value []string) (err error) {
	return instance.SetProperty("IdentityInfoType", (value))
}

// GetIdentityInfoType gets the value of IdentityInfoType for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyIdentityInfoType() (value []string, err error) {
	retValue, err := instance.GetProperty("IdentityInfoType")
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

// SetIdentityInfoValue sets the value of IdentityInfoValue for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyIdentityInfoValue(value []string) (err error) {
	return instance.SetProperty("IdentityInfoValue", (value))
}

// GetIdentityInfoValue gets the value of IdentityInfoValue for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyIdentityInfoValue() (value []string, err error) {
	retValue, err := instance.GetProperty("IdentityInfoValue")
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

// SetIsEntity sets the value of IsEntity for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyIsEntity(value bool) (err error) {
	return instance.SetProperty("IsEntity", (value))
}

// GetIsEntity gets the value of IsEntity for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyIsEntity() (value bool, err error) {
	retValue, err := instance.GetProperty("IsEntity")
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

// SetIsLargeBuildNumber sets the value of IsLargeBuildNumber for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyIsLargeBuildNumber(value bool) (err error) {
	return instance.SetProperty("IsLargeBuildNumber", (value))
}

// GetIsLargeBuildNumber gets the value of IsLargeBuildNumber for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyIsLargeBuildNumber() (value bool, err error) {
	retValue, err := instance.GetProperty("IsLargeBuildNumber")
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

// SetLanguages sets the value of Languages for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyLanguages(value []string) (err error) {
	return instance.SetProperty("Languages", (value))
}

// GetLanguages gets the value of Languages for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyLanguages() (value []string, err error) {
	retValue, err := instance.GetProperty("Languages")
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

// SetLargeBuildNumber sets the value of LargeBuildNumber for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyLargeBuildNumber(value uint64) (err error) {
	return instance.SetProperty("LargeBuildNumber", (value))
}

// GetLargeBuildNumber gets the value of LargeBuildNumber for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyLargeBuildNumber() (value uint64, err error) {
	retValue, err := instance.GetProperty("LargeBuildNumber")
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

// SetMajorVersion sets the value of MajorVersion for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyMajorVersion(value uint16) (err error) {
	return instance.SetProperty("MajorVersion", (value))
}

// GetMajorVersion gets the value of MajorVersion for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyMajorVersion() (value uint16, err error) {
	retValue, err := instance.GetProperty("MajorVersion")
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

// SetManufacturer sets the value of Manufacturer for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyManufacturer(value string) (err error) {
	return instance.SetProperty("Manufacturer", (value))
}

// GetManufacturer gets the value of Manufacturer for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyManufacturer() (value string, err error) {
	retValue, err := instance.GetProperty("Manufacturer")
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

// SetMinExtendedResourceTypeBuildNumber sets the value of MinExtendedResourceTypeBuildNumber for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyMinExtendedResourceTypeBuildNumber(value uint16) (err error) {
	return instance.SetProperty("MinExtendedResourceTypeBuildNumber", (value))
}

// GetMinExtendedResourceTypeBuildNumber gets the value of MinExtendedResourceTypeBuildNumber for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyMinExtendedResourceTypeBuildNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("MinExtendedResourceTypeBuildNumber")
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

// SetMinExtendedResourceTypeMajorVersion sets the value of MinExtendedResourceTypeMajorVersion for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyMinExtendedResourceTypeMajorVersion(value uint16) (err error) {
	return instance.SetProperty("MinExtendedResourceTypeMajorVersion", (value))
}

// GetMinExtendedResourceTypeMajorVersion gets the value of MinExtendedResourceTypeMajorVersion for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyMinExtendedResourceTypeMajorVersion() (value uint16, err error) {
	retValue, err := instance.GetProperty("MinExtendedResourceTypeMajorVersion")
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

// SetMinExtendedResourceTypeMinorVersion sets the value of MinExtendedResourceTypeMinorVersion for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyMinExtendedResourceTypeMinorVersion(value uint16) (err error) {
	return instance.SetProperty("MinExtendedResourceTypeMinorVersion", (value))
}

// GetMinExtendedResourceTypeMinorVersion gets the value of MinExtendedResourceTypeMinorVersion for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyMinExtendedResourceTypeMinorVersion() (value uint16, err error) {
	retValue, err := instance.GetProperty("MinExtendedResourceTypeMinorVersion")
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

// SetMinExtendedResourceTypeRevisionNumber sets the value of MinExtendedResourceTypeRevisionNumber for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyMinExtendedResourceTypeRevisionNumber(value uint16) (err error) {
	return instance.SetProperty("MinExtendedResourceTypeRevisionNumber", (value))
}

// GetMinExtendedResourceTypeRevisionNumber gets the value of MinExtendedResourceTypeRevisionNumber for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyMinExtendedResourceTypeRevisionNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("MinExtendedResourceTypeRevisionNumber")
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

// SetMinorVersion sets the value of MinorVersion for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyMinorVersion(value uint16) (err error) {
	return instance.SetProperty("MinorVersion", (value))
}

// GetMinorVersion gets the value of MinorVersion for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyMinorVersion() (value uint16, err error) {
	retValue, err := instance.GetProperty("MinorVersion")
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

// SetOtherExtendedResourceTypeDescription sets the value of OtherExtendedResourceTypeDescription for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyOtherExtendedResourceTypeDescription(value string) (err error) {
	return instance.SetProperty("OtherExtendedResourceTypeDescription", (value))
}

// GetOtherExtendedResourceTypeDescription gets the value of OtherExtendedResourceTypeDescription for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyOtherExtendedResourceTypeDescription() (value string, err error) {
	retValue, err := instance.GetProperty("OtherExtendedResourceTypeDescription")
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

// SetReleaseDate sets the value of ReleaseDate for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyReleaseDate(value string) (err error) {
	return instance.SetProperty("ReleaseDate", (value))
}

// GetReleaseDate gets the value of ReleaseDate for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyReleaseDate() (value string, err error) {
	retValue, err := instance.GetProperty("ReleaseDate")
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

// SetRevisionNumber sets the value of RevisionNumber for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyRevisionNumber(value uint16) (err error) {
	return instance.SetProperty("RevisionNumber", (value))
}

// GetRevisionNumber gets the value of RevisionNumber for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyRevisionNumber() (value uint16, err error) {
	retValue, err := instance.GetProperty("RevisionNumber")
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

// SetSerialNumber sets the value of SerialNumber for the instance
func (instance *CIM_SoftwareIdentity) SetPropertySerialNumber(value string) (err error) {
	return instance.SetProperty("SerialNumber", (value))
}

// GetSerialNumber gets the value of SerialNumber for the instance
func (instance *CIM_SoftwareIdentity) GetPropertySerialNumber() (value string, err error) {
	retValue, err := instance.GetProperty("SerialNumber")
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

// SetTargetOperatingSystems sets the value of TargetOperatingSystems for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyTargetOperatingSystems(value []string) (err error) {
	return instance.SetProperty("TargetOperatingSystems", (value))
}

// GetTargetOperatingSystems gets the value of TargetOperatingSystems for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyTargetOperatingSystems() (value []string, err error) {
	retValue, err := instance.GetProperty("TargetOperatingSystems")
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

// SetTargetOSTypes sets the value of TargetOSTypes for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyTargetOSTypes(value []uint16) (err error) {
	return instance.SetProperty("TargetOSTypes", (value))
}

// GetTargetOSTypes gets the value of TargetOSTypes for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyTargetOSTypes() (value []uint16, err error) {
	retValue, err := instance.GetProperty("TargetOSTypes")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint16)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint16 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint16(valuetmp))
	}

	return
}

// SetTargetTypes sets the value of TargetTypes for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyTargetTypes(value []string) (err error) {
	return instance.SetProperty("TargetTypes", (value))
}

// GetTargetTypes gets the value of TargetTypes for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyTargetTypes() (value []string, err error) {
	retValue, err := instance.GetProperty("TargetTypes")
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

// SetVersionString sets the value of VersionString for the instance
func (instance *CIM_SoftwareIdentity) SetPropertyVersionString(value string) (err error) {
	return instance.SetProperty("VersionString", (value))
}

// GetVersionString gets the value of VersionString for the instance
func (instance *CIM_SoftwareIdentity) GetPropertyVersionString() (value string, err error) {
	retValue, err := instance.GetProperty("VersionString")
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
