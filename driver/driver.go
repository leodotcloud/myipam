package driver

import (
	"github.com/Sirupsen/logrus"
	"github.com/docker/go-plugins-helpers/ipam"
)

type RancherIPAMDriver struct {
}

func init() {
	//TODO: Take this as a flag during startup
	logrus.SetLevel(logrus.DebugLevel)
}

func (d *RancherIPAMDriver) GetCapabilities() (*ipam.CapabilitiesResponse, error) {
	logrus.Info("Endpoint hit: GetCapabilities")

	resp := &ipam.CapabilitiesResponse{RequiresMACAddress: true}
	logrus.Debugf("CapabilitiesResponse: %+v", resp)

	return resp, nil
}

func (d *RancherIPAMDriver) GetDefaultAddressSpaces() (*ipam.AddressSpacesResponse, error) {
	logrus.Info("Endpoint hit: GetDefaultAddressSpaces")

	resp := &ipam.AddressSpacesResponse{
		LocalDefaultAddressSpace:  "rancher-local",
		GlobalDefaultAddressSpace: "rancher-global",
	}
	logrus.Debugf("AddressSpacesResponse: %+v", resp)

	return resp, nil
}

func (d *RancherIPAMDriver) RequestPool(r *ipam.RequestPoolRequest) (*ipam.RequestPoolResponse, error) {
	logrus.Info("Endpoint hit: RequestPool")
	logrus.Debugf("RequestPoolRequest: %+v", r)

	resp := &ipam.RequestPoolResponse{
		PoolID: "racher-ipam-poolid",
		Pool:   "10.42.0.0/16",
	}
	logrus.Debugf("RequestAddressResponse: %+v", resp)

	return resp, nil
}

func (d *RancherIPAMDriver) ReleasePool(r *ipam.ReleasePoolRequest) error {
	logrus.Info("Endpoint hit: ReleasePool")
	logrus.Debugf("ReleasePoolRequest: %+v", r)

	return nil
}

func (d *RancherIPAMDriver) RequestAddress(r *ipam.RequestAddressRequest) (*ipam.RequestAddressResponse, error) {
	logrus.Info("Endpoint hit: RequestAddress")
	logrus.Debugf("RequestAddressRequest: %+v", r)

	var resp *ipam.RequestAddressResponse

	if r.Options["RequestAddressType"] == "com.docker.network.gateway" {
		resp = &ipam.RequestAddressResponse{
			Address: "10.42.1.1/16",
		}
	} else {
		resp = &ipam.RequestAddressResponse{
			Address: "10.42.1.100/16",
		}
	}

	logrus.Debugf("RequestAddressResponse: %+v", resp)
	return resp, nil
}

func (d *RancherIPAMDriver) ReleaseAddress(r *ipam.ReleaseAddressRequest) error {
	logrus.Info("Endpoint hit: ReleaseAddress")
	logrus.Debugf("ReleaseAddressRequest: %+v", r)
	return nil
}
