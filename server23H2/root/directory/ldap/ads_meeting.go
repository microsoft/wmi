// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.directory.LDAP
//////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_meeting struct
type ads_meeting struct {
	*ds_top

	//
	DS_meetingAdvertiseScope string

	//
	DS_meetingApplication []string

	//
	DS_meetingBandwidth []int32

	//
	DS_meetingBlob Uint8Array

	//
	DS_meetingContactInfo string

	//
	DS_meetingDescription string

	//
	DS_meetingEndTime []string

	//
	DS_meetingID string

	//
	DS_meetingIP string

	//
	DS_meetingIsEncrypted string

	//
	DS_meetingKeyword []string

	//
	DS_meetingLanguage []string

	//
	DS_meetingLocation []string

	//
	DS_meetingMaxParticipants int32

	//
	DS_meetingName string

	//
	DS_meetingOriginator string

	//
	DS_meetingOwner string

	//
	DS_meetingProtocol []string

	//
	DS_meetingRating []string

	//
	DS_meetingRecurrence string

	//
	DS_meetingScope []string

	//
	DS_meetingStartTime []string

	//
	DS_meetingType string

	//
	DS_meetingURL []string
}

func Newads_meetingEx1(instance *cim.WmiInstance) (newInstance *ads_meeting, err error) {
	tmp, err := Newds_topEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_meeting{
		ds_top: tmp,
	}
	return
}

func Newads_meetingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_meeting, err error) {
	tmp, err := Newds_topEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_meeting{
		ds_top: tmp,
	}
	return
}

// SetDS_meetingAdvertiseScope sets the value of DS_meetingAdvertiseScope for the instance
func (instance *ads_meeting) SetPropertyDS_meetingAdvertiseScope(value string) (err error) {
	return instance.SetProperty("DS_meetingAdvertiseScope", (value))
}

// GetDS_meetingAdvertiseScope gets the value of DS_meetingAdvertiseScope for the instance
func (instance *ads_meeting) GetPropertyDS_meetingAdvertiseScope() (value string, err error) {
	retValue, err := instance.GetProperty("DS_meetingAdvertiseScope")
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

// SetDS_meetingApplication sets the value of DS_meetingApplication for the instance
func (instance *ads_meeting) SetPropertyDS_meetingApplication(value []string) (err error) {
	return instance.SetProperty("DS_meetingApplication", (value))
}

// GetDS_meetingApplication gets the value of DS_meetingApplication for the instance
func (instance *ads_meeting) GetPropertyDS_meetingApplication() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_meetingApplication")
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

// SetDS_meetingBandwidth sets the value of DS_meetingBandwidth for the instance
func (instance *ads_meeting) SetPropertyDS_meetingBandwidth(value []int32) (err error) {
	return instance.SetProperty("DS_meetingBandwidth", (value))
}

// GetDS_meetingBandwidth gets the value of DS_meetingBandwidth for the instance
func (instance *ads_meeting) GetPropertyDS_meetingBandwidth() (value []int32, err error) {
	retValue, err := instance.GetProperty("DS_meetingBandwidth")
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
		value = append(value, int32(valuetmp))
	}

	return
}

// SetDS_meetingBlob sets the value of DS_meetingBlob for the instance
func (instance *ads_meeting) SetPropertyDS_meetingBlob(value Uint8Array) (err error) {
	return instance.SetProperty("DS_meetingBlob", (value))
}

// GetDS_meetingBlob gets the value of DS_meetingBlob for the instance
func (instance *ads_meeting) GetPropertyDS_meetingBlob() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_meetingBlob")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(Uint8Array)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = Uint8Array(valuetmp)

	return
}

// SetDS_meetingContactInfo sets the value of DS_meetingContactInfo for the instance
func (instance *ads_meeting) SetPropertyDS_meetingContactInfo(value string) (err error) {
	return instance.SetProperty("DS_meetingContactInfo", (value))
}

