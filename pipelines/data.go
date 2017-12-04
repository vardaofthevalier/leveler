package pipelines

type RequiredFile struct {
	Path string,
	Mode string,
	IsDir bool,
}

type IntegrationCmdConfig struct {
	Executable string,
	Args []string,
	RequiredVariables []string,
	RequiredFiles []*RequiredFile,
}

var LocalIntegrationCmdConfig = &IntegrationCmdConfig{
	Executable: "cp",
	Args: []string{
		"{{.Integration.SrcPath}}",
		"{{.Integration.DestPath}}",
	},
}

var gitCredentialsInit = `
#!/bin/bash

# TODO: 
# - ensure that all required variables (HOME, VAULT_TOKEN) are set
# - retrieve git credentials from vault 
# - store SSH keypair under /data/ssh
# NOTE: the SSH key needs to be added to the repository ahead of time (probably via the integration CREATE method)

mkdir -p "${HOME}/.ssh" && \
ssh-keyscan -t rsa bitbucket.org >> "${HOME}/.ssh/known_hosts" && \
mv /data/ssh/id_rsa > "${HOME}/.ssh/id_rsa" && \
chmod 0600 "${HOME}/.ssh/id_rsa" && \
cat /data/ssh/id_rsa.pub > "${HOME}/.ssh/id_rsa.pub" && \
chmod 0644 "${HOME}/.ssh/id_rsa.pub" && \
cat "${HOME}/.ssh/id_rsa.pub" >> "${HOME}/.ssh/authorized_keys" && \
chmod 0700 "${HOME}/.ssh" && \
chmod 0600 "${HOME}/.ssh/authorized_keys"
`

var GitIntegrationCmdConfig = &IntegrationCmdConfig{
	RequiredVariables: []string{
		"HOME",
		"VAULT_TOKEN"
	},
	Init: gitCredentialsInit,
	Executable: "git",
	Args: []string{
		"clone",
		"--init",
		"--recursive",
		"{{.Integration.SrcPath}}",
		"{{.Integration.DestPath}}"
	},
}

var nexusCredentialsInit = `
#!/bin/bash

# TODO: 
# - ensure that all required variables (VAULT_TOKEN) are set
# - retrieve nexus username/password from vault  (NEXUS_USER, NEXUS_PASSWORD)
# - DO ONE OF TWO THINGS: 
#   1) store username and password as variables in a script (for use with curl); or 
# 	2) store username and password in a properties.gradle file for use with gradle 
# 	NOTE: can be provided on the mvn command line if using mvn
`

var NexusIntegrationCmdConfig = &IntegrationCmdConfig{  // TODO: figure out which method of interaction with Nexus is preferred for which scenarios -- This way is the most general, but also most difficult to implement correctly for Java projects.  It may be useful for non-Java projects, though.  
	RequiredVariables: []string{
		"VAULT_TOKEN",
	},
	Init: nexusCredentialsInit,
	Executable: "curl",
	Args: []string{
		"-v",
		"-u",
		"$NEXUS_USER:$NEXUS_PASSWORD",
		"--upload-file",
		"{{.Integration.SrcPath}}",
		"{{.Integration.DestPath}}",
	},
}

var dockerRegistryCredentialsInit = `
#!/bin/bash

# TODO:
# - ensure that all required variables (VAULT_TOKEN) are set
# - retrieve docker repo credentials from vault (certs)
# - write certs to file (location?)
`

var DockerRegistryIntegrationCmdConfig = &IntegrationCmdConfig{
	RequiredVariables: []string{
		"VAULT_TOKEN"
	},
	Init: dockerRegistryCredentialsInit,
	Executable: "docker",
	Args: []string{
		"push",
		"{{.Integration.DestPath}}",
	},
}

var awsCredentialsInit = `
#!/bin/bash

# TODO: 
# - ensure that all required variables (VAULT_TOKEN) are set
# - retrieve aws credentials from vault
# - store aws credentials as variables in a script
`

var AwsIntegrationCmdConfig = &IntegrationCmdConfig{
	RequiredVariables: []string{
		"VAULT_TOKEN",
	},
	Init: awsCredentialsInit,
	Executable: "aws",
	Args: []string{
		"s3",
		"cp",
		"{{.Integration.SrcPath}}",
		"{{.Integration.DestPath}}",
	},
}