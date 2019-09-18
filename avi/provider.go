/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"fmt"
	"github.com/avinetworks/sdk/go/clients"
	"github.com/avinetworks/sdk/go/session"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"log"
	"time"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"avi_username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_USERNAME", nil),
				Description: "Username for Avi Controller.",
			},
			"avi_controller": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_CONTROLLER", nil),
				Description: "Avi Controller hostname or IP address.",
			},
			"avi_password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_PASSWORD", nil),
				Description: "Password for Avi Controller.",
			},
			"avi_tenant": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_TENANT", nil),
				Description: "Avi tenant for Avi Controller.",
			},
			"avi_version": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_VERSION", nil),
				Description: "Avi version for Avi Controller.",
			},
			"avi_authtoken": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_AUTHTOKEN", nil),
				Description: "Avi token for Avi Controller.",
			},
			"avi_api_timeout": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_API_TIMEOUT", nil),
				Description: "Session timeout for Avi Controller.",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"avi_useraccountprofile":            dataSourceAviUserAccountProfile(),
			"avi_role":                          dataSourceAviRole(),
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
			"avi_serviceenginegroup":            dataSourceAviServiceEngineGroup(),
			"avi_networkservice":                dataSourceAviNetworkService(),
			"avi_dnspolicy":                     dataSourceAviDnsPolicy(),
			"avi_tenant":                        dataSourceAviTenant(),
			"avi_hardwaresecuritymodulegroup":   dataSourceAviHardwareSecurityModuleGroup(),
			"avi_upgradestatusinfo":             dataSourceAviUpgradeStatusInfo(),
			"avi_vrfcontext":                    dataSourceAviVrfContext(),
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
			"avi_applicationprofile":            dataSourceAviApplicationProfile(),
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
			"avi_healthmonitor":                 dataSourceAviHealthMonitor(),
			"avi_httppolicyset":                 dataSourceAviHTTPPolicySet(),
			"avi_serviceengine":                 dataSourceAviServiceEngine(),
			"avi_fileservice":                   dataSourceAviFileService(),
			"avi_server":                        dataSourceAviServer(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"avi_useraccountprofile":            resourceAviUserAccountProfile(),
			"avi_role":                          resourceAviRole(),
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
			"avi_serviceenginegroup":            resourceAviServiceEngineGroup(),
			"avi_networkservice":                resourceAviNetworkService(),
			"avi_dnspolicy":                     resourceAviDnsPolicy(),
			"avi_tenant":                        resourceAviTenant(),
			"avi_hardwaresecuritymodulegroup":   resourceAviHardwareSecurityModuleGroup(),
			"avi_upgradestatusinfo":             resourceAviUpgradeStatusInfo(),
			"avi_vrfcontext":                    resourceAviVrfContext(),
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
			"avi_applicationprofile":            resourceAviApplicationProfile(),
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
			"avi_healthmonitor":                 resourceAviHealthMonitor(),
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
		Version:    "18.2.2",
		AuthToken:  d.Get("avi_authtoken").(string),
		Timeout:    time.Duration(time.Duration(d.Get("avi_api_timeout").(int)) * time.Second),
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
		config.Timeout = time.Duration(time.Duration(timeout.(int)) * time.Second)
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	aviClient, err := clients.NewAviClient(
		config.Controller, config.Username,
		session.SetPassword(config.Password),
		session.SetTenant(config.Tenant),
		session.SetVersion(config.Version),
		session.SetAuthToken(config.AuthToken),
		session.SetInsecure, session.SetTimeout(config.Timeout))

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

func (c *Credentials) validate() error {
	var err *multierror.Error

	if c.Controller == "" {
		err = multierror.Append(err, fmt.Errorf("Avi Controller must be provided"))
	}

	if c.Password == "" && c.AuthToken == "" {
		err = multierror.Append(err, fmt.Errorf("Avi Controller password or authtoken must be provided"))
	}
	return err.ErrorOrNil()
}
