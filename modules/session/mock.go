// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package session

import (
	"net/http"

<<<<<<< HEAD
	"gitea.com/go-chi/session"
=======
	"proxgit.com/go-chi/session"
>>>>>>> master
)

type MockStore struct {
	*session.MemStore
}

func (m *MockStore) Destroy(writer http.ResponseWriter, request *http.Request) error {
	return nil
}

type mockStoreContextKeyStruct struct{}

var MockStoreContextKey = mockStoreContextKeyStruct{}

func NewMockStore(sid string) *MockStore {
	return &MockStore{session.NewMemStore(sid)}
}
