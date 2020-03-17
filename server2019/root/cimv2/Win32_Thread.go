// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 3/16/2020 using wmigen
//      Source root.CIMV2
//////////////////////////////////////////////
package cimv2

// Win32_Thread struct
type Win32_Thread struct {
	CIM_Thread

	//
	ElapsedTime uint64

	//
	PriorityBase uint32

	//
	StartAddress uint32

	//
	ThreadState uint32

	//
	ThreadWaitReason uint32
}

// SetElapsedTime sets the value of ElapsedTime for the instance
func (instance *Win32_Thread) SetPropertyElapsedTime(value uint64) (err error) {
	return instance.SetProperty("ElapsedTime", value)
}

// GetElapsedTime gets the value of ElapsedTime for the instance
func (instance *Win32_Thread) GetPropertyElapsedTime() (value uint64, err error) {
	retValue, err := instance.GetProperty("ElapsedTime")
	if err != nil {
		return
	}
	value, ok := retValue.(uint64)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetPriorityBase sets the value of PriorityBase for the instance
func (instance *Win32_Thread) SetPropertyPriorityBase(value uint32) (err error) {
	return instance.SetProperty("PriorityBase", value)
}

// GetPriorityBase gets the value of PriorityBase for the instance
func (instance *Win32_Thread) GetPropertyPriorityBase() (value uint32, err error) {
	retValue, err := instance.GetProperty("PriorityBase")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetStartAddress sets the value of StartAddress for the instance
func (instance *Win32_Thread) SetPropertyStartAddress(value uint32) (err error) {
	return instance.SetProperty("StartAddress", value)
}

// GetStartAddress gets the value of StartAddress for the instance
func (instance *Win32_Thread) GetPropertyStartAddress() (value uint32, err error) {
	retValue, err := instance.GetProperty("StartAddress")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThreadState sets the value of ThreadState for the instance
func (instance *Win32_Thread) SetPropertyThreadState(value uint32) (err error) {
	return instance.SetProperty("ThreadState", value)
}

// GetThreadState gets the value of ThreadState for the instance
func (instance *Win32_Thread) GetPropertyThreadState() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreadState")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}

// SetThreadWaitReason sets the value of ThreadWaitReason for the instance
func (instance *Win32_Thread) SetPropertyThreadWaitReason(value uint32) (err error) {
	return instance.SetProperty("ThreadWaitReason", value)
}

// GetThreadWaitReason gets the value of ThreadWaitReason for the instance
func (instance *Win32_Thread) GetPropertyThreadWaitReason() (value uint32, err error) {
	retValue, err := instance.GetProperty("ThreadWaitReason")
	if err != nil {
		return
	}
	value, ok := retValue.(uint32)
	if !ok {
		// TODO: Set an error
	}
	return
}
