// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.ADFS
//////////////////////////////////////////////
package adfs

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ProxyService struct
type ProxyService struct {
	*cim.WmiInstance

	//
	CongestionControlConnectionTimeout int32

	//
	CongestionControlEnabled bool

	//
	CongestionControlLatencyThreshold int32

	//
	CongestionControlMinWindowSize int32

	//
	ConnectionPoolScavengeInterval int32

	//
	ConnectionPoolSize int32

	//
	EventLogLevel int32

	//
	FarmBehavior int32

	//
	ForwardHttpProxyAddress string

	//
	HostHttpPort int32

	//
	HostHttpsPort int32

	//
	HostName string

	//
	HostNameUserTlsAuth string

	//
	IgnoreTokenBinding bool

	//
	ProxyTrustCertificateKeyLengthInBits int32

	//
	ProxyTrustCertThumbprint string

	//
	ProxyTrustRenewPeriod int32

	//
	TlsClientPort int32

	//
	UpdatedFarmBehaviorLevel int32
}

func NewProxyServiceEx1(instance *cim.WmiInstance) (newInstance *ProxyService, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &ProxyService{
		WmiInstance: tmp,
	}
	return
}

func NewProxyServiceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ProxyService, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ProxyService{
		WmiInstance: tmp,
	}
	return
}

// SetCongestionControlConnectionTimeout sets the value of CongestionControlConnectionTimeout for the instance
func (instance *ProxyService) SetPropertyCongestionControlConnectionTimeout(value int32) (err error) {
	return instance.SetProperty("CongestionControlConnectionTimeout", (value))
}

