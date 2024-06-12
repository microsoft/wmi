// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/12/2024 using wmigen
//      Source root.ServiceModel
//////////////////////////////////////////////
package servicemodel

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// LocalServiceSecuritySettings struct
type LocalServiceSecuritySettings struct {
	*cim.WmiInstance

	// A Boolean value that specifies whether replay attacks against the channel are detected and dealt with automatically.
	DetectReplays bool

	// A value that specifies the maximum number of pending security sessions that the service supports.
	InactivityTimeout string

	// A TimeSpan that specifies the lifetime issued to all new security cookies.
	IssuedCookieLifetime string

	// A value that specifies the maximum number of cookies that can be cached.
	MaxCachedCookies int32

	// A TimeSpan that specifies the maximum time difference between the system clocks of the two communicating parties.
	MaxClockSkew string

	// The maximum number of pending connections on the service.
	MaxPendingSessions int32

	// A value that specifies the number of security negotiations that can be active concurrently.
	MaxStatefulNegotiations int32

	// A TimeSpan that specifies the maximum duration for the security negotiation phase between server and client.
	NegotiationTimeout string

	// A Boolean value that specifies whether connections using WS-Reliable messaging will attempt to reconnect after transport failures.
	ReconnectTransportOnFailure bool

	// A value that specifies the number of cached nonces used for replay detection.
	ReplayCacheSize int32

	// A TimeSpan that specifies the duration in which individual message nonces are valid.
	ReplayWindow string

	// A TimeSpan that specifies the duration after which the initiator will renew the key for the security session.
	SessionKeyRenewalInterval string

	// A TimeSpan that specifies the time interval a previous session key is valid on incoming messages during a key renewal.
	SessionKeyRolloverInterval string

	// A TimeSpan that specifies the duration in which a time stamp is valid.
	TimestampValidityDuration string
}

func NewLocalServiceSecuritySettingsEx1(instance *cim.WmiInstance) (newInstance *LocalServiceSecuritySettings, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &LocalServiceSecuritySettings{
		WmiInstance: tmp,
	}
	return
}

func NewLocalServiceSecuritySettingsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *LocalServiceSecuritySettings, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &LocalServiceSecuritySettings{
		WmiInstance: tmp,
	}
	return
}

// SetDetectReplays sets the value of DetectReplays for the instance
func (instance *LocalServiceSecuritySettings) SetPropertyDetectReplays(value bool) (err error) {
	return instance.SetProperty("DetectReplays", (value))
}

