// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package service

import (
	"sync"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/base/instance"
	"github.com/microsoft/wmi/pkg/base/query"
	"github.com/microsoft/wmi/pkg/constant"
	"github.com/microsoft/wmi/pkg/errors"
	wmi "github.com/microsoft/wmi/pkg/wmiinstance"
	v2 "github.com/microsoft/wmi/server2019/root/virtualization/v2"
)

var (
	serviceStoreMap map[string]*SecurityService
	mux             sync.Mutex
)

func init() {
	serviceStoreMap = map[string]*SecurityService{}
}

type SecurityService struct {
	*v2.Msvm_SecurityService
}

// GetVirtualEthernetSwitchManagementService gets the VMMS Switch Management Service
func GetSecurityService(whost *host.WmiHost) (mgmt *SecurityService, err error) {
	if val, ok := serviceStoreMap[whost.HostName]; ok {
		mgmt = val
		return
	}

	mgmt, err = getService(whost)
	if err != nil {
		return
	}

	mux.Lock()
	defer mux.Unlock()
	if val, ok := serviceStoreMap[whost.HostName]; ok {
		mgmt.Close()
		mgmt = val
		return
	}

	serviceStoreMap[whost.HostName] = mgmt
	return
}

func getService(whost *host.WmiHost) (mgmt *SecurityService, err error) {
	creds := whost.GetCredential()
	query := query.NewWmiQuery("Msvm_SecurityService")
	// TODO: Regenerate wrappers that would take WmiHost directly
	vmmswmi, err := v2.NewMsvm_SecurityServiceEx6(whost.HostName, string(constant.Virtualization), creds.UserName, creds.Password, creds.Domain, query)
	if err != nil {
		return
	}

	mgmt = &SecurityService{vmmswmi}
	return
}

// SetKey
func (ss *SecurityService) SetKey(settings *wmi.WmiInstance) (err error) {
	method, err := ss.GetWmiMethod("SetKeyProtector")
	if err != nil {
		return
	}
	defer method.Close()
	embeddedInstance, err := settings.EmbeddedXMLInstance()
	if err != nil {
		return err
	}
	inparams := wmi.WmiMethodParamCollection{wmi.NewWmiMethodParam("SecuritySettingData", embeddedInstance)}
	inparams = wmi.WmiMethodParamCollection{wmi.NewWmiMethodParam("KeyProtector", embeddedInstance)}
	outparams := wmi.WmiMethodParamCollection{wmi.NewWmiMethodParam("Job", nil)}

	result, err := method.Execute(inparams, outparams)
	if err != nil {
		return
	}

	if result.ReturnValue == 0 {
		return
	}

	if result.ReturnValue != 4096 {
		err = errors.Wrapf(errors.Failed, "Method failed with [%d]", result.ReturnValue)
		return
	}

	val, ok := result.OutMethodParams["Job"]
	if !ok || val.Value == nil {
		err = errors.Wrapf(errors.NotFound, "Job")
		return
	}
	job, err := instance.GetWmiJob(ss.GetWmiHost(), string(constant.Virtualization), val.Value.(string))
	if err != nil {
		return
	}
	defer job.Close()
	return job.WaitForJobCompletion(result.ReturnValue)
}
