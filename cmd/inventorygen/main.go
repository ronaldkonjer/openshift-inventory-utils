package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/ronaldkonjer/openshift-inventory-utils/inventory"
)

func main() {
	masterArg := flag.String("masters", "", "Comma-separated domain list for master node")
	etcdArg := flag.String("etcd", "", "Comma-separated domain list for etcd node")
	nfsArg := flag.String("nfs", "", "Comma-separated domain list for nfs node")
	nodesArg := flag.String("nodes", "", "Comma-separated domain list for nodes")
	inventoryPath := flag.String("inventory", "", "Inventory file on which the new inventory based")
	flag.Parse()

	nodes := parseNodeArg(*nodesArg)
	masters := parseNodeArg(*masterArg)
	etcd := parseNodeArg(*etcdArg)
	nfs := parseNodeArg(*nfsArg)

	inventory, err := inventory.Generate(nodes, masters, etcd, nfs *inventoryPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(inventory)
}

func parseNodeArg(arg string) []string {
	if arg == "" {
		return nil
	}
	return strings.Split(arg, ",")
}
