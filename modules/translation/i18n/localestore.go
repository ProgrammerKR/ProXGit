// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package i18n

import (
	"errors"
	"fmt"
	"html/template"
	"slices"

<<<<<<< HEAD
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/modules/setting"
=======
	"code.proxgit.io/proxgit/modules/log"
	"code.proxgit.io/proxgit/modules/setting"
>>>>>>> master
)

// This file implements the static LocaleStore that will not watch for changes

type locale struct {
	store       *localeStore
	langName    string
	idxToMsgMap map[int]string // the map idx is generated by store's trKeyToIdxMap
}

var _ Locale = (*locale)(nil)

type localeStore struct {
	// After initializing has finished, these fields are read-only.
	langNames []string
	langDescs []string

	localeMap     map[string]*locale
	trKeyToIdxMap map[string]int

	defaultLang string
}

// NewLocaleStore creates a static locale store
func NewLocaleStore() LocaleStore {
	return &localeStore{localeMap: make(map[string]*locale), trKeyToIdxMap: make(map[string]int)}
}

// AddLocaleByIni adds locale by ini into the store
func (store *localeStore) AddLocaleByIni(langName, langDesc string, source, moreSource []byte) error {
	if _, ok := store.localeMap[langName]; ok {
		return errors.New("lang has already been added")
	}

	store.langNames = append(store.langNames, langName)
	store.langDescs = append(store.langDescs, langDesc)

	l := &locale{store: store, langName: langName, idxToMsgMap: make(map[int]string)}
	store.localeMap[l.langName] = l

	iniFile, err := setting.NewConfigProviderForLocale(source, moreSource)
	if err != nil {
		return fmt.Errorf("unable to load ini: %w", err)
	}

	for _, section := range iniFile.Sections() {
		for _, key := range section.Keys() {
			var trKey string
			if section.Name() == "" || section.Name() == "DEFAULT" {
				trKey = key.Name()
			} else {
				trKey = section.Name() + "." + key.Name()
			}
			idx, ok := store.trKeyToIdxMap[trKey]
			if !ok {
				idx = len(store.trKeyToIdxMap)
				store.trKeyToIdxMap[trKey] = idx
			}
			l.idxToMsgMap[idx] = key.Value()
		}
	}

	return nil
}

func (store *localeStore) HasLang(langName string) bool {
	_, ok := store.localeMap[langName]
	return ok
}

func (store *localeStore) ListLangNameDesc() (names, desc []string) {
	return store.langNames, store.langDescs
}

// SetDefaultLang sets default language as a fallback
func (store *localeStore) SetDefaultLang(lang string) {
	store.defaultLang = lang
}

// Locale returns the locale for the lang or the default language
func (store *localeStore) Locale(lang string) (Locale, bool) {
	l, found := store.localeMap[lang]
	if !found {
		var ok bool
		l, ok = store.localeMap[store.defaultLang]
		if !ok {
			// no default - return an empty locale
			l = &locale{store: store, idxToMsgMap: make(map[int]string)}
		}
	}
	return l, found
}

func (store *localeStore) Close() error {
	return nil
}

func (l *locale) TrString(trKey string, trArgs ...any) string {
	format := trKey

	idx, ok := l.store.trKeyToIdxMap[trKey]
	if ok {
		if msg, ok := l.idxToMsgMap[idx]; ok {
			format = msg // use the found translation
		} else if def, ok := l.store.localeMap[l.store.defaultLang]; ok {
			// try to use default locale's translation
			if msg, ok := def.idxToMsgMap[idx]; ok {
				format = msg
			}
		}
	}

	msg, err := Format(format, trArgs...)
	if err != nil {
		log.Error("Error whilst formatting %q in %s: %v", trKey, l.langName, err)
	}
	return msg
}

func (l *locale) TrHTML(trKey string, trArgs ...any) template.HTML {
	args := slices.Clone(trArgs)
	for i, v := range args {
		switch v := v.(type) {
		case nil, bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, template.HTML:
			// for most basic types (including template.HTML which is safe), just do nothing and use it
		case string:
			args[i] = template.HTMLEscapeString(v)
		case fmt.Stringer:
			args[i] = template.HTMLEscapeString(v.String())
		default:
			args[i] = template.HTMLEscapeString(fmt.Sprint(v))
		}
	}
	return template.HTML(l.TrString(trKey, args...))
}

// HasKey returns whether a key is present in this locale or not
func (l *locale) HasKey(trKey string) bool {
	idx, ok := l.store.trKeyToIdxMap[trKey]
	if !ok {
		return false
	}
	_, ok = l.idxToMsgMap[idx]
	return ok
}
