/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/avinetworks/sdk/go/models"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strconv"
)

func ResourceAviPoolServerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"pool_ref": {
			Type:     schema.TypeString,
			Required: true,
		},
		"ip": {
			Type:     schema.TypeString,
			Required: true,
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "V4",
		},
		"autoscaling_group_name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"availability_zone": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"discovered_networks": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDiscoveredNetworkSchema(),
		},
		"enabled": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"external_orchestration_id": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"external_uuid": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"hostname": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"location": {
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceGeoLocationSchema(),
		},
		"mac_address": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"nw_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"port": {
			Type:     schema.TypeInt,
			Optional: true,
		},
		"prst_hdr_val": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"ratio": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"resolve_server_by_dns": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"rewrite_host_header": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"server_node": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"static": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"verify_network": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"vm_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviServerCreateOrUpdate,
		Read:   ResourceAviServerRead,
		Update: resourceAviServerCreateOrUpdate,
		Delete: resourceAviServerDelete,
		Schema: ResourceAviPoolServerSchema(),
	}
}

func resourceAviServerCreateOrUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AviClient)
	err, pUUID, poolObj, pserver := resourceAviServerReadApi(d, meta)
	//added check for err and poolObj.
	if err != nil || poolObj == nil {
		log.Printf("[ERROR] resourceAviServerCreateOrUpdate Error during fetching pool object using pool_ref %v", err)
		return err
	}
	if pserver == nil {
		// not found
		newServer := models.Server{}
		if port, ok := d.GetOk("port"); ok {
			portV := int32(port.(int))
			newServer.Port = &portV
		}
		pserver = &newServer
	}
	log.Printf("[INFO] resourceAviServerCreateOrUpdate pool %v server %v", pUUID, pserver)
	//set other attributes from server.
	if hostname, ok := d.GetOk("hostname"); ok {
		hostnameStr := hostname.(string)
		pserver.Hostname = &hostnameStr
	}
	if AutoscalingGroupName, ok := d.GetOk("autoscaling_group_name"); ok {
		agV := AutoscalingGroupName.(string)
		pserver.AutoscalingGroupName = &agV
	}
	if AvailabilityZone, ok := d.GetOk("availability_zone"); ok {
		az := AvailabilityZone.(string)
		pserver.AvailabilityZone = &az
	}
	if Description, ok := d.GetOk("description"); ok {
		desc := Description.(string)
		pserver.Description = &desc
	}
	if DiscoveredNetworks, ok := d.GetOk("discovered_networks"); ok {
		pserver.DiscoveredNetworks = DiscoveredNetworks.([]*models.DiscoveredNetwork)
	}
	if Enabled, ok := d.GetOk("enabled"); ok {
		en := Enabled.(bool)
		pserver.Enabled = &en
	}
	if ExternalOrchestrationID, ok := d.GetOk("external_orchestration_id"); ok {
		extOrc := ExternalOrchestrationID.(string)
		pserver.ExternalOrchestrationID = &extOrc
	}
	if ExternalUUID, ok := d.GetOk("external_uuid"); ok {
		extUUID := ExternalUUID.(string)
		pserver.ExternalUUID = &extUUID
	}
	if Location, ok := d.GetOk("location"); ok {
		pserver.Location = Location.(*models.GeoLocation)
	}
	if MacAddress, ok := d.GetOk("mac_address"); ok {
		mac := MacAddress.(string)
		pserver.MacAddress = &mac
	}
	if NwRef, ok := d.GetOk("nw_ref"); ok {
		nRef := NwRef.(string)
		pserver.NwRef = &nRef
	}
	if PrstHdrVal, ok := d.GetOk("prst_hdr_val"); ok {
		pHdrVal := PrstHdrVal.(string)
		pserver.PrstHdrVal = &pHdrVal
	}
	if Ratio, ok := d.GetOk("ratio"); ok {
		r := int32(Ratio.(int))
		pserver.Ratio = &r
	}
	if ResolveServerByDNS, ok := d.GetOk("resolve_server_by_dns"); ok {
		resolveSvrByDns := ResolveServerByDNS.(bool)
		pserver.ResolveServerByDNS = &resolveSvrByDns
	}
	if RewriteHostHeader, ok := d.GetOk("rewrite_host_header"); ok {
		rewriteHostHdr := RewriteHostHeader.(bool)
		pserver.RewriteHostHeader = &rewriteHostHdr
	}
	if ServerNode, ok := d.GetOk("server_node"); ok {
		sn := ServerNode.(string)
		pserver.ServerNode = &sn
	}
	if Static, ok := d.GetOk("static"); ok {
		s := Static.(bool)
		pserver.Static = &s
	}
	if VerifyNetwork, ok := d.GetOk("verify_network"); ok {
		verifyNet := VerifyNetwork.(bool)
		pserver.VerifyNetwork = &verifyNet
	}
	if VMRef, ok := d.GetOk("vm_ref"); ok {
		vmRefStr := VMRef.(string)
		pserver.VMRef = &vmRefStr
	}
	if t, ok := d.GetOk("type"); ok {
		tStr := t.(string)
		ip := d.Get("ip").(string)
		pserver.IP = &models.IPAddr{Type: &tStr, Addr: &ip}
	}

	uri := "api/pool/" + pUUID
	var response interface{}
	patchPool := models.Pool{}
	patchPool.Name = poolObj.Name
	patchPool.TenantRef = poolObj.TenantRef
	patchPool.CloudRef = poolObj.CloudRef
	patchPool.Servers = append(patchPool.Servers, pserver)
	err = client.AviSession.Patch(uri, patchPool, "add", response)
	log.Printf("[INFO] resourceAviServerCreateOrUpdate pool %v poolobj %v err %v response %v",
		pUUID, patchPool, err, response)
	if err == nil {
		err = ResourceAviServerRead(d, meta)
	}
	return err
}

