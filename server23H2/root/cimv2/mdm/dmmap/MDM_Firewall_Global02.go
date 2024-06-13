// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// MDM_Firewall_Global02 struct
type MDM_Firewall_Global02 struct {
	*cim.WmiInstance

	//
	BinaryVersionSupported string

	//
	CRLcheck int32

	//
	CurrentProfiles int32

	//
	DisableStatefulFtp bool

	//
	EnablePacketQueue int32

	//
	InstanceID string

	//
	IPsecExempt int32

	//
	OpportunisticallyMatchAuthSetPerKM bool

	//
	ParentID string

	//
	PolicyVersion string

	//
	PolicyVersionSupported int32

	//
	PresharedKeyEncoding int32

	//
	SaIdleTime int32
}

func NewMDM_Firewall_Global02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Firewall_Global02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Firewall_Global02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Firewall_Global02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Firewall_Global02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Firewall_Global02{
		WmiInstance: tmp,
	}
	return
}

// SetBinaryVersionSupported sets the value of BinaryVersionSupported for the instance
func (instance *MDM_Firewall_Global02) SetPropertyBinaryVersionSupported(value string) (err error) {
	return instance.SetProperty("BinaryVersionSupported", (value))
}

// GetBinaryVersionSupported gets the value of BinaryVersionSupported for the instance
func (instance *MDM_Firewall_Global02) GetPropertyBinaryVersionSupported() (value string, err error) {
	retValue, err := instance.GetProperty("BinaryVersionSupported")
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

// SetCRLcheck sets the value of CRLcheck for the instance
func (instance *MDM_Firewall_Global02) SetPropertyCRLcheck(value int32) (err error) {
	return instance.SetProperty("CRLcheck", (value))
}

// GetCRLcheck gets the value of CRLcheck for the instance
func (instance *MDM_Firewall_Global02) GetPropertyCRLcheck() (value int32, err error) {
	retValue, err := instance.GetProperty("CRLcheck")
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

// SetCurrentProfiles sets the value of CurrentProfiles for the instance
func (instance *MDM_Firewall_Global02) SetPropertyCurrentProfiles(value int32) (err error) {
	return instance.SetProperty("CurrentProfiles", (value))
}

// GetCurrentProfiles gets the value of CurrentProfiles for the instance
func (instance *MDM_Firewall_Global02) GetPropertyCurrentProfiles() (value int32, err error) {
	retValue, err := instance.GetProperty("CurrentProfiles")
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

// SetDisableStatefulFtp sets the value of DisableStatefulFtp for the instance
func (instance *MDM_Firewall_Global02) SetPropertyDisableStatefulFtp(value bool) (err error) {
	return instance.SetProperty("DisableStatefulFtp", (value))
}

// GetDisableStatefulFtp gets the value of DisableStatefulFtp for the instance
func (instance *MDM_Firewall_Global02) GetPropertyDisableStatefulFtp() (value bool, err error) {
	retValue, err := instance.GetProperty("DisableStatefulFtp")
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

// SetEnablePacketQueue sets the value of EnablePacketQueue for the instance
func (instance *MDM_Firewall_Global02) SetPropertyEnablePacketQueue(value int32) (err error) {
	return instance.SetProperty("EnablePacketQueue", (value))
}

// GetEnablePacketQueue gets the value of EnablePacketQueue for the instance
func (instance *MDM_Firewall_Global02) GetPropertyEnablePacketQueue() (value int32, err error) {
	retValue, err := instance.GetProperty("EnablePacketQueue")
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

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Firewall_Global02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", (value))
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Firewall_Global02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
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

// SetIPsecExempt sets the value of IPsecExempt for the instance
func (instance *MDM_Firewall_Global02) SetPropertyIPsecExempt(value int32) (err error) {
	return instance.SetProperty("IPsecExempt", (value))
}

// GetIPsecExempt gets the value of IPsecExempt for the instance
func (instance *MDM_Firewall_Global02) GetPropertyIPsecExempt() (value int32, err error) {
	retValue, err := instance.GetProperty("IPsecExempt")
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

// SetOpportunisticallyMatchAuthSetPerKM sets the value of OpportunisticallyMatchAuthSetPerKM for the instance
func (instance *MDM_Firewall_Global02) SetPropertyOpportunisticallyMatchAuthSetPerKM(value bool) (err error) {
	return instance.SetProperty("OpportunisticallyMatchAuthSetPerKM", (value))
}

// GetOpportunisticallyMatchAuthSetPerKM gets the value of OpportunisticallyMatchAuthSetPerKM for the instance
func (instance *MDM_Firewall_Global02) GetPropertyOpportunisticallyMatchAuthSetPerKM() (value bool, err error) {
	retValue, err := instance.GetProperty("OpportunisticallyMatchAuthSetPerKM")
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

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Firewall_Global02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", (value))
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Firewall_Global02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
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

// SetPolicyVersion sets the value of PolicyVersion for the instance
func (instance *MDM_Firewall_Global02) SetPropertyPolicyVersion(value string) (err error) {
	return instance.SetProperty("PolicyVersion", (value))
}

// GetPolicyVersion gets the value of PolicyVersion for the instance
func (instance *MDM_Firewall_Global02) GetPropertyPolicyVersion() (value string, err error) {
	retValue, err := instance.GetProperty("PolicyVersion")
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

// SetPolicyVersionSupported sets the value of PolicyVersionSupported for the instance
func (instance *MDM_Firewall_Global02) SetPropertyPolicyVersionSupported(value int32) (err error) {
	return instance.SetProperty("PolicyVersionSupported", (value))
}

// GetPolicyVersionSupported gets the value of PolicyVersionSupported for the instance
func (instance *MDM_Firewall_Global02) GetPropertyPolicyVersionSupported() (value int32, err error) {
	retValue, err := instance.GetProperty("PolicyVersionSupported")
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

// SetPresharedKeyEncoding sets the value of PresharedKeyEncoding for the instance
func (instance *MDM_Firewall_Global02) SetPropertyPresharedKeyEncoding(value int32) (err error) {
	return instance.SetProperty("PresharedKeyEncoding", (value))
}

// GetPresharedKeyEncoding gets the value of PresharedKeyEncoding for the instance
func (instance *MDM_Firewall_Global02) GetPropertyPresharedKeyEncoding() (value int32, err error) {
	retValue, err := instance.GetProperty("PresharedKeyEncoding")
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

// SetSaIdleTime sets the value of SaIdleTime for the instance
func (instance *MDM_Firewall_Global02) SetPropertySaIdleTime(value int32) (err error) {
	return instance.SetProperty("SaIdleTime", (value))
}

// GetSaIdleTime gets the value of SaIdleTime for the instance
func (instance *MDM_Firewall_Global02) GetPropertySaIdleTime() (value int32, err error) {
	retValue, err := instance.GetProperty("SaIdleTime")
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
