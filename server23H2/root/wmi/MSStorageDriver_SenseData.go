// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSStorageDriver_SenseData struct
type MSStorageDriver_SenseData struct {
	*cim.WmiInstance

	// Additional Sense Code
	additionalSenseCode uint8

	// Additional Sense Code Qualifier
	additionalSenseCodeQualifier uint8

	// Additional Sense Length
	additionalSenseLength uint8

	// Command Specific Information
	commandSpecificInformation []uint8

	// End Of Media
	endOfMedia bool

	// Error Code
	errorCode uint8

	// Field Replaceable Unit Code
	fieldReplaceableUnitCode uint8

	// File Mark
	fileMark bool

	// Incorrect Length
	incorrectLength bool

	// Information
	information []uint8

	// Reserved
	reserved bool

	// Segment Number
	segmentNumber uint8

	// Sense Key
	senseKey uint8

	// Sense Key Specific
	senseKeySpecific []uint8

	// Valid
	valid bool
}

func NewMSStorageDriver_SenseDataEx1(instance *cim.WmiInstance) (newInstance *MSStorageDriver_SenseData, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MSStorageDriver_SenseData{
		WmiInstance: tmp,
	}
	return
}

func NewMSStorageDriver_SenseDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSStorageDriver_SenseData, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSStorageDriver_SenseData{
		WmiInstance: tmp,
	}
	return
}

// SetadditionalSenseCode sets the value of additionalSenseCode for the instance
func (instance *MSStorageDriver_SenseData) SetPropertyadditionalSenseCode(value uint8) (err error) {
	return instance.SetProperty("additionalSenseCode", (value))
}

