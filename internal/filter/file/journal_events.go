package file

import (
	"edpad2/internal/local/display"
	"edpad2/internal/router"
	"encoding/json"
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
)

type ApproachBodyEvent struct {
	Body          string    `json:"Body,omitempty"`
	BodyID        int       `json:"BodyID,omitempty"`
	StarSystem    string    `json:"StarSystem,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evApproachBody(ev *ApproachBodyEvent) {
	return
}

type ApproachSettlementEvent struct {
	BodyID        int       `json:"BodyID,omitempty"`
	BodyName      string    `json:"BodyName,omitempty"`
	Latitude      float64   `json:"Latitude,omitempty"`
	Longitude     float64   `json:"Longitude,omitempty"`
	MarketID      int       `json:"MarketID,omitempty"`
	Name          string    `json:"Name,omitempty"`
	NameLocalised string    `json:"Name_Localised,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evApproachSettlement(ev *ApproachSettlementEvent) {
	return
}

type BackpackEvent struct {
	Components  []interface{} `json:"Components,omitempty"`
	Consumables []struct {
		Count         int    `json:"Count,omitempty"`
		Name          string `json:"Name,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
		OwnerID       int    `json:"OwnerID,omitempty"`
	} `json:"Consumables,omitempty"`
	Data      []interface{} `json:"Data,omitempty"`
	Items     []interface{} `json:"Items,omitempty"`
	Event     string        `json:"event,omitempty"`
	Timestamp time.Time     `json:"timestamp,omitempty"`
}

func (h *handler) evBackpack(ev *BackpackEvent) {
	return
}

type BackPackEvent interface{}

func (h *handler) evBackPack(ev *BackPackEvent) {
	return
}

type BountyEvent struct {
	Reward  int `json:"Reward,omitempty"`
	Rewards []struct {
		Faction string `json:"Faction,omitempty"`
		Reward  int    `json:"Reward,omitempty"`
	} `json:"Rewards,omitempty"`
	Target                 string    `json:"Target,omitempty"`
	TargetLocalised        string    `json:"Target_Localised,omitempty"`
	TotalReward            int       `json:"TotalReward,omitempty"`
	VictimFaction          string    `json:"VictimFaction,omitempty"`
	VictimFactionLocalised string    `json:"VictimFaction_Localised,omitempty"`
	Event                  string    `json:"event,omitempty"`
	Timestamp              time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evBounty(ev *BountyEvent) {
	return
}

type BuyAmmoEvent struct {
	Cost      int       `json:"Cost,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evBuyAmmo(ev *BuyAmmoEvent) {
	return
}

type BuyDronesEvent struct {
	BuyPrice  int       `json:"BuyPrice,omitempty"`
	Count     int       `json:"Count,omitempty"`
	TotalCost int       `json:"TotalCost,omitempty"`
	Type      string    `json:"Type,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evBuyDrones(ev *BuyDronesEvent) {
	return
}

type BuyExplorationDataEvent struct {
	Cost      int       `json:"Cost,omitempty"`
	System    string    `json:"System,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evBuyExplorationData(ev *BuyExplorationDataEvent) {
	return
}

type BuyMicroResourcesEvent struct {
	Category      string    `json:"Category,omitempty"`
	Count         int       `json:"Count,omitempty"`
	MarketID      int       `json:"MarketID,omitempty"`
	Name          string    `json:"Name,omitempty"`
	NameLocalised string    `json:"Name_Localised,omitempty"`
	Price         int       `json:"Price,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evBuyMicroResources(ev *BuyMicroResourcesEvent) {
	return
}

type BuySuitEvent struct {
	Name          string    `json:"Name,omitempty"`
	NameLocalised string    `json:"Name_Localised,omitempty"`
	Price         int       `json:"Price,omitempty"`
	SuitID        int       `json:"SuitID,omitempty"`
	SuitMods      []string  `json:"SuitMods,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evBuySuit(ev *BuySuitEvent) {
	return
}

type BuyTradeDataEvent struct {
	Cost      int       `json:"Cost,omitempty"`
	System    string    `json:"System,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evBuyTradeData(ev *BuyTradeDataEvent) {
	return
}

type BuyWeaponEvent struct {
	Class         int           `json:"Class,omitempty"`
	Name          string        `json:"Name,omitempty"`
	NameLocalised string        `json:"Name_Localised,omitempty"`
	Price         int           `json:"Price,omitempty"`
	SuitModuleID  int           `json:"SuitModuleID,omitempty"`
	WeaponMods    []interface{} `json:"WeaponMods,omitempty"`
	Event         string        `json:"event,omitempty"`
	Timestamp     time.Time     `json:"timestamp,omitempty"`
}

func (h *handler) evBuyWeapon(ev *BuyWeaponEvent) {
	return
}

type CargoEvent struct {
	Count     int `json:"Count,omitempty"`
	Inventory []struct {
		Count         int    `json:"Count,omitempty"`
		MissionID     int    `json:"MissionID,omitempty"`
		Name          string `json:"Name,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
		Stolen        int    `json:"Stolen,omitempty"`
	} `json:"Inventory,omitempty"`
	Vessel    string    `json:"Vessel,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCargo(ev *CargoEvent) {
	return
}

type CargoDepotEvent struct {
	CargoType           string    `json:"CargoType,omitempty"`
	CargoTypeLocalised  string    `json:"CargoType_Localised,omitempty"`
	Count               int       `json:"Count,omitempty"`
	EndMarketID         int       `json:"EndMarketID,omitempty"`
	ItemsCollected      int       `json:"ItemsCollected,omitempty"`
	ItemsDelivered      int       `json:"ItemsDelivered,omitempty"`
	MissionID           int       `json:"MissionID,omitempty"`
	Progress            float64   `json:"Progress,omitempty"`
	StartMarketID       int       `json:"StartMarketID,omitempty"`
	TotalItemsToDeliver int       `json:"TotalItemsToDeliver,omitempty"`
	UpdateType          string    `json:"UpdateType,omitempty"`
	Event               string    `json:"event,omitempty"`
	Timestamp           time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCargoDepot(ev *CargoDepotEvent) {
	return
}

type CarrierJumpEvent struct {
	Body      string `json:"Body,omitempty"`
	BodyID    int    `json:"BodyID,omitempty"`
	BodyType  string `json:"BodyType,omitempty"`
	Conflicts []struct {
		Faction1 struct {
			Name    string `json:"Name,omitempty"`
			Stake   string `json:"Stake,omitempty"`
			WonDays int    `json:"WonDays,omitempty"`
		} `json:"Faction1,omitempty"`
		Faction2 struct {
			Name    string `json:"Name,omitempty"`
			Stake   string `json:"Stake,omitempty"`
			WonDays int    `json:"WonDays,omitempty"`
		} `json:"Faction2,omitempty"`
		Status  string `json:"Status,omitempty"`
		WarType string `json:"WarType,omitempty"`
	} `json:"Conflicts,omitempty"`
	Docked   bool `json:"Docked,omitempty"`
	Factions []struct {
		ActiveStates []struct {
			State string `json:"State,omitempty"`
		} `json:"ActiveStates,omitempty"`
		Allegiance         string  `json:"Allegiance,omitempty"`
		FactionState       string  `json:"FactionState,omitempty"`
		Government         string  `json:"Government,omitempty"`
		Happiness          string  `json:"Happiness,omitempty"`
		HappinessLocalised string  `json:"Happiness_Localised,omitempty"`
		Influence          float64 `json:"Influence,omitempty"`
		MyReputation       float64 `json:"MyReputation,omitempty"`
		Name               string  `json:"Name,omitempty"`
	} `json:"Factions,omitempty"`
	MarketID         int       `json:"MarketID,omitempty"`
	Population       int       `json:"Population,omitempty"`
	StarPos          []float64 `json:"StarPos,omitempty"`
	StarSystem       string    `json:"StarSystem,omitempty"`
	StationEconomies []struct {
		Name          string  `json:"Name,omitempty"`
		NameLocalised string  `json:"Name_Localised,omitempty"`
		Proportion    float64 `json:"Proportion,omitempty"`
	} `json:"StationEconomies,omitempty"`
	StationEconomy          string `json:"StationEconomy,omitempty"`
	StationEconomyLocalised string `json:"StationEconomy_Localised,omitempty"`
	StationFaction          struct {
		Name string `json:"Name,omitempty"`
	} `json:"StationFaction,omitempty"`
	StationGovernment          string   `json:"StationGovernment,omitempty"`
	StationGovernmentLocalised string   `json:"StationGovernment_Localised,omitempty"`
	StationName                string   `json:"StationName,omitempty"`
	StationServices            []string `json:"StationServices,omitempty"`
	StationType                string   `json:"StationType,omitempty"`
	SystemAddress              int      `json:"SystemAddress,omitempty"`
	SystemAllegiance           string   `json:"SystemAllegiance,omitempty"`
	SystemEconomy              string   `json:"SystemEconomy,omitempty"`
	SystemEconomyLocalised     string   `json:"SystemEconomy_Localised,omitempty"`
	SystemFaction              struct {
		FactionState string `json:"FactionState,omitempty"`
		Name         string `json:"Name,omitempty"`
	} `json:"SystemFaction,omitempty"`
	SystemGovernment             string    `json:"SystemGovernment,omitempty"`
	SystemGovernmentLocalised    string    `json:"SystemGovernment_Localised,omitempty"`
	SystemSecondEconomy          string    `json:"SystemSecondEconomy,omitempty"`
	SystemSecondEconomyLocalised string    `json:"SystemSecondEconomy_Localised,omitempty"`
	SystemSecurity               string    `json:"SystemSecurity,omitempty"`
	SystemSecurityLocalised      string    `json:"SystemSecurity_Localised,omitempty"`
	Event                        string    `json:"event,omitempty"`
	Timestamp                    time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCarrierJump(ev *CarrierJumpEvent) {
	return
}

type CodexEntryEvent struct {
	Category                    string    `json:"Category,omitempty"`
	CategoryLocalised           string    `json:"Category_Localised,omitempty"`
	EntryID                     int       `json:"EntryID,omitempty"`
	IsNewEntry                  bool      `json:"IsNewEntry,omitempty"`
	Latitude                    float64   `json:"Latitude,omitempty"`
	Longitude                   float64   `json:"Longitude,omitempty"`
	Name                        string    `json:"Name,omitempty"`
	NameLocalised               string    `json:"Name_Localised,omitempty"`
	NearestDestination          string    `json:"NearestDestination,omitempty"`
	NearestDestinationLocalised string    `json:"NearestDestination_Localised,omitempty"`
	Region                      string    `json:"Region,omitempty"`
	RegionLocalised             string    `json:"Region_Localised,omitempty"`
	SubCategory                 string    `json:"SubCategory,omitempty"`
	SubCategoryLocalised        string    `json:"SubCategory_Localised,omitempty"`
	System                      string    `json:"System,omitempty"`
	SystemAddress               int       `json:"SystemAddress,omitempty"`
	VoucherAmount               int       `json:"VoucherAmount,omitempty"`
	Event                       string    `json:"event,omitempty"`
	Timestamp                   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCodexEntry(ev *CodexEntryEvent) {
	return
}

type CollectCargoEvent struct {
	Stolen        bool      `json:"Stolen,omitempty"`
	Type          string    `json:"Type,omitempty"`
	TypeLocalised string    `json:"Type_Localised,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCollectCargo(ev *CollectCargoEvent) {
	return
}

type CommanderEvent struct {
	Fid       string    `json:"FID,omitempty"`
	Name      string    `json:"Name,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCommander(ev *CommanderEvent) {
	return
}

type CommitCrimeEvent struct {
	Bounty    int       `json:"Bounty,omitempty"`
	CrimeType string    `json:"CrimeType,omitempty"`
	Faction   string    `json:"Faction,omitempty"`
	Fine      int       `json:"Fine,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCommitCrime(ev *CommitCrimeEvent) {
	return
}

type CommunityGoalEvent struct {
	CurrentGoals []struct {
		Bonus                int       `json:"Bonus,omitempty"`
		Cgid                 int       `json:"CGID,omitempty"`
		CurrentTotal         int       `json:"CurrentTotal,omitempty"`
		Expiry               time.Time `json:"Expiry,omitempty"`
		IsComplete           bool      `json:"IsComplete,omitempty"`
		MarketName           string    `json:"MarketName,omitempty"`
		NumContributors      int       `json:"NumContributors,omitempty"`
		PlayerContribution   int       `json:"PlayerContribution,omitempty"`
		PlayerInTopRank      bool      `json:"PlayerInTopRank,omitempty"`
		PlayerPercentileBand int       `json:"PlayerPercentileBand,omitempty"`
		SystemName           string    `json:"SystemName,omitempty"`
		TierReached          string    `json:"TierReached,omitempty"`
		Title                string    `json:"Title,omitempty"`
		TopRankSize          int       `json:"TopRankSize,omitempty"`
		TopTier              struct {
			Bonus string `json:"Bonus,omitempty"`
			Name  string `json:"Name,omitempty"`
		} `json:"TopTier,omitempty"`
	} `json:"CurrentGoals,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCommunityGoal(ev *CommunityGoalEvent) {
	return
}

type CommunityGoalDiscardEvent struct {
	Cgid      int       `json:"CGID,omitempty"`
	Name      string    `json:"Name,omitempty"`
	System    string    `json:"System,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCommunityGoalDiscard(ev *CommunityGoalDiscardEvent) {
	return
}

type CommunityGoalJoinEvent struct {
	Cgid      int       `json:"CGID,omitempty"`
	Name      string    `json:"Name,omitempty"`
	System    string    `json:"System,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCommunityGoalJoin(ev *CommunityGoalJoinEvent) {
	return
}

type CommunityGoalRewardEvent struct {
	Cgid      int       `json:"CGID,omitempty"`
	Name      string    `json:"Name,omitempty"`
	Reward    int       `json:"Reward,omitempty"`
	System    string    `json:"System,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCommunityGoalReward(ev *CommunityGoalRewardEvent) {
	return
}

type CreateSuitLoadoutEvent struct {
	LoadoutID   int    `json:"LoadoutID,omitempty"`
	LoadoutName string `json:"LoadoutName,omitempty"`
	Modules     []struct {
		Class               int           `json:"Class,omitempty"`
		ModuleName          string        `json:"ModuleName,omitempty"`
		ModuleNameLocalised string        `json:"ModuleName_Localised,omitempty"`
		SlotName            string        `json:"SlotName,omitempty"`
		SuitModuleID        int           `json:"SuitModuleID,omitempty"`
		WeaponMods          []interface{} `json:"WeaponMods,omitempty"`
	} `json:"Modules,omitempty"`
	SuitID            int           `json:"SuitID,omitempty"`
	SuitMods          []interface{} `json:"SuitMods,omitempty"`
	SuitName          string        `json:"SuitName,omitempty"`
	SuitNameLocalised string        `json:"SuitName_Localised,omitempty"`
	Event             string        `json:"event,omitempty"`
	Timestamp         time.Time     `json:"timestamp,omitempty"`
}

func (h *handler) evCreateSuitLoadout(ev *CreateSuitLoadoutEvent) {
	return
}

type CrewAssignEvent struct {
	CrewID    int       `json:"CrewID,omitempty"`
	Name      string    `json:"Name,omitempty"`
	Role      string    `json:"Role,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCrewAssign(ev *CrewAssignEvent) {
	return
}

type CrewFireEvent struct {
	CrewID    int       `json:"CrewID,omitempty"`
	Name      string    `json:"Name,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCrewFire(ev *CrewFireEvent) {
	return
}

type CrewHireEvent struct {
	CombatRank int       `json:"CombatRank,omitempty"`
	Cost       int       `json:"Cost,omitempty"`
	CrewID     int       `json:"CrewID,omitempty"`
	Faction    string    `json:"Faction,omitempty"`
	Name       string    `json:"Name,omitempty"`
	Event      string    `json:"event,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCrewHire(ev *CrewHireEvent) {
	return
}

type CrimeVictimEvent struct {
	Bounty    int       `json:"Bounty,omitempty"`
	CrimeType string    `json:"CrimeType,omitempty"`
	Fine      int       `json:"Fine,omitempty"`
	Offender  string    `json:"Offender,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evCrimeVictim(ev *CrimeVictimEvent) {
	return
}

