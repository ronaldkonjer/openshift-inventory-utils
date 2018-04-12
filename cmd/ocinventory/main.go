package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/ronaldkonjer/openshift-inventory-utils/inventory"
	"github.com/ronaldkonjer/openshift-inventory-utils/node"
)

func main() {
	clusterArg := flag.String("cluster", "", "Cluster name used in the tag: kubernetes.io/cluster/xxx")
	roleTagArg := flag.String("role-tag", "Role", "Tag name for specifying node types")
	inventoryPath := flag.String("inventory", "", "Inventory file on which the new inventory based")
	flag.Parse()

	nodes, errNodes := node.List(*clusterArg, "node", *roleTagArg)
	if errNodes != nil {
		log.Fatal(errNodes)
	}

	masters, errMasters := node.List(*clusterArg, "master", *roleTagArg)
	if errMasters != nil {
		log.Fatal(errMasters)
	}

	etcd, errEtcd := node.List(*clusterArg, "etcd", *roleTagArg)
	if errEtcd != nil {
		log.Fatal(errEtcd)
	}

	nfs, errNfs := node.List(*clusterArg, "nfs", *roleTagArg)
	if errNfs != nil {
		log.Fatal(errNfs)
	}

	inventory, err := inventory.Generate(nodes, masters, etcd, *inventoryPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(inventory)
}
