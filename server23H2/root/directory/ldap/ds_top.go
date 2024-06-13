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

// ds_top struct
type ds_top struct {
	*DS_LDAP_Root_Class

	//
	DS_adminDescription string

	//
	DS_adminDisplayName string

	//
	DS_allowedAttributes []string

	//
	DS_allowedAttributesEffective []string

	//
	DS_allowedChildClasses []string

	//
	DS_allowedChildClassesEffective []string

	//
	DS_bridgeheadServerListBL []string

	//
	DS_canonicalName []string

	//
	DS_cn string

	//
	DS_createTimeStamp string

	//
	DS_description []string

	//
	DS_directReports []string

	//
	DS_displayName string

	//
	DS_displayNamePrintable string

	//
	DS_distinguishedName string

	//
	DS_dSASignature Uint8Array

	//
	DS_dSCorePropagationData []string

	//
	DS_extensionName []string

	//
	DS_flags int32

	//
	DS_fromEntry []bool

	//
	DS_frsComputerReferenceBL []string

	//
	DS_fRSMemberReferenceBL []string

	//
	DS_fSMORoleOwner string

	//
	DS_instanceType int32

	//
	DS_isCriticalSystemObject bool

	//
	DS_isDeleted bool

	//
	DS_isPrivilegeHolder []string

	//
	DS_isRecycled bool

	//
	DS_lastKnownParent string

	//
	DS_managedObjects []string

	//
	DS_masteredBy []string

	//
	DS_memberOf []string

	//
	DS_modifyTimeStamp string

	//
	DS_mS_DS_ConsistencyChildCount int32

	//
	DS_mS_DS_ConsistencyGuid Uint8Array

	//
	DS_msCOM_PartitionSetLink []string

	//
	DS_msCOM_UserLink []string

	//
	DS_msDFSR_ComputerReferenceBL []string

	//
	DS_msDFSR_MemberReferenceBL []string

	//
	DS_msDS_Approx_Immed_Subordinates int32

	//
	DS_msDS_AuthenticatedToAccountlist []string

	//
	DS_msDS_ClaimSharesPossibleValuesWithBL []string

	//
	DS_msDS_CloudAnchor Uint8Array

	//
	DS_msDS_EnabledFeatureBL []string

	//
	DS_msDS_HostServiceAccountBL []string

	//
	DS_msDS_IsDomainFor []string

	//
	DS_msDS_IsFullReplicaFor []string

	//
	DS_msDS_IsPartialReplicaFor []string

	//
	DS_msDS_IsPrimaryComputerFor []string

	//
	DS_msDS_KrbTgtLinkBl []string

	//
	DS_msDS_LastKnownRDN string

	//
	DS_msDS_LocalEffectiveDeletionTime string

	//
	DS_msDS_LocalEffectiveRecycleTime string

	//
	DS_msDs_masteredBy []string

	//
	DS_msds_memberOfTransitive []string

	//
	DS_msDS_MembersForAzRoleBL []string

	//
	DS_msDS_MembersOfResourcePropertyListBL []string

	//
	DS_msds_memberTransitive []string

	//
	DS_msDS_NC_RO_Replica_Locations_BL []string

	//
	DS_msDS_NCReplCursors []string

	//
	DS_msDS_NCReplInboundNeighbors []string

	//
	DS_msDS_NCReplOutboundNeighbors []string

	//
	DS_msDS_NcType int32

	//
	DS_msDS_NonMembersBL []string

	//
	DS_msDS_ObjectReferenceBL []string

	//
	DS_msDS_ObjectSoa string

	//
	DS_msDS_OIDToGroupLinkBl []string

	//
	DS_msDS_OperationsForAzRoleBL []string

	//
	DS_msDS_OperationsForAzTaskBL []string

	//
	DS_msDS_parentdistname string

	//
	DS_msDS_PrincipalName string

	//
	DS_msDS_PSOApplied []string

	//
	DS_msDS_ReplAttributeMetaData []string

	//
	DS_msDS_ReplValueMetaData []string

	//
	DS_msDS_ReplValueMetaDataExt []string

	//
	DS_msDS_RevealedDSAs []string

	//
	DS_msDS_RevealedListBL []string

	//
	DS_msDS_SourceAnchor string

	//
	DS_msDS_TasksForAzRoleBL []string

	//
	DS_msDS_TasksForAzTaskBL []string

	//
	DS_msDS_TDOEgressBL []string

	//
	DS_msDS_TDOIngressBL []string

	//
	DS_msDS_ValueTypeReferenceBL []string

	//
	DS_msSFU30PosixMemberOf []string

	//
	DS_name string

	//
	DS_netbootSCPBL []string

	//
	DS_nonSecurityMemberBL []string

	//
	DS_nTSecurityDescriptor Uint8Array

	//
	DS_objectCategory string

	//
	DS_objectClass []string

	//
	DS_objectGUID Uint8Array

	//
	DS_objectVersion int32

	//
	DS_otherWellKnownObjects []DN_With_Binary

	//
	DS_ownerBL []string

	//
	DS_partialAttributeDeletionList Uint8Array

	//
	DS_partialAttributeSet Uint8Array

	//
	DS_possibleInferiors []string

	//
	DS_proxiedObjectName DN_With_Binary

	//
	DS_proxyAddresses []string

	//
	DS_queryPolicyBL []string

	//
	DS_replPropertyMetaData Uint8Array

	//
	DS_replUpToDateVector Uint8Array

	//
	DS_repsFrom []Uint8Array

	//
	DS_repsTo []Uint8Array

	//
	DS_revision int32

	//
	DS_sDRightsEffective int32

	//
	DS_serverReferenceBL []string

	//
	DS_showInAdvancedViewOnly bool

	//
	DS_siteObjectBL []string

	//
	DS_structuralObjectClass []string

	//
	DS_subRefs []string

	//
	DS_subSchemaSubEntry []string

	//
	DS_systemFlags int32

	//
	DS_url []string

	//
	DS_uSNChanged int64

	//
	DS_uSNCreated int64

	//
	DS_uSNDSALastObjRemoved int64

	//
	DS_USNIntersite int32

	//
	DS_uSNLastObjRem int64

	//
	DS_uSNSource int64

	//
	DS_wbemPath []string

	//
	DS_wellKnownObjects []DN_With_Binary

	//
	DS_whenChanged string

	//
	DS_whenCreated string

	//
	DS_wWWHomePage string
}

func Newds_topEx1(instance *cim.WmiInstance) (newInstance *ds_top, err error) {
	tmp, err := NewDS_LDAP_Root_ClassEx1(instance)

	if err != nil {
		return
	}
	newInstance = &ds_top{
		DS_LDAP_Root_Class: tmp,
	}
	return
}

func Newds_topEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *ds_top, err error) {
	tmp, err := NewDS_LDAP_Root_ClassEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &ds_top{
		DS_LDAP_Root_Class: tmp,
	}
	return
}

// SetDS_adminDescription sets the value of DS_adminDescription for the instance
func (instance *ds_top) SetPropertyDS_adminDescription(value string) (err error) {
	return instance.SetProperty("DS_adminDescription", (value))
}

