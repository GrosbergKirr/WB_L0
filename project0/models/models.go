package models

type Order struct {
	OrderUid          string   `json:"order_uid" validate:"required"`
	TrackNumber       string   `json:"track_number" validate:"required"`
	Entry             string   `json:"entry" validate:"required"`
	Delivery          Delivery `json:"delivery" validate:"required"`
	Payment           Payment  `json:"payment" validate:"required"`
	Items             []Items  `json:"items" validate:"required"`
	Locale            string   `json:"locale" validate:"required"`
	InternalSignature string   `json:"internal_signature" validate:"required"`
	CustomerID        string   `json:"customer_id" validate:"required"`
	DeliveryService   string   `json:"delivery_service" validate:"required"`
	Shardkey          int32    `json:"shardkey" validate:"required"`
	SmID              int64    `json:"sm_id" validate:"required"`
	DateCreated       string   `json:"date_created" validate:"required"`
	OofShard          int32    `json:"oof_shard" validate:"required"`
}

type Delivery struct {
	DelId   string `json:"-"`
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Zip     int    `json:"zip" validate:"required"`
	City    string `json:"city" validate:"required"`
	Address string `json:"address" validate:"required"`
	Region  string `json:"region" validate:"required"`
	Email   string `json:"email" validate:"required"`
}

type Payment struct {
	PayId        string `json:"-"`
	Transaction  string `json:"transaction" validate:"required"`
	RequestID    string `json:"request_id" validate:"required"`
	Currency     string `json:"currency" validate:"required"`
	Provider     string `json:"provider" validate:"required"`
	Amount       int32  `json:"amount" validate:"required"`
	PaymentDt    int32  `json:"payment_dt" validate:"required"`
	Bank         string `json:"bank" validate:"required"`
	DeliveryCost int32  `json:"delivery_cost" validate:"required"`
	GoodsTotal   int32  `json:"goods_total" validate:"required"`
	CustomFee    int32  `json:"custom_fee" validate:"required"`
}

type Items struct {
	ItemId      string `json:"-"`
	ChrtID      int64  `json:"chrt_id" validate:"required"`
	TrackNumber string `json:"track_number" validate:"required"`
	Price       int32  `json:"price" validate:"required"`
	Rid         string `json:"rid" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Sale        int32  `json:"sale" validate:"required"`
	Size        int32  `json:"size" validate:"required"`
	TotalPrice  int32  `json:"total_price" validate:"required"`
	NmID        int32  `json:"nm_id" validate:"required"`
	Brand       string `json:"brand" validate:"required"`
	Status      int32  `json:"status" validate:"required"`
}
