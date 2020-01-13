// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1130
	`ALTER`: {
		//line sql.y: 1131
		Category: hGroup,
		//line sql.y: 1132
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER
`,
	},
	//line sql.y: 1147
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1148
		Category: hDDL,
		//line sql.y: 1149
		Text: `
ALTER TABLE [IF EXISTS] <tablename> <command> [, ...]

Commands:
  ALTER TABLE ... ADD [COLUMN] [IF NOT EXISTS] <colname> <type> [<qualifiers...>]
  ALTER TABLE ... ADD <constraint>
  ALTER TABLE ... DROP [COLUMN] [IF EXISTS] <colname> [RESTRICT | CASCADE]
  ALTER TABLE ... DROP CONSTRAINT [IF EXISTS] <constraintname> [RESTRICT | CASCADE]
  ALTER TABLE ... ALTER [COLUMN] <colname> {SET DEFAULT <expr> | DROP DEFAULT}
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP NOT NULL
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP STORED
  ALTER TABLE ... ALTER [COLUMN] <colname> [SET DATA] TYPE <type> [COLLATE <collation>]
  ALTER TABLE ... ALTER PRIMARY KEY USING INDEX <name>
  ALTER TABLE ... RENAME TO <newname>
  ALTER TABLE ... RENAME [COLUMN] <colname> TO <newname>
  ALTER TABLE ... VALIDATE CONSTRAINT <constraintname>
  ALTER TABLE ... SPLIT AT <selectclause> [WITH EXPIRATION <expr>]
  ALTER TABLE ... UNSPLIT AT <selectclause>
  ALTER TABLE ... UNSPLIT ALL
  ALTER TABLE ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]
  ALTER TABLE ... INJECT STATISTICS ...  (experimental)
  ALTER TABLE ... PARTITION BY RANGE ( <name...> ) ( <rangespec> )
  ALTER TABLE ... PARTITION BY LIST ( <name...> ) ( <listspec> )
  ALTER TABLE ... PARTITION BY NOTHING
  ALTER TABLE ... CONFIGURE ZONE <zoneconfig>

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )]
  COLLATE <collationname>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1187
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1201
	`ALTER PARTITION`: {
		ShortDescription: `apply zone configurations to a partition`,
		//line sql.y: 1202
		Category: hDDL,
		//line sql.y: 1203
		Text: `
ALTER PARTITION <name> <command>

Commands:
  -- Alter a single partition which exists on any of a table's indexes.
  ALTER PARTITION <partition> OF TABLE <tablename> CONFIGURE ZONE <zoneconfig>

  -- Alter a partition of a specific index.
  ALTER PARTITION <partition> OF INDEX <tablename>@<indexname> CONFIGURE ZONE <zoneconfig>

  -- Alter all partitions with the same name across a table's indexes.
  ALTER PARTITION <partition> OF INDEX <tablename>@* CONFIGURE ZONE <zoneconfig>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1222
		SeeAlso: `WEBDOCS/configure-zone.html
`,
	},
	//line sql.y: 1227
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1228
		Category: hDDL,
		//line sql.y: 1229
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1231
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1238
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1239
		Category: hDDL,
		//line sql.y: 1240
		Text: `
ALTER SEQUENCE [IF EXISTS] <name>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START <start>]
  [[NO] CYCLE]
ALTER SEQUENCE [IF EXISTS] <name> RENAME TO <newname>
`,
	},
	//line sql.y: 1263
	`ALTER USER`: {
		ShortDescription: `change user properties`,
		//line sql.y: 1264
		Category: hPriv,
		//line sql.y: 1265
		Text: `
ALTER USER [IF EXISTS] <name> WITH PASSWORD <password>
`,
		//line sql.y: 1267
		SeeAlso: `CREATE USER
`,
	},
	//line sql.y: 1272
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1273
		Category: hDDL,
		//line sql.y: 1274
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1276
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1284
	`ALTER RANGE`: {
		ShortDescription: `change the parameters of a range`,
		//line sql.y: 1285
		Category: hDDL,
		//line sql.y: 1286
		Text: `
ALTER RANGE <zonename> <command>

Commands:
  ALTER RANGE ... CONFIGURE ZONE <zoneconfig>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1298
		SeeAlso: `ALTER TABLE
`,
	},
	//line sql.y: 1303
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1304
		Category: hDDL,
		//line sql.y: 1305
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause> [WITH EXPIRATION <expr>]
  ALTER INDEX ... UNSPLIT AT <selectclause>
  ALTER INDEX ... UNSPLIT ALL
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1321
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1821
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1822
		Category: hCCL,
		//line sql.y: 1823
		Text: `
BACKUP <targets...> TO <location...>
       [ AS OF SYSTEM TIME <expr> ]
       [ INCREMENTAL FROM <location...> ]
       [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Location:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1840
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1852
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1853
		Category: hCCL,
		//line sql.y: 1854
		Text: `
