Terraform Provider
==================
- [![Build Status](https://travis-ci.org/terraform-providers/terraform-provider-avi.svg?branch=master)](https://travis-ci.org/terraform-providers/terraform-provider-avi)
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Website: https://www.terraform.io
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-text.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.8 (to build the provider plugin)

Usage
---------------------

Create Avi Provider in terraform plan

```
# For example, restrict template version in 0.1.x
provider "avi" {
  avi_username = "admin"
  avi_tenant = "admin"
  avi_password = "something"
  avi_controller= "42.42.42.42"
}
```

Create Avi Pool Example. Here Pool depends on read only tenant data source and another health monitor defined as resource in the terraform plan

```
resource "avi_pool" "testpool" {
  name= "pool-42",
  health_monitor_refs= ["${avi_healthmonitor.test_hm_1.id}"]
  tenant_ref= "${data.avi_tenant.default_tenant.id}"
  cloud_ref= "${data.avi_cloud.default_cloud.id}"
  application_persistence_profile_ref= "${avi_applicationpersistenceprofile.test_applicationpersistenceprofile.id}"
  servers {
    ip= {
      type= "V4",
      addr= "10.90.64.66",
    }
    port= 8080
  }
  fail_action= {
    type= "FAIL_ACTION_CLOSE_CONN"
  }
}

```

Reference existing resources as readonly or data sources

```
data "avi_applicationprofile" "system_http_profile" {
  name= "System-HTTP"
}

application_profile_ref= "${data.avi_applicationprofile.system_https_profile.id}"

```


Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-avi`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone https://github.com/terraform-providers/terraform-provider-avi.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-


$ make
```

Using the provider
----------------------
## Fill in for each provider

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.8+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make bin
...
$ $GOPATH/bin/terraform-provider-avi
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