// GetDS_adminDescription gets the value of DS_adminDescription for the instance
func (instance *ds_top) GetPropertyDS_adminDescription() (value string, err error) {
	retValue, err := instance.GetProperty("DS_adminDescription")
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

// SetDS_adminDisplayName sets the value of DS_adminDisplayName for the instance
func (instance *ds_top) SetPropertyDS_adminDisplayName(value string) (err error) {
	return instance.SetProperty("DS_adminDisplayName", (value))
}

// GetDS_adminDisplayName gets the value of DS_adminDisplayName for the instance
func (instance *ds_top) GetPropertyDS_adminDisplayName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_adminDisplayName")
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

// SetDS_allowedAttributes sets the value of DS_allowedAttributes for the instance
func (instance *ds_top) SetPropertyDS_allowedAttributes(value []string) (err error) {
	return instance.SetProperty("DS_allowedAttributes", (value))
}

// GetDS_allowedAttributes gets the value of DS_allowedAttributes for the instance
func (instance *ds_top) GetPropertyDS_allowedAttributes() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_allowedAttributes")
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

// SetDS_allowedAttributesEffective sets the value of DS_allowedAttributesEffective for the instance
func (instance *ds_top) SetPropertyDS_allowedAttributesEffective(value []string) (err error) {
	return instance.SetProperty("DS_allowedAttributesEffective", (value))
}

// GetDS_allowedAttributesEffective gets the value of DS_allowedAttributesEffective for the instance
func (instance *ds_top) GetPropertyDS_allowedAttributesEffective() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_allowedAttributesEffective")
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

// SetDS_allowedChildClasses sets the value of DS_allowedChildClasses for the instance
func (instance *ds_top) SetPropertyDS_allowedChildClasses(value []string) (err error) {
	return instance.SetProperty("DS_allowedChildClasses", (value))
}

// GetDS_allowedChildClasses gets the value of DS_allowedChildClasses for the instance
func (instance *ds_top) GetPropertyDS_allowedChildClasses() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_allowedChildClasses")
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

// SetDS_allowedChildClassesEffective sets the value of DS_allowedChildClassesEffective for the instance
func (instance *ds_top) SetPropertyDS_allowedChildClassesEffective(value []string) (err error) {
	return instance.SetProperty("DS_allowedChildClassesEffective", (value))
}

// GetDS_allowedChildClassesEffective gets the value of DS_allowedChildClassesEffective for the instance
func (instance *ds_top) GetPropertyDS_allowedChildClassesEffective() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_allowedChildClassesEffective")
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

// SetDS_bridgeheadServerListBL sets the value of DS_bridgeheadServerListBL for the instance
func (instance *ds_top) SetPropertyDS_bridgeheadServerListBL(value []string) (err error) {
	return instance.SetProperty("DS_bridgeheadServerListBL", (value))
}

// GetDS_bridgeheadServerListBL gets the value of DS_bridgeheadServerListBL for the instance
func (instance *ds_top) GetPropertyDS_bridgeheadServerListBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_bridgeheadServerListBL")
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

// SetDS_canonicalName sets the value of DS_canonicalName for the instance
func (instance *ds_top) SetPropertyDS_canonicalName(value []string) (err error) {
	return instance.SetProperty("DS_canonicalName", (value))
}

// GetDS_canonicalName gets the value of DS_canonicalName for the instance
func (instance *ds_top) GetPropertyDS_canonicalName() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_canonicalName")
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

// SetDS_cn sets the value of DS_cn for the instance
func (instance *ds_top) SetPropertyDS_cn(value string) (err error) {
	return instance.SetProperty("DS_cn", (value))
}

// GetDS_cn gets the value of DS_cn for the instance
func (instance *ds_top) GetPropertyDS_cn() (value string, err error) {
	retValue, err := instance.GetProperty("DS_cn")
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

// SetDS_createTimeStamp sets the value of DS_createTimeStamp for the instance
func (instance *ds_top) SetPropertyDS_createTimeStamp(value string) (err error) {
	return instance.SetProperty("DS_createTimeStamp", (value))
}

// GetDS_createTimeStamp gets the value of DS_createTimeStamp for the instance
func (instance *ds_top) GetPropertyDS_createTimeStamp() (value string, err error) {
	retValue, err := instance.GetProperty("DS_createTimeStamp")
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

// SetDS_description sets the value of DS_description for the instance
func (instance *ds_top) SetPropertyDS_description(value []string) (err error) {
	return instance.SetProperty("DS_description", (value))
}

// GetDS_description gets the value of DS_description for the instance
func (instance *ds_top) GetPropertyDS_description() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_description")
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

// SetDS_directReports sets the value of DS_directReports for the instance
func (instance *ds_top) SetPropertyDS_directReports(value []string) (err error) {
	return instance.SetProperty("DS_directReports", (value))
}

// GetDS_directReports gets the value of DS_directReports for the instance
func (instance *ds_top) GetPropertyDS_directReports() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_directReports")
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

// SetDS_displayName sets the value of DS_displayName for the instance
func (instance *ds_top) SetPropertyDS_displayName(value string) (err error) {
	return instance.SetProperty("DS_displayName", (value))
}

// GetDS_displayName gets the value of DS_displayName for the instance
func (instance *ds_top) GetPropertyDS_displayName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_displayName")
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

// SetDS_displayNamePrintable sets the value of DS_displayNamePrintable for the instance
func (instance *ds_top) SetPropertyDS_displayNamePrintable(value string) (err error) {
	return instance.SetProperty("DS_displayNamePrintable", (value))
}

// GetDS_displayNamePrintable gets the value of DS_displayNamePrintable for the instance
func (instance *ds_top) GetPropertyDS_displayNamePrintable() (value string, err error) {
	retValue, err := instance.GetProperty("DS_displayNamePrintable")
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

// SetDS_distinguishedName sets the value of DS_distinguishedName for the instance
func (instance *ds_top) SetPropertyDS_distinguishedName(value string) (err error) {
	return instance.SetProperty("DS_distinguishedName", (value))
}

// GetDS_distinguishedName gets the value of DS_distinguishedName for the instance
func (instance *ds_top) GetPropertyDS_distinguishedName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_distinguishedName")
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

// SetDS_dSASignature sets the value of DS_dSASignature for the instance
func (instance *ds_top) SetPropertyDS_dSASignature(value Uint8Array) (err error) {
	return instance.SetProperty("DS_dSASignature", (value))
}

// GetDS_dSASignature gets the value of DS_dSASignature for the instance
func (instance *ds_top) GetPropertyDS_dSASignature() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_dSASignature")
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

// SetDS_dSCorePropagationData sets the value of DS_dSCorePropagationData for the instance
func (instance *ds_top) SetPropertyDS_dSCorePropagationData(value []string) (err error) {
	return instance.SetProperty("DS_dSCorePropagationData", (value))
}

// GetDS_dSCorePropagationData gets the value of DS_dSCorePropagationData for the instance
func (instance *ds_top) GetPropertyDS_dSCorePropagationData() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_dSCorePropagationData")
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

// SetDS_extensionName sets the value of DS_extensionName for the instance
func (instance *ds_top) SetPropertyDS_extensionName(value []string) (err error) {
	return instance.SetProperty("DS_extensionName", (value))
}

// GetDS_extensionName gets the value of DS_extensionName for the instance
func (instance *ds_top) GetPropertyDS_extensionName() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_extensionName")
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

// SetDS_flags sets the value of DS_flags for the instance
func (instance *ds_top) SetPropertyDS_flags(value int32) (err error) {
	return instance.SetProperty("DS_flags", (value))
}

// GetDS_flags gets the value of DS_flags for the instance
func (instance *ds_top) GetPropertyDS_flags() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_flags")
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

// SetDS_fromEntry sets the value of DS_fromEntry for the instance
func (instance *ds_top) SetPropertyDS_fromEntry(value []bool) (err error) {
	return instance.SetProperty("DS_fromEntry", (value))
}

// GetDS_fromEntry gets the value of DS_fromEntry for the instance
func (instance *ds_top) GetPropertyDS_fromEntry() (value []bool, err error) {
	retValue, err := instance.GetProperty("DS_fromEntry")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(bool)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " bool is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, bool(valuetmp))
	}

	return
}

// SetDS_frsComputerReferenceBL sets the value of DS_frsComputerReferenceBL for the instance
func (instance *ds_top) SetPropertyDS_frsComputerReferenceBL(value []string) (err error) {
	return instance.SetProperty("DS_frsComputerReferenceBL", (value))
}

// GetDS_frsComputerReferenceBL gets the value of DS_frsComputerReferenceBL for the instance
func (instance *ds_top) GetPropertyDS_frsComputerReferenceBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_frsComputerReferenceBL")
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

// SetDS_fRSMemberReferenceBL sets the value of DS_fRSMemberReferenceBL for the instance
func (instance *ds_top) SetPropertyDS_fRSMemberReferenceBL(value []string) (err error) {
	return instance.SetProperty("DS_fRSMemberReferenceBL", (value))
}

// GetDS_fRSMemberReferenceBL gets the value of DS_fRSMemberReferenceBL for the instance
func (instance *ds_top) GetPropertyDS_fRSMemberReferenceBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_fRSMemberReferenceBL")
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

// SetDS_fSMORoleOwner sets the value of DS_fSMORoleOwner for the instance
func (instance *ds_top) SetPropertyDS_fSMORoleOwner(value string) (err error) {
	return instance.SetProperty("DS_fSMORoleOwner", (value))
}

// GetDS_fSMORoleOwner gets the value of DS_fSMORoleOwner for the instance
func (instance *ds_top) GetPropertyDS_fSMORoleOwner() (value string, err error) {
	retValue, err := instance.GetProperty("DS_fSMORoleOwner")
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

// SetDS_instanceType sets the value of DS_instanceType for the instance
func (instance *ds_top) SetPropertyDS_instanceType(value int32) (err error) {
	return instance.SetProperty("DS_instanceType", (value))
}

// GetDS_instanceType gets the value of DS_instanceType for the instance
func (instance *ds_top) GetPropertyDS_instanceType() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_instanceType")
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

// SetDS_isCriticalSystemObject sets the value of DS_isCriticalSystemObject for the instance
func (instance *ds_top) SetPropertyDS_isCriticalSystemObject(value bool) (err error) {
	return instance.SetProperty("DS_isCriticalSystemObject", (value))
}

// GetDS_isCriticalSystemObject gets the value of DS_isCriticalSystemObject for the instance
func (instance *ds_top) GetPropertyDS_isCriticalSystemObject() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_isCriticalSystemObject")
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

// SetDS_isDeleted sets the value of DS_isDeleted for the instance
func (instance *ds_top) SetPropertyDS_isDeleted(value bool) (err error) {
	return instance.SetProperty("DS_isDeleted", (value))
}

// GetDS_isDeleted gets the value of DS_isDeleted for the instance
func (instance *ds_top) GetPropertyDS_isDeleted() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_isDeleted")
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

// SetDS_isPrivilegeHolder sets the value of DS_isPrivilegeHolder for the instance
func (instance *ds_top) SetPropertyDS_isPrivilegeHolder(value []string) (err error) {
	return instance.SetProperty("DS_isPrivilegeHolder", (value))
}

// GetDS_isPrivilegeHolder gets the value of DS_isPrivilegeHolder for the instance
func (instance *ds_top) GetPropertyDS_isPrivilegeHolder() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_isPrivilegeHolder")
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

// SetDS_isRecycled sets the value of DS_isRecycled for the instance
func (instance *ds_top) SetPropertyDS_isRecycled(value bool) (err error) {
	return instance.SetProperty("DS_isRecycled", (value))
}

// GetDS_isRecycled gets the value of DS_isRecycled for the instance
func (instance *ds_top) GetPropertyDS_isRecycled() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_isRecycled")
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

// SetDS_lastKnownParent sets the value of DS_lastKnownParent for the instance
func (instance *ds_top) SetPropertyDS_lastKnownParent(value string) (err error) {
	return instance.SetProperty("DS_lastKnownParent", (value))
}

// GetDS_lastKnownParent gets the value of DS_lastKnownParent for the instance
func (instance *ds_top) GetPropertyDS_lastKnownParent() (value string, err error) {
	retValue, err := instance.GetProperty("DS_lastKnownParent")
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

// SetDS_managedObjects sets the value of DS_managedObjects for the instance
func (instance *ds_top) SetPropertyDS_managedObjects(value []string) (err error) {
	return instance.SetProperty("DS_managedObjects", (value))
}

// GetDS_managedObjects gets the value of DS_managedObjects for the instance
func (instance *ds_top) GetPropertyDS_managedObjects() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_managedObjects")
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

// SetDS_masteredBy sets the value of DS_masteredBy for the instance
func (instance *ds_top) SetPropertyDS_masteredBy(value []string) (err error) {
	return instance.SetProperty("DS_masteredBy", (value))
}

// GetDS_masteredBy gets the value of DS_masteredBy for the instance
func (instance *ds_top) GetPropertyDS_masteredBy() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_masteredBy")
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

// SetDS_memberOf sets the value of DS_memberOf for the instance
func (instance *ds_top) SetPropertyDS_memberOf(value []string) (err error) {
	return instance.SetProperty("DS_memberOf", (value))
}

// GetDS_memberOf gets the value of DS_memberOf for the instance
func (instance *ds_top) GetPropertyDS_memberOf() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_memberOf")
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

// SetDS_modifyTimeStamp sets the value of DS_modifyTimeStamp for the instance
func (instance *ds_top) SetPropertyDS_modifyTimeStamp(value string) (err error) {
	return instance.SetProperty("DS_modifyTimeStamp", (value))
}

// GetDS_modifyTimeStamp gets the value of DS_modifyTimeStamp for the instance
func (instance *ds_top) GetPropertyDS_modifyTimeStamp() (value string, err error) {
	retValue, err := instance.GetProperty("DS_modifyTimeStamp")
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

// SetDS_mS_DS_ConsistencyChildCount sets the value of DS_mS_DS_ConsistencyChildCount for the instance
func (instance *ds_top) SetPropertyDS_mS_DS_ConsistencyChildCount(value int32) (err error) {
	return instance.SetProperty("DS_mS_DS_ConsistencyChildCount", (value))
}

// GetDS_mS_DS_ConsistencyChildCount gets the value of DS_mS_DS_ConsistencyChildCount for the instance
func (instance *ds_top) GetPropertyDS_mS_DS_ConsistencyChildCount() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_mS_DS_ConsistencyChildCount")
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

// SetDS_mS_DS_ConsistencyGuid sets the value of DS_mS_DS_ConsistencyGuid for the instance
func (instance *ds_top) SetPropertyDS_mS_DS_ConsistencyGuid(value Uint8Array) (err error) {
	return instance.SetProperty("DS_mS_DS_ConsistencyGuid", (value))
}

// GetDS_mS_DS_ConsistencyGuid gets the value of DS_mS_DS_ConsistencyGuid for the instance
func (instance *ds_top) GetPropertyDS_mS_DS_ConsistencyGuid() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_mS_DS_ConsistencyGuid")
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

// SetDS_msCOM_PartitionSetLink sets the value of DS_msCOM_PartitionSetLink for the instance
func (instance *ds_top) SetPropertyDS_msCOM_PartitionSetLink(value []string) (err error) {
	return instance.SetProperty("DS_msCOM_PartitionSetLink", (value))
}

// GetDS_msCOM_PartitionSetLink gets the value of DS_msCOM_PartitionSetLink for the instance
func (instance *ds_top) GetPropertyDS_msCOM_PartitionSetLink() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msCOM_PartitionSetLink")
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

// SetDS_msCOM_UserLink sets the value of DS_msCOM_UserLink for the instance
func (instance *ds_top) SetPropertyDS_msCOM_UserLink(value []string) (err error) {
	return instance.SetProperty("DS_msCOM_UserLink", (value))
}

// GetDS_msCOM_UserLink gets the value of DS_msCOM_UserLink for the instance
func (instance *ds_top) GetPropertyDS_msCOM_UserLink() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msCOM_UserLink")
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

// SetDS_msDFSR_ComputerReferenceBL sets the value of DS_msDFSR_ComputerReferenceBL for the instance
func (instance *ds_top) SetPropertyDS_msDFSR_ComputerReferenceBL(value []string) (err error) {
	return instance.SetProperty("DS_msDFSR_ComputerReferenceBL", (value))
}

// GetDS_msDFSR_ComputerReferenceBL gets the value of DS_msDFSR_ComputerReferenceBL for the instance
func (instance *ds_top) GetPropertyDS_msDFSR_ComputerReferenceBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_ComputerReferenceBL")
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

// SetDS_msDFSR_MemberReferenceBL sets the value of DS_msDFSR_MemberReferenceBL for the instance
func (instance *ds_top) SetPropertyDS_msDFSR_MemberReferenceBL(value []string) (err error) {
	return instance.SetProperty("DS_msDFSR_MemberReferenceBL", (value))
}

// GetDS_msDFSR_MemberReferenceBL gets the value of DS_msDFSR_MemberReferenceBL for the instance
func (instance *ds_top) GetPropertyDS_msDFSR_MemberReferenceBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDFSR_MemberReferenceBL")
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

// SetDS_msDS_Approx_Immed_Subordinates sets the value of DS_msDS_Approx_Immed_Subordinates for the instance
func (instance *ds_top) SetPropertyDS_msDS_Approx_Immed_Subordinates(value int32) (err error) {
	return instance.SetProperty("DS_msDS_Approx_Immed_Subordinates", (value))
}

// GetDS_msDS_Approx_Immed_Subordinates gets the value of DS_msDS_Approx_Immed_Subordinates for the instance
func (instance *ds_top) GetPropertyDS_msDS_Approx_Immed_Subordinates() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_Approx_Immed_Subordinates")
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

// SetDS_msDS_AuthenticatedToAccountlist sets the value of DS_msDS_AuthenticatedToAccountlist for the instance
func (instance *ds_top) SetPropertyDS_msDS_AuthenticatedToAccountlist(value []string) (err error) {
	return instance.SetProperty("DS_msDS_AuthenticatedToAccountlist", (value))
}

// GetDS_msDS_AuthenticatedToAccountlist gets the value of DS_msDS_AuthenticatedToAccountlist for the instance
func (instance *ds_top) GetPropertyDS_msDS_AuthenticatedToAccountlist() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_AuthenticatedToAccountlist")
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

// SetDS_msDS_ClaimSharesPossibleValuesWithBL sets the value of DS_msDS_ClaimSharesPossibleValuesWithBL for the instance
func (instance *ds_top) SetPropertyDS_msDS_ClaimSharesPossibleValuesWithBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_ClaimSharesPossibleValuesWithBL", (value))
}

// GetDS_msDS_ClaimSharesPossibleValuesWithBL gets the value of DS_msDS_ClaimSharesPossibleValuesWithBL for the instance
func (instance *ds_top) GetPropertyDS_msDS_ClaimSharesPossibleValuesWithBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ClaimSharesPossibleValuesWithBL")
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

// SetDS_msDS_CloudAnchor sets the value of DS_msDS_CloudAnchor for the instance
func (instance *ds_top) SetPropertyDS_msDS_CloudAnchor(value Uint8Array) (err error) {
	return instance.SetProperty("DS_msDS_CloudAnchor", (value))
}

// GetDS_msDS_CloudAnchor gets the value of DS_msDS_CloudAnchor for the instance
func (instance *ds_top) GetPropertyDS_msDS_CloudAnchor() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_msDS_CloudAnchor")
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

// SetDS_msDS_EnabledFeatureBL sets the value of DS_msDS_EnabledFeatureBL for the instance
func (instance *ds_top) SetPropertyDS_msDS_EnabledFeatureBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_EnabledFeatureBL", (value))
}

// GetDS_msDS_EnabledFeatureBL gets the value of DS_msDS_EnabledFeatureBL for the instance
func (instance *ds_top) GetPropertyDS_msDS_EnabledFeatureBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_EnabledFeatureBL")
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

// SetDS_msDS_HostServiceAccountBL sets the value of DS_msDS_HostServiceAccountBL for the instance
func (instance *ds_top) SetPropertyDS_msDS_HostServiceAccountBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_HostServiceAccountBL", (value))
}

// GetDS_msDS_HostServiceAccountBL gets the value of DS_msDS_HostServiceAccountBL for the instance
func (instance *ds_top) GetPropertyDS_msDS_HostServiceAccountBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_HostServiceAccountBL")
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

// SetDS_msDS_IsDomainFor sets the value of DS_msDS_IsDomainFor for the instance
func (instance *ds_top) SetPropertyDS_msDS_IsDomainFor(value []string) (err error) {
	return instance.SetProperty("DS_msDS_IsDomainFor", (value))
}

// GetDS_msDS_IsDomainFor gets the value of DS_msDS_IsDomainFor for the instance
func (instance *ds_top) GetPropertyDS_msDS_IsDomainFor() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_IsDomainFor")
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

// SetDS_msDS_IsFullReplicaFor sets the value of DS_msDS_IsFullReplicaFor for the instance
func (instance *ds_top) SetPropertyDS_msDS_IsFullReplicaFor(value []string) (err error) {
	return instance.SetProperty("DS_msDS_IsFullReplicaFor", (value))
}

// GetDS_msDS_IsFullReplicaFor gets the value of DS_msDS_IsFullReplicaFor for the instance
func (instance *ds_top) GetPropertyDS_msDS_IsFullReplicaFor() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_IsFullReplicaFor")
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

// SetDS_msDS_IsPartialReplicaFor sets the value of DS_msDS_IsPartialReplicaFor for the instance
func (instance *ds_top) SetPropertyDS_msDS_IsPartialReplicaFor(value []string) (err error) {
	return instance.SetProperty("DS_msDS_IsPartialReplicaFor", (value))
}

// GetDS_msDS_IsPartialReplicaFor gets the value of DS_msDS_IsPartialReplicaFor for the instance
func (instance *ds_top) GetPropertyDS_msDS_IsPartialReplicaFor() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_IsPartialReplicaFor")
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

// SetDS_msDS_IsPrimaryComputerFor sets the value of DS_msDS_IsPrimaryComputerFor for the instance
func (instance *ds_top) SetPropertyDS_msDS_IsPrimaryComputerFor(value []string) (err error) {
	return instance.SetProperty("DS_msDS_IsPrimaryComputerFor", (value))
}

// GetDS_msDS_IsPrimaryComputerFor gets the value of DS_msDS_IsPrimaryComputerFor for the instance
func (instance *ds_top) GetPropertyDS_msDS_IsPrimaryComputerFor() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_IsPrimaryComputerFor")
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

// SetDS_msDS_KrbTgtLinkBl sets the value of DS_msDS_KrbTgtLinkBl for the instance
func (instance *ds_top) SetPropertyDS_msDS_KrbTgtLinkBl(value []string) (err error) {
	return instance.SetProperty("DS_msDS_KrbTgtLinkBl", (value))
}

// GetDS_msDS_KrbTgtLinkBl gets the value of DS_msDS_KrbTgtLinkBl for the instance
func (instance *ds_top) GetPropertyDS_msDS_KrbTgtLinkBl() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_KrbTgtLinkBl")
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

// SetDS_msDS_LastKnownRDN sets the value of DS_msDS_LastKnownRDN for the instance
func (instance *ds_top) SetPropertyDS_msDS_LastKnownRDN(value string) (err error) {
	return instance.SetProperty("DS_msDS_LastKnownRDN", (value))
}

// GetDS_msDS_LastKnownRDN gets the value of DS_msDS_LastKnownRDN for the instance
func (instance *ds_top) GetPropertyDS_msDS_LastKnownRDN() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_LastKnownRDN")
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

// SetDS_msDS_LocalEffectiveDeletionTime sets the value of DS_msDS_LocalEffectiveDeletionTime for the instance
func (instance *ds_top) SetPropertyDS_msDS_LocalEffectiveDeletionTime(value string) (err error) {
	return instance.SetProperty("DS_msDS_LocalEffectiveDeletionTime", (value))
}

// GetDS_msDS_LocalEffectiveDeletionTime gets the value of DS_msDS_LocalEffectiveDeletionTime for the instance
func (instance *ds_top) GetPropertyDS_msDS_LocalEffectiveDeletionTime() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_LocalEffectiveDeletionTime")
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

// SetDS_msDS_LocalEffectiveRecycleTime sets the value of DS_msDS_LocalEffectiveRecycleTime for the instance
func (instance *ds_top) SetPropertyDS_msDS_LocalEffectiveRecycleTime(value string) (err error) {
	return instance.SetProperty("DS_msDS_LocalEffectiveRecycleTime", (value))
}

// GetDS_msDS_LocalEffectiveRecycleTime gets the value of DS_msDS_LocalEffectiveRecycleTime for the instance
func (instance *ds_top) GetPropertyDS_msDS_LocalEffectiveRecycleTime() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_LocalEffectiveRecycleTime")
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

// SetDS_msDs_masteredBy sets the value of DS_msDs_masteredBy for the instance
func (instance *ds_top) SetPropertyDS_msDs_masteredBy(value []string) (err error) {
	return instance.SetProperty("DS_msDs_masteredBy", (value))
}

// GetDS_msDs_masteredBy gets the value of DS_msDs_masteredBy for the instance
func (instance *ds_top) GetPropertyDS_msDs_masteredBy() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDs_masteredBy")
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

// SetDS_msds_memberOfTransitive sets the value of DS_msds_memberOfTransitive for the instance
func (instance *ds_top) SetPropertyDS_msds_memberOfTransitive(value []string) (err error) {
	return instance.SetProperty("DS_msds_memberOfTransitive", (value))
}

// GetDS_msds_memberOfTransitive gets the value of DS_msds_memberOfTransitive for the instance
func (instance *ds_top) GetPropertyDS_msds_memberOfTransitive() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msds_memberOfTransitive")
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

// SetDS_msDS_MembersForAzRoleBL sets the value of DS_msDS_MembersForAzRoleBL for the instance
func (instance *ds_top) SetPropertyDS_msDS_MembersForAzRoleBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_MembersForAzRoleBL", (value))
}

// GetDS_msDS_MembersForAzRoleBL gets the value of DS_msDS_MembersForAzRoleBL for the instance
func (instance *ds_top) GetPropertyDS_msDS_MembersForAzRoleBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_MembersForAzRoleBL")
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

// SetDS_msDS_MembersOfResourcePropertyListBL sets the value of DS_msDS_MembersOfResourcePropertyListBL for the instance
func (instance *ds_top) SetPropertyDS_msDS_MembersOfResourcePropertyListBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_MembersOfResourcePropertyListBL", (value))
}

// GetDS_msDS_MembersOfResourcePropertyListBL gets the value of DS_msDS_MembersOfResourcePropertyListBL for the instance
func (instance *ds_top) GetPropertyDS_msDS_MembersOfResourcePropertyListBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_MembersOfResourcePropertyListBL")
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

// SetDS_msds_memberTransitive sets the value of DS_msds_memberTransitive for the instance
func (instance *ds_top) SetPropertyDS_msds_memberTransitive(value []string) (err error) {
	return instance.SetProperty("DS_msds_memberTransitive", (value))
}

// GetDS_msds_memberTransitive gets the value of DS_msds_memberTransitive for the instance
func (instance *ds_top) GetPropertyDS_msds_memberTransitive() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msds_memberTransitive")
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

// SetDS_msDS_NC_RO_Replica_Locations_BL sets the value of DS_msDS_NC_RO_Replica_Locations_BL for the instance
func (instance *ds_top) SetPropertyDS_msDS_NC_RO_Replica_Locations_BL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_NC_RO_Replica_Locations_BL", (value))
}

