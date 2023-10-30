// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
	"log"
	"time"
)

func ResourceClusterSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"nodes": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     ResourceClusterNodeSchema(),
		},
		"rejoin_nodes_automatically": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
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
		"virtual_ip": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceIpAddrSchema(),
		},
		"cluster_state": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"up_since": {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
					"progress": {
						Type:         schema.TypeString,
						Optional:     true,
						Computed:     true,
						ValidateFunc: validateInteger,
					},
					"state": {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
					},
				},
			},
		},
	}
}

func resourceAviCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviClusterCreate,
		Read:   ResourceAviClusterRead,
		Update: resourceAviClusterUpdate,
		Delete: resourceAviClusterDelete,
		Schema: ResourceClusterSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceClusterImporter,
		},
	}
}

func ResourceClusterImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceClusterSchema()
	return ResourceImporter(d, m, "cluster", s)
}

func ResourceAviClusterRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceClusterSchema()
	err := readClusterState(d, meta, s)
	if err == nil {
		err := APIRead(d, meta, "cluster", s)
		if err != nil {
			log.Printf("[ERROR] in reading object %v\n", err)
			return err
		}
	} else {
		log.Printf("[ERROR] in Updateing cluster state object %v\n", err)
	}
	return err
}

func resourceAviClusterCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceClusterSchema()
	err := APICreateOrUpdate(d, meta, "cluster", s)
	// Added wait for cluster initialization process as cluster initialization starts after few seconds.
	// This is necessary to store correct state of initialized cluster.
	time.Sleep(90 * time.Second)
	if err == nil {
		err = ResourceAviClusterRead(d, meta)
	}
	return err
}

func resourceAviClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceClusterSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "cluster", s)
	// Added wait for cluster initialization process as cluster initialization starts after few seconds.
	// This is necessary to store correct state of initialized cluster.
	time.Sleep(90 * time.Second)
	if err == nil {
		err = ResourceAviClusterRead(d, meta)
	}
	return err
}

func resourceAviClusterDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARNING] WE can not delete cluster.")
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "cluster")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}

// Function to get the cluster state and update the exact cluster state into d
func readClusterState(d *schema.ResourceData, meta interface{}, s map[string]*schema.Schema) error {
	client := meta.(*clients.AviClient)
	var err error
	var robj interface{}
	if err = client.AviSession.Get("api/cluster/runtime", &robj); err == nil {
		if localData, err := SchemaToAviData(d, s); err == nil {
			if modAPIRes, err := SetDefaultsInAPIRes(robj, localData, s); err == nil {
				if modAPIRes, err = PreprocessAPIRes(modAPIRes, s); err == nil {
					if _, err := APIDataToSchema(modAPIRes, d, s); err != nil {
						log.Printf("[ERROR] Converting APIDataToSchema object %v\n", err)
						return err
					}
				} else {
					log.Printf("[ERROR] Update Cluster State in PreprocessAPIRes api response object %v\n", err)
					return err
				}
			} else {
				log.Printf("[ERROR] Update Cluster State in modifying api response object %v\n", err)
				return err
			}
		} else {
			return err
		}
	}
	return err
}
