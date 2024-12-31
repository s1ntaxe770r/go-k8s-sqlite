package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/s1ntaxe770r/why-try-this-at-home/pkg/bst"
	"github.com/s1ntaxe770r/why-try-this-at-home/pkg/k8s"
	"github.com/s1ntaxe770r/why-try-this-at-home/pkg/squeal"
)

func main() {
	tree := bst.BinarySearchTree{}
	client, err := k8s.NewK8sClient()
	if err != nil {
		log.Fatal(err)
	}

	k8s.CreatePods(client, 7)

	// Get all the pods
	pods, err := k8s.GetPods(client)
	if err != nil {
		log.Fatal(err)
	}
	for _, pod := range pods {
		value, _ := strconv.Atoi(pod.Labels["value"])
		tree.Insert(tree.Root, value)
	}

	fmt.Println(tree.LevelOrder())

	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "./level_order.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Save the level-order traversal to the database
	err = squeal.Write(db, tree.LevelOrder())
	if err != nil {
		log.Fatal(err)
	}

}
