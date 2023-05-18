package datatypes

import (
	"context"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"strings"
)

type NullTypedJSON[T any] struct {
	Value T
	Valid bool
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (j *NullTypedJSON[T]) Scan(value interface{}) error {
	j.Valid = false
	if value == nil {
		return nil
	}
	var bytes []byte
	switch v := value.(type) {
	case []byte:
		if len(v) > 0 {
			bytes = make([]byte, len(v))
			copy(bytes, v)
		}
	case string:
		bytes = []byte(v)
	default:
		return fmt.Errorf("typed json must scan from string or bytes value")
	}

	if err := json.Unmarshal(bytes, &j.Value); err != nil {
		return err
	}
	j.Valid = true
	return nil
}

// MarshalJSON to output non base64 encoded []byte
func (j NullTypedJSON[T]) MarshalJSON() ([]byte, error) {
	if !j.Valid {
		return nil, nil
	}
	return json.Marshal(j.Value)
}

// UnmarshalJSON to deserialize []byte
func (j *NullTypedJSON[T]) UnmarshalJSON(data []byte) error {
	j.Valid = false
	if len(data) == 0 {
		return nil
	}
	if err := json.Unmarshal(data, &j.Value); err != nil {
		return err
	}
	j.Valid = true
	return nil
}

// GormDataType gorm common data type
func (NullTypedJSON[T]) GormDataType() string {
	return "json"
}

// GormDBDataType gorm db data type
func (NullTypedJSON[T]) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	return "JSON"
}

func (j NullTypedJSON[T]) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	data, _ := j.MarshalJSON()

	if v, ok := db.Dialector.(*mysql.Dialector); ok && !strings.Contains(v.ServerVersion, "MariaDB") {
		return gorm.Expr("CAST(? AS JSON)", string(data))
	}

	return gorm.Expr("?", string(data))
}
