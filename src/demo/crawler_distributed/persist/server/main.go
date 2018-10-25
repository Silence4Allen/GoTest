package main

import (
	"GoTest/src/demo/crawler_distributed/rpcsupport"
	"GoTest/src/demo/crawler_distributed/persist"
	"gopkg.in/olivere/elastic.v6"
	"log"
)

func main() {
	log.Fatal(serveRpc(":1234", "dating_profile"))

}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host,&persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
