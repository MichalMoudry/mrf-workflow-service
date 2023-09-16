package query

import _ "embed"

var (
	//go:embed scripts/commands/CreateApp.sql
	CreateApp string
	//go:embed scripts/commands/DeleteApp.sql
	DeleteApp string
	//go:embed scripts/commands/UpdateApp.sql
	UpdateApp string
	//go:embed scripts/commands/CreateWorkflow.sql
	CreateWorkflow string

	//go:embed scripts/queries/GetApp.sql
	GetApp string
	//go:embed scripts/queries/GetApps.sql
	GetApps string
)
