package retrievalmarket_test

import (
	"bytes"
	"testing"

	"github.com/filecoin-project/specs-actors/actors/abi"
	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	basicnode "github.com/ipld/go-ipld-prime/node/basic"
	"github.com/stretchr/testify/assert"

	"github.com/LIUYAN-0626/test01-go-fil-markets/retrievalmarket"
	"github.com/LIUYAN-0626/test01-go-fil-markets/shared"
	tut "github.com/LIUYAN-0626/test01-go-fil-markets/shared_testutil"
)

func TestParamsMarshalUnmarshal(t *testing.T) {
	pieceCid := tut.GenerateCids(1)[0]

	allSelector := shared.AllSelector()
	params := retrievalmarket.NewParamsV1(abi.NewTokenAmount(123), 456, 789, allSelector, &pieceCid)

	buf := new(bytes.Buffer)
	err := params.MarshalCBOR(buf)
	assert.NoError(t, err)

	unmarshalled := &retrievalmarket.Params{}
	err = unmarshalled.UnmarshalCBOR(buf)
	assert.NoError(t, err)

	assert.Equal(t, params, *unmarshalled)

	nb := basicnode.Style.Any.NewBuilder()
	err = dagcbor.Decoder(nb, bytes.NewBuffer(unmarshalled.Selector.Raw))
	assert.NoError(t, err)
	sel := nb.Build()
	assert.Equal(t, sel, allSelector)
}
