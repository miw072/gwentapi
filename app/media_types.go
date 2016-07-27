//************************************************************************//
// API "gwentapi": Application Media Types
//
// Generated with goagen v0.2.dev, command line:
// $ goagen
// --design=github.com/tri125/gwentapi/design
// --out=$(GOPATH)\src\github.com\tri125\gwentapi
// --version=v0.2.dev
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import (
	"github.com/goadesign/goa"
	"time"
)

// A card (default view)
//
// Identifier: application/vnd.gwentapi.card+json
type GwentapiCard struct {
	// Faction of the card
	Faction *GwentapiFactionLink `form:"faction" json:"faction" xml:"faction"`
	// Flavor text of the card
	Flavor *string `form:"flavor,omitempty" json:"flavor,omitempty" xml:"flavor,omitempty"`
	// API href for making requests on the card
	Href string `form:"href" json:"href" xml:"href"`
	// Unique card ID
	ID string `form:"id" json:"id" xml:"id"`
	// Name of the card
	Name string `form:"name" json:"name" xml:"name"`
	// Rarity of the card
	Rarity *GwentapiRarityLink `form:"rarity" json:"rarity" xml:"rarity"`
	// Rows where the card can be played
	Rows []string `form:"rows,omitempty" json:"rows,omitempty" xml:"rows,omitempty"`
	// Strength of the card
	Strength *int `form:"strength,omitempty" json:"strength,omitempty" xml:"strength,omitempty"`
	// Subtypes of the card
	Subtypes GwentapiTypeLinkCollection `form:"subtypes,omitempty" json:"subtypes,omitempty" xml:"subtypes,omitempty"`
	// Text of the card detailing its abilities and how it plays
	Text *string `form:"text,omitempty" json:"text,omitempty" xml:"text,omitempty"`
	// Type of the card
	Type *GwentapiTypeLink `form:"type" json:"type" xml:"type"`
}

// Validate validates the GwentapiCard media type instance.
func (mt *GwentapiCard) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Type == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}
	if mt.Faction == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "faction"))
	}
	if mt.Rarity == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "rarity"))
	}

	if mt.Faction != nil {
		if err2 := mt.Faction.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.Rarity != nil {
		if err2 := mt.Rarity.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if err2 := mt.Subtypes.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	if mt.Type != nil {
		if err2 := mt.Type.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// A card (link view)
//
// Identifier: application/vnd.gwentapi.card+json
type GwentapiCardLink struct {
	// API href for making requests on the card
	Href string `form:"href" json:"href" xml:"href"`
	// Name of the card
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the GwentapiCardLink media type instance.
func (mt *GwentapiCardLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return
}

// GwentapiCardCollection is the media type for an array of GwentapiCard (default view)
//
// Identifier: application/vnd.gwentapi.card+json; type=collection
type GwentapiCardCollection []*GwentapiCard

// Validate validates the GwentapiCardCollection media type instance.
func (mt GwentapiCardCollection) Validate() (err error) {
	for _, e := range mt {
		if e.ID == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "id"))
		}
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}
		if e.Type == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "type"))
		}
		if e.Faction == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "faction"))
		}
		if e.Rarity == nil {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "rarity"))
		}

		if e.Faction != nil {
			if err2 := e.Faction.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if e.Rarity != nil {
			if err2 := e.Rarity.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
		if err2 := e.Subtypes.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
		if e.Type != nil {
			if err2 := e.Type.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// GwentapiCardCollection is the media type for an array of GwentapiCard (link view)
//
// Identifier: application/vnd.gwentapi.card+json; type=collection
type GwentapiCardLinkCollection []*GwentapiCardLink

// Validate validates the GwentapiCardLinkCollection media type instance.
func (mt GwentapiCardLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}

	}
	return
}

// A card faction (default view)
//
// Identifier: application/vnd.gwentapi.faction+json
type GwentapiFaction struct {
	// API href for making requests on the faction
	Href string `form:"href" json:"href" xml:"href"`
	// Unique faction ID
	ID string `form:"id" json:"id" xml:"id"`
	// Name of the faction
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the GwentapiFaction media type instance.
func (mt *GwentapiFaction) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return
}

// A card faction (link view)
//
// Identifier: application/vnd.gwentapi.faction+json
type GwentapiFactionLink struct {
	// API href for making requests on the faction
	Href string `form:"href" json:"href" xml:"href"`
	// Name of the faction
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the GwentapiFactionLink media type instance.
func (mt *GwentapiFactionLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return
}

// GwentapiFactionCollection is the media type for an array of GwentapiFaction (default view)
//
// Identifier: application/vnd.gwentapi.faction+json; type=collection
type GwentapiFactionCollection []*GwentapiFaction

// Validate validates the GwentapiFactionCollection media type instance.
func (mt GwentapiFactionCollection) Validate() (err error) {
	for _, e := range mt {
		if e.ID == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "id"))
		}
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}

	}
	return
}