// GetCongestionControlConnectionTimeout gets the value of CongestionControlConnectionTimeout for the instance
func (instance *ProxyService) GetPropertyCongestionControlConnectionTimeout() (value int32, err error) {
	retValue, err := instance.GetProperty("CongestionControlConnectionTimeout")
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

// SetCongestionControlEnabled sets the value of CongestionControlEnabled for the instance
func (instance *ProxyService) SetPropertyCongestionControlEnabled(value bool) (err error) {
	return instance.SetProperty("CongestionControlEnabled", (value))
}

// GetCongestionControlEnabled gets the value of CongestionControlEnabled for the instance
func (instance *ProxyService) GetPropertyCongestionControlEnabled() (value bool, err error) {
	retValue, err := instance.GetProperty("CongestionControlEnabled")
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

// SetCongestionControlLatencyThreshold sets the value of CongestionControlLatencyThreshold for the instance
func (instance *ProxyService) SetPropertyCongestionControlLatencyThreshold(value int32) (err error) {
	return instance.SetProperty("CongestionControlLatencyThreshold", (value))
}

// GetCongestionControlLatencyThreshold gets the value of CongestionControlLatencyThreshold for the instance
func (instance *ProxyService) GetPropertyCongestionControlLatencyThreshold() (value int32, err error) {
	retValue, err := instance.GetProperty("CongestionControlLatencyThreshold")
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

// SetCongestionControlMinWindowSize sets the value of CongestionControlMinWindowSize for the instance
func (instance *ProxyService) SetPropertyCongestionControlMinWindowSize(value int32) (err error) {
	return instance.SetProperty("CongestionControlMinWindowSize", (value))
}

// GetCongestionControlMinWindowSize gets the value of CongestionControlMinWindowSize for the instance
func (instance *ProxyService) GetPropertyCongestionControlMinWindowSize() (value int32, err error) {
	retValue, err := instance.GetProperty("CongestionControlMinWindowSize")
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

// SetConnectionPoolScavengeInterval sets the value of ConnectionPoolScavengeInterval for the instance
func (instance *ProxyService) SetPropertyConnectionPoolScavengeInterval(value int32) (err error) {
	return instance.SetProperty("ConnectionPoolScavengeInterval", (value))
}

// GetConnectionPoolScavengeInterval gets the value of ConnectionPoolScavengeInterval for the instance
func (instance *ProxyService) GetPropertyConnectionPoolScavengeInterval() (value int32, err error) {
	retValue, err := instance.GetProperty("ConnectionPoolScavengeInterval")
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

// SetConnectionPoolSize sets the value of ConnectionPoolSize for the instance
func (instance *ProxyService) SetPropertyConnectionPoolSize(value int32) (err error) {
	return instance.SetProperty("ConnectionPoolSize", (value))
}

// GetConnectionPoolSize gets the value of ConnectionPoolSize for the instance
func (instance *ProxyService) GetPropertyConnectionPoolSize() (value int32, err error) {
	retValue, err := instance.GetProperty("ConnectionPoolSize")
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

// SetEventLogLevel sets the value of EventLogLevel for the instance
func (instance *ProxyService) SetPropertyEventLogLevel(value int32) (err error) {
	return instance.SetProperty("EventLogLevel", (value))
}

// GetEventLogLevel gets the value of EventLogLevel for the instance
func (instance *ProxyService) GetPropertyEventLogLevel() (value int32, err error) {
	retValue, err := instance.GetProperty("EventLogLevel")
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

// SetFarmBehavior sets the value of FarmBehavior for the instance
func (instance *ProxyService) SetPropertyFarmBehavior(value int32) (err error) {
	return instance.SetProperty("FarmBehavior", (value))
}

// GetFarmBehavior gets the value of FarmBehavior for the instance
func (instance *ProxyService) GetPropertyFarmBehavior() (value int32, err error) {
	retValue, err := instance.GetProperty("FarmBehavior")
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

// SetForwardHttpProxyAddress sets the value of ForwardHttpProxyAddress for the instance
func (instance *ProxyService) SetPropertyForwardHttpProxyAddress(value string) (err error) {
	return instance.SetProperty("ForwardHttpProxyAddress", (value))
}

// GetForwardHttpProxyAddress gets the value of ForwardHttpProxyAddress for the instance
func (instance *ProxyService) GetPropertyForwardHttpProxyAddress() (value string, err error) {
	retValue, err := instance.GetProperty("ForwardHttpProxyAddress")
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

// SetHostHttpPort sets the value of HostHttpPort for the instance
func (instance *ProxyService) SetPropertyHostHttpPort(value int32) (err error) {
	return instance.SetProperty("HostHttpPort", (value))
}

// GetHostHttpPort gets the value of HostHttpPort for the instance
func (instance *ProxyService) GetPropertyHostHttpPort() (value int32, err error) {
	retValue, err := instance.GetProperty("HostHttpPort")
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

// SetHostHttpsPort sets the value of HostHttpsPort for the instance
func (instance *ProxyService) SetPropertyHostHttpsPort(value int32) (err error) {
	return instance.SetProperty("HostHttpsPort", (value))
}

// GetHostHttpsPort gets the value of HostHttpsPort for the instance
func (instance *ProxyService) GetPropertyHostHttpsPort() (value int32, err error) {
	retValue, err := instance.GetProperty("HostHttpsPort")
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

// SetHostName sets the value of HostName for the instance
func (instance *ProxyService) SetPropertyHostName(value string) (err error) {
	return instance.SetProperty("HostName", (value))
}

// GetHostName gets the value of HostName for the instance
func (instance *ProxyService) GetPropertyHostName() (value string, err error) {
	retValue, err := instance.GetProperty("HostName")
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

// SetHostNameUserTlsAuth sets the value of HostNameUserTlsAuth for the instance
func (instance *ProxyService) SetPropertyHostNameUserTlsAuth(value string) (err error) {
	return instance.SetProperty("HostNameUserTlsAuth", (value))
}

// GetHostNameUserTlsAuth gets the value of HostNameUserTlsAuth for the instance
func (instance *ProxyService) GetPropertyHostNameUserTlsAuth() (value string, err error) {
	retValue, err := instance.GetProperty("HostNameUserTlsAuth")
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

// SetIgnoreTokenBinding sets the value of IgnoreTokenBinding for the instance
func (instance *ProxyService) SetPropertyIgnoreTokenBinding(value bool) (err error) {
	return instance.SetProperty("IgnoreTokenBinding", (value))
}

// GetIgnoreTokenBinding gets the value of IgnoreTokenBinding for the instance
func (instance *ProxyService) GetPropertyIgnoreTokenBinding() (value bool, err error) {
	retValue, err := instance.GetProperty("IgnoreTokenBinding")
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

// SetProxyTrustCertificateKeyLengthInBits sets the value of ProxyTrustCertificateKeyLengthInBits for the instance
func (instance *ProxyService) SetPropertyProxyTrustCertificateKeyLengthInBits(value int32) (err error) {
	return instance.SetProperty("ProxyTrustCertificateKeyLengthInBits", (value))
}

// GetProxyTrustCertificateKeyLengthInBits gets the value of ProxyTrustCertificateKeyLengthInBits for the instance
func (instance *ProxyService) GetPropertyProxyTrustCertificateKeyLengthInBits() (value int32, err error) {
	retValue, err := instance.GetProperty("ProxyTrustCertificateKeyLengthInBits")
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

// SetProxyTrustCertThumbprint sets the value of ProxyTrustCertThumbprint for the instance
func (instance *ProxyService) SetPropertyProxyTrustCertThumbprint(value string) (err error) {
	return instance.SetProperty("ProxyTrustCertThumbprint", (value))
}

// GetProxyTrustCertThumbprint gets the value of ProxyTrustCertThumbprint for the instance
func (instance *ProxyService) GetPropertyProxyTrustCertThumbprint() (value string, err error) {
	retValue, err := instance.GetProperty("ProxyTrustCertThumbprint")
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

// SetProxyTrustRenewPeriod sets the value of ProxyTrustRenewPeriod for the instance
func (instance *ProxyService) SetPropertyProxyTrustRenewPeriod(value int32) (err error) {
	return instance.SetProperty("ProxyTrustRenewPeriod", (value))
}

// GetProxyTrustRenewPeriod gets the value of ProxyTrustRenewPeriod for the instance
func (instance *ProxyService) GetPropertyProxyTrustRenewPeriod() (value int32, err error) {
	retValue, err := instance.GetProperty("ProxyTrustRenewPeriod")
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

// SetTlsClientPort sets the value of TlsClientPort for the instance
func (instance *ProxyService) SetPropertyTlsClientPort(value int32) (err error) {
	return instance.SetProperty("TlsClientPort", (value))
}

// GetTlsClientPort gets the value of TlsClientPort for the instance
func (instance *ProxyService) GetPropertyTlsClientPort() (value int32, err error) {
	retValue, err := instance.GetProperty("TlsClientPort")
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

// SetUpdatedFarmBehaviorLevel sets the value of UpdatedFarmBehaviorLevel for the instance
func (instance *ProxyService) SetPropertyUpdatedFarmBehaviorLevel(value int32) (err error) {
	return instance.SetProperty("UpdatedFarmBehaviorLevel", (value))
}

// GetUpdatedFarmBehaviorLevel gets the value of UpdatedFarmBehaviorLevel for the instance
func (instance *ProxyService) GetPropertyUpdatedFarmBehaviorLevel() (value int32, err error) {
	retValue, err := instance.GetProperty("UpdatedFarmBehaviorLevel")
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
