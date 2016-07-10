package controllers

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var DBCon *sql.DB

func NewDBConnection(DSN string) (*sql.DB, error) {
	var dbError error
	var db *sql.DB

	db, dbError = sql.Open("mysql", DSN)
	return db, dbError
}

func FetchType(id string) (TypeModel, error) {
	var myType TypeModel
	row := DBCon.QueryRow("SELECT id, name FROM Types WHERE id=?", id)
	err := row.Scan(&myType.ID, &myType.Name)
	if err != nil {
		myType.ID = ""
		myType.Name = ""
	}

	return myType, err
}

func FetchRow(id string) (RowModel, error) {
	var combatRow RowModel
	row := DBCon.QueryRow("SELECT id, name FROM Rows WHERE id=?", id)
	err := row.Scan(&combatRow.ID, &combatRow.Name)
	if err != nil {
		combatRow.ID = ""
		combatRow.Name = ""
	}

	return combatRow, err
}

func FetchRarity(id string) (RarityModel, error) {
	var rarity RarityModel
	row := DBCon.QueryRow("SELECT id, name FROM Rarities WHERE id=?", id)
	err := row.Scan(&rarity.ID, &rarity.Name)
	if err != nil {
		rarity.ID = ""
		rarity.Name = ""
	}

	return rarity, err
}

func FetchPatch(id string) (PatchModel, error) {
	var patch PatchModel
	var changeLog sql.NullString
	var emptyString = ""

	row := DBCon.QueryRow("SELECT id, version, releaseDate, changelog FROM Patches WHERE id=?", id)
	err := row.Scan(&patch.ID, &patch.Version, &patch.ReleaseDate, &changeLog)
	if err != nil {
		patch.ID = ""
		patch.Version = ""
		patch.ReleaseDate = time.Time{}
		patch.Changelog = &emptyString
	}

	if changeLog.Valid {
		patch.Changelog = &changeLog.String
	} else {
		patch.Changelog = &emptyString
	}

	return patch, err
}

func FetchLatestPatch() (PatchModel, error) {
	var patch PatchModel
	var changeLog sql.NullString
	var emptyString = ""

	row := DBCon.QueryRow("SELECT id, version, releaseDate, changelog FROM Patches WHERE releaseDate = (SELECT MAX(releaseDate) FROM Patches)")
	err := row.Scan(&patch.ID, &patch.Version, &patch.ReleaseDate, &changeLog)
	if err != nil {
		patch.ID = ""
		patch.Version = ""
		patch.ReleaseDate = time.Time{}
		patch.Changelog = &emptyString
	}

	if changeLog.Valid {
		patch.Changelog = &changeLog.String
	} else {
		patch.Changelog = &emptyString
	}

	return patch, err
}

func FetchGlyph(id string) (GlyphModel, error) {
	var glyph GlyphModel
	row := DBCon.QueryRow("SELECT id, name, isWeatherGlyph, text FROM Glyphs WHERE id=?", id)
	err := row.Scan(&glyph.ID, &glyph.Name, &glyph.IsWeatherGlyph, &glyph.Text)
	if err != nil {
		glyph.ID = ""
		glyph.Name = ""
		glyph.IsWeatherGlyph = false
		glyph.Text = ""
	}

	return glyph, err
}

func FetchFaction(id string) (FactionModel, error) {
	var faction FactionModel
	row := DBCon.QueryRow("SELECT id, name FROM Factions WHERE id=?", id)
	err := row.Scan(&faction.ID, &faction.Name)
	if err != nil {
		faction.ID = ""
		faction.Name = ""
	}

	return faction, err
}

func FetchIllustrator(id string) (IllustratorModel, error) {
	var illustrator IllustratorModel
	row := DBCon.QueryRow("SELECT id, name FROM Illustrators WHERE id=?", id)
	err := row.Scan(&illustrator.ID, &illustrator.Name)
	if err != nil {
		illustrator.ID = ""
		illustrator.Name = ""
	}

	return illustrator, err
}

func FetchCard(id string) (CardModel, error) {
	var card CardModel
	var rarity RarityModel
	var faction FactionModel
	var cardType TypeModel

	var flavor, text sql.NullString
	var strength sql.NullInt64

	row := DBCon.QueryRow("SELECT c.name, c.id, r.id, r.name, f.id, f.name, t.id, t.name, strength, text, flavor FROM Cards AS c INNER JOIN Rarities AS r ON c.idRarity = r.idRarity INNER JOIN Factions AS f ON c.idFaction = f.idFaction INNER JOIN Types AS t ON c.idType = t.idType WHERE c.id =?", id)

	err := row.Scan(&card.Name, &card.ID, &rarity.ID, &rarity.Name, &faction.ID, &faction.Name, &cardType.ID, &cardType.Name, &strength, &text, &flavor)

	if err != nil {
		card.Faction.ID = ""
		card.Faction.Name = ""
		card.Flavor = nil
		card.ID = ""
		card.Name = ""
		card.Rarity.ID = ""
		card.Rarity.Name = ""
		card.Rows = nil
		card.Strength = nil
		card.Subtypes = nil
		card.Text = nil
		card.TypeCard.ID = ""
		card.TypeCard.Name = ""
	} else {
		card.Faction = faction
		card.Rarity = rarity
		card.TypeCard = cardType

		cardRows, _ := fetchCardRows(id)
		cardSubtypes, _ := fetchCardSubTypes(id)

		card.Rows = cardRows
		card.Subtypes = cardSubtypes

		if strength.Valid {
			var converted = int(strength.Int64)
			card.Strength = &converted
		} else {
			card.Strength = nil
		}

		if text.Valid {
			card.Text = &text.String
		} else {
			card.Text = nil
		}

		if flavor.Valid {
			card.Flavor = &flavor.String
		} else {
			card.Flavor = nil
		}
	}

	return card, err
}

func fetchCardRows(id string) ([]*RowModel, error) {
	var cardRows []*RowModel

	rows, err := DBCon.Query("SELECT r.id, r.name FROM Rows AS r INNER JOIN CardsRows AS cr ON r.idRow = cr.idRow INNER JOIN Cards AS c ON c.idCard = cr.idCard WHERE c.id =?", id)
	defer rows.Close()

	if err != nil {
		return cardRows, err
	}

	for rows.Next() {
		var rowModel RowModel
		err := rows.Scan(&rowModel.ID, &rowModel.Name)

		if err != nil {
			continue
		}
		cardRows = append(cardRows, &rowModel)
	}

	return cardRows, err
}

func fetchCardSubTypes(id string) ([]*TypeModel, error) {
	var cardSubtypes []*TypeModel

	rows, err := DBCon.Query("SELECT t.id, t.name FROM Types AS t INNER JOIN CardsSubtypes AS st ON t.idType = st.idType INNER JOIN Cards AS c ON c.idCard = st.idCard WHERE c.id =?", id)
	defer rows.Close()

	if err != nil {
		return cardSubtypes, err
	}

	for rows.Next() {
		var typeModel TypeModel
		err := rows.Scan(&typeModel.ID, &typeModel.Name)

		if err != nil {
			continue
		}
		cardSubtypes = append(cardSubtypes, &typeModel)
	}

	return cardSubtypes, err
}