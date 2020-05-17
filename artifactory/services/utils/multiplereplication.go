package utils

type MultipleReplicationBody struct {
	RepoKey				   string        `json:"repoKey"`
	EnableEventReplication bool          `json:"enableEventReplication"`
	CronExp                string        `json:"cronExp"`
	Replications           []Replication `json:"replications"`
}

type Replication struct {
    Url                    string   `json:"url"`
    Username               string   `json:"username"`
    Password               string   `json:"password"`
    EnableEventReplication bool     `json:"enableEventReplication""`
    SocketTimeoutMillis    int      `json:"socketTimeoutMillis"`
    Enabled                bool     `json:"enabled"`
    SyncDeletes            bool     `json:"syncDeletes"`
    SyncProperties         bool     `json:"syncProperties"`
    SyncStatistics         bool     `json:"syncStatistics"`
    PathPrefix             string   `json:"pathPrefix"`
	RepoKey				   string   `json:"repoKey"`

}

type MultipleReplicationParams struct {
	RepoKey                string
	CronExp                string
	EnableEventReplication bool
	Replications           []Replication
}

func CreateMultipleReplicationBody(params MultipleReplicationParams) *MultipleReplicationBody {
	return &MultipleReplicationBody{
		CronExp:                params.CronExp,
		RepoKey:                params.RepoKey,
		EnableEventReplication: params.EnableEventReplication,
		Replications:           params.Replications,
	}
}
