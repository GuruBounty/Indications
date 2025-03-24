package repository

import (
	"context"
	"database/sql"

	"indication/internal/domain"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type NumLSRepository struct {
	db *sqlx.DB
}

func NewNumLSRepository(db *sqlx.DB) *NumLSRepository {
	return &NumLSRepository{db: db}
}

func (n *NumLSRepository) GetObjectsByNumLS(ctx context.Context, ls int64) ([]domain.LS_Object, error) {
	var numls []domain.LS_Object

	query := `SELECT ls.num_ls, 
	t.type, adr.address,
	md.day_night_type, 
	md.device_type, 
	md.last_metering,
	md.device_number,
	md.device_guid
	FROM ls_objects AS ls
	JOIN addresses as adr ON ls.address_id = adr.id
	JOIN object_types as t ON ls.object_type_id = t.id
	JOIN metering_devices md ON ls.id = md.ls_object_id
	WHERE ls.num_ls = $1
	`
	err := n.db.SelectContext(ctx, &numls, query, ls)

	if err == sql.ErrNoRows {
		return []domain.LS_Object{}, nil
	} else if err != nil {
		return nil, err
	}
	// }
	// if err != nil {
	// 	// if err == sql.ErrNoRows {
	// 	// 	return nil, nil
	// 	// }
	// 	return nil, err
	// }

	return numls, err
}

func (n NumLSRepository) SetMeterIndicationByGUID(ctx context.Context, guid string, meter float32, request int64) (int64, error) {
	//var id int64

	query := `UPDATE metering_devices 
	SET last_metering = $1, updated_at = NOW()
	WHERE device_guid = $2
	`
	result, err := n.db.ExecContext(ctx, query, meter, guid)
	if err != nil {
		return 0, err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowAffected == 0 {
		return 0, nil
	}
	return rowAffected, nil
}
