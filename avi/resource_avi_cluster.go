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
			Optional: true,
			Elem:     ResourceClusterNodeSchema(),
		},
		"rejoin_nodes_automatically": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
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
						Type:     schema.TypeInt,
						Optional: true,
						Computed: true,
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
	err := read_cluster_state(d, meta, s)
	if err == nil {
		err := ApiRead(d, meta, "cluster", s)
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
	err := ApiCreateOrUpdate(d, meta, "cluster", s)
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
	err = ApiCreateOrUpdate(d, meta, "cluster", s)
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
	return nil
}

// Function to get the cluster state and update the exact cluster state into d
func read_cluster_state(d *schema.ResourceData, meta interface{}, s map[string]*schema.Schema) error {
	client := meta.(*clients.AviClient)
	var err error
	var robj interface{}
	if err = client.AviSession.Get("api/cluster/runtime", &robj); err == nil {
		if local_data, err := SchemaToAviData(d, s); err == nil {
			if mod_api_res, err := SetDefaultsInAPIRes(robj, local_data, s); err == nil {
				if _, err := ApiDataToSchema(mod_api_res, d, s); err != nil {
					log.Printf("[ERROR] Converting ApiDataToSchema object %v\n", err)
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
