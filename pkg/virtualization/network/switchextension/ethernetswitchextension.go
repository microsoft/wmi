// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package switchextension

import (
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

type EthernetSwitchExtension struct {
	*v2.Msvm_EthernetSwitchExtension
}

// NewEthernetSwitchExtension
func NewEthernetSwitchExtension(instance *wmi.WmiInstance) (*EthernetSwitchExtension, error) {
	wmivm, err := v2.NewMsvm_EthernetSwitchExtensionEx1(instance)
	if err != nil {
		return nil, err
	}
	return &EthernetSwitchExtension{wmivm}, nil
}

func (iese *EthernetSwitchExtension) CloneEx1() (newese *EthernetSwitchExtension, err error) {
	tmp, err := iese.Clone()
	if err != nil {
		return
	}
	return NewEthernetSwitchExtension(tmp)
}
func (iese *EthernetSwitchExtension) GetFeatureSettingData() (wmi.WmiInstanceCollection, error) {
	instance, err := iese.GetRelated("Msvm_InstalledEthernetSwitchExtension")
	if err != nil {
		return nil, err
	}

	installedSwitchExtension, err := NewInstalledEthernetSwitchExtension(instance)
	if err != nil {
		instance.Close()
		return nil, err
	}
	defer installedSwitchExtension.Close()

	return installedSwitchExtension.GetFeatureSettingData()
}

func (ese *EthernetSwitchExtension) IsEnabled() (state bool, err error) {
	retValue, err := ese.GetProperty("EnabledState")
	if err != nil {
		return
	}
	value := retValue.(int32)

	return (value == int32(v2.EnabledLogicalElement_RequestedState_Enabled)), nil
}

func (ese *EthernetSwitchExtension) Enable() (err error) {
	return ese.EnableAsync(-1)
}

func (ese *EthernetSwitchExtension) EnableAsync(timeoutSeconds int16) (err error) {
	return ese.changeState(v2.EnabledLogicalElement_RequestedState_Enabled, timeoutSeconds)
}

func (ese *EthernetSwitchExtension) Disable() (err error) {
	return ese.DisableAsync(-1)
}

func (ese *EthernetSwitchExtension) DisableAsync(timeoutSeconds int16) (err error) {
	return ese.changeState(v2.EnabledLogicalElement_RequestedState_Disabled, timeoutSeconds)
}

//
func (ese *EthernetSwitchExtension) changeState(state v2.EnabledLogicalElement_RequestedState, timeoutSeconds int16) (err error) {
	method, err := ese.GetWmiMethod("RequestStateChange")
	if err != nil {
		return
	}
	defer method.Close()

	inparams := wmi.WmiMethodParamCollection{}
	inparams = append(inparams, wmi.NewWmiMethodParam("RequestedState", int32(state)))

	outparams := wmi.WmiMethodParamCollection{wmi.NewWmiMethodParam("Job", nil)}

	result, err := method.Execute(inparams, outparams)
	if err != nil {
		return
	}

	returnVal := result.ReturnValue
	if returnVal != 0 && returnVal != 4096 {
		err = errors.Wrapf(errors.Failed, "Method failed with [%d]", result.ReturnValue)
		return
	}

	if result.ReturnValue == 0 {
		return
	}

	val := result.OutMethodParams["Job"]
	job, err := instance.GetWmiJob(ese.GetWmiHost(), string(constant.Virtualization), val.Value.(string))
	if err != nil {
		return
	}
	defer job.Close()
	err = job.WaitForJobCompletion(result.ReturnValue, timeoutSeconds)
	return
}
