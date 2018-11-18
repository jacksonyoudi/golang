package persist

import (
	"context"
	"encoding/json"
	"golang/spider/engine"
	"golang/spider/model"
	"gopkg.in/olivere/elastic.v6"
	"log"
)

func ItemSaver() chan engine.Item {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("item: save, %d, %v", itemCount, item)
			itemCount++
		}
	}()
	return out
}

func Save(client *elastic.Client, item engine.Item)  (id string, err error) {

	resp, err := client.Index().Index("dating_profile").
		Type(item.Type).Id(item.Id).BodyJson(item).
		Do(context.Background())

	if err != nil {
		return "" , err
	}

	return resp.Id, nil
}

func Get(client *elastic.Client, Id string) interface{} {
	resp ,err := client.Get().Index("dating_profile").
		Type("zhenai").
		Id(Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}


	var actual model.Profile
	json.Unmarshal(*resp.Source, &actual)

	return actual
}
