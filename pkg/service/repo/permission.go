// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

// Auto generated by 'go run gen_helper.go', DO NOT EDIT.

package repo

import (
	"context"

	"openpitrix.io/openpitrix/pkg/constants"
	"openpitrix.io/openpitrix/pkg/db"
	"openpitrix.io/openpitrix/pkg/gerr"
	"openpitrix.io/openpitrix/pkg/models"
	"openpitrix.io/openpitrix/pkg/pi"
	"openpitrix.io/openpitrix/pkg/util/senderutil"
)

func CheckReposPermission(ctx context.Context, resourceIds []string) ([]*models.Repo, error) {
	if len(resourceIds) == 0 {
		return nil, nil
	}
	var sender = senderutil.GetSenderFromContext(ctx)
	var repos []*models.Repo
	_, err := pi.Global().DB(ctx).
		Select(models.RepoColumns...).
		From(constants.TableRepo).
		Where(db.Eq(constants.ColumnRepoId, resourceIds)).Load(&repos)
	if err != nil {
		return nil, gerr.NewWithDetail(ctx, gerr.Internal, err, gerr.ErrorInternalError)
	}
	if sender != nil && !sender.IsGlobalAdmin() {
		for _, repo := range repos {
			if repo.Owner != sender.UserId {
				return nil, gerr.New(ctx, gerr.PermissionDenied, gerr.ErrorResourceAccessDenied, repo.RepoId)
			}
		}
	}
	if len(repos) == 0 {
		return nil, gerr.New(ctx, gerr.NotFound, gerr.ErrorResourceNotFound, resourceIds)
	}
	return repos, nil
}

func CheckRepoPermission(ctx context.Context, resourceId string) (*models.Repo, error) {
	if len(resourceId) == 0 {
		return nil, nil
	}
	var sender = senderutil.GetSenderFromContext(ctx)
	var repos []*models.Repo
	_, err := pi.Global().DB(ctx).
		Select(models.RepoColumns...).
		From(constants.TableRepo).
		Where(db.Eq(constants.ColumnRepoId, resourceId)).Load(&repos)
	if err != nil {
		return nil, gerr.NewWithDetail(ctx, gerr.Internal, err, gerr.ErrorInternalError)
	}
	if sender != nil && !sender.IsGlobalAdmin() {
		for _, repo := range repos {
			if repo.Owner != sender.UserId {
				return nil, gerr.New(ctx, gerr.PermissionDenied, gerr.ErrorResourceAccessDenied, repo.RepoId)
			}
		}
	}
	if len(repos) == 0 {
		return nil, gerr.New(ctx, gerr.NotFound, gerr.ErrorResourceNotFound, resourceId)
	}
	return repos[0], nil
}
