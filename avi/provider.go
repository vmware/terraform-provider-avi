// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

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
			"avi_useraccountprofile":            dataSourceAviUserAccountProfile(),
			"avi_role":                          dataSourceAviRole(),
			"avi_user":                          dataSourceAviUser(),
			"avi_objectaccesspolicy":            dataSourceAviObjectAccessPolicy(),
			"avi_natpolicy":                     dataSourceAviNatPolicy(),
			"avi_image":                         dataSourceAviImage(),
			"avi_ipaddrgroup":                   dataSourceAviIpAddrGroup(),
			"avi_microservicegroup":             dataSourceAviMicroServiceGroup(),
			"avi_stringgroup":                   dataSourceAviStringGroup(),
			"avi_trafficcloneprofile":           dataSourceAviTrafficCloneProfile(),
			"avi_webhook":                       dataSourceAviWebhook(),
			"avi_authprofile":                   dataSourceAviAuthProfile(),
			"avi_sslkeyandcertificate":          dataSourceAviSSLKeyAndCertificate(),
			"avi_sslprofile":                    dataSourceAviSSLProfile(),
			"avi_pkiprofile":                    dataSourceAviPKIProfile(),
			"avi_certificatemanagementprofile":  dataSourceAviCertificateManagementProfile(),
			"avi_ssopolicy":                     dataSourceAviSSOPolicy(),
			"avi_l4policyset":                   dataSourceAviL4PolicySet(),
			"avi_scheduler":                     dataSourceAviScheduler(),
			"avi_backupconfiguration":           dataSourceAviBackupConfiguration(),
			"avi_tenant":                        dataSourceAviTenant(),
			"avi_serviceenginegroup":            dataSourceAviServiceEngineGroup(),
			"avi_networkservice":                dataSourceAviNetworkService(),
			"avi_dnspolicy":                     dataSourceAviDnsPolicy(),
			"avi_hardwaresecuritymodulegroup":   dataSourceAviHardwareSecurityModuleGroup(),
			"avi_upgradestatussummary":          dataSourceAviUpgradeStatusSummary(),
			"avi_upgradestatusinfo":             dataSourceAviUpgradeStatusInfo(),
			"avi_vrfcontext":                    dataSourceAviVrfContext(),
			"avi_testsedatastorelevel1":         dataSourceAviTestSeDatastoreLevel1(),
			"avi_testsedatastorelevel2":         dataSourceAviTestSeDatastoreLevel2(),
			"avi_testsedatastorelevel3":         dataSourceAviTestSeDatastoreLevel3(),
			"avi_securitypolicy":                dataSourceAviSecurityPolicy(),
			"avi_protocolparser":                dataSourceAviProtocolParser(),
			"avi_cloudproperties":               dataSourceAviCloudProperties(),
			"avi_applicationpersistenceprofile": dataSourceAviApplicationPersistenceProfile(),
			"avi_backup":                        dataSourceAviBackup(),
			"avi_networksecuritypolicy":         dataSourceAviNetworkSecurityPolicy(),
			"avi_seproperties":                  dataSourceAviSeProperties(),
			"avi_pingaccessagent":               dataSourceAviPingAccessAgent(),
			"avi_gslbgeodbprofile":              dataSourceAviGslbGeoDbProfile(),
			"avi_gslbservice":                   dataSourceAviGslbService(),
			"avi_gslb":                          dataSourceAviGslb(),
			"avi_cluster":                       dataSourceAviCluster(),
			"avi_clusterclouddetails":           dataSourceAviClusterCloudDetails(),
			"avi_wafpolicy":                     dataSourceAviWafPolicy(),
			"avi_wafcrs":                        dataSourceAviWafCRS(),
			"avi_wafprofile":                    dataSourceAviWafProfile(),
			"avi_wafpolicypsmgroup":             dataSourceAviWafPolicyPSMGroup(),
			"avi_snmptrapprofile":               dataSourceAviSnmpTrapProfile(),
			"avi_systemconfiguration":           dataSourceAviSystemConfiguration(),
			"avi_controllersite":                dataSourceAviControllerSite(),
			"avi_networkprofile":                dataSourceAviNetworkProfile(),
			"avi_errorpagebody":                 dataSourceAviErrorPageBody(),
			"avi_errorpageprofile":              dataSourceAviErrorPageProfile(),
			"avi_customerportalinfo":            dataSourceAviCustomerPortalInfo(),
			"avi_controllerproperties":          dataSourceAviControllerProperties(),
			"avi_healthmonitor":                 dataSourceAviHealthMonitor(),
			"avi_virtualservice":                dataSourceAviVirtualService(),
			"avi_vsvip":                         dataSourceAviVsVip(),
			"avi_analyticsprofile":              dataSourceAviAnalyticsProfile(),
			"avi_cloud":                         dataSourceAviCloud(),
			"avi_cloudconnectoruser":            dataSourceAviCloudConnectorUser(),
			"avi_portalfileupload":              dataSourceAviPortalFileUpload(),
			"avi_alertsyslogconfig":             dataSourceAviAlertSyslogConfig(),
			"avi_alertscriptconfig":             dataSourceAviAlertScriptConfig(),
			"avi_alertconfig":                   dataSourceAviAlertConfig(),
			"avi_actiongroupconfig":             dataSourceAviActionGroupConfig(),
			"avi_alertemailconfig":              dataSourceAviAlertEmailConfig(),
			"avi_vsdatascriptset":               dataSourceAviVSDataScriptSet(),
			"avi_controllerportalregistration":  dataSourceAviControllerPortalRegistration(),
			"avi_customipamdnsprofile":          dataSourceAviCustomIpamDnsProfile(),
			"avi_ipamdnsproviderprofile":        dataSourceAviIpamDnsProviderProfile(),
			"avi_poolgroup":                     dataSourceAviPoolGroup(),
			"avi_prioritylabels":                dataSourceAviPriorityLabels(),
			"avi_poolgroupdeploymentpolicy":     dataSourceAviPoolGroupDeploymentPolicy(),
			"avi_pool":                          dataSourceAviPool(),
			"avi_network":                       dataSourceAviNetwork(),
			"avi_serverautoscalepolicy":         dataSourceAviServerAutoScalePolicy(),
			"avi_autoscalelaunchconfig":         dataSourceAviAutoScaleLaunchConfig(),
			"avi_applicationprofile":            dataSourceAviApplicationProfile(),
			"avi_httppolicyset":                 dataSourceAviHTTPPolicySet(),
			"avi_serviceengine":                 dataSourceAviServiceEngine(),
			"avi_fileservice":                   dataSourceAviFileService(),
			"avi_server":                        dataSourceAviServer(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"avi_useraccountprofile":            resourceAviUserAccountProfile(),
			"avi_role":                          resourceAviRole(),
			"avi_user":                          resourceAviUser(),
			"avi_objectaccesspolicy":            resourceAviObjectAccessPolicy(),
			"avi_natpolicy":                     resourceAviNatPolicy(),
			"avi_image":                         resourceAviImage(),
			"avi_ipaddrgroup":                   resourceAviIpAddrGroup(),
			"avi_microservicegroup":             resourceAviMicroServiceGroup(),
			"avi_stringgroup":                   resourceAviStringGroup(),
			"avi_trafficcloneprofile":           resourceAviTrafficCloneProfile(),
			"avi_webhook":                       resourceAviWebhook(),
			"avi_authprofile":                   resourceAviAuthProfile(),
			"avi_sslkeyandcertificate":          resourceAviSSLKeyAndCertificate(),
			"avi_sslprofile":                    resourceAviSSLProfile(),
			"avi_pkiprofile":                    resourceAviPKIProfile(),
			"avi_certificatemanagementprofile":  resourceAviCertificateManagementProfile(),
			"avi_ssopolicy":                     resourceAviSSOPolicy(),
			"avi_l4policyset":                   resourceAviL4PolicySet(),
			"avi_scheduler":                     resourceAviScheduler(),
			"avi_backupconfiguration":           resourceAviBackupConfiguration(),
			"avi_tenant":                        resourceAviTenant(),
			"avi_serviceenginegroup":            resourceAviServiceEngineGroup(),
			"avi_networkservice":                resourceAviNetworkService(),
			"avi_dnspolicy":                     resourceAviDnsPolicy(),
			"avi_hardwaresecuritymodulegroup":   resourceAviHardwareSecurityModuleGroup(),
			"avi_upgradestatussummary":          resourceAviUpgradeStatusSummary(),
			"avi_upgradestatusinfo":             resourceAviUpgradeStatusInfo(),
			"avi_vrfcontext":                    resourceAviVrfContext(),
			"avi_testsedatastorelevel1":         resourceAviTestSeDatastoreLevel1(),
			"avi_testsedatastorelevel2":         resourceAviTestSeDatastoreLevel2(),
			"avi_testsedatastorelevel3":         resourceAviTestSeDatastoreLevel3(),
			"avi_securitypolicy":                resourceAviSecurityPolicy(),
			"avi_protocolparser":                resourceAviProtocolParser(),
			"avi_cloudproperties":               resourceAviCloudProperties(),
			"avi_applicationpersistenceprofile": resourceAviApplicationPersistenceProfile(),
			"avi_backup":                        resourceAviBackup(),
			"avi_networksecuritypolicy":         resourceAviNetworkSecurityPolicy(),
			"avi_seproperties":                  resourceAviSeProperties(),
			"avi_pingaccessagent":               resourceAviPingAccessAgent(),
			"avi_gslbgeodbprofile":              resourceAviGslbGeoDbProfile(),
			"avi_gslbservice":                   resourceAviGslbService(),
			"avi_gslb":                          resourceAviGslb(),
			"avi_cluster":                       resourceAviCluster(),
			"avi_clusterclouddetails":           resourceAviClusterCloudDetails(),
			"avi_wafpolicy":                     resourceAviWafPolicy(),
			"avi_wafcrs":                        resourceAviWafCRS(),
			"avi_wafprofile":                    resourceAviWafProfile(),
			"avi_wafpolicypsmgroup":             resourceAviWafPolicyPSMGroup(),
			"avi_snmptrapprofile":               resourceAviSnmpTrapProfile(),
			"avi_systemconfiguration":           resourceAviSystemConfiguration(),
			"avi_controllersite":                resourceAviControllerSite(),
			"avi_networkprofile":                resourceAviNetworkProfile(),
			"avi_errorpagebody":                 resourceAviErrorPageBody(),
			"avi_errorpageprofile":              resourceAviErrorPageProfile(),
			"avi_customerportalinfo":            resourceAviCustomerPortalInfo(),
			"avi_controllerproperties":          resourceAviControllerProperties(),
			"avi_healthmonitor":                 resourceAviHealthMonitor(),
			"avi_virtualservice":                resourceAviVirtualService(),
			"avi_vsvip":                         resourceAviVsVip(),
			"avi_analyticsprofile":              resourceAviAnalyticsProfile(),
			"avi_cloud":                         resourceAviCloud(),
			"avi_cloudconnectoruser":            resourceAviCloudConnectorUser(),
			"avi_portalfileupload":              resourceAviPortalFileUpload(),
			"avi_alertsyslogconfig":             resourceAviAlertSyslogConfig(),
			"avi_alertscriptconfig":             resourceAviAlertScriptConfig(),
			"avi_alertconfig":                   resourceAviAlertConfig(),
			"avi_actiongroupconfig":             resourceAviActionGroupConfig(),
			"avi_alertemailconfig":              resourceAviAlertEmailConfig(),
			"avi_vsdatascriptset":               resourceAviVSDataScriptSet(),
			"avi_controllerportalregistration":  resourceAviControllerPortalRegistration(),
			"avi_customipamdnsprofile":          resourceAviCustomIpamDnsProfile(),
			"avi_ipamdnsproviderprofile":        resourceAviIpamDnsProviderProfile(),
			"avi_poolgroup":                     resourceAviPoolGroup(),
			"avi_prioritylabels":                resourceAviPriorityLabels(),
			"avi_poolgroupdeploymentpolicy":     resourceAviPoolGroupDeploymentPolicy(),
			"avi_pool":                          resourceAviPool(),
			"avi_network":                       resourceAviNetwork(),
			"avi_serverautoscalepolicy":         resourceAviServerAutoScalePolicy(),
			"avi_autoscalelaunchconfig":         resourceAviAutoScaleLaunchConfig(),
			"avi_applicationprofile":            resourceAviApplicationProfile(),
			"avi_httppolicyset":                 resourceAviHTTPPolicySet(),
			"avi_serviceengine":                 resourceAviServiceEngine(),
			"avi_useraccount":                   resourceAviUserAccount(),
			"avi_fileservice":                   resourceAviFileService(),
			"avi_server":                        resourceAviServer(),
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
