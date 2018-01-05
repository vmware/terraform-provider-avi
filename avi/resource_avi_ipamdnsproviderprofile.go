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

func ResourceIpamDnsProviderProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allocate_ip_in_vrf": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"aws_profile": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceIpamDnsAwsProfileSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"azure_profile": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceIpamDnsAzureProfileSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"custom_profile": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceIpamDnsCustomProfileSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"gcp_profile": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceIpamDnsGCPProfileSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"infoblox_profile": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceIpamDnsInfobloxProfileSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"internal_profile": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceIpamDnsInternalProfileSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"openstack_profile": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceIpamDnsOpenstackProfileSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"proxy_configuration": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceProxyConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"type": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"uuid": &schema.Schema{
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
	}
}

func ResourceAviIpamDnsProviderProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceIpamDnsProviderProfileSchema()
	err := ApiRead(d, meta, "ipamdnsproviderprofile", s)
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
	err := ApiCreateOrUpdate(d, meta, "ipamdnsproviderprofile", s)
	if err == nil {
		err = ResourceAviIpamDnsProviderProfileRead(d, meta)
	}
	return err
}

func resourceAviIpamDnsProviderProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "ipamdnsproviderprofile"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviIpamDnsProviderProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
