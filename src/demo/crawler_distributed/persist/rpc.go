package persist

import (
	"GoTest/src/demo/crawler/engine"
	"gopkg.in/olivere/elastic.v6"
	"GoTest/src/demo/crawler/persist"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.save(s.Client, s.Index, item)
	if err == nil {
		*result = "ok"
	}
	return err
}
