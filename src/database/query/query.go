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
	//go:embed scripts/commands/UpdateWorkflow.sql
	UpdateWorkflow string
	//go:embed scripts/commands/DeleteWorkflow.sql
	DeleteWorkflow string
	//go:embed scripts/commands/CreateTemplate.sql
	CreateTemplate string
	//go:embed scripts/commands/DeleteTemplate.sql
	DeleteTemplate string
	//go:embed scripts/commands/CreateField.sql
	CreateField string
	//go:embed scripts/commands/DeleteField.sql
	DeleteField string

	//go:embed scripts/queries/GetApp.sql
	GetApp string
	//go:embed scripts/queries/GetApps.sql
	GetApps string
	//go:embed scripts/queries/GetWorkflow.sql
	GetWorkflow string
	//go:embed scripts/queries/GetWorkflows.sql
	GetWorkflows string
)
