package main

type Item struct {
	Match      string  `json:"match"`
	Reward     float64 `json:"reward"`
	RewardType string  `json:"reward_type"`
}

type OrderItem struct {
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type Order struct {
	Number      int         `json:"order,string"`
	Goods       []OrderItem `json:"goods"`
	WaitAccrual float64     `json:"-"`
}
