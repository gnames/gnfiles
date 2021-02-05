package metadata

import (
	"time"
)

type Action int

const (
	NoAction Action = iota
	Upload
	Download
)

type MetaData struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Action Action `json:"action"`
	Info
	PastInfo []Info `json:"pastInfo"`
}

type Info struct {
	ID      string    `json:"cid"`
	SHA     string    `json:"sha"`
	AddTime time.Time `json:"dateTime"`
	Size    int64     `json:"size"`
}

type MetaFiles map[string]*MetaData

func (mfLoc MetaFiles) Sync(mfExo MetaFiles) MetaFiles {
	res := make(map[string]*MetaData)
	paths := unionPaths(mfLoc, mfExo)
	for _, v := range paths {
		valLoc, okLoc := mfLoc[v]
		valExo, okExo := mfExo[v]

		if !okLoc && valExo.Info.SHA != "" {
			pastInfo := append(valExo.PastInfo, valExo.Info)
			res[v] = &MetaData{
				Name:     valExo.Name,
				Path:     valExo.Path,
				Action:   NoAction,
				PastInfo: pastInfo,
			}
			continue
		}

		if !okExo && valLoc.Info.SHA != "" {
			valExo.Info.AddTime = time.Now()
			res[v] = &MetaData{
				Name:     valLoc.Name,
				Path:     valLoc.Path,
				Action:   Upload,
				Info:     valLoc.Info,
				PastInfo: valExo.PastInfo,
			}
			res[v].Info.AddTime = time.Now()
			continue
		}

		if valLoc.SHA != "" && valLoc.SHA == valExo.SHA {
			res[v] = valLoc
			res[v].Action = NoAction
			continue
		}

		if valExo.SHA != "" && valLoc.SHA != valExo.SHA {
			res[v] = &MetaData{
				Name:     valLoc.Name,
				Path:     valLoc.Path,
				Action:   Upload,
				Info:     valLoc.Info,
				PastInfo: append(valExo.PastInfo, valExo.Info),
			}
			continue
		}
	}
	return res
}

func unionPaths(mf1, mf2 MetaFiles) []string {
	pathSet := make(map[string]struct{})

	for k := range mf1 {
		pathSet[k] = struct{}{}
	}
	for k := range mf2 {
		pathSet[k] = struct{}{}
	}

	res := make([]string, 0, len(pathSet))
	for k := range pathSet {
		res = append(res, k)
	}

	return res
}
