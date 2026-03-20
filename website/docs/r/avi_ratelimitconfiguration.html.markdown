<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_ratelimitconfiguration"
sidebar_current: "docs-avi-resource-ratelimitconfiguration"
description: |-
  Creates and manages Avi RateLimitConfiguration.
---

# avi_ratelimitconfiguration

The RateLimitConfiguration resource allows the creation and management of Avi RateLimitConfiguration

## Example Usage

```hcl
resource "avi_ratelimitconfiguration" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `http_methods` - (Required) List of http method(s) of the resources that need to be rate limited. Enum options - HTTP_METHOD_GET, HTTP_METHOD_HEAD, HTTP_METHOD_PUT, HTTP_METHOD_DELETE, HTTP_METHOD_POST, HTTP_METHOD_OPTIONS, HTTP_METHOD_TRACE, HTTP_METHOD_CONNECT, HTTP_METHOD_PATCH, HTTP_METHOD_PROPFIND, HTTP_METHOD_PROPPATCH, HTTP_METHOD_MKCOL, HTTP_METHOD_COPY, HTTP_METHOD_MOVE, HTTP_METHOD_LOCK, HTTP_METHOD_UNLOCK. Field introduced in 31.2.1. Minimum of 1 items required. Maximum of 5 items allowed. Allowed with any value in enterprise, enterprise with cloud services edition. 
* `name` - (Required) Name of the rate limit configuration(unique). Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition. 
* `resource` - (Required) Ratelimitresource which needs to be rate limited. Enum options - RATE_LIMIT_VIRTUALSERVICE, RATE_LIMIT_POOL, RATE_LIMIT_LOGIN, RATE_LIMIT_AUTHTOKEN, RATE_LIMIT_HEALTHMONITOR. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition. 
* `token_refill_rate` - (Required) Token refill rate. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition. 
* `burst` - (Optional) The maximum request per second(rps) user intends to support for this category.this is not guaranteed as this will be the minimum of the rps supported by the resources in the category and this value.if user doesn't provide then it will be minimum value of the resources in this category. Allowed values are 1-1000. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition. 
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `description` - (Optional) Description for the rate limit configuration. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition. 
* `tenant_ref` - (Optional) Tenant ref for the auth rate limit configuration. It is a reference to an object of type tenant. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition. 
* `type` - (Optional) Type of the rate limiter, for now we only support api categorization based. Enum options - RATE_LIMITER_API_CATEGORY. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition. Changing this value forces the resource to be recreated.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the rate limit configuration. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.

