# Terraform Provider for Avi (NSX Advanced Load Balancer)

Terraform provider that manages **NSX Advanced Load Balancer / Avi Vantage** objects Test
(pools, virtual services, health monitors, certs, GSLB, WAF, etc.) via the Avi Controller REST API.
Maintained by Broadcom/VMware (`vmware/terraform-provider-avi`), built on top of
[`hashicorp/terraform-plugin-sdk/v2`](https://github.com/hashicorp/terraform-plugin-sdk) and
[`vmware/alb-sdk`](https://github.com/vmware/alb-sdk) (the Go client for the Avi Controller API).

- Provider version tracks Avi Controller releases
- License: Mozilla Public License 2.0 (see `LICENSE`).
- Registry namespace (TF 0.13+): `vmware.com/avi/avi`.

## Repository layout

```
.
├── main.go                      # plugin entrypoint: plugin.Serve(avi.Provider)
├── avi/                         # provider implementation (single Go package)
│   ├── provider.go              #   schema.Provider{}: auth schema + Resource/DataSource maps (~195/136 objects)
│   ├── utils.go                 #   generic CRUD engine (APICreate/APIRead/APIUpdate/APIDelete/...),
│   │                             #   schema<->API marshaling (SchemaToAviData/APIDataToSchema), importer
│   ├── diff_suppress_func.go    #   ValidateFunc helpers (validateInteger/Bool/Float) + sensitive-field diff suppression
│   ├── resource_avi_rest_dependants.go  # ~46k lines, AUTO-GENERATED nested/sub-object schemas (skipped by lint)
│   ├── resource_avi_<object>.go        # one file per top-level API object: schema + thin Create/Read/Update/Delete
│   ├── datasource_avi_<object>.go      # matching read-only data source per object
│   ├── resource_avi_<object>_test.go   # acceptance tests (TF_ACC=1), one set per resource
│   └── data_source_avi_<object>_test.go
├── website/docs/{r,d}/*.html.markdown  # Terraform Registry docs, one page per resource/data source
├── examples/                    # example .tf configs (aws, azure, gcp, nsxt, openstack, vmware, pool, waf_V2, ...)
├── modules/                     # reusable TF modules (services/vmware_deploy, nia/pool - Consul-Terraform-Sync)
├── scripts/                     # gofmtcheck.sh, errcheck.sh, changelog-links.sh, gogetcookie.sh
├── GNUmakefile                  # build/test/lint targets
└── .github/workflows/           # golangci-lint.yml, release.yml (goreleaser)
```

## Architecture

- **`provider.go`** is the single source of truth for what's registered: `Schema` (connection
  params), `DataSourcesMap`, `ResourcesMap`. Every Avi API object type has one entry in each map,
  wired to a `resourceAvi<Object>()` / `dataSourceAvi<Object>()` constructor.
- **Generic CRUD**: individual `resource_avi_*.go` files mostly just define the Terraform schema
  (`Resource<Object>Schema()`) and delegate Create/Read/Update/Delete to the generic helpers in
  `utils.go` (`APICreate`, `APIRead`, `APIUpdate`, `APIDelete`), passing the API object type string
  and schema. This is what lets ~195 resource types share one CRUD implementation instead of each
  hand-rolling REST calls.
- **`resource_avi_rest_dependants.go`** holds schemas for nested/sub-objects referenced by the
  top-level resources (e.g. `ResourceACSubjectInfoSchema`). It's large and mechanically generated —
  it's excluded from `golangci-lint` (`skip-files` in `.golangci.yml`) and should not be hand-edited
  piecemeal.
- **Codegen note**: git history is dominated by "Updated assets for terraform" commits synced from
  an internal Jenkins job (no generator script lives in this repo). Assume most
  `resource_avi_*.go` / `datasource_avi_*.go` / `website/docs/**` files are regenerated from the
  Avi Controller's Swagger/API spec upstream — hand edits to generated sections are likely to be
  overwritten on the next sync. Manual, durable changes belong in `utils.go`, `provider.go`,
  `diff_suppress_func.go`, examples, or genuinely hand-maintained resource logic.
- **Auth**: `provider.go` supports username/password, API token (`avi_authtoken`), and VMware Cloud
  CSP token (`avi_csp_token` / `avi_csp_host`) auth against the Avi Controller, using lazy
  authentication (session is only established when a resource actually needs it).
- **Special-cased objects**: `utils.go` maintains a `postNotAllowed` list (`IsPostNotAllowed`) for
  API objects that can't be created via POST (e.g. singleton/system objects like `cluster`,
  `useraccount`) — Create for these effectively becomes a Read/Update against an existing object.
- **Sensitive field diffs**: `AVI_SUPPRESS_SENSITIVE_FIELDS_DIFF` env var
  (`diff_suppress_func.go`) can suppress plan diffs on sensitive fields the API masks on read.

## Requirements

