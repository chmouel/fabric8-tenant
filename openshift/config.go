package openshift

import (
	"fmt"
	"net/http"
)

type Config struct {
	MasterURL      string
	MasterUser     string
	Token          string
	HttpTransport  *http.Transport
	TemplateDir    string
	MavenRepoURL   string
	TeamVersion    string
	CheVersion     string
	JenkinsVersion string
	LogCallback    LogCallback
}

type LogCallback func(message string)

func (c Config) WithToken(token string) Config {
	return Config{MasterURL: c.MasterURL, MasterUser: c.MasterUser, Token: token, HttpTransport: c.HttpTransport, TemplateDir: c.TemplateDir, MavenRepoURL: c.MavenRepoURL, TeamVersion: c.TeamVersion}
}

func (c Config) WithUserSettings(cheVersion string, jenkinsVersion string, teamVersion string, mavenRepoURL string) Config {
	if len(cheVersion) > 0 || len(jenkinsVersion) > 0 || len(teamVersion) > 0 || len(mavenRepoURL) > 0 {
		copy := c
		if cheVersion != "" {
			copy.CheVersion = cheVersion
		}
		if jenkinsVersion != "" {
			copy.JenkinsVersion = jenkinsVersion
		}
		if teamVersion != "" {
			copy.TeamVersion = teamVersion
		}
		if mavenRepoURL != "" {
			copy.MavenRepoURL = mavenRepoURL
		}
		return copy
	}
	return c
}

func (c Config) GetLogCallback() LogCallback {
	if c.LogCallback == nil {
		return nilLogCallback
	}
	return c.LogCallback
}

func nilLogCallback(string) {
}

type multiError struct {
	Message string
	Errors  []error
}

func (m multiError) Error() string {
	s := m.Message + "\n"
	for _, err := range m.Errors {
		s += fmt.Sprintf("%v\n", err)
	}
	return s
}

func (m *multiError) String() string {
	return m.Error()
}
