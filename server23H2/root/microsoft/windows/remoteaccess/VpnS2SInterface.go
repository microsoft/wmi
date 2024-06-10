// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.Microsoft.Windows.RemoteAccess
//
// ////////////////////////////////////////////
package remoteaccess

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// VpnS2SInterface struct
type VpnS2SInterface struct {
	*cim.WmiInstance

	//
	AdminStatus bool

	//
	AuthenticationMethod string

	//
	Certificate []uint8

	//
	ConnectionState string

	//
	Destination []string

	//
	EapMethod string

	//
	EnableQoS uint32

	//
	IdleDisconnect uint32

	//
	InitiateConfigPayload bool

	//
	InterfaceType string

	//
	InternalIPv4 bool

	//
	InternalIPv6 bool

	//
	IPv4Subnet []string

	//
	IPv4TriggerFilter []string

	//
	IPv4TriggerFilterAction uint32

	//
	IPv6Subnet []string

	//
	IPv6TriggerFilter []string

	//
	IPv6TriggerFilterAction uint32

	//
	LastDisconnectReason uint32

	//
	LastError uint32

	//
	LocalVpnTrafficSelector []VpnTrafficSelector

	//
	MMSALifeTime uint32

	//
	Name string

	//
	NetworkOutageTime uint32

	//
	NumberOfTries uint32

	//
	Persistent bool

	//
	PostConnectionIPv4Subnet []string

	//
	PostConnectionIPv6Subnet []string

	//
	PromoteAlternate bool

	//
	Protocol string

	//
	RemoteVpnTrafficSelector []VpnTrafficSelector

	//
	ResponderAuthenticationMethod string

	//
	RetryInterval uint32

	//
	RxBandwidthKbps uint64

	//
	SADataSizeForRenegotiation uint32

	//
	SALifeTime uint32

	//
	SourceIpAddress string

	//
	TxBandwidthKbps uint64

	//
	UnReachabilityReasons string

	//
	UserName string
}

func NewVpnS2SInterfaceEx1(instance *cim.WmiInstance) (newInstance *VpnS2SInterface, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &VpnS2SInterface{
		WmiInstance: tmp,
	}
	return
}

func NewVpnS2SInterfaceEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *VpnS2SInterface, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &VpnS2SInterface{
		WmiInstance: tmp,
	}
	return
}

// SetAdminStatus sets the value of AdminStatus for the instance
func (instance *VpnS2SInterface) SetPropertyAdminStatus(value bool) (err error) {
	return instance.SetProperty("AdminStatus", (value))
}

