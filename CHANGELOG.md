## 32.1.4
FEATURES:
- Added new resources and data sources: `avi_akoamkocluster`, `avi_albservicesstatus`, `avi_apipath`, `avi_apipolicy`, `avi_apischema`, `avi_apispecgenerate`, `avi_labelprofile`, `avi_sessionkeyforwarder`.
- Added new data source `avi_tlsprofile`.
- `avi_virtualservice`: added `api_policy_ref` and `label_profile_ref`.
- `avi_serviceenginegroup`: added `log_streaming_tls_config`, `seg_mode` and `waap_mode_config`.
- `avi_systemconfiguration`: added `ai_assistant_enabled`, `allow_private_ips`, `gslb_client_cert_ref` and `truststore_ca_certs_pem`.
- `avi_controllerproperties`: added `ai_assistant_project_key`, `ipgls_client_cache_size`, `ipgls_client_cache_ttl_minutes`, `promoted_log_fields` and `vs_se_license_reservation_fail`.
- `avi_gslb`: added `tls_config`; `avi_gslbcrmruntime`: added `tls_connection_err`.
- `avi_analyticsprofile`: added `waap_classification_disabled_penalty`.
- `avi_backupconfiguration`: added `known_hosts_file_ref`.
- `avi_jwtserverprofile`: added `allowed_algorithms`.
- `avi_fileobject`: added API spec generation progress fields (`api_spec_detail`, `state`, `progress`, `history`, etc.).
- `avi_licensestatus`: added `ssp_license_reported_at`.

IMPROVEMENTS:
- Updated Avi Go SDK (`alb-sdk`) to the latest version.
- Upgraded dependencies: `golang.org/x/crypto`, `golang.org/x/net`, `golang.org/x/sys`, `golang.org/x/text`, `google.golang.org/grpc` and `google.golang.org/protobuf`.
- Updated resource and data source documentation.

BREAKING CHANGES:
- `avi_systemconfiguration`: `syslog_servers` replaced by `syslog_server_settings`.
- `avi_controllerproperties`: removed `bm_use_ansible`.
- Removed `avi_techsupportmessage` resource and data source.
- Removed `avi_albservicesfiledownload` resource (data source replaced by `avi_sessionkeyforwarder`).

BUG FIXES:
- AV-273506: Fixed terraform functional test failures.
- Fixed functional test failures for controllerproperties, customipamdnsprofile, user and useraccount tests.
- AV-293347: Resolve `tenant_ref` via user-tenant-list to avoid dependency on Tenant-read RBAC permission.

## 0.3.0 (Unreleased)
## 0.2.3 (June 16, 2020)
Features:
- Updated Avi Go SDK to latest.
- Update Avi 18.2.9 release

Bugs:
- AV-85581 Added seproperties into post not allowed list
- Fix for issue #65

## 0.2.2 (May 12, 2020)
FEATURES:
- Lazy Authentication support. The provider uses lazy authentication so Avi controllers can be created in a plan and it allows password reset.
- Added examples for VMware Horizon and VMware vsphere

BUG FIXES:
- AV-79347: Terraform: Unable to reconcile state of resources modified outside Terraform
- AV-76349: Avi terraform enable=false is ignored.

## 0.2.1 (November 20, 2019)
FEATURES:
- Added new resource and datasource for upgradestatussummary object.
- Added new resource and datasource for testsedatastorelevel1 object.
- Added new resource and datasource for testsedatastorelevel2 object.
- Added new resource and datasource for testsedatastorelevel3 object.

IMPROVEMENTS:
- Updated terraform provider avi with latest gosdk (AV-72035)
- Added cluster status field to cluster datasource and resource to monitor the cluster status.

BUG FIXES
- Fixed cluster deletion erroring out due to delete not allowed on cluster.

## 0.2.0 (October 01, 2019)
FEATURES:
- Added new resource and datasource for controllerportalregistration object.
- Added new resource and datasource for customerportalinfo object.
- Added new resource and datasource for image object. This is used in managing software updates to the Avi Controller.
- Added new resource and datasource for portalfileupload object.
- Added new datasource for fetching upgrade status.


IMPROVEMENTS:
- Updated terraform provider avi with latest gosdk (AV-69259). It has updated options for session handling and fixes the Avi Session and CSRF Token being set during login.
- Added api_timeout property to avi provider definition (AV-68694)


BUG FIXES:
- Miscallaneous bug fixes regarding handling of resources that do not have UUID like useraccount, cluster. It also fixed the crash that happens when any API does not return a JSON value by design.
- Cleanup Terraform VSVIP examples. It is recommended to use proper vsvip resource instead of embedding them inside Virtualservice resource which was supported for legacy purposes. (AV-64687)
- Go FMT changes.


## 0.1.2 (July 12, 2019)
## 0.1.1 (June 21, 2017)

NOTES:

Bumping the provider version to get around provider caching issues - still same functionality

## 0.1.0 (June 21, 2017)

NOTES:

* Same functionality as that of Terraform 0.9.8. Repacked as part of [Provider Splitout](https://www.hashicorp.com/blog/upcoming-provider-changes-in-terraform-0-10/)
