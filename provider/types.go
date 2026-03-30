<<<<<<< ours
// Package provider defines optional capability interfaces for music providers.
// Providers implement the base playlist.Provider interface and may additionally
// implement any of the interfaces here to expose extended features (browsing,
// searching, scrobbling, etc.). The UI discovers capabilities at runtime via
// type assertions.
package provider

// AlbumInfo describes an album in a provider's catalog.
type AlbumInfo struct {
	ID        string
	Name      string
	Artist    string
	ArtistID  string
	Year      int
	SongCount int
	Genre     string
}

// ArtistInfo describes an artist in a provider's catalog.
=======
// Package provider defines optional capability interfaces and shared browse
// types used by providers beyond the base playlist.Provider contract.
package provider

// ArtistInfo describes one artist in a provider's browse catalog.
>>>>>>> theirs
type ArtistInfo struct {
	ID         string
	Name       string
	AlbumCount int
}

<<<<<<< ours
// SortType describes one sort option for album listing.
type SortType struct {
	ID    string // e.g. "alphabeticalByName"
	Label string // e.g. "By Name"
}

// ProviderMeta key constants used across providers and the UI.
const (
	MetaNavidromeID = "navidrome.id"
=======
// AlbumInfo describes one album in a provider's browse catalog.
type AlbumInfo struct {
	ID         string
	Name       string
	Artist     string
	ArtistID   string
	Year       int
	TrackCount int
	Genre      string
}

// SortType describes one album sort option exposed by a provider.
type SortType struct {
	ID    string
	Label string
}

// ProviderMeta keys shared across providers and the UI.
const (
	MetaNavidromeID = "navidrome.id"
	MetaJellyfinID  = "jellyfin.id"
>>>>>>> theirs
)