RESTORE <targets...> FROM <location...>
        [ AS OF SYSTEM TIME <expr> ]
        [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Locations:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 1870
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1904
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1905
		Category: hCCL,
		//line sql.y: 1906
		Text: `
-- Import both schema and table data:
IMPORT [ TABLE <tablename> FROM ]
       <format> <datafile>
       [ WITH <option> [= <value>] [, ...] ]

-- Import using specific schema, use only table data from external file:
IMPORT TABLE <tablename>
       { ( <elements> ) | CREATE USING <schemafile> }
       <format>
       DATA ( <datafile> [, ...] )
       [ WITH <option> [= <value>] [, ...] ]

Formats:
   CSV
   DELIMITED
   MYSQLDUMP
   PGCOPY
   PGDUMP

Options:
   distributed = '...'
   sstsize = '...'
   temp = '...'
   delimiter = '...'      [CSV, PGCOPY-specific]
   nullif = '...'         [CSV, PGCOPY-specific]
   comment = '...'        [CSV-specific]

`,
		//line sql.y: 1934
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1978
	`EXPORT`: {
		ShortDescription: `export data to file in a distributed manner`,
		//line sql.y: 1979
		Category: hCCL,
		//line sql.y: 1980
		Text: `
EXPORT INTO <format> <datafile> [WITH <option> [= value] [,...]] FROM <query>

Formats:
   CSV

Options:
   delimiter = '...'   [CSV-specific]

`,
		//line sql.y: 1989
		SeeAlso: `SELECT
`,
	},
	//line sql.y: 2083
	`CANCEL`: {
		//line sql.y: 2084
		Category: hGroup,
		//line sql.y: 2085
		Text: `CANCEL JOBS, CANCEL QUERIES, CANCEL SESSIONS
`,
	},
	//line sql.y: 2092
	`CANCEL JOBS`: {
		ShortDescription: `cancel background jobs`,
		//line sql.y: 2093
		Category: hMisc,
		//line sql.y: 2094
		Text: `
CANCEL JOBS <selectclause>
CANCEL JOB <jobid>
`,
		//line sql.y: 2097
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 2115
	`CANCEL QUERIES`: {
		ShortDescription: `cancel running queries`,
		//line sql.y: 2116
		Category: hMisc,
		//line sql.y: 2117
		Text: `
CANCEL QUERIES [IF EXISTS] <selectclause>
CANCEL QUERY [IF EXISTS] <expr>
`,
		//line sql.y: 2120
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 2151
	`CANCEL SESSIONS`: {
		ShortDescription: `cancel open sessions`,
		//line sql.y: 2152
		Category: hMisc,
		//line sql.y: 2153
		Text: `
CANCEL SESSIONS [IF EXISTS] <selectclause>
CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 2156
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 2226
	`CREATE`: {
		//line sql.y: 2227
		Category: hGroup,
		//line sql.y: 2228
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE
`,
	},
	//line sql.y: 2309
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic`,
		//line sql.y: 2310
		Category: hMisc,
		//line sql.y: 2311
		Text: `
CREATE STATISTICS <statisticname>
  [ON <colname> [, ...]]
  FROM <tablename> [AS OF SYSTEM TIME <expr>]
`,
	},
	//line sql.y: 2454
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 2455
		Category: hDML,
		//line sql.y: 2456
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 2460
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 2480
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 2481
		Category: hCfg,
		//line sql.y: 2482
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 2494
	`DROP`: {
		//line sql.y: 2495
		Category: hGroup,
		//line sql.y: 2496
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE
`,
	},
	//line sql.y: 2513
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 2514
		Category: hDDL,
		//line sql.y: 2515
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2516
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2528
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 2529
		Category: hDDL,
		//line sql.y: 2530
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2531
		SeeAlso: `DROP
`,
	},
	//line sql.y: 2543
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 2544
		Category: hDDL,
		//line sql.y: 2545
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2546
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 2558
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 2559
		Category: hDDL,
		//line sql.y: 2560
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2561
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2581
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 2582
		Category: hDDL,
		//line sql.y: 2583
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 2584
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 2604
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2605
		Category: hPriv,
		//line sql.y: 2606
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2607
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 2619
	`DROP ROLE`: {
		ShortDescription: `remove a role`,
		//line sql.y: 2620
		Category: hPriv,
		//line sql.y: 2621
		Text: `DROP ROLE [IF EXISTS] <role> [, ...]
`,
		//line sql.y: 2622
		SeeAlso: `CREATE ROLE, SHOW ROLES
`,
	},
	//line sql.y: 2646
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 2647
		Category: hMisc,
		//line sql.y: 2648
		Text: `
EXPLAIN <statement>
EXPLAIN ([PLAN ,] <planoptions...> ) <statement>
EXPLAIN [ANALYZE] (DISTSQL) <statement>
EXPLAIN ANALYZE [(DISTSQL)] <statement>

Explainable statements:
    SELECT, CREATE, DROP, ALTER, INSERT, UPSERT, UPDATE, DELETE,
    SHOW, EXPLAIN

Plan options:
    TYPES, VERBOSE, OPT

`,
		//line sql.y: 2661
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2744
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2745
		Category: hMisc,
		//line sql.y: 2746
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2747
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2778
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2779
		Category: hMisc,
		//line sql.y: 2780
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2781
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2811
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2812
		Category: hMisc,
		//line sql.y: 2813
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2814
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2834
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2835
		Category: hPriv,
		//line sql.y: 2836
		Text: `
Grant privileges:
  GRANT {ALL | <privileges...> } ON <targets...> TO <grantees...>
Grant role membership (CCL only):
  GRANT <roles...> TO <grantees...> [WITH ADMIN OPTION]

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, ...]
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 2849
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2865
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2866
		Category: hPriv,
		//line sql.y: 2867
		Text: `
Revoke privileges:
  REVOKE {ALL | <privileges...> } ON <targets...> FROM <grantees...>
Revoke role membership (CCL only):
  REVOKE [ADMIN OPTION FOR] <roles...> FROM <grantees...>

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, <databasename>]...
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 2880
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2934
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2935
		Category: hCfg,
		//line sql.y: 2936
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2937
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2949
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2950
		Category: hCfg,
		//line sql.y: 2951
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2952
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2961
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 2962
		Category: hCfg,
		//line sql.y: 2963
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 2966
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2987
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 2988
		Category: hExperimental,
		//line sql.y: 2989
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2997
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 3003
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 3004
		Category: hExperimental,
		//line sql.y: 3005
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 3013
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 3021
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 3022
		Category: hExperimental,
		//line sql.y: 3023
		Text: `
SCRUB TABLE <tablename>
            [AS OF SYSTEM TIME <expr>]
            [WITH OPTIONS <option> [, ...]]

Options:
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX (<index>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT (<constraint>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS PHYSICAL
`,
		//line sql.y: 3034
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 3089
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 3090
		Category: hCfg,
		//line sql.y: 3091
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 3092
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3113
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 3114
		Category: hCfg,
		//line sql.y: 3115
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { on | off | cluster | local | kv | results } [,...]

