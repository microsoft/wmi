// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source root.msdtc
//////////////////////////////////////////////
package msdtc

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/errors"
	cim "github.com/microsoft/wmi/pkg/wmiinstance"
	"reflect"
)

// DtcTransactionsStatistics struct
type DtcTransactionsStatistics struct {
	*cim.WmiInstance

	//
	Aborted uint32

	//
	AbortedMax uint32

	//
	Committed uint32

	//
	CommittedMax uint32

	//
	ForcedAbort uint32

	//
	ForcedCommit uint32

	//
	Heuristic uint32

	//
	HeuristicMax uint32

	//
	InDoubt uint32

	//
	InDoubtMax uint32

	//
	Open uint32

	//
	OpenMax uint32

	//
	SinglePhaseInDoubt uint32
}

func NewDtcTransactionsStatisticsEx1(instance *cim.WmiInstance) (newInstance *DtcTransactionsStatistics, err error) {
	tmp, err := instance, nil

	if err != nil {
		return
	}
	newInstance = &DtcTransactionsStatistics{
		WmiInstance: tmp,
	}
	return
}

func NewDtcTransactionsStatisticsEx6(hostName string,
	wmiNamespace string,
	userName string,
	password string,
	domainName string,
	query *query.WmiQuery) (newInstance *DtcTransactionsStatistics, err error) {
	tmp, err := instance.GetWmiInstance(hostName, wmiNamespace, userName, password, domainName, query)

	if err != nil {
		return
	}
	newInstance = &DtcTransactionsStatistics{
		WmiInstance: tmp,
	}
	return
}

// SetAborted sets the value of Aborted for the instance
func (instance *DtcTransactionsStatistics) SetPropertyAborted(value uint32) (err error) {
	return instance.SetProperty("Aborted", (value))
}

// GetAborted gets the value of Aborted for the instance
func (instance *DtcTransactionsStatistics) GetPropertyAborted() (value uint32, err error) {
	retValue, err := instance.GetProperty("Aborted")
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

// SetAbortedMax sets the value of AbortedMax for the instance
func (instance *DtcTransactionsStatistics) SetPropertyAbortedMax(value uint32) (err error) {
	return instance.SetProperty("AbortedMax", (value))
}

// GetAbortedMax gets the value of AbortedMax for the instance
func (instance *DtcTransactionsStatistics) GetPropertyAbortedMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("AbortedMax")
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

// SetCommitted sets the value of Committed for the instance
func (instance *DtcTransactionsStatistics) SetPropertyCommitted(value uint32) (err error) {
	return instance.SetProperty("Committed", (value))
}

// GetCommitted gets the value of Committed for the instance
func (instance *DtcTransactionsStatistics) GetPropertyCommitted() (value uint32, err error) {
	retValue, err := instance.GetProperty("Committed")
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

// SetCommittedMax sets the value of CommittedMax for the instance
func (instance *DtcTransactionsStatistics) SetPropertyCommittedMax(value uint32) (err error) {
	return instance.SetProperty("CommittedMax", (value))
}

// GetCommittedMax gets the value of CommittedMax for the instance
func (instance *DtcTransactionsStatistics) GetPropertyCommittedMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("CommittedMax")
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

// SetForcedAbort sets the value of ForcedAbort for the instance
func (instance *DtcTransactionsStatistics) SetPropertyForcedAbort(value uint32) (err error) {
	return instance.SetProperty("ForcedAbort", (value))
}

// GetForcedAbort gets the value of ForcedAbort for the instance
func (instance *DtcTransactionsStatistics) GetPropertyForcedAbort() (value uint32, err error) {
	retValue, err := instance.GetProperty("ForcedAbort")
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

// SetForcedCommit sets the value of ForcedCommit for the instance
func (instance *DtcTransactionsStatistics) SetPropertyForcedCommit(value uint32) (err error) {
	return instance.SetProperty("ForcedCommit", (value))
}

// GetForcedCommit gets the value of ForcedCommit for the instance
func (instance *DtcTransactionsStatistics) GetPropertyForcedCommit() (value uint32, err error) {
	retValue, err := instance.GetProperty("ForcedCommit")
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

// SetHeuristic sets the value of Heuristic for the instance
func (instance *DtcTransactionsStatistics) SetPropertyHeuristic(value uint32) (err error) {
	return instance.SetProperty("Heuristic", (value))
}

// GetHeuristic gets the value of Heuristic for the instance
func (instance *DtcTransactionsStatistics) GetPropertyHeuristic() (value uint32, err error) {
	retValue, err := instance.GetProperty("Heuristic")
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

// SetHeuristicMax sets the value of HeuristicMax for the instance
func (instance *DtcTransactionsStatistics) SetPropertyHeuristicMax(value uint32) (err error) {
	return instance.SetProperty("HeuristicMax", (value))
}

// GetHeuristicMax gets the value of HeuristicMax for the instance
func (instance *DtcTransactionsStatistics) GetPropertyHeuristicMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("HeuristicMax")
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

// SetInDoubt sets the value of InDoubt for the instance
func (instance *DtcTransactionsStatistics) SetPropertyInDoubt(value uint32) (err error) {
	return instance.SetProperty("InDoubt", (value))
}

// GetInDoubt gets the value of InDoubt for the instance
func (instance *DtcTransactionsStatistics) GetPropertyInDoubt() (value uint32, err error) {
	retValue, err := instance.GetProperty("InDoubt")
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

// SetInDoubtMax sets the value of InDoubtMax for the instance
func (instance *DtcTransactionsStatistics) SetPropertyInDoubtMax(value uint32) (err error) {
	return instance.SetProperty("InDoubtMax", (value))
}

// GetInDoubtMax gets the value of InDoubtMax for the instance
func (instance *DtcTransactionsStatistics) GetPropertyInDoubtMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("InDoubtMax")
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

// SetOpen sets the value of Open for the instance
func (instance *DtcTransactionsStatistics) SetPropertyOpen(value uint32) (err error) {
	return instance.SetProperty("Open", (value))
}

// GetOpen gets the value of Open for the instance
func (instance *DtcTransactionsStatistics) GetPropertyOpen() (value uint32, err error) {
	retValue, err := instance.GetProperty("Open")
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

// SetOpenMax sets the value of OpenMax for the instance
func (instance *DtcTransactionsStatistics) SetPropertyOpenMax(value uint32) (err error) {
	return instance.SetProperty("OpenMax", (value))
}

// GetOpenMax gets the value of OpenMax for the instance
func (instance *DtcTransactionsStatistics) GetPropertyOpenMax() (value uint32, err error) {
	retValue, err := instance.GetProperty("OpenMax")
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

// SetSinglePhaseInDoubt sets the value of SinglePhaseInDoubt for the instance
func (instance *DtcTransactionsStatistics) SetPropertySinglePhaseInDoubt(value uint32) (err error) {
	return instance.SetProperty("SinglePhaseInDoubt", (value))
}

// GetSinglePhaseInDoubt gets the value of SinglePhaseInDoubt for the instance
func (instance *DtcTransactionsStatistics) GetPropertySinglePhaseInDoubt() (value uint32, err error) {
	retValue, err := instance.GetProperty("SinglePhaseInDoubt")
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
