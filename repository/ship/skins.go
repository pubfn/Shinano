package ship

import "database/sql"

func ApplySkin(b *sql.DB, id int, skinId int) error {
	_, err := b.Exec("update shiplist set skinid = ? where id = ?", skinId, id)
	if err != nil {
		return err
	}
	return nil
}

func GetAllSkins(b *sql.DB) ([]uint32, error) {
	rows, err := b.Query("select * from skinlist")
	skins := []uint32{}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var skinid uint32
		if err := rows.Scan(&skinid); err != nil {
			return nil, err
		}
		skins = append(skins, skinid)
	}

	return skins, err
}

func BuySkin(b *sql.DB, skinId int) error {
	_, err := b.Exec("insert into skinlist (skinid) values (?)", skinId)
	if err != nil {
		return err
	}

	return nil
}