`,
		//line sql.y: 3121
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 3138
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 3139
		Category: hTxn,
		//line sql.y: 3140
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 3147
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 3339
	`SHOW`: {
		//line sql.y: 3340
		Category: hGroup,
		//line sql.y: 3341
		Text: `
SHOW BACKUP, SHOW CLUSTER SETTING, SHOW COLUMNS, SHOW CONSTRAINTS,
SHOW CREATE, SHOW DATABASES, SHOW HISTOGRAM, SHOW INDEXES, SHOW
PARTITIONS, SHOW JOBS, SHOW QUERIES, SHOW RANGE, SHOW RANGES,
SHOW ROLES, SHOW SCHEMAS, SHOW SEQUENCES, SHOW SESSION, SHOW SESSIONS,
SHOW STATISTICS, SHOW SYNTAX, SHOW TABLES, SHOW TRACE SHOW TRANSACTION, SHOW USERS
`,
	},
	//line sql.y: 3377
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 3378
		Category: hCfg,
		//line sql.y: 3379
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 3380
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 3401
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics (experimental)`,
		//line sql.y: 3402
		Category: hExperimental,
		//line sql.y: 3403
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 3410
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 3423
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram (experimental)`,
		//line sql.y: 3424
		Category: hExperimental,
		//line sql.y: 3425
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 3429
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 3442
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 3443
		Category: hCCL,
		//line sql.y: 3444
		Text: `SHOW BACKUP [SCHEMAS|FILES|RANGES] <location>
`,
		//line sql.y: 3445
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 3480
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 3481
		Category: hCfg,
		//line sql.y: 3482
		Text: `
