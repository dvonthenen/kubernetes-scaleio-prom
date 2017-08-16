package server

import (
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/codedellemc/goscaleio"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	numOfStoragePools = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "NumOfStoragePools",
			Help:      "NumOfStoragePools",
		},
		[]string{"system", "domain", "pool"},
	)
	protectedCapacityInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "ProtectedCapacityInKb",
			Help:      "ProtectedCapacityInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	movingCapacityInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "MovingCapacityInKb",
			Help:      "MovingCapacityInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	snapCapacityInUseOccupiedInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "SnapCapacityInUseOccupiedInKb",
			Help:      "SnapCapacityInUseOccupiedInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	snapCapacityInUseInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "SnapCapacityInUseInKb",
			Help:      "SnapCapacityInUseInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	activeFwdRebuildCapacityInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "ActiveFwdRebuildCapacityInKb",
			Help:      "ActiveFwdRebuildCapacityInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	degradedHealthyVacInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "DegradedHealthyVacInKb",
			Help:      "DegradedHealthyVacInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	activeMovingRebalanceJobs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "ActiveMovingRebalanceJobs",
			Help:      "ActiveMovingRebalanceJobs",
		},
		[]string{"system", "domain", "pool"},
	)
	maxCapacityInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "MaxCapacityInKb",
			Help:      "MaxCapacityInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	pendingBckRebuildCapacityInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "PendingBckRebuildCapacityInKb",
			Help:      "PendingBckRebuildCapacityInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	activeMovingOutFwdRebuildJobs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "ActiveMovingOutFwdRebuildJobs",
			Help:      "ActiveMovingOutFwdRebuildJobs",
		},
		[]string{"system", "domain", "pool"},
	)
	capacityLimitInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "CapacityLimitInKb",
			Help:      "CapacityLimitInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	secondaryVacInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "SecondaryVacInKb",
			Help:      "SecondaryVacInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	pendingFwdRebuildCapacityInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "PendingFwdRebuildCapacityInKb",
			Help:      "PendingFwdRebuildCapacityInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	thinCapacityInUseInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "ThinCapacityInUseInKb",
			Help:      "ThinCapacityInUseInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	atRestCapacityInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "AtRestCapacityInKb",
			Help:      "AtRestCapacityInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	activeMovingInBckRebuildJobs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "ActiveMovingInBckRebuildJobs",
			Help:      "ActiveMovingInBckRebuildJobs",
		},
		[]string{"system", "domain", "pool"},
	)
	degradedHealthyCapacityInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "DegradedHealthyCapacityInKb",
			Help:      "DegradedHealthyCapacityInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	numOfScsiInitiators = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "NumOfScsiInitiators",
			Help:      "NumOfScsiInitiators",
		},
		[]string{"system", "domain", "pool"},
	)
	numOfUnmappedVolumes = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "NumOfUnmappedVolumes",
			Help:      "NumOfUnmappedVolumes",
		},
		[]string{"system", "domain", "pool"},
	)
	failedCapacityInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "FailedCapacityInKb",
			Help:      "FailedCapacityInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	numOfVolumes = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "NumOfVolumes",
			Help:      "NumOfVolumes",
		},
		[]string{"system", "domain", "pool"},
	)
	activeBckRebuildCapacityInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "ActiveBckRebuildCapacityInKb",
			Help:      "ActiveBckRebuildCapacityInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	failedVacInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "FailedVacInKb",
			Help:      "FailedVacInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	pendingMovingCapacityInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "PendingMovingCapacityInKb",
			Help:      "PendingMovingCapacityInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	activeMovingInRebalanceJobs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "ActiveMovingInRebalanceJobs",
			Help:      "ActiveMovingInRebalanceJobs",
		},
		[]string{"system", "domain", "pool"},
	)
	pendingMovingInRebalanceJobs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "PendingMovingInRebalanceJobs",
			Help:      "PendingMovingInRebalanceJobs",
		},
		[]string{"system", "domain", "pool"},
	)
	degradedFailedVacInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "DegradedFailedVacInKb",
			Help:      "DegradedFailedVacInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	numOfSnapshots = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "NumOfSnapshots",
			Help:      "NumOfSnapshots",
		},
		[]string{"system", "domain", "pool"},
	)
	rebalanceCapacityInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "RebalanceCapacityInKb",
			Help:      "RebalanceCapacityInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	numOfSdc = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "NumOfSdc",
			Help:      "NumOfSdc",
		},
		[]string{"system", "domain", "pool"},
	)
	activeMovingInFwdRebuildJobs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "ActiveMovingInFwdRebuildJobs",
			Help:      "ActiveMovingInFwdRebuildJobs",
		},
		[]string{"system", "domain", "pool"},
	)
	numOfVtrees = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "NumOfVtrees",
			Help:      "NumOfVtrees",
		},
		[]string{"system", "domain", "pool"},
	)
	thickCapacityInUseInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "ThickCapacityInUseInKb",
			Help:      "ThickCapacityInUseInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	protectedVacInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "ProtectedVacInKb",
			Help:      "ProtectedVacInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	pendingMovingInBckRebuildJobs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "PendingMovingInBckRebuildJobs",
			Help:      "PendingMovingInBckRebuildJobs",
		},
		[]string{"system", "domain", "pool"},
	)
	capacityAvailableForVolumeAllocationInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "CapacityAvailableForVolumeAllocationInKb",
			Help:      "CapacityAvailableForVolumeAllocationInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	pendingRebalanceCapacityInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "PendingRebalanceCapacityInKb",
			Help:      "PendingRebalanceCapacityInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	pendingMovingRebalanceJobs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "PendingMovingRebalanceJobs",
			Help:      "PendingMovingRebalanceJobs",
		},
		[]string{"system", "domain", "pool"},
	)
	numOfProtectionDomains = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "NumOfProtectionDomains",
			Help:      "NumOfProtectionDomains",
		},
		[]string{"system", "domain", "pool"},
	)
	numOfSds = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "NumOfSds",
			Help:      "NumOfSds",
		},
		[]string{"system", "domain", "pool"},
	)
	capacityInUseInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "CapacityInUseInKb",
			Help:      "CapacityInUseInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	degradedFailedCapacityInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "DegradedFailedCapacityInKb",
			Help:      "DegradedFailedCapacityInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	numOfThinBaseVolumes = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "NumOfThinBaseVolumes",
			Help:      "NumOfThinBaseVolumes",
		},
		[]string{"system", "domain", "pool"},
	)
	pendingMovingOutFwdRebuildJobs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "PendingMovingOutFwdRebuildJobs",
			Help:      "PendingMovingOutFwdRebuildJobs",
		},
		[]string{"system", "domain", "pool"},
	)
	pendingMovingOutBckRebuildJobs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "PendingMovingOutBckRebuildJobs",
			Help:      "PendingMovingOutBckRebuildJobs",
		},
		[]string{"system", "domain", "pool"},
	)
	numOfVolumesInDeletion = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "NumOfVolumesInDeletion",
			Help:      "NumOfVolumesInDeletion",
		},
		[]string{"system", "domain", "pool"},
	)
	numOfDevices = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "NumOfDevices",
			Help:      "NumOfDevices",
		},
		[]string{"system", "domain", "pool"},
	)
	inUseVacInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "InUseVacInKb",
			Help:      "InUseVacInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	unreachableUnusedCapacityInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "UnreachableUnusedCapacityInKb",
			Help:      "UnreachableUnusedCapacityInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	spareCapacityInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "SpareCapacityInKb",
			Help:      "SpareCapacityInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	activeMovingOutBckRebuildJobs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "ActiveMovingOutBckRebuildJobs",
			Help:      "ActiveMovingOutBckRebuildJobs",
		},
		[]string{"system", "domain", "pool"},
	)
	primaryVacInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "PrimaryVacInKb",
			Help:      "PrimaryVacInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	numOfThickBaseVolumes = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "NumOfThickBaseVolumes",
			Help:      "NumOfThickBaseVolumes",
		},
		[]string{"system", "domain", "pool"},
	)
	bckRebuildCapacityInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "BckRebuildCapacityInKb",
			Help:      "BckRebuildCapacityInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	numOfMappedToAllVolumes = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "NumOfMappedToAllVolumes",
			Help:      "NumOfMappedToAllVolumes",
		},
		[]string{"system", "domain", "pool"},
	)
	activeMovingCapacityInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "ActiveMovingCapacityInKb",
			Help:      "ActiveMovingCapacityInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	pendingMovingInFwdRebuildJobs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "PendingMovingInFwdRebuildJobs",
			Help:      "PendingMovingInFwdRebuildJobs",
		},
		[]string{"system", "domain", "pool"},
	)
	activeRebalanceCapacityInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "ActiveRebalanceCapacityInKb",
			Help:      "ActiveRebalanceCapacityInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	rmcacheSizeInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "RmcacheSizeInKb",
			Help:      "RmcacheSizeInKb",
		},
		[]string{"system", "domain", "pool"},
	)
	fwdRebuildCapacityInKb = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "scaleio",
			Name:      "FwdRebuildCapacityInKb",
			Help:      "FwdRebuildCapacityInKb",
		},
		[]string{"system", "domain", "pool"},
	)
)

