package tests

import (
	"strings"
	"testing"

	"github.com/jfrog/jfrog-client-go/artifactory/services"
	"github.com/jfrog/jfrog-client-go/artifactory/services/utils"
	"github.com/stretchr/testify/assert"
)

var (
	// TrimSuffix cannot be constants
	// we can declare them as top-level variables
	repoKey string = strings.TrimSuffix(RtTargetRepo, "/")
)

func TestMultipleReplication(t *testing.T) {
	err := createMultipleReplication()
	if err != nil {
		t.Error(err)
	}
	err = getMlitplePushReplication(t, GetMultipleReplicationConfig())
	if err != nil {
		t.Error(err)
	}
	err = deleteMultipleReplication(t)
	if err != nil {
		t.Error(err)
	}
	err = getMultiplePushReplication(t, nil)
	assert.Error(t, err)
}

func createMultipleReplication() error {
	params := services.NewCreateMultipleReplicationParams()
	// Those fields are required
	params.RepoKey = repoKey
	params.CronExp = "0 0/9 14 * * ?"
	params.Username = "anonymous"
	params.Password = "password"
	params.Replications = append(params.Replications,utils.Replication{`
	Url:"http://www.jfrog.com/1"
	Username:               "anonymous",
	Password:               "password",
	EnableEventReplication: false,
	SocketTimeoutMillis:    100,
	Enabled:                true,
	SyncDeletes:            false,
	SyncProperties:         false,
	SyncStatistics:         false,
	PathPrefix:             "",
	 `})
	params.Replications = append(params.Replications,utils.Replication{`
	Url:"http://www.jfrog.com/2"
	Username:               "anonymous",
	Password:               "password",
	EnableEventReplication: false,
	SocketTimeoutMillis:    100,
	Enabled:                true,
	SyncDeletes:            false,
	SyncProperties:         false,
	SyncStatistics:         false,
	PathPrefix:             "",
	 `})
	return testsCreateReplicationService.CreateReplication(params)
}

func getMultiplePushReplication(t *testing.T, expected []utils.MultipleReplicationParams) error {
	multipleReplicationConf, err := testsMultipleReplicationGetService.GetMultipleReplication(repoKey)
	if err != nil {
		return err
	}

	assert.Len(t, expected, 1, "Error in the test input. Probably a bug. Expecting only 1 replication. Got %d.", len(expected))
	assert.Len(t, multipleReplicationConf, 1, "Expected to fetch only 1 replication. Got %d.", len(multipleReplicationConf))

	// Artifactory may return the password encrypted. We therefore remove it,
	// before we can properly compare 'replicationConf' and 'expected'.
	multipleReplicationConf[0].Replications[0].Password = ""
	expected[0].Replications[0].Password = ""

	assert.ElementsMatch(t, multipleReplicationConf, expected)
	return nil
}

func deleteMultipleReplication(t *testing.T) error {
	err := testsMultipleReplicationDeleteService.DeleteMultipleReplication(repoKey, repoUrl)
	if err != nil {
		return err
	}
	return nil
}

func GetMutipleReplicationConfig() []utils.MultipleReplicationParams {
	return []utils.MultipleReplicationParams{
		{
			RepoKey:                repoKey,
			cronExp:"0 0/9 14 * * ?",
			enableEventReplication:true,
			replications: {
			Url:                    "http://www.jfrog.com/1",
			Username:               "anonymous",
			Password:               "password",
			EnableEventReplication: false,
			SocketTimeoutMillis:    100,
			Enabled:                true,
			SyncDeletes:            false,
			SyncProperties:         false,
			SyncStatistics:         false,
			PathPrefix:             "",
		},
		{
			RepoKey:                repoKey,
			cronExp:"0 0/9 14 * * ?",
			enableEventReplication:true,
			replications: {
			Url:                    "http://www.jfrog.com/2",
			Username:               "anonymous",
			Password:               "password",
			EnableEventReplication: false,
			SocketTimeoutMillis:    100,
			Enabled:                true,
			SyncDeletes:            false,
			SyncProperties:         false,
			SyncStatistics:         false,
			PathPrefix:             "",
		}
	}
}
