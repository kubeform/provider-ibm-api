/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"

	apigatewayv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/apigateway/v1alpha1"
	appv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/app/v1alpha1"
	cdnv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/cdn/v1alpha1"
	certificatev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/certificate/v1alpha1"
	cisv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/cis/v1alpha1"
	cmv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/cm/v1alpha1"
	computev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/compute/v1alpha1"
	containerv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/container/v1alpha1"
	cosv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/cos/v1alpha1"
	crv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/cr/v1alpha1"
	databasev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/database/v1alpha1"
	dlv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/dl/v1alpha1"
	dnsv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/dns/v1alpha1"
	enterprisev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/enterprise/v1alpha1"
	eventv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/event/v1alpha1"
	firewallv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/firewall/v1alpha1"
	functionv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/function/v1alpha1"
	hardwarev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/hardware/v1alpha1"
	hpcsv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/hpcs/v1alpha1"
	iamv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/iam/v1alpha1"
	ipsecv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/ipsec/v1alpha1"
	isv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/is/v1alpha1"
	kmsv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/kms/v1alpha1"
	kpv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/kp/v1alpha1"
	lbv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/lb/v1alpha1"
	lbaasv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/lbaas/v1alpha1"
	multiv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/multi/v1alpha1"
	networkv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/network/v1alpha1"
	obv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/ob/v1alpha1"
	objectv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/object/v1alpha1"
	orgv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/org/v1alpha1"
	piv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/pi/v1alpha1"
	pnv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/pn/v1alpha1"
	resourcev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/resource/v1alpha1"
	satellitev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/satellite/v1alpha1"
	schematicsv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/schematics/v1alpha1"
	securityv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/security/v1alpha1"
	servicev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/service/v1alpha1"
	spacev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/space/v1alpha1"
	sslv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/ssl/v1alpha1"
	storagev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/storage/v1alpha1"
	subnetv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/subnet/v1alpha1"
	tgv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/tg/v1alpha1"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ApigatewayV1alpha1() apigatewayv1alpha1.ApigatewayV1alpha1Interface
	AppV1alpha1() appv1alpha1.AppV1alpha1Interface
	CdnV1alpha1() cdnv1alpha1.CdnV1alpha1Interface
	CertificateV1alpha1() certificatev1alpha1.CertificateV1alpha1Interface
	CisV1alpha1() cisv1alpha1.CisV1alpha1Interface
	CmV1alpha1() cmv1alpha1.CmV1alpha1Interface
	ComputeV1alpha1() computev1alpha1.ComputeV1alpha1Interface
	ContainerV1alpha1() containerv1alpha1.ContainerV1alpha1Interface
	CosV1alpha1() cosv1alpha1.CosV1alpha1Interface
	CrV1alpha1() crv1alpha1.CrV1alpha1Interface
	DatabaseV1alpha1() databasev1alpha1.DatabaseV1alpha1Interface
	DlV1alpha1() dlv1alpha1.DlV1alpha1Interface
	DnsV1alpha1() dnsv1alpha1.DnsV1alpha1Interface
	EnterpriseV1alpha1() enterprisev1alpha1.EnterpriseV1alpha1Interface
	EventV1alpha1() eventv1alpha1.EventV1alpha1Interface
	FirewallV1alpha1() firewallv1alpha1.FirewallV1alpha1Interface
	FunctionV1alpha1() functionv1alpha1.FunctionV1alpha1Interface
	HardwareV1alpha1() hardwarev1alpha1.HardwareV1alpha1Interface
	HpcsV1alpha1() hpcsv1alpha1.HpcsV1alpha1Interface
	IamV1alpha1() iamv1alpha1.IamV1alpha1Interface
	IpsecV1alpha1() ipsecv1alpha1.IpsecV1alpha1Interface
	IsV1alpha1() isv1alpha1.IsV1alpha1Interface
	KmsV1alpha1() kmsv1alpha1.KmsV1alpha1Interface
	KpV1alpha1() kpv1alpha1.KpV1alpha1Interface
	LbV1alpha1() lbv1alpha1.LbV1alpha1Interface
	LbaasV1alpha1() lbaasv1alpha1.LbaasV1alpha1Interface
	MultiV1alpha1() multiv1alpha1.MultiV1alpha1Interface
	NetworkV1alpha1() networkv1alpha1.NetworkV1alpha1Interface
	ObV1alpha1() obv1alpha1.ObV1alpha1Interface
	ObjectV1alpha1() objectv1alpha1.ObjectV1alpha1Interface
	OrgV1alpha1() orgv1alpha1.OrgV1alpha1Interface
	PiV1alpha1() piv1alpha1.PiV1alpha1Interface
	PnV1alpha1() pnv1alpha1.PnV1alpha1Interface
	ResourceV1alpha1() resourcev1alpha1.ResourceV1alpha1Interface
	SatelliteV1alpha1() satellitev1alpha1.SatelliteV1alpha1Interface
	SchematicsV1alpha1() schematicsv1alpha1.SchematicsV1alpha1Interface
	SecurityV1alpha1() securityv1alpha1.SecurityV1alpha1Interface
	ServiceV1alpha1() servicev1alpha1.ServiceV1alpha1Interface
	SpaceV1alpha1() spacev1alpha1.SpaceV1alpha1Interface
	SslV1alpha1() sslv1alpha1.SslV1alpha1Interface
	StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface
	SubnetV1alpha1() subnetv1alpha1.SubnetV1alpha1Interface
	TgV1alpha1() tgv1alpha1.TgV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	apigatewayV1alpha1  *apigatewayv1alpha1.ApigatewayV1alpha1Client
	appV1alpha1         *appv1alpha1.AppV1alpha1Client
	cdnV1alpha1         *cdnv1alpha1.CdnV1alpha1Client
	certificateV1alpha1 *certificatev1alpha1.CertificateV1alpha1Client
	cisV1alpha1         *cisv1alpha1.CisV1alpha1Client
	cmV1alpha1          *cmv1alpha1.CmV1alpha1Client
	computeV1alpha1     *computev1alpha1.ComputeV1alpha1Client
	containerV1alpha1   *containerv1alpha1.ContainerV1alpha1Client
	cosV1alpha1         *cosv1alpha1.CosV1alpha1Client
	crV1alpha1          *crv1alpha1.CrV1alpha1Client
	databaseV1alpha1    *databasev1alpha1.DatabaseV1alpha1Client
	dlV1alpha1          *dlv1alpha1.DlV1alpha1Client
	dnsV1alpha1         *dnsv1alpha1.DnsV1alpha1Client
	enterpriseV1alpha1  *enterprisev1alpha1.EnterpriseV1alpha1Client
	eventV1alpha1       *eventv1alpha1.EventV1alpha1Client
	firewallV1alpha1    *firewallv1alpha1.FirewallV1alpha1Client
	functionV1alpha1    *functionv1alpha1.FunctionV1alpha1Client
	hardwareV1alpha1    *hardwarev1alpha1.HardwareV1alpha1Client
	hpcsV1alpha1        *hpcsv1alpha1.HpcsV1alpha1Client
	iamV1alpha1         *iamv1alpha1.IamV1alpha1Client
	ipsecV1alpha1       *ipsecv1alpha1.IpsecV1alpha1Client
	isV1alpha1          *isv1alpha1.IsV1alpha1Client
	kmsV1alpha1         *kmsv1alpha1.KmsV1alpha1Client
	kpV1alpha1          *kpv1alpha1.KpV1alpha1Client
	lbV1alpha1          *lbv1alpha1.LbV1alpha1Client
	lbaasV1alpha1       *lbaasv1alpha1.LbaasV1alpha1Client
	multiV1alpha1       *multiv1alpha1.MultiV1alpha1Client
	networkV1alpha1     *networkv1alpha1.NetworkV1alpha1Client
	obV1alpha1          *obv1alpha1.ObV1alpha1Client
	objectV1alpha1      *objectv1alpha1.ObjectV1alpha1Client
	orgV1alpha1         *orgv1alpha1.OrgV1alpha1Client
	piV1alpha1          *piv1alpha1.PiV1alpha1Client
	pnV1alpha1          *pnv1alpha1.PnV1alpha1Client
	resourceV1alpha1    *resourcev1alpha1.ResourceV1alpha1Client
	satelliteV1alpha1   *satellitev1alpha1.SatelliteV1alpha1Client
	schematicsV1alpha1  *schematicsv1alpha1.SchematicsV1alpha1Client
	securityV1alpha1    *securityv1alpha1.SecurityV1alpha1Client
	serviceV1alpha1     *servicev1alpha1.ServiceV1alpha1Client
	spaceV1alpha1       *spacev1alpha1.SpaceV1alpha1Client
	sslV1alpha1         *sslv1alpha1.SslV1alpha1Client
	storageV1alpha1     *storagev1alpha1.StorageV1alpha1Client
	subnetV1alpha1      *subnetv1alpha1.SubnetV1alpha1Client
	tgV1alpha1          *tgv1alpha1.TgV1alpha1Client
}

