// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package artifactregistry_test

import (
	"context"

	artifactregistry "cloud.google.com/go/artifactregistry/apiv1"
	artifactregistrypb "cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb"
	"google.golang.org/api/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
)

func ExampleNewClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleNewRESTClient() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewRESTClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleClient_ListDockerImages() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.ListDockerImagesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#ListDockerImagesRequest.
	}
	it := c.ListDockerImages(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GetDockerImage() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.GetDockerImageRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#GetDockerImageRequest.
	}
	resp, err := c.GetDockerImage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListMavenArtifacts() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.ListMavenArtifactsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#ListMavenArtifactsRequest.
	}
	it := c.ListMavenArtifacts(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GetMavenArtifact() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.GetMavenArtifactRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#GetMavenArtifactRequest.
	}
	resp, err := c.GetMavenArtifact(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListNpmPackages() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.ListNpmPackagesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#ListNpmPackagesRequest.
	}
	it := c.ListNpmPackages(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GetNpmPackage() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.GetNpmPackageRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#GetNpmPackageRequest.
	}
	resp, err := c.GetNpmPackage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListPythonPackages() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.ListPythonPackagesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#ListPythonPackagesRequest.
	}
	it := c.ListPythonPackages(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GetPythonPackage() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.GetPythonPackageRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#GetPythonPackageRequest.
	}
	resp, err := c.GetPythonPackage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ImportAptArtifacts() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.ImportAptArtifactsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#ImportAptArtifactsRequest.
	}
	op, err := c.ImportAptArtifacts(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ImportYumArtifacts() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.ImportYumArtifactsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#ImportYumArtifactsRequest.
	}
	op, err := c.ImportYumArtifacts(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListRepositories() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.ListRepositoriesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#ListRepositoriesRequest.
	}
	it := c.ListRepositories(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GetRepository() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.GetRepositoryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#GetRepositoryRequest.
	}
	resp, err := c.GetRepository(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateRepository() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.CreateRepositoryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#CreateRepositoryRequest.
	}
	op, err := c.CreateRepository(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateRepository() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.UpdateRepositoryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#UpdateRepositoryRequest.
	}
	resp, err := c.UpdateRepository(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteRepository() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.DeleteRepositoryRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#DeleteRepositoryRequest.
	}
	op, err := c.DeleteRepository(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ListPackages() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.ListPackagesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#ListPackagesRequest.
	}
	it := c.ListPackages(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GetPackage() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.GetPackageRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#GetPackageRequest.
	}
	resp, err := c.GetPackage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeletePackage() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.DeletePackageRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#DeletePackageRequest.
	}
	op, err := c.DeletePackage(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ListVersions() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.ListVersionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#ListVersionsRequest.
	}
	it := c.ListVersions(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GetVersion() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.GetVersionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#GetVersionRequest.
	}
	resp, err := c.GetVersion(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteVersion() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.DeleteVersionRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#DeleteVersionRequest.
	}
	op, err := c.DeleteVersion(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ListFiles() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.ListFilesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#ListFilesRequest.
	}
	it := c.ListFiles(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GetFile() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.GetFileRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#GetFileRequest.
	}
	resp, err := c.GetFile(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListTags() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.ListTagsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#ListTagsRequest.
	}
	it := c.ListTags(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GetTag() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.GetTagRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#GetTagRequest.
	}
	resp, err := c.GetTag(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateTag() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.CreateTagRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#CreateTagRequest.
	}
	resp, err := c.CreateTag(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateTag() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.UpdateTagRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#UpdateTagRequest.
	}
	resp, err := c.UpdateTag(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteTag() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.DeleteTagRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#DeleteTagRequest.
	}
	err = c.DeleteTag(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_SetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.SetIamPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#SetIamPolicyRequest.
	}
	resp, err := c.SetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetIamPolicy() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.GetIamPolicyRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#GetIamPolicyRequest.
	}
	resp, err := c.GetIamPolicy(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_TestIamPermissions() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &iampb.TestIamPermissionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/iam/v1#TestIamPermissionsRequest.
	}
	resp, err := c.TestIamPermissions(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetProjectSettings() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.GetProjectSettingsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#GetProjectSettingsRequest.
	}
	resp, err := c.GetProjectSettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateProjectSettings() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.UpdateProjectSettingsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#UpdateProjectSettingsRequest.
	}
	resp, err := c.UpdateProjectSettings(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetVPCSCConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.GetVPCSCConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#GetVPCSCConfigRequest.
	}
	resp, err := c.GetVPCSCConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_UpdateVPCSCConfig() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &artifactregistrypb.UpdateVPCSCConfigRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/artifactregistry/apiv1/artifactregistrypb#UpdateVPCSCConfigRequest.
	}
	resp, err := c.UpdateVPCSCConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetLocation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &locationpb.GetLocationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/location#GetLocationRequest.
	}
	resp, err := c.GetLocation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListLocations() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &locationpb.ListLocationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/cloud/location#ListLocationsRequest.
	}
	it := c.ListLocations(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_GetOperation() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := artifactregistry.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &longrunningpb.GetOperationRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/google.golang.org/genproto/googleapis/longrunning#GetOperationRequest.
	}
	resp, err := c.GetOperation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
