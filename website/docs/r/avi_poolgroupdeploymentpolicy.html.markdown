############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_poolgroupdeploymentpolicy"
sidebar_current: "docs-avi-resource-poolgroupdeploymentpolicy"
description: |-
  Creates and manages Avi PoolGroupDeploymentPolicy.
---

# avi_poolgroupdeploymentpolicy

The PoolGroupDeploymentPolicy resource allows the creation and management of Avi PoolGroupDeploymentPolicy

## Example Usage

```hcl
resource "avi_poolgroupdeploymentpolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the pool group deployment policy.
* `auto_disable_old_prod_pools` - (Optional) It will automatically disable old production pools once there is a new production candidate.
* `description` - (Optional) User defined description for the object.
* `evaluation_duration` - (Optional) Duration of evaluation period for automatic deployment. Allowed values are 60-86400. Unit is sec.
* `labels` - (Optional) Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `rules` - (Optional) List of list.
* `scheme` - (Optional) Deployment scheme. Enum options - BLUE_GREEN, CANARY.
* `target_test_traffic_ratio` - (Optional) Target traffic ratio before pool is made production. Allowed values are 1-100. Unit is ratio.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `test_traffic_ratio_rampup` - (Optional) Ratio of the traffic that is sent to the pool under test. Test ratio of 100 means blue green. Allowed values are 1-100.
* `webhook_ref` - (Optional) Webhook configured with url that avi controller will pass back information about pool group, old and new pool information and current deployment rule results. It is a reference to an object of type webhook. Field introduced in 17.1.1.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the pool group deployment policy.

