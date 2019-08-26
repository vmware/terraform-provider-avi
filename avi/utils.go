/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/avinetworks/sdk/go/session"
	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var post_not_allowed = [...]string{"systemconfiguration", "cluster"}

// It takes the terraform plan data and schema and converts it into Avi JSON
// It recursively resolves the data type of the terraform schema and converts scalar to scalar, Set to dictionary,
// and list to list.
func SchemaToAviData(d interface{}, s map[string]*schema.Schema) (interface{}, error) {
	switch d.(type) {
	default:
	case map[string]interface{}:
		m := make(map[string]interface{})
		for k, v := range d.(map[string]interface{}) {
			if obj, err := SchemaToAviData(v, nil); err == nil && obj != nil && obj != "" {
				m[k] = obj
			} else if err != nil {
				log.Printf("[ERROR] SchemaToAviData %v in parsing k: %v v: %v", err, k, v)
			}
		}
		return m, nil
	case []interface{}:
		var objList []interface{}
		varray := d.([]interface{})
		for i := 0; i < len(varray); i++ {
			obj, err := SchemaToAviData(varray[i], nil)
			if err == nil && obj != nil {
				objList = append(objList, obj)
			}
		}
		if len(objList) == 0 {
			return nil, nil
		}
		return objList, nil

	case *schema.Set:
		if len(d.(*schema.Set).List()) == 0 {
			return nil, nil
		}
		obj, err := SchemaToAviData(d.(*schema.Set).List()[0], nil)
		return obj, err

	case *schema.ResourceData:
		// In this case the top level schema should be present.
		m := make(map[string]interface{})
		r := d.(*schema.ResourceData)
		for k, v := range s {
			if data, ok := r.GetOk(k); ok {
				if obj, err := SchemaToAviData(data, nil); err == nil && obj != nil && obj != "" {
					m[k] = obj
				} else if err != nil {
					log.Printf("[ERROR] SchemaToAviData %v in converting k: %v v: %v", err, k, v)
				}
			}
		}
		return m, nil
	}
	/** Return the same object as there is nothing special about **/
	return d, nil
}

func CommonHash(v interface{}) int {
	return hashcode.String("avi")
}

// It sets default values in the terraform resources to avoid diffs for scalars.
func SetDefaultsInAPIRes(api_res interface{}, d_local interface{}, s map[string]*schema.Schema) (interface{}, error) {
	if api_res == nil {
		log.Printf("[ERROR] SetDefaultsInAPIRes got nil for %v", s)
		return api_res, nil
	}
	switch d_local.(type) {
	default:
	case map[string]interface{}:
		for k, v := range d_local.(map[string]interface{}) {
			switch v.(type) {
			//Getting key, value for given d_local
			default:
				if _, ok := api_res.(map[string]interface{})[k]; !ok {
					//Cheking if field is present in schema
					if dval, ok := s[k]; ok {
						//Getting default values from schema
						default_val, err := dval.DefaultValue()
						if err != nil {
							log.Printf("[ERROR] SetDefaultsInAPIRes did not get default value from schema err %v %v",
								err, default_val)
						} else {
							if default_val != nil {
								api_res.(map[string]interface{})[k] = default_val
							}
						}
					}
				}
			//d_local nested dictionary.
			case map[string]interface{}:
				if s2, ok := s[k]; ok {
					switch s2.Elem.(type) {
					default:
					case *schema.Resource:
						if api_res.(map[string]interface{})[k] != nil {
							api_res1, err := SetDefaultsInAPIRes(api_res.(map[string]interface{})[k], v, s2.Elem.(*schema.Resource).Schema)
							if err != nil {
								log.Printf("[ERROR] SetDefaultsInAPIRes %v", err)
							} else {
								api_res.(map[string]interface{})[k] = api_res1
							}
						} else {
							api_res.(map[string]interface{})[k] = v
						}
					}
				}
			//d_local is array of dictionaries.
			case []interface{}:
				var objList []interface{}
				if api_res.(map[string]interface{})[k] != nil {
					varray2 := api_res.(map[string]interface{})[k].([]interface{})
					//getting schema for nested object.
					s2, err := s[k]
					var dst, src []interface{}
					//As err returned is boolean value
					if !err {
						log.Printf("[ERROR] SetDefaultsInAPIRes in fetching k %v err %v", k, err)
					}
					if len(varray2) > len(v.([]interface{})) {
						dst = varray2
						src = v.([]interface{})
					} else {
						dst = v.([]interface{})
						src = varray2
					}
					for x, y := range src {
						switch s2.Elem.(type) {
						default:
						case *schema.Resource:
							obj, err := SetDefaultsInAPIRes(dst[x], y, s2.Elem.(*schema.Resource).Schema)
							if err != nil {
								log.Printf("[ERROR] SetDefaultsInAPIRes err %v in x %v y %v", err, x, y)
							} else {
								objList = append(objList, obj)
							}
						case *schema.Schema:
							objList = append(objList, src[x])
						}
					}
				}
				api_res.(map[string]interface{})[k] = objList
			}
		}
	}
	return api_res, nil
}

