package nexusiq

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	restRemediationByApp = "api/v2/components/remediation/application/%s?stageId=%s"
	restRemediationByOrg = "api/v2/components/remediation/organization/%s?stageId=%s"

	remediationTypeNoViolations = "next-no-violations"
	remediationTypeNonFailing   = "next-non-failing"
)

type remediationData struct {
	Component Component `json:"component"`
}

// Equals does a deep equality check against two remediationData objects
func (a *remediationData) Equals(b *remediationData) (_ bool) {
	if a == b {
		return true
	}

	if !a.Component.Equals(&b.Component) {
		return
	}

	return true
}

type remediationVersionChange struct {
	Type string          `json:"type"`
	Data remediationData `json:"data"`
}

// Equals does a deep equality check against two remediationVersionChange objects
func (a *remediationVersionChange) Equals(b *remediationVersionChange) (_ bool) {
	if a == b {
		return true
	}

	if a.Type != b.Type {
		return
	}

	if !a.Data.Equals(&b.Data) {
		return
	}

	return true
}

// Remediation encapsulates the remediation information for a component
type Remediation struct {
	VersionChanges []remediationVersionChange `json:"versionChanges"`
}

// Equals does a deep equality check against two Remediation objects
func (a *Remediation) Equals(b *Remediation) (_ bool) {
	if a == b {
		return true
	}

	if len(a.VersionChanges) != len(b.VersionChanges) {
		return
	}

	for i, v := range a.VersionChanges {
		if !v.Equals(&b.VersionChanges[i]) {
			return
		}
	}

	return true
}

type remediationResponse struct {
	Remediation Remediation `json:"remediation"`
}

func getRemediation(iq IQ, component Component, endpoint string) (Remediation, error) {
	request, err := json.Marshal(component)
	if err != nil {
		return Remediation{}, fmt.Errorf("could not build the request: %v", err)
	}

	body, _, err := iq.Post(endpoint, bytes.NewBuffer(request))
	if err != nil {
		return Remediation{}, fmt.Errorf("could not get remediation: %v", err)
	}

	var results remediationResponse
	if err = json.Unmarshal(body, &results); err != nil {
		return Remediation{}, fmt.Errorf("could not parse remediation response: %v", err)
	}

	return results.Remediation, nil
}

// GetRemediationByApp retrieves the remediation information on a component based on an application's policies
func GetRemediationByApp(iq IQ, component Component, stage, applicationID string) (Remediation, error) {
	app, err := GetApplicationByPublicID(iq, applicationID)
	if err != nil {
		return Remediation{}, fmt.Errorf("could not get application: %v", err)
	}

	endpoint := fmt.Sprintf(restRemediationByApp, app.ID, stage)

	return getRemediation(iq, component, endpoint)
}

// GetRemediationByOrg retrieves the remediation information on a component based on an organization's policies
func GetRemediationByOrg(iq IQ, component Component, stage, organizationName string) (Remediation, error) {
	org, err := GetOrganizationByName(iq, organizationName)
	if err != nil {
		return Remediation{}, fmt.Errorf("could not get organization: %v", err)
	}

	endpoint := fmt.Sprintf(restRemediationByOrg, org.ID, stage)

	return getRemediation(iq, component, endpoint)
}
