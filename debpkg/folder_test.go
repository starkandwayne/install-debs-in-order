package debpkg

import (
	"testing"
)

func TestNewDebianPackagesFromFolder(t *testing.T) {
	folder, err := NewDebianPackagesFromFolder("/app/fixtures/debs/archives/")
	if err != nil {
		t.Error("Should not have error, got: ", err)
	}
	if folder.FolderPath != "/app/fixtures/debs/archives" {
		t.Error("Expected FolderPath to be trimmed, got ", folder.FolderPath)
	}
	if folder.PackageNameToFileNames["tree"] != "tree_1.7.0-5_amd64.deb" {
		t.Error("Expected tree to map to tree_1.7.0-5_amd64.deb, got ", folder.PackageNameToFileNames["tree"])
	}
	treePackage := folder.FileNamesToPackages["tree_1.7.0-5_amd64.deb"]
	if treePackage == nil {
		t.Error("Expected tree_1.7.0-5_amd64.deb to map to a package, got nothing")
	}
	if treePackage.PackageName != "tree" {
		t.Errorf("Expected tree_1.7.0-5_amd64.deb to map to a package called 'tree', got %#v", treePackage)
	}
}

func TestRemovePreinstalledPackages(t *testing.T) {
	folder, err := NewDebianPackagesFromFolder("/app/fixtures/debs/archives/")
	if err != nil {
		t.Error("Should not have error, got: ", err)
	}

	folder.RemovePreinstalledPackages()

	treePackage := folder.FileNamesToPackages["tree_1.7.0-5_amd64.deb"]
	if len(treePackage.UninstalledDependencies) != 0 {
		t.Error("Expected tree package to have no internal dependencies, got", treePackage.UninstalledDependencies)
	}

	dbusPackage := folder.FileNamesToPackages["dbus_1.10.26-0+deb9u1_amd64.deb"]
	if len(dbusPackage.UninstalledDependencies) != 2 {
		t.Error("Expected dbus package to have 2 internal dependencies, got", treePackage.UninstalledDependencies)
	}
}
