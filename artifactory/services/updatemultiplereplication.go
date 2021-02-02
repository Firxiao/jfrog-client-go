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

type UpdateMultipleReplicationService struct {
	client     *jfroghttpclient.JfrogHttpClient
	ArtDetails auth.ServiceDetails
}

func NewUpdateMultipleReplicationService(client *jfroghttpclient.JfrogHttpClient) *UpdateMultipleReplicationService {
	return &UpdateMultipleReplicationService{client: client}
}

func (rs *UpdateMultipleReplicationService) GetJfrogHttpClient() *jfroghttpclient.JfrogHttpClient {
	return rs.client
}

func (rs *UpdateMultipleReplicationService) performRequest(params *utils.MultipleReplicationBody) error {
	content, err := json.Marshal(params)
	if err != nil {
		return errorutils.CheckError(err)
	}
	httpClientsDetails := rs.ArtDetails.CreateHttpClientDetails()
	utils.SetContentType("application/vnd.org.jfrog.artifactory.replications.MultipleReplicationConfigRequest+json", &httpClientsDetails.Headers)
	var url = rs.ArtDetails.GetUrl() + "api/replications/multiple/" + params.RepoKey
	var resp *http.Response
	var body []byte
	log.Info("Update replication...")
	operationString := "updating"
	resp, body, err = rs.client.SendPost(url, content, &httpClientsDetails)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return errorutils.CheckError(errors.New("Artifactory response: " + resp.Status + "\n" + clientutils.IndentJson(body)))
	}
	log.Debug("Artifactory response:", resp.Status)
	log.Info("Done " + operationString + " repository.")
	return nil
}

func (rs *UpdateMultipleReplicationService) UpdateMultipleReplication(params UpdateMultipleReplicationParams) error {
	return rs.performRequest(utils.CreateMultipleReplicationBody(params.MultipleReplicationParams))
}

func NewUpdateMultipleReplicationParams() UpdateMultipleReplicationParams {
	return UpdateMultipleReplicationParams{}
}

type UpdateMultipleReplicationParams struct {
	utils.MultipleReplicationParams
}