---
layout: "avi"
page_title: "Provider: AVI"
sidebar_current: "docs-avi-index"
description: |-
  The AVI provider is used to interact with the many resources supported by AVI. The provider needs to be configured with the proper credentials before it can be used.
---

# AVI Provider

The AVI provider is used to interact with [AVI Controller](https://avinetworks.com/).
The provider needs
to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

# Integrating Terraform with Avi Vantage
Each Avi Vantage REST resource is exposed as a resource in Terraform. For instance, you can setup a pool by using avi_pool as the resource type in Terraform. There are more than 50 different REST resources supported.

Each Terraform resource maps directly to a corresponding Avi Vantage API as defined in the Swagger Spec (available on any Avi Vantage Controller at https://<controller>/swagger/) or the API documentation (available [here](https://avinetworks.com/docs/latest/api-guide/)). There is a 1:1 mapping between the fields in the Terraform schema and the corresponding Avi Vantage object definition.

Avi Terraform provider is a native integration into the Terraform to setup all Avi configuration that is part of Avi REST API. With the Avi Terraform provider, a Terraform plan written in HCL syntax can specify any Avi configuration and it will be reflected on the Avi Controller. For instance, you can use a Terraform plan with the following code to setup a pool in Avi Vantage.

```hcl
resource "avi_pool" "testpool" {
  name= "pool-42"
  health_monitor_refs= [avi_healthmonitor.test_hm_1.id]
  tenant_ref= data.avi_tenant.default_tenant.id
  cloud_ref= data.avi_cloud.default_cloud.id
  servers {
    ip {
      type= "V4"
      addr= "x.x.x.x"
    }
    port= 8080
  }
  fail_action {
    type= "FAIL_ACTION_CLOSE_CONN"
  }
}
```
The above code breaks down as follows::
* The line below declares a resource of type ‘pool’ with name ‘testpool’:

```hcl
  resource "avi_pool" "testpool" {
```
* The line below shows a field that is a reference to another Terraform resource declared elsewhere in the plan:

```hcl 
  health_monitor_refs= [avi_healthmonitor.test_hm_1.id]
```
* The line below shows a field that is a reference to a read-only data source corresponding to an existing Avi object not defined through the Terraform plan:

```hcl
  tenant_ref= data.avi_tenant.default_tenant.id
```

* Read-only data sources can be imported either by name or by UUID. For example, the “System-HTTP” application profile can be imported as a data source by name using the following:

```hcl
data "avi_applicationprofile" "system_http_profile" {
  name= "System-HTTP"
}

// can be referenced as data.avi_applicationprofile.system_http_profile.id
```
* Alternatively, a data source can be imported based on the object’s UUID rather than it’s name, for example:
```hcl
data "avi_applicationprofile" "system_http_profile" {
uuid= "applicationprofile-xxxxxxx"
}
```
# Usage Avi Terraform
The following are the steps to use the provider:
1. Terraform module must declare which providers it requires, so that Terraform can install and use them. Provider requirements are declared in a required_providers block. Starting with Terraform version 0.13+, Avi Terraform provider has been migrated to Terraform registry. In order to use it, you need to add the below block in versions.tf file.

```hcl
terraform {
  required_providers {
    avi = {
      source  = "vmware/avi"
      version = "21.1.1"
    }
  }
}
```
where,
* source — Source of Terraform provider. Terraform will pull the released Avi Terraform provider from Terraform registry from this namespace.

* version — Avi Terraform provider release version in Terraform registry. If this field is skipped then Terraform will pull the latest release of Avi Terraform provider from Terraform registry.

2. Create Avi provider in Terraform plan. For instance, to restrict template version in 0.1.x, you can use the following CLI:
```hcl
provider "avi" {
  avi_username = "admin"
  avi_tenant = "admin"
  avi_password = "password"
  avi_controller = "x.x.x.x"
  avi_version = "21.1.1"
}  
```
3. Create Avi virtual service and pool example. Here the VS and Pool depends on read only tenant, cloud, applicationprofile, sslkeyandcertificate, sslprofile data sources and another networkprofile, vsvip, health monitor defined as a resource in the Terraform plan.
```hcl
data "avi_tenant" "default_tenant" {
  name= "admin"
 }
data "avi_applicationprofile" "system_http_profile" {
  name= "System-HTTP"
 }

data "avi_cloud" "default_cloud" {
  name= "Defaul-Cloud"
 }

data "avi_sslkeyandcertificate" "system_default_cert" {
    name= "System-Default-Cert"
}

data "avi_sslprofile" "system_standard_sslprofile" {
    name= "System-Standard"
}

resource "avi_networkprofile" "network_profile1" {
  name = "tf-network-profile"
  profile {
    type = "PROTOCOL_TYPE_TCP_PROXY"
    tcp_proxy_profile {
      ignore_time_wait = false
      time_wait_delay = 2000
      max_retransmissions = 8
      max_syn_retransmissions = 8
      automatic = true
      receive_window = 64
      nagles_algorithm = false
      ip_dscp = 0
      reorder_threshold = 10
      min_rexmt_timeout = 100
      idle_connection_type = "KEEP_ALIVE"
      idle_connection_timeout = 600
      use_interface_mtu = true
      cc_algo = "CC_ALGO_NEW_RENO"
      aggressive_congestion_avoidance = false
      slow_start_scaling_factor = 1
      congestion_recovery_scaling_factor = 2
      reassembly_queue_size = 0
      keepalive_in_halfclose_state = true
      auto_window_growth = true
    }
  }
  connection_mirror = false
  tenant_ref = data.avi_tenant.default_tenant.id
}

resource "avi_healthmonitor" "test_hm_1" {
  name       = "terraform-monitor"
  type       = "HEALTH_MONITOR_HTTP"
  tenant_ref = data.avi_tenant.default_tenant.id
}

resource "avi_pool" "testpool" {
  name= "pool-42"
  health_monitor_refs= [avi_healthmonitor.test_hm_1.id]
  tenant_ref= data.avi_tenant.default_tenant.id
  cloud_ref= data.avi_cloud.default_cloud.id
  servers {
    ip {
      type= "V4"
      addr= "x.x.x.x"
    }
    port= 8080
  }
  fail_action {
    type= "FAIL_ACTION_CLOSE_CONN"
  }
}
resource "avi_vsvip" "test_vsvip" {
  name = "terraform-vip-1"
  vip {
    vip_id = "0"
    ip_address {
      type = "V4"
      addr = "x.x.x.x"
    }
  }
  cloud_ref = data.avi_cloud.default_cloud.id
  tenant_ref = data.avi_tenant.default_tenant.id
}
resource "avi_virtualservice" "https_vs" {
  name                          = "tf_vs"
  pool_ref                      = avi_pool.testpool.id
  tenant_ref                    = data.avi_tenant.default_tenant.id
  vsvip_ref                     = avi_vsvip.test_vsvip.id
  cloud_ref                     = data.avi_cloud.default_cloud.id
  ssl_key_and_certificate_refs  = [data.avi_sslkeyandcertificate.system_default_cert.id]
  ssl_profile_ref               = data.avi_sslprofile.system_standard_sslprofile.id
  application_profile_ref       = data.avi_applicationprofile.system_http_profile.id
  network_profile_ref           = avi_networkprofile.network_profile1.id
  cloud_type                    = "CLOUD_VCENTER"
  services {
    port           = 8443
    enable_ssl     = true
  }
}
 ```

## Authentication

The AVI provider offers following means of providing credentials for
authentication:

- Static credentials
- Environment variable

### Static credentials ###

Static credentials can be provided by adding an `avi_username`, `avi_password`, `avi_controller_ip` and `avi_tenant` in-line in the
AVI provider block:

Usage:

```hcl
// Configure the AVI provider
provider "avi" {
  avi_username = "admin"
  avi_tenant = "admin"
  avi_password = "password"
  avi_controller = "x.x.x.x"
  avi_version = "21.1.1"
} 
```
```hcl
// Configure the AVI provider is csp authentication is needed
provider "avi" {
  avi_csp_host = "csp_host"
  avi_tenant = "admin"
  avi_csp_token = "csp_token"
  avi_controller = "x.x.x.x"
  avi_version = "21.1.1"
} 
```
### Environment variables

You can provide your credentials via the `AVI_USERNAME`, `AVI_PASSWORD`, `AVI_CONTROLLER` , `AVI_VERSION` and `AVI_TENANT`
environment variables, representing your AVI username, password, controller, api version and tenant, respectively.

```hcl
provider "avi" {}
```
### Usage:

```sh
$ export AVI_USERNAME = username
$ export AVI_PASSWORD = password
$ export AVI_CONTROLLER = x.x.x.x
$ export AVI_TENANT = foo
$ export AVI_VERSION=21.1.1
$ terraform init
$ terraform plan
```

You can export `AVI_SUPPRESS_SENSITIVE_FIELDS_DIFF` environment variable as `AVI_SUPPRESS_SENSITIVE_FIELDS_DIFF=true`
if you want Terraform to suppress the difference for sensitive fields during the plan update. However, if you want to
intentionally update the sensitive fields in the plan update then you need to un-export the environment variable or set
it to false, i.e. export `AVI_SUPPRESS_SENSITIVE_FIELDS_DIFF=false` or unset `AVI_SUPPRESS_SENSITIVE_FIELDS_DIFF`.

### Usage:
```sh
$ export AVI_SUPPRESS_SENSITIVE_FIELDS_DIFF = true
```

# Examples
| Name                   | Link       | Description |
|------------------------|------------|-------------|
| Virtual Services             | [Basic VS](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/basic) <br>[HTTPS VS](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/virtualservice/https_vs) <br> [HTTP VS](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/virtualservice/http_vs) <br> [HTTP VS](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/virtualservice/http_vs) <br> [DNS VS](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/virtualservice/dns_vs) <br> [TCP VS](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/virtualservice/tcp_vs) | Using these examples we can configure HTTPS Virtual Service on AVI controller with dependent resources including Pool, Pool Group, SSL Certificates, Health Monitor, VsVip, Application Persistence Profile, Network Profile and datasources including Application Profiles, Tenant, Cloud, Service Engine Group, Network Profile, Analytics Profile, SSL certificates, SSl Profiles, VRF Context. |
| VMware Environment    | [Vcenter Cloud](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/cloud/vmware) <br> [Vcenter Controller Deployment](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/vmware) | Using Vcenter Cloud example we can configure cloud of type CLOUD_VCENTER on AVI controller. <br> Using Controller Deployment examples we can deploy AVI Controller on the Vcenter using Ova or Ovf or content library. |
| NSXT Environment    | [nsx-t Cloud](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/nsxt/avi_nsxt_cloud) | Using this example we can integrate AVI Solution with NSX-T. This Terraform plan will create NSX-T cloud using AVI Terraform provider. |
| GCP Environment     | [GCP Cloud and Controller](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/gcp) | These examples deploy AVI controller on GCP, configures AVI controller cluster and cloud of type GCP. |
| Azure Environment   | [Azure Cloud and Controller](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/azure) | These examples deploy AVI controller on Azure, configures AVI controller cluster, cloud of type Azure, Virtual Services and its dependent resources. |
| AWS Environment     | [AWS Cloud and Controller](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/aws) | These examples deploy AVI controller on AWS, configures AVI controller cluster, cloud of type AWS, Virtual Services and its dependent resources. |
| WAF                 | [Waf Profile/Policy](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/waf_V2) | Using this example we can configure WAF Policy on AVI controller with resources including Waf Profile, WAF Policy, Virtual Service, Pool, Pool Group, SSL profile, VsVip, Network Profile and datasources including WAF Application Signature Provider, Application Profiles, Tenant, Cloud, SSL certificates, SSl Profiles, VRF Context. |
| Horizon             | [Horizon VS](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/Horizon) | This Terraform example can be used to configure AVI for Horizon in a Shared VIP with L7 and L4 Virtual Services. |
| Openstack Environment | [Openstack Cloud and Controller](https://github.com/vmware/terraform-provider-avi/tree/eng/examples/openstack) | Using this example we can launch AVI controller instance in openstack environment and configure application(Virtual Service) on controller. |

