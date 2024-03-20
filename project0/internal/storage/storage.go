package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, err
	}
	//cr1, err := db.Exec("create table if not exists" +
	//	"orders (order_uid varchar, track_number varchar, entry varchar, delivery varchar," +
	//	"payment varchar, items varchar," +
	//	"locale varchar, internal_signature varchar," +
	//	"customer_id varchar, delivery_service varchar," +
	//	"shardkey integer, sm_id integer, date_created timestamp, oof_shard integer," +
	//	"PRIMARY KEY(order_uid))")
	//
	//cr2, err := db.Exec("create table if not exists " +
	//	"delivery (del_id varchar,name varchar, phone varchar, zip integer, city varchar," +
	//	"address varchar, region varchar, email varchar," +
	//	"PRIMARY KEY(del_id))")
	//
	//cr3, err := db.Exec("create table if not exists " +
	//	"payment(pay_id varchar, transaction varchar, request_id varchar, currency varchar, provider varchar," +
	//	"amount varchar, payment_dt  integer, bank varchar, delivery_cost varchar," +
	//	"goods_total integer, custom_fee integer," +
	//	"primary key(pay_id))")
	//
	//cr4, err := db.Exec("create table if not exists " +
	//	"items (item_id varchar, chrt_id integer, track_number varchar, price integer," +
	//	"rid varchar, name varchar,sale integer, size integer," +
	//	"total_price integer, nm_id integer, brand varchar, status integer," +
	//	"primary key(item_id))")
	//
	//_ = cr1
	//_ = cr2
	//_ = cr3
	//_ = cr4
	//if err != nil {
	//	return nil, err
	//}

	return &Storage{Db: db}, nil
}
