package storage

import (
	"encoding/json"
	"log"
	"project0/models"
)

func (s *Storage) InsertJson(uid string, dataj []byte) {

	insstmt, err := s.Db.Prepare("INSERT INTO data_ord VALUES($1, $2)")
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

	dt, err := s.Db.Query("select orders from data_ord")
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
