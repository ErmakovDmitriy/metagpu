package pkg

import pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"

type DeviceManager interface {
	CacheDevices()
	ListDevices() []*pluginapi.Device
	ListMetaDevices() []*pluginapi.Device
	ParseRealDeviceId(metaDevicesIds []string) (realDeviceId string)
}