// GwentapiFactionCollection is the media type for an array of GwentapiFaction (link view)
//
// Identifier: application/vnd.gwentapi.faction+json; type=collection
type GwentapiFactionLinkCollection []*GwentapiFactionLink

// Validate validates the GwentapiFactionLinkCollection media type instance.
func (mt GwentapiFactionLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}

	}
	return
}

// A glyph (default view)
//
// Identifier: application/vnd.gwentapi.glyph+json
type GwentapiGlyph struct {
	// API href for making requests on the glyph
	Href string `form:"href" json:"href" xml:"href"`
	// Unique glyph ID
	ID string `form:"id" json:"id" xml:"id"`
	// Indicate whether or not the glyph is a weather glyph
	IsWeatherGlyph bool `form:"isWeatherGlyph" json:"isWeatherGlyph" xml:"isWeatherGlyph"`
	// Name of the glyph
	Name string `form:"name" json:"name" xml:"name"`
	// Text of the glyph
	Text string `form:"text" json:"text" xml:"text"`
}

// Validate validates the GwentapiGlyph media type instance.
func (mt *GwentapiGlyph) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Text == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "text"))
	}

	return
}

// A glyph (link view)
//
// Identifier: application/vnd.gwentapi.glyph+json
type GwentapiGlyphLink struct {
	// API href for making requests on the glyph
	Href string `form:"href" json:"href" xml:"href"`
	// Name of the glyph
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the GwentapiGlyphLink media type instance.
func (mt *GwentapiGlyphLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return
}

// GwentapiGlyphCollection is the media type for an array of GwentapiGlyph (default view)
//
// Identifier: application/vnd.gwentapi.glyph+json; type=collection
type GwentapiGlyphCollection []*GwentapiGlyph

// Validate validates the GwentapiGlyphCollection media type instance.
func (mt GwentapiGlyphCollection) Validate() (err error) {
	for _, e := range mt {
		if e.ID == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "id"))
		}
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}
		if e.Text == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "text"))
		}

	}
	return
}

// GwentapiGlyphCollection is the media type for an array of GwentapiGlyph (link view)
//
// Identifier: application/vnd.gwentapi.glyph+json; type=collection
type GwentapiGlyphLinkCollection []*GwentapiGlyphLink

// Validate validates the GwentapiGlyphLinkCollection media type instance.
func (mt GwentapiGlyphLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}

	}
	return
}

// A game patch (default view)
//
// Identifier: application/vnd.gwentapi.patch+json
type GwentapiPatch struct {
	// API href for making requests on the patch
	Href string `form:"href" json:"href" xml:"href"`
	// Unique Patch ID
	ID string `form:"id" json:"id" xml:"id"`
	// Release date of the patch
	ReleaseDate time.Time `form:"releaseDate" json:"releaseDate" xml:"releaseDate"`
	// Patch version
	Version string `form:"version" json:"version" xml:"version"`
}

// Validate validates the GwentapiPatch media type instance.
func (mt *GwentapiPatch) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Version == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}

	return
}

// A game patch (full view)
//
// Identifier: application/vnd.gwentapi.patch+json
type GwentapiPatchFull struct {
	// Patch changelog
	Changelog *string `form:"changelog,omitempty" json:"changelog,omitempty" xml:"changelog,omitempty"`
	// API href for making requests on the patch
	Href string `form:"href" json:"href" xml:"href"`
	// Unique Patch ID
	ID string `form:"id" json:"id" xml:"id"`
	// Release date of the patch
	ReleaseDate time.Time `form:"releaseDate" json:"releaseDate" xml:"releaseDate"`
	// Patch version
	Version string `form:"version" json:"version" xml:"version"`
}

// Validate validates the GwentapiPatchFull media type instance.
func (mt *GwentapiPatchFull) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Version == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}

	return
}

// A game patch (link view)
//
// Identifier: application/vnd.gwentapi.patch+json
type GwentapiPatchLink struct {
	// API href for making requests on the patch
	Href string `form:"href" json:"href" xml:"href"`
	// Patch version
	Version string `form:"version" json:"version" xml:"version"`
}

// Validate validates the GwentapiPatchLink media type instance.
func (mt *GwentapiPatchLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Version == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "version"))
	}

	return
}

// GwentapiPatchCollection is the media type for an array of GwentapiPatch (default view)
//
// Identifier: application/vnd.gwentapi.patch+json; type=collection
type GwentapiPatchCollection []*GwentapiPatch

// Validate validates the GwentapiPatchCollection media type instance.
func (mt GwentapiPatchCollection) Validate() (err error) {
	for _, e := range mt {
		if e.ID == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "id"))
		}
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Version == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "version"))
		}

	}
	return
}

// GwentapiPatchCollection is the media type for an array of GwentapiPatch (full view)
//
// Identifier: application/vnd.gwentapi.patch+json; type=collection
type GwentapiPatchFullCollection []*GwentapiPatchFull

