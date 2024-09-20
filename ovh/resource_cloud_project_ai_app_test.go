package ovh

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"text/template"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

var TestAccAiDeployConfig = `
resource "ovh_cloud_project_ai_app" "my_test_app" {
	service_name = "{{.OpenstackTenant}}"
	region = "{{.Region}}"
	resources = {
		cpu = 1
		flavor = "{{.Flavor}}"
	}
	image = "{{.Image}}"
	{{- if .DefaultHttpPort }}
	default_http_port={{.DefaultHttpPort }}
	{{- end}}
	{{- if .DefaultHttpPort }}
	grpc_port={{.GrpcPort }}
	{{- end}}
	{{- if .EnvVars }}
	env_vars = [
	{{- range $env_var := .EnvVars}}
		{
			name = "{{.Name}}"
			{{ if not (eq .Value nil) -}}value = "{{.Value}}"{{ else -}}value = null{{- end }}
		},
	{{- end}}
	]
	{{- end}}
}
`

type AIDeployAppResource struct {
	OpenstackTenant string
	Region          string
	Flavor          string
	Image           string
	EnvVars         []envVar
	DefaultHttpPort *int32
	GrpcPort        *int32
}

type envVar struct {
	Name  string
	Value *string
}

// compileMessageTemplate generate a string from a given string template and Arguments
func compileMessageTemplate(stringTemplate string, data any) string {
	newTemplate, err := template.New("terraform_resource").Parse(stringTemplate)
	if err != nil {
		panic(fmt.Sprintf("Could not create template from %s", stringTemplate))
	}
	var compiledTemplate bytes.Buffer
	err = newTemplate.Execute(&compiledTemplate, data)
	if err != nil {
		panic(err)
	}
	return compiledTemplate.String()
}

func TestAccAiDeploy_basic(t *testing.T) {
	serviceName := os.Getenv("OVH_CLOUD_PROJECT_SERVICE_TEST")

	envVar1Value := "testEnvVarValue"
	envVar2Value := "testEnvVarValue2"
	initialStateEnvVars := []envVar{{
		Name:  "testEnvVarName",
		Value: &envVar1Value,
	}}
	updatedStateEnvVars := []envVar{{
		// setting value to nil should remove it from app env vars
		Name:  "testEnvVarName",
		Value: nil,
	}, {
		Name:  "testEnvVarName2",
		Value: &envVar2Value,
	}}
	initialState := AIDeployAppResource{
		OpenstackTenant: serviceName,
		Region:          "GRA",
		Flavor:          "ai1-1-cpu",
		Image:           "k248cdcu.gra7.container-registry.ovh.net/public/grpc-gateway-server-image:v1",
		EnvVars:         initialStateEnvVars,
	}
	defaultHttpPortUpdate := int32(8081)
	grpcPortUpdate := int32(8082)
	updatedState := AIDeployAppResource{
		OpenstackTenant: initialState.OpenstackTenant,
		Region:          initialState.Region,
		Flavor:          initialState.Flavor,
		Image:           "k248cdcu.gra7.container-registry.ovh.net/public/grpc-gateway-server-image:http-8081-grpc-8082",
		EnvVars:         updatedStateEnvVars,
		DefaultHttpPort: &defaultHttpPortUpdate,
		GrpcPort:        &grpcPortUpdate,
	}

	initialAppConfig := compileMessageTemplate(TestAccAiDeployConfig, initialState)
	updatedAppConfig := compileMessageTemplate(TestAccAiDeployConfig, updatedState)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheckCloud(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: initialAppConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ovh_cloud_project_ai_app.my_test_app", "service_name", initialState.OpenstackTenant),
					resource.TestCheckResourceAttr("ovh_cloud_project_ai_app.my_test_app", "region", initialState.Region),
					resource.TestCheckResourceAttr("ovh_cloud_project_ai_app.my_test_app", "image", initialState.Image),
					resource.TestCheckResourceAttr("ovh_cloud_project_ai_app.my_test_app", "env_vars.0.name", initialStateEnvVars[0].Name),
					resource.TestCheckResourceAttr("ovh_cloud_project_ai_app.my_test_app", "env_vars.0.value", *initialStateEnvVars[0].Value),
				),
			},
			{
				Config: updatedAppConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ovh_cloud_project_ai_app.my_test_app", "service_name", updatedState.OpenstackTenant),
					resource.TestCheckResourceAttr("ovh_cloud_project_ai_app.my_test_app", "region", updatedState.Region),
					resource.TestCheckResourceAttr("ovh_cloud_project_ai_app.my_test_app", "image", updatedState.Image),
					resource.TestCheckResourceAttr("ovh_cloud_project_ai_app.my_test_app", "env_vars.0.name", updatedStateEnvVars[1].Name),
					resource.TestCheckResourceAttr("ovh_cloud_project_ai_app.my_test_app", "env_vars.0.value", *updatedStateEnvVars[1].Value),
					resource.TestCheckResourceAttr("ovh_cloud_project_ai_app.my_test_app", "DefaultHttpPort", fmt.Sprintf("%d", updatedState.DefaultHttpPort)),
					resource.TestCheckResourceAttr("ovh_cloud_project_ai_app.my_test_app", "GrpcPort", fmt.Sprintf("%d", updatedState.GrpcPort)),
				),
			},
		},
	})
}
