package main

import (
	"fmt"
	"strings"
	"testing"

	api "k8s.io/api/apps/v1"
	yamlv2 "k8s.io/apimachinery/pkg/util/yaml"
)

func TestHasReadiness(t *testing.T) {
	output, err := deploy(true)
	if err != nil {
		t.Errorf("Error: %v", err)
		return
	}
	section, err := getSection("deployment", output)
	if err != nil {
		t.Errorf("error: %v", err)
		return
	}
	bytes := []byte(strings.Join(section, "\n"))
	deployment := api.Deployment{}
	if err := yamlv2.Unmarshal(bytes, &deployment); err != nil {
		t.Errorf("Error unmarshalling deployment: %v", err)
		return
	}
	if deployment.Spec.Template.Spec.Containers[0].ReadinessProbe == nil {
		t.Errorf("No readiness probe found")
	}
}

func getSection(section string, output []string) ([]string, error) {
	var found = false
	var buffer []string = make([]string, 0)
	marker := "kind: " + strings.Title(section)
	for _, line := range output {
		if strings.HasPrefix(line, "---") {
			if found {
				return buffer, nil
			}
			buffer = make([]string, 0)
		} else if strings.Contains(line, marker) {
			found = true
		} else {
			buffer = append(buffer, line)
		}
	}
	if found {
		return buffer, nil
	} else {
		return nil, fmt.Errorf("Section %s not found", section)
	}
}
