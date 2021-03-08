############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_ssopolicy"
sidebar_current: "docs-avi-resource-ssopolicy"
description: |-
  Creates and manages Avi SSOPolicy.
---

# avi_ssopolicy

The SSOPolicy resource allows the creation and management of Avi SSOPolicy

## Example Usage

```hcl
resource "avi_ssopolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `authentication_policy` - (Required) Authentication policy settings. Field introduced in 18.2.1.
* `name` - (Required) Name of the sso policy. Field introduced in 18.2.3.
* `authorization_policy` - (Optional) Authorization policy settings. Field introduced in 18.2.5.
* `labels` - (Optional) Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `tenant_ref` - (Optional) Uuid of the tenant. It is a reference to an object of type tenant. Field introduced in 18.2.3.
* `type` - (Optional) Sso policy type. Enum options - SSO_TYPE_SAML, SSO_TYPE_PINGACCESS, SSO_TYPE_JWT, SSO_TYPE_LDAP. Field introduced in 18.2.5.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the sso policy. Field introduced in 18.2.3.

