// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.RSOP.NSBEDCFEBE_839A_4FE9_8DE6_B18030C86529.Computer
//
// ////////////////////////////////////////////
package computer

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// RSOP_IEConnectionDialUpSettings struct
type RSOP_IEConnectionDialUpSettings struct {
	*cim.WmiInstance

	//
	alternateOffset uint32

	//
	alternatePhoneNumbers string

	//
	areaCode string

	//
	autoDialDll string

	//
	autoDialFunction string

	//
	channels uint32

	//
	connectionName string

	//
	countryCode uint32

	//
	countryID uint32

	//
	customAuthenticationKey uint32

	//
	customDialDll string

	//
	deviceName string

	//
	deviceType string

	//
	dialExtraPercent uint32

	//
	dialExtraSampleSeconds uint32

	//
	dialMode uint32

	//
	encryptionType uint32

	//
	frameSize uint32

	//
	framingProtocol uint32

	//
	guidID string

	//
	hangUpExtraPercent uint32

	//
	hangUpExtraSampleSeconds uint32

	//
	idleDisconnectSeconds uint32

	//
	ipAddress string

	//
	ipDNSAddress string

	//
	ipDNSAddressAlternate string

	//
	ipWINSAddress string

	//
	ipWINSAddressAlternate string

	//
	localPhoneNumber string

	//
	netProtocols uint32

	//
	options uint32

	//
	options2 uint32

	//
	options3 uint32

	//
	rasEntryData []uint8

	//
	rasEntryDataSize uint32

	//
	reserved1 uint32

	//
	reserved2 uint32

	//
	rsopID string

	//
	rsopPrecedence uint32

	//
	scriptFile string

	//
	subEntries uint32

	//
	_type uint32

	//
	vpnStrategy int32

	//
	windowsVersion uint32

	//
	x25Address string

	//
	x25Facilities string

	//
	x25PadType string

	//
	x25UserData string
}

func NewRSOP_IEConnectionDialUpSettingsEx1(instance *cim.WmiInstance) (newInstance *RSOP_IEConnectionDialUpSettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &RSOP_IEConnectionDialUpSettings{
		WmiInstance: tmp,
	}
	return
}

func NewRSOP_IEConnectionDialUpSettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *RSOP_IEConnectionDialUpSettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &RSOP_IEConnectionDialUpSettings{
		WmiInstance: tmp,
	}
	return
}

// SetalternateOffset sets the value of alternateOffset for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyalternateOffset(value uint32) (err error) {
	return instance.SetProperty("alternateOffset", (value))
}

// GetalternateOffset gets the value of alternateOffset for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyalternateOffset() (value uint32, err error) {
	retValue, err := instance.GetProperty("alternateOffset")
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

// SetalternatePhoneNumbers sets the value of alternatePhoneNumbers for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyalternatePhoneNumbers(value string) (err error) {
	return instance.SetProperty("alternatePhoneNumbers", (value))
}

// GetalternatePhoneNumbers gets the value of alternatePhoneNumbers for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyalternatePhoneNumbers() (value string, err error) {
	retValue, err := instance.GetProperty("alternatePhoneNumbers")
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

// SetareaCode sets the value of areaCode for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyareaCode(value string) (err error) {
	return instance.SetProperty("areaCode", (value))
}

// GetareaCode gets the value of areaCode for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyareaCode() (value string, err error) {
	retValue, err := instance.GetProperty("areaCode")
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

// SetautoDialDll sets the value of autoDialDll for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyautoDialDll(value string) (err error) {
	return instance.SetProperty("autoDialDll", (value))
}

// GetautoDialDll gets the value of autoDialDll for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyautoDialDll() (value string, err error) {
	retValue, err := instance.GetProperty("autoDialDll")
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

// SetautoDialFunction sets the value of autoDialFunction for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyautoDialFunction(value string) (err error) {
	return instance.SetProperty("autoDialFunction", (value))
}

// GetautoDialFunction gets the value of autoDialFunction for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyautoDialFunction() (value string, err error) {
	retValue, err := instance.GetProperty("autoDialFunction")
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

// Setchannels sets the value of channels for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertychannels(value uint32) (err error) {
	return instance.SetProperty("channels", (value))
}

// Getchannels gets the value of channels for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertychannels() (value uint32, err error) {
	retValue, err := instance.GetProperty("channels")
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

// SetconnectionName sets the value of connectionName for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyconnectionName(value string) (err error) {
	return instance.SetProperty("connectionName", (value))
}

