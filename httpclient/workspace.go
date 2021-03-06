package httpclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nylo-andry/togglsheet"
)

const workspaceURL = "api/v8/workspaces"

// WorkspaceAPI is the client to the workspaces endpoint of Toggl.
type WorkspaceAPI struct {
	baseURL string
	config  *togglsheet.Config
	client  *http.Client
}

// WorkspaceResponse represents what is returned by the workspace endpoint.
type WorkspaceResponse struct {
	Workspaces []togglsheet.Workspace
}

// UnmarshalJSON instructs how to read JSON data.
func (w *WorkspaceResponse) UnmarshalJSON(bs []byte) error {
	return json.Unmarshal(bs, &w.Workspaces)
}

// NewWorkspaceAPI returns a new instance of the WorkspaceAPI
func NewWorkspaceAPI(baseURL string, config *togglsheet.Config, client *http.Client) *WorkspaceAPI {
	return &WorkspaceAPI{baseURL, config, client}
}

// GetWorkspaces returns the workspaces that can be exported.
func (t *WorkspaceAPI) GetWorkspaces() (*WorkspaceResponse, error) {
	req, err := t.buildRequest()

	if err != nil {
		return nil, err
	}

	res, err := t.client.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	workspaceResponse := WorkspaceResponse{}
	err = json.Unmarshal(body, &workspaceResponse)
	if err != nil {
		return nil, err
	}

	return &workspaceResponse, nil
}

func (t *WorkspaceAPI) buildRequest() (*http.Request, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/%v", t.baseURL, workspaceURL), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(t.config.APIToken, "api_token")

	return req, nil
}
