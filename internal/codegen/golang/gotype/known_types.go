package gotype

import (
	"github.com/jackc/pgtype"
	"github.com/jschaf/pggen/internal/pg/pgoid"
)

// FindKnownTypePgx returns the native pgx type, like pgtype.Text, if known, for
// a Postgres OID. If there is no known type, returns nil.
func FindKnownTypePgx(oid pgtype.OID) (Type, bool) {
	typ, ok := knownTypesByOID[oid]
	return typ.pgNative, ok
}

// FindKnownTypeNullable returns the nullable type, like *string, if known, for
// a Postgres OID. Falls back to the pgNative type. If there is no known type
// for the OID, returns nil.
func FindKnownTypeNullable(oid pgtype.OID) (Type, bool) {
	typ, ok := knownTypesByOID[oid]
	if !ok {
		return nil, false
	}
	if typ.nullable != nil {
		return typ.nullable, true
	}
	return typ.pgNative, true
}

// FindKnownTypeNonNullable returns the non-nullable type like string, if known,
// for a Postgres OID. Falls back to the nullable type and pgNative type. If
// there is no known type for the OID, returns nil.
func FindKnownTypeNonNullable(oid pgtype.OID) (Type, bool) {
	typ, ok := knownTypesByOID[oid]
	if !ok {
		return nil, false
	}
	if typ.nonNullable != nil {
		return typ.nonNullable, true
	}
	if typ.nullable != nil {
		return typ.nullable, true
	}
	return typ.pgNative, true
}

// Native go types are not prefixed.
//goland:noinspection GoUnusedGlobalVariable
var (
	Bool          = MustParseKnownType("bool")
	Boolp         = MustParseKnownType("*bool")
	Int           = MustParseKnownType("int")
	Intp          = MustParseKnownType("*int")
	IntSlice      = MustParseKnownType("[]int")
	IntpSlice     = MustParseKnownType("[]*int")
	Int16         = MustParseKnownType("int16")
	Int16p        = MustParseKnownType("*int16")
	Int16Slice    = MustParseKnownType("[]int16")
	Int16pSlice   = MustParseKnownType("[]*int16")
	Int32         = MustParseKnownType("int32")
	Int32p        = MustParseKnownType("*int32")
	Int32Slice    = MustParseKnownType("[]int32")
	Int32pSlice   = MustParseKnownType("[]*int32")
	Int64         = MustParseKnownType("int64")
	Int64p        = MustParseKnownType("*int64")
	Int64Slice    = MustParseKnownType("[]int64")
	Int64pSlice   = MustParseKnownType("[]*int64")
	Uint          = MustParseKnownType("uint")
	UintSlice     = MustParseKnownType("[]uint")
	Uint16        = MustParseKnownType("uint16")
	Uint16Slice   = MustParseKnownType("[]uint16")
	Uint32        = MustParseKnownType("uint32")
	Uint32Slice   = MustParseKnownType("[]uint32")
	Uint64        = MustParseKnownType("uint64")
	Uint64Slice   = MustParseKnownType("[]uint64")
	String        = MustParseKnownType("string")
	Stringp       = MustParseKnownType("*string")
	StringSlice   = MustParseKnownType("[]string")
	StringpSlice  = MustParseKnownType("[]*string")
	Float32       = MustParseKnownType("float32")
	Float32p      = MustParseKnownType("*float32")
	Float32Slice  = MustParseKnownType("[]float32")
	Float32pSlice = MustParseKnownType("[]*float32")
	Float64       = MustParseKnownType("float64")
	Float64p      = MustParseKnownType("*float64")
	Float64Slice  = MustParseKnownType("[]float64")
	Float64pSlice = MustParseKnownType("[]*float64")
	ByteSlice     = MustParseKnownType("[]byte")
)