// GetconnectionName gets the value of connectionName for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyconnectionName() (value string, err error) {
	retValue, err := instance.GetProperty("connectionName")
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

// SetcountryCode sets the value of countryCode for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertycountryCode(value uint32) (err error) {
	return instance.SetProperty("countryCode", (value))
}

// GetcountryCode gets the value of countryCode for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertycountryCode() (value uint32, err error) {
	retValue, err := instance.GetProperty("countryCode")
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

// SetcountryID sets the value of countryID for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertycountryID(value uint32) (err error) {
	return instance.SetProperty("countryID", (value))
}

// GetcountryID gets the value of countryID for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertycountryID() (value uint32, err error) {
	retValue, err := instance.GetProperty("countryID")
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

// SetcustomAuthenticationKey sets the value of customAuthenticationKey for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertycustomAuthenticationKey(value uint32) (err error) {
	return instance.SetProperty("customAuthenticationKey", (value))
}

// GetcustomAuthenticationKey gets the value of customAuthenticationKey for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertycustomAuthenticationKey() (value uint32, err error) {
	retValue, err := instance.GetProperty("customAuthenticationKey")
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

// SetcustomDialDll sets the value of customDialDll for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertycustomDialDll(value string) (err error) {
	return instance.SetProperty("customDialDll", (value))
}

// GetcustomDialDll gets the value of customDialDll for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertycustomDialDll() (value string, err error) {
	retValue, err := instance.GetProperty("customDialDll")
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

// SetdeviceName sets the value of deviceName for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertydeviceName(value string) (err error) {
	return instance.SetProperty("deviceName", (value))
}

// GetdeviceName gets the value of deviceName for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertydeviceName() (value string, err error) {
	retValue, err := instance.GetProperty("deviceName")
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

// SetdeviceType sets the value of deviceType for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertydeviceType(value string) (err error) {
	return instance.SetProperty("deviceType", (value))
}

// GetdeviceType gets the value of deviceType for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertydeviceType() (value string, err error) {
	retValue, err := instance.GetProperty("deviceType")
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

// SetdialExtraPercent sets the value of dialExtraPercent for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertydialExtraPercent(value uint32) (err error) {
	return instance.SetProperty("dialExtraPercent", (value))
}

// GetdialExtraPercent gets the value of dialExtraPercent for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertydialExtraPercent() (value uint32, err error) {
	retValue, err := instance.GetProperty("dialExtraPercent")
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

// SetdialExtraSampleSeconds sets the value of dialExtraSampleSeconds for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertydialExtraSampleSeconds(value uint32) (err error) {
	return instance.SetProperty("dialExtraSampleSeconds", (value))
}

// GetdialExtraSampleSeconds gets the value of dialExtraSampleSeconds for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertydialExtraSampleSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("dialExtraSampleSeconds")
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

// SetdialMode sets the value of dialMode for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertydialMode(value uint32) (err error) {
	return instance.SetProperty("dialMode", (value))
}

// GetdialMode gets the value of dialMode for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertydialMode() (value uint32, err error) {
	retValue, err := instance.GetProperty("dialMode")
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

// SetencryptionType sets the value of encryptionType for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyencryptionType(value uint32) (err error) {
	return instance.SetProperty("encryptionType", (value))
}

// GetencryptionType gets the value of encryptionType for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyencryptionType() (value uint32, err error) {
	retValue, err := instance.GetProperty("encryptionType")
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

// SetframeSize sets the value of frameSize for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyframeSize(value uint32) (err error) {
	return instance.SetProperty("frameSize", (value))
}

// GetframeSize gets the value of frameSize for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyframeSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("frameSize")
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

// SetframingProtocol sets the value of framingProtocol for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyframingProtocol(value uint32) (err error) {
	return instance.SetProperty("framingProtocol", (value))
}

// GetframingProtocol gets the value of framingProtocol for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyframingProtocol() (value uint32, err error) {
	retValue, err := instance.GetProperty("framingProtocol")
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

// SetguidID sets the value of guidID for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyguidID(value string) (err error) {
	return instance.SetProperty("guidID", (value))
}

// GetguidID gets the value of guidID for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyguidID() (value string, err error) {
	retValue, err := instance.GetProperty("guidID")
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

// SethangUpExtraPercent sets the value of hangUpExtraPercent for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyhangUpExtraPercent(value uint32) (err error) {
	return instance.SetProperty("hangUpExtraPercent", (value))
}

