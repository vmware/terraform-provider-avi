// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceALBServicesConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_signature_config": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceAppSignatureConfigSchema(),
		},
		"asset_contact": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceALBServicesUserSchema(),
		},
		"case_config": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceCaseConfigSchema(),
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"feature_opt_in_status": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourcePortalFeatureOptInSchema(),
		},
		"ip_reputation_config": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceIpReputationConfigSchema(),
		},
		"mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "MYVMWARE",
		},
		"polling_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "10",
			ValidateFunc: validateInteger,
		},
		"portal_url": {
			Type:     schema.TypeString,
			Required: true,
		},
		"saas_licensing_config": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSaasLicensingInfoSchema(),
		},
		"split_proxy_configuration": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceProxyConfigurationSchema(),
		},
		"use_split_proxy": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"use_tls": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"user_agent_db_config": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceUserAgentDBConfigSchema(),
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"waf_config": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceWafCrsConfigSchema(),
		},
	}
}

func resourceAviALBServicesConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviALBServicesConfigCreate,
		Read:   ResourceAviALBServicesConfigRead,
		Update: resourceAviALBServicesConfigUpdate,
		Delete: resourceAviALBServicesConfigDelete,
		Schema: ResourceALBServicesConfigSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceALBServicesConfigImporter,
		},
	}
}

func ResourceALBServicesConfigImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceALBServicesConfigSchema()
	return ResourceImporter(d, m, "albservicesconfig", s)
}

func ResourceAviALBServicesConfigRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesConfigSchema()
	err := APIRead(d, meta, "albservicesconfig", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviALBServicesConfigCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesConfigSchema()
	err := APICreateOrUpdate(d, meta, "albservicesconfig", s)
	if err == nil {
		err = ResourceAviALBServicesConfigRead(d, meta)
	}
	return err
}

func resourceAviALBServicesConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesConfigSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "albservicesconfig", s)
	if err == nil {
		err = ResourceAviALBServicesConfigRead(d, meta)
	}
	return err
}

func resourceAviALBServicesConfigDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "albservicesconfig"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviALBServicesConfigDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
