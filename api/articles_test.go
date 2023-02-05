package api_test

import (
	"context"
	"testing"

	"github.com/jtom38/newsbot/portal/api"
	"github.com/jtom38/newsbot/portal/services"
)

func TestArticlesList(t *testing.T) {
	ctx := context.Background()
	cfg := services.NewConfigClient()
	c := api.NewArticlesClient(cfg.MustGet(services.Config_API_Address))

	_, err := c.List(ctx, api.ArticlesListParam{})
	if err != nil {
		t.Error(err)
	}
}

func TestArticlesGet(t *testing.T) {
	ctx := context.Background()
	cfg := services.NewConfigClient()
	c := api.NewArticlesClient(cfg.MustGet(services.Config_API_Address))

	res, err := c.List(ctx, api.ArticlesListParam{})
	if err != nil {
		t.Error(err)
	}

	single, err := c.Get(ctx, res[0].ID)
	if err != nil {
		t.Error(err)
	}

	if single.ID != res[0].ID {
		t.Error("got the wrong record back")
	}
}

func TestArticlesBySourceId(t *testing.T) {
	ctx := context.Background()
	cfg := services.NewConfigClient()
	c := api.NewArticlesClient(cfg.MustGet(services.Config_API_Address))

	res, err := c.List(ctx, api.ArticlesListParam{})
	if err != nil {
		t.Error(err)
	}

	sourceRecords, err := c.ListBySourceId(ctx, res[0].SourceID, 0)
	if err != nil {
		t.Error(err)
	}

	if len(*sourceRecords) == 0 {
		t.Error("did not get the expected results")
	}
}

func TestArticlesGetDetails(t *testing.T) {
	ctx := context.Background()
	cfg := services.NewConfigClient()
	c := api.NewArticlesClient(cfg.MustGet(services.Config_API_Address))

	res, err := c.List(ctx, api.ArticlesListParam{})
	if err != nil {
		t.Error(err)
	}

	item, err := c.GetDetails(ctx, res[0].ID)
	if err != nil {
		t.Error(err)
	}

	if item.Source.ID != res[0].SourceID {
		t.Error("did not get the expected results")
	}
}
