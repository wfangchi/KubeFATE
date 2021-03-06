/*
 * Copyright 2019-2021 VMware, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 * http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
package cli

import (
	"github.com/FederatedAI/KubeFATE/k8s-deploy/pkg/api"
	"github.com/urfave/cli/v2"
)

func serviceCommand() *cli.Command {
	return &cli.Command{
		Name:   "service",
		Usage:  "Start KubeFATE as service",
		Action: serviceRun,
	}
}

func serviceRun(c *cli.Context) error {
	return api.Run()
}
