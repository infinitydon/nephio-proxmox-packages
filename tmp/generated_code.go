
package controller

import (
	"context"
	"fmt"
	"time"
	"encoding/base64"
	
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	schedulingv1 "k8s.io/api/scheduling/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/utils/ptr"
)

func deleteMeAfterDeletingUnusedImportedModules() {
	/*
		It is written to handle the error "Module Imported but not used",
		The user can delete the non-required modules from import and then delete this function also
	*/
	_ = time.Now()
	_ = &unstructured.Unstructured{}
	_ = corev1.Service{}
	_ = metav1.ObjectMeta{}
	_ = appsv1.Deployment{}
	_ = rbacv1.Role{}
	_ = schedulingv1.PriorityClass{}
	_ = intstr.FromInt(4)
	_, _ = resource.ParseQuantity("")
	_ = context.TODO()
	_ = fmt.Sprintf("")
	_ = ptr.To(32)
}

func int32Ptr(val int) *int32 {
	var a int32
	a = int32(val)
	return &a
}

func int64Ptr(val int) *int64 {
	var a int64
	a = int64(val)
	return &a
}

func intPtr(val int) *int {
	a := val
	return &a
}

func int16Ptr(val int) *int16 {
	var a int16
	a = int16(val)
	return &a
}

func boolPtr(val bool) *bool {
	a := val
	return &a
}

func stringPtr(val string) *string {
	a := val
	return &a
}

func getDataForSecret(encodedVal string) []byte {
	/*
		Concept: Based on my Understanding, corev1.Secret requires the actual data(not encoded) as secret-Data
		But in general terms, we put encoded values in secret-data, which make sense (why to write actual value in readable format)
		This function takes the encodedVal and decodes it and returns
	*/
	decodeVal, err := base64.StdEncoding.DecodeString(encodedVal)
	if err != nil {
		fmt.Println("Unable to decode the SecretVal ", encodedVal, " || This Secret Will Probably would give error during deployment| Kindly Check")
		return []byte(encodedVal)
	}
	return decodeVal
}


