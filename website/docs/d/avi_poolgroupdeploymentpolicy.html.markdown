
############################################################################
#
# AVI CONFIDENTIAL
# __________________
#
# [2013] - [2019] Avi Networks Incorporated
# All Rights Reserved.
#
# NOTICE: All information contained herein is, and remains the property
# of Avi Networks Incorporated and its suppliers, if any. The intellectual
# and technical concepts contained herein are proprietary to Avi Networks
# Incorporated, and its suppliers and are covered by U.S. and Foreign
# Patents, patents in process, and are protected by trade secret or
# copyright law, and other laws. Dissemination of this information or
# reproduction of this material is strictly forbidden unless prior written
# permission is obtained from Avi Networks Incorporated.
###

---
layout: "avi"
page_title: "AVI: avi_poolgroupdeploymentpolicy"
sidebar_current: "docs-avi-datasource-poolgroupdeploymentpolicy"
description: |-
  Get information of Avi PoolGroupDeploymentPolicy.
---

# avi_poolgroupdeploymentpolicy

This data source is used to to get avi_poolgroupdeploymentpolicy objects.

## Example Usage

```hcl
data "PoolGroupDeploymentPolicy" "foo_PoolGroupDeploymentPolicy" {
    uuid = "PoolGroupDeploymentPolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search PoolGroupDeploymentPolicy by name.
* `uuid` - (Optional) Search PoolGroupDeploymentPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `auto_disable_old_prod_pools` - It will automatically disable old production pools once there is a new production candidate.
* `description` - General description.
* `evaluation_duration` - Duration of evaluation period for automatic deployment.
* `name` - The name of the pool group deployment policy.
* `rules` - General description.
* `scheme` - Deployment scheme.
* `target_test_traffic_ratio` - Target traffic ratio before pool is made production.
* `tenant_ref` - It is a reference to an object of type tenant.
* `test_traffic_ratio_rampup` - Ratio of the traffic that is sent to the pool under test.
* `uuid` - Uuid of the pool group deployment policy.
* `webhook_ref` - Webhook configured with url that avi controller will pass back information about pool group, old and new pool information and current deployment rule results.

