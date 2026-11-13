package repository

import (
	"database/sql"

	"Shinano/protobuf"
)

func GetItems(b *sql.DB) ([]*protobuf.ITEMINFO, error) {
	itemList := []*protobuf.ITEMINFO{}

	items, err := b.Query("select * from inventory")
	if err != nil {
		return nil, err
	}
	defer items.Close()

	for items.Next() {
		buf := &protobuf.ITEMINFO{}
		if err = items.Scan(&buf.Id, &buf.Count); err != nil {
			return nil, err
		}
		itemList = append(itemList, buf)
	}

	return itemList, nil
}
