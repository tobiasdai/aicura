package nexus

import (
	"encoding/json"
	"net/url"
)

type MavenComponentService service

/*
func (m *MavenComponentService) list(repoType RepositoryType, unmarshalValue interface{}) error {
	req, err := m.client.get(m.client.appendVersion(repositoriesPath), "")
	if err != nil {
		return err
	}
	var jsonRepos json.RawMessage
	_, err = m.client.do(req, &jsonRepos)
	if err != nil {
		return err
	}
	filtered, err := filterRepository(jsonRepos, RepositoryFormatMaven2, repoType)
	if err != nil {
		return err
	}
	if filtered != nil {
		err = json.Unmarshal(filtered, unmarshalValue)
		return err
	}
	return nil
}*/

func (m *MavenComponentService) ListAll(repositoryName string) (*ComponentResponse, error) {
	components := &ComponentResponse{}

	parsedURL, err := url.ParseQuery("repository=" + repositoryName)
	if err != nil {
		return nil, err
	}

	req, err := m.client.get(m.client.appendVersion(ComponentsPath), parsedURL.Encode())
	if err != nil {
		return nil, err
	}
	var jsonRepos json.RawMessage
	_, err = m.client.do(req, &jsonRepos)
	if err != nil {
		return nil, err
	}

	if jsonRepos != nil {
		err = json.Unmarshal(jsonRepos, &components)
		if err != nil {
			return nil, err
		}
	}
	return components, nil
}

