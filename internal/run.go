package internal

import (
	"github.com/evebot-tools/database"
	"github.com/evebot-tools/utils"
	"github.com/imroc/req/v3"
	"github.com/kamva/mgm/v3"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
)

const (
	zkillboardURL = "https://redisq.zkillboard.com/listen.php?queueID=evebot-tools-prod"
)

func init() {
	log.Info().Msg("Initializing")
	utils.InitMongoDBClient()
}

func Run() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	client := req.C()
	for {
		select {
		case <-quit:
			log.Fatal().Msg("SIGINT received, exiting")
		default:
			resp, err := client.R().Get(zkillboardURL)
			if err != nil {
				log.Err(err).Msg("failed to get killmail api response")
			}
			var data KillmailResponse
			err = resp.UnmarshalJson(&data)
			if err != nil {
				log.Err(err).Msg("failed to unmarshal json")
			} else {
				if data == (KillmailResponse{}) {
					log.Info().Msg("Timeout reached. No killmails to process.")
				} else {
					go ProcessKillmail(&data)
				}
			}
		}
	}
}

func ProcessKillmail(resp *KillmailResponse) {
	r := resp.Package
	var attackers []database.KillmailAttackers
	var items []database.KillmailItems
	for _, a := range r.Killmail.Attackers {
		attackers = append(attackers, database.KillmailAttackers{
			AllianceID:     a.AllianceID,
			CharacterID:    a.CharacterID,
			CorporationID:  a.CorporationID,
			DamageDone:     a.DamageDone,
			FactionID:      a.FactionID,
			FinalBlow:      a.FinalBlow,
			SecurityStatus: a.SecurityStatus,
			ShipTypeID:     a.ShipTypeID,
			WeaponTypeID:   a.WeaponTypeID,
		})
	}
	for _, i := range r.Killmail.Victim.Items {
		items = append(items, database.KillmailItems{
			Flag:       i.Flag,
			ItemTypeID: i.ItemTypeID,
			Items: []struct {
				Flag            int `json:"flag" bson:"flag"`
				ItemTypeID      int `json:"itemTypeID" bson:"itemTypeID"`
				QuantityDropped int `json:"quantityDropped" bson:"quantityDropped"`
				Singleton       int `json:"singleton" bson:"singleton"`
			}(i.Items),
			QuantityDestroyed: i.QuantityDestroyed,
			QuantityDropped:   i.QuantityDropped,
			Singleton:         i.Singleton,
		})
	}
	record := database.Killmail{
		KillmailID:    r.Killmail.KillmailID,
		KillmailTime:  r.Killmail.KillmailTime,
		SolarSystemID: r.Killmail.SolarSystemID,
		WarID:         r.Killmail.WarID,
		DamageTaken:   r.Killmail.Victim.DamageTaken,
		TotalValue:    r.Zkb.TotalValue,
		Victim: database.KillmailVictim{
			AllianceID:    r.Killmail.Victim.AllianceID,
			CharacterID:   r.Killmail.Victim.CharacterID,
			CorporationID: r.Killmail.Victim.CorporationID,
			FactionID:     r.Killmail.Victim.FactionID,
		},
		Items: items,
		Position: database.KillmailPosition{
			X: r.Killmail.Victim.Position.X,
			Y: r.Killmail.Victim.Position.Y,
			Z: r.Killmail.Victim.Position.Z,
		},
		Attackers: attackers,
		ZkillboardData: database.KillmailZkb{
			Awox:           r.Zkb.Awox,
			DestroyedValue: r.Zkb.DestroyedValue,
			DroppedValue:   r.Zkb.DroppedValue,
			FittedValue:    r.Zkb.FittedValue,
			Hash:           r.Zkb.Hash,
			Href:           r.Zkb.Href,
			Labels:         r.Zkb.Labels,
			LocationID:     r.Zkb.LocationID,
			Npc:            r.Zkb.Npc,
			Points:         r.Zkb.Points,
			Solo:           r.Zkb.Solo,
		},
	}
	err := mgm.Coll(&database.Killmail{}).Create(&record)
	if err != nil {
		log.Err(err).Msg("failed to create record")
	}
}
