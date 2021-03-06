/*
Copyright 2019 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	metadataClient "github.com/google/knative-gcp/pkg/gclient/metadata"
)

const (
	clusterNameAttr = "cluster-name"
	ProjectIDEnvKey = "PROJECT_ID"
)

// ProjectIDEnvConfig is a struct to parse project ID from env var
type ProjectIDEnvConfig struct {
	ProjectID string `envconfig:"PROJECT_ID"`
}

// ProjectID returns the project ID for a particular resource.
func ProjectID(project string, client metadataClient.Client) (string, error) {
	// If project is set, then return that one.
	if project != "" {
		return project, nil
	}
	// Otherwise, ask GKE metadata server.
	projectID, err := client.ProjectID()
	if err != nil {
		return "", err
	}
	return projectID, nil
}

// ClusterName returns the cluster name for a particular resource.
func ClusterName(clusterName string, client metadataClient.Client) (string, error) {
	// If clusterName is set, then return that one.
	if clusterName != "" {
		return clusterName, nil
	}
	clusterName, err := client.InstanceAttributeValue(clusterNameAttr)
	if err != nil {
		return "", err
	}
	return clusterName, nil
}
