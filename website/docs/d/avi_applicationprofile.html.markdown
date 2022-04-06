<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_applicationprofile"
sidebar_current: "docs-avi-datasource-applicationprofile"
description: |-
  Get information of Avi ApplicationProfile.
---

# avi_applicationprofile

This data source is used to to get avi_applicationprofile objects.

## Example Usage

```hcl
data "avi_applicationprofile" "foo_applicationprofile" {
    uuid = "applicationprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ApplicationProfile by name.
* `uuid` - (Optional) Search ApplicationProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `app_service_type` - Specifies app service type for an application. Enum options - APP_SERVICE_TYPE_L7_HORIZON, APP_SERVICE_TYPE_L4_BLAST, APP_SERVICE_TYPE_L4_PCOIP. Field introduced in 21.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `cloud_config_cksum` - Checksum of application profiles. Internally set by cloud connector. Field introduced in 17.2.14, 18.1.5, 18.2.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `created_by` - Name of the application profile creator. Field introduced in 17.2.14, 18.1.5, 18.2.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `description` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `dns_service_profile` - Specifies various dns service related controls for virtual service. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `dos_rl_profile` - Specifies various security related controls for virtual service. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `http_profile` - Specifies the http application proxy profile parameters. Allowed in enterprise with any value edition, basic edition, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `name` - The name of the application profile. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `preserve_client_ip` - Specifies if client ip needs to be preserved for backend connection. Not compatible with connection multiplexing. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `preserve_client_port` - Specifies if we need to preserve client port while preserving client ip for backend connections. Field introduced in 17.2.7. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `preserve_dest_ip_port` - Specifies if destination ip and port needs to be preserved for backend connection. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials(allowed values- false) edition, basic(allowed values- false) edition, enterprise with cloud services edition.
* `sip_service_profile` - Specifies various sip service related controls for virtual service. Field introduced in 17.2.8, 18.1.3, 18.2.1. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `tcp_app_profile` - Specifies the tcp application proxy profile parameters. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `type` - Specifies which application layer proxy is enabled for the virtual service. Enum options - APPLICATION_PROFILE_TYPE_L4, APPLICATION_PROFILE_TYPE_HTTP, APPLICATION_PROFILE_TYPE_SYSLOG, APPLICATION_PROFILE_TYPE_DNS, APPLICATION_PROFILE_TYPE_SSL, APPLICATION_PROFILE_TYPE_SIP. Allowed in enterprise with any value edition, essentials(allowed values- application_profile_type_l4) edition, basic(allowed values- application_profile_type_l4,application_profile_type_http) edition, enterprise with cloud services edition.
* `uuid` - Uuid of the application profile. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

