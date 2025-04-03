// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceGslbHSMRuntimeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cluster_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
		"local_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceLocalInfoSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
		"send_interval": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"site_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
	}
}

func resourceAviGslbHSMRuntime() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviGslbHSMRuntimeCreate,
		Read:   ResourceAviGslbHSMRuntimeRead,
		Update: resourceAviGslbHSMRuntimeUpdate,
		Delete: resourceAviGslbHSMRuntimeDelete,
		Schema: ResourceGslbHSMRuntimeSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceGslbHSMRuntimeImporter,
		},
	}
}

func ResourceGslbHSMRuntimeImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceGslbHSMRuntimeSchema()
	return ResourceImporter(d, m, "gslbhsmruntime", s)
}

func ResourceAviGslbHSMRuntimeRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbHSMRuntimeSchema()
	err := APIRead(d, meta, "gslbhsmruntime", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviGslbHSMRuntimeCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbHSMRuntimeSchema()
	err := APICreate(d, meta, "gslbhsmruntime", s)
	if err == nil {
		err = ResourceAviGslbHSMRuntimeRead(d, meta)
	}
	return err
}

func resourceAviGslbHSMRuntimeUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbHSMRuntimeSchema()
	var err error
	err = APIUpdate(d, meta, "gslbhsmruntime", s)
	if err == nil {
		err = ResourceAviGslbHSMRuntimeRead(d, meta)
	}
	return err
}

func resourceAviGslbHSMRuntimeDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "gslbhsmruntime")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