type DatalinkScanEvent struct {
	Message          string    `json:"Message,omitempty"`
	MessageLocalised string    `json:"Message_Localised,omitempty"`
	Event            string    `json:"event,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDatalinkScan(ev *DatalinkScanEvent) {
	return
}

type DatalinkVoucherEvent struct {
	PayeeFaction  string    `json:"PayeeFaction,omitempty"`
	Reward        int       `json:"Reward,omitempty"`
	VictimFaction string    `json:"VictimFaction,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDatalinkVoucher(ev *DatalinkVoucherEvent) {
	return
}

type DataScannedEvent struct {
	Type          string    `json:"Type,omitempty"`
	TypeLocalised string    `json:"Type_Localised,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDataScanned(ev *DataScannedEvent) {
	return
}

type DeleteSuitLoadoutEvent struct {
	LoadoutID         int       `json:"LoadoutID,omitempty"`
	LoadoutName       string    `json:"LoadoutName,omitempty"`
	SuitID            int       `json:"SuitID,omitempty"`
	SuitName          string    `json:"SuitName,omitempty"`
	SuitNameLocalised string    `json:"SuitName_Localised,omitempty"`
	Event             string    `json:"event,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDeleteSuitLoadout(ev *DeleteSuitLoadoutEvent) {
	return
}

