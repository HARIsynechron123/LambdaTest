package test

import (
	"fmt"
	"testing"
	//"time"

	//http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformAwsLambdaFunctionExample(t *testing.T) {
	t.Parallel()

	// Construct the terraform options with default retryable errors to handle the most common
	// retryable errors in terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../LambdaFunc",
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created.
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the IP of the instance
	//lambda := terraform.Output(t, terraformOptions, "lambda")

	// Invoke the function, so we can test its output
	out :=aws.InvokeFunction(t, "eu-west-1", "hello_lambda", "")
 	
	fmt.Printf("%q",out)

	assert.Contains(t, string(out), "Hello from Lambda!")
}