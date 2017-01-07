// generated by internal/cmd/gentokens/main.go

package schemalex

type TokenType int

// Token represents a token
type Token struct {
	Type  TokenType
	Value string
	Pos   int
	Line  int
	Col   int
	EOF   bool
}

func NewToken(t TokenType, v string) *Token {
	return &Token{Type: t, Value: v}
}

const (
	ILLEGAL TokenType = iota
	EOF
	SPACE
	IDENT
	BACKTICK_IDENT
	DOUBLE_QUOTE_IDENT
	SINGLE_QUOTE_IDENT
	NUMBER
	LPAREN        // (
	RPAREN        // )
	COMMA         // ,
	SEMICOLON     // ;
	DOT           // .
	SLASH         // /
	ASTERISK      // *
	DASH          // -
	PLUS          // +
	SINGLE_QUOTE  // '
	DOUBLE_QUOTE  // "
	EQUAL         // =
	COMMENT_IDENT // // /*   */, --, #
	ACTION
	AUTO_INCREMENT
	AVG_ROW_LENGTH
	BIGINT
	BINARY
	BIT
	BLOB
	BTREE
	CASCADE
	CHAR
	CHARACTER
	CHECK
	CHECKSUM
	COLLATE
	COMMENT
	COMPACT
	COMPRESSED
	CONNECTION
	CONSTRAINT
	CREATE
	CURRENT_TIMESTAMP
	DATA
	DATABASE
	DATE
	DATETIME
	DECIMAL
	DEFAULT
	DELAY_KEY_WRITE
	DELETE
	DIRECTORY
	DISK
	DOUBLE
	DROP
	DYNAMIC
	ENGINE
	ENUM
	EXISTS
	FIRST
	FIXED
	FLOAT
	FOREIGN
	FULL
	FULLTEXT
	HASH
	IF
	INDEX
	INSERT_METHOD
	INT
	INTEGER
	KEY
	KEY_BLOCK_SIZE
	LAST
	LONGBLOB
	LONGTEXT
	MATCH
	MAX_ROWS
	MEDIUMBLOB
	MEDIUMINT
	MEDIUMTEXT
	MEMORY
	MIN_ROWS
	NO
	NOT
	NULL
	NUMERIC
	ON
	PACK_KEYS
	PARTIAL
	PASSWORD
	PRIMARY
	REAL
	REDUNDANT
	REFERENCES
	RESTRICT
	ROW_FORMAT
	SET
	SIMPLE
	SMALLINT
	SPARTIAL
	STATS_AUTO_RECALC
	STATS_PERSISTENT
	STATS_SAMPLE_PAGES
	STORAGE
	TABLE
	TABLESPACE
	TEMPORARY
	TEXT
	TIME
	TIMESTAMP
	TINYBLOB
	TINYINT
	TINYTEXT
	UNION
	UNIQUE
	UNSIGNED
	UPDATE
	USE
	USING
	VARBINARY
	VARCHAR
	YEAR
	ZEROFILL
)

var keywordIdentMap = map[string]TokenType{
	"ACTION":             ACTION,
	"AUTO_INCREMENT":     AUTO_INCREMENT,
	"AVG_ROW_LENGTH":     AVG_ROW_LENGTH,
	"BIGINT":             BIGINT,
	"BINARY":             BINARY,
	"BIT":                BIT,
	"BLOB":               BLOB,
	"BTREE":              BTREE,
	"CASCADE":            CASCADE,
	"CHAR":               CHAR,
	"CHARACTER":          CHARACTER,
	"CHECK":              CHECK,
	"CHECKSUM":           CHECKSUM,
	"COLLATE":            COLLATE,
	"COMMENT":            COMMENT,
	"COMPACT":            COMPACT,
	"COMPRESSED":         COMPRESSED,
	"CONNECTION":         CONNECTION,
	"CONSTRAINT":         CONSTRAINT,
	"CREATE":             CREATE,
	"CURRENT_TIMESTAMP":  CURRENT_TIMESTAMP,
	"DATA":               DATA,
	"DATABASE":           DATABASE,
	"DATE":               DATE,
	"DATETIME":           DATETIME,
	"DECIMAL":            DECIMAL,
	"DEFAULT":            DEFAULT,
	"DELAY_KEY_WRITE":    DELAY_KEY_WRITE,
	"DELETE":             DELETE,
	"DIRECTORY":          DIRECTORY,
	"DISK":               DISK,
	"DOUBLE":             DOUBLE,
	"DROP":               DROP,
	"DYNAMIC":            DYNAMIC,
	"ENGINE":             ENGINE,
	"ENUM":               ENUM,
	"EXISTS":             EXISTS,
	"FIRST":              FIRST,
	"FIXED":              FIXED,
	"FLOAT":              FLOAT,
	"FOREIGN":            FOREIGN,
	"FULL":               FULL,
	"FULLTEXT":           FULLTEXT,
	"HASH":               HASH,
	"IF":                 IF,
	"INDEX":              INDEX,
	"INSERT_METHOD":      INSERT_METHOD,
	"INT":                INT,
	"INTEGER":            INTEGER,
	"KEY":                KEY,
	"KEY_BLOCK_SIZE":     KEY_BLOCK_SIZE,
	"LAST":               LAST,
	"LONGBLOB":           LONGBLOB,
	"LONGTEXT":           LONGTEXT,
	"MATCH":              MATCH,
	"MAX_ROWS":           MAX_ROWS,
	"MEDIUMBLOB":         MEDIUMBLOB,
	"MEDIUMINT":          MEDIUMINT,
	"MEDIUMTEXT":         MEDIUMTEXT,
	"MEMORY":             MEMORY,
	"MIN_ROWS":           MIN_ROWS,
	"NO":                 NO,
	"NOT":                NOT,
	"NULL":               NULL,
	"NUMERIC":            NUMERIC,
	"ON":                 ON,
	"PACK_KEYS":          PACK_KEYS,
	"PARTIAL":            PARTIAL,
	"PASSWORD":           PASSWORD,
	"PRIMARY":            PRIMARY,
	"REAL":               REAL,
	"REDUNDANT":          REDUNDANT,
	"REFERENCES":         REFERENCES,
	"RESTRICT":           RESTRICT,
	"ROW_FORMAT":         ROW_FORMAT,
	"SET":                SET,
	"SIMPLE":             SIMPLE,
	"SMALLINT":           SMALLINT,
	"SPARTIAL":           SPARTIAL,
	"STATS_AUTO_RECALC":  STATS_AUTO_RECALC,
	"STATS_PERSISTENT":   STATS_PERSISTENT,
	"STATS_SAMPLE_PAGES": STATS_SAMPLE_PAGES,
	"STORAGE":            STORAGE,
	"TABLE":              TABLE,
	"TABLESPACE":         TABLESPACE,
	"TEMPORARY":          TEMPORARY,
	"TEXT":               TEXT,
	"TIME":               TIME,
	"TIMESTAMP":          TIMESTAMP,
	"TINYBLOB":           TINYBLOB,
	"TINYINT":            TINYINT,
	"TINYTEXT":           TINYTEXT,
	"UNION":              UNION,
	"UNIQUE":             UNIQUE,
	"UNSIGNED":           UNSIGNED,
	"UPDATE":             UPDATE,
	"USE":                USE,
	"USING":              USING,
	"VARBINARY":          VARBINARY,
	"VARCHAR":            VARCHAR,
	"YEAR":               YEAR,
	"ZEROFILL":           ZEROFILL,
}
