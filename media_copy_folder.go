package imagekit

import (
	"context"
	"errors"
)

//
// REQUESTS
//

type CopyFolderRequest struct {
	// SourceFolderPath is the full path to the source folder you want to copy.
	//
	// For example - /path/of/source/folder
	SourceFolderPath string `json:"sourceFolderPath"`
	// DestinationPath is the full path to the destination folder where you want to copy the source folder into.
	//
	// For example - /path/of/destination/folder/
	DestinationPath string `json:"destinationPath"`
}

//
// RESPONSES
//

type CopyFolderResponse struct {
	JobID string `json:"jobId"`
}

//
// METHODS
//

// CopyFolder will copy one folder into another.
func (s *MediaService) CopyFolder(ctx context.Context, r *CopyFolderRequest) (*CopyFolderResponse, error) {
	if r == nil {
		return nil, errors.New("request is empty")
	}

	// Prepare request
	req, err := s.client.request("POST", "v1/bulkJobs/copyFolder", nil, requestTypeAPI)
	if err != nil {
		return nil, err
	}

	// Submit the request
	res := new(CopyFolderResponse)

	err = s.client.do(ctx, req, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