// It takes the Avi JSON data and fills in the terraform data during API read.
// It takes input as the top level schema and it uses that to properly create the corresponding terraform resource data
// It also checks whether a given Avi key is defined in the schema before attempting to fill the data.
func ApiDataToSchema(adata interface{}, d interface{}, t map[string]*schema.Schema) (interface{}, error) {
	switch adata.(type) {
	default:
	case map[string]interface{}:
		// resolve d interface into a set
		if t == nil {
			var m map[string]interface{}
			m = map[string]interface{}{}
			for k, v := range adata.(map[string]interface{}) {
				if obj, err := ApiDataToSchema(v, nil, nil); err == nil {
					m[k] = obj
				} else if err != nil {
					log.Printf("[ERROR] ApiDataToSchema %v in converting k: %v v: %v", err, k, v)
				}
			}
			objs := []interface{}{}
			objs = append(objs, m)
			s := schema.NewSet(CommonHash, objs)
			return s, nil
		} else {
			for k, v := range adata.(map[string]interface{}) {
				if _, ok := t[k]; ok {
					// found in the schema
					if obj, err := ApiDataToSchema(v, nil, nil); err == nil {
						switch d.(type) {
						default:
						case *schema.ResourceData:
							err := d.(*schema.ResourceData).Set(k, obj)
							if err != nil {
								log.Printf("[ERROR] ApiDataToSchema %v in setting %v", err, obj)
							}
						case map[string]interface{}:
							d.(map[string]interface{})[k] = obj
						}
					}
				}
			}
			return d, nil
		}
	case []interface{}:
		var objList []interface{}
		varray := adata.([]interface{})
		for i := 0; i < len(varray); i++ {
			obj, err := ApiDataToSchema(varray[i], nil, nil)
			if err == nil {
				switch obj.(type) {
				default:
					objList = append(objList, obj)
				case *schema.Set:
					objList = append(objList, obj.(*schema.Set).List()[0])
				}
			} else {
				log.Printf("[ERROR] ApiDataToSchema %v", err)
			}
		}
		return objList, nil
		/** Return the same object as there is nothing special about **/
	}
	return adata, nil
}

func SetIDFromObj(d *schema.ResourceData, robj interface{}) {
	uuid := robj.(map[string]interface{})["uuid"].(string)
	d.Set("uuid", uuid)
	if url, ok := robj.(map[string]interface{})["url"].(string); ok && url != "" {
		d.SetId(url)
	} else {
		d.SetId(uuid)
	}
}