// GetDS_msDS_NC_RO_Replica_Locations_BL gets the value of DS_msDS_NC_RO_Replica_Locations_BL for the instance
func (instance *ds_top) GetPropertyDS_msDS_NC_RO_Replica_Locations_BL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_NC_RO_Replica_Locations_BL")
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

// SetDS_msDS_NCReplCursors sets the value of DS_msDS_NCReplCursors for the instance
func (instance *ds_top) SetPropertyDS_msDS_NCReplCursors(value []string) (err error) {
	return instance.SetProperty("DS_msDS_NCReplCursors", (value))
}

// GetDS_msDS_NCReplCursors gets the value of DS_msDS_NCReplCursors for the instance
func (instance *ds_top) GetPropertyDS_msDS_NCReplCursors() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_NCReplCursors")
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

// SetDS_msDS_NCReplInboundNeighbors sets the value of DS_msDS_NCReplInboundNeighbors for the instance
func (instance *ds_top) SetPropertyDS_msDS_NCReplInboundNeighbors(value []string) (err error) {
	return instance.SetProperty("DS_msDS_NCReplInboundNeighbors", (value))
}

// GetDS_msDS_NCReplInboundNeighbors gets the value of DS_msDS_NCReplInboundNeighbors for the instance
func (instance *ds_top) GetPropertyDS_msDS_NCReplInboundNeighbors() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_NCReplInboundNeighbors")
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

