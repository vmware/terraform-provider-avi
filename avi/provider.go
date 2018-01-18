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
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"avi_username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_USERNAME", nil),
				Description: "Username for Avi Controller.",
			},
			"avi_controller": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_CONTROLLER", nil),
				Description: "Avi Controller hostname or IP address.",
			},
			"avi_password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_PASSWORD", nil),
				Description: "Password for Avi Controller.",
			},
			"avi_tenant": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AVI_TENANT", nil),
				Description: "Avi tenant for Avi Controller.",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"avi_useraccountprofile":            dataSourceAviUserAccountProfile(),
			"avi_role":                          dataSourceAviRole(),
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
			"avi_scheduler":                     dataSourceAviScheduler(),
			"avi_backupconfiguration":           dataSourceAviBackupConfiguration(),
			"avi_serviceenginegroup":            dataSourceAviServiceEngineGroup(),
			"avi_tenant":                        dataSourceAviTenant(),
			"avi_dnspolicy":                     dataSourceAviDnsPolicy(),
			"avi_hardwaresecuritymodulegroup":   dataSourceAviHardwareSecurityModuleGroup(),
			"avi_applicationpersistenceprofile": dataSourceAviApplicationPersistenceProfile(),
			"avi_vrfcontext":                    dataSourceAviVrfContext(),
			"avi_cloudproperties":               dataSourceAviCloudProperties(),
			"avi_customipamdnsprofile":          dataSourceAviCustomIpamDnsProfile(),
			"avi_ipamdnsproviderprofile":        dataSourceAviIpamDnsProviderProfile(),
			"avi_backup":                        dataSourceAviBackup(),
			"avi_networksecuritypolicy":         dataSourceAviNetworkSecurityPolicy(),
			"avi_seproperties":                  dataSourceAviSeProperties(),
			"avi_gslbgeodbprofile":              dataSourceAviGslbGeoDbProfile(),
			"avi_gslbservice":                   dataSourceAviGslbService(),
			"avi_gslb":                          dataSourceAviGslb(),
			"avi_cluster":                       dataSourceAviCluster(),
			"avi_clusterclouddetails":           dataSourceAviClusterCloudDetails(),
			"avi_wafpolicy":                     dataSourceAviWafPolicy(),
			"avi_wafprofile":                    dataSourceAviWafProfile(),
			"avi_snmptrapprofile":               dataSourceAviSnmpTrapProfile(),
			"avi_systemconfiguration":           dataSourceAviSystemConfiguration(),
			"avi_networkprofile":                dataSourceAviNetworkProfile(),
			"avi_errorpagebody":                 dataSourceAviErrorPageBody(),
			"avi_errorpageprofile":              dataSourceAviErrorPageProfile(),
			"avi_serverautoscalepolicy":         dataSourceAviServerAutoScalePolicy(),
			"avi_autoscalelaunchconfig":         dataSourceAviAutoScaleLaunchConfig(),
			"avi_healthmonitor":                 dataSourceAviHealthMonitor(),
			"avi_analyticsprofile":              dataSourceAviAnalyticsProfile(),
			"avi_cloud":                         dataSourceAviCloud(),
			"avi_cloudconnectoruser":            dataSourceAviCloudConnectorUser(),
			"avi_virtualservice":                dataSourceAviVirtualService(),
			"avi_vsvip":                         dataSourceAviVsVip(),
			"avi_alertsyslogconfig":             dataSourceAviAlertSyslogConfig(),
			"avi_alertscriptconfig":             dataSourceAviAlertScriptConfig(),
			"avi_alertconfig":                   dataSourceAviAlertConfig(),
			"avi_actiongroupconfig":             dataSourceAviActionGroupConfig(),
			"avi_alertemailconfig":              dataSourceAviAlertEmailConfig(),
			"avi_vsdatascriptset":               dataSourceAviVSDataScriptSet(),
			"avi_poolgroup":                     dataSourceAviPoolGroup(),
			"avi_prioritylabels":                dataSourceAviPriorityLabels(),
			"avi_poolgroupdeploymentpolicy":     dataSourceAviPoolGroupDeploymentPolicy(),
			"avi_pool":                          dataSourceAviPool(),
			"avi_network":                       dataSourceAviNetwork(),
			"avi_controllerproperties":          dataSourceAviControllerProperties(),
			"avi_applicationprofile":            dataSourceAviApplicationProfile(),
			"avi_httppolicyset":                 dataSourceAviHTTPPolicySet(),
			"avi_serviceengine":                 dataSourceAviServiceEngine(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"avi_useraccountprofile":            resourceAviUserAccountProfile(),
			"avi_role":                          resourceAviRole(),
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
			"avi_scheduler":                     resourceAviScheduler(),
			"avi_backupconfiguration":           resourceAviBackupConfiguration(),
			"avi_serviceenginegroup":            resourceAviServiceEngineGroup(),
			"avi_tenant":                        resourceAviTenant(),
			"avi_dnspolicy":                     resourceAviDnsPolicy(),
			"avi_hardwaresecuritymodulegroup":   resourceAviHardwareSecurityModuleGroup(),
			"avi_applicationpersistenceprofile": resourceAviApplicationPersistenceProfile(),
			"avi_vrfcontext":                    resourceAviVrfContext(),
			"avi_cloudproperties":               resourceAviCloudProperties(),
			"avi_customipamdnsprofile":          resourceAviCustomIpamDnsProfile(),
			"avi_ipamdnsproviderprofile":        resourceAviIpamDnsProviderProfile(),
			"avi_backup":                        resourceAviBackup(),
			"avi_networksecuritypolicy":         resourceAviNetworkSecurityPolicy(),
			"avi_seproperties":                  resourceAviSeProperties(),
			"avi_gslbgeodbprofile":              resourceAviGslbGeoDbProfile(),
			"avi_gslbservice":                   resourceAviGslbService(),
			"avi_gslb":                          resourceAviGslb(),
			"avi_cluster":                       resourceAviCluster(),
			"avi_clusterclouddetails":           resourceAviClusterCloudDetails(),
			"avi_wafpolicy":                     resourceAviWafPolicy(),
			"avi_wafprofile":                    resourceAviWafProfile(),
			"avi_snmptrapprofile":               resourceAviSnmpTrapProfile(),
			"avi_systemconfiguration":           resourceAviSystemConfiguration(),
			"avi_networkprofile":                resourceAviNetworkProfile(),
			"avi_errorpagebody":                 resourceAviErrorPageBody(),
			"avi_errorpageprofile":              resourceAviErrorPageProfile(),
			"avi_serverautoscalepolicy":         resourceAviServerAutoScalePolicy(),
			"avi_autoscalelaunchconfig":         resourceAviAutoScaleLaunchConfig(),
			"avi_healthmonitor":                 resourceAviHealthMonitor(),
			"avi_analyticsprofile":              resourceAviAnalyticsProfile(),
			"avi_cloud":                         resourceAviCloud(),
			"avi_cloudconnectoruser":            resourceAviCloudConnectorUser(),
			"avi_virtualservice":                resourceAviVirtualService(),
			"avi_vsvip":                         resourceAviVsVip(),
			"avi_alertsyslogconfig":             resourceAviAlertSyslogConfig(),
			"avi_alertscriptconfig":             resourceAviAlertScriptConfig(),
			"avi_alertconfig":                   resourceAviAlertConfig(),
			"avi_actiongroupconfig":             resourceAviActionGroupConfig(),
			"avi_alertemailconfig":              resourceAviAlertEmailConfig(),
			"avi_vsdatascriptset":               resourceAviVSDataScriptSet(),
			"avi_poolgroup":                     resourceAviPoolGroup(),
			"avi_prioritylabels":                resourceAviPriorityLabels(),
			"avi_poolgroupdeploymentpolicy":     resourceAviPoolGroupDeploymentPolicy(),
			"avi_pool":                          resourceAviPool(),
			"avi_network":                       resourceAviNetwork(),
			"avi_controllerproperties":          resourceAviControllerProperties(),
			"avi_applicationprofile":            resourceAviApplicationProfile(),
			"avi_httppolicyset":                 resourceAviHTTPPolicySet(),
			"avi_serviceengine":                 resourceAviServiceEngine(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Credentials{
		Username:   d.Get("avi_username").(string),
		Password:   d.Get("avi_password").(string),
		Controller: d.Get("avi_controller").(string),
		Tenant:     d.Get("avi_tenant").(string),
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	aviClient, err := clients.NewAviClient(
		config.Controller, config.Username,
		session.SetPassword(config.Password),
		session.SetTenant(config.Tenant),
		session.SetVersion("17.2.6"),
		session.SetInsecure)

	log.Println("Avi Client created ")

	return aviClient, err
}

type Credentials struct {
	Username   string
	Password   string
	Controller string
	Port       string
	Tenant     string
}

func (c *Credentials) validate() error {
	var err *multierror.Error

	if c.Controller == "" {
		err = multierror.Append(err, fmt.Errorf("Avi Controller must be provided"))
	}

	if c.Username == "" {
		err = multierror.Append(err, fmt.Errorf("Avi Controller username must be provided"))
	}

	if c.Password == "" {
		err = multierror.Append(err, fmt.Errorf("Avi Controller password must be provided"))
	}
	return err.ErrorOrNil()
}
