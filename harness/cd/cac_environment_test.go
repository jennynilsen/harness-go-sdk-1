package cd

import (
	"fmt"
	"testing"

	"github.com/harness/harness-go-sdk/harness/cd/cac"
	"github.com/harness/harness-go-sdk/harness/utils"
	"github.com/stretchr/testify/require"
)

func TestUpsertEnvironment(t *testing.T) {
	c := getClient()
	appName := fmt.Sprintf("%s-%s-app", t.Name(), utils.RandStringBytes(5))
	envName := fmt.Sprintf("%s-%s-svc", t.Name(), utils.RandStringBytes(5))

	app, err := createApplication(appName)
	require.NoError(t, err)

	env, err := createEnvironment(app.Id, envName)
	require.NoError(t, err)
	require.NotNil(t, env)

	require.Equal(t, env.Name, envName)
	require.Equal(t, env.ApplicationId, app.Id)

	err = c.ConfigAsCodeClient.DeleteEnvironment(app.Name, envName)
	require.NoError(t, err)
}

func TestUpsertEnvironment_WithVariables(t *testing.T) {
	c := getClient()
	appName := fmt.Sprintf("%s-%s-app", t.Name(), utils.RandStringBytes(5))
	envName := fmt.Sprintf("%s-%s-svc", t.Name(), utils.RandStringBytes(5))

	secret, err := createEncryptedTextSecret(envName, "secret_value")
	require.NoError(t, err)

	app, err := createApplication(appName)
	require.NoError(t, err)

	input := cac.NewEntity(cac.ObjectTypes.Environment).(*cac.Environment)
	input.Name = envName
	input.ApplicationId = app.Id
	input.EnvironmentType = cac.EnvironmentTypes.Prod
	input.ConfigMapYamlByServiceTemplateName = &map[string]interface{}{}
	input.VariableOverrides = []*cac.VariableOverride{
		{
			Name:      "var",
			Value:     "foo",
			ValueType: cac.VariableOverrideValueTypes.Text,
		},
		{
			Name:      "secret_variable",
			Value:     secret.Name,
			ValueType: cac.VariableOverrideValueTypes.EncryptedText,
		},
	}

	env, err := c.ConfigAsCodeClient.UpsertEnvironment(input)
	require.NoError(t, err)
	require.Equal(t, env.Name, envName)
	require.Equal(t, env.ApplicationId, app.Id)
	require.Equal(t, 2, len(env.VariableOverrides))

	err = c.ConfigAsCodeClient.DeleteEnvironment(app.Name, envName)
	require.NoError(t, err)
}

func TestGetEnvironmentById(t *testing.T) {
	c := getClient()

	appName := fmt.Sprintf("%s-%s-apps", t.Name(), utils.RandStringBytes(5))
	envName := fmt.Sprintf("%s-%s-svc", t.Name(), utils.RandStringBytes(5))

	app, err := createApplication(appName)
	require.NoError(t, err)

	env, err := createEnvironment(app.Id, envName)
	require.NoError(t, err)
	require.NotNil(t, env)

	envLookup, err := c.ConfigAsCodeClient.GetEnvironmentById(app.Id, env.Id)
	require.NoError(t, err)
	require.NotNil(t, envLookup)
	require.Equal(t, env, envLookup)
}

func TestGetEnvironmentByName(t *testing.T) {
	c := getClient()

	appName := fmt.Sprintf("%s-%s-app", t.Name(), utils.RandStringBytes(5))
	envName := fmt.Sprintf("%s-%s-app", t.Name(), utils.RandStringBytes(5))

	app, err := createApplication(appName)
	require.NoError(t, err)

	env, err := createEnvironment(app.Id, envName)
	require.NoError(t, err)
	require.NotNil(t, env)

	envLookup, err := c.ConfigAsCodeClient.GetEnvironmentByName(app.Id, env.Name)
	require.NoError(t, err)
	require.NotNil(t, envLookup)
	require.Equal(t, env, envLookup)
}

func createEnvironment(applicationId string, name string) (*cac.Environment, error) {
	input := cac.NewEntity(cac.ObjectTypes.Environment).(*cac.Environment)
	input.Name = name
	input.ApplicationId = applicationId
	input.EnvironmentType = cac.EnvironmentTypes.Prod
	input.ConfigMapYamlByServiceTemplateName = &map[string]interface{}{}

	c := getClient()
	return c.ConfigAsCodeClient.UpsertEnvironment(input)

}