var registerMetrics sync.Once

func init() {
	registerMetrics.Do(func() {
		prometheus.MustRegister(numOfStoragePools)
		prometheus.MustRegister(protectedCapacityInKb)
		prometheus.MustRegister(movingCapacityInKb)
		prometheus.MustRegister(snapCapacityInUseOccupiedInKb)
		prometheus.MustRegister(snapCapacityInUseInKb)
		prometheus.MustRegister(activeFwdRebuildCapacityInKb)
		prometheus.MustRegister(degradedHealthyVacInKb)
		prometheus.MustRegister(activeMovingRebalanceJobs)
		prometheus.MustRegister(maxCapacityInKb)
		prometheus.MustRegister(pendingBckRebuildCapacityInKb)
		prometheus.MustRegister(activeMovingOutFwdRebuildJobs)
		prometheus.MustRegister(capacityLimitInKb)
		prometheus.MustRegister(secondaryVacInKb)
		prometheus.MustRegister(pendingFwdRebuildCapacityInKb)
		prometheus.MustRegister(thinCapacityInUseInKb)
		prometheus.MustRegister(atRestCapacityInKb)
		prometheus.MustRegister(activeMovingInBckRebuildJobs)
		prometheus.MustRegister(degradedHealthyCapacityInKb)
		prometheus.MustRegister(numOfScsiInitiators)
		prometheus.MustRegister(numOfUnmappedVolumes)
		prometheus.MustRegister(failedCapacityInKb)
		prometheus.MustRegister(numOfVolumes)
		prometheus.MustRegister(activeBckRebuildCapacityInKb)
		prometheus.MustRegister(failedVacInKb)
		prometheus.MustRegister(pendingMovingCapacityInKb)
		prometheus.MustRegister(activeMovingInRebalanceJobs)
		prometheus.MustRegister(pendingMovingInRebalanceJobs)
		prometheus.MustRegister(degradedFailedVacInKb)
		prometheus.MustRegister(numOfSnapshots)
		prometheus.MustRegister(rebalanceCapacityInKb)
		prometheus.MustRegister(numOfSdc)
		prometheus.MustRegister(activeMovingInFwdRebuildJobs)
		prometheus.MustRegister(numOfVtrees)
		prometheus.MustRegister(thickCapacityInUseInKb)
		prometheus.MustRegister(protectedVacInKb)
		prometheus.MustRegister(pendingMovingInBckRebuildJobs)
		prometheus.MustRegister(capacityAvailableForVolumeAllocationInKb)
		prometheus.MustRegister(pendingRebalanceCapacityInKb)
		prometheus.MustRegister(pendingMovingRebalanceJobs)
		prometheus.MustRegister(numOfProtectionDomains)
		prometheus.MustRegister(numOfSds)
		prometheus.MustRegister(capacityInUseInKb)
		prometheus.MustRegister(degradedFailedCapacityInKb)
		prometheus.MustRegister(numOfThinBaseVolumes)
		prometheus.MustRegister(pendingMovingOutFwdRebuildJobs)
		prometheus.MustRegister(pendingMovingOutBckRebuildJobs)
		prometheus.MustRegister(numOfVolumesInDeletion)
		prometheus.MustRegister(numOfDevices)
		prometheus.MustRegister(inUseVacInKb)
		prometheus.MustRegister(unreachableUnusedCapacityInKb)
		prometheus.MustRegister(spareCapacityInKb)
		prometheus.MustRegister(activeMovingOutBckRebuildJobs)
		prometheus.MustRegister(primaryVacInKb)
		prometheus.MustRegister(numOfThickBaseVolumes)
		prometheus.MustRegister(bckRebuildCapacityInKb)
		prometheus.MustRegister(numOfMappedToAllVolumes)
		prometheus.MustRegister(activeMovingCapacityInKb)
		prometheus.MustRegister(pendingMovingInFwdRebuildJobs)
		prometheus.MustRegister(activeRebalanceCapacityInKb)
		prometheus.MustRegister(rmcacheSizeInKb)
		prometheus.MustRegister(fwdRebuildCapacityInKb)
	})
}

