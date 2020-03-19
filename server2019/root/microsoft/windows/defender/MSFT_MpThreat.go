// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/19/2020 using wmigen
//      Source root.Microsoft.Windows.Defender
//////////////////////////////////////////////
package defender

import (
	"github.com/microsoft/wmi/pkg/base/query"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
)

// MSFT_MpThreat struct
type MSFT_MpThreat struct {
	*BaseStatus

	//
	CategoryID uint8

	//
	DidThreatExecute bool

	//
	IsActive bool

	//
	Resources []string

	//
	RollupStatus uint32

	//
	SchemaVersion string

	//
	SeverityID uint8

	//
	ThreatID int64

	//
	ThreatName string

	//
	TypeID uint8
}

func NewMSFT_MpThreatEx1(instance *cim.WmiInstance) (newInstance *MSFT_MpThreat, err error) {
	tmp, err := NewBaseStatusEx1(instance)

	if err != nil {
		return
	}
	newInstance = &MSFT_MpThreat{
		BaseStatus: tmp,
	}
	return
}

func NewMSFT_MpThreatEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *MSFT_MpThreat, err error) {
	tmp, err := NewBaseStatusEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &MSFT_MpThreat{
		BaseStatus: tmp,
	}
	return
}

// SetCategoryID sets the value of CategoryID for the instance
func (instance *MSFT_MpThreat) SetPropertyCategoryID(value uint8) (err error) {
	return instance.SetProperty("CategoryID", value)
}

// GetCategoryID gets the value of CategoryID for the instance
func (instance *MSFT_MpThreat) GetPropertyCategoryID() (value uint8, err error) {
	retValue, err := instance.GetProperty("CategoryID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetDidThreatExecute sets the value of DidThreatExecute for the instance
func (instance *MSFT_MpThreat) SetPropertyDidThreatExecute(value bool) (err error) {
	return instance.SetProperty("DidThreatExecute", value)
}

// GetDidThreatExecute gets the value of DidThreatExecute for the instance
func (instance *MSFT_MpThreat) GetPropertyDidThreatExecute() (value bool, err error) {
	retValue, err := instance.GetProperty("DidThreatExecute")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetIsActive sets the value of IsActive for the instance
func (instance *MSFT_MpThreat) SetPropertyIsActive(value bool) (err error) {
	return instance.SetProperty("IsActive", value)
}

// GetIsActive gets the value of IsActive for the instance
func (instance *MSFT_MpThreat) GetPropertyIsActive() (value bool, err error) {
	retValue, err := instance.GetProperty("IsActive")
	if err != nil {
		return
	}
	value, ok := retValue.(bool)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetResources sets the value of Resources for the instance
func (instance *MSFT_MpThreat) SetPropertyResources(value []string) (err error) {
	return instance.SetProperty("Resources", value)
}

// GetResources gets the value of Resources for the instance
func (instance *MSFT_MpThreat) GetPropertyResources() (value []string, err error) {
	retValue, err := instance.GetProperty("Resources")
	if err != nil {
		return
	}
	value, ok := retValue.([]string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetRollupStatus sets the value of RollupStatus for the instance
func (instance *MSFT_MpThreat) SetPropertyRollupStatus(value uint32) (err error) {
	return instance.SetProperty("RollupStatus", value)
}

// GetRollupStatus gets the value of RollupStatus for the instance
func (instance *MSFT_MpThreat) GetPropertyRollupStatus() (value uint32, err error) {
	retValue, err := instance.GetProperty("RollupStatus")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSchemaVersion sets the value of SchemaVersion for the instance
func (instance *MSFT_MpThreat) SetPropertySchemaVersion(value string) (err error) {
	return instance.SetProperty("SchemaVersion", value)
}

// GetSchemaVersion gets the value of SchemaVersion for the instance
func (instance *MSFT_MpThreat) GetPropertySchemaVersion() (value string, err error) {
	retValue, err := instance.GetProperty("SchemaVersion")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetSeverityID sets the value of SeverityID for the instance
func (instance *MSFT_MpThreat) SetPropertySeverityID(value uint8) (err error) {
	return instance.SetProperty("SeverityID", value)
}

// GetSeverityID gets the value of SeverityID for the instance
func (instance *MSFT_MpThreat) GetPropertySeverityID() (value uint8, err error) {
	retValue, err := instance.GetProperty("SeverityID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThreatID sets the value of ThreatID for the instance
func (instance *MSFT_MpThreat) SetPropertyThreatID(value int64) (err error) {
	return instance.SetProperty("ThreatID", value)
}

// GetThreatID gets the value of ThreatID for the instance
func (instance *MSFT_MpThreat) GetPropertyThreatID() (value int64, err error) {
	retValue, err := instance.GetProperty("ThreatID")
	if err != nil {
		return
	}
	value, ok := retValue.(int64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThreatName sets the value of ThreatName for the instance
func (instance *MSFT_MpThreat) SetPropertyThreatName(value string) (err error) {
	return instance.SetProperty("ThreatName", value)
}

// GetThreatName gets the value of ThreatName for the instance
func (instance *MSFT_MpThreat) GetPropertyThreatName() (value string, err error) {
	retValue, err := instance.GetProperty("ThreatName")
	if err != nil {
		return
	}
	value, ok := retValue.(string)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetTypeID sets the value of TypeID for the instance
func (instance *MSFT_MpThreat) SetPropertyTypeID(value uint8) (err error) {
	return instance.SetProperty("TypeID", value)
}

// GetTypeID gets the value of TypeID for the instance
func (instance *MSFT_MpThreat) GetPropertyTypeID() (value uint8, err error) {
	retValue, err := instance.GetProperty("TypeID")
	if err != nil {
		return
	}
	value, ok := retValue.(uint8)
	if !ok {
		// TODO: Set an error
	}
	return
}

//

// <param name="ReturnValue" type="uint32 "></param>
func (instance *MSFT_MpThreat) Remove() (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("Remove")
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}
