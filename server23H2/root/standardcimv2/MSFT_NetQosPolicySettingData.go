// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/14/2024 using wmigen
//      Source root.StandardCimv2
//////////////////////////////////////////////
package standardcimv2

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MSFT_NetQosPolicySettingData struct
type MSFT_NetQosPolicySettingData struct {
	*MSFT_NetSettingData

	//
	AppPathNameMatchCondition string

	//
	DSCPAction int8

	//
	IPDstPortEndMatchCondition uint16

	//
	IPDstPortStartMatchCondition uint16

	//
	IPDstPrefixMatchCondition string

	//
	IPPortMatchCondition uint16

	//
	IPProtocolMatchCondition uint32

	//
	IPSrcPortEndMatchCondition uint16

	//
	IPSrcPortStartMatchCondition uint16

	//
	IPSrcPrefixMatchCondition string

	//
	JobObjectMatchCondition string

	//
	MinBandwidthWeightAction uint8

	//
	Name string

	//
	NetDirectPortMatchCondition uint16

	//
	NetworkProfile uint32

	//
	Owner string

	//
	Precedence uint32

	//
	PriorityValue8021Action int8

	//
	TemplateMatchCondition uint32

	//
	ThrottleRateAction uint64

	//
	URIMatchCondition string

	//
	URIRecursiveMatchCondition bool

	//
	UserMatchCondition string

	//
	Version string
}

func NewMSFT_NetQosPolicySettingDataEx1(instance *cim.WmiInstance) (newInstance *MSFT_NetQosPolicySettingData, err error) {
	tmp, err := NewMSFT_NetSettingDataEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetQosPolicySettingData{
		MSFT_NetSettingData: tmp,
	}
	return
}

func NewMSFT_NetQosPolicySettingDataEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_NetQosPolicySettingData, err error) {
	tmp, err := NewMSFT_NetSettingDataEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_NetQosPolicySettingData{
		MSFT_NetSettingData: tmp,
	}
	return
}

// SetAppPathNameMatchCondition sets the value of AppPathNameMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyAppPathNameMatchCondition(value string) (err error) {
	return instance.SetProperty("AppPathNameMatchCondition", (value))
}

// GetAppPathNameMatchCondition gets the value of AppPathNameMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyAppPathNameMatchCondition() (value string, err error) {
	retValue, err := instance.GetProperty("AppPathNameMatchCondition")
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

// SetDSCPAction sets the value of DSCPAction for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyDSCPAction(value int8) (err error) {
	return instance.SetProperty("DSCPAction", (value))
}

// GetDSCPAction gets the value of DSCPAction for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyDSCPAction() (value int8, err error) {
	retValue, err := instance.GetProperty("DSCPAction")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int8(valuetmp)

	return
}

// SetIPDstPortEndMatchCondition sets the value of IPDstPortEndMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyIPDstPortEndMatchCondition(value uint16) (err error) {
	return instance.SetProperty("IPDstPortEndMatchCondition", (value))
}

// GetIPDstPortEndMatchCondition gets the value of IPDstPortEndMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyIPDstPortEndMatchCondition() (value uint16, err error) {
	retValue, err := instance.GetProperty("IPDstPortEndMatchCondition")
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

// SetIPDstPortStartMatchCondition sets the value of IPDstPortStartMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyIPDstPortStartMatchCondition(value uint16) (err error) {
	return instance.SetProperty("IPDstPortStartMatchCondition", (value))
}

// GetIPDstPortStartMatchCondition gets the value of IPDstPortStartMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyIPDstPortStartMatchCondition() (value uint16, err error) {
	retValue, err := instance.GetProperty("IPDstPortStartMatchCondition")
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

// SetIPDstPrefixMatchCondition sets the value of IPDstPrefixMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyIPDstPrefixMatchCondition(value string) (err error) {
	return instance.SetProperty("IPDstPrefixMatchCondition", (value))
}

// GetIPDstPrefixMatchCondition gets the value of IPDstPrefixMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyIPDstPrefixMatchCondition() (value string, err error) {
	retValue, err := instance.GetProperty("IPDstPrefixMatchCondition")
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

// SetIPPortMatchCondition sets the value of IPPortMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyIPPortMatchCondition(value uint16) (err error) {
	return instance.SetProperty("IPPortMatchCondition", (value))
}

// GetIPPortMatchCondition gets the value of IPPortMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyIPPortMatchCondition() (value uint16, err error) {
	retValue, err := instance.GetProperty("IPPortMatchCondition")
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

// SetIPProtocolMatchCondition sets the value of IPProtocolMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyIPProtocolMatchCondition(value uint32) (err error) {
	return instance.SetProperty("IPProtocolMatchCondition", (value))
}

// GetIPProtocolMatchCondition gets the value of IPProtocolMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyIPProtocolMatchCondition() (value uint32, err error) {
	retValue, err := instance.GetProperty("IPProtocolMatchCondition")
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

// SetIPSrcPortEndMatchCondition sets the value of IPSrcPortEndMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyIPSrcPortEndMatchCondition(value uint16) (err error) {
	return instance.SetProperty("IPSrcPortEndMatchCondition", (value))
}

// GetIPSrcPortEndMatchCondition gets the value of IPSrcPortEndMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyIPSrcPortEndMatchCondition() (value uint16, err error) {
	retValue, err := instance.GetProperty("IPSrcPortEndMatchCondition")
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

// SetIPSrcPortStartMatchCondition sets the value of IPSrcPortStartMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyIPSrcPortStartMatchCondition(value uint16) (err error) {
	return instance.SetProperty("IPSrcPortStartMatchCondition", (value))
}