// It is generic API to create and update any Avi REST resource. It handles special situations with cloud
// and tenant filters as objects may already be present. If the resource does not exist it will try to
// create it. In case, it is present then automatically converts to PUT semantics.
func ApiCreateOrUpdate(d *schema.ResourceData, meta interface{}, objType string, s map[string]*schema.Schema,
	opts ...bool) error {
	client := meta.(*clients.AviClient)
	var robj interface{}
	obj := d

	usePatchForUpdate := false
	if len(opts) > 0 {
		usePatchForUpdate = opts[0]
	}

	if data, err := SchemaToAviData(obj, s); err == nil {
		path := "api/" + objType
		specialobj := IsPostNotAllowed(objType)
		if specialobj {
			path = path + "?skip_default=true"
			err = client.AviSession.Put(path, data, &robj)
			if err != nil {
				log.Printf("[ERROR] ApiCreateOrUpdate: PUT on %v Error %v path %v id %v\n", objType, err, path,
					d.Id())
			} else {
				SetIDFromObj(d, robj)
			}
		} else if uuid, ok := d.GetOk("uuid"); ok {
			path = path + "/" + uuid.(string) + "?skip_default=true"
			if !usePatchForUpdate {
				err = client.AviSession.Put(path, data, &robj)
			} else {
				err = client.AviSession.Patch(path, data, "replace", &robj)
			}
			if err != nil {
				log.Printf("[ERROR] ApiCreateOrUpdate: PUT Error %v path %v id %v\n", err, path, d.Id())
			}
		} else {
			if name, ok := d.GetOk("name"); ok {
				var existing_obj interface{}
				if cloudRef, ok := d.GetOk("cloud_ref"); ok && strings.Contains(cloudRef.(string),
					"api/cloud/") {
					cloudUUID := strings.SplitN(cloudRef.(string), "api/cloud/", 2)[1]
					// strip the # if it exists
					cloudUUID = strings.Split(cloudUUID, "#")[0]
					log.Printf("[INFO] ApiCreateOrUpdate: using cloud %v for obj %v name %s \n",
						cloudUUID, objType, name)
					err = client.AviSession.GetObject(objType, session.SetName(name.(string)),
						session.SetResult(&existing_obj), session.SetCloudUUID(cloudUUID),
						session.SetSkipDefault(true))
					if err != nil {
						log.Printf("[ERROR] ApiCreateOrUpdate: GET Error %v path %v id %v\n", err, path, d.Id())
					}
				} else {
					log.Printf("[INFO] ApiCreateOrUpdate: reading obj %v name %s \n",
						objType, name)
					err = client.AviSession.GetObject(objType, session.SetName(name.(string)),
						session.SetResult(&existing_obj), session.SetSkipDefault(true))
					if err != nil {
						log.Printf("[ERROR] ApiCreateOrUpdate: GET Error %v path %v id %v\n", err, path, d.Id())
					}
				}

				if existing_obj == nil {
					// object not found
					log.Printf("[INFO] ApiCreateOrUpdate: Creating obj type %v schema %v data %v\n", objType, d,
						data)
					err = client.AviSession.Post(path, data, &robj)
					if err == nil && robj != nil {
						SetIDFromObj(d, robj)
					} else {
						log.Printf("[ERROR] ApiCreateOrUpdate creation failed %v object with name %v\n", err,
							name)
					}
				} else {
					// found existing object.
					SetIDFromObj(d, existing_obj)
					uuid = existing_obj.(map[string]interface{})["uuid"].(string)
					path = path + "/" + uuid.(string) + "?skip_default=true"
					if !usePatchForUpdate {
						err = client.AviSession.Put(path, data, &robj)
					} else {
						err = client.AviSession.Patch(path, data, "replace", &robj)
					}
					if err != nil {
						log.Printf("[ERROR] ApiCreateOrUpdate: PUT Error %v path %v id %v\n", err, path, d.Id())
					}
				}
			} else {
				log.Printf("[INFO] ApiCreateOrUpdate: Creating obj %v schema %v data %v\n", objType, d, data)
				err = client.AviSession.Post(path, data, &robj)
				if err != nil {
					log.Printf("[ERROR] ApiCreateOrUpdate creation failed %v\n", err)
				} else {
					SetIDFromObj(d, robj)
				}
			}
		}
		return err
	} else {
		log.Printf("[ERROR] ApiCreateOrUpdate: Error %v", err)
		return err
	}
	return nil
}