// GethangUpExtraPercent gets the value of hangUpExtraPercent for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyhangUpExtraPercent() (value uint32, err error) {
	retValue, err := instance.GetProperty("hangUpExtraPercent")
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

// SethangUpExtraSampleSeconds sets the value of hangUpExtraSampleSeconds for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyhangUpExtraSampleSeconds(value uint32) (err error) {
	return instance.SetProperty("hangUpExtraSampleSeconds", (value))
}

// GethangUpExtraSampleSeconds gets the value of hangUpExtraSampleSeconds for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyhangUpExtraSampleSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("hangUpExtraSampleSeconds")
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

// SetidleDisconnectSeconds sets the value of idleDisconnectSeconds for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyidleDisconnectSeconds(value uint32) (err error) {
	return instance.SetProperty("idleDisconnectSeconds", (value))
}

// GetidleDisconnectSeconds gets the value of idleDisconnectSeconds for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyidleDisconnectSeconds() (value uint32, err error) {
	retValue, err := instance.GetProperty("idleDisconnectSeconds")
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

// SetipAddress sets the value of ipAddress for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyipAddress(value string) (err error) {
	return instance.SetProperty("ipAddress", (value))
}

// GetipAddress gets the value of ipAddress for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyipAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ipAddress")
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

// SetipDNSAddress sets the value of ipDNSAddress for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyipDNSAddress(value string) (err error) {
	return instance.SetProperty("ipDNSAddress", (value))
}

// GetipDNSAddress gets the value of ipDNSAddress for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyipDNSAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ipDNSAddress")
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

// SetipDNSAddressAlternate sets the value of ipDNSAddressAlternate for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyipDNSAddressAlternate(value string) (err error) {
	return instance.SetProperty("ipDNSAddressAlternate", (value))
}

// GetipDNSAddressAlternate gets the value of ipDNSAddressAlternate for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyipDNSAddressAlternate() (value string, err error) {
	retValue, err := instance.GetProperty("ipDNSAddressAlternate")
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

// SetipWINSAddress sets the value of ipWINSAddress for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyipWINSAddress(value string) (err error) {
	return instance.SetProperty("ipWINSAddress", (value))
}

// GetipWINSAddress gets the value of ipWINSAddress for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyipWINSAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ipWINSAddress")
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

// SetipWINSAddressAlternate sets the value of ipWINSAddressAlternate for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyipWINSAddressAlternate(value string) (err error) {
	return instance.SetProperty("ipWINSAddressAlternate", (value))
}

// GetipWINSAddressAlternate gets the value of ipWINSAddressAlternate for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyipWINSAddressAlternate() (value string, err error) {
	retValue, err := instance.GetProperty("ipWINSAddressAlternate")
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

// SetlocalPhoneNumber sets the value of localPhoneNumber for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertylocalPhoneNumber(value string) (err error) {
	return instance.SetProperty("localPhoneNumber", (value))
}

// GetlocalPhoneNumber gets the value of localPhoneNumber for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertylocalPhoneNumber() (value string, err error) {
	retValue, err := instance.GetProperty("localPhoneNumber")
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

// SetnetProtocols sets the value of netProtocols for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertynetProtocols(value uint32) (err error) {
	return instance.SetProperty("netProtocols", (value))
}

// GetnetProtocols gets the value of netProtocols for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertynetProtocols() (value uint32, err error) {
	retValue, err := instance.GetProperty("netProtocols")
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

// Setoptions sets the value of options for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyoptions(value uint32) (err error) {
	return instance.SetProperty("options", (value))
}

// Getoptions gets the value of options for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyoptions() (value uint32, err error) {
	retValue, err := instance.GetProperty("options")
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

// Setoptions2 sets the value of options2 for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyoptions2(value uint32) (err error) {
	return instance.SetProperty("options2", (value))
}

// Getoptions2 gets the value of options2 for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyoptions2() (value uint32, err error) {
	retValue, err := instance.GetProperty("options2")
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

// Setoptions3 sets the value of options3 for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyoptions3(value uint32) (err error) {
	return instance.SetProperty("options3", (value))
}

// Getoptions3 gets the value of options3 for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyoptions3() (value uint32, err error) {
	retValue, err := instance.GetProperty("options3")
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

// SetrasEntryData sets the value of rasEntryData for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyrasEntryData(value []uint8) (err error) {
	return instance.SetProperty("rasEntryData", (value))
}

// GetrasEntryData gets the value of rasEntryData for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyrasEntryData() (value []uint8, err error) {
	retValue, err := instance.GetProperty("rasEntryData")
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