// ApigatewayV1alpha1 retrieves the ApigatewayV1alpha1Client
func (c *Clientset) ApigatewayV1alpha1() apigatewayv1alpha1.ApigatewayV1alpha1Interface {
	return c.apigatewayV1alpha1
}

// AppV1alpha1 retrieves the AppV1alpha1Client
func (c *Clientset) AppV1alpha1() appv1alpha1.AppV1alpha1Interface {
	return c.appV1alpha1
}

// CdnV1alpha1 retrieves the CdnV1alpha1Client
func (c *Clientset) CdnV1alpha1() cdnv1alpha1.CdnV1alpha1Interface {
	return c.cdnV1alpha1
}

// CertificateV1alpha1 retrieves the CertificateV1alpha1Client
func (c *Clientset) CertificateV1alpha1() certificatev1alpha1.CertificateV1alpha1Interface {
	return c.certificateV1alpha1
}

// CisV1alpha1 retrieves the CisV1alpha1Client
func (c *Clientset) CisV1alpha1() cisv1alpha1.CisV1alpha1Interface {
	return c.cisV1alpha1
}

// CmV1alpha1 retrieves the CmV1alpha1Client
func (c *Clientset) CmV1alpha1() cmv1alpha1.CmV1alpha1Interface {
	return c.cmV1alpha1
}

