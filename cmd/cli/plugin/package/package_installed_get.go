// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"github.com/spf13/cobra"
	apierrors "k8s.io/apimachinery/pkg/api/errors"

	"github.com/vmware-tanzu-private/core/pkg/v1/cli/component"
	"github.com/vmware-tanzu-private/core/pkg/v1/tkg/log"
	"github.com/vmware-tanzu-private/core/pkg/v1/tkg/tkgpackageclient"
	"github.com/vmware-tanzu-private/core/pkg/v1/tkg/tkgpackagedatamodel"
)

var packageInstalledGetOp = tkgpackagedatamodel.NewPackageGetOptions()

var packageInstalledGetCmd = &cobra.Command{
	Use:   "get NAME",
	Short: "Get details for an installed package",
	Args:  cobra.ExactArgs(1),
	RunE:  packageInstalledGet,
}

func init() {
	packageInstalledGetCmd.Flags().StringVarP(&packageInstalledGetOp.Namespace, "namespace", "n", "default", "Namespace for installed package CR")
	packageInstalledCmd.AddCommand(packageInstalledGetCmd)
}

func packageInstalledGet(cmd *cobra.Command, args []string) error {
	pkgClient, err := tkgpackageclient.NewTKGPackageClient(packageInstalledGetOp.KubeConfig)
	if err != nil {
		return err
	}
	pkgName = args[0]
	packageInstalledGetOp.PackageName = pkgName
	t, err := component.NewOutputWriterWithSpinner(cmd.OutOrStdout(), outputFormat,
		fmt.Sprintf("Retrieving installation details for %s...", pkgName), true)
	if err != nil {
		return err
	}

	pkg, err := pkgClient.GetPackageInstall(packageInstalledGetOp)

	if err != nil {
		t.StopSpinner()
		if apierrors.IsNotFound(err) {
			log.Infof("failed to find installed package '%s'", pkgName)
		} else {
			return err
		}
	}
	t.AddRow("NAME", pkg.Name)
	t.AddRow("PACKAGE-NAME", pkg.Spec.PackageRef.RefName)
	t.AddRow("PACKAGE-VERSION", pkg.Spec.PackageRef.VersionSelection.Constraints)
	t.AddRow("STATUS", pkg.Status.FriendlyDescription)
	t.AddRow("CONDITIONS", pkg.Status.Conditions)

	t.RenderWithSpinner()

	return nil
}