// SetDS_msDS_NCReplOutboundNeighbors sets the value of DS_msDS_NCReplOutboundNeighbors for the instance
func (instance *ds_top) SetPropertyDS_msDS_NCReplOutboundNeighbors(value []string) (err error) {
	return instance.SetProperty("DS_msDS_NCReplOutboundNeighbors", (value))
}

// GetDS_msDS_NCReplOutboundNeighbors gets the value of DS_msDS_NCReplOutboundNeighbors for the instance
func (instance *ds_top) GetPropertyDS_msDS_NCReplOutboundNeighbors() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_NCReplOutboundNeighbors")
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

// SetDS_msDS_NcType sets the value of DS_msDS_NcType for the instance
func (instance *ds_top) SetPropertyDS_msDS_NcType(value int32) (err error) {
	return instance.SetProperty("DS_msDS_NcType", (value))
}

// GetDS_msDS_NcType gets the value of DS_msDS_NcType for the instance
func (instance *ds_top) GetPropertyDS_msDS_NcType() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_msDS_NcType")
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

// SetDS_msDS_NonMembersBL sets the value of DS_msDS_NonMembersBL for the instance
func (instance *ds_top) SetPropertyDS_msDS_NonMembersBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_NonMembersBL", (value))
}

// GetDS_msDS_NonMembersBL gets the value of DS_msDS_NonMembersBL for the instance
func (instance *ds_top) GetPropertyDS_msDS_NonMembersBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_NonMembersBL")
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

