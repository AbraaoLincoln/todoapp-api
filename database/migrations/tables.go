package migrations

const (
	ddl_project_table = "CREATE TABLE IF NOT EXISTS project (id INTERGER PRIMARY KEY, name TEXT, color TEXT, created_at TEXT, modified_at TEXT)"
	ddl_task_table    = "CREATE TABLE IF NOT EXISTS task (id INTERGER PRIMARY KEY, name TEXT, status TEXT, description TEXT, created_at TEXT, modified_at TEXT, project_id INTEGER)"
	ddl_subtask_table = "CREATE TABLE IF NOT EXISTS subtask (id INTERGER PRIMARY KEY, name TEXT, status TEXT, created_at TEXT, modified_at TEXT, task_id INTEGER)"
)

func GetTables() []string {
	return []string{ddl_project_table, ddl_task_table, ddl_subtask_table}
}