// ComputeV1alpha1 retrieves the ComputeV1alpha1Client
func (c *Clientset) ComputeV1alpha1() computev1alpha1.ComputeV1alpha1Interface {
	return c.computeV1alpha1
}

// ContainerV1alpha1 retrieves the ContainerV1alpha1Client
func (c *Clientset) ContainerV1alpha1() containerv1alpha1.ContainerV1alpha1Interface {
	return c.containerV1alpha1
}

// CosV1alpha1 retrieves the CosV1alpha1Client
func (c *Clientset) CosV1alpha1() cosv1alpha1.CosV1alpha1Interface {
	return c.cosV1alpha1
}

// CrV1alpha1 retrieves the CrV1alpha1Client
func (c *Clientset) CrV1alpha1() crv1alpha1.CrV1alpha1Interface {
	return c.crV1alpha1
}

// DatabaseV1alpha1 retrieves the DatabaseV1alpha1Client
func (c *Clientset) DatabaseV1alpha1() databasev1alpha1.DatabaseV1alpha1Interface {
	return c.databaseV1alpha1
}

// DlV1alpha1 retrieves the DlV1alpha1Client
func (c *Clientset) DlV1alpha1() dlv1alpha1.DlV1alpha1Interface {
	return c.dlV1alpha1
}

