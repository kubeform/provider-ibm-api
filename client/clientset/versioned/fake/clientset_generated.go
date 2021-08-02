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

package fake

import (
	clientset "kubeform.dev/provider-ibm-api/client/clientset/versioned"
	apigatewayv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/apigateway/v1alpha1"
	fakeapigatewayv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/apigateway/v1alpha1/fake"
	appv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/app/v1alpha1"
	fakeappv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/app/v1alpha1/fake"
	cdnv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/cdn/v1alpha1"
	fakecdnv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/cdn/v1alpha1/fake"
	certificatev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/certificate/v1alpha1"
	fakecertificatev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/certificate/v1alpha1/fake"
	cisv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/cis/v1alpha1"
	fakecisv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/cis/v1alpha1/fake"
	cmv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/cm/v1alpha1"
	fakecmv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/cm/v1alpha1/fake"
	computev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/compute/v1alpha1"
	fakecomputev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/compute/v1alpha1/fake"
	containerv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/container/v1alpha1"
	fakecontainerv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/container/v1alpha1/fake"
	cosv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/cos/v1alpha1"
	fakecosv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/cos/v1alpha1/fake"
	crv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/cr/v1alpha1"
	fakecrv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/cr/v1alpha1/fake"
	databasev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/database/v1alpha1"
	fakedatabasev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/database/v1alpha1/fake"
	dlv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/dl/v1alpha1"
	fakedlv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/dl/v1alpha1/fake"
	dnsv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/dns/v1alpha1"
	fakednsv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/dns/v1alpha1/fake"
	enterprisev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/enterprise/v1alpha1"
	fakeenterprisev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/enterprise/v1alpha1/fake"
	eventv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/event/v1alpha1"
	fakeeventv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/event/v1alpha1/fake"
	firewallv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/firewall/v1alpha1"
	fakefirewallv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/firewall/v1alpha1/fake"
	functionv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/function/v1alpha1"
	fakefunctionv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/function/v1alpha1/fake"
	hardwarev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/hardware/v1alpha1"
	fakehardwarev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/hardware/v1alpha1/fake"
	hpcsv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/hpcs/v1alpha1"
	fakehpcsv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/hpcs/v1alpha1/fake"
	iamv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/iam/v1alpha1"
	fakeiamv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/iam/v1alpha1/fake"
	ipsecv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/ipsec/v1alpha1"
	fakeipsecv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/ipsec/v1alpha1/fake"
	isv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/is/v1alpha1"
	fakeisv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/is/v1alpha1/fake"
	kmsv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/kms/v1alpha1"
	fakekmsv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/kms/v1alpha1/fake"
	kpv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/kp/v1alpha1"
	fakekpv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/kp/v1alpha1/fake"
	lbv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/lb/v1alpha1"
	fakelbv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/lb/v1alpha1/fake"
	lbaasv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/lbaas/v1alpha1"
	fakelbaasv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/lbaas/v1alpha1/fake"
	multiv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/multi/v1alpha1"
	fakemultiv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/multi/v1alpha1/fake"
	networkv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/network/v1alpha1"
	fakenetworkv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/network/v1alpha1/fake"
	obv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/ob/v1alpha1"
	fakeobv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/ob/v1alpha1/fake"
	objectv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/object/v1alpha1"
	fakeobjectv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/object/v1alpha1/fake"
	orgv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/org/v1alpha1"
	fakeorgv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/org/v1alpha1/fake"
	piv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/pi/v1alpha1"
	fakepiv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/pi/v1alpha1/fake"
	pnv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/pn/v1alpha1"
	fakepnv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/pn/v1alpha1/fake"
	resourcev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/resource/v1alpha1"
	fakeresourcev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/resource/v1alpha1/fake"
	satellitev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/satellite/v1alpha1"
	fakesatellitev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/satellite/v1alpha1/fake"
	schematicsv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/schematics/v1alpha1"
	fakeschematicsv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/schematics/v1alpha1/fake"
	securityv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/security/v1alpha1"
	fakesecurityv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/security/v1alpha1/fake"
	servicev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/service/v1alpha1"
	fakeservicev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/service/v1alpha1/fake"
	spacev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/space/v1alpha1"
	fakespacev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/space/v1alpha1/fake"
	sslv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/ssl/v1alpha1"
	fakesslv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/ssl/v1alpha1/fake"
	storagev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/storage/v1alpha1"
	fakestoragev1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/storage/v1alpha1/fake"
	subnetv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/subnet/v1alpha1"
	fakesubnetv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/subnet/v1alpha1/fake"
	tgv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/tg/v1alpha1"
	faketgv1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/tg/v1alpha1/fake"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// ApigatewayV1alpha1 retrieves the ApigatewayV1alpha1Client
func (c *Clientset) ApigatewayV1alpha1() apigatewayv1alpha1.ApigatewayV1alpha1Interface {
	return &fakeapigatewayv1alpha1.FakeApigatewayV1alpha1{Fake: &c.Fake}
}

// AppV1alpha1 retrieves the AppV1alpha1Client
func (c *Clientset) AppV1alpha1() appv1alpha1.AppV1alpha1Interface {
	return &fakeappv1alpha1.FakeAppV1alpha1{Fake: &c.Fake}
}

// CdnV1alpha1 retrieves the CdnV1alpha1Client
func (c *Clientset) CdnV1alpha1() cdnv1alpha1.CdnV1alpha1Interface {
	return &fakecdnv1alpha1.FakeCdnV1alpha1{Fake: &c.Fake}
}

// CertificateV1alpha1 retrieves the CertificateV1alpha1Client
func (c *Clientset) CertificateV1alpha1() certificatev1alpha1.CertificateV1alpha1Interface {
	return &fakecertificatev1alpha1.FakeCertificateV1alpha1{Fake: &c.Fake}
}

// CisV1alpha1 retrieves the CisV1alpha1Client
func (c *Clientset) CisV1alpha1() cisv1alpha1.CisV1alpha1Interface {
	return &fakecisv1alpha1.FakeCisV1alpha1{Fake: &c.Fake}
}

// CmV1alpha1 retrieves the CmV1alpha1Client
func (c *Clientset) CmV1alpha1() cmv1alpha1.CmV1alpha1Interface {
	return &fakecmv1alpha1.FakeCmV1alpha1{Fake: &c.Fake}
}

// ComputeV1alpha1 retrieves the ComputeV1alpha1Client
func (c *Clientset) ComputeV1alpha1() computev1alpha1.ComputeV1alpha1Interface {
	return &fakecomputev1alpha1.FakeComputeV1alpha1{Fake: &c.Fake}
}

// ContainerV1alpha1 retrieves the ContainerV1alpha1Client
func (c *Clientset) ContainerV1alpha1() containerv1alpha1.ContainerV1alpha1Interface {
	return &fakecontainerv1alpha1.FakeContainerV1alpha1{Fake: &c.Fake}
}

// CosV1alpha1 retrieves the CosV1alpha1Client
func (c *Clientset) CosV1alpha1() cosv1alpha1.CosV1alpha1Interface {
	return &fakecosv1alpha1.FakeCosV1alpha1{Fake: &c.Fake}
}

// CrV1alpha1 retrieves the CrV1alpha1Client
func (c *Clientset) CrV1alpha1() crv1alpha1.CrV1alpha1Interface {
	return &fakecrv1alpha1.FakeCrV1alpha1{Fake: &c.Fake}
}

// DatabaseV1alpha1 retrieves the DatabaseV1alpha1Client
func (c *Clientset) DatabaseV1alpha1() databasev1alpha1.DatabaseV1alpha1Interface {
	return &fakedatabasev1alpha1.FakeDatabaseV1alpha1{Fake: &c.Fake}
}

// DlV1alpha1 retrieves the DlV1alpha1Client
func (c *Clientset) DlV1alpha1() dlv1alpha1.DlV1alpha1Interface {
	return &fakedlv1alpha1.FakeDlV1alpha1{Fake: &c.Fake}
}

// DnsV1alpha1 retrieves the DnsV1alpha1Client
func (c *Clientset) DnsV1alpha1() dnsv1alpha1.DnsV1alpha1Interface {
	return &fakednsv1alpha1.FakeDnsV1alpha1{Fake: &c.Fake}
}

// EnterpriseV1alpha1 retrieves the EnterpriseV1alpha1Client
func (c *Clientset) EnterpriseV1alpha1() enterprisev1alpha1.EnterpriseV1alpha1Interface {
	return &fakeenterprisev1alpha1.FakeEnterpriseV1alpha1{Fake: &c.Fake}
}

// EventV1alpha1 retrieves the EventV1alpha1Client
func (c *Clientset) EventV1alpha1() eventv1alpha1.EventV1alpha1Interface {
	return &fakeeventv1alpha1.FakeEventV1alpha1{Fake: &c.Fake}
}

// FirewallV1alpha1 retrieves the FirewallV1alpha1Client
func (c *Clientset) FirewallV1alpha1() firewallv1alpha1.FirewallV1alpha1Interface {
	return &fakefirewallv1alpha1.FakeFirewallV1alpha1{Fake: &c.Fake}
}

// FunctionV1alpha1 retrieves the FunctionV1alpha1Client
func (c *Clientset) FunctionV1alpha1() functionv1alpha1.FunctionV1alpha1Interface {
	return &fakefunctionv1alpha1.FakeFunctionV1alpha1{Fake: &c.Fake}
}

// HardwareV1alpha1 retrieves the HardwareV1alpha1Client
func (c *Clientset) HardwareV1alpha1() hardwarev1alpha1.HardwareV1alpha1Interface {
	return &fakehardwarev1alpha1.FakeHardwareV1alpha1{Fake: &c.Fake}
}

// HpcsV1alpha1 retrieves the HpcsV1alpha1Client
func (c *Clientset) HpcsV1alpha1() hpcsv1alpha1.HpcsV1alpha1Interface {
	return &fakehpcsv1alpha1.FakeHpcsV1alpha1{Fake: &c.Fake}
}

// IamV1alpha1 retrieves the IamV1alpha1Client
func (c *Clientset) IamV1alpha1() iamv1alpha1.IamV1alpha1Interface {
	return &fakeiamv1alpha1.FakeIamV1alpha1{Fake: &c.Fake}
}

// IpsecV1alpha1 retrieves the IpsecV1alpha1Client
func (c *Clientset) IpsecV1alpha1() ipsecv1alpha1.IpsecV1alpha1Interface {
	return &fakeipsecv1alpha1.FakeIpsecV1alpha1{Fake: &c.Fake}
}

// IsV1alpha1 retrieves the IsV1alpha1Client
func (c *Clientset) IsV1alpha1() isv1alpha1.IsV1alpha1Interface {
	return &fakeisv1alpha1.FakeIsV1alpha1{Fake: &c.Fake}
}

// KmsV1alpha1 retrieves the KmsV1alpha1Client
func (c *Clientset) KmsV1alpha1() kmsv1alpha1.KmsV1alpha1Interface {
	return &fakekmsv1alpha1.FakeKmsV1alpha1{Fake: &c.Fake}
}

// KpV1alpha1 retrieves the KpV1alpha1Client
func (c *Clientset) KpV1alpha1() kpv1alpha1.KpV1alpha1Interface {
	return &fakekpv1alpha1.FakeKpV1alpha1{Fake: &c.Fake}
}

// LbV1alpha1 retrieves the LbV1alpha1Client
func (c *Clientset) LbV1alpha1() lbv1alpha1.LbV1alpha1Interface {
	return &fakelbv1alpha1.FakeLbV1alpha1{Fake: &c.Fake}
}

// LbaasV1alpha1 retrieves the LbaasV1alpha1Client
func (c *Clientset) LbaasV1alpha1() lbaasv1alpha1.LbaasV1alpha1Interface {
	return &fakelbaasv1alpha1.FakeLbaasV1alpha1{Fake: &c.Fake}
}

// MultiV1alpha1 retrieves the MultiV1alpha1Client
func (c *Clientset) MultiV1alpha1() multiv1alpha1.MultiV1alpha1Interface {
	return &fakemultiv1alpha1.FakeMultiV1alpha1{Fake: &c.Fake}
}

// NetworkV1alpha1 retrieves the NetworkV1alpha1Client
func (c *Clientset) NetworkV1alpha1() networkv1alpha1.NetworkV1alpha1Interface {
	return &fakenetworkv1alpha1.FakeNetworkV1alpha1{Fake: &c.Fake}
}

// ObV1alpha1 retrieves the ObV1alpha1Client
func (c *Clientset) ObV1alpha1() obv1alpha1.ObV1alpha1Interface {
	return &fakeobv1alpha1.FakeObV1alpha1{Fake: &c.Fake}
}

// ObjectV1alpha1 retrieves the ObjectV1alpha1Client
func (c *Clientset) ObjectV1alpha1() objectv1alpha1.ObjectV1alpha1Interface {
	return &fakeobjectv1alpha1.FakeObjectV1alpha1{Fake: &c.Fake}
}

// OrgV1alpha1 retrieves the OrgV1alpha1Client
func (c *Clientset) OrgV1alpha1() orgv1alpha1.OrgV1alpha1Interface {
	return &fakeorgv1alpha1.FakeOrgV1alpha1{Fake: &c.Fake}
}

// PiV1alpha1 retrieves the PiV1alpha1Client
func (c *Clientset) PiV1alpha1() piv1alpha1.PiV1alpha1Interface {
	return &fakepiv1alpha1.FakePiV1alpha1{Fake: &c.Fake}
}

// PnV1alpha1 retrieves the PnV1alpha1Client
func (c *Clientset) PnV1alpha1() pnv1alpha1.PnV1alpha1Interface {
	return &fakepnv1alpha1.FakePnV1alpha1{Fake: &c.Fake}
}

// ResourceV1alpha1 retrieves the ResourceV1alpha1Client
func (c *Clientset) ResourceV1alpha1() resourcev1alpha1.ResourceV1alpha1Interface {
	return &fakeresourcev1alpha1.FakeResourceV1alpha1{Fake: &c.Fake}
}

// SatelliteV1alpha1 retrieves the SatelliteV1alpha1Client
func (c *Clientset) SatelliteV1alpha1() satellitev1alpha1.SatelliteV1alpha1Interface {
	return &fakesatellitev1alpha1.FakeSatelliteV1alpha1{Fake: &c.Fake}
}

// SchematicsV1alpha1 retrieves the SchematicsV1alpha1Client
func (c *Clientset) SchematicsV1alpha1() schematicsv1alpha1.SchematicsV1alpha1Interface {
	return &fakeschematicsv1alpha1.FakeSchematicsV1alpha1{Fake: &c.Fake}
}

// SecurityV1alpha1 retrieves the SecurityV1alpha1Client
func (c *Clientset) SecurityV1alpha1() securityv1alpha1.SecurityV1alpha1Interface {
	return &fakesecurityv1alpha1.FakeSecurityV1alpha1{Fake: &c.Fake}
}

// ServiceV1alpha1 retrieves the ServiceV1alpha1Client
func (c *Clientset) ServiceV1alpha1() servicev1alpha1.ServiceV1alpha1Interface {
	return &fakeservicev1alpha1.FakeServiceV1alpha1{Fake: &c.Fake}
}

// SpaceV1alpha1 retrieves the SpaceV1alpha1Client
func (c *Clientset) SpaceV1alpha1() spacev1alpha1.SpaceV1alpha1Interface {
	return &fakespacev1alpha1.FakeSpaceV1alpha1{Fake: &c.Fake}
}

// SslV1alpha1 retrieves the SslV1alpha1Client
func (c *Clientset) SslV1alpha1() sslv1alpha1.SslV1alpha1Interface {
	return &fakesslv1alpha1.FakeSslV1alpha1{Fake: &c.Fake}
}

// StorageV1alpha1 retrieves the StorageV1alpha1Client
func (c *Clientset) StorageV1alpha1() storagev1alpha1.StorageV1alpha1Interface {
	return &fakestoragev1alpha1.FakeStorageV1alpha1{Fake: &c.Fake}
}

// SubnetV1alpha1 retrieves the SubnetV1alpha1Client
func (c *Clientset) SubnetV1alpha1() subnetv1alpha1.SubnetV1alpha1Interface {
	return &fakesubnetv1alpha1.FakeSubnetV1alpha1{Fake: &c.Fake}
}

// TgV1alpha1 retrieves the TgV1alpha1Client
func (c *Clientset) TgV1alpha1() tgv1alpha1.TgV1alpha1Interface {
	return &faketgv1alpha1.FakeTgV1alpha1{Fake: &c.Fake}
}
