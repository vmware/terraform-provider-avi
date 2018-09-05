---
layout: "avi"
page_title: "Provider: AVI"
sidebar_current: "docs-avi-index"
description: |-
  The AVI provider is used to interact with the many resources supported by AVI. The provider needs to be configured with the proper credentials before it can be used.
---

# AVI Provider

The AVI provider is used to interact with [AVI Controller](https://avinetworks.com/).
The provider needs
to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
// Configure the AVI provider
provider "avi" {
    avi_username   = "${var.avi_username}"
    avi_password   = "${var.avi_password}"
    avi_controller = "${var.avi_controller_ip}"
    avi_tenant     = "foo"
}
```

# Create a user account
```hcl
resource "avi_useraccount" "avi_user" {
  username = "${var.avi_username}"
  old_password = "${var.avi_current_password}"
  password = "${var.avi_new_password}"
}
```

## Authentication

The AVI provider offers following means of providing credentials for
authentication:

- Static credentials
- Environment variable

### Static credentials ###

Static credentials can be provided by adding an `avi_username`, `avi_password`, `avi_controller_ip` and `avi_tenant` in-line in the
AVI provider block:

Usage:

```hcl
// Configure the AVI provider
provider "avi" {
    avi_username   = "username"
    avi_password   = "password"
    avi_controller = "10.0.0.3"
    avi_tenant     = "tenant"
}
```

### Environment variables

You can provide your credentials via the `AVI_USERNAME`, `AVI_PASSWORD`, `AVI_CONTROLLER` and `AVI_TENANT`
environment variables, representing your AVI username, password, controller and tenant, respectively.

```hcl
provider "avi" {}
```
Usage:

```sh
$ export AVI_USERNAME = username
$ export AVI_PASSWORD = password
$ export AVI_CONTROLLER = 10.0.0.3
$ export AVI_TENANT = foo
$ terraform init
$ terraform plan
```
