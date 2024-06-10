// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 9/18/2020 using wmigen
//      Source root.StandardCimv2.mlnx
//////////////////////////////////////////////
package mlnx

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// CIM_Setting struct
type CIM_Setting struct {
	*CIM_ManagedElement

	//
	SettingID string
}

func NewCIM_SettingEx1(instance *cim.WmiInstance) (newInstance *CIM_Setting, err error) {
	tmp, err := NewCIM_ManagedElementEx1(instance)

	if err != nil {
		return
	}
	newInstance = &CIM_Setting{
		CIM_ManagedElement: tmp,
	}
	return
}

func NewCIM_SettingEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *CIM_Setting, err error) {
	tmp, err := NewCIM_ManagedElementEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &CIM_Setting{
		CIM_ManagedElement: tmp,
	}
	return
}

// SetSettingID sets the value of SettingID for the instance
func (instance *CIM_Setting) SetPropertySettingID(value string) (err error) {
	return instance.SetProperty("SettingID", (value))
}

// GetSettingID gets the value of SettingID for the instance
func (instance *CIM_Setting) GetPropertySettingID() (value string, err error) {
	retValue, err := instance.GetProperty("SettingID")
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

//

// <param name="MSE" type="CIM_ManagedSystemElement "></param>
// <param name="MustBeCompletedBy" type="string "></param>
// <param name="TimeToApply" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_Setting) VerifyOKToApplyToMSE( /* IN */ MSE CIM_ManagedSystemElement,
	/* IN */ TimeToApply string,
	/* IN */ MustBeCompletedBy string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("VerifyOKToApplyToMSE", MSE, TimeToApply, MustBeCompletedBy)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="MSE" type="CIM_ManagedSystemElement "></param>
// <param name="MustBeCompletedBy" type="string "></param>
// <param name="TimeToApply" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_Setting) ApplyToMSE( /* IN */ MSE CIM_ManagedSystemElement,
	/* IN */ TimeToApply string,
	/* IN */ MustBeCompletedBy string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ApplyToMSE", MSE, TimeToApply, MustBeCompletedBy)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>
// <param name="MustBeCompletedBy" type="string "></param>
// <param name="TimeToApply" type="string "></param>

// <param name="CanNotApply" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_Setting) VerifyOKToApplyToCollection( /* IN */ Collection CIM_CollectionOfMSEs,
	/* IN */ TimeToApply string,
	/* IN */ MustBeCompletedBy string,
	/* OUT */ CanNotApply []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("VerifyOKToApplyToCollection", Collection, TimeToApply, MustBeCompletedBy)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>
// <param name="ContinueOnError" type="bool "></param>
// <param name="MustBeCompletedBy" type="string "></param>
// <param name="TimeToApply" type="string "></param>

// <param name="CanNotApply" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_Setting) ApplyToCollection( /* IN */ Collection CIM_CollectionOfMSEs,
	/* IN */ TimeToApply string,
	/* IN */ ContinueOnError bool,
	/* IN */ MustBeCompletedBy string,
	/* OUT */ CanNotApply []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ApplyToCollection", Collection, TimeToApply, ContinueOnError, MustBeCompletedBy)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="MSE" type="CIM_ManagedSystemElement "></param>
// <param name="MustBeCompletedBy" type="string "></param>
// <param name="PropertiesToApply" type="string []"></param>
// <param name="TimeToApply" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_Setting) VerifyOKToApplyIncrementalChangeToMSE( /* IN */ MSE CIM_ManagedSystemElement,
	/* IN */ TimeToApply string,
	/* IN */ MustBeCompletedBy string,
	/* IN */ PropertiesToApply []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("VerifyOKToApplyIncrementalChangeToMSE", MSE, TimeToApply, MustBeCompletedBy, PropertiesToApply)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="MSE" type="CIM_ManagedSystemElement "></param>
// <param name="MustBeCompletedBy" type="string "></param>
// <param name="PropertiesToApply" type="string []"></param>
// <param name="TimeToApply" type="string "></param>

// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_Setting) ApplyIncrementalChangeToMSE( /* IN */ MSE CIM_ManagedSystemElement,
	/* IN */ TimeToApply string,
	/* IN */ MustBeCompletedBy string,
	/* IN */ PropertiesToApply []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethodWithReturn("ApplyIncrementalChangeToMSE", MSE, TimeToApply, MustBeCompletedBy, PropertiesToApply)
	if err != nil {
		return
	}
	result = uint32(retVal)
	return

}

//

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>
// <param name="MustBeCompletedBy" type="string "></param>
// <param name="PropertiesToApply" type="string []"></param>
// <param name="TimeToApply" type="string "></param>

// <param name="CanNotApply" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_Setting) VerifyOKToApplyIncrementalChangeToCollection( /* IN */ Collection CIM_CollectionOfMSEs,
	/* IN */ TimeToApply string,
	/* IN */ MustBeCompletedBy string,
	/* IN */ PropertiesToApply []string,
	/* OUT */ CanNotApply []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("VerifyOKToApplyIncrementalChangeToCollection", Collection, TimeToApply, MustBeCompletedBy, PropertiesToApply)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}

//

// <param name="Collection" type="CIM_CollectionOfMSEs "></param>
// <param name="ContinueOnError" type="bool "></param>
// <param name="MustBeCompletedBy" type="string "></param>
// <param name="PropertiesToApply" type="string []"></param>
// <param name="TimeToApply" type="string "></param>

// <param name="CanNotApply" type="string []"></param>
// <param name="ReturnValue" type="uint32 "></param>
func (instance *CIM_Setting) ApplyIncrementalChangeToCollection( /* IN */ Collection CIM_CollectionOfMSEs,
	/* IN */ TimeToApply string,
	/* IN */ ContinueOnError bool,
	/* IN */ MustBeCompletedBy string,
	/* IN */ PropertiesToApply []string,
	/* OUT */ CanNotApply []string) (result uint32, err error) {
	retVal, err := instance.InvokeMethod("ApplyIncrementalChangeToCollection", Collection, TimeToApply, ContinueOnError, MustBeCompletedBy, PropertiesToApply)
	if err != nil {
		return
	}
	retValue := retVal[0].(int32)
	result = uint32(retValue)
	return

}