// GetAdminStatus gets the value of AdminStatus for the instance
func (instance *VpnS2SInterface) GetPropertyAdminStatus() (value bool, err error) {
	retValue, err := instance.GetProperty("AdminStatus")
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

// SetAuthenticationMethod sets the value of AuthenticationMethod for the instance
func (instance *VpnS2SInterface) SetPropertyAuthenticationMethod(value string) (err error) {
	return instance.SetProperty("AuthenticationMethod", (value))
}

// GetAuthenticationMethod gets the value of AuthenticationMethod for the instance
func (instance *VpnS2SInterface) GetPropertyAuthenticationMethod() (value string, err error) {
	retValue, err := instance.GetProperty("AuthenticationMethod")
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

// SetCertificate sets the value of Certificate for the instance
func (instance *VpnS2SInterface) SetPropertyCertificate(value []uint8) (err error) {
	return instance.SetProperty("Certificate", (value))
}

// GetCertificate gets the value of Certificate for the instance
func (instance *VpnS2SInterface) GetPropertyCertificate() (value []uint8, err error) {
	retValue, err := instance.GetProperty("Certificate")
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

// SetConnectionState sets the value of ConnectionState for the instance
func (instance *VpnS2SInterface) SetPropertyConnectionState(value string) (err error) {
	return instance.SetProperty("ConnectionState", (value))
}

// GetConnectionState gets the value of ConnectionState for the instance
func (instance *VpnS2SInterface) GetPropertyConnectionState() (value string, err error) {
	retValue, err := instance.GetProperty("ConnectionState")
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

// SetDestination sets the value of Destination for the instance
func (instance *VpnS2SInterface) SetPropertyDestination(value []string) (err error) {
	return instance.SetProperty("Destination", (value))
}

// GetDestination gets the value of Destination for the instance
func (instance *VpnS2SInterface) GetPropertyDestination() (value []string, err error) {
	retValue, err := instance.GetProperty("Destination")
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

// SetEapMethod sets the value of EapMethod for the instance
func (instance *VpnS2SInterface) SetPropertyEapMethod(value string) (err error) {
	return instance.SetProperty("EapMethod", (value))
}

// GetEapMethod gets the value of EapMethod for the instance
func (instance *VpnS2SInterface) GetPropertyEapMethod() (value string, err error) {
	retValue, err := instance.GetProperty("EapMethod")
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

// SetEnableQoS sets the value of EnableQoS for the instance
func (instance *VpnS2SInterface) SetPropertyEnableQoS(value uint32) (err error) {
	return instance.SetProperty("EnableQoS", (value))
}

// GetEnableQoS gets the value of EnableQoS for the instance
func (instance *VpnS2SInterface) GetPropertyEnableQoS() (value uint32, err error) {
	retValue, err := instance.GetProperty("EnableQoS")
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

// SetIdleDisconnect sets the value of IdleDisconnect for the instance
func (instance *VpnS2SInterface) SetPropertyIdleDisconnect(value uint32) (err error) {
	return instance.SetProperty("IdleDisconnect", (value))
}

// GetIdleDisconnect gets the value of IdleDisconnect for the instance
func (instance *VpnS2SInterface) GetPropertyIdleDisconnect() (value uint32, err error) {
	retValue, err := instance.GetProperty("IdleDisconnect")
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

// SetInitiateConfigPayload sets the value of InitiateConfigPayload for the instance
func (instance *VpnS2SInterface) SetPropertyInitiateConfigPayload(value bool) (err error) {
	return instance.SetProperty("InitiateConfigPayload", (value))
}

// GetInitiateConfigPayload gets the value of InitiateConfigPayload for the instance
func (instance *VpnS2SInterface) GetPropertyInitiateConfigPayload() (value bool, err error) {
	retValue, err := instance.GetProperty("InitiateConfigPayload")
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

// SetInterfaceType sets the value of InterfaceType for the instance
func (instance *VpnS2SInterface) SetPropertyInterfaceType(value string) (err error) {
	return instance.SetProperty("InterfaceType", (value))
}

// GetInterfaceType gets the value of InterfaceType for the instance
func (instance *VpnS2SInterface) GetPropertyInterfaceType() (value string, err error) {
	retValue, err := instance.GetProperty("InterfaceType")
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

// SetInternalIPv4 sets the value of InternalIPv4 for the instance
func (instance *VpnS2SInterface) SetPropertyInternalIPv4(value bool) (err error) {
	return instance.SetProperty("InternalIPv4", (value))
}

// GetInternalIPv4 gets the value of InternalIPv4 for the instance
func (instance *VpnS2SInterface) GetPropertyInternalIPv4() (value bool, err error) {
	retValue, err := instance.GetProperty("InternalIPv4")
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

// SetInternalIPv6 sets the value of InternalIPv6 for the instance
func (instance *VpnS2SInterface) SetPropertyInternalIPv6(value bool) (err error) {
	return instance.SetProperty("InternalIPv6", (value))
}

// GetInternalIPv6 gets the value of InternalIPv6 for the instance
func (instance *VpnS2SInterface) GetPropertyInternalIPv6() (value bool, err error) {
	retValue, err := instance.GetProperty("InternalIPv6")
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

// SetIPv4Subnet sets the value of IPv4Subnet for the instance
func (instance *VpnS2SInterface) SetPropertyIPv4Subnet(value []string) (err error) {
	return instance.SetProperty("IPv4Subnet", (value))
}

// GetIPv4Subnet gets the value of IPv4Subnet for the instance
func (instance *VpnS2SInterface) GetPropertyIPv4Subnet() (value []string, err error) {
	retValue, err := instance.GetProperty("IPv4Subnet")
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

// SetIPv4TriggerFilter sets the value of IPv4TriggerFilter for the instance
func (instance *VpnS2SInterface) SetPropertyIPv4TriggerFilter(value []string) (err error) {
	return instance.SetProperty("IPv4TriggerFilter", (value))
}

// GetIPv4TriggerFilter gets the value of IPv4TriggerFilter for the instance
func (instance *VpnS2SInterface) GetPropertyIPv4TriggerFilter() (value []string, err error) {
	retValue, err := instance.GetProperty("IPv4TriggerFilter")
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

// SetIPv4TriggerFilterAction sets the value of IPv4TriggerFilterAction for the instance
func (instance *VpnS2SInterface) SetPropertyIPv4TriggerFilterAction(value uint32) (err error) {
	return instance.SetProperty("IPv4TriggerFilterAction", (value))
}

// GetIPv4TriggerFilterAction gets the value of IPv4TriggerFilterAction for the instance
func (instance *VpnS2SInterface) GetPropertyIPv4TriggerFilterAction() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv4TriggerFilterAction")
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

// SetIPv6Subnet sets the value of IPv6Subnet for the instance
func (instance *VpnS2SInterface) SetPropertyIPv6Subnet(value []string) (err error) {
	return instance.SetProperty("IPv6Subnet", (value))
}

// GetIPv6Subnet gets the value of IPv6Subnet for the instance
func (instance *VpnS2SInterface) GetPropertyIPv6Subnet() (value []string, err error) {
	retValue, err := instance.GetProperty("IPv6Subnet")
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

// SetIPv6TriggerFilter sets the value of IPv6TriggerFilter for the instance
func (instance *VpnS2SInterface) SetPropertyIPv6TriggerFilter(value []string) (err error) {
	return instance.SetProperty("IPv6TriggerFilter", (value))
}

// GetIPv6TriggerFilter gets the value of IPv6TriggerFilter for the instance
func (instance *VpnS2SInterface) GetPropertyIPv6TriggerFilter() (value []string, err error) {
	retValue, err := instance.GetProperty("IPv6TriggerFilter")
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

// SetIPv6TriggerFilterAction sets the value of IPv6TriggerFilterAction for the instance
func (instance *VpnS2SInterface) SetPropertyIPv6TriggerFilterAction(value uint32) (err error) {
	return instance.SetProperty("IPv6TriggerFilterAction", (value))
}

// GetIPv6TriggerFilterAction gets the value of IPv6TriggerFilterAction for the instance
func (instance *VpnS2SInterface) GetPropertyIPv6TriggerFilterAction() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPv6TriggerFilterAction")
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

// SetLastDisconnectReason sets the value of LastDisconnectReason for the instance
func (instance *VpnS2SInterface) SetPropertyLastDisconnectReason(value uint32) (err error) {
	return instance.SetProperty("LastDisconnectReason", (value))
}

// GetLastDisconnectReason gets the value of LastDisconnectReason for the instance
func (instance *VpnS2SInterface) GetPropertyLastDisconnectReason() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastDisconnectReason")
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

// SetLastError sets the value of LastError for the instance
func (instance *VpnS2SInterface) SetPropertyLastError(value uint32) (err error) {
	return instance.SetProperty("LastError", (value))
}

// GetLastError gets the value of LastError for the instance
func (instance *VpnS2SInterface) GetPropertyLastError() (value uint32, err error) {
	retValue, err := instance.GetProperty("LastError")
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

// SetLocalVpnTrafficSelector sets the value of LocalVpnTrafficSelector for the instance
func (instance *VpnS2SInterface) SetPropertyLocalVpnTrafficSelector(value []VpnTrafficSelector) (err error) {
	return instance.SetProperty("LocalVpnTrafficSelector", (value))
}

// GetLocalVpnTrafficSelector gets the value of LocalVpnTrafficSelector for the instance
func (instance *VpnS2SInterface) GetPropertyLocalVpnTrafficSelector() (value []VpnTrafficSelector, err error) {
	retValue, err := instance.GetProperty("LocalVpnTrafficSelector")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(VpnTrafficSelector)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " VpnTrafficSelector is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, VpnTrafficSelector(valuetmp))
	}

	return
}

// SetMMSALifeTime sets the value of MMSALifeTime for the instance
func (instance *VpnS2SInterface) SetPropertyMMSALifeTime(value uint32) (err error) {
	return instance.SetProperty("MMSALifeTime", (value))
}

// GetMMSALifeTime gets the value of MMSALifeTime for the instance
func (instance *VpnS2SInterface) GetPropertyMMSALifeTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("MMSALifeTime")
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
func (instance *VpnS2SInterface) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *VpnS2SInterface) GetPropertyName() (value string, err error) {
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

// SetNetworkOutageTime sets the value of NetworkOutageTime for the instance
func (instance *VpnS2SInterface) SetPropertyNetworkOutageTime(value uint32) (err error) {
	return instance.SetProperty("NetworkOutageTime", (value))
}

// GetNetworkOutageTime gets the value of NetworkOutageTime for the instance
func (instance *VpnS2SInterface) GetPropertyNetworkOutageTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("NetworkOutageTime")
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

// SetNumberOfTries sets the value of NumberOfTries for the instance
func (instance *VpnS2SInterface) SetPropertyNumberOfTries(value uint32) (err error) {
	return instance.SetProperty("NumberOfTries", (value))
}

// GetNumberOfTries gets the value of NumberOfTries for the instance
func (instance *VpnS2SInterface) GetPropertyNumberOfTries() (value uint32, err error) {
	retValue, err := instance.GetProperty("NumberOfTries")
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

// SetPersistent sets the value of Persistent for the instance
func (instance *VpnS2SInterface) SetPropertyPersistent(value bool) (err error) {
	return instance.SetProperty("Persistent", (value))
}

// GetPersistent gets the value of Persistent for the instance
func (instance *VpnS2SInterface) GetPropertyPersistent() (value bool, err error) {
	retValue, err := instance.GetProperty("Persistent")
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

// SetPostConnectionIPv4Subnet sets the value of PostConnectionIPv4Subnet for the instance
func (instance *VpnS2SInterface) SetPropertyPostConnectionIPv4Subnet(value []string) (err error) {
	return instance.SetProperty("PostConnectionIPv4Subnet", (value))
}

// GetPostConnectionIPv4Subnet gets the value of PostConnectionIPv4Subnet for the instance
func (instance *VpnS2SInterface) GetPropertyPostConnectionIPv4Subnet() (value []string, err error) {
	retValue, err := instance.GetProperty("PostConnectionIPv4Subnet")
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

// SetPostConnectionIPv6Subnet sets the value of PostConnectionIPv6Subnet for the instance
func (instance *VpnS2SInterface) SetPropertyPostConnectionIPv6Subnet(value []string) (err error) {
	return instance.SetProperty("PostConnectionIPv6Subnet", (value))
}

// GetPostConnectionIPv6Subnet gets the value of PostConnectionIPv6Subnet for the instance
func (instance *VpnS2SInterface) GetPropertyPostConnectionIPv6Subnet() (value []string, err error) {
	retValue, err := instance.GetProperty("PostConnectionIPv6Subnet")
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

// SetPromoteAlternate sets the value of PromoteAlternate for the instance
func (instance *VpnS2SInterface) SetPropertyPromoteAlternate(value bool) (err error) {
	return instance.SetProperty("PromoteAlternate", (value))
}

// GetPromoteAlternate gets the value of PromoteAlternate for the instance
func (instance *VpnS2SInterface) GetPropertyPromoteAlternate() (value bool, err error) {
	retValue, err := instance.GetProperty("PromoteAlternate")
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

// SetProtocol sets the value of Protocol for the instance
func (instance *VpnS2SInterface) SetPropertyProtocol(value string) (err error) {
	return instance.SetProperty("Protocol", (value))
}

// GetProtocol gets the value of Protocol for the instance
func (instance *VpnS2SInterface) GetPropertyProtocol() (value string, err error) {
	retValue, err := instance.GetProperty("Protocol")
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

// SetRemoteVpnTrafficSelector sets the value of RemoteVpnTrafficSelector for the instance
func (instance *VpnS2SInterface) SetPropertyRemoteVpnTrafficSelector(value []VpnTrafficSelector) (err error) {
	return instance.SetProperty("RemoteVpnTrafficSelector", (value))
}

// GetRemoteVpnTrafficSelector gets the value of RemoteVpnTrafficSelector for the instance
func (instance *VpnS2SInterface) GetPropertyRemoteVpnTrafficSelector() (value []VpnTrafficSelector, err error) {
	retValue, err := instance.GetProperty("RemoteVpnTrafficSelector")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(VpnTrafficSelector)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " VpnTrafficSelector is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, VpnTrafficSelector(valuetmp))
	}

	return
}

// SetResponderAuthenticationMethod sets the value of ResponderAuthenticationMethod for the instance
func (instance *VpnS2SInterface) SetPropertyResponderAuthenticationMethod(value string) (err error) {
	return instance.SetProperty("ResponderAuthenticationMethod", (value))
}

// GetResponderAuthenticationMethod gets the value of ResponderAuthenticationMethod for the instance
func (instance *VpnS2SInterface) GetPropertyResponderAuthenticationMethod() (value string, err error) {
	retValue, err := instance.GetProperty("ResponderAuthenticationMethod")
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

// SetRetryInterval sets the value of RetryInterval for the instance
func (instance *VpnS2SInterface) SetPropertyRetryInterval(value uint32) (err error) {
	return instance.SetProperty("RetryInterval", (value))
}

// GetRetryInterval gets the value of RetryInterval for the instance
func (instance *VpnS2SInterface) GetPropertyRetryInterval() (value uint32, err error) {
	retValue, err := instance.GetProperty("RetryInterval")
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

// SetRxBandwidthKbps sets the value of RxBandwidthKbps for the instance
func (instance *VpnS2SInterface) SetPropertyRxBandwidthKbps(value uint64) (err error) {
	return instance.SetProperty("RxBandwidthKbps", (value))
}

// GetRxBandwidthKbps gets the value of RxBandwidthKbps for the instance
func (instance *VpnS2SInterface) GetPropertyRxBandwidthKbps() (value uint64, err error) {
	retValue, err := instance.GetProperty("RxBandwidthKbps")
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

// SetSADataSizeForRenegotiation sets the value of SADataSizeForRenegotiation for the instance
func (instance *VpnS2SInterface) SetPropertySADataSizeForRenegotiation(value uint32) (err error) {
	return instance.SetProperty("SADataSizeForRenegotiation", (value))
}

// GetSADataSizeForRenegotiation gets the value of SADataSizeForRenegotiation for the instance
func (instance *VpnS2SInterface) GetPropertySADataSizeForRenegotiation() (value uint32, err error) {
	retValue, err := instance.GetProperty("SADataSizeForRenegotiation")
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

// SetSALifeTime sets the value of SALifeTime for the instance
func (instance *VpnS2SInterface) SetPropertySALifeTime(value uint32) (err error) {
	return instance.SetProperty("SALifeTime", (value))
}

// GetSALifeTime gets the value of SALifeTime for the instance
func (instance *VpnS2SInterface) GetPropertySALifeTime() (value uint32, err error) {
	retValue, err := instance.GetProperty("SALifeTime")
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

// SetSourceIpAddress sets the value of SourceIpAddress for the instance
func (instance *VpnS2SInterface) SetPropertySourceIpAddress(value string) (err error) {
	return instance.SetProperty("SourceIpAddress", (value))
}

// GetSourceIpAddress gets the value of SourceIpAddress for the instance
func (instance *VpnS2SInterface) GetPropertySourceIpAddress() (value string, err error) {
	retValue, err := instance.GetProperty("SourceIpAddress")
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

// SetTxBandwidthKbps sets the value of TxBandwidthKbps for the instance
func (instance *VpnS2SInterface) SetPropertyTxBandwidthKbps(value uint64) (err error) {
	return instance.SetProperty("TxBandwidthKbps", (value))
}

// GetTxBandwidthKbps gets the value of TxBandwidthKbps for the instance
func (instance *VpnS2SInterface) GetPropertyTxBandwidthKbps() (value uint64, err error) {
	retValue, err := instance.GetProperty("TxBandwidthKbps")
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

// SetUnReachabilityReasons sets the value of UnReachabilityReasons for the instance
func (instance *VpnS2SInterface) SetPropertyUnReachabilityReasons(value string) (err error) {
	return instance.SetProperty("UnReachabilityReasons", (value))
}

// GetUnReachabilityReasons gets the value of UnReachabilityReasons for the instance
func (instance *VpnS2SInterface) GetPropertyUnReachabilityReasons() (value string, err error) {
	retValue, err := instance.GetProperty("UnReachabilityReasons")
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

// SetUserName sets the value of UserName for the instance
func (instance *VpnS2SInterface) SetPropertyUserName(value string) (err error) {
	return instance.SetProperty("UserName", (value))
}

// GetUserName gets the value of UserName for the instance
func (instance *VpnS2SInterface) GetPropertyUserName() (value string, err error) {
	retValue, err := instance.GetProperty("UserName")
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
