package paths

import "path/filepath"

func RootPath(root, dir string) string {
	return filepath.Join(root, dir)
}

func IPFSPath(id string) string {
	return "/ipfs/" + id
}

func IPNSPath(keyID string) string {
	return "/ipns/" + keyID
}

func MetaPath(dir string) string {
	return filepath.Join(dir, "_META.json")
}
