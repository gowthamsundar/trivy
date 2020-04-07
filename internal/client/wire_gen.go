// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package client

import (
	"context"
	"github.com/aquasecurity/fanal/analyzer"
	"github.com/aquasecurity/fanal/cache"
	"github.com/aquasecurity/fanal/extractor/docker"
	"github.com/aquasecurity/trivy-db/pkg/db"
	"github.com/aquasecurity/trivy/pkg/rpc/client"
	"github.com/aquasecurity/trivy/pkg/scanner"
	"github.com/aquasecurity/trivy/pkg/types"
	"github.com/aquasecurity/trivy/pkg/vulnerability"
	"time"
)

// Injectors from inject.go:

func initializeDockerScanner(ctx context.Context, imageName string, layerCache cache.ImageCache, customHeaders client.CustomHeaders, url client.RemoteURL, timeout time.Duration) (scanner.Scanner, func(), error) {
	scannerScanner := client.NewProtobufClient(url)
	clientScanner := client.NewScanner(customHeaders, scannerScanner)
	dockerOption, err := types.GetDockerOption(timeout)
	if err != nil {
		return scanner.Scanner{}, nil, err
	}
	extractor, cleanup, err := docker.NewDockerExtractor(ctx, imageName, dockerOption)
	if err != nil {
		return scanner.Scanner{}, nil, err
	}
	config := analyzer.New(extractor, layerCache)
	scanner2 := scanner.NewScanner(clientScanner, config)
	return scanner2, func() {
		cleanup()
	}, nil
}

func initializeArchiveScanner(ctx context.Context, filePath string, layerCache cache.ImageCache, customHeaders client.CustomHeaders, url client.RemoteURL, timeout time.Duration) (scanner.Scanner, func(), error) {
	scannerScanner := client.NewProtobufClient(url)
	clientScanner := client.NewScanner(customHeaders, scannerScanner)
	dockerOption, err := types.GetDockerOption(timeout)
	if err != nil {
		return scanner.Scanner{}, nil, err
	}
	extractor, err := docker.NewDockerArchiveExtractor(ctx, filePath, dockerOption)
	if err != nil {
		return scanner.Scanner{}, nil, err
	}
	config := analyzer.New(extractor, layerCache)
	scanner2 := scanner.NewScanner(clientScanner, config)
	return scanner2, func() {
	}, nil
}

func initializeVulnerabilityClient() vulnerability.Client {
	config := db.Config{}
	vulnerabilityClient := vulnerability.NewClient(config)
	return vulnerabilityClient
}
