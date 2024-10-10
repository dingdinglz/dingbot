package bot

import (
	"github.com/LagrangeDev/LagrangeGo/client"
	"github.com/LagrangeDev/LagrangeGo/client/auth"
)

var (
	DingQQBot  *client.QQClient
	DeviceInfo *auth.DeviceInfo
	FirstLogin bool = false
)

type DeviceStruct struct {
	GUID          string `json:"guid"`
	DeviceName    string `json:"device_name"`
	SystemKernel  string `json:"system_kernel"`
	KernelVersion string `json:"kernel_version"`
}
