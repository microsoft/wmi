// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

// Author:
//
//	Auto Generated on 6/6/2024 using wmigen
//	Source root.directory.LDAP
//
// ////////////////////////////////////////////
package ldap

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// ads_trusteddomain struct
type ads_trusteddomain struct {
	*ds_leaf

	//
	DS_additionalTrustedServiceNames []string

	//
	DS_domainCrossRef string

	//
	DS_domainIdentifier int32

	//
	DS_flatName string

	//
	DS_initialAuthIncoming string

	//
	DS_initialAuthOutgoing string

	//
	DS_mS_DS_CreatorSID Uint8Array

	//
	DS_msDS_EgressClaimsTransformationPolicy string

	//
	DS_msDS_IngressClaimsTransformationPolicy string

	//
	DS_msDS_SupportedEncryptionTypes int32

	//
	DS_msDS_TrustForestTrustInfo Uint8Array

	//
	DS_securityIdentifier Uint8Array

	//
	DS_trustAttributes int32

	//
	DS_trustAuthIncoming Uint8Array

	//
	DS_trustAuthOutgoing Uint8Array

	//
	DS_trustDirection int32

	//
	DS_trustPartner string

	//
	DS_trustPosixOffset int32

	//
	DS_trustType int32
}

func Newads_trusteddomainEx1(instance *cim.WmiInstance) (newInstance *ads_trusteddomain, err error) {
	tmp, err := Newds_leafEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ads_trusteddomain{
		ds_leaf: tmp,
	}
	return
}

func Newads_trusteddomainEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ads_trusteddomain, err error) {
	tmp, err := Newds_leafEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ads_trusteddomain{
		ds_leaf: tmp,
	}
	return
}

// SetDS_additionalTrustedServiceNames sets the value of DS_additionalTrustedServiceNames for the instance
func (instance *ads_trusteddomain) SetPropertyDS_additionalTrustedServiceNames(value []string) (err error) {
	return instance.SetProperty("DS_additionalTrustedServiceNames", (value))
}

// GetDS_additionalTrustedServiceNames gets the value of DS_additionalTrustedServiceNames for the instance
func (instance *ads_trusteddomain) GetPropertyDS_additionalTrustedServiceNames() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_additionalTrustedServiceNames")
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

// SetDS_domainCrossRef sets the value of DS_domainCrossRef for the instance
func (instance *ads_trusteddomain) SetPropertyDS_domainCrossRef(value string) (err error) {
	return instance.SetProperty("DS_domainCrossRef", (value))
}