SHOW CLUSTER SETTING <var>
SHOW [ PUBLIC | ALL ] CLUSTER SETTINGS
`,
		//line sql.y: 3485
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3511
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 3512
		Category: hDDL,
		//line sql.y: 3513
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 3514
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 3522
	`SHOW PARTITIONS`: {
		ShortDescription: `list partition information`,
		//line sql.y: 3523
		Category: hDDL,
		//line sql.y: 3524
		Text: `SHOW PARTITIONS FROM { TABLE <table> | INDEX <index> | DATABASE <database> }
`,
		//line sql.y: 3525
		SeeAlso: `WEBDOCS/show-partitions.html
`,
	},
	//line sql.y: 3545
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 3546
		Category: hDDL,
		//line sql.y: 3547
		Text: `SHOW DATABASES
`,
		//line sql.y: 3548
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 3556
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 3557
		Category: hPriv,
		//line sql.y: 3558
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 3564
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 3577
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 3578
		Category: hDDL,
		//line sql.y: 3579
		Text: `SHOW INDEXES FROM { <tablename> | DATABASE <database_name> } [WITH COMMENT]
`,
		//line sql.y: 3580
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 3610
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 3611
		Category: hDDL,
		//line sql.y: 3612
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 3613
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 3626
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 3627
		Category: hMisc,
		//line sql.y: 3628
		Text: `SHOW [ALL] [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 3629
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 3650
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 3651
		Category: hMisc,
		//line sql.y: 3652
		Text: `
SHOW [AUTOMATIC] JOBS
SHOW JOB <jobid>
`,
		//line sql.y: 3655
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3695
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3696
		Category: hMisc,
		//line sql.y: 3697
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
`,
		//line sql.y: 3699
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3722
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3723
		Category: hMisc,
		//line sql.y: 3724
		Text: `SHOW [ALL] [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3725
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3738
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 3739
		Category: hDDL,
		//line sql.y: 3740
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ] [WITH COMMENT]
`,
		//line sql.y: 3741
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 3773
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 3774
		Category: hDDL,
		//line sql.y: 3775
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 3787
	`SHOW SEQUENCES`: {
		ShortDescription: `list sequences`,
		//line sql.y: 3788
		Category: hDDL,
		//line sql.y: 3789
		Text: `SHOW SEQUENCES [FROM <databasename> ]
`,
	},
	//line sql.y: 3801
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 3802
		Category: hMisc,
		//line sql.y: 3803
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 3812
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 3813
		Category: hCfg,
		//line sql.y: 3814
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 3815
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 3834
	`SHOW CREATE`: {
		ShortDescription: `display the CREATE statement for a table, sequence or view`,
		//line sql.y: 3835
		Category: hDDL,
		//line sql.y: 3836
		Text: `SHOW CREATE [ TABLE | SEQUENCE | VIEW ] <tablename>
`,
		//line sql.y: 3837
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 3855
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3856
		Category: hPriv,
		//line sql.y: 3857
		Text: `SHOW USERS
`,
		//line sql.y: 3858
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 3866
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 3867
		Category: hPriv,
		//line sql.y: 3868
		Text: `SHOW ROLES
`,
		//line sql.y: 3869
		SeeAlso: `CREATE ROLE, DROP ROLE
`,
	},
	//line sql.y: 3925
	`SHOW RANGE`: {
		ShortDescription: `show range information for a row`,
		//line sql.y: 3926
		Category: hMisc,
		//line sql.y: 3927
		Text: `
SHOW RANGE FROM TABLE <tablename> FOR ROW (row, value, ...)
SHOW RANGE FROM INDEX [ <tablename> @ ] <indexname> FOR ROW (row, value, ...)
`,
	},
	//line sql.y: 3948
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 3949
		Category: hMisc,
		//line sql.y: 3950
		Text: `
SHOW RANGES FROM TABLE <tablename>
SHOW RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 4187
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 4188
		Category: hMisc,
		//line sql.y: 4189
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 4192
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 4209
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 4210
		Category: hDDL,
		//line sql.y: 4211
		Text: `
CREATE [[GLOBAL | LOCAL] {TEMPORARY | TEMP}] TABLE [IF NOT EXISTS] <tablename> ( <elements...> ) [<interleave>]
CREATE [[GLOBAL | LOCAL] {TEMPORARY | TEMP}] TABLE [IF NOT EXISTS] <tablename> [( <colnames...> )] AS <source>

Table elements:
   <name> <type> [<qualifiers...>]
   [UNIQUE | INVERTED] INDEX [<name>] ( <colname> [ASC | DESC] [, ...] )
                           [STORING ( <colnames...> )] [<interleave>]
   FAMILY [<name>] ( <colnames...> )
   [CONSTRAINT <name>] <constraint>

Table constraints:
   PRIMARY KEY ( <colnames...> )
   FOREIGN KEY ( <colnames...> ) REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
   UNIQUE ( <colnames... ) [STORING ( <colnames...> )] [<interleave>]
   CHECK ( <expr> )

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
  COLLATE <collationname>
  AS ( <expr> ) STORED

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 4238
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 4976
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 4977
		Category: hDDL,
		//line sql.y: 4978
		Text: `
CREATE [TEMPORARY | TEMP] SEQUENCE <seqname>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START [WITH] <start>]
  [CACHE <cache>]
  [NO CYCLE]
  [VIRTUAL]

