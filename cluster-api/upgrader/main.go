package main

import (
	"flag"
	"fmt"
	"time"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	clusapiclnt "k8s.io/kube-deploy/cluster-api/client"
	"k8s.io/apimachinery/pkg/util/wait"
	machinesv1 "k8s.io/kube-deploy/cluster-api/api/machines/v1alpha1"
)

var (
	kubeconfig = flag.String("kubeconfig", "", "path to kubeconfig file")
	kubeversion = flag.String("version", "", "The kubernetes version to be upgraded to")
)

func main() {
	flag.Parse()

	if *kubeversion == "" {
		fmt.Errorf("You need to specify kubernetes version being upgraded to.")
	}

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	client, err := clusapiclnt.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// Fetch Cluster object.
	clusters, err := client.Clusters().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	} else if len(clusters.Items) != 1 {
		panic("There is zero or more than one cluster returned!")
	}

	// Modify the control plane version
	clusters.Items[0].Spec.KubernetesVersion.Version = *kubeversion

	// Update the cluster.
	cluster, err := client.Clusters().Update(&clusters.Items[0])
	if err != nil {
		panic(err.Error())
	}

	// Polling the cluster until new control plan is ready.
	err = wait.Poll(5*time.Second, 10*time.Minute, func() (bool, error) {
		cluster, err = client.Clusters().Get(cluster.Name, metav1.GetOptions{})
		//if err == nil && cluster.Status.Ready {
		//	return true, err
		//}
		//return false, err
		return true, nil
	})
	if err != nil {
		fmt.Println("Failed to upgraded control plan. Error = %v", err)
		return
	} else {
		fmt.Println("Successfully upgraded control plan.")
	}

	// Now continue to update all the node's state.
	machine_list, err := client.Machines().List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// Polling the cluster until nodes are updated.
	errors := make(chan error, len(machine_list.Items))
	for i, _ := range machine_list.Items {
		go func(mach *machinesv1.Machine) {
			mach.Spec.Versions.Kubelet = *kubeversion
			fmt.Println("Updating machine:", mach.Name)
			new_machine, err := client.Machines().Update(mach)
			if err == nil {
				err = wait.Poll(5*time.Second, 10*time.Minute, func() (bool, error) {
					new_machine, err = client.Machines().Get(new_machine.Name, metav1.GetOptions{})
					//if err == nil && new_machine.Status.Ready {
					//	return true, err
					//}
					//return false, err
					if err != nil {
						panic(err.Error())
					}
					return true, nil
				})
			}
			if err != nil {
				panic(err.Error())
			}
			errors <- err
		} (&machine_list.Items[i])
	}

	for i := 0; i < len(machine_list.Items); i++ {
		if err = <-errors; err != nil {
			panic(err.Error())
		}
	}

	fmt.Println("Successfully upgraded the cluster to version: ", *kubeversion)
}
