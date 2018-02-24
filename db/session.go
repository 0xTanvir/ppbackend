package db

import (
	"crypto/tls"
	"crypto/x509"
	"net"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

// prepare structural logging.
var log = logrus.WithFields(logrus.Fields{"component": "db"})

// DB type encapsulate mongodb drive service.
type DB struct {
	session *mgo.Session
}

// Dial mongo server.
func Dial(uri string) (*DB, error) {
	d := &DB{}
	var err error
	if viper.GetBool("db.tls.enable") {
		var dialInfo *mgo.DialInfo
		dialInfo, err = mgo.ParseURL(uri)
		if err != nil {
			return nil, err
		}
		dialInfo.DialServer = d.getDialCallback()
		d.session, err = mgo.DialWithInfo(dialInfo)
		if err != nil {
			return nil, err
		}
	} else {
		d.session, err = mgo.Dial(uri)
	}

	if err != nil {
		return nil, err
	}
	return d, nil
}

// Clone session from current connection of mongo.
func (d *DB) Clone() *mgo.Session {
	return d.session.Clone()
}

// getDialCallback function for SSL connection.
func (d *DB) getDialCallback() func(addr *mgo.ServerAddr) (net.Conn, error) {
	tlsConfig := &tls.Config{
		RootCAs: x509.NewCertPool(),
	}

	if ok := tlsConfig.RootCAs.AppendCertsFromPEM(
		[]byte(viper.GetString("db.tls.pem"))); !ok {
		log.Warnf("can't append PEM")
	}

	return func(addr *mgo.ServerAddr) (net.Conn, error) {
		log.Debugf("dial ssl: %v", addr.String())
		return tls.Dial("tcp", addr.String(), tlsConfig)
	}
}

// Close current open session of mongodb.
func (d *DB) Close() {
	d.session.Close()
}