// DnsV1alpha1 retrieves the DnsV1alpha1Client
func (c *Clientset) DnsV1alpha1() dnsv1alpha1.DnsV1alpha1Interface {
	return c.dnsV1alpha1
}

// EnterpriseV1alpha1 retrieves the EnterpriseV1alpha1Client
func (c *Clientset) EnterpriseV1alpha1() enterprisev1alpha1.EnterpriseV1alpha1Interface {
	return c.enterpriseV1alpha1
}

// EventV1alpha1 retrieves the EventV1alpha1Client
func (c *Clientset) EventV1alpha1() eventv1alpha1.EventV1alpha1Interface {
	return c.eventV1alpha1
}

// FirewallV1alpha1 retrieves the FirewallV1alpha1Client
func (c *Clientset) FirewallV1alpha1() firewallv1alpha1.FirewallV1alpha1Interface {
	return c.firewallV1alpha1
}

// FunctionV1alpha1 retrieves the FunctionV1alpha1Client
func (c *Clientset) FunctionV1alpha1() functionv1alpha1.FunctionV1alpha1Interface {
	return c.functionV1alpha1
}

// HardwareV1alpha1 retrieves the HardwareV1alpha1Client
func (c *Clientset) HardwareV1alpha1() hardwarev1alpha1.HardwareV1alpha1Interface {
	return c.hardwareV1alpha1
}

// HpcsV1alpha1 retrieves the HpcsV1alpha1Client
func (c *Clientset) HpcsV1alpha1() hpcsv1alpha1.HpcsV1alpha1Interface {
	return c.hpcsV1alpha1
}

// IamV1alpha1 retrieves the IamV1alpha1Client
func (c *Clientset) IamV1alpha1() iamv1alpha1.IamV1alpha1Interface {
	return c.iamV1alpha1
}

// IpsecV1alpha1 retrieves the IpsecV1alpha1Client
func (c *Clientset) IpsecV1alpha1() ipsecv1alpha1.IpsecV1alpha1Interface {
	return c.ipsecV1alpha1
}

// IsV1alpha1 retrieves the IsV1alpha1Client
func (c *Clientset) IsV1alpha1() isv1alpha1.IsV1alpha1Interface {
	return c.isV1alpha1
}

// KmsV1alpha1 retrieves the KmsV1alpha1Client
func (c *Clientset) KmsV1alpha1() kmsv1alpha1.KmsV1alpha1Interface {
	return c.kmsV1alpha1
}

// KpV1alpha1 retrieves the KpV1alpha1Client
func (c *Clientset) KpV1alpha1() kpv1alpha1.KpV1alpha1Interface {
	return c.kpV1alpha1
}

// LbV1alpha1 retrieves the LbV1alpha1Client
func (c *Clientset) LbV1alpha1() lbv1alpha1.LbV1alpha1Interface {
	return c.lbV1alpha1
}

// LbaasV1alpha1 retrieves the LbaasV1alpha1Client
func (c *Clientset) LbaasV1alpha1() lbaasv1alpha1.LbaasV1alpha1Interface {
	return c.lbaasV1alpha1
}

// MultiV1alpha1 retrieves the MultiV1alpha1Client
func (c *Clientset) MultiV1alpha1() multiv1alpha1.MultiV1alpha1Interface {
	return c.multiV1alpha1
}

// NetworkV1alpha1 retrieves the NetworkV1alpha1Client
func (c *Clientset) NetworkV1alpha1() networkv1alpha1.NetworkV1alpha1Interface {
	return c.networkV1alpha1
}

// ObV1alpha1 retrieves the ObV1alpha1Client
func (c *Clientset) ObV1alpha1() obv1alpha1.ObV1alpha1Interface {
	return c.obV1alpha1
}

// ObjectV1alpha1 retrieves the ObjectV1alpha1Client
func (c *Clientset) ObjectV1alpha1() objectv1alpha1.ObjectV1alpha1Interface {
	return c.objectV1alpha1
}

