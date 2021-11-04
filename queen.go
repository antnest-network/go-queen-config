package config

type Queen struct {
	Pika  Pika
	Mine  Mine
	Chain Chain
}

type Pika struct {
	Addrs    []string
	Password string
}

type Mine struct {
	Copies                         int
	BlockSize                      int
	BlockPrice                     string
	PushRandomBlockIntervalSeconds int
	InspectionEnabled              bool
}

type Chain struct {
	Endpoint                  string
	ChequebookFactoryContract string
	LockerContract            string
	IssuerContract            string
	IssuerCronEnabled         bool
	IssuerCronIntervalSeconds int
}
