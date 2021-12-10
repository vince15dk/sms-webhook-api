package scheme

type EvalMatches struct {
	Value  float64 `json:"value"`
	Metric string  `json:"metric"`
	Tags   Tags`json:"tags"`
}

type Tags struct {
	AppKubernetesIoComponent string `json:"app_kubernetes_io_component"`
	AppKubernetesIoInstance  string `json:"app_kubernetes_io_instance"`
	AppKubernetesIoName      string `json:"app_kubernetes_io_name"`
	Container                string `json:"container"`
	Filename                 string `json:"filename"`
	Hostname                 string `json:"hostname"`
	Job                      string `json:"job"`
	Namespace                string `json:"namespace"`
	Pod                      string `json:"pod"`
	PodTemplateHash          string `json:"pod_template_hash"`
	Stream                   string `json:"stream"`
}

type GrafanaLog struct {
	Title       string        `json:"title"`
	RuleID      int           `json:"ruleId"`
	RuleName    string        `json:"ruleName"`
	State       string        `json:"state"`
	EvalMatches []EvalMatches `json:"evalMatches"`
	OrgID       int           `json:"orgId"`
	DashboardID int           `json:"dashboardId"`
	PanelID     int           `json:"panelId"`
	RuleURL     string        `json:"ruleUrl"`
	Message     string        `json:"message"`
}