// OrgV1alpha1 retrieves the OrgV1alpha1Client
func (c *Clientset) OrgV1alpha1() orgv1alpha1.OrgV1alpha1Interface {
	return c.orgV1alpha1
}

// PiV1alpha1 retrieves the PiV1alpha1Client
func (c *Clientset) PiV1alpha1() piv1alpha1.PiV1alpha1Interface {
	return c.piV1alpha1
}

// PnV1alpha1 retrieves the PnV1alpha1Client
func (c *Clientset) PnV1alpha1() pnv1alpha1.PnV1alpha1Interface {
	return c.pnV1alpha1
}

// ResourceV1alpha1 retrieves the ResourceV1alpha1Client
func (c *Clientset) ResourceV1alpha1() resourcev1alpha1.ResourceV1alpha1Interface {
	return c.resourceV1alpha1
}

// SatelliteV1alpha1 retrieves the SatelliteV1alpha1Client
func (c *Clientset) SatelliteV1alpha1() satellitev1alpha1.SatelliteV1alpha1Interface {
	return c.satelliteV1alpha1
}

// SchematicsV1alpha1 retrieves the SchematicsV1alpha1Client
func (c *Clientset) SchematicsV1alpha1() schematicsv1alpha1.SchematicsV1alpha1Interface {
	return c.schematicsV1alpha1
}

// SecurityV1alpha1 retrieves the SecurityV1alpha1Client
func (c *Clientset) SecurityV1alpha1() securityv1alpha1.SecurityV1alpha1Interface {
	return c.securityV1alpha1
}

// ServiceV1alpha1 retrieves the ServiceV1alpha1Client
func (c *Clientset) ServiceV1alpha1() servicev1alpha1.ServiceV1alpha1Interface {
	return c.serviceV1alpha1
}

// SpaceV1alpha1 retrieves the SpaceV1alpha1Client
func (c *Clientset) SpaceV1alpha1() spacev1alpha1.SpaceV1alpha1Interface {
	return c.spaceV1alpha1
}

// SslV1alpha1 retrieves the SslV1alpha1Client
func (c *Clientset) SslV1alpha1() sslv1alpha1.SslV1alpha1Interface {
	return c.sslV1alpha1
}

// StorageV1alpha1 retrieves the StorageV1alpha1Client
func (c *Clientset) StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface {
	return c.storageV1alpha1
}

// SubnetV1alpha1 retrieves the SubnetV1alpha1Client
func (c *Clientset) SubnetV1alpha1() subnetv1alpha1.SubnetV1alpha1Interface {
	return c.subnetV1alpha1
}

