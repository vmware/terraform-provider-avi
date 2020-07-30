############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_httppolicyset"
sidebar_current: "docs-avi-resource-httppolicyset"
description: |-
  Creates and manages Avi HTTPPolicySet.
---

# avi_httppolicyset

The HTTPPolicySet resource allows the creation and management of Avi HTTPPolicySet

## Example Usage

```hcl
resource "avi_httppolicyset" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the http policy set.
* `cloud_config_cksum` - (Optional) Checksum of cloud configuration for pool.
* `created_by` - (Optional) Creator name.
* `description` - (Optional) User defined description for the object.
* `http_request_policy` - (Optional) Http request policy for the virtual service.
* `http_response_policy` - (Optional) Http response policy for the virtual service.
* `http_security_policy` - (Optional) Http security policy for the virtual service.
* `is_internal_policy` - (Optional) Boolean flag to set is_internal_policy.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the http policy set.

