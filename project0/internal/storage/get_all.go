package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"project0/models"
)

func (s *Storage) OldGet() []models.Order {

	// Get array of all order_uids:

	uids, err := s.Db.Query("select order_uid from orders")
	if err != nil {
		log.Fatalf("get uids query mistake%s", err)
	}

	var orderUids []string

	for uids.Next() {
		var id string
		err := uids.Scan(&id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		orderUids = append(orderUids, id)
	}

	var ordersArray []models.Order

	for id := range orderUids {

		delstmt, err := s.Db.Prepare("select * from delivery where del_id = $1")

		if err != nil {
			log.Fatalf("DELIVERY prepare mistake%s", err)
		}
		delivery, err := delstmt.Query(orderUids[id])

		if err != nil {
			log.Fatalf("DELIVERY query mistake%s", err)
		}
		d := models.Delivery{}
		for delivery.Next() {
			err := delivery.Scan(&d.DelId, &d.Name, &d.Phone, &d.Zip, &d.City, &d.Address, &d.Region, &d.Email)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

		paystmt, err := s.Db.Prepare("select * from payment where pay_id = $1")

		if err != nil {
			log.Fatalf("ITEMS prepare mistake%s", err)
		}
		payment, err := paystmt.Query(orderUids[id])

		if err != nil {
			log.Fatalf("PAYMENT query mistake%s", err)
		}

		p := models.Payment{}
		for payment.Next() {
			err := payment.Scan(&p.PayId, &p.Transaction, &p.RequestID, &p.Currency, &p.Provider,
				&p.Amount, &p.Bank, &p.DeliveryCost, &p.GoodsTotal, &p.CustomFee)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

		itemstmt, err := s.Db.Prepare("select * from items where item_id = $1 ")

		if err != nil {
			log.Fatalf("ITEMS prepare mistake%s", err)
		}
		items, err := itemstmt.Query(orderUids[id])

		if err != nil {
			log.Fatalf("ITEMS query mistake%s", err)
		}

		var itemslst []models.Items

		for items.Next() {
			i := models.Items{}
			err := items.Scan(&i.ItemId, &i.ChrtID, &i.TrackNumber, &i.Price, &i.Rid, &i.Name,
				&i.Sale, &i.Size, &i.TotalPrice, &i.NmID, &i.Brand, &i.Status)
			if err != nil {
				fmt.Println(err)
				continue
			}
			itemslst = append(itemslst, i)
		}

		orderstmt, err := s.Db.Prepare("select * from orders where order_uid = $1")

		if err != nil {
			log.Fatalf("ORDER prepare mistake%s", err)
		}
		order, err := orderstmt.Query(orderUids[id])

		if err != nil {
			log.Fatalf("ORDER query mistake%s", err)
		}

		ord := models.Order{Delivery: d, Payment: p, Items: itemslst}

		for order.Next() {
			err := order.Scan(&ord.OrderUid, &ord.TrackNumber, &ord.Entry, &ord.Locale, &ord.InternalSignature,
				&ord.CustomerID, &ord.DeliveryService, &ord.Shardkey, &ord.SmID, &ord.DateCreated, &ord.OofShard)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
		ordersArray = append(ordersArray, ord)
	}

	return ordersArray
}

func (s *Storage) InsertJson(uid string, dataj []byte) {

	insstmt, err := s.Db.Prepare("INSERT INTO data VALUES($1, $2)")
	if err != nil {
		log.Fatalf("insert prepare mistake\n%s", err)
	}

	dt, err := insstmt.Exec(uid, dataj)
	if err != nil {
		log.Fatalf("insert exec mistake%s", err)
	}
	_ = dt

}
func (s *Storage) GetAllFromDb() []models.Order {

	dt, err := s.Db.Query("select data from data")
	if err != nil {
		log.Fatalf("insert prepare mistake\n%s", err)
	}
	if err != nil {
		log.Fatalf("insert exec mistake%s", err)
	}

	var ord []models.Order

	for dt.Next() {
		b := []byte{}
		err := dt.Scan(&b)
		if err != nil {
			panic(err)
		}
		s := models.Order{}

		err = json.Unmarshal(b, &s)
		if err != nil {
			panic(err)
		}
		ord = append(ord, s)

	}
	return ord
}