// TgV1alpha1 retrieves the TgV1alpha1Client
func (c *Clientset) TgV1alpha1() tgv1alpha1.TgV1alpha1Interface {
	return c.tgV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.apigatewayV1alpha1, err = apigatewayv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.appV1alpha1, err = appv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.cdnV1alpha1, err = cdnv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.certificateV1alpha1, err = certificatev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.cisV1alpha1, err = cisv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.cmV1alpha1, err = cmv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.computeV1alpha1, err = computev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.containerV1alpha1, err = containerv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.cosV1alpha1, err = cosv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.crV1alpha1, err = crv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.databaseV1alpha1, err = databasev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.dlV1alpha1, err = dlv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.dnsV1alpha1, err = dnsv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.enterpriseV1alpha1, err = enterprisev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.eventV1alpha1, err = eventv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.firewallV1alpha1, err = firewallv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.functionV1alpha1, err = functionv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.hardwareV1alpha1, err = hardwarev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.hpcsV1alpha1, err = hpcsv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.iamV1alpha1, err = iamv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.ipsecV1alpha1, err = ipsecv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.isV1alpha1, err = isv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.kmsV1alpha1, err = kmsv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.kpV1alpha1, err = kpv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.lbV1alpha1, err = lbv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.lbaasV1alpha1, err = lbaasv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.multiV1alpha1, err = multiv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.networkV1alpha1, err = networkv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.obV1alpha1, err = obv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.objectV1alpha1, err = objectv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.orgV1alpha1, err = orgv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.piV1alpha1, err = piv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.pnV1alpha1, err = pnv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.resourceV1alpha1, err = resourcev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.satelliteV1alpha1, err = satellitev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.schematicsV1alpha1, err = schematicsv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.securityV1alpha1, err = securityv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.serviceV1alpha1, err = servicev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.spaceV1alpha1, err = spacev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.sslV1alpha1, err = sslv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.storageV1alpha1, err = storagev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.subnetV1alpha1, err = subnetv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.tgV1alpha1, err = tgv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.apigatewayV1alpha1 = apigatewayv1alpha1.NewForConfigOrDie(c)
	cs.appV1alpha1 = appv1alpha1.NewForConfigOrDie(c)
	cs.cdnV1alpha1 = cdnv1alpha1.NewForConfigOrDie(c)
	cs.certificateV1alpha1 = certificatev1alpha1.NewForConfigOrDie(c)
	cs.cisV1alpha1 = cisv1alpha1.NewForConfigOrDie(c)
	cs.cmV1alpha1 = cmv1alpha1.NewForConfigOrDie(c)
	cs.computeV1alpha1 = computev1alpha1.NewForConfigOrDie(c)
	cs.containerV1alpha1 = containerv1alpha1.NewForConfigOrDie(c)
	cs.cosV1alpha1 = cosv1alpha1.NewForConfigOrDie(c)
	cs.crV1alpha1 = crv1alpha1.NewForConfigOrDie(c)
	cs.databaseV1alpha1 = databasev1alpha1.NewForConfigOrDie(c)
	cs.dlV1alpha1 = dlv1alpha1.NewForConfigOrDie(c)
	cs.dnsV1alpha1 = dnsv1alpha1.NewForConfigOrDie(c)
	cs.enterpriseV1alpha1 = enterprisev1alpha1.NewForConfigOrDie(c)
	cs.eventV1alpha1 = eventv1alpha1.NewForConfigOrDie(c)
	cs.firewallV1alpha1 = firewallv1alpha1.NewForConfigOrDie(c)
	cs.functionV1alpha1 = functionv1alpha1.NewForConfigOrDie(c)
	cs.hardwareV1alpha1 = hardwarev1alpha1.NewForConfigOrDie(c)
	cs.hpcsV1alpha1 = hpcsv1alpha1.NewForConfigOrDie(c)
	cs.iamV1alpha1 = iamv1alpha1.NewForConfigOrDie(c)
	cs.ipsecV1alpha1 = ipsecv1alpha1.NewForConfigOrDie(c)
	cs.isV1alpha1 = isv1alpha1.NewForConfigOrDie(c)
	cs.kmsV1alpha1 = kmsv1alpha1.NewForConfigOrDie(c)
	cs.kpV1alpha1 = kpv1alpha1.NewForConfigOrDie(c)
	cs.lbV1alpha1 = lbv1alpha1.NewForConfigOrDie(c)
	cs.lbaasV1alpha1 = lbaasv1alpha1.NewForConfigOrDie(c)
	cs.multiV1alpha1 = multiv1alpha1.NewForConfigOrDie(c)
	cs.networkV1alpha1 = networkv1alpha1.NewForConfigOrDie(c)
	cs.obV1alpha1 = obv1alpha1.NewForConfigOrDie(c)
	cs.objectV1alpha1 = objectv1alpha1.NewForConfigOrDie(c)
	cs.orgV1alpha1 = orgv1alpha1.NewForConfigOrDie(c)
	cs.piV1alpha1 = piv1alpha1.NewForConfigOrDie(c)
	cs.pnV1alpha1 = pnv1alpha1.NewForConfigOrDie(c)
	cs.resourceV1alpha1 = resourcev1alpha1.NewForConfigOrDie(c)
	cs.satelliteV1alpha1 = satellitev1alpha1.NewForConfigOrDie(c)
	cs.schematicsV1alpha1 = schematicsv1alpha1.NewForConfigOrDie(c)
	cs.securityV1alpha1 = securityv1alpha1.NewForConfigOrDie(c)
	cs.serviceV1alpha1 = servicev1alpha1.NewForConfigOrDie(c)
	cs.spaceV1alpha1 = spacev1alpha1.NewForConfigOrDie(c)
	cs.sslV1alpha1 = sslv1alpha1.NewForConfigOrDie(c)
	cs.storageV1alpha1 = storagev1alpha1.NewForConfigOrDie(c)
	cs.subnetV1alpha1 = subnetv1alpha1.NewForConfigOrDie(c)
	cs.tgV1alpha1 = tgv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.apigatewayV1alpha1 = apigatewayv1alpha1.New(c)
	cs.appV1alpha1 = appv1alpha1.New(c)
	cs.cdnV1alpha1 = cdnv1alpha1.New(c)
	cs.certificateV1alpha1 = certificatev1alpha1.New(c)
	cs.cisV1alpha1 = cisv1alpha1.New(c)
	cs.cmV1alpha1 = cmv1alpha1.New(c)
	cs.computeV1alpha1 = computev1alpha1.New(c)
	cs.containerV1alpha1 = containerv1alpha1.New(c)
	cs.cosV1alpha1 = cosv1alpha1.New(c)
	cs.crV1alpha1 = crv1alpha1.New(c)
	cs.databaseV1alpha1 = databasev1alpha1.New(c)
	cs.dlV1alpha1 = dlv1alpha1.New(c)
	cs.dnsV1alpha1 = dnsv1alpha1.New(c)
	cs.enterpriseV1alpha1 = enterprisev1alpha1.New(c)
	cs.eventV1alpha1 = eventv1alpha1.New(c)
	cs.firewallV1alpha1 = firewallv1alpha1.New(c)
	cs.functionV1alpha1 = functionv1alpha1.New(c)
	cs.hardwareV1alpha1 = hardwarev1alpha1.New(c)
	cs.hpcsV1alpha1 = hpcsv1alpha1.New(c)
	cs.iamV1alpha1 = iamv1alpha1.New(c)
	cs.ipsecV1alpha1 = ipsecv1alpha1.New(c)
	cs.isV1alpha1 = isv1alpha1.New(c)
	cs.kmsV1alpha1 = kmsv1alpha1.New(c)
	cs.kpV1alpha1 = kpv1alpha1.New(c)
	cs.lbV1alpha1 = lbv1alpha1.New(c)
	cs.lbaasV1alpha1 = lbaasv1alpha1.New(c)
	cs.multiV1alpha1 = multiv1alpha1.New(c)
	cs.networkV1alpha1 = networkv1alpha1.New(c)
	cs.obV1alpha1 = obv1alpha1.New(c)
	cs.objectV1alpha1 = objectv1alpha1.New(c)
	cs.orgV1alpha1 = orgv1alpha1.New(c)
	cs.piV1alpha1 = piv1alpha1.New(c)
	cs.pnV1alpha1 = pnv1alpha1.New(c)
	cs.resourceV1alpha1 = resourcev1alpha1.New(c)
	cs.satelliteV1alpha1 = satellitev1alpha1.New(c)
	cs.schematicsV1alpha1 = schematicsv1alpha1.New(c)
	cs.securityV1alpha1 = securityv1alpha1.New(c)
	cs.serviceV1alpha1 = servicev1alpha1.New(c)
	cs.spaceV1alpha1 = spacev1alpha1.New(c)
	cs.sslV1alpha1 = sslv1alpha1.New(c)
	cs.storageV1alpha1 = storagev1alpha1.New(c)
	cs.subnetV1alpha1 = subnetv1alpha1.New(c)
	cs.tgV1alpha1 = tgv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}