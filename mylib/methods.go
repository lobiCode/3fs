package methods

type Arith int

func (t *Arith) ParseMsisdn(msisdn *string, reply *string) error {

	*reply = "Hello. You are the man"
	return nil
}
