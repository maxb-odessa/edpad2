package file

import (
	"encoding/json"
	"time"
)

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
