package services

import (
	"errors"
	"net/http"

	"github.com/jfrog/jfrog-client-go/auth"
	"github.com/jfrog/jfrog-client-go/http/jfroghttpclient"

	clientutils "github.com/jfrog/jfrog-client-go/utils"
	"github.com/jfrog/jfrog-client-go/utils/errorutils"
	"github.com/jfrog/jfrog-client-go/utils/log"
)

type DeleteMultipleReplicationService struct {
	client     *jfroghttpclient.JfrogHttpClient
	ArtDetails auth.ServiceDetails
}

func NewDeleteMultipleReplicationService(client *jfroghttpclient.JfrogHttpClient) *DeleteMultipleReplicationService {
	return &DeleteMultipleReplicationService{client: client}
}

func (drs *DeleteMultipleReplicationService) GetJfrogHttpClient() *jfroghttpclient.JfrogHttpClient {
	return drs.client
}

func (drs *DeleteMultipleReplicationService) DeleteMultipleReplication(repoKey, repoUrl string) error {
	httpClientsDetails := drs.ArtDetails.CreateHttpClientDetails()
	log.Info("Deleting multiple replication job...")
	resp, body, err := drs.client.SendDelete(drs.ArtDetails.GetUrl()+"api/replications/"+repoKey+"?url="+repoUrl, nil, &httpClientsDetails)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return errorutils.CheckError(errors.New("Artifactory response: " + resp.Status + "\n" + clientutils.IndentJson(body)))
	}
	log.Debug("Artifactory response:", resp.Status)
	log.Info("Done Deleting multiple replication job.")
	return nil
}