// GetDS_domainCrossRef gets the value of DS_domainCrossRef for the instance
func (instance *ads_trusteddomain) GetPropertyDS_domainCrossRef() (value string, err error) {
	retValue, err := instance.GetProperty("DS_domainCrossRef")
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

// SetDS_domainIdentifier sets the value of DS_domainIdentifier for the instance
func (instance *ads_trusteddomain) SetPropertyDS_domainIdentifier(value int32) (err error) {
	return instance.SetProperty("DS_domainIdentifier", (value))
}

// GetDS_domainIdentifier gets the value of DS_domainIdentifier for the instance
func (instance *ads_trusteddomain) GetPropertyDS_domainIdentifier() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_domainIdentifier")
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

// SetDS_flatName sets the value of DS_flatName for the instance
func (instance *ads_trusteddomain) SetPropertyDS_flatName(value string) (err error) {
	return instance.SetProperty("DS_flatName", (value))
}

// GetDS_flatName gets the value of DS_flatName for the instance
func (instance *ads_trusteddomain) GetPropertyDS_flatName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_flatName")
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

// SetDS_initialAuthIncoming sets the value of DS_initialAuthIncoming for the instance
func (instance *ads_trusteddomain) SetPropertyDS_initialAuthIncoming(value string) (err error) {
	return instance.SetProperty("DS_initialAuthIncoming", (value))
}

// GetDS_initialAuthIncoming gets the value of DS_initialAuthIncoming for the instance
func (instance *ads_trusteddomain) GetPropertyDS_initialAuthIncoming() (value string, err error) {
	retValue, err := instance.GetProperty("DS_initialAuthIncoming")
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

// SetDS_initialAuthOutgoing sets the value of DS_initialAuthOutgoing for the instance
func (instance *ads_trusteddomain) SetPropertyDS_initialAuthOutgoing(value string) (err error) {
	return instance.SetProperty("DS_initialAuthOutgoing", (value))
}

// GetDS_initialAuthOutgoing gets the value of DS_initialAuthOutgoing for the instance
func (instance *ads_trusteddomain) GetPropertyDS_initialAuthOutgoing() (value string, err error) {
	retValue, err := instance.GetProperty("DS_initialAuthOutgoing")
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

// SetDS_mS_DS_CreatorSID sets the value of DS_mS_DS_CreatorSID for the instance
func (instance *ads_trusteddomain) SetPropertyDS_mS_DS_CreatorSID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_mS_DS_CreatorSID", (value))
}

// GetDS_mS_DS_CreatorSID gets the value of DS_mS_DS_CreatorSID for the instance
func (instance *ads_trusteddomain) GetPropertyDS_mS_DS_CreatorSID() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_mS_DS_CreatorSID")
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

// SetDS_msDS_EgressClaimsTransformationPolicy sets the value of DS_msDS_EgressClaimsTransformationPolicy for the instance
func (instance *ads_trusteddomain) SetPropertyDS_msDS_EgressClaimsTransformationPolicy(value string) (err error) {
	return instance.SetProperty("DS_msDS_EgressClaimsTransformationPolicy", (value))
}

// GetDS_msDS_EgressClaimsTransformationPolicy gets the value of DS_msDS_EgressClaimsTransformationPolicy for the instance
func (instance *ads_trusteddomain) GetPropertyDS_msDS_EgressClaimsTransformationPolicy() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_EgressClaimsTransformationPolicy")
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

// SetDS_msDS_IngressClaimsTransformationPolicy sets the value of DS_msDS_IngressClaimsTransformationPolicy for the instance
func (instance *ads_trusteddomain) SetPropertyDS_msDS_IngressClaimsTransformationPolicy(value string) (err error) {
	return instance.SetProperty("DS_msDS_IngressClaimsTransformationPolicy", (value))
}

// GetDS_msDS_IngressClaimsTransformationPolicy gets the value of DS_msDS_IngressClaimsTransformationPolicy for the instance
func (instance *ads_trusteddomain) GetPropertyDS_msDS_IngressClaimsTransformationPolicy() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_IngressClaimsTransformationPolicy")
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

// SetDS_msDS_SupportedEncryptionTypes sets the value of DS_msDS_SupportedEncryptionTypes for the instance
func (instance *ads_trusteddomain) SetPropertyDS_msDS_SupportedEncryptionTypes(value int32) (err error) {
	return instance.SetProperty("DS_msDS_SupportedEncryptionTypes", (value))
}

// GetDS_msDS_SupportedEncryptionTypes gets the value of DS_msDS_SupportedEncryptionTypes for the instance
func (instance *ads_trusteddomain) GetPropertyDS_msDS_SupportedEncryptionTypes() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_SupportedEncryptionTypes")
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

// SetDS_msDS_TrustForestTrustInfo sets the value of DS_msDS_TrustForestTrustInfo for the instance
func (instance *ads_trusteddomain) SetPropertyDS_msDS_TrustForestTrustInfo(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_TrustForestTrustInfo", (value))
}

// GetDS_msDS_TrustForestTrustInfo gets the value of DS_msDS_TrustForestTrustInfo for the instance
func (instance *ads_trusteddomain) GetPropertyDS_msDS_TrustForestTrustInfo() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_TrustForestTrustInfo")
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

// SetDS_securityIdentifier sets the value of DS_securityIdentifier for the instance
func (instance *ads_trusteddomain) SetPropertyDS_securityIdentifier(value Uint8Array) (err error) {
	return instance.SetProperty("DS_securityIdentifier", (value))
}

// GetDS_securityIdentifier gets the value of DS_securityIdentifier for the instance
func (instance *ads_trusteddomain) GetPropertyDS_securityIdentifier() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_securityIdentifier")
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

// SetDS_trustAttributes sets the value of DS_trustAttributes for the instance
func (instance *ads_trusteddomain) SetPropertyDS_trustAttributes(value int32) (err error) {
	return instance.SetProperty("DS_trustAttributes", (value))
}

// GetDS_trustAttributes gets the value of DS_trustAttributes for the instance
func (instance *ads_trusteddomain) GetPropertyDS_trustAttributes() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_trustAttributes")
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

// SetDS_trustAuthIncoming sets the value of DS_trustAuthIncoming for the instance
func (instance *ads_trusteddomain) SetPropertyDS_trustAuthIncoming(value Uint8Array) (err error) {
	return instance.SetProperty("DS_trustAuthIncoming", (value))
}

// GetDS_trustAuthIncoming gets the value of DS_trustAuthIncoming for the instance
func (instance *ads_trusteddomain) GetPropertyDS_trustAuthIncoming() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_trustAuthIncoming")
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

// SetDS_trustAuthOutgoing sets the value of DS_trustAuthOutgoing for the instance
func (instance *ads_trusteddomain) SetPropertyDS_trustAuthOutgoing(value Uint8Array) (err error) {
	return instance.SetProperty("DS_trustAuthOutgoing", (value))
}

// GetDS_trustAuthOutgoing gets the value of DS_trustAuthOutgoing for the instance
func (instance *ads_trusteddomain) GetPropertyDS_trustAuthOutgoing() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_trustAuthOutgoing")
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

// SetDS_trustDirection sets the value of DS_trustDirection for the instance
func (instance *ads_trusteddomain) SetPropertyDS_trustDirection(value int32) (err error) {
	return instance.SetProperty("DS_trustDirection", (value))
}

// GetDS_trustDirection gets the value of DS_trustDirection for the instance
func (instance *ads_trusteddomain) GetPropertyDS_trustDirection() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_trustDirection")
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

// SetDS_trustPartner sets the value of DS_trustPartner for the instance
func (instance *ads_trusteddomain) SetPropertyDS_trustPartner(value string) (err error) {
	return instance.SetProperty("DS_trustPartner", (value))
}

// GetDS_trustPartner gets the value of DS_trustPartner for the instance
func (instance *ads_trusteddomain) GetPropertyDS_trustPartner() (value string, err error) {
	retValue, err := instance.GetProperty("DS_trustPartner")
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

// SetDS_trustPosixOffset sets the value of DS_trustPosixOffset for the instance
func (instance *ads_trusteddomain) SetPropertyDS_trustPosixOffset(value int32) (err error) {
	return instance.SetProperty("DS_trustPosixOffset", (value))
}

// GetDS_trustPosixOffset gets the value of DS_trustPosixOffset for the instance
func (instance *ads_trusteddomain) GetPropertyDS_trustPosixOffset() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_trustPosixOffset")
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

// SetDS_trustType sets the value of DS_trustType for the instance
func (instance *ads_trusteddomain) SetPropertyDS_trustType(value int32) (err error) {
	return instance.SetProperty("DS_trustType", (value))
}

// GetDS_trustType gets the value of DS_trustType for the instance
func (instance *ads_trusteddomain) GetPropertyDS_trustType() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_trustType")
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
