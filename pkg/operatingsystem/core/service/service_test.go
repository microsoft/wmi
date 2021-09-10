package service

import (
	"fmt"
	"testing"

	"github.com/microsoft/wmi/pkg/base/host"
	"github.com/microsoft/wmi/pkg/errors"
)

func TestWin32_Service_Exists1(t *testing.T) {
	whost := host.NewWmiLocalHost()
	win32svc, err := GetWin32Service(whost, "BFE")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Win32Service is installed\n")
	isRunning, err := win32svc.IsRunning()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Win32Service is running : %t\n", isRunning)
}

func TestWin32_Service_Exists2(t *testing.T) {
	whost := host.NewWmiLocalHost()
	win32svc, err := GetWin32Service(whost, "ClusSvc")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Win32Service is installed\n")
	isRunning, err := win32svc.IsRunning()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Win32Service is running : %t\n", isRunning)
}

func TestWin32_Service_Doesnot_Exist(t *testing.T) {
	whost := host.NewWmiLocalHost()
	_, err := GetWin32Service(whost, "InvalidService")

	if errors.IsNotFound(err) {
		fmt.Println(err)
	} else {
		t.Fatal(err)
	}
}

func TestWin32_WaitForService(t *testing.T) {
	serviceName := "DeviceInstall"
	whost := host.NewWmiLocalHost()
	win32svc1, err := GetWin32Service(whost, serviceName)
	if err != nil {
		t.Fatal(err)
	}
	defer win32svc1.Close()

	_, err = win32svc1.StopService()
	if err != nil {
		t.Fatal(err)
	}
	_, err = win32svc1.StartService()
	if err != nil {
		t.Fatal(err)
	}

	err = WaitForWin32Service(whost, serviceName, 300)
	if err != nil {
		t.Fatal(err)
	}

	win32svc2, err := GetWin32Service(whost, serviceName)
	if err != nil {
		t.Fatal(err)
	}
	defer win32svc2.Close()

	isRunning, err := win32svc2.IsRunning()
	if err != nil {
		t.Fatal(err)
	}

	if !isRunning {
		t.Fatal(err)
	}

	fmt.Printf("Win32Service is running : %t\n", isRunning)

}