// SetDS_msDS_ObjectReferenceBL sets the value of DS_msDS_ObjectReferenceBL for the instance
func (instance *ds_top) SetPropertyDS_msDS_ObjectReferenceBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_ObjectReferenceBL", (value))
}

// GetDS_msDS_ObjectReferenceBL gets the value of DS_msDS_ObjectReferenceBL for the instance
func (instance *ds_top) GetPropertyDS_msDS_ObjectReferenceBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ObjectReferenceBL")
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

// SetDS_msDS_ObjectSoa sets the value of DS_msDS_ObjectSoa for the instance
func (instance *ds_top) SetPropertyDS_msDS_ObjectSoa(value string) (err error) {
	return instance.SetProperty("DS_msDS_ObjectSoa", (value))
}

// GetDS_msDS_ObjectSoa gets the value of DS_msDS_ObjectSoa for the instance
func (instance *ds_top) GetPropertyDS_msDS_ObjectSoa() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ObjectSoa")
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

// SetDS_msDS_OIDToGroupLinkBl sets the value of DS_msDS_OIDToGroupLinkBl for the instance
func (instance *ds_top) SetPropertyDS_msDS_OIDToGroupLinkBl(value []string) (err error) {
	return instance.SetProperty("DS_msDS_OIDToGroupLinkBl", (value))
}

// GetDS_msDS_OIDToGroupLinkBl gets the value of DS_msDS_OIDToGroupLinkBl for the instance
func (instance *ds_top) GetPropertyDS_msDS_OIDToGroupLinkBl() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_OIDToGroupLinkBl")
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

// SetDS_msDS_OperationsForAzRoleBL sets the value of DS_msDS_OperationsForAzRoleBL for the instance
func (instance *ds_top) SetPropertyDS_msDS_OperationsForAzRoleBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_OperationsForAzRoleBL", (value))
}

// GetDS_msDS_OperationsForAzRoleBL gets the value of DS_msDS_OperationsForAzRoleBL for the instance
func (instance *ds_top) GetPropertyDS_msDS_OperationsForAzRoleBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_OperationsForAzRoleBL")
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

// SetDS_msDS_OperationsForAzTaskBL sets the value of DS_msDS_OperationsForAzTaskBL for the instance
func (instance *ds_top) SetPropertyDS_msDS_OperationsForAzTaskBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_OperationsForAzTaskBL", (value))
}

// GetDS_msDS_OperationsForAzTaskBL gets the value of DS_msDS_OperationsForAzTaskBL for the instance
func (instance *ds_top) GetPropertyDS_msDS_OperationsForAzTaskBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_OperationsForAzTaskBL")
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

// SetDS_msDS_parentdistname sets the value of DS_msDS_parentdistname for the instance
func (instance *ds_top) SetPropertyDS_msDS_parentdistname(value string) (err error) {
	return instance.SetProperty("DS_msDS_parentdistname", (value))
}

// GetDS_msDS_parentdistname gets the value of DS_msDS_parentdistname for the instance
func (instance *ds_top) GetPropertyDS_msDS_parentdistname() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_parentdistname")
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

// SetDS_msDS_PrincipalName sets the value of DS_msDS_PrincipalName for the instance
func (instance *ds_top) SetPropertyDS_msDS_PrincipalName(value string) (err error) {
	return instance.SetProperty("DS_msDS_PrincipalName", (value))
}

// GetDS_msDS_PrincipalName gets the value of DS_msDS_PrincipalName for the instance
func (instance *ds_top) GetPropertyDS_msDS_PrincipalName() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_PrincipalName")
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

// SetDS_msDS_PSOApplied sets the value of DS_msDS_PSOApplied for the instance
func (instance *ds_top) SetPropertyDS_msDS_PSOApplied(value []string) (err error) {
	return instance.SetProperty("DS_msDS_PSOApplied", (value))
}

// GetDS_msDS_PSOApplied gets the value of DS_msDS_PSOApplied for the instance
func (instance *ds_top) GetPropertyDS_msDS_PSOApplied() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_PSOApplied")
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