type DiedEvent struct {
	KillerName string    `json:"KillerName,omitempty"`
	KillerRank string    `json:"KillerRank,omitempty"`
	KillerShip string    `json:"KillerShip,omitempty"`
	Event      string    `json:"event,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDied(ev *DiedEvent) {
	return
}

type DisembarkEvent struct {
	Body          string    `json:"Body,omitempty"`
	BodyID        int       `json:"BodyID,omitempty"`
	ID            int       `json:"ID,omitempty"`
	MarketID      int       `json:"MarketID,omitempty"`
	Multicrew     bool      `json:"Multicrew,omitempty"`
	OnPlanet      bool      `json:"OnPlanet,omitempty"`
	OnStation     bool      `json:"OnStation,omitempty"`
	Srv           bool      `json:"SRV,omitempty"`
	StarSystem    string    `json:"StarSystem,omitempty"`
	StationName   string    `json:"StationName,omitempty"`
	StationType   string    `json:"StationType,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	Taxi          bool      `json:"Taxi,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDisembark(ev *DisembarkEvent) {
	return
}

type DockedEvent struct {
	ActiveFine        bool    `json:"ActiveFine,omitempty"`
	DistFromStarLs    float64 `json:"DistFromStarLS,omitempty"`
	MarketID          int     `json:"MarketID,omitempty"`
	Multicrew         bool    `json:"Multicrew,omitempty"`
	StarSystem        string  `json:"StarSystem,omitempty"`
	StationAllegiance string  `json:"StationAllegiance,omitempty"`
	StationEconomies  []struct {
		Name          string  `json:"Name,omitempty"`
		NameLocalised string  `json:"Name_Localised,omitempty"`
		Proportion    float64 `json:"Proportion,omitempty"`
	} `json:"StationEconomies,omitempty"`
	StationEconomy          string `json:"StationEconomy,omitempty"`
	StationEconomyLocalised string `json:"StationEconomy_Localised,omitempty"`
	StationFaction          struct {
		FactionState string `json:"FactionState,omitempty"`
		Name         string `json:"Name,omitempty"`
	} `json:"StationFaction,omitempty"`
	StationGovernment          string    `json:"StationGovernment,omitempty"`
	StationGovernmentLocalised string    `json:"StationGovernment_Localised,omitempty"`
	StationName                string    `json:"StationName,omitempty"`
	StationServices            []string  `json:"StationServices,omitempty"`
	StationType                string    `json:"StationType,omitempty"`
	SystemAddress              int       `json:"SystemAddress,omitempty"`
	Taxi                       bool      `json:"Taxi,omitempty"`
	Wanted                     bool      `json:"Wanted,omitempty"`
	Event                      string    `json:"event,omitempty"`
	Timestamp                  time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDocked(ev *DockedEvent) {
	return
}

type DockFighterEvent struct {
	ID        int       `json:"ID,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDockFighter(ev *DockFighterEvent) {
	return
}

type DockingCancelledEvent struct {
	MarketID    int       `json:"MarketID,omitempty"`
	StationName string    `json:"StationName,omitempty"`
	StationType string    `json:"StationType,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDockingCancelled(ev *DockingCancelledEvent) {
	return
}

type DockingDeniedEvent struct {
	MarketID    int       `json:"MarketID,omitempty"`
	Reason      string    `json:"Reason,omitempty"`
	StationName string    `json:"StationName,omitempty"`
	StationType string    `json:"StationType,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDockingDenied(ev *DockingDeniedEvent) {
	return
}

type DockingGrantedEvent struct {
	LandingPad  int       `json:"LandingPad,omitempty"`
	MarketID    int       `json:"MarketID,omitempty"`
	StationName string    `json:"StationName,omitempty"`
	StationType string    `json:"StationType,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDockingGranted(ev *DockingGrantedEvent) {
	return
}

type DockingRequestedEvent struct {
	MarketID    int       `json:"MarketID,omitempty"`
	StationName string    `json:"StationName,omitempty"`
	StationType string    `json:"StationType,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDockingRequested(ev *DockingRequestedEvent) {
	return
}

type DockingTimeoutEvent interface{}

func (h *handler) evDockingTimeout(ev *DockingTimeoutEvent) {
	return
}

type DockSRVEvent struct {
	ID        int       `json:"ID,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evDockSRV(ev *DockSRVEvent) {
	return
}

type EjectCargoEvent struct {
	Abandoned     bool      `json:"Abandoned,omitempty"`
	Count         int       `json:"Count,omitempty"`
	Type          string    `json:"Type,omitempty"`
	TypeLocalised string    `json:"Type_Localised,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evEjectCargo(ev *EjectCargoEvent) {
	return
}

type EmbarkEvent struct {
	Body          string    `json:"Body,omitempty"`
	BodyID        int       `json:"BodyID,omitempty"`
	ID            int       `json:"ID,omitempty"`
	MarketID      int       `json:"MarketID,omitempty"`
	Multicrew     bool      `json:"Multicrew,omitempty"`
	OnPlanet      bool      `json:"OnPlanet,omitempty"`
	OnStation     bool      `json:"OnStation,omitempty"`
	Srv           bool      `json:"SRV,omitempty"`
	StarSystem    string    `json:"StarSystem,omitempty"`
	StationName   string    `json:"StationName,omitempty"`
	StationType   string    `json:"StationType,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	Taxi          bool      `json:"Taxi,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evEmbark(ev *EmbarkEvent) {
	return
}

type EngineerContributionEvent struct {
	Commodity          string    `json:"Commodity,omitempty"`
	CommodityLocalised string    `json:"Commodity_Localised,omitempty"`
	Engineer           string    `json:"Engineer,omitempty"`
	EngineerID         int       `json:"EngineerID,omitempty"`
	Quantity           int       `json:"Quantity,omitempty"`
	TotalQuantity      int       `json:"TotalQuantity,omitempty"`
	Type               string    `json:"Type,omitempty"`
	Event              string    `json:"event,omitempty"`
	Timestamp          time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evEngineerContribution(ev *EngineerContributionEvent) {
	return
}

type EngineerCraftEvent struct {
	ApplyExperimentalEffect     string `json:"ApplyExperimentalEffect,omitempty"`
	BlueprintID                 int    `json:"BlueprintID,omitempty"`
	BlueprintName               string `json:"BlueprintName,omitempty"`
	Engineer                    string `json:"Engineer,omitempty"`
	EngineerID                  int    `json:"EngineerID,omitempty"`
	ExperimentalEffect          string `json:"ExperimentalEffect,omitempty"`
	ExperimentalEffectLocalised string `json:"ExperimentalEffect_Localised,omitempty"`
	Ingredients                 []struct {
		Count         int    `json:"Count,omitempty"`
		Name          string `json:"Name,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
	} `json:"Ingredients,omitempty"`
	Level     int `json:"Level,omitempty"`
	Modifiers []struct {
		Label         string  `json:"Label,omitempty"`
		LessIsGood    int     `json:"LessIsGood,omitempty"`
		OriginalValue float64 `json:"OriginalValue,omitempty"`
		Value         float64 `json:"Value,omitempty"`
	} `json:"Modifiers,omitempty"`
	Module    string    `json:"Module,omitempty"`
	Quality   float64   `json:"Quality,omitempty"`
	Slot      string    `json:"Slot,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evEngineerCraft(ev *EngineerCraftEvent) {
	return
}

type EngineerProgressEvent struct {
	Engineer   string `json:"Engineer,omitempty"`
	EngineerID int    `json:"EngineerID,omitempty"`
	Engineers  []struct {
		Engineer     string `json:"Engineer,omitempty"`
		EngineerID   int    `json:"EngineerID,omitempty"`
		Progress     string `json:"Progress,omitempty"`
		Rank         int    `json:"Rank,omitempty"`
		RankProgress int    `json:"RankProgress,omitempty"`
	} `json:"Engineers,omitempty"`
	Progress  string    `json:"Progress,omitempty"`
	Rank      int       `json:"Rank,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evEngineerProgress(ev *EngineerProgressEvent) {
	return
}

type EscapeInterdictionEvent struct {
	Interdictor string    `json:"Interdictor,omitempty"`
	IsPlayer    bool      `json:"IsPlayer,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evEscapeInterdiction(ev *EscapeInterdictionEvent) {
	return
}

type FetchRemoteModuleEvent struct {
	ServerID            int       `json:"ServerId,omitempty"`
	Ship                string    `json:"Ship,omitempty"`
	ShipID              int       `json:"ShipID,omitempty"`
	StorageSlot         int       `json:"StorageSlot,omitempty"`
	StoredItem          string    `json:"StoredItem,omitempty"`
	StoredItemLocalised string    `json:"StoredItem_Localised,omitempty"`
	TransferCost        int       `json:"TransferCost,omitempty"`
	TransferTime        int       `json:"TransferTime,omitempty"`
	Event               string    `json:"event,omitempty"`
	Timestamp           time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evFetchRemoteModule(ev *FetchRemoteModuleEvent) {
	return
}

type FileheaderEvent struct {
	Odyssey     bool      `json:"Odyssey,omitempty"`
	Build       string    `json:"build,omitempty"`
	Event       string    `json:"event,omitempty"`
	Gameversion string    `json:"gameversion,omitempty"`
	Language    string    `json:"language,omitempty"`
	Part        int       `json:"part,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evFileheader(ev *FileheaderEvent) {
	return
}

type FriendsEvent struct {
	Name      string    `json:"Name,omitempty"`
	Status    string    `json:"Status,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evFriends(ev *FriendsEvent) {
	return
}

type FSDJumpEvent struct {
	Body      string `json:"Body,omitempty"`
	BodyID    int    `json:"BodyID,omitempty"`
	BodyType  string `json:"BodyType,omitempty"`
	BoostUsed int    `json:"BoostUsed,omitempty"`
	Conflicts []struct {
		Faction1 struct {
			Name    string `json:"Name,omitempty"`
			Stake   string `json:"Stake,omitempty"`
			WonDays int    `json:"WonDays,omitempty"`
		} `json:"Faction1,omitempty"`
		Faction2 struct {
			Name    string `json:"Name,omitempty"`
			Stake   string `json:"Stake,omitempty"`
			WonDays int    `json:"WonDays,omitempty"`
		} `json:"Faction2,omitempty"`
		Status  string `json:"Status,omitempty"`
		WarType string `json:"WarType,omitempty"`
	} `json:"Conflicts,omitempty"`
	Factions []struct {
		ActiveStates []struct {
			State string `json:"State,omitempty"`
		} `json:"ActiveStates,omitempty"`
		Allegiance         string  `json:"Allegiance,omitempty"`
		FactionState       string  `json:"FactionState,omitempty"`
		Government         string  `json:"Government,omitempty"`
		HappiestSystem     bool    `json:"HappiestSystem,omitempty"`
		Happiness          string  `json:"Happiness,omitempty"`
		HappinessLocalised string  `json:"Happiness_Localised,omitempty"`
		HomeSystem         bool    `json:"HomeSystem,omitempty"`
		Influence          float64 `json:"Influence,omitempty"`
		MyReputation       float64 `json:"MyReputation,omitempty"`
		Name               string  `json:"Name,omitempty"`
		PendingStates      []struct {
			State string `json:"State,omitempty"`
			Trend int    `json:"Trend,omitempty"`
		} `json:"PendingStates,omitempty"`
		RecoveringStates []struct {
			State string `json:"State,omitempty"`
			Trend int    `json:"Trend,omitempty"`
		} `json:"RecoveringStates,omitempty"`
		SquadronFaction bool `json:"SquadronFaction,omitempty"`
	} `json:"Factions,omitempty"`
	FuelLevel              float64   `json:"FuelLevel,omitempty"`
	FuelUsed               float64   `json:"FuelUsed,omitempty"`
	JumpDist               float64   `json:"JumpDist,omitempty"`
	Multicrew              bool      `json:"Multicrew,omitempty"`
	Population             int       `json:"Population,omitempty"`
	PowerplayState         string    `json:"PowerplayState,omitempty"`
	Powers                 []string  `json:"Powers,omitempty"`
	StarPos                []float64 `json:"StarPos,omitempty"`
	StarSystem             string    `json:"StarSystem,omitempty"`
	SystemAddress          int       `json:"SystemAddress,omitempty"`
	SystemAllegiance       string    `json:"SystemAllegiance,omitempty"`
	SystemEconomy          string    `json:"SystemEconomy,omitempty"`
	SystemEconomyLocalised string    `json:"SystemEconomy_Localised,omitempty"`
	SystemFaction          *struct {
		FactionState string `json:"FactionState,omitempty"`
		Name         string `json:"Name,omitempty"`
	} `json:"SystemFaction,omitempty"`
	SystemGovernment             string    `json:"SystemGovernment,omitempty"`
	SystemGovernmentLocalised    string    `json:"SystemGovernment_Localised,omitempty"`
	SystemSecondEconomy          string    `json:"SystemSecondEconomy,omitempty"`
	SystemSecondEconomyLocalised string    `json:"SystemSecondEconomy_Localised,omitempty"`
	SystemSecurity               string    `json:"SystemSecurity,omitempty"`
	SystemSecurityLocalised      string    `json:"SystemSecurity_Localised,omitempty"`
	Taxi                         bool      `json:"Taxi,omitempty"`
	Event                        string    `json:"event,omitempty"`
	Timestamp                    time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evFSDJump(ev *FSDJumpEvent) {
	return
}

type FSDTargetEvent struct {
	Name                  string    `json:"Name,omitempty"`
	RemainingJumpsInRoute int       `json:"RemainingJumpsInRoute,omitempty"`
	StarClass             string    `json:"StarClass,omitempty"`
	SystemAddress         int       `json:"SystemAddress,omitempty"`
	Event                 string    `json:"event,omitempty"`
	Timestamp             time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evFSDTarget(ev *FSDTargetEvent) {
	return
}

type FSSAllBodiesFoundEvent struct {
	Count         int       `json:"Count,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	SystemName    string    `json:"SystemName,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evFSSAllBodiesFound(ev *FSSAllBodiesFoundEvent) {
	return
}

type FSSDiscoveryScanEvent struct {
	BodyCount     int       `json:"BodyCount,omitempty"`
	NonBodyCount  int       `json:"NonBodyCount,omitempty"`
	Progress      float64   `json:"Progress,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	SystemName    string    `json:"SystemName,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evFSSDiscoveryScan(ev *FSSDiscoveryScanEvent) {
	return
}

type FSSSignalDiscoveredEvent struct {
	IsStation                bool      `json:"IsStation,omitempty"`
	SignalName               string    `json:"SignalName,omitempty"`
	SignalNameLocalised      string    `json:"SignalName_Localised,omitempty"`
	SpawningFaction          string    `json:"SpawningFaction,omitempty"`
	SpawningFactionLocalised string    `json:"SpawningFaction_Localised,omitempty"`
	SpawningState            string    `json:"SpawningState,omitempty"`
	SpawningStateLocalised   string    `json:"SpawningState_Localised,omitempty"`
	SystemAddress            int       `json:"SystemAddress,omitempty"`
	ThreatLevel              int       `json:"ThreatLevel,omitempty"`
	TimeRemaining            float64   `json:"TimeRemaining,omitempty"`
	UssType                  string    `json:"USSType,omitempty"`
	USSTypeLocalised         string    `json:"USSType_Localised,omitempty"`
	Event                    string    `json:"event,omitempty"`
	Timestamp                time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evFSSSignalDiscovered(ev *FSSSignalDiscoveredEvent) {
	return
}

type FuelScoopEvent struct {
	Scooped   float64   `json:"Scooped,omitempty"`
	Total     float64   `json:"Total,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evFuelScoop(ev *FuelScoopEvent) {
	return
}

type HeatDamageEvent interface{}

func (h *handler) evHeatDamage(ev *HeatDamageEvent) {
	return
}

type HeatWarningEvent interface{}

func (h *handler) evHeatWarning(ev *HeatWarningEvent) {
	return
}

type HullDamageEvent struct {
	Fighter     bool      `json:"Fighter,omitempty"`
	Health      float64   `json:"Health,omitempty"`
	PlayerPilot bool      `json:"PlayerPilot,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evHullDamage(ev *HullDamageEvent) {
	return
}

type InterdictedEvent struct {
	CombatRank  int       `json:"CombatRank,omitempty"`
	Faction     string    `json:"Faction,omitempty"`
	Interdictor string    `json:"Interdictor,omitempty"`
	IsPlayer    bool      `json:"IsPlayer,omitempty"`
	Submitted   bool      `json:"Submitted,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evInterdicted(ev *InterdictedEvent) {
	return
}

type InvitedToSquadronEvent struct {
	SquadronName string    `json:"SquadronName,omitempty"`
	Event        string    `json:"event,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evInvitedToSquadron(ev *InvitedToSquadronEvent) {
	return
}

type JetConeBoostEvent struct {
	BoostValue float64   `json:"BoostValue,omitempty"`
	Event      string    `json:"event,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evJetConeBoost(ev *JetConeBoostEvent) {
	return
}

type LaunchFighterEvent struct {
	ID               int       `json:"ID,omitempty"`
	Loadout          string    `json:"Loadout,omitempty"`
	PlayerControlled bool      `json:"PlayerControlled,omitempty"`
	Event            string    `json:"event,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evLaunchFighter(ev *LaunchFighterEvent) {
	return
}

type LaunchSRVEvent struct {
	ID               int       `json:"ID,omitempty"`
	Loadout          string    `json:"Loadout,omitempty"`
	PlayerControlled bool      `json:"PlayerControlled,omitempty"`
	Event            string    `json:"event,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evLaunchSRV(ev *LaunchSRVEvent) {
	return
}

type LeaveBodyEvent struct {
	Body          string    `json:"Body,omitempty"`
	BodyID        int       `json:"BodyID,omitempty"`
	StarSystem    string    `json:"StarSystem,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evLeaveBody(ev *LeaveBodyEvent) {
	return
}

type LiftoffEvent struct {
	Body                        string    `json:"Body,omitempty"`
	BodyID                      int       `json:"BodyID,omitempty"`
	Latitude                    float64   `json:"Latitude,omitempty"`
	Longitude                   float64   `json:"Longitude,omitempty"`
	Multicrew                   bool      `json:"Multicrew,omitempty"`
	NearestDestination          string    `json:"NearestDestination,omitempty"`
	NearestDestinationLocalised string    `json:"NearestDestination_Localised,omitempty"`
	OnPlanet                    bool      `json:"OnPlanet,omitempty"`
	OnStation                   bool      `json:"OnStation,omitempty"`
	PlayerControlled            bool      `json:"PlayerControlled,omitempty"`
	StarSystem                  string    `json:"StarSystem,omitempty"`
	SystemAddress               int       `json:"SystemAddress,omitempty"`
	Taxi                        bool      `json:"Taxi,omitempty"`
	Event                       string    `json:"event,omitempty"`
	Timestamp                   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evLiftoff(ev *LiftoffEvent) {
	return
}

type LoadGameEvent struct {
	Commander     string    `json:"Commander,omitempty"`
	Credits       int       `json:"Credits,omitempty"`
	Fid           string    `json:"FID,omitempty"`
	FuelCapacity  float64   `json:"FuelCapacity,omitempty"`
	FuelLevel     float64   `json:"FuelLevel,omitempty"`
	GameMode      string    `json:"GameMode,omitempty"`
	Horizons      bool      `json:"Horizons,omitempty"`
	Loan          int       `json:"Loan,omitempty"`
	Odyssey       bool      `json:"Odyssey,omitempty"`
	Ship          string    `json:"Ship,omitempty"`
	ShipID        int       `json:"ShipID,omitempty"`
	ShipIdent     string    `json:"ShipIdent,omitempty"`
	ShipName      string    `json:"ShipName,omitempty"`
	ShipLocalised string    `json:"Ship_Localised,omitempty"`
	StartLanded   bool      `json:"StartLanded,omitempty"`
	Build         string    `json:"build,omitempty"`
	Event         string    `json:"event,omitempty"`
	Gameversion   string    `json:"gameversion,omitempty"`
	Language      string    `json:"language,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evLoadGame(ev *LoadGameEvent) {
	return
}

type LoadoutEvent struct {
	CargoCapacity int `json:"CargoCapacity,omitempty"`
	FuelCapacity  struct {
		Main    float64 `json:"Main,omitempty"`
		Reserve float64 `json:"Reserve,omitempty"`
	} `json:"FuelCapacity,omitempty"`
	HullHealth   float64 `json:"HullHealth,omitempty"`
	HullValue    int     `json:"HullValue,omitempty"`
	MaxJumpRange float64 `json:"MaxJumpRange,omitempty"`
	Modules      []struct {
		AmmoInClip   int `json:"AmmoInClip,omitempty"`
		AmmoInHopper int `json:"AmmoInHopper,omitempty"`
		Engineering  *struct {
			BlueprintID                 int    `json:"BlueprintID,omitempty"`
			BlueprintName               string `json:"BlueprintName,omitempty"`
			Engineer                    string `json:"Engineer,omitempty"`
			EngineerID                  int    `json:"EngineerID,omitempty"`
			ExperimentalEffect          string `json:"ExperimentalEffect,omitempty"`
			ExperimentalEffectLocalised string `json:"ExperimentalEffect_Localised,omitempty"`
			Level                       int    `json:"Level,omitempty"`
			Modifiers                   []struct {
				Label         string  `json:"Label,omitempty"`
				LessIsGood    int     `json:"LessIsGood,omitempty"`
				OriginalValue float64 `json:"OriginalValue,omitempty"`
				Value         float64 `json:"Value,omitempty"`
			} `json:"Modifiers,omitempty"`
			Quality float64 `json:"Quality,omitempty"`
		} `json:"Engineering,omitempty"`
		Health   float64 `json:"Health,omitempty"`
		Item     string  `json:"Item,omitempty"`
		On       bool    `json:"On,omitempty"`
		Priority int     `json:"Priority,omitempty"`
		Slot     string  `json:"Slot,omitempty"`
		Value    int     `json:"Value,omitempty"`
	} `json:"Modules,omitempty"`
	ModulesValue int       `json:"ModulesValue,omitempty"`
	Rebuy        int       `json:"Rebuy,omitempty"`
	Ship         string    `json:"Ship,omitempty"`
	ShipID       int       `json:"ShipID,omitempty"`
	ShipIdent    string    `json:"ShipIdent,omitempty"`
	ShipName     string    `json:"ShipName,omitempty"`
	UnladenMass  float64   `json:"UnladenMass,omitempty"`
	Event        string    `json:"event,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evLoadout(ev *LoadoutEvent) {
	return
}

type LoadoutEquipModuleEvent struct {
	Class               int           `json:"Class,omitempty"`
	LoadoutID           int           `json:"LoadoutID,omitempty"`
	LoadoutName         string        `json:"LoadoutName,omitempty"`
	ModuleName          string        `json:"ModuleName,omitempty"`
	ModuleNameLocalised string        `json:"ModuleName_Localised,omitempty"`
	SlotName            string        `json:"SlotName,omitempty"`
	SuitID              int           `json:"SuitID,omitempty"`
	SuitModuleID        int           `json:"SuitModuleID,omitempty"`
	SuitName            string        `json:"SuitName,omitempty"`
	SuitNameLocalised   string        `json:"SuitName_Localised,omitempty"`
	WeaponMods          []interface{} `json:"WeaponMods,omitempty"`
	Event               string        `json:"event,omitempty"`
	Timestamp           time.Time     `json:"timestamp,omitempty"`
}

func (h *handler) evLoadoutEquipModule(ev *LoadoutEquipModuleEvent) {
	return
}

type LocationEvent struct {
	Body      string `json:"Body,omitempty"`
	BodyID    int    `json:"BodyID,omitempty"`
	BodyType  string `json:"BodyType,omitempty"`
	Conflicts []struct {
		Faction1 struct {
			Name    string `json:"Name,omitempty"`
			Stake   string `json:"Stake,omitempty"`
			WonDays int    `json:"WonDays,omitempty"`
		} `json:"Faction1,omitempty"`
		Faction2 struct {
			Name    string `json:"Name,omitempty"`
			Stake   string `json:"Stake,omitempty"`
			WonDays int    `json:"WonDays,omitempty"`
		} `json:"Faction2,omitempty"`
		Status  string `json:"Status,omitempty"`
		WarType string `json:"WarType,omitempty"`
	} `json:"Conflicts,omitempty"`
	DistFromStarLs float64 `json:"DistFromStarLS,omitempty"`
	Docked         bool    `json:"Docked,omitempty"`
	Factions       []struct {
		ActiveStates []struct {
			State string `json:"State,omitempty"`
		} `json:"ActiveStates,omitempty"`
		Allegiance         string  `json:"Allegiance,omitempty"`
		FactionState       string  `json:"FactionState,omitempty"`
		Government         string  `json:"Government,omitempty"`
		HappiestSystem     bool    `json:"HappiestSystem,omitempty"`
		Happiness          string  `json:"Happiness,omitempty"`
		HappinessLocalised string  `json:"Happiness_Localised,omitempty"`
		HomeSystem         bool    `json:"HomeSystem,omitempty"`
		Influence          float64 `json:"Influence,omitempty"`
		MyReputation       float64 `json:"MyReputation,omitempty"`
		Name               string  `json:"Name,omitempty"`
		PendingStates      []struct {
			State string `json:"State,omitempty"`
			Trend int    `json:"Trend,omitempty"`
		} `json:"PendingStates,omitempty"`
		RecoveringStates []struct {
			State string `json:"State,omitempty"`
			Trend int    `json:"Trend,omitempty"`
		} `json:"RecoveringStates,omitempty"`
		SquadronFaction bool `json:"SquadronFaction,omitempty"`
	} `json:"Factions,omitempty"`
	InSrv             bool      `json:"InSRV,omitempty"`
	Latitude          float64   `json:"Latitude,omitempty"`
	Longitude         float64   `json:"Longitude,omitempty"`
	MarketID          int       `json:"MarketID,omitempty"`
	Multicrew         bool      `json:"Multicrew,omitempty"`
	OnFoot            bool      `json:"OnFoot,omitempty"`
	Population        int       `json:"Population,omitempty"`
	PowerplayState    string    `json:"PowerplayState,omitempty"`
	Powers            []string  `json:"Powers,omitempty"`
	StarPos           []float64 `json:"StarPos,omitempty"`
	StarSystem        string    `json:"StarSystem,omitempty"`
	StationAllegiance string    `json:"StationAllegiance,omitempty"`
	StationEconomies  []struct {
		Name          string  `json:"Name,omitempty"`
		NameLocalised string  `json:"Name_Localised,omitempty"`
		Proportion    float64 `json:"Proportion,omitempty"`
	} `json:"StationEconomies,omitempty"`
	StationEconomy          string `json:"StationEconomy,omitempty"`
	StationEconomyLocalised string `json:"StationEconomy_Localised,omitempty"`
	StationFaction          *struct {
		FactionState string `json:"FactionState,omitempty"`
		Name         string `json:"Name,omitempty"`
	} `json:"StationFaction,omitempty"`
	StationGovernment          string   `json:"StationGovernment,omitempty"`
	StationGovernmentLocalised string   `json:"StationGovernment_Localised,omitempty"`
	StationName                string   `json:"StationName,omitempty"`
	StationServices            []string `json:"StationServices,omitempty"`
	StationType                string   `json:"StationType,omitempty"`
	SystemAddress              int      `json:"SystemAddress,omitempty"`
	SystemAllegiance           string   `json:"SystemAllegiance,omitempty"`
	SystemEconomy              string   `json:"SystemEconomy,omitempty"`
	SystemEconomyLocalised     string   `json:"SystemEconomy_Localised,omitempty"`
	SystemFaction              *struct {
		FactionState string `json:"FactionState,omitempty"`
		Name         string `json:"Name,omitempty"`
	} `json:"SystemFaction,omitempty"`
	SystemGovernment             string    `json:"SystemGovernment,omitempty"`
	SystemGovernmentLocalised    string    `json:"SystemGovernment_Localised,omitempty"`
	SystemSecondEconomy          string    `json:"SystemSecondEconomy,omitempty"`
	SystemSecondEconomyLocalised string    `json:"SystemSecondEconomy_Localised,omitempty"`
	SystemSecurity               string    `json:"SystemSecurity,omitempty"`
	SystemSecurityLocalised      string    `json:"SystemSecurity_Localised,omitempty"`
	Taxi                         bool      `json:"Taxi,omitempty"`
	Event                        string    `json:"event,omitempty"`
	Timestamp                    time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evLocation(ev *LocationEvent) {
	return
}

type MarketEvent struct {
	CarrierDockingAccess string    `json:"CarrierDockingAccess,omitempty"`
	MarketID             int       `json:"MarketID,omitempty"`
	StarSystem           string    `json:"StarSystem,omitempty"`
	StationName          string    `json:"StationName,omitempty"`
	StationType          string    `json:"StationType,omitempty"`
	Event                string    `json:"event,omitempty"`
	Timestamp            time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMarket(ev *MarketEvent) {
	return
}

type MarketBuyEvent struct {
	BuyPrice      int       `json:"BuyPrice,omitempty"`
	Count         int       `json:"Count,omitempty"`
	MarketID      int       `json:"MarketID,omitempty"`
	TotalCost     int       `json:"TotalCost,omitempty"`
	Type          string    `json:"Type,omitempty"`
	TypeLocalised string    `json:"Type_Localised,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMarketBuy(ev *MarketBuyEvent) {
	return
}

type MarketSellEvent struct {
	AvgPricePaid  int       `json:"AvgPricePaid,omitempty"`
	BlackMarket   bool      `json:"BlackMarket,omitempty"`
	Count         int       `json:"Count,omitempty"`
	IllegalGoods  bool      `json:"IllegalGoods,omitempty"`
	MarketID      int       `json:"MarketID,omitempty"`
	SellPrice     int       `json:"SellPrice,omitempty"`
	StolenGoods   bool      `json:"StolenGoods,omitempty"`
	TotalSale     int       `json:"TotalSale,omitempty"`
	Type          string    `json:"Type,omitempty"`
	TypeLocalised string    `json:"Type_Localised,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMarketSell(ev *MarketSellEvent) {
	return
}

type MaterialCollectedEvent struct {
	Category      string    `json:"Category,omitempty"`
	Count         int       `json:"Count,omitempty"`
	Name          string    `json:"Name,omitempty"`
	NameLocalised string    `json:"Name_Localised,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMaterialCollected(ev *MaterialCollectedEvent) {
	return
}

type MaterialDiscoveredEvent struct {
	Category        string    `json:"Category,omitempty"`
	DiscoveryNumber int       `json:"DiscoveryNumber,omitempty"`
	Name            string    `json:"Name,omitempty"`
	NameLocalised   string    `json:"Name_Localised,omitempty"`
	Event           string    `json:"event,omitempty"`
	Timestamp       time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMaterialDiscovered(ev *MaterialDiscoveredEvent) {
	return
}

type MaterialsEvent struct {
	Encoded []struct {
		Count         int    `json:"Count,omitempty"`
		Name          string `json:"Name,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
	} `json:"Encoded,omitempty"`
	Manufactured []struct {
		Count         int    `json:"Count,omitempty"`
		Name          string `json:"Name,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
	} `json:"Manufactured,omitempty"`
	Raw []struct {
		Count int    `json:"Count,omitempty"`
		Name  string `json:"Name,omitempty"`
	} `json:"Raw,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMaterials(ev *MaterialsEvent) {
	return
}

type MaterialTradeEvent struct {
	MarketID int `json:"MarketID,omitempty"`
	Paid     struct {
		Category          string `json:"Category,omitempty"`
		Material          string `json:"Material,omitempty"`
		MaterialLocalised string `json:"Material_Localised,omitempty"`
		Quantity          int    `json:"Quantity,omitempty"`
	} `json:"Paid,omitempty"`
	Received struct {
		Category          string `json:"Category,omitempty"`
		Material          string `json:"Material,omitempty"`
		MaterialLocalised string `json:"Material_Localised,omitempty"`
		Quantity          int    `json:"Quantity,omitempty"`
	} `json:"Received,omitempty"`
	TraderType string    `json:"TraderType,omitempty"`
	Event      string    `json:"event,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMaterialTrade(ev *MaterialTradeEvent) {
	return
}

type MissionAbandonedEvent struct {
	MissionID int       `json:"MissionID,omitempty"`
	Name      string    `json:"Name,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMissionAbandoned(ev *MissionAbandonedEvent) {
	return
}

type MissionAcceptedEvent struct {
	Commodity           string    `json:"Commodity,omitempty"`
	CommodityLocalised  string    `json:"Commodity_Localised,omitempty"`
	Count               int       `json:"Count,omitempty"`
	DestinationStation  string    `json:"DestinationStation,omitempty"`
	DestinationSystem   string    `json:"DestinationSystem,omitempty"`
	Donation            string    `json:"Donation,omitempty"`
	Expiry              time.Time `json:"Expiry,omitempty"`
	Faction             string    `json:"Faction,omitempty"`
	Influence           string    `json:"Influence,omitempty"`
	LocalisedName       string    `json:"LocalisedName,omitempty"`
	MissionID           int       `json:"MissionID,omitempty"`
	Name                string    `json:"Name,omitempty"`
	PassengerCount      int       `json:"PassengerCount,omitempty"`
	PassengerType       string    `json:"PassengerType,omitempty"`
	PassengerViPs       bool      `json:"PassengerVIPs,omitempty"`
	PassengerWanted     bool      `json:"PassengerWanted,omitempty"`
	Reputation          string    `json:"Reputation,omitempty"`
	Reward              int       `json:"Reward,omitempty"`
	Target              string    `json:"Target,omitempty"`
	TargetFaction       string    `json:"TargetFaction,omitempty"`
	TargetType          string    `json:"TargetType,omitempty"`
	TargetTypeLocalised string    `json:"TargetType_Localised,omitempty"`
	TargetLocalised     string    `json:"Target_Localised,omitempty"`
	Wing                bool      `json:"Wing,omitempty"`
	Event               string    `json:"event,omitempty"`
	Timestamp           time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMissionAccepted(ev *MissionAcceptedEvent) {
	return
}

type MissionCompletedEvent struct {
	Commodity       string `json:"Commodity,omitempty"`
	CommodityReward []struct {
		Count         int    `json:"Count,omitempty"`
		Name          string `json:"Name,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
	} `json:"CommodityReward,omitempty"`
	CommodityLocalised string `json:"Commodity_Localised,omitempty"`
	Count              int    `json:"Count,omitempty"`
	DestinationStation string `json:"DestinationStation,omitempty"`
	DestinationSystem  string `json:"DestinationSystem,omitempty"`
	Donated            int    `json:"Donated,omitempty"`
	Donation           string `json:"Donation,omitempty"`
	Faction            string `json:"Faction,omitempty"`
	FactionEffects     []struct {
		Effects []struct {
			Effect          string `json:"Effect,omitempty"`
			EffectLocalised string `json:"Effect_Localised,omitempty"`
			Trend           string `json:"Trend,omitempty"`
		} `json:"Effects,omitempty"`
		Faction   string `json:"Faction,omitempty"`
		Influence []struct {
			Influence     string `json:"Influence,omitempty"`
			SystemAddress int    `json:"SystemAddress,omitempty"`
			Trend         string `json:"Trend,omitempty"`
		} `json:"Influence,omitempty"`
		Reputation      string `json:"Reputation,omitempty"`
		ReputationTrend string `json:"ReputationTrend,omitempty"`
	} `json:"FactionEffects,omitempty"`
	MaterialsReward []struct {
		Category          string `json:"Category,omitempty"`
		CategoryLocalised string `json:"Category_Localised,omitempty"`
		Count             int    `json:"Count,omitempty"`
		Name              string `json:"Name,omitempty"`
		NameLocalised     string `json:"Name_Localised,omitempty"`
	} `json:"MaterialsReward,omitempty"`
	MissionID            int       `json:"MissionID,omitempty"`
	Name                 string    `json:"Name,omitempty"`
	NewDestinationSystem string    `json:"NewDestinationSystem,omitempty"`
	Reward               int       `json:"Reward,omitempty"`
	Target               string    `json:"Target,omitempty"`
	TargetFaction        string    `json:"TargetFaction,omitempty"`
	TargetType           string    `json:"TargetType,omitempty"`
	TargetTypeLocalised  string    `json:"TargetType_Localised,omitempty"`
	TargetLocalised      string    `json:"Target_Localised,omitempty"`
	Event                string    `json:"event,omitempty"`
	Timestamp            time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMissionCompleted(ev *MissionCompletedEvent) {
	return
}

type MissionFailedEvent struct {
	MissionID int       `json:"MissionID,omitempty"`
	Name      string    `json:"Name,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMissionFailed(ev *MissionFailedEvent) {
	return
}

type MissionRedirectedEvent struct {
	MissionID             int       `json:"MissionID,omitempty"`
	Name                  string    `json:"Name,omitempty"`
	NewDestinationStation string    `json:"NewDestinationStation,omitempty"`
	NewDestinationSystem  string    `json:"NewDestinationSystem,omitempty"`
	OldDestinationStation string    `json:"OldDestinationStation,omitempty"`
	OldDestinationSystem  string    `json:"OldDestinationSystem,omitempty"`
	Event                 string    `json:"event,omitempty"`
	Timestamp             time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMissionRedirected(ev *MissionRedirectedEvent) {
	return
}

type MissionsEvent struct {
	Active []struct {
		Expires          int    `json:"Expires,omitempty"`
		MissionID        int    `json:"MissionID,omitempty"`
		Name             string `json:"Name,omitempty"`
		PassengerMission bool   `json:"PassengerMission,omitempty"`
	} `json:"Active,omitempty"`
	Complete []interface{} `json:"Complete,omitempty"`
	Failed   []struct {
		Expires          int    `json:"Expires,omitempty"`
		MissionID        int    `json:"MissionID,omitempty"`
		Name             string `json:"Name,omitempty"`
		PassengerMission bool   `json:"PassengerMission,omitempty"`
	} `json:"Failed,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMissions(ev *MissionsEvent) {
	return
}

type ModuleBuyEvent struct {
	BuyItem             string    `json:"BuyItem,omitempty"`
	BuyItemLocalised    string    `json:"BuyItem_Localised,omitempty"`
	BuyPrice            int       `json:"BuyPrice,omitempty"`
	MarketID            int       `json:"MarketID,omitempty"`
	SellItem            string    `json:"SellItem,omitempty"`
	SellItemLocalised   string    `json:"SellItem_Localised,omitempty"`
	SellPrice           int       `json:"SellPrice,omitempty"`
	Ship                string    `json:"Ship,omitempty"`
	ShipID              int       `json:"ShipID,omitempty"`
	Slot                string    `json:"Slot,omitempty"`
	StoredItem          string    `json:"StoredItem,omitempty"`
	StoredItemLocalised string    `json:"StoredItem_Localised,omitempty"`
	Event               string    `json:"event,omitempty"`
	Timestamp           time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evModuleBuy(ev *ModuleBuyEvent) {
	return
}

type ModuleInfoEvent interface{}

func (h *handler) evModuleInfo(ev *ModuleInfoEvent) {
	return
}

type ModuleRetrieveEvent struct {
	EngineerModifications  string    `json:"EngineerModifications,omitempty"`
	Hot                    bool      `json:"Hot,omitempty"`
	Level                  int       `json:"Level,omitempty"`
	MarketID               int       `json:"MarketID,omitempty"`
	Quality                float64   `json:"Quality,omitempty"`
	RetrievedItem          string    `json:"RetrievedItem,omitempty"`
	RetrievedItemLocalised string    `json:"RetrievedItem_Localised,omitempty"`
	Ship                   string    `json:"Ship,omitempty"`
	ShipID                 int       `json:"ShipID,omitempty"`
	Slot                   string    `json:"Slot,omitempty"`
	SwapOutItem            string    `json:"SwapOutItem,omitempty"`
	SwapOutItemLocalised   string    `json:"SwapOutItem_Localised,omitempty"`
	Event                  string    `json:"event,omitempty"`
	Timestamp              time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evModuleRetrieve(ev *ModuleRetrieveEvent) {
	return
}

type ModuleSellEvent struct {
	MarketID          int       `json:"MarketID,omitempty"`
	SellItem          string    `json:"SellItem,omitempty"`
	SellItemLocalised string    `json:"SellItem_Localised,omitempty"`
	SellPrice         int       `json:"SellPrice,omitempty"`
	Ship              string    `json:"Ship,omitempty"`
	ShipID            int       `json:"ShipID,omitempty"`
	Slot              string    `json:"Slot,omitempty"`
	Event             string    `json:"event,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evModuleSell(ev *ModuleSellEvent) {
	return
}

type ModuleSellRemoteEvent struct {
	SellItem          string    `json:"SellItem,omitempty"`
	SellItemLocalised string    `json:"SellItem_Localised,omitempty"`
	SellPrice         int       `json:"SellPrice,omitempty"`
	ServerID          int       `json:"ServerId,omitempty"`
	Ship              string    `json:"Ship,omitempty"`
	ShipID            int       `json:"ShipID,omitempty"`
	StorageSlot       int       `json:"StorageSlot,omitempty"`
	Event             string    `json:"event,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evModuleSellRemote(ev *ModuleSellRemoteEvent) {
	return
}

type ModuleStoreEvent struct {
	EngineerModifications string    `json:"EngineerModifications,omitempty"`
	Hot                   bool      `json:"Hot,omitempty"`
	Level                 int       `json:"Level,omitempty"`
	MarketID              int       `json:"MarketID,omitempty"`
	Quality               float64   `json:"Quality,omitempty"`
	Ship                  string    `json:"Ship,omitempty"`
	ShipID                int       `json:"ShipID,omitempty"`
	Slot                  string    `json:"Slot,omitempty"`
	StoredItem            string    `json:"StoredItem,omitempty"`
	StoredItemLocalised   string    `json:"StoredItem_Localised,omitempty"`
	Event                 string    `json:"event,omitempty"`
	Timestamp             time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evModuleStore(ev *ModuleStoreEvent) {
	return
}

type ModuleSwapEvent struct {
	FromItem          string    `json:"FromItem,omitempty"`
	FromItemLocalised string    `json:"FromItem_Localised,omitempty"`
	FromSlot          string    `json:"FromSlot,omitempty"`
	MarketID          int       `json:"MarketID,omitempty"`
	Ship              string    `json:"Ship,omitempty"`
	ShipID            int       `json:"ShipID,omitempty"`
	ToItem            string    `json:"ToItem,omitempty"`
	ToSlot            string    `json:"ToSlot,omitempty"`
	Event             string    `json:"event,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evModuleSwap(ev *ModuleSwapEvent) {
	return
}

type MultiSellExplorationDataEvent struct {
	BaseValue  int `json:"BaseValue,omitempty"`
	Bonus      int `json:"Bonus,omitempty"`
	Discovered []struct {
		NumBodies  int    `json:"NumBodies,omitempty"`
		SystemName string `json:"SystemName,omitempty"`
	} `json:"Discovered,omitempty"`
	TotalEarnings int       `json:"TotalEarnings,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMultiSellExplorationData(ev *MultiSellExplorationDataEvent) {
	return
}

type MusicEvent struct {
	MusicTrack string    `json:"MusicTrack,omitempty"`
	Event      string    `json:"event,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evMusic(ev *MusicEvent) {
	return
}

type NavRouteEvent interface{}

func (h *handler) evNavRoute(ev *NavRouteEvent) {
	return
}

type NewCommanderEvent struct {
	Fid       string    `json:"FID,omitempty"`
	Name      string    `json:"Name,omitempty"`
	Package   string    `json:"Package,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evNewCommander(ev *NewCommanderEvent) {
	return
}

type NpcCrewPaidWageEvent struct {
	Amount      int       `json:"Amount,omitempty"`
	NpcCrewID   int       `json:"NpcCrewId,omitempty"`
	NpcCrewName string    `json:"NpcCrewName,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evNpcCrewPaidWage(ev *NpcCrewPaidWageEvent) {
	return
}

type OutfittingEvent struct {
	MarketID    int       `json:"MarketID,omitempty"`
	StarSystem  string    `json:"StarSystem,omitempty"`
	StationName string    `json:"StationName,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evOutfitting(ev *OutfittingEvent) {
	return
}

type PassengersEvent struct {
	Manifest []struct {
		Count     int    `json:"Count,omitempty"`
		MissionID int    `json:"MissionID,omitempty"`
		Type      string `json:"Type,omitempty"`
		Vip       bool   `json:"VIP,omitempty"`
		Wanted    bool   `json:"Wanted,omitempty"`
	} `json:"Manifest,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evPassengers(ev *PassengersEvent) {
	return
}

type PayBountiesEvent struct {
	Amount           int       `json:"Amount,omitempty"`
	BrokerPercentage float64   `json:"BrokerPercentage,omitempty"`
	Faction          string    `json:"Faction,omitempty"`
	FactionLocalised string    `json:"Faction_Localised,omitempty"`
	ShipID           int       `json:"ShipID,omitempty"`
	Event            string    `json:"event,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evPayBounties(ev *PayBountiesEvent) {
	return
}

type PayFinesEvent struct {
	AllFines  bool      `json:"AllFines,omitempty"`
	Amount    int       `json:"Amount,omitempty"`
	Faction   string    `json:"Faction,omitempty"`
	ShipID    int       `json:"ShipID,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evPayFines(ev *PayFinesEvent) {
	return
}

type PowerplayEvent struct {
	Merits      int       `json:"Merits,omitempty"`
	Power       string    `json:"Power,omitempty"`
	Rank        int       `json:"Rank,omitempty"`
	TimePledged int       `json:"TimePledged,omitempty"`
	Votes       int       `json:"Votes,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evPowerplay(ev *PowerplayEvent) {
	return
}

type PowerplayJoinEvent struct {
	Power     string    `json:"Power,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evPowerplayJoin(ev *PowerplayJoinEvent) {
	return
}

type PowerplayLeaveEvent struct {
	Power     string    `json:"Power,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evPowerplayLeave(ev *PowerplayLeaveEvent) {
	return
}

type PowerplaySalaryEvent struct {
	Amount    int       `json:"Amount,omitempty"`
	Power     string    `json:"Power,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evPowerplaySalary(ev *PowerplaySalaryEvent) {
	return
}

type ProgressEvent struct {
	Cqc          int       `json:"CQC,omitempty"`
	Combat       int       `json:"Combat,omitempty"`
	Empire       int       `json:"Empire,omitempty"`
	Exobiologist int       `json:"Exobiologist,omitempty"`
	Explore      int       `json:"Explore,omitempty"`
	Federation   int       `json:"Federation,omitempty"`
	Soldier      int       `json:"Soldier,omitempty"`
	Trade        int       `json:"Trade,omitempty"`
	Event        string    `json:"event,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evProgress(ev *ProgressEvent) {
	return
}

type PromotionEvent struct {
	Empire       int       `json:"Empire,omitempty"`
	Exobiologist int       `json:"Exobiologist,omitempty"`
	Explore      int       `json:"Explore,omitempty"`
	Federation   int       `json:"Federation,omitempty"`
	Trade        int       `json:"Trade,omitempty"`
	Event        string    `json:"event,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evPromotion(ev *PromotionEvent) {
	return
}

type RankEvent struct {
	Cqc          int       `json:"CQC,omitempty"`
	Combat       int       `json:"Combat,omitempty"`
	Empire       int       `json:"Empire,omitempty"`
	Exobiologist int       `json:"Exobiologist,omitempty"`
	Explore      int       `json:"Explore,omitempty"`
	Federation   int       `json:"Federation,omitempty"`
	Soldier      int       `json:"Soldier,omitempty"`
	Trade        int       `json:"Trade,omitempty"`
	Event        string    `json:"event,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evRank(ev *RankEvent) {
	return
}

type RebootRepairEvent struct {
	Modules   []interface{} `json:"Modules,omitempty"`
	Event     string        `json:"event,omitempty"`
	Timestamp time.Time     `json:"timestamp,omitempty"`
}

func (h *handler) evRebootRepair(ev *RebootRepairEvent) {
	return
}

type ReceiveTextEvent struct {
	Channel          string    `json:"Channel,omitempty"`
	From             string    `json:"From,omitempty"`
	FromLocalised    string    `json:"From_Localised,omitempty"`
	Message          string    `json:"Message,omitempty"`
	MessageLocalised string    `json:"Message_Localised,omitempty"`
	Event            string    `json:"event,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evReceiveText(ev *ReceiveTextEvent) {
	return
}

type RedeemVoucherEvent struct {
	Amount           int     `json:"Amount,omitempty"`
	BrokerPercentage float64 `json:"BrokerPercentage,omitempty"`
	Faction          string  `json:"Faction,omitempty"`
	Factions         []struct {
		Amount  int    `json:"Amount,omitempty"`
		Faction string `json:"Faction,omitempty"`
	} `json:"Factions,omitempty"`
	Type      string    `json:"Type,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evRedeemVoucher(ev *RedeemVoucherEvent) {
	return
}

type RefuelAllEvent struct {
	Amount    float64   `json:"Amount,omitempty"`
	Cost      int       `json:"Cost,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evRefuelAll(ev *RefuelAllEvent) {
	return
}

type RefuelPartialEvent struct {
	Amount    float64   `json:"Amount,omitempty"`
	Cost      int       `json:"Cost,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evRefuelPartial(ev *RefuelPartialEvent) {
	return
}

type RenameSuitLoadoutEvent struct {
	LoadoutID         int       `json:"LoadoutID,omitempty"`
	LoadoutName       string    `json:"LoadoutName,omitempty"`
	SuitID            int       `json:"SuitID,omitempty"`
	SuitName          string    `json:"SuitName,omitempty"`
	SuitNameLocalised string    `json:"SuitName_Localised,omitempty"`
	Event             string    `json:"event,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evRenameSuitLoadout(ev *RenameSuitLoadoutEvent) {
	return
}

type RepairEvent struct {
	Cost      int       `json:"Cost,omitempty"`
	Item      string    `json:"Item,omitempty"`
	Items     []string  `json:"Items,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evRepair(ev *RepairEvent) {
	return
}

type RepairAllEvent struct {
	Cost      int       `json:"Cost,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evRepairAll(ev *RepairAllEvent) {
	return
}

type ReputationEvent struct {
	Alliance   float64   `json:"Alliance,omitempty"`
	Empire     float64   `json:"Empire,omitempty"`
	Federation float64   `json:"Federation,omitempty"`
	Event      string    `json:"event,omitempty"`
	Timestamp  time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evReputation(ev *ReputationEvent) {
	return
}

type ReservoirReplenishedEvent struct {
	FuelMain      float64   `json:"FuelMain,omitempty"`
	FuelReservoir float64   `json:"FuelReservoir,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evReservoirReplenished(ev *ReservoirReplenishedEvent) {
	return
}

type ResurrectEvent struct {
	Bankrupt  bool      `json:"Bankrupt,omitempty"`
	Cost      int       `json:"Cost,omitempty"`
	Option    string    `json:"Option,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evResurrect(ev *ResurrectEvent) {
	return
}

type SAAScanCompleteEvent struct {
	BodyID           int       `json:"BodyID,omitempty"`
	BodyName         string    `json:"BodyName,omitempty"`
	EfficiencyTarget int       `json:"EfficiencyTarget,omitempty"`
	ProbesUsed       int       `json:"ProbesUsed,omitempty"`
	SystemAddress    int       `json:"SystemAddress,omitempty"`
	Event            string    `json:"event,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evSAAScanComplete(ev *SAAScanCompleteEvent) {
	return
}

type SAASignalsFoundEvent struct {
	BodyID   int    `json:"BodyID,omitempty"`
	BodyName string `json:"BodyName,omitempty"`
	Signals  []struct {
		Count         int    `json:"Count,omitempty"`
		Type          string `json:"Type,omitempty"`
		TypeLocalised string `json:"Type_Localised,omitempty"`
	} `json:"Signals,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evSAASignalsFound(ev *SAASignalsFoundEvent) {
	return
}

type ScanEvent struct {
	AbsoluteMagnitude     float64 `json:"AbsoluteMagnitude,omitempty"`
	AgeMy                 int     `json:"Age_MY,omitempty"`
	Atmosphere            string  `json:"Atmosphere,omitempty"`
	AtmosphereComposition []struct {
		Name    string  `json:"Name,omitempty"`
		Percent float64 `json:"Percent,omitempty"`
	} `json:"AtmosphereComposition,omitempty"`
	AtmosphereType string  `json:"AtmosphereType,omitempty"`
	AxialTilt      float64 `json:"AxialTilt,omitempty"`
	BodyID         int     `json:"BodyID,omitempty"`
	BodyName       string  `json:"BodyName,omitempty"`
	Composition    *struct {
		Ice   float64 `json:"Ice,omitempty"`
		Metal float64 `json:"Metal,omitempty"`
		Rock  float64 `json:"Rock,omitempty"`
	} `json:"Composition,omitempty"`
	DistanceFromArrivalLs float64 `json:"DistanceFromArrivalLS,omitempty"`
	Eccentricity          float64 `json:"Eccentricity,omitempty"`
	Landable              bool    `json:"Landable,omitempty"`
	Luminosity            string  `json:"Luminosity,omitempty"`
	MassEm                float64 `json:"MassEM,omitempty"`
	Materials             []struct {
		Name    string  `json:"Name,omitempty"`
		Percent float64 `json:"Percent,omitempty"`
	} `json:"Materials,omitempty"`
	OrbitalInclination float64 `json:"OrbitalInclination,omitempty"`
	OrbitalPeriod      float64 `json:"OrbitalPeriod,omitempty"`
	Parents            []struct {
		Null   int `json:"Null,omitempty"`
		Planet int `json:"Planet,omitempty"`
		Ring   int `json:"Ring,omitempty"`
		Star   int `json:"Star,omitempty"`
	} `json:"Parents,omitempty"`
	Periapsis    float64 `json:"Periapsis,omitempty"`
	PlanetClass  string  `json:"PlanetClass,omitempty"`
	Radius       float64 `json:"Radius,omitempty"`
	ReserveLevel string  `json:"ReserveLevel,omitempty"`
	Rings        []struct {
		InnerRad  json.Number `json:"InnerRad,omitempty"`
		MassMt    json.Number `json:"MassMT,omitempty"`
		Name      string      `json:"Name,omitempty"`
		OuterRad  float64     `json:"OuterRad,omitempty"`
		RingClass string      `json:"RingClass,omitempty"`
	} `json:"Rings,omitempty"`
	RotationPeriod     float64   `json:"RotationPeriod,omitempty"`
	ScanType           string    `json:"ScanType,omitempty"`
	SemiMajorAxis      float64   `json:"SemiMajorAxis,omitempty"`
	StarSystem         string    `json:"StarSystem,omitempty"`
	StarType           string    `json:"StarType,omitempty"`
	StellarMass        float64   `json:"StellarMass,omitempty"`
	Subclass           int       `json:"Subclass,omitempty"`
	SurfaceGravity     float64   `json:"SurfaceGravity,omitempty"`
	SurfacePressure    float64   `json:"SurfacePressure,omitempty"`
	SurfaceTemperature float64   `json:"SurfaceTemperature,omitempty"`
	SystemAddress      int       `json:"SystemAddress,omitempty"`
	TerraformState     string    `json:"TerraformState,omitempty"`
	TidalLock          bool      `json:"TidalLock,omitempty"`
	Volcanism          string    `json:"Volcanism,omitempty"`
	WasDiscovered      bool      `json:"WasDiscovered,omitempty"`
	WasMapped          bool      `json:"WasMapped,omitempty"`
	Event              string    `json:"event,omitempty"`
	Timestamp          time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evScan(ev *ScanEvent) {
	m := &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort: display.VP_SRV,
			Text:     fmt.Sprintf("%+v", ev),
			Append:   false,
			SubTitle: "(123)",
		},
	}
	h.connector.ToRouterCh <- m
	return
}

type ScannedEvent struct {
	ScanType  string    `json:"ScanType,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evScanned(ev *ScannedEvent) {
	return
}

type ScanOrganicEvent struct {
	Body             int       `json:"Body,omitempty"`
	Genus            string    `json:"Genus,omitempty"`
	GenusLocalised   string    `json:"Genus_Localised,omitempty"`
	ScanType         string    `json:"ScanType,omitempty"`
	Species          string    `json:"Species,omitempty"`
	SpeciesLocalised string    `json:"Species_Localised,omitempty"`
	SystemAddress    int       `json:"SystemAddress,omitempty"`
	Event            string    `json:"event,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evScanOrganic(ev *ScanOrganicEvent) {
	return
}

type SellDronesEvent struct {
	Count     int       `json:"Count,omitempty"`
	SellPrice int       `json:"SellPrice,omitempty"`
	TotalSale int       `json:"TotalSale,omitempty"`
	Type      string    `json:"Type,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evSellDrones(ev *SellDronesEvent) {
	return
}

type SellExplorationDataEvent struct {
	BaseValue     int           `json:"BaseValue,omitempty"`
	Bonus         int           `json:"Bonus,omitempty"`
	Discovered    []interface{} `json:"Discovered,omitempty"`
	Systems       []string      `json:"Systems,omitempty"`
	TotalEarnings int           `json:"TotalEarnings,omitempty"`
	Event         string        `json:"event,omitempty"`
	Timestamp     time.Time     `json:"timestamp,omitempty"`
}

func (h *handler) evSellExplorationData(ev *SellExplorationDataEvent) {
	return
}

type SellOrganicDataEvent struct {
	BioData []struct {
		Bonus            int    `json:"Bonus,omitempty"`
		Genus            string `json:"Genus,omitempty"`
		GenusLocalised   string `json:"Genus_Localised,omitempty"`
		Species          string `json:"Species,omitempty"`
		SpeciesLocalised string `json:"Species_Localised,omitempty"`
		Value            int    `json:"Value,omitempty"`
	} `json:"BioData,omitempty"`
	MarketID  int       `json:"MarketID,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evSellOrganicData(ev *SellOrganicDataEvent) {
	return
}

type SellSuitEvent struct {
	Name          string        `json:"Name,omitempty"`
	NameLocalised string        `json:"Name_Localised,omitempty"`
	Price         int           `json:"Price,omitempty"`
	SuitID        int           `json:"SuitID,omitempty"`
	SuitMods      []interface{} `json:"SuitMods,omitempty"`
	Event         string        `json:"event,omitempty"`
	Timestamp     time.Time     `json:"timestamp,omitempty"`
}

func (h *handler) evSellSuit(ev *SellSuitEvent) {
	return
}

type SendTextEvent struct {
	Message   string    `json:"Message,omitempty"`
	Sent      bool      `json:"Sent,omitempty"`
	To        string    `json:"To,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evSendText(ev *SendTextEvent) {
	return
}

type SetUserShipNameEvent struct {
	Ship         string    `json:"Ship,omitempty"`
	ShipID       int       `json:"ShipID,omitempty"`
	UserShipID   string    `json:"UserShipId,omitempty"`
	UserShipName string    `json:"UserShipName,omitempty"`
	Event        string    `json:"event,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evSetUserShipName(ev *SetUserShipNameEvent) {
	return
}

type ShieldStateEvent struct {
	ShieldsUp bool      `json:"ShieldsUp,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evShieldState(ev *ShieldStateEvent) {
	return
}

type ShipLockerEvent struct {
	Components  []interface{} `json:"Components,omitempty"`
	Consumables []struct {
		Count         int    `json:"Count,omitempty"`
		Name          string `json:"Name,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
		OwnerID       int    `json:"OwnerID,omitempty"`
	} `json:"Consumables,omitempty"`
	Data      []interface{} `json:"Data,omitempty"`
	Items     []interface{} `json:"Items,omitempty"`
	Event     string        `json:"event,omitempty"`
	Timestamp time.Time     `json:"timestamp,omitempty"`
}

func (h *handler) evShipLocker(ev *ShipLockerEvent) {
	return
}

type ShipLockerMaterialsEvent struct {
	Components  []interface{} `json:"Components,omitempty"`
	Consumables []struct {
		Count         int    `json:"Count,omitempty"`
		Name          string `json:"Name,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
		OwnerID       int    `json:"OwnerID,omitempty"`
	} `json:"Consumables,omitempty"`
	Data      []interface{} `json:"Data,omitempty"`
	Items     []interface{} `json:"Items,omitempty"`
	Event     string        `json:"event,omitempty"`
	Timestamp time.Time     `json:"timestamp,omitempty"`
}

func (h *handler) evShipLockerMaterials(ev *ShipLockerMaterialsEvent) {
	return
}

type ShipTargetedEvent struct {
	Bounty             int       `json:"Bounty,omitempty"`
	Faction            string    `json:"Faction,omitempty"`
	HullHealth         float64   `json:"HullHealth,omitempty"`
	LegalStatus        string    `json:"LegalStatus,omitempty"`
	PilotName          string    `json:"PilotName,omitempty"`
	PilotNameLocalised string    `json:"PilotName_Localised,omitempty"`
	PilotRank          string    `json:"PilotRank,omitempty"`
	ScanStage          int       `json:"ScanStage,omitempty"`
	ShieldHealth       float64   `json:"ShieldHealth,omitempty"`
	Ship               string    `json:"Ship,omitempty"`
	ShipLocalised      string    `json:"Ship_Localised,omitempty"`
	SquadronID         string    `json:"SquadronID,omitempty"`
	TargetLocked       bool      `json:"TargetLocked,omitempty"`
	Event              string    `json:"event,omitempty"`
	Timestamp          time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evShipTargeted(ev *ShipTargetedEvent) {
	return
}

type ShipyardEvent struct {
	MarketID    int       `json:"MarketID,omitempty"`
	StarSystem  string    `json:"StarSystem,omitempty"`
	StationName string    `json:"StationName,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evShipyard(ev *ShipyardEvent) {
	return
}

type ShipyardBuyEvent struct {
	MarketID          int       `json:"MarketID,omitempty"`
	SellOldShip       string    `json:"SellOldShip,omitempty"`
	SellPrice         int       `json:"SellPrice,omitempty"`
	SellShipID        int       `json:"SellShipID,omitempty"`
	ShipPrice         int       `json:"ShipPrice,omitempty"`
	ShipType          string    `json:"ShipType,omitempty"`
	ShipTypeLocalised string    `json:"ShipType_Localised,omitempty"`
	StoreOldShip      string    `json:"StoreOldShip,omitempty"`
	StoreShipID       int       `json:"StoreShipID,omitempty"`
	Event             string    `json:"event,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evShipyardBuy(ev *ShipyardBuyEvent) {
	return
}

type ShipyardNewEvent struct {
	NewShipID         int       `json:"NewShipID,omitempty"`
	ShipType          string    `json:"ShipType,omitempty"`
	ShipTypeLocalised string    `json:"ShipType_Localised,omitempty"`
	Event             string    `json:"event,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evShipyardNew(ev *ShipyardNewEvent) {
	return
}

type ShipyardSellEvent struct {
	MarketID          int       `json:"MarketID,omitempty"`
	SellShipID        int       `json:"SellShipID,omitempty"`
	ShipPrice         int       `json:"ShipPrice,omitempty"`
	ShipType          string    `json:"ShipType,omitempty"`
	ShipTypeLocalised string    `json:"ShipType_Localised,omitempty"`
	Event             string    `json:"event,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evShipyardSell(ev *ShipyardSellEvent) {
	return
}

type ShipyardSwapEvent struct {
	MarketID          int       `json:"MarketID,omitempty"`
	ShipID            int       `json:"ShipID,omitempty"`
	ShipType          string    `json:"ShipType,omitempty"`
	ShipTypeLocalised string    `json:"ShipType_Localised,omitempty"`
	StoreOldShip      string    `json:"StoreOldShip,omitempty"`
	StoreShipID       int       `json:"StoreShipID,omitempty"`
	Event             string    `json:"event,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evShipyardSwap(ev *ShipyardSwapEvent) {
	return
}

type ShipyardTransferEvent struct {
	Distance          float64   `json:"Distance,omitempty"`
	MarketID          int       `json:"MarketID,omitempty"`
	ShipID            int       `json:"ShipID,omitempty"`
	ShipMarketID      int       `json:"ShipMarketID,omitempty"`
	ShipType          string    `json:"ShipType,omitempty"`
	ShipTypeLocalised string    `json:"ShipType_Localised,omitempty"`
	System            string    `json:"System,omitempty"`
	TransferPrice     int       `json:"TransferPrice,omitempty"`
	TransferTime      int       `json:"TransferTime,omitempty"`
	Event             string    `json:"event,omitempty"`
	Timestamp         time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evShipyardTransfer(ev *ShipyardTransferEvent) {
	return
}

type ShutdownEvent interface{}

func (h *handler) evShutdown(ev *ShutdownEvent) {
	return
}

type SquadronStartupEvent struct {
	CurrentRank  int       `json:"CurrentRank,omitempty"`
	SquadronName string    `json:"SquadronName,omitempty"`
	Event        string    `json:"event,omitempty"`
	Timestamp    time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evSquadronStartup(ev *SquadronStartupEvent) {
	return
}

type StartJumpEvent struct {
	JumpType      string    `json:"JumpType,omitempty"`
	StarClass     string    `json:"StarClass,omitempty"`
	StarSystem    string    `json:"StarSystem,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evStartJump(ev *StartJumpEvent) {
	return
}

type StatisticsEvent struct {
	BankAccount struct {
		CurrentWealth          int `json:"Current_Wealth,omitempty"`
		InsuranceClaims        int `json:"Insurance_Claims,omitempty"`
		OwnedShipCount         int `json:"Owned_Ship_Count,omitempty"`
		PremiumStockBought     int `json:"Premium_Stock_Bought,omitempty"`
		SpentOnAmmoConsumables int `json:"Spent_On_Ammo_Consumables,omitempty"`
		SpentOnFuel            int `json:"Spent_On_Fuel,omitempty"`
		SpentOnInsurance       int `json:"Spent_On_Insurance,omitempty"`
		SpentOnOutfitting      int `json:"Spent_On_Outfitting,omitempty"`
		SpentOnPremiumStock    int `json:"Spent_On_Premium_Stock,omitempty"`
		SpentOnRepairs         int `json:"Spent_On_Repairs,omitempty"`
		SpentOnShips           int `json:"Spent_On_Ships,omitempty"`
		SpentOnSuitConsumables int `json:"Spent_On_Suit_Consumables,omitempty"`
		SpentOnSuits           int `json:"Spent_On_Suits,omitempty"`
		SpentOnWeapons         int `json:"Spent_On_Weapons,omitempty"`
		SuitsOwned             int `json:"Suits_Owned,omitempty"`
		WeaponsOwned           int `json:"Weapons_Owned,omitempty"`
	} `json:"Bank_Account,omitempty"`
	Combat struct {
		AssassinationProfits   int `json:"Assassination_Profits,omitempty"`
		Assassinations         int `json:"Assassinations,omitempty"`
		BountiesClaimed        int `json:"Bounties_Claimed,omitempty"`
		BountyHuntingProfit    int `json:"Bounty_Hunting_Profit,omitempty"`
		CombatBondProfits      int `json:"Combat_Bond_Profits,omitempty"`
		CombatBonds            int `json:"Combat_Bonds,omitempty"`
		ConflictZoneHigh       int `json:"ConflictZone_High,omitempty"`
		ConflictZoneHighWins   int `json:"ConflictZone_High_Wins,omitempty"`
		ConflictZoneLow        int `json:"ConflictZone_Low,omitempty"`
		ConflictZoneLowWins    int `json:"ConflictZone_Low_Wins,omitempty"`
		ConflictZoneMedium     int `json:"ConflictZone_Medium,omitempty"`
		ConflictZoneMediumWins int `json:"ConflictZone_Medium_Wins,omitempty"`
		ConflictZoneTotal      int `json:"ConflictZone_Total,omitempty"`
		ConflictZoneTotalWins  int `json:"ConflictZone_Total_Wins,omitempty"`
		DropshipsBooked        int `json:"Dropships_Booked,omitempty"`
		DropshipsCancelled     int `json:"Dropships_Cancelled,omitempty"`
		DropshipsTaken         int `json:"Dropships_Taken,omitempty"`
		HighestSingleReward    int `json:"Highest_Single_Reward,omitempty"`
		OnFootScavsKilled      int `json:"OnFoot_Scavs_Killed,omitempty"`
		OnFootSkimmersKilled   int `json:"OnFoot_Skimmers_Killed,omitempty"`
		SettlementConquered    int `json:"Settlement_Conquered,omitempty"`
		SettlementDefended     int `json:"Settlement_Defended,omitempty"`
		SkimmersKilled         int `json:"Skimmers_Killed,omitempty"`
	} `json:"Combat,omitempty"`
	Crafting *struct {
		CountOfUsedEngineers  int `json:"Count_Of_Used_Engineers,omitempty"`
		RecipesGenerated      int `json:"Recipes_Generated,omitempty"`
		RecipesGeneratedRank1 int `json:"Recipes_Generated_Rank_1,omitempty"`
		RecipesGeneratedRank2 int `json:"Recipes_Generated_Rank_2,omitempty"`
		RecipesGeneratedRank3 int `json:"Recipes_Generated_Rank_3,omitempty"`
		RecipesGeneratedRank4 int `json:"Recipes_Generated_Rank_4,omitempty"`
		RecipesGeneratedRank5 int `json:"Recipes_Generated_Rank_5,omitempty"`
		SuitModsApplied       int `json:"Suit_Mods_Applied,omitempty"`
		SuitModsAppliedFull   int `json:"Suit_Mods_Applied_Full,omitempty"`
		SuitsUpgraded         int `json:"Suits_Upgraded,omitempty"`
		SuitsUpgradedFull     int `json:"Suits_Upgraded_Full,omitempty"`
		WeaponModsApplied     int `json:"Weapon_Mods_Applied,omitempty"`
		WeaponModsAppliedFull int `json:"Weapon_Mods_Applied_Full,omitempty"`
		WeaponsUpgraded       int `json:"Weapons_Upgraded,omitempty"`
		WeaponsUpgradedFull   int `json:"Weapons_Upgraded_Full,omitempty"`
	} `json:"Crafting,omitempty"`
	Crew *struct {
		NpcCrewDied       int `json:"NpcCrew_Died,omitempty"`
		NpcCrewFired      int `json:"NpcCrew_Fired,omitempty"`
		NpcCrewHired      int `json:"NpcCrew_Hired,omitempty"`
		NpcCrewTotalWages int `json:"NpcCrew_TotalWages,omitempty"`
	} `json:"Crew,omitempty"`
	Crime struct {
		BountiesReceived         int `json:"Bounties_Received,omitempty"`
		CitizensMurdered         int `json:"Citizens_Murdered,omitempty"`
		DataStolen               int `json:"Data_Stolen,omitempty"`
		Fines                    int `json:"Fines,omitempty"`
		GoodsStolen              int `json:"Goods_Stolen,omitempty"`
		GuardsMurdered           int `json:"Guards_Murdered,omitempty"`
		HighestBounty            int `json:"Highest_Bounty,omitempty"`
		MalwareUploaded          int `json:"Malware_Uploaded,omitempty"`
		Notoriety                int `json:"Notoriety,omitempty"`
		OmnipolMurdered          int `json:"Omnipol_Murdered,omitempty"`
		ProductionSabotage       int `json:"Production_Sabotage,omitempty"`
		ProductionTheft          int `json:"Production_Theft,omitempty"`
		ProfilesCloned           int `json:"Profiles_Cloned,omitempty"`
		SampleStolen             int `json:"Sample_Stolen,omitempty"`
		SettlementsStateShutdown int `json:"Settlements_State_Shutdown,omitempty"`
		TotalBounties            int `json:"Total_Bounties,omitempty"`
		TotalFines               int `json:"Total_Fines,omitempty"`
		TotalMurders             int `json:"Total_Murders,omitempty"`
		TotalStolen              int `json:"Total_Stolen,omitempty"`
		TurretsDestroyed         int `json:"Turrets_Destroyed,omitempty"`
		TurretsOverloaded        int `json:"Turrets_Overloaded,omitempty"`
		TurretsTotal             int `json:"Turrets_Total,omitempty"`
		ValueStolenStateChange   int `json:"Value_Stolen_StateChange,omitempty"`
	} `json:"Crime,omitempty"`
	Exobiology *struct {
		FirstLogged               int `json:"First_Logged,omitempty"`
		FirstLoggedProfits        int `json:"First_Logged_Profits,omitempty"`
		OrganicData               int `json:"Organic_Data,omitempty"`
		OrganicDataProfits        int `json:"Organic_Data_Profits,omitempty"`
		OrganicGenus              int `json:"Organic_Genus,omitempty"`
		OrganicGenusEncountered   int `json:"Organic_Genus_Encountered,omitempty"`
		OrganicPlanets            int `json:"Organic_Planets,omitempty"`
		OrganicSpecies            int `json:"Organic_Species,omitempty"`
		OrganicSpeciesEncountered int `json:"Organic_Species_Encountered,omitempty"`
		OrganicSystems            int `json:"Organic_Systems,omitempty"`
		OrganicVariantEncountered int `json:"Organic_Variant_Encountered,omitempty"`
	} `json:"Exobiology,omitempty"`
	Exploration struct {
		EfficientScans            int         `json:"Efficient_Scans,omitempty"`
		ExplorationProfits        int         `json:"Exploration_Profits,omitempty"`
		FirstFootfalls            int         `json:"First_Footfalls,omitempty"`
		GreatestDistanceFromStart json.Number `json:"Greatest_Distance_From_Start,omitempty"`
		HighestPayout             int         `json:"Highest_Payout,omitempty"`
		PlanetFootfalls           int         `json:"Planet_Footfalls,omitempty"`
		PlanetsScannedToLevel2    int         `json:"Planets_Scanned_To_Level_2,omitempty"`
		PlanetsScannedToLevel3    int         `json:"Planets_Scanned_To_Level_3,omitempty"`
		SettlementsVisited        int         `json:"Settlements_Visited,omitempty"`
		ShuttleDistanceTravelled  int         `json:"Shuttle_Distance_Travelled,omitempty"`
		ShuttleJourneys           int         `json:"Shuttle_Journeys,omitempty"`
		SpentOnShuttles           int         `json:"Spent_On_Shuttles,omitempty"`
		SystemsVisited            int         `json:"Systems_Visited,omitempty"`
		TimePlayed                int         `json:"Time_Played,omitempty"`
		TotalHyperspaceDistance   int         `json:"Total_Hyperspace_Distance,omitempty"`
		TotalHyperspaceJumps      int         `json:"Total_Hyperspace_Jumps,omitempty"`
	} `json:"Exploration,omitempty"`
	MaterialTraderStats *struct {
		AssetsTradedIn         int `json:"Assets_Traded_In,omitempty"`
		AssetsTradedOut        int `json:"Assets_Traded_Out,omitempty"`
		EncodedMaterialsTraded int `json:"Encoded_Materials_Traded,omitempty"`
		Grade1MaterialsTraded  int `json:"Grade_1_Materials_Traded,omitempty"`
		Grade2MaterialsTraded  int `json:"Grade_2_Materials_Traded,omitempty"`
		Grade3MaterialsTraded  int `json:"Grade_3_Materials_Traded,omitempty"`
		Grade4MaterialsTraded  int `json:"Grade_4_Materials_Traded,omitempty"`
		Grade5MaterialsTraded  int `json:"Grade_5_Materials_Traded,omitempty"`
		MaterialsTraded        int `json:"Materials_Traded,omitempty"`
		RawMaterialsTraded     int `json:"Raw_Materials_Traded,omitempty"`
		TradesCompleted        int `json:"Trades_Completed,omitempty"`
	} `json:"Material_Trader_Stats,omitempty"`
	Mining struct {
		MaterialsCollected int `json:"Materials_Collected,omitempty"`
		MiningProfits      int `json:"Mining_Profits,omitempty"`
		QuantityMined      int `json:"Quantity_Mined,omitempty"`
	} `json:"Mining,omitempty"`
	Multicrew *struct {
		MulticrewCreditsTotal     int `json:"Multicrew_Credits_Total,omitempty"`
		MulticrewFighterTimeTotal int `json:"Multicrew_Fighter_Time_Total,omitempty"`
		MulticrewFinesTotal       int `json:"Multicrew_Fines_Total,omitempty"`
		MulticrewGunnerTimeTotal  int `json:"Multicrew_Gunner_Time_Total,omitempty"`
		MulticrewTimeTotal        int `json:"Multicrew_Time_Total,omitempty"`
	} `json:"Multicrew,omitempty"`
	Passengers struct {
		PassengersMissionsAccepted    int `json:"Passengers_Missions_Accepted,omitempty"`
		PassengersMissionsBulk        int `json:"Passengers_Missions_Bulk,omitempty"`
		PassengersMissionsDelivered   int `json:"Passengers_Missions_Delivered,omitempty"`
		PassengersMissionsDisgruntled int `json:"Passengers_Missions_Disgruntled,omitempty"`
		PassengersMissionsEjected     int `json:"Passengers_Missions_Ejected,omitempty"`
		PassengersMissionsVip         int `json:"Passengers_Missions_VIP,omitempty"`
	} `json:"Passengers,omitempty"`
	SearchAndRescue struct {
		MaglocksOpened            int `json:"Maglocks_Opened,omitempty"`
		PanelsOpened              int `json:"Panels_Opened,omitempty"`
		SalvageIllegalPoi         int `json:"Salvage_Illegal_POI,omitempty"`
		SalvageIllegalSettlements int `json:"Salvage_Illegal_Settlements,omitempty"`
		SalvageLegalPoi           int `json:"Salvage_Legal_POI,omitempty"`
		SalvageLegalSettlements   int `json:"Salvage_Legal_Settlements,omitempty"`
		SearchRescueCount         int `json:"SearchRescue_Count,omitempty"`
		SearchRescueProfit        int `json:"SearchRescue_Profit,omitempty"`
		SearchRescueTraded        int `json:"SearchRescue_Traded,omitempty"`
		SettlementsStateFireOut   int `json:"Settlements_State_FireOut,omitempty"`
		SettlementsStateReboot    int `json:"Settlements_State_Reboot,omitempty"`
	} `json:"Search_And_Rescue,omitempty"`
	Smuggling struct {
		AverageProfit            json.Number `json:"Average_Profit,omitempty"`
		BlackMarketsProfits      int         `json:"Black_Markets_Profits,omitempty"`
		BlackMarketsTradedWith   int         `json:"Black_Markets_Traded_With,omitempty"`
		HighestSingleTransaction int         `json:"Highest_Single_Transaction,omitempty"`
		ResourcesSmuggled        int         `json:"Resources_Smuggled,omitempty"`
	} `json:"Smuggling,omitempty"`
	Trading struct {
		AssetsSold               int         `json:"Assets_Sold,omitempty"`
		AverageProfit            json.Number `json:"Average_Profit,omitempty"`
		DataSold                 int         `json:"Data_Sold,omitempty"`
		GoodsSold                int         `json:"Goods_Sold,omitempty"`
		HighestSingleTransaction int         `json:"Highest_Single_Transaction,omitempty"`
		MarketProfits            int         `json:"Market_Profits,omitempty"`
		MarketsTradedWith        int         `json:"Markets_Traded_With,omitempty"`
		ResourcesTraded          int         `json:"Resources_Traded,omitempty"`
	} `json:"Trading,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evStatistics(ev *StatisticsEvent) {
	return
}

type StoredModulesEvent struct {
	Items []struct {
		BuyPrice              int     `json:"BuyPrice,omitempty"`
		EngineerModifications string  `json:"EngineerModifications,omitempty"`
		Hot                   bool    `json:"Hot,omitempty"`
		InTransit             bool    `json:"InTransit,omitempty"`
		Level                 int     `json:"Level,omitempty"`
		MarketID              int     `json:"MarketID,omitempty"`
		Name                  string  `json:"Name,omitempty"`
		NameLocalised         string  `json:"Name_Localised,omitempty"`
		Quality               float64 `json:"Quality,omitempty"`
		StarSystem            string  `json:"StarSystem,omitempty"`
		StorageSlot           int     `json:"StorageSlot,omitempty"`
		TransferCost          int     `json:"TransferCost,omitempty"`
		TransferTime          int     `json:"TransferTime,omitempty"`
	} `json:"Items,omitempty"`
	MarketID    int       `json:"MarketID,omitempty"`
	StarSystem  string    `json:"StarSystem,omitempty"`
	StationName string    `json:"StationName,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evStoredModules(ev *StoredModulesEvent) {
	return
}

type StoredShipsEvent struct {
	MarketID  int `json:"MarketID,omitempty"`
	ShipsHere []struct {
		Hot               bool   `json:"Hot,omitempty"`
		Name              string `json:"Name,omitempty"`
		ShipID            int    `json:"ShipID,omitempty"`
		ShipType          string `json:"ShipType,omitempty"`
		ShipTypeLocalised string `json:"ShipType_Localised,omitempty"`
		Value             int    `json:"Value,omitempty"`
	} `json:"ShipsHere,omitempty"`
	ShipsRemote []struct {
		Hot               bool   `json:"Hot,omitempty"`
		InTransit         bool   `json:"InTransit,omitempty"`
		Name              string `json:"Name,omitempty"`
		ShipID            int    `json:"ShipID,omitempty"`
		ShipMarketID      int    `json:"ShipMarketID,omitempty"`
		ShipType          string `json:"ShipType,omitempty"`
		ShipTypeLocalised string `json:"ShipType_Localised,omitempty"`
		StarSystem        string `json:"StarSystem,omitempty"`
		TransferPrice     int    `json:"TransferPrice,omitempty"`
		TransferTime      int    `json:"TransferTime,omitempty"`
		Value             int    `json:"Value,omitempty"`
	} `json:"ShipsRemote,omitempty"`
	StarSystem  string    `json:"StarSystem,omitempty"`
	StationName string    `json:"StationName,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evStoredShips(ev *StoredShipsEvent) {
	return
}

type SuitLoadoutEvent struct {
	LoadoutID   int    `json:"LoadoutID,omitempty"`
	LoadoutName string `json:"LoadoutName,omitempty"`
	Modules     []struct {
		Class               int           `json:"Class,omitempty"`
		ModuleName          string        `json:"ModuleName,omitempty"`
		ModuleNameLocalised string        `json:"ModuleName_Localised,omitempty"`
		SlotName            string        `json:"SlotName,omitempty"`
		SuitModuleID        int           `json:"SuitModuleID,omitempty"`
		WeaponMods          []interface{} `json:"WeaponMods,omitempty"`
	} `json:"Modules,omitempty"`
	SuitID            int           `json:"SuitID,omitempty"`
	SuitMods          []interface{} `json:"SuitMods,omitempty"`
	SuitName          string        `json:"SuitName,omitempty"`
	SuitNameLocalised string        `json:"SuitName_Localised,omitempty"`
	Event             string        `json:"event,omitempty"`
	Timestamp         time.Time     `json:"timestamp,omitempty"`
}

func (h *handler) evSuitLoadout(ev *SuitLoadoutEvent) {
	return
}

type SupercruiseEntryEvent struct {
	Multicrew     bool      `json:"Multicrew,omitempty"`
	StarSystem    string    `json:"StarSystem,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	Taxi          bool      `json:"Taxi,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evSupercruiseEntry(ev *SupercruiseEntryEvent) {
	return
}

type SupercruiseExitEvent struct {
	Body          string    `json:"Body,omitempty"`
	BodyID        int       `json:"BodyID,omitempty"`
	BodyType      string    `json:"BodyType,omitempty"`
	Multicrew     bool      `json:"Multicrew,omitempty"`
	StarSystem    string    `json:"StarSystem,omitempty"`
	SystemAddress int       `json:"SystemAddress,omitempty"`
	Taxi          bool      `json:"Taxi,omitempty"`
	Event         string    `json:"event,omitempty"`
	Timestamp     time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evSupercruiseExit(ev *SupercruiseExitEvent) {
	return
}

type SwitchSuitLoadoutEvent struct {
	LoadoutID   int    `json:"LoadoutID,omitempty"`
	LoadoutName string `json:"LoadoutName,omitempty"`
	Modules     []struct {
		Class               int           `json:"Class,omitempty"`
		ModuleName          string        `json:"ModuleName,omitempty"`
		ModuleNameLocalised string        `json:"ModuleName_Localised,omitempty"`
		SlotName            string        `json:"SlotName,omitempty"`
		SuitModuleID        int           `json:"SuitModuleID,omitempty"`
		WeaponMods          []interface{} `json:"WeaponMods,omitempty"`
	} `json:"Modules,omitempty"`
	SuitID            int           `json:"SuitID,omitempty"`
	SuitMods          []interface{} `json:"SuitMods,omitempty"`
	SuitName          string        `json:"SuitName,omitempty"`
	SuitNameLocalised string        `json:"SuitName_Localised,omitempty"`
	Event             string        `json:"event,omitempty"`
	Timestamp         time.Time     `json:"timestamp,omitempty"`
}

func (h *handler) evSwitchSuitLoadout(ev *SwitchSuitLoadoutEvent) {
	return
}

type SynthesisEvent struct {
	Materials []struct {
		Count int    `json:"Count,omitempty"`
		Name  string `json:"Name,omitempty"`
	} `json:"Materials,omitempty"`
	Name      string    `json:"Name,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evSynthesis(ev *SynthesisEvent) {
	return
}

type SystemsShutdownEvent interface{}

func (h *handler) evSystemsShutdown(ev *SystemsShutdownEvent) {
	return
}

type TechnologyBrokerEvent struct {
	BrokerType  string `json:"BrokerType,omitempty"`
	Commodities []struct {
		Count         int    `json:"Count,omitempty"`
		Name          string `json:"Name,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
	} `json:"Commodities,omitempty"`
	ItemsUnlocked []struct {
		Name          string `json:"Name,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
	} `json:"ItemsUnlocked,omitempty"`
	MarketID  int `json:"MarketID,omitempty"`
	Materials []struct {
		Category      string `json:"Category,omitempty"`
		Count         int    `json:"Count,omitempty"`
		Name          string `json:"Name,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
	} `json:"Materials,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evTechnologyBroker(ev *TechnologyBrokerEvent) {
	return
}

type TouchdownEvent struct {
	Body                        string    `json:"Body,omitempty"`
	BodyID                      int       `json:"BodyID,omitempty"`
	Latitude                    float64   `json:"Latitude,omitempty"`
	Longitude                   float64   `json:"Longitude,omitempty"`
	Multicrew                   bool      `json:"Multicrew,omitempty"`
	NearestDestination          string    `json:"NearestDestination,omitempty"`
	NearestDestinationLocalised string    `json:"NearestDestination_Localised,omitempty"`
	OnPlanet                    bool      `json:"OnPlanet,omitempty"`
	OnStation                   bool      `json:"OnStation,omitempty"`
	PlayerControlled            bool      `json:"PlayerControlled,omitempty"`
	StarSystem                  string    `json:"StarSystem,omitempty"`
	SystemAddress               int       `json:"SystemAddress,omitempty"`
	Taxi                        bool      `json:"Taxi,omitempty"`
	Event                       string    `json:"event,omitempty"`
	Timestamp                   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evTouchdown(ev *TouchdownEvent) {
	return
}

type TransferMicroResourcesEvent struct {
	Transfers []struct {
		Category       string `json:"Category,omitempty"`
		Direction      string `json:"Direction,omitempty"`
		LockerNewCount int    `json:"LockerNewCount,omitempty"`
		LockerOldCount int    `json:"LockerOldCount,omitempty"`
		Name           string `json:"Name,omitempty"`
		NameLocalised  string `json:"Name_Localised,omitempty"`
	} `json:"Transfers,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evTransferMicroResources(ev *TransferMicroResourcesEvent) {
	return
}

type UnderAttackEvent struct {
	Target    string    `json:"Target,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evUnderAttack(ev *UnderAttackEvent) {
	return
}

type UndockedEvent struct {
	MarketID    int       `json:"MarketID,omitempty"`
	Multicrew   bool      `json:"Multicrew,omitempty"`
	StationName string    `json:"StationName,omitempty"`
	StationType string    `json:"StationType,omitempty"`
	Taxi        bool      `json:"Taxi,omitempty"`
	Event       string    `json:"event,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evUndocked(ev *UndockedEvent) {
	return
}

type USSDropEvent struct {
	UssThreat        int       `json:"USSThreat,omitempty"`
	UssType          string    `json:"USSType,omitempty"`
	USSTypeLocalised string    `json:"USSType_Localised,omitempty"`
	Event            string    `json:"event,omitempty"`
	Timestamp        time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evUSSDrop(ev *USSDropEvent) {
	return
}

type VehicleSwitchEvent struct {
	To        string    `json:"To,omitempty"`
	Event     string    `json:"event,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}

func (h *handler) evVehicleSwitch(ev *VehicleSwitchEvent) {
	return
}

func (h *handler) handleEvent(evName string, mData map[string]interface{}) {

	switch evName {
	case "ApproachBody":
		ev := new(ApproachBodyEvent)
		mapstructure.Decode(mData, ev)
		h.evApproachBody(ev)
	case "ApproachSettlement":
		ev := new(ApproachSettlementEvent)
		mapstructure.Decode(mData, ev)
		h.evApproachSettlement(ev)
	case "Backpack":
		ev := new(BackpackEvent)
		mapstructure.Decode(mData, ev)
		h.evBackpack(ev)
	case "BackPack":
		ev := new(BackPackEvent)
		mapstructure.Decode(mData, ev)
		h.evBackPack(ev)
	case "Bounty":
		ev := new(BountyEvent)
		mapstructure.Decode(mData, ev)
		h.evBounty(ev)
	case "BuyAmmo":
		ev := new(BuyAmmoEvent)
		mapstructure.Decode(mData, ev)
		h.evBuyAmmo(ev)
	case "BuyDrones":
		ev := new(BuyDronesEvent)
		mapstructure.Decode(mData, ev)
		h.evBuyDrones(ev)
	case "BuyExplorationData":
		ev := new(BuyExplorationDataEvent)
		mapstructure.Decode(mData, ev)
		h.evBuyExplorationData(ev)
	case "BuyMicroResources":
		ev := new(BuyMicroResourcesEvent)
		mapstructure.Decode(mData, ev)
		h.evBuyMicroResources(ev)
	case "BuySuit":
		ev := new(BuySuitEvent)
		mapstructure.Decode(mData, ev)
		h.evBuySuit(ev)
	case "BuyTradeData":
		ev := new(BuyTradeDataEvent)
		mapstructure.Decode(mData, ev)
		h.evBuyTradeData(ev)
	case "BuyWeapon":
		ev := new(BuyWeaponEvent)
		mapstructure.Decode(mData, ev)
		h.evBuyWeapon(ev)
	case "Cargo":
		ev := new(CargoEvent)
		mapstructure.Decode(mData, ev)
		h.evCargo(ev)
	case "CargoDepot":
		ev := new(CargoDepotEvent)
		mapstructure.Decode(mData, ev)
		h.evCargoDepot(ev)
	case "CarrierJump":
		ev := new(CarrierJumpEvent)
		mapstructure.Decode(mData, ev)
		h.evCarrierJump(ev)
	case "CodexEntry":
		ev := new(CodexEntryEvent)
		mapstructure.Decode(mData, ev)
		h.evCodexEntry(ev)
	case "CollectCargo":
		ev := new(CollectCargoEvent)
		mapstructure.Decode(mData, ev)
		h.evCollectCargo(ev)
	case "Commander":
		ev := new(CommanderEvent)
		mapstructure.Decode(mData, ev)
		h.evCommander(ev)
	case "CommitCrime":
		ev := new(CommitCrimeEvent)
		mapstructure.Decode(mData, ev)
		h.evCommitCrime(ev)
	case "CommunityGoal":
		ev := new(CommunityGoalEvent)
		mapstructure.Decode(mData, ev)
		h.evCommunityGoal(ev)
	case "CommunityGoalDiscard":
		ev := new(CommunityGoalDiscardEvent)
		mapstructure.Decode(mData, ev)
		h.evCommunityGoalDiscard(ev)
	case "CommunityGoalJoin":
		ev := new(CommunityGoalJoinEvent)
		mapstructure.Decode(mData, ev)
		h.evCommunityGoalJoin(ev)
	case "CommunityGoalReward":
		ev := new(CommunityGoalRewardEvent)
		mapstructure.Decode(mData, ev)
		h.evCommunityGoalReward(ev)
	case "CreateSuitLoadout":
		ev := new(CreateSuitLoadoutEvent)
		mapstructure.Decode(mData, ev)
		h.evCreateSuitLoadout(ev)
	case "CrewAssign":
		ev := new(CrewAssignEvent)
		mapstructure.Decode(mData, ev)
		h.evCrewAssign(ev)
	case "CrewFire":
		ev := new(CrewFireEvent)
		mapstructure.Decode(mData, ev)
		h.evCrewFire(ev)
	case "CrewHire":
		ev := new(CrewHireEvent)
		mapstructure.Decode(mData, ev)
		h.evCrewHire(ev)
	case "CrimeVictim":
		ev := new(CrimeVictimEvent)
		mapstructure.Decode(mData, ev)
		h.evCrimeVictim(ev)
	case "DatalinkScan":
		ev := new(DatalinkScanEvent)
		mapstructure.Decode(mData, ev)
		h.evDatalinkScan(ev)
	case "DatalinkVoucher":
		ev := new(DatalinkVoucherEvent)
		mapstructure.Decode(mData, ev)
		h.evDatalinkVoucher(ev)
	case "DataScanned":
		ev := new(DataScannedEvent)
		mapstructure.Decode(mData, ev)
		h.evDataScanned(ev)
	case "DeleteSuitLoadout":
		ev := new(DeleteSuitLoadoutEvent)
		mapstructure.Decode(mData, ev)
		h.evDeleteSuitLoadout(ev)
	case "Died":
		ev := new(DiedEvent)
		mapstructure.Decode(mData, ev)
		h.evDied(ev)
	case "Disembark":
		ev := new(DisembarkEvent)
		mapstructure.Decode(mData, ev)
		h.evDisembark(ev)
	case "Docked":
		ev := new(DockedEvent)
		mapstructure.Decode(mData, ev)
		h.evDocked(ev)
	case "DockFighter":
		ev := new(DockFighterEvent)
		mapstructure.Decode(mData, ev)
		h.evDockFighter(ev)
	case "DockingCancelled":
		ev := new(DockingCancelledEvent)
		mapstructure.Decode(mData, ev)
		h.evDockingCancelled(ev)
	case "DockingDenied":
		ev := new(DockingDeniedEvent)
		mapstructure.Decode(mData, ev)
		h.evDockingDenied(ev)
	case "DockingGranted":
		ev := new(DockingGrantedEvent)
		mapstructure.Decode(mData, ev)
		h.evDockingGranted(ev)
	case "DockingRequested":
		ev := new(DockingRequestedEvent)
		mapstructure.Decode(mData, ev)
		h.evDockingRequested(ev)
	case "DockingTimeout":
		ev := new(DockingTimeoutEvent)
		mapstructure.Decode(mData, ev)
		h.evDockingTimeout(ev)
	case "DockSRV":
		ev := new(DockSRVEvent)
		mapstructure.Decode(mData, ev)
		h.evDockSRV(ev)
	case "EjectCargo":
		ev := new(EjectCargoEvent)
		mapstructure.Decode(mData, ev)
		h.evEjectCargo(ev)
	case "Embark":
		ev := new(EmbarkEvent)
		mapstructure.Decode(mData, ev)
		h.evEmbark(ev)
	case "EngineerContribution":
		ev := new(EngineerContributionEvent)
		mapstructure.Decode(mData, ev)
		h.evEngineerContribution(ev)
	case "EngineerCraft":
		ev := new(EngineerCraftEvent)
		mapstructure.Decode(mData, ev)
		h.evEngineerCraft(ev)
	case "EngineerProgress":
		ev := new(EngineerProgressEvent)
		mapstructure.Decode(mData, ev)
		h.evEngineerProgress(ev)
	case "EscapeInterdiction":
		ev := new(EscapeInterdictionEvent)
		mapstructure.Decode(mData, ev)
		h.evEscapeInterdiction(ev)
	case "FetchRemoteModule":
		ev := new(FetchRemoteModuleEvent)
		mapstructure.Decode(mData, ev)
		h.evFetchRemoteModule(ev)
	case "Fileheader":
		ev := new(FileheaderEvent)
		mapstructure.Decode(mData, ev)
		h.evFileheader(ev)
	case "Friends":
		ev := new(FriendsEvent)
		mapstructure.Decode(mData, ev)
		h.evFriends(ev)
	case "FSDJump":
		ev := new(FSDJumpEvent)
		mapstructure.Decode(mData, ev)
		h.evFSDJump(ev)
	case "FSDTarget":
		ev := new(FSDTargetEvent)
		mapstructure.Decode(mData, ev)
		h.evFSDTarget(ev)
	case "FSSAllBodiesFound":
		ev := new(FSSAllBodiesFoundEvent)
		mapstructure.Decode(mData, ev)
		h.evFSSAllBodiesFound(ev)
	case "FSSDiscoveryScan":
		ev := new(FSSDiscoveryScanEvent)
		mapstructure.Decode(mData, ev)
		h.evFSSDiscoveryScan(ev)
	case "FSSSignalDiscovered":
		ev := new(FSSSignalDiscoveredEvent)
		mapstructure.Decode(mData, ev)
		h.evFSSSignalDiscovered(ev)
	case "FuelScoop":
		ev := new(FuelScoopEvent)
		mapstructure.Decode(mData, ev)
		h.evFuelScoop(ev)
	case "HeatDamage":
		ev := new(HeatDamageEvent)
		mapstructure.Decode(mData, ev)
		h.evHeatDamage(ev)
	case "HeatWarning":
		ev := new(HeatWarningEvent)
		mapstructure.Decode(mData, ev)
		h.evHeatWarning(ev)
	case "HullDamage":
		ev := new(HullDamageEvent)
		mapstructure.Decode(mData, ev)
		h.evHullDamage(ev)
	case "Interdicted":
		ev := new(InterdictedEvent)
		mapstructure.Decode(mData, ev)
		h.evInterdicted(ev)
	case "InvitedToSquadron":
		ev := new(InvitedToSquadronEvent)
		mapstructure.Decode(mData, ev)
		h.evInvitedToSquadron(ev)
	case "JetConeBoost":
		ev := new(JetConeBoostEvent)
		mapstructure.Decode(mData, ev)
		h.evJetConeBoost(ev)
	case "LaunchFighter":
		ev := new(LaunchFighterEvent)
		mapstructure.Decode(mData, ev)
		h.evLaunchFighter(ev)
	case "LaunchSRV":
		ev := new(LaunchSRVEvent)
		mapstructure.Decode(mData, ev)
		h.evLaunchSRV(ev)
	case "LeaveBody":
		ev := new(LeaveBodyEvent)
		mapstructure.Decode(mData, ev)
		h.evLeaveBody(ev)
	case "Liftoff":
		ev := new(LiftoffEvent)
		mapstructure.Decode(mData, ev)
		h.evLiftoff(ev)
	case "LoadGame":
		ev := new(LoadGameEvent)
		mapstructure.Decode(mData, ev)
		h.evLoadGame(ev)
	case "Loadout":
		ev := new(LoadoutEvent)
		mapstructure.Decode(mData, ev)
		h.evLoadout(ev)
	case "LoadoutEquipModule":
		ev := new(LoadoutEquipModuleEvent)
		mapstructure.Decode(mData, ev)
		h.evLoadoutEquipModule(ev)
	case "Location":
		ev := new(LocationEvent)
		mapstructure.Decode(mData, ev)
		h.evLocation(ev)
	case "Market":
		ev := new(MarketEvent)
		mapstructure.Decode(mData, ev)
		h.evMarket(ev)
	case "MarketBuy":
		ev := new(MarketBuyEvent)
		mapstructure.Decode(mData, ev)
		h.evMarketBuy(ev)
	case "MarketSell":
		ev := new(MarketSellEvent)
		mapstructure.Decode(mData, ev)
		h.evMarketSell(ev)
	case "MaterialCollected":
		ev := new(MaterialCollectedEvent)
		mapstructure.Decode(mData, ev)
		h.evMaterialCollected(ev)
	case "MaterialDiscovered":
		ev := new(MaterialDiscoveredEvent)
		mapstructure.Decode(mData, ev)
		h.evMaterialDiscovered(ev)
	case "Materials":
		ev := new(MaterialsEvent)
		mapstructure.Decode(mData, ev)
		h.evMaterials(ev)
	case "MaterialTrade":
		ev := new(MaterialTradeEvent)
		mapstructure.Decode(mData, ev)
		h.evMaterialTrade(ev)
	case "MissionAbandoned":
		ev := new(MissionAbandonedEvent)
		mapstructure.Decode(mData, ev)
		h.evMissionAbandoned(ev)
	case "MissionAccepted":
		ev := new(MissionAcceptedEvent)
		mapstructure.Decode(mData, ev)
		h.evMissionAccepted(ev)
	case "MissionCompleted":
		ev := new(MissionCompletedEvent)
		mapstructure.Decode(mData, ev)
		h.evMissionCompleted(ev)
	case "MissionFailed":
		ev := new(MissionFailedEvent)
		mapstructure.Decode(mData, ev)
		h.evMissionFailed(ev)
	case "MissionRedirected":
		ev := new(MissionRedirectedEvent)
		mapstructure.Decode(mData, ev)
		h.evMissionRedirected(ev)
	case "Missions":
		ev := new(MissionsEvent)
		mapstructure.Decode(mData, ev)
		h.evMissions(ev)
	case "ModuleBuy":
		ev := new(ModuleBuyEvent)
		mapstructure.Decode(mData, ev)
		h.evModuleBuy(ev)
	case "ModuleInfo":
		ev := new(ModuleInfoEvent)
		mapstructure.Decode(mData, ev)
		h.evModuleInfo(ev)
	case "ModuleRetrieve":
		ev := new(ModuleRetrieveEvent)
		mapstructure.Decode(mData, ev)
		h.evModuleRetrieve(ev)
	case "ModuleSell":
		ev := new(ModuleSellEvent)
		mapstructure.Decode(mData, ev)
		h.evModuleSell(ev)
	case "ModuleSellRemote":
		ev := new(ModuleSellRemoteEvent)
		mapstructure.Decode(mData, ev)
		h.evModuleSellRemote(ev)
	case "ModuleStore":
		ev := new(ModuleStoreEvent)
		mapstructure.Decode(mData, ev)
		h.evModuleStore(ev)
	case "ModuleSwap":
		ev := new(ModuleSwapEvent)
		mapstructure.Decode(mData, ev)
		h.evModuleSwap(ev)
	case "MultiSellExplorationData":
		ev := new(MultiSellExplorationDataEvent)
		mapstructure.Decode(mData, ev)
		h.evMultiSellExplorationData(ev)
	case "Music":
		ev := new(MusicEvent)
		mapstructure.Decode(mData, ev)
		h.evMusic(ev)
	case "NavRoute":
		ev := new(NavRouteEvent)
		mapstructure.Decode(mData, ev)
		h.evNavRoute(ev)
	case "NewCommander":
		ev := new(NewCommanderEvent)
		mapstructure.Decode(mData, ev)
		h.evNewCommander(ev)
	case "NpcCrewPaidWage":
		ev := new(NpcCrewPaidWageEvent)
		mapstructure.Decode(mData, ev)
		h.evNpcCrewPaidWage(ev)
	case "Outfitting":
		ev := new(OutfittingEvent)
		mapstructure.Decode(mData, ev)
		h.evOutfitting(ev)
	case "Passengers":
		ev := new(PassengersEvent)
		mapstructure.Decode(mData, ev)
		h.evPassengers(ev)
	case "PayBounties":
		ev := new(PayBountiesEvent)
		mapstructure.Decode(mData, ev)
		h.evPayBounties(ev)
	case "PayFines":
		ev := new(PayFinesEvent)
		mapstructure.Decode(mData, ev)
		h.evPayFines(ev)
	case "Powerplay":
		ev := new(PowerplayEvent)
		mapstructure.Decode(mData, ev)
		h.evPowerplay(ev)
	case "PowerplayJoin":
		ev := new(PowerplayJoinEvent)
		mapstructure.Decode(mData, ev)
		h.evPowerplayJoin(ev)
	case "PowerplayLeave":
		ev := new(PowerplayLeaveEvent)
		mapstructure.Decode(mData, ev)
		h.evPowerplayLeave(ev)
	case "PowerplaySalary":
		ev := new(PowerplaySalaryEvent)
		mapstructure.Decode(mData, ev)
		h.evPowerplaySalary(ev)
	case "Progress":
		ev := new(ProgressEvent)
		mapstructure.Decode(mData, ev)
		h.evProgress(ev)
	case "Promotion":
		ev := new(PromotionEvent)
		mapstructure.Decode(mData, ev)
		h.evPromotion(ev)
	case "Rank":
		ev := new(RankEvent)
		mapstructure.Decode(mData, ev)
		h.evRank(ev)
	case "RebootRepair":
		ev := new(RebootRepairEvent)
		mapstructure.Decode(mData, ev)
		h.evRebootRepair(ev)
	case "ReceiveText":
		ev := new(ReceiveTextEvent)
		mapstructure.Decode(mData, ev)
		h.evReceiveText(ev)
	case "RedeemVoucher":
		ev := new(RedeemVoucherEvent)
		mapstructure.Decode(mData, ev)
		h.evRedeemVoucher(ev)
	case "RefuelAll":
		ev := new(RefuelAllEvent)
		mapstructure.Decode(mData, ev)
		h.evRefuelAll(ev)
	case "RefuelPartial":
		ev := new(RefuelPartialEvent)
		mapstructure.Decode(mData, ev)
		h.evRefuelPartial(ev)
	case "RenameSuitLoadout":
		ev := new(RenameSuitLoadoutEvent)
		mapstructure.Decode(mData, ev)
		h.evRenameSuitLoadout(ev)
	case "Repair":
		ev := new(RepairEvent)
		mapstructure.Decode(mData, ev)
		h.evRepair(ev)
	case "RepairAll":
		ev := new(RepairAllEvent)
		mapstructure.Decode(mData, ev)
		h.evRepairAll(ev)
	case "Reputation":
		ev := new(ReputationEvent)
		mapstructure.Decode(mData, ev)
		h.evReputation(ev)
	case "ReservoirReplenished":
		ev := new(ReservoirReplenishedEvent)
		mapstructure.Decode(mData, ev)
		h.evReservoirReplenished(ev)
	case "Resurrect":
		ev := new(ResurrectEvent)
		mapstructure.Decode(mData, ev)
		h.evResurrect(ev)
	case "SAAScanComplete":
		ev := new(SAAScanCompleteEvent)
		mapstructure.Decode(mData, ev)
		h.evSAAScanComplete(ev)
	case "SAASignalsFound":
		ev := new(SAASignalsFoundEvent)
		mapstructure.Decode(mData, ev)
		h.evSAASignalsFound(ev)
	case "Scan":
		ev := new(ScanEvent)
		mapstructure.Decode(mData, ev)
		h.evScan(ev)
	case "Scanned":
		ev := new(ScannedEvent)
		mapstructure.Decode(mData, ev)
		h.evScanned(ev)
	case "ScanOrganic":
		ev := new(ScanOrganicEvent)
		mapstructure.Decode(mData, ev)
		h.evScanOrganic(ev)
	case "SellDrones":
		ev := new(SellDronesEvent)
		mapstructure.Decode(mData, ev)
		h.evSellDrones(ev)
	case "SellExplorationData":
		ev := new(SellExplorationDataEvent)
		mapstructure.Decode(mData, ev)
		h.evSellExplorationData(ev)
	case "SellOrganicData":
		ev := new(SellOrganicDataEvent)
		mapstructure.Decode(mData, ev)
		h.evSellOrganicData(ev)
	case "SellSuit":
		ev := new(SellSuitEvent)
		mapstructure.Decode(mData, ev)
		h.evSellSuit(ev)
	case "SendText":
		ev := new(SendTextEvent)
		mapstructure.Decode(mData, ev)
		h.evSendText(ev)
	case "SetUserShipName":
		ev := new(SetUserShipNameEvent)
		mapstructure.Decode(mData, ev)
		h.evSetUserShipName(ev)
	case "ShieldState":
		ev := new(ShieldStateEvent)
		mapstructure.Decode(mData, ev)
		h.evShieldState(ev)
	case "ShipLocker":
		ev := new(ShipLockerEvent)
		mapstructure.Decode(mData, ev)
		h.evShipLocker(ev)
	case "ShipLockerMaterials":
		ev := new(ShipLockerMaterialsEvent)
		mapstructure.Decode(mData, ev)
		h.evShipLockerMaterials(ev)
	case "ShipTargeted":
		ev := new(ShipTargetedEvent)
		mapstructure.Decode(mData, ev)
		h.evShipTargeted(ev)
	case "Shipyard":
		ev := new(ShipyardEvent)
		mapstructure.Decode(mData, ev)
		h.evShipyard(ev)
	case "ShipyardBuy":
		ev := new(ShipyardBuyEvent)
		mapstructure.Decode(mData, ev)
		h.evShipyardBuy(ev)
	case "ShipyardNew":
		ev := new(ShipyardNewEvent)
		mapstructure.Decode(mData, ev)
		h.evShipyardNew(ev)
	case "ShipyardSell":
		ev := new(ShipyardSellEvent)
		mapstructure.Decode(mData, ev)
		h.evShipyardSell(ev)
	case "ShipyardSwap":
		ev := new(ShipyardSwapEvent)
		mapstructure.Decode(mData, ev)
		h.evShipyardSwap(ev)
	case "ShipyardTransfer":
		ev := new(ShipyardTransferEvent)
		mapstructure.Decode(mData, ev)
		h.evShipyardTransfer(ev)
	case "Shutdown":
		ev := new(ShutdownEvent)
		mapstructure.Decode(mData, ev)
		h.evShutdown(ev)
	case "SquadronStartup":
		ev := new(SquadronStartupEvent)
		mapstructure.Decode(mData, ev)
		h.evSquadronStartup(ev)
	case "StartJump":
		ev := new(StartJumpEvent)
		mapstructure.Decode(mData, ev)
		h.evStartJump(ev)
	case "Statistics":
		ev := new(StatisticsEvent)
		mapstructure.Decode(mData, ev)
		h.evStatistics(ev)
	case "StoredModules":
		ev := new(StoredModulesEvent)
		mapstructure.Decode(mData, ev)
		h.evStoredModules(ev)
	case "StoredShips":
		ev := new(StoredShipsEvent)
		mapstructure.Decode(mData, ev)
		h.evStoredShips(ev)
	case "SuitLoadout":
		ev := new(SuitLoadoutEvent)
		mapstructure.Decode(mData, ev)
		h.evSuitLoadout(ev)
	case "SupercruiseEntry":
		ev := new(SupercruiseEntryEvent)
		mapstructure.Decode(mData, ev)
		h.evSupercruiseEntry(ev)
	case "SupercruiseExit":
		ev := new(SupercruiseExitEvent)
		mapstructure.Decode(mData, ev)
		h.evSupercruiseExit(ev)
	case "SwitchSuitLoadout":
		ev := new(SwitchSuitLoadoutEvent)
		mapstructure.Decode(mData, ev)
		h.evSwitchSuitLoadout(ev)
	case "Synthesis":
		ev := new(SynthesisEvent)
		mapstructure.Decode(mData, ev)
		h.evSynthesis(ev)
	case "SystemsShutdown":
		ev := new(SystemsShutdownEvent)
		mapstructure.Decode(mData, ev)
		h.evSystemsShutdown(ev)
	case "TechnologyBroker":
		ev := new(TechnologyBrokerEvent)
		mapstructure.Decode(mData, ev)
		h.evTechnologyBroker(ev)
	case "Touchdown":
		ev := new(TouchdownEvent)
		mapstructure.Decode(mData, ev)
		h.evTouchdown(ev)
	case "TransferMicroResources":
		ev := new(TransferMicroResourcesEvent)
		mapstructure.Decode(mData, ev)
		h.evTransferMicroResources(ev)
	case "UnderAttack":
		ev := new(UnderAttackEvent)
		mapstructure.Decode(mData, ev)
		h.evUnderAttack(ev)
	case "Undocked":
		ev := new(UndockedEvent)
		mapstructure.Decode(mData, ev)
		h.evUndocked(ev)
	case "USSDrop":
		ev := new(USSDropEvent)
		mapstructure.Decode(mData, ev)
		h.evUSSDrop(ev)
	case "VehicleSwitch":
		ev := new(VehicleSwitchEvent)
		mapstructure.Decode(mData, ev)
		h.evVehicleSwitch(ev)
	default:
		return
	}

}