func ApiRead(d *schema.ResourceData, meta interface{}, objType string, s map[string]*schema.Schema) error {
	client := meta.(*clients.AviClient)
	var obj interface{}
	var path string
	uuid := ""
	url := ""
	specialobj := IsPostNotAllowed(objType)
	log.Printf("[DEBUG] ApiRead reading object with objType %v id %v\n", objType, d.Id())
	if d.Id() != "" {
		// extract the uuid from it.
		log.Printf("[DEBUG] ApiRead reading object with objType %v id %v \n", objType, d.Id())
		url_parts := strings.Split(d.Id(), "/")
		uuid = url_parts[len(url_parts)-1]
		// need to strip #xxx if present
		uuid = strings.Split(uuid, "#")[0]
	} else if u, ok := d.GetOk("uuid"); ok {
		uuid = u.(string)
		log.Printf("[DEBUG] ApiRead reading object with uuid %v \n", uuid)
	}
	if uuid != "" {
		if specialobj {
			path = "api/" + objType
		} else {
			path = "api/" + objType + "/" + uuid + "?skip_default=true"
		}
		log.Printf("[DEBUG] ApiRead reading object with id %v path %v\n", uuid, path)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			log.Printf("[ERROR] ApiRead object with uuid %v not found err %v\n", uuid, err)
			return nil
		}
	} else if name, ok := d.GetOk("name"); ok {
		var err error
		if cloudRef, ok := d.GetOk("cloud_ref"); ok && strings.Contains(cloudRef.(string), "api/cloud/") {
			cloudUUID := strings.SplitN(cloudRef.(string), "api/cloud/", 2)[1]
			cloudUUID = strings.Split(cloudUUID, "#")[0]
			log.Printf("[DEBUG] ApiRead using cloud %v obj %v name %v\n", cloudUUID,
				objType, name)
			err = client.AviSession.GetObject(objType, session.SetName(name.(string)),
				session.SetResult(&obj), session.SetCloudUUID(cloudUUID), session.SetSkipDefault(true))
		} else {
			log.Printf("[DEBUG] ApiRead using name %v \n", name)
			err = client.AviSession.GetObject(objType, session.SetName(name.(string)),
				session.SetResult(&obj), session.SetSkipDefault(true))
		}
		if err != nil {
			d.SetId("")
			log.Printf("[ERROR] ApiRead object with name %v:%v not found err %v\n", objType, name, err)
			return nil
		}
	} else if specialobj {
		path := "api/" + objType
		log.Printf("[DEBUG] ApiRead reading special object with path %v\n", path)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			log.Printf("[ERROR] ApiRead special object with path %v not found err %v\n", path, err)
			return nil
		}
	} else {
		d.SetId("")
		log.Printf("[ERROR] ApiRead not found %v\n", d.Get("uuid"))
		return nil
	}
	if objType == "cluster" {
		update_cluster_state(d, meta, s)
	}
	if local_data, err := SchemaToAviData(d, s); err == nil {
		mod_api_res, err := SetDefaultsInAPIRes(obj, local_data, s)
		if err != nil {
			log.Printf("[ERROR] ApiRead in modifying api response object %v\n", err)
		}
		if _, err := ApiDataToSchema(mod_api_res, d, s); err == nil {
			if mod_api_res.(map[string]interface{})["uuid"] != nil {
				uuid = mod_api_res.(map[string]interface{})["uuid"].(string)
			}
			if mod_api_res.(map[string]interface{})["url"] != nil {
				url = mod_api_res.(map[string]interface{})["url"].(string)
			}
			//url = strings.SplitN(url, "#", 2)[0]
			if url != "" {
				d.SetId(url)
				log.Printf("[DEBUG] ApiRead read object with id %v\n", url)
			} else {
				d.SetId(uuid)
				log.Printf("[DEBUG] ApiRead read object with id %v\n", uuid)
			}
		} else {
			log.Printf("[ERROR] ApiRead in setting read object %v\n", err)
		}
		log.Printf("[DEBUG] type: %v local_data : %v", objType, local_data)
		log.Printf("[DEBUG] type: %v mod_api_res: %v", objType, mod_api_res)
	}

	return nil
}

