package datatypes

import (
	"context"
	"database/sql/driver"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"math"
	"reflect"
)

type Blob []byte

func (b Blob) Value(ctx context.Context) (driver.Value, error) {
	// driver.Value is either nil,
	// a type handled by a database driver's NamedValueChecker interface,
	// or an instance of one of these types:
	//
	//	int64
	//	float64
	//	bool
	//	[]byte
	//	string
	//	time.Time
	//
	return []byte(b), nil
}

func (b *Blob) Scan(src interface{}) error {
	// The src value will be of one of the following types:
	//
	//    int64
	//    float64
	//    bool
	//    []byte
	//    string
	//    time.Time
	//    nil - for NULL values
	//
	switch val := src.(type) {
	case nil:
		*b = nil
	case []byte:
		*b = append([]byte(nil), val...)
	case string:
		*b = append([]byte(nil), val...)
	default:
		return fmt.Errorf("cannot scan value of %s type into Blob", reflect.TypeOf(src))
	}
	return nil
}

func (b Blob) MarshalJSON() ([]byte, error) {
	base64Str := base64.StdEncoding.EncodeToString(b)
	return json.Marshal(base64Str)
}

// UnmarshalJSON to deserialize []byte
func (b *Blob) UnmarshalJSON(data []byte) (err error) {
	var base64Str string
	err = json.Unmarshal(data, &base64Str)
	if err != nil {
		*b = nil
	} else {
		*b, err = base64.StdEncoding.DecodeString(base64Str)
	}
	return
}

// GormDataType gorm common data type
func (b Blob) GormDataType() string {
	return string(schema.Bytes)
}

// GormDBDataType gorm db data type
func (b Blob) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	if field.Size > 0 && field.Size < 65536 {
		return fmt.Sprintf("varbinary(%d)", field.Size)
	}
	if field.Size >= 65536 && field.Size <= int(math.Pow(2, 24)) {
		return "mediumblob"
	}
	return "longblob"
}

func (b Blob) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	// 当 DryRun
	if !db.DryRun {
		return gorm.Expr("?", []byte(b))
	}

	if len(b) == 0 {
		return gorm.Expr("NULL")
	}

	hexStr := hex.EncodeToString(b)
	return gorm.Expr("X?", hexStr)
}
