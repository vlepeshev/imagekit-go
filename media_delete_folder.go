package imagekit

import (
	"context"
	"errors"
)

//
// REQUESTS
//

type DeleteFolderRequest struct {
	// FolderPath is a full path to the folder you want to delete.
	//
	// For example folder/to/delete/
	FolderPath string `json:"folderPath"`
}

//
// METHODS
//

// DeleteFolder will delete the specified folder and all nested files & folders.
//
// This action is cannot be undone.
func (s *MediaService) DeleteFolder(ctx context.Context, r *DeleteFolderRequest) error {
	if r == nil {
		return errors.New("request is empty")
	}

	// Prepare request
	req, err := s.client.request("DELETE", "v1/files/folder", nil, requestTypeAPI)
	if err != nil {
		return err
	}

	err = s.client.do(ctx, req, nil)
	if err != nil {
		return err
	}

	return nil
}
