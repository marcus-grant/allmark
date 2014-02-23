// Copyright 2013 Andreas Koch. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debughandler

import (
	"fmt"
	"github.com/andreaskoch/allmark2/common/index"
	"github.com/andreaskoch/allmark2/common/logger"
	"net/http"
)

func New(logger logger.Logger, itemIndex *index.ItemIndex, fileIndex *index.FileIndex) *DebugHandler {
	return &DebugHandler{
		logger:    logger,
		itemIndex: itemIndex,
		fileIndex: fileIndex,
	}
}

type DebugHandler struct {
	logger    logger.Logger
	itemIndex *index.ItemIndex
	fileIndex *index.FileIndex
}

func (handler *DebugHandler) Func() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, route := range handler.itemIndex.Routes() {
			fmt.Fprintf(w, "%q\n", route.Value())
		}

		for _, route := range handler.fileIndex.Routes() {
			fmt.Fprintf(w, "%q\n", route.Value())
		}
	}
}
