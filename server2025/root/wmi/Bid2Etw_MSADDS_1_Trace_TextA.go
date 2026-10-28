// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 10/28/2026 using wmigen
//      Source root.WMI
//////////////////////////////////////////////
package wmi

import (
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// Bid2Etw_MSADDS_1_Trace_TextA struct
type Bid2Etw_MSADDS_1_Trace_TextA struct {
	*Bid2Etw_MSADDS_1_Trace

	// Module ID
	ModID uint32

	// Text StringA
	msgStr interface{}
}

func NewBid2Etw_MSADDS_1_Trace_TextAEx1(instance *cim.WmiInstance) (newInstance *Bid2Etw_MSADDS_1_Trace_TextA, err error) {
	tmp, err := NewBid2Etw_MSADDS_1_TraceEx1(instance)

	if err != nil {
		return
	}
	newInstance = &Bid2Etw_MSADDS_1_Trace_TextA{
		Bid2Etw_MSADDS_1_Trace: tmp,
	}
	return
}

func NewBid2Etw_MSADDS_1_Trace_TextAEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *Bid2Etw_MSADDS_1_Trace_TextA, err error) {
	tmp, err := NewBid2Etw_MSADDS_1_TraceEx6(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &Bid2Etw_MSADDS_1_Trace_TextA{
		Bid2Etw_MSADDS_1_Trace: tmp,
	}
	return
}

// SetModID sets the value of ModID for the instance
func (instance *Bid2Etw_MSADDS_1_Trace_TextA) SetPropertyModID(value uint32) (err error) {
	return instance.SetProperty("ModID", (value))
}

// GetModID gets the value of ModID for the instance
func (instance *Bid2Etw_MSADDS_1_Trace_TextA) GetPropertyModID() (value uint32, err error) {
	retValue, err := instance.GetProperty("ModID")
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

// SetmsgStr sets the value of msgStr for the instance
func (instance *Bid2Etw_MSADDS_1_Trace_TextA) SetPropertymsgStr(value interface{}) (err error) {
	return instance.SetProperty("msgStr", (value))
}

// GetmsgStr gets the value of msgStr for the instance
func (instance *Bid2Etw_MSADDS_1_Trace_TextA) GetPropertymsgStr() (value interface{}, err error) {
	retValue, err := instance.GetProperty("msgStr")
	if err != nil {
		return
	}
	if retValue == nil {
		// Doesn't have any value. Return empty
		return
	}

	valuetmp, ok := retValue.(interface{})
	if !ok {
		err = errors.Wrapf(errors.InvalidType, " interface{} is Invalid. Expected %s", reflect.TypeOf(retValue))
		return
	}

	value = interface{}(valuetmp)

	return
}
