package nexusrm

import (
	"bytes"
	"fmt"
	"text/template"
)

const groovyCreateHostedMaven = "repository.createMavenHosted('{{.Name}}'{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateHostedNpm = "repository.createNpmHosted('{{.Name}}'{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateHostedNuget = "repository.createNugetHosted('{{.Name}}'{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateHostedApt = "repository.createAptHosted('{{.Name}}'{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateHostedDocker = "repository.createDockerHosted('{{.Name}}'{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateHostedGolang = "repository.createGolangHosted('{{.Name}}'{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateHostedRaw = "repository.createRawHosted('{{.Name}}'{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateHostedRubygems = "repository.createRubygemsHosted('{{.Name}}'{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateHostedBower = "repository.createBowerHosted('{{.Name}}'{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateHostedPypi = "repository.createPypiHosted('{{.Name}}'{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateHostedYum = "repository.createYumHosted('{{.Name}}'{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateHostedGitLfs = "repository.createGitLfsHosted('{{.Name}}'{{with .BlobStore}}, '{{.}}'{{end}})"

type repositoryHosted struct {
	Name, BlobStore             string
	StrictContentTypeValidation bool
	// versionPolicy VersionPolicy
	// writePolicy WritePolicy
	// layoutPolicy LayoutPolicy
}

const groovyCreateProxyMaven = "repository.createMavenProxy('{{.Name}}'{{with .RemoteURL}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateProxyNpm = "repository.createNpmProxy('{{.Name}}'{{with .RemoteURL}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateProxyNuget = "repository.createNugetProxy('{{.Name}}'{{with .RemoteURL}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateProxyApt = "repository.createAptProxy('{{.Name}}'{{with .RemoteURL}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateProxyDocker = "repository.createDockerProxy('{{.Name}}'{{with .RemoteURL}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateProxyGolang = "repository.createGolangProxy('{{.Name}}'{{with .RemoteURL}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateProxyRaw = "repository.createRawProxy('{{.Name}}'{{with .RemoteURL}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateProxyRubygems = "repository.createRubygemsProxy('{{.Name}}'{{with .RemoteURL}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateProxyBower = "repository.createBowerProxy('{{.Name}}'{{with .RemoteURL}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateProxyPypi = "repository.createPypiProxy('{{.Name}}'{{with .RemoteURL}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateProxyYum = "repository.createYumProxy('{{.Name}}'{{with .RemoteURL}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateProxyGitLfs = "repository.createGitLfsProxy('{{.Name}}'{{with .RemoteURL}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"

type repositoryProxy struct {
	Name, RemoteURL, BlobStore  string
	StrictContentTypeValidation bool
	// versionPolicy VersionPolicy
	// layoutPolicy LayoutPolicy
}

const groovyCreateGroupMaven = "repository.createMavenGroup('{{.Name}}'{{with .Members}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateGroupNpm = "repository.createNpmGroup('{{.Name}}'{{with .Members}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateGroupNuget = "repository.createNugetGroup('{{.Name}}'{{with .Members}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateGroupApt = "repository.createAptGroup('{{.Name}}'{{with .Members}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateGroupDocker = "repository.createDockerGroup('{{.Name}}'{{with .Members}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateGroupGolang = "repository.createGolangGroup('{{.Name}}'{{with .Members}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateGroupRaw = "repository.createRawGroup('{{.Name}}'{{with .Members}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateGroupRubygems = "repository.createRubygemsGroup('{{.Name}}'{{with .Members}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateGroupBower = "repository.createBowerGroup('{{.Name}}'{{with .Members}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateGroupPypi = "repository.createPypiGroup('{{.Name}}'{{with .Members}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateGroupYum = "repository.createYumGroup('{{.Name}}'{{with .Members}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"
const groovyCreateGroupGitLfs = "repository.createGitLfsGroup('{{.Name}}'{{with .Members}}, '{{.}}'{{end}}{{with .BlobStore}}, '{{.}}'{{end}})"

type repositoryGroup struct {
	Name, BlobStore string
	Members         []string
}

// CreateHostedRepository creates a hosted repository of the indicated format
func CreateHostedRepository(rm RM, format repositoryFormat, config repositoryHosted) error {
	var groovyTmpl string
	switch format {
	case Maven:
		groovyTmpl = groovyCreateHostedMaven
	case Npm:
		groovyTmpl = groovyCreateHostedNpm
	case Nuget:
		groovyTmpl = groovyCreateHostedNuget
	case Apt:
		groovyTmpl = groovyCreateHostedApt
	case Docker:
		groovyTmpl = groovyCreateHostedDocker
	case Golang:
		groovyTmpl = groovyCreateHostedGolang
	case Raw:
		groovyTmpl = groovyCreateHostedRaw
	case Rubygems:
		groovyTmpl = groovyCreateHostedRubygems
	case Bower:
		groovyTmpl = groovyCreateHostedBower
	case Pypi:
		groovyTmpl = groovyCreateHostedPypi
	case Yum:
		groovyTmpl = groovyCreateHostedYum
	case GitLfs:
		groovyTmpl = groovyCreateHostedGitLfs
	}

	tmpl, err := template.New("hosted").Parse(groovyTmpl)
	if err != nil {
		return fmt.Errorf("could not parse template: %v", err)
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, config)
	if err != nil {
		return fmt.Errorf("could not create hosted repository from template: %v", err)
	}

	_, err = ScriptRunOnce(rm, newAnonGroovyScript(buf.String()), nil)
	return fmt.Errorf("could not create hosted repository: %v", err)
}

// CreateProxyRepository creates a proxy repository of the indicated format
func CreateProxyRepository(rm RM, format repositoryFormat, config repositoryProxy) error {
	var groovyTmpl string
	switch format {
	case Maven:
		groovyTmpl = groovyCreateProxyMaven
	case Npm:
		groovyTmpl = groovyCreateProxyNpm
	case Nuget:
		groovyTmpl = groovyCreateProxyNuget
	case Apt:
		groovyTmpl = groovyCreateProxyApt
	case Docker:
		groovyTmpl = groovyCreateProxyDocker
	case Golang:
		groovyTmpl = groovyCreateProxyGolang
	case Raw:
		groovyTmpl = groovyCreateProxyRaw
	case Rubygems:
		groovyTmpl = groovyCreateProxyRubygems
	case Bower:
		groovyTmpl = groovyCreateProxyBower
	case Pypi:
		groovyTmpl = groovyCreateProxyPypi
	case Yum:
		groovyTmpl = groovyCreateProxyYum
	case GitLfs:
		groovyTmpl = groovyCreateProxyGitLfs
	}

	tmpl, err := template.New("proxy").Parse(groovyTmpl)
	if err != nil {
		return fmt.Errorf("could not parse template: %v", err)
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, config)
	if err != nil {
		return fmt.Errorf("could not create proxy repository from template: %v", err)
	}

	_, err = ScriptRunOnce(rm, newAnonGroovyScript(buf.String()), nil)
	return fmt.Errorf("could not create proxy repository: %v", err)
}

// CreateGroupRepository creates a group repository of the indicated format
func CreateGroupRepository(rm RM, format repositoryFormat, config repositoryGroup) error {
	var groovyTmpl string
	switch format {
	case Maven:
		groovyTmpl = groovyCreateGroupMaven
	case Npm:
		groovyTmpl = groovyCreateGroupNpm
	case Nuget:
		groovyTmpl = groovyCreateGroupNuget
	case Apt:
		groovyTmpl = groovyCreateGroupApt
	case Docker:
		groovyTmpl = groovyCreateGroupDocker
	case Golang:
		groovyTmpl = groovyCreateGroupGolang
	case Raw:
		groovyTmpl = groovyCreateGroupRaw
	case Rubygems:
		groovyTmpl = groovyCreateGroupRubygems
	case Bower:
		groovyTmpl = groovyCreateGroupBower
	case Pypi:
		groovyTmpl = groovyCreateGroupPypi
	case Yum:
		groovyTmpl = groovyCreateGroupYum
	case GitLfs:
		groovyTmpl = groovyCreateGroupGitLfs
	}

	tmpl, err := template.New("group").Parse(groovyTmpl)
	if err != nil {
		return fmt.Errorf("could not parse template: %v", err)
	}

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, config)
	if err != nil {
		return fmt.Errorf("could not create group repository from template: %v", err)
	}

	_, err = ScriptRunOnce(rm, newAnonGroovyScript(buf.String()), nil)
	return fmt.Errorf("could not create group repository: %v", err)
}
