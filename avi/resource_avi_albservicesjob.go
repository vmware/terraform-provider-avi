// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceALBServicesJobSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"command": {
			Type:     schema.TypeString,
			Required: true,
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"end_time": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceTimeStampSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"params": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceALBServicesJobParamSchema(),
		},
		"pulse_job_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"pulse_sync_status": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"result": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"start_time": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceTimeStampSchema(),
		},
		"status": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "PENDING",
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"token": {
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

func resourceAviALBServicesJob() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviALBServicesJobCreate,
		Read:   ResourceAviALBServicesJobRead,
		Update: resourceAviALBServicesJobUpdate,
		Delete: resourceAviALBServicesJobDelete,
		Schema: ResourceALBServicesJobSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceALBServicesJobImporter,
		},
	}
}

func ResourceALBServicesJobImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceALBServicesJobSchema()
	return ResourceImporter(d, m, "albservicesjob", s)
}

func ResourceAviALBServicesJobRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesJobSchema()
	err := APIRead(d, meta, "albservicesjob", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviALBServicesJobCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesJobSchema()
	err := APICreateOrUpdate(d, meta, "albservicesjob", s)
	if err == nil {
		err = ResourceAviALBServicesJobRead(d, meta)
	}
	return err
}

func resourceAviALBServicesJobUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesJobSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "albservicesjob", s)
	if err == nil {
		err = ResourceAviALBServicesJobRead(d, meta)
	}
	return err
}

func resourceAviALBServicesJobDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "albservicesjob"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviALBServicesJobDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
