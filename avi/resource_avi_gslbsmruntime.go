// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceGslbSMRuntimeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_leader": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"cluster_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"dns_configs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDNSConfigSchema(),
		},
		"dns_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceGslbDnsInfoSchema(),
		},
		"enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"events": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceEventInfoSchema(),
		},
		"health_monitor_info": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"leader_cluster_uuid": {
			Type:     schema.TypeString,
			Required: true,
		},
		"member_type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "GSLB_PASSIVE_MEMBER",
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"node_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"num_of_retries": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"obj_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"oper_status": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceOperationalStatusSchema(),
		},
		"remote_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceRemoteInfoSchema(),
		},
		"role": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "GSLB_NOT_A_MEMBER",
		},
		"site_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"site_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"state": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SITE_STATE_NULL",
		},
		"sw_version": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "Not-Initialized",
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"view_id": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
	}
}

func resourceAviGslbSMRuntime() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviGslbSMRuntimeCreate,
		Read:   ResourceAviGslbSMRuntimeRead,
		Update: resourceAviGslbSMRuntimeUpdate,
		Delete: resourceAviGslbSMRuntimeDelete,
		Schema: ResourceGslbSMRuntimeSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceGslbSMRuntimeImporter,
		},
	}
}

func ResourceGslbSMRuntimeImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceGslbSMRuntimeSchema()
	return ResourceImporter(d, m, "gslbsmruntime", s)
}

func ResourceAviGslbSMRuntimeRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbSMRuntimeSchema()
	err := APIRead(d, meta, "gslbsmruntime", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviGslbSMRuntimeCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbSMRuntimeSchema()
	err := APICreate(d, meta, "gslbsmruntime", s)
	if err == nil {
		err = ResourceAviGslbSMRuntimeRead(d, meta)
	}
	return err
}

func resourceAviGslbSMRuntimeUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbSMRuntimeSchema()
	var err error
	err = APIUpdate(d, meta, "gslbsmruntime", s)
	if err == nil {
		err = ResourceAviGslbSMRuntimeRead(d, meta)
	}
	return err
}

func resourceAviGslbSMRuntimeDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "gslbsmruntime")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
