// ----------------
// Minecraft API @see https://docs.curseforge.com/?go#curseforge-core-api-fingerprints
// ----------------

package api

import "github.com/lyssar/msdcli/curseforge/schemas"

// GetMinecraftVersions @see https://docs.curseforge.com/?go#get-minecraft-versions
// TODO Implement
func (api CurseforgeApi) GetMinecraftVersions(sortDescending string) (response schemas.ApiResponseOfListOfMinecraftGameVersion, err error) {
	return
}

// GetSpecificMinecraftVersion @see https://docs.curseforge.com/?go#get-specific-minecraft-version
// TODO Implement
func (api CurseforgeApi) GetSpecificMinecraftVersion(gameVersionString string) (response schemas.ApiResponseOfMinecraftGameVersion, err error) {
	return
}

// GetMinecraftModLoaders @see https://docs.curseforge.com/?go#get-minecraft-modloaders
// TODO Implement
func (api CurseforgeApi) GetMinecraftModLoaders(version string, includeAll string) (response schemas.ApiResponseOfListOfMinecraftModLoaderIndex, err error) {
	return
}

// GetSpecificMinecraftModLoader @see https://docs.curseforge.com/?go#get-specific-minecraft-modloader
// TODO Implement
func (api CurseforgeApi) GetSpecificMinecraftModLoader(modLoaderName string) (response schemas.ApiResponseOfMinecraftModLoaderVersion, err error) {
	return
}
