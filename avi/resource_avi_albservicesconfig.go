// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
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
		"operations_config": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceOperationsConfigSchema(),
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
			Required: true,
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
	err := APICreate(d, meta, "albservicesconfig", s)
	if err == nil {
		err = ResourceAviALBServicesConfigRead(d, meta)
	}
	return err
}

func resourceAviALBServicesConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesConfigSchema()
	var err error
	err = APIUpdate(d, meta, "albservicesconfig", s)
	if err == nil {
		err = ResourceAviALBServicesConfigRead(d, meta)
	}
	return err
}

func resourceAviALBServicesConfigDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "albservicesconfig")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
