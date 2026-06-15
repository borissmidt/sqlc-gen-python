package python

type Config struct {
	EmitExactTableNames         bool     `json:"emit_exact_table_names"`
	EmitSyncQuerier             bool     `json:"emit_sync_querier"`
	EmitAsyncQuerier            bool     `json:"emit_async_querier"`
	Package                     string   `json:"package"`
	Out                         string   `json:"out"`
	EmitPydanticModels          bool     `json:"emit_pydantic_models"`
	EmitStrEnum                 bool     `json:"emit_str_enum"`
	QueryParameterLimit         *int32   `json:"query_parameter_limit"`
	InflectionExcludeTableNames []string `json:"inflection_exclude_table_names"`

	// DomainOverrides maps a PostgreSQL type name (typically a DOMAIN, whose
	// definition sqlc does not pass to plugins) to a fully-qualified Python
	// type. For example {"job_status": "my.module.JobStatus"} emits a
	// "import my.module" and annotates the column as "my.module.JobStatus".
	DomainOverrides map[string]string `json:"domain_overrides"`
}
