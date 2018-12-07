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
data "ApplicationProfile" "foo_ApplicationProfile" {
    uuid = "ApplicationProfile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ApplicationProfile by name.
* `uuid` - (Optional) Search ApplicationProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cloud_config_cksum` - Checksum of application profiles.
* `created_by` - Name of the application profile creator.
* `description` - General description.
* `dns_service_profile` - Specifies various dns service related controls for virtual service.
* `dos_rl_profile` - Specifies various security related controls for virtual service.
* `http_profile` - Specifies the http application proxy profile parameters.
* `name` - The name of the application profile.
* `preserve_client_ip` - Specifies if client ip needs to be preserved for backend connection.
* `preserve_client_port` - Specifies if we need to preserve client port while preserving client ip for backend connections.
* `sip_service_profile` - Specifies various sip service related controls for virtual service.
* `tcp_app_profile` - Specifies the tcp application proxy profile parameters.
* `tenant_ref` - It is a reference to an object of type tenant.
* `type` - Specifies which application layer proxy is enabled for the virtual service.
* `uuid` - Uuid of the application profile.

