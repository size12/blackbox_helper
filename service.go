package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type Service struct {
	Client *resty.Client
}

func NewService(cfg Config) *Service {
	r := resty.New()
	r.SetBaseURL(cfg.AccrualSystemAddress)
	r.SetHeader("Content-Type", "application/json")
	return &Service{Client: r}
}

func (s *Service) AddGoods(goods []string) error {
	for _, el := range goods {
		item := Item{
			Match:      el,
			Reward:     10,
			RewardType: "%",
		}

		b, err := json.Marshal(item)

		if err != nil {
			return err
		}

		res, err := s.Client.R().SetBody(b).Post("/api/goods")

		if err != nil {
			return err
		}

		code := res.StatusCode()

		if code == http.StatusConflict {
			return errors.New("sent item which is already in system")
		}

		if code != http.StatusOK {
			return errors.New("didn't receive 200 OK status code")
		}

	}
	return nil
}

func (s *Service) AddOrders(orders []Order) error {
	for _, order := range orders {

		b, err := json.Marshal(order)

		if err != nil {
			return err
		}

		res, err := s.Client.R().SetBody(b).Post("/api/orders")

		if err != nil {
			return err
		}

		code := res.StatusCode()

		if code == http.StatusConflict {
			return errors.New("sent item which is already in system")
		}

		if code != http.StatusAccepted {
			return errors.New("didn't receive 202 Accepted status code")
		}

	}
	return nil
}