// SetDS_msDS_ReplAttributeMetaData sets the value of DS_msDS_ReplAttributeMetaData for the instance
func (instance *ds_top) SetPropertyDS_msDS_ReplAttributeMetaData(value []string) (err error) {
	return instance.SetProperty("DS_msDS_ReplAttributeMetaData", (value))
}

// GetDS_msDS_ReplAttributeMetaData gets the value of DS_msDS_ReplAttributeMetaData for the instance
func (instance *ds_top) GetPropertyDS_msDS_ReplAttributeMetaData() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ReplAttributeMetaData")
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

// SetDS_msDS_ReplValueMetaData sets the value of DS_msDS_ReplValueMetaData for the instance
func (instance *ds_top) SetPropertyDS_msDS_ReplValueMetaData(value []string) (err error) {
	return instance.SetProperty("DS_msDS_ReplValueMetaData", (value))
}

// GetDS_msDS_ReplValueMetaData gets the value of DS_msDS_ReplValueMetaData for the instance
func (instance *ds_top) GetPropertyDS_msDS_ReplValueMetaData() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ReplValueMetaData")
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

// SetDS_msDS_ReplValueMetaDataExt sets the value of DS_msDS_ReplValueMetaDataExt for the instance
func (instance *ds_top) SetPropertyDS_msDS_ReplValueMetaDataExt(value []string) (err error) {
	return instance.SetProperty("DS_msDS_ReplValueMetaDataExt", (value))
}

// GetDS_msDS_ReplValueMetaDataExt gets the value of DS_msDS_ReplValueMetaDataExt for the instance
func (instance *ds_top) GetPropertyDS_msDS_ReplValueMetaDataExt() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ReplValueMetaDataExt")
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

// SetDS_msDS_RevealedDSAs sets the value of DS_msDS_RevealedDSAs for the instance
func (instance *ds_top) SetPropertyDS_msDS_RevealedDSAs(value []string) (err error) {
	return instance.SetProperty("DS_msDS_RevealedDSAs", (value))
}

// GetDS_msDS_RevealedDSAs gets the value of DS_msDS_RevealedDSAs for the instance
func (instance *ds_top) GetPropertyDS_msDS_RevealedDSAs() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_RevealedDSAs")
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

// SetDS_msDS_RevealedListBL sets the value of DS_msDS_RevealedListBL for the instance
func (instance *ds_top) SetPropertyDS_msDS_RevealedListBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_RevealedListBL", (value))
}

// GetDS_msDS_RevealedListBL gets the value of DS_msDS_RevealedListBL for the instance
func (instance *ds_top) GetPropertyDS_msDS_RevealedListBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_RevealedListBL")
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

// SetDS_msDS_SourceAnchor sets the value of DS_msDS_SourceAnchor for the instance
func (instance *ds_top) SetPropertyDS_msDS_SourceAnchor(value string) (err error) {
	return instance.SetProperty("DS_msDS_SourceAnchor", (value))
}

// GetDS_msDS_SourceAnchor gets the value of DS_msDS_SourceAnchor for the instance
func (instance *ds_top) GetPropertyDS_msDS_SourceAnchor() (value string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_SourceAnchor")
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

// SetDS_msDS_TasksForAzRoleBL sets the value of DS_msDS_TasksForAzRoleBL for the instance
func (instance *ds_top) SetPropertyDS_msDS_TasksForAzRoleBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_TasksForAzRoleBL", (value))
}

// GetDS_msDS_TasksForAzRoleBL gets the value of DS_msDS_TasksForAzRoleBL for the instance
func (instance *ds_top) GetPropertyDS_msDS_TasksForAzRoleBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_TasksForAzRoleBL")
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

// SetDS_msDS_TasksForAzTaskBL sets the value of DS_msDS_TasksForAzTaskBL for the instance
func (instance *ds_top) SetPropertyDS_msDS_TasksForAzTaskBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_TasksForAzTaskBL", (value))
}

// GetDS_msDS_TasksForAzTaskBL gets the value of DS_msDS_TasksForAzTaskBL for the instance
func (instance *ds_top) GetPropertyDS_msDS_TasksForAzTaskBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_TasksForAzTaskBL")
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

// SetDS_msDS_TDOEgressBL sets the value of DS_msDS_TDOEgressBL for the instance
func (instance *ds_top) SetPropertyDS_msDS_TDOEgressBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_TDOEgressBL", (value))
}

// GetDS_msDS_TDOEgressBL gets the value of DS_msDS_TDOEgressBL for the instance
func (instance *ds_top) GetPropertyDS_msDS_TDOEgressBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_TDOEgressBL")
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

// SetDS_msDS_TDOIngressBL sets the value of DS_msDS_TDOIngressBL for the instance
func (instance *ds_top) SetPropertyDS_msDS_TDOIngressBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_TDOIngressBL", (value))
}

// GetDS_msDS_TDOIngressBL gets the value of DS_msDS_TDOIngressBL for the instance
func (instance *ds_top) GetPropertyDS_msDS_TDOIngressBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_TDOIngressBL")
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

// SetDS_msDS_ValueTypeReferenceBL sets the value of DS_msDS_ValueTypeReferenceBL for the instance
func (instance *ds_top) SetPropertyDS_msDS_ValueTypeReferenceBL(value []string) (err error) {
	return instance.SetProperty("DS_msDS_ValueTypeReferenceBL", (value))
}

// GetDS_msDS_ValueTypeReferenceBL gets the value of DS_msDS_ValueTypeReferenceBL for the instance
func (instance *ds_top) GetPropertyDS_msDS_ValueTypeReferenceBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msDS_ValueTypeReferenceBL")
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

// SetDS_msSFU30PosixMemberOf sets the value of DS_msSFU30PosixMemberOf for the instance
func (instance *ds_top) SetPropertyDS_msSFU30PosixMemberOf(value []string) (err error) {
	return instance.SetProperty("DS_msSFU30PosixMemberOf", (value))
}

// GetDS_msSFU30PosixMemberOf gets the value of DS_msSFU30PosixMemberOf for the instance
func (instance *ds_top) GetPropertyDS_msSFU30PosixMemberOf() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_msSFU30PosixMemberOf")
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

// SetDS_name sets the value of DS_name for the instance
func (instance *ds_top) SetPropertyDS_name(value string) (err error) {
	return instance.SetProperty("DS_name", (value))
}

