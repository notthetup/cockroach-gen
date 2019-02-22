// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1061
	`ALTER`: {
		//line sql.y: 1062
		Category: hGroup,
		//line sql.y: 1063
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER
`,
	},
	//line sql.y: 1077
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1078
		Category: hDDL,
		//line sql.y: 1079
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
  ALTER TABLE ... RENAME TO <newname>
  ALTER TABLE ... RENAME [COLUMN] <colname> TO <newname>
  ALTER TABLE ... VALIDATE CONSTRAINT <constraintname>
  ALTER TABLE ... SPLIT AT <selectclause>
  ALTER TABLE ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]
  ALTER TABLE ... INJECT STATISTICS ...  (experimental)
  ALTER TABLE ... PARTITION BY RANGE ( <name...> ) ( <rangespec> )
  ALTER TABLE ... PARTITION BY LIST ( <name...> ) ( <listspec> )
  ALTER TABLE ... PARTITION BY NOTHING
  ALTER TABLE ... CONFIGURE ZONE <zoneconfig>
  ALTER PARTITION ... OF TABLE ... CONFIGURE ZONE <zoneconfig>

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
		//line sql.y: 1115
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1129
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1130
		Category: hDDL,
		//line sql.y: 1131
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1133
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1140
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1141
		Category: hDDL,
		//line sql.y: 1142
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
	//line sql.y: 1167
	`ALTER USER`: {
		ShortDescription: `change user properties`,
		//line sql.y: 1168
		Category: hPriv,
		//line sql.y: 1169
		Text: `
ALTER USER [IF EXISTS] <name> WITH PASSWORD <password>
`,
		//line sql.y: 1171
		SeeAlso: `CREATE USER
`,
	},
	//line sql.y: 1176
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1177
		Category: hDDL,
		//line sql.y: 1178
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1180
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1188
	`ALTER RANGE`: {
		ShortDescription: `change the parameters of a range`,
		//line sql.y: 1189
		Category: hDDL,
		//line sql.y: 1190
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
		//line sql.y: 1202
		SeeAlso: `ALTER TABLE
`,
	},
	//line sql.y: 1207
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1208
		Category: hDDL,
		//line sql.y: 1209
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause>
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

`,
		//line sql.y: 1217
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1608
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1609
		Category: hCCL,
		//line sql.y: 1610
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
		//line sql.y: 1627
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1635
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1636
		Category: hCCL,
		//line sql.y: 1637
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
		//line sql.y: 1653
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1671
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1672
		Category: hCCL,
		//line sql.y: 1673
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
   MYSQLOUTFILE
   MYSQLDUMP (mysqldump's SQL output)
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
		//line sql.y: 1701
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1735
	`EXPORT`: {
		ShortDescription: `export data to file in a distributed manner`,
		//line sql.y: 1736
		Category: hCCL,
		//line sql.y: 1737
		Text: `
EXPORT INTO <format> <datafile> [WITH <option> [= value] [,...]] FROM <query>

Formats:
   CSV

Options:
   delimiter = '...'   [CSV-specific]

`,
		//line sql.y: 1746
		SeeAlso: `SELECT
`,
	},
	//line sql.y: 1839
	`CANCEL`: {
		//line sql.y: 1840
		Category: hGroup,
		//line sql.y: 1841
		Text: `CANCEL JOBS, CANCEL QUERIES, CANCEL SESSIONS
`,
	},
	//line sql.y: 1848
	`CANCEL JOBS`: {
		ShortDescription: `cancel background jobs`,
		//line sql.y: 1849
		Category: hMisc,
		//line sql.y: 1850
		Text: `
CANCEL JOBS <selectclause>
CANCEL JOB <jobid>
`,
		//line sql.y: 1853
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 1871
	`CANCEL QUERIES`: {
		ShortDescription: `cancel running queries`,
		//line sql.y: 1872
		Category: hMisc,
		//line sql.y: 1873
		Text: `
CANCEL QUERIES [IF EXISTS] <selectclause>
CANCEL QUERY [IF EXISTS] <expr>
`,
		//line sql.y: 1876
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 1907
	`CANCEL SESSIONS`: {
		ShortDescription: `cancel open sessions`,
		//line sql.y: 1908
		Category: hMisc,
		//line sql.y: 1909
		Text: `
CANCEL SESSIONS [IF EXISTS] <selectclause>
CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 1912
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 1985
	`CREATE`: {
		//line sql.y: 1986
		Category: hGroup,
		//line sql.y: 1987
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE
`,
	},
	//line sql.y: 2068
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic (experimental)`,
		//line sql.y: 2069
		Category: hExperimental,
		//line sql.y: 2070
		Text: `
CREATE STATISTICS <statisticname>
  [ON <colname> [, ...]]
  FROM <tablename> [AS OF SYSTEM TIME <expr>]
`,
	},
	//line sql.y: 2160
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 2161
		Category: hDML,
		//line sql.y: 2162
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 2166
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 2181
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 2182
		Category: hCfg,
		//line sql.y: 2183
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 2195
	`DROP`: {
		//line sql.y: 2196
		Category: hGroup,
		//line sql.y: 2197
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE
`,
	},
	//line sql.y: 2214
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 2215
		Category: hDDL,
		//line sql.y: 2216
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2217
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2229
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 2230
		Category: hDDL,
		//line sql.y: 2231
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2232
		SeeAlso: `DROP
`,
	},
	//line sql.y: 2244
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 2245
		Category: hDDL,
		//line sql.y: 2246
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2247
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 2259
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 2260
		Category: hDDL,
		//line sql.y: 2261
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2262
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2282
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 2283
		Category: hDDL,
		//line sql.y: 2284
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 2285
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 2305
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2306
		Category: hPriv,
		//line sql.y: 2307
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2308
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 2320
	`DROP ROLE`: {
		ShortDescription: `remove a role`,
		//line sql.y: 2321
		Category: hPriv,
		//line sql.y: 2322
		Text: `DROP ROLE [IF EXISTS] <role> [, ...]
`,
		//line sql.y: 2323
		SeeAlso: `CREATE ROLE, SHOW ROLES
`,
	},
	//line sql.y: 2347
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 2348
		Category: hMisc,
		//line sql.y: 2349
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
		//line sql.y: 2362
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2422
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2423
		Category: hMisc,
		//line sql.y: 2424
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2425
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2456
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2457
		Category: hMisc,
		//line sql.y: 2458
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2459
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2489
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2490
		Category: hMisc,
		//line sql.y: 2491
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2492
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2512
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2513
		Category: hPriv,
		//line sql.y: 2514
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
		//line sql.y: 2527
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2543
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2544
		Category: hPriv,
		//line sql.y: 2545
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
		//line sql.y: 2558
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2613
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2614
		Category: hCfg,
		//line sql.y: 2615
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2616
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2628
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2629
		Category: hCfg,
		//line sql.y: 2630
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2631
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2640
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 2641
		Category: hCfg,
		//line sql.y: 2642
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 2645
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2666
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 2667
		Category: hExperimental,
		//line sql.y: 2668
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2676
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 2682
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 2683
		Category: hExperimental,
		//line sql.y: 2684
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2692
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 2700
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 2701
		Category: hExperimental,
		//line sql.y: 2702
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
		//line sql.y: 2713
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 2769
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 2770
		Category: hCfg,
		//line sql.y: 2771
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 2772
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2793
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 2794
		Category: hCfg,
		//line sql.y: 2795
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { on | off | cluster | local | kv | results } [,...]

`,
		//line sql.y: 2801
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2818
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 2819
		Category: hTxn,
		//line sql.y: 2820
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 2827
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 3010
	`SHOW`: {
		//line sql.y: 3011
		Category: hGroup,
		//line sql.y: 3012
		Text: `
SHOW BACKUP, SHOW CLUSTER SETTING, SHOW COLUMNS, SHOW CONSTRAINTS,
SHOW CREATE, SHOW DATABASES, SHOW HISTOGRAM, SHOW INDEXES, SHOW JOBS,
SHOW QUERIES, SHOW ROLES, SHOW SESSION, SHOW SESSIONS, SHOW STATISTICS,
SHOW SYNTAX, SHOW TABLES, SHOW TRACE SHOW TRANSACTION, SHOW USERS
`,
	},
	//line sql.y: 3044
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 3045
		Category: hCfg,
		//line sql.y: 3046
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 3047
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 3068
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics (experimental)`,
		//line sql.y: 3069
		Category: hExperimental,
		//line sql.y: 3070
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 3077
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 3092
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram (experimental)`,
		//line sql.y: 3093
		Category: hExperimental,
		//line sql.y: 3094
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 3098
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 3112
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 3113
		Category: hCCL,
		//line sql.y: 3114
		Text: `SHOW BACKUP [FILES|RANGES] <location>
`,
		//line sql.y: 3115
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 3142
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 3143
		Category: hCfg,
		//line sql.y: 3144
		Text: `
SHOW CLUSTER SETTING <var>
SHOW ALL CLUSTER SETTINGS
`,
		//line sql.y: 3147
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3164
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 3165
		Category: hDDL,
		//line sql.y: 3166
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 3167
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 3176
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 3177
		Category: hDDL,
		//line sql.y: 3178
		Text: `SHOW DATABASES
`,
		//line sql.y: 3179
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 3187
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 3188
		Category: hPriv,
		//line sql.y: 3189
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 3195
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 3208
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 3209
		Category: hDDL,
		//line sql.y: 3210
		Text: `SHOW INDEXES FROM <tablename>
`,
		//line sql.y: 3211
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 3232
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 3233
		Category: hDDL,
		//line sql.y: 3234
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 3235
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 3250
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 3251
		Category: hMisc,
		//line sql.y: 3252
		Text: `SHOW [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 3253
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 3269
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 3270
		Category: hMisc,
		//line sql.y: 3271
		Text: `SHOW JOBS
`,
		//line sql.y: 3272
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3280
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3281
		Category: hMisc,
		//line sql.y: 3282
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
`,
		//line sql.y: 3284
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3307
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3308
		Category: hMisc,
		//line sql.y: 3309
		Text: `SHOW [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3310
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3326
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 3327
		Category: hDDL,
		//line sql.y: 3328
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ]
`,
		//line sql.y: 3329
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 3361
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 3362
		Category: hDDL,
		//line sql.y: 3363
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 3375
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 3376
		Category: hMisc,
		//line sql.y: 3377
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 3386
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 3387
		Category: hCfg,
		//line sql.y: 3388
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 3389
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 3408
	`SHOW CREATE`: {
		ShortDescription: `display the CREATE statement for a table, sequence or view`,
		//line sql.y: 3409
		Category: hDDL,
		//line sql.y: 3410
		Text: `SHOW CREATE [ TABLE | SEQUENCE | VIEW ] <tablename>
`,
		//line sql.y: 3411
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 3431
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3432
		Category: hPriv,
		//line sql.y: 3433
		Text: `SHOW USERS
`,
		//line sql.y: 3434
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 3442
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 3443
		Category: hPriv,
		//line sql.y: 3444
		Text: `SHOW ROLES
`,
		//line sql.y: 3445
		SeeAlso: `CREATE ROLE, DROP ROLE
`,
	},
	//line sql.y: 3492
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 3493
		Category: hMisc,
		//line sql.y: 3494
		Text: `
SHOW EXPERIMENTAL_RANGES FROM TABLE <tablename>
SHOW EXPERIMENTAL_RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 3732
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 3733
		Category: hMisc,
		//line sql.y: 3734
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 3737
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3754
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 3755
		Category: hDDL,
		//line sql.y: 3756
		Text: `
CREATE TABLE [IF NOT EXISTS] <tablename> ( <elements...> ) [<interleave>]
CREATE TABLE [IF NOT EXISTS] <tablename> [( <colnames...> )] AS <source>

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
		//line sql.y: 3783
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 4371
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 4372
		Category: hDDL,
		//line sql.y: 4373
		Text: `
CREATE SEQUENCE <seqname>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START [WITH] <start>]
  [CACHE <cache>]
  [NO CYCLE]
  [VIRTUAL]

`,
		//line sql.y: 4383
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 4430
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 4431
		Category: hDML,
		//line sql.y: 4432
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 4433
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 4441
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 4442
		Category: hPriv,
		//line sql.y: 4443
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 4444
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 4466
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 4467
		Category: hPriv,
		//line sql.y: 4468
		Text: `CREATE ROLE [IF NOT EXISTS] <name>
`,
		//line sql.y: 4469
		SeeAlso: `DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 4487
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 4488
		Category: hDDL,
		//line sql.y: 4489
		Text: `CREATE VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 4490
		SeeAlso: `CREATE TABLE, SHOW CREATE, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 4524
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 4525
		Category: hDDL,
		//line sql.y: 4526
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 4534
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 4776
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 4777
		Category: hTxn,
		//line sql.y: 4778
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 4779
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 4787
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 4788
		Category: hMisc,
		//line sql.y: 4789
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 4792
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 4809
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 4810
		Category: hTxn,
		//line sql.y: 4811
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 4812
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 4827
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 4828
		Category: hTxn,
		//line sql.y: 4829
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 4837
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 4850
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 4851
		Category: hTxn,
		//line sql.y: 4852
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 4855
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 4879
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 4880
		Category: hTxn,
		//line sql.y: 4881
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 4882
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 5000
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 5001
		Category: hDDL,
		//line sql.y: 5002
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 5003
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 5072
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 5073
		Category: hDML,
		//line sql.y: 5074
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5079
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 5098
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 5099
		Category: hDML,
		//line sql.y: 5100
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 5104
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 5211
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 5212
		Category: hDML,
		//line sql.y: 5213
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5220
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 5394
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 5395
		Category: hDML,
		//line sql.y: 5396
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 5407
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 5408
		Category: hDML,
		//line sql.y: 5409
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
		//line sql.y: 5421
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 5496
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 5497
		Category: hDML,
		//line sql.y: 5498
		Text: `TABLE <tablename>
`,
		//line sql.y: 5499
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 5779
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 5780
		Category: hDML,
		//line sql.y: 5781
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 5782
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 5887
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 5888
		Category: hDML,
		//line sql.y: 5889
		Text: `
Data sources:
  <tablename> [ @ { <idxname> | <indexhint> } ]
  <tablefunc> ( <exprs...> )
  ( { <selectclause> | <source> } )
  <source> [AS] <alias> [( <colnames...> )]
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> ON <expr>
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> USING ( <colnames...> )
  <source> NATURAL { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source>
  <source> CROSS JOIN <source>
  <source> WITH ORDINALITY
  '[' EXPLAIN ... ']'
  '[' SHOW ... ']'

Index flags:
  '{' FORCE_INDEX = <idxname> [, ...] '}'
  '{' NO_INDEX_JOIN [, ...] '}'

`,
		//line sql.y: 5907
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