/*
// Before Uncommenting the following function, Make sure the data-type of r is same as of your Reconciler,
// Replace "YourKindReconciler" with the type of your Reconciler
func (r *YourKindReconciler)CreateAll(){
 	var err error
	namespaceProvided := "open5gscns"
	
	for _, resource := range GetConfigMap(){
		if resource.ObjectMeta.Namespace == ""{
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Create(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Creating resource of GetConfigMap()| Error --> |", err)
		}
	} 
			
	for _, resource := range GetService(){
		if resource.ObjectMeta.Namespace == ""{
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Create(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Creating resource of GetService()| Error --> |", err)
		}
	} 
			
	for _, resource := range GetDeployment(){
		if resource.ObjectMeta.Namespace == ""{
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Create(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Creating resource of GetDeployment()| Error --> |", err)
		}
	} 
			
	for _, resource := range GetRoleBinding(){
		if resource.ObjectMeta.Namespace == ""{
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Create(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Creating resource of GetRoleBinding()| Error --> |", err)
		}
	} 
			
	for _, resource := range GetRole(){
		if resource.ObjectMeta.Namespace == ""{
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Create(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Creating resource of GetRole()| Error --> |", err)
		}
	} 
			
	for _, resource := range GetServiceAccount(){
		if resource.ObjectMeta.Namespace == ""{
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Create(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Creating resource of GetServiceAccount()| Error --> |", err)
		}
	} 
			
	for _, resource := range GetNetworkAttachmentDefinition(){
		if resource.GetNamespace() == ""{
			resource.SetNamespace(namespaceProvided)
		}
		err = r.Create(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Creating resource of GetNetworkAttachmentDefinition()| Error --> |", err)
		}
	} 
			
}
*/

	
/*
// Before Uncommenting the following function, Make sure the data-type of r is same as of your Reconciler,
// Replace "YourKindReconciler" with the type of your Reconciler
func (r *YourKindReconciler)DeleteAll(){
 	var err error
	namespaceProvided := "open5gscns"
	
	for _, resource := range GetConfigMap(){
		if resource.ObjectMeta.Namespace == ""{
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetConfigMap()| Error --> |", err)
		}
	} 
			
	for _, resource := range GetService(){
		if resource.ObjectMeta.Namespace == ""{
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetService()| Error --> |", err)
		}
	} 
			
	for _, resource := range GetDeployment(){
		if resource.ObjectMeta.Namespace == ""{
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetDeployment()| Error --> |", err)
		}
	} 
			
	for _, resource := range GetRoleBinding(){
		if resource.ObjectMeta.Namespace == ""{
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetRoleBinding()| Error --> |", err)
		}
	} 
			
	for _, resource := range GetRole(){
		if resource.ObjectMeta.Namespace == ""{
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetRole()| Error --> |", err)
		}
	} 
			
	for _, resource := range GetServiceAccount(){
		if resource.ObjectMeta.Namespace == ""{
			resource.ObjectMeta.Namespace = namespaceProvided
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetServiceAccount()| Error --> |", err)
		}
	} 
			
	for _, resource := range GetNetworkAttachmentDefinition(){
		if resource.GetNamespace() == ""{
			resource.SetNamespace(namespaceProvided)
		}
		err = r.Delete(context.TODO(), resource)
		if err != nil {
			fmt.Println("Erorr During Deleting resource of GetNetworkAttachmentDefinition()| Error --> |", err)
		}
	} 
			
}
*/

	
func GetConfigMap() []*corev1.ConfigMap{
	
	configMap1 := &corev1.ConfigMap{
		Data : map[string]string{
			"amf.yaml" : "logger:\n" + 
 "    file: /var/log/open5gs/amf.log\n" + 
 "    #level: debug\n" + 
 "    #domain: sbi\n" + 
 "\n" + 
 "amf:\n" + 
 "    sbi:\n" + 
 "    - addr: 0.0.0.0\n" + 
 "      advertise: release-name-amf.open5gscns.svc.cluster.local\n" + 
 "    discovery:\n" + 
 "       delegated: no            \n" + 
 "    ngap:\n" + 
 "      dev: \n" + 
 "    metrics:\n" + 
 "      - addr: 0.0.0.0\n" + 
 "        port: 9091          \n" + 
 "    guami:\n" + 
 "      - plmn_id:\n" + 
 "          mcc: 208\n" + 
 "          mnc: 93\n" + 
 "        amf_id:\n" + 
 "          region: 2\n" + 
 "          set: 1\n" + 
 "    tai:\n" + 
 "      - plmn_id:\n" + 
 "          mcc: 208\n" + 
 "          mnc: 93\n" + 
 "        tac: 7\n" + 
 "    plmn_support:\n" + 
 "    - plmn_id:\n" + 
 "        mcc: 208\n" + 
 "        mnc: 93\n" + 
 "      s_nssai:\n" + 
 "      - sst: 1\n" + 
 "        sd: 1\n" + 
 "    security:\n" + 
 "        integrity_order : [ NIA2, NIA1, NIA0 ]\n" + 
 "        ciphering_order : [ NEA0, NEA1, NEA2 ]\n" + 
 "    network_name:\n" + 
 "        full: Open5GS\n" + 
 "    amf_name: open5gs-amf\n" + 
 "sbi:\n" + 
 "    server:\n" + 
 "      no_tls: true\n" + 
 "    client:\n" + 
 "      no_tls: true\n" + 
 "scp:\n" + 
 " sbi:\n" + 
 "   name: release-name-scp.open5gscns.svc.cluster.local\n" + 
 "nrf:\n" + 
 " sbi:\n" + 
 "   name: release-name-nrf.open5gscns.svc.cluster.local\n" + 
 "",
		},
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "amf",
			},
			Name  : "release-name-amf-config", 
		}, 
		TypeMeta  : metav1.TypeMeta{
			Kind  : "ConfigMap", 
			APIVersion  : "v1", 
		}, 
	}
	
		
	configMap2 := &corev1.ConfigMap{
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "ConfigMap", 
		}, 
		Data : map[string]string{
			"ausf.yaml" : "logger:\n" + 
 "    file: /var/log/open5gs/ausf.log\n" + 
 "\n" + 
 "ausf:\n" + 
 "  sbi:\n" + 
 "    - addr: 0.0.0.0\n" + 
 "      advertise: release-name-ausf.open5gscns.svc.cluster.local\n" + 
 "  discovery:\n" + 
 "     delegated: no           \n" + 
 "sbi:\n" + 
 "sbi:\n" + 
 "    server:\n" + 
 "      no_tls: true\n" + 
 "    client:\n" + 
 "      no_tls: true              \n" + 
 "scp:\n" + 
 " sbi:\n" + 
 "   name: release-name-scp.open5gscns.svc.cluster.local\n" + 
 "nrf:\n" + 
 " sbi:\n" + 
 "   name: release-name-nrf.open5gscns.svc.cluster.local\n" + 
 "",
		},
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "ausf",
			},
			Name  : "release-name-ausf-config", 
		}, 
	}
	
		
	configMap3 := &corev1.ConfigMap{
		Data : map[string]string{
			"bsf.yaml" : "logger:\n" + 
 "    file: /var/log/open5gs/bsf.log\n" + 
 "\n" + 
 "bsf:\n" + 
 "  sbi:\n" + 
 "    - addr: 0.0.0.0\n" + 
 "      advertise: release-name-bsf.open5gscns.svc.cluster.local\n" + 
 "  discovery:\n" + 
 "     delegated: no           \n" + 
 "sbi:\n" + 
 "sbi:\n" + 
 "    server:\n" + 
 "      no_tls: true\n" + 
 "    client:\n" + 
 "      no_tls: true              \n" + 
 "scp:\n" + 
 " sbi:\n" + 
 "   name: release-name-scp.open5gscns.svc.cluster.local\n" + 
 "nrf:\n" + 
 " sbi:\n" + 
 "   name: release-name-nrf.open5gscns.svc.cluster.local\n" + 
 "",
		},
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "bsf",
			},
			Name  : "release-name-bsf-config", 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "ConfigMap", 
		}, 
	}
	
		
	configMap4 := &corev1.ConfigMap{
		Data : map[string]string{
			"nrf.yaml" : "logger:\n" + 
 "    file: /var/log/open5gs/nrf.log\n" + 
 "\n" + 
 "nrf:\n" + 
 "  sbi:\n" + 
 "  - dev: eth0\n" + 
 "sbi:\n" + 
 "    server:\n" + 
 "      no_tls: true\n" + 
 "    client:\n" + 
 "      no_tls: true\n" + 
 "",
		},
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "nrf",
			},
			Name  : "release-name-nrf-config", 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "ConfigMap", 
		}, 
	}
	
		
	configMap5 := &corev1.ConfigMap{
		Data : map[string]string{
			"nssf.yaml" : "logger:\n" + 
 "    file: /var/log/open5gs/nssf.log   \n" + 
 "\n" + 
 "nssf:\n" + 
 "  sbi:\n" + 
 "    - addr: 0.0.0.0\n" + 
 "      advertise: release-name-nssf.open5gscns.svc.cluster.local\n" + 
 "  discovery:\n" + 
 "     delegated: no           \n" + 
 "  nsi:\n" + 
 "  - addr: release-name-scp.open5gscns.svc.cluster.local\n" + 
 "    port: 80\n" + 
 "    s_nssai:\n" + 
 "      sst: \"1\"\n" + 
 "      sd: \"1\"\n" + 
 "sbi:\n" + 
 "    server:\n" + 
 "      no_tls: true\n" + 
 "    client:\n" + 
 "      no_tls: true\n" + 
 "scp:\n" + 
 " sbi:\n" + 
 "   name: release-name-scp.open5gscns.svc.cluster.local\n" + 
 "nrf:\n" + 
 " sbi:\n" + 
 "   name: release-name-nrf.open5gscns.svc.cluster.local\n" + 
 "",
		},
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "nssf",
			},
			Name  : "release-name-nssf-config", 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "ConfigMap", 
		}, 
	}
	
		
	configMap6 := &corev1.ConfigMap{
		Data : map[string]string{
			"pcf.yaml" : "logger:\n" + 
 "    file: /var/log/open5gs/pcf.log\n" + 
 "\n" + 
 "db_uri: mongodb://release-name-mongodb-svc/open5gs\n" + 
 "\n" + 
 "pcf:\n" + 
 "  sbi:\n" + 
 "    - addr: 0.0.0.0\n" + 
 "      advertise: release-name-pcf.open5gscns.svc.cluster.local\n" + 
 "  discovery:\n" + 
 "     delegated: no           \n" + 
 "scp:\n" + 
 " sbi:\n" + 
 "   name: release-name-scp.open5gscns.svc.cluster.local\n" + 
 "nrf:\n" + 
 " sbi:\n" + 
 "   name: release-name-nrf.open5gscns.svc.cluster.local \n" + 
 "sbi:\n" + 
 "    server:\n" + 
 "      no_tls: true\n" + 
 "    client:\n" + 
 "      no_tls: true\n" + 
 "",
		},
		ObjectMeta  : metav1.ObjectMeta{
			Name  : "release-name-pcf-config", 
			Labels : map[string]string{
				"epc-mode" : "pcf",
			},
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "ConfigMap", 
		}, 
	}
	
		
	configMap7 := &corev1.ConfigMap{
		Data : map[string]string{
			"scp.yaml" : "logger:\n" + 
 "    file: /var/log/open5gs/scp.log\n" + 
 "\n" + 
 "db_uri: mongodb://release-name-mongodb-svc/open5gs\n" + 
 "\n" + 
 "scp:\n" + 
 "  sbi:\n" + 
 "    - addr: 0.0.0.0\n" + 
 "      advertise: release-name-scp.open5gscns.svc.cluster.local\n" + 
 "nrf:\n" + 
 " sbi:\n" + 
 "   name: release-name-nrf.open5gscns.svc.cluster.local \n" + 
 "sbi:\n" + 
 "    server:\n" + 
 "      no_tls: true\n" + 
 "    client:\n" + 
 "      no_tls: true\n" + 
 "",
		},
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "scp",
			},
			Name  : "release-name-scp-config", 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "ConfigMap", 
		}, 
	}
	
		
	configMap8 := &corev1.ConfigMap{
		Data : map[string]string{
			"smf.yaml" : "logger:\n" + 
 "    file: /var/log/open5gs/smf.log\n" + 
 "\n" + 
 "parameter:\n" + 
 "    no_ipv6: true\n" + 
 "smf:\n" + 
 "    sbi:\n" + 
 "    - addr: 0.0.0.0\n" + 
 "      advertise: release-name-smf.open5gscns.svc.cluster.local\n" + 
 "    discovery:\n" + 
 "       delegated: no           \n" + 
 "    metrics:\n" + 
 "      - addr: 0.0.0.0\n" + 
 "        port: 9091          \n" + 
 "    pfcp:\n" + 
 "       dev: n4\n" + 
 "    gtpc:\n" + 
 "       dev: n4\n" + 
 "    gtpu:\n" + 
 "       dev: eth0\n" + 
 "    subnet:\n" + 
 "     - addr: 10.45.0.1/16\n" + 
 "       dnn: \n" + 
 "    dns:\n" + 
 "      - 8.8.8.8\n" + 
 "      - 8.8.4.4\n" + 
 "    mtu: 1400\n" + 
 "sbi:\n" + 
 "    server:\n" + 
 "      no_tls: true\n" + 
 "    client:\n" + 
 "      no_tls: true   \n" + 
 "scp:\n" + 
 " sbi:\n" + 
 "   name: release-name-scp.open5gscns.svc.cluster.local\n" + 
 "nrf:\n" + 
 " sbi:\n" + 
 "   name: release-name-nrf.open5gscns.svc.cluster.local \n" + 
 "\n" + 
 "upf:\n" + 
 "  pfcp:\n" + 
 "    - name: 192.168.4.2\n" + 
 "      dnn:\n" + 
 "",
		},
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "smf",
			},
			Name  : "release-name-smf-config", 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "ConfigMap", 
		}, 
	}
	
		
	configMap9 := &corev1.ConfigMap{
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "ue-import",
			},
			Name  : "release-name-ue-provisioning", 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "ConfigMap", 
		}, 
		Data : map[string]string{
			"account.js" : "db = db.getSiblingDB('open5gs')\n" + 
 "cursor = db.accounts.find()\n" + 
 "if ( cursor.count() == 0 ) {\n" + 
 "    db.accounts.insert({ salt: 'f5c15fa72622d62b6b790aa8569b9339729801ab8bda5d13997b5db6bfc1d997', hash: '402223057db5194899d2e082aeb0802f6794622e1cbc47529c419e5a603f2cc592074b4f3323b239ffa594c8b756d5c70a4e1f6ecd3f9f0d2d7328c4cf8b1b766514effff0350a90b89e21eac54cd4497a169c0c7554a0e2cd9b672e5414c323f76b8559bc768cba11cad2ea3ae704fb36abc8abc2619231ff84ded60063c6e1554a9777a4a464ef9cfdfa90ecfdacc9844e0e3b2f91b59d9ff024aec4ea1f51b703a31cda9afb1cc2c719a09cee4f9852ba3cf9f07159b1ccf8133924f74df770b1a391c19e8d67ffdcbbef4084a3277e93f55ac60d80338172b2a7b3f29cfe8a36738681794f7ccbe9bc98f8cdeded02f8a4cd0d4b54e1d6ba3d11792ee0ae8801213691848e9c5338e39485816bb0f734b775ac89f454ef90992003511aa8cceed58a3ac2c3814f14afaaed39cbaf4e2719d7213f81665564eec02f60ede838212555873ef742f6666cc66883dcb8281715d5c762fb236d72b770257e7e8d86c122bb69028a34cf1ed93bb973b440fa89a23604cd3fefe85fbd7f55c9b71acf6ad167228c79513f5cfe899a2e2cc498feb6d2d2f07354a17ba74cecfbda3e87d57b147e17dcc7f4c52b802a8e77f28d255a6712dcdc1519e6ac9ec593270bfcf4c395e2531a271a841b1adefb8516a07136b0de47c7fd534601b16f0f7a98f1dbd31795feb97da59e1d23c08461cf37d6f2877d0f2e437f07e25015960f63', username: 'admin', roles: [ 'admin' ], \"__v\" : 0})\n" + 
 "}\n" + 
 "",
			"open5gs-dbctl" : "#!/bin/bash\n" + 
 "\n" + 
 "version=0.10.1\n" + 
 "\n" + 
 "display_help() {\n" + 
 "  echo \"open5gs-dbctl: Open5GS Database Configuration Tool ($version)\"\n" + 
 "  echo \"FLAGS: --db_uri=mongodb://localhost\"\n" + 
 "  echo \"COMMANDS:\" >&2\n" + 
 "  echo \"   add {imsi key opc}: adds a user to the database with default values\"\n" + 
 "  echo \"   add {imsi ip key opc}: adds a user to the database with default values and a IPv4 address for the UE\"\n" + 
 "  echo \"   addT1 {imsi key opc}: adds a user to the database with 3 differents apns\"\n" + 
 "  echo \"   addT1 {imsi ip key opc}: adds a user to the database with 3 differents apns and the same IPv4 address for the each apn\"\n" + 
 "  echo \"   remove {imsi}: removes a user from the database\"\n" + 
 "  echo \"   reset: WIPES OUT the database and restores it to an empty default\"\n" + 
 "  echo \"   static_ip {imsi ip4}: adds a static IP assignment to an already-existing user\"\n" + 
 "  echo \"   static_ip6 {imsi ip6}: adds a static IPv6 assignment to an already-existing user\"\n" + 
 "  echo \"   type {imsi type}: changes the PDN-Type of the first PDN: 0 = IPv4, 1 = IPv6, 2 = IPv4v6, 3 = v4 OR v6\"\n" + 
 "  echo \"   help: displays this message and exits\"\n" + 
 "  echo \"   default values are as follows: APN \\\"internet\\\", dl_bw/ul_bw 1 Gbps, PGW address is 127.0.0.3, IPv4 only\"\n" + 
 "  echo \"   add_ue_with_apn {imsi key opc apn}: adds a user to the database with a specific apn,\"\n" + 
 "  echo \"   add_ue_with_slice {imsi key opc apn sst sd}: adds a user to the database with a specific apn, sst and sd\"\n" + 
 "  echo \"   update_apn {imsi apn slice_num}: adds an APN to the slice number slice_num of an existent UE\"\n" + 
 "  echo \"   update_slice {imsi apn sst sd}: adds an slice to an existent UE\"\n" + 
 "  echo \"   showall: shows the list of subscriber in the db\"\n" + 
 "  echo \"   showpretty: shows the list of subscriber in the db in a pretty json tree format\"\n" + 
 "  echo \"   showfiltered: shows {imsi key opc apn ip} information of subscriber\"\n" + 
 "\n" + 
 "\n" + 
 "}\n" + 
 "\n" + 
 "#while test $# -gt 0; do\n" + 
 "#  case \"$1\" in\n" + 
 "#    --db_uri*)\n" + 
 "#      DB_URI=`echo $1 | sed -e 's/^[^=]*=//g'`\n" + 
 "#      shift\n" + 
 "#      ;;\n" + 
 "#    *)\n" + 
 "#      break\n" + 
 "#      ;;\n" + 
 "#  esac\n" + 
 "#done\n" + 
 "\n" + 
 "DB_URI=\"${DB_URI:-mongodb://localhost/open5gs}\"\n" + 
 "\n" + 
 "if [ \"$#\" -lt 1 ]; then\n" + 
 "  display_help\n" + 
 "  exit 1\n" + 
 "fi\n" + 
 "\n" + 
 "if [ \"$1\" = \"help\" ]; then\n" + 
 "  display_help\n" + 
 "  exit 1\n" + 
 "fi\n" + 
 "\n" + 
 "if [ \"$1\" = \"add\" ]; then\n" + 
 "  if [ \"$#\" -eq 4 ]; then\n" + 
 "      IMSI=$2\n" + 
 "      KI=$3\n" + 
 "      OPC=$4\n" + 
 "\n" + 
 "      mongo --quiet --eval \"db.subscribers.update( { \\\"imsi\\\" : \\\"$IMSI\\\" },\n" + 
 "          { \\$setOnInsert:\n" + 
 "              {\n" + 
 "                  \\\"imsi\\\" : \\\"$IMSI\\\",\n" + 
 "                  \\\"subscribed_rau_tau_timer\\\" : NumberInt(12),\n" + 
 "                  \\\"network_access_mode\\\" : NumberInt(0),\n" + 
 "                  \\\"subscriber_status\\\" : NumberInt(0),\n" + 
 "                  \\\"access_restriction_data\\\" : NumberInt(32),\n" + 
 "                  \\\"slice\\\" :\n" + 
 "                  [{\n" + 
 "                      \\\"sst\\\" : NumberInt(1),\n" + 
 "                      \\\"default_indicator\\\" : true,\n" + 
 "                      \\\"_id\\\" : new ObjectId(),\n" + 
 "                      \\\"session\\\" :\n" + 
 "                      [{\n" + 
 "                          \\\"name\\\" : \\\"internet\\\",\n" + 
 "                          \\\"type\\\" : NumberInt(3),\n" + 
 "                          \\\"_id\\\" : new ObjectId(),\n" + 
 "                          \\\"pcc_rule\\\" : [],\n" + 
 "                          \\\"ambr\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"uplink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                              \\\"downlink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                          },\n" + 
 "                          \\\"qos\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"index\\\" : NumberInt(9),\n" + 
 "                              \\\"arp\\\" :\n" + 
 "                              {\n" + 
 "                                  \\\"priority_level\\\" : NumberInt(8),\n" + 
 "                                  \\\"pre_emption_capability\\\" : NumberInt(1),\n" + 
 "                                  \\\"pre_emption_vulnerability\\\" : NumberInt(1),\n" + 
 "                              },\n" + 
 "                          },\n" + 
 "                      }],\n" + 
 "                  }],\n" + 
 "                  \\\"ambr\\\" :\n" + 
 "                  {\n" + 
 "                      \\\"uplink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3),},\n" + 
 "                      \\\"downlink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                  },\n" + 
 "                  \\\"security\\\" :\n" + 
 "                  {\n" + 
 "                      \\\"k\\\" : \\\"$KI\\\",\n" + 
 "                      \\\"amf\\\" : \\\"8000\\\",\n" + 
 "                      \\\"op\\\" : null,\n" + 
 "                      \\\"opc\\\" : \\\"$OPC\\\"\n" + 
 "                  },\n" + 
 "                  \\\"__v\\\" : 0\n" + 
 "              },\n" + 
 "          },\n" + 
 "          upsert=true);\" $DB_URI\n" + 
 "      exit $?\n" + 
 "  fi\n" + 
 "\n" + 
 "  if [ \"$#\" -eq 5 ]; then\n" + 
 "      IMSI=$2\n" + 
 "      IP=$3\n" + 
 "      KI=$4\n" + 
 "      OPC=$5\n" + 
 "\n" + 
 "      mongo --quiet --eval \"db.subscribers.update( { \\\"imsi\\\" : \\\"$IMSI\\\" },\n" + 
 "          { \\$setOnInsert:\n" + 
 "              {\n" + 
 "                  \\\"imsi\\\" : \\\"$IMSI\\\",\n" + 
 "                  \\\"subscribed_rau_tau_timer\\\" : NumberInt(12),\n" + 
 "                  \\\"network_access_mode\\\" : NumberInt(0),\n" + 
 "                  \\\"subscriber_status\\\" : NumberInt(0),\n" + 
 "                  \\\"access_restriction_data\\\" : NumberInt(32),\n" + 
 "                  \\\"slice\\\" :\n" + 
 "                  [{\n" + 
 "                      \\\"sst\\\" : NumberInt(1),\n" + 
 "                      \\\"default_indicator\\\" : true,\n" + 
 "                      \\\"_id\\\" : new ObjectId(),\n" + 
 "                      \\\"session\\\" :\n" + 
 "                      [{\n" + 
 "                          \\\"name\\\" : \\\"internet\\\",\n" + 
 "                          \\\"type\\\" : NumberInt(3),\n" + 
 "                          \\\"_id\\\" : new ObjectId(),\n" + 
 "                          \\\"pcc_rule\\\" : [],\n" + 
 "                          \\\"ue\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"addr\\\" : \\\"$IP\\\",\n" + 
 "                          },\n" + 
 "                          \\\"ambr\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"uplink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                              \\\"downlink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                          },\n" + 
 "                          \\\"qos\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"index\\\" : NumberInt(9),\n" + 
 "                              \\\"arp\\\" :\n" + 
 "                              {\n" + 
 "                                  \\\"priority_level\\\" : NumberInt(8),\n" + 
 "                                  \\\"pre_emption_capability\\\" : NumberInt(1),\n" + 
 "                                  \\\"pre_emption_vulnerability\\\" : NumberInt(1),\n" + 
 "                              },\n" + 
 "                          },\n" + 
 "                      }],\n" + 
 "                  }],\n" + 
 "                  \\\"ambr\\\" :\n" + 
 "                  {\n" + 
 "                      \\\"uplink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3),},\n" + 
 "                      \\\"downlink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                  },\n" + 
 "                  \\\"security\\\" :\n" + 
 "                  {\n" + 
 "                      \\\"k\\\" : \\\"$KI\\\",\n" + 
 "                      \\\"amf\\\" : \\\"8000\\\",\n" + 
 "                      \\\"op\\\" : null,\n" + 
 "                      \\\"opc\\\" : \\\"$OPC\\\"\n" + 
 "                  },\n" + 
 "                  \\\"__v\\\" : 0\n" + 
 "              },\n" + 
 "          },\n" + 
 "          upsert=true);\" $DB_URI\n" + 
 "      exit $?\n" + 
 "  fi\n" + 
 "\n" + 
 "  echo \"open5gs-dbctl: incorrect number of args, format is \\\"open5gs-dbctl add imsi key opc\\\"\"\n" + 
 "  exit 1\n" + 
 "fi\n" + 
 "\n" + 
 "if [ \"$1\" = \"addT1\" ]; then\n" + 
 "  if [ \"$#\" -eq 4 ]; then\n" + 
 "      IMSI=$2\n" + 
 "      KI=$3\n" + 
 "      OPC=$4\n" + 
 "\n" + 
 "      mongo --quiet --eval \"db.subscribers.update( { \\\"imsi\\\" : \\\"$IMSI\\\" },\n" + 
 "          { \\$setOnInsert:\n" + 
 "              {\n" + 
 "                  \\\"imsi\\\" : \\\"$IMSI\\\",\n" + 
 "                  \\\"subscribed_rau_tau_timer\\\" : NumberInt(12),\n" + 
 "                  \\\"network_access_mode\\\" : NumberInt(0),\n" + 
 "                  \\\"subscriber_status\\\" : NumberInt(0),\n" + 
 "                  \\\"access_restriction_data\\\" : NumberInt(32),\n" + 
 "                  \\\"slice\\\" :\n" + 
 "                  [{\n" + 
 "                      \\\"sst\\\" : NumberInt(1),\n" + 
 "                      \\\"default_indicator\\\" : true,\n" + 
 "                      \\\"_id\\\" : new ObjectId(),\n" + 
 "                      \\\"session\\\" :\n" + 
 "                      [{\n" + 
 "                          \\\"name\\\" : \\\"internet\\\",\n" + 
 "                          \\\"type\\\" : NumberInt(3),\n" + 
 "                          \\\"_id\\\" : new ObjectId(),\n" + 
 "                          \\\"pcc_rule\\\" : [],\n" + 
 "                          \\\"ambr\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"uplink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                              \\\"downlink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                          },\n" + 
 "                          \\\"qos\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"index\\\" : NumberInt(9),\n" + 
 "                              \\\"arp\\\" :\n" + 
 "                              {\n" + 
 "                                  \\\"priority_level\\\" : NumberInt(8),\n" + 
 "                                  \\\"pre_emption_capability\\\" : NumberInt(1),\n" + 
 "                                  \\\"pre_emption_vulnerability\\\" : NumberInt(1),\n" + 
 "                              },\n" + 
 "                          },\n" + 
 "                      },{\n" + 
 "                          \\\"name\\\" : \\\"internet1\\\",\n" + 
 "                          \\\"type\\\" : NumberInt(3),\n" + 
 "                          \\\"_id\\\" : new ObjectId(),\n" + 
 "                          \\\"pcc_rule\\\" : [],\n" + 
 "                          \\\"ambr\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"uplink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                              \\\"downlink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                          },\n" + 
 "                          \\\"qos\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"index\\\" : NumberInt(9),\n" + 
 "                              \\\"arp\\\" :\n" + 
 "                              {\n" + 
 "                                  \\\"priority_level\\\" : NumberInt(8),\n" + 
 "                                  \\\"pre_emption_capability\\\" : NumberInt(1),\n" + 
 "                                  \\\"pre_emption_vulnerability\\\" : NumberInt(1),\n" + 
 "                              },\n" + 
 "                          },\n" + 
 "                      },{\n" + 
 "                          \\\"name\\\" : \\\"internet2\\\",\n" + 
 "                          \\\"type\\\" : NumberInt(3),\n" + 
 "                          \\\"_id\\\" : new ObjectId(),\n" + 
 "                          \\\"pcc_rule\\\" : [],\n" + 
 "                          \\\"ambr\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"uplink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                              \\\"downlink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                          },\n" + 
 "                          \\\"qos\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"index\\\" : NumberInt(9),\n" + 
 "                              \\\"arp\\\" :\n" + 
 "                              {\n" + 
 "                                  \\\"priority_level\\\" : NumberInt(8),\n" + 
 "                                  \\\"pre_emption_capability\\\" : NumberInt(1),\n" + 
 "                                  \\\"pre_emption_vulnerability\\\" : NumberInt(1),\n" + 
 "                              },\n" + 
 "                          },\n" + 
 "                      }],\n" + 
 "                  }],\n" + 
 "                  \\\"ambr\\\" :\n" + 
 "                  {\n" + 
 "                      \\\"uplink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3),},\n" + 
 "                      \\\"downlink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                  },\n" + 
 "                  \\\"security\\\" :\n" + 
 "                  {\n" + 
 "                      \\\"k\\\" : \\\"$KI\\\",\n" + 
 "                      \\\"amf\\\" : \\\"8000\\\",\n" + 
 "                      \\\"op\\\" : null,\n" + 
 "                      \\\"opc\\\" : \\\"$OPC\\\"\n" + 
 "                  },\n" + 
 "                  \\\"__v\\\" : 0\n" + 
 "              },\n" + 
 "          },\n" + 
 "          upsert=true);\" $DB_URI\n" + 
 "      exit $?\n" + 
 "  fi\n" + 
 "\n" + 
 "  if [ \"$#\" -eq 5 ]; then\n" + 
 "      IMSI=$2\n" + 
 "      IP=$3\n" + 
 "      KI=$4\n" + 
 "      OPC=$5\n" + 
 "\n" + 
 "      mongo --quiet --eval \"db.subscribers.update( { \\\"imsi\\\" : \\\"$IMSI\\\" },\n" + 
 "          { \\$setOnInsert:\n" + 
 "              {\n" + 
 "                  \\\"imsi\\\" : \\\"$IMSI\\\",\n" + 
 "                  \\\"subscribed_rau_tau_timer\\\" : NumberInt(12),\n" + 
 "                  \\\"network_access_mode\\\" : NumberInt(0),\n" + 
 "                  \\\"subscriber_status\\\" : NumberInt(0),\n" + 
 "                  \\\"access_restriction_data\\\" : NumberInt(32),\n" + 
 "                  \\\"slice\\\" :\n" + 
 "                  [{\n" + 
 "                      \\\"sst\\\" : NumberInt(1),\n" + 
 "                      \\\"default_indicator\\\" : true,\n" + 
 "                      \\\"_id\\\" : new ObjectId(),\n" + 
 "                      \\\"session\\\" :\n" + 
 "                      [{\n" + 
 "                          \\\"name\\\" : \\\"internet\\\",\n" + 
 "                          \\\"type\\\" : NumberInt(3),\n" + 
 "                          \\\"_id\\\" : new ObjectId(),\n" + 
 "                          \\\"pcc_rule\\\" : [],\n" + 
 "                          \\\"ue\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"addr\\\" : \\\"$IP\\\",\n" + 
 "                          },\n" + 
 "                          \\\"ambr\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"uplink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                              \\\"downlink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                          },\n" + 
 "                          \\\"qos\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"index\\\" : NumberInt(9),\n" + 
 "                              \\\"arp\\\" :\n" + 
 "                              {\n" + 
 "                                  \\\"priority_level\\\" : NumberInt(8),\n" + 
 "                                  \\\"pre_emption_capability\\\" : NumberInt(1),\n" + 
 "                                  \\\"pre_emption_vulnerability\\\" : NumberInt(1),\n" + 
 "                              },\n" + 
 "                          },\n" + 
 "                      },{\n" + 
 "                          \\\"name\\\" : \\\"internet1\\\",\n" + 
 "                          \\\"type\\\" : NumberInt(3),\n" + 
 "                          \\\"_id\\\" : new ObjectId(),\n" + 
 "                          \\\"pcc_rule\\\" : [],\n" + 
 "                          \\\"ue\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"addr\\\" : \\\"$IP\\\",\n" + 
 "                          },\n" + 
 "                          \\\"ambr\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"uplink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                              \\\"downlink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                          },\n" + 
 "                          \\\"qos\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"index\\\" : NumberInt(9),\n" + 
 "                              \\\"arp\\\" :\n" + 
 "                              {\n" + 
 "                                  \\\"priority_level\\\" : NumberInt(8),\n" + 
 "                                  \\\"pre_emption_capability\\\" : NumberInt(1),\n" + 
 "                                  \\\"pre_emption_vulnerability\\\" : NumberInt(1),\n" + 
 "                              },\n" + 
 "                          },\n" + 
 "                      },{\n" + 
 "                          \\\"name\\\" : \\\"internet2\\\",\n" + 
 "                          \\\"type\\\" : NumberInt(3),\n" + 
 "                          \\\"_id\\\" : new ObjectId(),\n" + 
 "                          \\\"pcc_rule\\\" : [],\n" + 
 "                          \\\"ue\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"addr\\\" : \\\"$IP\\\",\n" + 
 "                          },\n" + 
 "                          \\\"ambr\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"uplink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                              \\\"downlink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                          },\n" + 
 "                          \\\"qos\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"index\\\" : NumberInt(9),\n" + 
 "                              \\\"arp\\\" :\n" + 
 "                              {\n" + 
 "                                  \\\"priority_level\\\" : NumberInt(8),\n" + 
 "                                  \\\"pre_emption_capability\\\" : NumberInt(1),\n" + 
 "                                  \\\"pre_emption_vulnerability\\\" : NumberInt(1),\n" + 
 "                              },\n" + 
 "                          },\n" + 
 "                      }],\n" + 
 "                  }],\n" + 
 "                  \\\"ambr\\\" :\n" + 
 "                  {\n" + 
 "                      \\\"uplink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3),},\n" + 
 "                      \\\"downlink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                  },\n" + 
 "                  \\\"security\\\" :\n" + 
 "                  {\n" + 
 "                      \\\"k\\\" : \\\"$KI\\\",\n" + 
 "                      \\\"amf\\\" : \\\"8000\\\",\n" + 
 "                      \\\"op\\\" : null,\n" + 
 "                      \\\"opc\\\" : \\\"$OPC\\\"\n" + 
 "                  },\n" + 
 "                  \\\"__v\\\" : 0\n" + 
 "              },\n" + 
 "          },\n" + 
 "          upsert=true);\" $DB_URI\n" + 
 "      exit $?\n" + 
 "  fi\n" + 
 "\n" + 
 "  echo \"open5gs-dbctl: incorrect number of args, format is \\\"open5gs-dbctl add imsi key opc\\\"\"\n" + 
 "  exit 1\n" + 
 "fi\n" + 
 "\n" + 
 "if [ \"$1\" = \"remove\" ]; then\n" + 
 "  if [ \"$#\" -ne 2 ]; then\n" + 
 "      echo \"open5gs-dbctl: incorrect number of args, format is \\\"open5gs-dbctl remove imsi\\\"\"\n" + 
 "      exit 1\n" + 
 "  fi\n" + 
 "\n" + 
 "  IMSI=$2\n" + 
 "  mongo --quiet --eval \"db.subscribers.remove({\\\"imsi\\\": \\\"$IMSI\\\"});\" $DB_URI\n" + 
 "  exit $?\n" + 
 "fi\n" + 
 "\n" + 
 "if [ \"$1\" = \"reset\" ]; then\n" + 
 "  if [ \"$#\" -ne 1 ]; then\n" + 
 "      echo \"open5gs-dbctl: incorrect number of args, format is \\\"open5gs-dbctl reset\\\"\"\n" + 
 "      exit 1\n" + 
 "  fi\n" + 
 "\n" + 
 "  mongo --quiet --eval \"db.subscribers.remove({});\" $DB_URI\n" + 
 "  exit $?\n" + 
 "fi\n" + 
 "\n" + 
 "if [ \"$1\" = \"static_ip\" ]; then\n" + 
 "  if [ \"$#\" -ne 3 ]; then\n" + 
 "      echo \"open5gs-dbctl: incorrect number of args, format is \\\"open5gs-dbctl static_ip imsi ip\\\"\"\n" + 
 "      exit 1\n" + 
 "  fi\n" + 
 "  IMSI=$2\n" + 
 "  IP=$3\n" + 
 "\n" + 
 "  mongo --quiet --eval \"db.subscribers.update({\\\"imsi\\\": \\\"$IMSI\\\"},{\\$set: { \\\"slice.0.session.0.ue.addr\\\": \\\"$IP\\\" }});\" $DB_URI\n" + 
 "  exit $?\n" + 
 "fi\n" + 
 "\n" + 
 "if [ \"$1\" = \"static_ip6\" ]; then\n" + 
 "  if [ \"$#\" -ne 3 ]; then\n" + 
 "      echo \"open5gs-dbctl: incorrect number of args, format is \\\"open5gs-dbctl static_ip6 imsi ip\\\"\"\n" + 
 "      exit 1\n" + 
 "  fi\n" + 
 "  IMSI=$2\n" + 
 "  IP=$3\n" + 
 "\n" + 
 "  mongo --eval \"db.subscribers.update({\\\"imsi\\\": \\\"$IMSI\\\"},{\\$set: { \\\"slice.0.session.0.ue.addr6\\\": \\\"$IP\\\" }});\" $DB_URI\n" + 
 "  exit $?\n" + 
 "fi\n" + 
 "\n" + 
 "if [ \"$1\" = \"type\" ]; then\n" + 
 "  if [ \"$#\" -ne 3 ]; then\n" + 
 "      echo \"open5gs-dbctl: incorrect number of args, format is \\\"open5gs-dbctl type imsi type\\\"\"\n" + 
 "      exit 1\n" + 
 "  fi\n" + 
 "  IMSI=$2\n" + 
 "  TYPE=$3\n" + 
 "\n" + 
 "  mongo --quiet --eval \"db.subscribers.update({\\\"imsi\\\": \\\"$IMSI\\\"},{\\$set: { \\\"slice.0.session.0.type\\\": NumberInt($TYPE) }});\" $DB_URI\n" + 
 "  exit $?\n" + 
 "fi\n" + 
 "\n" + 
 "if [ \"$1\" = \"add_ue_with_apn\" ]; then\n" + 
 "  if [ \"$#\" -eq 5 ]; then\n" + 
 "      IMSI=$2\n" + 
 "      KI=$3\n" + 
 "      OPC=$4\n" + 
 "      APN=$5\n" + 
 "\n" + 
 "      mongo --quiet --eval \"db.subscribers.update( { \\\"imsi\\\" : \\\"$IMSI\\\" },\n" + 
 "          { \\$setOnInsert:\n" + 
 "              {\n" + 
 "                  \\\"imsi\\\" : \\\"$IMSI\\\",\n" + 
 "                  \\\"subscribed_rau_tau_timer\\\" : NumberInt(12),\n" + 
 "                  \\\"network_access_mode\\\" : NumberInt(0),\n" + 
 "                  \\\"subscriber_status\\\" : NumberInt(0),\n" + 
 "                  \\\"access_restriction_data\\\" : NumberInt(32),\n" + 
 "                  \\\"slice\\\" :\n" + 
 "                  [{\n" + 
 "                      \\\"sst\\\" : NumberInt(1),\n" + 
 "                      \\\"default_indicator\\\" : true,\n" + 
 "                      \\\"_id\\\" : new ObjectId(),\n" + 
 "                      \\\"session\\\" :\n" + 
 "                      [{\n" + 
 "                          \\\"name\\\" : \\\"$APN\\\",\n" + 
 "                          \\\"type\\\" : NumberInt(3),\n" + 
 "                          \\\"_id\\\" : new ObjectId(),\n" + 
 "                          \\\"pcc_rule\\\" : [],\n" + 
 "                          \\\"ambr\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"uplink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                              \\\"downlink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                          },\n" + 
 "                          \\\"qos\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"index\\\" : NumberInt(9),\n" + 
 "                              \\\"arp\\\" :\n" + 
 "                              {\n" + 
 "                                  \\\"priority_level\\\" : NumberInt(8),\n" + 
 "                                  \\\"pre_emption_capability\\\" : NumberInt(1),\n" + 
 "                                  \\\"pre_emption_vulnerability\\\" : NumberInt(1),\n" + 
 "                              },\n" + 
 "                          },\n" + 
 "                      }],\n" + 
 "                  }],\n" + 
 "                  \\\"ambr\\\" :\n" + 
 "                  {\n" + 
 "                      \\\"uplink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3),},\n" + 
 "                      \\\"downlink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                  },\n" + 
 "                  \\\"security\\\" :\n" + 
 "                  {\n" + 
 "                      \\\"k\\\" : \\\"$KI\\\",\n" + 
 "                      \\\"amf\\\" : \\\"8000\\\",\n" + 
 "                      \\\"op\\\" : null,\n" + 
 "                      \\\"opc\\\" : \\\"$OPC\\\"\n" + 
 "                  },\n" + 
 "                  \\\"__v\\\" : 0\n" + 
 "              },\n" + 
 "          },\n" + 
 "          upsert=true);\" $DB_URI\n" + 
 "      exit $?\n" + 
 "  fi\n" + 
 "\n" + 
 "  echo \"open5gs-dbctl: incorrect number of args, format is \\\"open5gs-dbctl add_ue_with_apn imsi key opc apn\\\"\"\n" + 
 "  exit 1\n" + 
 "fi\n" + 
 "\n" + 
 "if [ \"$1\" = \"add_ue_with_slice\" ]; then\n" + 
 "  if [ \"$#\" -eq 7 ]; then\n" + 
 "      IMSI=$2\n" + 
 "      KI=$3\n" + 
 "      OPC=$4\n" + 
 "      APN=$5\n" + 
 "      SST=$6\n" + 
 "      SD=$7\n" + 
 "\n" + 
 "      mongo --quiet --eval \"db.subscribers.update( { \\\"imsi\\\" : \\\"$IMSI\\\" },\n" + 
 "          { \\$setOnInsert:\n" + 
 "              {\n" + 
 "                  \\\"imsi\\\" : \\\"$IMSI\\\",\n" + 
 "                  \\\"subscribed_rau_tau_timer\\\" : NumberInt(12),\n" + 
 "                  \\\"network_access_mode\\\" : NumberInt(0),\n" + 
 "                  \\\"subscriber_status\\\" : NumberInt(0),\n" + 
 "                  \\\"access_restriction_data\\\" : NumberInt(32),\n" + 
 "                  \\\"slice\\\" :\n" + 
 "                  [{\n" + 
 "                      \\\"sst\\\" : NumberInt($SST),\n" + 
 "                      \\\"sd\\\" : \\\"$SD\\\",\n" + 
 "                      \\\"default_indicator\\\" : true,\n" + 
 "                      \\\"_id\\\" : new ObjectId(),\n" + 
 "                      \\\"session\\\" :\n" + 
 "                      [{\n" + 
 "                          \\\"name\\\" : \\\"$APN\\\",\n" + 
 "                          \\\"type\\\" : NumberInt(3),\n" + 
 "                          \\\"_id\\\" : new ObjectId(),\n" + 
 "                          \\\"pcc_rule\\\" : [],\n" + 
 "                          \\\"ambr\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"uplink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                              \\\"downlink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                          },\n" + 
 "                          \\\"qos\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"index\\\" : NumberInt(9),\n" + 
 "                              \\\"arp\\\" :\n" + 
 "                              {\n" + 
 "                                  \\\"priority_level\\\" : NumberInt(8),\n" + 
 "                                  \\\"pre_emption_capability\\\" : NumberInt(1),\n" + 
 "                                  \\\"pre_emption_vulnerability\\\" : NumberInt(1),\n" + 
 "                              },\n" + 
 "                          },\n" + 
 "                      }],\n" + 
 "                  }],\n" + 
 "                  \\\"ambr\\\" :\n" + 
 "                  {\n" + 
 "                      \\\"uplink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3),},\n" + 
 "                      \\\"downlink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                  },\n" + 
 "                  \\\"security\\\" :\n" + 
 "                  {\n" + 
 "                      \\\"k\\\" : \\\"$KI\\\",\n" + 
 "                      \\\"amf\\\" : \\\"8000\\\",\n" + 
 "                      \\\"op\\\" : null,\n" + 
 "                      \\\"opc\\\" : \\\"$OPC\\\"\n" + 
 "                  },\n" + 
 "                  \\\"__v\\\" : 0\n" + 
 "              },\n" + 
 "          },\n" + 
 "          upsert=true);\" $DB_URI\n" + 
 "      exit $?\n" + 
 "  fi\n" + 
 "\n" + 
 "  echo \"open5gs-dbctl: incorrect number of args, format is \\\"open5gs-dbctl add_ue_with_slice imsi key opc apn sst sd\\\"\"\n" + 
 "  exit 1\n" + 
 "fi\n" + 
 "\n" + 
 "if [ \"$1\" = \"update_apn\" ]; then\n" + 
 "  if [ \"$#\" -eq 4 ]; then\n" + 
 "      IMSI=$2\n" + 
 "      APN=$3\n" + 
 "      SLICE_NUM=$4\n" + 
 "\n" + 
 "      mongo --quiet --eval \"db.subscribers.updateOne({ \\\"imsi\\\": \\\"$IMSI\\\"},\n" + 
 "          {\\$push: { \\\"slice.$SLICE_NUM.session\\\":\n" + 
 "                          {\n" + 
 "                          \\\"name\\\" : \\\"$APN\\\",\n" + 
 "                          \\\"type\\\" : NumberInt(3),\n" + 
 "                          \\\"_id\\\" : new ObjectId(),\n" + 
 "                          \\\"pcc_rule\\\" : [],\n" + 
 "                          \\\"ambr\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"uplink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                              \\\"downlink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                          },\n" + 
 "                          \\\"qos\\\" :\n" + 
 "                          {\n" + 
 "                              \\\"index\\\" : NumberInt(9),\n" + 
 "                              \\\"arp\\\" :\n" + 
 "                              {\n" + 
 "                                  \\\"priority_level\\\" : NumberInt(8),\n" + 
 "                                  \\\"pre_emption_capability\\\" : NumberInt(1),\n" + 
 "                                  \\\"pre_emption_vulnerability\\\" : NumberInt(1),\n" + 
 "                              },\n" + 
 "                          },\n" + 
 "                          }\n" + 
 "                  }\n" + 
 "          });\" $DB_URI\n" + 
 "      exit $?\n" + 
 "  fi\n" + 
 "\n" + 
 "  echo \"open5gs-dbctl: incorrect number of args, format is \\\"open5gs-dbctl update_apn imsi apn num_slice\\\"\"\n" + 
 "  exit 1\n" + 
 "fi\n" + 
 "\n" + 
 "if [ \"$1\" = \"update_slice\" ]; then\n" + 
 "  if [ \"$#\" -eq 5 ]; then\n" + 
 "      IMSI=$2\n" + 
 "      APN=$3\n" + 
 "      SST=$4\n" + 
 "      SD=$5\n" + 
 "\n" + 
 "      mongo --eval \"db.subscribers.updateOne({ \\\"imsi\\\": \\\"$IMSI\\\"},\n" + 
 "          {\\$push: { \\\"slice\\\":\n" + 
 "\n" + 
 "                          {\n" + 
 "                          \\\"sst\\\" : NumberInt($SST),\n" + 
 "                          \\\"sd\\\" : \\\"$SD\\\",\n" + 
 "                          \\\"default_indicator\\\" : false,\n" + 
 "                          \\\"_id\\\" : new ObjectId(),\n" + 
 "                          \\\"session\\\" :\n" + 
 "                          [{\n" + 
 "                              \\\"name\\\" : \\\"$APN\\\",\n" + 
 "                              \\\"type\\\" : NumberInt(3),\n" + 
 "                              \\\"_id\\\" : new ObjectId(),\n" + 
 "                              \\\"pcc_rule\\\" : [],\n" + 
 "                              \\\"ambr\\\" :\n" + 
 "                              {\n" + 
 "                                  \\\"uplink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                                  \\\"downlink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                              },\n" + 
 "                              \\\"qos\\\" :\n" + 
 "                              {\n" + 
 "                                  \\\"index\\\" : NumberInt(9),\n" + 
 "                                  \\\"arp\\\" :\n" + 
 "                                  {\n" + 
 "                                      \\\"priority_level\\\" : NumberInt(8),\n" + 
 "                                      \\\"pre_emption_capability\\\" : NumberInt(1),\n" + 
 "                                      \\\"pre_emption_vulnerability\\\" : NumberInt(1),\n" + 
 "                                  },\n" + 
 "                              },\n" + 
 "                            }]\n" + 
 "                          }\n" + 
 "\n" + 
 "                  }\n" + 
 "          });\" $DB_URI\n" + 
 "      exit $?\n" + 
 "  fi\n" + 
 "\n" + 
 "  echo \"open5gs-dbctl: incorrect number of args, format is \\\"open5gs-dbctl update_slice imsi apn sst sd\\\"\"\n" + 
 "  exit 1\n" + 
 "fi\n" + 
 "if [ \"$1\" = \"showall\" ]; then\n" + 
 "  mongo --quiet --eval \"db.subscribers.find()\" $DB_URI\n" + 
 "      exit $?\n" + 
 "fi\n" + 
 "if [ \"$1\" = \"showpretty\" ]; then\n" + 
 "  mongo --quiet --eval \"db.subscribers.find({},{_id:0}).pretty().toArray()\" $DB_URI\n" + 
 "      exit $?\n" + 
 "fi\n" + 
 "if [ \"$1\" = \"showfiltered\" ]; then\n" + 
 "  mongo --quiet --eval \"db.subscribers.find({},{'_id':0,'imsi':1,'security.k':1, 'security.opc':1,'slice.session.name':1,'slice.session.ue.addr':1})\" $DB_URI\n" + 
 "      exit $?\n" + 
 "fi\n" + 
 "  echo \"open5gs-dbctl: incorrect number of args, format is \\\"open5gs-dbctl add imsi key opc\\\"\"\n" + 
 "  exit 1\n" + 
 "fi\n" + 
 "\n" + 
 "display_help\n" + 
 "",
			"ue.sh" : "IMSI=\"208930000000031\"\n" + 
 "KI=\"0C0A34601D4F07677303652C0462535B\"\n" + 
 "OPC=\"63bfa50ee6523365ff14c1f45f88737d\"\n" + 
 "APN=\"internet\"\n" + 
 "SST=\"1\"\n" + 
 "SD=\"1\"\n" + 
 "DB_URI=\"mongodb://release-name-mongodb-svc/open5gs\"\n" + 
 "\n" + 
 "mongo --quiet --eval \"db.subscribers.update( { \\\"imsi\\\" : \\\"$IMSI\\\" },\n" + 
 "    { \\$setOnInsert:\n" + 
 "        {\n" + 
 "            \\\"imsi\\\" : \\\"$IMSI\\\",\n" + 
 "            \\\"subscribed_rau_tau_timer\\\" : NumberInt(12),\n" + 
 "            \\\"network_access_mode\\\" : NumberInt(0),\n" + 
 "            \\\"subscriber_status\\\" : NumberInt(0),\n" + 
 "            \\\"access_restriction_data\\\" : NumberInt(32),\n" + 
 "            \\\"slice\\\" :\n" + 
 "            [{\n" + 
 "                \\\"sst\\\" : NumberInt($SST),\n" + 
 "                \\\"sd\\\" : \\\"$SD\\\",\n" + 
 "                \\\"default_indicator\\\" : true,\n" + 
 "                \\\"_id\\\" : new ObjectId(),\n" + 
 "                \\\"session\\\" :\n" + 
 "                [{\n" + 
 "                    \\\"name\\\" : \\\"$APN\\\",\n" + 
 "                    \\\"type\\\" : NumberInt(3),\n" + 
 "                    \\\"_id\\\" : new ObjectId(),\n" + 
 "                    \\\"pcc_rule\\\" : [],\n" + 
 "                    \\\"ipv4_framed_routes\\\": [\n" + 
 "                    \\\"192.168.20.0/24\\\"\n" + 
 "                    ],\n" + 
 "                    \\\"ambr\\\" :\n" + 
 "                    {\n" + 
 "                        \\\"uplink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                        \\\"downlink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "                    },\n" + 
 "                    \\\"qos\\\" :\n" + 
 "                    {\n" + 
 "                        \\\"index\\\" : NumberInt(9),\n" + 
 "                        \\\"arp\\\" :\n" + 
 "                        {\n" + 
 "                            \\\"priority_level\\\" : NumberInt(8),\n" + 
 "                            \\\"pre_emption_capability\\\" : NumberInt(1),\n" + 
 "                            \\\"pre_emption_vulnerability\\\" : NumberInt(1),\n" + 
 "                        },\n" + 
 "                    },\n" + 
 "                }],\n" + 
 "            }],\n" + 
 "            \\\"ambr\\\" :\n" + 
 "            {\n" + 
 "                \\\"uplink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3),},\n" + 
 "                \\\"downlink\\\" : { \\\"value\\\": NumberInt(1), \\\"unit\\\" : NumberInt(3) },\n" + 
 "            },\n" + 
 "            \\\"security\\\" :\n" + 
 "            {\n" + 
 "                \\\"k\\\" : \\\"$KI\\\",\n" + 
 "                \\\"amf\\\" : \\\"8000\\\",\n" + 
 "                \\\"op\\\" : null,\n" + 
 "                \\\"opc\\\" : \\\"$OPC\\\"\n" + 
 "            },\n" + 
 "            \\\"__v\\\" : 0\n" + 
 "        },\n" + 
 "    },\n" + 
 "    upsert=true);\" $DB_URI      \n" + 
 "",
		},
	}
	
		
	configMap10 := &corev1.ConfigMap{
		Data : map[string]string{
			"udm.yaml" : "logger:\n" + 
 "    file: /var/log/open5gs/udm.log\n" + 
 "\n" + 
 "udm:\n" + 
 "  sbi:\n" + 
 "    - addr: 0.0.0.0\n" + 
 "      advertise: release-name-udm.open5gscns.svc.cluster.local\n" + 
 "  discovery:\n" + 
 "     delegated: no           \n" + 
 "sbi:\n" + 
 "    server:\n" + 
 "      no_tls: true\n" + 
 "    client:\n" + 
 "      no_tls: true                \n" + 
 "scp:\n" + 
 " sbi:\n" + 
 "   name: release-name-scp.open5gscns.svc.cluster.local\n" + 
 "nrf:\n" + 
 " sbi:\n" + 
 "   name: release-name-nrf.open5gscns.svc.cluster.local\n" + 
 "",
		},
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "udm",
			},
			Name  : "release-name-udm-config", 
		}, 
		TypeMeta  : metav1.TypeMeta{
			Kind  : "ConfigMap", 
			APIVersion  : "v1", 
		}, 
	}
	
		
	configMap11 := &corev1.ConfigMap{
		Data : map[string]string{
			"udr.yaml" : "logger:\n" + 
 "    file: /var/log/open5gs/udr.log\n" + 
 "\n" + 
 "db_uri: mongodb://release-name-mongodb-svc/open5gs\n" + 
 "udr:\n" + 
 "  sbi:\n" + 
 "    - addr: 0.0.0.0\n" + 
 "      advertise: release-name-udr.open5gscns.svc.cluster.local\n" + 
 "  discovery:\n" + 
 "     delegated: no           \n" + 
 "scp:\n" + 
 " sbi:\n" + 
 "   name: release-name-scp.open5gscns.svc.cluster.local\n" + 
 "nrf:\n" + 
 " sbi:\n" + 
 "   name: release-name-nrf.open5gscns.svc.cluster.local \n" + 
 "sbi:\n" + 
 "    server:\n" + 
 "      no_tls: true\n" + 
 "    client:\n" + 
 "      no_tls: true\n" + 
 "",
		},
		ObjectMeta  : metav1.ObjectMeta{
			Name  : "release-name-udr-config", 
			Labels : map[string]string{
				"epc-mode" : "udr",
			},
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "ConfigMap", 
		}, 
	}
	
		
	return []*corev1.ConfigMap{configMap1, configMap2, configMap3, configMap4, configMap5, configMap6, configMap7, configMap8, configMap9, configMap10, configMap11, }
}
	
	
func GetService() []*corev1.Service{
	
	service1 := &corev1.Service{
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "amf",
			},
			Name  : "release-name-amf", 
		}, 
		Spec  : corev1.ServiceSpec{
			Ports  : []corev1.ServicePort{

				corev1.ServicePort{
					Name  : "sbi", 
					Port  : 80, 
					Protocol  : corev1.Protocol("TCP"), 
					TargetPort  : intstr.IntOrString{
						IntVal  : 80, 
					}, 
				},
				corev1.ServicePort{
					Port  : 38412, 
					Protocol  : corev1.Protocol("SCTP"), 
					TargetPort  : intstr.IntOrString{
						IntVal  : 38412, 
					}, 
					Name  : "ngap", 
				},

			}, 
			PublishNotReadyAddresses  : false, 
			Selector : map[string]string{
				"epc-mode" : "amf",
			},
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "Service", 
		}, 
	}
	
		
	service2 := &corev1.Service{
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "ausf",
			},
			Name  : "release-name-ausf", 
		}, 
		Spec  : corev1.ServiceSpec{
			PublishNotReadyAddresses  : false, 
			Selector : map[string]string{
				"epc-mode" : "ausf",
			},
			Ports  : []corev1.ServicePort{

				corev1.ServicePort{
					Name  : "sbi", 
					Port  : 80, 
					Protocol  : corev1.Protocol("TCP"), 
					TargetPort  : intstr.IntOrString{
						IntVal  : 80, 
					}, 
				},

			}, 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "Service", 
		}, 
	}
	
		
	service3 := &corev1.Service{
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "bsf",
			},
			Name  : "release-name-bsf", 
		}, 
		Spec  : corev1.ServiceSpec{
			PublishNotReadyAddresses  : false, 
			Selector : map[string]string{
				"epc-mode" : "bsf",
			},
			Ports  : []corev1.ServicePort{

				corev1.ServicePort{
					Protocol  : corev1.Protocol("TCP"), 
					TargetPort  : intstr.IntOrString{
						IntVal  : 80, 
					}, 
					Port  : 80, 
				},

			}, 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "Service", 
		}, 
	}
	
		
	service4 := &corev1.Service{
		ObjectMeta  : metav1.ObjectMeta{
			Name  : "release-name-mongodb-svc", 
		}, 
		Spec  : corev1.ServiceSpec{
			Selector : map[string]string{
				"app" : "open5gs-mongodb",
			},
			Ports  : []corev1.ServicePort{

				corev1.ServicePort{
					Port  : 27017, 
				},

			}, 
			PublishNotReadyAddresses  : false, 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "Service", 
		}, 
	}
	
		
	service5 := &corev1.Service{
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "nrf",
			},
			Name  : "release-name-nrf", 
		}, 
		Spec  : corev1.ServiceSpec{
			Ports  : []corev1.ServicePort{

				corev1.ServicePort{
					TargetPort  : intstr.IntOrString{
						IntVal  : 80, 
					}, 
					Port  : 80, 
					Protocol  : corev1.Protocol("TCP"), 
				},

			}, 
			PublishNotReadyAddresses  : false, 
			Selector : map[string]string{
				"epc-mode" : "nrf",
			},
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "Service", 
		}, 
	}
	
		
	service6 := &corev1.Service{
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "nssf",
			},
			Name  : "release-name-nssf", 
		}, 
		Spec  : corev1.ServiceSpec{
			Ports  : []corev1.ServicePort{

				corev1.ServicePort{
					Port  : 80, 
					Protocol  : corev1.Protocol("TCP"), 
					TargetPort  : intstr.IntOrString{
						IntVal  : 80, 
					}, 
				},

			}, 
			PublishNotReadyAddresses  : false, 
			Selector : map[string]string{
				"epc-mode" : "nssf",
			},
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "Service", 
		}, 
	}
	
		
	service7 := &corev1.Service{
		Spec  : corev1.ServiceSpec{
			Ports  : []corev1.ServicePort{

				corev1.ServicePort{
					Port  : 80, 
					Protocol  : corev1.Protocol("TCP"), 
					TargetPort  : intstr.IntOrString{
						IntVal  : 80, 
					}, 
				},

			}, 
			PublishNotReadyAddresses  : false, 
			Selector : map[string]string{
				"epc-mode" : "pcf",
			},
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "Service", 
		}, 
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "pcf",
			},
			Name  : "release-name-pcf", 
		}, 
	}
	
		
	service8 := &corev1.Service{
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "Service", 
		}, 
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "scp",
			},
			Name  : "release-name-scp", 
		}, 
		Spec  : corev1.ServiceSpec{
			Selector : map[string]string{
				"epc-mode" : "scp",
			},
			Ports  : []corev1.ServicePort{

				corev1.ServicePort{
					Port  : 80, 
					Protocol  : corev1.Protocol("TCP"), 
					TargetPort  : intstr.IntOrString{
						IntVal  : 80, 
					}, 
				},

			}, 
			PublishNotReadyAddresses  : false, 
		}, 
	}
	
		
	service9 := &corev1.Service{
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "smf",
			},
			Name  : "release-name-smf", 
		}, 
		Spec  : corev1.ServiceSpec{
			Ports  : []corev1.ServicePort{

				corev1.ServicePort{
					Port  : 80, 
					Protocol  : corev1.Protocol("TCP"), 
					TargetPort  : intstr.IntOrString{
						IntVal  : 80, 
					}, 
				},

			}, 
			PublishNotReadyAddresses  : false, 
			Selector : map[string]string{
				"epc-mode" : "smf",
			},
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "Service", 
		}, 
	}
	
		
	service10 := &corev1.Service{
		Spec  : corev1.ServiceSpec{
			PublishNotReadyAddresses  : false, 
			Selector : map[string]string{
				"epc-mode" : "udm",
			},
			Ports  : []corev1.ServicePort{

				corev1.ServicePort{
					Port  : 80, 
					Protocol  : corev1.Protocol("TCP"), 
					TargetPort  : intstr.IntOrString{
						IntVal  : 80, 
					}, 
				},

			}, 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "Service", 
		}, 
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "udm",
			},
			Name  : "release-name-udm", 
		}, 
	}
	
		
	service11 := &corev1.Service{
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "udr",
			},
			Name  : "release-name-udr", 
		}, 
		Spec  : corev1.ServiceSpec{
			Ports  : []corev1.ServicePort{

				corev1.ServicePort{
					TargetPort  : intstr.IntOrString{
						IntVal  : 80, 
					}, 
					Port  : 80, 
					Protocol  : corev1.Protocol("TCP"), 
				},

			}, 
			PublishNotReadyAddresses  : false, 
			Selector : map[string]string{
				"epc-mode" : "udr",
			},
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "Service", 
		}, 
	}
	
		
	service12 := &corev1.Service{
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "webui",
			},
			Name  : "release-name-webui", 
		}, 
		Spec  : corev1.ServiceSpec{
			Ports  : []corev1.ServicePort{

				corev1.ServicePort{
					Port  : 80, 
					TargetPort  : intstr.IntOrString{
						IntVal  : 3000, 
					}, 
				},

			}, 
			PublishNotReadyAddresses  : false, 
			Selector : map[string]string{
				"epc-mode" : "webui",
			},
			Type  : corev1.ServiceType("ClusterIP"), 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "Service", 
		}, 
	}
	
		
	return []*corev1.Service{service1, service2, service3, service4, service5, service6, service7, service8, service9, service10, service11, service12, }
}
	
	
func GetDeployment() []*appsv1.Deployment{
	
	deployment1 := &appsv1.Deployment{
		Spec  : appsv1.DeploymentSpec{
			Paused  : false, 
			Replicas  : int32Ptr(1), 
			Selector  : &metav1.LabelSelector{
				MatchLabels : map[string]string{
					"epc-mode" : "amf",
				},
			}, 
			Template  : corev1.PodTemplateSpec{
				ObjectMeta  : metav1.ObjectMeta{
					Annotations : map[string]string{
						"k8s.v1.cni.cncf.io/networks" : "[ { \"name\": \"release-name-bridge-nad\", \"interface\": \"n2\", \"ips\": [\"192.168.2.1/24\" ] } ]",
					},
					Labels : map[string]string{
						"epc-mode" : "amf",
						"epc-prom" : "enabled",
					},
				}, 
				Spec  : corev1.PodSpec{
					HostIPC  : false, 
					HostNetwork  : false, 
					HostPID  : false, 
					InitContainers  : []corev1.Container{

						corev1.Container{
							Name  : "wait-for-scp", 
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
							Args  : []string{

								"pod",
								"-lepc-mode=scp",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
						},
						corev1.Container{
							StdinOnce  : false, 
							TTY  : false, 
							Args  : []string{

								"pod",
								"-lepc-mode=nrf",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "wait-for-nrf", 
							Stdin  : false, 
						},

					}, 
					ServiceAccountName  : "release-name-k8s-wait-for", 
					Volumes  : []corev1.Volume{

						corev1.Volume{
							Name  : "release-name-amf-config", 
							VolumeSource  : corev1.VolumeSource{
								ConfigMap  : &corev1.ConfigMapVolumeSource{
									LocalObjectReference  : corev1.LocalObjectReference{
										Name  : "release-name-amf-config", 
									}, 
								}, 
							}, 
						},

					}, 
					Containers  : []corev1.Container{

						corev1.Container{
							Name  : "amf", 
							Ports  : []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort  : 9091, 
									Name  : "prom", 
									Protocol  : corev1.Protocol("TCP"), 
								},

							}, 
							Stdin  : false, 
							StdinOnce  : false, 
							Command  : []string{

								"/bin/sh",
								"-c",

							}, 
							Image  : "registry.gitlab.com/infinitydon/registry/open5gs-aio:v2.6.2", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							VolumeMounts  : []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath  : "/open5gs/config-map/amf.yaml", 
									Name  : "release-name-amf-config", 
									ReadOnly  : false, 
									SubPath  : "amf.yaml", 
								},

							}, 
							Args  : []string{

								"open5gs-amfd -c /open5gs/config-map/amf.yaml;",

							}, 
							SecurityContext  : &corev1.SecurityContext{
								Capabilities  : &corev1.Capabilities{
									Add  : []corev1.Capability{

										corev1.Capability("NET_ADMIN"),

									}, 
								}, 
							}, 
							TTY  : false, 
						},

					}, 
				}, 
			}, 
		}, 
		TypeMeta  : metav1.TypeMeta{
			Kind  : "Deployment", 
			APIVersion  : "apps/v1", 
		}, 
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "amf",
			},
			Name  : "release-name-amf-deployment", 
		}, 
	}
	
		
	deployment2 := &appsv1.Deployment{
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "ausf",
			},
			Name  : "release-name-ausf-deployment", 
		}, 
		Spec  : appsv1.DeploymentSpec{
			Paused  : false, 
			Replicas  : int32Ptr(1), 
			Selector  : &metav1.LabelSelector{
				MatchLabels : map[string]string{
					"epc-mode" : "ausf",
				},
			}, 
			Template  : corev1.PodTemplateSpec{
				ObjectMeta  : metav1.ObjectMeta{
					Labels : map[string]string{
						"epc-mode" : "ausf",
					},
				}, 
				Spec  : corev1.PodSpec{
					Volumes  : []corev1.Volume{

						corev1.Volume{
							Name  : "release-name-ausf-config", 
							VolumeSource  : corev1.VolumeSource{
								ConfigMap  : &corev1.ConfigMapVolumeSource{
									LocalObjectReference  : corev1.LocalObjectReference{
										Name  : "release-name-ausf-config", 
									}, 
								}, 
							}, 
						},

					}, 
					Containers  : []corev1.Container{

						corev1.Container{
							Image  : "registry.gitlab.com/infinitydon/registry/open5gs-aio:v2.6.2", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "ausf", 
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
							VolumeMounts  : []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath  : "/open5gs/config-map/ausf.yaml", 
									Name  : "release-name-ausf-config", 
									ReadOnly  : false, 
									SubPath  : "ausf.yaml", 
								},

							}, 
							Command  : []string{

								"open5gs-ausfd",
								"-c",
								"/open5gs/config-map/ausf.yaml",

							}, 
						},

					}, 
					HostIPC  : false, 
					HostNetwork  : false, 
					HostPID  : false, 
					InitContainers  : []corev1.Container{

						corev1.Container{
							TTY  : false, 
							Args  : []string{

								"pod",
								"-lepc-mode=scp",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "wait-for-scp", 
							Stdin  : false, 
							StdinOnce  : false, 
						},
						corev1.Container{
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
							Args  : []string{

								"pod",
								"-lepc-mode=nrf",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "wait-for-nrf", 
						},

					}, 
					ServiceAccountName  : "release-name-k8s-wait-for", 
				}, 
			}, 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "apps/v1", 
			Kind  : "Deployment", 
		}, 
	}
	
		
	deployment3 := &appsv1.Deployment{
		Spec  : appsv1.DeploymentSpec{
			Paused  : false, 
			Replicas  : int32Ptr(1), 
			Selector  : &metav1.LabelSelector{
				MatchLabels : map[string]string{
					"epc-mode" : "bsf",
				},
			}, 
			Template  : corev1.PodTemplateSpec{
				ObjectMeta  : metav1.ObjectMeta{
					Labels : map[string]string{
						"epc-mode" : "bsf",
					},
				}, 
				Spec  : corev1.PodSpec{
					HostIPC  : false, 
					HostNetwork  : false, 
					HostPID  : false, 
					InitContainers  : []corev1.Container{

						corev1.Container{
							Args  : []string{

								"pod",
								"-lepc-mode=scp",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "wait-for-scp", 
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
						},
						corev1.Container{
							Args  : []string{

								"pod",
								"-lepc-mode=nrf",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "wait-for-nrf", 
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
						},

					}, 
					ServiceAccountName  : "release-name-k8s-wait-for", 
					Volumes  : []corev1.Volume{

						corev1.Volume{
							Name  : "release-name-bsf-config", 
							VolumeSource  : corev1.VolumeSource{
								ConfigMap  : &corev1.ConfigMapVolumeSource{
									LocalObjectReference  : corev1.LocalObjectReference{
										Name  : "release-name-bsf-config", 
									}, 
								}, 
							}, 
						},

					}, 
					Containers  : []corev1.Container{

						corev1.Container{
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "bsf", 
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
							VolumeMounts  : []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath  : "/open5gs/config-map/bsf.yaml", 
									Name  : "release-name-bsf-config", 
									ReadOnly  : false, 
									SubPath  : "bsf.yaml", 
								},

							}, 
							Command  : []string{

								"open5gs-bsfd",
								"-c",
								"/open5gs/config-map/bsf.yaml",

							}, 
							Image  : "registry.gitlab.com/infinitydon/registry/open5gs-aio:v2.6.2", 
						},

					}, 
				}, 
			}, 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "apps/v1", 
			Kind  : "Deployment", 
		}, 
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "bsf",
			},
			Name  : "release-name-bsf-deployment", 
		}, 
	}
	
		
	deployment4 := &appsv1.Deployment{
		ObjectMeta  : metav1.ObjectMeta{
			Name  : "release-name-mongodb", 
		}, 
		Spec  : appsv1.DeploymentSpec{
			Selector  : &metav1.LabelSelector{
				MatchLabels : map[string]string{
					"app" : "open5gs-mongodb",
				},
			}, 
			Template  : corev1.PodTemplateSpec{
				ObjectMeta  : metav1.ObjectMeta{
					Labels : map[string]string{
						"app" : "open5gs-mongodb",
					},
				}, 
				Spec  : corev1.PodSpec{
					Volumes  : []corev1.Volume{

						corev1.Volume{
							Name  : "mongodb-persistent-storage", 
						},

					}, 
					Containers  : []corev1.Container{

						corev1.Container{
							Name  : "open5gs-mongodb", 
							Ports  : []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort  : 27017, 
									Name  : "mongodb", 
								},

							}, 
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
							VolumeMounts  : []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath  : "/data/db", 
									Name  : "mongodb-persistent-storage", 
									ReadOnly  : false, 
								},

							}, 
							Image  : "mongo:6.0.1", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
						},

					}, 
					HostIPC  : false, 
					HostNetwork  : false, 
					HostPID  : false, 
				}, 
			}, 
			Paused  : false, 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "apps/v1", 
			Kind  : "Deployment", 
		}, 
	}
	
		
	deployment5 := &appsv1.Deployment{
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "apps/v1", 
			Kind  : "Deployment", 
		}, 
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "nrf",
			},
			Name  : "release-name-nrf-deployment", 
		}, 
		Spec  : appsv1.DeploymentSpec{
			Paused  : false, 
			Replicas  : int32Ptr(1), 
			Selector  : &metav1.LabelSelector{
				MatchLabels : map[string]string{
					"epc-mode" : "nrf",
				},
			}, 
			Template  : corev1.PodTemplateSpec{
				ObjectMeta  : metav1.ObjectMeta{
					Labels : map[string]string{
						"epc-mode" : "nrf",
					},
				}, 
				Spec  : corev1.PodSpec{
					HostPID  : false, 
					Volumes  : []corev1.Volume{

						corev1.Volume{
							Name  : "release-name-nrf-config", 
							VolumeSource  : corev1.VolumeSource{
								ConfigMap  : &corev1.ConfigMapVolumeSource{
									LocalObjectReference  : corev1.LocalObjectReference{
										Name  : "release-name-nrf-config", 
									}, 
								}, 
							}, 
						},

					}, 
					Containers  : []corev1.Container{

						corev1.Container{
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "nrf", 
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
							VolumeMounts  : []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath  : "/open5gs/config-map/nrf.yaml", 
									Name  : "release-name-nrf-config", 
									ReadOnly  : false, 
									SubPath  : "nrf.yaml", 
								},

							}, 
							Command  : []string{

								"open5gs-nrfd",
								"-c",
								"/open5gs/config-map/nrf.yaml",

							}, 
							Image  : "registry.gitlab.com/infinitydon/registry/open5gs-aio:v2.6.2", 
						},

					}, 
					HostIPC  : false, 
					HostNetwork  : false, 
				}, 
			}, 
		}, 
	}
	
		
	deployment6 := &appsv1.Deployment{
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "nssf",
			},
			Name  : "release-name-nssf-deployment", 
		}, 
		Spec  : appsv1.DeploymentSpec{
			Selector  : &metav1.LabelSelector{
				MatchLabels : map[string]string{
					"epc-mode" : "nssf",
				},
			}, 
			Template  : corev1.PodTemplateSpec{
				ObjectMeta  : metav1.ObjectMeta{
					Labels : map[string]string{
						"epc-mode" : "nssf",
					},
				}, 
				Spec  : corev1.PodSpec{
					HostNetwork  : false, 
					HostPID  : false, 
					InitContainers  : []corev1.Container{

						corev1.Container{
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
							Args  : []string{

								"pod",
								"-lepc-mode=scp",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "wait-for-scp", 
						},
						corev1.Container{
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
							Args  : []string{

								"pod",
								"-lepc-mode=nrf",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "wait-for-nrf", 
						},

					}, 
					ServiceAccountName  : "release-name-k8s-wait-for", 
					Volumes  : []corev1.Volume{

						corev1.Volume{
							Name  : "release-name-nssf-config", 
							VolumeSource  : corev1.VolumeSource{
								ConfigMap  : &corev1.ConfigMapVolumeSource{
									LocalObjectReference  : corev1.LocalObjectReference{
										Name  : "release-name-nssf-config", 
									}, 
								}, 
							}, 
						},

					}, 
					Containers  : []corev1.Container{

						corev1.Container{
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "nssf", 
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
							VolumeMounts  : []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath  : "/open5gs/config-map/nssf.yaml", 
									Name  : "release-name-nssf-config", 
									ReadOnly  : false, 
									SubPath  : "nssf.yaml", 
								},

							}, 
							Command  : []string{

								"open5gs-nssfd",
								"-c",
								"/open5gs/config-map/nssf.yaml",

							}, 
							Image  : "registry.gitlab.com/infinitydon/registry/open5gs-aio:v2.6.2", 
						},

					}, 
					HostIPC  : false, 
				}, 
			}, 
			Paused  : false, 
			Replicas  : int32Ptr(1), 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "apps/v1", 
			Kind  : "Deployment", 
		}, 
	}
	
		
	deployment7 := &appsv1.Deployment{
		Spec  : appsv1.DeploymentSpec{
			Paused  : false, 
			Replicas  : int32Ptr(1), 
			Selector  : &metav1.LabelSelector{
				MatchLabels : map[string]string{
					"epc-mode" : "pcf",
				},
			}, 
			Template  : corev1.PodTemplateSpec{
				Spec  : corev1.PodSpec{
					Volumes  : []corev1.Volume{

						corev1.Volume{
							Name  : "release-name-pcf-config", 
							VolumeSource  : corev1.VolumeSource{
								ConfigMap  : &corev1.ConfigMapVolumeSource{
									LocalObjectReference  : corev1.LocalObjectReference{
										Name  : "release-name-pcf-config", 
									}, 
								}, 
							}, 
						},

					}, 
					Containers  : []corev1.Container{

						corev1.Container{
							Command  : []string{

								"open5gs-pcfd",
								"-c",
								"/open5gs/config-map/pcf.yaml",

							}, 
							Image  : "registry.gitlab.com/infinitydon/registry/open5gs-aio:v2.6.2", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "pcf", 
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
							VolumeMounts  : []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath  : "/open5gs/config-map/pcf.yaml", 
									Name  : "release-name-pcf-config", 
									ReadOnly  : false, 
									SubPath  : "pcf.yaml", 
								},

							}, 
						},

					}, 
					HostIPC  : false, 
					HostNetwork  : false, 
					HostPID  : false, 
					InitContainers  : []corev1.Container{

						corev1.Container{
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "wait-for-scp", 
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
							Args  : []string{

								"pod",
								"-lepc-mode=scp",

							}, 
						},
						corev1.Container{
							Args  : []string{

								"pod",
								"-lapp=open5gs-mongodb",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "wait-for-mongo", 
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
						},
						corev1.Container{
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
							Args  : []string{

								"pod",
								"-lepc-mode=nrf",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "wait-for-nrf", 
						},

					}, 
					ServiceAccountName  : "release-name-k8s-wait-for", 
				}, 
				ObjectMeta  : metav1.ObjectMeta{
					Labels : map[string]string{
						"epc-mode" : "pcf",
					},
				}, 
			}, 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "apps/v1", 
			Kind  : "Deployment", 
		}, 
		ObjectMeta  : metav1.ObjectMeta{
			Name  : "release-name-pcf-deployment", 
			Labels : map[string]string{
				"epc-mode" : "pcf",
			},
		}, 
	}
	
		
	deployment8 := &appsv1.Deployment{
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "scp",
			},
			Name  : "release-name-scp-deployment", 
		}, 
		Spec  : appsv1.DeploymentSpec{
			Replicas  : int32Ptr(1), 
			Selector  : &metav1.LabelSelector{
				MatchLabels : map[string]string{
					"epc-mode" : "scp",
				},
			}, 
			Template  : corev1.PodTemplateSpec{
				ObjectMeta  : metav1.ObjectMeta{
					Labels : map[string]string{
						"epc-mode" : "scp",
					},
				}, 
				Spec  : corev1.PodSpec{
					HostNetwork  : false, 
					HostPID  : false, 
					InitContainers  : []corev1.Container{

						corev1.Container{
							StdinOnce  : false, 
							TTY  : false, 
							Args  : []string{

								"pod",
								"-lepc-mode=nrf",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "wait-for-nrf", 
							Stdin  : false, 
						},
						corev1.Container{
							TTY  : false, 
							Args  : []string{

								"pod",
								"-lapp=open5gs-mongodb",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "wait-for-mongo", 
							Stdin  : false, 
							StdinOnce  : false, 
						},

					}, 
					ServiceAccountName  : "release-name-k8s-wait-for", 
					Volumes  : []corev1.Volume{

						corev1.Volume{
							Name  : "release-name-scp-config", 
							VolumeSource  : corev1.VolumeSource{
								ConfigMap  : &corev1.ConfigMapVolumeSource{
									LocalObjectReference  : corev1.LocalObjectReference{
										Name  : "release-name-scp-config", 
									}, 
								}, 
							}, 
						},

					}, 
					Containers  : []corev1.Container{

						corev1.Container{
							TTY  : false, 
							VolumeMounts  : []corev1.VolumeMount{

								corev1.VolumeMount{
									ReadOnly  : false, 
									SubPath  : "scp.yaml", 
									MountPath  : "/open5gs/config-map/scp.yaml", 
									Name  : "release-name-scp-config", 
								},

							}, 
							Command  : []string{

								"open5gs-scpd",
								"-c",
								"/open5gs/config-map/scp.yaml",

							}, 
							Image  : "registry.gitlab.com/infinitydon/registry/open5gs-aio:v2.6.2", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "scp", 
							Stdin  : false, 
							StdinOnce  : false, 
						},

					}, 
					HostIPC  : false, 
				}, 
			}, 
			Paused  : false, 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "apps/v1", 
			Kind  : "Deployment", 
		}, 
	}
	
		
	deployment9 := &appsv1.Deployment{
		ObjectMeta  : metav1.ObjectMeta{
			Name  : "release-name-smf-deployment", 
			Labels : map[string]string{
				"epc-mode" : "smf",
			},
		}, 
		Spec  : appsv1.DeploymentSpec{
			Selector  : &metav1.LabelSelector{
				MatchLabels : map[string]string{
					"epc-mode" : "smf",
				},
			}, 
			Template  : corev1.PodTemplateSpec{
				ObjectMeta  : metav1.ObjectMeta{
					Labels : map[string]string{
						"epc-mode" : "smf",
						"epc-prom" : "enabled",
					},
					Annotations : map[string]string{
						"k8s.v1.cni.cncf.io/networks" : "[ { \"name\": \"release-name-bridge-nad\", \"interface\": \"n4\", \"ips\": [\"192.168.4.1/24\" ] } ]",
					},
				}, 
				Spec  : corev1.PodSpec{
					HostPID  : false, 
					InitContainers  : []corev1.Container{

						corev1.Container{
							Args  : []string{

								"pod",
								"-lepc-mode=scp",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "wait-for-scp", 
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
						},
						corev1.Container{
							Args  : []string{

								"pod",
								"-lepc-mode=nrf",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "wait-for-nrf", 
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
						},

					}, 
					ServiceAccountName  : "release-name-k8s-wait-for", 
					Volumes  : []corev1.Volume{

						corev1.Volume{
							Name  : "release-name-smf-config", 
							VolumeSource  : corev1.VolumeSource{
								ConfigMap  : &corev1.ConfigMapVolumeSource{
									LocalObjectReference  : corev1.LocalObjectReference{
										Name  : "release-name-smf-config", 
									}, 
								}, 
							}, 
						},

					}, 
					Containers  : []corev1.Container{

						corev1.Container{
							Ports  : []corev1.ContainerPort{

								corev1.ContainerPort{
									ContainerPort  : 9091, 
									Name  : "prom", 
									Protocol  : corev1.Protocol("TCP"), 
								},

							}, 
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
							VolumeMounts  : []corev1.VolumeMount{

								corev1.VolumeMount{
									MountPath  : "/open5gs/config-map/smf.yaml", 
									Name  : "release-name-smf-config", 
									ReadOnly  : false, 
									SubPath  : "smf.yaml", 
								},

							}, 
							Command  : []string{

								"/bin/sh",
								"-c",

							}, 
							Name  : "smf", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							SecurityContext  : &corev1.SecurityContext{
								Capabilities  : &corev1.Capabilities{
									Add  : []corev1.Capability{

										corev1.Capability("NET_ADMIN"),

									}, 
								}, 
							}, 
							Args  : []string{

								"open5gs-smfd -c /open5gs/config-map/smf.yaml;",

							}, 
							Image  : "registry.gitlab.com/infinitydon/registry/open5gs-aio:v2.6.2", 
						},

					}, 
					HostIPC  : false, 
					HostNetwork  : false, 
				}, 
			}, 
			Paused  : false, 
			Replicas  : int32Ptr(1), 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "apps/v1", 
			Kind  : "Deployment", 
		}, 
	}
	
		
	deployment10 := &appsv1.Deployment{
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "ue-import",
			},
			Name  : "release-name-ue-import", 
		}, 
		Spec  : appsv1.DeploymentSpec{
			Paused  : false, 
			Replicas  : int32Ptr(1), 
			Selector  : &metav1.LabelSelector{
				MatchLabels : map[string]string{
					"epc-mode" : "ue-import",
				},
			}, 
			Template  : corev1.PodTemplateSpec{
				Spec  : corev1.PodSpec{
					HostIPC  : false, 
					HostNetwork  : false, 
					HostPID  : false, 
					InitContainers  : []corev1.Container{

						corev1.Container{
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
							Args  : []string{

								"pod",
								"-lapp=open5gs-mongodb",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "wait-for-mongo", 
						},

					}, 
					ServiceAccountName  : "release-name-k8s-wait-for", 
					Volumes  : []corev1.Volume{

						corev1.Volume{
							Name  : "account-config", 
							VolumeSource  : corev1.VolumeSource{
								ConfigMap  : &corev1.ConfigMapVolumeSource{
									LocalObjectReference  : corev1.LocalObjectReference{
										Name  : "release-name-ue-provisioning", 
									}, 
								}, 
							}, 
						},
						corev1.Volume{
							Name  : "db-script", 
							VolumeSource  : corev1.VolumeSource{
								ConfigMap  : &corev1.ConfigMapVolumeSource{
									LocalObjectReference  : corev1.LocalObjectReference{
										Name  : "release-name-ue-provisioning", 
									}, 
								}, 
							}, 
						},
						corev1.Volume{
							Name  : "frame-route-1", 
							VolumeSource  : corev1.VolumeSource{
								ConfigMap  : &corev1.ConfigMapVolumeSource{
									LocalObjectReference  : corev1.LocalObjectReference{
										Name  : "release-name-ue-provisioning", 
									}, 
								}, 
							}, 
						},
						corev1.Volume{
							Name  : "frame-route-2", 
							VolumeSource  : corev1.VolumeSource{
								ConfigMap  : &corev1.ConfigMapVolumeSource{
									LocalObjectReference  : corev1.LocalObjectReference{
										Name  : "release-name-ue-provisioning", 
									}, 
								}, 
							}, 
						},

					}, 
					Containers  : []corev1.Container{

						corev1.Container{
							Args  : []string{

								"mongo mongodb://release-name-mongodb-svc/open5gs /tmp/account.js; cp -rf /tmp/open5gs-dbctl /usr/local/bin/open5gs-dbctl; chmod +x /usr/local/bin/open5gs-dbctl; bash -x /tmp/ue.sh; sleep infinity;",

							}, 
							Command  : []string{

								"/bin/sh",
								"-c",

							}, 
							Image  : "free5gmano/nextepc-mongodb:latest", 
							Stdin  : false, 
							StdinOnce  : false, 
							Env  : []corev1.EnvVar{

								corev1.EnvVar{
									Name  : "DB_URI", 
									Value  : "mongodb://release-name-mongodb-svc/open5gs", 
								},

							}, 
							Name  : "mongo", 
							TTY  : false, 
							VolumeMounts  : []corev1.VolumeMount{

								corev1.VolumeMount{
									Name  : "account-config", 
									ReadOnly  : false, 
									SubPath  : "account.js", 
									MountPath  : "/tmp/account.js", 
								},
								corev1.VolumeMount{
									MountPath  : "/tmp/open5gs-dbctl", 
									Name  : "db-script", 
									ReadOnly  : false, 
									SubPath  : "open5gs-dbctl", 
								},
								corev1.VolumeMount{
									MountPath  : "/tmp/ue.sh", 
									Name  : "frame-route-1", 
									ReadOnly  : false, 
									SubPath  : "ue.sh", 
								},

							}, 
						},

					}, 
				}, 
				ObjectMeta  : metav1.ObjectMeta{
					Labels : map[string]string{
						"epc-mode" : "ue-import",
					},
				}, 
			}, 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "apps/v1", 
			Kind  : "Deployment", 
		}, 
	}
	
		
	deployment11 := &appsv1.Deployment{
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "udm",
			},
			Name  : "release-name-udm-deployment", 
		}, 
		Spec  : appsv1.DeploymentSpec{
			Paused  : false, 
			Replicas  : int32Ptr(1), 
			Selector  : &metav1.LabelSelector{
				MatchLabels : map[string]string{
					"epc-mode" : "udm",
				},
			}, 
			Template  : corev1.PodTemplateSpec{
				ObjectMeta  : metav1.ObjectMeta{
					Labels : map[string]string{
						"epc-mode" : "udm",
					},
				}, 
				Spec  : corev1.PodSpec{
					HostPID  : false, 
					InitContainers  : []corev1.Container{

						corev1.Container{
							Args  : []string{

								"pod",
								"-lepc-mode=scp",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "wait-for-scp", 
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
						},
						corev1.Container{
							Name  : "wait-for-nrf", 
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
							Args  : []string{

								"pod",
								"-lepc-mode=nrf",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
						},

					}, 
					ServiceAccountName  : "release-name-k8s-wait-for", 
					Volumes  : []corev1.Volume{

						corev1.Volume{
							Name  : "release-name-udm-config", 
							VolumeSource  : corev1.VolumeSource{
								ConfigMap  : &corev1.ConfigMapVolumeSource{
									LocalObjectReference  : corev1.LocalObjectReference{
										Name  : "release-name-udm-config", 
									}, 
								}, 
							}, 
						},

					}, 
					Containers  : []corev1.Container{

						corev1.Container{
							VolumeMounts  : []corev1.VolumeMount{

								corev1.VolumeMount{
									ReadOnly  : false, 
									SubPath  : "udm.yaml", 
									MountPath  : "/open5gs/config-map/udm.yaml", 
									Name  : "release-name-udm-config", 
								},

							}, 
							Command  : []string{

								"open5gs-udmd",
								"-c",
								"/open5gs/config-map/udm.yaml",

							}, 
							Image  : "registry.gitlab.com/infinitydon/registry/open5gs-aio:v2.6.2", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "udm", 
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
						},

					}, 
					HostIPC  : false, 
					HostNetwork  : false, 
				}, 
			}, 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "apps/v1", 
			Kind  : "Deployment", 
		}, 
	}
	
		
	deployment12 := &appsv1.Deployment{
		ObjectMeta  : metav1.ObjectMeta{
			Name  : "release-name-udr-deployment", 
			Labels : map[string]string{
				"epc-mode" : "udr",
			},
		}, 
		Spec  : appsv1.DeploymentSpec{
			Paused  : false, 
			Replicas  : int32Ptr(1), 
			Selector  : &metav1.LabelSelector{
				MatchLabels : map[string]string{
					"epc-mode" : "udr",
				},
			}, 
			Template  : corev1.PodTemplateSpec{
				ObjectMeta  : metav1.ObjectMeta{
					Labels : map[string]string{
						"epc-mode" : "udr",
					},
				}, 
				Spec  : corev1.PodSpec{
					InitContainers  : []corev1.Container{

						corev1.Container{
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
							Args  : []string{

								"pod",
								"-lepc-mode=scp",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "wait-for-scp", 
						},
						corev1.Container{
							Name  : "wait-for-mongo", 
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
							Args  : []string{

								"pod",
								"-lapp=open5gs-mongodb",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
						},
						corev1.Container{
							Name  : "wait-for-nrf", 
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
							Args  : []string{

								"pod",
								"-lepc-mode=nrf",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
						},

					}, 
					ServiceAccountName  : "release-name-k8s-wait-for", 
					Volumes  : []corev1.Volume{

						corev1.Volume{
							Name  : "release-name-udr-config", 
							VolumeSource  : corev1.VolumeSource{
								ConfigMap  : &corev1.ConfigMapVolumeSource{
									LocalObjectReference  : corev1.LocalObjectReference{
										Name  : "release-name-udr-config", 
									}, 
								}, 
							}, 
						},

					}, 
					Containers  : []corev1.Container{

						corev1.Container{
							Image  : "registry.gitlab.com/infinitydon/registry/open5gs-aio:v2.6.2", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "udr", 
							Stdin  : false, 
							StdinOnce  : false, 
							TTY  : false, 
							VolumeMounts  : []corev1.VolumeMount{

								corev1.VolumeMount{
									ReadOnly  : false, 
									SubPath  : "udr.yaml", 
									MountPath  : "/open5gs/config-map/udr.yaml", 
									Name  : "release-name-udr-config", 
								},

							}, 
							Command  : []string{

								"open5gs-udrd",
								"-c",
								"/open5gs/config-map/udr.yaml",

							}, 
						},

					}, 
					HostIPC  : false, 
					HostNetwork  : false, 
					HostPID  : false, 
				}, 
			}, 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "apps/v1", 
			Kind  : "Deployment", 
		}, 
	}
	
		
	deployment13 := &appsv1.Deployment{
		Spec  : appsv1.DeploymentSpec{
			Template  : corev1.PodTemplateSpec{
				ObjectMeta  : metav1.ObjectMeta{
					Labels : map[string]string{
						"epc-mode" : "webui",
					},
				}, 
				Spec  : corev1.PodSpec{
					HostNetwork  : false, 
					HostPID  : false, 
					InitContainers  : []corev1.Container{

						corev1.Container{
							StdinOnce  : false, 
							TTY  : false, 
							Args  : []string{

								"pod",
								"-lapp=open5gs-mongodb",

							}, 
							Image  : "groundnuty/k8s-wait-for:v1.6", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "wait-for-mongo", 
							Stdin  : false, 
						},

					}, 
					ServiceAccountName  : "release-name-k8s-wait-for", 
					Containers  : []corev1.Container{

						corev1.Container{
							TTY  : false, 
							Env  : []corev1.EnvVar{

								corev1.EnvVar{
									Name  : "DB_URI", 
									Value  : "mongodb://release-name-mongodb-svc/open5gs", 
								},
								corev1.EnvVar{
									Value  : "production", 
									Name  : "NODE_ENV", 
								},
								corev1.EnvVar{
									Name  : "HOSTNAME", 
									Value  : "0.0.0.0", 
								},

							}, 
							Image  : "registry.gitlab.com/infinitydon/registry/open5gs-webui:v2.6.2", 
							ImagePullPolicy  : corev1.PullPolicy("IfNotPresent"), 
							Name  : "webui", 
							Stdin  : false, 
							StdinOnce  : false, 
						},

					}, 
					HostIPC  : false, 
				}, 
			}, 
			Paused  : false, 
			Replicas  : int32Ptr(1), 
			Selector  : &metav1.LabelSelector{
				MatchLabels : map[string]string{
					"epc-mode" : "webui",
				},
			}, 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "apps/v1", 
			Kind  : "Deployment", 
		}, 
		ObjectMeta  : metav1.ObjectMeta{
			Labels : map[string]string{
				"epc-mode" : "webui",
			},
			Name  : "release-name-webui", 
		}, 
	}
	
		
	return []*appsv1.Deployment{deployment1, deployment2, deployment3, deployment4, deployment5, deployment6, deployment7, deployment8, deployment9, deployment10, deployment11, deployment12, deployment13, }
}
	
	
func GetRoleBinding() []*rbacv1.RoleBinding{
	
	roleBinding1 := &rbacv1.RoleBinding{
		ObjectMeta  : metav1.ObjectMeta{
			Name  : "release-name-k8s-wait-for", 
		}, 
		RoleRef  : rbacv1.RoleRef{
			APIGroup  : "rbac.authorization.k8s.io", 
			Kind  : "Role", 
			Name  : "release-name-k8s-wait-for", 
		}, 
		Subjects  : []rbacv1.Subject{

			rbacv1.Subject{
				Kind  : "ServiceAccount", 
				Name  : "release-name-k8s-wait-for", 
				Namespace  : "open5gscns", 
			},

		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "rbac.authorization.k8s.io/v1", 
			Kind  : "RoleBinding", 
		}, 
	}
	
		
	return []*rbacv1.RoleBinding{roleBinding1, }
}
	
	
func GetRole() []*rbacv1.Role{
	
	role1 := &rbacv1.Role{
		ObjectMeta  : metav1.ObjectMeta{
			Name  : "release-name-k8s-wait-for", 
		}, 
		Rules  : []rbacv1.PolicyRule{

			rbacv1.PolicyRule{
				APIGroups  : []string{

					"",

				}, 
				Resources  : []string{

					"pods",

				}, 
				Verbs  : []string{

					"get",
					"watch",
					"list",

				}, 
			},

		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "rbac.authorization.k8s.io/v1", 
			Kind  : "Role", 
		}, 
	}
	
		
	return []*rbacv1.Role{role1, }
}
	
	
func GetServiceAccount() []*corev1.ServiceAccount{
	
	serviceAccount1 := &corev1.ServiceAccount{
		ObjectMeta  : metav1.ObjectMeta{
			Name  : "release-name-k8s-wait-for", 
			Namespace  : "open5gscns", 
		}, 
		TypeMeta  : metav1.TypeMeta{
			APIVersion  : "v1", 
			Kind  : "ServiceAccount", 
		}, 
	}
	
		
	return []*corev1.ServiceAccount{serviceAccount1, }
}
	
	
func GetNetworkAttachmentDefinition() []*unstructured.Unstructured{
	
	networkAttachmentDefinition1 := &unstructured.Unstructured{
	Object: map[string]any{
		"apiVersion": "k8s.cni.cncf.io/v1",
		"kind": "NetworkAttachmentDefinition",
		"metadata": map[string]any{
			"name": "release-name-bridge-nad",
			},
		"spec": map[string]any{
			"config": "{ \"cniVersion\": \"0.3.1\", \"plugins\": [ { \"type\": \"macvlan\", \"capabilities\": { \"ips\": true }, \"master\": \"eth1\", \"mode\": \"bridge\", \"ipam\": { \"type\": \"static\" } }] }",
			},
		},
	}
	
		
	return []*unstructured.Unstructured{networkAttachmentDefinition1, }
}
	
	