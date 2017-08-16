package server

import (
	log "github.com/Sirupsen/logrus"
	goscaleio "github.com/codedellemc/goscaleio"
)

func (s *RestServer) getScaleioClient() (*goscaleio.Client, error) {
	log.Debugln("getScaleioClient ENTER")

	if s.scaleIO != nil {
		log.Infoln("Reusing ScaleIO Client")
		log.Debugln("getScaleioClient ENTER")
		return s.scaleIO, nil
	}

	log.Infoln("Endpoint:", s.Config.Endpoint)
	log.Infoln("APIVersion:", s.Config.APIVersion)

	client, err := goscaleio.NewClientWithArgs(s.Config.Endpoint, s.Config.APIVersion, true, false)
	if err != nil {
		log.Errorln("NewClientWithArgs Error:", err)
		log.Debugln("getScaleioClient LEAVE")
		return nil, err
	}

	_, err = client.Authenticate(&goscaleio.ConfigConnect{
		Endpoint: s.Config.Endpoint,
		Username: s.Config.Username,
		Password: s.Config.Password,
	})
	if err != nil {
		log.Errorln("Authenticate Error:", err)
		log.Debugln("getScaleioClient LEAVE")
		return nil, err
	}
	log.Infoln("Successfuly logged in to ScaleIO Gateway at", client.SIOEndpoint.String())

	s.scaleIO = client

	log.Debugln("getScaleioClient Succeeded")
	log.Debugln("getScaleioClient LEAVE")

	return client, nil
}
