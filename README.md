Terraform Provider
==================
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Website: https://www.terraform.io
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-text.svg" width="600px">

Developing the Provider or Use Locally Built Provider
----------------------------------------------------
If you wish to work on the provider or want to use the locally built provider,
you'll first need [Go](http://www.golang.org) and [Terraform](https://www.terraform.io) installed on your machine.
You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin`
to your `$PATH`.

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.12.x/0.13+ (0.11.x or lower is incompatible)
-	[Go](https://golang.org/doc/install) 1.13 (to build the provider plugin)

Building The Provider (Terraform v0.12+)
----------------------------------------

Clone repository to: `$GOPATH/src/github.com/vmware/terraform-provider-avi`

```sh
$ mkdir -p $GOPATH/src/github.com/vmware; cd $GOPATH/src/github.com/vmware
$ git clone https://github.com/vmware/terraform-provider-avi.git
```

Enter the provider directory and build the provider.

```sh
$ cd $GOPATH/src/github.com/vmware/terraform-provider-avi
$ make
```
For Terraform v0.12.x to use a locally built version of a provider add following to ~/.terraformrc on Linux/Unix.

```
providers {
  "avi" = "$GOPATH/bin/terraform-provider-avi"
}
```
Or copy provider binary in ~./.terraform/plugins/linux_amd64/
```shell
$ mkdir -p ~./.terraform/plugins/linux_amd64/
$ cp $GOPATH/bin/terraform-provider-avi ~./.terraform/plugins/linux_amd64/
```

Building The Provider (Terraform v0.13+)
----------------------------------------

Clone repository to: `$GOPATH/src/github.com/vmware/terraform-provider-avi`

```sh
$ mkdir -p $GOPATH/src/github.com/vmware; cd $GOPATH/src/github.com/vmware
$ git clone https://github.com/vmware/terraform-provider-avi.git
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/vmware/terraform-provider-avi
$ make build13
```
This will put the provider binary in the ~/.terraform.d/plugins/vmware.com/avi/avi/<provider_version>/$(GOOS)_$(GOARCH)
directory.

For Terraform v0.13+, to use a locally built version of a provider you must add the following snippet to every
terraform plan.
```
terraform {
  required_providers {
    avi = {
      source  = "vmware.com/avi/avi"
      version = "<provider_version>"
    }
  }
}
```

------
Usage
------

Create Avi Provider in terraform plan

```
provider "avi" {
  avi_username = "admin"
  avi_tenant = "admin"
  avi_password = "password"
  avi_controller = "x.x.x.x"
  avi_version = "20.1.4"
}
```

Create Avi Pool Example. Here Pool depends on read only tenant data source and another health monitor defined as
resource in the terraform plan

```
data "avi_tenant" "default_tenant" {
  name= "admin"
}
data "avi_cloud" "default_cloud" {
  name= "Default-Cloud"
}

resource "avi_applicationpersistenceprofile" "test_applicationpersistenceprofile" {
  name             = "terraform-app-pers-profile"
  tenant_ref       = data.avi_tenant.default_tenant.id
  persistence_type = "PERSISTENCE_TYPE_CLIENT_IP_ADDRESS"
}

resource "avi_healthmonitor" "test_hm_1" {
  name       = "terraform-monitor"
  type       = "HEALTH_MONITOR_HTTP"
  tenant_ref = data.avi_tenant.default_tenant.id
}

resource "avi_pool" "testpool" {
  name= "pool-42"
  health_monitor_refs = [avi_healthmonitor.test_hm_1.id]
  tenant_ref = data.avi_tenant.default_tenant.id
  cloud_ref = data.avi_cloud.default_cloud.id
  application_persistence_profile_ref= avi_applicationpersistenceprofile.test_applicationpersistenceprofile.id
  servers {
    ip {
      type = "V4"
      addr = "10.90.64.66"
    }
    port= 8080
  }
  fail_action {
    type = "FAIL_ACTION_CLOSE_CONN"
  }
}
```

Reference existing resources as readonly or data sources

```
data "avi_applicationprofile" "system_http_profile" {
  name= "System-HTTP"
}

application_profile_ref= data.avi_applicationprofile.system_https_profile.id

```
-----------------

Test The Provider
-----------------
In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`. 
Running the tests for a provider requires version 0.12.26 or higher of the Terraform CLI.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```