// Validate validates the GwentapiPatchFullCollection media type instance.
func (mt GwentapiPatchFullCollection) Validate() (err error) {
	for _, e := range mt {
		if e.ID == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "id"))
		}
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Version == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "version"))
		}

	}
	return
}

// GwentapiPatchCollection is the media type for an array of GwentapiPatch (link view)
//
// Identifier: application/vnd.gwentapi.patch+json; type=collection
type GwentapiPatchLinkCollection []*GwentapiPatchLink

// Validate validates the GwentapiPatchLinkCollection media type instance.
func (mt GwentapiPatchLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Version == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "version"))
		}

	}
	return
}

// A card rarity (default view)
//
// Identifier: application/vnd.gwentapi.rarity+json
type GwentapiRarity struct {
	// API href for making requests on the rarity
	Href string `form:"href" json:"href" xml:"href"`
	// Unique rarity ID
	ID string `form:"id" json:"id" xml:"id"`
	// Name of the rarity
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the GwentapiRarity media type instance.
func (mt *GwentapiRarity) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return
}

// A card rarity (link view)
//
// Identifier: application/vnd.gwentapi.rarity+json
type GwentapiRarityLink struct {
	// API href for making requests on the rarity
	Href string `form:"href" json:"href" xml:"href"`
	// Name of the rarity
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the GwentapiRarityLink media type instance.
func (mt *GwentapiRarityLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return
}

// GwentapiRarityCollection is the media type for an array of GwentapiRarity (default view)
//
// Identifier: application/vnd.gwentapi.rarity+json; type=collection
type GwentapiRarityCollection []*GwentapiRarity

// Validate validates the GwentapiRarityCollection media type instance.
func (mt GwentapiRarityCollection) Validate() (err error) {
	for _, e := range mt {
		if e.ID == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "id"))
		}
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}

	}
	return
}

// GwentapiRarityCollection is the media type for an array of GwentapiRarity (link view)
//
// Identifier: application/vnd.gwentapi.rarity+json; type=collection
type GwentapiRarityLinkCollection []*GwentapiRarityLink

// Validate validates the GwentapiRarityLinkCollection media type instance.
func (mt GwentapiRarityLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}

	}
	return
}

// Listing of all available resource endpoint (default view)
//
// Identifier: application/vnd.gwentapi.resource+json
type GwentapiResource struct {
	// API href for making requests on cards
	Cards string `form:"cards" json:"cards" xml:"cards"`
	// API href for making requests on factions
	Factions string `form:"factions" json:"factions" xml:"factions"`
	// API href for making requests on glyphs
	Glyphs string `form:"glyphs" json:"glyphs" xml:"glyphs"`
	// API href for making requests on patches
	Patches string `form:"patches" json:"patches" xml:"patches"`
	// API href for making requests on rarities
	Rarities string `form:"rarities" json:"rarities" xml:"rarities"`
	// API href for making requests on types
	Types string `form:"types" json:"types" xml:"types"`
}

// Validate validates the GwentapiResource media type instance.
func (mt *GwentapiResource) Validate() (err error) {
	if mt.Cards == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "cards"))
	}
	if mt.Factions == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "factions"))
	}
	if mt.Glyphs == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "glyphs"))
	}
	if mt.Rarities == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "rarities"))
	}
	if mt.Types == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "types"))
	}
	if mt.Patches == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "patches"))
	}

	return
}

// A card type (default view)
//
// Identifier: application/vnd.gwentapi.type+json
type GwentapiType struct {
	// API href for making requests on the type
	Href string `form:"href" json:"href" xml:"href"`
	// Unique type ID
	ID string `form:"id" json:"id" xml:"id"`
	// Name of the type
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the GwentapiType media type instance.
func (mt *GwentapiType) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return
}

// A card type (link view)
//
// Identifier: application/vnd.gwentapi.type+json
type GwentapiTypeLink struct {
	// API href for making requests on the type
	Href string `form:"href" json:"href" xml:"href"`
	// Name of the type
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the GwentapiTypeLink media type instance.
func (mt *GwentapiTypeLink) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return
}

// GwentapiTypeCollection is the media type for an array of GwentapiType (default view)
//
// Identifier: application/vnd.gwentapi.type+json; type=collection
type GwentapiTypeCollection []*GwentapiType

// Validate validates the GwentapiTypeCollection media type instance.
func (mt GwentapiTypeCollection) Validate() (err error) {
	for _, e := range mt {
		if e.ID == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "id"))
		}
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}

	}
	return
}

// GwentapiTypeCollection is the media type for an array of GwentapiType (link view)
//
// Identifier: application/vnd.gwentapi.type+json; type=collection
type GwentapiTypeLinkCollection []*GwentapiTypeLink

// Validate validates the GwentapiTypeLinkCollection media type instance.
func (mt GwentapiTypeLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e.Href == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "href"))
		}
		if e.Name == "" {
			err = goa.MergeErrors(err, goa.MissingAttributeError(`response[*]`, "name"))
		}

	}
	return
}