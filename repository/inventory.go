package repository

import (
	"database/sql"

	"Shinano/protobuf"
)

func GetResources(b *sql.DB) ([]*protobuf.RESOURCE, error) {
	var ResourceList []*protobuf.RESOURCE

	rows, err := b.Query("select * from resourcelist ")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var buf protobuf.RESOURCE

		err = rows.Scan(&buf.Type, &buf.Num)
		if err != nil {
			return nil, err
		}

		ResourceList = append(ResourceList, &buf)
	}

	return ResourceList, nil
}

func GetInventory(b *sql.DB) ([]*protobuf.ITEMINFO, error) {
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
