############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_applicationpersistenceprofile"
sidebar_current: "docs-avi-resource-applicationpersistenceprofile"
description: |-
  Creates and manages Avi ApplicationPersistenceProfile.
---

# avi_applicationpersistenceprofile

The ApplicationPersistenceProfile resource allows the creation and management of Avi ApplicationPersistenceProfile

## Example Usage

```hcl
resource "avi_applicationpersistenceprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A user-friendly name for the persistence profile.
* `persistence_type` - (Required) Method used to persist clients to the same server for a duration of time or a session.
* `app_cookie_persistence_profile` - (Optional) Specifies the application cookie persistence profile parameters.
* `description` - (Optional) User defined description for the object.
* `hdr_persistence_profile` - (Optional) Specifies the custom http header persistence profile parameters.
* `http_cookie_persistence_profile` - (Optional) Specifies the http cookie persistence profile parameters.
* `ip_persistence_profile` - (Optional) Specifies the client ip persistence profile parameters.
* `is_federated` - (Optional) This field describes the object's replication scope.
* `server_hm_down_recovery` - (Optional) Specifies behavior when a persistent server has been marked down by a health monitor.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the persistence profile.

