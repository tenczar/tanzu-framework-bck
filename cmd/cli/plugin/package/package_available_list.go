// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/vmware-tanzu-private/core/pkg/v1/cli/component"
	"github.com/vmware-tanzu-private/core/pkg/v1/tkg/kappclient"
)

var packageAvailableListCmd = &cobra.Command{
	Use:   "list or list PACKAGE_NAME",
	Short: "List available packages",
	Args:  cobra.MaximumNArgs(1),
	RunE:  packageAvailableList,
}

func init() {
	packageAvailableCmd.AddCommand(packageAvailableListCmd)
}

func packageAvailableList(cmd *cobra.Command, args []string) error {
	kc, err := kappclient.NewKappClient(packageAvailableOp.KubeConfig)
	if err != nil {
		return err
	}
	if packageAvailableOp.AllNamespaces {
		packageAvailableOp.Namespace = ""
	}

	if len(args) == 0 {
		t, err := component.NewOutputWriterWithSpinner(cmd.OutOrStdout(), outputFormat,
			"Retrieving available packages...", true)
		if err != nil {
			return err
		}
		packageMetadataList, err := kc.ListPackageMetadata(packageAvailableOp.Namespace)
		if err != nil {
			return err
		}
		if packageAvailableOp.AllNamespaces {
			t.SetKeys("NAME", "DISPLAY-NAME", "SHORT-DESCRIPTION", "NAMESPACE")
		} else {
			t.SetKeys("NAME", "DISPLAY-NAME", "SHORT-DESCRIPTION")
		}
		for i := range packageMetadataList.Items {
			pkg := packageMetadataList.Items[i]
			if packageAvailableOp.AllNamespaces {
				t.AddRow(pkg.Name, pkg.Spec.DisplayName, pkg.Spec.ShortDescription, pkg.Namespace)
			} else {
				t.AddRow(pkg.Name, pkg.Spec.DisplayName, pkg.Spec.ShortDescription)
			}
		}
		t.RenderWithSpinner()
		return nil
	}
	t, err := component.NewOutputWriterWithSpinner(cmd.OutOrStdout(), outputFormat,
		fmt.Sprintf("Retrieving package versions for %s...", args[0]), true)
	if err != nil {
		return err
	}
	pkgs, err := kc.ListPackages(args[0], packageAvailableOp.Namespace)
	if err != nil {
		return err
	}
	if packageAvailableOp.AllNamespaces {
		t.SetKeys("NAME", "VERSION", "RELEASED-AT", "NAMESPACE")
	} else {
		t.SetKeys("NAME", "VERSION", "RELEASED-AT")
	}
	for i := range pkgs.Items {
		pkg := pkgs.Items[i]
		if packageAvailableOp.AllNamespaces {
			t.AddRow(pkg.Spec.RefName, pkg.Spec.Version, pkg.Spec.ReleasedAt, pkg.Namespace)
		} else {
			t.AddRow(pkg.Spec.RefName, pkg.Spec.Version, pkg.Spec.ReleasedAt)
		}
	}
	t.RenderWithSpinner()
	return nil
}