// pgtype types prefixed with "pg".
var (
	PgBool             = MustParseKnownType("github.com/jackc/pgtype.Bool")
	PgQChar            = MustParseKnownType("github.com/jackc/pgtype.QChar")
	PgName             = MustParseKnownType("github.com/jackc/pgtype.Name")
	PgInt8             = MustParseKnownType("github.com/jackc/pgtype.Int8")
	PgInt2             = MustParseKnownType("github.com/jackc/pgtype.Int2")
	PgInt4             = MustParseKnownType("github.com/jackc/pgtype.Int4")
	PgText             = MustParseKnownType("github.com/jackc/pgtype.Text")
	PgBytea            = MustParseKnownType("github.com/jackc/pgtype.Bytea")
	PgOID              = MustParseKnownType("github.com/jackc/pgtype.OID")
	PgTID              = MustParseKnownType("github.com/jackc/pgtype.TID")
	PgXID              = MustParseKnownType("github.com/jackc/pgtype.XID")
	PgCID              = MustParseKnownType("github.com/jackc/pgtype.CID")
	PgJSON             = MustParseKnownType("github.com/jackc/pgtype.JSON")
	PgPoint            = MustParseKnownType("github.com/jackc/pgtype.Point")
	PgLseg             = MustParseKnownType("github.com/jackc/pgtype.Lseg")
	PgPath             = MustParseKnownType("github.com/jackc/pgtype.Path")
	PgBox              = MustParseKnownType("github.com/jackc/pgtype.Box")
	PgPolygon          = MustParseKnownType("github.com/jackc/pgtype.Polygon")
	PgLine             = MustParseKnownType("github.com/jackc/pgtype.Line")
	PgCIDR             = MustParseKnownType("github.com/jackc/pgtype.CIDR")
	PgCIDRArray        = MustParseKnownType("github.com/jackc/pgtype.CIDRArray")
	PgFloat4           = MustParseKnownType("github.com/jackc/pgtype.Float4")
	PgFloat8           = MustParseKnownType("github.com/jackc/pgtype.Float8")
	PgUnknown          = MustParseKnownType("github.com/jackc/pgtype.Unknown")
	PgCircle           = MustParseKnownType("github.com/jackc/pgtype.Circle")
	PgMacaddr          = MustParseKnownType("github.com/jackc/pgtype.Macaddr")
	PgInet             = MustParseKnownType("github.com/jackc/pgtype.Inet")
	PgBoolArray        = MustParseKnownType("github.com/jackc/pgtype.BoolArray")
	PgByteaArray       = MustParseKnownType("github.com/jackc/pgtype.ByteaArray")
	PgInt2Array        = MustParseKnownType("github.com/jackc/pgtype.Int2Array")
	PgInt4Array        = MustParseKnownType("github.com/jackc/pgtype.Int4Array")
	PgTextArray        = MustParseKnownType("github.com/jackc/pgtype.TextArray")
	PgBPCharArray      = MustParseKnownType("github.com/jackc/pgtype.BPCharArray")
	PgVarcharArray     = MustParseKnownType("github.com/jackc/pgtype.VarcharArray")
	PgInt8Array        = MustParseKnownType("github.com/jackc/pgtype.Int8Array")
	PgFloat4Array      = MustParseKnownType("github.com/jackc/pgtype.Float4Array")
	PgFloat8Array      = MustParseKnownType("github.com/jackc/pgtype.Float8Array")
	PgACLItem          = MustParseKnownType("github.com/jackc/pgtype.ACLItem")
	PgACLItemArray     = MustParseKnownType("github.com/jackc/pgtype.ACLItemArray")
	PgInetArray        = MustParseKnownType("github.com/jackc/pgtype.InetArray")
	PgMacaddrArray     = MustParseKnownType("github.com/jackc/pgtype.MacaddrArray")
	PgBPChar           = MustParseKnownType("github.com/jackc/pgtype.BPChar")
	PgVarchar          = MustParseKnownType("github.com/jackc/pgtype.Varchar")
	PgDate             = MustParseKnownType("github.com/jackc/pgtype.Date")
	PgTime             = MustParseKnownType("github.com/jackc/pgtype.Time")
	PgTimestamp        = MustParseKnownType("github.com/jackc/pgtype.Timestamp")
	PgTimestampArray   = MustParseKnownType("github.com/jackc/pgtype.TimestampArray")
	PgDateArray        = MustParseKnownType("github.com/jackc/pgtype.DateArray")
	PgTimestamptz      = MustParseKnownType("github.com/jackc/pgtype.Timestamptz")
	PgTimestamptzArray = MustParseKnownType("github.com/jackc/pgtype.TimestamptzArray")
	PgInterval         = MustParseKnownType("github.com/jackc/pgtype.Interval")
	PgNumericArray     = MustParseKnownType("github.com/jackc/pgtype.NumericArray")
	PgBit              = MustParseKnownType("github.com/jackc/pgtype.Bit")
	PgVarbit           = MustParseKnownType("github.com/jackc/pgtype.Varbit")
	PgVoid             = &VoidType{}
	PgNumeric          = MustParseKnownType("github.com/jackc/pgtype.Numeric")
	PgRecord           = MustParseKnownType("github.com/jackc/pgtype.Record")
	PgUUID             = MustParseKnownType("github.com/jackc/pgtype.UUID")
	PgUUIDArray        = MustParseKnownType("github.com/jackc/pgtype.UUIDArray")
	PgJSONB            = MustParseKnownType("github.com/jackc/pgtype.JSONB")
	PgJSONBArray       = MustParseKnownType("github.com/jackc/pgtype.JSONBArray")
	PgInt4range        = MustParseKnownType("github.com/jackc/pgtype.Int4range")
	PgNumrange         = MustParseKnownType("github.com/jackc/pgtype.Numrange")
	PgTsrange          = MustParseKnownType("github.com/jackc/pgtype.Tsrange")
	PgTstzrange        = MustParseKnownType("github.com/jackc/pgtype.Tstzrange")
	PgDaterange        = MustParseKnownType("github.com/jackc/pgtype.Daterange")
	PgInt8range        = MustParseKnownType("github.com/jackc/pgtype.Int8range")
)

