/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"time"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/avinetworks/sdk/go/session"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"avi_username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_USERNAME", nil),
				Description: "Username for Avi Controller.",
			},
			"avi_controller": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_CONTROLLER", nil),
				Description: "Avi Controller hostname or IP address.",
			},
			"avi_password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_PASSWORD", nil),
				Description: "Password for Avi Controller.",
			},
			"avi_tenant": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_TENANT", nil),
				Description: "Avi tenant for Avi Controller.",
			},
			"avi_version": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_VERSION", nil),
				Description: "Avi version for Avi Controller.",
			},
			"avi_authtoken": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_AUTHTOKEN", nil),
				Description: "Avi token for Avi Controller.",
			},
			"avi_api_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_API_TIMEOUT", nil),
				Description: "Session timeout for Avi Controller.",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"avi_rmcloudopsproto":                 dataSourceAviRmCloudOpsProto(),
			"avi_snmptrapprofile":                 dataSourceAviSnmpTrapProfile(),
			"avi_gslb":                            dataSourceAviGslb(),
			"avi_gslbservice":                     dataSourceAviGslbService(),
			"avi_gslbgeodbprofile":                dataSourceAviGslbGeoDbProfile(),
			"avi_healthmonitor":                   dataSourceAviHealthMonitor(),
			"avi_albservicesconfig":               dataSourceAviALBServicesConfig(),
			"avi_network":                         dataSourceAviNetwork(),
			"avi_federationcheckpoint":            dataSourceAviFederationCheckpoint(),
			"avi_siteversion":                     dataSourceAviSiteVersion(),
			"avi_networkservice":                  dataSourceAviNetworkService(),
			"avi_securitypolicy":                  dataSourceAviSecurityPolicy(),
			"avi_vsvip":                           dataSourceAviVsVip(),
			"avi_virtualservice":                  dataSourceAviVirtualService(),
			"avi_serviceenginegroup":              dataSourceAviServiceEngineGroup(),
			"avi_errorpagebody":                   dataSourceAviErrorPageBody(),
			"avi_errorpageprofile":                dataSourceAviErrorPageProfile(),
			"avi_fileobject":                      dataSourceAviFileObject(),
			"avi_httppolicyset":                   dataSourceAviHTTPPolicySet(),
			"avi_systemlimits":                    dataSourceAviSystemLimits(),
			"avi_upgradestatusinfo":               dataSourceAviUpgradeStatusInfo(),
			"avi_upgradestatussummary":            dataSourceAviUpgradeStatusSummary(),
			"avi_licenseledgerdetails":            dataSourceAviLicenseLedgerDetails(),
			"avi_ssopolicy":                       dataSourceAviSSOPolicy(),
			"avi_authprofile":                     dataSourceAviAuthProfile(),
			"avi_controllerproperties":            dataSourceAviControllerProperties(),
			"avi_l4policyset":                     dataSourceAviL4PolicySet(),
			"avi_systemconfiguration":             dataSourceAviSystemConfiguration(),
			"avi_controllersite":                  dataSourceAviControllerSite(),
			"avi_alertsyslogconfig":               dataSourceAviAlertSyslogConfig(),
			"avi_alertemailconfig":                dataSourceAviAlertEmailConfig(),
			"avi_alertscriptconfig":               dataSourceAviAlertScriptConfig(),
			"avi_actiongroupconfig":               dataSourceAviActionGroupConfig(),
			"avi_alertconfig":                     dataSourceAviAlertConfig(),
			"avi_vrfcontext":                      dataSourceAviVrfContext(),
			"avi_applicationprofile":              dataSourceAviApplicationProfile(),
			"avi_ipaddrgroup":                     dataSourceAviIpAddrGroup(),
			"avi_stringgroup":                     dataSourceAviStringGroup(),
			"avi_microservicegroup":               dataSourceAviMicroServiceGroup(),
			"avi_sslprofile":                      dataSourceAviSSLProfile(),
			"avi_certificatemanagementprofile":    dataSourceAviCertificateManagementProfile(),
			"avi_sslkeyandcertificate":            dataSourceAviSSLKeyAndCertificate(),
			"avi_pkiprofile":                      dataSourceAviPKIProfile(),
			"avi_tenant":                          dataSourceAviTenant(),
			"avi_vsdatascriptset":                 dataSourceAviVSDataScriptSet(),
			"avi_cluster":                         dataSourceAviCluster(),
			"avi_clusterclouddetails":             dataSourceAviClusterCloudDetails(),
			"avi_role":                            dataSourceAviRole(),
			"avi_useraccountprofile":              dataSourceAviUserAccountProfile(),
			"avi_user":                            dataSourceAviUser(),
			"avi_analyticsprofile":                dataSourceAviAnalyticsProfile(),
			"avi_wafprofile":                      dataSourceAviWafProfile(),
			"avi_wafcrs":                          dataSourceAviWafCRS(),
			"avi_wafpolicypsmgroup":               dataSourceAviWafPolicyPSMGroup(),
			"avi_wafapplicationsignatureprovider": dataSourceAviWafApplicationSignatureProvider(),
			"avi_wafpolicy":                       dataSourceAviWafPolicy(),
			"avi_pool":                            dataSourceAviPool(),
			"avi_prioritylabels":                  dataSourceAviPriorityLabels(),
			"avi_poolgroupdeploymentpolicy":       dataSourceAviPoolGroupDeploymentPolicy(),
			"avi_poolgroup":                       dataSourceAviPoolGroup(),
			"avi_networksecuritypolicy":           dataSourceAviNetworkSecurityPolicy(),
			"avi_serviceengine":                   dataSourceAviServiceEngine(),
			"avi_icapprofile":                     dataSourceAviIcapProfile(),
			"avi_hardwaresecuritymodulegroup":     dataSourceAviHardwareSecurityModuleGroup(),
			"avi_autoscalelaunchconfig":           dataSourceAviAutoScaleLaunchConfig(),
			"avi_serverautoscalepolicy":           dataSourceAviServerAutoScalePolicy(),
			"avi_image":                           dataSourceAviImage(),
			"avi_cloud":                           dataSourceAviCloud(),
			"avi_cloudconnectoruser":              dataSourceAviCloudConnectorUser(),
			"avi_availabilityzone":                dataSourceAviAvailabilityZone(),
			"avi_vcenterserver":                   dataSourceAviVCenterServer(),
			"avi_networkprofile":                  dataSourceAviNetworkProfile(),
			"avi_dnspolicy":                       dataSourceAviDnsPolicy(),
			"avi_jwtserverprofile":                dataSourceAviJWTServerProfile(),
			"avi_nsxtsegmentruntime":              dataSourceAviNsxtSegmentRuntime(),
			"avi_ipamdnsproviderprofile":          dataSourceAviIpamDnsProviderProfile(),
			"avi_customipamdnsprofile":            dataSourceAviCustomIpamDnsProfile(),
			"avi_backupconfiguration":             dataSourceAviBackupConfiguration(),
			"avi_scheduler":                       dataSourceAviScheduler(),
			"avi_ipreputationdb":                  dataSourceAviIPReputationDB(),
			"avi_applicationpersistenceprofile":   dataSourceAviApplicationPersistenceProfile(),
			"avi_natpolicy":                       dataSourceAviNatPolicy(),
			"avi_trafficcloneprofile":             dataSourceAviTrafficCloneProfile(),
			"avi_testsedatastorelevel1":           dataSourceAviTestSeDatastoreLevel1(),
			"avi_testsedatastorelevel2":           dataSourceAviTestSeDatastoreLevel2(),
			"avi_testsedatastorelevel3":           dataSourceAviTestSeDatastoreLevel3(),
			"avi_albservicesfileupload":           dataSourceAviALBServicesFileUpload(),
			"avi_dynamicdnsrecord":                dataSourceAviDynamicDnsRecord(),
			"avi_seproperties":                    dataSourceAviSeProperties(),
			"avi_cloudproperties":                 dataSourceAviCloudProperties(),
			"avi_controllerportalregistration":    dataSourceAviControllerPortalRegistration(),
			"avi_protocolparser":                  dataSourceAviProtocolParser(),
			"avi_webhook":                         dataSourceAviWebhook(),
			"avi_backup":                          dataSourceAviBackup(),
			"avi_securitymanagerdata":             dataSourceAviSecurityManagerData(),
			"avi_pingaccessagent":                 dataSourceAviPingAccessAgent(),
			"avi_fileservice":                     dataSourceAviFileService(),
			"avi_server":                          dataSourceAviServer(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"avi_rmcloudopsproto":                 resourceAviRmCloudOpsProto(),
			"avi_snmptrapprofile":                 resourceAviSnmpTrapProfile(),
			"avi_gslb":                            resourceAviGslb(),
			"avi_gslbservice":                     resourceAviGslbService(),
			"avi_gslbgeodbprofile":                resourceAviGslbGeoDbProfile(),
			"avi_healthmonitor":                   resourceAviHealthMonitor(),
			"avi_albservicesconfig":               resourceAviALBServicesConfig(),
			"avi_network":                         resourceAviNetwork(),
			"avi_federationcheckpoint":            resourceAviFederationCheckpoint(),
			"avi_siteversion":                     resourceAviSiteVersion(),
			"avi_networkservice":                  resourceAviNetworkService(),
			"avi_securitypolicy":                  resourceAviSecurityPolicy(),
			"avi_vsvip":                           resourceAviVsVip(),
			"avi_virtualservice":                  resourceAviVirtualService(),
			"avi_serviceenginegroup":              resourceAviServiceEngineGroup(),
			"avi_errorpagebody":                   resourceAviErrorPageBody(),
			"avi_errorpageprofile":                resourceAviErrorPageProfile(),
			"avi_fileobject":                      resourceAviFileObject(),
			"avi_httppolicyset":                   resourceAviHTTPPolicySet(),
			"avi_systemlimits":                    resourceAviSystemLimits(),
			"avi_upgradestatusinfo":               resourceAviUpgradeStatusInfo(),
			"avi_upgradestatussummary":            resourceAviUpgradeStatusSummary(),
			"avi_licenseledgerdetails":            resourceAviLicenseLedgerDetails(),
			"avi_ssopolicy":                       resourceAviSSOPolicy(),
			"avi_authprofile":                     resourceAviAuthProfile(),
			"avi_controllerproperties":            resourceAviControllerProperties(),
			"avi_l4policyset":                     resourceAviL4PolicySet(),
			"avi_systemconfiguration":             resourceAviSystemConfiguration(),
			"avi_controllersite":                  resourceAviControllerSite(),
			"avi_alertsyslogconfig":               resourceAviAlertSyslogConfig(),
			"avi_alertemailconfig":                resourceAviAlertEmailConfig(),
			"avi_alertscriptconfig":               resourceAviAlertScriptConfig(),
			"avi_actiongroupconfig":               resourceAviActionGroupConfig(),
			"avi_alertconfig":                     resourceAviAlertConfig(),
			"avi_vrfcontext":                      resourceAviVrfContext(),
			"avi_applicationprofile":              resourceAviApplicationProfile(),
			"avi_ipaddrgroup":                     resourceAviIpAddrGroup(),
			"avi_stringgroup":                     resourceAviStringGroup(),
			"avi_microservicegroup":               resourceAviMicroServiceGroup(),
			"avi_sslprofile":                      resourceAviSSLProfile(),
			"avi_certificatemanagementprofile":    resourceAviCertificateManagementProfile(),
			"avi_sslkeyandcertificate":            resourceAviSSLKeyAndCertificate(),
			"avi_pkiprofile":                      resourceAviPKIProfile(),
			"avi_tenant":                          resourceAviTenant(),
			"avi_vsdatascriptset":                 resourceAviVSDataScriptSet(),
			"avi_cluster":                         resourceAviCluster(),
			"avi_clusterclouddetails":             resourceAviClusterCloudDetails(),
			"avi_role":                            resourceAviRole(),
			"avi_useraccountprofile":              resourceAviUserAccountProfile(),
			"avi_user":                            resourceAviUser(),
			"avi_analyticsprofile":                resourceAviAnalyticsProfile(),
			"avi_wafprofile":                      resourceAviWafProfile(),
			"avi_wafcrs":                          resourceAviWafCRS(),
			"avi_wafpolicypsmgroup":               resourceAviWafPolicyPSMGroup(),
			"avi_wafapplicationsignatureprovider": resourceAviWafApplicationSignatureProvider(),
			"avi_wafpolicy":                       resourceAviWafPolicy(),
			"avi_pool":                            resourceAviPool(),
			"avi_prioritylabels":                  resourceAviPriorityLabels(),
			"avi_poolgroupdeploymentpolicy":       resourceAviPoolGroupDeploymentPolicy(),
			"avi_poolgroup":                       resourceAviPoolGroup(),
			"avi_networksecuritypolicy":           resourceAviNetworkSecurityPolicy(),
			"avi_serviceengine":                   resourceAviServiceEngine(),
			"avi_icapprofile":                     resourceAviIcapProfile(),
			"avi_hardwaresecuritymodulegroup":     resourceAviHardwareSecurityModuleGroup(),
			"avi_autoscalelaunchconfig":           resourceAviAutoScaleLaunchConfig(),
			"avi_serverautoscalepolicy":           resourceAviServerAutoScalePolicy(),
			"avi_image":                           resourceAviImage(),
			"avi_cloud":                           resourceAviCloud(),
			"avi_cloudconnectoruser":              resourceAviCloudConnectorUser(),
			"avi_availabilityzone":                resourceAviAvailabilityZone(),
			"avi_vcenterserver":                   resourceAviVCenterServer(),
			"avi_networkprofile":                  resourceAviNetworkProfile(),
			"avi_dnspolicy":                       resourceAviDnsPolicy(),
			"avi_jwtserverprofile":                resourceAviJWTServerProfile(),
			"avi_nsxtsegmentruntime":              resourceAviNsxtSegmentRuntime(),
			"avi_ipamdnsproviderprofile":          resourceAviIpamDnsProviderProfile(),
			"avi_customipamdnsprofile":            resourceAviCustomIpamDnsProfile(),
			"avi_backupconfiguration":             resourceAviBackupConfiguration(),
			"avi_scheduler":                       resourceAviScheduler(),
			"avi_ipreputationdb":                  resourceAviIPReputationDB(),
			"avi_applicationpersistenceprofile":   resourceAviApplicationPersistenceProfile(),
			"avi_natpolicy":                       resourceAviNatPolicy(),
			"avi_trafficcloneprofile":             resourceAviTrafficCloneProfile(),
			"avi_testsedatastorelevel1":           resourceAviTestSeDatastoreLevel1(),
			"avi_testsedatastorelevel2":           resourceAviTestSeDatastoreLevel2(),
			"avi_testsedatastorelevel3":           resourceAviTestSeDatastoreLevel3(),
			"avi_albservicesfileupload":           resourceAviALBServicesFileUpload(),
			"avi_dynamicdnsrecord":                resourceAviDynamicDnsRecord(),
			"avi_seproperties":                    resourceAviSeProperties(),
			"avi_cloudproperties":                 resourceAviCloudProperties(),
			"avi_controllerportalregistration":    resourceAviControllerPortalRegistration(),
			"avi_protocolparser":                  resourceAviProtocolParser(),
			"avi_webhook":                         resourceAviWebhook(),
			"avi_backup":                          resourceAviBackup(),
			"avi_securitymanagerdata":             resourceAviSecurityManagerData(),
			"avi_pingaccessagent":                 resourceAviPingAccessAgent(),
			"avi_useraccount":                     resourceAviUserAccount(),
			"avi_fileservice":                     resourceAviFileService(),
			"avi_server":                          resourceAviServer(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Credentials{
		Username:   "admin",
		Password:   d.Get("avi_password").(string),
		Controller: d.Get("avi_controller").(string),
		Tenant:     "admin",
		Version:    "18.2.8",
		AuthToken:  d.Get("avi_authtoken").(string),
		Timeout:    time.Duration(d.Get("avi_api_timeout").(int)) * time.Second,
	}
	if username, ok := d.GetOk("avi_username"); ok {
		config.Username = username.(string)
	}
	if version, ok := d.GetOk("avi_version"); ok {
		config.Version = version.(string)
	}

	if tenant, ok := d.GetOk("avi_tenant"); ok {
		config.Tenant = tenant.(string)
	}

	if timeout, ok := d.GetOk("avi_api_timeout"); ok {
		config.Timeout = time.Duration(timeout.(int)) * time.Second
	}

	aviClient, err := clients.NewAviClient(
		config.Controller, config.Username,
		session.SetPassword(config.Password),
		session.SetTenant(config.Tenant),
		session.SetVersion(config.Version),
		session.SetAuthToken(config.AuthToken),
		session.SetInsecure, session.SetTimeout(config.Timeout),
		session.SetLazyAuthentication(true))

	log.Printf("Avi Client created for user %s tenant %s version %s\n",
		config.Username, config.Tenant, config.Version)
	return aviClient, err
}

type Credentials struct {
	Username   string
	Password   string
	Controller string
	Port       string
	Tenant     string
	Version    string
	AuthToken  string
	Timeout    time.Duration
}