// GetDS_meetingContactInfo gets the value of DS_meetingContactInfo for the instance
func (instance *ads_meeting) GetPropertyDS_meetingContactInfo() (value string, err error) {
	retValue, err := instance.GetProperty("DS_meetingContactInfo")
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

// SetDS_meetingDescription sets the value of DS_meetingDescription for the instance
func (instance *ads_meeting) SetPropertyDS_meetingDescription(value string) (err error) {
	return instance.SetProperty("DS_meetingDescription", (value))
}

// GetDS_meetingDescription gets the value of DS_meetingDescription for the instance
func (instance *ads_meeting) GetPropertyDS_meetingDescription() (value string, err error) {
	retValue, err := instance.GetProperty("DS_meetingDescription")
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

// SetDS_meetingEndTime sets the value of DS_meetingEndTime for the instance
func (instance *ads_meeting) SetPropertyDS_meetingEndTime(value []string) (err error) {
	return instance.SetProperty("DS_meetingEndTime", (value))
}

// GetDS_meetingEndTime gets the value of DS_meetingEndTime for the instance
func (instance *ads_meeting) GetPropertyDS_meetingEndTime() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_meetingEndTime")
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

// SetDS_meetingID sets the value of DS_meetingID for the instance
func (instance *ads_meeting) SetPropertyDS_meetingID(value string) (err error) {
	return instance.SetProperty("DS_meetingID", (value))
}

// GetDS_meetingID gets the value of DS_meetingID for the instance
func (instance *ads_meeting) GetPropertyDS_meetingID() (value string, err error) {
	retValue, err := instance.GetProperty("DS_meetingID")
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

// SetDS_meetingIP sets the value of DS_meetingIP for the instance
func (instance *ads_meeting) SetPropertyDS_meetingIP(value string) (err error) {
	return instance.SetProperty("DS_meetingIP", (value))
}

// GetDS_meetingIP gets the value of DS_meetingIP for the instance
func (instance *ads_meeting) GetPropertyDS_meetingIP() (value string, err error) {
	retValue, err := instance.GetProperty("DS_meetingIP")
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

// SetDS_meetingIsEncrypted sets the value of DS_meetingIsEncrypted for the instance
func (instance *ads_meeting) SetPropertyDS_meetingIsEncrypted(value string) (err error) {
	return instance.SetProperty("DS_meetingIsEncrypted", (value))
}

// GetDS_meetingIsEncrypted gets the value of DS_meetingIsEncrypted for the instance
func (instance *ads_meeting) GetPropertyDS_meetingIsEncrypted() (value string, err error) {
	retValue, err := instance.GetProperty("DS_meetingIsEncrypted")
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

// SetDS_meetingKeyword sets the value of DS_meetingKeyword for the instance
func (instance *ads_meeting) SetPropertyDS_meetingKeyword(value []string) (err error) {
	return instance.SetProperty("DS_meetingKeyword", (value))
}

// GetDS_meetingKeyword gets the value of DS_meetingKeyword for the instance
func (instance *ads_meeting) GetPropertyDS_meetingKeyword() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_meetingKeyword")
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

// SetDS_meetingLanguage sets the value of DS_meetingLanguage for the instance
func (instance *ads_meeting) SetPropertyDS_meetingLanguage(value []string) (err error) {
	return instance.SetProperty("DS_meetingLanguage", (value))
}

// GetDS_meetingLanguage gets the value of DS_meetingLanguage for the instance
func (instance *ads_meeting) GetPropertyDS_meetingLanguage() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_meetingLanguage")
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

// SetDS_meetingLocation sets the value of DS_meetingLocation for the instance
func (instance *ads_meeting) SetPropertyDS_meetingLocation(value []string) (err error) {
	return instance.SetProperty("DS_meetingLocation", (value))
}

// GetDS_meetingLocation gets the value of DS_meetingLocation for the instance
func (instance *ads_meeting) GetPropertyDS_meetingLocation() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_meetingLocation")
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

// SetDS_meetingMaxParticipants sets the value of DS_meetingMaxParticipants for the instance
func (instance *ads_meeting) SetPropertyDS_meetingMaxParticipants(value int32) (err error) {
	return instance.SetProperty("DS_meetingMaxParticipants", (value))
}

// GetDS_meetingMaxParticipants gets the value of DS_meetingMaxParticipants for the instance
func (instance *ads_meeting) GetPropertyDS_meetingMaxParticipants() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_meetingMaxParticipants")
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

	value = int32(valuetmp)

	return
}

// SetDS_meetingName sets the value of DS_meetingName for the instance
func (instance *ads_meeting) SetPropertyDS_meetingName(value string) (err error) {
	return instance.SetProperty("DS_meetingName", (value))
}

// GetDS_meetingName gets the value of DS_meetingName for the instance
func (instance *ads_meeting) GetPropertyDS_meetingName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_meetingName")
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

// SetDS_meetingOriginator sets the value of DS_meetingOriginator for the instance
func (instance *ads_meeting) SetPropertyDS_meetingOriginator(value string) (err error) {
	return instance.SetProperty("DS_meetingOriginator", (value))
}

// GetDS_meetingOriginator gets the value of DS_meetingOriginator for the instance
func (instance *ads_meeting) GetPropertyDS_meetingOriginator() (value string, err error) {
	retValue, err := instance.GetProperty("DS_meetingOriginator")
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

// SetDS_meetingOwner sets the value of DS_meetingOwner for the instance
func (instance *ads_meeting) SetPropertyDS_meetingOwner(value string) (err error) {
	return instance.SetProperty("DS_meetingOwner", (value))
}

// GetDS_meetingOwner gets the value of DS_meetingOwner for the instance
func (instance *ads_meeting) GetPropertyDS_meetingOwner() (value string, err error) {
	retValue, err := instance.GetProperty("DS_meetingOwner")
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

// SetDS_meetingProtocol sets the value of DS_meetingProtocol for the instance
func (instance *ads_meeting) SetPropertyDS_meetingProtocol(value []string) (err error) {
	return instance.SetProperty("DS_meetingProtocol", (value))
}

// GetDS_meetingProtocol gets the value of DS_meetingProtocol for the instance
func (instance *ads_meeting) GetPropertyDS_meetingProtocol() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_meetingProtocol")
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

// SetDS_meetingRating sets the value of DS_meetingRating for the instance
func (instance *ads_meeting) SetPropertyDS_meetingRating(value []string) (err error) {
	return instance.SetProperty("DS_meetingRating", (value))
}

// GetDS_meetingRating gets the value of DS_meetingRating for the instance
func (instance *ads_meeting) GetPropertyDS_meetingRating() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_meetingRating")
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

// SetDS_meetingRecurrence sets the value of DS_meetingRecurrence for the instance
func (instance *ads_meeting) SetPropertyDS_meetingRecurrence(value string) (err error) {
	return instance.SetProperty("DS_meetingRecurrence", (value))
}

// GetDS_meetingRecurrence gets the value of DS_meetingRecurrence for the instance
func (instance *ads_meeting) GetPropertyDS_meetingRecurrence() (value string, err error) {
	retValue, err := instance.GetProperty("DS_meetingRecurrence")
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

// SetDS_meetingScope sets the value of DS_meetingScope for the instance
func (instance *ads_meeting) SetPropertyDS_meetingScope(value []string) (err error) {
	return instance.SetProperty("DS_meetingScope", (value))
}

// GetDS_meetingScope gets the value of DS_meetingScope for the instance
func (instance *ads_meeting) GetPropertyDS_meetingScope() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_meetingScope")
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

// SetDS_meetingStartTime sets the value of DS_meetingStartTime for the instance
func (instance *ads_meeting) SetPropertyDS_meetingStartTime(value []string) (err error) {
	return instance.SetProperty("DS_meetingStartTime", (value))
}

// GetDS_meetingStartTime gets the value of DS_meetingStartTime for the instance
func (instance *ads_meeting) GetPropertyDS_meetingStartTime() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_meetingStartTime")
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

// SetDS_meetingType sets the value of DS_meetingType for the instance
func (instance *ads_meeting) SetPropertyDS_meetingType(value string) (err error) {
	return instance.SetProperty("DS_meetingType", (value))
}

// GetDS_meetingType gets the value of DS_meetingType for the instance
func (instance *ads_meeting) GetPropertyDS_meetingType() (value string, err error) {
	retValue, err := instance.GetProperty("DS_meetingType")
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

// SetDS_meetingURL sets the value of DS_meetingURL for the instance
func (instance *ads_meeting) SetPropertyDS_meetingURL(value []string) (err error) {
	return instance.SetProperty("DS_meetingURL", (value))
}

// GetDS_meetingURL gets the value of DS_meetingURL for the instance
func (instance *ads_meeting) GetPropertyDS_meetingURL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_meetingURL")
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