// GetIPSrcPortStartMatchCondition gets the value of IPSrcPortStartMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyIPSrcPortStartMatchCondition() (value uint16, err error) {
	retValue, err := instance.GetProperty("IPSrcPortStartMatchCondition")
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

// SetIPSrcPrefixMatchCondition sets the value of IPSrcPrefixMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyIPSrcPrefixMatchCondition(value string) (err error) {
	return instance.SetProperty("IPSrcPrefixMatchCondition", (value))
}

// GetIPSrcPrefixMatchCondition gets the value of IPSrcPrefixMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyIPSrcPrefixMatchCondition() (value string, err error) {
	retValue, err := instance.GetProperty("IPSrcPrefixMatchCondition")
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

// SetJobObjectMatchCondition sets the value of JobObjectMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyJobObjectMatchCondition(value string) (err error) {
	return instance.SetProperty("JobObjectMatchCondition", (value))
}

// GetJobObjectMatchCondition gets the value of JobObjectMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyJobObjectMatchCondition() (value string, err error) {
	retValue, err := instance.GetProperty("JobObjectMatchCondition")
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

// SetMinBandwidthWeightAction sets the value of MinBandwidthWeightAction for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyMinBandwidthWeightAction(value uint8) (err error) {
	return instance.SetProperty("MinBandwidthWeightAction", (value))
}

// GetMinBandwidthWeightAction gets the value of MinBandwidthWeightAction for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyMinBandwidthWeightAction() (value uint8, err error) {
	retValue, err := instance.GetProperty("MinBandwidthWeightAction")
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

// SetName sets the value of Name for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyName(value string) (err error) {
	return instance.SetProperty("Name", (value))
}

// GetName gets the value of Name for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyName() (value string, err error) {
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

// SetNetDirectPortMatchCondition sets the value of NetDirectPortMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyNetDirectPortMatchCondition(value uint16) (err error) {
	return instance.SetProperty("NetDirectPortMatchCondition", (value))
}

// GetNetDirectPortMatchCondition gets the value of NetDirectPortMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyNetDirectPortMatchCondition() (value uint16, err error) {
	retValue, err := instance.GetProperty("NetDirectPortMatchCondition")
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

// SetNetworkProfile sets the value of NetworkProfile for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyNetworkProfile(value uint32) (err error) {
	return instance.SetProperty("NetworkProfile", (value))
}

// GetNetworkProfile gets the value of NetworkProfile for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyNetworkProfile() (value uint32, err error) {
	retValue, err := instance.GetProperty("NetworkProfile")
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

// SetOwner sets the value of Owner for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyOwner(value string) (err error) {
	return instance.SetProperty("Owner", (value))
}

// GetOwner gets the value of Owner for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyOwner() (value string, err error) {
	retValue, err := instance.GetProperty("Owner")
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

// SetPrecedence sets the value of Precedence for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyPrecedence(value uint32) (err error) {
	return instance.SetProperty("Precedence", (value))
}

// GetPrecedence gets the value of Precedence for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyPrecedence() (value uint32, err error) {
	retValue, err := instance.GetProperty("Precedence")
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

// SetPriorityValue8021Action sets the value of PriorityValue8021Action for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyPriorityValue8021Action(value int8) (err error) {
	return instance.SetProperty("PriorityValue8021Action", (value))
}

// GetPriorityValue8021Action gets the value of PriorityValue8021Action for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyPriorityValue8021Action() (value int8, err error) {
	retValue, err := instance.GetProperty("PriorityValue8021Action")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int8)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int8 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int8(valuetmp)

	return
}

// SetTemplateMatchCondition sets the value of TemplateMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyTemplateMatchCondition(value uint32) (err error) {
	return instance.SetProperty("TemplateMatchCondition", (value))
}

// GetTemplateMatchCondition gets the value of TemplateMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyTemplateMatchCondition() (value uint32, err error) {
	retValue, err := instance.GetProperty("TemplateMatchCondition")
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

// SetThrottleRateAction sets the value of ThrottleRateAction for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyThrottleRateAction(value uint64) (err error) {
	return instance.SetProperty("ThrottleRateAction", (value))
}

// GetThrottleRateAction gets the value of ThrottleRateAction for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyThrottleRateAction() (value uint64, err error) {
	retValue, err := instance.GetProperty("ThrottleRateAction")
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

// SetURIMatchCondition sets the value of URIMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyURIMatchCondition(value string) (err error) {
	return instance.SetProperty("URIMatchCondition", (value))
}

// GetURIMatchCondition gets the value of URIMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyURIMatchCondition() (value string, err error) {
	retValue, err := instance.GetProperty("URIMatchCondition")
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

// SetURIRecursiveMatchCondition sets the value of URIRecursiveMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyURIRecursiveMatchCondition(value bool) (err error) {
	return instance.SetProperty("URIRecursiveMatchCondition", (value))
}

// GetURIRecursiveMatchCondition gets the value of URIRecursiveMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyURIRecursiveMatchCondition() (value bool, err error) {
	retValue, err := instance.GetProperty("URIRecursiveMatchCondition")
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

// SetUserMatchCondition sets the value of UserMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyUserMatchCondition(value string) (err error) {
	return instance.SetProperty("UserMatchCondition", (value))
}

// GetUserMatchCondition gets the value of UserMatchCondition for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyUserMatchCondition() (value string, err error) {
	retValue, err := instance.GetProperty("UserMatchCondition")
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

// SetVersion sets the value of Version for the instance
func (instance *MSFT_NetQosPolicySettingData) SetPropertyVersion(value string) (err error) {
	return instance.SetProperty("Version", (value))
}

// GetVersion gets the value of Version for the instance
func (instance *MSFT_NetQosPolicySettingData) GetPropertyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("Version")
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
