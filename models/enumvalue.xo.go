// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

// EnumValue represents a enum value.
type EnumValue struct {
	EnumValue  string // enum_value
	ConstValue int    // const_value
}

// PgEnumValues runs a custom query, returning results as EnumValue.
func PgEnumValues(db XODB, schema string, table string, enum string) ([]*EnumValue, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`e.enumlabel, ` + // ::varchar AS enum_value
		`e.enumsortorder ` + // ::integer AS const_value
		`FROM pg_type t ` +
		`JOIN ONLY pg_namespace n ON n.oid = t.typnamespace ` +
		`LEFT JOIN pg_enum e ON t.oid = e.enumtypid ` +
		`WHERE n.nspname = $1 AND t.typname = $2`

	// run query
	XOLog(sqlstr, schema, enum)
	q, err := db.Query(sqlstr, schema, enum)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*EnumValue{}
	for q.Next() {
		ev := EnumValue{}

		// scan
		err = q.Scan(&ev.EnumValue, &ev.ConstValue)
		if err != nil {
			return nil, err
		}

		res = append(res, &ev)
	}

	return res, nil
}
