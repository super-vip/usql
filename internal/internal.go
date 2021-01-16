package internal

// Code generated by gen.sh. DO NOT EDIT.

//go:generate ./gen.sh

// KnownBuildTags returns a map of known driver names to its respective build
// tags.
func KnownBuildTags() map[string]string {
	return map[string]string{
		"adodb":         "adodb",         // github.com/mattn/go-adodb
		"athena":        "athena",        // github.com/uber/athenadriver/go
		"avatica":       "avatica",       // github.com/apache/calcite-avatica-go/v5
		"bigquery":      "bigquery",      // gorm.io/driver/bigquery/driver
		"cassandra":     "cql",           // github.com/MichaelS11/go-cql-driver
		"clickhouse":    "clickhouse",    // github.com/ClickHouse/clickhouse-go
		"cosmos":        "cosmos",        // github.com/btnguyen2k/gocosmos
		"couchbase":     "n1ql",          // github.com/couchbase/go_n1ql
		"firebird":      "firebirdsql",   // github.com/nakagami/firebirdsql
		"genji":         "genji",         // github.com/genjidb/genji/sql/driver
		"godror":        "godror",        // github.com/godror/godror
		"h2":            "h2",            // github.com/jmrobles/h2go
		"hive":          "hive",          // sqlflow.org/gohive
		"ignite":        "ignite",        // github.com/amsokol/ignite-go-client/sql
		"impala":        "impala",        // github.com/bippio/go-impala
		"maxcompute":    "maxcompute",    // sqlflow.org/gomaxcompute
		"moderncsqlite": "moderncsqlite", // modernc.org/sqlite
		"mssql":         "mssql",         // github.com/denisenkom/go-mssqldb
		"mymysql":       "mymysql",       // github.com/ziutek/mymysql/godrv
		"mysql":         "mysql",         // github.com/go-sql-driver/mysql
		"odbc":          "odbc",          // github.com/alexbrainman/odbc
		"oracle":        "oracle",        // github.com/sijms/go-ora
		"pgx":           "pgx",           // github.com/jackc/pgx/v4/stdlib
		"postgres":      "postgres",      // github.com/lib/pq
		"presto":        "presto",        // github.com/prestodb/presto-go-client/presto
		"ql":            "ql",            // modernc.org/ql
		"sapase":        "tds",           // github.com/thda/tds
		"saphana":       "hdb",           // github.com/SAP/go-hdb/driver
		"snowflake":     "snowflake",     // github.com/snowflakedb/gosnowflake
		"spanner":       "spanner",       // github.com/rakyll/go-sql-driver-spanner
		"sqlite3":       "sqlite3",       // github.com/mattn/go-sqlite3
		"trino":         "trino",         // github.com/trinodb/trino-go-client/trino
		"vertica":       "vertica",       // github.com/vertica/vertica-sql-go
		"voltdb":        "voltdb",        // github.com/VoltDB/voltdb-client-go/voltdbclient
	}
}