func ResourceImporter(d *schema.ResourceData, meta interface{}, objType string, s map[string]*schema.Schema) ([]*schema.ResourceData, error) {
	log.Printf("[DEBUG] ResourceImporter obuType%v id %v\n", objType, d.Id())
	if d.Id() != "" {
		// return the ID based import
		return []*schema.ResourceData{d}, nil
	}
	var data interface{}
	client := meta.(*clients.AviClient)
	path := "api/" + objType + "?skip_default=true"
	err := client.AviSession.Get(path, &data)
	if err != nil {
		log.Printf("[ERROR] ResourceImporter %v in GET of path %v\n", err, path)
		return nil, err
	}
	count := int((data.(map[string]interface{})["count"]).(float64))
	log.Printf("[DEBUG] ResourceImporter read data with path %v -> count %v\n", path, count)
	if count > 0 {
		results := make([]*schema.ResourceData, count)
		apiResults := (data.(map[string]interface{})["results"]).([]interface{})
		for index := 0; index < count; index++ {
			obj := apiResults[index].(map[string]interface{})
			log.Printf("[DEBUG] ResourceImporter processing obj %v results %v\n", obj, results[index])
			result := new(schema.ResourceData)
			if _, err := ApiDataToSchema(obj, result, s); err == nil {
				log.Printf("[DEBUG] ResourceImporter Processing obj %v\n", obj)
				url := obj["url"].(string)
				uuid := obj["uuid"].(string)
				//url = strings.SplitN(url, "#", 2)[0]
				result.SetId(url)
				result.Set("uuid", uuid)
				result.SetType("avi_" + objType)
				results[index] = result
			}
		}
		return results, nil
	}
	return nil, nil
}

func ApiDeleteSystemDefaultCheck(d *schema.ResourceData) bool {
	var systemDefault bool
	var sysName string
	if sysdef, ok := d.GetOk("system_default"); ok {
		systemDefault = sysdef.(bool)
	}
	if name, ok := d.GetOk("name"); ok {
		sysName = name.(string)
	}
	if systemDefault || strings.HasPrefix(sysName, "System-") || sysName == "Default-Group" {
		d.SetId("")
		return true
	}
	return false
}

//Open given file as a file pointer
func mustOpen(f string) *os.File {
	r, err := os.Open(f)
	if err != nil {
		log.Printf("[ERROR] mustOpen Error while opening file %v", f)
		panic(err)
	}
	return r
}

func createFilePointer(path string) (*os.File, error) {
	// detect if file exists
	var _, err = os.Stat(path)
	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return nil, err
		}
		log.Printf("[INFO] createFilePointer File created %v", path)
		return file, err
	} else {
		// open file using READ & WRITE permission
		var file, err = os.OpenFile(path, os.O_RDWR, 0644)
		log.Printf("[INFO] createFilePointer File exist Reopening %v", path)
		return file, err
	}
}

