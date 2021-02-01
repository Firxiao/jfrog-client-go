package services

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/jfrog/jfrog-client-go/artifactory/services/utils"
	"github.com/jfrog/jfrog-client-go/auth"
	"github.com/jfrog/jfrog-client-go/http/jfroghttpclient"
	clientutils "github.com/jfrog/jfrog-client-go/utils"
	"github.com/jfrog/jfrog-client-go/utils/errorutils"
	"github.com/jfrog/jfrog-client-go/utils/log"
)

type GetMultipleReplicationService struct {
	client     *jfroghttpclient.JfrogHttpClient
	ArtDetails auth.ServiceDetails
}

func NewGetMultipleReplicationService(client *jfroghttpclient.JfrogHttpClient) *GetMultipleReplicationService {
	return &GetMultipleReplicationService{client: client}
}

func (drs *GetMultipleReplicationService) GetJfrogHttpClient() *jfroghttpclient.JfrogHttpClient {
	return drs.client
}

func (drs *GetMultipleReplicationService) GetMultipleReplication(repoKey string) ([]utils.MultipleReplicationParams, error) {
	body, err := drs.preform(repoKey)
	if err != nil {
		return nil, err
	}
	var multipleReplicationConf []utils.MultipleReplicationParams
	if err := json.Unmarshal(body, &multipleReplicationConf); err != nil {
		return nil, errorutils.CheckError(err)
	}
	return multipleReplicationConf, nil
}

func (drs *GetMultipleReplicationService) preform(repoKey string) ([]byte, error) {
	httpClientsDetails := drs.ArtDetails.CreateHttpClientDetails()
	log.Info("Retrieve Multi-push replication configuration...")
	resp, body, _, err := drs.client.SendGet(drs.ArtDetails.GetUrl()+"api/replications/"+repoKey, true, &httpClientsDetails)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errorutils.CheckError(errors.New("Artifactory response: " + resp.Status + "\n" + clientutils.IndentJson(body)))
	}
	log.Debug("Artifactory response:", resp.Status)
	log.Info("Done retrieve Multi-push replication job.")
	return body, nil
}
