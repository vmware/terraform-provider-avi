<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_applicationprofile"
sidebar_current: "docs-avi-resource-applicationprofile"
description: |-
  Creates and manages Avi ApplicationProfile.
---

# avi_applicationprofile

The ApplicationProfile resource allows the creation and management of Avi ApplicationProfile

## Example Usage

```hcl
resource "avi_applicationprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the application profile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `type` - (Required) Specifies which application layer proxy is enabled for the virtual service. Enum options - APPLICATION_PROFILE_TYPE_L4, APPLICATION_PROFILE_TYPE_HTTP, APPLICATION_PROFILE_TYPE_SYSLOG, APPLICATION_PROFILE_TYPE_DNS, APPLICATION_PROFILE_TYPE_SSL, APPLICATION_PROFILE_TYPE_SIP. Allowed in enterprise edition with any value, essentials edition(allowed values- application_profile_type_l4), basic edition(allowed values- application_profile_type_l4,application_profile_type_http), enterprise with cloud services edition.
* `app_service_type` - (Optional) Specifies app service type for an application. Enum options - APP_SERVICE_TYPE_L7_HORIZON, APP_SERVICE_TYPE_L4_BLAST, APP_SERVICE_TYPE_L4_PCOIP, APP_SERVICE_TYPE_L4_FTP. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `cloud_config_cksum` - (Optional) Checksum of application profiles. Internally set by cloud connector. Field introduced in 17.2.14, 18.1.5, 18.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `created_by` - (Optional) Name of the application profile creator. Field introduced in 17.2.14, 18.1.5, 18.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `description` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `dns_service_profile` - (Optional) Specifies various dns service related controls for virtual service. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `dos_rl_profile` - (Optional) Specifies various security related controls for virtual service. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `http_profile` - (Optional) Specifies the http application proxy profile parameters. Allowed in enterprise edition with any value, basic, enterprise with cloud services edition.
* `l4_ssl_profile` - (Optional) Specifies various l4 ssl service related controls for virtual service. Field introduced in 22.1.2. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `preserve_client_ip` - (Optional) Specifies if client ip needs to be preserved for backend connection. Not compatible with connection multiplexing. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `preserve_client_port` - (Optional) Specifies if we need to preserve client port while preserving client ip for backend connections. Field introduced in 17.2.7. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `preserve_dest_ip_port` - (Optional) Specifies if destination ip and port needs to be preserved for backend connection. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition.
* `sip_service_profile` - (Optional) Specifies various sip service related controls for virtual service. Field introduced in 17.2.8, 18.1.3, 18.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tcp_app_profile` - (Optional) Specifies the tcp application proxy profile parameters. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the application profile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

