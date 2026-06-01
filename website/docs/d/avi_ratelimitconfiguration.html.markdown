<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_ratelimitconfiguration"
sidebar_current: "docs-avi-datasource-ratelimitconfiguration"
description: |-
  Get information of Avi RateLimitConfiguration.
---

# avi_ratelimitconfiguration

This data source is used to to get avi_ratelimitconfiguration objects.

## Example Usage

```hcl
data "avi_ratelimitconfiguration" "foo_ratelimitconfiguration" {
    uuid = "ratelimitconfiguration-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search RateLimitConfiguration by name.
* `uuid` - (Optional) Search RateLimitConfiguration by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `burst` - The maximum request per second(rps) user intends to support for this category.this is not guaranteed as this will be the minimum of the rps supported by the resources in the category and this value.if user doesn't provide then it will be minimum value of the resources in this category. Allowed values are 1-1000. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `description` - Description for the rate limit configuration. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `http_methods` - List of http method(s) of the resources that need to be rate limited. Enum options - HTTP_METHOD_GET, HTTP_METHOD_HEAD, HTTP_METHOD_PUT, HTTP_METHOD_DELETE, HTTP_METHOD_POST, HTTP_METHOD_OPTIONS, HTTP_METHOD_TRACE, HTTP_METHOD_CONNECT, HTTP_METHOD_PATCH, HTTP_METHOD_PROPFIND, HTTP_METHOD_PROPPATCH, HTTP_METHOD_MKCOL, HTTP_METHOD_COPY, HTTP_METHOD_MOVE, HTTP_METHOD_LOCK, HTTP_METHOD_UNLOCK. Field introduced in 31.2.1. Minimum of 1 items required. Maximum of 5 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `name` - Name of the rate limit configuration(unique). Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `resource` - Ratelimitresource which needs to be rate limited. Enum options - RATE_LIMIT_VIRTUALSERVICE, RATE_LIMIT_POOL, RATE_LIMIT_LOGIN, RATE_LIMIT_AUTHTOKEN, RATE_LIMIT_HEALTHMONITOR. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - Tenant ref for the auth rate limit configuration. It is a reference to an object of type tenant. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `token_refill_rate` - Token refill rate. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `type` - Type of the rate limiter, for now we only support api categorization based. Enum options - RATE_LIMITER_API_CATEGORY. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of the rate limit configuration. Field introduced in 31.2.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