// knownGoType is the native pgtype type, the nullable and non-nullable types
// for a Postgres type.
//
// pgNative means a type that implements the pgx decoder methods directly.
// Such types are typically provided by the pgtype package. Used as the fallback
// type and for cases like composite types where we need a
// pgtype.ValueTranscoder.
//
// A nullable type is one that can represent a nullable column, like *string for
// a Postgres text type that can be null. A nullable type is nicer to work with
// than the corresponding pgNative type, i.e. "*string" is easier to work with
// than pgtype.Text{}.
//
// A nonNullable type is one that can represent a column that's never null, like
// "string" for a Postgres text type.
type knownGoType struct{ pgNative, nullable, nonNullable Type }

var knownTypesByOID = map[pgtype.OID]knownGoType{
	pgtype.BoolOID:             {PgBool, Boolp, Bool},
	pgtype.QCharOID:            {PgQChar, nil, nil},
	pgtype.NameOID:             {PgName, nil, nil},
	pgtype.Int8OID:             {PgInt8, Intp, Int},
	pgtype.Int2OID:             {PgInt2, Int16p, Int16},
	pgtype.Int4OID:             {PgInt4, Int32p, Int32},
	pgtype.TextOID:             {PgText, Stringp, String},
	pgtype.ByteaOID:            {PgBytea, PgBytea, ByteSlice},
	pgtype.OIDOID:              {PgOID, nil, nil},
	pgtype.TIDOID:              {PgTID, nil, nil},
	pgtype.XIDOID:              {PgXID, nil, nil},
	pgtype.CIDOID:              {PgCID, nil, nil},
	pgtype.JSONOID:             {PgJSON, nil, nil},
	pgtype.PointOID:            {PgPoint, nil, nil},
	pgtype.LsegOID:             {PgLseg, nil, nil},
	pgtype.PathOID:             {PgPath, nil, nil},
	pgtype.BoxOID:              {PgBox, nil, nil},
	pgtype.PolygonOID:          {PgPolygon, nil, nil},
	pgtype.LineOID:             {PgLine, nil, nil},
	pgtype.CIDROID:             {PgCIDR, nil, nil},
	pgtype.CIDRArrayOID:        {PgCIDRArray, nil, nil},
	pgtype.Float4OID:           {PgFloat4, nil, nil},
	pgtype.Float8OID:           {PgFloat8, nil, nil},
	pgoid.OIDArray:             {Uint32Slice, nil, nil},
	pgtype.UnknownOID:          {PgUnknown, nil, nil},
	pgtype.CircleOID:           {PgCircle, nil, nil},
	pgtype.MacaddrOID:          {PgMacaddr, nil, nil},
	pgtype.InetOID:             {PgInet, nil, nil},
	pgtype.BoolArrayOID:        {PgBoolArray, nil, nil},
	pgtype.ByteaArrayOID:       {PgByteaArray, nil, nil},
	pgtype.Int2ArrayOID:        {PgInt2Array, Int16pSlice, Int16Slice},
	pgtype.Int4ArrayOID:        {PgInt4Array, Int32pSlice, Int32Slice},
	pgtype.TextArrayOID:        {PgTextArray, StringSlice, nil},
	pgtype.BPCharArrayOID:      {PgBPCharArray, nil, nil},
	pgtype.VarcharArrayOID:     {PgVarcharArray, nil, nil},
	pgtype.Int8ArrayOID:        {PgInt8Array, IntpSlice, IntSlice},
	pgtype.Float4ArrayOID:      {PgFloat4Array, Float32pSlice, Float32Slice},
	pgtype.Float8ArrayOID:      {PgFloat8Array, Float64pSlice, Float64Slice},
	pgtype.ACLItemOID:          {PgACLItem, nil, nil},
	pgtype.ACLItemArrayOID:     {PgACLItemArray, nil, nil},
	pgtype.InetArrayOID:        {PgInetArray, nil, nil},
	pgoid.MacaddrArray:         {PgMacaddrArray, nil, nil},
	pgtype.BPCharOID:           {PgBPChar, nil, nil},
	pgtype.VarcharOID:          {PgVarchar, nil, nil},
	pgtype.DateOID:             {PgDate, nil, nil},
	pgtype.TimeOID:             {PgTime, nil, nil},
	pgtype.TimestampOID:        {PgTimestamp, nil, nil},
	pgtype.TimestampArrayOID:   {PgTimestampArray, nil, nil},
	pgtype.DateArrayOID:        {PgDateArray, nil, nil},
	pgtype.TimestamptzOID:      {PgTimestamptz, nil, nil},
	pgtype.TimestamptzArrayOID: {PgTimestamptzArray, nil, nil},
	pgtype.IntervalOID:         {PgInterval, nil, nil},
	pgtype.NumericArrayOID:     {PgNumericArray, nil, nil},
	pgtype.BitOID:              {PgBit, nil, nil},
	pgtype.VarbitOID:           {PgVarbit, nil, nil},
	pgoid.Void:                 {PgVoid, nil, nil},
	pgtype.NumericOID:          {PgNumeric, nil, nil},
	pgtype.RecordOID:           {PgRecord, nil, nil},
	pgtype.UUIDOID:             {PgUUID, nil, nil},
	pgtype.UUIDArrayOID:        {PgUUIDArray, nil, nil},
	pgtype.JSONBOID:            {PgJSONB, nil, nil},
	pgtype.JSONBArrayOID:       {PgJSONBArray, nil, nil},
	pgtype.Int4rangeOID:        {PgInt4range, nil, nil},
	pgtype.NumrangeOID:         {PgNumrange, nil, nil},
	pgtype.TsrangeOID:          {PgTsrange, nil, nil},
	pgtype.TstzrangeOID:        {PgTstzrange, nil, nil},
	pgtype.DaterangeOID:        {PgDaterange, nil, nil},
	pgtype.Int8rangeOID:        {PgInt8range, nil, nil},
}
