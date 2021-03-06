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
	v1alpha1 "kubeform.dev/provider-ibm-api/client/clientset/versioned/typed/cis/v1alpha1"

	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCisV1alpha1 struct {
	*testing.Fake
}

func (c *FakeCisV1alpha1) CacheSettingses(namespace string) v1alpha1.CacheSettingsInterface {
	return &FakeCacheSettingses{c, namespace}
}

func (c *FakeCisV1alpha1) CertificateOrders(namespace string) v1alpha1.CertificateOrderInterface {
	return &FakeCertificateOrders{c, namespace}
}

func (c *FakeCisV1alpha1) CertificateUploads(namespace string) v1alpha1.CertificateUploadInterface {
	return &FakeCertificateUploads{c, namespace}
}

func (c *FakeCisV1alpha1) Cises(namespace string) v1alpha1.CisInterface {
	return &FakeCises{c, namespace}
}

func (c *FakeCisV1alpha1) CustomPages(namespace string) v1alpha1.CustomPageInterface {
	return &FakeCustomPages{c, namespace}
}

func (c *FakeCisV1alpha1) DnsRecords(namespace string) v1alpha1.DnsRecordInterface {
	return &FakeDnsRecords{c, namespace}
}

func (c *FakeCisV1alpha1) DnsRecordsImports(namespace string) v1alpha1.DnsRecordsImportInterface {
	return &FakeDnsRecordsImports{c, namespace}
}

func (c *FakeCisV1alpha1) Domains(namespace string) v1alpha1.DomainInterface {
	return &FakeDomains{c, namespace}
}

func (c *FakeCisV1alpha1) DomainSettingses(namespace string) v1alpha1.DomainSettingsInterface {
	return &FakeDomainSettingses{c, namespace}
}

func (c *FakeCisV1alpha1) EdgeFunctionsActions(namespace string) v1alpha1.EdgeFunctionsActionInterface {
	return &FakeEdgeFunctionsActions{c, namespace}
}

func (c *FakeCisV1alpha1) EdgeFunctionsTriggers(namespace string) v1alpha1.EdgeFunctionsTriggerInterface {
	return &FakeEdgeFunctionsTriggers{c, namespace}
}

func (c *FakeCisV1alpha1) Filters(namespace string) v1alpha1.FilterInterface {
	return &FakeFilters{c, namespace}
}

func (c *FakeCisV1alpha1) Firewalls(namespace string) v1alpha1.FirewallInterface {
	return &FakeFirewalls{c, namespace}
}

func (c *FakeCisV1alpha1) GlobalLoadBalancers(namespace string) v1alpha1.GlobalLoadBalancerInterface {
	return &FakeGlobalLoadBalancers{c, namespace}
}

func (c *FakeCisV1alpha1) Healthchecks(namespace string) v1alpha1.HealthcheckInterface {
	return &FakeHealthchecks{c, namespace}
}

func (c *FakeCisV1alpha1) OriginPools(namespace string) v1alpha1.OriginPoolInterface {
	return &FakeOriginPools{c, namespace}
}

func (c *FakeCisV1alpha1) PageRules(namespace string) v1alpha1.PageRuleInterface {
	return &FakePageRules{c, namespace}
}

func (c *FakeCisV1alpha1) RangeApps(namespace string) v1alpha1.RangeAppInterface {
	return &FakeRangeApps{c, namespace}
}

func (c *FakeCisV1alpha1) RateLimits(namespace string) v1alpha1.RateLimitInterface {
	return &FakeRateLimits{c, namespace}
}

func (c *FakeCisV1alpha1) Routings(namespace string) v1alpha1.RoutingInterface {
	return &FakeRoutings{c, namespace}
}

func (c *FakeCisV1alpha1) TlsSettingses(namespace string) v1alpha1.TlsSettingsInterface {
	return &FakeTlsSettingses{c, namespace}
}

func (c *FakeCisV1alpha1) WafGroups(namespace string) v1alpha1.WafGroupInterface {
	return &FakeWafGroups{c, namespace}
}

func (c *FakeCisV1alpha1) WafPackages(namespace string) v1alpha1.WafPackageInterface {
	return &FakeWafPackages{c, namespace}
}

func (c *FakeCisV1alpha1) WafRules(namespace string) v1alpha1.WafRuleInterface {
	return &FakeWafRules{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCisV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
