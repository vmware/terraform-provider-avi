/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strings"
)

func ResourceSystemConfigurationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"admin_auth_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceAdminAuthConfigurationSchema(),
		},
		"default_license_tier": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "ENTERPRISE_18",
		},
		"dns_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceDNSConfigurationSchema(),
		},
		"dns_virtualservice_refs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"docker_mode": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"email_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceEmailConfigurationSchema(),
		},
		"global_tenant_config": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceTenantConfigurationSchema(),
		},
		"linux_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceLinuxConfigurationSchema(),
		},
		"mgmt_ip_access_control": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceMgmtIpAccessControlSchema(),
		},
		"ntp_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceNTPConfigurationSchema(),
		},
		"portal_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourcePortalConfigurationSchema(),
		},
		"proxy_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceProxyConfigurationSchema(),
		},
		"secure_channel_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceSecureChannelConfigurationSchema(),
		},
		"snmp_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceSnmpConfigurationSchema(),
		},
		"ssh_ciphers": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"ssh_hmacs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviSystemConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSystemConfigurationCreate,
		Read:   ResourceAviSystemConfigurationRead,
		Update: resourceAviSystemConfigurationUpdate,
		Delete: resourceAviSystemConfigurationDelete,
		Schema: ResourceSystemConfigurationSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceSystemConfigurationImporter,
		},
	}
}

func ResourceSystemConfigurationImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceSystemConfigurationSchema()
	return ResourceImporter(d, m, "systemconfiguration", s)
}

func ResourceAviSystemConfigurationRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSystemConfigurationSchema()
	err := ApiRead(d, meta, "systemconfiguration", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviSystemConfigurationCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSystemConfigurationSchema()
	err := ApiCreateOrUpdate(d, meta, "systemconfiguration", s)
	if err == nil {
		err = ResourceAviSystemConfigurationRead(d, meta)
	}
	return err
}

func resourceAviSystemConfigurationUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSystemConfigurationSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "systemconfiguration", s)
	if err == nil {
		err = ResourceAviSystemConfigurationRead(d, meta)
	}
	return err
}

func resourceAviSystemConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "systemconfiguration"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviSystemConfigurationDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
