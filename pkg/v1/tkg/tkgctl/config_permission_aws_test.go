// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package tkgctl

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"

	"github.com/vmware-tanzu-private/core/pkg/v1/tkg/fakes"
)

var _ = Describe("Unit tests for add region", func() {
	var (
		ctl           *tkgctl
		tkgClient     = &fakes.Client{}
		updaterClient = &fakes.TKGConfigUpdaterClient{}
		err           error
		configDir     string
	)

	JustBeforeEach(func() {
		configDir, err = ioutil.TempDir("", "test")
		err = os.MkdirAll(testingDir, 0o700)
		Expect(err).ToNot(HaveOccurred())
		prepareConfiDir(configDir)
		options := Options{
			ConfigDir: configDir,
		}
		c, createErr := New(options)
		Expect(createErr).ToNot(HaveOccurred())
		ctl, _ = c.(*tkgctl)
		ctl.tkgClient = tkgClient
		ctl.tkgConfigUpdaterClient = updaterClient

		err = ctl.CreateAWSCloudFormationStack("")
	})
	Context("when there is error on ensure configuration file", func() {
		BeforeEach(func() {
			updaterClient.DecodeCredentialsInViperReturns(errors.New("failed to decode"))
		})
		It("should return an error", func() {
			Expect(err).To(HaveOccurred())
		})
	})

	Context("when there is error on creating cloudformation stack ", func() {
		BeforeEach(func() {
			updaterClient.DecodeCredentialsInViperReturns(nil)
			tkgClient.CreateAWSCloudFormationStackReturns(errors.New("failed to create stack"))
		})
		It("should return an error", func() {
			Expect(err).To(HaveOccurred())
		})
	})

	Context("when the stack can be created successfully", func() {
		BeforeEach(func() {
			updaterClient.DecodeCredentialsInViperReturns(nil)
			tkgClient.CreateAWSCloudFormationStackReturns(nil)
		})
		It("should not return an error", func() {
			Expect(err).ToNot(HaveOccurred())
		})
	})

	AfterEach(func() {
		os.Remove(configDir)
	})
})