// GetadditionalSenseCode gets the value of additionalSenseCode for the instance
func (instance *MSStorageDriver_SenseData) GetPropertyadditionalSenseCode() (value uint8, err error) {
	retValue, err := instance.GetProperty("additionalSenseCode")
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

// SetadditionalSenseCodeQualifier sets the value of additionalSenseCodeQualifier for the instance
func (instance *MSStorageDriver_SenseData) SetPropertyadditionalSenseCodeQualifier(value uint8) (err error) {
	return instance.SetProperty("additionalSenseCodeQualifier", (value))
}

// GetadditionalSenseCodeQualifier gets the value of additionalSenseCodeQualifier for the instance
func (instance *MSStorageDriver_SenseData) GetPropertyadditionalSenseCodeQualifier() (value uint8, err error) {
	retValue, err := instance.GetProperty("additionalSenseCodeQualifier")
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

// SetadditionalSenseLength sets the value of additionalSenseLength for the instance
func (instance *MSStorageDriver_SenseData) SetPropertyadditionalSenseLength(value uint8) (err error) {
	return instance.SetProperty("additionalSenseLength", (value))
}

// GetadditionalSenseLength gets the value of additionalSenseLength for the instance
func (instance *MSStorageDriver_SenseData) GetPropertyadditionalSenseLength() (value uint8, err error) {
	retValue, err := instance.GetProperty("additionalSenseLength")
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

// SetcommandSpecificInformation sets the value of commandSpecificInformation for the instance
func (instance *MSStorageDriver_SenseData) SetPropertycommandSpecificInformation(value []uint8) (err error) {
	return instance.SetProperty("commandSpecificInformation", (value))
}

// GetcommandSpecificInformation gets the value of commandSpecificInformation for the instance
func (instance *MSStorageDriver_SenseData) GetPropertycommandSpecificInformation() (value []uint8, err error) {
	retValue, err := instance.GetProperty("commandSpecificInformation")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// SetendOfMedia sets the value of endOfMedia for the instance
func (instance *MSStorageDriver_SenseData) SetPropertyendOfMedia(value bool) (err error) {
	return instance.SetProperty("endOfMedia", (value))
}

// GetendOfMedia gets the value of endOfMedia for the instance
func (instance *MSStorageDriver_SenseData) GetPropertyendOfMedia() (value bool, err error) {
	retValue, err := instance.GetProperty("endOfMedia")
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

// SeterrorCode sets the value of errorCode for the instance
func (instance *MSStorageDriver_SenseData) SetPropertyerrorCode(value uint8) (err error) {
	return instance.SetProperty("errorCode", (value))
}

// GeterrorCode gets the value of errorCode for the instance
func (instance *MSStorageDriver_SenseData) GetPropertyerrorCode() (value uint8, err error) {
	retValue, err := instance.GetProperty("errorCode")
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

// SetfieldReplaceableUnitCode sets the value of fieldReplaceableUnitCode for the instance
func (instance *MSStorageDriver_SenseData) SetPropertyfieldReplaceableUnitCode(value uint8) (err error) {
	return instance.SetProperty("fieldReplaceableUnitCode", (value))
}

// GetfieldReplaceableUnitCode gets the value of fieldReplaceableUnitCode for the instance
func (instance *MSStorageDriver_SenseData) GetPropertyfieldReplaceableUnitCode() (value uint8, err error) {
	retValue, err := instance.GetProperty("fieldReplaceableUnitCode")
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

// SetfileMark sets the value of fileMark for the instance
func (instance *MSStorageDriver_SenseData) SetPropertyfileMark(value bool) (err error) {
	return instance.SetProperty("fileMark", (value))
}

// GetfileMark gets the value of fileMark for the instance
func (instance *MSStorageDriver_SenseData) GetPropertyfileMark() (value bool, err error) {
	retValue, err := instance.GetProperty("fileMark")
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

// SetincorrectLength sets the value of incorrectLength for the instance
func (instance *MSStorageDriver_SenseData) SetPropertyincorrectLength(value bool) (err error) {
	return instance.SetProperty("incorrectLength", (value))
}

// GetincorrectLength gets the value of incorrectLength for the instance
func (instance *MSStorageDriver_SenseData) GetPropertyincorrectLength() (value bool, err error) {
	retValue, err := instance.GetProperty("incorrectLength")
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

// Setinformation sets the value of information for the instance
func (instance *MSStorageDriver_SenseData) SetPropertyinformation(value []uint8) (err error) {
	return instance.SetProperty("information", (value))
}

// Getinformation gets the value of information for the instance
func (instance *MSStorageDriver_SenseData) GetPropertyinformation() (value []uint8, err error) {
	retValue, err := instance.GetProperty("information")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// Setreserved sets the value of reserved for the instance
func (instance *MSStorageDriver_SenseData) SetPropertyreserved(value bool) (err error) {
	return instance.SetProperty("reserved", (value))
}

// Getreserved gets the value of reserved for the instance
func (instance *MSStorageDriver_SenseData) GetPropertyreserved() (value bool, err error) {
	retValue, err := instance.GetProperty("reserved")
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

// SetsegmentNumber sets the value of segmentNumber for the instance
func (instance *MSStorageDriver_SenseData) SetPropertysegmentNumber(value uint8) (err error) {
	return instance.SetProperty("segmentNumber", (value))
}

// GetsegmentNumber gets the value of segmentNumber for the instance
func (instance *MSStorageDriver_SenseData) GetPropertysegmentNumber() (value uint8, err error) {
	retValue, err := instance.GetProperty("segmentNumber")
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

// SetsenseKey sets the value of senseKey for the instance
func (instance *MSStorageDriver_SenseData) SetPropertysenseKey(value uint8) (err error) {
	return instance.SetProperty("senseKey", (value))
}

// GetsenseKey gets the value of senseKey for the instance
func (instance *MSStorageDriver_SenseData) GetPropertysenseKey() (value uint8, err error) {
	retValue, err := instance.GetProperty("senseKey")
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

// SetsenseKeySpecific sets the value of senseKeySpecific for the instance
func (instance *MSStorageDriver_SenseData) SetPropertysenseKeySpecific(value []uint8) (err error) {
	return instance.SetProperty("senseKeySpecific", (value))
}

// GetsenseKeySpecific gets the value of senseKeySpecific for the instance
func (instance *MSStorageDriver_SenseData) GetPropertysenseKeySpecific() (value []uint8, err error) {
	retValue, err := instance.GetProperty("senseKeySpecific")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(uint8)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " uint8 is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, uint8(valuetmp))
	}

	return
}

// Setvalid sets the value of valid for the instance
func (instance *MSStorageDriver_SenseData) SetPropertyvalid(value bool) (err error) {
	return instance.SetProperty("valid", (value))
}

// Getvalid gets the value of valid for the instance
func (instance *MSStorageDriver_SenseData) GetPropertyvalid() (value bool, err error) {
	retValue, err := instance.GetProperty("valid")
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
