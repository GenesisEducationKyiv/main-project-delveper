package api

import (
	"net/http"
	"path"

	"github.com/GenesisEducationKyiv/main-project-delveper/internal/rate"
	"github.com/GenesisEducationKyiv/main-project-delveper/internal/subscription"
	"github.com/GenesisEducationKyiv/main-project-delveper/sys/filestore"
	"github.com/hashicorp/go-retryablehttp"
)

const (
	pathRate       = "/rate"
	pathSubscribe  = "/subscribe"
	pathSendEmails = "/sendEmails"
)

func WithRate(cfg ConfigAggregate) Route {
	return func(app *App) {
		grp := path.Join(cfg.Api.Path, cfg.Api.Version)

		client := new(http.Client)
		btcSvc := rate.NewBTCExchangeRateClient(client, cfg.Rate.Provider.RapidApi.Endpoint)

		svc := rate.NewService(btcSvc)
		h := rate.NewHandler(svc)

		app.web.Handle(http.MethodGet, grp, pathRate, h.Rate)
	}
}

func WithSubscription(cfg ConfigAggregate) Route {
	return func(app *App) {
		grp := path.Join(cfg.Api.Path, cfg.Api.Version)

		conn := filestore.New[subscription.Subscriber](cfg.Subscription.Repo.Data)

		client := retryablehttp.NewClient()
		client.RetryMax = cfg.Rate.Client.RetryMax
		btcSvc := rate.NewBTCExchangeRateClient(client.StandardClient(), cfg.Rate.Provider.RapidApi.Endpoint)

		svc := subscription.NewService(
			subscription.NewRepo(conn),
			rate.NewService(btcSvc),
			subscription.NewSender(cfg.Subscription.Sender.Address, cfg.Subscription.Sender.Key),
		)
		h := subscription.NewHandler(svc)

		app.web.Handle(http.MethodPost, grp, pathSendEmails, h.SendEmails)
		app.web.Handle(http.MethodPost, grp, pathSubscribe, h.Subscribe)
	}
}