// SetrasEntryDataSize sets the value of rasEntryDataSize for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyrasEntryDataSize(value uint32) (err error) {
	return instance.SetProperty("rasEntryDataSize", (value))
}

// GetrasEntryDataSize gets the value of rasEntryDataSize for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyrasEntryDataSize() (value uint32, err error) {
	retValue, err := instance.GetProperty("rasEntryDataSize")
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

// Setreserved1 sets the value of reserved1 for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyreserved1(value uint32) (err error) {
	return instance.SetProperty("reserved1", (value))
}

// Getreserved1 gets the value of reserved1 for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyreserved1() (value uint32, err error) {
	retValue, err := instance.GetProperty("reserved1")
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

// Setreserved2 sets the value of reserved2 for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyreserved2(value uint32) (err error) {
	return instance.SetProperty("reserved2", (value))
}

// Getreserved2 gets the value of reserved2 for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyreserved2() (value uint32, err error) {
	retValue, err := instance.GetProperty("reserved2")
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

// SetrsopID sets the value of rsopID for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyrsopID(value string) (err error) {
	return instance.SetProperty("rsopID", (value))
}

// GetrsopID gets the value of rsopID for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyrsopID() (value string, err error) {
	retValue, err := instance.GetProperty("rsopID")
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

// SetrsopPrecedence sets the value of rsopPrecedence for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyrsopPrecedence(value uint32) (err error) {
	return instance.SetProperty("rsopPrecedence", (value))
}

// GetrsopPrecedence gets the value of rsopPrecedence for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyrsopPrecedence() (value uint32, err error) {
	retValue, err := instance.GetProperty("rsopPrecedence")
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

// SetscriptFile sets the value of scriptFile for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyscriptFile(value string) (err error) {
	return instance.SetProperty("scriptFile", (value))
}

// GetscriptFile gets the value of scriptFile for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyscriptFile() (value string, err error) {
	retValue, err := instance.GetProperty("scriptFile")
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

// SetsubEntries sets the value of subEntries for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertysubEntries(value uint32) (err error) {
	return instance.SetProperty("subEntries", (value))
}

// GetsubEntries gets the value of subEntries for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertysubEntries() (value uint32, err error) {
	retValue, err := instance.GetProperty("subEntries")
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

// Settype sets the value of type for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertytype(value uint32) (err error) {
	return instance.SetProperty("type", (value))
}

// Gettype gets the value of type for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertytype() (value uint32, err error) {
	retValue, err := instance.GetProperty("type")
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

// SetvpnStrategy sets the value of vpnStrategy for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyvpnStrategy(value int32) (err error) {
	return instance.SetProperty("vpnStrategy", (value))
}

// GetvpnStrategy gets the value of vpnStrategy for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyvpnStrategy() (value int32, err error) {
	retValue, err := instance.GetProperty("vpnStrategy")
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

// SetwindowsVersion sets the value of windowsVersion for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertywindowsVersion(value uint32) (err error) {
	return instance.SetProperty("windowsVersion", (value))
}

// GetwindowsVersion gets the value of windowsVersion for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertywindowsVersion() (value uint32, err error) {
	retValue, err := instance.GetProperty("windowsVersion")
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

// Setx25Address sets the value of x25Address for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyx25Address(value string) (err error) {
	return instance.SetProperty("x25Address", (value))
}

// Getx25Address gets the value of x25Address for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyx25Address() (value string, err error) {
	retValue, err := instance.GetProperty("x25Address")
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

// Setx25Facilities sets the value of x25Facilities for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyx25Facilities(value string) (err error) {
	return instance.SetProperty("x25Facilities", (value))
}

// Getx25Facilities gets the value of x25Facilities for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyx25Facilities() (value string, err error) {
	retValue, err := instance.GetProperty("x25Facilities")
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

// Setx25PadType sets the value of x25PadType for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyx25PadType(value string) (err error) {
	return instance.SetProperty("x25PadType", (value))
}

// Getx25PadType gets the value of x25PadType for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyx25PadType() (value string, err error) {
	retValue, err := instance.GetProperty("x25PadType")
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

// Setx25UserData sets the value of x25UserData for the instance
func (instance *RSOP_IEConnectionDialUpSettings) SetPropertyx25UserData(value string) (err error) {
	return instance.SetProperty("x25UserData", (value))
}

// Getx25UserData gets the value of x25UserData for the instance
func (instance *RSOP_IEConnectionDialUpSettings) GetPropertyx25UserData() (value string, err error) {
	retValue, err := instance.GetProperty("x25UserData")
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
