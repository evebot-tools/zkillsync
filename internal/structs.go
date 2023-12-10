package internal

import "time"

type KillmailResponse struct {
	Package *struct {
		KillID   int `json:"killID"`
		Killmail struct {
			Attackers []struct {
				AllianceID     *int    `json:"alliance_id,omitempty"`
				CharacterID    *int    `json:"character_id,omitempty"`
				CorporationID  *int    `json:"corporation_id,omitempty"`
				DamageDone     int     `json:"damage_done"`
				FactionID      *int    `json:"faction_id,omitempty"`
				FinalBlow      bool    `json:"final_blow"`
				SecurityStatus float64 `json:"security_status"`
				ShipTypeID     *int    `json:"ship_type_id,omitempty"`
				WeaponTypeID   *int    `json:"weapon_type_id,omitempty"`
			} `json:"attackers"`
			KillmailID    int       `json:"killmail_id"`
			KillmailTime  time.Time `json:"killmail_time"`
			SolarSystemID int       `json:"solar_system_id"`
			Victim        struct {
				AllianceID    *int `json:"alliance_id,omitempty"`
				CharacterID   int  `json:"character_id"`
				CorporationID int  `json:"corporation_id"`
				DamageTaken   int  `json:"damage_taken"`
				FactionID     *int `json:"faction_id,omitempty"`
				Items         []struct {
					Flag       int `json:"flag"`
					ItemTypeID int `json:"item_type_id"`
					Items      []struct {
						Flag            int `json:"flag"`
						ItemTypeID      int `json:"item_type_id"`
						QuantityDropped int `json:"quantity_dropped"`
						Singleton       int `json:"singleton"`
					} `json:"items,omitempty"`
					QuantityDestroyed *int `json:"quantity_destroyed,omitempty"`
					QuantityDropped   *int `json:"quantity_dropped,omitempty"`
					Singleton         int  `json:"singleton"`
				} `json:"items"`
				Position struct {
					X float64 `json:"x"`
					Y float64 `json:"y"`
					Z float64 `json:"z"`
				} `json:"position"`
				ShipTypeID int `json:"ship_type_id"`
			} `json:"victim"`
			WarID *int `json:"war_id,omitempty"`
		} `json:"killmail"`
		Zkb struct {
			Awox           bool     `json:"awox"`
			DestroyedValue float64  `json:"destroyedValue"`
			DroppedValue   float64  `json:"droppedValue"`
			FittedValue    float64  `json:"fittedValue"`
			Hash           string   `json:"hash"`
			Href           string   `json:"href"`
			Labels         []string `json:"labels"`
			LocationID     int      `json:"locationID"`
			Npc            bool     `json:"npc"`
			Points         int      `json:"points"`
			Solo           bool     `json:"solo"`
			TotalValue     float64  `json:"totalValue"`
		} `json:"zkb"`
	} `json:"package"`
}
