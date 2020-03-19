// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.CIMV2.mdm.dmmap
//////////////////////////////////////////////
package dmmap

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MDM_Policy_Result01_Kerberos02 struct
type MDM_Policy_Result01_Kerberos02 struct {
	*cim.WmiInstance

	//
	AllowForestSearchOrder string

	//
	InstanceID string

	//
	KerberosClientSupportsClaimsCompoundArmor string

	//
	ParentID string

	//
	RequireKerberosArmoring string

	//
	RequireStrictKDCValidation string

	//
	SetMaximumContextTokenSize string

	//
	UPNNameHints string
}

func NewMDM_Policy_Result01_Kerberos02Ex1(instance *cim.WmiInstance) (newInstance *MDM_Policy_Result01_Kerberos02, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_Kerberos02{
		WmiInstance: tmp,
	}
	return
}

func NewMDM_Policy_Result01_Kerberos02Ex6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MDM_Policy_Result01_Kerberos02, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MDM_Policy_Result01_Kerberos02{
		WmiInstance: tmp,
	}
	return
}

// SetAllowForestSearchOrder sets the value of AllowForestSearchOrder for the instance
func (instance *MDM_Policy_Result01_Kerberos02) SetPropertyAllowForestSearchOrder(value string) (err error) {
	return instance.SetProperty("AllowForestSearchOrder", value)
}

// GetAllowForestSearchOrder gets the value of AllowForestSearchOrder for the instance
func (instance *MDM_Policy_Result01_Kerberos02) GetPropertyAllowForestSearchOrder() (value string, err error) {
	retValue, err := instance.GetProperty("AllowForestSearchOrder")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetInstanceID sets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Kerberos02) SetPropertyInstanceID(value string) (err error) {
	return instance.SetProperty("InstanceID", value)
}

// GetInstanceID gets the value of InstanceID for the instance
func (instance *MDM_Policy_Result01_Kerberos02) GetPropertyInstanceID() (value string, err error) {
	retValue, err := instance.GetProperty("InstanceID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetKerberosClientSupportsClaimsCompoundArmor sets the value of KerberosClientSupportsClaimsCompoundArmor for the instance
func (instance *MDM_Policy_Result01_Kerberos02) SetPropertyKerberosClientSupportsClaimsCompoundArmor(value string) (err error) {
	return instance.SetProperty("KerberosClientSupportsClaimsCompoundArmor", value)
}

// GetKerberosClientSupportsClaimsCompoundArmor gets the value of KerberosClientSupportsClaimsCompoundArmor for the instance
func (instance *MDM_Policy_Result01_Kerberos02) GetPropertyKerberosClientSupportsClaimsCompoundArmor() (value string, err error) {
	retValue, err := instance.GetProperty("KerberosClientSupportsClaimsCompoundArmor")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetParentID sets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Kerberos02) SetPropertyParentID(value string) (err error) {
	return instance.SetProperty("ParentID", value)
}

// GetParentID gets the value of ParentID for the instance
func (instance *MDM_Policy_Result01_Kerberos02) GetPropertyParentID() (value string, err error) {
	retValue, err := instance.GetProperty("ParentID")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequireKerberosArmoring sets the value of RequireKerberosArmoring for the instance
func (instance *MDM_Policy_Result01_Kerberos02) SetPropertyRequireKerberosArmoring(value string) (err error) {
	return instance.SetProperty("RequireKerberosArmoring", value)
}

// GetRequireKerberosArmoring gets the value of RequireKerberosArmoring for the instance
func (instance *MDM_Policy_Result01_Kerberos02) GetPropertyRequireKerberosArmoring() (value string, err error) {
	retValue, err := instance.GetProperty("RequireKerberosArmoring")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRequireStrictKDCValidation sets the value of RequireStrictKDCValidation for the instance
func (instance *MDM_Policy_Result01_Kerberos02) SetPropertyRequireStrictKDCValidation(value string) (err error) {
	return instance.SetProperty("RequireStrictKDCValidation", value)
}

// GetRequireStrictKDCValidation gets the value of RequireStrictKDCValidation for the instance
func (instance *MDM_Policy_Result01_Kerberos02) GetPropertyRequireStrictKDCValidation() (value string, err error) {
	retValue, err := instance.GetProperty("RequireStrictKDCValidation")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSetMaximumContextTokenSize sets the value of SetMaximumContextTokenSize for the instance
func (instance *MDM_Policy_Result01_Kerberos02) SetPropertySetMaximumContextTokenSize(value string) (err error) {
	return instance.SetProperty("SetMaximumContextTokenSize", value)
}

// GetSetMaximumContextTokenSize gets the value of SetMaximumContextTokenSize for the instance
func (instance *MDM_Policy_Result01_Kerberos02) GetPropertySetMaximumContextTokenSize() (value string, err error) {
	retValue, err := instance.GetProperty("SetMaximumContextTokenSize")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetUPNNameHints sets the value of UPNNameHints for the instance
func (instance *MDM_Policy_Result01_Kerberos02) SetPropertyUPNNameHints(value string) (err error) {
	return instance.SetProperty("UPNNameHints", value)
}

// GetUPNNameHints gets the value of UPNNameHints for the instance
func (instance *MDM_Policy_Result01_Kerberos02) GetPropertyUPNNameHints() (value string, err error) {
	retValue, err := instance.GetProperty("UPNNameHints")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}