//Function to make REST API call for upload and download.
func MultipartUploadOrDownload(d *schema.ResourceData, meta interface{}, s map[string]*schema.Schema) error {
	client := meta.(*clients.AviClient)
	uri := d.Get("uri").(string)
	local_file := d.Get("local_file").(string)
	var err error
	var res interface{}
	var license_file_name string
	switch upload := d.Get("upload").(bool); upload {
	case true:
		//Checking for license URI
		switch uri := d.Get("uri").(string); uri {
		case "license":
			data, err := ioutil.ReadFile(local_file)
			if err != nil {
				log.Panicf("failed reading data from file: %s", err)
			}
			str_data := string(data)
			license_data := map[string]string{
				"license_text": str_data,
			}
			uri = "/api/" + uri
			err = client.AviSession.Put(uri, license_data, &res)
			if err != nil {
				log.Printf("[ERROR] MultipartUploadOrDownload %v in PUT of URI %v\n", err, uri)
				return err
			}
			if res != nil {
				//RegX to fetch license name from controller response.
				re := regexp.MustCompile(`(?mi)((?:license.*[^"}]))`)
				result := res.(map[string]interface{})["result"].(string)
				matched_str := re.FindString(result)
				splited_str := strings.Split(matched_str, " ")
				//Iterating over string array if license name contains white spaces. Concatinating strings to get final license name
				for _, val := range splited_str[1:] {
					license_file_name = license_file_name + val
				}
				//Setting license name as an ID.
				d.SetId(license_file_name)
			}
		default:
			if _, err := os.Stat(local_file); os.IsNotExist(err) {
				log.Printf("[ERROR] MultipartUploadOrDownload File path does not exist %v", local_file)
				return err
			}
			local_file_ptr := mustOpen(local_file)
			err := client.AviSession.PostMultipartRequest("POST", uri, local_file_ptr)
			if err != nil {
				log.Printf("[ERROR] MultipartUploadOrDownload Error uploading file %v %v", local_file, err)
				return err
			}
			_, file := filepath.Split(local_file)
			d.SetId(file)
		}
	case false:
		//For multipart file download
		//File creation
		download_file_ptr, err := createFilePointer(local_file)
		if err != nil {
			log.Printf("[ERROR] MultipartUploadOrDownload Error for creation of file %v", local_file)
		}
		err = client.AviSession.GetMultipartRaw("GET", uri, download_file_ptr)
		if err != nil {
			log.Printf("[ERROR] MultipartUploadOrDownload Error downloaing file using uri %v %v", uri, err)
			return err
		}
		_, file := filepath.Split(local_file)
		d.SetId(file)
	default:
		//For multipart file download
		//File creation
		download_file_ptr, err := createFilePointer(local_file)
		if err != nil {
			log.Printf("[ERROR] MultipartUploadOrDownload Error for creation of file %v", local_file)
		}
		err = client.AviSession.GetMultipartRaw("GET", uri, download_file_ptr)
		if err != nil {
			log.Printf("[ERROR] MultipartUploadOrDownload Error downloaing file using uri %v %v", uri, err)
			return err
		}
		_, file := filepath.Split(local_file)
		d.SetId(file)
	}
	return err
}

func UUIDFromID(Id string) string {
	urlParts := strings.Split(Id, "/")
	idParts := urlParts[len(urlParts)-1]
	// need to strip #xxx if present
	nu := strings.Split(idParts, "#")
	return nu[0]
}

func IsPostNotAllowed(objtype string) bool {
	specialobj := false
	for _, o_type := range post_not_allowed {
		if o_type == objtype {
			log.Printf("[INFO] ApiRead: Found special object type %v", objtype)
			specialobj = true
			break
		}
	}
	return specialobj
}

func update_cluster_state(d *schema.ResourceData, meta interface{}, s map[string]*schema.Schema) error {
	client := meta.(*clients.AviClient)
	var err error
	var robj interface{}
	err = client.AviSession.Get("api/cluster/runtime", &robj)
	if err == nil {
		if local_data, err := SchemaToAviData(d, s); err == nil {
			mod_api_res, err := SetDefaultsInAPIRes(robj, local_data, s)
			if err != nil {
				log.Printf("[ERROR] Update Cluster State in modifying api response object %v\n", err)
			}
			if _, err := ApiDataToSchema(mod_api_res, d, s); err != nil {
				log.Printf("[ERROR] Converting ApiDataToSchema object %v\n", err)
			}
		}
	}
	return err
}