`,
		//line sql.y: 4988
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 5053
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 5054
		Category: hDML,
		//line sql.y: 5055
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 5056
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 5064
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 5065
		Category: hPriv,
		//line sql.y: 5066
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 5067
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 5089
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 5090
		Category: hPriv,
		//line sql.y: 5091
		Text: `CREATE ROLE [IF NOT EXISTS] <name>
`,
		//line sql.y: 5092
		SeeAlso: `DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 5110
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 5111
		Category: hDDL,
		//line sql.y: 5112
		Text: `CREATE [TEMPORARY | TEMP] VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 5113
		SeeAlso: `CREATE TABLE, SHOW CREATE, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 5148
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 5149
		Category: hDDL,
		//line sql.y: 5150
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 5158
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 5387
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 5388
		Category: hTxn,
		//line sql.y: 5389
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 5390
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5398
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 5399
		Category: hMisc,
		//line sql.y: 5400
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 5403
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 5420
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 5421
		Category: hTxn,
		//line sql.y: 5422
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 5423
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5438
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 5439
		Category: hTxn,
		//line sql.y: 5440
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 5448
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 5461
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 5462
		Category: hTxn,
		//line sql.y: 5463
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 5466
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 5490
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 5491
		Category: hTxn,
		//line sql.y: 5492
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 5493
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 5611
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 5612
		Category: hDDL,
		//line sql.y: 5613
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 5614
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 5683
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 5684
		Category: hDML,
		//line sql.y: 5685
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5690
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 5709
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 5710
		Category: hDML,
		//line sql.y: 5711
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 5715
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 5826
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 5827
		Category: hDML,
		//line sql.y: 5828
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5835
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 6035
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 6036
		Category: hDML,
		//line sql.y: 6037
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 6048
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 6049
		Category: hDML,
		//line sql.y: 6050
		Text: `
SELECT [DISTINCT [ ON ( <expr> [ , ... ] ) ] ]
       { <expr> [[AS] <name>] | [ [<dbname>.] <tablename>. ] * } [, ...]
       [ FROM <source> ]
       [ WHERE <expr> ]
       [ GROUP BY <expr> [ , ... ] ]
       [ HAVING <expr> ]
       [ WINDOW <name> AS ( <definition> ) ]
       [ { UNION | INTERSECT | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
       [ ORDER BY <expr> [ ASC | DESC ] [, ...] ]
       [ LIMIT { <expr> | ALL } ]
       [ OFFSET <expr> [ ROW | ROWS ] ]
`,
		//line sql.y: 6062
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 6137
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 6138
		Category: hDML,
		//line sql.y: 6139
		Text: `TABLE <tablename>
`,
		//line sql.y: 6140
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6448
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 6449
		Category: hDML,
		//line sql.y: 6450
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 6451
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6560
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 6561
		Category: hDML,
		//line sql.y: 6562
		Text: `
Data sources:
  <tablename> [ @ { <idxname> | <indexflags> } ]
  <tablefunc> ( <exprs...> )
  ( { <selectclause> | <source> } )
  <source> [AS] <alias> [( <colnames...> )]
  <source> [ <jointype> ] JOIN <source> ON <expr>
  <source> [ <jointype> ] JOIN <source> USING ( <colnames...> )
  <source> NATURAL [ <jointype> ] JOIN <source>
  <source> CROSS JOIN <source>
  <source> WITH ORDINALITY
  '[' EXPLAIN ... ']'
  '[' SHOW ... ']'

Index flags:
  '{' FORCE_INDEX = <idxname> [, ...] '}'
  '{' NO_INDEX_JOIN [, ...] '}'
  '{' IGNORE_FOREIGN_KEYS [, ...] '}'

Join types:
  { INNER | { LEFT | RIGHT | FULL } [OUTER] } [ { HASH | MERGE | LOOKUP } ]

`,
		//line sql.y: 6584
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
