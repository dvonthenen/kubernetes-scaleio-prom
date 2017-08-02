package server

import log "github.com/Sirupsen/logrus"

func (s *RestServer) initPrometheus() error {
	log.Debugln("initPrometheus ENTER")

	client, err := s.getScaleioClient()
	if err != nil {
		log.Errorln("getScaleioClient failed:", err)
		log.Debugln("initPrometheus LEAVE")

		return err
	}

	system, errSystem := client.FindSystem(s.Config.ScaleIOID, s.Config.ScaleIOName, "")
	if errSystem != nil {
		log.Errorln("FindSystem Error:", errSystem)
		log.Infoln("initPrometheus LEAVE")
		return nil, errSystem
	}

	domains, errDomain := system.GetProtectionDomain("")
	if errDomain != nil {
		log.Errorln("GetProtectionDomain Error:", errDomain)
		log.Infoln("initPrometheus LEAVE")
		return nil, errDomain
	}

	scaleIOID := s.Config.ScaleIOName
	if len(s.Config.ScaleIOID) > 0 {
		log.Debugln("Use ScaleIO SystemID")
		scaleIOID = s.Config.ScaleIOID
	}
	log.Debugln("ScaleIO Identity:", scaleIOID)

	for _, domain := range domains {
		storagePools, errPool := domain.GetStoragePool("")
		if errPool != nil {
			log.Errorln("GetStoragePool Error:", errPool)
			continue
		}

		for _, storagePool := range storagePools {
			stats, errStats := storagePool.GetStatistics()
			if errStats != nil {
				log.Errorln("GetStatistics Error:", errPool)
				continue
			}

			keyRoot := scaleIOID + "_" + domain.Name + "_" + storagePool.Name + "_"
		}
	}

	log.Debugln("initPrometheus Succeeded")
	log.Debugln("initPrometheus LEAVE")

	return nil
}