// GetDetectReplays gets the value of DetectReplays for the instance
func (instance *LocalServiceSecuritySettings) GetPropertyDetectReplays() (value bool, err error) {
	retValue, err := instance.GetProperty("DetectReplays")
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

// SetInactivityTimeout sets the value of InactivityTimeout for the instance
func (instance *LocalServiceSecuritySettings) SetPropertyInactivityTimeout(value string) (err error) {
	return instance.SetProperty("InactivityTimeout", (value))
}

// GetInactivityTimeout gets the value of InactivityTimeout for the instance
func (instance *LocalServiceSecuritySettings) GetPropertyInactivityTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("InactivityTimeout")
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

// SetIssuedCookieLifetime sets the value of IssuedCookieLifetime for the instance
func (instance *LocalServiceSecuritySettings) SetPropertyIssuedCookieLifetime(value string) (err error) {
	return instance.SetProperty("IssuedCookieLifetime", (value))
}

// GetIssuedCookieLifetime gets the value of IssuedCookieLifetime for the instance
func (instance *LocalServiceSecuritySettings) GetPropertyIssuedCookieLifetime() (value string, err error) {
	retValue, err := instance.GetProperty("IssuedCookieLifetime")
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

// SetMaxCachedCookies sets the value of MaxCachedCookies for the instance
func (instance *LocalServiceSecuritySettings) SetPropertyMaxCachedCookies(value int32) (err error) {
	return instance.SetProperty("MaxCachedCookies", (value))
}

// GetMaxCachedCookies gets the value of MaxCachedCookies for the instance
func (instance *LocalServiceSecuritySettings) GetPropertyMaxCachedCookies() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxCachedCookies")
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

// SetMaxClockSkew sets the value of MaxClockSkew for the instance
func (instance *LocalServiceSecuritySettings) SetPropertyMaxClockSkew(value string) (err error) {
	return instance.SetProperty("MaxClockSkew", (value))
}

// GetMaxClockSkew gets the value of MaxClockSkew for the instance
func (instance *LocalServiceSecuritySettings) GetPropertyMaxClockSkew() (value string, err error) {
	retValue, err := instance.GetProperty("MaxClockSkew")
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

// SetMaxPendingSessions sets the value of MaxPendingSessions for the instance
func (instance *LocalServiceSecuritySettings) SetPropertyMaxPendingSessions(value int32) (err error) {
	return instance.SetProperty("MaxPendingSessions", (value))
}

// GetMaxPendingSessions gets the value of MaxPendingSessions for the instance
func (instance *LocalServiceSecuritySettings) GetPropertyMaxPendingSessions() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxPendingSessions")
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

// SetMaxStatefulNegotiations sets the value of MaxStatefulNegotiations for the instance
func (instance *LocalServiceSecuritySettings) SetPropertyMaxStatefulNegotiations(value int32) (err error) {
	return instance.SetProperty("MaxStatefulNegotiations", (value))
}

// GetMaxStatefulNegotiations gets the value of MaxStatefulNegotiations for the instance
func (instance *LocalServiceSecuritySettings) GetPropertyMaxStatefulNegotiations() (value int32, err error) {
	retValue, err := instance.GetProperty("MaxStatefulNegotiations")
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

// SetNegotiationTimeout sets the value of NegotiationTimeout for the instance
func (instance *LocalServiceSecuritySettings) SetPropertyNegotiationTimeout(value string) (err error) {
	return instance.SetProperty("NegotiationTimeout", (value))
}

// GetNegotiationTimeout gets the value of NegotiationTimeout for the instance
func (instance *LocalServiceSecuritySettings) GetPropertyNegotiationTimeout() (value string, err error) {
	retValue, err := instance.GetProperty("NegotiationTimeout")
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

// SetReconnectTransportOnFailure sets the value of ReconnectTransportOnFailure for the instance
func (instance *LocalServiceSecuritySettings) SetPropertyReconnectTransportOnFailure(value bool) (err error) {
	return instance.SetProperty("ReconnectTransportOnFailure", (value))
}

// GetReconnectTransportOnFailure gets the value of ReconnectTransportOnFailure for the instance
func (instance *LocalServiceSecuritySettings) GetPropertyReconnectTransportOnFailure() (value bool, err error) {
	retValue, err := instance.GetProperty("ReconnectTransportOnFailure")
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

// SetReplayCacheSize sets the value of ReplayCacheSize for the instance
func (instance *LocalServiceSecuritySettings) SetPropertyReplayCacheSize(value int32) (err error) {
	return instance.SetProperty("ReplayCacheSize", (value))
}

// GetReplayCacheSize gets the value of ReplayCacheSize for the instance
func (instance *LocalServiceSecuritySettings) GetPropertyReplayCacheSize() (value int32, err error) {
	retValue, err := instance.GetProperty("ReplayCacheSize")
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

// SetReplayWindow sets the value of ReplayWindow for the instance
func (instance *LocalServiceSecuritySettings) SetPropertyReplayWindow(value string) (err error) {
	return instance.SetProperty("ReplayWindow", (value))
}

// GetReplayWindow gets the value of ReplayWindow for the instance
func (instance *LocalServiceSecuritySettings) GetPropertyReplayWindow() (value string, err error) {
	retValue, err := instance.GetProperty("ReplayWindow")
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

// SetSessionKeyRenewalInterval sets the value of SessionKeyRenewalInterval for the instance
func (instance *LocalServiceSecuritySettings) SetPropertySessionKeyRenewalInterval(value string) (err error) {
	return instance.SetProperty("SessionKeyRenewalInterval", (value))
}

// GetSessionKeyRenewalInterval gets the value of SessionKeyRenewalInterval for the instance
func (instance *LocalServiceSecuritySettings) GetPropertySessionKeyRenewalInterval() (value string, err error) {
	retValue, err := instance.GetProperty("SessionKeyRenewalInterval")
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

// SetSessionKeyRolloverInterval sets the value of SessionKeyRolloverInterval for the instance
func (instance *LocalServiceSecuritySettings) SetPropertySessionKeyRolloverInterval(value string) (err error) {
	return instance.SetProperty("SessionKeyRolloverInterval", (value))
}

// GetSessionKeyRolloverInterval gets the value of SessionKeyRolloverInterval for the instance
func (instance *LocalServiceSecuritySettings) GetPropertySessionKeyRolloverInterval() (value string, err error) {
	retValue, err := instance.GetProperty("SessionKeyRolloverInterval")
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

// SetTimestampValidityDuration sets the value of TimestampValidityDuration for the instance
func (instance *LocalServiceSecuritySettings) SetPropertyTimestampValidityDuration(value string) (err error) {
	return instance.SetProperty("TimestampValidityDuration", (value))
}

// GetTimestampValidityDuration gets the value of TimestampValidityDuration for the instance
func (instance *LocalServiceSecuritySettings) GetPropertyTimestampValidityDuration() (value string, err error) {
	retValue, err := instance.GetProperty("TimestampValidityDuration")
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
