package imagekit

import (
	"context"
	"errors"
)

//
// REQUESTS
//

type CopyFileRequest struct {
	// SourceFilePath is the full path of the file you want to copy.
	//
	// For example - /path/to/file.jpg
	SourceFilePath string `json:"sourceFilePath"`
	// DestinationPath is the full path to the folder you want to copy the above file into.
	//
	// For example - /folder/to/copy/into/
	DestinationPath string `json:"destinationPath"`
}

//
// METHODS
//

// CopyFile will copy a file from one folder to another.
func (s *MediaService) CopyFile(ctx context.Context, r *CopyFileRequest) error {
	if r == nil {
		return errors.New("request is empty")
	}

	// Prepare request
	req, err := s.client.request("POST", "v1/files/copy", nil, requestTypeAPI)
	if err != nil {
		return err
	}

	err = s.client.do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}
