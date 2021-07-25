package zerorpc

import (
	"context"
	pool "github.com/jolestar/go-commons-pool/v2"
	"github.com/sirupsen/logrus"
	"net"
)

type MyPooledObj struct {
	session *Session
}

type MyPooledObjFactory struct {
	serverAddr string
}

func (f *MyPooledObjFactory) MakeObject(ctx context.Context) (*pool.PooledObject, error) {
	conn, err := net.Dial("tcp", f.serverAddr)
	if err != nil {
		logrus.Errorf("create conn err:%v", err)
		return nil, err
	}
	return pool.NewPooledObject(
			&MyPooledObj{
				session: NewSession(conn),
			}),
		nil
}

func (f *MyPooledObjFactory) DestroyObject(ctx context.Context, object *pool.PooledObject) error {
	// do destroy
	myObj := object.Object.(*MyPooledObj)
	logrus.Infof("sessoin in poll destroyed, obj:%s", myObj.session.conn.LocalAddr())
	return myObj.session.Close()
}

func (f *MyPooledObjFactory) ValidateObject(ctx context.Context, object *pool.PooledObject) bool {
	// do validate
	myObj := object.Object.(*MyPooledObj)
	logrus.Infof("sessoin in poll validate, obj:%s", myObj.session.conn.LocalAddr())
	return !myObj.session.disabled
}

func (f *MyPooledObjFactory) ActivateObject(ctx context.Context, object *pool.PooledObject) error {
	// do activate
	myObj := object.Object.(*MyPooledObj)
	logrus.Debugf("sessoin in poll activate, obj:%s", myObj.session.conn.LocalAddr())
	return nil
}

func (f *MyPooledObjFactory) PassivateObject(ctx context.Context, object *pool.PooledObject) error {
	// do passivate
	myObj := object.Object.(*MyPooledObj)
	logrus.Debugf("sessoin in poll passivate, obj:%s", myObj.session.conn.LocalAddr())
	return nil
}