func (s *RestServer) getScaleIOStats() error {
	log.Debugln("getScaleIOStats ENTER")

	client, err := s.getScaleioClient()
	if err != nil {
		log.Errorln("getScaleioClient failed:", err)
		log.Debugln("getScaleIOStats LEAVE")

		return err
	}

	system, errSystem := client.FindSystem(s.Config.ScaleIOID, s.Config.ScaleIOName, "")
	if errSystem != nil {
		log.Errorln("FindSystem Error:", errSystem)
		log.Debugln("getScaleIOStats LEAVE")
		return errSystem
	}

	domains, errDomain := system.GetProtectionDomain("")
	if errDomain != nil {
		log.Errorln("GetProtectionDomain Error:", errDomain)
		log.Debugln("getScaleIOStats LEAVE")
		return errDomain
	}

	scaleIOID := s.Config.ScaleIOName
	if len(s.Config.ScaleIOID) > 0 {
		log.Debugln("Use ScaleIO SystemID")
		scaleIOID = s.Config.ScaleIOID
	}
	log.Debugln("ScaleIO Identity:", scaleIOID)

	for _, domain := range domains {
		pd := goscaleio.NewProtectionDomainEx(s.scaleIO, domain)
		storagePools, errPool := pd.GetStoragePool("")
		if errPool != nil {
			log.Errorln("GetStoragePool Error:", errPool)
			continue
		}

		for _, storagePool := range storagePools {
			sp := goscaleio.NewStoragePoolEx(s.scaleIO, storagePool)
			stats, errStats := sp.GetStatistics()
			if errStats != nil {
				log.Errorln("GetStatistics Error:", errPool)
				continue
			}

			numOfStoragePools.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.NumOfStoragePools))
			protectedCapacityInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.ProtectedCapacityInKb))
			movingCapacityInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.MovingCapacityInKb))
			snapCapacityInUseOccupiedInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.SnapCapacityInUseOccupiedInKb))
			snapCapacityInUseInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.SnapCapacityInUseInKb))
			activeFwdRebuildCapacityInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.ActiveFwdRebuildCapacityInKb))
			degradedHealthyVacInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.DegradedHealthyVacInKb))
			activeMovingRebalanceJobs.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.ActiveMovingRebalanceJobs))
			maxCapacityInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.MaxCapacityInKb))
			pendingBckRebuildCapacityInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.PendingBckRebuildCapacityInKb))
			activeMovingOutFwdRebuildJobs.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.ActiveMovingOutFwdRebuildJobs))
			capacityLimitInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.CapacityLimitInKb))
			secondaryVacInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.SecondaryVacInKb))
			pendingFwdRebuildCapacityInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.PendingFwdRebuildCapacityInKb))
			thinCapacityInUseInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.ThinCapacityInUseInKb))
			atRestCapacityInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.AtRestCapacityInKb))
			activeMovingInBckRebuildJobs.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.ActiveMovingInBckRebuildJobs))
			degradedHealthyCapacityInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.DegradedHealthyCapacityInKb))
			numOfScsiInitiators.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.NumOfScsiInitiators))
			numOfUnmappedVolumes.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.NumOfUnmappedVolumes))
			failedCapacityInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.FailedCapacityInKb))
			numOfVolumes.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.NumOfVolumes))
			activeBckRebuildCapacityInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.ActiveBckRebuildCapacityInKb))
			failedVacInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.FailedVacInKb))
			pendingMovingCapacityInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.PendingMovingCapacityInKb))
			activeMovingInRebalanceJobs.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.ActiveMovingInRebalanceJobs))
			pendingMovingInRebalanceJobs.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.PendingMovingInRebalanceJobs))
			degradedFailedVacInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.DegradedFailedVacInKb))
			numOfSnapshots.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.NumOfSnapshots))
			rebalanceCapacityInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.RebalanceCapacityInKb))
			numOfSdc.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.NumOfSdc))
			activeMovingInFwdRebuildJobs.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.ActiveMovingInFwdRebuildJobs))
			numOfVtrees.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.NumOfVtrees))
			thickCapacityInUseInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.ThickCapacityInUseInKb))
			protectedVacInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.ProtectedVacInKb))
			pendingMovingInBckRebuildJobs.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.PendingMovingInBckRebuildJobs))
			capacityAvailableForVolumeAllocationInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.CapacityAvailableForVolumeAllocationInKb))
			pendingRebalanceCapacityInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.PendingRebalanceCapacityInKb))
			pendingMovingRebalanceJobs.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.PendingMovingRebalanceJobs))
			numOfProtectionDomains.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.NumOfProtectionDomains))
			numOfSds.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.NumOfSds))
			capacityInUseInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.CapacityInUseInKb))
			degradedFailedCapacityInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.DegradedFailedCapacityInKb))
			numOfThinBaseVolumes.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.NumOfThinBaseVolumes))
			pendingMovingOutFwdRebuildJobs.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.PendingMovingOutFwdRebuildJobs))
			pendingMovingOutBckRebuildJobs.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.PendingMovingOutBckRebuildJobs))
			numOfVolumesInDeletion.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.NumOfVolumesInDeletion))
			numOfDevices.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.NumOfDevices))
			inUseVacInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.InUseVacInKb))
			unreachableUnusedCapacityInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.UnreachableUnusedCapacityInKb))
			spareCapacityInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.SpareCapacityInKb))
			activeMovingOutBckRebuildJobs.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.ActiveMovingOutBckRebuildJobs))
			primaryVacInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.PrimaryVacInKb))
			numOfThickBaseVolumes.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.NumOfThickBaseVolumes))
			bckRebuildCapacityInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.BckRebuildCapacityInKb))
			numOfMappedToAllVolumes.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.NumOfMappedToAllVolumes))
			activeMovingCapacityInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.ActiveMovingCapacityInKb))
			pendingMovingInFwdRebuildJobs.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.PendingMovingInFwdRebuildJobs))
			activeRebalanceCapacityInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.ActiveRebalanceCapacityInKb))
			rmcacheSizeInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.RmcacheSizeInKb))
			fwdRebuildCapacityInKb.WithLabelValues(scaleIOID, domain.Name, storagePool.Name).Set(float64(stats.FwdRebuildCapacityInKb))
		}
	}

	log.Debugln("getScaleIOStats Succeeded")
	log.Debugln("getScaleIOStats LEAVE")

	return nil
}