func ResourceAviServerRead(d *schema.ResourceData, meta interface{}) error {
	err, pUUID, _, pserver := resourceAviServerReadApi(d, meta)
	if err == nil && pserver != nil {
		//Set id to include port number. if port is not in tf then use 0
		var sUUID string
		portStr := "0"
		ip := d.Get("ip").(string)
		if port, ok := d.GetOk("port"); ok {
			portStr = strconv.Itoa(port.(int))
		}
		sUUID = pUUID + ":" + ip + ":" + portStr
		log.Printf("[INFO] pool %v ip %v port %v", pUUID, ip, portStr)
		d.SetId(sUUID)
		// Fill in the server information
		if pserver.Hostname != nil {
			d.Set("hostname", pserver.Hostname)
		}

		d.Set("enabled", pserver.Enabled)

		if pserver.AutoscalingGroupName != nil {
			d.Set("autoscaling_group_name", pserver.AutoscalingGroupName)
		}
		if pserver.AvailabilityZone != nil {
			d.Set("availability_zone", pserver.AvailabilityZone)
		}
		if pserver.Description != nil {
			d.Set("description", pserver.Description)
		}
		if pserver.DiscoveredNetworks != nil {
			d.Set("discovered_networks", pserver.DiscoveredNetworks)
		}

		if pserver.ExternalUUID != nil {
			d.Set("external_uuid", pserver.ExternalUUID)
		}
		if pserver.ExternalOrchestrationID != nil {
			d.Set("external_orchestration_id", pserver.ExternalOrchestrationID)
		}

		if pserver.Location != nil {
			d.Set("mac_address", pserver.Location)
		}
		if pserver.NwRef != nil {
			d.Set("nw_ref", pserver.NwRef)
		}
		if pserver.PrstHdrVal != nil {
			d.Set("prst_hdr_val", pserver.PrstHdrVal)
		}
		d.Set("ratio", pserver.Ratio)
		d.Set("resolve_server_by_dns", pserver.ResolveServerByDNS)
		d.Set("rewrite_host_header", pserver.RewriteHostHeader)
		if pserver.ServerNode != nil {
			d.Set("server_node", pserver.ServerNode)
		}
		d.Set("static", pserver.Static)
		d.Set("verify_network", pserver.VerifyNetwork)
		if pserver.VMRef != nil {
			d.Set("vm_ref", pserver.VMRef)
		}
		// Add more fields to read.
	} else if err != nil {
		d.SetId("")
		log.Printf("[ERROR] ResourceAviServerRead in reading object %v\n", err)
		return nil
	}
	return err
}

func resourceAviServerReadApi(d *schema.ResourceData, meta interface{}) (error, string, *models.Pool, *models.Server) {
	client := meta.(*clients.AviClient)
	pUUID := UUIDFromID(d.Get("pool_ref").(string))
	uri := "api/pool/" + pUUID
	var poolObj *models.Pool
	err := client.AviSession.Get(uri, &poolObj)
	if err != nil {
		log.Printf("[ERROR] pool uuid %v not found", pUUID)
		return err, pUUID, nil, nil
	}
	log.Printf("[INFO] found pool %v", poolObj.Name)
	ip := d.Get("ip").(string)
	port := d.Get("port")

	// find the server in the pool object.
	var matchedServer *models.Server = nil
	for i := 0; i < len(poolObj.Servers); i++ {
		sObj := poolObj.Servers[i]
		if *sObj.IP.Addr == ip {
			if port == nil && sObj.Port == nil {
				matchedServer = sObj
				break
			} else if port != nil && sObj.Port != nil {
				if int32(port.(int)) == *sObj.Port {
					matchedServer = sObj
					break
				}
			}
		}
	}
	return nil, pUUID, poolObj, matchedServer
}

func resourceAviServerDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.AviClient)
	err, pUUID, poolObj, pserver := resourceAviServerReadApi(d, meta)
	log.Printf("[DEBUG] pool %v %v server %v", pUUID, poolObj.Name, d.Id())
	if pserver != nil {
		uri := "api/pool/" + pUUID
		var response interface{}
		var patchPool = make(map[string]interface{})
		var servers = make([]models.Server, 1)
		servers[0] = *pserver
		patchPool["servers"] = servers
		err = client.AviSession.Patch(uri, patchPool, "delete", response)
		log.Printf("[INFO] pool %v server %v deleted err %v", patchPool, d.Id(), err)
	}
	d.SetId("")
	return err
}