// GetDS_name gets the value of DS_name for the instance
func (instance *ds_top) GetPropertyDS_name() (value string, err error) {
	retValue, err := instance.GetProperty("DS_name")
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

// SetDS_netbootSCPBL sets the value of DS_netbootSCPBL for the instance
func (instance *ds_top) SetPropertyDS_netbootSCPBL(value []string) (err error) {
	return instance.SetProperty("DS_netbootSCPBL", (value))
}

// GetDS_netbootSCPBL gets the value of DS_netbootSCPBL for the instance
func (instance *ds_top) GetPropertyDS_netbootSCPBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_netbootSCPBL")
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

// SetDS_nonSecurityMemberBL sets the value of DS_nonSecurityMemberBL for the instance
func (instance *ds_top) SetPropertyDS_nonSecurityMemberBL(value []string) (err error) {
	return instance.SetProperty("DS_nonSecurityMemberBL", (value))
}

// GetDS_nonSecurityMemberBL gets the value of DS_nonSecurityMemberBL for the instance
func (instance *ds_top) GetPropertyDS_nonSecurityMemberBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_nonSecurityMemberBL")
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

// SetDS_nTSecurityDescriptor sets the value of DS_nTSecurityDescriptor for the instance
func (instance *ds_top) SetPropertyDS_nTSecurityDescriptor(value Uint8Array) (err error) {
	return instance.SetProperty("DS_nTSecurityDescriptor", (value))
}

// GetDS_nTSecurityDescriptor gets the value of DS_nTSecurityDescriptor for the instance
func (instance *ds_top) GetPropertyDS_nTSecurityDescriptor() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_nTSecurityDescriptor")
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

// SetDS_objectCategory sets the value of DS_objectCategory for the instance
func (instance *ds_top) SetPropertyDS_objectCategory(value string) (err error) {
	return instance.SetProperty("DS_objectCategory", (value))
}

// GetDS_objectCategory gets the value of DS_objectCategory for the instance
func (instance *ds_top) GetPropertyDS_objectCategory() (value string, err error) {
	retValue, err := instance.GetProperty("DS_objectCategory")
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

// SetDS_objectClass sets the value of DS_objectClass for the instance
func (instance *ds_top) SetPropertyDS_objectClass(value []string) (err error) {
	return instance.SetProperty("DS_objectClass", (value))
}

// GetDS_objectClass gets the value of DS_objectClass for the instance
func (instance *ds_top) GetPropertyDS_objectClass() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_objectClass")
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

// SetDS_objectGUID sets the value of DS_objectGUID for the instance
func (instance *ds_top) SetPropertyDS_objectGUID(value Uint8Array) (err error) {
	return instance.SetProperty("DS_objectGUID", (value))
}

// GetDS_objectGUID gets the value of DS_objectGUID for the instance
func (instance *ds_top) GetPropertyDS_objectGUID() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_objectGUID")
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

// SetDS_objectVersion sets the value of DS_objectVersion for the instance
func (instance *ds_top) SetPropertyDS_objectVersion(value int32) (err error) {
	return instance.SetProperty("DS_objectVersion", (value))
}

// GetDS_objectVersion gets the value of DS_objectVersion for the instance
func (instance *ds_top) GetPropertyDS_objectVersion() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_objectVersion")
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

// SetDS_otherWellKnownObjects sets the value of DS_otherWellKnownObjects for the instance
func (instance *ds_top) SetPropertyDS_otherWellKnownObjects(value []DN_With_Binary) (err error) {
	return instance.SetProperty("DS_otherWellKnownObjects", (value))
}

// GetDS_otherWellKnownObjects gets the value of DS_otherWellKnownObjects for the instance
func (instance *ds_top) GetPropertyDS_otherWellKnownObjects() (value []DN_With_Binary, err error) {
	retValue, err := instance.GetProperty("DS_otherWellKnownObjects")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(DN_With_Binary)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " DN_With_Binary is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, DN_With_Binary(valuetmp))
	}

	return
}

// SetDS_ownerBL sets the value of DS_ownerBL for the instance
func (instance *ds_top) SetPropertyDS_ownerBL(value []string) (err error) {
	return instance.SetProperty("DS_ownerBL", (value))
}

// GetDS_ownerBL gets the value of DS_ownerBL for the instance
func (instance *ds_top) GetPropertyDS_ownerBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_ownerBL")
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

// SetDS_partialAttributeDeletionList sets the value of DS_partialAttributeDeletionList for the instance
func (instance *ds_top) SetPropertyDS_partialAttributeDeletionList(value Uint8Array) (err error) {
	return instance.SetProperty("DS_partialAttributeDeletionList", (value))
}

// GetDS_partialAttributeDeletionList gets the value of DS_partialAttributeDeletionList for the instance
func (instance *ds_top) GetPropertyDS_partialAttributeDeletionList() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_partialAttributeDeletionList")
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

// SetDS_partialAttributeSet sets the value of DS_partialAttributeSet for the instance
func (instance *ds_top) SetPropertyDS_partialAttributeSet(value Uint8Array) (err error) {
	return instance.SetProperty("DS_partialAttributeSet", (value))
}

// GetDS_partialAttributeSet gets the value of DS_partialAttributeSet for the instance
func (instance *ds_top) GetPropertyDS_partialAttributeSet() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_partialAttributeSet")
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

// SetDS_possibleInferiors sets the value of DS_possibleInferiors for the instance
func (instance *ds_top) SetPropertyDS_possibleInferiors(value []string) (err error) {
	return instance.SetProperty("DS_possibleInferiors", (value))
}

// GetDS_possibleInferiors gets the value of DS_possibleInferiors for the instance
func (instance *ds_top) GetPropertyDS_possibleInferiors() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_possibleInferiors")
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

// SetDS_proxiedObjectName sets the value of DS_proxiedObjectName for the instance
func (instance *ds_top) SetPropertyDS_proxiedObjectName(value DN_With_Binary) (err error) {
	return instance.SetProperty("DS_proxiedObjectName", (value))
}

// GetDS_proxiedObjectName gets the value of DS_proxiedObjectName for the instance
func (instance *ds_top) GetPropertyDS_proxiedObjectName() (value DN_With_Binary, err error) {
	retValue, err := instance.GetProperty("DS_proxiedObjectName")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(DN_With_Binary)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " DN_With_Binary is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = DN_With_Binary(valuetmp)

	return
}

// SetDS_proxyAddresses sets the value of DS_proxyAddresses for the instance
func (instance *ds_top) SetPropertyDS_proxyAddresses(value []string) (err error) {
	return instance.SetProperty("DS_proxyAddresses", (value))
}

// GetDS_proxyAddresses gets the value of DS_proxyAddresses for the instance
func (instance *ds_top) GetPropertyDS_proxyAddresses() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_proxyAddresses")
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

// SetDS_queryPolicyBL sets the value of DS_queryPolicyBL for the instance
func (instance *ds_top) SetPropertyDS_queryPolicyBL(value []string) (err error) {
	return instance.SetProperty("DS_queryPolicyBL", (value))
}

// GetDS_queryPolicyBL gets the value of DS_queryPolicyBL for the instance
func (instance *ds_top) GetPropertyDS_queryPolicyBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_queryPolicyBL")
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

// SetDS_replPropertyMetaData sets the value of DS_replPropertyMetaData for the instance
func (instance *ds_top) SetPropertyDS_replPropertyMetaData(value Uint8Array) (err error) {
	return instance.SetProperty("DS_replPropertyMetaData", (value))
}

// GetDS_replPropertyMetaData gets the value of DS_replPropertyMetaData for the instance
func (instance *ds_top) GetPropertyDS_replPropertyMetaData() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_replPropertyMetaData")
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

// SetDS_replUpToDateVector sets the value of DS_replUpToDateVector for the instance
func (instance *ds_top) SetPropertyDS_replUpToDateVector(value Uint8Array) (err error) {
	return instance.SetProperty("DS_replUpToDateVector", (value))
}

// GetDS_replUpToDateVector gets the value of DS_replUpToDateVector for the instance
func (instance *ds_top) GetPropertyDS_replUpToDateVector() (value Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_replUpToDateVector")
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

// SetDS_repsFrom sets the value of DS_repsFrom for the instance
func (instance *ds_top) SetPropertyDS_repsFrom(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_repsFrom", (value))
}

// GetDS_repsFrom gets the value of DS_repsFrom for the instance
func (instance *ds_top) GetPropertyDS_repsFrom() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_repsFrom")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Uint8Array)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Uint8Array(valuetmp))
	}

	return
}

// SetDS_repsTo sets the value of DS_repsTo for the instance
func (instance *ds_top) SetPropertyDS_repsTo(value []Uint8Array) (err error) {
	return instance.SetProperty("DS_repsTo", (value))
}

// GetDS_repsTo gets the value of DS_repsTo for the instance
func (instance *ds_top) GetPropertyDS_repsTo() (value []Uint8Array, err error) {
	retValue, err := instance.GetProperty("DS_repsTo")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(Uint8Array)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " Uint8Array is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, Uint8Array(valuetmp))
	}

	return
}

// SetDS_revision sets the value of DS_revision for the instance
func (instance *ds_top) SetPropertyDS_revision(value int32) (err error) {
	return instance.SetProperty("DS_revision", (value))
}

// GetDS_revision gets the value of DS_revision for the instance
func (instance *ds_top) GetPropertyDS_revision() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_revision")
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

// SetDS_sDRightsEffective sets the value of DS_sDRightsEffective for the instance
func (instance *ds_top) SetPropertyDS_sDRightsEffective(value int32) (err error) {
	return instance.SetProperty("DS_sDRightsEffective", (value))
}

// GetDS_sDRightsEffective gets the value of DS_sDRightsEffective for the instance
func (instance *ds_top) GetPropertyDS_sDRightsEffective() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_sDRightsEffective")
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

// SetDS_serverReferenceBL sets the value of DS_serverReferenceBL for the instance
func (instance *ds_top) SetPropertyDS_serverReferenceBL(value []string) (err error) {
	return instance.SetProperty("DS_serverReferenceBL", (value))
}

// GetDS_serverReferenceBL gets the value of DS_serverReferenceBL for the instance
func (instance *ds_top) GetPropertyDS_serverReferenceBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_serverReferenceBL")
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

