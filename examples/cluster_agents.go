//
// Copyright 2021, Timo Furrer <tuxtimo@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func clusterAgentsExample() {
	git, err := gitlab.NewClient("tokengoeshere")
	if err != nil {
		log.Fatal(err)
	}

	projectID := 33
	opt := &gitlab.RegisterAgentOptions{
		Name: gitlab.String("agent-2"),
	}

	// Register Cluster Agent
	clusterAgent, _, err := git.ClusterAgents.RegisterAgent(projectID, opt)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Cluster Agent: %+v\n", clusterAgent)

	// List Cluster Agents
	clusterAgents, _, err := git.ClusterAgents.ListAgents(projectID, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Cluster Agents: %+v", clusterAgents)
}
