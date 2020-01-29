/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
)

func ResourceIpamDnsProviderProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allocate_ip_in_vrf": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"aws_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpamDnsAwsProfileSchema(),
		},
		"azure_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpamDnsAzureProfileSchema(),
		},
		"custom_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpamDnsCustomProfileSchema(),
		},
		"gcp_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpamDnsGCPProfileSchema(),
		},
		"infoblox_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpamDnsInfobloxProfileSchema(),
		},
		"internal_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpamDnsInternalProfileSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"oci_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpamDnsOCIProfileSchema(),
		},
		"openstack_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpamDnsOpenstackProfileSchema(),
		},
		"proxy_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceProxyConfigurationSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tencent_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpamDnsTencentProfileSchema(),
		},
		"type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviIpamDnsProviderProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviIpamDnsProviderProfileCreate,
		Read:   ResourceAviIpamDnsProviderProfileRead,
		Update: resourceAviIpamDnsProviderProfileUpdate,
		Delete: resourceAviIpamDnsProviderProfileDelete,
		Schema: ResourceIpamDnsProviderProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceIpamDnsProviderProfileImporter,
		},
	}
}

func ResourceIpamDnsProviderProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceIpamDnsProviderProfileSchema()
	return ResourceImporter(d, m, "ipamdnsproviderprofile", s)
}

func ResourceAviIpamDnsProviderProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIpamDnsProviderProfileSchema()
	err := ApiRead(d, meta, "ipamdnsproviderprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviIpamDnsProviderProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIpamDnsProviderProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "ipamdnsproviderprofile", s)
	if err == nil {
		err = ResourceAviIpamDnsProviderProfileRead(d, meta)
	}
	return err
}

func resourceAviIpamDnsProviderProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIpamDnsProviderProfileSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "ipamdnsproviderprofile", s)
	if err == nil {
		err = ResourceAviIpamDnsProviderProfileRead(d, meta)
	}
	return err
}

func resourceAviIpamDnsProviderProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "ipamdnsproviderprofile"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviIpamDnsProviderProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