// SetDS_showInAdvancedViewOnly sets the value of DS_showInAdvancedViewOnly for the instance
func (instance *ds_top) SetPropertyDS_showInAdvancedViewOnly(value bool) (err error) {
	return instance.SetProperty("DS_showInAdvancedViewOnly", (value))
}

// GetDS_showInAdvancedViewOnly gets the value of DS_showInAdvancedViewOnly for the instance
func (instance *ds_top) GetPropertyDS_showInAdvancedViewOnly() (value bool, err error) {
	retValue, err := instance.GetProperty("DS_showInAdvancedViewOnly")
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

// SetDS_siteObjectBL sets the value of DS_siteObjectBL for the instance
func (instance *ds_top) SetPropertyDS_siteObjectBL(value []string) (err error) {
	return instance.SetProperty("DS_siteObjectBL", (value))
}

// GetDS_siteObjectBL gets the value of DS_siteObjectBL for the instance
func (instance *ds_top) GetPropertyDS_siteObjectBL() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_siteObjectBL")
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

// SetDS_structuralObjectClass sets the value of DS_structuralObjectClass for the instance
func (instance *ds_top) SetPropertyDS_structuralObjectClass(value []string) (err error) {
	return instance.SetProperty("DS_structuralObjectClass", (value))
}

// GetDS_structuralObjectClass gets the value of DS_structuralObjectClass for the instance
func (instance *ds_top) GetPropertyDS_structuralObjectClass() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_structuralObjectClass")
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

// SetDS_subRefs sets the value of DS_subRefs for the instance
func (instance *ds_top) SetPropertyDS_subRefs(value []string) (err error) {
	return instance.SetProperty("DS_subRefs", (value))
}

// GetDS_subRefs gets the value of DS_subRefs for the instance
func (instance *ds_top) GetPropertyDS_subRefs() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_subRefs")
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

// SetDS_subSchemaSubEntry sets the value of DS_subSchemaSubEntry for the instance
func (instance *ds_top) SetPropertyDS_subSchemaSubEntry(value []string) (err error) {
	return instance.SetProperty("DS_subSchemaSubEntry", (value))
}

// GetDS_subSchemaSubEntry gets the value of DS_subSchemaSubEntry for the instance
func (instance *ds_top) GetPropertyDS_subSchemaSubEntry() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_subSchemaSubEntry")
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

// SetDS_systemFlags sets the value of DS_systemFlags for the instance
func (instance *ds_top) SetPropertyDS_systemFlags(value int32) (err error) {
	return instance.SetProperty("DS_systemFlags", (value))
}

// GetDS_systemFlags gets the value of DS_systemFlags for the instance
func (instance *ds_top) GetPropertyDS_systemFlags() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_systemFlags")
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

// SetDS_url sets the value of DS_url for the instance
func (instance *ds_top) SetPropertyDS_url(value []string) (err error) {
	return instance.SetProperty("DS_url", (value))
}

// GetDS_url gets the value of DS_url for the instance
func (instance *ds_top) GetPropertyDS_url() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_url")
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

// SetDS_uSNChanged sets the value of DS_uSNChanged for the instance
func (instance *ds_top) SetPropertyDS_uSNChanged(value int64) (err error) {
	return instance.SetProperty("DS_uSNChanged", (value))
}

// GetDS_uSNChanged gets the value of DS_uSNChanged for the instance
func (instance *ds_top) GetPropertyDS_uSNChanged() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_uSNChanged")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_uSNCreated sets the value of DS_uSNCreated for the instance
func (instance *ds_top) SetPropertyDS_uSNCreated(value int64) (err error) {
	return instance.SetProperty("DS_uSNCreated", (value))
}

// GetDS_uSNCreated gets the value of DS_uSNCreated for the instance
func (instance *ds_top) GetPropertyDS_uSNCreated() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_uSNCreated")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_uSNDSALastObjRemoved sets the value of DS_uSNDSALastObjRemoved for the instance
func (instance *ds_top) SetPropertyDS_uSNDSALastObjRemoved(value int64) (err error) {
	return instance.SetProperty("DS_uSNDSALastObjRemoved", (value))
}

// GetDS_uSNDSALastObjRemoved gets the value of DS_uSNDSALastObjRemoved for the instance
func (instance *ds_top) GetPropertyDS_uSNDSALastObjRemoved() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_uSNDSALastObjRemoved")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_USNIntersite sets the value of DS_USNIntersite for the instance
func (instance *ds_top) SetPropertyDS_USNIntersite(value int32) (err error) {
	return instance.SetProperty("DS_USNIntersite", (value))
}

// GetDS_USNIntersite gets the value of DS_USNIntersite for the instance
func (instance *ds_top) GetPropertyDS_USNIntersite() (value int32, err error) {
	retValue, err := instance.GetProperty("DS_USNIntersite")
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

// SetDS_uSNLastObjRem sets the value of DS_uSNLastObjRem for the instance
func (instance *ds_top) SetPropertyDS_uSNLastObjRem(value int64) (err error) {
	return instance.SetProperty("DS_uSNLastObjRem", (value))
}

// GetDS_uSNLastObjRem gets the value of DS_uSNLastObjRem for the instance
func (instance *ds_top) GetPropertyDS_uSNLastObjRem() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_uSNLastObjRem")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_uSNSource sets the value of DS_uSNSource for the instance
func (instance *ds_top) SetPropertyDS_uSNSource(value int64) (err error) {
	return instance.SetProperty("DS_uSNSource", (value))
}

// GetDS_uSNSource gets the value of DS_uSNSource for the instance
func (instance *ds_top) GetPropertyDS_uSNSource() (value int64, err error) {
	retValue, err := instance.GetProperty("DS_uSNSource")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(int64)
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " int64 is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = int64(valuetmp)

	return
}

// SetDS_wbemPath sets the value of DS_wbemPath for the instance
func (instance *ds_top) SetPropertyDS_wbemPath(value []string) (err error) {
	return instance.SetProperty("DS_wbemPath", (value))
}

// GetDS_wbemPath gets the value of DS_wbemPath for the instance
func (instance *ds_top) GetPropertyDS_wbemPath() (value []string, err error) {
	retValue, err := instance.GetProperty("DS_wbemPath")
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

// SetDS_wellKnownObjects sets the value of DS_wellKnownObjects for the instance
func (instance *ds_top) SetPropertyDS_wellKnownObjects(value []DN_With_Binary) (err error) {
	return instance.SetProperty("DS_wellKnownObjects", (value))
}

// GetDS_wellKnownObjects gets the value of DS_wellKnownObjects for the instance
func (instance *ds_top) GetPropertyDS_wellKnownObjects() (value []DN_With_Binary, err error) {
	retValue, err := instance.GetProperty("DS_wellKnownObjects")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	for _, interfaceValue := range retValue.([]interface{}) {
		valuetmp, ok := interfaceValue.(DN_With_Binary)
		if !ok {
			err = errors.Wrapf(errors.InvalidType, " DN_With_Binary is Invalid. Expected %s", reflect.TypeOf(interfaceValue))
			return
		}
		value = append(value, DN_With_Binary(valuetmp))
	}

	return
}

// SetDS_whenChanged sets the value of DS_whenChanged for the instance
func (instance *ds_top) SetPropertyDS_whenChanged(value string) (err error) {
	return instance.SetProperty("DS_whenChanged", (value))
}

// GetDS_whenChanged gets the value of DS_whenChanged for the instance
func (instance *ds_top) GetPropertyDS_whenChanged() (value string, err error) {
	retValue, err := instance.GetProperty("DS_whenChanged")
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

// SetDS_whenCreated sets the value of DS_whenCreated for the instance
func (instance *ds_top) SetPropertyDS_whenCreated(value string) (err error) {
	return instance.SetProperty("DS_whenCreated", (value))
}

// GetDS_whenCreated gets the value of DS_whenCreated for the instance
func (instance *ds_top) GetPropertyDS_whenCreated() (value string, err error) {
	retValue, err := instance.GetProperty("DS_whenCreated")
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

// SetDS_wWWHomePage sets the value of DS_wWWHomePage for the instance
func (instance *ds_top) SetPropertyDS_wWWHomePage(value string) (err error) {
	return instance.SetProperty("DS_wWWHomePage", (value))
}

// GetDS_wWWHomePage gets the value of DS_wWWHomePage for the instance
func (instance *ds_top) GetPropertyDS_wWWHomePage() (value string, err error) {
	retValue, err := instance.GetProperty("DS_wWWHomePage")
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