- [Go](https://golang.org/doc/install) 1.25.6+ (see `go.mod`)
- [Terraform](https://www.terraform.io/downloads.html) 0.13+ (0.12 build path still exists but is legacy)
- Network access to an Avi Controller for acceptance testing

## Building the provider

Clone into your `GOPATH` (module-aware builds still expect this path for the historical
`~/.terraform.d/plugins` layout used below):

```sh
mkdir -p $GOPATH/src/github.com/vmware && cd $GOPATH/src/github.com/vmware
git clone https://github.com/vmware/terraform-provider-avi.git
cd terraform-provider-avi
```

**Terraform 0.13+ (recommended):**

```sh
make build13
```

Installs to `~/.terraform.d/plugins/vmware.com/avi/avi/<AVI_PROVIDER_VERSION>/$(GOOS)_$(GOARCH)/`
(version pinned by `AVI_PROVIDER_VERSION` in `GNUmakefile`). Reference it in a plan:

```hcl
terraform {
  required_providers {
    avi = {
      source  = "vmware.com/avi/avi"
      version = "<provider_version>"
    }
  }
}
```

**Terraform 0.12.x (legacy):**

```sh
make build   # go install; then wire up ~/.terraformrc `providers {}` block or
             # copy the binary into ~/.terraform/plugins/linux_amd64/
```

**Faster local iteration** without bumping the pinned version each time — use a
[dev override](https://developer.hashicorp.com/terraform/cli/config/config-file#development-overrides-for-provider-developers)
in `~/.terraformrc`:

```hcl
provider_installation {
  dev_overrides {
    "vmware.com/avi/avi" = "/path/to/GOPATH/bin"
  }
  direct {}
}
```

## Usage

```hcl
provider "avi" {
  avi_username   = "admin"
  avi_tenant     = "admin"
  avi_password   = "password"
  avi_controller = "x.x.x.x"
  avi_version    = "21.1.1"
}
```

Example: a pool depending on data sources and other resources in the same plan.

```hcl
data "avi_tenant" "default_tenant" {
  name = "admin"
}
data "avi_cloud" "default_cloud" {
  name = "Default-Cloud"
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
  name                                 = "pool-42"
  health_monitor_refs                  = [avi_healthmonitor.test_hm_1.id]
  tenant_ref                           = data.avi_tenant.default_tenant.id
  cloud_ref                            = data.avi_cloud.default_cloud.id
  application_persistence_profile_ref  = avi_applicationpersistenceprofile.test_applicationpersistenceprofile.id
  servers {
    ip {
      type = "V4"
      addr = "10.90.64.66"
    }
    port = 8080
  }
  fail_action {
    type = "FAIL_ACTION_CLOSE_CONN"
  }
}
```

Reference existing controller objects read-only via a data source:

```hcl
data "avi_applicationprofile" "system_http_profile" {
  name = "System-HTTP"
}
# use as: application_profile_ref = data.avi_applicationprofile.system_http_profile.id
```

More end-to-end examples (AWS/Azure/GCP/NSX-T/OpenStack/vSphere, WAF, GSLB, Horizon, autoscale) live
under `examples/`.

## Development workflow

```sh
make build13        # build + install locally (also runs fmtcheck)
make fmt             # gofmt -w over the tree
make fmtcheck        # scripts/gofmtcheck.sh — CI-equivalent format check
make vet             # go vet
make errcheck        # scripts/errcheck.sh
```

Adding/changing a resource:
1. Confirm whether the object type is actually generated upstream — check if a matching
   `website/docs/r/*.html.markdown` and Avi API object already exist. If this is a genuine new
   object type, add `resource_avi_<object>.go` + `datasource_avi_<object>.go` following the pattern
   in an existing simple resource (e.g. `resource_avi_healthmonitor.go`): a `Resource<Object>Schema()`
   map, a thin `resourceAvi<Object>()` wired to `utils.go`'s generic CRUD, and register both in
   `provider.go`'s `ResourcesMap`/`DataSourcesMap`.
2. Add the docs page under `website/docs/r/` (and `d/` for the data source).
3. Add acceptance tests (`resource_avi_<object>_test.go`, `data_source_avi_<object>_test.go`).

## Testing

```sh
make test            # unit tests, no live controller needed
make testacc         # acceptance tests — TF_ACC=1, creates REAL objects against a live Avi Controller
```

Acceptance tests need `AVI_CONTROLLER`, `AVI_USERNAME`, `AVI_PASSWORD` (or token equivalents) in
the environment, cost real controller resources, and can take a long time (`testacc` uses a 120m
timeout). Run a single test with `specific_test=TestAcc... make testacc`.

`conftest.py` / `test_terraform_acc.py` are pytest wrappers used by the internal CI pipeline to
shell out to `make testacc` and capture logs — not needed for local `go test` development.

## Linting / CI

- `.github/workflows/golangci-lint.yml` runs `golangci-lint` (enabled: `errcheck`, `gofmt`,
  `gosimple`, `ineffassign`, `nakedret`, `misspell`, `vet`, `vetshadow`; see `.golangci.yml`).
  `avi/resource_avi_rest_dependants.go` is excluded.
- `.github/workflows/release.yml` + `.goreleaser.yml` handle tagged releases (multi-OS/arch builds,
  checksums, GPG-signed, draft GitHub release).

## Contributing / ownership

- `CODEOWNERS`: `@mayank-avinetworks @manojkumarjain` own the whole repo by default.
- See `CONTRIBUTING.md` for branch/PR/commit-message conventions (CLA required).
- `CHANGELOG.md` tracks released versions; reference internal `AV-xxxxx` ticket IDs and/or GitHub
  issue numbers in bug-fix entries, matching existing style.
