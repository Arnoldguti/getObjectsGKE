package main

import (

	"fmt"
	 v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)


type Service struct {
	Name      string
	Port     int32
	ClusterIP  string
	App       string
}

type Srvs []Service
var services_type Srvs


//New declaration for slide
type Data struct {
	Name        string       //`json:"id"`
	ClusterIP  string       //`json:"lastname"`
	App  string             //`json:"username"`
}

type Data2 struct {
	Name        string       //`json:"id"`
	Containers  int       //`json:"lastname"`
	Reason  string             //`json:"username"`
	Message  string
}

type Datas []Data

func main() {


	//API for objects
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/pods", GetPods)
	router.HandleFunc("/services", Services)

	log.Fatal(http.ListenAndServe(":6060", router))



}


//Get pods from GKE
func GetPods(w http.ResponseWriter, r *http.Request) {
	var (
		//	user  Data
	users2 []Data2

	)


	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Fprintln(w, "ha habido un error")
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Fprintln(w, "ha habido un error")
	}

	pods, err := clientset.Core().Pods("").List(v1.ListOptions{})
	if err != nil {
		//panic(err.Error())
		fmt.Fprintln(w, "ha habido un error")
	}

   //Pods Objects
	for i := 0; i < len(pods.Items); i++ {

		users2 = append(users2, Data2{Name: pods.Items[i].Name , Containers:len(pods.Items[i].Spec.Containers) , Reason:pods.Items[i].Status.Reason,
			Message: string(pods.Items[i].Status.Phase)})
	}

    //JSON Response
	json.NewEncoder(w).Encode(users2)


}


//Get service from GKE
func Services(w http.ResponseWriter, r *http.Request) {

	config, err := rest.InClusterConfig()
	if err != nil {
		fmt.Fprintln(w, "ha habido un error")
	}
	//Clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Fprintln(w, "ha habido un error")
	}

	services, err := clientset.Core().Services("").List(v1.ListOptions{})
	if err != nil {
		//panic(err.Error())
		fmt.Fprintln(w, "ha habido un error")
	}

	for _, s := range services.Items {

			services_type = append(services_type, Service{Name: s.Name, Port: s.Spec.Ports[1].Port , ClusterIP:s.Spec.ClusterIP , App:s.Spec.ExternalName})

	}

	//JSON Response
	json.NewEncoder(w).Encode(services_type)


}

