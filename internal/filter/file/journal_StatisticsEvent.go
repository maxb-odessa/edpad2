package file

import (
	"encoding/json"
	"time"
)

type StatisticsEvent struct {
	BankAccount struct {
		CurrentWealth          int `mapstructure:"Current_Wealth,omitempty"`
		InsuranceClaims        int `mapstructure:"Insurance_Claims,omitempty"`
		OwnedShipCount         int `mapstructure:"Owned_Ship_Count,omitempty"`
		PremiumStockBought     int `mapstructure:"Premium_Stock_Bought,omitempty"`
		SpentOnAmmoConsumables int `mapstructure:"Spent_On_Ammo_Consumables,omitempty"`
		SpentOnFuel            int `mapstructure:"Spent_On_Fuel,omitempty"`
		SpentOnInsurance       int `mapstructure:"Spent_On_Insurance,omitempty"`
		SpentOnOutfitting      int `mapstructure:"Spent_On_Outfitting,omitempty"`
		SpentOnPremiumStock    int `mapstructure:"Spent_On_Premium_Stock,omitempty"`
		SpentOnRepairs         int `mapstructure:"Spent_On_Repairs,omitempty"`
		SpentOnShips           int `mapstructure:"Spent_On_Ships,omitempty"`
		SpentOnSuitConsumables int `mapstructure:"Spent_On_Suit_Consumables,omitempty"`
		SpentOnSuits           int `mapstructure:"Spent_On_Suits,omitempty"`
		SpentOnWeapons         int `mapstructure:"Spent_On_Weapons,omitempty"`
		SuitsOwned             int `mapstructure:"Suits_Owned,omitempty"`
		WeaponsOwned           int `mapstructure:"Weapons_Owned,omitempty"`
	} `mapstructure:"Bank_Account,omitempty"`
	Combat struct {
		AssassinationProfits   int `mapstructure:"Assassination_Profits,omitempty"`
		Assassinations         int `mapstructure:"Assassinations,omitempty"`
		BountiesClaimed        int `mapstructure:"Bounties_Claimed,omitempty"`
		BountyHuntingProfit    int `mapstructure:"Bounty_Hunting_Profit,omitempty"`
		CombatBondProfits      int `mapstructure:"Combat_Bond_Profits,omitempty"`
		CombatBonds            int `mapstructure:"Combat_Bonds,omitempty"`
		ConflictZoneHigh       int `mapstructure:"ConflictZone_High,omitempty"`
		ConflictZoneHighWins   int `mapstructure:"ConflictZone_High_Wins,omitempty"`
		ConflictZoneLow        int `mapstructure:"ConflictZone_Low,omitempty"`
		ConflictZoneLowWins    int `mapstructure:"ConflictZone_Low_Wins,omitempty"`
		ConflictZoneMedium     int `mapstructure:"ConflictZone_Medium,omitempty"`
		ConflictZoneMediumWins int `mapstructure:"ConflictZone_Medium_Wins,omitempty"`
		ConflictZoneTotal      int `mapstructure:"ConflictZone_Total,omitempty"`
		ConflictZoneTotalWins  int `mapstructure:"ConflictZone_Total_Wins,omitempty"`
		DropshipsBooked        int `mapstructure:"Dropships_Booked,omitempty"`
		DropshipsCancelled     int `mapstructure:"Dropships_Cancelled,omitempty"`
		DropshipsTaken         int `mapstructure:"Dropships_Taken,omitempty"`
		HighestSingleReward    int `mapstructure:"Highest_Single_Reward,omitempty"`
		OnFootScavsKilled      int `mapstructure:"OnFoot_Scavs_Killed,omitempty"`
		OnFootSkimmersKilled   int `mapstructure:"OnFoot_Skimmers_Killed,omitempty"`
		SettlementConquered    int `mapstructure:"Settlement_Conquered,omitempty"`
		SettlementDefended     int `mapstructure:"Settlement_Defended,omitempty"`
		SkimmersKilled         int `mapstructure:"Skimmers_Killed,omitempty"`
	} `mapstructure:"Combat,omitempty"`
	Crafting *struct {
		CountOfUsedEngineers  int `mapstructure:"Count_Of_Used_Engineers,omitempty"`
		RecipesGenerated      int `mapstructure:"Recipes_Generated,omitempty"`
		RecipesGeneratedRank1 int `mapstructure:"Recipes_Generated_Rank_1,omitempty"`
		RecipesGeneratedRank2 int `mapstructure:"Recipes_Generated_Rank_2,omitempty"`
		RecipesGeneratedRank3 int `mapstructure:"Recipes_Generated_Rank_3,omitempty"`
		RecipesGeneratedRank4 int `mapstructure:"Recipes_Generated_Rank_4,omitempty"`
		RecipesGeneratedRank5 int `mapstructure:"Recipes_Generated_Rank_5,omitempty"`
		SuitModsApplied       int `mapstructure:"Suit_Mods_Applied,omitempty"`
		SuitModsAppliedFull   int `mapstructure:"Suit_Mods_Applied_Full,omitempty"`
		SuitsUpgraded         int `mapstructure:"Suits_Upgraded,omitempty"`
		SuitsUpgradedFull     int `mapstructure:"Suits_Upgraded_Full,omitempty"`
		WeaponModsApplied     int `mapstructure:"Weapon_Mods_Applied,omitempty"`
		WeaponModsAppliedFull int `mapstructure:"Weapon_Mods_Applied_Full,omitempty"`
		WeaponsUpgraded       int `mapstructure:"Weapons_Upgraded,omitempty"`
		WeaponsUpgradedFull   int `mapstructure:"Weapons_Upgraded_Full,omitempty"`
	} `mapstructure:"Crafting,omitempty"`
	Crew *struct {
		NpcCrewDied       int `mapstructure:"NpcCrew_Died,omitempty"`
		NpcCrewFired      int `mapstructure:"NpcCrew_Fired,omitempty"`
		NpcCrewHired      int `mapstructure:"NpcCrew_Hired,omitempty"`
		NpcCrewTotalWages int `mapstructure:"NpcCrew_TotalWages,omitempty"`
	} `mapstructure:"Crew,omitempty"`
	Crime struct {
		BountiesReceived         int `mapstructure:"Bounties_Received,omitempty"`
		CitizensMurdered         int `mapstructure:"Citizens_Murdered,omitempty"`
		DataStolen               int `mapstructure:"Data_Stolen,omitempty"`
		Fines                    int `mapstructure:"Fines,omitempty"`
		GoodsStolen              int `mapstructure:"Goods_Stolen,omitempty"`
		GuardsMurdered           int `mapstructure:"Guards_Murdered,omitempty"`
		HighestBounty            int `mapstructure:"Highest_Bounty,omitempty"`
		MalwareUploaded          int `mapstructure:"Malware_Uploaded,omitempty"`
		Notoriety                int `mapstructure:"Notoriety,omitempty"`
		OmnipolMurdered          int `mapstructure:"Omnipol_Murdered,omitempty"`
		ProductionSabotage       int `mapstructure:"Production_Sabotage,omitempty"`
		ProductionTheft          int `mapstructure:"Production_Theft,omitempty"`
		ProfilesCloned           int `mapstructure:"Profiles_Cloned,omitempty"`
		SampleStolen             int `mapstructure:"Sample_Stolen,omitempty"`
		SettlementsStateShutdown int `mapstructure:"Settlements_State_Shutdown,omitempty"`
		TotalBounties            int `mapstructure:"Total_Bounties,omitempty"`
		TotalFines               int `mapstructure:"Total_Fines,omitempty"`
		TotalMurders             int `mapstructure:"Total_Murders,omitempty"`
		TotalStolen              int `mapstructure:"Total_Stolen,omitempty"`
		TurretsDestroyed         int `mapstructure:"Turrets_Destroyed,omitempty"`
		TurretsOverloaded        int `mapstructure:"Turrets_Overloaded,omitempty"`
		TurretsTotal             int `mapstructure:"Turrets_Total,omitempty"`
		ValueStolenStateChange   int `mapstructure:"Value_Stolen_StateChange,omitempty"`
	} `mapstructure:"Crime,omitempty"`
	Exobiology *struct {
		FirstLogged               int `mapstructure:"First_Logged,omitempty"`
		FirstLoggedProfits        int `mapstructure:"First_Logged_Profits,omitempty"`
		OrganicData               int `mapstructure:"Organic_Data,omitempty"`
		OrganicDataProfits        int `mapstructure:"Organic_Data_Profits,omitempty"`
		OrganicGenus              int `mapstructure:"Organic_Genus,omitempty"`
		OrganicGenusEncountered   int `mapstructure:"Organic_Genus_Encountered,omitempty"`
		OrganicPlanets            int `mapstructure:"Organic_Planets,omitempty"`
		OrganicSpecies            int `mapstructure:"Organic_Species,omitempty"`
		OrganicSpeciesEncountered int `mapstructure:"Organic_Species_Encountered,omitempty"`
		OrganicSystems            int `mapstructure:"Organic_Systems,omitempty"`
		OrganicVariantEncountered int `mapstructure:"Organic_Variant_Encountered,omitempty"`
	} `mapstructure:"Exobiology,omitempty"`
	Exploration struct {
		EfficientScans            int         `mapstructure:"Efficient_Scans,omitempty"`
		ExplorationProfits        int         `mapstructure:"Exploration_Profits,omitempty"`
		FirstFootfalls            int         `mapstructure:"First_Footfalls,omitempty"`
		GreatestDistanceFromStart json.Number `mapstructure:"Greatest_Distance_From_Start,omitempty"`
		HighestPayout             int         `mapstructure:"Highest_Payout,omitempty"`
		PlanetFootfalls           int         `mapstructure:"Planet_Footfalls,omitempty"`
		PlanetsScannedToLevel2    int         `mapstructure:"Planets_Scanned_To_Level_2,omitempty"`
		PlanetsScannedToLevel3    int         `mapstructure:"Planets_Scanned_To_Level_3,omitempty"`
		SettlementsVisited        int         `mapstructure:"Settlements_Visited,omitempty"`
		ShuttleDistanceTravelled  int         `mapstructure:"Shuttle_Distance_Travelled,omitempty"`
		ShuttleJourneys           int         `mapstructure:"Shuttle_Journeys,omitempty"`
		SpentOnShuttles           int         `mapstructure:"Spent_On_Shuttles,omitempty"`
		SystemsVisited            int         `mapstructure:"Systems_Visited,omitempty"`
		TimePlayed                int         `mapstructure:"Time_Played,omitempty"`
		TotalHyperspaceDistance   int         `mapstructure:"Total_Hyperspace_Distance,omitempty"`
		TotalHyperspaceJumps      int         `mapstructure:"Total_Hyperspace_Jumps,omitempty"`
	} `mapstructure:"Exploration,omitempty"`
	MaterialTraderStats *struct {
		AssetsTradedIn         int `mapstructure:"Assets_Traded_In,omitempty"`
		AssetsTradedOut        int `mapstructure:"Assets_Traded_Out,omitempty"`
		EncodedMaterialsTraded int `mapstructure:"Encoded_Materials_Traded,omitempty"`
		Grade1MaterialsTraded  int `mapstructure:"Grade_1_Materials_Traded,omitempty"`
		Grade2MaterialsTraded  int `mapstructure:"Grade_2_Materials_Traded,omitempty"`
		Grade3MaterialsTraded  int `mapstructure:"Grade_3_Materials_Traded,omitempty"`
		Grade4MaterialsTraded  int `mapstructure:"Grade_4_Materials_Traded,omitempty"`
		Grade5MaterialsTraded  int `mapstructure:"Grade_5_Materials_Traded,omitempty"`
		MaterialsTraded        int `mapstructure:"Materials_Traded,omitempty"`
		RawMaterialsTraded     int `mapstructure:"Raw_Materials_Traded,omitempty"`
		TradesCompleted        int `mapstructure:"Trades_Completed,omitempty"`
	} `mapstructure:"Material_Trader_Stats,omitempty"`
	Mining struct {
		MaterialsCollected int `mapstructure:"Materials_Collected,omitempty"`
		MiningProfits      int `mapstructure:"Mining_Profits,omitempty"`
		QuantityMined      int `mapstructure:"Quantity_Mined,omitempty"`
	} `mapstructure:"Mining,omitempty"`
	Multicrew *struct {
		MulticrewCreditsTotal     int `mapstructure:"Multicrew_Credits_Total,omitempty"`
		MulticrewFighterTimeTotal int `mapstructure:"Multicrew_Fighter_Time_Total,omitempty"`
		MulticrewFinesTotal       int `mapstructure:"Multicrew_Fines_Total,omitempty"`
		MulticrewGunnerTimeTotal  int `mapstructure:"Multicrew_Gunner_Time_Total,omitempty"`
		MulticrewTimeTotal        int `mapstructure:"Multicrew_Time_Total,omitempty"`
	} `mapstructure:"Multicrew,omitempty"`
	Passengers struct {
		PassengersMissionsAccepted    int `mapstructure:"Passengers_Missions_Accepted,omitempty"`
		PassengersMissionsBulk        int `mapstructure:"Passengers_Missions_Bulk,omitempty"`
		PassengersMissionsDelivered   int `mapstructure:"Passengers_Missions_Delivered,omitempty"`
		PassengersMissionsDisgruntled int `mapstructure:"Passengers_Missions_Disgruntled,omitempty"`
		PassengersMissionsEjected     int `mapstructure:"Passengers_Missions_Ejected,omitempty"`
		PassengersMissionsVip         int `mapstructure:"Passengers_Missions_VIP,omitempty"`
	} `mapstructure:"Passengers,omitempty"`
	SearchAndRescue struct {
		MaglocksOpened            int `mapstructure:"Maglocks_Opened,omitempty"`
		PanelsOpened              int `mapstructure:"Panels_Opened,omitempty"`
		SalvageIllegalPoi         int `mapstructure:"Salvage_Illegal_POI,omitempty"`
		SalvageIllegalSettlements int `mapstructure:"Salvage_Illegal_Settlements,omitempty"`
		SalvageLegalPoi           int `mapstructure:"Salvage_Legal_POI,omitempty"`
		SalvageLegalSettlements   int `mapstructure:"Salvage_Legal_Settlements,omitempty"`
		SearchRescueCount         int `mapstructure:"SearchRescue_Count,omitempty"`
		SearchRescueProfit        int `mapstructure:"SearchRescue_Profit,omitempty"`
		SearchRescueTraded        int `mapstructure:"SearchRescue_Traded,omitempty"`
		SettlementsStateFireOut   int `mapstructure:"Settlements_State_FireOut,omitempty"`
		SettlementsStateReboot    int `mapstructure:"Settlements_State_Reboot,omitempty"`
	} `mapstructure:"Search_And_Rescue,omitempty"`
	Smuggling struct {
		AverageProfit            json.Number `mapstructure:"Average_Profit,omitempty"`
		BlackMarketsProfits      int         `mapstructure:"Black_Markets_Profits,omitempty"`
		BlackMarketsTradedWith   int         `mapstructure:"Black_Markets_Traded_With,omitempty"`
		HighestSingleTransaction int         `mapstructure:"Highest_Single_Transaction,omitempty"`
		ResourcesSmuggled        int         `mapstructure:"Resources_Smuggled,omitempty"`
	} `mapstructure:"Smuggling,omitempty"`
	Trading struct {
		AssetsSold               int         `mapstructure:"Assets_Sold,omitempty"`
		AverageProfit            json.Number `mapstructure:"Average_Profit,omitempty"`
		DataSold                 int         `mapstructure:"Data_Sold,omitempty"`
		GoodsSold                int         `mapstructure:"Goods_Sold,omitempty"`
		HighestSingleTransaction int         `mapstructure:"Highest_Single_Transaction,omitempty"`
		MarketProfits            int         `mapstructure:"Market_Profits,omitempty"`
		MarketsTradedWith        int         `mapstructure:"Markets_Traded_With,omitempty"`
		ResourcesTraded          int         `mapstructure:"Resources_Traded,omitempty"`
	} `mapstructure:"Trading,omitempty"`
	Event     string    `mapstructure:"event,omitempty"`
	Timestamp time.Time `mapstructure:"timestamp,omitempty"`
}

func (h *handler) evStatistics(ev *StatisticsEvent) {
	return
}
