############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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

* `name` - (Required) The name of the application profile.
* `type` - (Required) Specifies which application layer proxy is enabled for the virtual service.
* `cloud_config_cksum` - (Optional) Checksum of application profiles.
* `created_by` - (Optional) Name of the application profile creator.
* `description` - (Optional) User defined description for the object.
* `dns_service_profile` - (Optional) Specifies various dns service related controls for virtual service.
* `dos_rl_profile` - (Optional) Specifies various security related controls for virtual service.
* `http_profile` - (Optional) Specifies the http application proxy profile parameters.
* `preserve_client_ip` - (Optional) Specifies if client ip needs to be preserved for backend connection.
* `preserve_client_port` - (Optional) Specifies if we need to preserve client port while preserving client ip for backend connections.
* `sip_service_profile` - (Optional) Specifies various sip service related controls for virtual service.
* `tcp_app_profile` - (Optional) Specifies the tcp application proxy profile parameters.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the application profile.

