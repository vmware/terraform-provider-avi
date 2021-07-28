/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
	"github.com/vmware/alb-sdk/go/session"
)

var postNotAllowed = [...]string{"systemconfiguration", "cluster", "seproperties", "controllerproperties"}

// It takes the terraform plan data and schema and converts it into Avi JSON
// It recursively resolves the data type of the terraform schema and converts scalar to scalar, Set to dictionary,
// and list to list.
func SchemaToAviData(d interface{}, s interface{}) (interface{}, error) {
	switch dType := d.(type) {
	default:
		// Convert schema data to expected data type for API payload
		if validateFunctionName := getValidateFunction(s.(*schema.Schema)); validateFunctionName != "" && d != "" {
			var err error
			switch validateFunctionName {
			case "validateInteger":
				if d, err = strconv.ParseInt(d.(string), 10, 64); err != nil {
					log.Printf("[ERROR] SchemaToAviData in converting string %v to integer. Error: %v", d, err)
				}
			case "validateBool":
				if d, err = strconv.ParseBool(d.(string)); err != nil {
					log.Printf("[ERROR] SchemaToAviData in converting string %v to bool. Error: %v", d, err)
				}
			case "validateFloat":
				if d, err = strconv.ParseFloat(d.(string), 64); err != nil {
					log.Printf("[ERROR] SchemaToAviData in converting string %v to float. Error: %v", d, err)
				}
			}
		}
	case map[string]interface{}:
		m := make(map[string]interface{})
		for k, v := range d.(map[string]interface{}) {
			if obj, err := SchemaToAviData(v, s.(map[string]*schema.Schema)[k]); err == nil && obj != nil && obj != "" {
				m[k] = obj
			} else if err != nil {
				log.Printf("[ERROR] SchemaToAviData %v in parsing k: %v v: %v type: %v", err, k, v, dType)
			}
		}
		return m, nil
	case []interface{}:
		var objList []interface{}
		varray := d.([]interface{})
		var listSchema interface{}
		switch sType := s.(*schema.Schema).Elem.(type) {
		default:
			log.Printf("%v", sType)
		case *schema.Resource:
			listSchema = s.(*schema.Schema).Elem.(*schema.Resource).Schema
		case *schema.Schema:
			listSchema = s.(*schema.Schema).Elem.(*schema.Schema)
		}
		for i := 0; i < len(varray); i++ {
			obj, err := SchemaToAviData(varray[i], listSchema)
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
		obj, err := SchemaToAviData(d.(*schema.Set).List()[0], s.(*schema.Schema).Elem.(*schema.Resource).Schema)
		return obj, err

	case *schema.ResourceData:
		// In this case the top level schema should be present.
		m := make(map[string]interface{})
		r := d.(*schema.ResourceData)
		for k, v := range s.(map[string]*schema.Schema) {
			if data, ok := r.GetOkExists(k); ok {
				if obj, err := SchemaToAviData(data, v); err == nil && obj != nil && obj != "" {
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
	return schema.HashString("avi")
}

// It sets default values in the terraform resources to avoid diffs for scalars.
func SetDefaultsInAPIRes(apiRes interface{}, dLocal interface{}, s map[string]*schema.Schema) (interface{}, error) {
	if apiRes == nil {
		log.Printf("[ERROR] SetDefaultsInAPIRes got nil for %v", s)
		return apiRes, nil
	}
	switch dLocal.(type) {
	default:
	case map[string]interface{}:
		for k, v := range dLocal.(map[string]interface{}) {
			switch v.(type) {
			//Getting key, value for given dLocal
			default:
				if _, ok := apiRes.(map[string]interface{})[k]; !ok {
					//Cheking if field is present in schema
					if dval, ok := s[k]; ok {
						//Getting default values from schema
						defaultVal, err := dval.DefaultValue()
						if err != nil {
							log.Printf("[ERROR] SetDefaultsInAPIRes did not get default value from schema err %v %v",
								err, defaultVal)
						} else {
							if defaultVal != nil {
								apiRes.(map[string]interface{})[k] = defaultVal
							}
						}
					}
				}
			//dLocal nested dictionary.
			case map[string]interface{}:
				if s2, ok := s[k]; ok {
					switch s2.Elem.(type) {
					default:
					case *schema.Resource:
						if apiRes.(map[string]interface{})[k] != nil {
							apiRes1, err := SetDefaultsInAPIRes(apiRes.(map[string]interface{})[k], v, s2.Elem.(*schema.Resource).Schema)
							if err != nil {
								log.Printf("[ERROR] SetDefaultsInAPIRes %v", err)
							} else {
								apiRes.(map[string]interface{})[k] = apiRes1
							}
						} else {
							apiRes.(map[string]interface{})[k] = v
						}
					}
				}
			//dLocal is array of dictionaries.
			case []interface{}:
				var objList []interface{}
				if apiRes.(map[string]interface{})[k] != nil {
					varray2 := apiRes.(map[string]interface{})[k].([]interface{})
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
				apiRes.(map[string]interface{})[k] = objList
			}
		}
	}
	return apiRes, nil
}

// It takes the Avi JSON data and fills in the terraform data during API read.
// It takes input as the top level schema and it uses that to properly create the corresponding terraform resource data
// It also checks whether a given Avi key is defined in the schema before attempting to fill the data.
func APIDataToSchema(adata interface{}, d interface{}, t map[string]*schema.Schema) (interface{}, error) {
	switch adata.(type) {
	default:
	case map[string]interface{}:
		// resolve d interface into a set
		if t == nil {
			m := map[string]interface{}{}
			for k, v := range adata.(map[string]interface{}) {
				if obj, err := APIDataToSchema(v, nil, nil); err == nil {
					m[k] = obj
				} else if err != nil {
					log.Printf("[ERROR] APIDataToSchema %v in converting k: %v v: %v", err, k, v)
				}
			}
			objs := []interface{}{}
			objs = append(objs, m)
			s := schema.NewSet(CommonHash, objs)
			return s, nil
		} else { //nolint
			for k, v := range adata.(map[string]interface{}) {
				if _, ok := t[k]; ok {
					// found in the schema
					if obj, err := APIDataToSchema(v, nil, nil); err == nil {
						switch dType := d.(type) {
						default:
						case *schema.ResourceData:
							if err := d.(*schema.ResourceData).Set(k, obj); err != nil {
								log.Printf("[ERROR] APIDataToSchema %v in setting %v   type %v", err, obj, dType)
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
			obj, err := APIDataToSchema(varray[i], nil, nil)
			if err == nil {
				switch objType := obj.(type) {
				default:
					log.Printf("Appending objtype %v to list %v", objType, objList)
					objList = append(objList, obj)
				case *schema.Set:
					objList = append(objList, obj.(*schema.Set).List()[0])
				}
			} else {
				log.Printf("[ERROR] APIDataToSchema %v", err)
			}
		}
		return objList, nil
		/** Return the same object as there is nothing special about **/
	}
	return adata, nil
}

func SetIDFromObj(d *schema.ResourceData, robj interface{}) {
	uuid := robj.(map[string]interface{})["uuid"].(string)
	_ = d.Set("uuid", uuid)
	if url, ok := robj.(map[string]interface{})["url"].(string); ok && url != "" {
		d.SetId(url)
	} else {
		d.SetId(uuid)
	}
}

// It is generic API to create and update any Avi REST resource. It handles special situations with cloud
// and tenant filters as objects may already be present. If the resource does not exist it will try to
// create it. In case, it is present then automatically converts to PUT semantics.
func APICreateOrUpdate(d *schema.ResourceData, meta interface{}, objType string, s map[string]*schema.Schema,
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
				log.Printf("[ERROR] APICreateOrUpdate: PUT on %v Error %v path %v id %v\n", objType, err, path,
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
				log.Printf("[ERROR] APICreateOrUpdate: PUT Error %v path %v id %v\n", err, path, d.Id())
			}
		} else {
			if name, ok := d.GetOk("name"); ok {
				var existingObj interface{}
				if cloudRef, ok := d.GetOk("cloud_ref"); ok && strings.Contains(cloudRef.(string),
					"api/cloud/") {
					cloudUUID := strings.SplitN(cloudRef.(string), "api/cloud/", 2)[1]
					// strip the # if it exists
					cloudUUID = strings.Split(cloudUUID, "#")[0]
					log.Printf("[INFO] APICreateOrUpdate: using cloud %v for obj %v name %s \n",
						cloudUUID, objType, name)
					err = client.AviSession.GetObject(objType, session.SetName(name.(string)),
						session.SetResult(&existingObj), session.SetCloudUUID(cloudUUID),
						session.SetSkipDefault(true))
					if err != nil {
						log.Printf("[ERROR] APICreateOrUpdate: GET Error %v path %v id %v\n", err, path, d.Id())
					}
				} else {
					log.Printf("[INFO] APICreateOrUpdate: reading obj %v name %s \n",
						objType, name)
					err = client.AviSession.GetObject(objType, session.SetName(name.(string)),
						session.SetResult(&existingObj), session.SetSkipDefault(true))
					if err != nil {
						log.Printf("[ERROR] APICreateOrUpdate: GET Error %v path %v id %v\n", err, path, d.Id())
					}
				}

				if existingObj == nil {
					// object not found
					log.Printf("[INFO] APICreateOrUpdate: Creating obj type %v schema %v data %v\n", objType, d,
						data)
					err = client.AviSession.Post(path, data, &robj)
					if err == nil && robj != nil {
						SetIDFromObj(d, robj)
					} else {
						log.Printf("[ERROR] APICreateOrUpdate creation failed %v object with name %v\n", err,
							name)
					}
				} else {
					// found existing object.
					SetIDFromObj(d, existingObj)
					uuid = existingObj.(map[string]interface{})["uuid"].(string)
					path = path + "/" + uuid.(string) + "?skip_default=true"
					if !usePatchForUpdate {
						err = client.AviSession.Put(path, data, &robj)
					} else {
						err = client.AviSession.Patch(path, data, "replace", &robj)
					}
					if err != nil {
						log.Printf("[ERROR] APICreateOrUpdate: PUT Error %v path %v id %v\n", err, path, d.Id())
					}
				}
			} else {
				log.Printf("[INFO] APICreateOrUpdate: Creating obj %v schema %v data %v\n", objType, d, data)
				err = client.AviSession.Post(path, data, &robj)
				if err != nil {
					log.Printf("[ERROR] APICreateOrUpdate creation failed %v\n", err)
				} else {
					SetIDFromObj(d, robj)
				}
			}
		}
		return err
	} else { //nolint
		log.Printf("[ERROR] APICreateOrUpdate: Error %v", err)
		return err
	}
}

func APIRead(d *schema.ResourceData, meta interface{}, objType string, s map[string]*schema.Schema) error {
	client := meta.(*clients.AviClient)
	var obj interface{}
	var path string
	uuid := ""
	url := ""
	specialobj := IsPostNotAllowed(objType)
	log.Printf("[DEBUG] APIRead reading object with objType %v id %v\n", objType, d.Id())
	if d.Id() != "" {
		// extract the uuid from it.
		log.Printf("[DEBUG] APIRead reading object with objType %v id %v \n", objType, d.Id())
		urlParts := strings.Split(d.Id(), "/")
		uuid = urlParts[len(urlParts)-1]
		// need to strip #xxx if present
		uuid = strings.Split(uuid, "#")[0]
	} else if u, ok := d.GetOk("uuid"); ok {
		uuid = u.(string)
		log.Printf("[DEBUG] APIRead reading object with uuid %v \n", uuid)
	}
	if uuid != "" {
		if specialobj {
			path = "api/" + objType
		} else {
			path = "api/" + objType + "/" + uuid + "?skip_default=true"
		}
		log.Printf("[DEBUG] APIRead reading object with id %v path %v\n", uuid, path)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			log.Printf("[ERROR] APIRead object with uuid %v not found err %v\n", uuid, err)
			return nil
		}
	} else if name, ok := d.GetOk("name"); ok {
		var err error
		if cloudRef, ok := d.GetOk("cloud_ref"); ok && strings.Contains(cloudRef.(string), "api/cloud/") {
			cloudUUID := strings.SplitN(cloudRef.(string), "api/cloud/", 2)[1]
			cloudUUID = strings.Split(cloudUUID, "#")[0]
			log.Printf("[DEBUG] APIRead using cloud %v obj %v name %v\n", cloudUUID,
				objType, name)
			err = client.AviSession.GetObject(objType, session.SetName(name.(string)),
				session.SetResult(&obj), session.SetCloudUUID(cloudUUID), session.SetSkipDefault(true))
		} else {
			log.Printf("[DEBUG] APIRead using name %v \n", name)
			err = client.AviSession.GetObject(objType, session.SetName(name.(string)),
				session.SetResult(&obj), session.SetSkipDefault(true))
		}
		if err != nil {
			d.SetId("")
			log.Printf("[ERROR] APIRead object with name %v:%v not found err %v\n", objType, name, err)
			return nil
		}
	} else if specialobj {
		path := "api/" + objType
		log.Printf("[DEBUG] APIRead reading special object with path %v\n", path)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			log.Printf("[ERROR] APIRead special object with path %v not found err %v\n", path, err)
			return nil
		}
	} else {
		d.SetId("")
		log.Printf("[ERROR] APIRead not found %v\n", d.Get("uuid"))
		return nil
	}
	if localData, err := SchemaToAviData(d, s); err == nil {
		modAPIRes, err := SetDefaultsInAPIRes(obj, localData, s)
		if err != nil {
			log.Printf("[ERROR] APIRead in modifying api response object %v\n", err)
		}
		modAPIRes, err = PreprocessAPIRes(obj, s)
		if err != nil {
			log.Printf("[ERROR] APIRead in modifying api response object for conversion %v\n", err)
		}
		if _, err := APIDataToSchema(modAPIRes, d, s); err == nil {
			if modAPIRes.(map[string]interface{})["uuid"] != nil {
				uuid = modAPIRes.(map[string]interface{})["uuid"].(string)
			}
			if modAPIRes.(map[string]interface{})["url"] != nil {
				url = modAPIRes.(map[string]interface{})["url"].(string)
			}
			//url = strings.SplitN(url, "#", 2)[0]
			if url != "" {
				d.SetId(url)
				log.Printf("[DEBUG] APIRead read object with id %v\n", url)
			} else {
				d.SetId(uuid)
				log.Printf("[DEBUG] APIRead read object with id %v\n", uuid)
			}
		} else {
			log.Printf("[ERROR] APIRead in setting read object %v\n", err)
		}
		log.Printf("[DEBUG] type: %v localData : %v", objType, localData)
		log.Printf("[DEBUG] type: %v modAPIRes: %v", objType, modAPIRes)
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
			if _, err := APIDataToSchema(obj, result, s); err == nil {
				log.Printf("[DEBUG] ResourceImporter Processing obj %v\n", obj)
				url := obj["url"].(string)
				uuid := obj["uuid"].(string)
				//url = strings.SplitN(url, "#", 2)[0]
				result.SetId(url)
				_ = result.Set("uuid", uuid)
				result.SetType("avi_" + objType)
				results[index] = result
			}
		}
		return results, nil
	}
	return nil, nil
}

func APIDeleteSystemDefaultCheck(d *schema.ResourceData) bool {
	var systemDefault bool
	var sysName string
	var err error
	if sysdef, ok := d.GetOk("system_default"); ok {
		if systemDefault, err = strconv.ParseBool(sysdef.(string)); err != nil {
			log.Printf("[ERROR] Error %v while converting system_default value %s to bool.", err, sysdef)
		}
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
	_, err := os.Stat(path)
	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return nil, err
		}
		log.Printf("[INFO] createFilePointer File created %v", path)
		return file, err
	} else { //nolint
		// open file using READ & WRITE permission
		file, err := os.OpenFile(path, os.O_RDWR, 0644)
		log.Printf("[INFO] createFilePointer File exist Reopening %v", path)
		return file, err
	}
}

//Function to make REST API call for upload and download.
func MultipartUploadOrDownload(d *schema.ResourceData, meta interface{}, s map[string]*schema.Schema) error {
	client := meta.(*clients.AviClient)
	uri := d.Get("uri").(string)
	localFile := d.Get("local_file").(string)
	var err error
	var res interface{}
	var licenseFileName string
	switch upload := d.Get("upload").(bool); upload {
	case true:
		//Checking for license URI
		switch uri := d.Get("uri").(string); uri {
		case "license":
			data, err := ioutil.ReadFile(localFile)
			if err != nil {
				log.Panicf("failed reading data from file: %s", err)
			}
			strData := string(data)
			licenseData := map[string]string{
				"license_text": strData,
			}
			uri = "/api/" + uri
			err = client.AviSession.Put(uri, licenseData, &res)
			if err != nil {
				log.Printf("[ERROR] MultipartUploadOrDownload %v in PUT of URI %v\n", err, uri)
				return err
			}
			if res != nil {
				//RegX to fetch license name from controller response.
				re := regexp.MustCompile(`(?mi)((?:license.*[^"}]))`)
				result := res.(map[string]interface{})["result"].(string)
				matchedStr := re.FindString(result)
				splitedStr := strings.Split(matchedStr, " ")
				//Iterating over string array if license name contains white spaces. Concatinating strings to get final license name
				for _, val := range splitedStr[1:] {
					licenseFileName = licenseFileName + val
				}
				//Setting license name as an ID.
				d.SetId(licenseFileName)
			}
		default:
			if _, err := os.Stat(localFile); os.IsNotExist(err) {
				log.Printf("[ERROR] MultipartUploadOrDownload File path does not exist %v", localFile)
				return err
			}
			localFilePtr := mustOpen(localFile)
			err := client.AviSession.PostMultipartRequest("POST", uri, localFilePtr)
			if err != nil {
				log.Printf("[ERROR] MultipartUploadOrDownload Error uploading file %v %v", localFile, err)
				return err
			}
			_, file := filepath.Split(localFile)
			d.SetId(file)
		}
	case false:
		//For multipart file download
		//File creation
		downloadFilePtr, err := createFilePointer(localFile)
		if err != nil {
			log.Printf("[ERROR] MultipartUploadOrDownload Error for creation of file %v", localFile)
		}
		err = client.AviSession.GetMultipartRaw("GET", uri, downloadFilePtr)
		if err != nil {
			log.Printf("[ERROR] MultipartUploadOrDownload Error downloaing file using uri %v %v", uri, err)
			return err
		}
		_, file := filepath.Split(localFile)
		d.SetId(file)
	default:
		//For multipart file download
		//File creation
		downloadFilePtr, err := createFilePointer(localFile)
		if err != nil {
			log.Printf("[ERROR] MultipartUploadOrDownload Error for creation of file %v", localFile)
		}
		err = client.AviSession.GetMultipartRaw("GET", uri, downloadFilePtr)
		if err != nil {
			log.Printf("[ERROR] MultipartUploadOrDownload Error downloaing file using uri %v %v", uri, err)
			return err
		}
		_, file := filepath.Split(localFile)
		d.SetId(file)
	}
	return err
}

func UUIDFromID(ID string) string {
	urlParts := strings.Split(ID, "/")
	idParts := urlParts[len(urlParts)-1]
	// need to strip #xxx if present
	nu := strings.Split(idParts, "#")
	return nu[0]
}

func IsPostNotAllowed(objtype string) bool {
	specialobj := false
	for _, oType := range postNotAllowed {
		if oType == objtype {
			log.Printf("[INFO] APIRead: Found special object type %v", objtype)
			specialobj = true
			break
		}
	}
	return specialobj
}

// Preprocess API response to convert int/bool/float data from respose to string
func PreprocessAPIRes(apiRes interface{}, s map[string]*schema.Schema) (interface{}, error) {
	if apiRes == nil {
		log.Printf("[ERROR] PreprocessAPIRes got nil for %v", s)
		return apiRes, nil
	}
	switch apiRes.(type) {
	default:
	case map[string]interface{}:
		for k, v := range apiRes.(map[string]interface{}) {
			switch v.(type) {
			//Getting key, value for response apiRes
			default:
				if _, ok := apiRes.(map[string]interface{})[k]; ok {
					if dval, ok := s[k]; ok {
						validateFunctionName := getValidateFunction(dval)
						if validateFunctionName != "" {
							switch validateFunctionName {
							default:
							case "validateInteger":
								if reflect.TypeOf(apiRes.(map[string]interface{})[k]).Kind() == reflect.Float64 {
									apiRes.(map[string]interface{})[k] = strconv.FormatFloat(apiRes.(map[string]interface{})[k].(float64), 'f', -1, 64)
								} else if reflect.TypeOf(apiRes.(map[string]interface{})[k]).Kind() == reflect.Int64 {
									apiRes.(map[string]interface{})[k] = strconv.FormatInt(apiRes.(map[string]interface{})[k].(int64), 10)
								}
							case "validateBool":
								if reflect.TypeOf(apiRes.(map[string]interface{})[k]).Kind() == reflect.Bool {
									apiRes.(map[string]interface{})[k] = strconv.FormatBool(apiRes.(map[string]interface{})[k].(bool))
								}
							case "validateFloat":
								if reflect.TypeOf(apiRes.(map[string]interface{})[k]).Kind() == reflect.Float64 {
									apiRes.(map[string]interface{})[k] = strconv.FormatFloat(apiRes.(map[string]interface{})[k].(float64), 'f', -1, 64)
								} else if reflect.TypeOf(apiRes.(map[string]interface{})[k]).Kind() == reflect.String {
									// Remove trailing zeroes from response
									if strFloat, err := strconv.ParseFloat(apiRes.(map[string]interface{})[k].(string), 64); err == nil {
										apiRes.(map[string]interface{})[k] = strconv.FormatFloat(strFloat, 'f', -1, 64)
									} else {
										log.Printf("[ERROR] PreprocessAPIRes in converting string %v to float. Error: %v", apiRes.(map[string]interface{})[k], err)
									}
								}
							}
						}
					}
				}
			//apiRes nested dictionary.
			case map[string]interface{}:
				if s2, ok := s[k]; ok {
					switch s2.Elem.(type) {
					default:
					case *schema.Resource:
						if apiRes.(map[string]interface{})[k] != nil {
							apiRes1, err := PreprocessAPIRes(v, s2.Elem.(*schema.Resource).Schema)
							if err != nil {
								log.Printf("[ERROR] PreprocessAPIRes %v", err)
							} else {
								apiRes.(map[string]interface{})[k] = apiRes1
							}
						} else {
							apiRes.(map[string]interface{})[k] = v
						}
					}
				}
			//apiRes is array of dictionaries.
			case []interface{}:
				var objList []interface{}
				if apiRes.(map[string]interface{})[k] != nil {
					//getting schema for nested object.
					s2, err := s[k]
					//As err returned is boolean value
					if !err {
						log.Printf("[ERROR] PreprocessAPIRes in fetching k %v err %v", k, err)
						continue
					}
					for x, y := range v.([]interface{}) {
						switch s2.Elem.(type) {
						default:
						case *schema.Resource:
							obj, err := PreprocessAPIRes(y, s2.Elem.(*schema.Resource).Schema)
							if err != nil {
								log.Printf("[ERROR] PreprocessAPIRes err %v in x %v y %v", err, x, y)
							} else {
								objList = append(objList, obj)
							}
						case *schema.Schema:
							objList = append(objList, v.([]interface{})[x])
						}
					}
				}
				apiRes.(map[string]interface{})[k] = objList
			}
		}
	}
	return apiRes, nil
}

func getValidateFunction(schemaVal *schema.Schema) string {
	if isValidateFunction := schemaVal.ValidateFunc; isValidateFunction != nil {
		validateFunctionRef := runtime.FuncForPC(reflect.ValueOf(schemaVal.ValidateFunc).Pointer()).Name()
		validateFunctionSplit := strings.Split(validateFunctionRef, ".")
		validateFunctionName := validateFunctionSplit[len(validateFunctionSplit)-1]
		return validateFunctionName
	}
	return ""
}
